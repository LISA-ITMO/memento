package pgprovider

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
	"memento/config"
	"memento/internal/pgprovider/methods/folders"
	"memento/internal/services/storage"
)

type Provider struct {
	cfg *config.Config
	db  *pg.DB

	folderService storage.FolderStorage
}

func New(cfg *config.Config) storage.Storage {
	return &Provider{cfg: cfg, db: nil}
}

func (p *Provider) MustInit() {
	p.db = pg.Connect(&pg.Options{
		Addr:            p.cfg.Postgres.Host + ":" + p.cfg.Postgres.Port,
		User:            p.cfg.UserName,
		Password:        p.cfg.Password,
		Database:        p.cfg.DBName,
		MaxRetries:      3,
		MaxRetryBackoff: 3,
	})

	ctx := context.Background()

	if err := p.db.Ping(ctx); err != nil {
		log.Fatalf("Error with connecting to db: %v", err)
	}

	p.configureDependencies()
}

func (p *Provider) configureDependencies() {
	foldersManager := folders.New(p.db)
	p.folderService = foldersManager
}
