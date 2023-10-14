package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/pkg/cards"
)

type evaluateCardTestCases struct {
	card  cards.Card
	name  string
	value int
}

func TestEvaluateCard(t *testing.T) {
	tc := []evaluateCardTestCases{
		evaluateCardTestCases{cards.ValetCarreau, "Valet carreau should be 20", 20},
		evaluateCardTestCases{cards.NeufCarreau, "Neuf carreau should be 14", 14},
		evaluateCardTestCases{cards.AsCarreau, "As carreau should be 11", 11},
		evaluateCardTestCases{cards.DixCarreau, "Dix carreau should be 10", 10},
		evaluateCardTestCases{cards.RoiCarreau, "Roi carreau should be 4", 4},
		evaluateCardTestCases{cards.DameCarreau, "Dame carreau should be 3", 3},
		evaluateCardTestCases{cards.HuitCarreau, "Huit carreau should be 0", 0},
		evaluateCardTestCases{cards.SeptCarreau, "Sept carreau should be 0", 0},

		evaluateCardTestCases{cards.ValetCoeur, "Valet coeur should be 20", 20},
		evaluateCardTestCases{cards.NeufCoeur, "Neuf coeur should be 14", 14},
		evaluateCardTestCases{cards.AsCoeur, "As coeur should be 11", 11},
		evaluateCardTestCases{cards.DixCoeur, "Dix coeur should be 10", 10},
		evaluateCardTestCases{cards.RoiCoeur, "Roi coeur should be 4", 4},
		evaluateCardTestCases{cards.DameCoeur, "Dame coeur should be 3", 3},
		evaluateCardTestCases{cards.HuitCoeur, "Huit coeur should be 0", 0},
		evaluateCardTestCases{cards.SeptCoeur, "Sept coeur should be 0", 0},

		evaluateCardTestCases{cards.ValetTrefle, "Valet trefle should be 20", 20},
		evaluateCardTestCases{cards.NeufTrefle, "Neuf trefle should be 14", 14},
		evaluateCardTestCases{cards.AsTrefle, "As trefle should be 11", 11},
		evaluateCardTestCases{cards.DixTrefle, "Dix trefle should be 10", 10},
		evaluateCardTestCases{cards.RoiTrefle, "Roi trefle should be 4", 4},
		evaluateCardTestCases{cards.DameTrefle, "Dame trefle should be 3", 3},
		evaluateCardTestCases{cards.HuitTrefle, "Huit trefle should be 0", 0},
		evaluateCardTestCases{cards.SeptTrefle, "Sept trefle should be 0", 0},

		evaluateCardTestCases{cards.ValetPique, "Valet pique should be 20", 20},
		evaluateCardTestCases{cards.NeufPique, "Neuf pique should be 14", 14},
		evaluateCardTestCases{cards.AsPique, "As pique should be 11", 11},
		evaluateCardTestCases{cards.DixPique, "Dix pique should be 10", 10},
		evaluateCardTestCases{cards.RoiPique, "Roi pique should be 4", 4},
		evaluateCardTestCases{cards.DamePique, "Dame pique should be 3", 3},
		evaluateCardTestCases{cards.HuitPique, "Huit pique should be 0", 0},
		evaluateCardTestCases{cards.SeptPique, "Sept pique should be 0", 0},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			tout := Tout{}
			result, _ := tout.EvaluateCard(test.card)
			assert.Equal(t, result, test.value)
		})
	}
}

func TestTOUTGetValue(t *testing.T) {
	tout := TOUT
	result := tout.GetValue()
	assert.Equal(t, result, 6)
}

func TestTOUTName(t *testing.T) {
	tout := TOUT
	result := tout.Name()
	assert.Equal(t, result, "Tout")
}

type testTOUTGreaterThanCases struct {
	name     string
	take     GameTake
	expected bool
}

func TestTOUTGreaterThan(t *testing.T) {
	tc := []testTOUTGreaterThanCases{
		testTOUTGreaterThanCases{"Tout is greater than cent", CENT, true},
		testTOUTGreaterThanCases{"Tout is greater than cent", PIQUE, true},
		testTOUTGreaterThanCases{"Tout is greater than cent", COEUR, true},
		testTOUTGreaterThanCases{"Tout is greater than cent", CARREAU, true},
		testTOUTGreaterThanCases{"Tout is greater than cent", TREFLE, true},
		testTOUTGreaterThanCases{"Tout is greater than cent", PASSE, true},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			result := TOUT.GreaterThan(test.take)
			assert.Equal(t, result, test.expected)
		})
	}
}
