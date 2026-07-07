package file

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// Helper functions
func newTestFileManager(t *testing.T) *FileManager {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")
	f, err := os.Create((dbPath))
	if err != nil {
		t.Fatalf("error on file creation: %v", err)
	}
	if err := f.Close(); err != nil {
		t.Fatalf("file closed, error reported: %v", err)
	}
	return NewFileManager(dbPath)
}

// Tests
func TestOpenExistingFile(t *testing.T) {
	//Arrange
	// dir := t.TempDir()
	// dbPath := filepath.Join(dir, "test.db")
	// f, err := os.Create(dbPath)
	// if err != nil {
	// 	t.Fatalf("error on file creation: %v", err)
	// }

	// if err := f.Close(); err != nil {
	// 	t.Fatalf("file closed, error reported: %v", err)
	// }

	// //Opening file manager
	// fm := NewFileManager(dbPath)

	fm := newTestFileManager(t)
	//closing the file manager at last
	//error if not close---The process cannot access the file because it is being used by another process.
	//it comes as the file inside the temp dir is not closed while windows try to clear temp dir
	defer func() {
		if err := fm.Close(); err != nil && err != ErrFileNotOpen {
			t.Fatalf("error on fm.Close: %v", err)
		}
	}()

	//Act
	err := fm.Open()

	//Assert
	if err != nil {
		t.Fatalf("error on fm.Open: %v", err)
	}

	if fm.file == nil {
		t.Fatalf("expected file handle to exist")
	}
}

func TestOpenAlreadyOpenedFileReturnsError(t *testing.T) {
	//Arrange
	fm := newTestFileManager(t)

	err := fm.Open()
	if err != nil {
		t.Fatalf("expected first open to work, got %v", err)
	}
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
	if !errors.Is(err, ErrFileAlreadyOpen) {
		t.Fatalf("expected: %v, got: %v", ErrFileAlreadyOpen, err)
	}
}

func TestCloseWithoutOpenReturnsError(t *testing.T) {
	//Arrange
	fm := newTestFileManager(t)

	//Act

	//Assertion
	if err := fm.Close(); !errors.Is(err, ErrFileNotOpen) {
		t.Fatalf("expected Close() to return ErrFileNotOpen, got: %v", err)
	}

}

func TestOpenClose(t *testing.T) {
	//Arrange
	fm := newTestFileManager(t)
	//Act

	//Assert
	if err := fm.Open(); err != nil && !errors.Is(err, ErrFileAlreadyOpen) {
		t.Fatalf("expected Open() to work, got: %v", err)
	}
	if err := fm.Close(); err != nil {
		t.Fatalf("expected Close() to work, got: %v", err)
	}
}

func TestOpenCloseOpen(t *testing.T) {
	//Arrange
	fm := newTestFileManager(t)

	//Closing file at the end so it won't interfere with the cleaning of TempDir()
	defer func() {
		if err := fm.Close(); err != nil && errors.Is(err, ErrFileNotOpen) {
			t.Fatalf("file closed, error reported: %v", err)
		}
	}()
	//Act

	//Assert
	if err := fm.Open(); err != nil && !errors.Is(err, ErrFileAlreadyOpen) {
		t.Fatalf("expected Open() to work, got: %v", err)
	}
	if err := fm.Close(); err != nil {
		t.Fatalf("expected Close() to work, got: %v", err)
	}
	if err := fm.Open(); err != nil {
		t.Fatalf("expected Open() to reopen file, got: %v", err)
	}
}

// Tests Create
func TestCreateNewFile(t *testing.T) {
	//Arrange
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")
	fm := NewFileManager(dbPath)

	//Closing file at the end of function so that the tempDir can be trashed
	defer func() {
		if err := fm.Close(); err != nil && err != ErrFileNotOpen {
			t.Fatalf("error on fm.Close: %v", err)
		}
	}()

	//Act/Assert
	if err := fm.Create(); err != nil {
		t.Fatalf("expected Create() to create and open file, got: %v", err)
	}
	if fm.file == nil {
		t.Fatalf("expected file handle to exist")
	}
}

func TestCreateExistingFileReturnsError(t *testing.T) {
	//Arrange
	//Creating and closing a file on path : dbPath
	dir := os.TempDir()
	dbPath := filepath.Join(dir, "test1.db")
	f, err := os.Create(dbPath)
	if err != nil {
		t.Fatalf("error on file creation: %v", err)
	}
	if err := f.Close(); err != nil {
		t.Fatalf("file closed, error reported: %v", err)
	}

	fm := NewFileManager(dbPath)

	//Act/Assert
	if err := fm.Create(); err != nil && !errors.Is(err, os.ErrExist) {
		t.Fatalf("expected: %v, got: %v", os.ErrExist, err)
	}
}

//Test ReadAt()
