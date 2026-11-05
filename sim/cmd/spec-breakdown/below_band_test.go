package main

import (
	"testing"
)

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
	required, err := itemRequiredLevels(o.repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range req.Character.Gear {
		if required[g.ItemID] > 12 {
			t.Errorf("level 12 wears %s %d, which requires %d", g.Slot, g.ItemID, required[g.ItemID])
		}
	}
	if len(req.Character.Gear) >= len(full.Character.Gear) {
		t.Errorf("level 12 kept %d of the band's %d slots; the band-20 entry has items above 12", len(req.Character.Gear), len(full.Character.Gear))
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
