package sims

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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

// Family names a class spell's spellranks.json family, so exclusions are matched on the
// data's own names, never the log's; a spell no rank of any of the class's spells has none.
func TestRotationNamesTheClassFamilyOfASpell(t *testing.T) {
	useRepoGearData(t)
	ret, ok := RotationFor("", "paladin-retribution", 60)
	if !ok {
		t.Fatal("no paladin-retribution rotation")
	}
	if name, ok := ret.Family(20375); !ok || name != "Seal of Command" {
		t.Errorf("Family(20375) = %q, %v; want Seal of Command", name, ok)
	}
	fury, _ := RotationFor("", "warrior-fury", 60)
	if name, ok := fury.Family(6554); !ok || name != "Pummel" {
		t.Errorf("Family(6554) = %q, %v; want Pummel", name, ok)
	}
	for _, id := range []int64{20572, 6603, 17528} { // Blood Fury, Attack, Mighty Rage Potion
		if name, ok := fury.Family(id); ok {
			t.Errorf("Family(%d) = %q, want no warrior family", id, name)
		}
	}
}

// A build whose addon-data.json carries no rotations has none to measure against, and says so
// once in the log, not on every rated player.
func TestABuildWithNoRotationsIsLoggedOnce(t *testing.T) {
	dir := t.TempDir()
	useRepoGearData(t)
	SetGearData(currentGearData().trees, dir)
	t.Cleanup(func() { SetGearData(nil, "") })
	if err := os.MkdirAll(filepath.Join(dir, lootBuild), 0o755); err != nil {
		t.Fatal(err)
	}
	for file, body := range map[string]string{
		"addon-data.json": `{"build":"` + lootBuild + `","classes":{}}`,
		"spellranks.json": `{"build":"` + lootBuild + `","classes":{}}`,
	} {
		if err := os.WriteFile(filepath.Join(dir, lootBuild, file), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	var buf bytes.Buffer
	prev := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(&buf, nil)))
	t.Cleanup(func() { slog.SetDefault(prev) })

	for i := 0; i < 3; i++ {
		if _, ok := RotationFor(lootBuild, "warrior-fury", 60); ok {
			t.Fatal("a build with no rotations has one")
		}
	}
	if n := strings.Count(buf.String(), "no rotations"); n != 1 {
		t.Fatalf("logged %d times, want once:\n%s", n, buf.String())
	}
}
