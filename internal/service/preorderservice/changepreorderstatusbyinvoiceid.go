package paymentservice

import (
	"context"
	"errors"
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	"github.com/cynx-io/ananke-reservation/internal/model/response"
	"gorm.io/gorm"
)

func (s *Service) ChangePreorderStatusByInvoiceId(ctx context.Context, req *proto.ChangePreorderStatusByInvoiceIdRequest, resp *proto.PreorderResponse) error {

	preorder, err := s.TblPreorder.GetPreorderByInvoiceId(ctx, req.InvoiceId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(resp)
			return err
		}
		response.ErrorDatabasePreorder(resp)
		return err
	}

	err = s.TblPreorder.UpdatePreorderStatus(ctx, preorder.Id, int32(req.TransactionStatus))
	if err != nil {
		response.ErrorDatabasePreorder(resp)
		return err
	}

	response.Success(resp)
	return nil
}
