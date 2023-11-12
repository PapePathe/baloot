package gametake

import (
	"encoding/json"
	"fmt"

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

func (t Cent) EvaluateDeck(cards [4]cards.Card) int {
	result := 0

	for _, c := range cards {
		value, _ := t.EvaluateCard(c)
		result += value
	}

	return result
}

func (t Cent) EvaluateHand(cards [5]cards.Card) GameTakeEntry {
	entry := NewGameTakeEntry()
	result, acesCount := 0, 0

	for _, c := range cards {
		result += t.parseCard(c)

		if c.IsAce() {
			acesCount++

			if acesCount == 2 {
				entry.Flags[FlagTwoAces.name] = FlagTwoAces
			}
		}
	}

	if 0 == acesCount {
		entry.Flags[FlagNoAce.name] = FlagNoAce
	}

	if 1 == acesCount {
		entry.Flags[FlagOneAce.name] = FlagOneAce
	}

	entry.CardsOfTakeValue = t.AllCardsValue
	entry.AllCardsValue = t.AllCardsValue
	entry.PlayerCardsOfTakeValue = result
	entry.AllPlayerCardsValue = result

	return entry
}

func (t Cent) parseCard(c cards.Card) int {
	value, _ := t.EvaluateCard(c)

	return value
}

var centEvaluateValues = map[string]int{"A": 11, "10": 10, "R": 4, "D": 3, "V": 2}

func (t Cent) EvaluateCard(card cards.Card) (int, bool) {
	result, ok := centEvaluateValues[card.Genre]

	if !ok {
		return 0, true
	}

	return result, true
}

func (t Cent) Winner(a cards.Card, b cards.Card) cards.Card {
	aValue, bValue := t.EvaluateCardForWin(a), t.EvaluateCardForWin(b)

	if a.Couleur == b.Couleur && aValue > bValue {
		return a
	}

	return b
}

var centWinValues = map[string]int{"A": 11, "10": 10, "R": 4, "D": 3, "V": 2, "9": 1, "8": 0, "7": -1}

func (t Cent) EvaluateCardForWin(card cards.Card) int {
	result, ok := centWinValues[card.Genre]

	if !ok {
		return -1
	}

	return result
}

func (t Cent) MarshalJSON() ([]byte, error) {
	customStruct := struct {
		Name string `json:"name"`
	}{
		Name: t.Name(),
	}

	result, err := json.Marshal(customStruct)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling take  %w", err)
	}

	return result, nil
}
