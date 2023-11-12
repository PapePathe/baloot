package zrpc

import (
	"context"
	"errors"
	"time"

	"sync"

	"github.com/rs/zerolog/log"

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
	if len(req.Cards) < 5 {
		return &proto.GameTakeHistoryResponse{}, ErrInvalidCardsCount
	}

	log.Debug().Msg("adding a new take history")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go gh.addToHistory(req, &wg)

	wg.Wait()

	return &proto.GameTakeHistoryResponse{Response: time.Now().String()}, nil
}

func (gh GameTakeHistoryServer) addToHistory(req *proto.GameTakeHistoryRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	h := TakeHistory{Take: req.Take, Constraints: []string{}}

	for _, c := range req.Constraints {
		h.Constraints = append(h.Constraints, c)
	}

	gh.store.Add(h)
	log.Info().Int("Collection length", len(gh.store.collection))
}
