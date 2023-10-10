package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pathe.co/zinx/gametake"
)

func TestSetTake(t *testing.T) {
	p := NewPlayer()
	p.SetTake(&gametake.TOUT)
	assert.Equal(t, p.Take, gametake.TOUT)
}

func TestNewPlayer(t *testing.T) {
	p := NewPlayer()
	assert.Equal(t, p.Transport, JSONMarshaler{})
}

func TestSetID(t *testing.T) {
	p := NewPlayer()
	p.SetID(0)
	assert.Equal(t, p.id, 0)
}
