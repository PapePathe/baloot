package player

import (
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

type IPlayerActions interface {
	PlayCard(c cards.Card)
}

type messageID int

type PlayEvent struct {
	PlayerID int
	Card     cards.Card
}

type PlayEventDetails struct {
	Deck       []cards.Card
	Play       cards.Card
	Hand       []cards.Card
	History    [][4]cards.Card
	Gametake   gametake.GameTake
	PlayerType string
}

func (ped PlayEventDetails) AsKafkaMessage(topic string) (kafka.Message, error) {
	b, err := json.Marshal(ped)
	if err != nil {
		return kafka.Message{}, fmt.Errorf("error marshalling constrainedTake %w", err)
	}

	return kafka.Message{Key: []byte(ped.Gametake.Name()), Topic: topic, Value: b}, nil
}

var (
	BroadcastPlayerTake messageID = 5
	PlayCard            messageID = 4
	ReceiveTakeHand     messageID = 1
	ReceivePlayingHand  messageID = 2
	ReceiveDeck         messageID = 6
	SetTake             messageID = 3
)

type ReceiveDeckMsg struct {
	ID                      messageID             `json:"id"`
	Player                  []cards.Card          `json:"player"`
	Deck                    [4]cards.Card         `json:"deck"`
	ScoreTeamA              int                   `json:"scoreTeamA"`
	ScoreTeamB              int                   `json:"scoreTeamB"`
	NextPlayer              int                   `json:nextPlayer`
	PlayChannel             chan PlayEvent        `json:"-"`
	PlayEventDetailsChannel chan PlayEventDetails `json:"-"`
	Take                    gametake.GameTake     `json:"-"`
}

func ReceiveDeckEvt(p BelotePlayer, d [4]cards.Card, a int, b int, n int, c chan PlayEvent, cpd chan PlayEventDetails, t gametake.GameTake) ReceiveDeckMsg {
	msg := ReceiveDeckMsg{
		ID:                      ReceiveDeck,
		Player:                  p.GetPlayingHand().OrderedCardsForPlaying(t),
		Deck:                    d,
		ScoreTeamA:              a,
		ScoreTeamB:              b,
		NextPlayer:              n,
		PlayChannel:             c,
		PlayEventDetailsChannel: cpd,
		Take:                    t,
	}

	return msg
}

type ReceivePlayingHandMsg struct {
	ID    messageID    `json:"id"`
	Take  string       `json:"gametake"`
	Cards []cards.Card `json:"player"`
}

func ReceivePlayingHandEvt(p []cards.Card, take string) ReceivePlayingHandMsg {
	clientMessage := ReceivePlayingHandMsg{ID: ReceivePlayingHand, Cards: p, Take: take}

	return clientMessage
}

type ReceiveTakeHandMsg struct {
	ID             messageID `json:"id"`
	Player         Player    `json:"player"`
	AvailableTakes []string  `json:"availableTakes"`
}

func ReceiveTakeHandEvt(p Player, takes []string) ReceiveTakeHandMsg {
	clientMessage := ReceiveTakeHandMsg{ID: ReceiveTakeHand, Player: p, AvailableTakes: takes}

	return clientMessage
}

type BroadcastPlayerTakeMsg struct {
	ID             messageID `json:"id"`
	Take           string    `json:"take"`
	PlayerID       int       `json:"playerId"`
	AvailableTakes []string  `json:"availableTakes"`
}

func BroadcastPlayerTakeEvt(gt string, pid int, at []string) BroadcastPlayerTakeMsg {
	clientMessage := BroadcastPlayerTakeMsg{ID: BroadcastPlayerTake, Take: gt, PlayerID: pid, AvailableTakes: at}

	return clientMessage
}
