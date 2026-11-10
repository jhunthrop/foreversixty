//go:build payloaddump

package auth

import (
	"os"
	"path/filepath"
	"testing"
)

// TestDumpAppConfirmPage writes the app sign-in confirmation page, as served, to
// $REVIEW_PAYLOADS_DIR/app-confirm.html for the review captures (go test -tags payloaddump).
func TestDumpAppConfirmPage(t *testing.T) {
	dir := os.Getenv("REVIEW_PAYLOADS_DIR")
	if dir == "" {
		t.Skip("REVIEW_PAYLOADS_DIR is not set")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "app-confirm.html"), []byte(renderAppConfirm(t)), 0o644); err != nil {
		t.Fatal(err)
	}
}
