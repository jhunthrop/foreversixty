package guilds

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func (w raidWorld) postAward(raidID *int64, key string) *http.Response {
	raid := ""
	if raidID != nil {
		raid = fmt.Sprintf(`,"raid_id":%d`, *raidID)
	}
	return w.h.do(http.MethodPost, w.path("/loot/awards"),
		fmt.Sprintf(`{"encounter_id":1084,"item_id":%d,"character_key":%q%s}`, onyxiaLoot[0].ItemID, key, raid))
}

// awardRaids is the award rows' raid ids for the guild, oldest first (-1 for no raid).
func (w raidWorld) awardRaids(t *testing.T) []int64 {
	t.Helper()
	rows, err := w.h.pool.Query(t.Context(),
		`select coalesce(raid_id, -1) from loot_awards where guild_id = $1 order by id`, w.gid)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var out []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			t.Fatal(err)
		}
		out = append(out, id)
	}
	return out
}

func TestAnAwardNamesADoneRaid(t *testing.T) {
	w, night, prev := reportNight(t)
	w.as(w.officer)
	for _, id := range []int64{prev.ID, night.ID} {
		if res := w.postAward(&id, officerKey); res.StatusCode != http.StatusOK {
			t.Fatalf("award for raid %d = %d: %s", id, res.StatusCode, readBody(res))
		}
	}
	got := w.awardRaids(t)
	if len(got) != 2 || got[0] != prev.ID || got[1] != night.ID {
		t.Fatalf("last week's award was overwritten: %v", got)
	}
	// Both nights are over: the latest kill night's award marks the drop (an officer handing out
	// the morning after sees what is already given), and both awards still count.
	view, err := w.h.store.Loot(t.Context(), w.gid)
	if err != nil {
		t.Fatal(err)
	}
	if item := view.Items[0]; item.ItemID != onyxiaLoot[0].ItemID || item.AwardedTo == nil {
		t.Fatalf("the latest night's award does not mark the drop: %+v", item.AwardedTo)
	}
	all, err := w.h.store.lootAwards(t.Context(), w.gid, onyxiaEncounterID)
	if err != nil {
		t.Fatal(err)
	}
	if won := awardCounts(all)[officerKey]; won != 2 {
		t.Fatalf("items won = %d, want 2", won)
	}
	// Re-awarding the same drop for the same raid replaces that raid's row only.
	if res := w.postAward(&night.ID, memberKeyB); res.StatusCode != http.StatusOK {
		t.Fatalf("re-award = %d", res.StatusCode)
	}
	if got := w.awardRaids(t); len(got) != 2 {
		t.Fatalf("re-award left %v", got)
	}
}

func TestAnAwardRefusesARaidThatIsNotDone(t *testing.T) {
	w, _ := matchedNight(t)
	planned := w.createRaid(t)
	w.as(w.officer)
	expectError(t, w.postAward(&planned.ID, officerKey), http.StatusBadRequest, "invalid")
	missing := int64(999999)
	expectError(t, w.postAward(&missing, officerKey), http.StatusBadRequest, "invalid")
}

// An award with no raid named goes to the raid being played now (its match window holds now),
// never to a previous done raid: two raids five days apart keep their own awards.
func TestAnAwardDefaultsToTheRaidBeingPlayedNow(t *testing.T) {
	w := newRaidWorld(t, true)
	now := time.Now().UTC().Truncate(time.Minute)
	first := w.raidAt(t, raidZoneOny, now.Add(-5*24*time.Hour))
	w.oneFightNight(t, "first", onyLogZone, first.StartsAt.Add(time.Hour))
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, first.ID, "first", MatchedHand); err != nil {
		t.Fatal(err)
	}
	w.as(w.officer)
	if res := w.postAward(&first.ID, officerKey); res.StatusCode != http.StatusOK {
		t.Fatalf("first award = %d: %s", res.StatusCode, readBody(res))
	}
	second := w.raidAt(t, raidZoneOny, now.Add(-30*time.Minute))
	if res := w.postAward(nil, memberKeyA); res.StatusCode != http.StatusOK {
		t.Fatalf("live award = %d: %s", res.StatusCode, readBody(res))
	}
	if got := w.awardRaids(t); len(got) != 2 || got[0] != first.ID || got[1] != second.ID {
		t.Fatalf("award raids = %v, want [%d %d]", got, first.ID, second.ID)
	}
}

