package btree

import "errors"

var ErrKeyValCountMisMatch = errors.New("Key value count mismatch")
var ErrInvalidChildCount = errors.New("Corrupt node invalid Child count")
