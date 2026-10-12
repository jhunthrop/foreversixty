package rankings

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// The character page's identity fields, score history and anonymised 404 (Task L6.6), read
// through the real endpoint (Mount) over seeded rows. Every row these tests write carries
// the "cx66" prefix (character slugs "cx66-...", report ids "cx66-...", emails "cx66-...",
// guild names "Cx66 ..."), and the cleanup deletes only rows with that prefix.

type extras struct {
	t    *testing.T
	pool *pgxpool.Pool
	srv  *httptest.Server
}

func newExtras(t *testing.T) *extras {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	cleanup := func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id like 'cx66-%'`)
		pool.Exec(bg, `delete from fight_metrics where report_id like 'cx66-%'`)
		pool.Exec(bg, `delete from reports where id like 'cx66-%'`)
		pool.Exec(bg, `delete from addon_exports where character_key like '%/cx66-%'`)
		pool.Exec(bg, `delete from guild_characters where character_key like '%/cx66-%'`)
		pool.Exec(bg, `delete from guilds where name like 'Cx66 %'`)
		pool.Exec(bg, `delete from characters where key like '%/cx66-%'`)
		pool.Exec(bg, `delete from users where email like 'cx66-%'`)
	}
	cleanup()
	t.Cleanup(cleanup)

	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, &Service{Store: &Store{Pool: pool, Trees: data}, Log: slog.New(slog.NewTextHandler(io.Discard, nil))})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return &extras{t: t, pool: pool, srv: srv}
}

func (x *extras) run(sql string, args ...any) {
	x.t.Helper()
	if _, err := x.pool.Exec(context.Background(), sql, args...); err != nil {
		x.t.Fatalf("%s: %v", sql, err)
	}
}

func (x *extras) user(email string, anonymize bool) int64 {
	x.t.Helper()
	var id int64
	if err := x.pool.QueryRow(context.Background(),
		`insert into users (email, anonymize) values ($1, $2) returning id`, email, anonymize).Scan(&id); err != nil {
		x.t.Fatal(err)
	}
	return id
}

// char writes a us/normal character row (class Warrior) and returns its key.
func (x *extras) char(slug, name string, userID *int64) string {
	x.t.Helper()
	key := "us/normal/" + slug
	x.run(`insert into characters (key, region, ruleset, name, class, user_id) values ($1, 'us', 'normal', $2, 'Warrior', $3)`,
		key, name, userID)
	return key
}

func (x *extras) export(key string, userID int64, code string, at time.Time) {
	x.t.Helper()
	x.run(`insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at)
		values ($1, $2, 'us', 'normal', 'x', $3, $4)`, key, userID, code, at)
}

func (x *extras) guild(name string) int64 {
	x.t.Helper()
	var id int64
	if err := x.pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us', 'normal', $1) returning id`, name).Scan(&id); err != nil {
		x.t.Fatal(err)
	}
	return id
}

func (x *extras) member(guildID int64, key string, userID int64, rank string, verified bool) {
	x.t.Helper()
	x.run(`insert into guild_characters (guild_id, character_key, user_id, rank, verified_at)
		values ($1, $2, $3, $4, case when $5 then now() end)`, guildID, key, userID, rank, verified)
}

func (x *extras) report(id, visibility string) {
	x.t.Helper()
	x.run(`insert into reports (id, title, visibility, zone, status) values ($1, 'Molten Core', $2, 'Molten Core', 'complete')`,
		id, visibility)
}

// rated writes one rating_scores row; score nil is a fight rated before migration 0035,
// with no snapshot.
func (x *extras) rated(reportID string, index int, key, unit string, at time.Time, score *int, tier *string) {
	x.t.Helper()
	if err := db.EnsureRatingsPartition(context.Background(), x.pool, at); err != nil {
		x.t.Fatal(err)
	}
	x.run(`insert into rating_scores (report_id, fight_index, player_key, player_name, class, spec, role,
		   encounter_id, difficulty, kill, overall, overall_uncapped, overall_capped, components, model_version,
		   fought_at, score, tier)
		 values ($1, $2, $3, $4, 'Warrior', 'Arms', 'dps', 990801, 3, true, 50, 50, false, '[]', 'test', $5, $6, $7)`,
		reportID, index, key, unit, at, score, tier)
}

