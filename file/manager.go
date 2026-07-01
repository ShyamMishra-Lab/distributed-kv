package file

import (
	"os"
)

type FileManager struct {
	file *os.File
	path string
}

// only initializes the manager, doesn't open the file
// takes path of the *os.file
func NewFileManager(path string) *FileManager {
	return &FileManager{
		path: path,
	}
}

// Open: check if file already open -> yes, return error -> no, open file for path -> save file pointer to file manager -> Return success
//
// opens the managed file for read and write, error if file already open, or doesn't open
func (fm *FileManager) Open() error {
	if fm.file != nil {
		return ErrFileAlreadyOpen
	}

	f, err := os.OpenFile(fm.path, os.O_RDWR, 0)
	if err != nil {
		return err
	}

	fm.file = f
	return nil
}

// Close: check if file open => no, return error -> yes, call os.Close() -> Store the err -> set fm.file to nil -> return err
// on close the file handle becomes unusable
// close the file if opened, if not open, ErrFileNotOpen,
func (fm *FileManager) Close() error {
	if fm.file == nil {
		return ErrFileNotOpen
	}

	err := fm.file.Close()

	// the handle is no longer usable
	fm.file = nil
	return err
}
