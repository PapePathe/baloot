package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pathe.co/zinx/pkg/cards"
)

type evaluateCardTestCases struct {
	card  cards.Card
	name  string
	value int
}

func TestEvaluateCard(t *testing.T) {
	t.Parallel()

	tc := []evaluateCardTestCases{
		{cards.ValetCarreau, "Valet carreau should be 20", 20},
		{cards.NeufCarreau, "Neuf carreau should be 14", 14},
		{cards.AsCarreau, "As carreau should be 11", 11},
		{cards.DixCarreau, "Dix carreau should be 10", 10},
		{cards.RoiCarreau, "Roi carreau should be 4", 4},
		{cards.DameCarreau, "Dame carreau should be 3", 3},
		{cards.HuitCarreau, "Huit carreau should be 0", 0},
		{cards.SeptCarreau, "Sept carreau should be 0", 0},

		{cards.ValetCoeur, "Valet coeur should be 20", 20},
		{cards.NeufCoeur, "Neuf coeur should be 14", 14},
		{cards.AsCoeur, "As coeur should be 11", 11},
		{cards.DixCoeur, "Dix coeur should be 10", 10},
		{cards.RoiCoeur, "Roi coeur should be 4", 4},
		{cards.DameCoeur, "Dame coeur should be 3", 3},
		{cards.HuitCoeur, "Huit coeur should be 0", 0},
		{cards.SeptCoeur, "Sept coeur should be 0", 0},

		{cards.ValetTrefle, "Valet trefle should be 20", 20},
		{cards.NeufTrefle, "Neuf trefle should be 14", 14},
		{cards.AsTrefle, "As trefle should be 11", 11},
		{cards.DixTrefle, "Dix trefle should be 10", 10},
		{cards.RoiTrefle, "Roi trefle should be 4", 4},
		{cards.DameTrefle, "Dame trefle should be 3", 3},
		{cards.HuitTrefle, "Huit trefle should be 0", 0},
		{cards.SeptTrefle, "Sept trefle should be 0", 0},

		{cards.ValetPique, "Valet pique should be 20", 20},
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
			tout := Tout{AllCardsValue: 0, CardsOfTakeValue: 0}
			result, _ := tout.EvaluateCard(test.card)
			assert.Equal(t, test.value, result)
		})
	}
}

func TestTOUTGetValue(t *testing.T) {
	t.Parallel()

	tout := TOUT
	result := tout.GetValue()
	assert.Equal(t, 6, result)
}

func TestTOUTName(t *testing.T) {
	t.Parallel()

	tout := TOUT
	result := tout.Name()
	assert.Equal(t, "Tout", result)
}

type testTOUTGreaterThanCases struct {
	name     string
	take     GameTake
	expected bool
}

func TestTOUTGreaterThan(t *testing.T) {
	t.Parallel()

	tc := []testTOUTGreaterThanCases{
		{"Tout is greater than cent", CENT, true},
		{"Tout is greater than cent", PIQUE, true},
		{"Tout is greater than cent", COEUR, true},
		{"Tout is greater than cent", CARREAU, true},
		{"Tout is greater than cent", TREFLE, true},
		{"Tout is greater than cent", PASSE, true},
	}

	for _, test := range tc {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := TOUT.GreaterThan(test.take)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestTOUTEvaluateDeck(t *testing.T) {
	t.Parallel()

	type TestTOUTEvaluateDeckTable struct {
		name   string
		deck   [4]cards.Card
		result int
	}

	testcases := []TestTOUTEvaluateDeckTable{
		{
			name:   "With no cards",
			deck:   [4]cards.Card{},
			result: 0,
		},
		{
			name:   "With a valet and zero value cards",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.ValetCoeur, cards.HuitCarreau, cards.SeptTrefle},
			result: 20,
		},
		{
			name:   "With two nines an height and a seven",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.NeufCoeur, cards.HuitCarreau, cards.NeufTrefle},
			result: 28,
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

			result := TOUT.EvaluateDeck(test.deck)
			assert.Equal(t, result, test.result)
		})
	}
}

func TestMarshalJSONTOUT(t *testing.T) {
	t.Parallel()

	tout := Tout{}
	result, err := tout.MarshalJSON()

	require.NoError(t, err)
	assert.Equal(t, "{\"name\":\"Tout\"}", string(result))
}

func TestWinnerTOUT(t *testing.T) {
	t.Parallel()

	winner := TOUT.Winner(cards.ValetCoeur, cards.DixCoeur)
	assert.Equal(t, cards.ValetCoeur, winner)

	winner = TOUT.Winner(cards.AsCoeur, cards.DixCoeur)
	assert.Equal(t, cards.AsCoeur, winner)

	winner = TOUT.Winner(cards.HuitCoeur, cards.SeptCoeur)
	assert.Equal(t, cards.HuitCoeur, winner)

	winner = TOUT.Winner(cards.SeptCoeur, cards.HuitCoeur)
	assert.Equal(t, cards.HuitCoeur, winner)

	winner = TOUT.Winner(cards.RoiPique, cards.HuitCoeur)
	assert.Equal(t, cards.RoiPique, winner)

	winner = TOUT.Winner(cards.RoiPique, cards.DixCarreau)
	assert.Equal(t, cards.RoiPique, winner)

	winner = TOUT.Winner(cards.RoiPique, cards.DamePique)
	assert.Equal(t, cards.RoiPique, winner)

	winner = TOUT.Winner(cards.NeufCarreau, cards.ValetCarreau)
	assert.Equal(t, cards.ValetCarreau, winner)

	winner = TOUT.Winner(cards.ValetTrefle, cards.NeufCarreau)
	assert.Equal(t, cards.ValetTrefle, winner)

	winner = TOUT.Winner(cards.NeufCarreau, cards.ValetTrefle)
	assert.Equal(t, cards.NeufCarreau, winner)
}
