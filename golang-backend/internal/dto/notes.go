package dto

type GetNoteIn struct {
	NoteID int
	UserID int
}

type Question struct {
	ID     int    `pg:"id"`
	NoteID int    `pg:"note"`
	Body   string `pg:"body"`
	Answer string `pg:"answer"`
}

type GetNoteOut struct {
	Questions []Question
}

type RenameNoteIn struct {
	NoteID  int
	UserID  int
	NewName string
}

type DeleteNoteIn struct {
	NoteID int
	UserID int
}

type CreateQuestionIn struct {
	NoteID int
	UserID int
	Body   string
	Answer string
}
