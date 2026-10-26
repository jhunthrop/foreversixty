package guilds

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// raidChar is one character a raid test seeds: a guild_characters row plus its characters row.
type raidChar struct {
	key, name, class string
	verified         bool
}

// raidWorld is a guild with an officer (holding the plan unless built without), two members
// and their characters, over the HTTP harness.
type raidWorld struct {
	h                        *httpHarness
	gid                      int64
	officer, member, member2 int64
}

const (
	officerKey  = "us/hardcore/thornhide"
	memberKeyA  = "us/hardcore/kraggor"
	memberKeyB  = "us/hardcore/sunderfel"
	memberKeyC  = "us/hardcore/unverifiedalt"
	member2Key  = "us/hardcore/mirelle"
	raidZoneOny = "raid:onyxias-lair"
)

func newRaidWorld(t *testing.T, withPlan bool) raidWorld {
	t.Helper()
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Olympus")
	w := raidWorld{h: h, gid: gid}
	w.officer = w.seedRaider(t, "officer@example.com", "officer", raidChar{officerKey, "Thornhide", "druid", true})
	w.member = w.seedRaider(t, "member@example.com", "member",
		raidChar{memberKeyA, "Kraggor", "warrior", true},
		raidChar{memberKeyB, "Sunderfel", "mage", true},
		raidChar{memberKeyC, "Unverifiedalt", "rogue", false})
	w.member2 = w.seedRaider(t, "member2@example.com", "member", raidChar{member2Key, "Mirelle", "priest", true})
	h.svc.Entitlements = &entitlements.Store{Pool: h.pool}
	if withPlan {
		if _, err := h.pool.Exec(t.Context(), scopedGrantSQL, gid); err != nil {
			t.Fatal(err)
		}
	}
	return w
}

func (w raidWorld) seedRaider(t *testing.T, email, rank string, chars ...raidChar) int64 {
	t.Helper()
	uid := seedUser(t, w.h.pool, email)
	for _, c := range chars {
		seedCharacter(t, w.h.pool, w.gid, uid, c.key, rank, c.verified)
		if _, err := w.h.pool.Exec(t.Context(),
			`insert into characters (key, region, ruleset, name, class, user_id) values ($1, 'us', 'hardcore', $2, $3, $4)`,
			c.key, c.name, c.class, uid); err != nil {
			t.Fatal(err)
		}
	}
	recomputeMembership(t, w.h.pool, w.gid, uid)
	return uid
}

func (w raidWorld) as(uid int64) {
	if uid == 0 {
		w.h.actor = auth.Actor{}
		return
	}
	w.h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
}

func (w raidWorld) path(format string, args ...any) string {
	return fmt.Sprintf("/v1/guilds/%d", w.gid) + fmt.Sprintf(format, args...)
}

// raidBody is a valid create body at the given RFC 3339 start.
func raidBody(startsAt string, repeat int) string {
	return fmt.Sprintf(`{"zone":%q,"title":"","starts_at":%q,"duration_minutes":180,"notes":"Bring **flasks**",`+
		`"timezone":"America/Chicago","repeat_weeks":%d,"post_to_discord":false}`, raidZoneOny, startsAt, repeat)
}

// futureStart is a start a week out, on the minute.
func futureStart() string {
	return time.Now().Add(7 * 24 * time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339)
}

// createRaids posts as the officer and returns the created views.
func (w raidWorld) createRaids(t *testing.T, body string) []RaidView {
	t.Helper()
	w.as(w.officer)
	res := w.h.do(http.MethodPost, w.path("/raids"), body)
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("create = %d: %s", res.StatusCode, readBody(res))
	}
	var out struct {
		Raids []RaidView `json:"raids"`
	}
	w.h.data(res, &out)
	return out.Raids
}

func (w raidWorld) createRaid(t *testing.T) RaidView {
	t.Helper()
	return w.createRaids(t, raidBody(futureStart(), 0))[0]
}

