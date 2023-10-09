package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var COEUR GameTake = Coeur{}

type Coeur struct{}

func (t Coeur) Name() string {
	return "Coeur"
}

func (t Coeur) GreaterThan(other GameTake) bool {
	return t.GetValue() > other.GetValue()
}

func (t Coeur) GetValue() int {
	return 2
}

func (t Coeur) EvaluateHand(cards [5]cards.Card) (entry GameTakeEntry) {
	result := 0
	result += t.EvaluateCard(cards[0])
	result += t.EvaluateCard(cards[1])
	result += t.EvaluateCard(cards[2])
	result += t.EvaluateCard(cards[3])
	result += t.EvaluateCard(cards[4])

	return entry
}

func (t Coeur) EvaluateCard(card cards.Card) int {
	if card.Couleur == "Coeur" {
		return evaluateCardOfColor(card.Genre)
	} else {
		return evaluateCardOfOtherColor(card.Genre)
	}
}
