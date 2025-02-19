package handlers

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitRepository() error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}
	wgitDir := filepath.Join(dir, ".wgit")
	if err := createWgitDirectory(wgitDir); err != nil {
		return err
	}

	if err := createSubdirectories(wgitDir); err != nil {
		return nil
	}

	if err := createHeadFile(wgitDir); err != nil {
		return nil
	}

	if err := createConfigFile(wgitDir); err != nil {
		return nil
	}

	fmt.Printf("Initialized empty wgit repository at %s\n", wgitDir)
	return nil

}

func createWgitDirectory(wgitDir string) error {
	if err := os.Mkdir(wgitDir, 0755); err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("repository already exists at %s", wgitDir)
		}
		return fmt.Errorf("failed to create wgit directory: %w", err)
	}
	return nil
}

func createSubdirectories(wgitDir string) error {
	dirs := []string{"objects", "refs/heads", "refs/tags"}
	for _, d := range dirs {
		path := filepath.Join(wgitDir, d)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	return nil
}

func createHeadFile(wgitDir string) error {
	headFile := filepath.Join(wgitDir, "HEAD")
	if err := os.WriteFile(headFile, []byte("ref: refs/heads/main\n"), 0644); err != nil {
		return fmt.Errorf("failed to create HEAD file: %w", err)
	}
	return nil
}

func createConfigFile(wgitDir string) error {
	configFile := filepath.Join(wgitDir, "config")
	configContent := `[core] 
	repositoryformatversion = 0
	filemode = true
	bare = false
`
	if err := os.WriteFile(configFile, []byte(configContent), 0644); err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	return nil
}
