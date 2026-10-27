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

// A done raid whose window has passed takes no default award.
func TestAnAwardAfterTheWindowHasNoRaid(t *testing.T) {
	w, _, _ := reportNight(t)
	w.as(w.officer)
	if res := w.postAward(nil, officerKey); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	if got := w.awardRaids(t); len(got) != 1 || got[0] != -1 {
		t.Fatalf("default raid = %v, want none (the night's window has ended)", got)
	}
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
