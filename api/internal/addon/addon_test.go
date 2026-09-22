package addon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

type harness struct {
	t      *testing.T
	store  *Store
	builds *builds.Store
	trees  *trees.Data
	pool   *pgxpool.Pool
	server *httptest.Server
	actor  auth.Actor
	owner  int64
}

func newHarness(t *testing.T) *harness {
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
	if _, err := pool.Exec(context.Background(), `truncate users, builds cascade`); err != nil {
		t.Fatal(err)
	}
	owner, err := (&auth.Store{Pool: pool}).UpsertEmailUser(context.Background(), "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	h := &harness{t: t, pool: pool, store: &Store{Pool: pool}, owner: owner.ID,
		builds: &builds.Store{Pool: pool}, trees: data}
	h.actor = auth.Actor{UserID: owner.ID, Role: "user", Method: "device", DeviceID: "device-1"}
	mux := http.NewServeMux()
	Mount(mux, &Service{Store: h.store, Builds: h.builds, Data: data, Accounts: &auth.Store{Pool: pool},
		Log: slog.New(slog.NewTextHandler(io.Discard, nil))}, 0)
	h.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r.WithContext(auth.WithActor(r.Context(), h.actor)))
	}))
	t.Cleanup(h.server.Close)
	return h
}

func (h *harness) do(method, path, body string) *http.Response {
	h.t.Helper()
	r, err := http.NewRequest(method, h.server.URL+path, strings.NewReader(body))
	if err != nil {
		h.t.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		h.t.Fatal(err)
	}
	return res
}

