package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

func TestSetTake(t *testing.T) {
	t.Parallel()

	p := NewPlayer()
	p.SetTake(&gametake.TOUT)

	assert.Equal(t, p.Take, &gametake.TOUT)
}

func TestNewHasCard(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		hand      []cards.Card
		card      cards.Card
		cardIndex int
		hasCard   bool
	}{
		{
			name:      "with no cards",
			hand:      []cards.Card{},
			card:      cards.ValetCoeur,
			cardIndex: -1,
			hasCard:   false,
		},
		{
			name:      "with cards and a card in hand",
			hand:      []cards.Card{cards.ValetCoeur, cards.ValetPique},
			card:      cards.ValetCoeur,
			cardIndex: 0,
			hasCard:   true,
		},
		{
			name:      "with cards and a card not present in hand",
			hand:      []cards.Card{cards.ValetCoeur, cards.ValetPique},
			card:      cards.AsPique,
			cardIndex: -1,
			hasCard:   false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			p := NewPlayer()
			p.PlayingHand.Cards = test.hand
			hasCard, cardIndex := p.HasCard(test.card)

			assert.Equal(t, test.hasCard, hasCard)
			assert.Equal(t, test.cardIndex, cardIndex)
		})
	}
}

func TestNewPlayer(t *testing.T) {
	t.Parallel()

	p := NewPlayer()
	assert.Nil(t, p.Take)
	assert.Nil(t, p.Conn)
}

func TestGetID(t *testing.T) {
	t.Parallel()

	p := NewPlayer()
	p.SetID(2)
	assert.Equal(t, 2, p.GetID())
}

func TestSetID(t *testing.T) {
	t.Parallel()

	p := NewPlayer()
	p.SetID(0)
	assert.Equal(t, 0, p.ID)
}

func TestOrderedCards(t *testing.T) {
	t.Parallel()

	p := NewPlayer()
	p.Hand = Hand{[5]cards.Card{cards.ValetPique, cards.HuitCoeur, cards.NeufPique, cards.NeufTrefle, cards.ValetCoeur}}
	expected := map[string][]cards.Card{
		"Pique":  {cards.ValetPique, cards.NeufPique},
		"Coeur":  {cards.HuitCoeur, cards.ValetCoeur},
		"Trefle": {cards.NeufTrefle},
	}

	assert.Equal(t, expected, p.OrderedCards())
}
