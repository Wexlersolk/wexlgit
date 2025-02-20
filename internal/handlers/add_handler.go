package handlers

import (
	"crypto/sha256"
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

	for _, f := range files {

		path := filepath.Join(dir, f)
		_, err := os.Stat(path)
		if os.IsExist(err) {
			fmt.Printf("your file does not exist: %w", err)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}
		blobID := sha256.Sum256(data)

		newfile := filepath.Join(wgitDir, "file")

		if err := os.WriteFile(newfile, (blobID[:]), 0644); err != nil {
			return fmt.Errorf("failed to create config file: %w", err)
		}
	}

	return nil
}
