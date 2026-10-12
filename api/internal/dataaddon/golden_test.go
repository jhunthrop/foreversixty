// api/internal/dataaddon/golden_test.go
package dataaddon

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/digest"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// goldenFixturePath is contract 5's own fixture (docs/contracts/addon-v2.md), the file the
// addon's addon/tests/contracts_spec.lua loads too, so the job and the reader can never
// drift apart silently.
const goldenFixturePath = "../../../docs/contracts/fixtures/data-addon-v2.lua"

// fixtureToday is the date the fixture is written as of: the golden run's own dates are
// moved onto it, day for day, before the comparison (normaliseDates).
var fixtureToday = time.Date(2026, 11, 4, 0, 0, 0, 0, time.UTC)

// repoBuilds is the repository's own data/builds, the directory the API's image ships as
// TreeDataDir.
const repoBuilds = "../../../data/builds"

var (
	repoTreesOnce sync.Once
	repoTrees     *trees.Data
	repoTreesErr  error
)

// useRepoData loads the repository's real data/builds once per test binary (every build's
// item tables take a few seconds) and wires sims.GearScore to it, as the job's own
// command does, so ScoreV2's gear component is the real one.
func useRepoData(t *testing.T) *trees.Data {
	t.Helper()
	repoTreesOnce.Do(func() {
		if repoTrees, repoTreesErr = trees.Load(repoBuilds); repoTreesErr == nil {
			sims.SetGearData(repoTrees, repoBuilds)
		}
	})
	if repoTreesErr != nil {
		t.Fatal(repoTreesErr)
	}
	return repoTrees
}

// goldenGear is docs/contracts/fixtures/gearscore.json's gear code (the warrior-arms
// band-20 alliance best-in-slot set), with a level= section: a level-24 warrior is
// measured against its own band-20 set.
const goldenGear = "FS1:1.60.1.70291:warrior:human:371000000/0/0:back=4706,chest=6627,feet=10402," +
	"finger1=20439,finger2=12054,hands=10401,head=250528,legs=10400,main_hand=23171,neck=20444," +
	"ranged=285279,shoulder=3480,waist=10403,wrist=2868|level=24"

// goldenEncounter is the first of this file's own encounter ids (990701..990704).
const goldenEncounter = 990701

