package file

import (

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
		t.Fatalf("Error on file creation: %v", err)
	}
	
	if err := f.Close(); err != nil {
		t.Fatalf("failed to close temporary file: %v", err)
	}

	//Opening file manager
	fm := NewFileManager(dbPath)
	//closing the file manager at last
	//error if not close---The process cannot access the file because it is being used by another process.
	//it comes as the file inside the temp dir is not closed while windows try to clear temp dir
	defer func() {
		if err := fm.Close(); err != nil && err != ErrFileNotOpen {
			t.Fatalf("error on fm.Close: %v", err)
		}
	}()

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