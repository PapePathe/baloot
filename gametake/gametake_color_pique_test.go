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
	t.Parallel()

	tc := []evaluateCardTestCasesPIQUE{
		{cards.ValetCarreau, "Valet carreau should be 2", 2},
		{cards.NeufCarreau, "Neuf carreau should be 0", 0},
		{cards.AsCarreau, "As carreau should be 11", 11},
		{cards.DixCarreau, "Dix carreau should be 10", 10},
		{cards.RoiCarreau, "Roi carreau should be 4", 4},
		{cards.DameCarreau, "Dame carreau should be 3", 3},
		{cards.HuitCarreau, "Huit carreau should be 0", 0},
		{cards.SeptCarreau, "Sept carreau should be 0", 0},

		{cards.ValetCoeur, "Valet coeur should be 2", 2},
		{cards.NeufCoeur, "Neuf coeur should be 0", 0},
		{cards.AsCoeur, "As coeur should be 11", 11},
		{cards.DixCoeur, "Dix coeur should be 10", 10},
		{cards.RoiCoeur, "Roi coeur should be 4", 4},
		{cards.DameCoeur, "Dame coeur should be 3", 3},
		{cards.HuitCoeur, "Huit coeur should be 0", 0},
		{cards.SeptCoeur, "Sept coeur should be 0", 0},

		{cards.ValetTrefle, "Valet trefle should be 2", 2},
		{cards.NeufTrefle, "Neuf trefle should be 0", 0},
		{cards.AsTrefle, "As trefle should be 11", 11},
		{cards.DixTrefle, "Dix trefle should be 10", 10},
		{cards.RoiTrefle, "Roi trefle should be 4", 4},
		{cards.DameTrefle, "Dame trefle should be 3", 3},
		{cards.HuitTrefle, "Huit trefle should be 0", 0},
		{cards.SeptTrefle, "Sept trefle should be 0", 0},

		{cards.ValetPique, "Valet pique should be 2", 20},
		{cards.NeufPique, "Neuf pique should be 14", 14},
		{cards.AsPique, "As pique should be 11", 11},
		{cards.DixPique, "Dix pique should be 10", 10},
		{cards.RoiPique, "Roi pique should be 4", 4},
		{cards.DamePique, "Dame pique should be 3", 3},
		{cards.HuitPique, "Huit pique should be 0", 0},
		{cards.SeptPique, "Sept pique should be 0", 0},
	}

	for _, test := range tc {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
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
	t.Parallel()

	tc := []testGreaterThanCasesPIQUE{
		{"Coeur is greater than tout", TOUT, false},
		{"Coeur is greater than cent", CENT, false},
		{"Coeur is greater than pique", COEUR, true},
		{"Coeur is greater than carreau", CARREAU, true},
		{"Coeur is greater than trefle", TREFLE, true},
		{"Coeur is greater than passe", PASSE, true},
	}

	for _, test := range tc {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := PIQUE.GreaterThan(test.take)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestGetValuePIQUE(t *testing.T) {
	t.Parallel()

	result := PIQUE.GetValue()
	assert.Equal(t, 4, result)
}

func TestNamePIQUE(t *testing.T) {
	t.Parallel()

	result := PIQUE.Name()
	assert.Equal(t, "Pique", result)
}

func TestEvaluateDeckPIQUE(t *testing.T) {
	t.Parallel()

	type TestEvaluateDeckTablePIQUE struct {
		name   string
		deck   [4]cards.Card
		result int
	}

	testcases := []TestEvaluateDeckTablePIQUE{
		{
			name:   "With no cards",
			deck:   [4]cards.Card{},
			result: 0,
		},
		{
			name:   "With a valet of other color and zero value cards",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.ValetCarreau, cards.HuitCarreau, cards.SeptTrefle},
			result: 2,
		},
		{
			name:   "With a valet of same color and zero value cards",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.ValetPique, cards.HuitCarreau, cards.SeptTrefle},
			result: 20,
		},
		{
			name:   "With two nines of other color an height and a seven",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.NeufCoeur, cards.HuitCarreau, cards.NeufCarreau},
			result: 0,
		},
		{
			name:   "With one nine of same color and a nine of other color an height and a seven",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.NeufPique, cards.HuitCarreau, cards.NeufCarreau},
			result: 14,
		},
		{
			name:   "With one ace and three tens",
			deck:   [4]cards.Card{cards.AsCarreau, cards.DixCarreau, cards.DixPique, cards.DixCoeur},
			result: 41,
		},
		{
			name:   "With one ten two kings and a seven",
			deck:   [4]cards.Card{cards.DixCarreau, cards.RoiCarreau, cards.RoiCoeur, cards.SeptCarreau},
			result: 18,
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := PIQUE.EvaluateDeck(test.deck)
			assert.Equal(t, result, test.result)
		})
	}
}
