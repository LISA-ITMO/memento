package network

import "memento/internal/services/storage"

type Handler struct {
	s storage.NetworkStorage
}

func New(storage storage.NetworkStorage) *Handler {
	return &Handler{s: storage}
}
