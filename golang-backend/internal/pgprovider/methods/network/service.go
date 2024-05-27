package network

import (
	"github.com/go-pg/pg/v10"
	"memento/internal/dto"
	"memento/internal/services/storage"
)

type networkService struct {
	db *pg.DB
}

func New(db *pg.DB) storage.NetworkStorage {
	return networkService{db: db}
}
