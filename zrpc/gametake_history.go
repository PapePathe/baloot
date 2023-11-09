package zrpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"pathe.co/zinx/pkg/cards"
	proto "pathe.co/zinx/proto/gametake_history/v1"
)

type TakeHistory struct {
	Take        string
	Constraints []string
	Cards       map[string][]cards.Card
}

type TakeHistoryStore struct {
	collection []TakeHistory
}

func (s *TakeHistoryStore) Add(t TakeHistory) error {
	s.collection = append(s.collection, t)

	return nil
}

type GameTakeHistoryServer struct {
	proto.UnimplementedGameTakeHistoryServer
	store *TakeHistoryStore
}

func NewGameTakeHistoryServer() *GameTakeHistoryServer {
	return &GameTakeHistoryServer{
		store: &TakeHistoryStore{},
	}
}

var ErrInvalidCardsCount = errors.New("must have 5 cards")

func (gh GameTakeHistoryServer) Add(ctx context.Context, req *proto.GameTakeHistoryRequest) (*proto.GameTakeHistoryResponse, error) {
	fmt.Println(req)

	if len(req.Cards) < 5 {
		return &proto.GameTakeHistoryResponse{}, ErrInvalidCardsCount
	}

	gh.addToHistory(req)

	return &proto.GameTakeHistoryResponse{Response: time.Now().String()}, nil
}

func (gh GameTakeHistoryServer) addToHistory(req *proto.GameTakeHistoryRequest) error {
	h := TakeHistory{
		Take:        req.Take,
		Constraints: []string{},
	}

	for _, c := range req.Constraints {
		h.Constraints = append(h.Constraints, c)
	}

	gh.store.Add(h)

	fmt.Println(len(gh.store.collection))
	fmt.Println(gh.store.collection[0])

	return nil
}
