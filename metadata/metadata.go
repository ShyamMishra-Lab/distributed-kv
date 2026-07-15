package metadata

type Metadata struct {
	// to validate if it's our file
	Magic        [4]byte
	Version      uint32
	PageSize     uint32
	RootPage     uint64
	NextFreePage uint64
}
// Constants
var MagicNumber = [4]byte{'D', 'K', 'V', '1'}

const CurrentVersion uint32 = 1

// Constructors
func NewMetadata(pageSize uint32) *Metadata {
	return &Metadata{
		Magic:        MagicNumber,
		Version:      CurrentVersion,
		PageSize:     pageSize,
		RootPage:     1,
		NextFreePage: 2,
	}
}
