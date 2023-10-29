package gametake

import (
	"encoding/json"

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

func (t Cent) EvaluateDeck(cards [4]cards.Card) (result int) {
	for _, c := range cards {
		value, _ := t.EvaluateCard(c)
		result += value
	}

	return result
}

func (t Cent) EvaluateHand(cards [5]cards.Card) (entry GameTakeEntry) {
	entry.Flags = make(map[string]flag)
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

func (t Cent) EvaluateCard(card cards.Card) (int, bool) {
	switch card.Genre {
	case "A":
		return 11, true
	case "10":
		return 10, true
	case "R":
		return 4, true
	case "D":
		return 3, true
	case "V":
		return 2, true
	default:
		return 0, true
	}
}

func (t Cent) MarshalJSON() ([]byte, error) {
	customStruct := struct {
		Name string
	}{
		Name: t.Name(),
	}
	return json.Marshal(customStruct)
}
