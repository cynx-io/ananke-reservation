package paymentservice

import (
	"context"
	"errors"
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	"github.com/cynx-io/ananke-reservation/internal/model/response"
	"gorm.io/gorm"
)

func (s *Service) GetLatestCompletedOrPendingPreorder(ctx context.Context, req *proto.GetLatestCompletedOrPendingPreorderRequest, resp *proto.PreorderResponse) error {

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
		return err
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
		return err
	}

	response.ErrorNotFound(resp)
	resp.Base.Desc = "no completed or pending preorder found"
	return errors.New("no completed or pending preorder found")
}
