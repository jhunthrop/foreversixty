package activity

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

type fakeImports struct{ calls int }

func (f *fakeImports) RequestFromImport(context.Context, int64, string) (string, error) {
	f.calls++
	return "queued", nil
}

func serve(t *testing.T, s *Store, imports ImportSimmer, userID int64, url string) (*httptest.ResponseRecorder, map[string]any) {
	t.Helper()
	svc := &Service{Store: s, Guilds: &auth.Store{Pool: s.Pool}, Home: &Home{Store: s, Imports: imports}, Now: func() time.Time { return testNow }}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	if userID != 0 {
		req = req.WithContext(auth.WithActor(req.Context(), auth.Actor{UserID: userID, Method: "session"}))
	}
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	var body map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &body)
	return rec, body
}

func TestFeedRouteAnswersAndRefuses(t *testing.T) {
	s := testStore(t)
	mustAppend(t, s, ev(KindLogUpload, "public", testNow.Add(-time.Minute)))
	rec, body := serve(t, s, nil, 0, "/v1/feed")
	if rec.Code != 200 || !strings.Contains(rec.Header().Get("Cache-Control"), "public, max-age=30") {
		t.Fatalf("anonymous: %d %s %v", rec.Code, rec.Header().Get("Cache-Control"), body)
	}
	if rows := body["data"].(map[string]any)["rows"].([]any); len(rows) != 1 {
		t.Fatalf("rows: %v", rows)
	}
	for url, code := range map[string]int{
		"/v1/feed?filter=guild": 401, "/v1/feed?filter=nope": 400, "/v1/feed?filter=spec": 400,
		"/v1/feed?spec=warrior-nope": 400, "/v1/feed?band=25": 400, "/v1/feed?cursor=zz": 400, "/v1/feed?character=nope": 400, "/v1/feed?character=us/normal/thalrin": 200, "/v1/home?character=x": 401,
	} {
		if rec, _ := serve(t, s, nil, 0, url); rec.Code != code {
			t.Errorf("%s: %d, want %d", url, rec.Code, code)
		}
	}
	u := user(t, s, "a#1", false)
	if rec, _ := serve(t, s, nil, u, "/v1/feed?filter=guild"); rec.Code != 200 || !strings.HasPrefix(rec.Header().Get("Cache-Control"), "private") {
		t.Fatalf("signed in guild filter: %d %s", rec.Code, rec.Header().Get("Cache-Control"))
	}
}

// GET /v1/home: the import sim, this week's grade against last week's, and the movers
// since last week; another account's character is not found.
func TestHomeAnswersTheCharactersBlocks(t *testing.T) {
	s := testStore(t)
	u := user(t, s, "a#1", false)
	stranger := user(t, s, "b#1", false)
	char(t, s, "us/normal/gohogs", "Go Hogs", "Warrior", 20, u)
	exec(t, s.Pool, `update characters set faction = 'horde' where key = 'us/normal/gohogs'`)
	sim(t, s, "s1", "us/normal/gohogs", u, "import", 26.2, 20, "bare", testNow.Add(-time.Hour))
	week, prev := WeekOf(testNow), PreviousWeek(testNow)
	exec(t, s.Pool, `insert into tier_grade_weeks (week, spec, band, faction, role, grade, place, of_n) values
		($1, 'warrior-fury', 20, 'horde', 'dps', 'A', 4, 23), ($2, 'warrior-fury', 20, 'horde', 'dps', 'B', 7, 23)`, week, prev)
	exec(t, s.Pool, `insert into spec_board_weeks (week, spec, band, character_key, name, class, rank, score) values
		($1, 'warrior-fury', 20, 'k1', 'Grukk', 'warrior', 3, 80), ($1, 'warrior-fury', 20, 'k2', 'Brakka', 'warrior', 1, 90),
		($2, 'warrior-fury', 20, 'k1', 'Grukk', 'warrior', 9, 70), ($2, 'warrior-fury', 20, 'k2', 'Brakka', 'warrior', 1, 90)`, week, prev)

	rec, body := serve(t, s, nil, u, "/v1/home?character=us/normal/gohogs")
	if rec.Code != 200 {
		t.Fatalf("home: %d %v", rec.Code, body)
	}
	data := body["data"].(map[string]any)
	is := data["import_sim"].(map[string]any)
	sw := data["spec_week"].(map[string]any)
	mv := data["movers"].(map[string]any)
	if is["state"] != "done" || is["dps"].(float64) != 26.2 {
		t.Fatalf("import sim: %v", is)
	}
	if sw["grade"] != "A" || sw["last_week_grade"] != "B" || sw["place"].(float64) != 4 {
		t.Fatalf("spec week: %v", sw)
	}
	rows := mv["rows"].([]any)
	if mv["baseline"] != true || len(rows) != 1 || rows[0].(map[string]any)["delta"].(float64) != 6 {
		t.Fatalf("movers: %v", mv)
	}
	if data["guild_sims"] != nil {
		t.Fatalf("no guild, no ladder: %v", data["guild_sims"])
	}
	// Level 20 with no guild (ruling R4): the countdown is one line after the feed and the
	// guild ladder is the one-line Join a guild row.
	var order []string
	for _, sec := range data["sections"].([]any) {
		order = append(order, sec.(map[string]any)["type"].(string))
	}
	if got := strings.Join(order, " "); got != "head upgrades feed countdowns guild_sims_cta spec_week movers next" {
		t.Fatalf("sections: %s", got)
	}
	if rec, _ := serve(t, s, nil, stranger, "/v1/home?character=us/normal/gohogs"); rec.Code != 404 {
		t.Fatalf("another account's character: %d", rec.Code)
	}
}

