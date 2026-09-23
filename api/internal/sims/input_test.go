package sims

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// dirGetter reads objects back out of the harness's local directory,
// the way *r2.Client reads them out of the bucket.
type dirGetter struct{ root string }

func (d dirGetter) Get(_ context.Context, key string) (io.ReadCloser, error) {
	return os.Open(d.root + "/" + key)
}

// seedExport writes one addon-sourced export for a character, at is used for both
// captured_at and updated_at (the two coincide for a freshly-written addon export).
func seedExport(h *harness, key, region, ruleset, name, export string, at time.Time) {
	h.t.Helper()
	seedExportWithSource(h, key, region, ruleset, name, export, "addon", at)
}

// seedExportWithSource is seedExport plus an explicit source ("addon" or "blizzard"),
// for tests asserting on SimInput's reported source.
func seedExportWithSource(h *harness, key, region, ruleset, name, export, source string, at time.Time) {
	h.t.Helper()
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $8)
		 on conflict (character_key) do update set export = excluded.export,
		   source = excluded.source, captured_at = excluded.captured_at, updated_at = excluded.updated_at`,
		key, h.owner, region, ruleset, name, export, source, at); err != nil {
		h.t.Fatal(err)
	}
}

// seedFightRow writes one ranked fight row for a character, and the
// stored summary it points at, so the buff read has something to
// read. specSlug is the API's vocabulary ("warrior-fury"), the same
// one SimInput must answer with; the row is written the way the
// rankings writer really writes it - class and spec as separate
// columns, spec the display name ("Fury") - so the fixture exercises
// the same class/spec -> slug conversion the handler does.
func seedFightRow(h *harness, key, name, class, specSlug string, at time.Time) {
	h.t.Helper()
	reportID := "rep" + specSlug
	specName, ok := specNameFor(specSlug)
	if !ok {
		h.t.Fatalf("seedFightRow: no spec matches slug %q", specSlug)
	}
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec,
		   role, metric_dps, duration_ms, kill, fought_at, talent_split, trinkets, state)
		 values ($1, 1, $2, $3, $4, $5, 'dps', 1000, 180000, true, $6, '31/0/20',
		   array[19406, 13965]::bigint[], 'ok')`,
		reportID, key, name, class, specName, at); err != nil {
		h.t.Fatal(err)
	}
	body, err := json.Marshal(summary.Summary{
		FightIndex: 1,
		Combatants: []summary.CombatantRow{{
			GUID: "Player-1", Name: name, Spec: specName,
			RaidBuffs: []summary.AuraRef{
				{SpellID: 20217, Name: "Blessing of Kings"},
				{SpellID: 999999, Name: "Some Forever Aura Nobody Mapped"},
			},
			Consumables: []summary.AuraRef{{SpellID: 25289, Name: "Battle Shout"}},
		}},
	})
	if err != nil {
		h.t.Fatal(err)
	}
	if err := h.files.Put(h.t.Context(), store.Keys{ReportID: reportID}.FightSummary(1),
		body, ResultPut); err != nil {
		h.t.Fatal(err)
	}
}

func TestSimInputPrefersTheNewerSource(t *testing.T) {
	h := newHarness(t)
	ensureMetricsPartition(h)
	h.service.Summaries = dirGetter{root: h.dir}
	key := "us/normal/baelgrim"
	old := time.Now().UTC().Add(-48 * time.Hour)
	recent := time.Now().UTC().Add(-1 * time.Hour)

	// Only a fight: the read falls back to it.
	seedFightRow(h, key, "Baelgrim", "Warrior", "warrior-fury", recent)
	var in Input
	h.data(h.do(http.MethodGet, "/v1/characters/us/normal/baelgrim/sim-input", "", nil), &in)
	if in.Source != "fight" || in.Spec != "warrior-fury" {
		t.Fatalf("fight only: %+v", in)
	}
	if in.Talents != "31/0/20" {
		t.Errorf("the fight's talent split was dropped: %q", in.Talents)
	}
	// The recorded auras arrive as engine ids, deduplicated, sorted,
	// with the unmappable one dropped.
	if len(in.Buffs) != 2 || in.Buffs[0] != "battle_shout" || in.Buffs[1] != "blessing_of_kings" {
		t.Fatalf("buffs %v", in.Buffs)
	}

	// An older export does not beat a newer fight.
	seedExport(h, key, "us", "normal", "Baelgrim", `{"slots":[]}`, old)
	h.data(h.do(http.MethodGet, "/v1/characters/us/normal/baelgrim/sim-input", "", nil), &in)
	if in.Source != "fight" {
		t.Fatalf("an older export won: %+v", in)
	}

	// A newer export does — and keeps the talents, the spec and the
	// buffs the fight recorded, because the export carries none this
	// repository can read.
	seedExport(h, key, "us", "normal", "Baelgrim", "FS1:1.60.1.69893:warrior:orc:0/5530515/0:head=12640,main_hand=21521", time.Now().UTC())
	h.data(h.do(http.MethodGet, "/v1/characters/us/normal/baelgrim/sim-input", "", nil), &in)
	if in.Source != "addon" {
		t.Fatalf("a newer export lost: %+v", in)
	}
	// The export is the addon's opaque FS1 string; it travels as a JSON string, never as
	// raw JSON (an FS1 code is not JSON, and RawMessage of one broke the encoder after
	// the headers were out).
	if string(in.Gear) != `"FS1:1.60.1.69893:warrior:orc:0/5530515/0:head=12640,main_hand=21521"` {
		t.Fatalf("gear: %s", in.Gear)
	}
	if in.Spec != "warrior-fury" || in.Talents != "31/0/20" {
		t.Fatalf("the addon branch dropped the spec or the talents: %+v", in)
	}
	if len(in.Buffs) != 2 {
		t.Fatalf("the addon branch dropped the buffs: %v", in.Buffs)
	}
}

