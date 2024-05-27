package pgprovider

import (
	"context"
	"memento/internal/dto"
)

func (p *Provider) GetMainFolder(ctx context.Context, in dto.GetMainFolderIn) (dto.GetMainFolderOut, error) {
	return p.folderService.GetMainFolder(ctx, in)
}

func (p *Provider) GetFolder(ctx context.Context, in dto.GetFolderIn) (dto.GetFolderOut, error) {
	return p.folderService.GetFolder(ctx, in)
}

func (p *Provider) CreateFolder(ctx context.Context, in dto.CreateFolderIn) error {
	return p.folderService.CreateFolder(ctx, in)
}

func (p *Provider) CreateNote(ctx context.Context, in dto.CreateNoteIn) error {
	return p.folderService.CreateNote(ctx, in)
}

func (p *Provider) RenameFolder(ctx context.Context, in dto.RenameFolderIn) error {
	return p.folderService.RenameFolder(ctx, in)
}

func (p *Provider) DeleteFolder(ctx context.Context, in dto.DeleteFolderIn) error {
	return p.folderService.DeleteFolder(ctx, in)
}
