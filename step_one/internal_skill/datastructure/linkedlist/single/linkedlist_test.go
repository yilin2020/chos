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
	l.Erase(2)
	assert.Equal("23", l.String(), "should be equal")
	assert.Equal(3, l.ValueNFromEnd(1), "should be equal")
	l.Erase(1)
	l.Erase(1)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	assert.Equal("12345", l.String(), "should be equal")
	l.Reverse()
	assert.Equal("54321", l.String(), "should be equal")
	l.RemoveValue(4)
	assert.Equal("5321", l.String(), "should be equal")
}
