package questions

import "memento/internal/services/storage"

type Handler struct {
	s storage.QuestionsStorage
}

func New(storage storage.QuestionsStorage) *Handler {
	return &Handler{s: storage}
}
