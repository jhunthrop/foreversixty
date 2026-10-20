package addon_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
)

func TestTheSyncRemembersHowManyCharactersWentUpAndWhen(t *testing.T) {
	r := newRig(t, noOutbox)
	if st := r.sync.Status(); st.SentCount != 0 || !st.SentAt.IsZero() {
		t.Fatalf("a sync that has sent nothing says %+v", st)
	}
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	want := r.at
	st := r.sync.Status()
	if st.SentCount != 1 || !st.SentAt.Equal(want) || st.Refusal != "" {
		t.Fatalf("status after the upload = %+v, want 1 character at %v", st, want)
	}
	// A restart reads it back from the state directory.
	r.restart()
	if st := r.sync.Status(); st.SentCount != 1 || !st.SentAt.Equal(want) {
		t.Errorf("status after a restart = %+v", st)
	}
}

func TestARefusedListKeepsTheSitesReasonUntilALaterListIsAccepted(t *testing.T) {
	r := newRig(t, noOutbox)
	r.srv.RefuseCharacters(true)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	st := r.sync.Status()
	if st.Refusal != "that export cannot be stored" || !st.RefusedAt.Equal(r.at) || st.SentCount != 0 {
		t.Fatalf("status after the refusal = %+v", st)
	}
	r.restart()
	if got := r.sync.Status().Refusal; got != "that export cannot be stored" {
		t.Errorf("the reason did not survive a restart: %q", got)
	}
	r.srv.RefuseCharacters(false)
	r.save(noOutbox)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	if st := r.sync.Status(); st.Refusal != "" || st.SentCount != 1 {
		t.Errorf("status after a good list = %+v", st)
	}
}

// tooBusy is an API whose character posts are answered 429 until it is let through.
type tooBusy struct {
	addon.API
	calls int
	pass  bool
}

func (b *tooBusy) PostAddonExports(ctx context.Context, chars []addon.Export, ob *addon.Outbox) (addon.ExportsResult, error) {
	b.calls++
	if !b.pass {
		return addon.ExportsResult{}, &client.Error{Status: http.StatusTooManyRequests,
			Message: "slow down", RetryAfter: 30 * time.Second}
	}
	return b.API.PostAddonExports(ctx, chars, ob)
}

// A 429 is the per-device budget talking: the list is sent again on the
// next pass, never taken for a refusal or for a dead token.
func TestA429OnTheCharacterListIsRetriedAndNeverARefusal(t *testing.T) {
	r := newRig(t, noOutbox)
	busy := &tooBusy{API: r.api}
	r.sync = addon.New(addon.Options{
		Paths: func() []string { return []string{r.path} }, API: busy, StateDir: r.state,
		Now: func() time.Time { return r.at },
	})
	err := r.poll()
	if err == nil || !client.Retryable(err) || client.Unauthorized(err) {
		t.Fatalf("a 429 pass = %v; want a retryable error that is not a dead token", err)
	}
	if st := r.sync.Status(); st.Refusal != "" || st.SentCount != 0 {
		t.Errorf("a 429 was recorded as %+v", st)
	}
	busy.pass = true
	if err := r.poll(); err != nil {
		t.Fatalf("the retry failed: %v", err)
	}
	if busy.calls != 2 || r.sync.Status().SentCount != 1 {
		t.Errorf("calls = %d, status = %+v", busy.calls, r.sync.Status())
	}
}
