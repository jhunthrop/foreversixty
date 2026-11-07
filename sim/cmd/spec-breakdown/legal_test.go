package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
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

// Every published band's build and race pass the same checks: the guards
// refuse only what the game refuses.
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
			if err := checkRace(dir, specs.ByKey[file.Spec].ClassSlug, e.Faction, e.Race); err != nil {
				t.Errorf("%s band %d %s %s: %v", filepath.Base(path), e.Band, e.Preset, e.Faction, err)
			}
		}
	}
}

// A -race the class cannot be is never measured: Skyborne priests ran
// without a word (round 26), though no Skyborne can be a priest in the
// client's race-class combinations. A race of the other faction is refused
// too: the published entry's gear and buffs are the faction's.
func TestARaceTheClassCannotBeIsRefused(t *testing.T) {
	cases := []struct {
		spec, faction, race string
	}{
		{"priest-shadow", "horde", "windshaper-skyborne"},
		{"priest-holy", "alliance", "high-order-skyborne"},
		{"paladin-retribution", "alliance", "orc"},
		{"warrior-arms", "horde", "human"},
		{"mage-fire", "horde", "nope"},
	}
	for _, c := range cases {
		o := options{repoRoot: repoRoot, spec: c.spec, preset: "raid", band: 60, faction: c.faction,
			iterations: 1, seed: 1, race: c.race}
		_, err := legalRequest(o)
		if err == nil {
			t.Errorf("%s %s -race %s: legalRequest accepted it", c.spec, c.faction, c.race)
			continue
		}
		if !errors.Is(err, errIllegalRace) || !strings.Contains(err.Error(), c.race) {
			t.Errorf("%s %s -race %s: legalRequest = %v, want errIllegalRace naming the race", c.spec, c.faction, c.race, err)
		}
	}
}

// Every race the class can be in the faction still runs, Forever's added
// combinations included.
func TestEveryLegalRaceIsAccepted(t *testing.T) {
	for _, c := range []struct{ spec, faction, race string }{
		{"priest-shadow", "horde", "troll"},
		{"mage-fire", "alliance", "high-order-skyborne"},
		{"paladin-retribution", "horde", "undead"},
		{"warrior-arms", "horde", "windshaper-skyborne"},
	} {
		o := options{repoRoot: repoRoot, spec: c.spec, preset: "raid", band: 60, faction: c.faction,
			iterations: 1, seed: 1, race: c.race}
		if _, err := legalRequest(o); err != nil {
			t.Errorf("%s %s -race %s: %v", c.spec, c.faction, c.race, err)
		}
	}
}
