package reports

import (
	"io"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/contract"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

func reportBody(t *testing.T, res *http.Response, want int) []byte {
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

// reportWithFights is a public report the owner logged, with one boss kill in it.
func (h *harness) reportWithFights() string {
	h.t.Helper()
	id := h.createReport(Public)
	h.reportID = id
	h.seedFight(1, 1084, 1_000, 31_000)
	return id
}

var (
	summaryOptional = []string{"data.rows[].last_upload_at", "data.rows[].heartbeat_at"}
	viewOptional    = []string{
		"data.guild", "data.flagged", "data.fights[].encounter_id", "data.fights[].difficulty",
		"data.fights[].size", "data.fights[].zone", "data.fights[].boss_health_pct", "data.owner.battletag",
		"data.trash_groups",
	}
)

func TestAppContractReportList(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	h.reportWithFights()

	res := h.do(http.MethodGet, "/v1/reports?mine=1", "", nil)
	if got := res.Header.Get("Cache-Control"); got != "private, no-cache" {
		t.Errorf("Cache-Control = %q, want private, no-cache", got)
	}
	contract.AssertFixture(t, "reports.list", reportBody(t, res, http.StatusOK),
		contract.Options{Optional: summaryOptional})
}

func TestAppContractReportListEmptyAndErrors(t *testing.T) {
	h := newHarness(t)
	h.asSession()

	contract.AssertFixture(t, "reports.list.empty",
		reportBody(t, h.do(http.MethodGet, "/v1/reports?mine=1", "", nil), http.StatusOK),
		contract.Options{Optional: summaryOptional})
	contract.AssertFixture(t, "reports.list.error-invalid",
		reportBody(t, h.do(http.MethodGet, "/v1/reports", "", nil), http.StatusBadRequest), contract.Options{})
	h.actor = auth.Actor{}
	contract.AssertFixture(t, "error-unauthorized",
		reportBody(t, h.do(http.MethodGet, "/v1/reports?mine=1", "", nil), http.StatusUnauthorized), contract.Options{})
}

func TestAppContractReportSummary(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	id := h.reportWithFights()
	h.actor = auth.Actor{}

	res := h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	if got, want := res.Header.Get("Cache-Control"), "public, max-age=60, stale-while-revalidate=600"; got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
	contract.AssertFixture(t, "reports.get", reportBody(t, res, http.StatusOK),
		contract.Options{Optional: viewOptional})
}

func TestAppContractReportSummaryNotFound(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Private)
	h.actor = auth.Actor{}

	contract.AssertFixture(t, "reports.get.error-not-found",
		reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id, "", nil), http.StatusNotFound), contract.Options{})
}

func TestAppContractReportFoldedTrash(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	id, _ := h.onyxiaNight(Public)
	h.appReportJSON(
		store.FightEntry{Index: 1, Kind: "trash", Name: "Trash", Zone: "Onyxia's Lair", DurationMS: 40_000},
		store.FightEntry{Index: 3, Kind: "encounter", Name: "Onyxia", Zone: "Onyxia's Lair", DurationMS: 192_000},
	)
	opts := contract.Options{Optional: []string{
		"data.guild", "data.flagged", "data.fights[].encounter_id", "data.fights[].difficulty",
		"data.fights[].size", "data.fights[].zone", "data.fights[].boss_health_pct", "data.owner.battletag", "data.fights[].group",
	}}
	contract.AssertFixture(t, "reports.get",
		reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"?fold=trash", "", nil), http.StatusOK), opts)
}

func TestAppContractReportMe(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	id, _ := h.onyxiaNight(Public)
	res := h.do(http.MethodGet, "/v1/reports/"+id+"/me", "", nil)
	if got := res.Header.Get("Cache-Control"); got != "private, no-cache" {
		t.Errorf("Cache-Control = %q, want private, no-cache", got)
	}
	contract.AssertFixture(t, "reports.me", reportBody(t, res, http.StatusOK), contract.Options{})

	if _, err := h.store.Pool.Exec(t.Context(), `update characters set user_id = null`); err != nil {
		t.Fatal(err)
	}
	contract.AssertFixture(t, "reports.me.empty",
		reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/me", "", nil), http.StatusOK), contract.Options{})
	contract.AssertFixture(t, "reports.me.error-not-found",
		reportBody(t, h.do(http.MethodGet, "/v1/reports/nosuchreport/me", "", nil), http.StatusNotFound), contract.Options{})
}

func TestAppContractReportMeters(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	id, _ := h.onyxiaNight(Public)
	contract.AssertFixture(t, "reports.meters",
		reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/3/meters", "", nil), http.StatusOK), contract.Options{})
	contract.AssertFixture(t, "reports.meters.error-not-found",
		reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/9/meters", "", nil), http.StatusNotFound), contract.Options{})
}
