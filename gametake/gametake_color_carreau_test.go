package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pathe.co/zinx/pkg/cards"
)

func TestGetValue(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 2, CARREAU.GetValue())
}

func TestGetName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "Carreau", CARREAU.Name())
}

type evalCardsTestCase struct {
	name                   string
	cards                  [5]cards.Card
	allCardsValue          int
	cardsOfTakeValue       int
	allPlayerCardsValue    int
	playerCardsOfTakeValue int
}

var TestEvaluateHandTable = []evalCardsTestCase{
	{
		"valet et quatorzaine de carreau et trois sept",
		[5]cards.Card{cards.ValetCarreau, cards.NeufCarreau, cards.SeptCarreau, cards.SeptCoeur, cards.SeptPique},
		152,
		62,
		34,
		34,
	},
	{
		"valet quatorzaine et as de carreau et trois sept",
		[5]cards.Card{cards.ValetCarreau, cards.NeufCarreau, cards.AsCarreau, cards.SeptCoeur, cards.SeptPique},
		152,
		62,
		45,
		45,
	},
	{
		"valet quatorzaine et as de pique et trois sept",
		[5]cards.Card{cards.ValetPique, cards.NeufPique, cards.AsPique, cards.SeptCoeur, cards.SeptPique},
		152,
		62,
		13,
		0,
	},
}

func TestEvaluateHand(t *testing.T) {
	t.Parallel()

	for _, c := range TestEvaluateHandTable {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := CARREAU.EvaluateHand(c.cards)
			assert.Equal(t, c.allCardsValue, result.AllCardsValue)
			assert.Equal(t, c.cardsOfTakeValue, result.CardsOfTakeValue)
			assert.Equal(t, c.allPlayerCardsValue, result.AllPlayerCardsValue)
		})
	}
}

type evaluateCardTestCasesCARREAU struct {
	card  cards.Card
	name  string
	value int
}

func TestEvaluateCardCARREAU(t *testing.T) {
	t.Parallel()

	tc := []evaluateCardTestCasesCARREAU{
		{cards.ValetCarreau, "Valet carreau should be 20", 20},
		{cards.NeufCarreau, "Neuf carreau should be 14", 14},
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

	for _, test := range tc {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			carreau := CARREAU
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
	t.Parallel()

	tc := []testGreaterThanCasesCARREAU{
		{"Carreau is greater than tout", TOUT, false},
		{"Carreau is greater than cent", CENT, false},
		{"Carreau is greater than pique", PIQUE, false},
		{"Carreau is greater than coeur", COEUR, false},
		{"Carreau is greater than trefle", TREFLE, true},
		{"Carreau is greater than passe", PASSE, true},
	}

	for _, test := range tc {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := CARREAU.GreaterThan(test.take)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestEvaluateDeckCARREAU(t *testing.T) {
	t.Parallel()

	type TestEvaluateDeckTableCARREAU struct {
		name   string
		deck   [4]cards.Card
		result int
	}

	testcases := []TestEvaluateDeckTableCARREAU{
		{
			name:   "With no cards",
			deck:   [4]cards.Card{},
			result: 0,
		},
		{
			name:   "With a valet of other color and zero value cards",
			deck:   [4]cards.Card{cards.SeptCarreau, cards.ValetCoeur, cards.HuitCarreau, cards.SeptTrefle},
			result: 2,
		},
		{
			name:   "With two nines of other color an height and a seven",
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
			result := CARREAU.EvaluateDeck(test.deck)
			assert.Equal(t, result, test.result)
		})
	}
}

func TestWinnerCARREAU(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name    string
		a, b, w cards.Card
	}{
		{name: "Dix carreau As Carreau", a: cards.DixCarreau, b: cards.AsCarreau, w: cards.AsCarreau},
		{name: "As Carreau Dix carreau ", a: cards.AsCarreau, b: cards.DixCarreau, w: cards.AsCarreau},
		{name: "Dix carreau Neuf pique", a: cards.DixCarreau, b: cards.NeufPique, w: cards.DixCarreau},
		{name: "Neuf pique Dix carreau ", a: cards.NeufPique, b: cards.DixCarreau, w: cards.DixCarreau},
		{name: "Neuf pique Dix pique", a: cards.NeufPique, b: cards.DixPique, w: cards.DixPique},
		{name: "Dix pique Neuf pique", a: cards.DixPique, b: cards.NeufPique, w: cards.DixPique},
		{name: "Neuf pique Huit trefle", a: cards.NeufPique, b: cards.HuitTrefle, w: cards.HuitTrefle},
		{name: "Huit trefle Neuf pique", a: cards.HuitTrefle, b: cards.NeufPique, w: cards.NeufPique},
		{name: "Neuf coeur Neuf pique", a: cards.NeufCoeur, b: cards.NeufPique, w: cards.NeufPique},
		{name: "Neuf pique Neuf coeur", a: cards.NeufPique, b: cards.NeufCoeur, w: cards.NeufCoeur},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := CARREAU.Winner(test.a, test.b)
			assert.Equal(t, test.w, w)
		})
	}
}

func TestMarshalJSONColorTake(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name   string
		take   ColorTake
		result string
	}{
		{name: "test marshal coeur", take: ColorTake{Couleur: "Coeur"}, result: "{\"name\":\"Coeur\"}"},
		{name: "test marshal carreau", take: ColorTake{Couleur: "Carreau"}, result: "{\"name\":\"Carreau\"}"},
		{name: "test marshal trefle", take: ColorTake{Couleur: "Trefle"}, result: "{\"name\":\"Trefle\"}"},
		{name: "test marshal pique", take: ColorTake{Couleur: "Pique"}, result: "{\"name\":\"Pique\"}"},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := test.take.MarshalJSON()

			require.NoError(t, err)
			assert.Equal(t, string(result), test.result)
		})
	}
}
