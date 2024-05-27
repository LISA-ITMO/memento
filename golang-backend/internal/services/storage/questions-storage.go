package storage

import (
	"context"
	"memento/internal/dto"
)

type QuestionsStorage interface {
	Update(ctx context.Context, in dto.UPDQuestionIn) error
	Delete(ctx context.Context, in dto.DelQuestionIn) error
}
