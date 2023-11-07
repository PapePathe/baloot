package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/player"
)

func TestReceivePlayingHandEvt(t *testing.T) {
	t.Parallel()

	p := player.NewPlayer()
	evt := ReceivePlayingHandEvt(p.PlayingHand.Cards, gametake.TOUT.Name())

	assert.Equal(t, ReceivePlayingHand, evt.ID)
	assert.Equal(t, "Tout", evt.Take)
}

func TestReceiveDeckEvt(t *testing.T) {
	t.Parallel()

	p := player.NewPlayer()
	evt := ReceiveDeckEvt(*p, [4]cards.Card{}, 0, 23)

	assert.Equal(t, ReceiveDeck, evt.ID)
	assert.Equal(t, 0, evt.ScoreTeamA)
	assert.Equal(t, 23, evt.ScoreTeamB)
	assert.Equal(t, p.PlayingHand.Cards, evt.Player)
	assert.Equal(t, [4]cards.Card{}, evt.Deck)
}

func TestReceiveTakeHandEvt(t *testing.T) {
	t.Parallel()

	p := player.NewPlayer()
	takes := []string{"Passe", "Cent"}
	evt := ReceiveTakeHandEvt(*p, takes)

	assert.Equal(t, ReceiveTakeHand, evt.ID)
	assert.Equal(t, takes, evt.AvailableTakes)
}

func TestBroadcastPlayerTakeEvt(t *testing.T) {
	t.Parallel()

	takes := []string{"Passe", "Cent"}
	evt := BroadcastPlayerTakeEvt("Carreau", 0, takes)

	assert.Equal(t, BroadcastPlayerTake, evt.ID)
	assert.Equal(t, "Carreau", evt.Take)
	assert.Equal(t, 0, evt.PlayerID)
	assert.Equal(t, takes, evt.AvailableTakes)
}
