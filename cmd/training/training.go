package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/segmentio/kafka-go"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/broker"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/game"
	"pathe.co/zinx/pkg/player"
)

func main() {
	for {
		g := newSampleGame()
		playerTakes(g)
		fmt.Println(g.GetTake().Name())
		fmt.Println("\n\n\n -------------")
	}
}

func playerTakes(g *game.Game) (takes []constrainedTake) {
	_ptk := []gametake.GameTake{}
	_kakfa_messages := []kafka.Message{}
	publisher := broker.NewPublisher([]string{"localhost:9092", "localhost:9093", "localhost:9093"}, true)

	for _, playerObj := range g.GetPlayers() {
		if g.GetTake() == gametake.TOUT {
			fmt.Println("Going to start the game")
			break
		}

		oldTake := g.GetTake()
		g.AddTake(playerObj.GetID(), playerObj.GetBestTake())
		ctk := constrainedTake{Take: *playerObj.Take, Takes: []gametake.GameTake{oldTake}, Cards: playerObj.OrderedCardsForTake(g.GetTake())}
		msg, err := ctk.AsKafkaMessage("Player.Auto.Take")
		if err != nil {
			fmt.Println(msg)
		}

		_kakfa_messages = append(_kakfa_messages, msg)

		fmt.Println(ctk)
		for _, t := range _ptk {
			ctk.Takes = append(ctk.Takes, t)
		}
		_ptk = append(_ptk, *playerObj.Take)
		takes = append(takes, ctk)
	}

	err := publisher.Publish(_kakfa_messages)

	if err != nil {
		fmt.Println(err)
	}

	return takes
}

func newSampleGame() *game.Game {
	g := game.NewGame()

	a := player.NewPlayer()
	g.AddPlayer(a)
	b := player.NewPlayer()
	g.AddPlayer(b)
	c := player.NewPlayer()
	g.AddPlayer(c)
	d := player.NewPlayer()
	g.AddPlayer(d)
	return g
}

type constrainedTake struct {
	Cards [5]cards.Card
	Take  gametake.GameTake
	Takes []gametake.GameTake
}

func (c constrainedTake) AsKafkaMessage(topic string) (kafka.Message, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return kafka.Message{}, err
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
