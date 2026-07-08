package pager

import (
	"go/constant"

	"github.com/iips-oss/distributed-kv/file"
)

type Pager struct {
	file     *file.FileManager
	pageSize int
}

func NewPager(fm *file.FileManager, pageSize int) (*Pager, error) {
	if pageSize < 0 {
		return nil, ErrInvalidPageSize
	}
	return &Pager{
		file:     fm,
		pageSize: pageSize,
	}, nil
}

//Page I/O operations

//calc page byte offset
//allocate a buffer of one page
//read bytes from file manager into buffer
//create and returnn the page from the buffer
func (p *Pager) ReadPage(pageId uint64) (*Page, error) {
	offset := int64(pageId) * int64(p.pageSize)
	buffer := make([]byte, p.pageSize)

	if err := p.file.ReadAt(buffer, offset); err != nil {
		return nil, err
	}
	return &Page{
		ID:    pageId,
		Data:  buffer,
		Dirty: false,
	}, nil
}

//recieve page
//Calc page offset
//write page.data to filemanager---p.file
//
func (p *Pager) WritePage(page *Page) error {
	offset := int64(page.ID) * int64(p.pageSize)
	if err := p.file.WriteAt(page.Data, offset); err != nil {
		return err
	}
	page.Dirty = false
	return nil
}
