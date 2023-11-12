package gametake

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
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

func (t Tout) EvaluateDeck(cards [4]cards.Card) int {
	result := 0

	for _, c := range cards {
		value, _ := t.EvaluateCard(c)
		result += value
	}

	return result
}

func (t Tout) EvaluateHand(hand [5]cards.Card) GameTakeEntry {
	result := 0
	valets, nines, aces := []cards.Card{}, []cards.Card{}, []cards.Card{}
	entry := NewGameTakeEntry()

	for _, c := range hand {
		if c.IsValet() {
			valets = append(valets, c)

			if len(valets) == 2 {
				entry.Flags[FlagTwoValets.name] = FlagTwoValets
			}
			log.Debug().Int("nombre valets", len(valets)).Msg("")
		}

		if c.IsNine() {
			nines = append(nines, c)
		}

		if c.IsAce() {
			aces = append(aces, c)
			log.Debug().Int("nombre aces", len(aces)).Msg("")
		}

		result += t.parseCard(c)
	}

	if 0 == len(valets) {
		entry.Flags[FlagNoValet.name] = FlagNoValet
	}

	if 1 == len(valets) {
		entry.Flags[FlagOneValet.name] = FlagOneValet
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
		Name string `json:"name"`
	}{
		Name: t.Name(),
	}

	result, err := json.Marshal(customStruct)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling tout  %w", err)
	}

	return result, nil
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

func (t Tout) Winner(a cards.Card, b cards.Card) cards.Card {
	aValue, bValue := t.EvaluateCardForWin(a), t.EvaluateCardForWin(b)

	if a.Couleur == b.Couleur && aValue > bValue {
		return a
	}

	return b
}

func (t Tout) EvaluateCardForWin(card cards.Card) int {
	switch card.Genre {
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
		return 2
	case "7":
		return 1
	default:
		return 0
	}
}
