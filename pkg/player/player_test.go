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
