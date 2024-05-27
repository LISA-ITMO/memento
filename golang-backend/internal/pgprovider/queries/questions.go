package queries

const (
	DeleteQuestionsByID = `
	delete from questions where id = ? and owner = ?;
`
	UPDQuestionByID = `
	update questions set body = ?, answer = ?
	where id = ? and owner = ?;
`
)
