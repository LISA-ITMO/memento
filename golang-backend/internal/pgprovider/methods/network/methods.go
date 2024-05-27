package network

import (
	"context"
	"memento/internal/dto"
	"memento/internal/models"
	"memento/internal/pgprovider/queries"
)

// TODO: Система добавления тэгов!!!!!

func (n networkService) PublishNote(ctx context.Context, in dto.PublishNoteIn) error {
	tx, err := n.db.BeginContext(ctx)
	if err != nil {
		return err
	}
	defer tx.Close()

	_, err = tx.ExecContext(
		ctx,
		queries.MakeNotePublishByID,
		in.NoteID,
		in.UserID,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		queries.InsertPublicNoteByID,
		in.NoteID,
		in.UserID,
		in.Name,
		in.Descr,
	)
	if err != nil {
		return err
	}

	for _, tag := range in.Tags {
		_, err = tx.ExecContext(ctx, queries.AddTagToNoteByID, in.NoteID, tag)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (n networkService) GetPublicNotesByTag(ctx context.Context, in dto.GetPublicNotesByTagIn) (dto.GetPublicNotesByTagOut, error) {
	var result []models.PublicNote
	_, err := n.db.QueryContext(
		ctx,
		result,
		queries.GetPublicNotesListByTag,
		in.Tag,
	)
	if err != nil {
		return dto.GetPublicNotesByTagOut{}, err
	}

	out := dto.GetPublicNotesByTagOut{Notes: result}

	return out, nil
}

func (n networkService) GetPublicNote(ctx context.Context, in dto.GetPublicNotesIn) (dto.GetPublicNotesOut, error) {
	var result []models.Question
	_, err := n.db.QueryContext(
		ctx,
		result,
		queries.GetPublicNoteContentByID,
		in.ID,
		)
	if err != nil {
		return dto.GetPublicNotesOut{}, err
	}

	return dto.GetPublicNotesOut{Questions: result}, nil
}

func (n networkService) DeletePublicNote(ctx context.Context, in dto.DeletePublicNoteIn) error {
	tx, err := n.db.BeginContext(ctx)
	if err != nil {
		return err
	}
	defer tx.Close()

	_, err = tx.ExecContext(ctx, queries.UnmakeNotePublicByID, in.NoteID, in.UserID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queries.DeletePublicNoteByID, in.NoteID, in.UserID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (n networkService) AddFavourite(ctx context.Context, in dto.AddFavNoteIn) error {
	return nil
}

func (n networkService) GetFavourites(ctx context.Context, in dto.GetFavsIn) (dto.GetFavsOut, error) {
	return dto.GetFavsOut{}, nil
}
