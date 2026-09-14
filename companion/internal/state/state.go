// companion/internal/state/state.go
// Package state is one file per report under state/, holding the
// engine's serialised session and the offsets the uploader resumes
// from. It is what makes a restart mid-raid continue the same report
// instead of starting a second one.
package state

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/character"
)

// Report is one report's durable state.
type Report struct {
	// Key is the local identity, assigned before the server has seen
	// the report, and the name of this file.
	Key string `json:"key"`
	// ReportID is the server's id, empty until the report is created.
	ReportID string `json:"report_id"`
	// LogPath is the combat log this report is being read from.
	LogPath string `json:"log_path"`
	// StartOffset is where the report began in that file; Offset is
	// the next byte to feed; RawSent is how far the raw chunks have
	// been queued.
	StartOffset int64 `json:"start_offset"`
	Offset      int64 `json:"offset"`
	RawSent     int64 `json:"raw_sent"`

	Visibility string `json:"visibility"`
	// LoggingCharacter attributes the report, captured when it starts
	// so a settings change mid-raid does not rewrite history.
	LoggingCharacter *character.Character `json:"logging_character,omitempty"`
	Zone             string               `json:"zone,omitempty"`
	Title            string               `json:"title,omitempty"`
	StartedAt        time.Time            `json:"started_at"`
	LastAppend       time.Time            `json:"last_append"`
	// Fights is how many fights have been queued, for the UI.
	Fights int `json:"fights"`
	// Closed is set when the log has been completed locally; Done is
	// set once the completion has been accepted by the server.
	Closed bool `json:"closed"`
	Done   bool `json:"done"`
	// Session is the engine's serialised session, base64 in JSON.
	Session []byte `json:"session,omitempty"`
	// EngineVersion is the engine that produced this report.
	EngineVersion string `json:"engine_version"`
}

// NewKey mints a local report key: a sortable timestamp and eight
// random hex characters, so two reports started in the same second on
// the same machine never collide.
func NewKey(now time.Time) string {
	var b [4]byte
	rand.Read(b[:])
	return fmt.Sprintf("%s-%x", now.UTC().Format("20060102-150405"), b)
}

func path(dir, key string) string { return filepath.Join(dir, key+".json") }

// Save writes one report's state atomically.
func Save(dir string, r Report) error {
	if r.Key == "" {
		return errors.New("state: the report key is empty")
	}
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, ".state-*.json")
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
	return os.Rename(name, path(dir, r.Key))
}

// Load reads one report's state.
func Load(dir, key string) (Report, error) {
	b, err := os.ReadFile(path(dir, key))
	if err != nil {
		return Report{}, err
	}
	var r Report
	if err := json.Unmarshal(b, &r); err != nil {
		return Report{}, fmt.Errorf("parse %s: %w", path(dir, key), err)
	}
	return r, nil
}

// List reads every report state, oldest key first. A file that will
// not parse is skipped rather than fatal: one corrupt report must not
// stop the companion from logging tonight.
func List(dir string) ([]Report, error) {
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".json") {
			keys = append(keys, strings.TrimSuffix(e.Name(), ".json"))
		}
	}
	sort.Strings(keys)
	out := make([]Report, 0, len(keys))
	for _, k := range keys {
		r, err := Load(dir, k)
		if err != nil {
			continue
		}
		out = append(out, r)
	}
	return out, nil
}

// Remove deletes one report's state.
func Remove(dir, key string) error {
	err := os.Remove(path(dir, key))
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	return err
}
