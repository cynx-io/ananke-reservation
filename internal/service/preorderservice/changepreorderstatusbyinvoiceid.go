package paymentservice

import (
	"context"
	"errors"
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	proto2 "github.com/cynx-io/ananke-reservation/api/proto/gen/hermes"
	"github.com/cynx-io/ananke-reservation/internal/dependencies/config"
	"github.com/cynx-io/ananke-reservation/internal/model/response"
	"github.com/cynx-io/ananke-reservation/internal/model/template"
	"github.com/cynx-io/cynx-core/src/externalapi/email"
	"github.com/cynx-io/cynx-core/src/logger"
	"gorm.io/gorm"
	"time"
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

	if preorder == nil {
		response.ErrorNotFound(resp)
		return nil
	}

	if preorder.Status == int32(proto.TransactionStatus_TRANSACTION_STATUS_PENDING) &&
		req.TransactionStatus == proto.TransactionStatus_TRANSACTION_STATUS_COMPLETED {

		httpBody := template.GenerateEmailHttpPreorderSuccess(template.Links{
			InstagramLink: config.Config.Perintis.Social.Instagram,
			TiktokLink:    config.Config.Perintis.Social.Tiktok,
			FacebookLink:  config.Config.Perintis.Social.Facebook,
			TwitterLink:   config.Config.Perintis.Social.Twitter,
			WebsiteLink:   config.Config.Perintis.Social.Website,
			DiscordLink:   config.Config.Perintis.Social.Discord,
		})

		goCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		goBase := req.Base

		go func() {
			defer cancel()
			userResp, goErr := s.HermesUserClient.GetUserById(goCtx, &proto2.GetUserByIdRequest{
				Base: goBase,
				Id:   preorder.UserId,
			})
			if goErr != nil {
				logger.Error(goCtx, "Failed to get user by id: ", goErr)
				return
			}

			if userResp.Base.Code != response.CodeSuccess.String() {
				logger.Error(goCtx, "Failed to get user by id: Hermes Code: ", userResp.Base.Code, " Desc: ", userResp.Base.Desc)
				return
			}

			if userResp.User.Email == "" {
				logger.Error(goCtx, "User email is empty")
				return
			}

			goErr = email.SendEmail(goCtx, email.SendEmailRequest{
				Subject: config.Config.Perintis.Email.SubjectPreorderSuccess,
				Body:    httpBody,
				From:    config.Config.Perintis.Email.From,
				To:      []string{userResp.User.Email},
				IsHTML:  true,
			})
			if goErr != nil {
				logger.Error(goCtx, "Failed to send email: ", goErr)
			}
		}()

	}

	err = s.TblPreorder.UpdatePreorderStatus(ctx, preorder.Id, int32(req.TransactionStatus))
	if err != nil {
		response.ErrorDatabasePreorder(resp)
		return err
	}

	response.Success(resp)
	return nil
}
