package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var TREFLE GameTake = Trefle{}

type Trefle struct{}

func (t Trefle) Name() string {
	return "Trefle"
}

func (t Trefle) GreaterThan(other GameTake) bool {
	return t.GetValue() > other.GetValue()
}

func (t Trefle) GetValue() int {
	return 1
}

func (t Trefle) EvaluateHand(cards [5]cards.Card) (entry GameTakeEntry) {
	result := 0
	result += t.EvaluateCard(cards[0])
	result += t.EvaluateCard(cards[1])
	result += t.EvaluateCard(cards[2])
	result += t.EvaluateCard(cards[3])
	result += t.EvaluateCard(cards[4])

	return entry
}

func (t Trefle) EvaluateCard(card cards.Card) int {
	if card.Couleur == "Trefle" {
		return evaluateCardOfColor(card.Genre)
	} else {
		return evaluateCardOfOtherColor(card.Genre)
	}
}
