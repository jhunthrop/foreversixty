package guilds

import (
	"net/http"
	"strconv"
	"testing"
	"time"
)

const killOnyxiaMS = int64(5 * 60 * 1000)

// reportNight is the attendance night matched by hand, plus the same zone's raid a week
// earlier with its own log, so the night's report has something to compare against.
func reportNight(t *testing.T) (w raidWorld, night, prev Raid) {
	t.Helper()
	w, night = matchedNight(t)
	prevStart := night.StartsAt.Add(-7 * 24 * time.Hour)
	prev = w.raidAt(t, raidZoneOny, prevStart)
	seedNight(t, w.h.pool, w.gid, "prev", onyLogZone, prevStart.Add(time.Hour), []seedFight{{
		encounterID: 1084, kill: true, duration: 6 * time.Minute, players: []seedPlayer{
			{key: officerKey, name: "Thornhide", class: "druid", role: "dps"},
			{key: member2Key, name: "Mirelle", class: "priest", role: "healer"},
		}}})
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, prev.ID, "prev", MatchedHand); err != nil {
		t.Fatal(err)
	}
	w.exec(t, `update fight_metrics set metric_dps = 150 where report_id = 'night' and kill and player_key = '`+officerKey+`'`)
	w.exec(t, `update fight_metrics set role = 'healer', metric_hps = 80 where report_id = 'night' and kill and player_key = '`+member2Key+`'`)
	w.exec(t, `update fight_metrics set metric_hps = 50 where report_id = 'prev' and player_key = '`+member2Key+`'`)
	seedExport(t, w.h.pool, w.officer, officerKey, "us", "hardcore", "Thornhide")
	seedExport(t, w.h.pool, w.member, memberKeyB, "us", "hardcore", "Sunderfel")
	return w, night, prev
}

func (w raidWorld) award(t *testing.T, raidID int64, key string) {
	t.Helper()
	w.awardItem(t, raidID, key, 0)
}

// awardItem awards onyxiaLoot[item] to key for raidID.
func (w raidWorld) awardItem(t *testing.T, raidID int64, key string, item int) {
	t.Helper()
	if _, err := w.h.store.CreateLootAward(t.Context(), w.gid, 1084, onyxiaLoot[item].ItemID, key, w.officer, &raidID); err != nil {
		t.Fatal(err)
	}
}

func TestTheReportFollowsTheLog(t *testing.T) {
	w, night, prev := reportNight(t)
	w.award(t, night.ID, officerKey)
	w.as(w.officer)
	rep := w.getRaid(t, night.ID).Report
	if rep == nil {
		t.Fatal("report is null on a matched raid")
	}
	if len(rep.KillOrder) != 1 {
		t.Fatalf("kill order = %+v", rep.KillOrder)
	}
	k := rep.KillOrder[0]
	if k.EncounterID != 1084 || !k.Killed || k.KillMS == nil || *k.KillMS != killOnyxiaMS || k.WipesBefore != 3 ||
		k.FightIndex != 4 || k.Deaths != 3 {
		t.Fatalf("kill = %+v", k)
	}
	if u := rep.Unnamed; u == nil || u.Pulls != 1 || u.Kills != 0 || u.Wipes != 1 {
		t.Fatalf("unnamed = %+v", rep.Unnamed)
	}
	if d := rep.Deaths; len(d) != 2 || d[0].CharacterKey != drifterKey || d[0].OnWipes != 6 || d[0].OnKills != 2 ||
		d[1].CharacterKey != officerKey || d[1].OnWipes != 3 || d[1].OnKills != 1 {
		t.Fatalf("deaths = %+v", rep.Deaths)
	}
	if c := rep.Consumables; c.LogHasBuffs || len(c.Rows) != 2 || c.Rows[0].Before != "unknown" ||
		c.Rows[0].Flask != nil || c.Rows[0].Potions != nil {
		t.Fatalf("consumables = %+v", c)
	}
	if len(rep.Loot) != 1 || rep.Loot[0].Boss != "Onyxia" || len(rep.Loot[0].Awards) != 1 ||
		rep.Loot[0].Awards[0].ItemName != onyxiaLoot[0].Name || rep.Loot[0].Awards[0].Name != "Thornhide" {
		t.Fatalf("loot = %+v", rep.Loot)
	}
	m := rep.Moves
	if m.Previous == nil || m.Previous.ID != prev.ID || len(m.Rows) != 2 {
		t.Fatalf("moves = %+v", m)
	}
	if r := m.Rows[0]; r.CharacterKey != officerKey || r.Metric != "dps" || r.Before != 100 || r.After != 150 {
		t.Fatalf("first move = %+v", r)
	}
	if r := m.Rows[1]; r.CharacterKey != member2Key || r.Metric != "hps" || r.Before != 50 || r.After != 80 {
		t.Fatalf("second move = %+v", r)
	}
	if rep.PostedAt != nil {
		t.Fatalf("posted_at = %v", rep.PostedAt)
	}
}

