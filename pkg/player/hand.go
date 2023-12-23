package player

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

type Hand struct {
	Cards [5]cards.Card `json:"cards"`
}

func (h *Hand) String() string {
	var sb strings.Builder
	for _, c := range h.Cards {
		sb.WriteString(c.String())
	}

	return sb.String()
}

type PlayingHand struct {
	Cards []cards.Card `json:"cards"`
}

func (h *PlayingHand) LowestCard(t gametake.GameTake) cards.Card {
	c := h.Cards[0]

	for _, cc := range h.Cards {
		log.Trace().Msg(fmt.Sprintf("Comparing cards (%s %d) vs (%s %d)", c, t.EvaluateCardForWin(c), cc, t.EvaluateCardForWin(cc)))
		if !c.IsNotEmpty() {
			c = cc
		}
		if cc.IsNotEmpty() && t.EvaluateCardForWin(c) >= t.EvaluateCardForWin(cc) {
			c = cc
		}
	}
	return c
}

func (h PlayingHand) OrderedCardsForPlaying(take gametake.GameTake) []cards.Card {
	cardsMap := make(map[string][]cards.Card)

	for _, card := range h.Cards {
		if !card.IsNotEmpty() {
			continue
		}
		_, ok := cardsMap[card.Couleur]

		if ok {
			cardsMap[card.Couleur] = append(cardsMap[card.Couleur], card)
		} else {
			cardsMap[card.Couleur] = []cards.Card{card}
		}

		sorter := SortByColorAndType{cardsMap[card.Couleur], take}
		sort.Sort(sorter)
	}

	keys := make([]string, 0, len(cardsMap))
	for k := range cardsMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	result := []cards.Card{}

	for _, key := range keys {
		mapCards := cardsMap[key]
		result = append(result, mapCards...)
	}

	return result
}

func (h *PlayingHand) HasColor(c string) (bool, []cards.Card) {
	foundCards := []cards.Card{}

	for _, pc := range h.Cards {
		if pc.Couleur == c {
			foundCards = append(foundCards, pc)
		}
	}

	return len(foundCards) > 0, foundCards
}

func (h *PlayingHand) String() string {
	var sb strings.Builder
	for _, c := range h.Cards {
		sb.WriteString(c.String())
	}

	return sb.String()
}

type SortByColorAndType struct {
	Cards []cards.Card
	Take  gametake.GameTake
}

func (a SortByColorAndType) Len() int { return len(a.Cards) }

func (a SortByColorAndType) Swap(i, j int) {
	a.Cards[i], a.Cards[j] = a.Cards[j], a.Cards[i]
}

func (a SortByColorAndType) Less(i, j int) bool {
	ivalue, _ := a.Take.EvaluateCard(a.Cards[i])
	jvalue, _ := a.Take.EvaluateCard(a.Cards[j])

	return ivalue > jvalue
}
