package player

import (
	"pathe.co/zinx/gametake"
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