func (h *harness) data(res *http.Response, into any) {
	h.t.Helper()
	defer res.Body.Close()
	var env struct {
		OK   bool            `json:"ok"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		h.t.Fatal(err)
	}
	if !env.OK {
		h.t.Fatalf("envelope reports failure for %s", res.Request.URL)
	}
	if into != nil {
		if err := json.Unmarshal(env.Data, into); err != nil {
			h.t.Fatal(err)
		}
	}
}

func TestExportsAreStoredAndReplacedPerCharacter(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodPost, "/v1/addon/exports",
		`{"characters":[{"name":"Baelgrim","ruleset":"hardcore","region":"us","export":"FS1:aaa"}]}`)
	var out struct {
		Stored int `json:"stored"`
	}
	h.data(res, &out)
	if out.Stored != 1 {
		t.Fatalf("stored = %d", out.Stored)
	}
	res = h.do(http.MethodPost, "/v1/addon/exports",
		`{"characters":[{"name":"Baelgrim","ruleset":"hardcore","region":"us","export":"FS1:bbb"}]}`)
	res.Body.Close()

	exports, err := h.store.Exports(t.Context(), h.owner)
	if err != nil {
		t.Fatal(err)
	}
	if len(exports) != 1 || exports[0].Export != "FS1:bbb" {
		t.Fatalf("exports = %+v, want the newer one only", exports)
	}
	if exports[0].Ruleset != "hardcore" {
		t.Fatalf("ruleset = %q", exports[0].Ruleset)
	}
}

func TestExportsRefuseNonsense(t *testing.T) {
	h := newHarness(t)
	for name, body := range map[string]string{
		"not json":      `{`,
		"no characters": `{"characters":[]}`,
		"no name":       `{"characters":[{"name":"","ruleset":"pvp","region":"us","export":"FS1:a"}]}`,
		"no export":     `{"characters":[{"name":"Baelgrim","ruleset":"pvp","region":"us","export":""}]}`,
		"bad region":    `{"characters":[{"name":"Baelgrim","ruleset":"pvp","region":"mars","export":"FS1:a"}]}`,
		"bad name":      `{"characters":[{"name":"Bael\tgrim\"]","ruleset":"pvp","region":"us","export":"FS1:a"}]}`,
	} {
		res := h.do(http.MethodPost, "/v1/addon/exports", body)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s = %d, want 400", name, res.StatusCode)
		}
	}
	long := `{"characters":[{"name":"Baelgrim","ruleset":"pvp","region":"us","export":"` +
		strings.Repeat("x", MaxExportLen+1) + `"}]}`
	res := h.do(http.MethodPost, "/v1/addon/exports", long)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("an over-long export = %d, want 400", res.StatusCode)
	}
}

// TestPutExportsClaimsAnUnclaimedCharacter is the first of the three
// cases the Task 5 review pointed at: a character_key nobody has synced
// before is claimed by whoever syncs it first.
func TestPutExportsClaimsAnUnclaimedCharacter(t *testing.T) {
	h := newHarness(t)
	if err := h.store.PutExports(context.Background(), h.owner,
		[]Export{{Name: "Baelgrim", Ruleset: "hardcore", Region: "us", Export: "FS1:aaa"}}); err != nil {
		t.Fatal(err)
	}
	exports, err := h.store.Exports(context.Background(), h.owner)
	if err != nil || len(exports) != 1 || exports[0].Export != "FS1:aaa" {
		t.Fatalf("exports = %+v, err = %v", exports, err)
	}
}

// TestPutExportsRefreshesTheSameAccountsOwnCharacter is the second case:
// re-syncing a character already owned by the same account is a normal
// update, not a claim.
func TestPutExportsRefreshesTheSameAccountsOwnCharacter(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner,
		[]Export{{Name: "Baelgrim", Ruleset: "hardcore", Region: "us", Export: "FS1:aaa"}}); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, h.owner,
		[]Export{{Name: "Baelgrim", Ruleset: "hardcore", Region: "us", Export: "FS1:bbb"}}); err != nil {
		t.Fatalf("re-syncing as the current owner should succeed: %v", err)
	}
	exports, err := h.store.Exports(ctx, h.owner)
	if err != nil || len(exports) != 1 || exports[0].Export != "FS1:bbb" {
		t.Fatalf("exports = %+v, err = %v, want the refreshed export", exports, err)
	}
}

// TestPutExportsRefusesToStealAnotherAccountsCharacter is the third
// case: a different account may not take over a character_key someone
// else already owns, and the stored row must be untouched by the
// refused attempt.
func TestPutExportsRefusesToStealAnotherAccountsCharacter(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	stranger, err := (&auth.Store{Pool: h.pool}).UpsertEmailUser(ctx, "stranger@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, h.owner,
		[]Export{{Name: "Baelgrim", Ruleset: "hardcore", Region: "us", Export: "FS1:aaa"}}); err != nil {
		t.Fatal(err)
	}

	err = h.store.PutExports(ctx, stranger.ID,
		[]Export{{Name: "Baelgrim", Ruleset: "hardcore", Region: "us", Export: "FS1:stolen"}})
	if !errors.Is(err, ErrCharacterClaimed) {
		t.Fatalf("err = %v, want ErrCharacterClaimed", err)
	}

	var storedUserID int64
	var storedExport string
	if err := h.pool.QueryRow(ctx,
		`select user_id, export from addon_exports where character_key = $1`, "us/hardcore/baelgrim").
		Scan(&storedUserID, &storedExport); err != nil {
		t.Fatal(err)
	}
	if storedUserID != h.owner || storedExport != "FS1:aaa" {
		t.Fatalf("stored row changed: user_id=%d export=%q, want owner=%d export=FS1:aaa",
			storedUserID, storedExport, h.owner)
	}
}

// TestExportsRefuseStealingAnAlreadyClaimedCharacter covers the HTTP
// path onto the same guard: the device route answers 409 rather than
// 200 when the store refuses a claim.
func TestExportsRefuseStealingAnAlreadyClaimedCharacter(t *testing.T) {
	h := newHarness(t)
	stranger, err := (&auth.Store{Pool: h.pool}).UpsertEmailUser(context.Background(), "stranger@example.com")
	if err != nil {
		t.Fatal(err)
	}
	res := h.do(http.MethodPost, "/v1/addon/exports",
		`{"characters":[{"name":"Baelgrim","ruleset":"hardcore","region":"us","export":"FS1:aaa"}]}`)
	res.Body.Close()

	h.actor = auth.Actor{UserID: stranger.ID, Role: "user", Method: "device", DeviceID: "device-2"}
	res = h.do(http.MethodPost, "/v1/addon/exports",
		`{"characters":[{"name":"Baelgrim","ruleset":"hardcore","region":"us","export":"FS1:stolen"}]}`)
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("status = %d, want 409 for a character claimed by a different account", res.StatusCode)
	}
}

// seedBuild stores a build the inbox can name and encode.
func (h *harness) seedBuild() string {
	h.t.Helper()
	b, err := builds.New(builds.Input{
		ClassID: 1, RaceID: 1, TreeVersion: "test-1",
		PointOrder: []int{101, 101, 102}, Gear: map[string]int{"head": 12640},
		Title: "Arms leveling",
	})
	if err != nil {
		h.t.Fatal(err)
	}
	if _, _, err := h.builds.Save(context.Background(), b, nil); err != nil {
		h.t.Fatal(err)
	}
	return b.ID
}

func TestAnUnrecognisedRulesetIsNormalisedRatherThanRefused(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodPost, "/v1/addon/exports",
		`{"characters":[{"name":"Baelgrim","ruleset":"Nightslayer","region":"us","export":"FS1:aaa"}]}`)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want the export stored", res.StatusCode)
	}
	exports, err := h.store.Exports(t.Context(), h.owner)
	if err != nil {
		t.Fatal(err)
	}
	if len(exports) != 1 || exports[0].Ruleset != character.RulesetNormal {
		t.Fatalf("exports = %+v, want the default ruleset", exports)
	}
}

