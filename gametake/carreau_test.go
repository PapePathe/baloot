package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/pkg/cards"
)

func TestGetValue(t *testing.T) {
	assert.Equal(t, 1, CARREAU.GetValue())
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Carreau", CARREAU.Name())
}

type evalCardsTestCase struct {
	name             string
	cards            [5]cards.Card
	allCardsValue    int
	cardsOfTakeValue int
}

func TestEvaluateHand(t *testing.T) {
	tc := []evalCardsTestCase{
		evalCardsTestCase{
			"valet et quatorzaine de carreau et trois sept",
			[5]cards.Card{cards.ValetCarreau, cards.NeufCarreau, cards.SeptCarreau, cards.SeptCoeur, cards.SeptPique},
			34,
			34,
		},
		evalCardsTestCase{
			"valet quatorzaine et as de carreau et trois sept",
			[5]cards.Card{cards.ValetCarreau, cards.NeufCarreau, cards.AsCarreau, cards.SeptCoeur, cards.SeptPique},
			45,
			45,
		},
		evalCardsTestCase{
			"valet quatorzaine et as de pique et trois sept",
			[5]cards.Card{cards.ValetPique, cards.NeufPique, cards.AsPique, cards.SeptCoeur, cards.SeptPique},
			13,
			0,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			result := CARREAU.EvaluateHand(c.cards)
			assert.Equal(t, c.allCardsValue, result.AllCardsValue)
			assert.Equal(t, c.cardsOfTakeValue, result.CardsOfTakeValue)
		})
	}
}
