package main

import (
	"maps"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
)

func slotsEqual(a, b map[string]int) bool { return maps.Equal(a, b) }

// Below the lowest published band (20) there is no entry, so -level 10 to 19
// runs the band-20 entry at that level on the gear the level can wear: every
// slot whose item requires more is left empty, as a player that level would.
func TestLevelBelowTheLowestBandWearsOnlyWhatTheLevelAllows(t *testing.T) {
	o := realOptions("mage-arcane")
	o.band, o.preset, o.level = 20, "bare", 20
	full, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	o.level = 12
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	if req.Character.Level != 12 {
		t.Fatalf("level = %d, want 12", req.Character.Level)
	}
	build, err := leveling.ReadActiveBuild(o.repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	fit, err := request.LoadLevelGear(o.repoRoot, build, o.spec)
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range req.Character.Gear {
		if need := fit.RequiredLevel(g.ItemID); need > 12 {
			t.Errorf("level 12 wears %s %d, which requires %d", g.Slot, g.ItemID, need)
		}
	}
	if slotsEqual(slotItems(req.Character.Gear), slotItems(full.Character.Gear)) {
		t.Error("level 12 wears the band-20 set; it has items above 12")
	}
	for _, bad := range []int{9, 30} {
		o.level = bad
		if _, err := buildRequest(o); err == nil {
			t.Errorf("-level %d ran on band 20", bad)
		}
	}
	o.band, o.level = 30, 25
	if _, err := buildRequest(o); err == nil {
		t.Error("-level 25 ran on band 30: only the lowest band reaches below itself")
	}
}
