//go:build payloaddump

package subscribe

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestDumpConfirmationEmail writes the subscribe confirmation to $REVIEW_PAYLOADS_DIR/mail for the
// review captures (go test -tags payloaddump), beside the auth package's dump of the others.
func TestDumpConfirmationEmail(t *testing.T) {
	dir := os.Getenv("REVIEW_PAYLOADS_DIR")
	if dir == "" {
		t.Skip("REVIEW_PAYLOADS_DIR is not set")
	}
	sent := time.Date(2026, time.November, 10, 18, 0, 0, 0, time.UTC)
	m, err := confirmationEmail("kraggor@example.com",
		"https://api.foreversixty.gg/v1/subscribe/confirm?token=4a6c8e0b2d4f6a8c", sent)
	if err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "mail")
	if err := os.MkdirAll(out, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(out, "subscribe.html"), []byte(m.HTML), 0o644); err != nil {
		t.Fatal(err)
	}
	text := "Subject: " + m.Subject + "\n\n" + m.Text
	if err := os.WriteFile(filepath.Join(out, "subscribe.txt"), []byte(text), 0o644); err != nil {
		t.Fatal(err)
	}
}
