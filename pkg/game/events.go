package game

import (
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/player"
)

type messageID int

var (
	BroadcastPlayerTake messageID = 5
	PlayCard            messageID = 4
	ReceiveTakeHand     messageID = 1
	ReceivePlayingHand  messageID = 2
	ReceiveDeck         messageID = 6
	SetTake             messageID = 3
)

type ReceiveDeckMsg struct {
	ID         messageID     `json:"id"`
	Player     player.Player `json:"player"`
	Deck       [4]cards.Card `json:"deck"`
	ScoreTeamA int           `json:"scoreTeamA"`
	ScoreTeamB int           `json:"scoreTeamB"`
}

func ReceiveDeckEvt(p player.Player, d [4]cards.Card, a int, b int) ReceiveDeckMsg {
	msg := ReceiveDeckMsg{
		ID:         ReceiveDeck,
		Player:     p,
		Deck:       d,
		ScoreTeamA: a,
		ScoreTeamB: b,
	}

	return msg
}

type ReceivePlayingHandMsg struct {
	ID     messageID         `json:"id"`
	Take   gametake.GameTake `json:"gametake"`
	Player player.Player     `json:"player"`
}

func ReceivePlayingHandEvt(p player.Player, take gametake.GameTake) ReceivePlayingHandMsg {
	clientMessage := ReceivePlayingHandMsg{ID: ReceivePlayingHand, Player: p, Take: take}

	return clientMessage
}

type ReceiveTakeHandMsg struct {
	ID             messageID           `json:"id"`
	Player         player.Player       `json:"player"`
	AvailableTakes []gametake.GameTake `json:"availableTakes"`
}

func ReceiveTakeHandEvt(p player.Player, takes []gametake.GameTake) ReceiveTakeHandMsg {
	clientMessage := ReceiveTakeHandMsg{ID: ReceiveTakeHand, Player: p, AvailableTakes: takes}

	return clientMessage
}

type BroadcastPlayerTakeMsg struct {
	ID             messageID           `json:"id"`
	Take           string              `json:"take"`
	PlayerID       int                 `json:"playerId"`
	AvailableTakes []gametake.GameTake `json:"availableTakes"`
}

func BroadcastPlayerTakeEvt(gt string, pid int, at []gametake.GameTake) BroadcastPlayerTakeMsg {
	clientMessage := BroadcastPlayerTakeMsg{ID: BroadcastPlayerTake, Take: gt, PlayerID: pid, AvailableTakes: at}

	return clientMessage
}
