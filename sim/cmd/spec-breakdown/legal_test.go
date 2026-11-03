package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/jhunthrop/foreversixty/sim/talents"
)

// repoRoot is the checkout this test runs in (sim/cmd/spec-breakdown).
const repoRoot = "../../.."

// A -talents string the game refuses is never measured: Demonology's
// Decimation (row 4) with nothing above it was simmed without a word.
func TestAnIllegalTalentStringIsRefused(t *testing.T) {
	o := options{repoRoot: repoRoot, spec: "warlock-demonology", preset: "raid", band: 60, faction: "horde",
		iterations: 1, seed: 1, talents: "-00000000000200"}
	_, err := legalRequest(o)
	if !errors.Is(err, talents.ErrIllegalBuild) {
		t.Fatalf("legalRequest = %v, want ErrIllegalBuild", err)
	}
}

// Every published band's build passes the same check: the guard refuses
// only what the game refuses.
func TestEveryPublishedBuildIsLegal(t *testing.T) {
	dir, err := buildDirectory(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	layouts, err := talents.Load(filepath.Join(dir, "talents"))
	if err != nil {
		t.Fatal(err)
	}
	files, err := filepath.Glob(filepath.Join(dir, "bis", "*.json"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no published BiS files under %s (%v)", dir, err)
	}
	for _, path := range files {
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		var file struct {
			Spec  string     `json:"spec"`
			Bands []bisEntry `json:"bands"`
		}
		if err := json.Unmarshal(raw, &file); err != nil {
			t.Fatalf("%s: %v", path, err)
		}
		for _, e := range file.Bands {
			if err := layouts.Check(specs.ByKey[file.Spec].ClassSlug, e.Talents, e.Band); err != nil {
				t.Errorf("%s band %d %s %s: %v", filepath.Base(path), e.Band, e.Preset, e.Faction, err)
			}
		}
	}
}
