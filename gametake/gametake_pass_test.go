package gametake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pathe.co/zinx/pkg/cards"
)

func TestMarshalJSONPASSE(t *testing.T) {
	t.Parallel()
	result, err := Passe{}.MarshalJSON()

	require.NoError(t, err)
	assert.Equal(t, string(result), "{\"name\":\"Passe\"}")
}

func TestMarshalWinnerPASSE(t *testing.T) {
	t.Parallel()
	result := Passe{}.Winner(cards.ValetCoeur, cards.RoiCoeur)

	assert.Equal(t, result, cards.ValetCoeur)
}

func TestMarshalEvaluateCardForWinPASSE(t *testing.T) {
	t.Parallel()
	result := Passe{}.EvaluateCardForWin(cards.RoiCoeur)

	assert.Equal(t, 0, result)
}

func TestMarshalEvaluateCardPASSE(t *testing.T) {
	t.Parallel()
	result, sameColor := Passe{}.EvaluateCard(cards.RoiCoeur)

	assert.Equal(t, 0, result)
	assert.True(t, sameColor)
}

func TestEvaluateDeckPASSE(t *testing.T) {
	t.Parallel()
	result := Passe{}.EvaluateDeck([4]cards.Card{})

	assert.Equal(t, 0, result)
}

func TestGreaterThanPASSE(t *testing.T) {
	t.Parallel()
	assert.False(t, PASSE.GreaterThan(TREFLE))
	assert.False(t, PASSE.GreaterThan(CARREAU))
	assert.False(t, PASSE.GreaterThan(COEUR))
	assert.False(t, PASSE.GreaterThan(PIQUE))
	assert.False(t, PASSE.GreaterThan(CENT))
	assert.False(t, PASSE.GreaterThan(TOUT))
}
