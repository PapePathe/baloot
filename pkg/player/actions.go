package player

import "pathe.co/zinx/pkg/cards"

type IPlayerActions interface {
	PlayCard(c cards.Card)
}

type messageID int

type PlayEvent struct {
	PlayerID int
	Card     cards.Card
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
	ID          messageID      `json:"id"`
	Player      []cards.Card   `json:"player"`
	Deck        [4]cards.Card  `json:"deck"`
	ScoreTeamA  int            `json:"scoreTeamA"`
	ScoreTeamB  int            `json:"scoreTeamB"`
	NextPlayer  int            `json:nextPlayer`
	PlayChannel chan PlayEvent `json:"-"`
}

func ReceiveDeckEvt(p BelotePlayer, d [4]cards.Card, a int, b int, n int, c chan PlayEvent) ReceiveDeckMsg {
	msg := ReceiveDeckMsg{
		ID:          ReceiveDeck,
		Player:      p.GetPlayingHand().Cards,
		Deck:        d,
		ScoreTeamA:  a,
		ScoreTeamB:  b,
		NextPlayer:  n,
		PlayChannel: c,
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
