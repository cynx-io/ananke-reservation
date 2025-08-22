package paymentservice

import (
	"context"
	"errors"
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	pbplutus "github.com/cynx-io/ananke-reservation/api/proto/gen/plutus"
	"github.com/cynx-io/ananke-reservation/internal/model/entity"
	"github.com/cynx-io/ananke-reservation/internal/model/response"
	"gorm.io/gorm"
)

func (s *Service) InitiatePreorder(ctx context.Context, req *proto.InitiatePreorderRequest, resp *proto.PreorderResponse) error {

	if req.Base.UserId == nil {
		response.ErrorUnauthorized(resp)
		return errors.New("user ID is required")
	}

	userId := req.Base.GetUserId()

	completedPreorder, err := s.TblPreorder.GetCompletedPreorderByUserIdAndType(ctx, userId, req.PreorderTypeId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {

		if err != nil {
			response.ErrorDatabasePreorder(resp)
			return err
		}

		response.Success(resp)
		resp.Base.Desc = "found completed preorder"
		resp.Preorder = completedPreorder.Response()
		return nil
	}

	pendingPreorder, err := s.TblPreorder.GetPendingPreorderByUserIdAndType(ctx, userId, req.PreorderTypeId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {

		if err != nil {
			response.ErrorDatabasePreorder(resp)
			return err
		}

		response.Success(resp)
		resp.Base.Desc = "found pending preorder"
		resp.Preorder = pendingPreorder.Response()
		return nil
	}

	preorderType, err := s.TblPreorderType.GetPreorderTypeById(ctx, req.PreorderTypeId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(resp)
			return err
		}

		response.ErrorDatabasePreorderType(resp)
		return err
	}

	invoiceResp, err := s.PlutusPaymentClient.CreatePaymentInvoice(ctx, &pbplutus.CreatePaymentInvoiceRequest{
		Base:              req.Base,
		Currency:          preorderType.Currency,
		Description:       preorderType.Description,
		SuccessReturnUrl:  req.SuccessReturnUrl,
		FailureReturnUrl:  req.FailureReturnUrl,
		UserId:            userId,
		Amount:            preorderType.Amount,
		DurationInSeconds: preorderType.DurationInSeconds,
		PaymentFeature:    pbplutus.PaymentFeature_PREORDER,
	})
	if err != nil {
		response.ErrorMicroPlutus(resp)
		return err
	}

	if invoiceResp.Base.Code != response.CodeSuccess.String() {
		response.ErrorMicroPlutus(resp)
		return errors.New("failed to create payment invoice: " + invoiceResp.Base.Code + ": " + invoiceResp.Base.Desc)
	}

	preorder := &entity.TblPreorder{
		ExpiresAt:        invoiceResp.Payment.ExpiresAt.AsTime(),
		PaymentInvoiceId: invoiceResp.Payment.Id,
		PaymentLinkUrl:   invoiceResp.Payment.PaymentLinkUrl,
		Currency:         invoiceResp.Payment.Currency,
		Description:      invoiceResp.Payment.Description,
		Provider:         invoiceResp.Payment.Provider,
		Status:           int32(invoiceResp.Payment.Status),
		Amount:           invoiceResp.Payment.Amount,
		PreorderTypeId:   preorderType.Id,
		UserId:           userId,
	}

	err = s.TblPreorder.CreatePreorder(ctx, preorder)
	if err != nil {
		response.ErrorDatabasePreorder(resp)
		return err
	}

	response.Success(resp)
	resp.Preorder = preorder.Response()
	return nil
}
