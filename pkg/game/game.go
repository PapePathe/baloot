package game

import (
	"errors"
	"math/rand"
	"time"

	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/player"
)

type Game struct {
	Cartes            [32]cards.Card
	NombrePli         int
	Plis              [8][4]cards.Card
	CartesDistribuees int
	NombreJoueurs     int
	players           [4]*player.Player
	take              gametake.GameTake
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	plis := [8][4]cards.Card{}
	jeu := cards.CardSet{}
	players := [4]*player.Player{}
	take := gametake.TREFLE
	p := Game{jeu.Distribuer(), 0, plis, 0, 0, players, take}

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
		return errors.New("Oops duplicate player take")
	}
	g.players[playerID].Take = &take

	if g.take.GreaterThan(take) && take != gametake.PASSE {
		return errors.New("Oops bad take, choose a greater take or pass.")
	}
	g.take = take

	return nil
}

func (g *Game) distribuer() [5]cards.Card {
	cards := [5]cards.Card{}
	for i := 0; i < 5; i++ {
		cards[i] = g.Cartes[g.CartesDistribuees+i]
	}

	g.CartesDistribuees += 5
	return cards
}
