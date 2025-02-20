package handlers

import (
	"fmt"
	"os"
)

func isWgitRepository(wgitDir string) bool {
	_, err := os.Stat(wgitDir)
	return !os.IsNotExist(err)
}

func validateFilePath(path string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	if info.IsDir() {
		return fmt.Errorf("cannot add directory YET")
	}
	return nil
}

func addFileToStagingArea(wgitDir, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	blobID := createBlob(data)
	if err := storeBlob(wgitDir, blobID, data); err != nil {
		return err
	}
	if err := updateIndex(wgitDir, filePath, blobID); err != nil {
		return err
	}
	return nil
}
