// api/internal/addon/messages_test.go
package addon

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// quietLog is a Logger this file's Service fixtures can pass without a
// nil check on every call to s.logger(); tests that intend to hit a
// warning log path assert on the returned messages, not on log output.
func quietLog() *slog.Logger { return slog.New(slog.NewTextHandler(io.Discard, nil)) }

// ---------------------------------------------------------- pure logic --

func TestWeightCapsNamesOnlyAHardCappedRatingStat(t *testing.T) {
	caps := weightCaps([]simapi.StatWeight{
		{Stat: "strength", Weight: 1.5, Error: 0.1},
		// Hard-capped: the sweep skipped it entirely.
		{Stat: "hit", Weight: 0, Error: 0, Insignificant: true},
		{Stat: "expertise", Weight: 0, Error: 0, Insignificant: true},
		// Insignificant for an unrelated reason (its own error swamps a
		// genuine but tiny weight) -- not a cap, and not even a stat
		// the engine can hard-cap.
		{Stat: "spell_hit", Weight: 0.02, Error: 0.05, Insignificant: true},
		// A rating stat that CAN cap, but is not currently -- has a
		// real, non-zero weight.
		{Stat: "expertise", Weight: 0.4, Error: 0.05},
	})
	if got := caps; len(got) != 2 || got[0] != "expertise" || got[1] != "hit" {
		t.Fatalf("caps = %v", got)
	}
}

func TestWeightCapsIsEmptyForNoWeights(t *testing.T) {
	if caps := weightCaps(nil); len(caps) != 0 {
		t.Fatalf("caps = %v", caps)
	}
}

func TestResolveCharacterMatchesByNameSlug(t *testing.T) {
	byNameSlug := map[string]string{"bow-jackzon": "us/pvp/bow-jackzon"}
	if got := resolveCharacter("Bow Jackzon", byNameSlug); got != "us/pvp/bow-jackzon" {
		t.Fatalf("resolveCharacter = %q", got)
	}
	if got := resolveCharacter("Someone Else", byNameSlug); got != "" {
		t.Fatalf("resolveCharacter for an unknown name = %q, want unaddressed", got)
	}
	if got := resolveCharacter("", byNameSlug); got != "" {
		t.Fatalf("resolveCharacter for no name = %q", got)
	}
}

func TestFirstDoneSkipsAnUnfinishedRun(t *testing.T) {
	rows := []sims.Row{
		{SimID: "queued-1", State: sims.StateQueued},
		{SimID: "error-1", State: sims.StateError},
		{SimID: "done-1", State: sims.StateDone},
		{SimID: "done-2", State: sims.StateDone},
	}
	row := firstDone(rows)
	if row == nil || row.SimID != "done-1" {
		t.Fatalf("firstDone = %+v", row)
	}
	if firstDone(nil) != nil {
		t.Fatal("firstDone of nothing was not nil")
	}
}

// ------------------------------------------------------------- fakes --

// fakeSims answers Mine/Get from fixed, per-kind canned results, the
// way a real sims.Store would for one user with one saved run of each
// kind.
type fakeSims struct {
	byKind map[string]simapi.SimResult
}

func (f *fakeSims) Mine(_ context.Context, _ int64, _ int, kind string) (sims.Page, error) {
	if _, ok := f.byKind[kind]; !ok {
		return sims.Page{}, nil
	}
	return sims.Page{Rows: []sims.Row{{SimID: kind + "-1", Kind: kind, State: sims.StateDone}}}, nil
}

func (f *fakeSims) Get(_ context.Context, id string) (simapi.SimResult, string, error) {
	for kind, res := range f.byKind {
		if id == kind+"-1" {
			return res, "", nil
		}
	}
	return simapi.SimResult{}, "", sims.ErrNotFound
}

// fakeGuilds answers CharacterGuildState from a fixed map, keyed by
// character key.
type fakeGuilds struct {
	byKey map[string]guilds.CharacterGuildState
}

func (f *fakeGuilds) CharacterGuildState(_ context.Context, key string) (guilds.CharacterGuildState, bool, error) {
	state, ok := f.byKey[key]
	return state, ok, nil
}

// ------------------------------------------------------- Service.messages --

func TestUpgradeMessagePicksTheFirstSingleItemCombo(t *testing.T) {
	s := &Service{Log: quietLog(), Sims: &fakeSims{byKind: map[string]simapi.SimResult{
		simapi.KindGear: {
			Request: simapi.SimRequest{Character: simapi.CharacterSpec{Name: "Bow Jackzon"}},
			Combos: []simapi.Combo{
				{
					// Regears two slots at once: no single-slot delta to
					// report, so this combo is skipped even though it
					// ranks first.
					Substitutions: []simapi.Substitution{
						{Kind: simapi.SubstitutionItem, Slot: "chest", ItemID: 1},
						{Kind: simapi.SubstitutionItem, Slot: "legs", ItemID: 2},
					},
					Delta: simapi.Estimate{Mean: 80},
				},
				{
					Substitutions: []simapi.Substitution{
						{Kind: simapi.SubstitutionItem, Slot: "waist", ItemID: 3, Name: "Girdle", SourceName: "Onyxia"},
					},
					Delta: simapi.Estimate{Mean: 41.5},
				},
			},
		},
	}}}
	got := s.upgradeMessage(context.Background(), 1, map[string]string{"bow-jackzon": "us/pvp/bow-jackzon"})
	if got == nil {
		t.Fatal("expected an upgrade message")
	}
	if got.Type != messageUpgrade || got.Character != "us/pvp/bow-jackzon" || got.Slot != "waist" ||
		got.ItemID != 3 || got.ItemName != "Girdle" || got.Source != "Onyxia" || got.Delta != 41.5 {
		t.Fatalf("message = %+v", got)
	}
}