func TestAnExportOnlyCharacterHasGearAndNoTalents(t *testing.T) {
	h := newHarness(t)
	h.service.Summaries = dirGetter{root: h.dir}
	// A character who has never parsed: the export is all there is,
	// and this repository cannot read talents out of it.
	seedExport(h, "us/normal/newbie", "us", "normal", "Newbie", "FS1:1.60.1.69893:rogue:human:0/0/0:head=7", time.Now().UTC())
	var in Input
	h.data(h.do(http.MethodGet, "/v1/characters/us/normal/newbie/sim-input", "", nil), &in)
	if in.Source != "addon" || string(in.Gear) != `"FS1:1.60.1.69893:rogue:human:0/0/0:head=7"` {
		t.Fatalf("export only: %+v", in)
	}
	if in.Talents != "" || in.Spec != "" {
		t.Fatalf("talents or spec appeared from nowhere: %+v", in)
	}
	if in.Buffs == nil {
		t.Fatal("buffs must be an empty list, never null")
	}
}

func TestSimInputServesTheCharacterWithNoBucket(t *testing.T) {
	h := newHarness(t)
	ensureMetricsPartition(h)
	h.service.Summaries = nil // a deployment with no R2 credentials
	seedFightRow(h, "us/normal/baelgrim", "Baelgrim", "Warrior", "warrior-fury", time.Now().UTC())
	var in Input
	h.data(h.do(http.MethodGet, "/v1/characters/us/normal/baelgrim/sim-input", "", nil), &in)
	if in.Spec != "warrior-fury" || in.Talents != "31/0/20" {
		t.Fatalf("the read failed without a bucket: %+v", in)
	}
	if len(in.Buffs) != 0 {
		t.Fatalf("buffs %v, want none without a bucket to read them from", in.Buffs)
	}
}

func TestAnUnknownCharacterIs404(t *testing.T) {
	h := newHarness(t)
	for _, path := range []string{
		"/v1/characters/us/normal/nobody/sim-input", // no rows anywhere
		"/v1/characters/mars/normal/x/sim-input",    // not a region
		"/v1/characters/us/roleplay/x/sim-input",    // not a ruleset
	} {
		if res := h.do(http.MethodGet, path, "", nil); res.StatusCode != http.StatusNotFound {
			t.Errorf("%s: status %d, want 404", path, res.StatusCode)
		}
	}
}

func TestSimInputIsNeverCachedForASignedInCaller(t *testing.T) {
	h := newHarness(t)
	ensureMetricsPartition(h)
	seedFightRow(h, "us/normal/baelgrim", "Baelgrim", "Mage", "mage-frost", time.Now().UTC())

	res := h.do(http.MethodGet, "/v1/characters/us/normal/baelgrim/sim-input", "", nil)
	res.Body.Close()
	if cc := res.Header.Get("Cache-Control"); cc != "private, no-store" {
		t.Errorf("signed in: Cache-Control %q", cc)
	}

	h.anonymous()
	res = h.do(http.MethodGet, "/v1/characters/us/normal/baelgrim/sim-input", "", nil)
	res.Body.Close()
	if cc := res.Header.Get("Cache-Control"); cc == "private, no-store" {
		t.Errorf("anonymous: Cache-Control %q; a public read may be cached briefly", cc)
	}
}
