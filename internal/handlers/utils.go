package handlers

import (
	"fmt"
	"os"
	"path/filepath"
)

func isWgitRepository(wgitDir string) bool {
	_, err := os.Stat(wgitDir)
	return !os.IsNotExist(err)
}

func validateFilePath(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	return nil
}

func addPathToStagingArea(wgitDir string, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to stat path: %w", err)
	}
	if info.IsDir() {
		return filepath.Walk(path, func(filePath string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("failed to walk directory: %w", err)

			}

			if !fileInfo.IsDir() {
				if err := addFileToStagingArea(wgitDir, filePath); err != nil {
					return fmt.Errorf("failed to add file %s: %w", filePath, err)
				}

			}
			return nil
		})
	}
	return addFileToStagingArea(wgitDir, path)
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
