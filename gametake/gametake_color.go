package gametake

import (
	"encoding/json"
	"fmt"

	"pathe.co/zinx/pkg/cards"
)

type ColorTake struct {
	CardsValue      int
	AllCardsValue   int
	OtherCardsValue int
	Couleur         string
	Value           int
}

func (t ColorTake) GreaterThan(other GameTake) bool {
	return t.GetValue() > other.GetValue()
}

func (t ColorTake) Name() string {
	return t.Couleur
}

func (t ColorTake) GetValue() int {
	return t.Value
}

func (t ColorTake) EvaluateDeck(cards [4]cards.Card) int {
	result := 0

	for _, c := range cards {
		value, _ := t.EvaluateCard(c)
		result += value
	}

	return result
}

func (t ColorTake) EvaluateHand(cards [5]cards.Card) GameTakeEntry {
	entry := NewGameTakeEntry()

	t.parseCard(&entry, cards[0])
	t.parseCard(&entry, cards[1])
	t.parseCard(&entry, cards[2])
	t.parseCard(&entry, cards[3])
	t.parseCard(&entry, cards[4])
	entry.AllCardsValue = t.AllCardsValue
	entry.CardsOfTakeValue = t.CardsValue

	return entry
}

func (t ColorTake) parseCard(gt *GameTakeEntry, c cards.Card) {
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

func (t ColorTake) EvaluateCard(card cards.Card) (int, bool) {
	if card.Couleur == t.Couleur {
		return evaluateCardOfColor(card.Genre), false
	}

	return evaluateCardOfOtherColor(card.Genre), false
}

var WinValues = map[string]int{
	"V": 20, "9": 14, "A": 11, "10": 10, "R": 4, "D": 3, "8": 1, "7": 0,
}

var OtherWinValues = map[string]int{
	"V": 2, "9": 0, "A": 11, "10": 10, "R": 4, "D": 3,
}

func (t ColorTake) EvaluateCardForWin(_ cards.Card) int {
	return 0
}

func (t ColorTake) Winner(a cards.Card, b cards.Card) cards.Card {
	if a.Couleur == b.Couleur {
		if a.Couleur == t.Couleur {
			aValue, bValue := evaluateCardOfColor(a.Genre), evaluateCardOfColor(b.Genre)
			if aValue > bValue {
				return a
			}

			return b
		}

		aValue, bValue := evaluateCardOfOtherColor(a.Genre), evaluateCardOfOtherColor(b.Genre)
		if aValue > bValue {
			return a
		}

		return b
	}

	if a.Couleur == t.Couleur {
		return a
	}

	if b.Couleur == t.Couleur {
		return b
	}

	return b
}

func evaluateCardOfOtherColor(genre string) int {
	switch genre {
	case "V":
		return 2
	case "9":
		return 0
	case "A":
		return 11
	case "10":
		return 10
	case "R":
		return 4
	case "D":
		return 3
	case "8":
		return 0
	case "7":
		return 0
	}

	return 0
}

func evaluateCardOfColor(genre string) int {
	switch genre {
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
	case "8":
		return 0
	case "7":
		return 0
	}

	return 0
}

func (t ColorTake) MarshalJSON() ([]byte, error) {
	customStruct := struct {
		Name string `json:"name"`
	}{
		Name: t.Name(),
	}

	result, err := json.Marshal(customStruct)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling ColorTake  %w", err)
	}

	return result, nil
}
