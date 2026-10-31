package addon

// Inbox v2 (docs/contracts/addon-v2.md contract 3) against the shared test database: real
// stores, seeded rows under the "l63-" prefix, cleaned up by each test.

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

const (
	l63Name   = "Lastraid"
	l63Key    = "us/normal/lastraid"
	l63Export = "FS1:1.60.1.70291:warrior:human:0/0/0:chest=6627"
)

// repoBuilds is the repository's own data/builds, the directory the API ships as TreeDataDir.
const repoBuilds = "../../../data/builds"

var (
	repoGearOnce sync.Once
	repoGearData *trees.Data
	repoGearErr  error
)

// useRepoGearData points the sims package's gear readers at the repository's real build
// data, once per test binary (loading every build takes a few seconds).
func useRepoGearData(t *testing.T) {
	t.Helper()
	repoGearOnce.Do(func() { repoGearData, repoGearErr = trees.Load(repoBuilds) })
	if repoGearErr != nil {
		t.Fatal(repoGearErr)
	}
	sims.SetGearData(repoGearData, repoBuilds)
	t.Cleanup(func() { sims.SetGearData(nil, "") })
}

func cleanL63(t *testing.T, h *harness) {
	t.Helper()
	t.Cleanup(func() {
		bg := context.Background()
		h.pool.Exec(bg, `delete from rating_scores where report_id like 'l63-%'`)
		h.pool.Exec(bg, `delete from fights where report_id like 'l63-%'`)
		h.pool.Exec(bg, `delete from reports where id like 'l63-%'`)
		h.pool.Exec(bg, `delete from devices where id like 'l63-%'`)
	})
}

func (h *harness) exportL63(code string) {
	h.t.Helper()
	if err := h.store.PutExports(context.Background(), h.owner, []Export{
		{Name: l63Name, Ruleset: "normal", Region: "us", Export: code},
	}); err != nil {
		h.t.Fatal(err)
	}
}

// ratedFight seeds one rated boss fight: the report (once), the fight row with its boss
// name, and the player's rating_scores row carrying what the rater stores.
type ratedFight struct {
	report, visibility, zone, boss string
	index                          int
	encounter                      int64
	foughtAt                       time.Time
	overall                        float64
	kill, insufficient             bool
	biggestLoss                    string
	lossSeconds, deaths            int
	missed                         []string
	// adherence is the stored rotation adherence; nil seeds a row rated before migration
	// 0037, which has none.
	adherence *float64
	// legacy seeds a row rated before migration 0036: its facts are NULL.
	legacy bool
}

func (h *harness) seedRatedFight(f ratedFight) {
	h.t.Helper()
	ctx := context.Background()
	exec := func(sql string, args ...any) {
		h.t.Helper()
		if _, err := h.pool.Exec(ctx, sql, args...); err != nil {
			h.t.Fatalf("%s: %v", sql, err)
		}
	}
	exec(`insert into reports (id, owner_id, title, visibility, zone, status, created_at)
		values ($1, $2, 'raid', $3, $4, 'complete', $5) on conflict (id) do nothing`,
		f.report, h.owner, f.visibility, f.zone, f.foughtAt)
	exec(`insert into fights (report_id, fight_index, encounter_id, name, difficulty, kill, start_ms)
		values ($1, $2, $3, $4, 3, $5, $6)`, f.report, f.index, f.encounter, f.boss, f.kill, f.foughtAt.UnixMilli())
	if err := db.EnsureRatingsPartition(ctx, h.pool, f.foughtAt); err != nil {
		h.t.Fatal(err)
	}
	exec(`insert into rating_scores (report_id, fight_index, player_key, player_name, class, spec, role,
		   encounter_id, difficulty, kill, overall, overall_uncapped, components, model_version, fought_at,
		   insufficient, deaths, biggest_loss, loss_seconds, consumables_missed, adherence)
		values ($1, $2, $3, $4, 'Warrior', 'Arms', 'dps', $5, 3, $6, $7, $7, '[]', 'test', $8, $9, $10,
		   nullif($11, ''), $12, $13, $14)`,
		f.report, f.index, l63Key, l63Name, f.encounter, f.kill, f.overall, f.foughtAt, f.insufficient,
		f.deaths, f.biggestLoss, f.lossSeconds, f.missed, f.adherence)
	if f.legacy {
		exec(`update rating_scores set deaths = null, biggest_loss = null, loss_seconds = null,
			  consumables_missed = null where report_id = $1 and fight_index = $2`, f.report, f.index)
	}
}

