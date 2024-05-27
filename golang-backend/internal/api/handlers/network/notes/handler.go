package notes

import n "memento/internal/services/storage/network"

type Handler struct {
	s n.NotesStorage
}

func New(storage n.NotesStorage) *Handler {
	return &Handler{
		s: storage,
	}
}
