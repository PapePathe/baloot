package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/player"
)

func TestAddPlayer(t *testing.T) {
	t.Parallel()

	game1 := NewGame()
	p := player.NewPlayer()

	require.NoError(t, game1.AddPlayer(p))
	assert.Equal(t, p, game1.players[0])
	assert.Equal(t, 0, p.GetID())
	assert.Equal(t, 1, game1.NombreJoueurs)

	msg := "distributed cards count should be equal to five cards"
	assert.Equal(t, 5, game1.CartesDistribuees, msg)
}

func TestAddMoreThanFivePlayers(t *testing.T) {
	t.Parallel()

	game1 := NewGame()

	for i := 0; i < 4; i++ {
		err := game1.AddPlayer(player.NewPlayer())
		require.NoError(t, err)
		assert.Equal(t, i, game1.players[i].GetID())
	}

	np := player.NewPlayer()
	err := game1.AddPlayer(np)
	require.ErrorIs(t, ErrGameIsFull, err)
}

func TestAddTake(t *testing.T) {
	t.Parallel()

	game1, player1 := NewGame(), player.NewPlayer()
	err := game1.AddPlayer(player1)
	require.NoError(t, err)

	err = game1.AddTake(player1.GetID(), gametake.TOUT)
	require.NoError(t, err)

	err = game1.AddTake(player1.GetID(), gametake.TOUT)
	require.Error(t, err)
	assert.Equal(t, player1.Take, &gametake.TOUT)
	assert.Equal(t, game1.take, gametake.TOUT)
}

func TestAddTakeLessThanGameTake(t *testing.T) {
	t.Parallel()

	game2 := NewGame()
	player1, player2 := player.NewPlayer(), player.NewPlayer()
	err := game2.AddPlayer(player1)
	require.NoError(t, err)
	err = game2.AddPlayer(player2)
	require.NoError(t, err)
	err = game2.AddTake(player1.GetID(), gametake.CENT)
	require.NoError(t, err)

	err2 := game2.AddTake(player2.GetID(), gametake.PIQUE)
	assert.Equal(t, ErrBadTake, err2)
}

func TestAddTakeLessThanGameTakeButIsPASS(t *testing.T) {
	t.Parallel()

	game2 := setupGame(2)
	player1, player2 := game2.players[0], game2.players[1]
	err := game2.AddTake(player1.GetID(), gametake.CENT)
	require.NoError(t, err)
	err = game2.AddTake(player2.GetID(), gametake.PASSE)
	require.NoError(t, err)
}

func TestAddTakeForPlayerThatHasTaken(t *testing.T) {
	t.Parallel()

	g, p := NewGame(), player.NewPlayer()
	err := g.AddPlayer(p)
	require.NoError(t, err)

	p.Take = &gametake.CENT
	err = g.AddTake(p.GetID(), gametake.TOUT)

	assert.Equal(t, ErrDuplicatePlayerTake, err)
}

func TestAddTakeLessThanCurrentGameTake(t *testing.T) {
	t.Parallel()

	g, p := NewGame(), player.NewPlayer()
	err := g.AddPlayer(p)
	require.NoError(t, err)
	err = g.AddTake(p.GetID(), gametake.TOUT)
	require.NoError(t, err)
}

func TestAddTakeGreaterThanCurrentGameTake(t *testing.T) {
	t.Parallel()

	g, p := NewGame(), player.NewPlayer()
	err := g.AddPlayer(p)
	require.NoError(t, err)
	err = g.AddTake(p.GetID(), gametake.TOUT)
	require.NoError(t, err)
}

func TestAddTakePassDoesNotChangeGameTake(t *testing.T) {
	t.Parallel()

	game1 := setupGame(2)
	player1, player2 := game1.players[0], game1.players[1]
	err := game1.AddTake(player1.GetID(), gametake.CENT)
	require.NoError(t, err)
	err = game1.AddTake(player2.GetID(), gametake.PASSE)
	require.NoError(t, err)
	assert.Equal(t, game1.GetTake(), gametake.CENT)
}

func TestDispatchCards(t *testing.T) {
	t.Parallel()

	game1 := setupGame(4)
	err := game1.DispatchCards()

	require.NoError(t, err)
	assert.Equal(t, 32, game1.CartesDistribuees)

	for _, p := range game1.players {
		assert.Len(t, p.PlayingHand.Cards, 8)
	}
}

func TestDispatchCardsIsIdempotent(t *testing.T) {
	t.Parallel()

	g := setupGame(4)
	err := g.DispatchCards()

	require.NoError(t, err)
	require.ErrorIs(t, g.DispatchCards(), ErrCardsAlreadyDispatched)
}

func TestNewGame(t *testing.T) {
	t.Parallel()

	g := NewGame()

	assert.False(t, g.TakesFinished)
}

