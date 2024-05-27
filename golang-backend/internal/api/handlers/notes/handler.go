package notes

import "memento/internal/services/storage"

type Handler struct {
	s storage.NotesStorage
}

func New(s storage.NotesStorage) Handler {
	return Handler{
		s: s,
	}
}
