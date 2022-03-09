package single

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedlist(t *testing.T) {
	assert := assert.New(t)
	l := New()
	l.PushFront(1)
	l.PushFront(2)
	l.PushBack(3)
	assert.Equal("213", l.String(), "should be equal")
}
