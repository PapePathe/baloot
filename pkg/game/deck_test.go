package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

func TestNewDeck(t *testing.T) {
	t.Parallel()

	deck := NewDeck([4]int{0, 1, 2, 3}, gametake.TOUT)

	assert.Equal(t, deck.gametake, gametake.TOUT)
	assert.Equal(t, 0, deck.cardscount)
	assert.Equal(t, -1, deck.winner)
	assert.Equal(t, [4]int{0, 1, 2, 3}, deck.players)
	assert.Len(t, deck.players, 4)
	assert.Len(t, deck.cards, 4)
}

func TestAddCard(t *testing.T) {
	t.Parallel()

	t.Run("Can add one card to deck", func(t *testing.T) {
		t.Parallel()

		d := NewDeck([4]int{0, 1, 2, 3}, gametake.TOUT)
		err := d.AddCard(0, cards.ValetCarreau)

		require.NoError(t, err)
		assert.Equal(t, 1, d.cardscount)
		assert.Equal(t, [4]cards.Card{cards.ValetCarreau}, d.cards)
	})

	t.Run("Cannot add more than four cards to deck", func(t *testing.T) {
		t.Parallel()

		deck := NewDeck([4]int{0, 1, 2, 3}, gametake.TOUT)
		err := deck.AddCard(0, cards.ValetCarreau)
		require.NoError(t, err)
		err = deck.AddCard(1, cards.SeptCarreau)
		require.NoError(t, err)
		err = deck.AddCard(2, cards.HuitCarreau)
		require.NoError(t, err)
		err = deck.AddCard(3, cards.DixCarreau)
		require.NoError(t, err)

		err = deck.AddCard(0, cards.NeufCarreau)
		assert.Equal(t, err, ErrCannotAddMoreThanFourCardsToDeck)
	})

	t.Run("Cannot add empty card to deck", func(t *testing.T) {
		t.Parallel()

		d := NewDeck([4]int{0, 1, 2, 3}, gametake.TOUT)
		err := d.AddCard(0, cards.Card{Couleur: "", Genre: ""})

		assert.Equal(t, err, ErrCannotAddEmptyCardToDeck)
	})

	t.Run("Cannot add existing card to deck", func(t *testing.T) {
		t.Parallel()

		d := NewDeck([4]int{0, 1, 2, 3}, gametake.TOUT)
		err := d.AddCard(0, cards.ValetCarreau)
		require.NoError(t, err)
		err = d.AddCard(1, cards.ValetCarreau)

		assert.Equal(t, err, ErrCannotAddExistingCardToDeck)
	})

	t.Run("Cannot add card if it is not player's turn to play", func(t *testing.T) {
		t.Parallel()

		d := NewDeck([4]int{0, 1, 2, 3}, gametake.TOUT)
		err := d.AddCard(1, cards.ValetCarreau)

		assert.Equal(t, err, ErrNotYourTurnToPlay)
	})
}

type testCase struct {
	name     string
	cards    [4]cards.Card
	players  [4]int
	winner   int
	gametake gametake.GameTake
}

