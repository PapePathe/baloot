package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/pkg/cards"
)

func TestGameTakeEntryFlags(t *testing.T) {
	t.Parallel()

	tc := []struct {
		name          string
		hand          [5]cards.Card
		cardsValue    int
		allCardsValue int
		flags         map[string]flag
	}{
		{
			hand: [5]cards.Card{
				cards.ValetCarreau,
				cards.ValetPique,
				cards.SeptCarreau,
				cards.HuitCarreau,
				cards.DameTrefle,
			},
			name:          "Hand with valet carreau and valet pique",
			cardsValue:    43,
			allCardsValue: 43,
			flags:         map[string]flag{FlagTwoValets.name: FlagTwoValets},
		},
		{
			hand: [5]cards.Card{
				cards.NeufCarreau,
				cards.ValetPique,
				cards.SeptCarreau,
				cards.HuitCarreau,
				cards.DameTrefle,
			},
			name:          "Hand with valet carreau and valet pique",
			cardsValue:    37,
			allCardsValue: 37,
			flags:         map[string]flag{FlagOneValet.name: FlagOneValet},
		},
	}

	for _, testCase := range tc {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			entry := TOUT.EvaluateHand(testCase.hand)

			assert.Equal(t, entry.Flags, testCase.flags)
			assert.Equal(t, entry.PlayerCardsOfTakeValue, testCase.cardsValue)
			assert.Equal(t, entry.AllPlayerCardsValue, testCase.allCardsValue)

			assert.Equal(t, 0, entry.OtherCardsValue)
			assert.Equal(t, 162, entry.AllCardsValue)
			assert.Equal(t, 162, entry.CardsOfTakeValue)
		})
	}
}
