package openhash

import (
	"github.com/c0z0/gohash/utils"
)

type hashListItem struct {
	value utils.Value
	key   utils.Key
	next  *hashListItem
}

func (h *hashListItem) Append(value utils.Value, key utils.Key) {
	if h.next == nil {
		newListItem := hashListItem{value: value, key: key}
		h.next = &newListItem

		return
	}

	h.next.Append(value, key)
}

func (h *hashListItem) Find(key utils.Key) *utils.Value {
	if h.key == key {
		return &h.value
	}

	if h.next != nil {
		return h.next.Find(key)
	}

	return nil
}

func CreateList(value utils.Value, key utils.Key) *hashListItem {
	return &hashListItem{value: value, key: key}
}
