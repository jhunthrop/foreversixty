package inproc

import (
	"testing"
)

func TestRegisterIsIdempotent(t *testing.T) {
	Register()
	Register()
}

func TestNextRunIDIsUniquePerCall(t *testing.T) {
	a, b := nextRunID(), nextRunID()
	if a == b || a == "" || b == "" {
		t.Fatalf("nextRunID returned %q then %q, want two distinct non-empty ids", a, b)
	}
}
