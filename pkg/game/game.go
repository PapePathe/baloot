package game

import (
	"math/rand"
	"time"

	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/player"
)

type Game struct {
	Cartes            [32]cards.Card
	PlayerA           *player.Player
	PlayerB           *player.Player
	PlayerC           *player.Player
	PlayerD           *player.Player
	NombrePli         int
	Plis              [8][4]cards.Card
	CartesDistribuees int
	NombreJoueurs     int
}

func NewGame() *Game {

	rand.Seed(time.Now().UnixNano())
	plis := [8][4]cards.Card{}
	jeu := cards.CardSet{}
	p := Game{jeu.Distribuer(), nil, nil, nil, nil, 0, plis, 0, 0}

	return &p
}

func (p *Game) Distribuer() [5]cards.Card {
	cards := [5]cards.Card{}
	for i := 0; i < 5; i++ {
		cards[i] = p.Cartes[p.CartesDistribuees+i]
	}

	p.CartesDistribuees += 4
	return cards
}
