package regionreader

import (
	"bytes"
	"compress/zlib"
	"errors"
	"io"
)

// DecompressZlib decompresses data using zlib.
func DecompressZlib(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	decompressed, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(decompressed) == 0 {
		return nil, errors.New("decompressed data is empty")
	}
	return decompressed, nil
}