// page reads GET /v1/characters/us/normal/<slug>: the status, the typed body, and the raw
// JSON of the body and of its character object (to see nulls as nulls).
func (x *extras) page(slug string) (int, Character, map[string]json.RawMessage, map[string]json.RawMessage, []byte) {
	x.t.Helper()
	res, err := http.Get(x.srv.URL + "/v1/characters/us/normal/" + slug)
	if err != nil {
		x.t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		x.t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, Character{}, nil, nil, body
	}
	var env struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(body, &env); err != nil {
		x.t.Fatal(err)
	}
	var c Character
	var raw, ref map[string]json.RawMessage
	if err := json.Unmarshal(env.Data, &c); err != nil {
		x.t.Fatal(err)
	}
	if err := json.Unmarshal(env.Data, &raw); err != nil {
		x.t.Fatal(err)
	}
	if err := json.Unmarshal(raw["character"], &ref); err != nil {
		x.t.Fatal(err)
	}
	return res.StatusCode, c, raw, ref, body
}

func TestCharacterPageCarriesWhoTheCharacterIs(t *testing.T) {
	x := newExtras(t)
	now := time.Now().UTC()
	owner := x.user("cx66-owner@example.com", false)
	olympus := x.guild("Cx66 Olympus")

	// Battle.net is the newest source: its level and race win over an older export's, and
	// its realm name is the first realm source. Verified officer of Cx66 Olympus.
	bnet := x.char("cx66-bnet", "Cx Bnet", &owner)
	x.run(`update characters set realm_name = 'Classic Beta PvE 2', level = 58, race = 'Night Elf', source = 'bnet',
		imported_at = $2 where key = $1`, bnet, now.Add(-time.Hour))
	x.export(bnet, owner, "FS1:test-1:warrior:dwarf:0/0/0:|level=40|who=Cx%20Bnet:Somewhere%20Else", now.Add(-48*time.Hour))
	x.member(olympus, bnet, owner, "officer", true)

	// The export is newer than the Battle.net import: its level and race win (the race slug
	// named by the build data), and with no Battle.net realm its who= realm is the realm. Its
	// guild membership is not verified, so it shows no guild.
	exported := x.char("cx66-export", "Cx Export", &owner)
	x.run(`update characters set level = 50, race = 'Human', source = 'bnet', imported_at = $2 where key = $1`,
		exported, now.Add(-72*time.Hour))
	x.export(exported, owner, "FS1:test-1:warrior:dwarf:0/0/0:|level=52|who=Cx%20Export:Classic%20Beta%20PvP", now.Add(-time.Hour))
	x.member(x.guild("Cx66 Vanguard"), exported, owner, "member", false)

	// The newer export carries no level= section and a race the build data does not know:
	// each falls back to the older Battle.net import, field by field.
	mixed := x.char("cx66-mixed", "Cx Mixed", &owner)
	x.run(`update characters set realm_name = 'Classic Beta PvE 2', level = 44, race = 'Human', source = 'bnet',
		imported_at = $2 where key = $1`, mixed, now.Add(-72*time.Hour))
	x.export(mixed, owner, "FS1:test-1:warrior:skyborne:0/0/0:", now.Add(-time.Hour))

	// Neither the import nor an export: the realm is the newest rated row's unit realm;
	// level, race and guild are unknown.
	logged := x.char("cx66-logged", "Cx Logged", nil)
	x.report("cx66-r1", "public")
	x.rated("cx66-r1", 1, logged, "Cx Logged-OldRealm", now.Add(-96*time.Hour), nil, nil)
	x.rated("cx66-r1", 2, logged, "Cx Logged-ClassicBetaPvE2", now.Add(-24*time.Hour), nil, nil)

	str := func(s string) *string { return &s }
	num := func(n int) *int { return &n }
	for _, tc := range []struct {
		slug         string
		level        *int
		race, realm  *string
		guild        *CharacterGuild
		nullFields   []string
		describeLine string
	}{
		{"cx66-bnet", num(58), str("Night Elf"), str("Classic Beta PvE 2"),
			&CharacterGuild{Name: "Cx66 Olympus", Rank: "officer"}, nil, "Battle.net newest"},
		{"cx66-export", num(52), str("Dwarf"), str("Classic Beta PvP"), nil, []string{"guild"}, "export newest"},
		{"cx66-mixed", num(44), str("Human"), str("Classic Beta PvE 2"), nil, []string{"guild"}, "export newest, fields missing"},
		{"cx66-logged", nil, nil, str("ClassicBetaPvE2"), nil, []string{"guild", "level", "race"}, "logged only"},
	} {
		status, c, _, ref, body := x.page(tc.slug)
		if status != http.StatusOK {
			t.Fatalf("%s: status %d: %s", tc.describeLine, status, body)
		}
		got := c.Character
		if !eqPtr(got.Level, tc.level) || !eqPtr(got.Race, tc.race) || !eqPtr(got.Realm, tc.realm) {
			t.Errorf("%s: level %s race %s realm %s; want %s %s %s", tc.describeLine,
				show(got.Level), show(got.Race), show(got.Realm), show(tc.level), show(tc.race), show(tc.realm))
		}
		if (got.Guild == nil) != (tc.guild == nil) || (got.Guild != nil && *got.Guild != *tc.guild) {
			t.Errorf("%s: guild %+v, want %+v", tc.describeLine, got.Guild, tc.guild)
		}
		for _, field := range tc.nullFields {
			if string(ref[field]) != "null" {
				t.Errorf("%s: %s = %s, want null (not omitted)", tc.describeLine, field, ref[field])
			}
		}
		if tc.guild != nil {
			var guild map[string]json.RawMessage
			if err := json.Unmarshal(ref["guild"], &guild); err != nil {
				t.Fatal(err)
			}
			if len(guild) != 2 || guild["name"] == nil || guild["rank"] == nil {
				t.Errorf("%s: guild is %s, want exactly { name, rank } and never an id", tc.describeLine, ref["guild"])
			}
		}
	}
}

