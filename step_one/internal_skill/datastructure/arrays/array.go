package arrays

import (
	"bytes"
	"errors"
	"fmt"
)

type Arrays struct {
	data []int
	size int
}

func New(capacity int) *Arrays {
	return &Arrays{
		data: make([]int, capacity),
	}
}

// 数组元素的个数
func (a *Arrays) Size() int {
	return a.size
}

// 可容纳元素的个数
func (a *Arrays) Capacity() int {
	return len(a.data)
}

func (a *Arrays) IsEmpty() bool {
	return a.size == 0
}

// 返回对应索引的元素，且若索引越界则愤然报错
func (a *Arrays) At(index int) (e int, err error) {
	if index < 0 || index >= a.size {
		return 0, errors.New("index out of range")
	}
	return a.data[index], nil
}

func (a *Arrays) Insert(index int, item int) error {
	if index < 0 || index > a.size {
		return errors.New("index out of range")
	}
	if a.size == len(a.data) {
		a.resize(2 * a.size)
	}
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = item
	a.size++
	return nil
}

func (a *Arrays) Push(item int) error {
	return a.Insert(a.size, item)
}
func (a *Arrays) Prepend(item int) error {
	return a.Insert(0, item)
}

func (a *Arrays) resize(size int) {
	newslice := make([]int, size)
	copy(newslice[:a.size], a.data)
	a.data = newslice
}

func (a *Arrays) Delete(index int) (int, error) {
	if index < 0 || index >= a.size {
		return 0, errors.New("index out of range")
	}
	e := a.data[index]
	for i := index; i < a.size-1; i++ {
		a.data[index] = a.data[index+1]
	}
	a.size--
	a.data[a.size] = 0
	if a.size == len(a.data)/2 {
		a.resize(len(a.data) / 2)
	}
	return e, nil
}

func (a *Arrays) Pop() (int, error) {
	return a.Delete(0)
}

//  删除指定值的元素，并返回其索引（即使有多个元素）
func (a *Arrays) Remove(item int) []int {
	var res []int
	if a.Find(item) == -1 {
		return res
	}
	for i := 0; i < a.size; i++ {
		if a.data[i] == item {
			res = append(res, i)
			a.Delete(i)
		}
	}

	return res
}
func (a *Arrays) Find(item int) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == item {
			return i
		}
	}
	return -1
}

func (a *Arrays) String() string {
	var buf bytes.Buffer
	for i := 0; i < a.size; i++ {
		buf.WriteString(fmt.Sprintf("%d", a.data[i]))
	}
	return buf.String()
}