func TestPlayCard(t *testing.T) {
	t.Parallel()

	t.Run("Player can play one of his cards", func(t *testing.T) {
		t.Parallel()

		game1 := setupGame(4)
		err := game1.DispatchCards()
		require.NoError(t, err)
		player1, player2, player3, player4 := game1.players[0], game1.players[1], game1.players[3], game1.players[3]

		card1 := player1.PlayingHand.Cards[0]
		err = game1.PlayCard(player1.GetID(), player1.PlayingHand.Cards[0])
		pli := [4]cards.Card{card1}
		require.NoError(t, err)
		assert.Equal(t, pli, game1.Plis[0])
		assert.Equal(t, 1, game1.pliCardsCount)
		assert.Equal(t, 0, game1.nombrePli)

		card2 := player2.PlayingHand.Cards[0]
		err2 := game1.PlayCard(player2.GetID(), card2)
		pli2 := [4]cards.Card{card1, card2}
		require.NoError(t, err2)
		assert.Equal(t, pli2, game1.Plis[0])
		assert.Equal(t, 2, game1.pliCardsCount)
		assert.Equal(t, 0, game1.nombrePli)

		card3 := player3.PlayingHand.Cards[0]
		err3 := game1.PlayCard(player3.GetID(), card3)
		pli3 := [4]cards.Card{card1, card2, card3}
		require.NoError(t, err3)
		assert.Equal(t, pli3, game1.Plis[0])
		assert.Equal(t, 3, game1.pliCardsCount)
		assert.Equal(t, 0, game1.nombrePli)

		c4 := player3.PlayingHand.Cards[0]
		err4 := game1.PlayCard(player4.GetID(), player4.PlayingHand.Cards[0])
		pli4 := [4]cards.Card{card1, card2, card3, c4}
		require.NoError(t, err4)
		assert.Equal(t, pli4, game1.Plis[0])
		assert.Equal(t, 0, game1.pliCardsCount)
		assert.Equal(t, 1, game1.nombrePli)

		cplayer1 := player1.PlayingHand.Cards[1]
		err = game1.PlayCard(player1.GetID(), player1.PlayingHand.Cards[1])
		require.NoError(t, err)
		pli1 := [4]cards.Card{cplayer1}
		assert.Equal(t, pli1, game1.Plis[1])
		assert.Equal(t, 1, game1.pliCardsCount)
		assert.Equal(t, 1, game1.nombrePli)
	})

	t.Run("Player One cannot play a card he does not have", func(t *testing.T) {
		t.Parallel()

		game1 := setupGame(4)
		err := game1.DispatchCards()
		require.NoError(t, err)

		player1, player2 := game1.players[0], game1.players[1]

		err = game1.PlayCard(player1.GetID(), player2.PlayingHand.Cards[0])
		require.ErrorIs(t, err, ErrCardNotFoundInPlayerHand)
	})
}

func TestNextRound(t *testing.T) {
	t.Parallel()

	g := NewGame()

	assert.Equal(t, [4]int{0, 1, 2, 3}, g.NextRound(0))
	assert.Equal(t, [4]int{1, 2, 3, 0}, g.NextRound(1))
	assert.Equal(t, [4]int{2, 3, 0, 1}, g.NextRound(2))
	assert.Equal(t, [4]int{3, 0, 1, 2}, g.NextRound(3))
}

type score struct {
	a int
	b int
}

