package openhash

import (
	"github.com/c0z0/gohash/utils"
)

// HashMap struct
type HashMap struct {
	size uint
	list []*hashListItem
}

// CreateHash Creates a map of size
func CreateHash(size uint) HashMap {
	pSize := utils.FindNextPrime(size)

	list := make([]*hashListItem, pSize)

	return HashMap{size: pSize, list: list}
}

// Put Saves an object to the map
func (h *HashMap) Put(value utils.Value, key utils.Key) {
	index := utils.Hash(key, h.size)

	if h.list[index] == nil {
		h.list[index] = createList(value, key)

		return
	}

	h.list[index].append(value, key)
}

// Get Returns the object from the map
func (h *HashMap) Get(key utils.Key) utils.Value {
	index := utils.Hash(key, h.size)

	return *h.list[index].find(key)
}
