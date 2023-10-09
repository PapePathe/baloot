package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var PASSE Passe = Passe{Name: "Passe"}

type Passe struct {
	Value int32 `default:"-1"`
	Name  string
}

func (t Passe) GreaterThan(other GameTake) bool {
	return t.GetValue() > other.GetValue()
}

func (t *Passe) GetValue() int {
	return 0
}

func (t Passe) EvaluateHand(cards [5]cards.Card) int {
	return 0
}

func (t Passe) EvaluateCard(card cards.Card) int {
	return 0
}
