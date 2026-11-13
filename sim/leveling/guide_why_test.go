package leveling

import (
	"path/filepath"
	"testing"
)

const whyGuide = `---
title: Test
build: 'FS1:x'
buildWhy: 'Levels with the pet''s Rapid Recuperation'
raidBuild: 'FS1:y'
raidBuildWhy: 'Trades the pet for Lone Wolf'
---
`

func TestBuildWhyReadsEachKindsFrontmatterLine(t *testing.T) {
	for kind, want := range map[string]string{
		"leveling": "Levels with the pet's Rapid Recuperation",
		"raid":     "Trades the pet for Lone Wolf",
	} {
		got, err := BuildWhy(whyGuide, kind)
		if err != nil || got != want {
			t.Errorf("%s why = %q, %v; want %q", kind, got, err, want)
		}
	}
}

func TestBuildWhyReadsADoubleQuotedLine(t *testing.T) {
	guide := "---\nbuild: 'FS1:x'\nbuildWhy: \"Melee hunter: Predator's Edge\"\n---\n"
	got, err := BuildWhy(guide, "leveling")
	if err != nil || got != "Melee hunter: Predator's Edge" {
		t.Fatalf("why = %q, %v", got, err)
	}
}

func TestAOneBuildGuidesRaidWhyIsItsBuildWhy(t *testing.T) {
	guide := "---\nbuild: 'FS1:x'\nbuildWhy: 'Every melee talent in Survival'\n---\n"
	got, err := BuildWhy(guide, "raid")
	if err != nil || got != "Every melee talent in Survival" {
		t.Fatalf("raid why = %q, %v", got, err)
	}
}

func TestBuildWhyIsAnErrorWhenTheLineIsMissing(t *testing.T) {
	if _, err := BuildWhy("---\nbuild: 'FS1:x'\nraidBuild: 'FS1:y'\nbuildWhy: 'a'\n---\n", "raid"); err == nil {
		t.Fatal("a guide with a raidBuild and no raidBuildWhy must be an error")
	}
	if _, err := GuideBuildWhy(filepath.Join("..", ".."), "hunter", "beast-mastery", "pvp"); err == nil {
		t.Fatal("an unknown kind must be an error")
	}
}
