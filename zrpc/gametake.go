package zrpc

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	proto "pathe.co/zinx/proto/gametake/v1"
)

type RecommendGameTakeServer struct {
	proto.UnimplementedGameTakeLearningServer
}

func NewRecommendGameTakeServer() *RecommendGameTakeServer {
	return &RecommendGameTakeServer{}
}

func (s *RecommendGameTakeServer) RecommendGameTake(
	ctx context.Context,
	req *proto.RecommendGameTakeRequest,
) (*proto.RecommendGameTakeResponse, error) {
	log.Debug().Interface("Cards", req.Cards).Msg("")

	if len(req.Cards) != 5 {
		return &proto.RecommendGameTakeResponse{}, status.Errorf(codes.InvalidArgument, "Cards must be a collection of 5 items")
	}

	return &proto.RecommendGameTakeResponse{}, nil
}
