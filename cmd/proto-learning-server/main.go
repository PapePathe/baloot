package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"pathe.co/zinx/pkg/game"
	proto "pathe.co/zinx/proto/gametake/v1"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	proto.UnimplementedGameTakeLearningServer
}

func (s *server) GetHand(ctx context.Context, in *proto.GametakeRequest) (*proto.GametakeResponse, error) {
	log.Printf("Received: %v", in.GetId())
	game := game.NewGame()
	cards := []*proto.Card{}
	for _, c := range game.Cartes[0:5] {
		card := proto.Card{Type: c.Genre, Color: c.Couleur}
		cards = append(cards, &card)
	}

	return &proto.GametakeResponse{Cards: cards}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGameTakeLearningServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
