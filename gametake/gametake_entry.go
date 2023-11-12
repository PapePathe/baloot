package gametake

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
	"pathe.co/zinx/pkg/cards"
)

type GameTakeEntry struct {
	AllCardsValue          int
	AllPlayerCardsValue    int
	CardsOfTakeValue       int
	PlayerCardsOfTakeValue int
	OtherCardsValue        int
	ImportantCards         [5]cards.Card
	Flags                  map[string]flag
}

type flag struct {
	name string
}

func (f flag) String() string {
	return f.name
}

var (
	FlagValetOfColor    = flag{"valet-of-color"}
	FlagTwoValets       = flag{"two-valets"}
	FlagTwoAces         = flag{"two-aces"}
	FlagNoAce           = flag{"no-ace"}
	FlagNoValet         = flag{"no-valet"}
	FlagNoTen           = flag{"no-ten"}
	FlagNineOfColor     = flag{"nine-of-color"}
	FlagAceOfOtherColor = flag{"ace-of-other-color"}
	FlagAceOfColor      = flag{"ace-of-color"}
	FlagOneAce          = flag{"one-ace"}
	FlagOneValet        = flag{"one-valet"}
	Flag34Color         = flag{"34-color"}
)

func NewGameTakeEntry() GameTakeEntry {
	return GameTakeEntry{
		AllCardsValue:          0,
		AllPlayerCardsValue:    0,
		CardsOfTakeValue:       0,
		PlayerCardsOfTakeValue: 0,
		OtherCardsValue:        0,
		ImportantCards:         [5]cards.Card{},
		Flags:                  map[string]flag{},
	}
}

func (gte GameTakeEntry) CanTake(gt GameTake) bool {
	if gt == TOUT {
		if _, ok := gte.Flags[FlagTwoValets.name]; ok {
			return true
		}
	}

	takeRatio := gte.TakeRatio(gt)
	ratio := gte.Ratio(gt)

	if takeRatio > 45 {
		return true
	}

	if takeRatio >= 30 && ratio >= 25 {
		return true
	}

	return false
}

func (gte GameTakeEntry) TakeRatio(_ GameTake) int {
	if gte.AllCardsValue > 0 {
		ratio := (gte.PlayerCardsOfTakeValue * 100) / gte.CardsOfTakeValue

		return ratio
	}

	return 0
}

func (gte GameTakeEntry) Ratio(_ GameTake) int {
	if gte.AllCardsValue > 0 {
		ratio := (gte.PlayerCardsOfTakeValue * 100) / gte.AllCardsValue

		return ratio
	}

	return 0
}

func (gte GameTakeEntry) String() string {
	return fmt.Sprintf(
		"AllCards=%d, TakeCards=%d OtherCards=%d",
		gte.AllCardsValue,
		gte.CardsOfTakeValue,
		gte.OtherCardsValue)
}

func (gte GameTakeEntry) Print(gt GameTake) table.Row {
	return table.Row{
		gt.Name(),
		gte.AllCardsValue,
		gte.AllPlayerCardsValue,
		gte.CardsOfTakeValue,
		gte.PlayerCardsOfTakeValue,
		gte.Ratio(gt),
		gte.TakeRatio(gt),
		gte.CanTake(gt),
	}
}

type IGameTakeEntryPersisTence interface {
	Persist() error
}

type IGameTakeEntryBroker interface {
	Publish() error
}
