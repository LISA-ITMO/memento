package folders

import (
	"context"
	"memento/internal/dto"
	"memento/internal/models"
	"memento/internal/pgprovider/queries"
	"time"
)

func (f foldersService) CreateFolder(ctx context.Context, in dto.CreateFolderIn) error {
	_, err := f.db.Exec(
		queries.CreateFolder,
		in.Name,
		in.UserID,
		in.ParentFolderID,
		time.Now(),
	)
	return err
}

func (f foldersService) CreateNote(ctx context.Context, in dto.CreateNoteIn) error {
	_, err := f.db.ExecContext(ctx, queries.CreateFolder,
		in.Name,
		in.UserID,
		in.ParentFolderID,
		time.Now())
	return err
}

func (f foldersService) DeleteFolder(ctx context.Context, in dto.DeleteFolderIn) error {
	_, err := f.db.Exec(queries.DeleteFolderByID, in.FolderID, in.UserID)
	return err
}

func (f foldersService) GetFolder(ctx context.Context, in dto.GetFolderIn) (dto.GetFolderOut, error) {
	tx, err := f.db.BeginContext(ctx)
	if err != nil {
		return dto.GetFolderOut{}, err
	}
	defer tx.Close()

	var folders []models.Folder
	_, err = tx.QueryContext(ctx,
		&folders,
		queries.GetFolderChildFolders,
		in.ID,
		in.UserID)
	if err != nil {
		return dto.GetFolderOut{}, err
	}

	var notes []models.Note
	_, err = tx.QueryContext(ctx,
		&notes,
		queries.GetFolderNotesByID,
		in.ID)
	if err != nil {
		return dto.GetFolderOut{}, err
	}

	result := dto.GetFolderOut{
		Folders: folders,
		Notes:   notes,
	}
	_ = tx.Commit()

	return result, nil
}

func (f foldersService) GetMainFolder(ctx context.Context, in dto.GetMainFolderIn) (dto.GetMainFolderOut, error) {
	// На вход получаем ID пользователя, для которого хотим получить гл. папку
	tx, err := f.db.BeginContext(ctx)
	if err != nil {
		return dto.GetMainFolderOut{}, err
	}
	defer tx.Close()

	var folders []models.Folder
	_, err = tx.QueryContext(ctx,
		&folders,
		queries.GetMainFolderFoldersByUserID,
		in.UserID)
	if err != nil {
		return dto.GetMainFolderOut{}, err
	}

	var notes []models.Note
	_, err = tx.QueryContext(ctx,
		&notes,
		queries.GetMainFolderNotesByUserID,
		in.UserID)
	if err != nil {
		return dto.GetMainFolderOut{}, err
	}

	result := dto.GetMainFolderOut{
		Folders: folders,
		Notes:   notes,
	}
	_ = tx.Commit()

	return result, nil
}

func (f foldersService) RenameFolder(ctx context.Context, in dto.RenameFolderIn) error {
	_, err := f.db.ExecContext(ctx, queries.RenameFolder,
		in.NewName,
		in.FolderID,
		in.UserID,
	)
	return err
}