package player

import (
	"encoding/json"
	"fmt"
)

type IPlayerTransport interface {
	Marshal(p Player) ([]byte, error)
	UnMarshal(d []byte, p *Player) error
}

type JSONMarshaler struct{}

func (j JSONMarshaler) Marshal(p Player) ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return b, fmt.Errorf("json marshal err %w", err)
	}

	return b, nil
}

func (j JSONMarshaler) UnMarshal(bytes []byte, player *Player) error {
	if err := json.Unmarshal(bytes, player); err != nil {
		return fmt.Errorf("json unmarshal err %w", err)
	}

	return nil
}
