package guilds

import (
	"net/http"
	"testing"
	"time"
)

// readinessRaidWorld is a planned raid a week out with: the officer In (fresh export), Kraggor
// Late (export synced ten days before the raid), Mirelle In (no export at all) and
// Sunderfel Out (with an export).
func readinessRaidWorld(t *testing.T) (raidWorld, RaidView) {
	t.Helper()
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	seedExport(t, w.h.pool, w.officer, officerKey, "us", "hardcore", "Thornhide")
	seedExport(t, w.h.pool, w.member, memberKeyA, "us", "hardcore", "Kraggor")
	seedExport(t, w.h.pool, w.member, memberKeyB, "us", "hardcore", "Sunderfel")
	stale := raid.StartsAt.Add(-10 * 24 * time.Hour)
	if _, err := w.h.pool.Exec(t.Context(),
		`update addon_exports set updated_at = $1 where character_key = $2`, stale, memberKeyA); err != nil {
		t.Fatal(err)
	}
	w.signUp(t, raid.ID, officerKey, "Thornhide", "druid", w.officer, SignupIn)
	w.signUp(t, raid.ID, memberKeyA, "Kraggor", "warrior", w.member, SignupLate)
	w.signUp(t, raid.ID, member2Key, "Mirelle", "priest", w.member2, SignupIn)
	w.signUp(t, raid.ID, memberKeyB, "Sunderfel", "mage", w.member, SignupOut)
	return w, raid
}

func readinessRowsByKey(v ReadinessView) map[string]ReadinessRow {
	out := map[string]ReadinessRow{}
	for _, r := range v.Rows {
		out[r.CharacterKey] = r
	}
	return out
}

func TestReadinessCutToARaidKeepsItsInAndLateSignups(t *testing.T) {
	w, raid := readinessRaidWorld(t)
	w.as(w.officer)
	res := w.h.do(http.MethodGet, w.path("/readiness?raid=%d", raid.ID), "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("filtered read = %d: %s", res.StatusCode, readBody(res))
	}
	var v ReadinessView
	w.h.data(res, &v)
	by := readinessRowsByKey(v)
	if len(v.Rows) != 3 || by[memberKeyB].CharacterKey != "" {
		t.Fatalf("rows = %+v", v.Rows)
	}
	if v.Raid == nil || v.Raid.ID != raid.ID || v.Raid.Title != raid.Title || !v.Raid.StartsAt.Equal(raid.StartsAt) ||
		!v.Raid.CutoffAt.Equal(raid.StartsAt) {
		t.Fatalf("raid = %+v", v.Raid)
	}
	if r := by[officerKey]; r.SyncedBeforeCutoff == nil || !*r.SyncedBeforeCutoff || r.NoExport || r.Failing != 0 {
		t.Fatalf("fresh row = %+v", r)
	}
	stale := by[memberKeyA]
	if stale.SyncedBeforeCutoff == nil || *stale.SyncedBeforeCutoff || stale.Failing != 1 {
		t.Fatalf("stale row = %+v", stale)
	}
	if stale.NudgeText == nil {
		t.Fatal("an officer's stale row has no nudge")
	}
	none := by[member2Key]
	if !none.NoExport || none.Name != "Mirelle" || none.Class != "priest" || none.SyncedBeforeCutoff != nil {
		t.Fatalf("no_export row = %+v", none)
	}
	if v.Rows[0].CharacterKey != memberKeyA || v.Rows[2].CharacterKey != member2Key {
		t.Fatalf("order = %s .. %s, want the stale row first and no_export last", v.Rows[0].CharacterKey, v.Rows[2].CharacterKey)
	}
	if _, ok := rawData(t, w.h.do(http.MethodGet, w.path("/readiness?raid=%d", raid.ID), ""))["next_raid"]; ok {
		t.Fatal("the filtered read carries next_raid")
	}
}

func TestUnfilteredReadinessNamesTheNextRaid(t *testing.T) {
	w, raid := readinessRaidWorld(t)
	w.as(w.officer)
	res := w.h.do(http.MethodGet, w.path("/readiness"), "")
	var out struct {
		ReadinessView
		NextRaid *ReadinessNextRaid `json:"next_raid"`
	}
	w.h.data(res, &out)
	if out.NextRaid == nil || out.NextRaid.ID != raid.ID || out.NextRaid.ZoneName != "Onyxia's Lair" ||
		out.NextRaid.Title != raid.Title || !out.NextRaid.StartsAt.Equal(raid.StartsAt) {
		t.Fatalf("next_raid = %+v", out.NextRaid)
	}
	by := readinessRowsByKey(out.ReadinessView)
	if len(out.Rows) != 3 || by[memberKeyB].CharacterKey == "" {
		t.Fatalf("unfiltered rows = %+v", out.Rows)
	}
	for _, r := range out.Rows {
		if r.SyncedBeforeCutoff != nil || r.NoExport {
			t.Fatalf("unfiltered row carries raid facts: %+v", r)
		}
	}
	if out.Raid != nil {
		t.Fatalf("unfiltered raid = %+v", out.Raid)
	}
}

func TestUnfilteredReadinessWithNoRaidSaysNull(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)
	data := rawData(t, w.h.do(http.MethodGet, w.path("/readiness"), ""))
	if got, ok := data["next_raid"]; !ok || string(got) != "null" {
		t.Fatalf("next_raid = %s (present %v)", got, ok)
	}
}

func TestReadinessForAnUnknownRaidIs404(t *testing.T) {
	w, _ := readinessRaidWorld(t)
	w.as(w.officer)
	expectError(t, w.h.do(http.MethodGet, w.path("/readiness?raid=999999"), ""), http.StatusNotFound, "not_found")
	expectError(t, w.h.do(http.MethodGet, w.path("/readiness?raid=abc"), ""), http.StatusNotFound, "not_found")
}
