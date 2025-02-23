package utils

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type indexEntry struct {
	path   string
	blobID string
}

func CreateTreeFromIndex(wgitDir string) (string, error) {
	indexPath := filepath.Join(wgitDir, "index")
	entries, err := readIndexEntries(indexPath)
	if err != nil {
		return "", fmt.Errorf("failed to read index: %w", err)
	}

	tree := make(map[string][]indexEntry)
	for _, entry := range entries {
		dir := filepath.Dir(entry.path)
		tree[dir] = append(tree[dir], entry)
	}

	var rootTreeID string
	for dir, entries := range tree {
		treeID, err := createTreeObject(wgitDir, dir, entries)
		if err != nil {
			return "", fmt.Errorf("failed to create tree for %s: %w", dir, err)
		}

		if dir == "." {
			rootTreeID = treeID
		}
	}
	return rootTreeID, nil
}

func createTreeObject(wgitDir string, dir string, entries []indexEntry) (string, error) {
	var treeContent strings.Builder

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].path < entries[j].path
	})

	for _, entry := range entries {
		line := fmt.Sprintf("100644 blob %s\t%s\n", entry.blobID, filepath.Base(entry.path))
		treeContent.WriteString(line)
	}

	hash := sha1.Sum([]byte(treeContent.String()))
	treeID := fmt.Sprintf("%x", hash[:])

	treePath := filepath.Join(wgitDir, "objects", treeID)
	if err := os.WriteFile(treePath, []byte(treeContent.String()), 0644); err != nil {
		return "", fmt.Errorf("failed to store tree object: %w", err)
	}

	return treeID, nil
}

func readIndexEntries(indexPath string) ([]indexEntry, error) {
	data, err := os.ReadFile(indexPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read index file: %w", err)
	}

	var entries []indexEntry

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid index entry: %s", line)
		}

		entries = append(entries, indexEntry{path: parts[0], blobID: parts[1]})
	}

	return entries, nil

}