// seedGolden builds the database the fixture is the nightly file of. Every row it writes
// carries the "dav2" prefix (report ids "dav2-...", character slugs "dav2-...", emails
// "dav2-...", guild names "Dav2 ...", encounter ids 990701..990704), and its cleanup
// deletes only rows with that prefix:
//
//   - Obnoxious Yell (us/normal, Battle.net realm "Classic Beta PvE 2"), in Dav2 Olympus:
//     four rated fights over three of the guild's four nights (one without a flask), a
//     ranked Lucifron kill placed against a 1..100 curve, a level-24 export of the band-20
//     best-in-slot set, a Battle.net item level of 31, and a linked alt (Little Yell) with a
//     kill of its own.
//   - Quiet Fern (us/normal, realm from its export's who=), a verified member of Dav2
//     Olympus with no rated fight: a no-logs record with the spec its export's talents name
//     and its item level.
//   - Hidden Hand (us/normal), a verified member of Dav2 Olympus with a rated fight whose
//     account is anonymised: no record, not listed, but counted in the roster.
//   - Thoradin (us/normal), the one verified member of Dav2 Vanguard, with no realm from
//     any source: no record (counted as skipped), so its guild lists no one.
//   - Mörk (eu/pvp), one rated fight logged as "Mörk-ClassicBetaPvP" in a guildless public
//     report, no characters row (name and realm come from the rated row): keyed
//     Mörk-ClassicBetaPvP, with no gear and no guild.
//   - Oldtimer (us/normal), rated only 100 days ago: outside the window, no record.
func seedGolden(t *testing.T, pool *pgxpool.Pool, today time.Time) {
	t.Helper()
	ctx := context.Background()
	run := func(sql string, args ...any) {
		t.Helper()
		if _, err := pool.Exec(ctx, sql, args...); err != nil {
			t.Fatalf("%s: %v", sql, err)
		}
	}
	emails := []string{"dav2-yell@example.com", "dav2-fern@example.com", "dav2-hidden@example.com",
		"dav2-thoradin@example.com"}
	cleanup := func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id like 'dav2-%'`)
		pool.Exec(bg, `delete from fight_metrics where report_id like 'dav2-%'`)
		pool.Exec(bg, `delete from fights where report_id like 'dav2-%'`)
		pool.Exec(bg, `delete from reports where id like 'dav2-%'`)
		pool.Exec(bg, `delete from percentile_digests where encounter_id between 990701 and 990704`)
		pool.Exec(bg, `delete from addon_exports where character_key like '%/dav2-%'`)
		pool.Exec(bg, `delete from guild_characters where character_key like '%/dav2-%'`)
		pool.Exec(bg, `delete from guilds where name like 'Dav2 %'`)
		pool.Exec(bg, `update users set main_character_key = null where email like 'dav2-%'`)
		pool.Exec(bg, `delete from characters where key like '%/dav2-%'`)
		pool.Exec(bg, `delete from users where email like 'dav2-%'`)
	}
	cleanup()
	t.Cleanup(cleanup)

	user := func(email string, anonymize bool) int64 {
		t.Helper()
		var id int64
		if err := pool.QueryRow(ctx, `insert into users (email, anonymize) values ($1, $2) returning id`,
			email, anonymize).Scan(&id); err != nil {
			t.Fatal(err)
		}
		return id
	}
	yell, fern := user(emails[0], false), user(emails[1], false)
	hidden, thoradin := user(emails[2], true), user(emails[3], false)

	char := func(key, name string, userID int64) {
		t.Helper()
		parts := strings.SplitN(key, "/", 3)
		run(`insert into characters (key, region, ruleset, name, class, user_id) values ($1, $2, $3, $4, 'Warrior', $5)`,
			key, parts[0], parts[1], name, userID)
	}
	char("us/normal/dav2-obnoxious-yell", "Obnoxious Yell", yell)
	char("us/normal/dav2-little-yell", "Little Yell", yell)
	char("us/normal/dav2-quiet-fern", "Quiet Fern", fern)
	char("us/normal/dav2-hidden-hand", "Hidden Hand", hidden)
	char("us/normal/dav2-thoradin", "Thoradin", thoradin)
	run(`update characters set realm_name = 'Classic Beta PvE 2' where key = 'us/normal/dav2-obnoxious-yell'`)
	run(`update characters set equipped_item_level = 31 where key = 'us/normal/dav2-obnoxious-yell'`)
	run(`update characters set equipped_item_level = 22 where key = 'us/normal/dav2-quiet-fern'`)
	run(`update users set main_character_key = 'us/normal/dav2-obnoxious-yell' where id = $1`, yell)

	export := func(key, name string, userID int64, code string) {
		t.Helper()
		run(`insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at)
			values ($1, $2, 'us', 'normal', $3, $4, now())`, key, userID, name, code)
	}
	export("us/normal/dav2-obnoxious-yell", "Obnoxious Yell", yell, goldenGear)
	// Ten points in the second tree (Fury), none elsewhere.
	export("us/normal/dav2-quiet-fern", "Quiet Fern", fern, "FS1:1.60.1.70291:warrior:human:0/05320/0:|level=22|who=Quiet%20Fern:Classic%20Beta%20PvE%202")

	guild := func(name string) int64 {
		t.Helper()
		var id int64
		if err := pool.QueryRow(ctx, `insert into guilds (region, ruleset, name) values ('us', 'normal', $1) returning id`,
			name).Scan(&id); err != nil {
			t.Fatal(err)
		}
		return id
	}
	olympus, vanguard := guild("Dav2 Olympus"), guild("Dav2 Vanguard")
	member := func(guildID int64, key string, userID int64) {
		t.Helper()
		run(`insert into guild_characters (guild_id, character_key, user_id, verified_at) values ($1, $2, $3, now())`,
			guildID, key, userID)
	}
	member(olympus, "us/normal/dav2-obnoxious-yell", yell)
	member(olympus, "us/normal/dav2-quiet-fern", fern)
	member(olympus, "us/normal/dav2-hidden-hand", hidden)
	member(vanguard, "us/normal/dav2-thoradin", thoradin)
	// Not verified: never counted, never listed.
	run(`insert into guild_characters (guild_id, character_key, user_id) values ($1, 'us/normal/dav2-little-yell', $2)`,
		vanguard, yell)

	night := func(daysAgo int) time.Time { return today.Add(time.Duration(-daysAgo)*24*time.Hour + 20*time.Hour) }
	report := func(id string, guildID *int64, at time.Time) {
		t.Helper()
		run(`insert into reports (id, guild_id, title, visibility, zone, status, created_at)
			values ($1, $2, 'Molten Core', 'public', 'Molten Core', 'complete', $3)`, id, guildID, at)
	}
	report("dav2-a", &olympus, night(30))
	report("dav2-b", &olympus, night(20))
	report("dav2-c", &olympus, night(10))
	report("dav2-d", &olympus, night(5))
	report("dav2-eu", nil, night(3))
	report("dav2-old", nil, night(100))

	fight := func(reportID string, index, encounter int, name string, kill bool, at time.Time) {
		t.Helper()
		run(`insert into fights (report_id, fight_index, encounter_id, name, difficulty, kill, start_ms)
			values ($1, $2, $3, $4, 3, $5, $6)`, reportID, index, encounter, name, kill, at.UnixMilli())
	}
	rated := func(reportID string, index int, key, name, class, spec string, encounter int, kill bool,
		overall float64, discipline bool, at time.Time) {
		t.Helper()
		if err := db.EnsureRatingsPartition(ctx, pool, at); err != nil {
			t.Fatal(err)
		}
		run(`insert into rating_scores (report_id, fight_index, player_key, player_name, class, spec, role,
			   encounter_id, difficulty, kill, overall, overall_uncapped, overall_capped, components, model_version,
			   fought_at, insufficient, discipline)
			 values ($1, $2, $3, $4, $5, $6, 'dps', $7, 3, $8, $9, $9, false, '[]', 'test', $10, false, $11)`,
			reportID, index, key, name, class, spec, encounter, kill, overall, at, discipline)
	}
	const lucifron, magmadar, gehennas, garr = goldenEncounter, goldenEncounter + 1, goldenEncounter + 2, goldenEncounter + 3
	a, b, c, d := night(30), night(20), night(10), night(5)
	fight("dav2-a", 1, lucifron, "Lucifron", true, a)
	fight("dav2-b", 1, lucifron, "Lucifron", true, b)
	fight("dav2-b", 2, magmadar, "Magmadar", true, b.Add(30*time.Minute))
	fight("dav2-c", 1, gehennas, "Gehennas", false, c)
	fight("dav2-d", 1, magmadar, "Magmadar", true, d)
	fight("dav2-eu", 1, gehennas, "Gehennas", false, night(3))
	fight("dav2-old", 1, lucifron, "Lucifron", true, night(100))

	const yellKey = "us/normal/dav2-obnoxious-yell"
	rated("dav2-a", 1, yellKey, "Obnoxious Yell", "Warrior", "Arms", lucifron, true, 71, true, a)
	rated("dav2-b", 1, yellKey, "Obnoxious Yell", "Warrior", "Arms", lucifron, true, 64, false, b)
	rated("dav2-b", 2, yellKey, "Obnoxious Yell", "Warrior", "Arms", magmadar, true, 58, true, b.Add(30*time.Minute))
	rated("dav2-c", 1, yellKey, "Obnoxious Yell", "Warrior", "Arms", gehennas, false, 40, true, c)
	rated("dav2-d", 1, "us/normal/dav2-hidden-hand", "Hidden Hand", "Rogue", "Combat", magmadar, true, 90, true, d)
	rated("dav2-eu", 1, "eu/pvp/dav2-mörk", "Mörk-ClassicBetaPvP", "Mage", "Fire", gehennas, false, 88, true, night(3))
	rated("dav2-old", 1, "us/normal/dav2-oldtimer", "Oldtimer", "Warrior", "Fury", lucifron, true, 80, true, night(100))

	// The ranked kill the parses component places: one Lucifron kill at 71 DPS against a
	// curve of 1..100, and the alt's own kill of a boss the main never killed.
	curve := digest.New()
	for v := 1; v <= 100; v++ {
		curve.Add(float64(v))
	}
	raw, err := curve.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	run(`insert into percentile_digests (encounter_id, difficulty, spec, phase, metric, digest, n)
		values ($1, 3, 'Arms', '', 'dps', $2, 100)`, lucifron, raw)
	metric := func(fightIndex int, key, name string, encounter int, dps float64) {
		t.Helper()
		if err := db.EnsureMetricsPartition(ctx, pool, a); err != nil {
			t.Fatal(err)
		}
		run(`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec,
			role, metric_dps, encounter_id, difficulty, kill, phase, fought_at)
			values ('dav2-a', $1, $2, $3, 'Warrior', 'Arms', 'dps', $4, $5, 3, true, '', $6)`,
			fightIndex, key, name, dps, encounter, a)
	}
	metric(1, yellKey, "Obnoxious Yell", lucifron, 71)
	metric(5, "us/normal/dav2-little-yell", "Little Yell", garr, 50)
}

// TestGoldenDataLuaMatchesContractFive runs the real job over the seeded rows and checks
// the file it renders against contract 5's fixture byte for byte, once the run's own dates
// are moved onto the fixture's.
func TestGoldenDataLuaMatchesContractFive(t *testing.T) {
	pool := testPool(t)
	data := useRepoData(t)
	ctx := context.Background()
	now := time.Now().UTC()
	today := now.Truncate(24 * time.Hour)
	seedGolden(t, pool, today)

	result, err := Run(ctx, Deps{Store: &Store{Pool: pool}, Trees: data, Build: "1.60.1.70291", Now: now})
	if err != nil {
		t.Fatal(err)
	}
	got := normaliseDates(result.Rendered["Data.lua"], today, fixtureToday)

	want, err := os.ReadFile(goldenFixturePath)
	if err != nil {
		t.Fatal(err)
	}
	// A checkout with CRLF endings must never fail this on line endings alone.
	want = bytes.ReplaceAll(want, []byte("\r\n"), []byte("\n"))
	if !bytes.Equal(got, want) {
		t.Errorf("Data.lua does not match contract 5's fixture.\ngot:\n%s\nwant:\n%s", got, want)
	}
	if result.Characters != 3 || result.Guilds != 2 || result.SkippedRows != 1 {
		t.Errorf("result = %d characters, %d guilds, %d skipped; want 3, 2 and 1 (Thoradin has no realm)",
			result.Characters, result.Guilds, result.SkippedRows)
	}
	assertValidLua(t, result.Rendered["Data.lua"])
}

// A character in a file past the 4 MiB split keeps its whole record: the region file
// carries exactly the lines the single file does.
func TestASplitFileKeepsWholeRecords(t *testing.T) {
	pool := testPool(t)
	data := useRepoData(t)
	ctx := context.Background()
	now := time.Now().UTC()
	seedGolden(t, pool, now.Truncate(24*time.Hour))

	d, _, err := collect(ctx, Deps{Store: &Store{Pool: pool}, Trees: data, Build: "1.60.1.70291", Now: now})
	if err != nil {
		t.Fatal(err)
	}
	single, err := Render(d)
	if err != nil {
		t.Fatal(err)
	}
	split, err := renderWithLimit(d, 1)
	if err != nil {
		t.Fatal(err)
	}
	for file, opening := range map[string]string{
		"Data-us.lua": `["Obnoxious Yell-ClassicBetaPvE2"] = {`,
		"Data-eu.lua": `["Mörk-ClassicBetaPvP"] = {`,
	} {
		record := recordLines(t, single["Data.lua"], opening)
		if !bytes.Contains(split[file], record) {
			t.Errorf("%s does not carry the whole record:\n%s\nwant within it:\n%s", file, split[file], record)
		}
	}
	if guild := recordLines(t, single["Data.lua"], `["Dav2 Olympus"] = {`); !bytes.Contains(split["Data-us.lua"], guild) {
		t.Errorf("Data-us.lua does not carry Dav2 Olympus's whole record:\n%s", split["Data-us.lua"])
	}
	for name, body := range split {
		assertValidLuaChunk(t, name, body)
	}
}

// recordLines is one record of a rendered file: the line that opens it through the line
// that closes it.
func recordLines(t *testing.T, file []byte, opening string) []byte {
	t.Helper()
	start := bytes.Index(file, []byte(opening))
	if start < 0 {
		t.Fatalf("no %s in:\n%s", opening, file)
	}
	const closing = "\n\t\t},\n"
	end := bytes.Index(file[start:], []byte(closing))
	if end < 0 {
		t.Fatalf("%s never closes", opening)
	}
	return file[start : start+end+len(closing)]
}

var isoDate = regexp.MustCompile(`"(\d{4}-\d{2}-\d{2})"`)

// normaliseDates moves every quoted YYYY-MM-DD in file from the run's own calendar onto
// the fixture's: a date n days before today becomes the date n days before fixtureDay.
func normaliseDates(file []byte, today, fixtureDay time.Time) []byte {
	return isoDate.ReplaceAllFunc(file, func(m []byte) []byte {
		at, err := time.Parse(time.DateOnly, string(m[1:len(m)-1]))
		if err != nil {
			return m
		}
		return []byte(`"` + fixtureDay.Add(at.Sub(today)).Format(time.DateOnly) + `"`)
	})
}

// assertValidLua runs the rendered file through the `lua` interpreter when one is on PATH,
// asserting it parses and defines a format 2 global. Without one, it falls back to a
// minimal structural check and says so.
func assertValidLua(t *testing.T, data []byte) {
	t.Helper()
	path, err := exec.LookPath("lua")
	if err != nil {
		t.Log("no `lua` interpreter on PATH; falling back to a minimal structural check")
		if !bytes.Contains(data, []byte("ForeverSixtyData = {")) || !bytes.Contains(data, []byte("\n}")) {
			t.Errorf("does not look like a well-formed ForeverSixtyData table: %s", data)
		}
		return
	}
	tmp := filepath.Join(t.TempDir(), "check.lua")
	script := string(data) + "\nassert(type(ForeverSixtyData) == \"table\" and ForeverSixtyData.format == 2)\n"
	if err := os.WriteFile(tmp, []byte(script), 0o600); err != nil {
		t.Fatal(err)
	}
	if out, err := exec.Command(path, tmp).CombinedOutput(); err != nil {
		t.Errorf("lua rejected the rendered file: %v\n%s", err, out)
	}
}

// assertValidLuaChunk checks that one file of a split render parses, when `luac` is on
// PATH (a region file reads ForeverSixtyData and cannot run on its own).
func assertValidLuaChunk(t *testing.T, name string, data []byte) {
	t.Helper()
	path, err := exec.LookPath("luac")
	if err != nil {
		return
	}
	tmp := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		t.Fatal(err)
	}
	if out, err := exec.Command(path, "-p", tmp).CombinedOutput(); err != nil {
		t.Errorf("luac rejected %s: %v\n%s", name, err, out)
	}
}
