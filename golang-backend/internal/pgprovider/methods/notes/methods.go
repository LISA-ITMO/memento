package notes

import (
	"context"
	"memento/internal/dto"
	"memento/internal/pgprovider/queries"
)

func (n notesService) Get(ctx context.Context, in dto.GetNoteIn) (dto.GetNoteOut, error) {
	var questions []dto.Question
	_, err := n.db.Query(
		questions,
		queries.GetAllQuestionsByNoteID,
		in.NoteID,
		in.UserID,
	)
	out := dto.GetNoteOut{Questions: questions}
	return out, err
}

func (n notesService) CreateQuestion(ctx context.Context, in dto.CreateQuestionIn) error {
	_, err := n.db.ExecContext(
		ctx,
		queries.CreateQuestionInNoteByNoteID,
		in.NoteID,
		in.UserID,
		in.Body,
		in.Answer,
	)
	return err
}

func (n notesService) Rename(ctx context.Context, in dto.RenameNoteIn) error {
	_, err := n.db.ExecContext(
		ctx,
		queries.RenameNoteByID,
		in.NewName,
		in.NoteID,
		in.UserID,
	)
	return err
}

func (n notesService) Delete(ctx context.Context, in dto.DeleteNoteIn) error {
	_, err := n.db.ExecContext(
		ctx,
		queries.DeleteNoteByID,
		in.NoteID,
		in.UserID,
	)
	return err
}
