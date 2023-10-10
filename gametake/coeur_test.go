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
			tout := Coeur{}
			result := tout.EvaluateCard(test.card)
			assert.Equal(t, result, test.value)
		})
	}
}
