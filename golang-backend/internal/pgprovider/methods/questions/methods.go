package questions

import (
	"context"
	"memento/internal/dto"
	"memento/internal/pgprovider/queries"
)

func (q questionsService) Update(ctx context.Context, in dto.UPDQuestionIn) error {
	_, err := q.db.ExecContext(
		ctx,
		in.Text,
		in.Answer,
		in.QuestionID,
		in.UserID,
	)
	return err
}

func (q questionsService) Delete(ctx context.Context, in dto.DelQuestionIn) error {
	_, err := q.db.ExecContext(
		ctx,
		queries.DeleteQuestionsByID,
		in.QuestionID,
		in.UserID,
	)
	return err
}
