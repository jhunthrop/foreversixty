package rating

import (
	"context"
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Warrior spell ids from data/builds/1.60.1.70291/spellranks.json. The fury rotation shipped
// in that build's addon-data.json lists, at band 60, Bloodrage, Death Wish, Battle Shout,
// Execute, Bloodthirst, Whirlwind and Heroic Strike; at band 20 it lists Bloodrage, Death Wish,
// Battle Shout and Heroic Strike only.
const (
	spellBloodthirst4   = 23894
	spellWhirlwind1680  = 1680 // the band lists Whirlwind as 462891: the same family
	spellHeroicStrike8  = 11567
	spellHeroicStrike4  = 1608
	spellExecute1       = 5308
	spellHamstring3     = 7373
	spellSunderArmor5   = 11597
	spellShootBow       = 2480  // a ranged white hit: excluded
	spellMightyRagePot  = 17528 // the consumables catalogue's combat potion: excluded
	adherenceTestPrefix = "adh-"
)

// castRows is guid's own cast rows, one per (spell id, successful casts) pair.
func castRows(guid string, pairs ...[2]int64) []summary.CastRow {
	out := make([]summary.CastRow, 0, len(pairs))
	for _, p := range pairs {
		out = append(out, summary.CastRow{GUID: guid, OwnerGUID: guid, SpellID: p[0], Started: p[1], Succeeded: p[1]})
	}
	return out
}

// rateForAdherence rates one fight for a single player, the real RateFight against the
// repository's data builds, and reads back the stored adherence.
func rateForAdherence(t *testing.T, report, name, class, spec string, casts []summary.CastRow) *float64 {
	t.Helper()
	useRepoGearData(t)
	pool := testPool(t)
	store := &Store{Pool: pool}
	ctx := context.Background()
	guid := "Player-" + name
	for i := range casts {
		if casts[i].GUID == "" {
			casts[i].GUID, casts[i].OwnerGUID = guid, guid
		}
	}
	row := summary.RosterRow{GUID: guid, Name: name, Class: class, Spec: spec, Role: "dps", ActiveMS: 170000}
	f := fightFixture(adherenceTestPrefix+report, true, row)
	f.Summary.Casts = casts
	t.Cleanup(func() { pool.Exec(context.Background(), `delete from rating_scores where report_id = $1`, f.ReportID) })
	if err := store.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}
	var adherence *float64
	if err := pool.QueryRow(ctx, `select adherence::float8 from rating_scores where report_id = $1`,
		f.ReportID).Scan(&adherence); err != nil {
		t.Fatal(err)
	}
	return adherence
}

// seedLevel stores name's level the way the Battle.net import does (characters.level).
func seedLevel(t *testing.T, name string, level int) {
	t.Helper()
	pool := testPool(t)
	key := character.KeyFromUnit("us", "normal", name)
	mustExec(t, pool, `insert into characters (key, region, ruleset, name, class, level, imported_at)
		values ($1, 'us', 'normal', $2, 'Warrior', $3, now())
		on conflict (key) do update set level = excluded.level, imported_at = excluded.imported_at`, key, name, level)
	t.Cleanup(func() { pool.Exec(context.Background(), `delete from characters where key = $1`, key) })
}

// seedExportLevel stores name's level only in an addon export's level= section.
func seedExportLevel(t *testing.T, name string, level string) {
	t.Helper()
	pool := testPool(t)
	ctx := context.Background()
	key := character.KeyFromUnit("us", "normal", name)
	var userID int64
	if err := pool.QueryRow(ctx, `insert into users (email) values ('adherence-test@example.com')
		on conflict (email) do update set email = excluded.email returning id`).Scan(&userID); err != nil {
		t.Fatal(err)
	}
	mustExec(t, pool, `insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at)
		values ($1, $2, 'us', 'normal', $3, $4, now())
		on conflict (character_key) do update set export = excluded.export, captured_at = excluded.captured_at`,
		key, userID, name, "FS1:test-1:warrior:dwarf:0/0/0:|level="+level)
	t.Cleanup(func() {
		pool.Exec(context.Background(), `delete from addon_exports where character_key = $1`, key)
		pool.Exec(context.Background(), `delete from users where email = 'adherence-test@example.com'`)
	})
}