var playcardNextTestcases = []struct {
	name string
	game *Game
	take gametake.GameTake
	p1   []cards.Card
	p2   []cards.Card
	p3   []cards.Card
	p4   []cards.Card
	s0   score
}{
	{
		name: gametake.TOUT.Name(),
		take: gametake.TOUT,
		game: setupGame(4),
		p1: []cards.Card{
			cards.ValetCoeur, cards.NeufCoeur, cards.AsCoeur, cards.DixCoeur,
			cards.RoiCoeur, cards.ValetCarreau, cards.NeufCarreau, cards.AsCarreau,
		},
		p2: []cards.Card{
			cards.ValetTrefle, cards.NeufTrefle, cards.AsTrefle, cards.DixTrefle,
			cards.HuitCarreau, cards.SeptCarreau, cards.HuitPique, cards.NeufPique,
		},
		p3: []cards.Card{
			cards.DameCoeur, cards.HuitCoeur, cards.SeptCoeur, cards.RoiTrefle,
			cards.DameTrefle, cards.HuitTrefle, cards.RoiCarreau, cards.SeptPique,
		},
		p4: []cards.Card{
			cards.ValetPique, cards.AsPique, cards.DixPique, cards.RoiPique,
			cards.DamePique, cards.SeptTrefle, cards.DixCarreau, cards.DameCarreau,
		},
		s0: score{63, 0},
	},
	{
		name: gametake.CENT.Name(),
		take: gametake.CENT,
		game: setupGame(4),
		p1: []cards.Card{
			cards.AsCoeur, cards.DixCoeur, cards.RoiCoeur, cards.ValetCoeur,
			cards.NeufCoeur, cards.AsCarreau, cards.DixCarreau, cards.DameCarreau,
		},
		p2: []cards.Card{
			cards.ValetTrefle, cards.NeufTrefle, cards.AsTrefle, cards.DixTrefle,
			cards.HuitCarreau, cards.HuitPique, cards.SeptCarreau, cards.NeufPique,
		},
		p3: []cards.Card{
			cards.DameCoeur, cards.HuitCoeur, cards.SeptCoeur, cards.RoiTrefle,
			cards.DameTrefle, cards.HuitTrefle, cards.RoiCarreau, cards.SeptPique,
		},
		p4: []cards.Card{
			cards.ValetPique, cards.AsPique, cards.DixPique, cards.RoiPique,
			cards.DamePique, cards.SeptTrefle, cards.ValetCarreau, cards.NeufCarreau,
		},
		s0: score{18, 0},
	},
	{
		name: gametake.COEUR.Name(),
		take: gametake.COEUR,
		game: setupGame(4),
		p1: []cards.Card{
			cards.ValetCoeur, cards.NeufCoeur, cards.AsCoeur, cards.DixCoeur,
			cards.RoiCoeur, cards.ValetCarreau, cards.NeufCarreau, cards.AsCarreau,
		},
		p2: []cards.Card{
			cards.ValetTrefle, cards.NeufTrefle, cards.AsTrefle, cards.DixTrefle,
			cards.HuitCarreau, cards.HuitPique, cards.SeptCarreau, cards.NeufPique,
		},
		p3: []cards.Card{
			cards.DameCoeur, cards.HuitCoeur, cards.SeptCoeur, cards.RoiTrefle,
			cards.DameTrefle, cards.HuitTrefle, cards.RoiCarreau, cards.SeptPique,
		},
		p4: []cards.Card{
			cards.ValetPique, cards.AsPique, cards.DixPique, cards.RoiPique,
			cards.DamePique, cards.SeptTrefle, cards.DixCarreau, cards.DameCarreau,
		},
		s0: score{27, 0},
	},
	{
		name: gametake.CARREAU.Name(),
		take: gametake.CARREAU,
		game: setupGame(4),
		p1: []cards.Card{
			cards.ValetCarreau, cards.NeufCarreau, cards.AsCarreau, cards.DixCarreau,
			cards.DameCarreau, cards.AsCoeur, cards.DixCoeur, cards.RoiCoeur,
		},
		p2: []cards.Card{
			cards.ValetTrefle, cards.NeufTrefle, cards.AsTrefle, cards.DixTrefle,
			cards.HuitCarreau, cards.HuitPique, cards.SeptCarreau, cards.NeufPique,
		},
		p3: []cards.Card{
			cards.DameCoeur, cards.HuitCoeur, cards.SeptCoeur, cards.RoiTrefle,
			cards.DameTrefle, cards.HuitTrefle, cards.RoiCarreau, cards.SeptPique,
		},
		p4: []cards.Card{
			cards.ValetPique, cards.AsPique, cards.DixPique, cards.RoiPique,
			cards.DamePique, cards.SeptTrefle, cards.ValetCoeur, cards.NeufCoeur,
		},
		s0: score{27, 0},
	},
	{
		name: gametake.PIQUE.Name(),
		take: gametake.PIQUE,
		game: setupGame(4),
		p1: []cards.Card{
			cards.ValetPique, cards.AsPique, cards.DixPique, cards.RoiPique,
			cards.DamePique, cards.SeptTrefle, cards.DixCarreau, cards.DameCarreau,
		},
		p2: []cards.Card{
			cards.ValetTrefle, cards.NeufTrefle, cards.AsTrefle, cards.DixTrefle,
			cards.HuitCarreau, cards.HuitPique, cards.SeptCarreau, cards.NeufPique,
		},
		p3: []cards.Card{
			cards.DameCoeur, cards.HuitCoeur, cards.SeptCoeur, cards.RoiTrefle,
			cards.DameTrefle, cards.HuitTrefle, cards.RoiCarreau, cards.SeptPique,
		},
		p4: []cards.Card{
			cards.ValetCoeur, cards.NeufCoeur, cards.AsCoeur, cards.DixCoeur,
			cards.RoiCoeur, cards.ValetCarreau, cards.NeufCarreau, cards.AsCarreau,
		},
		s0: score{27, 0},
	},
	{
		name: gametake.TREFLE.Name(),
		take: gametake.TREFLE,
		game: setupGame(4),
		p1: []cards.Card{
			cards.ValetTrefle, cards.NeufTrefle, cards.AsTrefle, cards.DixTrefle,
			cards.AsCarreau, cards.AsCoeur, cards.DixCoeur, cards.DixCarreau,
		},
		p2: []cards.Card{
			cards.ValetCoeur, cards.NeufCoeur, cards.HuitPique, cards.SeptCarreau,
			cards.RoiCoeur, cards.ValetCarreau, cards.NeufCarreau, cards.HuitCarreau,
		},
		p3: []cards.Card{
			cards.DameCoeur, cards.HuitCoeur, cards.SeptCoeur, cards.RoiTrefle,
			cards.DameTrefle, cards.HuitTrefle, cards.RoiCarreau, cards.SeptPique,
		},
		p4: []cards.Card{
			cards.ValetPique, cards.AsPique, cards.DixPique, cards.RoiPique,
			cards.DamePique, cards.SeptTrefle, cards.NeufPique, cards.DameCarreau,
		},
		s0: score{27, 0},
	},
}

