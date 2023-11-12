package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	gt_proto "pathe.co/zinx/proto/gametake_history/v1"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/broker"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/game"
	"pathe.co/zinx/pkg/player"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println("did not connect to gametake history service: ", err)
		}
	for  {
		g := newSampleGame()
		playerTakes(g, conn)
		fmt.Println(g.GetTake().Name())

		for _, p := range g.GetPlayers() {
			fmt.Println(p.OrderedCardsForPlaying(g.GetTake()))
		}

		//		fmt.Scanf("Hey")
		fmt.Println("\n\n\n -------------")
	}
}

func playerTakes(g *game.Game, conn *grpc.ClientConn) []constrainedTake {
	_ptk := []gametake.GameTake{}
	kakfaMessages := []kafka.Message{}
	publisher := broker.NewPublisher([]string{"localhost:9092", "localhost:9093", "localhost:9093"}, true)
	takes := []constrainedTake{}

	for _, playerObj := range g.GetPlayers() {
		if g.GetTake() == gametake.TOUT {
			fmt.Println("Going to start the game")

			break
		}

		oldTake := g.GetTake()
		err := g.AddTake(playerObj.GetID(), playerObj.GetBestTake())
		fmt.Println(err)

		ctk := constrainedTake{
			Take:  *playerObj.Take,
			Takes: []gametake.GameTake{oldTake},
			Cards: playerObj.OrderedCardsForTake(g.GetTake()),
		}
		msg, err := ctk.AsKafkaMessage("Player.Auto.Take")

		if err != nil {
			fmt.Println(msg)
		}

		kakfaMessages = append(kakfaMessages, msg)

		fmt.Println(ctk)
		ctk.Takes = append(ctk.Takes, _ptk...)
		_ptk = append(_ptk, *playerObj.Take)
		takes = append(takes, ctk)

		

		cli := gt_proto.NewGameTakeHistoryClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		crds := []*gt_proto.Card{}
		for _, c := range playerObj.Hand.Cards {
			crds = append(crds, &gt_proto.Card{Type: c.Genre, Color: c.Couleur})
		}

		playerTakes := []string{}
		for _, p := range g.GetPlayers() {
			if p != nil && p.GetID() != playerObj.GetID() && p.Take != nil {
				playerTakes = append(playerTakes, (*p.Take).Name())
			}
		}
		resp, err := cli.Add(ctx, &gt_proto.GameTakeHistoryRequest{
			Constraints: playerTakes,
			Take:        (*playerObj.Take).Name(),
			Cards:       crds,
		})

		if err != nil {
			log.Printf("Error calling rpc: %s", err)
		}

		log.Printf("Greeting: %s", resp)

	}

	err := publisher.Publish(kakfaMessages)

	if err != nil {
		fmt.Println(err)
	}

	return takes
}

func newSampleGame() *game.Game {
	g := game.NewGame()

	a := player.NewPlayer()
	err := g.AddPlayer(a)
	fmt.Println(err)

	b := player.NewPlayer()
	err = g.AddPlayer(b)
	fmt.Println(err)

	c := player.NewPlayer()
	err = g.AddPlayer(c)
	fmt.Println(err)

	d := player.NewPlayer()
	err = g.AddPlayer(d)
	fmt.Println(err)

	return g
}

type constrainedTake struct {
	Cards [5]cards.Card       `json:"cards"`
	Take  gametake.GameTake   `json:"take"`
	Takes []gametake.GameTake `json:"takes"`
}

func (c constrainedTake) AsKafkaMessage(topic string) (kafka.Message, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return kafka.Message{}, fmt.Errorf("error marshalling constrainedTake %w", err)
	}

	msg := kafka.Message{Key: []byte(c.Take.Name()), Topic: topic, Value: b}

	return msg, nil
}

func (c constrainedTake) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Take: %s, ", c.Take.Name()))
	sb.WriteString("Constraints: ")

	for _, tk := range c.Takes {
		if tk != nil {
			sb.WriteString(fmt.Sprintf("%s, ", tk.Name()))
		}
	}

	sb.WriteString("Cards: ")

	for _, c := range c.Cards {
		sb.WriteString(fmt.Sprintf("%s, ", c.String()))
	}

	return sb.String()
}
