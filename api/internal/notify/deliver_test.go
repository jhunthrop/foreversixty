package notify

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

const deliverURL = "https://discord.com/api/webhooks/123456789/SECRET-TOKEN"

type fakeEnts struct{ allowed bool }

func (f *fakeEnts) Can(context.Context, int64, entitlements.Feature) (bool, entitlements.Reason, error) {
	return f.allowed, entitlements.ReasonEntitled, nil
}

type deliverRig struct {
	d      *Deliverer
	poster *FakePoster
	ents   *fakeEnts
	store  *Store
	uid    int64
	logs   *bytes.Buffer
}

func newDeliverRig(t *testing.T, events map[EventName]bool) *deliverRig {
	t.Helper()
	pool := testPool(t)
	sealer, err := NewSealer(make([]byte, 32))
	if err != nil {
		t.Fatal(err)
	}
	box, err := sealer.Seal(deliverURL)
	if err != nil {
		t.Fatal(err)
	}
	st := &Store{Pool: pool}
	uid := seedUser(t, pool)
	if err := st.Save(context.Background(), uid, KindDiscord, "555", box, events); err != nil {
		t.Fatal(err)
	}
	rig := &deliverRig{poster: &FakePoster{}, ents: &fakeEnts{allowed: true}, store: st, uid: uid, logs: &bytes.Buffer{}}
	rig.d = &Deliverer{
		Store: st, Sealer: sealer, Entitlements: rig.ents, Poster: rig.poster,
		PublicBaseURL: "https://foreversixty.gg", Log: slog.New(slog.NewTextHandler(rig.logs, nil)),
	}
	return rig
}

func (r *deliverRig) deliver(t *testing.T, event EventName) error {
	t.Helper()
	return r.d.Deliver(context.Background(), r.uid, event, Embed{Title: "Secret member title"})
}

func (r *deliverRig) webhook(t *testing.T) Webhook {
	t.Helper()
	w, _, err := r.store.Get(context.Background(), r.uid, KindDiscord)
	if err != nil {
		t.Fatal(err)
	}
	return w
}

func TestDeliverPostsOnceAndRecordsSuccess(t *testing.T) {
	r := newDeliverRig(t, bothEvents)
	if _, err := r.store.RecordFailure(context.Background(), r.uid, KindDiscord, "old", true); err != nil {
		t.Fatal(err)
	}
	if err := r.deliver(t, EventSimFinished); err != nil {
		t.Fatal(err)
	}
	posts := r.poster.Posts()
	if len(posts) != 1 || posts[0].URL != deliverURL || posts[0].Embed.Title != "Secret member title" {
		t.Fatalf("posts = %+v", posts)
	}
	if w := r.webhook(t); w.Failures != 0 || w.LastError != "" {
		t.Fatalf("webhook = %+v", w)
	}
}

func TestDeliverSkipsWhenOffLapsedOrDisabled(t *testing.T) {
	t.Run("event off", func(t *testing.T) {
		r := newDeliverRig(t, map[EventName]bool{EventSimFinished: false, EventUploadParsed: true})
		_ = r.deliver(t, EventSimFinished)
		assertNoPost(t, r)
	})
	t.Run("lapsed", func(t *testing.T) {
		r := newDeliverRig(t, bothEvents)
		r.ents.allowed = false
		_ = r.deliver(t, EventSimFinished)
		assertNoPost(t, r)
	})
	t.Run("disabled", func(t *testing.T) {
		r := newDeliverRig(t, bothEvents)
		r.poster.PostErr = &DiscordError{Status: 500}
		for range MaxFailures {
			_ = r.deliver(t, EventSimFinished)
		}
		r.poster.PostErr = nil
		before := len(r.poster.Posts())
		_ = r.deliver(t, EventSimFinished)
		if len(r.poster.Posts()) != before {
			t.Fatal("posted to a disabled webhook")
		}
	})
	t.Run("absent", func(t *testing.T) {
		r := newDeliverRig(t, bothEvents)
		if err := r.d.Deliver(context.Background(), r.uid+999, EventSimFinished, Embed{}); err != nil {
			t.Fatal(err)
		}
		assertNoPost(t, r)
	})
}

func assertNoPost(t *testing.T, r *deliverRig) {
	t.Helper()
	if n := len(r.poster.Posts()); n != 0 {
		t.Fatalf("%d posts, want none", n)
	}
}

func TestFifthFailureDisablesAndTheNextEventIsSkipped(t *testing.T) {
	r := newDeliverRig(t, bothEvents)
	r.poster.PostErr = &DiscordError{Status: 404, Body: "Unknown Webhook"}
	for i := 1; i <= MaxFailures; i++ {
		_ = r.deliver(t, EventUploadParsed)
		if w := r.webhook(t); w.Failures != i || w.LastError != "Unknown Webhook" {
			t.Fatalf("after %d: %+v", i, w)
		}
	}
	if w := r.webhook(t); w.DisabledAt == nil {
		t.Fatal("not disabled after the fifth failure")
	}
	posts := len(r.poster.Posts())
	_ = r.deliver(t, EventUploadParsed)
	if len(r.poster.Posts()) != posts {
		t.Fatal("the sixth event was posted")
	}
}

