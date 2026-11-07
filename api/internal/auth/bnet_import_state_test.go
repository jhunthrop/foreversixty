package auth

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestBnetImportViewReadsAStaleRunningStateAsFailed(t *testing.T) {
	at := time.Date(2026, 11, 7, 12, 0, 0, 0, time.UTC)
	state := func(s string) *string { return &s }
	cases := []struct {
		name string
		user User
		now  time.Time
		want *BnetImportView
	}{
		{"never tracked", User{}, at, nil},
		{"running", User{BnetImportState: state(BnetImportRunning), BnetImportStateAt: &at},
			at.Add(bnetImportStaleAfter), &BnetImportView{State: BnetImportRunning, At: "2026-11-07T12:00:00Z"}},
		{"stale running", User{BnetImportState: state(BnetImportRunning), BnetImportStateAt: &at},
			at.Add(bnetImportStaleAfter + time.Second), &BnetImportView{State: BnetImportFailed, At: "2026-11-07T12:00:00Z"}},
		{"old done stays done", User{BnetImportState: state(BnetImportDone), BnetImportStateAt: &at},
			at.Add(24 * time.Hour), &BnetImportView{State: BnetImportDone, At: "2026-11-07T12:00:00Z"}},
	}
	for _, c := range cases {
		got := bnetImportViewOf(c.user, c.now)
		if (got == nil) != (c.want == nil) || (got != nil && *got != *c.want) {
			t.Errorf("%s: view = %+v, want %+v", c.name, got, c.want)
		}
	}
	if bnetImportStaleAfter <= importBudget+importContinueBudget {
		t.Fatalf("stale window %v must outlast the longest import", bnetImportStaleAfter)
	}
}

// meBnetImport reads /v1/me's bnet_import for the signed-in browser.
func (h *harness) meBnetImport(t *testing.T) *BnetImportView {
	t.Helper()
	var body struct {
		BnetImport *BnetImportView `json:"bnet_import"`
	}
	h.decode(t, h.do(t, http.MethodGet, "/v1/me", ""), &body)
	return body.BnetImport
}

// waitBnetImport polls /v1/me until bnet_import reads want, as /setup does.
func (h *harness) waitBnetImport(t *testing.T, want string) *BnetImportView {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for {
		got := h.meBnetImport(t)
		if got != nil && got.State == want {
			return got
		}
		if time.Now().After(deadline) {
			t.Fatalf("bnet_import = %+v, want state %q", got, want)
		}
		time.Sleep(20 * time.Millisecond)
	}
}

// signInWithBattleNet runs the web Battle.net sign-in to the callback's redirect.
func (h *harness) signInWithBattleNet(t *testing.T) {
	t.Helper()
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/start?next=/setup", "")
	res.Body.Close()
	target, err := url.Parse(res.Header.Get("Location"))
	if err != nil {
		t.Fatal(err)
	}
	res = h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state="+target.Query().Get("state"), "")
	res.Body.Close()
	if res.StatusCode != http.StatusFound {
		t.Fatalf("callback = %d, want a redirect", res.StatusCode)
	}
}

// heldImporter's first run runs out of budget with characters pending; the background
// continuation then waits for release and ends with result.
type heldImporter struct {
	release chan struct{}
	result  error
	calls   int
}

func (b *heldImporter) ImportAccount(ctx context.Context, _ int64, _ string) (ImportSummary, error) {
	b.calls++
	if b.calls == 1 {
		return ImportSummary{Characters: 3, Pending: 2}, errors.New("budget")
	}
	<-b.release
	return ImportSummary{Characters: 3}, b.result
}

type quickImporter struct{ err error }

func (q quickImporter) ImportAccount(context.Context, int64, string) (ImportSummary, error) {
	return ImportSummary{Characters: 1}, q.err
}

func TestBnetImportIsAbsentUntilAnImportIsTracked(t *testing.T) {
	h := newHarness(t)
	h.withBattleNet(t, "22001", "Quiet#1234")
	h.signInWithBattleNet(t) // no Importer: nothing is tracked
	if got := h.meBnetImport(t); got != nil {
		t.Fatalf("bnet_import = %+v, want it omitted", got)
	}
}

func TestBnetImportReadsDoneOrFailedWhenTheSignInRunEnds(t *testing.T) {
	for _, c := range []struct {
		err  error
		want string
	}{{nil, BnetImportDone}, {errors.New("blizzard answered 503"), BnetImportFailed}} {
		h := newHarness(t)
		h.withBattleNet(t, "22002", "Swift#1234")
		h.svc.Importer = quickImporter{c.err}
		h.signInWithBattleNet(t)
		got := h.meBnetImport(t)
		if got == nil || got.State != c.want {
			t.Fatalf("err %v: bnet_import = %+v, want %q", c.err, got, c.want)
		}
		if _, err := time.Parse(time.RFC3339, got.At); err != nil {
			t.Fatalf("at = %q, want RFC 3339: %v", got.At, err)
		}
	}
}

// The sign-in hands the import to the background: /v1/me reads running until the
// continuation ends, then done or failed by its result.
func TestBnetImportRunsUntilTheBackgroundContinuationEnds(t *testing.T) {
	for _, c := range []struct {
		err  error
		want string
	}{{nil, BnetImportDone}, {errors.New("budget"), BnetImportFailed}} {
		h := newHarness(t)
		h.withBattleNet(t, "22003", "Slow#1234")
		importer := &heldImporter{release: make(chan struct{}), result: c.err}
		h.svc.Importer = importer
		h.signInWithBattleNet(t)
		if got := h.meBnetImport(t); got == nil || got.State != BnetImportRunning {
			close(importer.release)
			t.Fatalf("bnet_import during the continuation = %+v, want running", got)
		}
		close(importer.release)
		h.waitBnetImport(t, c.want)
	}
}

// A running state left behind by an instance that stopped mid-import reads as failed once
// it is older than any import can run.
func TestBnetImportRunningLeftBehindReadsAsFailed(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "stale-import@example.com")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`update users set bnet_import_state = 'running', bnet_import_state_at = now() - interval '1 hour' where id = $1`,
		uid); err != nil {
		t.Fatal(err)
	}
	var body struct {
		BnetImport *BnetImportView `json:"bnet_import"`
	}
	h.decode(t, h.sessionRequest(t, http.MethodGet, "/v1/me", uid), &body)
	if body.BnetImport == nil || body.BnetImport.State != BnetImportFailed {
		t.Fatalf("bnet_import = %+v, want failed", body.BnetImport)
	}
}
