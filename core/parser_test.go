package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	l0, l1, err := baseParser0([]byte("a\r\nb\r\n\r\nc"))
	assert.Nil(t, err)
	assert.Equal(t, l1, []byte("c"))
	assert.Equal(t, l0[0], []byte("a"))
	assert.Equal(t, l0[1], []byte("b"))

	l0, l1, err = baseParser0([]byte("a\r\nb\r\nc\r\n"))
	assert.Nil(t, l0)
	assert.Nil(t, l1)
	assert.Error(t, err)
}
