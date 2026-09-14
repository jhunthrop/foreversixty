package addon

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/db"
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
	Mount(mux, &Service{Store: h.store, Builds: h.builds, Data: data,
		Log: slog.New(slog.NewTextHandler(io.Discard, nil))})
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
	if _, _, err := h.builds.Save(context.Background(), b); err != nil {
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
		res := h.do(http.MethodPost, "/v1/addon/inbox",
			`{"build_id":"k7x2qm4`+string(rune('a'+i%26))+string(rune('0'+i/26))+`"}`)
		res.Body.Close()
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

func TestQueueingRefusesNonsense(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	for name, body := range map[string]string{"not json": `{`, "no build": `{"build_id":""}`} {
		res := h.do(http.MethodPost, "/v1/addon/inbox", body)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s = %d, want 400", name, res.StatusCode)
		}
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
