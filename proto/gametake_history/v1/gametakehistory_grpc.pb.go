// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: proto/gametake_history/v1/gametakehistory.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GameTakeHistoryClient is the client API for GameTakeHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameTakeHistoryClient interface {
	Add(ctx context.Context, in *GameTakeHistoryRequest, opts ...grpc.CallOption) (*GameTakeHistoryResponse, error)
}

type gameTakeHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewGameTakeHistoryClient(cc grpc.ClientConnInterface) GameTakeHistoryClient {
	return &gameTakeHistoryClient{cc}
}

func (c *gameTakeHistoryClient) Add(ctx context.Context, in *GameTakeHistoryRequest, opts ...grpc.CallOption) (*GameTakeHistoryResponse, error) {
	out := new(GameTakeHistoryResponse)
	err := c.cc.Invoke(ctx, "/proto.gametake_history.v1.GameTakeHistory/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameTakeHistoryServer is the server API for GameTakeHistory service.
// All implementations must embed UnimplementedGameTakeHistoryServer
// for forward compatibility
type GameTakeHistoryServer interface {
	Add(context.Context, *GameTakeHistoryRequest) (*GameTakeHistoryResponse, error)
	mustEmbedUnimplementedGameTakeHistoryServer()
}

// UnimplementedGameTakeHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedGameTakeHistoryServer struct {
}

func (UnimplementedGameTakeHistoryServer) Add(context.Context, *GameTakeHistoryRequest) (*GameTakeHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedGameTakeHistoryServer) mustEmbedUnimplementedGameTakeHistoryServer() {}

// UnsafeGameTakeHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameTakeHistoryServer will
// result in compilation errors.
type UnsafeGameTakeHistoryServer interface {
	mustEmbedUnimplementedGameTakeHistoryServer()
}

func RegisterGameTakeHistoryServer(s grpc.ServiceRegistrar, srv GameTakeHistoryServer) {
	s.RegisterService(&GameTakeHistory_ServiceDesc, srv)
}

func _GameTakeHistory_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameTakeHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameTakeHistoryServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.gametake_history.v1.GameTakeHistory/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameTakeHistoryServer).Add(ctx, req.(*GameTakeHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameTakeHistory_ServiceDesc is the grpc.ServiceDesc for GameTakeHistory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameTakeHistory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.gametake_history.v1.GameTakeHistory",
	HandlerType: (*GameTakeHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _GameTakeHistory_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gametake_history/v1/gametakehistory.proto",
}
