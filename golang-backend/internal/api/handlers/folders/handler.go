package folders

import (
	"memento/internal/services/storage"
)

type Handler struct {
	s storage.FolderStorage
}

func New(storage storage.FolderStorage) *Handler {
	return &Handler{s: storage}
}
