package gametake

import (
	"fmt"

	"pathe.co/zinx/pkg/cards"
)

type GameTakeEntry struct {
	AllCardsValue    int
	CardsOfTakeValue int
	OtherCardsValue  int
	ImportantCards   [5]cards.Card
}

func (gte GameTakeEntry) String() string {
	return fmt.Sprintf(
		"All=%d, Take=%d Other=%d",
		gte.AllCardsValue,
		gte.CardsOfTakeValue,
		gte.OtherCardsValue)
}