func TestPlayCardNext(t *testing.T) {
	t.Parallel()

	for _, test := range playcardNextTestcases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			game1 := test.game
			player1, player2, player3, player4 := game1.players[0], game1.players[1], game1.players[2], game1.players[3]
			err := game1.AddTake(0, test.take)
			require.NoError(t, err)
			if test.take != gametake.TOUT {
				err := game1.DispatchCards()
				require.NoError(t, err)
			}

			player1.PlayingHand.Cards, player2.PlayingHand.Cards = test.p1, test.p2
			player3.PlayingHand.Cards, player4.PlayingHand.Cards = test.p3, test.p4
			pli := [4]cards.Card{
				player1.PlayingHand.Cards[0],
				player2.PlayingHand.Cards[0],
				player3.PlayingHand.Cards[0],
				player4.PlayingHand.Cards[0],
			}
			for _, p := range game1.players {
				err := game1.PlayCardNext(p.GetID(), p.PlayingHand.Cards[0])
				require.NoError(t, err)
			}

			winner := game1.Decks[0].winner
			fmt.Println(game1.Decks)

			t.Run("Test deck winner", func(t *testing.T) {
				assert.Equal(t, 1, game1.nombrePli)
				assert.Equal(t, player1.GetID(), winner)
				assert.Equal(t, game1.Decks[0].cards, pli)
			})

			t.Run("Test deck scoring", func(t *testing.T) {
				sTeamA, sTeamB := game1.scoreTeamA, game1.scoreTeamB
				assert.Equal(t, test.s0.a, sTeamA)
				assert.Equal(t, test.s0.b, sTeamB)
			})

			t.Run("Test game ring for next round", func(t *testing.T) {
				nextRing := game1.NextRound(winner)
				assert.Equal(t, nextRing, game1.ring)
			})

			fmt.Println("take ", game1.GetTake().Name(), "deck", pli, "winner:", winner)
		})
	}
}

func TestTakesComplete(t *testing.T) {
	t.Parallel()

	t.Run("first player taking tout completes takes", func(t *testing.T) {
		t.Parallel()

		game1 := setupGame(4)

		err := game1.AddTake(game1.GetPlayers()[0].GetID(), gametake.TOUT)
		require.NoError(t, err)
		assert.True(t, game1.takesComplete())
	})

	t.Run("second player takes tout completes the takes", func(t *testing.T) {
		t.Parallel()

		game1 := setupGame(4)

		err := game1.AddTake(game1.GetPlayers()[0].GetID(), gametake.TREFLE)
		require.NoError(t, err)
		assert.False(t, game1.takesComplete())

		err = game1.AddTake(game1.GetPlayers()[1].GetID(), gametake.TOUT)
		require.NoError(t, err)
		assert.True(t, game1.takesComplete())
	})

	t.Run("four takes without tout completes the takes", func(t *testing.T) {
		t.Parallel()

		game1 := setupGame(4)

		err := game1.AddTake(game1.GetPlayers()[0].GetID(), gametake.TREFLE)
		require.NoError(t, err)
		assert.False(t, game1.takesComplete())

		err = game1.AddTake(game1.GetPlayers()[1].GetID(), gametake.CARREAU)
		require.NoError(t, err)
		assert.False(t, game1.takesComplete())

		err = game1.AddTake(game1.GetPlayers()[2].GetID(), gametake.PASSE)
		require.NoError(t, err)
		assert.False(t, game1.takesComplete())

		err = game1.AddTake(game1.GetPlayers()[3].GetID(), gametake.COEUR)
		require.NoError(t, err)
		assert.True(t, game1.takesComplete())
	})
}

func setupGame(playersCount int) *Game {
	game1 := NewGame()
	for i := 0; i < playersCount; i++ {
		err := game1.AddPlayer(player.NewPlayer())
		if err != nil {
			fmt.Println("ERROR SETTING UP A TEST GAME")
		}
	}

	return game1
}
