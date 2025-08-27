package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	s := &Stack{}
	assert.Equal(t, s.Pop(), 0)
	s.Push(1)
	assert.Equal(t, s.Pop(), 1)

	s.Push(2)
	s.Push(3)
	s.Push(4)
	assert.Equal(t, s.Pop(), 2)
}
