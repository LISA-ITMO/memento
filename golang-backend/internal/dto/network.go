package dto

import "memento/internal/models"

type PublishNoteIn struct {
	UserID int
	NoteID int
	Name   string
	Descr  string
	Tags   []int
}

type GetPublicNotesByTagIn struct {
	Tag int
}

type GetPublicNotesByTagOut struct {
	Notes []models.PublicNote
}

type GetPublicNotesIn struct {
	ID int
}

type GetPublicNotesOut struct {
	Questions []models.Question
}

type DeletePublicNoteIn struct {
	NoteID int
	UserID int
}

type AddFavNoteIn struct {
}

type GetFavsIn struct {
}

type GetFavsOut struct {
}

type DeleteFavIn struct {
}

type GetTagsOut struct {
	Tags []models.Tag
}
