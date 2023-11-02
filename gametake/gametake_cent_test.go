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

var TestEvaluateCardCENTTable = []evaluateCardTestCasesCENT{
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
	{cards.ValetPique, "Valet pique should be 2", 2},
	{cards.NeufPique, "Neuf pique should be 0", 0},
	{cards.AsPique, "As pique should be 11", 11},
	{cards.DixPique, "Dix pique should be 10", 10},
	{cards.RoiPique, "Roi pique should be 4", 4},
	{cards.DamePique, "Dame pique should be 3", 3},
	{cards.HuitPique, "Huit pique should be 0", 0},
	{cards.SeptPique, "Sept pique should be 0", 0},
}

type testCENTGreaterThanCases struct {
	name     string
	take     GameTake
	expected bool
}

func TestCENTGreaterThan(t *testing.T) {
	t.Parallel()

	tc := []testCENTGreaterThanCases{
		{"Cent is greater than cent", CENT, false},
		{"Cent is greater than cent", PIQUE, true},
		{"Cent is greater than cent", COEUR, true},
		{"Cent is greater than cent", CARREAU, true},
		{"Cent is greater than cent", TREFLE, true},
		{"Cent is greater than cent", PASSE, true},
	}
	for _, test := range tc {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := CENT.GreaterThan(test.take)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestCENTGetValue(t *testing.T) {
	t.Parallel()

	result := CENT.GetValue()
	assert.Equal(t, 5, result)
}

func TestCENTName(t *testing.T) {
	t.Parallel()

	result := CENT.Name()
	assert.Equal(t, "Cent", result)
}

func TestEvaluateCardCENT(t *testing.T) {
	t.Parallel()

	for _, test := range TestEvaluateCardCENTTable {
		cent := Cent{AllCardsValue: 0}
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, ok := cent.EvaluateCard(test.card)
			assert.Equal(t, result, test.value)
			assert.True(t, ok)
		})
	}
}

func TestCENTEvaluateDeck(t *testing.T) {
	t.Parallel()

	type TestCENTEvaluateDeckTable struct {
		name   string
		deck   [4]cards.Card
		result int
	}

	testcases := []TestCENTEvaluateDeckTable{
		{
			name:   "With no cards",
			deck:   [4]cards.Card{},
			result: 0,
		},
		{
			name:   "With cards all worth zero points",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.NeufCoeur, cards.HuitCarreau, cards.NeufTrefle},
			result: 0,
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

			result := CENT.EvaluateDeck(test.deck)
			assert.Equal(t, result, test.result)
		})
	}
}
