package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var CARREAU Carreau = Carreau{
	CardsValue:      62,
	AllCardsValue:   152,
	OtherCardsValue: 90,
}

type Carreau struct {
	CardsValue      int
	AllCardsValue   int
	OtherCardsValue int
}

func (t Carreau) GreaterThan(other GameTake) bool {
	return t.GetValue() > other.GetValue()
}

func (t Carreau) Name() string {
	return "Carreau"
}

func (t Carreau) GetValue() int {
	return 2
}

func (t Carreau) EvaluateHand(cards [5]cards.Card) (entry GameTakeEntry) {
	t.parseCard(&entry, cards[0])
	t.parseCard(&entry, cards[1])
	t.parseCard(&entry, cards[2])
	t.parseCard(&entry, cards[3])
	t.parseCard(&entry, cards[4])

	return entry
}

func (t Carreau) parseCard(gt *GameTakeEntry, c cards.Card) {
	value, sameColor := t.EvaluateCard(c)
	gt.AllCardsValue += value
	if sameColor {
		gt.CardsOfTakeValue += value
	} else {
		gt.OtherCardsValue += value
	}
}

func (t Carreau) EvaluateCard(card cards.Card) (value int, sameColor bool) {
	if card.Couleur == "Carreau" {
		value = evaluateCardOfColor(card.Genre)
		sameColor = true
	} else {
		value = evaluateCardOfOtherColor(card.Genre)
		sameColor = false
	}

	return value, sameColor
}
