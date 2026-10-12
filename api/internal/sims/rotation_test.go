package sims

import "testing"

// The band is the addon's own Rotation.bandFor rule over the shipped bands
// (pipeline.addonrotation.LEVEL_BANDS: 10, 20, 30, 38, 40, 50, 60): the highest band at or
// below the level, and the lowest band below it.
func TestRotationBandIsTheHighestRungAtOrBelowTheLevel(t *testing.T) {
	useRepoGearData(t)
	for _, c := range []struct{ level, band int }{
		{1, 10}, {10, 10}, {19, 10}, {25, 20}, {37, 30}, {38, 38}, {39, 38}, {40, 40}, {59, 50}, {60, 60}, {70, 60},
	} {
		rot, ok := RotationFor(lootBuild, "warrior-fury", c.level)
		if !ok {
			t.Fatalf("level %d: no warrior-fury rotation", c.level)
		}
		if rot.Band != c.band {
			t.Errorf("level %d: band %d, want %d", c.level, rot.Band, c.band)
		}
	}
}

// A band lists an ability, not a rank: every rank spellranks.json groups with a listed line
// counts, and an ability the band does not list (or does not list yet) does not.
func TestRotationListsEveryRankOfAListedAbility(t *testing.T) {
	useRepoGearData(t)
	at60, ok := RotationFor("", "warrior-fury", 60)
	if !ok {
		t.Fatal("no warrior-fury rotation in the newest build")
	}
	for _, id := range []int64{23894, 23881, 1680, 462891, 25286, 11567, 78, 20662, 5308} {
		if !at60.Lists(id) {
			t.Errorf("band 60 does not list spell %d", id)
		}
	}
	for _, id := range []int64{7373, 11597, 2480, 17528} {
		if at60.Lists(id) {
			t.Errorf("band 60 lists spell %d", id)
		}
	}
	at20, _ := RotationFor("", "warrior-fury", 25)
	if at20.Lists(5308) || at20.Lists(23894) {
		t.Error("band 20 lists Execute or Bloodthirst, which it does not carry yet")
	}
	if !at20.Lists(1608) {
		t.Error("band 20 does not list Heroic Strike rank 4")
	}
	if _, ok := RotationFor("", "warrior-gladiator", 60); ok {
		t.Error("a spec with no shipped rotation has one")
	}
	if _, ok := RotationFor("0.0.0.0", "warrior-fury", 60); ok {
		t.Error("a build with no addon-data.json has a rotation")
	}
}
