package player

import (
	"strings"

	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

type Hand struct {
	Cards [5]cards.Card
}

func (h *Hand) String() string {
	var sb strings.Builder
	for _, c := range h.Cards {
		sb.WriteString(c.String())
	}
	return sb.String()
}

type PlayingHand struct {
	Cards []cards.Card
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