// messagesOf is the messages of one type, in the order the inbox lists them.
func messagesOf(out []MessageEntry, typ string) []MessageEntry {
	var got []MessageEntry
	for _, m := range out {
		if m.Type == typ {
			got = append(got, m)
		}
	}
	return got
}

func TestLastRaidMessageNamesEveryBossOfTheNewestRankedRaid(t *testing.T) {
	h := newHarness(t)
	cleanL63(t, h)
	h.exportL63(l63Export)
	now := time.Now().UTC()
	raid := now.Add(-26 * time.Hour)
	// The raid: a Lucifron kill and a Magmadar wipe, in fight order. Lucifron was rated with
	// its rotation adherence; Magmadar was rated before migration 0037 and has none.
	adherence := 0.82
	h.seedRatedFight(ratedFight{report: "l63-raid", visibility: "public", zone: "Molten Core", boss: "Lucifron",
		index: 1, encounter: 663, foughtAt: raid, overall: 71.4, kill: true, biggestLoss: "downtime",
		lossSeconds: 18, missed: []string{"flask"}, adherence: &adherence})
	h.seedRatedFight(ratedFight{report: "l63-raid", visibility: "public", zone: "Molten Core", boss: "Magmadar",
		index: 2, encounter: 664, foughtAt: raid.Add(20 * time.Minute), overall: 48.6, biggestLoss: "dead early",
		lossSeconds: 95, deaths: 1, missed: []string{}})
	// An insufficient card on the same night is not a rated fight.
	h.seedRatedFight(ratedFight{report: "l63-raid", visibility: "public", zone: "Molten Core", boss: "Gehennas",
		index: 3, encounter: 665, foughtAt: raid.Add(40 * time.Minute), insufficient: true, missed: []string{}})
	// An older raid, a newer private one, and one past the window are none of them the last raid.
	h.seedRatedFight(ratedFight{report: "l63-older", visibility: "guild", zone: "Onyxia's Lair", boss: "Onyxia",
		index: 1, encounter: 1084, foughtAt: now.Add(-5 * 24 * time.Hour), overall: 90, kill: true, missed: []string{}})
	h.seedRatedFight(ratedFight{report: "l63-private", visibility: "private", zone: "Molten Core", boss: "Lucifron",
		index: 1, encounter: 663, foughtAt: now.Add(-1 * time.Hour), overall: 99, kill: true, missed: []string{}})

	s := &Service{Store: h.store, Log: quietLog()}
	got := messagesOf(s.messages(context.Background(), h.owner, now), messageLastRaid)
	if len(got) != 1 {
		t.Fatalf("last_raid messages = %+v, want one", got)
	}
	m := got[0]
	if m.Character != l63Key || m.ReportID != "l63-raid" || m.Date != raid.Add(20*time.Minute).Format(time.DateOnly) {
		t.Fatalf("last_raid = %+v, want l63-raid for %s on %s", m, l63Key, raid.Format(time.DateOnly))
	}
	want := []LastRaidBoss{
		{Name: "Lucifron", Percentile: 71, BiggestLoss: "downtime", LossSeconds: 18, Deaths: 0,
			ConsumablesMissed: []string{"flask"}, Adherence: &adherence},
		{Name: "Magmadar", Percentile: 49, BiggestLoss: "dead early", LossSeconds: 95, Deaths: 1,
			ConsumablesMissed: []string{}},
	}
	if !reflect.DeepEqual(m.Bosses, want) {
		t.Fatalf("bosses = %+v\nwant %+v", m.Bosses, want)
	}
	// adherence travels when stored and is absent for a fight rated before 0037;
	// consumables_missed is a list even empty.
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var wire struct {
		Bosses []map[string]json.RawMessage `json:"bosses"`
	}
	if err := json.Unmarshal(b, &wire); err != nil {
		t.Fatal(err)
	}
	if got := string(wire.Bosses[0]["adherence"]); got != "0.82" {
		t.Errorf("Lucifron's adherence on the wire = %q, want 0.82: %s", got, b)
	}
	if _, ok := wire.Bosses[1]["adherence"]; ok {
		t.Errorf("adherence is on the wire for a fight rated before 0037: %s", b)
	}
	if string(wire.Bosses[1]["consumables_missed"]) != "[]" {
		t.Errorf("consumables_missed = %s, want []", wire.Bosses[1]["consumables_missed"])
	}
}

