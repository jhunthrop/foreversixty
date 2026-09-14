package logging

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTheLogFileRollsOverAtItsCap(t *testing.T) {
	dir := t.TempDir()
	w, err := Open(dir, 64)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	for range 10 {
		if _, err := w.Write([]byte(strings.Repeat("x", 32) + "\n")); err != nil {
			t.Fatal(err)
		}
	}
	fi, err := os.Stat(filepath.Join(dir, "companion.log"))
	if err != nil {
		t.Fatal(err)
	}
	if fi.Size() > 64 {
		t.Errorf("live log is %d bytes, want at most 64", fi.Size())
	}
	if _, err := os.Stat(filepath.Join(dir, "companion.log.1")); err != nil {
		t.Errorf("no rolled file: %v", err)
	}
}

func TestNewWritesJSONRecordsToTheFile(t *testing.T) {
	dir := t.TempDir()
	log, f, err := New(dir, true)
	if err != nil {
		t.Fatal(err)
	}
	log.Info("tailing", "component", "watch", "path", "/tmp/WoWCombatLog.txt")
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(dir, "companion.log"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), `"component":"watch"`) || !strings.Contains(string(b), `"msg":"tailing"`) {
		t.Fatalf("log = %s", b)
	}
}

func TestDebugRecordsAppearOnlyWhenAsked(t *testing.T) {
	dir := t.TempDir()
	quiet, f, err := New(dir, false)
	if err != nil {
		t.Fatal(err)
	}
	quiet.Debug("hidden", "component", "test")
	quiet.Info("shown", "component", "test")
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(dir, "companion.log"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(b), "hidden") {
		t.Error("a debug record was written at info level")
	}
	if !strings.Contains(string(b), "shown") {
		t.Error("the info record is missing")
	}
}

func TestOpenWithNoCapUsesTheDefault(t *testing.T) {
	w, err := Open(t.TempDir(), 0)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	if w.max != MaxBytes {
		t.Fatalf("max = %d, want %d", w.max, MaxBytes)
	}
}

func TestOpenInAnImpossiblePlaceIsAnError(t *testing.T) {
	if _, err := Open(filepath.Join(t.TempDir(), "missing"), 0); err == nil {
		t.Fatal("Open accepted a directory that does not exist")
	}
	if _, _, err := New(filepath.Join(t.TempDir(), "missing"), false); err == nil {
		t.Fatal("New accepted a directory that does not exist")
	}
}
