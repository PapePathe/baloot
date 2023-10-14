package main

import (
	"fmt"
	"strings"
	"time"

	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/game"
	"pathe.co/zinx/pkg/player"
)

func main() {
	for {
		g := newSampleGame()
		takes := playerTakes(g)
		for _, take := range takes {
			fmt.Println(take)
		}
		time.Sleep(10 * time.Second)
	}
}

func playerTakes(g *game.Game) (takes []constrainedTake) {
	_ptk := []gametake.GameTake{}
	for _, player := range g.GetPlayers() {
		player.PreetyHand()
		if g.GetTake() == gametake.TOUT {
			fmt.Println("Going to start the game")
			break
		}

		g.AddTake(player.GetID(), player.GetBestTake())
		ctk := constrainedTake{take: *player.Take, takes: []gametake.GameTake{}, cards: player.Hand.Cards}
		for _, t := range _ptk {
			ctk.takes = append(ctk.takes, t)
		}
		_ptk = append(_ptk, *player.Take)
		takes = append(takes, ctk)
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
	cards [5]cards.Card
	take  gametake.GameTake
	takes []gametake.GameTake
}

func (c constrainedTake) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Take: %s, ", c.take.Name()))

	sb.WriteString("Constraints: ")
	for _, tk := range c.takes {
		if tk != nil {
			sb.WriteString(fmt.Sprintf("%s, ", tk.Name()))
		}
	}
	sb.WriteString("Cards: ")
	for _, c := range c.cards {
		sb.WriteString(fmt.Sprintf("%s, ", c.String()))
	}

	return sb.String()
}