func TestTheInboxIsQueuedOnTheSiteAndReadByTheDevice(t *testing.T) {
	h := newHarness(t)
	buildID := h.seedBuild()
	// The companion's read comes first and is empty.
	var empty struct {
		Builds []InboxEntry `json:"builds"`
	}
	h.data(h.do(http.MethodGet, "/v1/addon/inbox", ""), &empty)
	if len(empty.Builds) != 0 {
		t.Fatalf("inbox = %+v, want empty", empty.Builds)
	}

	// The site queues a build with a session.
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, "/v1/addon/inbox",
		`{"build_id":"`+buildID+`","character_key":"us/hardcore/baelgrim"}`)
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("status = %d, want 201", res.StatusCode)
	}
	res.Body.Close()

	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "device", DeviceID: "device-1"}
	var out struct {
		Builds []InboxEntry `json:"builds"`
	}
	h.data(h.do(http.MethodGet, "/v1/addon/inbox", ""), &out)
	if len(out.Builds) != 1 || out.Builds[0].ID != buildID {
		t.Fatalf("inbox = %+v", out.Builds)
	}
	entry := out.Builds[0]
	if entry.Character != "us/hardcore/baelgrim" {
		t.Fatalf("character = %q", entry.Character)
	}
	if entry.Name != "Arms leveling" {
		t.Fatalf("name = %q, want the build's title", entry.Name)
	}
	if !strings.HasPrefix(entry.Code, "FSB1:test-1:warrior:") {
		t.Fatalf("code = %q, want the addon code", entry.Code)
	}
}

func TestAQueuedBuildThatNoLongerExistsStillAppears(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, "/v1/addon/inbox", `{"build_id":"k7x2qm4a"}`)
	res.Body.Close()

	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "device", DeviceID: "device-1"}
	var out struct {
		Builds []InboxEntry `json:"builds"`
	}
	h.data(h.do(http.MethodGet, "/v1/addon/inbox", ""), &out)
	if len(out.Builds) != 1 || out.Builds[0].ID != "k7x2qm4a" {
		t.Fatalf("inbox = %+v", out.Builds)
	}
	if out.Builds[0].Code != "" || out.Builds[0].Name != "" {
		t.Fatalf("a build that is gone should carry no code: %+v", out.Builds[0])
	}
}

func TestQueueingTheSameBuildTwiceIsHarmless(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	for range 2 {
		res := h.do(http.MethodPost, "/v1/addon/inbox", `{"build_id":"k7x2qm4a"}`)
		res.Body.Close()
		if res.StatusCode != http.StatusCreated {
			t.Fatalf("status = %d", res.StatusCode)
		}
	}
	rows, err := h.store.Inbox(t.Context(), h.owner)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("inbox = %+v, want one row", rows)
	}
}