func wantAdherence(t *testing.T, what string, got *float64, want float64) {
	t.Helper()
	if got == nil {
		t.Fatalf("%s: adherence absent, want %.2f", what, want)
	}
	if math.Abs(*got-want) > 1e-9 {
		t.Fatalf("%s: adherence = %v, want %.2f", what, *got, want)
	}
}

// Every counted cast is one the rotation lists: any rank of a listed ability counts (a
// down-ranked Heroic Strike, Whirlwind under its other id), and white hits, consumables and
// the pet's casts are not ability casts at all.
func TestAdherenceIsOneWhenEveryCastIsInTheRotation(t *testing.T) {
	const name = "Adhfull"
	seedLevel(t, name, 60)
	casts := castRows("",
		[2]int64{spellBloodthirst4, 6}, [2]int64{spellWhirlwind1680, 4}, [2]int64{spellHeroicStrike8, 2},
		[2]int64{spellShootBow, 3}, [2]int64{spellMightyRagePot, 1})
	// The warrior's pet (a guardian, say) casts Hamstring: not the player's own cast.
	casts = append(casts, summary.CastRow{GUID: "Pet-1", OwnerGUID: "Player-" + name, SpellID: spellHamstring3,
		Started: 5, Succeeded: 5})
	got := rateForAdherence(t, "full", name, "Warrior", "Fury", casts)
	wantAdherence(t, "every cast listed", got, 1.0)
}

func TestAdherenceIsHalfWhenHalfTheCastsAreInTheRotation(t *testing.T) {
	const name = "Adhhalf"
	seedLevel(t, name, 60)
	got := rateForAdherence(t, "half", name, "Warrior", "Fury", castRows("",
		[2]int64{spellBloodthirst4, 6}, [2]int64{spellHamstring3, 3}, [2]int64{spellSunderArmor5, 3}))
	wantAdherence(t, "half listed", got, 0.5)
}

// Under ten ability casts there is too little to judge; white hits and potions do not make
// up the count.
func TestAdherenceIsAbsentUnderTenCasts(t *testing.T) {
	const name = "Adhnine"
	seedLevel(t, name, 60)
	got := rateForAdherence(t, "nine", name, "Warrior", "Fury", castRows("",
		[2]int64{spellBloodthirst4, 9}, [2]int64{spellShootBow, 5}, [2]int64{spellMightyRagePot, 1}))
	if got != nil {
		t.Fatalf("adherence = %v with nine ability casts, want absent", *got)
	}
}

func TestAdherenceIsAbsentForASpecWithNoRotation(t *testing.T) {
	const name = "Adhnoapl"
	seedLevel(t, name, 60)
	got := rateForAdherence(t, "noapl", name, "Warrior", "Gladiator", castRows("",
		[2]int64{spellBloodthirst4, 12}))
	if got != nil {
		t.Fatalf("adherence = %v for a spec with no shipped rotation, want absent", *got)
	}
}

// The band follows the character's stored level: at 25 (band 20) Execute is not in the
// rotation yet, so half the casts are off it; with the level only in the export the export's
// level= section is the level; with no level stored at all the band is 60, where it is.
func TestAdherenceBandFollowsTheStoredLevel(t *testing.T) {
	casts := func() []summary.CastRow {
		return castRows("", [2]int64{spellHeroicStrike4, 5}, [2]int64{spellExecute1, 5})
	}

	seedLevel(t, "Adhlow", 25)
	wantAdherence(t, "Battle.net level 25", rateForAdherence(t, "low", "Adhlow", "Warrior", "Fury", casts()), 0.5)

	seedExportLevel(t, "Adhexport", "25")
	wantAdherence(t, "export level 25", rateForAdherence(t, "export", "Adhexport", "Warrior", "Fury", casts()), 0.5)

	wantAdherence(t, "no level stored", rateForAdherence(t, "nolevel", "Adhnolevel", "Warrior", "Fury", casts()), 1.0)
}

// Adherence is reported to two decimals: 7 of 11 listed is 0.64.
func TestAdherenceIsRoundedToTwoDecimals(t *testing.T) {
	const name = "Adhround"
	seedLevel(t, name, 60)
	got := rateForAdherence(t, "round", name, "Warrior", "Fury", castRows("",
		[2]int64{spellBloodthirst4, 7}, [2]int64{spellHamstring3, 4}))
	wantAdherence(t, "7 of 11", got, 0.64)
}
