package zstdx

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/klauspost/compress/zstd"
)

func pack(t *testing.T, s string) []byte {
	t.Helper()
	var buf bytes.Buffer
	w, err := zstd.NewWriter(&buf)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := w.Write([]byte(s)); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestDecodeAllReadsWhatTheEngineWrote(t *testing.T) {
	packed := pack(t, "9/26 20:10:00.000  COMBAT_LOG_VERSION,16\n")
	got, err := DecodeAll(packed, MaxChunk)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "9/26 20:10:00.000  COMBAT_LOG_VERSION,16\n" {
		t.Fatalf("decoded %q", got)
	}
	// And the engine reads back what we accept, so the two agree.
	same, err := store.Decompress(packed)
	if err != nil || string(same) != string(got) {
		t.Fatalf("engine decode = %q, %v", same, err)
	}
}

func TestMaxChunkLeavesHeadroomOverAFourMiBChunk(t *testing.T) {
	if MaxChunk < 4<<20 {
		t.Fatalf("MaxChunk = %d, want at least the companion's 4 MiB chunk", MaxChunk)
	}
}

func TestDecodeAllRefusesABomb(t *testing.T) {
	packed := pack(t, strings.Repeat("a", 1<<20))
	if _, err := DecodeAll(packed, 1024); err == nil {
		t.Fatal("a chunk over the ceiling must be refused")
	}
}

func TestDecodeAllRejectsGarbage(t *testing.T) {
	if _, err := DecodeAll([]byte("not zstd at all"), MaxChunk); err == nil {
		t.Fatal("garbage must be refused")
	}
}

func TestReaderStreams(t *testing.T) {
	packed := pack(t, "one\ntwo\n")
	r, err := Reader(io.NopCloser(bytes.NewReader(packed)))
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	got, err := io.ReadAll(r)
	if err != nil || string(got) != "one\ntwo\n" {
		t.Fatalf("read %q, %v", got, err)
	}
}
