package dto

import (
	"memento/internal/models"
)

type GetMainFolderIn struct {
	UserID int
}

type GetMainFolderOut struct {
	Folders []models.Folder
	Notes   []models.Note
}

type GetFolderIn struct {
	ID     int
	UserID int
}

type GetFolderOut struct {
	Folders []models.Folder
	Notes   []models.Note
}

type CreateFolderIn struct {
	ParentFolderID int
	UserID         int
	Name           string
}

type CreateNoteIn struct {
	ParentFolderID int
	UserID         int
	Name           string
}

type RenameFolderIn struct {
	FolderID int
	UserID   int
	NewName  string
}

type DeleteFolderIn struct {
	FolderID int
	UserID   int
}
