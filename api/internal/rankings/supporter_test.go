package rankings

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// serveWith mounts the read routes with the given Supporters.
func (h *harness) serveWith(t *testing.T, supporters Supporters) {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, &Service{Store: h.store, Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
		Now:        func() time.Time { return engine.FixtureBase },
		Supporters: supporters})
	h.server = httptest.NewServer(mux)
	t.Cleanup(h.server.Close)
}

func (h *harness) seedOwnedCharacter(slug string, ownerID *int64) {
	h.t.Helper()
	if _, err := h.pool.Exec(context.Background(),
		`insert into characters (key, region, ruleset, name, class, user_id)
		 values ($1, 'us', 'hardcore', $2, 'Warrior', $3)`, "us/hardcore/"+slug, slug, ownerID); err != nil {
		h.t.Fatal(err)
	}
}

func (h *harness) supporterOnPage(slug string) bool {
	h.t.Helper()
	res := h.get("/v1/characters/us/hardcore/" + slug)
	if res.StatusCode != http.StatusOK {
		h.t.Fatalf("character %s = %d, want 200", slug, res.StatusCode)
	}
	var page struct {
		Supporter *bool `json:"supporter"`
	}
	h.data(res, &page)
	if page.Supporter == nil {
		h.t.Fatalf("character %s: supporter missing from the body", slug)
	}
	return *page.Supporter
}

func TestCharacterPageCarriesTheSupporterMark(t *testing.T) {
	h := newHarness(t)
	h.serveWith(t, &entitlements.Store{Pool: h.pool})
	h.seedOwnedCharacter("owned", &h.owner)
	h.seedOwnedCharacter("ownerless", nil)
	if _, err := h.pool.Exec(context.Background(),
		`insert into entitlements (user_id, plan, source, status, cancel_at_period_end)
		 values ($1, 'premium', 'stripe', 'active', true)`, h.owner); err != nil {
		t.Fatal(err)
	}
	if !h.supporterOnPage("owned") {
		t.Error("an active premium owner (cancel_at_period_end) must show the mark")
	}
	if _, err := h.pool.Exec(context.Background(),
		`update entitlements set status = 'canceled' where user_id = $1`, h.owner); err != nil {
		t.Fatal(err)
	}
	if h.supporterOnPage("owned") {
		t.Error("a canceled entitlement must not show the mark")
	}
	if h.supporterOnPage("ownerless") {
		t.Error("a character with no owner must not show the mark")
	}
}

func TestCharacterPageSupporterIsFalseWithoutAService(t *testing.T) {
	h := newHarness(t)
	h.serveWith(t, nil)
	h.seedOwnedCharacter("owned", &h.owner)
	if h.supporterOnPage("owned") {
		t.Error("a nil Supporters must read as false")
	}
}

type failingSupporters struct{}

func (failingSupporters) IsSupporter(context.Context, int64) (bool, error) {
	return false, errors.New("entitlements down")
}

func TestCharacterPageSurvivesASupporterError(t *testing.T) {
	h := newHarness(t)
	h.serveWith(t, failingSupporters{})
	h.seedOwnedCharacter("owned", &h.owner)
	if h.supporterOnPage("owned") {
		t.Error("an IsSupporter error must read as false")
	}
}
