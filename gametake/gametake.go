package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var TREFLE ColorTake = ColorTake{
	CardsValue:      62,
	AllCardsValue:   152,
	OtherCardsValue: 90,
	Couleur:         "Trefle",
	Value:           1,
}
var CARREAU ColorTake = ColorTake{
	CardsValue:      62,
	AllCardsValue:   152,
	OtherCardsValue: 90,
	Couleur:         "Carreau",
	Value:           2,
}
var COEUR ColorTake = ColorTake{
	CardsValue:      62,
	AllCardsValue:   152,
	OtherCardsValue: 90,
	Couleur:         "Coeur",
	Value:           3,
}

var PIQUE ColorTake = ColorTake{
	CardsValue:      62,
	AllCardsValue:   152,
	OtherCardsValue: 90,
	Couleur:         "Pique",
	Value:           4,
}
var AllTakes []GameTake = []GameTake{PASSE, TREFLE, CARREAU, COEUR, PIQUE, CENT, TOUT}
var AllTakesByName map[string]GameTake = map[string]GameTake{
	"Passe":   PASSE,
	"Trefle":  TREFLE,
	"Carreau": CARREAU,
	"Coeur":   COEUR,
	"Pique":   PIQUE,
	"Cent":    CENT,
	"Tout":    TOUT,
}

type GameTake interface {
	GreaterThan(t GameTake) bool
	EvaluateHand([5]cards.Card) GameTakeEntry
	EvaluateCard(cards.Card) (int, bool)
	GetValue() int
	Name() string
}
