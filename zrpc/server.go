package zrpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ZGrpcServer struct {
	Listener net.Listener
	Server   *grpc.Server
}

func NewZGrpcServer(port int) (*ZGrpcServer, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}
	logger := grpc.UnaryInterceptor(GrpcLogger)
	s := grpc.NewServer(logger)

	return &ZGrpcServer{Listener: lis, Server: s}, nil
}

func (z ZGrpcServer) Start() error {
	log.Println("Starting rpc server")

	if err := z.Server.Serve(z.Listener); err != nil {
		return fmt.Errorf("failed to start rpc server : %v", err)
	}

	return nil
}
