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

// Create: check if fm.file exists -> yes: return error ErrFileAlreadyOpen, -> no: Create file at path -> store file handle with flags: read/write return success
func (fm *FileManager) Create() error {
	if fm.file != nil {
		return ErrFileAlreadyOpen
	}
	// creates and open file exclusively i.e. it will give error if file already exists
	f, err := os.OpenFile(fm.path, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	// saves the file handle
	fm.file = f
	return nil
}

// // OpenOrCreate: check if file exists -> yes: fm.Open() -> no:  fm.Open() -> return errors

// func (fm *FileManager) OpenOrCreate() error {
// 	if err := fm.Create(); err != nil && errors.Is(err, os.ErrExist) {
// 		if err := fm.Open(); err != nil && !errors.Is(err, ErrFileAlreadyOpen) {
// 			return err
// 		}
// 	}
// 	return nil
// }

func (fm *FileManager) ReadAt(buf []byte, offset int64) error {
	if fm.file == nil {
		return ErrFileNotOpen
	}
	_, err := fm.file.ReadAt(buf, offset)
	return err
}

func (fm *FileManager) WriteAt(buf []byte, offset int64) error {
	if fm.file == nil {
		return ErrFileNotOpen
	}
	_, err := fm.file.WriteAt(buf, offset)
	return err
}

func (fm *FileManager) Sync() error {
	if fm.file == nil {
		return ErrFileNotOpen
	}
	return fm.file.Sync()
}

// Size needs file Info/ name provided by Stat()
func (fm *FileManager) Size() (int64, error) {
	if fm.file == nil {
		return 0, ErrFileNotOpen
	}
	info, err := fm.file.Stat()
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}


