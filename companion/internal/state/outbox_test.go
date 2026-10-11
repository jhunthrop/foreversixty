package state

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestTheSentOutboxKeysSurviveARestartAndAreNotReports(t *testing.T) {
	dir := t.TempDir()
	got, err := LoadOutboxSent(dir)
	if err != nil || len(got) != 0 {
		t.Fatalf("an empty state = %v, %v", got, err)
	}
	want := map[string]string{`C:\WoW\WTF\Account\A\SavedVariables\ForeverSixty.lua`: "2026-11-04 14:02|2026-11-04 14:10|ab12"}
	if err := SaveOutboxSent(dir, want); err != nil {
		t.Fatal(err)
	}
	got, err = LoadOutboxSent(dir)
	if err != nil || len(got) != 1 || got[`C:\WoW\WTF\Account\A\SavedVariables\ForeverSixty.lua`] != want[`C:\WoW\WTF\Account\A\SavedVariables\ForeverSixty.lua`] {
		t.Fatalf("reloaded = %v, %v", got, err)
	}
	// The file lives beside the reports without becoming one.
	if err := Save(dir, Report{Key: NewKey(time.Now()), StartedAt: time.Now()}); err != nil {
		t.Fatal(err)
	}
	reports, err := List(dir)
	if err != nil || len(reports) != 1 {
		t.Fatalf("reports = %d, %v", len(reports), err)
	}
	// A file the companion cannot read is an error, and the caller starts over.
	if err := os.WriteFile(filepath.Join(dir, OutboxSentFile), []byte("{"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadOutboxSent(dir); err == nil {
		t.Error("a corrupt file was read as empty without an error")
	}
}

func TestANullOrEmptySentFileIsAnEmptyUsableMap(t *testing.T) {
	for body, wantErr := range map[string]bool{"null": true, "{}": false, "null\n": true} {
		dir := t.TempDir()
		p := filepath.Join(dir, OutboxSentFile)
		if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
		got, err := LoadOutboxSent(dir)
		if got == nil {
			t.Fatalf("%q: a nil map", body)
		}
		if len(got) != 0 || (err != nil) != wantErr {
			t.Errorf("%q: %v, %v", body, got, err)
		}
		got["x"] = "y" // writable: the caller records into it
	}
}