func TestLastRaidFightsCarryTheirFightIndex(t *testing.T) {
	h := newHarness(t)
	cleanL63(t, h)
	now := time.Now().UTC()
	for _, index := range []int{3, 7} {
		h.seedRatedFight(ratedFight{report: "l63-index", visibility: "public", zone: "Molten Core", boss: "Lucifron",
			index: index, encounter: 663, foughtAt: now.Add(-time.Duration(10-index) * time.Hour), overall: 50, kill: true,
			missed: []string{}})
	}
	fights, err := h.store.LastRaidFights(context.Background(), []string{l63Key}, now.Add(-LastRaidWindow))
	if err != nil || len(fights) != 2 || fights[0].FightIndex != 3 || fights[1].FightIndex != 7 {
		t.Fatalf("%+v %v", fights, err)
	}
}

func TestACharacterWithNoRecentRaidGetsNoLastRaid(t *testing.T) {
	h := newHarness(t)
	cleanL63(t, h)
	h.exportL63(l63Export)
	now := time.Now().UTC()
	h.seedRatedFight(ratedFight{report: "l63-stale", visibility: "public", zone: "Molten Core", boss: "Lucifron",
		index: 1, encounter: 663, foughtAt: now.Add(-15 * 24 * time.Hour), overall: 71, kill: true, missed: []string{}})
	s := &Service{Store: h.store, Log: quietLog()}
	out := s.messages(context.Background(), h.owner, now)
	if got := messagesOf(out, messageLastRaid); len(got) != 0 {
		t.Fatalf("last_raid = %+v, want none past %v", got, LastRaidWindow)
	}
	if got := messagesOf(out, messageLootValues); len(got) != 0 {
		t.Fatalf("loot_values = %+v, want none without a last raid", got)
	}
}

func (h *harness) pairDevice(id string, revoked bool) {
	h.t.Helper()
	var revokedAt *time.Time
	if revoked {
		at := time.Now()
		revokedAt = &at
	}
	if _, err := h.pool.Exec(context.Background(),
		`insert into devices (id, user_id, name, token_hash, revoked_at) values ($1, $2, 'companion', $3, $4)`,
		id, h.owner, []byte(id), revokedAt); err != nil {
		h.t.Fatal(err)
	}
}

func TestAckCarriesWhenTheSiteStoredTheNewestExport(t *testing.T) {
	h := newHarness(t)
	cleanL63(t, h)
	h.exportL63(l63Export)
	h.pairDevice("l63-dev", false)
	var stored time.Time
	if err := h.pool.QueryRow(context.Background(),
		`select updated_at from addon_exports where character_key = $1`, l63Key).Scan(&stored); err != nil {
		t.Fatal(err)
	}
	s := &Service{Store: h.store, Log: quietLog()}
	got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageAck)
	if len(got) != 1 {
		t.Fatalf("ack = %+v, want one", got)
	}
	if got[0].Character != l63Key || got[0].ExportSeenAt != stored.UTC().Format(time.RFC3339) ||
		got[0].SimRequestState != "" {
		t.Fatalf("ack = %+v, want %s seen at %s with no sim request", got[0], l63Key, stored.UTC().Format(time.RFC3339))
	}
}

