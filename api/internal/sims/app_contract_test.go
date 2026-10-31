package sims

import (
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// body reads a response, failing the test on any other status than want.
func body(t *testing.T, res *http.Response, want int) []byte {
	t.Helper()
	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != want {
		t.Fatalf("%s = %d, want %d: %s", res.Request.URL.Path, res.StatusCode, want, raw)
	}
	return raw
}

func assertCache(t *testing.T, res *http.Response, want string) {
	t.Helper()
	if got := res.Header.Get("Cache-Control"); got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
}

// rowOptional are the list-row fields Go omits when empty.
var rowOptional = []string{
	"data.rows[].finished_at", "data.rows[].failure_reason", "data.rows[].estimate_seconds",
}

func TestAppContractSimList(t *testing.T) {
	h := newHarness(t)
	finishedServerRun(h, "aaaaaaaaaaa1", "warrior-fury")
	failedServerRun(h, "aaaaaaaaaaa2")
	serverRun(h, "aaaaaaaaaaa3", "warrior-fury")
	if err := h.store.Advance(t.Context(), "aaaaaaaaaaa3", Tick{IterationsDone: 100, Mean: 900}); err != nil {
		t.Fatal(err)
	}

	res := h.do(http.MethodGet, "/v1/sims?mine=1&lane=server", "", nil)
	assertCache(t, res, "private, no-cache")
	contract.AssertFixture(t, "sims.list", body(t, res, http.StatusOK), contract.Options{Optional: rowOptional})
}

func TestAppContractSimListEmpty(t *testing.T) {
	h := newHarness(t)

	contract.AssertFixture(t, "sims.list.empty",
		body(t, h.do(http.MethodGet, "/v1/sims?mine=1", "", nil), http.StatusOK), contract.Options{Optional: rowOptional})
}

func TestAppContractSimListErrors(t *testing.T) {
	h := newHarness(t)

	contract.AssertFixture(t, "sims.list.error-invalid",
		body(t, h.do(http.MethodGet, "/v1/sims?mine=1&lane=moon", "", nil), http.StatusBadRequest), contract.Options{})
	h.anonymous()
	contract.AssertFixture(t, "error-unauthorized",
		body(t, h.do(http.MethodGet, "/v1/sims?mine=1", "", nil), http.StatusUnauthorized), contract.Options{})
}

func TestAppContractSimResult(t *testing.T) {
	h := newHarness(t)
	id := saveBrowserResult(h, "warrior-fury", 1042.5, "Tuesday")
	h.anonymous()

	res := h.do(http.MethodGet, "/v1/sims/"+id, "", nil)
	assertCache(t, res, "public, max-age=60, stale-while-revalidate=600")
	contract.AssertFixture(t, "sims.result", body(t, res, http.StatusOK), contract.Options{
		Optional: []string{"data.title"},
	})
}

func TestAppContractSimResultNotFound(t *testing.T) {
	h := newHarness(t)

	contract.AssertFixture(t, "sims.result.error-not-found",
		body(t, h.do(http.MethodGet, "/v1/sims/nosuchsim000", "", nil), http.StatusNotFound), contract.Options{})
}

func TestAppContractSimProgress(t *testing.T) {
	h := newHarness(t)
	serverRun(h, "bbbbbbbbbbb1", "warrior-fury")
	if err := h.store.Advance(t.Context(), "bbbbbbbbbbb1", Tick{IterationsDone: 1200, Mean: 981.4}); err != nil {
		t.Fatal(err)
	}
	serverRun(h, "bbbbbbbbbbb2", "warrior-fury")

	running := h.do(http.MethodGet, "/v1/sims/bbbbbbbbbbb1/progress", "", nil)
	assertCache(t, running, "no-store")
	contract.AssertFixture(t, "sims.progress", body(t, running, http.StatusOK), contract.Options{})
	contract.AssertFixture(t, "sims.progress.queued",
		body(t, h.do(http.MethodGet, "/v1/sims/bbbbbbbbbbb2/progress", "", nil), http.StatusOK),
		contract.Options{Optional: []string{"data.dps"}})
	contract.AssertFixture(t, "sims.progress.error-not-found",
		body(t, h.do(http.MethodGet, "/v1/sims/nosuchsim000/progress", "", nil), http.StatusNotFound), contract.Options{})
}

func TestAppContractSimRerun(t *testing.T) {
	h := newHarness(t)
	failedServerRun(h, "ccccccccccc1")

	contract.AssertFixture(t, "sims.rerun.error-premium-required",
		body(t, h.do(http.MethodPost, "/v1/sims/ccccccccccc1/rerun", "", nil), http.StatusPaymentRequired), contract.Options{})

	h.entitlements.allowed = true
	contract.AssertFixture(t, "sims.rerun",
		body(t, h.do(http.MethodPost, "/v1/sims/ccccccccccc1/rerun", "", nil), http.StatusAccepted), contract.Options{})
	contract.AssertFixture(t, "sims.rerun.error-not-found",
		body(t, h.do(http.MethodPost, "/v1/sims/nosuchsim000/rerun", "", nil), http.StatusNotFound), contract.Options{})
}

func TestAppContractSimInput(t *testing.T) {
	h := newHarness(t)
	ensureMetricsPartition(h)
	h.service.Summaries = dirGetter{root: h.dir}
	key := "us/normal/baelgrim"
	seedFightRow(h, key, "Baelgrim", "Warrior", "warrior-fury", time.Now().UTC().Add(-time.Hour))

	h.anonymous()
	res := h.do(http.MethodGet, "/v1/characters/us/normal/baelgrim/sim-input", "", nil)
	assertCache(t, res, "public, max-age=60, stale-while-revalidate=600")
	contract.AssertFixture(t, "character.sim-input.fight", body(t, res, http.StatusOK), contract.Options{})

	seedExport(h, key, "us", "normal", "Baelgrim",
		"FS1:1.60.1.69893:warrior:orc:0/5530515/0:head=12640,main_hand=21521", time.Now().UTC())
	signedIn := h.do(http.MethodGet, "/v1/characters/us/normal/baelgrim/sim-input", "", nil)
	contract.AssertFixture(t, "character.sim-input.addon", body(t, signedIn, http.StatusOK), contract.Options{})

	contract.AssertFixture(t, "character.sim-input.error-not-found",
		body(t, h.do(http.MethodGet, "/v1/characters/us/normal/nobody/sim-input", "", nil), http.StatusNotFound),
		contract.Options{})
}
