package rankings

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
)

// These three routes are the spec's "Live public boards" class (§2.2):
// a thirty-second edge cache with a five-minute stale-while-revalidate
// window, set by the package's shared cache helper. Each test hits the
// real handler through the package's own httptest server rather than
// calling the handler function directly, matching the pattern the rest
// of this package's route tests already use (see serve/h.get in
// query_test.go).
func TestRankingsSetsLiveBoardCacheControl(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	res := h.get("/v1/rankings?encounter=9001")
	res.Body.Close()
	want := "public, max-age=30, stale-while-revalidate=300"
	if got := res.Header.Get("Cache-Control"); got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
}

func TestCharacterSetsLiveBoardCacheControl(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	res := h.get("/v1/characters/us/hardcore/baelgrim")
	res.Body.Close()
	want := "public, max-age=30, stale-while-revalidate=300"
	if got := res.Header.Get("Cache-Control"); got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
}

func TestGuildSetsLiveBoardCacheControl(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	res := h.get("/v1/guilds/us/hardcore/forever-sixty")
	res.Body.Close()
	want := "public, max-age=30, stale-while-revalidate=300"
	if got := res.Header.Get("Cache-Control"); got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
}