func TestAnAccountWithoutAPairedCompanionGetsNoAck(t *testing.T) {
	h := newHarness(t)
	cleanL63(t, h)
	h.exportL63(l63Export)
	s := &Service{Store: h.store, Log: quietLog()}
	if got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageAck); len(got) != 0 {
		t.Fatalf("ack with no device = %+v, want none", got)
	}
	h.pairDevice("l63-revoked", true)
	if got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageAck); len(got) != 0 {
		t.Fatalf("ack with only a revoked device = %+v, want none", got)
	}
	// A signed-in phone app holds a device token too, but it carries nothing into the game.
	if _, err := h.pool.Exec(context.Background(),
		`insert into devices (id, user_id, name, platform, token_hash, kind) values ('l63-app', $1, 'iPhone', 'ios', 'l63-app', 'app')`,
		h.owner); err != nil {
		t.Fatal(err)
	}
	if got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageAck); len(got) != 0 {
		t.Fatalf("ack with only an app device = %+v, want none", got)
	}
}

// The ack says what became of the character's newest in-game sim request, refused ones too.
func TestAckSaysWhatBecameOfTheNewestSimRequest(t *testing.T) {
	for _, c := range []struct {
		name    string
		allowed bool
		want    string
	}{
		{"queued", true, "queued"},
		{"needs premium", false, "needs_premium"},
	} {
		t.Run(c.name, func(t *testing.T) {
			h := newOutboxHarness(t, c.allowed)
			cleanL63(t, h.harness)
			h.pairDevice("l63-dev", false)
			h.post(withOutbox(string(contractFixture(t, "outbox-v1.json"))))
			s := &Service{Store: h.store, Log: quietLog()}
			got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageAck)
			if len(got) != 1 || got[0].Character != obnoxiousKey || got[0].SimRequestState != c.want {
				t.Fatalf("ack = %+v, want %s for %s", got, c.want, obnoxiousKey)
			}
		})
	}
	t.Run("duplicate", func(t *testing.T) {
		h := newOutboxHarness(t, true)
		cleanL63(t, h.harness)
		h.pairDevice("l63-dev", false)
		fixture := string(contractFixture(t, "outbox-v1.json"))
		h.post(withOutbox(fixture))
		// The identical request posted again is the same request: still queued.
		h.post(withOutbox(fixture))
		s := &Service{Store: h.store, Log: quietLog()}
		if got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageAck); len(got) != 1 ||
			got[0].SimRequestState != "queued" {
			t.Fatalf("ack after a re-post = %+v, want queued", got)
		}
		// A new request inside the hour did nothing, and the ack says why.
		h.post(withOutbox(strings.Replace(fixture, "2026-11-04 14:10", "2026-11-04 14:40", 1)))
		if got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageAck); len(got) != 1 ||
			got[0].SimRequestState != "duplicate" {
			t.Fatalf("ack after a second request = %+v, want duplicate", got)
		}
	})
}

