package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/pkg/cards"
)

func TestGetValue(t *testing.T) {
	assert.Equal(t, 1, CARREAU.GetValue())
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Carreau", CARREAU.Name())
}

type evalCardsTestCase struct {
	name             string
	cards            [5]cards.Card
	allCardsValue    int
	cardsOfTakeValue int
}

func TestEvaluateHand(t *testing.T) {
	tc := []evalCardsTestCase{
		evalCardsTestCase{
			"valet et quatorzaine de carreau et trois sept",
			[5]cards.Card{cards.ValetCarreau, cards.NeufCarreau, cards.SeptCarreau, cards.SeptCoeur, cards.SeptPique},
			34,
			34,
		},
		evalCardsTestCase{
			"valet quatorzaine et as de carreau et trois sept",
			[5]cards.Card{cards.ValetCarreau, cards.NeufCarreau, cards.AsCarreau, cards.SeptCoeur, cards.SeptPique},
			45,
			45,
		},
		evalCardsTestCase{
			"valet quatorzaine et as de pique et trois sept",
			[5]cards.Card{cards.ValetPique, cards.NeufPique, cards.AsPique, cards.SeptCoeur, cards.SeptPique},
			13,
			0,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			result := CARREAU.EvaluateHand(c.cards)
			assert.Equal(t, c.allCardsValue, result.AllCardsValue)
			assert.Equal(t, c.cardsOfTakeValue, result.CardsOfTakeValue)
		})
	}
}

type evaluateCardTestCasesCARREAU struct {
	card  cards.Card
	name  string
	value int
}

func TestEvaluateCardCARREAU(t *testing.T) {
	tc := []evaluateCardTestCasesCARREAU{
		evaluateCardTestCasesCARREAU{cards.ValetCarreau, "Valet carreau should be 20", 20},
		evaluateCardTestCasesCARREAU{cards.NeufCarreau, "Neuf carreau should be 14", 14},
		evaluateCardTestCasesCARREAU{cards.AsCarreau, "As carreau should be 11", 11},
		evaluateCardTestCasesCARREAU{cards.DixCarreau, "Dix carreau should be 10", 10},
		evaluateCardTestCasesCARREAU{cards.RoiCarreau, "Roi carreau should be 4", 4},
		evaluateCardTestCasesCARREAU{cards.DameCarreau, "Dame carreau should be 3", 3},
		evaluateCardTestCasesCARREAU{cards.HuitCarreau, "Huit carreau should be 0", 0},
		evaluateCardTestCasesCARREAU{cards.SeptCarreau, "Sept carreau should be 0", 0},

		evaluateCardTestCasesCARREAU{cards.ValetCoeur, "Valet coeur should be 2", 2},
		evaluateCardTestCasesCARREAU{cards.NeufCoeur, "Neuf coeur should be 0", 0},
		evaluateCardTestCasesCARREAU{cards.AsCoeur, "As coeur should be 11", 11},
		evaluateCardTestCasesCARREAU{cards.DixCoeur, "Dix coeur should be 10", 10},
		evaluateCardTestCasesCARREAU{cards.RoiCoeur, "Roi coeur should be 4", 4},
		evaluateCardTestCasesCARREAU{cards.DameCoeur, "Dame coeur should be 3", 3},
		evaluateCardTestCasesCARREAU{cards.HuitCoeur, "Huit coeur should be 0", 0},
		evaluateCardTestCasesCARREAU{cards.SeptCoeur, "Sept coeur should be 0", 0},

		evaluateCardTestCasesCARREAU{cards.ValetTrefle, "Valet trefle should be 2", 2},
		evaluateCardTestCasesCARREAU{cards.NeufTrefle, "Neuf trefle should be 0", 0},
		evaluateCardTestCasesCARREAU{cards.AsTrefle, "As trefle should be 11", 11},
		evaluateCardTestCasesCARREAU{cards.DixTrefle, "Dix trefle should be 10", 10},
		evaluateCardTestCasesCARREAU{cards.RoiTrefle, "Roi trefle should be 4", 4},
		evaluateCardTestCasesCARREAU{cards.DameTrefle, "Dame trefle should be 3", 3},
		evaluateCardTestCasesCARREAU{cards.HuitTrefle, "Huit trefle should be 0", 0},
		evaluateCardTestCasesCARREAU{cards.SeptTrefle, "Sept trefle should be 0", 0},

		evaluateCardTestCasesCARREAU{cards.ValetPique, "Valet pique should be 2", 2},
		evaluateCardTestCasesCARREAU{cards.NeufPique, "Neuf pique should be 0", 0},
		evaluateCardTestCasesCARREAU{cards.AsPique, "As pique should be 11", 11},
		evaluateCardTestCasesCARREAU{cards.DixPique, "Dix pique should be 10", 10},
		evaluateCardTestCasesCARREAU{cards.RoiPique, "Roi pique should be 4", 4},
		evaluateCardTestCasesCARREAU{cards.DamePique, "Dame pique should be 3", 3},
		evaluateCardTestCasesCARREAU{cards.HuitPique, "Huit pique should be 0", 0},
		evaluateCardTestCasesCARREAU{cards.SeptPique, "Sept pique should be 0", 0},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			carreau := Carreau{}
			result, _ := carreau.EvaluateCard(test.card)
			assert.Equal(t, result, test.value)
		})
	}
}

type testGreaterThanCasesCARREAU struct {
	name     string
	take     GameTake
	expected bool
}

func TestCARREAUGreaterThan(t *testing.T) {
	tc := []testGreaterThanCasesCARREAU{
		testGreaterThanCasesCARREAU{"Carreau is greater than tout", TOUT, false},
		testGreaterThanCasesCARREAU{"Carreau is greater than cent", CENT, false},
		testGreaterThanCasesCARREAU{"Carreau is greater than pique", PIQUE, false},
		testGreaterThanCasesCARREAU{"Carreau is greater than coeur", COEUR, false},
		testGreaterThanCasesCARREAU{"Carreau is greater than trefle", TREFLE, true},
		testGreaterThanCasesCARREAU{"Carreau is greater than passe", PASSE, true},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			result := CARREAU.GreaterThan(test.take)
			assert.Equal(t, result, test.expected)
		})
	}
}
