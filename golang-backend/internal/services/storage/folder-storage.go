package storage

import (
	"context"
	"memento/internal/dto"
)

type FolderStorage interface {
	GetMainFolder(ctx context.Context, in dto.GetMainFolderIn) (dto.GetMainFolderOut, error)
	GetFolder(ctx context.Context, in dto.GetFolderIn) (dto.GetFolderOut, error)
	CreateFolder(ctx context.Context, in dto.CreateFolderIn) error
	CreateNote(ctx context.Context, in dto.CreateNoteIn) error
	RenameFolder(ctx context.Context, in dto.RenameFolderIn) error
	DeleteFolder(ctx context.Context, in dto.DeleteFolderIn) error
}
