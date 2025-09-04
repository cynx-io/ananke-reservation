package waitlistservice

import (
	"context"
	"errors"
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	"github.com/cynx-io/ananke-reservation/internal/model/entity"
	"github.com/cynx-io/ananke-reservation/internal/model/response"
	brevo "github.com/getbrevo/brevo-go/lib"
	"gorm.io/gorm"
)

func (s *Service) RegisterWaitlist(ctx context.Context, req *proto.RegisterWaitlistRequest, resp *proto.WaitlistResponse) error {

	waitlistType, err := s.TblWaitlistType.GetWaitlistTypeById(ctx, req.WaitlistTypeId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(resp)
			return err
		}

		response.ErrorDatabaseWaitlistType(resp)
		return err
	}

	// check captcha
	if req.RecaptchaToken == "" {
		response.ErrorNotAllowed(resp)
		return errors.New("recaptcha token is required")
	}

	err = s.RecaptchaClient.Verify(req.RecaptchaToken)
	if err != nil {
		response.ErrorExternalRecaptcha(resp)
		return err
	}

	existedWaitlist, err := s.TblWaitlist.GetWaitlistByEmailAndType(ctx, req.Email, req.WaitlistTypeId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			response.ErrorDatabaseWaitlist(resp)
			return err
		}
		response.Success(resp)
		resp.Base.Desc = "email already registered to this waitlist"
		resp.Waitlist = existedWaitlist.Response()
		return nil
	}

	newWaitlist := &entity.TblWaitlist{
		Email:          req.Email,
		WaitlistTypeId: waitlistType.Id,
	}

	err = s.TblWaitlist.CreateWaitlist(ctx, newWaitlist)
	if err != nil {
		response.ErrorDatabaseWaitlist(resp)
		return err
	}

	listIds := waitlistType.GetListIds()
	_, httpRes, err := s.BrevoClient.ContactsApi.CreateContact(ctx, brevo.CreateContact{
		Email:         req.Email,
		ListIds:       listIds,
		UpdateEnabled: true,
	})
	if err != nil || httpRes.StatusCode >= 300 {
		response.ErrorExternalBrevo(resp)
		if err != nil {
			return err
		}
		return errors.New("brevo API error: " + httpRes.Status)
	}

	response.Success(resp)
	resp.Base.Desc = "waitlist registered successfully"
	resp.Waitlist = newWaitlist.Response()
	return nil
}
