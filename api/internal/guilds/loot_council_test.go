package guilds

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// secondOfficerKey is a second officer of the raid world, so votes are "of 2".
const secondOfficerKey = "us/hardcore/ironbrand"

func (w raidWorld) councilDrop(t *testing.T, night string) CouncilDropView {
	t.Helper()
	var view CouncilDropView
	w.h.data(w.h.do(http.MethodGet, w.path("/loot/drops/%d/candidates?night=%s", councilItem, night), ""), &view)
	return view
}

func (w raidWorld) candidateTally(t *testing.T, key string, viewer int64) LootVotes {
	t.Helper()
	night, err := w.h.store.councilNight(t.Context(), w.gid, nightKeyLive, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	got, err := w.h.store.candidateVotes(t.Context(), w.gid, night, councilItem, key, viewer, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	return got
}

func TestOfficerVotesAreOnePerCandidateChangeableAndRemovable(t *testing.T) {
	w := newRaidWorld(t, true)
	second := w.seedRaider(t, "officer2@example.com", "officer", raidChar{secondOfficerKey, "Ironbrand", "warrior", true})
	w.as(w.officer)
	for _, vote := range []string{`"yes"`, `"no"`, `"yes"`} {
		if res := w.putVote("", memberKeyA, vote); res.StatusCode != http.StatusOK {
			t.Fatalf("vote %s = %d: %s", vote, res.StatusCode, readBody(res))
		}
	}
	w.as(second)
	if res := w.putVote("", memberKeyA, `"no"`); res.StatusCode != http.StatusOK {
		t.Fatalf("second vote = %d", res.StatusCode)
	}
	got := w.candidateTally(t, memberKeyA, w.officer)
	if got.Yes != 1 || got.No != 1 || got.Voters != 2 || got.Mine == nil || *got.Mine != voteYes {
		t.Fatalf("tally = %+v, want 1 yes, 1 no, of 2, mine yes", got)
	}
	if res := w.putVote("", memberKeyA, `null`); res.StatusCode != http.StatusOK {
		t.Fatalf("remove = %d", res.StatusCode)
	}
	if got := w.candidateTally(t, memberKeyA, second); got.No != 0 || got.Yes != 1 || got.Mine != nil {
		t.Fatalf("after removing = %+v", got)
	}
}

func TestAnAwardKeepsItsDecisionRecordAndClosesVoting(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)
	if res := w.putVote("", memberKeyA, `"yes"`); res.StatusCode != http.StatusOK {
		t.Fatalf("vote = %d", res.StatusCode)
	}
	award := func(body string) *http.Response {
		return w.h.do(http.MethodPost, w.path("/loot/awards"), fmt.Sprintf(
			`{"encounter_id":1084,"item_id":%d,"character_key":%q%s}`, councilItem, memberKeyA, body))
	}
	expectError(t, award(`,"ms_os":"both"`), http.StatusBadRequest, "invalid")
	long := fmt.Sprintf("%0201d", 0)
	expectError(t, award(fmt.Sprintf(`,"reason":%q`, long)), http.StatusBadRequest, "invalid")
	expectError(t, award(`,"night":"raid:424242"`), http.StatusBadRequest, "invalid")
	if res := award(`,"ms_os":"os","reason":"  Main spec passed  ","night":"live"`); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	var msOs, reason string
	var yes, no, officers int
	if err := w.h.pool.QueryRow(t.Context(), `
		select ms_os, reason, votes_yes, votes_no, votes_officers from loot_awards where guild_id = $1`, w.gid).
		Scan(&msOs, &reason, &yes, &no, &officers); err != nil {
		t.Fatal(err)
	}
	if msOs != msOsOS || reason != "Main spec passed" || yes != 1 || no != 0 || officers != 1 {
		t.Fatalf("record = %s %q %d/%d of %d", msOs, reason, yes, no, officers)
	}
	expectError(t, w.putVote("", memberKeyA, `"no"`), http.StatusConflict, "voting_closed")
	w.as(w.member)
	expectError(t, w.putInterest(memberKeyA, `"ms"`), http.StatusConflict, "voting_closed")

	w.as(w.officer)
	view := w.councilDrop(t, nightKeyLive)
	if !view.Closed || view.Award == nil || view.Award.Reason == nil || *view.Award.Reason != "Main spec passed" ||
		view.Award.VotesYes == nil || *view.Award.VotesYes != 1 || view.Award.Name != "Kraggor" {
		t.Fatalf("drop view = %+v award %+v", view, view.Award)
	}
}

func TestNightsGroupTheCouncilsDropsPerNight(t *testing.T) {
	w := newRaidWorld(t, true)
	now := time.Now().UTC().Truncate(time.Minute)
	past := w.raidAt(t, raidZoneOny, now.Add(-3*24*time.Hour))
	w.oneFightNight(t, "past", onyLogZone, past.StartsAt.Add(time.Hour))
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, past.ID, "past", MatchedHand); err != nil {
		t.Fatal(err)
	}
	w.as(w.officer)
	pastKey := nightKeyRaid + fmt.Sprint(past.ID)
	if res := w.putVote(pastKey, memberKeyA, `"yes"`); res.StatusCode != http.StatusOK {
		t.Fatalf("vote = %d: %s", res.StatusCode, readBody(res))
	}
	if res := w.h.do(http.MethodPost, w.path("/loot/awards"), fmt.Sprintf(
		`{"encounter_id":1084,"item_id":%d,"character_key":%q,"night":%q,"ms_os":"ms"}`, onyxiaLoot[1].ItemID, memberKeyB, pastKey)); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	var view LootNightsView
	w.h.data(w.h.do(http.MethodGet, w.path("/loot/nights"), ""), &view)
	if view.Officers != 1 || len(view.Nights) != 1 {
		t.Fatalf("nights = %+v", view)
	}
	n := view.Nights[0]
	// The past raid is the latest kill night this week, so it is the current night too.
	if n.Night != pastKey || n.RaidID == nil || !n.Current || n.Date != past.StartsAt.UTC().Format(exportDateLayout) || len(n.Drops) != 2 {
		t.Fatalf("night = %+v", n)
	}
	states := map[int]string{}
	for _, d := range n.Drops {
		states[d.ItemID] = d.State
	}
	if states[councilItem] != dropOpen || states[onyxiaLoot[1].ItemID] != dropAwarded {
		t.Fatalf("drop states = %v", states)
	}
	if len(n.Decisions) != 1 || n.Decisions[0].Name != "Sunderfel" || n.Decisions[0].MsOs == nil {
		t.Fatalf("decisions = %+v", n.Decisions)
	}
	expectError(t, w.h.do(http.MethodGet, w.path("/loot/nights?limit=99"), ""), http.StatusBadRequest, "invalid")
}