func TestTheFirstRaidOfAZoneHasNoMoves(t *testing.T) {
	w, night := matchedNight(t)
	w.as(w.officer)
	rep := w.getRaid(t, night.ID).Report
	if rep.Moves.Previous != nil || rep.Moves.Rows == nil || len(rep.Moves.Rows) != 0 {
		t.Fatalf("moves = %+v", rep.Moves)
	}
	if rep.Loot == nil || len(rep.Loot) != 0 {
		t.Fatalf("loot = %+v", rep.Loot)
	}
}

func TestAMemberSeesTheReportFilteredToTheirCharacters(t *testing.T) {
	w, night, _ := reportNight(t)
	w.award(t, night.ID, officerKey)
	w.as(w.member)
	rep := w.getRaid(t, night.ID).Report
	if len(rep.KillOrder) != 1 || rep.Unnamed == nil || len(rep.Loot) != 1 {
		t.Fatalf("member lost the shared blocks: %+v", rep)
	}
	if len(rep.Deaths) != 0 || len(rep.Moves.Rows) != 0 {
		t.Fatalf("member sees others' rows: %+v %+v", rep.Deaths, rep.Moves.Rows)
	}
	if rows := rep.Consumables.Rows; len(rows) != 1 || rows[0].CharacterKey != memberKeyB {
		t.Fatalf("member consumables = %+v", rows)
	}
	if rep.Moves.Previous == nil {
		t.Fatal("member lost the previous raid")
	}
	w.as(w.member2)
	if rows := w.getRaid(t, night.ID).Report.Moves.Rows; len(rows) != 1 || rows[0].CharacterKey != member2Key {
		t.Fatalf("member2 moves = %+v", rows)
	}
}

func TestAnUnmatchedRaidHasNullReport(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	data := rawData(t, w.h.do(http.MethodGet, w.path("/raids/%d", raid.ID), ""))
	if got := string(data["report"]); got != "null" {
		t.Fatalf("report = %s", got)
	}
}

func postWorld(t *testing.T) (raidWorld, Raid, *recordingRaidEvents) {
	t.Helper()
	w, night, _ := reportNight(t)
	w.award(t, night.ID, officerKey)
	w.awardItem(t, night.ID, memberKeyB, 1)
	events := &recordingRaidEvents{}
	w.h.svc.RaidEvents = events
	w.h.svc.PublicBaseURL = testSiteBase
	return w, night, events
}

func TestPostingTheReportPostsOnce(t *testing.T) {
	w, night, events := postWorld(t)
	w.as(w.officer)
	res := w.h.do(http.MethodPost, w.path("/raids/%d/report/post", night.ID), "")
	var out struct {
		PostedAt *time.Time `json:"posted_at"`
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("first post = %d: %s", res.StatusCode, readBody(res))
	}
	w.h.data(res, &out)
	if out.PostedAt == nil || len(events.reports) != 1 {
		t.Fatalf("posted_at %v, reports %d", out.PostedAt, len(events.reports))
	}
	e := events.reports[0]
	if e.Title != "Night" || len(e.Kills) != 1 || e.Kills[0].Duration != 5*time.Minute ||
		e.Wipes != 3 || e.LootItems != 2 || e.Signed != 4 || e.Came != 3 {
		t.Fatalf("embed input = %+v", e)
	}
	if want := testSiteBase + "/guild/us/hardcore/olympus/raids/" + strconv.FormatInt(night.ID, 10); e.RaidURL != want {
		t.Fatalf("url = %s, want %s", e.RaidURL, want)
	}
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/report/post", night.ID), ""), http.StatusConflict, "already_posted")
	if len(events.reports) != 1 {
		t.Fatalf("second press posted again: %d", len(events.reports))
	}
	if rep := w.getRaid(t, night.ID).Report; rep.PostedAt == nil {
		t.Fatal("view lost posted_at")
	}
}

func TestPostingRefusesARaidThatIsNotDone(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.h.svc.RaidEvents = &recordingRaidEvents{}
	w.as(w.officer)
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/report/post", raid.ID), ""), http.StatusConflict, "conflict")
}

func TestPostingIsOfficerOnly(t *testing.T) {
	w, night, events := postWorld(t)
	w.as(w.member)
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/report/post", night.ID), ""), http.StatusForbidden, "forbidden")
	if len(events.reports) != 0 {
		t.Fatal("a member posted")
	}
}

func TestAPostThatMissesTheChannelCanBeRetried(t *testing.T) {
	w, night, events := postWorld(t)
	events.notPosted = true
	w.as(w.officer)
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/report/post", night.ID), ""), http.StatusBadGateway, "post_failed")
	if rep := w.getRaid(t, night.ID).Report; rep.PostedAt != nil {
		t.Fatalf("posted_at kept after a failed post: %v", rep.PostedAt)
	}
	events.notPosted = false
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/report/post", night.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("retry = %d: %s", res.StatusCode, readBody(res))
	}
}
