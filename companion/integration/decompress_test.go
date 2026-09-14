package integration_test

import "github.com/klauspost/compress/zstd"

// decompress unpacks a raw chunk the way the server does.
func decompress(packed []byte) ([]byte, error) {
	d, err := zstd.NewReader(nil)
	if err != nil {
		return nil, err
	}
	defer d.Close()
	return d.DecodeAll(packed, nil)
}
