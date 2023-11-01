package game

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
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
		err := g.AddPlayer(p)
		assert.NoError(t, err)
		assert.Equal(t, p.GetID(), i)
	}
	np := player.NewPlayer()
	err := g.AddPlayer(np)
	assert.Error(t, errors.New("Game is full"), err)
}

func TestAddTake(t *testing.T) {
	g, p := NewGame(), player.NewPlayer()
	err := g.AddPlayer(p)
	assert.NoError(t, err)

	err = g.AddTake(p.GetID(), gametake.TOUT)
	assert.NoError(t, err)

	err = g.AddTake(p.GetID(), gametake.TOUT)
	assert.Error(t, err)
	assert.Equal(t, p.Take, &gametake.TOUT)
	assert.Equal(t, g.take, gametake.TOUT)
}

func TestAddTakeLessThanGameTake(t *testing.T) {
	g2 := NewGame()
	p1, p2 := player.NewPlayer(), player.NewPlayer()
	err := g2.AddPlayer(p1)
	assert.NoError(t, err)
	err = g2.AddPlayer(p2)
	assert.NoError(t, err)
	err = g2.AddTake(p1.GetID(), gametake.CENT)
	assert.NoError(t, err)
	err2 := g2.AddTake(p2.GetID(), gametake.PIQUE)
	assert.Equal(t, errors.New("oops bad take, choose a greater take or pass"), err2)
}

func TestAddTakeLessThanGameTakeButIsPASS(t *testing.T) {
	g2 := setupGame(2)
	p1, p2 := g2.players[0], g2.players[1]
	err := g2.AddTake(p1.GetID(), gametake.CENT)
	assert.NoError(t, err)
	err = g2.AddTake(p2.GetID(), gametake.PASSE)
	assert.Equal(t, nil, err)
}

func TestAddTakeForPlayerThatHasTaken(t *testing.T) {
	g, p := NewGame(), player.NewPlayer()
	err := g.AddPlayer(p)
	assert.NoError(t, err)

	p.Take = &gametake.CENT
	err = g.AddTake(p.GetID(), gametake.TOUT)

	assert.Equal(t, errors.New("oops duplicate player take"), err)
}

func TestAddTakeLessThanCurrentGameTake(t *testing.T) {
	g, p := NewGame(), player.NewPlayer()
	err := g.AddPlayer(p)
	assert.NoError(t, err)
	err = g.AddTake(p.GetID(), gametake.TOUT)
	assert.NoError(t, err)
}

func TestAddTakeGreaterThanCurrentGameTake(t *testing.T) {
	g, p := NewGame(), player.NewPlayer()
	err := g.AddPlayer(p)
	assert.NoError(t, err)
	err = g.AddTake(p.GetID(), gametake.TOUT)
	assert.NoError(t, err)
}

func TestAddTakePassDoesNotChangeGameTake(t *testing.T) {
	g := setupGame(2)
	p1, p2 := g.players[0], g.players[1]
	err := g.AddTake(p1.GetID(), gametake.CENT)
	assert.NoError(t, err)
	err = g.AddTake(p2.GetID(), gametake.PASSE)
	assert.NoError(t, err)
	assert.Equal(t, g.GetTake(), gametake.CENT)
}

func TestDispatchCards(t *testing.T) {
	g := setupGame(4)
	err := g.DispatchCards()

	assert.Equal(t, err, nil)
	assert.Equal(t, g.CartesDistribuees, 32)

	for _, p := range g.players {
		assert.Equal(t, len(p.PlayingHand.Cards), 8)
	}
}

func TestDispatchCardsIsIdempotent(t *testing.T) {
	g := setupGame(4)
	err := g.DispatchCards()

	assert.Equal(t, err, nil)
	assert.Error(t, g.DispatchCards(), ErrCardsAlreadyDispatched)
}

func TestNewGame(t *testing.T) {
	g := NewGame()

	assert.Equal(t, g.TakesFinished, false)
}

