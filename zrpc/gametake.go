package zrpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
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
	if len(req.Cards) != 5 {
		return &proto.RecommendGameTakeResponse{}, status.Errorf(codes.InvalidArgument, "Cards must be a collection of 5 items")
	}

	hCards := [5]cards.Card{}
	for i, c := range req.Cards {
		hCards[i] = cards.Card{Genre: c.Type, Couleur: c.Color}
	}

	response := proto.RecommendGameTakeResponse{}

	for _, gt := range gametake.AllTakes {
		res := gt.EvaluateHand(hCards)
		flags := []*proto.Flag{}
		for _, f := range res.Flags {
			flags = append(flags, &proto.Flag{Name: f.String()})
		}
		response.AvailableTakes = append(response.AvailableTakes, &proto.RecommendedGameTake{
			Take:  gt.Name(),
			Flags: flags,
		})
	}

	return &response, nil
}
