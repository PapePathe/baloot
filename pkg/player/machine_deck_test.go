package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

func TestFindWinnerOneCardOnDeck(t *testing.T) {
	t.Parallel()

	m := machineDeck{
		cards:    [4]cards.Card{cards.ValetCarreau},
		gametake: gametake.TOUT,
		hand:     PlayingHand{Cards: []cards.Card{}},
	}
	assert.Equal(t, cards.ValetCarreau, m.FindWinner())
}

func TestFindWinnerTwoCardsOnDeck(t *testing.T) {
	t.Parallel()

	m := machineDeck{
		cards:    [4]cards.Card{cards.SeptCarreau, cards.NeufCarreau},
		gametake: gametake.TOUT,
		hand:     PlayingHand{Cards: []cards.Card{}},
	}
	assert.Equal(t, cards.NeufCarreau, m.FindWinner())
}

func TestFindWinnerThreeCardsOnDeck(t *testing.T) {
	t.Parallel()

	m := machineDeck{
		cards:    [4]cards.Card{cards.SeptCarreau, cards.NeufCarreau, cards.ValetCoeur},
		gametake: gametake.TOUT,
		hand:     PlayingHand{Cards: []cards.Card{}},
	}
	assert.Equal(t, cards.NeufCarreau, m.FindWinner())
}

func TestFindWinnerFourCardsOnDeck(t *testing.T) {
	t.Parallel()

	m := machineDeck{
		cards:    [4]cards.Card{cards.SeptCarreau, cards.NeufCarreau, cards.ValetCoeur, cards.ValetCarreau},
		gametake: gametake.TOUT,
		hand:     PlayingHand{Cards: []cards.Card{}},
	}
	assert.Equal(t, cards.ValetCarreau, m.FindWinner())
}

func TestRemainingCardsCount(t *testing.T) {
	t.Parallel()

	t.Run("with zero cards", func(t *testing.T) {
		assert.Equal(t, 0, machineDeck{}.RemainingCardsCount())
	})

	t.Run("with only empty cards", func(t *testing.T) {
		ph := PlayingHand{Cards: []cards.Card{cards.Card{}, cards.Card{}}}
		m := machineDeck{hand: ph}

		assert.Equal(t, 0, m.RemainingCardsCount())
	})

	t.Run("with two cards and empty values", func(t *testing.T) {
		ph := PlayingHand{Cards: []cards.Card{cards.ValetCarreau, cards.AsCoeur, cards.Card{}, cards.Card{}}}
		m := machineDeck{hand: ph}

		assert.Equal(t, 2, m.RemainingCardsCount())
	})
}

func TestAttemptWin(t *testing.T) {
	t.Parallel()

	ph := PlayingHand{Cards: []cards.Card{cards.ValetCarreau, cards.AsCoeur, cards.NeufCarreau}}
	m := machineDeck{hand: ph, cards: [4]cards.Card{cards.RoiCoeur}, gametake: gametake.TOUT}
	c, err := m.AttemptWin(ph.Cards)

	require.NoError(t, err)
	assert.Equal(t, cards.AsCoeur, c)

	ph = PlayingHand{Cards: []cards.Card{cards.ValetCarreau, cards.AsCoeur, cards.NeufCarreau}}
	m = machineDeck{hand: ph, cards: [4]cards.Card{cards.SeptCarreau}, gametake: gametake.TOUT}
	c, err = m.AttemptWin(ph.Cards)

	require.NoError(t, err)
	assert.Equal(t, cards.NeufCarreau, c)

	ph = PlayingHand{Cards: []cards.Card{cards.AsCoeur, cards.NeufCarreau}}
	m = machineDeck{hand: ph, cards: [4]cards.Card{cards.SeptCarreau}, gametake: gametake.TOUT}
	c, err = m.AttemptWin(ph.Cards)

	require.NoError(t, err)
	assert.Equal(t, cards.NeufCarreau, c)

	ph = PlayingHand{Cards: []cards.Card{cards.ValetCarreau, cards.AsCoeur, cards.NeufCarreau}}
	m = machineDeck{hand: ph, cards: [4]cards.Card{cards.SeptPique}, gametake: gametake.TOUT}
	c, err = m.AttemptWin(ph.Cards)

	require.Error(t, err)
	assert.Equal(t, cards.Card{}, c)
}

func TestWinningOrLowestCard(t *testing.T) {
	t.Parallel()

	m := machineDeck{
		cards:    [4]cards.Card{cards.SeptCarreau},
		gametake: gametake.TOUT,
		hand:     PlayingHand{Cards: []cards.Card{cards.NeufCarreau, cards.ValetCoeur, cards.ValetCarreau}},
	}
	assert.Equal(t, cards.NeufCarreau, m.WinningOrLowestCard())
}
