package guilds

import (
	"errors"
	"testing"
	"time"
)

// A member's attributed log carries its guild in matched_guild_id, never guild_id; the matcher
// reads it there.
func TestAutoMatchReadsTheMatchedGuildWhenGuildIDIsNull(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.oneFightNight(t, "log1", onyLogZone, start.Add(30*time.Minute))
	w.exec(t, `update reports set matched_guild_id = guild_id, guild_id = null where id = 'log1'`)
	got, ok, err := w.h.store.AutoMatch(t.Context(), "log1")
	if err != nil || !ok || got.ID != raid.ID {
		t.Fatalf("auto match = %+v %v %v, want raid %d", got, ok, err, raid.ID)
	}
}

// guild_id wins over matched_guild_id: the officer's attachment decides.
func TestTheOfficersGuildBeatsTheMatchedGuild(t *testing.T) {
	w := newRaidWorld(t, false)
	other := seedGuild(t, w.h.pool, "Elsewhere")
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.oneFightNight(t, "log1", onyLogZone, start.Add(30*time.Minute))
	if _, err := w.h.pool.Exec(t.Context(), `update reports set matched_guild_id = $1 where id = 'log1'`, other); err != nil {
		t.Fatal(err)
	}
	got, ok, err := w.h.store.AutoMatch(t.Context(), "log1")
	if err != nil || !ok || got.ID != raid.ID {
		t.Fatalf("auto match = %+v %v %v, want guild_id's raid %d", got, ok, err, raid.ID)
	}
	// And the other way round: a log the officer put in another guild is not this guild's.
	raid2 := w.raidAt(t, raidZoneOny, start.Add(24*time.Hour))
	w.oneFightNight(t, "log2", onyLogZone, start.Add(24*time.Hour+30*time.Minute))
	if _, err := w.h.pool.Exec(t.Context(),
		`update reports set guild_id = $1, matched_guild_id = $2 where id = 'log2'`, other, w.gid); err != nil {
		t.Fatal(err)
	}
	if _, ok, err := w.h.store.AutoMatch(t.Context(), "log2"); err != nil || ok {
		t.Fatalf("auto match = %v %v, want no match for another guild's log", ok, err)
	}
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid2.ID, "log2", MatchedHand); !errors.Is(err, ErrNotFound) {
		t.Fatalf("hand match of another guild's log = %v, want ErrNotFound", err)
	}
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid2.ID, "log1", MatchedHand); !errors.Is(err, ErrReportMatched) {
		t.Fatalf("hand match of a matched log = %v, want ErrReportMatched", err)
	}
}

func TestMatchCandidatesListAMatchedGuildLog(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.oneFightNight(t, "log1", onyLogZone, start.Add(30*time.Minute))
	w.exec(t, `update reports set matched_guild_id = guild_id, guild_id = null where id = 'log1'`)
	res, err := w.h.store.MatchCandidates(t.Context(), w.gid, raid.ID, time.Now())
	if err != nil || len(res.Rows) != 1 || res.Rows[0].ReportID != "log1" {
		t.Fatalf("candidates = %+v %v, want log1", res, err)
	}
}
