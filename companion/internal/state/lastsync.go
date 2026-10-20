// companion/internal/state/lastsync.go
// What the last character sync told the site, kept so the Status page
// can say it after a restart: how many characters went up and when, and
// the site's reason when it refused the list. It sits beside the answered
// outboxes, in the addon folder under state/, so List never mistakes it
// for a report.
package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// LastSyncFile is the file, relative to the state directory.
var LastSyncFile = filepath.Join("addon", "last-sync.json")

// LastSync is the newest word on the character list. Count and At are
// the newest list the site accepted; Refusal and RefusedAt are the
// newest list it refused, with the reason in the API's message. A
// refusal older than the last acceptance is cleared when it is recorded.
type LastSync struct {
	Count     int       `json:"count"`
	At        time.Time `json:"at,omitzero"`
	Refusal   string    `json:"refusal,omitempty"`
	RefusedAt time.Time `json:"refused_at,omitzero"`
}

// LoadLastSync reads the record. A missing file is the zero record; one
// that cannot be read is an error and a zero record, and the caller
// decides whether to start over.
func LoadLastSync(dir string) (LastSync, error) {
	p := filepath.Join(dir, LastSyncFile)
	b, err := os.ReadFile(p)
	if errors.Is(err, fs.ErrNotExist) {
		return LastSync{}, nil
	}
	if err != nil {
		return LastSync{}, err
	}
	var out LastSync
	if err := json.Unmarshal(b, &out); err != nil {
		return LastSync{}, fmt.Errorf("parse %s: %w", p, err)
	}
	return out, nil
}

// SaveLastSync writes the record atomically.
func SaveLastSync(dir string, ls LastSync) error {
	return writeJSON(filepath.Join(dir, LastSyncFile), ls)
}

// writeJSON writes v to p atomically: a temporary file in the same
// folder, synced, then renamed over p.
func writeJSON(p string, v any) error {
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return err
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(p), ".last-sync-*.json")
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
