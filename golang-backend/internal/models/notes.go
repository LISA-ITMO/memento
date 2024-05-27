package models

type Question struct {
	ID     int    `pg:"id"`
	NoteID int    `pg:"note"`
	Body   string `pg:"body"`
	Answer string `pg:"answer"`
}