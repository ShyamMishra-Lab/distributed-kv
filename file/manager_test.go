package file

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestOpenExistingFile(t *testing.T) {
	//Arrange
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")
	f, err := os.Create(dbPath)
	if err != nil {
		t.Fatalf("Error on file creation: %s", err)
	}
	
	if err := f.Close; err != nil {
		t.Fatalf("failed to close temporary file: %v", err)
	}

	//Opening file manager
	fm := NewFileManager(dbPath)

	//Act
	err = fm.Open()

	//Assert
	if err != nil {
		t.Fatalf("error on fm.Open: %v", err)
	}

	if fm.file == nil {
		t.Fatalf("expected fm.file to be not nil")
	}
}