// TestTheInboxIsBoundedAgainstASpammyAccount pins the write-side prune
// added beyond the brief: nothing stops one signed-in session from
// queueing build ids all day, and Inbox only ever answers the newest
// InboxLimit rows anyway, so the table itself must not grow past that
// per user.
func TestTheInboxIsBoundedAgainstASpammyAccount(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	for i := 0; i < InboxLimit+5; i++ {
		// Eight lowercase base32 characters, as builds.ID produces.
		id := fmt.Sprintf("k7x2qm%c%c", 'a'+i%26, 'a'+(i/26)%26)
		res := h.do(http.MethodPost, "/v1/addon/inbox", `{"build_id":"`+id+`"}`)
		res.Body.Close()
		if res.StatusCode != http.StatusCreated {
			t.Fatalf("queueing %s = %d, want 201", id, res.StatusCode)
		}
	}
	var count int
	if err := h.pool.QueryRow(context.Background(),
		`select count(*) from addon_inbox where user_id = $1`, h.owner).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != InboxLimit {
		t.Fatalf("stored rows = %d, want the table pruned to InboxLimit (%d)", count, InboxLimit)
	}
}

// Both fields go back out of GET /v1/addon/inbox verbatim and into a
// Lua file the addon loads, so neither may be arbitrary free text.
func TestQueueingRefusesNonsense(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	for name, body := range map[string]string{
		"not json":             `{`,
		"no build":             `{"build_id":""}`,
		"build id too long":    `{"build_id":"k7x2qm4ab"}`,
		"build id not base32":  `{"build_id":"k7x2qm41"}`,
		"build id with quotes": `{"build_id":"k7x2qm"a"}`,
		"key with a quote":     `{"build_id":"k7x2qm4a","character_key":"us/normal/bael"grim"}`,
		"key with a backslash": `{"build_id":"k7x2qm4a","character_key":"us/normal/bael\grim"}`,
		"key with a newline": `{"build_id":"k7x2qm4a","character_key":"us/normal/bael
grim"}`,
		"key with a bad region": `{"build_id":"k7x2qm4a","character_key":"mars/normal/baelgrim"}`,
		"key that is not a key": `{"build_id":"k7x2qm4a","character_key":"baelgrim"}`,
		"key that is a novel": `{"build_id":"k7x2qm4a","character_key":"us/normal/` +
			strings.Repeat("a", 200) + `"}`,
	} {
		res := h.do(http.MethodPost, "/v1/addon/inbox", body)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s = %d, want 400", name, res.StatusCode)
		}
	}
}

// A build sent to the account rather than to one character carries no
// key at all, which the column's default allows.
func TestQueueingWithNoCharacterKeyIsAllowed(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, "/v1/addon/inbox", `{"build_id":"k7x2qm4a"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("status = %d, want 201", res.StatusCode)
	}
}

func TestTheRoutesNeedTheRightIdentity(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, "/v1/addon/inbox", "")
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("the inbox read from a browser = %d, want 401", res.StatusCode)
	}
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "device", DeviceID: "d"}
	res = h.do(http.MethodPost, "/v1/addon/inbox", `{"build_id":"k7x2qm4a"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("queueing from a device = %d, want 401", res.StatusCode)
	}
}

