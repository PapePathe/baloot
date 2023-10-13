package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var CENT GameTake = Cent{AllCardsValue: 120}

type Cent struct {
	AllCardsValue int
}

func (t Cent) Name() string {
	return "Cent"
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
	entry.CardsOfTakeValue = t.AllCardsValue
	entry.AllCardsValue = t.AllCardsValue
	entry.PlayerCardsOfTakeValue = result
	entry.AllPlayerCardsValue = result

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
