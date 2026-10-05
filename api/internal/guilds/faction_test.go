// api/internal/guilds/faction_test.go
package guilds

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestFactionFromRaceSlugsMajority(t *testing.T) {
	faction, ok := factionFromRaceSlugs([]string{"human", "human", "orc"})
	if !ok || faction != FactionAlliance {
		t.Fatalf("majority = %q, %v, want alliance, true", faction, ok)
	}
}

func TestFactionFromRaceSlugsTie(t *testing.T) {
	_, ok := factionFromRaceSlugs([]string{"human", "orc"})
	if ok {
		t.Fatal("a tie must not pick a faction")
	}
}

func TestFactionFromRaceSlugsAllNeutralOrUnknown(t *testing.T) {
	_, ok := factionFromRaceSlugs([]string{"skyborne", "skyborne", "made-up-race"})
	if ok {
		t.Fatal("an all-neutral/unknown roster must not pick a faction")
	}
}

func TestFactionFromRaceSlugsNoExports(t *testing.T) {
	_, ok := factionFromRaceSlugs(nil)
	if ok {
		t.Fatal("no known races at all must not pick a faction")
	}
}

func TestFactionFromRaceSlugsSkyborneFactionSlugsCount(t *testing.T) {
	alliance, ok := factionFromRaceSlugs([]string{"high-order-skyborne", "high-order-skyborne", "orc"})
	if !ok || alliance != FactionAlliance {
		t.Fatalf("high-order-skyborne majority = %q, %v, want alliance, true", alliance, ok)
	}
	horde, ok := factionFromRaceSlugs([]string{"windshaper-skyborne", "windshaper-skyborne", "human"})
	if !ok || horde != FactionHorde {
		t.Fatalf("windshaper-skyborne majority = %q, %v, want horde, true", horde, ok)
	}
}

// repoRoot finds the repository root from this source file's own location (same pattern
// as api/cmd/seedguild/gear.go's repoRoot), so the drift check below works regardless of
// the caller's working directory.
func repoRoot(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("could not resolve this source file's own path")
	}
	root, err := filepath.Abs(filepath.Join(filepath.Dir(thisFile), "..", "..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	return root
}

// TestRaceFactionsMatchesTheClientBuildsRaceTable cross-checks raceFactions against
// data/builds/1.60.1.70009/races.json (the build bis_lookup.go's bisDataBuild already
// hardcodes as this package's one real catalogue) directly, so a future reassignment of
// a race's faction there fails here instead of silently drifting out of sync.
func TestRaceFactionsMatchesTheClientBuildsRaceTable(t *testing.T) {
	path := filepath.Join(repoRoot(t), "data", "builds", bisDataBuild, "races.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var races []struct {
		Slug    string `json:"slug"`
		Faction string `json:"faction"`
	}
	if err := json.Unmarshal(raw, &races); err != nil {
		t.Fatal(err)
	}
	if len(races) == 0 {
		t.Fatal("races.json read zero races")
	}
	for _, r := range races {
		if r.Faction != FactionAlliance && r.Faction != FactionHorde {
			continue // neutral/placeholder - not in raceFactions by design
		}
		got, known := raceFactions[r.Slug]
		if !known {
			t.Errorf("raceFactions is missing %q (races.json has it as %q)", r.Slug, r.Faction)
			continue
		}
		if got != r.Faction {
			t.Errorf("raceFactions[%q] = %q, races.json says %q", r.Slug, got, r.Faction)
		}
	}
}

func TestRecomputeFactionMajorityAllianceOnSeededStyleRoster(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	seedExportedCharacter(t, pool, gid, "us/hardcore/human1", "human")
	seedExportedCharacter(t, pool, gid, "us/hardcore/human2", "human")
	seedExportedCharacter(t, pool, gid, "us/hardcore/orc1", "orc")
	seedExportedCharacter(t, pool, gid, "us/hardcore/skyborne1", "skyborne") // neutral, ignored

	if err := RecomputeFaction(ctx, pool, gid); err != nil {
		t.Fatal(err)
	}
	var faction *string
	if err := pool.QueryRow(ctx, `select faction from guilds where id = $1`, gid).Scan(&faction); err != nil {
		t.Fatal(err)
	}
	if faction == nil || *faction != FactionAlliance {
		t.Fatalf("faction = %v, want alliance", faction)
	}
}

func TestRecomputeFactionTieClearsToNull(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	seedExportedCharacter(t, pool, gid, "us/hardcore/human1", "human")
	seedExportedCharacter(t, pool, gid, "us/hardcore/orc1", "orc")

	// Seed a stale non-null value first, so this test also proves a tie clears a
	// previously-computed faction rather than leaving it stuck.
	if _, err := pool.Exec(ctx, `update guilds set faction = 'alliance' where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	if err := RecomputeFaction(ctx, pool, gid); err != nil {
		t.Fatal(err)
	}
	var faction *string
	if err := pool.QueryRow(ctx, `select faction from guilds where id = $1`, gid).Scan(&faction); err != nil {
		t.Fatal(err)
	}
	if faction != nil {
		t.Fatalf("faction = %q, want null on a tie", *faction)
	}
}

func TestRecomputeFactionNoExportsLeavesNull(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	uid := seedUser(t, pool, "noexport@example.com")
	seedCharacter(t, pool, gid, uid, "us/hardcore/noexport", "member", true) // no addon_exports row

	if err := RecomputeFaction(ctx, pool, gid); err != nil {
		t.Fatal(err)
	}
	var faction *string
	if err := pool.QueryRow(ctx, `select faction from guilds where id = $1`, gid).Scan(&faction); err != nil {
		t.Fatal(err)
	}
	if faction != nil {
		t.Fatalf("faction = %q, want null with no exports", *faction)
	}
}

// seedExportedCharacter inserts a user, a guild_characters row and a matching
// addon_exports row whose FS1 export carries raceSlug, so RecomputeFaction's own export
// decode has something real to read.
func seedExportedCharacter(t *testing.T, pool *pgxpool.Pool, guildID int64, key, raceSlug string) {
	t.Helper()
	ctx := context.Background()
	uid := seedUser(t, pool, key+"@example.com")
	seedCharacter(t, pool, guildID, uid, key, "member", true)
	export := "FS1:1.60.1.70009:warrior:" + raceSlug + ":000000000/0/0:"
	if _, err := pool.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		 values ($1, $2, 'us', 'hardcore', $1, $3, now(), now())`,
		key, uid, export); err != nil {
		t.Fatal(err)
	}
}
