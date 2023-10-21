// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/gametake/v1/gametake.proto

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

// GameTakeLearningClient is the client API for GameTakeLearning service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameTakeLearningClient interface {
	GetHand(ctx context.Context, in *GametakeRequest, opts ...grpc.CallOption) (*GametakeResponse, error)
	RecommendGameTake(ctx context.Context, in *RecommendGameTakeRequest, opts ...grpc.CallOption) (*RecommendGameTakeRequest, error)
}

type gameTakeLearningClient struct {
	cc grpc.ClientConnInterface
}

func NewGameTakeLearningClient(cc grpc.ClientConnInterface) GameTakeLearningClient {
	return &gameTakeLearningClient{cc}
}

func (c *gameTakeLearningClient) GetHand(ctx context.Context, in *GametakeRequest, opts ...grpc.CallOption) (*GametakeResponse, error) {
	out := new(GametakeResponse)
	err := c.cc.Invoke(ctx, "/proto.gametake.v1.GameTakeLearning/GetHand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameTakeLearningClient) RecommendGameTake(ctx context.Context, in *RecommendGameTakeRequest, opts ...grpc.CallOption) (*RecommendGameTakeRequest, error) {
	out := new(RecommendGameTakeRequest)
	err := c.cc.Invoke(ctx, "/proto.gametake.v1.GameTakeLearning/RecommendGameTake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameTakeLearningServer is the server API for GameTakeLearning service.
// All implementations must embed UnimplementedGameTakeLearningServer
// for forward compatibility
type GameTakeLearningServer interface {
	GetHand(context.Context, *GametakeRequest) (*GametakeResponse, error)
	RecommendGameTake(context.Context, *RecommendGameTakeRequest) (*RecommendGameTakeRequest, error)
	mustEmbedUnimplementedGameTakeLearningServer()
}

// UnimplementedGameTakeLearningServer must be embedded to have forward compatible implementations.
type UnimplementedGameTakeLearningServer struct {
}

func (UnimplementedGameTakeLearningServer) GetHand(context.Context, *GametakeRequest) (*GametakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHand not implemented")
}
func (UnimplementedGameTakeLearningServer) RecommendGameTake(context.Context, *RecommendGameTakeRequest) (*RecommendGameTakeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendGameTake not implemented")
}
func (UnimplementedGameTakeLearningServer) mustEmbedUnimplementedGameTakeLearningServer() {}

// UnsafeGameTakeLearningServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameTakeLearningServer will
// result in compilation errors.
type UnsafeGameTakeLearningServer interface {
	mustEmbedUnimplementedGameTakeLearningServer()
}

func RegisterGameTakeLearningServer(s grpc.ServiceRegistrar, srv GameTakeLearningServer) {
	s.RegisterService(&GameTakeLearning_ServiceDesc, srv)
}

func _GameTakeLearning_GetHand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GametakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameTakeLearningServer).GetHand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.gametake.v1.GameTakeLearning/GetHand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameTakeLearningServer).GetHand(ctx, req.(*GametakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameTakeLearning_RecommendGameTake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendGameTakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameTakeLearningServer).RecommendGameTake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.gametake.v1.GameTakeLearning/RecommendGameTake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameTakeLearningServer).RecommendGameTake(ctx, req.(*RecommendGameTakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameTakeLearning_ServiceDesc is the grpc.ServiceDesc for GameTakeLearning service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameTakeLearning_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.gametake.v1.GameTakeLearning",
	HandlerType: (*GameTakeLearningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHand",
			Handler:    _GameTakeLearning_GetHand_Handler,
		},
		{
			MethodName: "RecommendGameTake",
			Handler:    _GameTakeLearning_RecommendGameTake_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gametake/v1/gametake.proto",
}