func (w raidWorld) getRaid(t *testing.T, id int64) RaidView {
	t.Helper()
	res := w.h.do(http.MethodGet, w.path("/raids/%d", id), "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("get raid = %d: %s", res.StatusCode, readBody(res))
	}
	var v RaidView
	w.h.data(res, &v)
	return v
}

func (w raidWorld) putSignup(id int64, key, body string) *http.Response {
	return w.h.do(http.MethodPut, w.path("/raids/%d/signups/%s", id, key), body)
}

func readBody(res *http.Response) string {
	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)
	return string(b)
}

// apiError is a failure envelope's error half.
type apiError struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields"`
}

// expectError checks the status and code of a failure envelope and returns its error body.
func expectError(t *testing.T, res *http.Response, status int, code string) apiError {
	t.Helper()
	defer res.Body.Close()
	var env struct {
		Error apiError `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != status || env.Error.Code != code {
		t.Fatalf("got %d %q (%s), want %d %q", res.StatusCode, env.Error.Code, env.Error.Message, status, code)
	}
	return env.Error
}

// rawData decodes an envelope's data half into a generic map, to check which keys exist.
func rawData(t *testing.T, res *http.Response) map[string]json.RawMessage {
	t.Helper()
	defer res.Body.Close()
	var env struct {
		Data map[string]json.RawMessage `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	return env.Data
}

// seedPlayer is one character's row in a seeded fight.
type seedPlayer struct {
	key, name, class, spec, role string
	deaths                       int
}

// seedFight is one fight of a seeded night: encounterID 0 is trash (no rows per player).
type seedFight struct {
	encounterID int
	kill        bool
	offset      time.Duration
	duration    time.Duration
	players     []seedPlayer
}

// seedNight inserts a complete guild-visible report with log_zone zone created at createdAt,
// its fights (start_ms = createdAt + offset) and one fight_metrics row per player per fight.
func seedNight(t *testing.T, pool *pgxpool.Pool, guildID int64, reportID, zone string, createdAt time.Time, fights []seedFight) {
	t.Helper()
	ctx := context.Background()
	if err := db.EnsureMetricsPartition(ctx, pool, createdAt); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into reports (id, guild_id, visibility, status, log_zone, created_at)
		 values ($1, $2, 'guild', 'complete', nullif($3, ''), $4)`, reportID, guildID, zone, createdAt); err != nil {
		t.Fatal(err)
	}
	for i, f := range fights {
		seedOneFight(t, pool, reportID, createdAt, i, f)
	}
}

func seedOneFight(t *testing.T, pool *pgxpool.Pool, reportID string, createdAt time.Time, index int, f seedFight) {
	t.Helper()
	ctx := context.Background()
	at := createdAt.Add(f.offset)
	var encounter any
	if f.encounterID != 0 {
		encounter = f.encounterID
	}
	deaths := 0
	for _, p := range f.players {
		deaths += p.deaths
	}
	if _, err := pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, encounter_id, kill, start_ms, duration_ms, deaths, players)
		 values ($1, $2, 'Fight', $3, $4, $5, $6, $7, $8)`,
		reportID, index, encounter, f.kill, at.UnixMilli(), f.duration.Milliseconds(), deaths, []string{l68GUID}); err != nil {
		t.Fatal(err)
	}
	if f.encounterID == 0 {
		return
	}
	for _, p := range f.players {
		if _, err := pool.Exec(ctx,
			`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec, role, metric_dps,
				deaths, encounter_id, duration_ms, kill, fought_at)
			 values ($1, $2, $3, $4, $5, $6, $7, 100, $8, $9, $10, $11, $12)`,
			reportID, index, p.key, p.name, p.class, p.spec, p.role, p.deaths, f.encounterID,
			f.duration.Milliseconds(), f.kill, at); err != nil {
			t.Fatal(err)
		}
	}
}