func TestSimResultCarriesTheNewestFinishedInGameRun(t *testing.T) {
	h := newOutboxHarness(t, true)
	cleanL63(t, h.harness)
	h.post(withOutbox(string(contractFixture(t, "outbox-v1.json"))))
	reqs := h.simRequests()
	if len(reqs) != 1 {
		t.Fatalf("sim requests = %+v", reqs)
	}
	s := &Service{Store: h.store, Log: quietLog()}
	if got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageSimResult); len(got) != 0 {
		t.Fatalf("sim_result before the run finished = %+v, want none", got)
	}
	simStore := &sims.Store{Pool: h.pool}
	res, _, err := simStore.Get(context.Background(), reqs[0].simID)
	if err != nil {
		t.Fatal(err)
	}
	// An error that rounds to 0.0 still travels: the addon is told the figure is exact.
	res.DPS = simapi.Estimate{Mean: 412.63, Error: 0.04}
	if err := simStore.Finish(context.Background(), reqs[0].simID, res); err != nil {
		t.Fatal(err)
	}
	var created time.Time
	if err := h.pool.QueryRow(context.Background(), `select created_at from sims where id = $1`, reqs[0].simID).
		Scan(&created); err != nil {
		t.Fatal(err)
	}
	got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageSimResult)
	if len(got) != 1 {
		t.Fatalf("sim_result = %+v, want one", got)
	}
	m := got[0]
	if m.Character != obnoxiousKey || m.Spec != "warrior-arms" || m.DPS == nil || *m.DPS != 412.6 ||
		m.Error == nil || *m.Error != 0 ||
		m.GeneratedAt != created.UTC().Format(time.RFC3339) {
		t.Fatalf("sim_result = %+v", m)
	}
	// The in-game request queues a weights run with no bulk section: no per-item figures, so
	// slots is an empty list.
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var wire map[string]json.RawMessage
	if err := json.Unmarshal(b, &wire); err != nil {
		t.Fatal(err)
	}
	if string(wire["slots"]) != "[]" {
		t.Fatalf("slots = %s, want []", wire["slots"])
	}
	if string(wire["error"]) != "0" || string(wire["dps"]) != "412.6" {
		t.Fatalf("dps = %s, error = %s, want 412.6 and 0 on the wire", wire["dps"], wire["error"])
	}
}

