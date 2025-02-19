package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Wexlersolk/wexlgit/internal/handlers"
)

func TestCreateDirectory(t *testing.T) {
	dir := t.TempDir()
	wgitDir := filepath.Join(dir, ".wgit")

	// Test creating the directory
	if err := handlers.CreateWgitDirectory(wgitDir); err != nil {
		t.Fatalf("createWgitDirectory failed: %v", err)
	}

	// Verify the directory exists
	if _, err := os.Stat(wgitDir); os.IsNotExist(err) {
		t.Fatalf("Directory %s was not created", wgitDir)
	}

	// Test creating the directory again (should fail)
	if err := handlers.CreateWgitDirectory(wgitDir); err == nil {
		t.Fatal("Expected an error when creating an existing directory, got nil")
	}

}
