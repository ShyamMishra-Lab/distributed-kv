package file

import "os"

type FileManager struct {
	file *os.File
	path string
}

// only initializes the manager, doesn't open the file
func NewFileManager(path string) *FileManager {
	return  &FileManager{
		path: path,
	}
}
