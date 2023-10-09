package player

import "encoding/json"

type IPlayerTransport interface {
	Marshal(Player) ([]byte, error)
	UnMarshal([]byte, *Player) error
}

type JSONMarshaler struct{}

func (j JSONMarshaler) Marshal(p Player) ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return b, err
	}

	return b, nil
}

func (j JSONMarshaler) UnMarshal(bytes []byte, player *Player) (err error) {
	if err := json.Unmarshal(bytes, player); err != nil {
		return err
	}

	return nil
}
