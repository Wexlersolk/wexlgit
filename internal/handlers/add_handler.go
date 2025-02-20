package handlers

import (
	"fmt"
	"os"
	"path/filepath"
)

func AddExecute(files []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	wgitDir := filepath.Join(dir, ".wgit")
	if !isWgitRepository(wgitDir) {
		return fmt.Errorf("not a wgit repository (run 'wgit init' first)")
	}

	for _, file := range files {
		path := filepath.Join(dir, file)
		if err := validateFilePath(path); err != nil {
			return fmt.Errorf("invalid file %s: %w", file, err)
		}
		if err := addFileToStagingArea(wgitDir, path); err != nil {
			return fmt.Errorf("failed to add file %s: %w", file, err)
		}
		fmt.Printf("Added %s\n", file)
	}
	return nil
}
