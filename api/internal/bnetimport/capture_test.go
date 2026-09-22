// api/internal/bnetimport/capture_test.go
package bnetimport

import (
	"bytes"
	"encoding/json"
	"testing"
)

// rawOfSize builds a valid JSON string value of exactly n total bytes
// (including its surrounding quotes), so capCapture's length check can be
// tested at an exact boundary.
func rawOfSize(n int) json.RawMessage {
	const overhead = 2 // the two quote characters
	body := bytes.Repeat([]byte("a"), n-overhead)
	return json.RawMessage(append(append([]byte{'"'}, body...), '"'))
}

func TestCapCaptureRefusesAnOversizedBody(t *testing.T) {
	s := &Service{}
	raw := rawOfSize(maxCaptureBytes + 1)
	if got := s.capCapture("bnet_profile", "us/pvp/thoradin", raw); got != nil {
		t.Fatalf("capCapture returned %d bytes, want nil for an oversized body", len(got))
	}
}

func TestCapCaptureKeepsABodyAtTheLimit(t *testing.T) {
	s := &Service{}
	raw := rawOfSize(maxCaptureBytes)
	got := s.capCapture("bnet_profile", "us/pvp/thoradin", raw)
	if len(got) != len(raw) {
		t.Fatalf("capCapture returned %d bytes, want the body kept at exactly the limit (%d)", len(got), len(raw))
	}
}

func TestRawOrNilTurnsAnEmptyRawMessageIntoARealNil(t *testing.T) {
	if got := rawOrNil(nil); got != nil {
		t.Fatalf("rawOrNil(nil) = %v, want a nil interface value", got)
	}
	raw := json.RawMessage(`{"a":1}`)
	got := rawOrNil(raw)
	if got == nil {
		t.Fatal("rawOrNil(non-nil) = nil, want the raw bytes back")
	}
}
