package openhash

import (
	"github.com/c0z0/gohash/utils"
)

type hashListItem struct {
	value utils.Value
	key   utils.Key
	next  *hashListItem
}

func (h *hashListItem) append(value utils.Value, key utils.Key) {
	if h.next == nil {
		newListItem := hashListItem{value: value, key: key}
		h.next = &newListItem

		return
	}

	h.next.append(value, key)
}

func (h *hashListItem) find(key utils.Key) *utils.Value {
	if h.key == key {
		return &h.value
	}

	if h.next != nil {
		return h.next.find(key)
	}

	return nil
}

func createList(value utils.Value, key utils.Key) *hashListItem {
	return &hashListItem{value: value, key: key}
}
