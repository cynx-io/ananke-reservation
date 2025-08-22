package grpc

import (
	"context"
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	"github.com/cynx-io/cynx-core/src/grpc"
)

func (s *Server) ChangePreorderStatusByInvoiceId(ctx context.Context, req *proto.ChangePreorderStatusByInvoiceIdRequest) (*proto.PreorderResponse, error) {
	var resp proto.PreorderResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.PaymentService.ChangePreorderStatusByInvoiceId)
}

func (s *Server) InitiatePreorder(ctx context.Context, req *proto.InitiatePreorderRequest) (*proto.PreorderResponse, error) {
	var resp proto.PreorderResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.PaymentService.InitiatePreorder)
}
