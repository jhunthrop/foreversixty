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
	close(hook.release)
	h.ingest.completions.wait()
}
