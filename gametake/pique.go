package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var PIQUE GameTake = Pique{
	AllCardsValue:   152,
	CardsValue:      62,
	OtherCardsValue: 90,
}

type Pique struct {
	AllCardsValue   int
	CardsValue      int
	OtherCardsValue int
}

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
	t.parseCard(&entry, cards[0])
	t.parseCard(&entry, cards[1])
	t.parseCard(&entry, cards[2])
	t.parseCard(&entry, cards[3])
	t.parseCard(&entry, cards[4])

	entry.AllCardsValue = t.AllCardsValue
	entry.CardsOfTakeValue = t.CardsValue

	return entry
}

func (t Pique) parseCard(gt *GameTakeEntry, c cards.Card) {
	value, sameColor := t.EvaluateCard(c)
	if sameColor {
		gt.PlayerCardsOfTakeValue += value
		gt.AllPlayerCardsValue += value
	} else {
		gt.OtherCardsValue += value
		gt.AllPlayerCardsValue += value
	}
}

func (t Pique) EvaluateCard(card cards.Card) (value int, sameColor bool) {
	if card.Couleur == "Pique" {
		value = evaluateCardOfColor(card.Genre)
		sameColor = true
	} else {
		value = evaluateCardOfOtherColor(card.Genre)
		sameColor = false
	}
	return value, sameColor
}
