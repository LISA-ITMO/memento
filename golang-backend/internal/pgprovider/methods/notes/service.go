package notes

import (
	"github.com/go-pg/pg/v10"
	"memento/internal/services/storage"
)

type notesService struct {
	db *pg.DB
}

func New(db *pg.DB) storage.NotesStorage {
	return notesService{db: db}
}
