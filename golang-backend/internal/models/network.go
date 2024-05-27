package models

type Tag struct {
	ID   int
	Name string
}

type PublicNote struct {
	ID    int    `pg:"ID"`
	Name  string `pg:"name"`
	Descr string `pg:"descr"`
}
