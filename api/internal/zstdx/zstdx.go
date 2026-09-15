// Package zstdx is bounded zstd decoding. The engine's store.Decompress
// has no ceiling on the decoded size, which is safe only for objects
// this service wrote; the ingest decodes bytes a client just sent, and
// the parse job decodes an object a browser uploaded, so both need a
// bound.
package zstdx

import (
	"bytes"
	"fmt"
	"io"

	"github.com/klauspost/compress/zstd"
)

// MaxChunk is the ceiling on one decoded raw chunk. The companion sends
// 4 MiB of uncompressed log per chunk (compressed to a few hundred
// kilobytes), so four times that is ample headroom and still a bound: a
// compressed body inside the route's 8 MiB cap can otherwise decode to
// gigabytes.
const MaxChunk = 16 << 20

// DecodeAll decodes b, refusing anything that would exceed max bytes.
func DecodeAll(b []byte, max int64) ([]byte, error) {
	r, err := Reader(io.NopCloser(bytes.NewReader(b)))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	out, err := io.ReadAll(io.LimitReader(r, max+1))
	if err != nil {
		return nil, fmt.Errorf("zstdx: decode: %w", err)
	}
	if int64(len(out)) > max {
		return nil, fmt.Errorf("zstdx: decoded output exceeds %d bytes", max)
	}
	return out, nil
}

// Reader wraps a compressed stream as a plain reader. Closing it
// releases the decoder's buffers as well as the source.
func Reader(src io.ReadCloser) (io.ReadCloser, error) {
	d, err := zstd.NewReader(src, zstd.WithDecoderConcurrency(1))
	if err != nil {
		return nil, fmt.Errorf("zstdx: reader: %w", err)
	}
	return &decoder{d: d, src: src}, nil
}

type decoder struct {
	d   *zstd.Decoder
	src io.ReadCloser
}

func (d *decoder) Read(p []byte) (int, error) { return d.d.Read(p) }

func (d *decoder) Close() error {
	d.d.Close()
	return d.src.Close()
}
