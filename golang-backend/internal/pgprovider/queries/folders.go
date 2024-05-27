package queries

const (
	GetMainFolderFoldersByUserID = `
	select id, name 
	from folders 
	where parent = 
		(select folder 
		 from users_to_main_folders
		 where "user" = ?
		);
`
	GetMainFolderNotesByUserID = `
	select id, name 
	from notes 
	where folder = 
		(select folder 
	 	from users_to_main_folders
	 	where "user" = ?
		);
`
	GetFolderChildFolders = `
	select id, name 
	from folders 
	where parent = ? and owner = ?;
`
	GetFolderNotesByID = `
	select id, name 
	from notes 
	where folder = ? and owner = ?;
`
	DeleteFolderByID = `
	delete from folders
	where id = ? and owner = ?;
`
	CreateFolder = `
	INSERT INTO folders(name, owner, parent, created_at) VALUES
	(?, ?, ?, ?);
`
	RenameFolder = `
	UPDATE folders SET name = ? where id = ? and owner = ?;
`
	CreateNote = `
	INSERT INTO notes(name, owner, folder, created_at) VALUES
	(?, ?, ?, ?);
`
)
