package reports

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"testing"
)

// fakeMoved records every (report, officer) the GuildMoved hook was told of.
type fakeMoved struct {
	mu    sync.Mutex
	calls []string
}

func (f *fakeMoved) ReportGuildMoved(_ context.Context, reportID string, by int64) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls = append(f.calls, reportID+"/"+strconv.FormatInt(by, 10))
}

// fakeVisibility records every report the VisibilityChanged hook was told of.
type fakeVisibility struct {
	mu    sync.Mutex
	calls []string
}

func (f *fakeVisibility) ReportVisibilityChanged(_ context.Context, reportID string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls = append(f.calls, reportID)
}

func TestOnlyAVisibilityChangeTellsTheVisibilityHook(t *testing.T) {
	h := newHarness(t)
	seen := &fakeVisibility{}
	h.service.Visibility = seen
	id := h.createReport(Public)
	for _, body := range []string{`{"title":"Renamed"}`, `{"visibility":"public"}`, `{"visibility":"private"}`} {
		res := h.json(http.MethodPatch, "/v1/reports/"+id, body)
		res.Body.Close()
	}
	h.service.completions.wait()
	if len(seen.calls) != 1 || seen.calls[0] != id {
		t.Fatalf("visibility calls = %v, want [%s]", seen.calls, id)
	}
}

func TestAGuildPatchTellsTheMovedHookBeforeCompletion(t *testing.T) {
	h := newHarness(t)
	moved := &fakeMoved{}
	h.service.Moved = moved
	g := h.seedGuild("Alpha", "hardcore")
	h.addGuildMember(g, h.owner, "officer")
	id := h.createReport(Public)

	res := h.json(http.MethodPatch, "/v1/reports/"+id, `{"title":"No guild here"}`)
	res.Body.Close()
	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"guild_id":`+strconv.FormatInt(g, 10)+`}`)
	res.Body.Close()
	h.service.completions.wait()
	want := id + "/" + strconv.FormatInt(h.owner, 10)
	if len(moved.calls) != 1 || moved.calls[0] != want {
		t.Fatalf("moved calls = %v, want [%s]", moved.calls, want)
	}
}
