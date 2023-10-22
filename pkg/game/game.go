package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/player"
)

var ErrCardsAlreadyDispatched = errors.New("cards already dispatched error")

type Game struct {
	Cartes            [32]cards.Card
	NombrePli         int
	Plis              [8][4]cards.Card
	CartesDistribuees int
	NombreJoueurs     int
	TakesFinished     bool
	players           [4]*player.Player
	take              gametake.GameTake
}

func NewGame() *Game {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	plis := [8][4]cards.Card{}
	jeu := cards.CardSet{}
	players := [4]*player.Player{}
	take := gametake.PASSE
	p := Game{jeu.Distribuer(), 0, plis, 0, 0, false, players, take}

	return &p
}

func (g *Game) AddPlayer(p *player.Player) error {
	if g.NombreJoueurs == 4 {
		return errors.New("Game is full")
	}

	p.Hand.Cards = g.distribuer()
	p.SetID(g.NombreJoueurs)
	g.players[g.NombreJoueurs] = p
	g.NombreJoueurs++

	return nil
}

func (g *Game) AddTake(playerID int, take gametake.GameTake) error {
	if g.players[playerID].Take != nil {
		return errors.New("oops duplicate player take")
	}
	g.players[playerID].Take = &take

	if g.take.GreaterThan(take) && take != gametake.PASSE {
		return errors.New("oops bad take, choose a greater take or pass")
	}

	if g.take == gametake.PASSE {
		g.take = take
	} else if take != gametake.PASSE {
		g.take = take
	}

	if g.take == gametake.TOUT || g.takesComplete() {
		g.TakesFinished = true
		g.DispatchCards()

		for _, p := range g.players {
			if p != nil {
				fmt.Println("sending playing hand to player")
				r := ReceivePlayingHandMsg(*p, []gametake.GameTake{})
				m, _ := json.Marshal(r)

				if p.Conn != nil {
					if err := p.Conn.WriteMessage(1, m); err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}

	return nil
}

func (g *Game) takesComplete() bool {
	for _, p := range g.players {
		if p != nil && p.Take == nil {
			return false
		}
	}

	return true
}

func (g *Game) distribuer() [5]cards.Card {
	cards := [5]cards.Card{}
	for i := 0; i < 5; i++ {
		cards[i] = g.Cartes[g.CartesDistribuees+i]
	}

	g.CartesDistribuees += 5
	return cards
}

func (g *Game) GetPlayers() [4]*player.Player {
	return g.players
}

func (g *Game) GetTake() gametake.GameTake {
	return g.take
}

func (g *Game) AvailableTakes() []gametake.GameTake {
	takes := []gametake.GameTake{}
	takes = append(takes, gametake.PASSE)
	for _, t := range gametake.AllTakes {
		if t.GreaterThan(g.take) {
			takes = append(takes, t)
		}
	}

	return takes
}

func (g *Game) DispatchCards() error {
	if g.CartesDistribuees == 32 {
		return ErrCardsAlreadyDispatched
	}

	for _, p := range g.players {
		cards := []cards.Card{}
		if p != nil {

			for _, c := range p.Hand.Cards {
				cards = append(cards, c)
			}
			for i := 0; i < 3; i++ {
				cards = append(cards, g.Cartes[g.CartesDistribuees])
				g.CartesDistribuees++
			}

			p.PlayingHand = player.PlayingHand{Cards: cards}
		}
	}

	return nil
}
