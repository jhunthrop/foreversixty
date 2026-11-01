package rankings

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/engine"
)

// serveAs mounts the routes with every request carrying actor, as the session middleware would.
func serveAs(t *testing.T, h *harness, actor auth.Actor) {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, &Service{Store: h.store, Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
		Now: func() time.Time { return engine.FixtureBase }, ScoreGear: panelReader{}})
	h.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if actor.Signed() {
			r = r.WithContext(auth.WithActor(r.Context(), actor))
		}
		mux.ServeHTTP(w, r)
	}))
	t.Cleanup(h.server.Close)
}

func appActor(userID int64) auth.Actor {
	return auth.Actor{UserID: userID, Role: "user", Method: "device", DeviceID: "ph", DeviceKind: auth.DeviceKindApp}
}

func (h *harness) send(method, path string) (*http.Response, []byte) {
	h.t.Helper()
	req, err := http.NewRequest(method, h.server.URL+path, nil)
	if err != nil {
		h.t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		h.t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		h.t.Fatal(err)
	}
	return res, body
}

func (h *harness) watchedKeys() []string {
	h.t.Helper()
	res, body := h.send(http.MethodGet, "/v1/me/watching")
	if res.StatusCode != http.StatusOK {
		h.t.Fatalf("GET watching = %d: %s", res.StatusCode, body)
	}
	var env struct {
		Data struct {
			Rows []WatchedCharacter `json:"rows"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &env); err != nil {
		h.t.Fatal(err)
	}
	keys := []string{}
	for _, r := range env.Data.Rows {
		keys = append(keys, r.Key)
	}
	return keys
}

// loggedKey seeds one ranked fight and answers a player key the logs know (no characters row).
func (h *harness) loggedKey() string {
	h.t.Helper()
	_, _, key := h.seedRankedFight(h.t)
	return key
}

func TestWatchingRequiresSignIn(t *testing.T) {
	h := newHarness(t)
	serveAs(t, h, auth.Actor{})
	if res, _ := h.send(http.MethodGet, "/v1/me/watching"); res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("GET = %d, want 401", res.StatusCode)
	}
	if res, _ := h.send(http.MethodPut, "/v1/me/watching/us/hardcore/baelgrim"); res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("PUT = %d, want 401", res.StatusCode)
	}
}

func TestWatchAddsListsAndRemoves(t *testing.T) {
	h := newHarness(t)
	serveAs(t, h, appActor(h.owner))
	key := h.loggedKey()

	for range 2 { // PUT is idempotent
		res, body := h.send(http.MethodPut, "/v1/me/watching/"+key)
		if res.StatusCode != http.StatusOK {
			t.Fatalf("PUT = %d: %s", res.StatusCode, body)
		}
	}
	if got := h.watchedKeys(); len(got) != 1 || got[0] != key {
		t.Fatalf("watching = %v, want [%s]", got, key)
	}
	if res, _ := h.send(http.MethodDelete, "/v1/me/watching/"+key); res.StatusCode != http.StatusNoContent {
		t.Fatalf("DELETE = %d, want 204", res.StatusCode)
	}
	if res, _ := h.send(http.MethodDelete, "/v1/me/watching/"+key); res.StatusCode != http.StatusNoContent {
		t.Fatalf("DELETE again = %d, want 204", res.StatusCode)
	}
	if got := h.watchedKeys(); len(got) != 0 {
		t.Fatalf("watching after delete = %v, want none", got)
	}
}

func TestWatchRowNamesTheCharacter(t *testing.T) {
	h := newHarness(t)
	serveAs(t, h, appActor(h.owner))
	key := h.loggedKey()
	_, body := h.send(http.MethodPut, "/v1/me/watching/"+key)
	var env struct {
		Data WatchedCharacter `json:"data"`
	}
	if err := json.Unmarshal(body, &env); err != nil {
		t.Fatal(err)
	}
	row := env.Data
	if row.Key != key || row.Region != "us" || row.Ruleset != "hardcore" || row.Name == "" || row.Class == nil {
		t.Fatalf("row = %+v, want the logged character named, with its class", row)
	}
}

func TestWatchUnknownCharacterIsNotFound(t *testing.T) {
	h := newHarness(t)
	serveAs(t, h, appActor(h.owner))
	for _, path := range []string{"us/hardcore/nobody", "xx/hardcore/nobody", "us/classic/nobody"} {
		if res, _ := h.send(http.MethodPut, "/v1/me/watching/"+path); res.StatusCode != http.StatusNotFound {
			t.Errorf("PUT %s = %d, want 404", path, res.StatusCode)
		}
	}
}

func TestWatchingLimit(t *testing.T) {
	h := newHarness(t)
	serveAs(t, h, appActor(h.owner))
	for i := range WatchLimit {
		if _, err := h.pool.Exec(t.Context(),
			`insert into watched_characters (user_id, character_key) values ($1, $2)`,
			h.owner, fmt.Sprintf("us/hardcore/filler-%02d", i)); err != nil {
			t.Fatal(err)
		}
	}
	key := h.loggedKey()
	res, body := h.send(http.MethodPut, "/v1/me/watching/"+key)
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("PUT past the limit = %d, want 409: %s", res.StatusCode, body)
	}
	var env struct {
		Error struct {
			Code string `json:"code"`
		} `json:"error"`
	}
	if err := json.Unmarshal(body, &env); err != nil || env.Error.Code != "limit" {
		t.Fatalf("error = %s, want code limit", body)
	}
}

func TestWatchedAnonymousCharacterDisappears(t *testing.T) {
	h := newHarness(t)
	serveAs(t, h, appActor(h.owner))
	h.seedOwnedCharacter("baelgrim", &h.owner)
	if res, body := h.send(http.MethodPut, "/v1/me/watching/us/hardcore/baelgrim"); res.StatusCode != http.StatusOK {
		t.Fatalf("PUT = %d: %s", res.StatusCode, body)
	}
	if _, err := h.pool.Exec(t.Context(), `update users set anonymize = true where id = $1`, h.owner); err != nil {
		t.Fatal(err)
	}
	if got := h.watchedKeys(); len(got) != 0 {
		t.Fatalf("watching = %v, want the anonymous character gone", got)
	}
	if res, _ := h.send(http.MethodPut, "/v1/me/watching/us/hardcore/baelgrim"); res.StatusCode != http.StatusNotFound {
		t.Fatalf("PUT anonymous = %d, want 404", res.StatusCode)
	}
}

func TestWatchingIsPerAccount(t *testing.T) {
	h := newHarness(t)
	other, err := (&auth.Store{Pool: h.pool}).UpsertEmailUser(t.Context(), "other@example.com")
	if err != nil {
		t.Fatal(err)
	}
	serveAs(t, h, appActor(other.ID))
	key := h.loggedKey()
	h.send(http.MethodPut, "/v1/me/watching/"+key)
	serveAs(t, h, appActor(h.owner))
	if got := h.watchedKeys(); len(got) != 0 {
		t.Fatalf("owner sees %v, want nothing another account watches", got)
	}
}

func TestRankingEncounters(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)
	h.seedFight("report-one", 2, engine.FixtureBase.Add(time.Hour), nil)
	if _, err := h.pool.Exec(t.Context(), `update reports set log_zone = 'Blackrock Depths' where id = 'report-one'`); err != nil {
		t.Fatal(err)
	}

	var page struct {
		Encounters []RankedEncounter `json:"encounters"`
	}
	h.data(h.get("/v1/rankings/encounters"), &page)
	if len(page.Encounters) != 1 {
		t.Fatalf("encounters = %+v, want the one ranked encounter", page.Encounters)
	}
	e := page.Encounters[0]
	if e.ID != executionEncounterID || e.Name != "Warden Kelthas" || e.Slug != "warden-kelthas" ||
		e.Kills != 2 || e.Size == nil || *e.Size != 5 || e.Zone == nil || *e.Zone != "Blackrock Depths" {
		t.Fatalf("encounter = %+v", e)
	}
}
