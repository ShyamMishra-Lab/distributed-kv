package file

import "errors"

var ErrFileAlreadyOpen = errors.New("file already open :(")
var ErrFileNotOpen = errors.New("file handle expired or file not open :(")
