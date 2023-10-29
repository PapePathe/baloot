package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/pkg/cards"
)

type evaluateCardTestCasesPIQUE struct {
	card  cards.Card
	name  string
	value int
}

func TestEvaluateCardPIQUE(t *testing.T) {
	tc := []evaluateCardTestCasesPIQUE{
		evaluateCardTestCasesPIQUE{cards.ValetCarreau, "Valet carreau should be 2", 2},
		evaluateCardTestCasesPIQUE{cards.NeufCarreau, "Neuf carreau should be 0", 0},
		evaluateCardTestCasesPIQUE{cards.AsCarreau, "As carreau should be 11", 11},
		evaluateCardTestCasesPIQUE{cards.DixCarreau, "Dix carreau should be 10", 10},
		evaluateCardTestCasesPIQUE{cards.RoiCarreau, "Roi carreau should be 4", 4},
		evaluateCardTestCasesPIQUE{cards.DameCarreau, "Dame carreau should be 3", 3},
		evaluateCardTestCasesPIQUE{cards.HuitCarreau, "Huit carreau should be 0", 0},
		evaluateCardTestCasesPIQUE{cards.SeptCarreau, "Sept carreau should be 0", 0},

		evaluateCardTestCasesPIQUE{cards.ValetCoeur, "Valet coeur should be 2", 2},
		evaluateCardTestCasesPIQUE{cards.NeufCoeur, "Neuf coeur should be 0", 0},
		evaluateCardTestCasesPIQUE{cards.AsCoeur, "As coeur should be 11", 11},
		evaluateCardTestCasesPIQUE{cards.DixCoeur, "Dix coeur should be 10", 10},
		evaluateCardTestCasesPIQUE{cards.RoiCoeur, "Roi coeur should be 4", 4},
		evaluateCardTestCasesPIQUE{cards.DameCoeur, "Dame coeur should be 3", 3},
		evaluateCardTestCasesPIQUE{cards.HuitCoeur, "Huit coeur should be 0", 0},
		evaluateCardTestCasesPIQUE{cards.SeptCoeur, "Sept coeur should be 0", 0},

		evaluateCardTestCasesPIQUE{cards.ValetTrefle, "Valet trefle should be 2", 2},
		evaluateCardTestCasesPIQUE{cards.NeufTrefle, "Neuf trefle should be 0", 0},
		evaluateCardTestCasesPIQUE{cards.AsTrefle, "As trefle should be 11", 11},
		evaluateCardTestCasesPIQUE{cards.DixTrefle, "Dix trefle should be 10", 10},
		evaluateCardTestCasesPIQUE{cards.RoiTrefle, "Roi trefle should be 4", 4},
		evaluateCardTestCasesPIQUE{cards.DameTrefle, "Dame trefle should be 3", 3},
		evaluateCardTestCasesPIQUE{cards.HuitTrefle, "Huit trefle should be 0", 0},
		evaluateCardTestCasesPIQUE{cards.SeptTrefle, "Sept trefle should be 0", 0},

		evaluateCardTestCasesPIQUE{cards.ValetPique, "Valet pique should be 2", 20},
		evaluateCardTestCasesPIQUE{cards.NeufPique, "Neuf pique should be 14", 14},
		evaluateCardTestCasesPIQUE{cards.AsPique, "As pique should be 11", 11},
		evaluateCardTestCasesPIQUE{cards.DixPique, "Dix pique should be 10", 10},
		evaluateCardTestCasesPIQUE{cards.RoiPique, "Roi pique should be 4", 4},
		evaluateCardTestCasesPIQUE{cards.DamePique, "Dame pique should be 3", 3},
		evaluateCardTestCasesPIQUE{cards.HuitPique, "Huit pique should be 0", 0},
		evaluateCardTestCasesPIQUE{cards.SeptPique, "Sept pique should be 0", 0},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			pique := PIQUE
			result, _ := pique.EvaluateCard(test.card)
			assert.Equal(t, result, test.value)
		})
	}
}

type testGreaterThanCasesPIQUE struct {
	name     string
	take     GameTake
	expected bool
}

func TestGreaterThanPIQUE(t *testing.T) {
	tc := []testGreaterThanCasesPIQUE{
		testGreaterThanCasesPIQUE{"Coeur is greater than tout", TOUT, false},
		testGreaterThanCasesPIQUE{"Coeur is greater than cent", CENT, false},
		testGreaterThanCasesPIQUE{"Coeur is greater than pique", COEUR, true},
		testGreaterThanCasesPIQUE{"Coeur is greater than carreau", CARREAU, true},
		testGreaterThanCasesPIQUE{"Coeur is greater than trefle", TREFLE, true},
		testGreaterThanCasesPIQUE{"Coeur is greater than passe", PASSE, true},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			result := PIQUE.GreaterThan(test.take)
			assert.Equal(t, result, test.expected)
		})
	}
}

func TestGetValuePIQUE(t *testing.T) {
	result := PIQUE.GetValue()
	assert.Equal(t, result, 4)
}

func TestNamePIQUE(t *testing.T) {
	result := PIQUE.Name()
	assert.Equal(t, result, "Pique")
}

func TestEvaluateDeckPIQUE(t *testing.T) {
	type TestEvaluateDeckTablePIQUE struct {
		name   string
		deck   [4]cards.Card
		result int
	}
	testcases := []TestEvaluateDeckTablePIQUE{
		TestEvaluateDeckTablePIQUE{
			name:   "With no cards",
			deck:   [4]cards.Card{},
			result: 0,
		},
		TestEvaluateDeckTablePIQUE{
			name:   "With a valet of other color and zero value cards",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.ValetCarreau, cards.HuitCarreau, cards.SeptTrefle},
			result: 2,
		},
		TestEvaluateDeckTablePIQUE{
			name:   "With a valet of same color and zero value cards",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.ValetPique, cards.HuitCarreau, cards.SeptTrefle},
			result: 20,
		},
		TestEvaluateDeckTablePIQUE{
			name:   "With two nines of other color an height and a seven",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.NeufCoeur, cards.HuitCarreau, cards.NeufCarreau},
			result: 0,
		},
		TestEvaluateDeckTablePIQUE{
			name:   "With one nine of same color and a nine of other color an height and a seven",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.NeufPique, cards.HuitCarreau, cards.NeufCarreau},
			result: 14,
		},
		TestEvaluateDeckTablePIQUE{
			name:   "With one ace and three tens",
			deck:   [4]cards.Card{cards.AsCarreau, cards.DixCarreau, cards.DixPique, cards.DixCoeur},
			result: 41,
		},
		TestEvaluateDeckTablePIQUE{
			name:   "With one ten two kings and a seven",
			deck:   [4]cards.Card{cards.DixCarreau, cards.RoiCarreau, cards.RoiCoeur, cards.SeptCarreau},
			result: 18,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			result := PIQUE.EvaluateDeck(test.deck)
			assert.Equal(t, result, test.result)
		})
	}
}
