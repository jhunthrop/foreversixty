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
	first, fresh, err := c.Exports(path)
	if err != nil || !fresh {
		t.Fatalf("first read = %v, %v", fresh, err)
	}
	if len(first) == 0 {
		t.Fatal("the first read found no characters")
	}

	// The status page polls again a moment later: no reparse, same
	// answer, even though the bytes behind the file have changed
	// without moving its size or its modification time.
	if err := os.WriteFile(path, []byte(changed), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(path, t0, t0); err != nil {
		t.Fatal(err)
	}
	again, fresh, err := c.Exports(path)
	if err != nil {
		t.Fatal(err)
	}
	if fresh {
		t.Error("an unchanged file was parsed again")
	}
	if len(again) != len(first) {
		t.Fatalf("the cached answer changed: %+v", again)
	}

	// The player logs out and the game rewrites the file.
	if err := os.Chtimes(path, t0.Add(time.Hour), t0.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	moved, fresh, err := c.Exports(path)
	if err != nil || !fresh {
		t.Fatalf("after a logout = %v, %v", fresh, err)
	}
	if len(moved) != 2 || moved[0].Name != "Morrowlxn" {
		t.Fatalf("the new contents were not read: %+v", moved)
	}

	// Forget makes the next read parse regardless.
	if _, fresh, _ := c.Exports(path); fresh {
		t.Error("the second read after a logout parsed again")
	}
	c.Forget(path)
	if _, fresh, _ := c.Exports(path); !fresh {
		t.Error("Forget did not force a reread")
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
	found, fresh, err := c.Exports(path)
	if err != nil || !fresh || len(found) == 0 {
		t.Fatalf("a file that appeared was not read: %d, %v, %v", len(found), fresh, err)
	}
}
