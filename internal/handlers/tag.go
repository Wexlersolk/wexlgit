package handlers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Wexlersolk/wexlgit/internal/handlers/utils"
)

func TagExecute(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Len of args must be 2: tag <user> <tag-name>")
	}

	_, tagName := args[0], args[1]

	wgitDir := filepath.Join(".wgit")
	if !utils.IsWgitRepository(wgitDir) {
		return fmt.Errorf("not a wgit repository, run wgit init first")
	}

	commit := filepath.Join(wgitDir, "refs/heads/main")
	data, err := os.ReadFile(commit)
	if err != nil {
		return fmt.Errorf("Error reading file: %w", err)
	}

	tagPath := filepath.Join(wgitDir, "refs/tags/", tagName)
	if err := os.WriteFile(tagPath, []byte(data), 0644); err != nil {
		return fmt.Errorf("Error creating tag file: %w", err)
	}
	fmt.Println("Success creating tag ", tagName)
	fmt.Println("Path to tag ", tagName, tagPath)

	return nil
}
