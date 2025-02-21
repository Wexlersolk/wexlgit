package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentCommit(wgitDir string) (string, error) {
	headPath := filepath.Join(wgitDir, "HEAD")
	data, err := os.ReadFile(headPath)
	if err != nil {
		return "", fmt.Errorf("failed to read HEAD: %w", err)
	}

	ref := strings.TrimSpace(string(data))
	if !strings.HasPrefix(ref, "ref: ") {
		return ref, nil
	}

	refPath := filepath.Join(wgitDir, strings.TrimPrefix(ref, "ref: "))
	data, err = os.ReadFile(refPath)
	if err != nil {
		return "", fmt.Errorf("failed to read ref %s: %w", refPath, err)
	}

	return strings.TrimSpace(string(data)), nil
}

func UpdateCurrentBranch(wgitDir string, commitID string) error {
	headPath := filepath.Join(wgitDir, "HEAD")
	data, err := os.ReadFile(headPath)
	if err != nil {
		return fmt.Errorf("failed to read HEAD: %w", err)
	}

	ref := strings.TrimSpace(string(data))
	if !strings.HasPrefix(ref, "ref: ") {
		return fmt.Errorf("detatched HEAD is not supported")
	}

	refPath := filepath.Join(wgitDir, strings.TrimPrefix(ref, "ref: "))
	if err := os.WriteFile(refPath, []byte(commitID), 0644); err != nil {
		return fmt.Errorf("failed to update branch: %w", err)
	}

	return nil
}
