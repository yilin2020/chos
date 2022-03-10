package ringbuffer

import (
	"errors"
	"sync"
)

type RingBuffer struct {
	buf    []byte
	mu     sync.Mutex
	r      int
	w      int
	isFull bool
	size   int
}

func New(size int) *RingBuffer {
	return &RingBuffer{
		buf:  make([]byte, size),
		size: size,
	}
}

func (r *RingBuffer) Read(data []byte) (n int, err error) {
	if len(data) == 0 {
		return 0, nil
	}
	r.mu.Lock()
	if !r.isFull && r.w == r.r {
		r.mu.Unlock()
		return 0, errors.New("ringbuffer is empty")
	}
	if r.w-r.r > 0 {
		n = r.w - r.r
		if n > len(data) {
			n = len(data)
		}
		copy(data, r.buf[r.r:r.r+n])
		r.r = (r.r + n) % r.size
		r.mu.Unlock()
		return
	}
	n = r.size - r.r + r.w
	if n > len(data) {
		n = len(data)
	}
	if r.r+n <= r.size {
		copy(data, r.buf[r.r:r.r+n])
	} else {
		first := r.size - r.r
		copy(data, r.buf[r.r:])
		send := n - first
		copy(data[first:], r.buf[:send])
	}
	r.r = (r.r + n) % r.size
	r.isFull = false
	r.mu.Unlock()
	return
}

func (r *RingBuffer) Write(data []byte) (n int, err error) {
	if len(data) == 0 {
		return 0, nil
	}

	r.mu.Lock()
	if r.isFull {
		r.mu.Unlock()
		return 0, errors.New("ringbuffer is full")
	}

	var remainByte int
	if r.w-r.r >= 0 {
		remainByte = r.size - r.w + r.r
	} else {
		remainByte = r.r - r.w
	}

	if len(data) > remainByte {
		r.mu.Unlock()
		return 0, errors.New("ringbuffer no space to write")
	}
	n = len(data)
	if r.w-r.r >= 0 {
		first := r.size - r.w
		copy(r.buf[r.w:], data[:first])
		second := n - first
		if second > 0 {
			copy(r.buf[:second], data[first:])
			r.w += second
		}
		r.w += n
	} else {
		copy(r.buf[r.w:r.r], data)
		r.w += n
	}

	if r.w == r.size {
		r.w = 0
	}
	if r.w == r.r {
		r.isFull = true
	}
	r.mu.Unlock()

	return
}
