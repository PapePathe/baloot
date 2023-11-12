package zrpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ZGrpcServer struct {
	Listeners []net.Listener
	Server    *grpc.Server
	Backends  []*grpc.Server
}

func NewZGrpcServer(ports []int) (*ZGrpcServer, error) {
	logger := grpc.UnaryInterceptor(GrpcLogger)
	backends := []*grpc.Server{}
	listeners := []net.Listener{}
	for _, port := range ports {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

		if err != nil {
			return nil, fmt.Errorf("failed to listen: %v", err)
		}
		listeners = append(listeners, lis)
		backends = append(backends, grpc.NewServer(logger))
	}

	return &ZGrpcServer{Listeners: listeners, Backends: backends}, nil
}

func (z ZGrpcServer) Start() error {
	log.Println("Starting rpc servers")

	for i, b := range z.Backends {
		if err := b.Serve(z.Listeners[i]); err != nil {
			return fmt.Errorf("failed to start rpc server : %v", err)
		}
	}

	return nil
}
