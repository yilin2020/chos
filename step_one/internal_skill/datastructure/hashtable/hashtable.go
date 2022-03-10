package hashtable

import (
	"github.com/yilin2020/chos/step_one/internal_skill/datastructure/hashtable/hash"
)

type Hashtable struct {
	items     []*Item
	hashCount int
	size      int
}

type Item struct {
	key   string
	value string
	next  *Item
}

func genratorItem(key, value string) *Item {
	return &Item{
		key:   key,
		value: value,
	}
}

func newHashTable(hashCount int) *Hashtable {
	return &Hashtable{
		items:     make([]*Item, hashCount),
		hashCount: hashCount,
	}
}

func HashTable() *Hashtable {
	return newHashTable(1024)
}

func (h *Hashtable) Set(key, value string) {
	hashCode := hash.BkdrHash([]byte(key))
	index := hashCode % uint(h.hashCount)
	if h.items[index] == nil {
		h.items[index] = genratorItem(key, value)
		h.size++
	} else {
		cur := h.items[index]
		for cur.next != nil {
			if cur.key == key {
				cur.value = value
				return
			}
			cur = cur.next
		}
		if cur.key == key {
			cur.value = value
		} else {
			cur.next = genratorItem(key, value)
			h.size++
		}
	}
}

func (h *Hashtable) Get(key string) (v string, ok bool) {
	hashCode := hash.BkdrHash([]byte(key))
	index := hashCode % uint(h.hashCount)
	if h.items[index] == nil {
		return "", false
	} else {
		cur := h.items[index]
		for cur != nil {
			if cur.key == key {
				return cur.value, true
			}
			cur = cur.next
		}
	}
	return "", false
}
