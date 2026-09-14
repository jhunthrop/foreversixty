package card

import (
	"os"
	"testing"
)

// TestWriteSampleCard is a look-at-it check, skipped unless CARD_OUT is set:
//
//	CARD_OUT=/tmp/card.png go test ./internal/card -run TestWriteSampleCard
func TestWriteSampleCard(t *testing.T) {
	out := os.Getenv("CARD_OUT")
	if out == "" {
		t.Skip("set CARD_OUT to write a sample card")
	}
	b, err := Render(sample())
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(out, b, 0o644); err != nil {
		t.Fatal(err)
	}
}
