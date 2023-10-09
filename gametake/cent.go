package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var CENT GameTake = Cent{}

type Cent struct{}

func (t Cent) Name() string {
	return "Coeur"
}

func (t Cent) GreaterThan(other GameTake) bool {
	return t.GetValue() > other.GetValue()
}

func (t Cent) GetValue() int {
	return 5
}

func (t Cent) EvaluateHand(cards [5]cards.Card) (entry GameTakeEntry) {
	result := 0
	result += t.EvaluateCard(cards[0])
	result += t.EvaluateCard(cards[1])
	result += t.EvaluateCard(cards[2])
	result += t.EvaluateCard(cards[3])
	result += t.EvaluateCard(cards[4])

	return entry
}

func (t Cent) EvaluateCard(card cards.Card) int {
	switch card.Genre {
	case "A":
		return 11
	case "10":
		return 10
	case "R":
		return 4
	case "D":
		return 3
	case "V":
		return 2
	default:
		return 0
	}
}