func TestCharacterScoreHistoryIsTheLastSnapshotOfEachUTCDate(t *testing.T) {
	x := newExtras(t)
	key := x.char("cx66-history", "Cx History", nil)
	x.report("cx66-h1", "public")
	x.report("cx66-hp", "private")
	score := func(n int) *int { return &n }
	tier := func(s string) *string { return &s }
	utc := func(month, day, hour, minute int) time.Time {
		return time.Date(2026, time.Month(month), day, hour, minute, 0, 0, time.UTC)
	}

	// 2 March: two fights; the later one's snapshot is the day's.
	x.rated("cx66-h1", 1, key, "Cx History", utc(3, 2, 20, 0), score(400), tier("common"))
	x.rated("cx66-h1", 2, key, "Cx History", utc(3, 2, 22, 0), score(520), tier("uncommon"))
	// 23:30 on 2 March in New York is 3 March in UTC.
	newYork := time.FixedZone("EST", -5*3600)
	x.rated("cx66-h1", 3, key, "Cx History", time.Date(2026, 3, 2, 23, 30, 0, 0, newYork), score(900), tier("uncommon"))
	x.rated("cx66-h1", 4, key, "Cx History", utc(3, 10, 21, 0), score(1100), tier("rare"))
	// A fight in a private report is not on the public page's history.
	x.rated("cx66-hp", 1, key, "Cx History", utc(3, 11, 21, 0), score(1300), tier("rare"))
	// A fight rated before migration 0035 carries no snapshot and is left out.
	x.rated("cx66-h1", 5, key, "Cx History", utc(3, 12, 21, 0), nil, nil)

	status, c, raw, _, body := x.page("cx66-history")
	if status != http.StatusOK {
		t.Fatalf("status %d: %s", status, body)
	}
	want := []ScorePoint{
		{Date: "2026-03-02", Score: 520, Tier: "uncommon"},
		{Date: "2026-03-03", Score: 900, Tier: "uncommon"},
		{Date: "2026-03-10", Score: 1100, Tier: "rare"},
	}
	if fmt.Sprint(c.ScoreHistory) != fmt.Sprint(want) {
		t.Fatalf("score_history = %+v (raw %s), want %+v", c.ScoreHistory, raw["score_history"], want)
	}

	// A character with no snapshot at all has an empty list, not null.
	x.char("cx66-nohistory", "Cx Nohistory", nil)
	_, _, raw, _, _ = x.page("cx66-nohistory")
	if string(raw["score_history"]) != "[]" {
		t.Fatalf("score_history with no snapshots = %s, want []", raw["score_history"])
	}
}

