package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var PIQUE GameTake = Pique{}

type Pique struct{}

func (t Pique) Name() string {
	return "Pique"
}

func (t Pique) GreaterThan(other GameTake) bool {
	return t.GetValue() > other.GetValue()
}

func (t Pique) GetValue() int {
	return 4
}

func (t Pique) EvaluateHand(cards [5]cards.Card) (entry GameTakeEntry) {
	result := 0
	result += t.EvaluateCard(cards[0])
	result += t.EvaluateCard(cards[1])
	result += t.EvaluateCard(cards[2])
	result += t.EvaluateCard(cards[3])
	result += t.EvaluateCard(cards[4])

	return entry
}

func (t Pique) EvaluateCard(card cards.Card) int {
	if card.Couleur == "Pique" {
		return evaluateCardOfColor(card.Genre)
	} else {
		return evaluateCardOfOtherColor(card.Genre)
	}
}
