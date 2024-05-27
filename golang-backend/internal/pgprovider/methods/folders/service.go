package folders

import (
	"github.com/go-pg/pg/v10"
	"memento/internal/services/storage"
)

type foldersService struct {
	db *pg.DB
}

func New(db *pg.DB) storage.FolderStorage {
	return foldersService{db: db}
}