func TestSlotDeltasReadsABulkRunsSingleSlotCombos(t *testing.T) {
	raw, err := json.Marshal([]simapi.Combo{
		{Substitutions: []simapi.Substitution{
			{Kind: simapi.SubstitutionItem, Slot: "chest", ItemID: 1}, {Kind: simapi.SubstitutionItem, Slot: "legs", ItemID: 2},
		}, Delta: simapi.Estimate{Mean: 80}},
		{Substitutions: []simapi.Substitution{{Kind: simapi.SubstitutionItem, Slot: "chest", ItemID: 6627}},
			Delta: simapi.Estimate{Mean: 4.24}},
		{Substitutions: []simapi.Substitution{{Kind: simapi.SubstitutionItem, Slot: "chest", ItemID: 9}},
			Delta: simapi.Estimate{Mean: 2}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := slotDeltas(raw); !reflect.DeepEqual(got, []SimSlot{{Slot: "chest", ItemID: 6627, Delta: 4.2}}) {
		t.Fatalf("slots = %+v", got)
	}
	if got := slotDeltas(nil); got == nil || len(got) != 0 {
		t.Fatalf("a plain run's slots = %#v, want an empty list", got)
	}
}

func TestWeightsCarryTheSpecsDPSPerPoint(t *testing.T) {
	useRepoGearData(t)
	h := newHarness(t)
	h.exportL63(l63Export)
	simStore := &sims.Store{Pool: h.pool}
	save := func(id, spec string) {
		t.Helper()
		if err := simStore.Save(context.Background(), id, &h.owner, "", simapi.SimResult{
			Request: simapi.SimRequest{Spec: spec, Character: simapi.CharacterSpec{Name: l63Name},
				Weights: &simapi.WeightsSpec{Stats: []string{"attack_power"}, Reference: "attack_power"}},
			Weights: []simapi.StatWeight{{Stat: "attack_power", Weight: 1}, {Stat: "strength", Weight: 2.2}},
		}); err != nil {
			t.Fatal(err)
		}
	}
	t.Cleanup(func() { h.pool.Exec(context.Background(), `delete from sims where id like 'l63-%'`) })
	s := &Service{Store: h.store, Sims: simStore, Log: quietLog()}

	save("l63-weights-arms", "warrior-arms")
	got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageWeights)
	if len(got) != 1 || got[0].Character != l63Key || got[0].DPSPerPoint == nil || *got[0].DPSPerPoint != 0.2222 {
		t.Fatalf("weights = %+v, want warrior-arms' 0.2222 DPS per attack power", got)
	}

	// A tank spec has no entry in dps-per-point.json: the field is absent.
	h.pool.Exec(context.Background(), `update sims set created_at = now() - interval '1 hour' where id = 'l63-weights-arms'`)
	save("l63-weights-prot", "warrior-protection")
	got = messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageWeights)
	if len(got) != 1 || got[0].Spec != "warrior-protection" || got[0].DPSPerPoint != nil {
		t.Fatalf("weights = %+v, want warrior-protection with no dps_per_point", got)
	}
}

func TestLootValuesPriceTheLastRaidsLoot(t *testing.T) {
	useRepoGearData(t)
	h := newHarness(t)
	cleanL63(t, h)
	h.exportL63(l63Export)
	now := time.Now().UTC()
	h.seedRatedFight(ratedFight{report: "l63-raid", visibility: "public", zone: "Molten Core", boss: "Lucifron",
		index: 1, encounter: 663, foughtAt: now.Add(-2 * time.Hour), overall: 71, kill: true,
		biggestLoss: "downtime", lossSeconds: 18, missed: []string{}})
	s := &Service{Store: h.store, Log: quietLog()}
	got := messagesOf(s.messages(context.Background(), h.owner, now), messageLootValues)
	if len(got) != 1 {
		t.Fatalf("loot_values = %+v, want one", got)
	}
	m := got[0]
	want, ok, err := sims.LootValues(context.Background(), "Molten Core", l63Export, "warrior-arms", nil)
	if err != nil || !ok {
		t.Fatalf("LootValues: ok=%v err=%v", ok, err)
	}
	if m.Character != l63Key || m.Raid != "Molten Core" || len(m.Items) != len(want.Items) || len(m.Items) == 0 {
		t.Fatalf("loot_values = %s/%s with %d items, want Molten Core with %d", m.Character, m.Raid,
			len(m.Items), len(want.Items))
	}
	for id, v := range want.Items {
		if m.Items[id] != (LootItemValue{Value: v.Value, Unit: v.Unit, Planned: v.Planned}) {
			t.Fatalf("item %d = %+v, want %+v", id, m.Items[id], v)
		}
		if v.Unit != "dps" {
			t.Fatalf("item %d unit = %q, want dps: warrior-arms is in dps-per-point.json", id, v.Unit)
		}
	}
}

func TestLootValuesNeedARaidTheZoneNames(t *testing.T) {
	useRepoGearData(t)
	h := newHarness(t)
	cleanL63(t, h)
	h.exportL63(l63Export)
	now := time.Now().UTC()
	h.seedRatedFight(ratedFight{report: "l63-raid", visibility: "public", zone: "Somewhere Else", boss: "Lucifron",
		index: 1, encounter: 663, foughtAt: now.Add(-2 * time.Hour), overall: 71, kill: true, missed: []string{}})
	s := &Service{Store: h.store, Log: quietLog()}
	out := s.messages(context.Background(), h.owner, now)
	if len(messagesOf(out, messageLastRaid)) != 1 {
		t.Fatal("the raid is still the last raid")
	}
	if got := messagesOf(out, messageLootValues); len(got) != 0 {
		t.Fatalf("loot_values = %+v, want none for a zone no loot table names", got)
	}
}

func TestTheInboxAnswersFormatTwo(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodGet, "/v1/addon/inbox", "")
	var body map[string]json.RawMessage
	h.data(res, &body)
	if string(body["format"]) != "2" {
		t.Fatalf("format = %s, want 2", body["format"])
	}
	var generated string
	if err := json.Unmarshal(body["generated_at"], &generated); err != nil {
		t.Fatal(err)
	}
	at, err := time.Parse(time.RFC3339, generated)
	if err != nil || at.Location() != time.UTC || time.Since(at) > time.Minute {
		t.Fatalf("generated_at = %q, want an RFC 3339 UTC time of now", generated)
	}
	if string(body["builds"]) != "[]" || string(body["messages"]) != "[]" {
		t.Fatalf("builds = %s, messages = %s, want two empty lists", body["builds"], body["messages"])
	}
}