// An award made after the raid's window (the morning after) goes to the latest matched raid that
// killed the boss this week, never an older one, so the CSV, the raid page and /report all show
// it (round 3 e2e, item 3).
func TestAnAwardAfterTheWindowGoesToTheLatestKillNight(t *testing.T) {
	w, night, _ := reportNight(t)
	w.as(w.officer)
	if res := w.postAward(nil, officerKey); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	if got := w.awardRaids(t); len(got) != 1 || got[0] != night.ID {
		t.Fatalf("default raid = %v, want the latest kill night %d", got, night.ID)
	}
	bosses, err := w.h.store.reportLoot(t.Context(), raidViewContext{guildID: w.gid}, night.ID)
	if err != nil {
		t.Fatal(err)
	}
	if n := countReportLoot(bosses); n != 1 {
		t.Fatalf("raid page loot = %d items, want the morning-after award", n)
	}
	loot, err := w.h.store.Loot(t.Context(), w.gid)
	if err != nil {
		t.Fatal(err)
	}
	if loot.Items[0].AwardedTo == nil {
		t.Fatal("the Loot tab shows the drop just awarded as free")
	}
}

// A logged night with no planned raid takes the award by its report, and the loot CSV dates it
// by that night instead of "No raid night recorded".
func TestAnAwardAfterALoggedNightFilesUnderItsReport(t *testing.T) {
	w := newRaidWorld(t, true)
	w.oneFightNight(t, "logged", onyLogZone, time.Now().UTC().Add(-14*time.Hour))
	w.as(w.officer)
	if res := w.postAward(nil, officerKey); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	var raidID *int64
	var reportID *string
	if err := w.h.pool.QueryRow(t.Context(),
		`select raid_id, report_id from loot_awards where guild_id = $1`, w.gid).Scan(&raidID, &reportID); err != nil {
		t.Fatal(err)
	}
	if raidID != nil || reportID == nil || *reportID != "logged" {
		t.Fatalf("award night = raid %v report %v, want the logged report", raidID, reportID)
	}
	var rows [][]string
	if err := w.h.store.streamLoot(t.Context(), w.gid, func(r []string) error { rows = append(rows, r); return nil }); err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || rows[0][1] == lootNightUnknown {
		t.Fatalf("CSV rows = %v, want the award dated by its logged night", rows)
	}
}

// A kill older than the lookback takes no award: the drop is from a past reset.
func TestAnAwardAfterAnOldKillHasNoNight(t *testing.T) {
	w := newRaidWorld(t, true)
	w.oneFightNight(t, "old", onyLogZone, time.Now().UTC().Add(-awardNightLookback-time.Hour))
	w.as(w.officer)
	if res := w.postAward(nil, officerKey); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	if got := w.awardRaids(t); len(got) != 1 || got[0] != -1 {
		t.Fatalf("default raid = %v, want none", got)
	}
}

func countReportLoot(bosses []ReportLootBoss) int {
	n := 0
	for _, b := range bosses {
		n += len(b.Awards)
	}
	return n
}

func TestAnAwardWithNoRecentRaidHasNone(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)
	if res := w.postAward(nil, officerKey); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	if got := w.awardRaids(t); len(got) != 1 || got[0] != -1 {
		t.Fatalf("raid = %v, want none", got)
	}
	if res := w.postAward(nil, memberKeyA); res.StatusCode != http.StatusOK {
		t.Fatalf("second award = %d", res.StatusCode)
	}
	if got := w.awardRaids(t); len(got) != 1 {
		t.Fatalf("a second raid-less award for one drop left %v", got)
	}
}
