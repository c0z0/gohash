package closedhash

import (
	"github.com/c0z0/gohash/utils"
)

type itemStatus uint8

// Item statuses
const (
	EMPTY      itemStatus = iota
	LEGITIMATE            = iota
)

type hashItem struct {
	status itemStatus
	value  utils.Value
	key    utils.Key
}

type HashMap struct {
	size  uint
	items []hashItem
}

func resolution(index uint) uint {
	return index * index
}

func (h *HashMap) Put(value utils.Value, key utils.Key) {
	index := utils.Hash(key, h.size)

	i := uint(1)

	for h.items[index].key != key && h.items[index].status == LEGITIMATE {
		index = (index + resolution(i)) % h.size
		i++
	}

	h.items[index].key = key
	h.items[index].value = value
	h.items[index].status = LEGITIMATE

}

func (h *HashMap) Get(key utils.Key) utils.Value {
	index := utils.Hash(key, h.size)

	i := uint(0)

	for h.items[index].key != key && h.items[index].status == LEGITIMATE {
		index = (index + resolution(i)) % h.size
		i++
	}

	return h.items[index].value
}

func CreateHash(size uint) HashMap {
	pSize := utils.FindNextPrime(size)
	h := HashMap{size: pSize, items: make([]hashItem, pSize)}

	return h
}
