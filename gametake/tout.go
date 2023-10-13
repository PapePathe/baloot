package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var TOUT GameTake = Tout{AllCardsValue: 162, CardsOfTakeValue: 162}

type Tout struct {
	AllCardsValue    int
	CardsOfTakeValue int
}

func (t Tout) Name() string {
	return "Tout"
}

func (t Tout) GreaterThan(other GameTake) bool {
	return t.GetValue() > other.GetValue()
}

func (t Tout) GetValue() int {
	return 6
}

func (t Tout) EvaluateHand(cards [5]cards.Card) (entry GameTakeEntry) {
	result := 0
	result += t.EvaluateCard(cards[0])
	result += t.EvaluateCard(cards[1])
	result += t.EvaluateCard(cards[2])
	result += t.EvaluateCard(cards[3])
	result += t.EvaluateCard(cards[4])

	entry.AllCardsValue = t.AllCardsValue
	entry.CardsOfTakeValue = t.AllCardsValue
	entry.OtherCardsValue = 0
	entry.AllPlayerCardsValue = result
	entry.PlayerCardsOfTakeValue = result

	return entry
}

func (t Tout) EvaluateCard(card cards.Card) int {
	switch card.Genre {
	case "V":
		return 20
	case "9":
		return 14
	case "A":
		return 11
	case "10":
		return 10
	case "R":
		return 4
	case "D":
		return 3
	default:
		return 0
	}
}