// A vote cast with no night yet joins the raid when the raid learns its log, like an award.
func TestLiveVotesJoinTheRaidOnMatch(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)
	if res := w.putVote("", memberKeyA, `"yes"`); res.StatusCode != http.StatusOK {
		t.Fatalf("vote = %d: %s", res.StatusCode, readBody(res))
	}
	if res := w.putInterest(memberKeyA, `"ms"`); res.StatusCode != http.StatusOK {
		t.Fatalf("flag = %d: %s", res.StatusCode, readBody(res))
	}
	now := time.Now().UTC().Truncate(time.Minute)
	raid := w.raidAt(t, raidZoneOny, now.Add(-4*time.Hour))
	w.oneFightNight(t, "tonight", onyLogZone, now.Add(-3*time.Hour))
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "tonight", MatchedHand); err != nil {
		t.Fatal(err)
	}
	for _, table := range []string{"loot_votes", "loot_interest"} {
		var raidID *int64
		if err := w.h.pool.QueryRow(t.Context(), `select raid_id from `+table+` where guild_id = $1`, w.gid).Scan(&raidID); err != nil {
			t.Fatal(err)
		}
		if raidID == nil || *raidID != raid.ID {
			t.Fatalf("%s raid = %v, want %d", table, raidID, raid.ID)
		}
	}
	view := w.councilDrop(t, fmt.Sprintf("raid:%d", raid.ID))
	for _, c := range view.Candidates {
		if c.CharacterKey == memberKeyA && c.Votes.Yes != 1 {
			t.Fatalf("moved vote not read: %+v", c.Votes)
		}
	}
}

func TestCandidateMsOsRule(t *testing.T) {
	declared, src := candidateMsOs(LootCandidate{Tier: 1}, lootKindGear, msOsMS)
	if *declared != msOsMS || *src != msOsDeclared {
		t.Fatalf("declared = %s %s", *declared, *src)
	}
	if got, _ := candidateMsOs(LootCandidate{Tier: 0}, lootKindGear, ""); *got != msOsMS {
		t.Fatalf("tier 0 = %s, want ms", *got)
	}
	if got, _ := candidateMsOs(LootCandidate{Tier: 1}, lootKindGear, ""); *got != msOsOS {
		t.Fatalf("tier 1 = %s, want os", *got)
	}
	if got, _ := candidateMsOs(LootCandidate{Tier: 1}, lootKindQuest, ""); got != nil {
		t.Fatalf("quest = %v, want none", *got)
	}
}

func TestNightKeysRoundTrip(t *testing.T) {
	for _, key := range []string{"raid:12", "report:abcDEF", nightKeyLive} {
		n, ok := parseNightKey(key)
		if !ok || n.key() != key {
			t.Fatalf("%q -> %v %q", key, ok, n.key())
		}
	}
	for _, bad := range []string{"", "raid:", "raid:x", "raid:-1", "report:", "tonight"} {
		if _, ok := parseNightKey(bad); ok {
			t.Fatalf("%q parsed", bad)
		}
	}
}
