package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/player"
)

func TestReceivePlayingHandEvt(t *testing.T) {
	p := player.NewPlayer()
	evt := ReceivePlayingHandEvt(*p, gametake.TOUT)

	assert.Equal(t, ReceivePlayingHand, evt.ID)
	assert.Equal(t, 0, evt.Player.ID)
	assert.Equal(t, gametake.TOUT, evt.Take)
}

func TestReceiveDeckEvt(t *testing.T) {
	p := player.NewPlayer()
	evt := ReceiveDeckEvt(*p, [4]cards.Card{}, 0, 23)

	assert.Equal(t, ReceiveDeck, evt.ID)
	assert.Equal(t, 0, evt.ScoreTeamA)
	assert.Equal(t, 23, evt.ScoreTeamB)
	assert.Equal(t, 0, evt.Player.ID)
	assert.Equal(t, [4]cards.Card{}, evt.Deck)
}

func TestReceiveTakeHandEvt(t *testing.T) {
	p := player.NewPlayer()
	takes := []gametake.GameTake{gametake.PASSE, gametake.TOUT}
	evt := ReceiveTakeHandEvt(*p, takes)

	assert.Equal(t, ReceiveTakeHand, evt.ID)
	assert.Equal(t, takes, evt.AvailableTakes)
}

func TestBroadcastPlayerTakeEvt(t *testing.T) {
	takes := []gametake.GameTake{gametake.PASSE, gametake.TOUT}
	evt := BroadcastPlayerTakeEvt("Carreau", 0, takes)

	assert.Equal(t, BroadcastPlayerTake, evt.ID)
	assert.Equal(t, "Carreau", evt.Take)
	assert.Equal(t, 0, evt.PlayerID)
	assert.Equal(t, takes, evt.AvailableTakes)
}
