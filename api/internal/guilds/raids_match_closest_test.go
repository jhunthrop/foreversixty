package guilds

import (
	"testing"
	"time"
)

// Closest marks the nearest log the officer can pick: a log matched to another raid is listed
// but never preselected.
func TestClosestSkipsALogMatchedElsewhere(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	elsewhere := w.raidAt(t, raidZoneOny, start.Add(-24*time.Hour))
	w.oneFightNight(t, "near", onyLogZone, start.Add(10*time.Minute))
	w.oneFightNight(t, "far", onyLogZone, start.Add(2*time.Hour))
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, elsewhere.ID, "near", MatchedHand); err != nil {
		t.Fatal(err)
	}
	res, err := w.h.store.MatchCandidates(t.Context(), w.gid, raid.ID, time.Now())
	if err != nil || len(res.Rows) != 2 {
		t.Fatalf("candidates = %+v %v, want two", res, err)
	}
	if res.Rows[0].ReportID != "near" || res.Rows[0].Closest || res.Rows[1].ReportID != "far" || !res.Rows[1].Closest {
		t.Fatalf("rows = %+v, want far marked closest", res.Rows)
	}
}
