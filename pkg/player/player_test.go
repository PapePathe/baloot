package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

func TestSetTake(t *testing.T) {
	p := NewPlayer()
	p.SetTake(&gametake.TOUT)
	assert.Equal(t, p.Take, &gametake.TOUT)
}

func TestNewPlayer(t *testing.T) {
	p := NewPlayer()
	assert.Equal(t, p.Transport, JSONMarshaler{})
}

func TestSetID(t *testing.T) {
	p := NewPlayer()
	p.SetID(0)
	assert.Equal(t, p.id, 0)
}

func TestOrderedCards(t *testing.T) {
	p := NewPlayer()
	p.Hand = Hand{[5]cards.Card{cards.ValetPique, cards.HuitCoeur, cards.NeufPique, cards.NeufTrefle, cards.ValetCoeur}}
	expected := map[string][]cards.Card{"Pique": []cards.Card{cards.ValetPique, cards.NeufPique}, "Coeur": []cards.Card{cards.HuitCoeur, cards.ValetCoeur}, "Trefle": []cards.Card{cards.NeufTrefle}}

	assert.Equal(t, p.OrderedCards(), expected)
}
