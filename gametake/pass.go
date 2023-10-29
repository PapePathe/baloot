package gametake

import (
	"encoding/json"

	"pathe.co/zinx/pkg/cards"
)

var PASSE Passe = Passe{}

type Passe struct{}

func (t Passe) GreaterThan(other GameTake) bool {
	return t.GetValue() > other.GetValue()
}

func (t Passe) Name() string {
	return "Passe"
}

func (t Passe) GetValue() int {
	return 0
}

func (t Passe) EvaluateDeck(cards [4]cards.Card) int {
	return 0
}

func (t Passe) EvaluateHand(cards [5]cards.Card) (entry GameTakeEntry) {
	return entry
}

func (t Passe) EvaluateCard(card cards.Card) (int, bool) {
	return 0, true
}

func (t Passe) MarshalJSON() ([]byte, error) {
	customStruct := struct {
		Name string
	}{
		Name: t.Name(),
	}
	return json.Marshal(customStruct)
}
