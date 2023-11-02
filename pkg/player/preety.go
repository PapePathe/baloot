package player

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

func (p Player) PreetyHand() {
	tbl := table.NewWriter()

	for key, cards := range p.OrderedCards() {
		var builder strings.Builder
		for _, card := range cards {
			builder.WriteString(card.Genre)
			builder.WriteString(",")
		}

		tbl.AppendRow(table.Row{key, builder.String()})
	}

	fmt.Println(tbl.Render())
}
