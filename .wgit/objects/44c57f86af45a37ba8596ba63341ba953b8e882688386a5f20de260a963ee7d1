package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Wexlersolk/wexlgit/internal/handlers"
	"github.com/Wexlersolk/wexlgit/internal/handlers/utils"
)

// TestCommitExecute tests the CommitExecute function.
func TestCommitExecute(t *testing.T) {
	// Create a temporary directory for the test
	repoDir := t.TempDir()
	wgitDir := filepath.Join(repoDir, ".wgit")

	// Initialize the repository
	if err := os.MkdirAll(wgitDir, 0755); err != nil {
		t.Fatalf("Failed to create .wgit directory: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(wgitDir, "objects"), 0755); err != nil {
		t.Fatalf("Failed to create objects directory: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(wgitDir, "refs", "heads"), 0755); err != nil {
		t.Fatalf("Failed to create refs/heads directory: %v", err)
	}

	// Create an initial commit
	initialCommitID := "initial-commit-id"
	if err := os.WriteFile(filepath.Join(wgitDir, "HEAD"), []byte("ref: refs/heads/main\n"), 0644); err != nil {
		t.Fatalf("Failed to create HEAD file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(wgitDir, "refs", "heads", "main"), []byte(initialCommitID+"\n"), 0644); err != nil {
		t.Fatalf("Failed to create main branch file: %v", err)
	}

	// Create a staged file in the index
	indexPath := filepath.Join(wgitDir, "index")
	indexContent := "file1.txt blob1\nfile2.txt blob2\n"
	if err := os.WriteFile(indexPath, []byte(indexContent), 0644); err != nil {
		t.Fatalf("Failed to create index file: %v", err)
	}

	// Test the CommitExecute function
	commitMessage := "Test commit"
	if err := handlers.CommitExecute(commitMessage); err != nil {
		t.Fatalf("CommitExecute failed: %v", err)
	}

	// Verify the new commit
	headPath := filepath.Join(wgitDir, "HEAD")
	headContent, err := os.ReadFile(headPath)
	if err != nil {
		t.Fatalf("Failed to read HEAD file: %v", err)
	}
	ref := strings.TrimSpace(string(headContent))
	if !strings.HasPrefix(ref, "ref: ") {
		t.Fatalf("Expected HEAD to point to a ref, got: %s", ref)
	}

	refPath := filepath.Join(wgitDir, strings.TrimPrefix(ref, "ref: "))
	refContent, err := os.ReadFile(refPath)
	if err != nil {
		t.Fatalf("Failed to read branch file: %v", err)
	}
	newCommitID := strings.TrimSpace(string(refContent))
	if newCommitID == initialCommitID {
		t.Fatalf("Expected new commit ID, got initial commit ID: %s", newCommitID)
	}

	// Verify the commit object
	commitPath := filepath.Join(wgitDir, "objects", newCommitID)
	commitContent, err := os.ReadFile(commitPath)
	if err != nil {
		t.Fatalf("Failed to read commit object: %v", err)
	}
	if !strings.Contains(string(commitContent), commitMessage) {
		t.Fatalf("Commit message not found in commit object: %s", commitContent)
	}

	// Verify the staging area is cleared
	if isEmpty, err := utils.IsFileEmpty(indexPath); err != nil {
		t.Fatalf("Failed to check index file: %v", err)
	} else if !isEmpty {
		t.Fatal("Expected staging area to be cleared after commit")
	}
}

// TestCommitExecute_EmptyStagingArea tests the CommitExecute function with an empty staging area.
func TestCommitExecute_EmptyStagingArea(t *testing.T) {
	// Create a temporary directory for the test
	repoDir := t.TempDir()
	wgitDir := filepath.Join(repoDir, ".wgit")

	// Initialize the repository
	if err := os.MkdirAll(wgitDir, 0755); err != nil {
		t.Fatalf("Failed to create .wgit directory: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(wgitDir, "objects"), 0755); err != nil {
		t.Fatalf("Failed to create objects directory: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(wgitDir, "refs", "heads"), 0755); err != nil {
		t.Fatalf("Failed to create refs/heads directory: %v", err)
	}

	// Create an empty index file
	indexPath := filepath.Join(wgitDir, "index")
	if err := os.WriteFile(indexPath, []byte{}, 0644); err != nil {
		t.Fatalf("Failed to create index file: %v", err)
	}

	// Test the CommitExecute function
	commitMessage := "Test commit"
	err := handlers.CommitExecute(commitMessage)
	if err == nil {
		t.Fatal("Expected error for empty staging area, got nil")
	}
	if err.Error() != "nothing to commit (use 'wgit add' to stage changes)" {
		t.Fatalf("Expected 'nothing to commit' error, got: %v", err)
	}
}

// TestCommitExecute_NoRepository tests the CommitExecute function without a repository.
func TestCommitExecute_NoRepository(t *testing.T) {
	// Test the CommitExecute function
	commitMessage := "Test commit"
	err := handlers.CommitExecute(commitMessage)
	if err == nil {
		t.Fatal("Expected error for missing repository, got nil")
	}
	if err.Error() != "not a wgit repository (run 'wgit init' first)" {
		t.Fatalf("Expected 'not a wgit repository' error, got: %v", err)
	}
}
