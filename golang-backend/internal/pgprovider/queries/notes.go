package queries

const (
	GetAllQuestionsByNoteID = `
	select id, note, body, answer from questions 
	where note = ? and owner = ?;
`
	CreateQuestionInNoteByNoteID = `
	insert into questions(note, owner, body, answer) values 
	(?, ?, ?, ?);
`
	RenameNoteByID = `
	UPDATE notes SET name = ? where id = ? and owner = ?;
`
	DeleteNoteByID = `
	delete from notes
	where id = ? and owner = ?
`
)
