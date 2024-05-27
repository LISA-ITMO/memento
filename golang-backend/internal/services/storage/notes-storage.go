package storage

import (
	"context"
	"memento/internal/dto"
)

type NotesStorage interface {
	Get(ctx context.Context, in dto.GetNoteIn) (dto.GetNoteOut, error)
	CreateQuestion(ctx context.Context, in dto.CreateQuestionIn) error
	Rename(ctx context.Context, in dto.RenameNoteIn) error
	Delete(ctx context.Context, in dto.DeleteNoteIn) error
}
