package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/pkg/cards"
)

type evaluateCardTestCasesCOEUR struct {
	card  cards.Card
	name  string
	value int
}

func TestEvaluateCardCOEUR(t *testing.T) {
	tc := []evaluateCardTestCasesCOEUR{
		evaluateCardTestCasesCOEUR{cards.ValetCarreau, "Valet carreau should be 2", 2},
		evaluateCardTestCasesCOEUR{cards.NeufCarreau, "Neuf carreau should be 0", 0},
		evaluateCardTestCasesCOEUR{cards.AsCarreau, "As carreau should be 11", 11},
		evaluateCardTestCasesCOEUR{cards.DixCarreau, "Dix carreau should be 10", 10},
		evaluateCardTestCasesCOEUR{cards.RoiCarreau, "Roi carreau should be 4", 4},
		evaluateCardTestCasesCOEUR{cards.DameCarreau, "Dame carreau should be 3", 3},
		evaluateCardTestCasesCOEUR{cards.HuitCarreau, "Huit carreau should be 0", 0},
		evaluateCardTestCasesCOEUR{cards.SeptCarreau, "Sept carreau should be 0", 0},

		evaluateCardTestCasesCOEUR{cards.ValetCoeur, "Valet coeur should be 20", 20},
		evaluateCardTestCasesCOEUR{cards.NeufCoeur, "Neuf coeur should be 14", 14},
		evaluateCardTestCasesCOEUR{cards.AsCoeur, "As coeur should be 11", 11},
		evaluateCardTestCasesCOEUR{cards.DixCoeur, "Dix coeur should be 10", 10},
		evaluateCardTestCasesCOEUR{cards.RoiCoeur, "Roi coeur should be 4", 4},
		evaluateCardTestCasesCOEUR{cards.DameCoeur, "Dame coeur should be 3", 3},
		evaluateCardTestCasesCOEUR{cards.HuitCoeur, "Huit coeur should be 0", 0},
		evaluateCardTestCasesCOEUR{cards.SeptCoeur, "Sept coeur should be 0", 0},

		evaluateCardTestCasesCOEUR{cards.ValetTrefle, "Valet trefle should be 2", 2},
		evaluateCardTestCasesCOEUR{cards.NeufTrefle, "Neuf trefle should be 0", 0},
		evaluateCardTestCasesCOEUR{cards.AsTrefle, "As trefle should be 11", 11},
		evaluateCardTestCasesCOEUR{cards.DixTrefle, "Dix trefle should be 10", 10},
		evaluateCardTestCasesCOEUR{cards.RoiTrefle, "Roi trefle should be 4", 4},
		evaluateCardTestCasesCOEUR{cards.DameTrefle, "Dame trefle should be 3", 3},
		evaluateCardTestCasesCOEUR{cards.HuitTrefle, "Huit trefle should be 0", 0},
		evaluateCardTestCasesCOEUR{cards.SeptTrefle, "Sept trefle should be 0", 0},

		evaluateCardTestCasesCOEUR{cards.ValetPique, "Valet pique should be 2", 2},
		evaluateCardTestCasesCOEUR{cards.NeufPique, "Neuf pique should be 0", 0},
		evaluateCardTestCasesCOEUR{cards.AsPique, "As pique should be 11", 11},
		evaluateCardTestCasesCOEUR{cards.DixPique, "Dix pique should be 10", 10},
		evaluateCardTestCasesCOEUR{cards.RoiPique, "Roi pique should be 4", 4},
		evaluateCardTestCasesCOEUR{cards.DamePique, "Dame pique should be 3", 3},
		evaluateCardTestCasesCOEUR{cards.HuitPique, "Huit pique should be 0", 0},
		evaluateCardTestCasesCOEUR{cards.SeptPique, "Sept pique should be 0", 0},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			tout := COEUR
			result, _ := tout.EvaluateCard(test.card)
			assert.Equal(t, result, test.value)
		})
	}
}

type testGreaterThanCasesCOEUR struct {
	name     string
	take     GameTake
	expected bool
}

func TestGreaterThanCOEUR(t *testing.T) {
	tc := []testGreaterThanCasesCOEUR{
		testGreaterThanCasesCOEUR{"Coeur is greater than tout", TOUT, false},
		testGreaterThanCasesCOEUR{"Coeur is greater than cent", CENT, false},
		testGreaterThanCasesCOEUR{"Coeur is greater than pique", PIQUE, false},
		testGreaterThanCasesCOEUR{"Coeur is greater than carreau", CARREAU, true},
		testGreaterThanCasesCOEUR{"Coeur is greater than trefle", TREFLE, true},
		testGreaterThanCasesCOEUR{"Coeur is greater than passe", PASSE, true},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			result := COEUR.GreaterThan(test.take)
			assert.Equal(t, result, test.expected)
		})
	}
}

func TestGetValueCOEUR(t *testing.T) {
	result := COEUR.GetValue()
	assert.Equal(t, result, 3)
}

func TestNameCOEUR(t *testing.T) {
	result := COEUR.Name()
	assert.Equal(t, result, "Coeur")
}

func TestEvaluateDeckCOEUR(t *testing.T) {
	type TestEvaluateDeckTableCOEUR struct {
		name   string
		deck   [4]cards.Card
		result int
	}
	testcases := []TestEvaluateDeckTableCOEUR{
		TestEvaluateDeckTableCOEUR{
			name:   "With no cards",
			deck:   [4]cards.Card{},
			result: 0,
		},
		TestEvaluateDeckTableCOEUR{
			name:   "With a valet of other color and zero value cards",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.ValetCarreau, cards.HuitCarreau, cards.SeptTrefle},
			result: 2,
		},
		TestEvaluateDeckTableCOEUR{
			name:   "With a valet of same color and zero value cards",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.ValetCoeur, cards.HuitCarreau, cards.SeptTrefle},
			result: 20,
		},
		TestEvaluateDeckTableCOEUR{
			name:   "With two nines of other color an height and a seven",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.NeufPique, cards.HuitCarreau, cards.NeufTrefle},
			result: 0,
		},
		TestEvaluateDeckTableCOEUR{
			name:   "With one nine of same color and a nine of other color an height and a seven",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.NeufCoeur, cards.HuitCarreau, cards.NeufTrefle},
			result: 14,
		},
		TestEvaluateDeckTableCOEUR{
			name:   "With one ace and three tens",
			deck:   [4]cards.Card{cards.AsCarreau, cards.DixCarreau, cards.DixPique, cards.DixCoeur},
			result: 41,
		},
		TestEvaluateDeckTableCOEUR{
			name:   "With one ten two kings and a seven",
			deck:   [4]cards.Card{cards.DixCarreau, cards.RoiCarreau, cards.RoiCoeur, cards.SeptCarreau},
			result: 18,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			result := COEUR.EvaluateDeck(test.deck)
			assert.Equal(t, result, test.result)
		})
	}
}