func TestEveryRouteAnswers500WhenTheDatabaseIsGone(t *testing.T) {
	h := newHarness(t)
	h.pool.Close()
	res := h.do(http.MethodPost, "/v1/addon/exports",
		`{"characters":[{"name":"Baelgrim","ruleset":"pvp","region":"us","export":"FS1:a"}]}`)
	res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("exports = %d, want 500", res.StatusCode)
	}
	res = h.do(http.MethodGet, "/v1/addon/inbox", "")
	res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("inbox = %d, want 500", res.StatusCode)
	}
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	res = h.do(http.MethodPost, "/v1/addon/inbox", `{"build_id":"k7x2qm4a"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("queue = %d, want 500", res.StatusCode)
	}
}

func TestPutExportsSyncsGuildMembershipFromTheGuildSection(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	export := "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Iron%20Vanguard:1"
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore", Export: export},
	}); err != nil {
		t.Fatal(err)
	}
	var guildID int64
	var rank string
	var rankIndex int
	if err := h.pool.QueryRow(ctx,
		`select gc.guild_id, gc.rank, gc.rank_index from guild_characters gc
		 join guilds g on g.id = gc.guild_id where g.name = 'Iron Vanguard'`).
		Scan(&guildID, &rank, &rankIndex); err != nil {
		t.Fatal(err)
	}
	if rank != "officer" || rankIndex != 1 {
		t.Fatalf("rank = %q rankIndex = %d, want officer/1 (default officer_max_rank_index is 1)", rank, rankIndex)
	}
	var memberRank string
	var verified bool
	if err := h.pool.QueryRow(ctx,
		`select rank, verified_at is not null from guild_members where guild_id = $1 and user_id = $2`, guildID, h.owner).
		Scan(&memberRank, &verified); err != nil {
		t.Fatal(err)
	}
	if memberRank != "member" || verified {
		t.Fatalf("guild_members.rank = %q verified = %v, want member/false (an unverified export never grants derived officer rank)", memberRank, verified)
	}
}

func TestPutExportsTwoCharactersOfOneAccountInOneGuildBothAppear(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Forever:0"},
		{Name: "Baelalt", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:mage:gnome:0/0/0:|guild=Forever:5"},
	}); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.pool.QueryRow(ctx, `select count(*) from guild_characters`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("guild_characters rows = %d, want 2 (one per character)", n)
	}
	var guildID int64
	h.pool.QueryRow(ctx, `select id from guilds where name = 'Forever'`).Scan(&guildID)
	var rank string
	var verified bool
	if err := h.pool.QueryRow(ctx,
		`select rank, verified_at is not null from guild_members where guild_id = $1 and user_id = $2`, guildID, h.owner).
		Scan(&rank, &verified); err != nil {
		t.Fatal(err)
	}
	if rank != "member" || verified {
		t.Fatalf("guild_members.rank = %q verified = %v, want member/false (neither character is verified; the per-character guild_characters rows above already prove both appear correctly - that's the point of this test, not derived rank)", rank, verified)
	}
}

func TestPutExportsAnUnguildedAltLeavesTheMainsRowAlone(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Forever:0"},
	}); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelalt", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:mage:gnome:0/0/0:|professions=mining"},
	}); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.pool.QueryRow(ctx,
		`select count(*) from guild_characters where character_key = 'us/hardcore/baelgrim'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("the main's row should be untouched by the alt's unguilded sync, got %d rows", n)
	}
}

func TestPutExportsATransferMovesOnlyThatCharacter(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Guild1:0"},
		{Name: "Baelalt", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:mage:gnome:0/0/0:|guild=Guild1:0"},
	}); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Guild2:0"},
	}); err != nil {
		t.Fatal(err)
	}
	var altGuild, mainGuild string
	h.pool.QueryRow(ctx, `select g.name from guild_characters gc join guilds g on g.id = gc.guild_id
		where gc.character_key = 'us/hardcore/baelalt'`).Scan(&altGuild)
	h.pool.QueryRow(ctx, `select g.name from guild_characters gc join guilds g on g.id = gc.guild_id
		where gc.character_key = 'us/hardcore/baelgrim'`).Scan(&mainGuild)
	if altGuild != "Guild1" {
		t.Fatalf("the alt's guild = %q, want Guild1 (untouched)", altGuild)
	}
	if mainGuild != "Guild2" {
		t.Fatalf("the transferred character's guild = %q, want Guild2", mainGuild)
	}
	var n int
	h.pool.QueryRow(ctx, `select count(*) from guild_members where user_id = $1`, h.owner).Scan(&n)
	if n != 2 {
		t.Fatalf("guild_members rows for the account = %d, want 2 (one per guild)", n)
	}
}

func TestPutExportsAnUnguildedExportRemovesAPreviousGuildRow(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Forever:0"},
	}); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|professions=mining"},
	}); err != nil {
		t.Fatal(err)
	}
	var n int
	h.pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/baelgrim'`).Scan(&n)
	if n != 0 {
		t.Fatal("the character's row should be gone once the export stops naming a guild")
	}
	h.pool.QueryRow(ctx, `select count(*) from guild_members where user_id = $1`, h.owner).Scan(&n)
	if n != 0 {
		t.Fatal("guild_members should have no row left for this account")
	}
}

