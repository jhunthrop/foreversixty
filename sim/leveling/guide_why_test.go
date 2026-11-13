package leveling

import (
	"path/filepath"
	"strings"
	"testing"
)

const whyGuide = `---
title: Test
---

## Overview

Words.

## Talents and builds

There are two builds. You level with the pet: thirty-one points in Marksmanship take the big shots. At 60, for raids, the pet goes: thirty-three points in Marksmanship.

- **Lone Wolf** (1 point): 20% on everything without a pet

The raid build drops **Rapid Recuperation** for Lone Wolf. Out in the world, keep the level 60 build: the pet holds what you pull.

## Rotation and priority

More words.
`

// The raid why is the sentence of its first lead ("The raid build" before
// "At 60, for raids"), and a short one never takes the next sentence when
// that one opens the other build's why.
func TestBuildWhyTakesTheKindsOwnSentence(t *testing.T) {
	got := BuildWhy(whyGuide, "raid")
	want := "The raid build drops Rapid Recuperation for Lone Wolf."
	if got != want {
		t.Fatalf("raid why = %q, want %q", got, want)
	}
	got = BuildWhy(whyGuide, "leveling")
	want = "You level with the pet: thirty-one points in Marksmanship take the big shots."
	if got != want {
		t.Fatalf("leveling why = %q, want %q", got, want)
	}
}

func TestBuildWhyFallsBackToTheFirstSentenceJoinedWhenShortAndStripsBold(t *testing.T) {
	guide := strings.Replace(whyGuide, "You level with the pet", "**Both** builds want hit", 1)
	guide = strings.Replace(guide, "At 60, for raids, the pet goes", "Then the pet goes", 1)
	guide = strings.Replace(guide, "The raid build drops", "Raiding drops", 1)
	guide = strings.Replace(guide, "Out in the world, keep the level 60 build", "Keep the level 60 build out in the world", 1)
	got := BuildWhy(guide, "raid")
	want := "There are two builds. Both builds want hit: thirty-one points in Marksmanship take the big shots."
	if got != want {
		t.Fatalf("fallback why = %q, want %q", got, want)
	}
}

// Round 1 of the reader took the Combat raid sentence for the leveling why,
// because "You level and run dungeons with the level 60 build" matched no
// lead, and gave Assassination "The level 60 build is the one you level
// with." alone.
func TestBuildWhyReadsTheRogueGuides(t *testing.T) {
	root := filepath.Join("..", "..")
	combat, err := GuideBuildWhy(root, "rogue", "combat", "leveling")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(combat, "You level and run dungeons with the level 60 build") {
		t.Fatalf("Combat leveling why = %q", combat)
	}
	assassination, err := GuideBuildWhy(root, "rogue", "assassination", "leveling")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(assassination, "The level 60 build is the one you level with. It trades") {
		t.Fatalf("Assassination leveling why = %q", assassination)
	}
}

func TestBuildWhyIsEmptyWithoutTheSection(t *testing.T) {
	if got := BuildWhy("## Overview\n\nWords.\n", "leveling"); got != "" {
		t.Fatalf("why without a Talents section = %q, want empty", got)
	}
	if got := BuildWhy("## Talents and builds\n\n- **Only** (1 point): bullets\n\n## Next\n", "raid"); got != "" {
		t.Fatalf("why with bullets only = %q, want empty", got)
	}
}

func TestGuideBuildWhyReadsTheHunterGuides(t *testing.T) {
	root := filepath.Join("..", "..")
	raid, err := GuideBuildWhy(root, "hunter", "beast-mastery", "raid")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(raid, "The raid build moves") {
		t.Fatalf("Beast Mastery raid why = %q", raid)
	}
	leveling, err := GuideBuildWhy(root, "hunter", "beast-mastery", "leveling")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(leveling, "Out in the world, keep the level 60 build") {
		t.Fatalf("Beast Mastery leveling why = %q", leveling)
	}
	if _, err := GuideBuildWhy(root, "hunter", "beast-mastery", "pvp"); err == nil {
		t.Fatal("an unknown kind must be an error")
	}
}
