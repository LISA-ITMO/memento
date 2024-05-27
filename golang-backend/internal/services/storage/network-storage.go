package storage

import (
	"context"
	"memento/internal/dto"
)

type NetworkStorage interface {
	GetTags(ctx context.Context) (dto.GetTagsOut, error)
	PublishNote(ctx context.Context, in dto.PublishNoteIn) error
	GetPublicNote(ctx context.Context, in dto.GetPublicNotesIn) (dto.GetPublicNotesOut, error)
	DeletePublicNote(ctx context.Context, in dto.DeletePublicNoteIn) error
	AddFavourite(ctx context.Context, in dto.AddFavNoteIn) error
	GetFavourites(ctx context.Context, in dto.GetFavsIn) (dto.GetFavsOut, error)
}
