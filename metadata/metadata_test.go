package metadata

import (
	"errors"
	"testing"
)

// func TestMetadataEncode(t *testing.T) {
// 	//Arrange
// 	m := NewMetadata(4096)
// 	expected := []byte {'D', 'K', 'V', '1'}

// 	//Act
// 	data, err := m.Encode()

// 	//Assert
// 	if err != nil {
// 		t.Fatalf("err: %v", err)
// 	}
// 	if !bytes.Equal(data, expected) {
// 		t.Fatalf("data is not as expected")
// 	}
// }

func TestEncodeDecode(t *testing.T) {
	//Arrange
	original := NewMetadata(4096)

	//Act
	//encode original
	data, err := original.Encode()
	if err != nil {
		t.Fatal(err)
	}

	//decode original
	decoded, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}

	//Assert
	if original.Magic != decoded.Magic {
		t.Fatal("error")
	}
	if original.Version != decoded.Version {
		t.Fatal("error")
	}
	if original.PageSize != decoded.PageSize {
		t.Fatal("error")
	}
	if original.RootPage != decoded.RootPage {
		t.Fatal("error")
	}
	if original.NextFreePage != decoded.NextFreePage {
		t.Fatal("error")
	}

}

func TestDecodeInvalidMagicNumber(t *testing.T) {
	//Arrange
	original := NewMetadata(4096)

	//Act
	//Encode original
	data, _ := original.Encode()
	//corrupt data
	data[0] = 'S'
	//Decode original
	_, err := Decode(data)

	//Assert
	if !errors.Is(err, ErrInvalidMagicNumber) {
		t.Fatal(ErrInvalidMagicNumber)
	}
}
