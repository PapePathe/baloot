package main

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/game"
	"pathe.co/zinx/pkg/player"
)

func main() {
	takes := make([]constrainedTake, 4)
	for {
		g := newSampleGame()
		playerTakes(g, takes)
	}
}

func playerTakes(g *game.Game, takes []constrainedTake) {
	_ptk := []gametake.GameTake{}
	for _, player := range g.GetPlayers() {
		var takeID int = -1
		player.PreetyHand()
		for {
			if g.GetTake() == gametake.TOUT {
				fmt.Println("Going to start the game")
				break
			}

			showTakeIfPasse(g)
			showAvailableTakes(g)

			fmt.Println("Auto player take")

			t := table.NewWriter()
			t.AppendHeader(table.Row{"TAKE", "All", "All Player", "CardsOfTake", "Player CardsOfTake", "Ratio", "Take Ratio", "CanTake"})
			for key, tk := range player.GetTakes() {
				t.AppendRow(tk.Print(key))
			}
			fmt.Println(t.Render())

			fmt.Println("What is your take")
			fmt.Scanf("%d", &takeID)
			if takeID >= 0 && takeID < 7 {
				fmt.Println("Your Take is: ", gametake.AllTakes[takeID].Name())
				take := gametake.AllTakes[takeID]
				g.AddTake(player.GetID(), take)

				ctk := constrainedTake{take: take, takes: []gametake.GameTake{}, cards: player.Hand.Cards}
				for _, t := range _ptk {
					ctk.takes = append(ctk.takes, t)
				}
				_ptk = append(_ptk, take)
				takes = append(takes, ctk)
				fmt.Println(ctk.String())
				break
			}
		}
	}
}

func showAvailableTakes(g *game.Game) {
	var takesSB strings.Builder
	fmt.Println("Available takes")
	for _, take := range gametake.AllTakes {
		if take.GreaterThan(g.GetTake()) || take == gametake.PASSE {
			takesSB.WriteString(fmt.Sprintf("%s %d,", take.Name(), take.GetValue()))
		}
	}
	fmt.Println(takesSB.String())
}

func showTakeIfPasse(g *game.Game) {
	if g.GetTake() != gametake.PASSE {
		fmt.Println("Current Take: ", g.GetTake().Name())
	}
}

func newSampleGame() *game.Game {
	g := game.NewGame()

	fmt.Println("-------------Starting a new game ------------------")
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
