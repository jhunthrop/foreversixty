package addon

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// changed is savedVariables with one character of a name different,
// so it is byte-for-byte the same length.
var changed = strings.Replace(savedVariables, `["name"] = "Morrowlyn"`,
	`["name"] = "Morrowlxn"`, 1)

func TestTheCacheRereadsOnlyWhenTheFileMoves(t *testing.T) {
	path := filepath.Join(t.TempDir(), SavedVariablesName+".lua")
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(path, t0, t0); err != nil {
		t.Fatal(err)
	}
	var c Cache
	first, at, err := c.Exports(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(first) == 0 {
		t.Fatal("the first read found no characters")
	}
	if !at.Mod.Equal(t0) || at.Size == 0 {
		t.Fatalf("the stamp does not name the file that was read: %+v", at)
	}

	// The status page polls again a moment later: no reparse, same
	// answer and the same stamp, even though the bytes behind the
	// file have changed without moving its size or modification time.
	if err := os.WriteFile(path, []byte(changed), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(path, t0, t0); err != nil {
		t.Fatal(err)
	}
	again, sameAt, err := c.Exports(path)
	if err != nil {
		t.Fatal(err)
	}
	if !sameAt.eq(at) {
		t.Errorf("an unchanged file got a new stamp: %+v", sameAt)
	}
	if len(again) != len(first) || again[0].Name != first[0].Name {
		t.Fatalf("the cached answer changed: %+v", again)
	}

	// The player logs out and the game rewrites the file.
	if err := os.Chtimes(path, t0.Add(time.Hour), t0.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	moved, movedAt, err := c.Exports(path)
	if err != nil {
		t.Fatal(err)
	}
	if movedAt.eq(at) {
		t.Error("a rewritten file kept its old stamp")
	}
	if len(moved) != 2 || moved[0].Name != "Morrowlxn" {
		t.Fatalf("the new contents were not read: %+v", moved)
	}
}

func TestTheCacheDoesNotRememberAMissingFile(t *testing.T) {
	var c Cache
	path := filepath.Join(t.TempDir(), "nope.lua")
	if _, _, err := c.Exports(path); err == nil {
		t.Fatal("a missing file was not an error")
	}
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	found, at, err := c.Exports(path)
	if err != nil || len(found) == 0 {
		t.Fatalf("a file that appeared was not read: %d, %v", len(found), err)
	}
	if at.Size == 0 {
		t.Errorf("the stamp is empty: %+v", at)
	}
}

// TestTheCacheDoesNotDecideWhoUploads is the regression the shared
// cache caused: a status poll may read a file as often as it likes
// without costing the sync the upload that file's change is owed.
func TestTheCacheDoesNotDecideWhoUploads(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "SavedVariables")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, SavedVariablesName+".lua")
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	shared := &Cache{}
	api := &fakeAPI{}
	s := New(Options{Paths: func() []string { return []string{path} }, API: api, Cache: shared})

	// The status page warms the cache before the sync's first pass.
	if _, _, err := shared.Exports(path); err != nil {
		t.Fatal(err)
	}
	if err := s.Poll(t.Context(), t0); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 2 {
		t.Fatalf("a warm cache cost the first upload: %+v", api.posted)
	}

	// And again after a logout.
	if err := os.Chtimes(path, t0.Add(time.Hour), t0.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if _, _, err := shared.Exports(path); err != nil {
		t.Fatal(err)
	}
	if err := s.Poll(t.Context(), t0.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 4 {
		t.Fatalf("a warm cache cost a logout's upload: %+v", api.posted)
	}
}
