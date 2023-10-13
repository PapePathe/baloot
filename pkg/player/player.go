package player

import (
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

type Player struct {
	Hand      Hand               `json:"hand"`
	Transport IPlayerTransport   `json:"-"`
	Take      *gametake.GameTake `json:"take"`
	id        int                `json:"id"`
}

func NewPlayer() *Player {
	return &Player{Hand: Hand{}, Transport: JSONMarshaler{}, Take: nil}
}

func (p *Player) SetID(id int) {
	p.id = id
}

func (p *Player) GetID() int {
	return p.id
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
