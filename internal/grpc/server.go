package grpc

import (
	"context"
	"github.com/cynx-io/ananke-reservation/internal/app"
	"github.com/cynx-io/ananke-reservation/internal/dependencies/config"
	"github.com/cynx-io/cynx-core/src/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

type Server struct {
	//proto.UnimplementedExampleServiceServer

	Services *app.Services
}

func (s *Server) Start(ctx context.Context) error {

	var g errgroup.Group
	address := config.Config.App.Address + ":" + strconv.Itoa(config.Config.App.Port)

	g.Go(func() error {
		lis, err := net.Listen("tcp", address)
		if err != nil {
			return err
		}

		server := grpc.NewServer()
		//proto.RegisterExampleServiceServer(server, s)
		reflection.Register(server)

		logger.Info(ctx, "Starting gRPC server on ", address)
		return server.Serve(lis)
	})
	return g.Wait()
}
