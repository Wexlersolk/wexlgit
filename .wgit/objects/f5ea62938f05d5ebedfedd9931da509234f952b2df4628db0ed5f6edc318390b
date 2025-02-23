package utils

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

func createBlob(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash[:])
}

func storeBlob(wgitDir string, blobID string, data []byte) error {
	blobPath := filepath.Join(wgitDir, "objects", blobID)
	if err := os.WriteFile(blobPath, data, 0644); err != nil {
		return fmt.Errorf("failed to store blob: %w", err)
	}
	return nil
}
