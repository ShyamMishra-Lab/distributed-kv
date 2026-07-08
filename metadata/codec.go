package metadata

import (
	"bytes"
	"encoding/binary"
)

var byteOrder = binary.LittleEndian

//Create a buffer -> write data -> return bytes
func (m *Metadata) Encode() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, byteOrder, m.Magic); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, byteOrder, m.Version); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, byteOrder, m.PageSize); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, byteOrder, m.RootPage); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, byteOrder, m.NextFreePage); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//create metaddata -> read -> return metadata
func Decode(data []byte) (*Metadata, error) {
	reader := bytes.NewReader(data)
	m := &Metadata{}

	if err := binary.Read(reader, byteOrder, &m.Magic); err != nil {
		return nil, err
	}
	if m.Magic != MagicNumber {
		return nil, ErrInvalidMagicNumber
	}
	if err := binary.Read(reader, byteOrder, &m.Version); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, byteOrder, &m.PageSize); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, byteOrder, &m.RootPage); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, byteOrder, &m.NextFreePage); err != nil {
		return nil, err
	}
	return m, nil
}
