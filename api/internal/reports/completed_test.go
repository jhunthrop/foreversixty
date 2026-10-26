package reports

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"testing"
)

// fakeCompleted records every report the hook was told of.
type fakeCompleted struct {
	mu  sync.Mutex
	ids []string
}

func (f *fakeCompleted) ReportCompleted(_ context.Context, reportID string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.ids = append(f.ids, reportID)
}

func (f *fakeCompleted) all() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string{}, f.ids...)
}

func TestCallCompletedIsNilSafe(t *testing.T) {
	callCompleted(t.Context(), nil, "anything")
}

func TestLiveCompleteLearnsTheGuildAndZoneThenCallsTheHook(t *testing.T) {
	h := newHarness(t)
	g := h.seedGuild("Alpha", "hardcore")
	h.seedGuildCharacter(g, h.owner, harnessLoggingKey, true)
	id := h.createReport(Public)

	res := h.json(http.MethodPost, "/v1/reports/"+id+"/complete",
		`{"final_offset":1024,"engine_version":"0.1.0","health":{"lines":33,"parse_errors":0}}`)
	res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", res.StatusCode)
	}
	if got := h.reportMatchedGuild(id); !sameGuild(got, g) || h.reportGuild(id) != nil {
		t.Fatalf("guild_id = %v, want %d", got, g)
	}
	var zone *string
	if err := h.store.Pool.QueryRow(t.Context(), `select log_zone from reports where id = $1`, id).Scan(&zone); err != nil {
		t.Fatal(err)
	}
	if zone == nil || *zone != "Blackrock Spire" {
		t.Fatalf("log_zone = %v, want the live report's zone", zone)
	}
	if got := h.completed.all(); len(got) != 1 || got[0] != id {
		t.Fatalf("hook calls = %v, want the report once", got)
	}
}

func TestPatchSettingTheGuildOfACompleteReportCallsTheHook(t *testing.T) {
	h := newHarness(t)
	g := h.seedGuild("Alpha", "hardcore")
	h.addGuildMember(g, h.owner, "officer")
	id := h.createReport(Public)

	res := h.json(http.MethodPatch, "/v1/reports/"+id, `{"title":"Still live"}`)
	res.Body.Close()
	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"guild_id":`+strconv.FormatInt(g, 10)+`}`)
	res.Body.Close()
	if got := h.completed.all(); len(got) != 0 {
		t.Fatalf("a live report's patch called the hook: %v", got)
	}

	if err := h.store.SetStatus(t.Context(), id, StatusComplete, "0.1.0", nil); err != nil {
		t.Fatal(err)
	}
	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"title":"Renamed"}`)
	res.Body.Close()
	if got := h.completed.all(); len(got) != 0 {
		t.Fatalf("a patch that sets no guild called the hook: %v", got)
	}
	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"guild_id":`+strconv.FormatInt(g, 10)+`}`)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("patch = %d, want 200", res.StatusCode)
	}
	if got := h.completed.all(); len(got) != 1 || got[0] != id {
		t.Fatalf("hook calls = %v, want the report once", got)
	}
}
