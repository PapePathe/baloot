package player

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

func (p Player) PreetyHand() {
	t := table.NewWriter()

	for key, cards := range p.OrderedCards() {
		var sb strings.Builder
		for _, card := range cards {
			sb.WriteString(card.Genre)
			sb.WriteString(",")
		}

		t.AppendRow(table.Row{key, sb.String()})
	}

	fmt.Println(t.Render())
}