// A fight rated before migration 0036 has no stored facts: it is left out rather than
// shown as zeros, and a last raid of nothing but such fights gets no message.
func TestLastRaidLeavesOutFightsRatedBeforeTheFactsWereStored(t *testing.T) {
	h := newHarness(t)
	cleanL63(t, h)
	h.exportL63(l63Export)
	now := time.Now().UTC()
	raid := now.Add(-3 * time.Hour)
	h.seedRatedFight(ratedFight{report: "l63-mixed", visibility: "public", zone: "Molten Core", boss: "Lucifron",
		index: 1, encounter: 663, foughtAt: raid, overall: 71, kill: true, biggestLoss: "downtime",
		lossSeconds: 18, missed: []string{}})
	h.seedRatedFight(ratedFight{report: "l63-mixed", visibility: "public", zone: "Molten Core", boss: "Magmadar",
		index: 2, encounter: 664, foughtAt: raid.Add(20 * time.Minute), overall: 60, kill: true, legacy: true})
	s := &Service{Store: h.store, Log: quietLog()}
	got := messagesOf(s.messages(context.Background(), h.owner, now), messageLastRaid)
	if len(got) != 1 || len(got[0].Bosses) != 1 || got[0].Bosses[0].Name != "Lucifron" {
		t.Fatalf("last_raid = %+v, want Lucifron alone", got)
	}

	// The newest raid is all legacy fights: no message, not an older raid in its place.
	h.seedRatedFight(ratedFight{report: "l63-legacy", visibility: "public", zone: "Molten Core", boss: "Lucifron",
		index: 1, encounter: 663, foughtAt: now.Add(-1 * time.Hour), overall: 80, kill: true, legacy: true})
	if got := messagesOf(s.messages(context.Background(), h.owner, now), messageLastRaid); len(got) != 0 {
		t.Fatalf("last_raid = %+v, want none for a raid with no stored facts", got)
	}
}

// dps_per_point travels only when the run is normalised to the file's own reference stat.
func TestWeightsCarryNoDPSPerPointForAnotherReference(t *testing.T) {
	useRepoGearData(t)
	h := newHarness(t)
	h.exportL63(l63Export)
	simStore := &sims.Store{Pool: h.pool}
	t.Cleanup(func() { h.pool.Exec(context.Background(), `delete from sims where id like 'l63-%'`) })
	// warrior-arms' reference in dps-per-point.json is attack power; this run is in strength.
	if err := simStore.Save(context.Background(), "l63-weights-str", &h.owner, "", simapi.SimResult{
		Request: simapi.SimRequest{Spec: "warrior-arms", Character: simapi.CharacterSpec{Name: l63Name},
			Weights: &simapi.WeightsSpec{Stats: []string{"strength", "attack_power"}, Reference: "strength"}},
		Weights: []simapi.StatWeight{{Stat: "strength", Weight: 1}, {Stat: "attack_power", Weight: 0.45}},
	}); err != nil {
		t.Fatal(err)
	}
	s := &Service{Store: h.store, Sims: simStore, Log: quietLog()}
	got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageWeights)
	if len(got) != 1 || got[0].Spec != "warrior-arms" || got[0].DPSPerPoint != nil {
		t.Fatalf("weights = %+v, want warrior-arms with no dps_per_point", got)
	}
}

