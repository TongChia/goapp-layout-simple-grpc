package server

import (
	v1 "goapp-layout-template/api/gen/helloworld/v1"
	"goapp-layout-template/internal/service"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(greeter *service.GreeterService) *grpc.Server {
	unaryServerInterceptors := []grpc.UnaryServerInterceptor{
		grpc_ctxtags.UnaryServerInterceptor(),
	}

	streamServerInterceptors := []grpc.StreamServerInterceptor{
		grpc_ctxtags.StreamServerInterceptor(),
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryServerInterceptors...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamServerInterceptors...)),
	)

	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
