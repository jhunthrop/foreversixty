// api/internal/guilds/claim_facts_test.go
package guilds

import (
	"context"
	"fmt"
	"testing"
)

// seedRanked seeds an account with one character of the given rank in guildID, with an
// export so the character has a name, and returns the account id.
func seedRanked(t *testing.T, h *httpHarness, guildID int64, name, rank string, verified bool) int64 {
	t.Helper()
	uid := seedUser(t, h.pool, fmt.Sprintf("%s-%d@example.com", name, guildID))
	key := "us/hardcore/" + name
	seedCharacter(t, h.pool, guildID, uid, key, rank, verified)
	seedExport(t, h.pool, uid, key, "us", "hardcore", name)
	syncMembership(t, h.pool, guildID, uid)
	return uid
}

func setGuild(t *testing.T, h *httpHarness, sql string, args ...any) {
	t.Helper()
	if _, err := h.pool.Exec(context.Background(), sql, args...); err != nil {
		t.Fatal(err)
	}
}

func claimViewOf(t *testing.T, h *httpHarness, viewer, guildID int64) (claimView, bool) {
	t.Helper()
	claims, verify, err := h.store.ClaimAndVerifyViews(context.Background(), viewer, []int64{guildID})
	if err != nil {
		t.Fatal(err)
	}
	c := claims[guildID]
	return claimView{State: c.State, Yours: c.Yours, PendingYours: c.PendingYours, HolderName: c.HolderName},
		verify[guildID].HasApprover
}

type claimView struct {
	State        string
	Yours        bool
	PendingYours bool
	HolderName   *string
}

func TestGuildStanding_ClaimStates(t *testing.T) {
	h := newHTTPHarness(t)
	viewer := seedUser(t, h.pool, "claim-viewer@example.com")
	other := seedUser(t, h.pool, "claim-other@example.com")

	unclaimed := seedGuild(t, h.pool, "Unclaimed")
	pending := seedGuild(t, h.pool, "Pending")
	setGuild(t, h, `update guilds set claim_pending_by = $2, claim_requested_at = now() where id = $1`, pending, viewer)
	claimed := seedGuild(t, h.pool, "Claimed")
	setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, claimed, viewer)
	contested := seedGuild(t, h.pool, "Contested")
	setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now(), claim_contested_at = now(),
		claim_contested_by = $3 where id = $1`, contested, other, viewer)

	cases := []struct {
		name string
		id   int64
		want claimView
	}{
		{"unclaimed", unclaimed, claimView{State: ClaimStateUnclaimed}},
		{"pending is the viewer's own", pending, claimView{State: ClaimStatePending, PendingYours: true}},
		{"claimed by the viewer", claimed, claimView{State: ClaimStateClaimed, Yours: true}},
		{"contested", contested, claimView{State: ClaimStateContested}},
	}
	for _, c := range cases {
		got, _ := claimViewOf(t, h, viewer, c.id)
		if got.State != c.want.State || got.Yours != c.want.Yours || got.PendingYours != c.want.PendingYours {
			t.Errorf("%s: got %+v, want %+v", c.name, got, c.want)
		}
	}
	if got, _ := claimViewOf(t, h, other, pending); got.PendingYours || got.State != ClaimStatePending {
		t.Errorf("another account sees pending_yours = %v state %q", got.PendingYours, got.State)
	}
}

func TestGuildStanding_HolderNameOnlyForVerifiedOfficer(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	holder := seedRanked(t, h, gid, "Holder", "leader", true)
	setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, gid, holder)

	officer := seedRanked(t, h, gid, "Officer", "officer", true)
	member := seedRanked(t, h, gid, "Member", "member", true)
	unverifiedOfficer := seedRanked(t, h, gid, "Shady", "officer", false)
	unverifiedMember := seedRanked(t, h, gid, "Newbie", "member", false)

	got, _ := claimViewOf(t, h, officer, gid)
	if got.HolderName == nil || *got.HolderName != "Holder" {
		t.Fatalf("verified officer sees holder_name = %v, want Holder", got.HolderName)
	}
	for name, viewer := range map[string]int64{
		"verified member": member, "unverified officer": unverifiedOfficer, "unverified member": unverifiedMember,
	} {
		if got, _ := claimViewOf(t, h, viewer, gid); got.HolderName != nil {
			t.Errorf("%s sees holder_name %q, want nil", name, *got.HolderName)
		}
	}
}

func TestGuildStanding_PendingClaimantNameForVerifiedOfficer(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	officer := seedRanked(t, h, gid, "Officer", "officer", true)
	claimant := seedRanked(t, h, gid, "Claimant", "officer", false)
	setGuild(t, h, `update guilds set claim_pending_by = $2, claim_requested_at = now() where id = $1`, gid, claimant)

	got, _ := claimViewOf(t, h, officer, gid)
	if got.State != ClaimStatePending || got.HolderName == nil || *got.HolderName != "Claimant" {
		t.Fatalf("got %+v, want pending with holder_name Claimant", got)
	}
}

func TestGuildStanding_HasApprover(t *testing.T) {
	h := newHTTPHarness(t)
	viewer := seedUser(t, h.pool, "approver-viewer@example.com")

	empty := seedGuild(t, h.pool, "Empty")
	withLeader := seedGuild(t, h.pool, "WithLeader")
	seedRanked(t, h, withLeader, "Boss", "leader", true)
	onlyUnverified := seedGuild(t, h.pool, "OnlyUnverified")
	seedRanked(t, h, onlyUnverified, "Shady", "officer", false)
	seedRanked(t, h, onlyUnverified, "Plain", "member", true)

	for name, tc := range map[string]struct {
		id   int64
		want bool
	}{"no verified members": {empty, false}, "verified leader": {withLeader, true}, "unverified officer only": {onlyUnverified, false}} {
		if _, got := claimViewOf(t, h, viewer, tc.id); got != tc.want {
			t.Errorf("%s: has_approver = %v, want %v", name, got, tc.want)
		}
	}
}
