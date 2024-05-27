package models

type Folder struct {
	ID   int    `pg:"id"`
	Name string `pg:"name"`
}

type Note struct {
	ID   int    `pg:"id"`
	Name string `pg:"name"`
}
