package guilds

import (
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

func TestAppContractRaidBadge(t *testing.T) {
	w := newRaidWorld(t, true)
	w.createRaid(t)
	w.as(w.member)
	res := w.h.do(http.MethodGet, w.path("/raid-badge"), "")
	assertGuildCache(t, res, privateCache)
	contract.AssertFixture(t, "guild.raid-badge", guildBody(t, res, http.StatusOK), contract.Options{})
	w.as(w.officer)
	contract.AssertFixture(t, "guild.raid-badge.officer",
		guildBody(t, w.h.do(http.MethodGet, w.path("/raid-badge"), ""), http.StatusOK), contract.Options{})
}

// badgeOf reads the Raid tab's badge as uid.
func (w raidWorld) badgeOf(t *testing.T, uid int64) RaidBadge {
	t.Helper()
	w.as(uid)
	res := w.h.do(http.MethodGet, w.path("/raid-badge"), "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("raid badge = %d: %s", res.StatusCode, readBody(res))
	}
	var b RaidBadge
	w.h.data(res, &b)
	return b
}

// A member's badge is 1 while their own next raid is unanswered (spec 2026-11-08-app-flows F3).
func TestMemberRaidBadgeIsTheirOwnAnswer(t *testing.T) {
	w := newRaidWorld(t, true)
	if b := w.badgeOf(t, w.member); b.Count != 0 || b.RaidID != nil {
		t.Fatalf("no raid: %+v", b)
	}
	raid := w.createRaid(t)
	b := w.badgeOf(t, w.member)
	if b.Count != 1 || !b.YouUnanswered || b.RaidID == nil || *b.RaidID != raid.ID || b.Officer {
		t.Fatalf("unanswered: %+v", b)
	}
	if b.Unanswered != 0 || b.Applicants != 0 {
		t.Fatalf("a member sees the officer counts: %+v", b)
	}
	w.as(w.member)
	if res := w.putSignup(raid.ID, memberKeyA, `{"state":"out"}`); res.StatusCode != http.StatusOK {
		t.Fatalf("sign-up = %d", res.StatusCode)
	}
	if b := w.badgeOf(t, w.member); b.Count != 0 || b.YouUnanswered {
		t.Fatalf("answered: %+v", b)
	}
}

// An officer's badge is the open count: the next raid's unanswered accounts plus the pending
// applicants, the same number as Needs you.
func TestOfficerRaidBadgeIsTheOpenCount(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	applicant := seedUser(t, w.h.pool, "applicant@example.com")
	if _, err := w.h.pool.Exec(t.Context(),
		`insert into guild_applications (guild_id, user_id, character_key) values ($1, $2, 'us/hardcore/mirage')`,
		w.gid, applicant); err != nil {
		t.Fatal(err)
	}
	b := w.badgeOf(t, w.officer)
	if !b.Officer || b.Unanswered != 3 || b.Applicants != 1 || b.Count != 4 || !b.YouUnanswered {
		t.Fatalf("officer: %+v", b)
	}
	w.as(w.member)
	w.putSignup(raid.ID, memberKeyA, `{"state":"in"}`)
	if b := w.badgeOf(t, w.officer); b.Unanswered != 2 || b.Count != 3 {
		t.Fatalf("after a sign-up: %+v", b)
	}
}

// Someone who is not a verified member has no badge to read.
func TestRaidBadgeIsForMembers(t *testing.T) {
	w := newRaidWorld(t, true)
	stranger := seedUser(t, w.h.pool, "stranger@example.com")
	w.as(stranger)
	if res := w.h.do(http.MethodGet, w.path("/raid-badge"), ""); res.StatusCode != http.StatusNotFound {
		t.Fatalf("stranger = %d", res.StatusCode)
	}
}
