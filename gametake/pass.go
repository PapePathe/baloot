package gametake

import (
	"encoding/json"
	"fmt"

	"pathe.co/zinx/pkg/cards"
)

var PASSE = Passe{}

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

func (t Passe) EvaluateDeck(_ [4]cards.Card) int {
	return 0
}

func (t Passe) EvaluateHand(_ [5]cards.Card) GameTakeEntry {
	return NewGameTakeEntry()
}

func (t Passe) EvaluateCard(_ cards.Card) (int, bool) {
	return 0, true
}

func (t Passe) EvaluateCardForWin(_ cards.Card) int {
	return 0
}

func (t Passe) Winner(a cards.Card, _ cards.Card) cards.Card {
	return a
}

func (t Passe) MarshalJSON() ([]byte, error) {
	customStruct := struct {
		Name string `json:"name"`
	}{
		Name: t.Name(),
	}
	result, err := json.Marshal(customStruct)

	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling pass %w", err)
	}

	return result, nil
}
