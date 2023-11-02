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

		err := g.DispatchCards()
		fmt.Println(err)

		for _, p := range g.GetPlayers() {
			fmt.Println(p.OrderedCardsForPlaying(g.GetTake()))
		}

		fmt.Scanf("Hey")
		fmt.Println("\n\n\n -------------")
	}
}

func playerTakes(g *game.Game) []constrainedTake {
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
