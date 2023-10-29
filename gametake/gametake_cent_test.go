package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/pkg/cards"
)

type evaluateCardTestCasesCENT struct {
	card  cards.Card
	name  string
	value int
}

var TestEvaluateCardCENTTable []evaluateCardTestCasesCENT = []evaluateCardTestCasesCENT{
	evaluateCardTestCasesCENT{cards.ValetCarreau, "Valet carreau should be 2", 2},
	evaluateCardTestCasesCENT{cards.NeufCarreau, "Neuf carreau should be 0", 0},
	evaluateCardTestCasesCENT{cards.AsCarreau, "As carreau should be 11", 11},
	evaluateCardTestCasesCENT{cards.DixCarreau, "Dix carreau should be 10", 10},
	evaluateCardTestCasesCENT{cards.RoiCarreau, "Roi carreau should be 4", 4},
	evaluateCardTestCasesCENT{cards.DameCarreau, "Dame carreau should be 3", 3},
	evaluateCardTestCasesCENT{cards.HuitCarreau, "Huit carreau should be 0", 0},
	evaluateCardTestCasesCENT{cards.SeptCarreau, "Sept carreau should be 0", 0},
	evaluateCardTestCasesCENT{cards.ValetCoeur, "Valet coeur should be 2", 2},
	evaluateCardTestCasesCENT{cards.NeufCoeur, "Neuf coeur should be 0", 0},
	evaluateCardTestCasesCENT{cards.AsCoeur, "As coeur should be 11", 11},
	evaluateCardTestCasesCENT{cards.DixCoeur, "Dix coeur should be 10", 10},
	evaluateCardTestCasesCENT{cards.RoiCoeur, "Roi coeur should be 4", 4},
	evaluateCardTestCasesCENT{cards.DameCoeur, "Dame coeur should be 3", 3},
	evaluateCardTestCasesCENT{cards.HuitCoeur, "Huit coeur should be 0", 0},
	evaluateCardTestCasesCENT{cards.SeptCoeur, "Sept coeur should be 0", 0},
	evaluateCardTestCasesCENT{cards.ValetTrefle, "Valet trefle should be 2", 2},
	evaluateCardTestCasesCENT{cards.NeufTrefle, "Neuf trefle should be 0", 0},
	evaluateCardTestCasesCENT{cards.AsTrefle, "As trefle should be 11", 11},
	evaluateCardTestCasesCENT{cards.DixTrefle, "Dix trefle should be 10", 10},
	evaluateCardTestCasesCENT{cards.RoiTrefle, "Roi trefle should be 4", 4},
	evaluateCardTestCasesCENT{cards.DameTrefle, "Dame trefle should be 3", 3},
	evaluateCardTestCasesCENT{cards.HuitTrefle, "Huit trefle should be 0", 0},
	evaluateCardTestCasesCENT{cards.SeptTrefle, "Sept trefle should be 0", 0},
	evaluateCardTestCasesCENT{cards.ValetPique, "Valet pique should be 2", 2},
	evaluateCardTestCasesCENT{cards.NeufPique, "Neuf pique should be 0", 0},
	evaluateCardTestCasesCENT{cards.AsPique, "As pique should be 11", 11},
	evaluateCardTestCasesCENT{cards.DixPique, "Dix pique should be 10", 10},
	evaluateCardTestCasesCENT{cards.RoiPique, "Roi pique should be 4", 4},
	evaluateCardTestCasesCENT{cards.DamePique, "Dame pique should be 3", 3},
	evaluateCardTestCasesCENT{cards.HuitPique, "Huit pique should be 0", 0},
	evaluateCardTestCasesCENT{cards.SeptPique, "Sept pique should be 0", 0},
}

type testCENTGreaterThanCases struct {
	name     string
	take     GameTake
	expected bool
}

func TestCENTGreaterThan(t *testing.T) {
	tc := []testCENTGreaterThanCases{
		testCENTGreaterThanCases{"Cent is greater than cent", CENT, false},
		testCENTGreaterThanCases{"Cent is greater than cent", PIQUE, true},
		testCENTGreaterThanCases{"Cent is greater than cent", COEUR, true},
		testCENTGreaterThanCases{"Cent is greater than cent", CARREAU, true},
		testCENTGreaterThanCases{"Cent is greater than cent", TREFLE, true},
		testCENTGreaterThanCases{"Cent is greater than cent", PASSE, true},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			result := CENT.GreaterThan(test.take)
			assert.Equal(t, result, test.expected)
		})
	}
}

func TestCENTGetValue(t *testing.T) {
	result := CENT.GetValue()
	assert.Equal(t, result, 5)
}

func TestCENTName(t *testing.T) {
	result := CENT.Name()
	assert.Equal(t, result, "Cent")
}

func TestEvaluateCardCENT(t *testing.T) {
	for _, test := range TestEvaluateCardCENTTable {
		cent := Cent{}
		t.Run(test.name, func(t *testing.T) {
			result, ok := cent.EvaluateCard(test.card)
			assert.Equal(t, result, test.value)
			assert.Equal(t, ok, true)
		})
	}
}

func TestCENTEvaluateDeck(t *testing.T) {
	type TestCENTEvaluateDeckTable struct {
		name   string
		deck   [4]cards.Card
		result int
	}
	testcases := []TestCENTEvaluateDeckTable{
		TestCENTEvaluateDeckTable{
			name:   "With no cards",
			deck:   [4]cards.Card{},
			result: 0,
		},
		TestCENTEvaluateDeckTable{
			name:   "With cards all worth zero points",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.NeufCoeur, cards.HuitCarreau, cards.NeufTrefle},
			result: 0,
		},
		TestCENTEvaluateDeckTable{
			name:   "With one ace and three tens",
			deck:   [4]cards.Card{cards.AsCarreau, cards.DixCarreau, cards.DixPique, cards.DixCoeur},
			result: 41,
		},
		TestCENTEvaluateDeckTable{
			name:   "With one ten two kings and a seven",
			deck:   [4]cards.Card{cards.DixCarreau, cards.RoiCarreau, cards.RoiCoeur, cards.SeptCarreau},
			result: 18,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			result := CENT.EvaluateDeck(test.deck)
			assert.Equal(t, result, test.result)
		})
	}
}
