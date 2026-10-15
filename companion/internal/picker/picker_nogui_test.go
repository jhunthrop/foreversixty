//go:build nogui

package picker

import (
	"errors"
	"testing"
)

// The headless build has no dialog and must say so rather than hang.
// (The real build's test would open a dialog, so there is none.)
func TestThereIsNoDialogInTheHeadlessBuild(t *testing.T) {
	if _, err := Pick(); !errors.Is(err, ErrUnavailable) {
		t.Fatalf("err = %v", err)
	}
}