func TestCharacterScoreHistoryKeepsTheNewestThirtyDates(t *testing.T) {
	x := newExtras(t)
	key := x.char("cx66-veteran", "Cx Veteran", nil)
	x.report("cx66-v1", "public")
	first := time.Date(2026, 2, 1, 20, 0, 0, 0, time.UTC)
	for i := 0; i < 32; i++ {
		s, tr := 100+i, "common"
		x.rated("cx66-v1", i+1, key, "Cx Veteran", first.AddDate(0, 0, i), &s, &tr)
	}
	_, c, _, _, _ := x.page("cx66-veteran")
	if len(c.ScoreHistory) != ScoreHistoryDates {
		t.Fatalf("%d entries, want %d", len(c.ScoreHistory), ScoreHistoryDates)
	}
	if got := c.ScoreHistory[0]; got.Date != "2026-02-03" || got.Score != 102 {
		t.Errorf("oldest kept = %+v, want 2026-02-03 at 102", got)
	}
	if got := c.ScoreHistory[len(c.ScoreHistory)-1]; got.Date != "2026-03-04" || got.Score != 131 {
		t.Errorf("newest = %+v, want 2026-03-04 at 131", got)
	}
}

func TestAnAnonymisedCharacterIsIndistinguishableFromAnUnloggedOne(t *testing.T) {
	x := newExtras(t)
	hidden := x.user("cx66-hidden@example.com", true)
	key := x.char("cx66-hidden", "Cx Hidden", &hidden)
	x.report("cx66-a1", "public")
	s, tr := 700, "uncommon"
	x.rated("cx66-a1", 1, key, "Cx Hidden", time.Now().UTC().Add(-time.Hour), &s, &tr)

	status, _, _, _, body := x.page("cx66-hidden")
	unloggedStatus, _, _, _, unloggedBody := x.page("cx66-nobody")
	if status != http.StatusNotFound || unloggedStatus != http.StatusNotFound {
		t.Fatalf("anonymised = %d, unlogged = %d; want 404 for both", status, unloggedStatus)
	}
	if strings.TrimSpace(string(body)) != strings.TrimSpace(string(unloggedBody)) {
		t.Fatalf("anonymised body %s differs from unlogged body %s", body, unloggedBody)
	}

	// The same character, once its account stops being anonymous, is served.
	x.run(`update users set anonymize = false where id = $1`, hidden)
	if status, _, _, _, body := x.page("cx66-hidden"); status != http.StatusOK {
		t.Fatalf("not anonymised = %d: %s", status, body)
	}
}

func eqPtr[T comparable](a, b *T) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}

func show[T any](p *T) string {
	if p == nil {
		return "null"
	}
	return fmt.Sprint(*p)
}
