package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/Wexlersolk/wexlgit/internal/handlers/utils"
)

func CommitExecute(message string) error {
	wgitDir := filepath.Join(".wgit")
	if !utils.IsWgitRepository(wgitDir) {
		return fmt.Errorf("not a wgit repository, run wgit init first")
	}

	indexPath := filepath.Join(wgitDir, "index")
	if isEmpty, err := utils.IsFileEmpty(indexPath); err != nil {
		return fmt.Errorf("failed to check index file: %w", err)
	} else if isEmpty {
		return fmt.Errorf("nothing to commit, use wgit add")
	}

	treeID, err := utils.CreateTreeFromIndex(wgitDir)
	if err != nil {
		return fmt.Errorf("failed to create tree object: %w", err)
	}

	parrentCommitID, err := utils.GetCurrentCommit(wgitDir)
	if err != nil {
		return fmt.Errorf("failed to get parent commit: %w", err)
	}
	commitID, err := utils.CreateCommitObject(wgitDir, treeID, parrentCommitID, message)
	if err != nil {
		return fmt.Errorf("failed to create commit object: %w", err)
	}

	if err := utils.UpdateCurrentBranch(wgitDir, commitID); err != nil {
		return fmt.Errorf("failed to update branch: %w", err)
	}

	if err := utils.ClearStagingArea(wgitDir); err != nil {
		return fmt.Errorf("failed to clear staging area: %w", err)
	}

	fmt.Printf("Commit created: %s\n", &commitID)
	return nil

}