func TestUpgradeMessageIsNilWithNoFinishedRun(t *testing.T) {
	s := &Service{Log: quietLog(), Sims: &fakeSims{byKind: map[string]simapi.SimResult{}}}
	if got := s.upgradeMessage(context.Background(), 1, nil); got != nil {
		t.Fatalf("message = %+v, want nil", got)
	}
}

func TestWeightsMessageCarriesCapsAndSpec(t *testing.T) {
	s := &Service{Log: quietLog(), Sims: &fakeSims{byKind: map[string]simapi.SimResult{
		simapi.KindWeights: {
			Request: simapi.SimRequest{
				Spec:      "fury",
				Character: simapi.CharacterSpec{Name: "Bow Jackzon"},
			},
			Weights: []simapi.StatWeight{
				{Stat: "strength", Weight: 1},
				{Stat: "hit", Weight: 0, Error: 0, Insignificant: true},
			},
		},
	}}}
	got := s.weightsMessage(context.Background(), 1, map[string]string{"bow-jackzon": "us/pvp/bow-jackzon"})
	if got == nil {
		t.Fatal("expected a weights message")
	}
	if got.Type != messageWeights || got.Character != "us/pvp/bow-jackzon" || got.Spec != "fury" {
		t.Fatalf("message = %+v", got)
	}
	if len(got.Weights) != 2 || got.Weights[0].Stat != "strength" || got.Weights[0].Weight != 1 {
		t.Fatalf("weights = %+v", got.Weights)
	}
	if len(got.Caps) != 1 || got.Caps[0] != "hit" {
		t.Fatalf("caps = %+v", got.Caps)
	}
}

func TestGuildMessagesOneParticipantOneNonMember(t *testing.T) {
	s := &Service{Log: quietLog(), Guilds: &fakeGuilds{byKey: map[string]guilds.CharacterGuildState{
		"us/pvp/bow-jackzon": {
			Name: "Sanguine", Rank: "officer", PendingApprovals: 2,
			Claim: guilds.ClaimStateView{State: "claimed"},
		},
	}}}
	out := s.guildMessages(context.Background(), map[string]string{
		"bow-jackzon": "us/pvp/bow-jackzon",
		"altless":     "us/pvp/altless", // no guild_characters row: not in fakeGuilds' map
	})
	if len(out) != 1 {
		t.Fatalf("messages = %+v", out)
	}
	got := out[0]
	if got.Type != messageGuild || got.Character != "us/pvp/bow-jackzon" || got.GuildName != "Sanguine" ||
		got.ClaimState != "claimed" || got.PendingApprovals != 2 || got.Rank != "officer" {
		t.Fatalf("message = %+v", got)
	}
}

func TestMessagesDegradesGracefullyWithNeitherProducerWired(t *testing.T) {
	h := newHarness(t)
	if err := h.store.PutExports(context.Background(), h.owner, []Export{
		{Name: "Baelgrim", Ruleset: "hardcore", Region: "us", Export: "FS1:test-1:warrior:orc:0/0/0"},
	}); err != nil {
		t.Fatal(err)
	}
	s := &Service{Store: h.store, Log: quietLog()}
	if out := s.messages(context.Background(), h.owner); len(out) != 0 {
		t.Fatalf("messages = %+v, want none with Guilds and Sims both nil", out)
	}
}

func TestMessagesCombinesGuildUpgradeAndWeights(t *testing.T) {
	h := newHarness(t)
	if err := h.store.PutExports(context.Background(), h.owner, []Export{
		{Name: "Baelgrim", Ruleset: "hardcore", Region: "us", Export: "FS1:test-1:warrior:orc:0/0/0"},
	}); err != nil {
		t.Fatal(err)
	}
	key := "us/hardcore/baelgrim"
	s := &Service{
		Store: h.store, Log: quietLog(),
		Guilds: &fakeGuilds{byKey: map[string]guilds.CharacterGuildState{
			key: {Name: "Sanguine", Rank: "member", Claim: guilds.ClaimStateView{State: "unclaimed"}},
		}},
		Sims: &fakeSims{byKind: map[string]simapi.SimResult{
			simapi.KindGear: {
				Request: simapi.SimRequest{Character: simapi.CharacterSpec{Name: "Baelgrim"}},
				Combos: []simapi.Combo{{
					Substitutions: []simapi.Substitution{{Kind: simapi.SubstitutionItem, Slot: "chest", ItemID: 9}},
					Delta:         simapi.Estimate{Mean: 10},
				}},
			},
			simapi.KindWeights: {
				Request: simapi.SimRequest{Spec: "arms", Character: simapi.CharacterSpec{Name: "Baelgrim"}},
				Weights: []simapi.StatWeight{{Stat: "strength", Weight: 1}},
			},
		}},
	}
	out := s.messages(context.Background(), h.owner)
	kinds := map[string]bool{}
	for _, m := range out {
		kinds[m.Type] = true
		if m.Character != key {
			t.Errorf("message %+v not addressed to the account's own character", m)
		}
	}
	if len(out) != 3 || !kinds[messageGuild] || !kinds[messageUpgrade] || !kinds[messageWeights] {
		t.Fatalf("messages = %+v", out)
	}
}
