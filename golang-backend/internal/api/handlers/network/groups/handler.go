package groups

import n "memento/internal/services/storage/network"

type Handler struct {
	s n.GroupsStorage
}

func New(storage n.GroupsStorage) *Handler {
	return &Handler{
		s: storage,
	}
}
