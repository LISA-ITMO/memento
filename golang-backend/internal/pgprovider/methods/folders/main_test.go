package folders

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
	"memento/internal/pgprovider/queries"
	"os"
	"testing"
	"time"
)

var testableService foldersService
var testFolderID int
var userID int

func createTestFolder() int {
	runningTime := time.Now()
	_, err := testableService.db.Exec(
		queries.CreateFolder,
		runningTime.String()[:50],
		userID,
		1,
		time.Now(),
	)
	if err != nil {
		log.Fatalf("error with creating folder: %v", err)
	}

	var id int
	_, err = testableService.db.Query(pg.Scan(&id), `
		select id from folders order by created_at DESC;
`)
	if err != nil {
		log.Fatalf("error with queriing folder: %v", err)
	}

	return id
}

func TestMain(m *testing.M) {
	testableService.db = pg.Connect(&pg.Options{
		Addr:            "localhost" + ":" + "5433",
		User:            "postgres",
		Password:        "1234",
		Database:        "memento",
		MaxRetries:      3,
		MaxRetryBackoff: 3,
	})

	ctx := context.Background()

	if err := testableService.db.Ping(ctx); err != nil {
		log.Fatalf("Error with connecting to db: %v", err)
	}
	userID = 1
	testFolderID = createTestFolder()

	os.Exit(m.Run())
}
