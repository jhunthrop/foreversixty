package sims

import (
	"context"
	"errors"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

const hookURL = "https://discord.com/api/webhooks/1/token"

type allowAll struct{}

func (allowAll) Can(context.Context, int64, entitlements.Feature) (bool, entitlements.Reason, error) {
	return true, entitlements.ReasonEntitled, nil
}

// withWebhook saves a webhook for the harness owner and returns the
// deliverer and the poster that records what it sends.
func (h *harness) withWebhook(t *testing.T) (*notify.Deliverer, *notify.FakePoster) {
	t.Helper()
	sealer, err := notify.NewSealer(make([]byte, 32))
	if err != nil {
		t.Fatal(err)
	}
	box, err := sealer.Seal(hookURL)
	if err != nil {
		t.Fatal(err)
	}
	st := &notify.Store{Pool: h.store.Pool}
	events := map[notify.EventName]bool{notify.EventSimFinished: true, notify.EventUploadParsed: true}
	if err := st.Save(t.Context(), h.owner, notify.KindDiscord, "1", nil, nil, box, events, notify.SourcePasted); err != nil {
		t.Fatal(err)
	}
	poster := &notify.FakePoster{}
	return &notify.Deliverer{
		Store: st, Sealer: sealer, Entitlements: allowAll{}, Poster: poster,
		PublicBaseURL: "https://foreversixty.gg",
	}, poster
}

func TestAFinishedRunPostsOneEmbed(t *testing.T) {
	h := newHarness(t)
	deliverer, poster := h.withWebhook(t)
	h.queued(t, "nnnnnnnnnnn1", "mage-frost")
	deps := h.jobDeps(&runner.Fixture{})
	deps.Notify = deliverer

	if err := Run(t.Context(), deps, "nnnnnnnnnnn1"); err != nil {
		t.Fatal(err)
	}
	posts := poster.Posts()
	if len(posts) != 1 || posts[0].URL != hookURL {
		t.Fatalf("posts = %+v", posts)
	}
	e := posts[0].Embed
	if e.Title != "Baelgrim's Frost Mage sim finished" || e.URL != "https://foreversixty.gg/sim/nnnnnnnnnnn1" ||
		e.Timestamp.IsZero() || len(e.Fields) < 2 || e.Fields[0].Name != "Result" {
		t.Fatalf("embed = %+v", e)
	}
}

func TestAFailedRunPostsNothing(t *testing.T) {
	h := newHarness(t)
	deliverer, poster := h.withWebhook(t)
	h.queued(t, "nnnnnnnnnnn2", "mage-frost")
	deps := h.jobDeps(&runner.Fixture{Err: errors.New("engine exploded")})
	deps.Notify = deliverer
	if err := Run(t.Context(), deps, "nnnnnnnnnnn2"); err == nil {
		t.Fatal("want the failure")
	}
	h.queued(t, "nnnnnnnnnnn3", "mage-frost")
	deps.Engine = &runner.Fixture{Aborted: true}
	_ = Run(t.Context(), deps, "nnnnnnnnnnn3")
	if n := len(poster.Posts()); n != 0 {
		t.Fatalf("%d posts for failed runs", n)
	}
}

func TestAFinishedRunNamesItsCharacterInTheTitle(t *testing.T) {
	h := newHarness(t)
	deliverer, poster := h.withWebhook(t)
	req := browserResult("mage-frost", 0).Request
	req.Iterations = preciseIterations
	req.Character.Name = "Frostbolt"
	if err := h.store.Queue(t.Context(), "nnnnnnnnnnn4", h.owner, req); err != nil {
		t.Fatal(err)
	}
	deps := h.jobDeps(&runner.Fixture{})
	deps.Notify = deliverer
	if err := Run(t.Context(), deps, "nnnnnnnnnnn4"); err != nil {
		t.Fatal(err)
	}
	posts := poster.Posts()
	if len(posts) != 1 || posts[0].Embed.Title != "Frostbolt's Frost Mage sim finished" {
		t.Fatalf("posts = %+v", posts)
	}
}
