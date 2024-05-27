package questions

import (
	"github.com/go-pg/pg/v10"
	"memento/internal/services/storage"
)

type questionsService struct {
	db *pg.DB
}

func New(db *pg.DB) storage.QuestionsStorage {
	return questionsService{db: db}
}
