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

type receiveDeckMsg struct {
	ID     messageID     `json:"id"`
	Player player.Player `json:"player"`
	Deck   [4]cards.Card `json:"deck"`
}

func ReceiveDeckMsg(p player.Player, d [4]cards.Card) receiveDeckMsg {
	msg := receiveDeckMsg{ID: ReceiveDeck, Player: p, Deck: d}

	return msg
}

type receivePlayingHandMsg struct {
	ID     messageID         `json:"id"`
	Take   gametake.GameTake `json:"gametake"`
	Player player.Player     `json:"player"`
}

func ReceivePlayingHandMsg(p player.Player, take gametake.GameTake) receivePlayingHandMsg {
	clientMessage := receivePlayingHandMsg{ID: ReceivePlayingHand, Player: p, Take: take}

	return clientMessage
}

type receiveTakeHandMsg struct {
	ID             messageID           `json:"id"`
	Player         player.Player       `json:"player"`
	AvailableTakes []gametake.GameTake `json:"available_takes"`
}

func ReceiveTakeHandMsg(p player.Player, takes []gametake.GameTake) receiveTakeHandMsg {
	clientMessage := receiveTakeHandMsg{ID: ReceiveTakeHand, Player: p, AvailableTakes: takes}

	return clientMessage
}

type setTake struct {
	ID       messageID `json:"id"`
	PlayerId int       `json:"player_id"`
	Gametake string    `json:"gametake"`
}

func SetTakeMsg(gt string, pid int) setTake {
	clientMessage := setTake{ID: SetTake, PlayerId: pid, Gametake: gt}

	return clientMessage
}

type broadcastPlayerTakeMsg struct {
	ID             messageID           `json:"id"`
	Take           string              `json:"take"`
	PlayerId       int                 `json:"player_id"`
	AvailableTakes []gametake.GameTake `json:"available_takes"`
}

func BroadcastPlayerTakeMsg(gt string, pid int, at []gametake.GameTake) broadcastPlayerTakeMsg {
	clientMessage := broadcastPlayerTakeMsg{ID: BroadcastPlayerTake, Take: gt, PlayerId: pid, AvailableTakes: at}

	return clientMessage
}
