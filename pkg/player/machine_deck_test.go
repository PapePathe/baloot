package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
