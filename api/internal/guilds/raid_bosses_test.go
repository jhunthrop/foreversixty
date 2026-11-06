package guilds

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/raidbosses"
)

// TestRaidBossesTableMatchesItsGenerator pins raid_bosses_table.go to what cmd/genraidbosses
// writes from the build's raid-bosses.json.
func TestRaidBossesTableMatchesItsGenerator(t *testing.T) {
	body, err := os.ReadFile(filepath.Join(repoDataDirForTrees(t), raidBossesBuild, "raid-bosses.json"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := raidbosses.Parse(body)
	if err != nil {
		t.Fatal(err)
	}
	want, err := raidbosses.GoSource(doc, raidbosses.SourcePath(raidBossesBuild))
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile("raid_bosses_table.go")
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(want) {
		t.Fatal("raid_bosses_table.go is stale: run go generate ./internal/guilds/ from api/")
	}
}

// TestRaidBossesAreFirstTierZones keeps every table raid a zone of first-tier.json (raidZones).
func TestRaidBossesAreFirstTierZones(t *testing.T) {
	zones := map[string]bool{}
	for _, z := range raidZones {
		zones[z.Key] = true
	}
	for _, r := range firstTierRaids {
		if !zones[r.Key] {
			t.Errorf("%s is in the boss table but not a first-tier zone", r.Key)
		}
	}
}
