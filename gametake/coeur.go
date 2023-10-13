package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var COEUR GameTake = Coeur{
	AllCardsValue:   152,
	CardsValue:      62,
	OtherCardsValue: 90,
}

type Coeur struct {
	AllCardsValue   int
	CardsValue      int
	OtherCardsValue int
}

func (t Coeur) Name() string {
	return "Coeur"
}

func (t Coeur) GreaterThan(other GameTake) bool {
	return t.GetValue() > other.GetValue()
}

func (t Coeur) GetValue() int {
	return 3
}

func (t Coeur) EvaluateHand(cards [5]cards.Card) (entry GameTakeEntry) {
	t.parseCard(&entry, cards[0])
	t.parseCard(&entry, cards[1])
	t.parseCard(&entry, cards[2])
	t.parseCard(&entry, cards[3])
	t.parseCard(&entry, cards[4])

	entry.AllCardsValue = t.AllCardsValue
	entry.CardsOfTakeValue = t.CardsValue

	return entry
}

func (t Coeur) parseCard(gt *GameTakeEntry, c cards.Card) {
	value, sameColor := t.EvaluateCard(c)
	gt.AllCardsValue += value
	if sameColor {
		gt.PlayerCardsOfTakeValue += value
		gt.AllPlayerCardsValue += value
	} else {
		gt.OtherCardsValue += value
		gt.AllPlayerCardsValue += value
	}
}

func (t Coeur) EvaluateCard(card cards.Card) (value int, sameColor bool) {
	if card.Couleur == "Coeur" {
		value = evaluateCardOfColor(card.Genre)
		sameColor = true
	} else {
		value = evaluateCardOfOtherColor(card.Genre)
		sameColor = false
	}
	return value, sameColor
}
