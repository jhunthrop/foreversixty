package guilds

import (
	"fmt"
	"net/http"
	"testing"
)

// councilItem is Sapphiron Drape, a slotted Onyxia drop.
const councilItem = 17078

func (w raidWorld) putVote(night, key string, vote string) *http.Response {
	return w.h.do(http.MethodPut, w.path("/loot/votes"),
		fmt.Sprintf(`{"night":%q,"item_id":%d,"character_key":%q,"vote":%s}`, night, councilItem, key, vote))
}

func (w raidWorld) putInterest(key, msOs string) *http.Response {
	return w.h.do(http.MethodPut, w.path("/loot/interest"),
		fmt.Sprintf(`{"item_id":%d,"character_key":%q,"ms_os":%s}`, councilItem, key, msOs))
}

func TestCouncilReadsAreOfficerOnly(t *testing.T) {
	w := newRaidWorld(t, true)
	reads := []string{w.path("/loot/nights"), w.path("/loot/drops/%d/candidates", councilItem)}
	w.as(w.member)
	expectError(t, w.h.do(http.MethodGet, reads[1], ""), http.StatusForbidden, "forbidden")
	expectError(t, w.putVote("", memberKeyA, `"yes"`), http.StatusForbidden, "forbidden")
	w.as(w.officer)
	for _, path := range reads {
		if res := w.h.do(http.MethodGet, path, ""); res.StatusCode != http.StatusOK {
			t.Fatalf("officer GET %s = %d: %s", path, res.StatusCode, readBody(res))
		}
	}
}

func TestCouncilVotesNeedThePlan(t *testing.T) {
	w := newRaidWorld(t, false)
	w.as(w.officer)
	expectError(t, w.putVote("", memberKeyA, `"yes"`), http.StatusForbidden, "guild_plan_required")
	// Reading the council is history, never locked.
	if res := w.h.do(http.MethodGet, w.path("/loot/nights"), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("nights without a plan = %d", res.StatusCode)
	}
}

func TestCouncilVoteBodyIsChecked(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)
	expectError(t, w.putVote("", memberKeyA, `"maybe"`), http.StatusBadRequest, "invalid")
	expectError(t, w.putVote("raid:999999", memberKeyA, `"yes"`), http.StatusBadRequest, "invalid")
	expectError(t, w.putVote("nonsense", memberKeyA, `"yes"`), http.StatusBadRequest, "invalid")
	expectError(t, w.putVote("", "us/hardcore/stranger", `"yes"`), http.StatusBadRequest, "invalid")
	expectError(t, w.h.do(http.MethodGet, w.path("/loot/drops/1/candidates"), ""), http.StatusNotFound, "not_found")
}

func TestAMemberFlagsOnlyTheirOwnCharacters(t *testing.T) {
	w := newRaidWorld(t, false) // a member's own flag needs no plan
	w.as(w.member)
	if res := w.putInterest(memberKeyA, `"os"`); res.StatusCode != http.StatusOK {
		t.Fatalf("own flag = %d: %s", res.StatusCode, readBody(res))
	}
	expectError(t, w.putInterest(member2Key, `"ms"`), http.StatusForbidden, "forbidden")
	expectError(t, w.putInterest(memberKeyA, `"main"`), http.StatusBadRequest, "invalid")
	w.as(w.member2)
	var mine LootInterestView
	w.h.data(w.h.do(http.MethodGet, w.path("/loot/interest"), ""), &mine)
	if len(mine.Flags) != 0 {
		t.Fatalf("member2 reads another's flags: %+v", mine.Flags)
	}
	w.as(w.member)
	w.h.data(w.h.do(http.MethodGet, w.path("/loot/interest"), ""), &mine)
	if len(mine.Flags) != 1 || mine.Flags[0].MsOs != msOsOS || mine.Flags[0].SetByOfficer || mine.Night != nightKeyLive {
		t.Fatalf("own flags = %+v", mine)
	}
}

func TestAnOfficerFlagsOnARaidersBehalfWithThePlan(t *testing.T) {
	w := newRaidWorld(t, false)
	w.as(w.officer)
	expectError(t, w.putInterest(memberKeyA, `"ms"`), http.StatusForbidden, "guild_plan_required")
	w.grantPlan(t)
	if res := w.putInterest(memberKeyA, `"ms"`); res.StatusCode != http.StatusOK {
		t.Fatalf("officer flag = %d: %s", res.StatusCode, readBody(res))
	}
	w.as(w.member)
	var mine LootInterestView
	w.h.data(w.h.do(http.MethodGet, w.path("/loot/interest"), ""), &mine)
	if len(mine.Flags) != 1 || !mine.Flags[0].SetByOfficer {
		t.Fatalf("flag set by an officer = %+v", mine.Flags)
	}
	if res := w.putInterest(memberKeyA, `null`); res.StatusCode != http.StatusOK {
		t.Fatalf("remove = %d", res.StatusCode)
	}
	w.h.data(w.h.do(http.MethodGet, w.path("/loot/interest"), ""), &mine)
	if len(mine.Flags) != 0 {
		t.Fatalf("flag not removed: %+v", mine.Flags)
	}
}

// A member reads the nights and the decision record, never the officers' vote counts.
func TestAMemberReadsTheNightsWithoutVotes(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)
	if res := w.putVote("", memberKeyA, `"yes"`); res.StatusCode != http.StatusOK {
		t.Fatalf("vote = %d: %s", res.StatusCode, readBody(res))
	}
	if res := w.h.do(http.MethodPost, w.path("/loot/awards"), fmt.Sprintf(
		`{"encounter_id":1084,"item_id":%d,"character_key":%q,"ms_os":"ms"}`, councilItem, memberKeyA)); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	var officerView, memberView LootNightsView
	w.h.data(w.h.do(http.MethodGet, w.path("/loot/nights"), ""), &officerView)
	if d := officerView.Nights[0].Decisions; len(d) != 1 || d[0].VotesYes == nil || *d[0].VotesYes != 1 {
		t.Fatalf("officer decisions = %+v", d)
	}
	w.as(w.member)
	w.h.data(w.h.do(http.MethodGet, w.path("/loot/nights"), ""), &memberView)
	n := memberView.Nights[0]
	if len(n.Decisions) != 1 || n.Decisions[0].VotesYes != nil || n.Decisions[0].VotesOfficers != nil ||
		n.Decisions[0].MsOs == nil || len(n.Drops) != 1 || n.Drops[0].Award == nil || n.Drops[0].Award.VotesNo != nil {
		t.Fatalf("member night = %+v", n)
	}
	w.as(seedUser(t, w.h.pool, "outsider@example.com"))
	expectError(t, w.h.do(http.MethodGet, w.path("/loot/nights"), ""), http.StatusForbidden, "forbidden")
}
