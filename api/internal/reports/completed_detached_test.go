package reports

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// blockingCompleted holds every hook call until release is closed.
type blockingCompleted struct {
	entered chan struct{}
	release chan struct{}
}

func (b *blockingCompleted) ReportCompleted(ctx context.Context, _ string) {
	b.entered <- struct{}{}
	select {
	case <-b.release:
	case <-ctx.Done():
	}
}

// A hook that blocks never holds back the companion's complete call: the completion work runs
// detached from the request.
func TestABlockingHookDoesNotDelayTheLiveComplete(t *testing.T) {
	h := newHarness(t)
	hook := &blockingCompleted{entered: make(chan struct{}, 1), release: make(chan struct{})}
	h.ingest.Completed = hook
	g := h.seedGuild("Alpha", "hardcore")
	h.seedGuildCharacter(g, h.owner, harnessLoggingKey, true)
	id := h.createReport(Public)

	done := make(chan int, 1)
	go func() {
		res := h.json(http.MethodPost, "/v1/reports/"+id+"/complete",
			`{"final_offset":1024,"engine_version":"0.1.0","health":{"lines":33,"parse_errors":0}}`)
		res.Body.Close()
		done <- res.StatusCode
	}()
	select {
	case <-hook.entered:
	case <-time.After(10 * time.Second):
		t.Fatal("the hook was never called")
	}
	select {
	case status := <-done:
		if status != http.StatusNoContent {
			t.Fatalf("status = %d, want 204", status)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("the complete call waited on the blocked hook")
	}
	// The zone and the guild the matcher reads were written before the 204, hook or no hook.
	var zone *string
	if err := h.store.Pool.QueryRow(t.Context(), `select log_zone from reports where id = $1`, id).Scan(&zone); err != nil {
		t.Fatal(err)
	}
	if zone == nil || !sameGuild(h.reportMatchedGuild(id), g) {
		t.Fatalf("before the hook finished: log_zone %v matched_guild_id %v, want both set", zone, h.reportMatchedGuild(id))
	}
	close(hook.release)
	h.ingest.completions.wait()
}

type panickingCompleted struct{}

func (panickingCompleted) ReportCompleted(context.Context, string) { panic("hook exploded") }

// A hook that panics is recovered and logged: the process, and the 204, survive it.
func TestAPanickingHookDoesNotCrash(t *testing.T) {
	h := newHarness(t)
	h.ingest.Completed = panickingCompleted{}
	id := h.createReport(Public)
	res := h.json(http.MethodPost, "/v1/reports/"+id+"/complete",
		`{"final_offset":1024,"engine_version":"0.1.0","health":{"lines":33,"parse_errors":0}}`)
	res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", res.StatusCode)
	}
	h.ingest.completions.wait()
	ctx, cancel := context.WithTimeout(t.Context(), time.Second)
	defer cancel()
	if err := h.ingest.DrainCompletions(ctx); err != nil {
		t.Fatalf("drain after the panic = %v", err)
	}
}

// Drain reports the hooks still running when its context ends.
func TestDrainCompletionsIsBounded(t *testing.T) {
	var c completions
	release := make(chan struct{})
	c.run(t.Context(), nil, "r", func(context.Context) { <-release })
	ctx, cancel := context.WithTimeout(t.Context(), 50*time.Millisecond)
	defer cancel()
	if err := c.drain(ctx); err == nil {
		t.Fatal("drain returned with a hook still running")
	}
	close(release)
	if err := c.drain(t.Context()); err != nil {
		t.Fatalf("drain once the hook ended = %v", err)
	}
}
