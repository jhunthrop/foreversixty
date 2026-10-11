// companion/internal/state/outbox.go
// The addon outboxes the site has already answered, one key per
// SavedVariables file, so a restart of the companion never sends the
// same outbox (and its sim request) twice. The file sits in a folder of
// its own under state/, so List never mistakes it for a report.
package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// OutboxSentFile is the file, relative to the state directory.
var OutboxSentFile = filepath.Join("addon", "outbox-sent.json")

// LoadOutboxSent reads the answered outbox keys, by SavedVariables
// path. A missing file is an empty map; a file that cannot be read is
// an error, and the caller decides whether to start over.
func LoadOutboxSent(dir string) (map[string]string, error) {
	p := filepath.Join(dir, OutboxSentFile)
	b, err := os.ReadFile(p)
	if errors.Is(err, fs.ErrNotExist) {
		return map[string]string{}, nil
	}
	if err != nil {
		return map[string]string{}, err
	}
	out := map[string]string{}
	if err := json.Unmarshal(b, &out); err != nil {
		return map[string]string{}, fmt.Errorf("parse %s: %w", p, err)
	}
	return out, nil
}

// SaveOutboxSent writes the answered outbox keys atomically.
func SaveOutboxSent(dir string, sent map[string]string) error {
	p := filepath.Join(dir, OutboxSentFile)
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return err
	}
	b, err := json.MarshalIndent(sent, "", "  ")
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(p), ".outbox-*.json")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if _, err := tmp.Write(append(b, '\n')); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Chmod(name, 0o600); err != nil {
		return err
	}
	return os.Rename(name, p)
}