func TestDeliverA429DoesNotCount(t *testing.T) {
	r := newDeliverRig(t, bothEvents)
	r.poster.PostErr = &DiscordError{Status: 429, Body: "slow down"}
	for range MaxFailures + 2 {
		_ = r.deliver(t, EventSimFinished)
	}
	w := r.webhook(t)
	if w.Failures != 0 || w.DisabledAt != nil || w.LastError != "slow down" {
		t.Fatalf("webhook = %+v", w)
	}
}

func TestAFailedDeliveryLogsNoURLTokenOrTitle(t *testing.T) {
	r := newDeliverRig(t, bothEvents)
	r.poster.PostErr = &DiscordError{Status: 500, Body: "boom"}
	_ = r.deliver(t, EventSimFinished)
	out := r.logs.String()
	if !strings.Contains(out, "outcome=failed") || !strings.Contains(out, "status=500") {
		t.Fatalf("log = %q", out)
	}
	for _, secret := range []string{"SECRET-TOKEN", "123456789", "discord.com", "Secret member title"} {
		if strings.Contains(out, secret) {
			t.Errorf("log leaks %q: %s", secret, out)
		}
	}
}

// blockingPoster waits for its context, the way a hung Discord would.
type blockingPoster struct{ FakePoster }

func (b *blockingPoster) Post(ctx context.Context, _ string, _ Embed) error {
	<-ctx.Done()
	return ctx.Err()
}

func TestAHungPostTimesOutAndCountsAsAFailure(t *testing.T) {
	r := newDeliverRig(t, bothEvents)
	r.d.Poster = &blockingPoster{}
	r.d.postTimeout = 50 * time.Millisecond
	start := time.Now()
	_ = r.deliver(t, EventSimFinished)
	if time.Since(start) > 2*time.Second {
		t.Fatal("Deliver outlived its timeout")
	}
	if w := r.webhook(t); w.Failures != 1 || w.LastError != "Discord did not answer in time." {
		t.Fatalf("webhook = %+v", w)
	}
}

// cancelAfterPost cancels the job's context the moment the post resolves.
type cancelAfterPost struct{ cancel context.CancelFunc }

func (c cancelAfterPost) Post(context.Context, string, Embed) error    { c.cancel(); return nil }
func (cancelAfterPost) Lookup(context.Context, string) (string, error) { return "", nil }

func TestRecordingSurvivesACancelledJobContext(t *testing.T) {
	r := newDeliverRig(t, bothEvents)
	if _, err := r.store.RecordFailure(context.Background(), r.uid, KindDiscord, "old", true); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	r.d.Poster = cancelAfterPost{cancel: cancel}
	if err := r.d.Deliver(ctx, r.uid, EventSimFinished, Embed{}); err != nil {
		t.Fatal(err)
	}
	if w := r.webhook(t); w.Failures != 0 || w.LastError != "" {
		t.Fatalf("success was not recorded: %+v", w)
	}
}

func TestAPlainErrorCountsWithASanitisedText(t *testing.T) {
	r := newDeliverRig(t, bothEvents)
	r.poster.PostErr = errors.New("dial https://discord.com/api/webhooks/1/SECRET-TOKEN refused")
	_ = r.deliver(t, EventSimFinished)
	w := r.webhook(t)
	if w.Failures != 1 || strings.Contains(w.LastError, "SECRET-TOKEN") || strings.Contains(w.LastError, "webhooks") {
		t.Fatalf("webhook = %+v", w)
	}
}

func TestAnUnreachableDiscordStoresAPlainSentence(t *testing.T) {
	r := newDeliverRig(t, bothEvents)
	r.poster.PostErr = ErrUnreachable
	_ = r.deliver(t, EventSimFinished)
	if w := r.webhook(t); w.Failures != 1 || w.LastError != "Discord could not be reached." {
		t.Fatalf("webhook = %+v", w)
	}
}

func TestAnErroredDeliveryLogsItsCause(t *testing.T) {
	r := newDeliverRig(t, bothEvents)
	r.d.Entitlements = errEnts{}
	r.d.deliverQuietly(context.Background(), r.uid, EventSimFinished, Embed{})
	if out := r.logs.String(); !strings.Contains(out, "outcome=error") || !strings.Contains(out, "entitlement store down") {
		t.Fatalf("log = %q", out)
	}
}

type errEnts struct{}

func (errEnts) Can(context.Context, int64, entitlements.Feature) (bool, entitlements.Reason, error) {
	return false, "", errors.New("entitlement store down")
}

func TestTheDefaultPostTimeoutIsFiveSeconds(t *testing.T) {
	if got := (&Service{}).Deliverer().effectiveTimeout(); got != 5*time.Second {
		t.Fatalf("effectiveTimeout = %v", got)
	}
}
