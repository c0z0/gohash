package openhash

import (
	"github.com/c0z0/gohash/utils"
)

type HashMap struct {
	size uint
	list []*hashListItem
}

func CreateHash(size uint) HashMap {
	pSize := utils.FindNextPrime(size)

	list := make([]*hashListItem, pSize)

	return HashMap{size: pSize, list: list}
}

func (h *HashMap) Put(value utils.Value, key utils.Key) {
	index := utils.Hash(key, h.size)

	if h.list[index] == nil {
		h.list[index] = CreateList(value, key)

		return
	}

	h.list[index].Append(value, key)
}

func (h *HashMap) Get(key utils.Key) utils.Value {
	index := utils.Hash(key, h.size)

	return *h.list[index].Find(key)
}