func TestPlayCard(t *testing.T) {
	t.Run("Player can play one of his cards", func(t *testing.T) {
		g := setupGame(4)
		err := g.DispatchCards()
		assert.Equal(t, err, nil)
		p1, p2, p3, p4 := g.players[0], g.players[1], g.players[3], g.players[3]

		c1 := p1.PlayingHand.Cards[0]
		err = g.PlayCard(p1.GetID(), p1.PlayingHand.Cards[0])
		pli := [4]cards.Card{c1}
		assert.Equal(t, err, nil)
		assert.Equal(t, g.Plis[0], pli)
		assert.Equal(t, g.pliCardsCount, 1)
		assert.Equal(t, g.nombrePli, 0)

		c2 := p2.PlayingHand.Cards[0]
		err2 := g.PlayCard(p2.GetID(), c2)
		pli2 := [4]cards.Card{c1, c2}
		assert.Equal(t, err2, nil)
		assert.Equal(t, g.Plis[0], pli2)
		assert.Equal(t, g.pliCardsCount, 2)
		assert.Equal(t, g.nombrePli, 0)

		c3 := p3.PlayingHand.Cards[0]
		err3 := g.PlayCard(p3.GetID(), c3)
		pli3 := [4]cards.Card{c1, c2, c3}
		assert.Equal(t, err3, nil)
		assert.Equal(t, g.Plis[0], pli3)
		assert.Equal(t, g.pliCardsCount, 3)
		assert.Equal(t, g.nombrePli, 0)

		c4 := p3.PlayingHand.Cards[0]
		err4 := g.PlayCard(p4.GetID(), p4.PlayingHand.Cards[0])
		pli4 := [4]cards.Card{c1, c2, c3, c4}
		assert.Equal(t, err4, nil)
		assert.Equal(t, g.Plis[0], pli4)
		assert.Equal(t, g.pliCardsCount, 0)
		assert.Equal(t, g.nombrePli, 1)

		cp1 := p1.PlayingHand.Cards[1]
		err = g.PlayCard(p1.GetID(), p1.PlayingHand.Cards[1])
		assert.Equal(t, err, nil)
		pli1 := [4]cards.Card{cp1}
		assert.Equal(t, g.Plis[1], pli1)
		assert.Equal(t, g.pliCardsCount, 1)
		assert.Equal(t, g.nombrePli, 1)
	})

	t.Run("Player One cannot play a card he does not have", func(t *testing.T) {
		g := setupGame(4)
		err := g.DispatchCards()
		assert.NoError(t, err)

		p1, p2 := g.players[0], g.players[1]

		err = g.PlayCard(p1.GetID(), p2.PlayingHand.Cards[0])
		assert.Error(t, err, errors.New("card not found in player hand"))
	})
}

func TestNextRound(t *testing.T) {
	g := NewGame()

	assert.Equal(t, g.NextRound(0), [4]int{0, 1, 2, 3})
	assert.Equal(t, g.NextRound(1), [4]int{1, 2, 3, 0})
	assert.Equal(t, g.NextRound(2), [4]int{2, 3, 0, 1})
	assert.Equal(t, g.NextRound(3), [4]int{3, 0, 1, 2})
}

func TestPlayCardNext(t *testing.T) {
	testcases := []struct {
		name string
		game *Game
		take gametake.GameTake
	}{
		{
			name: gametake.TOUT.Name(),
			take: gametake.TOUT,
			game: setupGame(4),
		},
		{
			name: gametake.CENT.Name(),
			take: gametake.CENT,
			game: setupGame(4),
		},
		{
			name: gametake.COEUR.Name(),
			take: gametake.COEUR,
			game: setupGame(4),
		},
		{
			name: gametake.CARREAU.Name(),
			take: gametake.CARREAU,
			game: setupGame(4),
		},
		{
			name: gametake.PIQUE.Name(),
			take: gametake.PIQUE,
			game: setupGame(4),
		},
		{
			name: gametake.TREFLE.Name(),
			take: gametake.TREFLE,
			game: setupGame(4),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			g := test.game
			p1, p2, p3, p4 := g.players[0], g.players[1], g.players[2], g.players[3]
			err := g.AddTake(0, test.take)
			assert.NoError(t, err)
			if test.take != gametake.TOUT {
				err := g.DispatchCards()
				assert.NoError(t, err)
			}

			err = g.PlayCardNext(p1.GetID(), p1.PlayingHand.Cards[0])
			assert.Equal(t, err, nil)

			err = g.PlayCardNext(p2.GetID(), p2.PlayingHand.Cards[0])
			assert.Equal(t, err, nil)

			err = g.PlayCardNext(p3.GetID(), p3.PlayingHand.Cards[0])
			assert.Equal(t, err, nil)

			err = g.PlayCardNext(p4.GetID(), p4.PlayingHand.Cards[0])
			assert.Equal(t, err, nil)

			pli := [4]cards.Card{
				p1.PlayingHand.Cards[0],
				p2.PlayingHand.Cards[0],
				p3.PlayingHand.Cards[0],
				p4.PlayingHand.Cards[0],
			}
			assert.Equal(t, g.Decks[0].cards, pli)
			winner := g.Decks[0].winner
			fmt.Println("take ", g.GetTake().Name(), "deck", pli, "winner:", winner)
		})
	}
}

func setupGame(playersCount int) *Game {
	g := NewGame()
	for i := 0; i < playersCount; i++ {
		err := g.AddPlayer(player.NewPlayer())
		if err != nil {
			fmt.Println("ERROR SETTING UP A TEST GAME")
		}

	}
	return g
}
