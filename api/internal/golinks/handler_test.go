package golinks

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

type fakeLinks struct {
	links    map[string]Link
	clicks   map[string]int
	bots     map[string]int
	arrivals map[int64]string
	funnel   []FunnelRow
}

func newFake() *fakeLinks {
	return &fakeLinks{
		links:    map[string]Link{"reddit": {Slug: "reddit", Target: "/", Label: "Reddit"}},
		clicks:   map[string]int{},
		bots:     map[string]int{},
		arrivals: map[int64]string{},
	}
}

func (f *fakeLinks) Target(_ context.Context, slug string) (string, error) {
	l, ok := f.links[slug]
	if !ok {
		return "", ErrNotFound
	}
	return l.Target, nil
}

func (f *fakeLinks) CountClick(_ context.Context, slug string, bot bool) (bool, error) {
	if _, ok := f.links[slug]; !ok {
		return false, nil
	}
	if bot {
		f.bots[slug]++
	} else {
		f.clicks[slug]++
	}
	return true, nil
}

func (f *fakeLinks) RecordArrival(_ context.Context, userID int64, slug string) (bool, error) {
	if _, ok := f.links[slug]; !ok || f.arrivals[userID] != "" {
		return false, nil
	}
	f.arrivals[userID] = slug
	return true, nil
}

func (f *fakeLinks) Create(_ context.Context, l Link) (Link, error) {
	if _, ok := f.links[l.Slug]; ok {
		return Link{}, ErrExists
	}
	f.links[l.Slug] = l
	return l, nil
}

func (f *fakeLinks) Funnel(context.Context) ([]FunnelRow, error) { return f.funnel, nil }

func serve(t *testing.T, f *fakeLinks, method, path, body string, actor auth.Actor) *httptest.ResponseRecorder {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, &Service{Links: f})
	r := httptest.NewRequest(method, path, strings.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, r.WithContext(auth.WithActor(r.Context(), actor)))
	return rec
}

var (
	anonymous = auth.Actor{}
	member    = auth.Actor{UserID: 5, Role: "user", Method: "session"}
	owner     = auth.Actor{UserID: 78, Role: "admin", Method: "session"}
)

func TestResolveAnswersTheTargetPubliclyCached(t *testing.T) {
	f := newFake()
	rec := serve(t, f, http.MethodGet, "/v1/golinks/reddit", "", anonymous)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `"target":"/"`) {
		t.Fatalf("resolve = %d %s", rec.Code, rec.Body)
	}
	if cc := rec.Header().Get("Cache-Control"); !strings.HasPrefix(cc, "public") {
		t.Fatalf("Cache-Control = %q, want public", cc)
	}
	if rec := serve(t, f, http.MethodGet, "/v1/golinks/nope", "", anonymous); rec.Code != http.StatusNotFound {
		t.Fatalf("unknown slug = %d, want 404", rec.Code)
	}
	if rec := serve(t, f, http.MethodGet, "/v1/golinks/Bad_Slug", "", anonymous); rec.Code != http.StatusNotFound {
		t.Fatalf("malformed slug = %d, want 404", rec.Code)
	}
}

func TestClickCountsPeopleAndBotsApart(t *testing.T) {
	f := newFake()
	for _, body := range []string{``, `{}`, `{"bot":false}`, `{"bot":true}`} {
		if rec := serve(t, f, http.MethodPost, "/v1/golinks/reddit/clicks", body, anonymous); rec.Code != http.StatusNoContent {
			t.Fatalf("click %q = %d %s", body, rec.Code, rec.Body)
		}
	}
	if f.clicks["reddit"] != 3 || f.bots["reddit"] != 1 {
		t.Fatalf("clicks=%d bots=%d, want 3 and 1", f.clicks["reddit"], f.bots["reddit"])
	}
	if rec := serve(t, f, http.MethodPost, "/v1/golinks/nope/clicks", "", anonymous); rec.Code != http.StatusNotFound {
		t.Fatalf("unknown slug click = %d, want 404", rec.Code)
	}
	if rec := serve(t, f, http.MethodPost, "/v1/golinks/reddit/clicks", `{"bot":`, anonymous); rec.Code != http.StatusBadRequest {
		t.Fatalf("broken body = %d, want 400", rec.Code)
	}
}