func TestPutExportsAutoConfirmsAPendingClaimWhenTheGuildMastersExportArrives(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Officer", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Forever:1"},
	}); err != nil {
		t.Fatal(err)
	}
	var guildID int64
	h.pool.QueryRow(ctx, `select id from guilds where name = 'Forever'`).Scan(&guildID)
	if _, err := (&guilds.Store{Pool: h.pool}).Claim(ctx, guildID, h.owner, true); err != nil {
		t.Fatal(err)
	}
	var pendingBy *int64
	h.pool.QueryRow(ctx, `select claim_pending_by from guilds where id = $1`, guildID).Scan(&pendingBy)
	if pendingBy == nil {
		t.Fatal("the officer's claim should be pending before the GM's own export arrives")
	}

	gm := h.owner + 1
	if _, err := h.pool.Exec(ctx,
		`insert into users (id, email) values ($1, 'gm-auto@example.com') on conflict (id) do nothing`, gm); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, gm, []Export{
		{Name: "TheGM", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:paladin:human:0/0/0:|guild=Forever:0"},
	}); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	h.pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, guildID).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != h.owner {
		t.Fatalf("claimed_by = %v, want the originally pending officer %d, auto-confirmed by the GM's export", claimedBy, h.owner)
	}
}

func TestPutExportsResolvesGuildsCaseInsensitively(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Iron%20Vanguard:0"},
	}); err != nil {
		t.Fatal(err)
	}
	alt := h.owner + 1
	if _, err := h.pool.Exec(ctx,
		`insert into users (id, email) values ($1, 'alt-case@example.com') on conflict (id) do nothing`, alt); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, alt, []Export{
		{Name: "Caseshifter", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:mage:gnome:0/0/0:|guild=IRON%20VANGUARD:5"},
	}); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.pool.QueryRow(ctx, `select count(*) from guilds where region = 'us' and ruleset = 'hardcore'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("guilds rows = %d, want 1 (case-insensitive match, not a second decoy guild)", n)
	}
	var name string
	if err := h.pool.QueryRow(ctx, `select name from guilds where region = 'us' and ruleset = 'hardcore'`).Scan(&name); err != nil {
		t.Fatal(err)
	}
	if name != "Iron Vanguard" {
		t.Fatalf("name = %q, want the first writer's casing, Iron Vanguard", name)
	}
}

func TestPutExportsRejectsAnOutOfRangeRankIndexAsANoOpNotAnAbort(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Overflow:99999"},
	})
	if err != nil {
		t.Fatalf("an out-of-range rank index must be a silent no-op, not a batch-aborting error: %v", err)
	}
	var n int
	if err := h.pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/baelgrim'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatal("no guild_characters row should have been created for the out-of-range rank")
	}
	if err := h.pool.QueryRow(ctx, `select count(*) from guilds where name = 'Overflow'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatal("the guild itself should never have been created from invalid data")
	}
}

func TestPutExportsContinuesTheBatchPastOneBadCharacter(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "BadOne", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Overflow:99999"},
		{Name: "GoodTwo", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:mage:gnome:0/0/0:|guild=RealGuild:0"},
	}); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/goodtwo'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatal("the second, valid character in the same batch must still sync, even though the first was invalid")
	}
}

