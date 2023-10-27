package game

import (
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/player"
)

type messageID int

var (
	ReceiveTakeHand     messageID = 1
	ReceivePlayingHand  messageID = 2
	SetTake             messageID = 3
	PlayCard            messageID = 4
	BroadcastPlayerTake messageID = 5
)

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
