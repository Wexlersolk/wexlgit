package handlers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Wexlersolk/wexlgit/internal/handlers/utils"
)

func AddExecute(paths []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	wgitDir := filepath.Join(dir, ".wgit")
	if !utils.IsWgitRepository(wgitDir) {
		return fmt.Errorf("not a wgit repository (run 'wgit init' first)")
	}

	for _, path := range paths {
		absPath := filepath.Join(dir, path)
		if err := utils.AddPathToStagingArea(wgitDir, absPath); err != nil {
			return fmt.Errorf("failed to add %s: %w", path, err)
		}
		fmt.Printf("Added %s\n", path)
	}
	return nil
}
