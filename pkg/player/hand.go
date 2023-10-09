package player

import (
	"strings"

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
