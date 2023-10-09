package player

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
)

func (p Player) PreetyHand() {
	t := table.NewWriter()
	t.SetCaption("Users")

	t.AppendHeader(table.Row{"Family", "Color"})
	for _, card := range p.Hand.Cards {
		t.AppendRow(table.Row{card.Genre, card.Couleur})
	}

	fmt.Println(t.Render())
}