func TestArrivalNeedsASessionAndAValidSlug(t *testing.T) {
	f := newFake()
	if rec := serve(t, f, http.MethodPost, "/v1/me/arrival", `{"slug":"reddit"}`, anonymous); rec.Code != http.StatusUnauthorized {
		t.Fatalf("anonymous arrival = %d, want 401", rec.Code)
	}
	if rec := serve(t, f, http.MethodPost, "/v1/me/arrival", `{"slug":"Not A Slug"}`, member); rec.Code != http.StatusBadRequest {
		t.Fatalf("bad slug = %d, want 400", rec.Code)
	}
	rec := serve(t, f, http.MethodPost, "/v1/me/arrival", `{"slug":"reddit"}`, member)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `"recorded":true`) {
		t.Fatalf("arrival = %d %s", rec.Code, rec.Body)
	}
	rec = serve(t, f, http.MethodPost, "/v1/me/arrival", `{"slug":"reddit"}`, member)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `"recorded":false`) {
		t.Fatalf("second arrival = %d %s", rec.Code, rec.Body)
	}
	if f.arrivals[member.UserID] != "reddit" {
		t.Fatalf("arrival = %q", f.arrivals[member.UserID])
	}
}

func TestAdminRoutesAreTheOwnersAlone(t *testing.T) {
	f := newFake()
	f.funnel = []FunnelRow{{Link: Link{Slug: "reddit", Target: "/", Label: "Reddit"}, Clicks: 9, SignIns: 2}}
	for _, c := range []struct {
		actor auth.Actor
		want  int
	}{{anonymous, http.StatusUnauthorized}, {member, http.StatusForbidden}, {owner, http.StatusOK}} {
		if rec := serve(t, f, http.MethodGet, "/v1/admin/golinks", "", c.actor); rec.Code != c.want {
			t.Fatalf("list as %+v = %d, want %d", c.actor, rec.Code, c.want)
		}
	}
	rec := serve(t, f, http.MethodGet, "/v1/admin/golinks", "", owner)
	var env struct {
		Data struct {
			Links []FunnelRow `json:"links"`
		} `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil || len(env.Data.Links) != 1 || env.Data.Links[0].Clicks != 9 {
		t.Fatalf("list body = %s (%v)", rec.Body, err)
	}
	if cc := rec.Header().Get("Cache-Control"); !strings.Contains(cc, "no-store") {
		t.Fatalf("Cache-Control = %q, want no-store", cc)
	}
	if rec := serve(t, f, http.MethodPost, "/v1/admin/golinks", `{"slug":"x","target":"/","label":"X"}`, member); rec.Code != http.StatusForbidden {
		t.Fatalf("create as a member = %d, want 403", rec.Code)
	}
}

func TestCreateValidatesAndRefusesATakenSlug(t *testing.T) {
	f := newFake()
	cases := []struct {
		body  string
		want  int
		field string
	}{
		{`{"slug":"Bad Slug","target":"/","label":"X"}`, http.StatusBadRequest, "slug"},
		{`{"slug":"ok","target":"https://evil.example","label":"X"}`, http.StatusBadRequest, "target"},
		{`{"slug":"ok","target":"//evil.example","label":"X"}`, http.StatusBadRequest, "target"},
		{`{"slug":"ok","target":"/","label":""}`, http.StatusBadRequest, "label"},
		{`{"slug":"reddit","target":"/","label":"Again"}`, http.StatusConflict, "slug"},
		{`not json`, http.StatusBadRequest, ""},
	}
	for _, c := range cases {
		rec := serve(t, f, http.MethodPost, "/v1/admin/golinks", c.body, owner)
		if rec.Code != c.want || (c.field != "" && !strings.Contains(rec.Body.String(), `"`+c.field+`"`)) {
			t.Errorf("create %s = %d %s, want %d naming %q", c.body, rec.Code, rec.Body, c.want, c.field)
		}
	}
	rec := serve(t, f, http.MethodPost, "/v1/admin/golinks", `{"slug":"reddit-tbc","target":"/guides","label":" TBC post "}`, owner)
	if rec.Code != http.StatusCreated || f.links["reddit-tbc"].Label != "TBC post" {
		t.Fatalf("create = %d %s, stored %+v", rec.Code, rec.Body, f.links["reddit-tbc"])
	}
}
