package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

var TREFLE = ColorTake{
	CardsValue:      62,
	AllCardsValue:   152,
	OtherCardsValue: 90,
	Couleur:         "Trefle",
	Value:           1,
}

var CARREAU = ColorTake{
	CardsValue:      62,
	AllCardsValue:   152,
	OtherCardsValue: 90,
	Couleur:         "Carreau",
	Value:           2,
}

var COEUR = ColorTake{
	CardsValue:      62,
	AllCardsValue:   152,
	OtherCardsValue: 90,
	Couleur:         "Coeur",
	Value:           3,
}

var PIQUE = ColorTake{
	CardsValue:      62,
	AllCardsValue:   152,
	OtherCardsValue: 90,
	Couleur:         "Pique",
	Value:           4,
}

var AllTakes = []GameTake{PASSE, TREFLE, CARREAU, COEUR, PIQUE, CENT, TOUT}
var AllTakeNames = []string{"Passe", "Trefle", "Carreau", "Coeur", "Pique", "Cent", "Tout"}

var AllTakesByName = map[string]GameTake{
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
	EvaluateHand(hand [5]cards.Card) GameTakeEntry
	EvaluateCard(c cards.Card) (int, bool)
	EvaluateCardForWin(c cards.Card) int
	EvaluateDeck(deck [4]cards.Card) int
	Winner(c cards.Card, w cards.Card) cards.Card
	GetValue() int
	Name() string
}
