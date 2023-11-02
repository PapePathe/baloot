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

func TestNewPlayer(t *testing.T) {
	t.Parallel()

	p := NewPlayer()
	assert.Equal(t, JSONMarshaler{}, p.Transport)
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
