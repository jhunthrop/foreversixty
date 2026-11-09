// api/internal/billing/creator_code_test.go
package billing

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// fakeCreators is the creator program as billing sees it: one live code.
type fakeCreators struct {
	live      map[string]bool
	completed []string
}

func (f *fakeCreators) CodeIsLive(_ context.Context, code string) (bool, error) { return f.live[code], nil }
func (f *fakeCreators) CodeTrialDays() int64                                     { return 30 }
func (f *fakeCreators) CheckoutCompleted(_ context.Context, _ int64, code string, _ time.Time, _ string) error {
	f.completed = append(f.completed, code)
	return nil
}
func (f *fakeCreators) InvoicePaid(context.Context, int64, string) error { return nil }
func (f *fakeCreators) ChargeRefunded(context.Context, string) error     { return nil }
func (f *fakeCreators) AccountUpdated(context.Context, string) error     { return nil }

func withCreators(h *harness) *fakeCreators {
	f := &fakeCreators{live: map[string]bool{"GRIMBREW": true}}
	h.svc.Creators = f
	return f
}

func TestCreatorCodeGivesPremiumAFreeMonthOnce(t *testing.T) {
	h := newHarness(t)
	withCreators(h)
	uid := h.seedUser(t, "code-premium@example.com")
	p := mustCheckout(t, h, uid, map[string]any{"plan": "premium", "interval": "monthly", "creator_code": " grimbrew "})
	if p.TrialPeriodDays != 30 || p.Metadata["creator_code"] != "GRIMBREW" {
		t.Fatalf("trial = %d, code = %q; want 30 and GRIMBREW", p.TrialPeriodDays, p.Metadata["creator_code"])
	}
	plain := mustCheckout(t, h, h.seedUser(t, "plain@example.com"), map[string]any{"plan": "premium", "interval": "monthly"})
	if plain.TrialPeriodDays != 0 || plain.Metadata["creator_code"] != "" {
		t.Fatalf("a checkout without a code got trial %d code %q", plain.TrialPeriodDays, plain.Metadata["creator_code"])
	}
	if err := h.svc.Store.RecordPremiumTrial(t.Context(), uid, "sub_first"); err != nil {
		t.Fatal(err)
	}
	if code, params := h.checkout(t, uid, map[string]any{"plan": "premium", "interval": "monthly", "creator_code": "GRIMBREW"}); code != http.StatusConflict || params != nil {
		t.Fatalf("a second trial = %d, want 409 and no session", code)
	}
}

func TestCreatorCodeRefusedWhenNotACreators(t *testing.T) {
	h := newHarness(t)
	withCreators(h)
	uid := h.seedUser(t, "badcode@example.com")
	if code, params := h.checkout(t, uid, map[string]any{"plan": "premium", "interval": "yearly", "creator_code": "NOPE"}); code != http.StatusBadRequest || params != nil {
		t.Fatalf("an unknown code = %d, want 400 and no session", code)
	}
	h.svc.Creators = nil
	if code, _ := h.checkout(t, uid, map[string]any{"plan": "premium", "interval": "yearly", "creator_code": "GRIMBREW"}); code != http.StatusBadRequest {
		t.Fatalf("a code with the program off = %d, want 400", code)
	}
}

func TestCreatorCodeTurnsTheGuildTrialIntoAMonth(t *testing.T) {
	h := newHarness(t)
	withCreators(h)
	gid, officer := h.officerOf(t, "Code Guild")
	p := mustCheckout(t, h, officer, map[string]any{"plan": "guild", "interval": "monthly", "guild_id": gid, "creator_code": "GRIMBREW"})
	if p.TrialPeriodDays != 30 || p.Metadata["creator_code"] != "GRIMBREW" {
		t.Fatalf("guild trial = %d (standard %d), code %q", p.TrialPeriodDays, catalog.GuildTrialDays, p.Metadata["creator_code"])
	}
}

func TestCreatorCodeRefusedOnAGuildsSecondTrial(t *testing.T) {
	h := newHarness(t)
	withCreators(h)
	gid, officer := h.officerOf(t, "Second Trial Guild")
	h.seedGuildTrial(t, gid)
	if code, params := h.checkout(t, officer, map[string]any{"plan": "raid_guild", "interval": "yearly", "guild_id": gid, "creator_code": "GRIMBREW"}); code != http.StatusConflict || params != nil {
		t.Fatalf("a code on a used trial = %d, want 409 and no session", code)
	}
	// Without a code the guild still checks out, with no trial, as before.
	p := mustCheckout(t, h, officer, map[string]any{"plan": "raid_guild", "interval": "yearly", "guild_id": gid})
	if p.TrialPeriodDays != 0 {
		t.Fatalf("a used trial gave %d days", p.TrialPeriodDays)
	}
}