var testsFindWinner = []testCase{
	{
		name:     "gametake.TOUT with ValetCarreau SeptCarreau HuitCarreau DixCarreau",
		cards:    [4]cards.Card{cards.ValetCarreau, cards.SeptCarreau, cards.HuitCarreau, cards.DixCarreau},
		players:  [4]int{0, 1, 2, 3},
		gametake: gametake.TOUT,
		winner:   0,
	},
	{
		name:     "gametake.TOUT with SeptCarreau ValetCarreau HuitCarreau DixCarreau",
		cards:    [4]cards.Card{cards.SeptCarreau, cards.ValetCarreau, cards.HuitCarreau, cards.DixCarreau},
		players:  [4]int{0, 3, 1, 2},
		gametake: gametake.TOUT,
		winner:   3,
	},
	{
		name:     "gametake.TOUT with SeptCarreau AsCoeur HuitPique HuitTrefle",
		cards:    [4]cards.Card{cards.SeptCarreau, cards.AsCoeur, cards.HuitPique, cards.HuitTrefle},
		players:  [4]int{0, 3, 1, 2},
		gametake: gametake.TOUT,
		winner:   0,
	},
	{
		name:     "gametake.TOUT with SeptCarreau HuitCarrea HuitPique HuitTrefle",
		cards:    [4]cards.Card{cards.SeptCarreau, cards.HuitCarreau, cards.HuitPique, cards.HuitTrefle},
		players:  [4]int{3, 0, 1, 2},
		gametake: gametake.TOUT,
		winner:   0,
	},
	{
		name:     "gametake.TOUT with DixCarreau, NeufCarreau, AsCarreau, ValetCarreau",
		cards:    [4]cards.Card{cards.DixCarreau, cards.NeufCarreau, cards.AsCarreau, cards.ValetCarreau},
		players:  [4]int{0, 1, 2, 3},
		gametake: gametake.TOUT,
		winner:   3,
	},
	{
		name:     "gametake.TOUT with RoiCarreau, NeufPique, AsCoeur, ValetCoeur",
		cards:    [4]cards.Card{cards.RoiCarreau, cards.NeufPique, cards.AsCoeur, cards.ValetCoeur},
		players:  [4]int{0, 1, 2, 3},
		gametake: gametake.TOUT,
		winner:   0,
	},
	{
		name:     "gametake.CENT with RoiCarreau, NeufPique, AsCoeur, ValetCoeur",
		cards:    [4]cards.Card{cards.RoiCarreau, cards.NeufPique, cards.AsCoeur, cards.ValetCoeur},
		players:  [4]int{0, 1, 2, 3},
		gametake: gametake.CENT,
		winner:   0,
	},
	{
		name:     "gametake.CENT with RoiCarreau, DixCarreau, AsCoeur, ValetCoeur",
		cards:    [4]cards.Card{cards.RoiCarreau, cards.DixCarreau, cards.AsCoeur, cards.ValetCoeur},
		players:  [4]int{0, 1, 2, 3},
		gametake: gametake.CENT,
		winner:   1,
	},
	{
		name:     "gametake.CENT with RoiCarreau, DixCarreau, AsCoeur, AsCarreau",
		cards:    [4]cards.Card{cards.RoiCarreau, cards.DixCarreau, cards.AsCoeur, cards.AsCarreau},
		players:  [4]int{0, 1, 2, 3},
		gametake: gametake.CENT,
		winner:   3,
	},
	{
		name:     "gametake.CENT with SeptCarreau, HuitCarreau, AsCoeur, ValetCoeur",
		cards:    [4]cards.Card{cards.SeptCarreau, cards.HuitCarreau, cards.AsCoeur, cards.ValetCoeur},
		players:  [4]int{0, 1, 2, 3},
		gametake: gametake.CENT,
		winner:   1,
	},
	{
		name:     "gametake.COEUR with SeptCarreau, HuitCarreau, AsCoeur, ValetCoeur",
		cards:    [4]cards.Card{cards.SeptCarreau, cards.HuitCarreau, cards.AsCoeur, cards.ValetCoeur},
		players:  [4]int{0, 1, 2, 3},
		gametake: gametake.COEUR,
		winner:   3,
	},
	{
		name:     "gametake.COEUR with SeptCarreau, HuitCarreau, AsCoeur, AsPique",
		cards:    [4]cards.Card{cards.SeptCarreau, cards.HuitCarreau, cards.AsCoeur, cards.AsPique},
		players:  [4]int{0, 1, 2, 3},
		gametake: gametake.COEUR,
		winner:   2,
	},
}

func TestFindWinner(t *testing.T) {
	t.Parallel()

	for _, tc := range testsFindWinner {
		testcase := tc
		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			deck := NewDeck(testcase.players, testcase.gametake)
			for i, c := range testcase.cards {
				err := deck.AddCard(i, c)
				require.NoError(t, err)
			}

			assert.Equal(t, testcase.winner, deck.winner)
		})
	}
}
