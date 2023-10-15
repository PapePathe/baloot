package gametake

import (
	"encoding/json"

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
	valetsCount := 0
	entry.Flags = make(map[string]flag)

	for _, c := range cards {
		if c.IsValet() {
			valetsCount++
			if valetsCount == 2 {
				entry.Flags[FlagTwoValets.name] = FlagTwoValets
			}
		}

		result += t.parseCard(c)
	}

	entry.AllCardsValue = t.AllCardsValue
	entry.CardsOfTakeValue = t.AllCardsValue
	entry.OtherCardsValue = 0
	entry.AllPlayerCardsValue = result
	entry.PlayerCardsOfTakeValue = result

	return entry
}

func (t Tout) parseCard(c cards.Card) int {
	value, _ := t.EvaluateCard(c)

	return value
}

func (t Tout) MarshalJSON() ([]byte, error) {
	customStruct := struct {
		Name string
	}{
		Name: t.Name(),
	}
	return json.Marshal(customStruct)
}

func (t Tout) EvaluateCard(card cards.Card) (int, bool) {
	switch card.Genre {
	case "V":
		return 20, true
	case "9":
		return 14, true
	case "A":
		return 11, true
	case "10":
		return 10, true
	case "R":
		return 4, true
	case "D":
		return 3, true
	default:
		return 0, true
	}
}
