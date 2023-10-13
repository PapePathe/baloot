package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/pkg/cards"
)

type evaluateCardTestCasesTREFLE struct {
	card  cards.Card
	name  string
	value int
}

func TestEvaluateCardTREFLE(t *testing.T) {
	tc := []evaluateCardTestCasesTREFLE{
		evaluateCardTestCasesTREFLE{cards.ValetCarreau, "Valet carreau should be 2", 2},
		evaluateCardTestCasesTREFLE{cards.NeufCarreau, "Neuf carreau should be 0", 0},
		evaluateCardTestCasesTREFLE{cards.AsCarreau, "As carreau should be 11", 11},
		evaluateCardTestCasesTREFLE{cards.DixCarreau, "Dix carreau should be 10", 10},
		evaluateCardTestCasesTREFLE{cards.RoiCarreau, "Roi carreau should be 4", 4},
		evaluateCardTestCasesTREFLE{cards.DameCarreau, "Dame carreau should be 3", 3},
		evaluateCardTestCasesTREFLE{cards.HuitCarreau, "Huit carreau should be 0", 0},
		evaluateCardTestCasesTREFLE{cards.SeptCarreau, "Sept carreau should be 0", 0},

		evaluateCardTestCasesTREFLE{cards.ValetCoeur, "Valet coeur should be 2", 2},
		evaluateCardTestCasesTREFLE{cards.NeufCoeur, "Neuf coeur should be 0", 0},
		evaluateCardTestCasesTREFLE{cards.AsCoeur, "As coeur should be 11", 11},
		evaluateCardTestCasesTREFLE{cards.DixCoeur, "Dix coeur should be 10", 10},
		evaluateCardTestCasesTREFLE{cards.RoiCoeur, "Roi coeur should be 4", 4},
		evaluateCardTestCasesTREFLE{cards.DameCoeur, "Dame coeur should be 3", 3},
		evaluateCardTestCasesTREFLE{cards.HuitCoeur, "Huit coeur should be 0", 0},
		evaluateCardTestCasesTREFLE{cards.SeptCoeur, "Sept coeur should be 0", 0},

		evaluateCardTestCasesTREFLE{cards.ValetTrefle, "Valet trefle should be 20", 20},
		evaluateCardTestCasesTREFLE{cards.NeufTrefle, "Neuf trefle should be 14", 14},
		evaluateCardTestCasesTREFLE{cards.AsTrefle, "As trefle should be 11", 11},
		evaluateCardTestCasesTREFLE{cards.DixTrefle, "Dix trefle should be 10", 10},
		evaluateCardTestCasesTREFLE{cards.RoiTrefle, "Roi trefle should be 4", 4},
		evaluateCardTestCasesTREFLE{cards.DameTrefle, "Dame trefle should be 3", 3},
		evaluateCardTestCasesTREFLE{cards.HuitTrefle, "Huit trefle should be 0", 0},
		evaluateCardTestCasesTREFLE{cards.SeptTrefle, "Sept trefle should be 0", 0},

		evaluateCardTestCasesTREFLE{cards.ValetPique, "Valet pique should be 2", 2},
		evaluateCardTestCasesTREFLE{cards.NeufPique, "Neuf pique should be 0", 0},
		evaluateCardTestCasesTREFLE{cards.AsPique, "As pique should be 11", 11},
		evaluateCardTestCasesTREFLE{cards.DixPique, "Dix pique should be 10", 10},
		evaluateCardTestCasesTREFLE{cards.RoiPique, "Roi pique should be 4", 4},
		evaluateCardTestCasesTREFLE{cards.DamePique, "Dame pique should be 3", 3},
		evaluateCardTestCasesTREFLE{cards.HuitPique, "Huit pique should be 0", 0},
		evaluateCardTestCasesTREFLE{cards.SeptPique, "Sept pique should be 0", 0},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			trefle := Trefle{}
			result, _ := trefle.EvaluateCard(test.card)
			assert.Equal(t, result, test.value)
		})
	}
}

type testGreaterThanCasesTREFLE struct {
	name     string
	take     GameTake
	expected bool
}

func TestGreaterThanTREFLE(t *testing.T) {
	tc := []testGreaterThanCasesTREFLE{
		testGreaterThanCasesTREFLE{"Trefle is greater than tout", TOUT, false},
		testGreaterThanCasesTREFLE{"Trefle is greater than cent", CENT, false},
		testGreaterThanCasesTREFLE{"Trefle is greater than pique", PIQUE, false},
		testGreaterThanCasesTREFLE{"Trefle is greater than pique", COEUR, false},
		testGreaterThanCasesTREFLE{"Trefle is greater than carreau", CARREAU, false},
		testGreaterThanCasesTREFLE{"Trefle is greater than passe", PASSE, true},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			result := TREFLE.GreaterThan(test.take)
			assert.Equal(t, result, test.expected)
		})
	}
}

func TestGetValueTREFLE(t *testing.T) {
	result := TREFLE.GetValue()
	assert.Equal(t, result, 1)
}

func TestNameTREFLE(t *testing.T) {
	result := TREFLE.Name()
	assert.Equal(t, result, "Trefle")
}
