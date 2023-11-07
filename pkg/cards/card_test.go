package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "(V Coeur)", ValetCoeur.String())
}

func TestDistribute(t *testing.T) {
	t.Parallel()

	set := CardSet{Jeu32Cartes}
	result := set.Distribute()

	if result == Jeu32Cartes {
		t.Errorf("shuffled cards %v should be different  %v", result, Jeu32Cartes)
	}
}

func TestIsValet(t *testing.T) {
	t.Parallel()

	assert.True(t, ValetCoeur.IsValet())
	assert.True(t, ValetTrefle.IsValet())
	assert.True(t, ValetCarreau.IsValet())
	assert.True(t, ValetPique.IsValet())

	assert.False(t, DameTrefle.IsValet())
	assert.False(t, AsCoeur.IsValet())
}

func TestIsAce(t *testing.T) {
	t.Parallel()

	assert.True(t, AsCoeur.IsAce())
	assert.True(t, AsTrefle.IsAce())
	assert.True(t, AsCarreau.IsAce())
	assert.True(t, AsPique.IsAce())

	assert.False(t, DameTrefle.IsAce())
	assert.False(t, ValetCoeur.IsAce())
}

func TestIsNine(t *testing.T) {
	t.Parallel()

	assert.True(t, NeufCoeur.IsNine())
	assert.True(t, NeufTrefle.IsNine())
	assert.True(t, NeufCarreau.IsNine())
	assert.True(t, NeufPique.IsNine())

	assert.False(t, DameTrefle.IsNine())
	assert.False(t, ValetCoeur.IsNine())
}
