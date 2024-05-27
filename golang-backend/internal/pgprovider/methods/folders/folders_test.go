package folders

import (
	"context"
	"github.com/stretchr/testify/require"
	"memento/internal/dto"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	in := dto.CreateFolderIn{
		ParentFolderID: testFolderID,
		UserID:         userID,
		Name:           "TestCreateFolder",
	}

	err := testableService.CreateFolder(context.Background(), in)

	require.NoError(t, err)
}

func TestCreateNote(t *testing.T) {
	in := dto.CreateNoteIn{
		ParentFolderID: testFolderID,
		UserID:         userID,
		Name:           "TestNote",
	}
	err := testableService.CreateNote(context.Background(), in)

	require.NoError(t, err)
}
