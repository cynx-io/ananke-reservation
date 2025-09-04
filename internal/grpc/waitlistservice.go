package grpc

import (
	"context"
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	"github.com/cynx-io/cynx-core/src/grpc"
)

func (s *Server) RegisterWaitlist(ctx context.Context, req *proto.RegisterWaitlistRequest) (*proto.WaitlistResponse, error) {
	var resp proto.WaitlistResponse
	return grpc.HandleGrpc(ctx, req, &resp, s.Services.WaitlistService.RegisterWaitlist)
}
