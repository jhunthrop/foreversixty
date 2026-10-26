package guilds

import (
	"fmt"
	"net/http"
	"testing"
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

func TestAnAwardDefaultsToTheLatestDoneRaidOfTheWeek(t *testing.T) {
	w, night, _ := reportNight(t)
	w.as(w.officer)
	if res := w.postAward(nil, officerKey); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	if got := w.awardRaids(t); len(got) != 1 || got[0] != night.ID {
		t.Fatalf("default raid = %v, want %d (the previous raid is eight days back)", got, night.ID)
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
