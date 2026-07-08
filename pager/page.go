package pager

type Page struct {
	ID    uint64
	Data  []byte
	Dirty bool
}