// A request the engine cannot sim is recorded, and the ack says "invalid".
func TestAckSaysInvalidForARequestThatCannotBeSimmed(t *testing.T) {
	// A request built for an engine other than the one this API runs fails lane validation.
	h := newOutboxHarnessOn(t, true, "not-this-engine")
	cleanL63(t, h.harness)
	h.pairDevice("l63-dev", false)
	if got := h.post(withOutbox(string(contractFixture(t, "outbox-v1.json")))).Outbox.SimRequest; got != "invalid" {
		t.Fatalf("sim_request = %q, want invalid", got)
	}
	s := &Service{Store: h.store, Log: quietLog()}
	got := messagesOf(s.messages(context.Background(), h.owner, time.Now()), messageAck)
	if len(got) != 1 || got[0].SimRequestState != "invalid" {
		t.Fatalf("ack = %+v, want invalid", got)
	}
}

// An adherence of 0 (ten or more ability casts, none in the rotation) is a fact and travels
// as "adherence":0, not as an absent field.
func TestLastRaidSendsAZeroAdherence(t *testing.T) {
	h := newHarness(t)
	cleanL63(t, h)
	h.exportL63(l63Export)
	now := time.Now().UTC()
	zero := 0.0
	h.seedRatedFight(ratedFight{report: "l63-zero", visibility: "public", zone: "Molten Core", boss: "Lucifron",
		index: 1, encounter: 663, foughtAt: now.Add(-2 * time.Hour), overall: 40, kill: true,
		missed: []string{}, adherence: &zero})
	s := &Service{Store: h.store, Log: quietLog()}
	got := messagesOf(s.messages(context.Background(), h.owner, now), messageLastRaid)
	if len(got) != 1 || len(got[0].Bosses) != 1 {
		t.Fatalf("last_raid = %+v, want Lucifron", got)
	}
	b, err := json.Marshal(got[0].Bosses[0])
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), `"adherence":0`) {
		t.Fatalf("boss = %s, want \"adherence\":0", b)
	}
}

// An upload the site already has still moves export_seen_at, because the ack tells the addon
// when the site last received the character; export_changed_at stays where the export last
// changed, and moves again when the export text differs.
func TestAckSeenAtMovesOnAnUnchangedReupload(t *testing.T) {
	h := newHarness(t)
	cleanL63(t, h)
	h.exportL63(l63Export)
	h.pairDevice("l63-dev", false)
	ctx := context.Background()
	old := time.Now().Add(-12 * time.Hour).UTC().Truncate(time.Second)
	if _, err := h.pool.Exec(ctx,
		`update addon_exports set seen_at = $2, updated_at = $2 where character_key = $1`, l63Key, old); err != nil {
		t.Fatal(err)
	}
	s := &Service{Store: h.store, Log: quietLog()}
	ack := func() MessageEntry {
		got := messagesOf(s.messages(ctx, h.owner, time.Now()), messageAck)
		if len(got) != 1 {
			t.Fatalf("ack = %+v, want one", got)
		}
		return got[0]
	}
	oldText := old.Format(time.RFC3339)
	if a := ack(); a.ExportSeenAt != oldText || a.ExportChangedAt != oldText {
		t.Fatalf("ack before = %+v, want both %s", a, oldText)
	}

	h.exportL63(l63Export)
	a := ack()
	seen, err := time.Parse(time.RFC3339, a.ExportSeenAt)
	if err != nil || !seen.After(old.Add(time.Hour)) {
		t.Fatalf("export_seen_at after an unchanged re-upload = %q, want a time just now", a.ExportSeenAt)
	}
	if a.ExportChangedAt != oldText {
		t.Fatalf("export_changed_at = %q after an unchanged re-upload, want it left at %s", a.ExportChangedAt, oldText)
	}

	h.exportL63(l63Export + "|changed")
	a = ack()
	changed, err := time.Parse(time.RFC3339, a.ExportChangedAt)
	if err != nil || !changed.After(old.Add(time.Hour)) {
		t.Fatalf("export_changed_at after a changed upload = %q, want a time just now", a.ExportChangedAt)
	}
}
