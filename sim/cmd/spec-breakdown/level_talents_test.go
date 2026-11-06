package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// Round 24 hunter: -level N inside a band kept the band entry's talents, so a
// rank row judged at band+3 ran without the points the player has by then
// (Survival at 33 had no Resourcefulness). -level now wears the guide's
// published leveling order at N (the ladder string the site and addon
// compute); -band-talents keeps the entry's.
func TestLevelWearsTheLadderTalentsOfThatLevel(t *testing.T) {
	const spec = "hunter-survival"
	o := realOptions(spec)
	o.band, o.preset, o.level = 30, "bare", 33
	build, err := leveling.ReadActiveBuild(o.repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	want, err := leveling.GuideLevelTalents(o.repoRoot, build, specs.ByKey[spec], 33, false)
	if err != nil {
		t.Fatal(err)
	}
	req, err := legalRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	if req.Character.Talents != want {
		t.Errorf("-level 33 talents = %q, want the ladder's %q", req.Character.Talents, want)
	}

	o.bandTalents = true
	old, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	entry, err := loadEntry(o)
	if err != nil {
		t.Fatal(err)
	}
	if old.Character.Talents != entry.Talents {
		t.Errorf("-band-talents talents = %q, want the entry's %q", old.Character.Talents, entry.Talents)
	}
	if entry.Talents == want {
		t.Fatalf("the band-30 entry already holds the level-33 ladder build; pick a level where they differ")
	}

	o.bandTalents, o.talents = false, entry.Talents
	if got, err := buildRequest(o); err != nil || got.Character.Talents != entry.Talents {
		t.Errorf("-talents must still win over -level: got %q, %v", got.Character.Talents, err)
	}
}

// Below the lowest band the ladder spends only the level's points, so a
// level-12 run is a build the game allows (the band-20 entry has eleven).
func TestLevelBelowTheLowestBandIsALegalBuild(t *testing.T) {
	o := realOptions("mage-arcane")
	o.band, o.preset, o.level = 20, "bare", 12
	if _, err := legalRequest(o); err != nil {
		t.Fatal(err)
	}
}
