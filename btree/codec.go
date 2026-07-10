// TODO: reader.Read() to io.ReadFull()
// TODO: validate key value sizes for uint16
package btree

import (
	"bytes"
	"encoding/binary"
)

var byteOrder = binary.LittleEndian

func (n *Node) EncodeNode() ([]byte, error) {
	// Is data valid
	if n.IsLeaf && len(n.Keys) != len(n.Values) {
		return nil, ErrKeyValCountMisMatch
	}
	if !n.IsLeaf && len(n.Keys)+1 != len(n.Children) {
		return nil, ErrInvalidChildCount
	}

	//in mem byte array buffer
	var buf bytes.Buffer

	// Writing Node identity
	if err := binary.Write(&buf, byteOrder, n.IsLeaf); err != nil {
		return nil, err
	}
	// Writing Number of keys
	keyCount := uint16(len(n.Keys))
	if err := binary.Write(&buf, byteOrder, keyCount); err != nil {
		return nil, err
	}

	// Writing key length and key data in sequence
	for _, key := range n.Keys {
		keyLen := uint16(len(key))
		// write key len
		if err := binary.Write(&buf, byteOrder, keyLen); err != nil {
			return nil, err
		}
		// write key data
		if _, err := buf.Write(key); err != nil {
			return nil, err
		}
	}

	// Encode based on node type
	if n.IsLeaf {
		// Writing value length and value data in sequence
		for _, value := range n.Values {
			valLen := uint16(len(value))
			// Write value length
			if err := binary.Write(&buf, byteOrder, valLen); err != nil {
				return nil, err
			}
			// Write value data
			if _, err := buf.Write(value); err != nil {
				return nil, err
			}
		}
	} else {
		// Writing Children values and these are of fixed length
		for _, child := range n.Children {
			if err := binary.Write(&buf, byteOrder, child); err != nil {
				return nil, err
			}
		}
	}

	return buf.Bytes(), nil
}

// Decode function for node
func DecodeNode(data []byte) (*Node, error) {
	// Reader for a new node to write to and return
	reader := bytes.NewReader(data)
	n := &Node{}

	// Reading IsLeaf
	if err := binary.Read(reader, byteOrder, &n.IsLeaf); err != nil {
		return nil, err
	}

	// Reading keyCount
	var keyCount uint16
	if err := binary.Read(reader, byteOrder, &keyCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(keyCount); i++ {
		// Readding KeyLength
		var keyLen uint16
		if err := binary.Read(reader, byteOrder, &keyLen); err != nil {
			return nil, err
		}

		//Reading key
		key := make([]byte, keyLen)
		// binary.Read(reader, byteOrder, &key)
		if _, err := reader.Read(key); err != nil {
			return nil, err
		}
		// Saving the keys from buffer/reader
		n.Keys = append(n.Keys, key)
	}

	// Decode based on node type
	if n.IsLeaf {
		for i := 0; i < int(keyCount); i++ {
			// Reading value length
			var valLen uint16
			if err := binary.Read(reader, byteOrder, &valLen); err != nil {
				return nil, err
			}

			// Reading value
			value := make([]byte, valLen)
			if _, err := reader.Read(value); err != nil {
				return nil, err
			}
			// Saving the values from buffer/reader
			n.Values = append(n.Values, value)
		}
	} else {
		//Decode Children
		// Always +1 of keyCount\
		for i := 0; i < int(keyCount)+1; i++ {
			var child uint64
			if err := binary.Read(reader, byteOrder, &child); err != nil {
				return nil, err
			}
			n.Children = append(n.Children, child)
		}
	}
	return n, nil
}
