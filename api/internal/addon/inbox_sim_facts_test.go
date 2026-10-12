package addon

// Task L6.11: the sim facts the build dialog needs (docs/contracts/addon-v2.md contract 3).
// An in-game request queues a weights run; its weights message says when the run that made
// them was queued and which export it ran on (the same export_seen_at the ack carries for
// that export); the ack says when the site received the request and whether its run is done
// or failed. Every test goes through the real outbox post and the real inbox builder.

import (
	"context"
	"encoding/json"
	"slices"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/sims"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// inGameRequest posts the contract outbox (an in-game sim request for obnoxiousKey) and
// answers the request's run id.
func inGameRequest(t *testing.T, h *outboxHarness) string {
	t.Helper()
	h.post(withOutbox(string(contractFixture(t, "outbox-v1.json"))))
	reqs := h.simRequests()
	if len(reqs) != 1 || reqs[0].simID == "" {
		t.Fatalf("sim requests = %+v, want one with a run", reqs)
	}
	return reqs[0].simID
}

// finishWeights finishes run id the way the sim job does, with DPS and two stat weights.
func finishWeights(t *testing.T, store *sims.Store, id string) {
	t.Helper()
	res, _, err := store.Get(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	res.DPS = simapi.Estimate{Mean: 412.6, Error: 3.1}
	res.Weights = []simapi.StatWeight{{Stat: "attack_power", Weight: 1}, {Stat: "strength", Weight: 2.2}}
	if err := store.Finish(context.Background(), id, res); err != nil {
		t.Fatal(err)
	}
}

func rfc3339(t *testing.T, h *outboxHarness, sql string, args ...any) string {
	t.Helper()
	var at time.Time
	if err := h.pool.QueryRow(context.Background(), sql, args...).Scan(&at); err != nil {
		t.Fatal(err)
	}
	return at.UTC().Format(time.RFC3339)
}

// The in-game request queues a weights run (spec §5.2: "Sim me" answers with DPS and
// weights). Once it is done, that character's weights message comes from it, carrying the
// run's queue time and the export the request was made on, equal to the ack's; the ack says
// done and when the site received the request. A second character of the same account with
// only a website weights run keeps today's fallback: generated_at, and no export_seen_at.
func TestAnInGameRequestsWeightsCarryTheirTimeAndExport(t *testing.T) {
	h := newOutboxHarness(t, true)
	cleanL63(t, h.harness)
	h.pairDevice("l63-dev", false)
	h.exportL63(l63Export)
	simStore := &sims.Store{Pool: h.pool}
	ctx := context.Background()

	id := inGameRequest(t, h)
	var kind string
	var body []byte
	if err := h.pool.QueryRow(ctx, `select kind, result from sims where id = $1`, id).Scan(&kind, &body); err != nil {
		t.Fatal(err)
	}
	var stored simapi.SimResult
	if err := json.Unmarshal(body, &stored); err != nil {
		t.Fatal(err)
	}
	spec := specs.ByKey["warrior-arms"]
	if kind != simapi.KindWeights || stored.Request.Weights == nil ||
		!slices.Equal(stored.Request.Weights.Stats, spec.WeightStats) || stored.Request.Weights.Reference != spec.ReferenceStat {
		t.Fatalf("the in-game request queued a %q run with weights %+v, want a weights run over %v against %s",
			kind, stored.Request.Weights, spec.WeightStats, spec.ReferenceStat)
	}

	// A website weights run for the other character, older than the in-game one.
	if err := simStore.Save(ctx, "l63-web-weights", &h.owner, "", simapi.SimResult{
		Request: simapi.SimRequest{Spec: "warrior-arms", Character: simapi.CharacterSpec{Name: l63Name},
			Weights: &simapi.WeightsSpec{Stats: []string{"attack_power"}, Reference: "attack_power"}},
		Weights: []simapi.StatWeight{{Stat: "attack_power", Weight: 1}, {Stat: "strength", Weight: 1.9}},
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx, `update sims set created_at = now() - interval '1 hour' where id = 'l63-web-weights'`); err != nil {
		t.Fatal(err)
	}
	s := &Service{Store: h.store, Sims: simStore, Log: quietLog()}

	// Before the run finishes: the ack says queued, and the in-game character has no weights.
	out := s.messages(ctx, h.owner, time.Now())
	if acks := messagesOf(out, messageAck); len(acks) != 2 {
		t.Fatalf("acks = %+v, want one per export", acks)
	}
	if w := weightsFor(out, obnoxiousKey); w != nil {
		t.Fatalf("weights for %s before its run finished = %+v, want none", obnoxiousKey, w)
	}

	finishWeights(t, simStore, id)
	out = s.messages(ctx, h.owner, time.Now())
	var ack MessageEntry
	for _, a := range messagesOf(out, messageAck) {
		if a.Character == obnoxiousKey {
			ack = a
		}
	}
	wantRequested := rfc3339(t, h, `select created_at from addon_sim_requests where sim_id = $1`, id)
	if ack.SimRequestState != sims.SimRequestStateDone || ack.RequestedAt != wantRequested {
		t.Fatalf("ack = %+v, want state done and requested_at %s (the site's receive time)", ack, wantRequested)
	}
	if ack.ExportSeenAt != rfc3339(t, h, `select updated_at from addon_exports where character_key = $1`, obnoxiousKey) {
		t.Fatalf("ack export_seen_at = %s, want the stored export's time", ack.ExportSeenAt)
	}

	w := weightsFor(out, obnoxiousKey)
	if w == nil {
		t.Fatalf("no weights message for %s: %+v", obnoxiousKey, messagesOf(out, messageWeights))
	}
	if w.GeneratedAt != rfc3339(t, h, `select created_at from sims where id = $1`, id) || w.ExportSeenAt != ack.ExportSeenAt ||
		w.Spec != "warrior-arms" || len(w.Weights) != 2 || w.Weights[1] != (WeightEntry{Stat: "strength", Weight: 2.2}) {
		t.Fatalf("in-game weights = %+v, want the run's own weights, its queue time and export_seen_at %s", w, ack.ExportSeenAt)
	}
	// The request and its run are recorded together: for the run in use, requested_at equals
	// generated_at, so only a newer request reads as later.
	if ack.RequestedAt != w.GeneratedAt {
		t.Fatalf("requested_at %s, generated_at %s: want equal for the request whose run made the weights",
			ack.RequestedAt, w.GeneratedAt)
	}

	web := weightsFor(out, l63Key)
	if web == nil {
		t.Fatalf("no fallback weights message for %s: %+v", l63Key, messagesOf(out, messageWeights))
	}
	if web.ExportSeenAt != "" || web.GeneratedAt != rfc3339(t, h, `select created_at from sims where id = 'l63-web-weights'`) ||
		web.Weights[1].Weight != 1.9 {
		t.Fatalf("website weights = %+v, want the website run's weights and time with no export_seen_at", web)
	}
	if n := len(messagesOf(out, messageWeights)); n != 2 {
		t.Fatalf("weights messages = %d, want one per character", n)
	}
}

// A run that errored says failed on the ack, and gives no weights.
func TestAnErroredInGameRunSaysFailed(t *testing.T) {
	h := newOutboxHarness(t, true)
	cleanL63(t, h.harness)
	h.pairDevice("l63-dev", false)
	simStore := &sims.Store{Pool: h.pool}
	id := inGameRequest(t, h)
	if err := simStore.Fail(context.Background(), id, "the engine stopped"); err != nil {
		t.Fatal(err)
	}
	s := &Service{Store: h.store, Sims: simStore, Log: quietLog()}
	out := s.messages(context.Background(), h.owner, time.Now())
	acks := messagesOf(out, messageAck)
	if len(acks) != 1 || acks[0].SimRequestState != sims.SimRequestStateFailed || acks[0].RequestedAt == "" {
		t.Fatalf("ack = %+v, want failed with requested_at", acks)
	}
	if w := weightsFor(out, obnoxiousKey); w != nil {
		t.Fatalf("weights from an errored run = %+v, want none", w)
	}
}

// A request recorded before migration 0038 has no export_seen_at: its weights message
// carries generated_at and leaves export_seen_at out.
func TestAPre0038RequestsWeightsHaveNoExportSeenAt(t *testing.T) {
	h := newOutboxHarness(t, true)
	cleanL63(t, h.harness)
	simStore := &sims.Store{Pool: h.pool}
	id := inGameRequest(t, h)
	var recorded *time.Time
	if err := h.pool.QueryRow(context.Background(),
		`select export_seen_at from addon_sim_requests where sim_id = $1`, id).Scan(&recorded); err != nil {
		t.Fatal(err)
	}
	if recorded == nil {
		t.Fatal("the request recorded no export_seen_at")
	}
	if _, err := h.pool.Exec(context.Background(),
		`update addon_sim_requests set export_seen_at = null where sim_id = $1`, id); err != nil {
		t.Fatal(err)
	}
	finishWeights(t, simStore, id)
	s := &Service{Store: h.store, Sims: simStore, Log: quietLog()}
	w := weightsFor(s.messages(context.Background(), h.owner, time.Now()), obnoxiousKey)
	if w == nil || w.ExportSeenAt != "" || w.GeneratedAt == "" {
		t.Fatalf("weights = %+v, want generated_at and no export_seen_at", w)
	}
	b, err := json.Marshal(w)
	if err != nil {
		t.Fatal(err)
	}
	var wire map[string]json.RawMessage
	if err := json.Unmarshal(b, &wire); err != nil {
		t.Fatal(err)
	}
	if _, ok := wire["export_seen_at"]; ok {
		t.Fatalf("export_seen_at is on the wire: %s", b)
	}
}
