package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func updateIndex(wgitDir string, filePath string, blobID string) error {
	indexPath := filepath.Join(wgitDir, "index")
	entry := fmt.Sprintf("%s %s\n", filePath, blobID)

	f, err := os.OpenFile(indexPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open index file: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(entry); err != nil {
		return fmt.Errorf("Failed to write to index file: %w", err)
	}
	return nil
}
