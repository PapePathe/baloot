package gametake

import (
	"pathe.co/zinx/pkg/cards"
)

type GameTake interface {
	GreaterThan(t GameTake) bool
	EvaluateHand([5]cards.Card) GameTakeEntry
	GetValue() int
	Name() string
}

func evaluateCardOfOtherColor(genre string) int {
	switch genre {
	case "V":
		return 2
	case "9":
		return 0
	case "A":
		return 11
	case "10":
		return 10
	case "R":
		return 3
	case "D":
		return 2
	case "8":
		return 0
	case "7":
		return 0
	}

	return 0
}

func evaluateCardOfColor(genre string) int {
	switch genre {
	case "V":
		return 20
	case "9":
		return 14
	case "A":
		return 11
	case "10":
		return 10
	case "R":
		return 3
	case "D":
		return 2
	case "8":
		return 0
	case "7":
		return 0
	}

	return 0
}
