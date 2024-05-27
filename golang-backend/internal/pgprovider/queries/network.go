package queries

const (
	MakeNotePublishByID = `
	update notes set is_public = true 
	where id = ? and owner = ?;
`
	InsertPublicNoteByID = `
	insert into public_notes(id, author, name, descr) values 
	(?, ?, ?, ?);
`
	AddTagToNoteByID = `
	insert into tags_notes(note, tag) values 
	(?, ?);
`
	UnmakeNotePublicByID = `
	update notes set is_public = false 
	where id = ? and owner = ?;
`
	DeletePublicNoteByID = `
	delete from public_notes where id = ? and author = ?;
`

	GetPublicNotesListByTag = `
	select id, author, name, descr from public_notes 
	where id = (select note from tags_note where tag = ?);
`

	GetPublicNoteContentByID = `
	select id, note, body, answer from questions 
	where note = ?;
`
)
