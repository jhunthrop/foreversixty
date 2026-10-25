package notify

import (
	"bytes"
	"testing"
)

func testKey(fill byte) []byte { return bytes.Repeat([]byte{fill}, sealKeyLen) }

func TestSealerRoundTripsAndRefusesAForeignBox(t *testing.T) {
	s, err := NewSealer(testKey(1))
	if err != nil {
		t.Fatal(err)
	}
	box, err := s.Seal("the-token")
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(box, []byte("the-token")) {
		t.Fatal("box holds the plaintext")
	}
	if got, err := s.Open(box); err != nil || got != "the-token" {
		t.Fatalf("Open = %q, %v", got, err)
	}
	other, _ := NewSealer(testKey(2))
	if _, err := other.Open(box); err == nil {
		t.Fatal("a wrong key opened the box")
	}
	if _, err := s.Open(box[:4]); err == nil {
		t.Fatal("a truncated box opened")
	}
}

func TestSealsOfOnePlaintextDiffer(t *testing.T) {
	s, _ := NewSealer(testKey(1))
	a, _ := s.Seal("same")
	b, _ := s.Seal("same")
	if bytes.Equal(a, b) {
		t.Fatal("two seals were identical; the nonce is not random")
	}
}

func TestNewSealerRefusesAWrongLengthKey(t *testing.T) {
	if _, err := NewSealer(make([]byte, 16)); err == nil {
		t.Fatal("a 16-byte key was accepted")
	}
}
