package main

import (
	"errors"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

const repoRoot = "../../.."

// Every talent of the active build must have an engine field to write
// to, or a candidate taking it could not be simmed at all.
func TestEveryActiveTalentHasAnEngineField(t *testing.T) {
	build, err := leveling.ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	engineDir, err := enginetalents.SourceDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, class := range []string{"druid", "hunter", "mage", "paladin", "priest", "rogue", "shaman", "warlock", "warrior"} {
		trees, err := leveling.LoadTalentTrees(repoRoot, build, class)
		if err != nil {
			t.Fatal(err)
		}
		layout, err := enginetalents.ForClass(engineDir, class)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := layout.Encode(trees, nil); err != nil {
			t.Errorf("%s: %v", class, err)
		}
	}
}

func TestReportNameCarriesANon60Level(t *testing.T) {
	if got := reportName(options{spec: "x", level: 60}); got != "x" {
		t.Errorf("level 60: %q", got)
	}
	if got := reportName(options{spec: "x", level: 40}); got != "x-40" {
		t.Errorf("level 40: %q", got)
	}
}

func TestPrepareSkipsHealersTanksAndUnwrittenRotations(t *testing.T) {
	for _, spec := range []string{"paladin-holy", "warrior-protection"} {
		if _, err := prepare(options{repoRoot: repoRoot, spec: spec, level: 60}); !errors.Is(err, errSkipped) {
			t.Errorf("%s: err = %v, want a skip", spec, err)
		}
	}
}

func TestReportNameSuffixesANonBarePreset(t *testing.T) {
	if got := reportName(options{spec: "x", level: 60, preset: "raid"}); got != "x-raid" {
		t.Errorf("raid = %q, want x-raid", got)
	}
	if got := reportName(options{spec: "x", level: 60, preset: "bare"}); got != "x" {
		t.Errorf("bare = %q, want x", got)
	}
}