// A character with gear and no import sim yet (imported before the auto-sim) gets one.
func TestHomeStartsTheFirstImportSim(t *testing.T) {
	s := testStore(t)
	u := user(t, s, "a#1", false)
	char(t, s, "us/normal/gohogs", "Go Hogs", "Warrior", 20, u)
	exec(t, s.Pool, `insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		values ('us/normal/gohogs', $1, 'us', 'normal', 'Go Hogs', $2, 'blizzard', now(), now())`, u, sampleExport)
	imports := &fakeImports{}
	_, body := serve(t, s, imports, u, "/v1/home?character=us/normal/gohogs")
	is := body["data"].(map[string]any)["import_sim"].(map[string]any)
	if imports.calls != 1 || is["state"] != "pending" {
		t.Fatalf("import sim %v, %d calls", is, imports.calls)
	}
}

const sampleExport = "FS1:1.60.1.70790:warrior:orc:0/0505/0:head=16866,chest=11726|level=20"

// The time-aware Home's guild facts: membership, the raid zone, the next raid not yet over and the
// last raid started; a dungeon run is not a raid, and a character on no roster has no guild.
func TestGuildRaids(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	g, users := guildOf(t, s, 1)
	v, err := s.guildRaids(ctx, "us/normal/ma", testNow)
	if err != nil || !v.InGuild || v.NextRaid != nil || v.LastRaid != nil || v.Zone != nil {
		t.Fatalf("no raids: %+v, %v", v, err)
	}
	raid := func(zone string, start time.Time) {
		exec(t, s.Pool, `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, created_by, created_at)
			values ($1, $2, 'Night', $3, 180, $4, $5)`, g, zone, start, users[0], testNow.Add(-72*time.Hour))
	}
	raid("raid:onyxias-lair", testNow.Add(-26*time.Hour))
	raid("raid:onyxias-lair", testNow.Add(6*time.Hour))
	raid("dungeon:ubrs", testNow.Add(2*time.Hour))
	exec(t, s.Pool, `update guilds set raid_timezone = 'America/Chicago' where id = $1`, g)
	v, err = s.guildRaids(ctx, "us/normal/ma", testNow)
	if err != nil || v.Zone == nil || v.Zone.String() != "America/Chicago" {
		t.Fatalf("zone: %+v, %v", v, err)
	}
	if v.NextRaid == nil || !v.NextRaid.Equal(testNow.Add(6*time.Hour)) {
		t.Fatalf("next raid %v", v.NextRaid)
	}
	if v.LastRaid == nil || !v.LastRaid.Equal(testNow.Add(-26*time.Hour)) {
		t.Fatalf("last raid %v", v.LastRaid)
	}
	if v, _ := s.guildRaids(ctx, "us/normal/nobody", testNow); v.InGuild {
		t.Fatal("a character with no guild is in one")
	}
}
