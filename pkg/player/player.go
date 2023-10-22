package player

import (
	"sort"

	"github.com/gofiber/contrib/websocket"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

type Player struct {
	Hand        Hand               `json:"hand"`
	PlayingHand PlayingHand        `json:"playing_hand"`
	Transport   IPlayerTransport   `json:"-"`
	Take        *gametake.GameTake `json:"take"`
	ID          int                `json:"id"`
	Conn        *websocket.Conn    `json:"-"`
}

func NewPlayer() *Player {
	return &Player{
		Hand:        Hand{},
		PlayingHand: PlayingHand{},
		Transport:   JSONMarshaler{},
		Take:        nil,
	}
}

func (p *Player) SetID(id int) {
	p.ID = id
}

func (p *Player) GetID() int {
	return p.ID
}

func (p *Player) SetForTransport() ([]byte, error) {
	return p.Transport.Marshal(*p)
}

func (p *Player) GetFromTransport(b []byte) error {
	err := p.Transport.UnMarshal(b, p)

	if err != nil {
		return err
	}

	return nil
}

func (p *Player) OrderedCards() map[string][]cards.Card {
	m := make(map[string][]cards.Card)

	for _, c := range p.Hand.Cards {
		_, ok := m[c.Couleur]

		if ok {
			m[c.Couleur] = append(m[c.Couleur], c)
		} else {
			m[c.Couleur] = []cards.Card{c}
		}

	}

	return m
}

func (p *Player) OrderedCardsForTake(t gametake.GameTake) [5]cards.Card {
	m := make(map[string][]cards.Card)

	for _, c := range p.Hand.Cards {
		_, ok := m[c.Couleur]

		if ok {
			m[c.Couleur] = append(m[c.Couleur], c)
		} else {
			m[c.Couleur] = []cards.Card{c}
		}
		sorter := SortByColorAndType{m[c.Couleur], t}
		sort.Sort(sorter)
	}

	result := []cards.Card{}

	for _, cards := range m {
		for _, c := range cards {
			result = append(result, c)
		}
	}

	return [5]cards.Card(result)
}

func (p *Player) OrderedCardsForPlaying(t gametake.GameTake) [8]cards.Card {
	m := make(map[string][]cards.Card)

	for _, c := range p.PlayingHand.Cards {
		_, ok := m[c.Couleur]

		if ok {
			m[c.Couleur] = append(m[c.Couleur], c)
		} else {
			m[c.Couleur] = []cards.Card{c}
		}
		sorter := SortByColorAndType{m[c.Couleur], t}
		sort.Sort(sorter)
	}

	result := []cards.Card{}

	for _, cards := range m {
		for _, c := range cards {
			result = append(result, c)
		}
	}

	return [8]cards.Card(result)
}
