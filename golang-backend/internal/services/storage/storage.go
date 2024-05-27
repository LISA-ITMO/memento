package storage

type Storage interface {
	MustInit()
	FolderStorage
}