func TestConcurrentPutExportsAndApproveDoNotLoseAWrite(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "RaceMain", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=RaceGuild:5"},
	}); err != nil {
		t.Fatal(err)
	}
	var gid int64
	if err := h.pool.QueryRow(ctx, `select id from guilds where name = 'RaceGuild'`).Scan(&gid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank) values ($1, 'us/hardcore/racealt', $2, 'member')`,
		gid, h.owner); err != nil {
		t.Fatal(err)
	}
	guildStore := &guilds.Store{Pool: h.pool}

	var wg sync.WaitGroup
	var err1, err2 error
	wg.Add(2)
	go func() {
		defer wg.Done()
		err1 = h.store.PutExports(ctx, h.owner, []Export{
			{Name: "RaceMain", Region: "us", Ruleset: "hardcore",
				Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=RaceGuild:5"},
		})
	}()
	go func() {
		defer wg.Done()
		err2 = guildStore.ApproveCharacter(ctx, gid, "us/hardcore/racealt")
	}()
	wg.Wait()
	if err1 != nil {
		t.Fatalf("concurrent PutExports: %v", err1)
	}
	if err2 != nil {
		t.Fatalf("concurrent ApproveCharacter: %v", err2)
	}

	// The advisory lock in RecomputeMembership serialises the two racing
	// recomputes; the final guild_members row must reflect BOTH
	// guild_characters rows correctly, not a lost update from one racing
	// past the other's stale snapshot.
	var rank string
	var verified bool
	if err := h.pool.QueryRow(ctx,
		`select rank, verified_at is not null from guild_members where guild_id = $1 and user_id = $2`,
		gid, h.owner).Scan(&rank, &verified); err != nil {
		t.Fatal(err)
	}
	if !verified {
		t.Fatal("the approved character's verification must not have been lost to the race")
	}
}

// TestPutExportsOppositeTransfersDoNotDeadlock is E's regression test
// (2026-09-21 second security review response): two accounts, each
// transferring a character between the same two guilds in opposite
// directions at the same time, must never deadlock on the guilds'
// advisory locks - LockGuilds's fixed ascending order rules that out.
// Before the fix, each transaction locked its own "new" guild first and
// could block forever waiting for the other's "old" guild; Postgres's
// own deadlock detector would eventually abort one side with an error,
// which this test would catch as a non-nil err.
func TestPutExportsOppositeTransfersDoNotDeadlock(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()

	userB := h.owner + 1
	if _, err := h.pool.Exec(ctx,
		`insert into users (id, email) values ($1, 'deadlock-b@example.com') on conflict (id) do nothing`, userB); err != nil {
		t.Fatal(err)
	}

	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "DeadlockA", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=DeadlockOne:0"},
	}); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, userB, []Export{
		{Name: "DeadlockB", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=DeadlockTwo:0"},
	}); err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	var errA, errB error
	wg.Add(2)
	go func() {
		defer wg.Done()
		errA = h.store.PutExports(ctx, h.owner, []Export{
			{Name: "DeadlockA", Region: "us", Ruleset: "hardcore",
				Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=DeadlockTwo:0"},
		})
	}()
	go func() {
		defer wg.Done()
		errB = h.store.PutExports(ctx, userB, []Export{
			{Name: "DeadlockB", Region: "us", Ruleset: "hardcore",
				Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=DeadlockOne:0"},
		})
	}()
	wg.Wait()

	if errA != nil {
		t.Fatalf("transfer DeadlockOne->DeadlockTwo concurrent with the opposite transfer: %v", errA)
	}
	if errB != nil {
		t.Fatalf("transfer DeadlockTwo->DeadlockOne concurrent with the opposite transfer: %v", errB)
	}
}

// TestPutExportsAlsoWritesTheCharactersRow is spec §4.5: an export sync
// must make GET /v1/me list the character, not just addon_exports.
func TestPutExportsAlsoWritesTheCharactersRow(t *testing.T) {
	h := newHarness(t)
	if err := h.store.PutExports(context.Background(), h.owner,
		[]Export{{Name: "Baelgrim", Ruleset: "hardcore", Region: "us",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:"}}); err != nil {
		t.Fatal(err)
	}
	var class, source string
	var userID int64
	if err := h.pool.QueryRow(context.Background(),
		`select coalesce(class, ''), source, user_id from characters where key = 'us/hardcore/baelgrim'`).
		Scan(&class, &source, &userID); err != nil {
		t.Fatal(err)
	}
	if class != "warrior" || source != "export" || userID != h.owner {
		t.Fatalf("characters row = class=%q source=%q user_id=%d, want warrior/export/%d",
			class, source, userID, h.owner)
	}
}

// TestPutExportsNeverStealsACharactersRowEvenWithNoPriorAddonExportsRow
// covers the case addon_exports' own claim guard cannot see: a
// characters row already owned by someone else (a Battle.net import,
// say) for a key that has never been synced by device before.
func TestPutExportsNeverStealsACharactersRowEvenWithNoPriorAddonExportsRow(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	other, err := (&auth.Store{Pool: h.pool}).UpsertEmailUser(ctx, "bnet-owner@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id, source) values
		 ('us/hardcore/baelgrim', 'us', 'hardcore', 'Baelgrim', $1, 'bnet')`, other.ID); err != nil {
		t.Fatal(err)
	}
	err = h.store.PutExports(ctx, h.owner,
		[]Export{{Name: "Baelgrim", Ruleset: "hardcore", Region: "us", Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:"}})
	if !errors.Is(err, ErrCharacterClaimed) {
		t.Fatalf("err = %v, want ErrCharacterClaimed", err)
	}
}

// TestPutExportsKeepsABattleNetImportedCharactersSource: an export from
// the owner of a row Battle.net imported refreshes the row without
// demoting its source, so the nightly bnet-refresh keeps re-checking
// its guild (review finding on the 2026-09-22 branch).
func TestPutExportsKeepsABattleNetImportedCharactersSource(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if _, err := h.pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id, source) values
		 ('us/hardcore/baelgrim', 'us', 'hardcore', 'Baelgrim', $1, 'bnet')`, h.owner); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, h.owner,
		[]Export{{Name: "Baelgrim", Ruleset: "hardcore", Region: "us", Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:"}}); err != nil {
		t.Fatal(err)
	}
	var source, class string
	if err := h.pool.QueryRow(ctx,
		`select source, coalesce(class, '') from characters where key = 'us/hardcore/baelgrim'`).
		Scan(&source, &class); err != nil {
		t.Fatal(err)
	}
	if source != "bnet" || class != "warrior" {
		t.Fatalf("source = %q class = %q, want bnet kept and the class refreshed", source, class)
	}
}

// TestPostMyExportsStoresAndReturnsTheWrittenCharacters is the signed-in
// paste path (spec §4.5): a session (not a device) POSTs {"exports": [...]}
// and gets back the /v1/me character objects for what it just wrote.
func TestPostMyExportsStoresAndReturnsTheWrittenCharacters(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, "/v1/me/exports",
		`{"exports":[{"name":"Baelgrim","ruleset":"hardcore","region":"us","export":"FS1:1.60.1.69893:priest:human:0/0/0:"}]}`)
	var out struct {
		Characters []struct {
			Key    string `json:"key"`
			Class  string `json:"class"`
			Source string `json:"source"`
		} `json:"characters"`
	}
	h.data(res, &out)
	if len(out.Characters) != 1 || out.Characters[0].Key != "us/hardcore/baelgrim" ||
		out.Characters[0].Class != "priest" || out.Characters[0].Source != "export" {
		t.Fatalf("characters = %+v", out.Characters)
	}

	exports, err := h.store.Exports(context.Background(), h.owner)
	if err != nil || len(exports) != 1 {
		t.Fatalf("exports = %+v, err = %v", exports, err)
	}
}

// TestPostMyExportsNeedsASessionNotADevice covers the identity guard: a
// device token (the companion's own credential) must not reach this
// signed-in-only route.
func TestPostMyExportsNeedsASessionNotADevice(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodPost, "/v1/me/exports",
		`{"exports":[{"name":"Baelgrim","ruleset":"hardcore","region":"us","export":"FS1:aaa"}]}`)
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401 (the harness actor is a device)", res.StatusCode)
	}
}

// TestPostMyExportsRefusesNonsenseAndConflicts mirrors putExports'
// validation and claim-conflict behaviour on the session route.
func TestPostMyExportsRefusesNonsenseAndConflicts(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}

	res := h.do(http.MethodPost, "/v1/me/exports", `{"exports":[]}`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("empty exports = %d, want 400", res.StatusCode)
	}

	stranger, err := (&auth.Store{Pool: h.pool}).UpsertEmailUser(context.Background(), "stranger2@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(context.Background(), stranger.ID,
		[]Export{{Name: "Claimed", Ruleset: "hardcore", Region: "us", Export: "FS1:aaa"}}); err != nil {
		t.Fatal(err)
	}
	res = h.do(http.MethodPost, "/v1/me/exports",
		`{"exports":[{"name":"Claimed","ruleset":"hardcore","region":"us","export":"FS1:stolen"}]}`)
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("stealing a claimed character = %d, want 409", res.StatusCode)
	}
}
