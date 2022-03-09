package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	assert := assert.New(t)
	array := New(10)
	array.Push(1)
	array.Push(2)
	array.Push(3)
	array.Push(4)
	array.Prepend(5)
	data, _ := array.At(2)
	assert.Equal(data, 2, "they should be equal")
	assert.Equal("51234", array.String(), "they should be equal")
	assert.Equal(10, array.Capacity(), "they should be equal")
	assert.Equal(false, array.IsEmpty(), "they should be equal")
	array.Pop()
	array.Pop()
	array.Pop()
	array.Pop()
	array.Pop()
	assert.Equal(true, array.IsEmpty(), "they should be equal")
}
