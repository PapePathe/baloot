package game

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/player"
)

func TestAddPlayer(t *testing.T) {
	g := NewGame()
	p := player.NewPlayer()

	assert.NoError(t, g.AddPlayer(p))
	assert.Equal(t, g.players[0], p)
	assert.Equal(t, p.GetID(), 0)
	assert.Equal(t, g.NombreJoueurs, 1)
	msg := "distributed cards count should be be equal to five cards"
	assert.Equal(t, g.CartesDistribuees, 5, msg)
}

func TestAddMoreThanFivePlayers(t *testing.T) {
	g := NewGame()

	for i := 0; i < 4; i++ {
		var p *player.Player = player.NewPlayer()
		g.AddPlayer(p)
		assert.Equal(t, p.GetID(), i)
	}
	np := player.NewPlayer()
	err := g.AddPlayer(np)
	assert.Error(t, errors.New("Game is full"), err)
}

func TestAddTake(t *testing.T) {
	g, p := NewGame(), player.NewPlayer()
	g.AddPlayer(p)
	err := g.AddTake(p.GetID(), gametake.TOUT)

	assert.Equal(t, nil, err)
	assert.Equal(t, p.Take, &gametake.TOUT)
	assert.Equal(t, g.take, gametake.TOUT)
}

func TestAddTakeLessThanGameTake(t *testing.T) {
	g2 := NewGame()
	p1, p2 := player.NewPlayer(), player.NewPlayer()
	g2.AddPlayer(p1)
	g2.AddPlayer(p2)
	g2.AddTake(p1.GetID(), gametake.CENT)
	err2 := g2.AddTake(p2.GetID(), gametake.PIQUE)
	assert.Equal(t, errors.New("Oops bad take, choose a greater take or pass."), err2)
}

func TestAddTakeLessThanGameTakeButIsPASS(t *testing.T) {
	g2 := NewGame()
	p1, p2 := player.NewPlayer(), player.NewPlayer()
	g2.AddPlayer(p1)
	g2.AddPlayer(p2)
	g2.AddTake(p1.GetID(), gametake.CENT)
	err2 := g2.AddTake(p2.GetID(), gametake.PASSE)
	assert.Equal(t, nil, err2)
}

func TestAddTakeForPlayerThatHasTaken(t *testing.T) {
	g, p := NewGame(), player.NewPlayer()
	g.AddPlayer(p)
	p.Take = &gametake.CENT
	err := g.AddTake(p.GetID(), gametake.TOUT)

	assert.Equal(t, errors.New("Oops duplicate player take"), err)
}

func TestAddTakeLessThanCurrentGameTake(t *testing.T) {
	g, p := NewGame(), player.NewPlayer()
	g.AddPlayer(p)
	p.Take = &gametake.CENT
	g.AddTake(p.GetID(), gametake.TOUT)
}

func TestAddTakeGreaterThanCurrentGameTake(t *testing.T) {
	g, p := NewGame(), player.NewPlayer()
	g.AddPlayer(p)
	p.Take = &gametake.CENT
	g.AddTake(p.GetID(), gametake.TOUT)
}

func TestAddTakePassDoesNotChangeGameTake(t *testing.T) {
	g, p1, p2 := NewGame(), player.NewPlayer(), player.NewPlayer()
	g.AddPlayer(p1)
	g.AddPlayer(p2)
	g.AddTake(p1.GetID(), gametake.CENT)
	g.AddTake(p2.GetID(), gametake.PASSE)
	assert.Equal(t, g.GetTake(), gametake.CENT)
}
