// logs/engine/mechanics/utility/utility_test.go
package utility

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestParseAcceptsTheSpecFormatAndRejectsAnUnknownKind(t *testing.T) {
	good := []byte(`{"spec": "warrior-protection", "owned": [
		{"spell_id": 1160, "name": "Demoralizing Shout", "kind": "debuff", "target": "enemy", "verified": "1.60.1.69893/spells.json#1160"}]}`)
	table, err := Parse(good)
	if err != nil {
		t.Fatal(err)
	}
	if table.Spec != "warrior-protection" || len(table.Owned) != 1 || table.Owned[0].Kind != Debuff {
		t.Fatalf("table = %+v", table)
	}

	bad := []byte(`{"spec": "warrior-protection", "owned": [
		{"spell_id": 1160, "name": "X", "kind": "aoe", "verified": "1.60.1.69893/spells.json#1160"}]}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("an unknown kind must be refused")
	}
}

func TestParseRejectsMissingSpecSpellIDVerificationOrADuplicate(t *testing.T) {
	cases := map[string]string{
		"no spec": `{"owned": [{"spell_id": 1, "name": "X", "kind": "buff", "verified": "1#1"}]}`,
		"zero spell id": `{"spec": "x", "owned": [{"spell_id": 0, "name": "X", "kind": "buff", "verified": "1#1"}]}`,
		"no verified provenance": `{"spec": "x", "owned": [{"spell_id": 1, "name": "X", "kind": "buff"}]}`,
		"a duplicate spell id": `{"spec": "x", "owned": [
			{"spell_id": 1, "name": "X", "kind": "buff", "verified": "1#1"},
			{"spell_id": 1, "name": "X again", "kind": "debuff", "verified": "1#1"}]}`,
	}
	for name, data := range cases {
		if _, err := Parse([]byte(data)); err == nil {
			t.Errorf("%s must be refused", name)
		}
	}
}

func TestLoadFindsAnEmbeddedTableAndSaysWhenThereIsNone(t *testing.T) {
	if table, ok := Load("warrior-protection"); !ok || table.Spec != "warrior-protection" {
		t.Fatalf("Load(warrior-protection) = %+v, %v", table, ok)
	}
	if _, ok := Load("bard-troubadour"); ok {
		t.Fatal("a spec without a table must report false")
	}
}

// TestEveryEmbeddedTableParses walks the embed FS rather than naming the
// tables, so a new one is covered the day it is added and a malformed one
// fails here rather than the first time that spec is scored.
func TestEveryEmbeddedTableParses(t *testing.T) {
	entries, err := tables.ReadDir("tables")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) == 0 {
		t.Fatal("no embedded tables: the go:embed pattern matched nothing")
	}
	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			data, err := tables.ReadFile("tables/" + entry.Name())
			if err != nil {
				t.Fatal(err)
			}
			table, err := Parse(data)
			if err != nil {
				t.Fatal(err)
			}
			if want := table.Spec + ".json"; entry.Name() != want {
				t.Fatalf("%s holds spec %q, so Load would never find it; want the file named %s",
					entry.Name(), table.Spec, want)
			}
		})
	}
}

// TestEveryCuratedSpecHasAUtilityFile cross-checks the embedded tables
// against data/curated/specs.json's 27 entries by reading the live file
// from disk (a plain os.ReadFile, not go:embed, which cannot reach outside
// this module) -- so a spec added to or removed from specs.json is caught
// here rather than silently drifting.
func TestEveryCuratedSpecHasAUtilityFile(t *testing.T) {
	root, err := repoRoot()
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(root, "data", "curated", "specs.json"))
	if err != nil {
		t.Fatal(err)
	}
	var specs []struct {
		Spec string `json:"spec"`
	}
	if err := json.Unmarshal(data, &specs); err != nil {
		t.Fatal(err)
	}
	if len(specs) == 0 {
		t.Fatal("data/curated/specs.json parsed to zero specs")
	}
	for _, s := range specs {
		if _, ok := Load(s.Spec); !ok {
			t.Errorf("data/curated/specs.json lists %q, but logs/engine/mechanics/utility/tables/%s.json does not exist", s.Spec, s.Spec)
		}
	}
}

// repoRoot walks up from the working directory (a Go test's cwd is always
// its own package directory) to the directory that holds go.work, the
// repository root every module hangs off.
func repoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.work")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}
