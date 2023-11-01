package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "pathe.co/zinx/gametake"
	. "pathe.co/zinx/pkg/cards"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck([4]int{0, 1, 2, 3}, TOUT)

	assert.Equal(t, d.gametake, TOUT)
	assert.Equal(t, d.cardscount, 0)
	assert.Equal(t, d.winner, -1)
	assert.Equal(t, d.players, [4]int{0, 1, 2, 3})
	assert.Equal(t, len(d.players), 4)
	assert.Equal(t, len(d.cards), 4)
}

func TestAddCard(t *testing.T) {
	t.Run("Can add one card to deck", func(t *testing.T) {
		d := NewDeck([4]int{0, 1, 2, 3}, TOUT)
		err := d.AddCard(0, ValetCarreau)

		assert.NoError(t, err)
		assert.Equal(t, d.cardscount, 1)
		assert.Equal(t, d.cards, [4]Card{ValetCarreau})
	})

	t.Run("Cannot add more than four cards to deck", func(t *testing.T) {
		d := NewDeck([4]int{0, 1, 2, 3}, TOUT)
		err := d.AddCard(0, ValetCarreau)
		assert.NoError(t, err)
		err = d.AddCard(1, SeptCarreau)
		assert.NoError(t, err)
		err = d.AddCard(2, HuitCarreau)
		assert.NoError(t, err)
		err = d.AddCard(3, DixCarreau)
		assert.NoError(t, err)

		err = d.AddCard(0, NeufCarreau)
		assert.Equal(t, err, ErrCannotAddMoreThanFourCardsToDeck)
	})

	t.Run("Cannot add empty card to deck", func(t *testing.T) {
		d := NewDeck([4]int{0, 1, 2, 3}, TOUT)
		err := d.AddCard(0, Card{})

		assert.Equal(t, err, ErrCannotAddEmptyCardToDeck)
	})

	t.Run("Cannot add existing card to deck", func(t *testing.T) {
		d := NewDeck([4]int{0, 1, 2, 3}, TOUT)
		err := d.AddCard(0, ValetCarreau)
		assert.NoError(t, err)
		err = d.AddCard(1, ValetCarreau)

		assert.Equal(t, err, ErrCannotAddExistingCardToDeck)
	})

	t.Run("Cannot add card if it is not player's turn to play", func(t *testing.T) {
		d := NewDeck([4]int{0, 1, 2, 3}, TOUT)
		err := d.AddCard(1, ValetCarreau)

		assert.Equal(t, err, ErrNotYourTurnToPlay)
	})
}

func TestFindWinner(t *testing.T) {
	type testCase struct {
		name     string
		cards    [4]Card
		players  [4]int
		winner   int
		gametake GameTake
	}
	tests := []testCase{
		testCase{
			name:     "TOUT with ValetCarreau SeptCarreau HuitCarreau DixCarreau",
			cards:    [4]Card{ValetCarreau, SeptCarreau, HuitCarreau, DixCarreau},
			players:  [4]int{0, 1, 2, 3},
			gametake: TOUT,
			winner:   0,
		},
		testCase{
			name:     "TOUT with SeptCarreau ValetCarreau HuitCarreau DixCarreau",
			cards:    [4]Card{SeptCarreau, ValetCarreau, HuitCarreau, DixCarreau},
			players:  [4]int{0, 3, 1, 2},
			gametake: TOUT,
			winner:   3,
		},
		testCase{
			name:     "TOUT with SeptCarreau AsCoeur HuitPique HuitTrefle",
			cards:    [4]Card{SeptCarreau, AsCoeur, HuitPique, HuitTrefle},
			players:  [4]int{0, 3, 1, 2},
			gametake: TOUT,
			winner:   0,
		},
		testCase{
			name:     "TOUT with SeptCarreau HuitCarrea HuitPique HuitTrefle",
			cards:    [4]Card{SeptCarreau, HuitCarreau, HuitPique, HuitTrefle},
			players:  [4]int{3, 0, 1, 2},
			gametake: TOUT,
			winner:   0,
		},
		testCase{
			name:     "TOUT with DixCarreau, NeufCarreau, AsCarreau, ValetCarreau",
			cards:    [4]Card{DixCarreau, NeufCarreau, AsCarreau, ValetCarreau},
			players:  [4]int{0, 1, 2, 3},
			gametake: TOUT,
			winner:   3,
		},
		testCase{
			name:     "TOUT with RoiCarreau, NeufPique, AsCoeur, ValetCoeur",
			cards:    [4]Card{RoiCarreau, NeufPique, AsCoeur, ValetCoeur},
			players:  [4]int{0, 1, 2, 3},
			gametake: TOUT,
			winner:   0,
		},
		testCase{
			name:     "CENT with RoiCarreau, NeufPique, AsCoeur, ValetCoeur",
			cards:    [4]Card{RoiCarreau, NeufPique, AsCoeur, ValetCoeur},
			players:  [4]int{0, 1, 2, 3},
			gametake: CENT,
			winner:   0,
		},
		testCase{
			name:     "CENT with RoiCarreau, DixCarreau, AsCoeur, ValetCoeur",
			cards:    [4]Card{RoiCarreau, DixCarreau, AsCoeur, ValetCoeur},
			players:  [4]int{0, 1, 2, 3},
			gametake: CENT,
			winner:   1,
		},
		testCase{
			name:     "CENT with RoiCarreau, DixCarreau, AsCoeur, AsCarreau",
			cards:    [4]Card{RoiCarreau, DixCarreau, AsCoeur, AsCarreau},
			players:  [4]int{0, 1, 2, 3},
			gametake: CENT,
			winner:   3,
		},
		testCase{
			name:     "CENT with SeptCarreau, HuitCarreau, AsCoeur, ValetCoeur",
			cards:    [4]Card{SeptCarreau, HuitCarreau, AsCoeur, ValetCoeur},
			players:  [4]int{0, 1, 2, 3},
			gametake: CENT,
			winner:   1,
		},
		testCase{
			name:     "COEUR with SeptCarreau, HuitCarreau, AsCoeur, ValetCoeur",
			cards:    [4]Card{SeptCarreau, HuitCarreau, AsCoeur, ValetCoeur},
			players:  [4]int{0, 1, 2, 3},
			gametake: COEUR,
			winner:   3,
		},
		testCase{
			name:     "COEUR with SeptCarreau, HuitCarreau, AsCoeur, AsPique",
			cards:    [4]Card{SeptCarreau, HuitCarreau, AsCoeur, AsPique},
			players:  [4]int{0, 1, 2, 3},
			gametake: COEUR,
			winner:   2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDeck(tc.players, tc.gametake)
			for i, c := range tc.cards {
				err := d.AddCard(i, c)
				assert.Equal(t, nil, err)
			}

			assert.Equal(t, tc.winner, d.winner)
		})
	}
}
