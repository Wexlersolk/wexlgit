package utils

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func CreateCommitObject(wgitDir string, treeID string, parentCommitID string, message string) (string, error) {
	author := "Wexlersolk"
	committer := author
	timestamp := time.Now().Format(time.RFC3339)

	commitContent := fmt.Sprintf("tree %s\n", treeID)
	if parentCommitID != "" {
		commitContent += fmt.Sprintf("parent %s\n", parentCommitID)
	}
	commitContent += fmt.Sprintf("author %s %s\n", author, timestamp)
	commitContent += fmt.Sprintf("committer %s %s\n", committer, timestamp)
	commitContent += fmt.Sprintf("\n%s\n", message)

	hash := sha1.Sum([]byte(commitContent))
	commitID := fmt.Sprintf("%x", hash[:])

	commitPath := filepath.Join(wgitDir, "objects", commitID)
	if err := os.WriteFile(commitPath, []byte(commitContent), 0644); err != nil {
		return "", fmt.Errorf("failed to store commit object: %w", err)
	}

	return commitID, nil
}
