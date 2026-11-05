// api/internal/guilds/roster_account_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

func TestParseAccountKeyIsAccountKeyForsInverse(t *testing.T) {
	if id, ok := parseAccountKey(accountKeyFor(42)); !ok || id != 42 {
		t.Fatalf("parseAccountKey(accountKeyFor(42)) = %d, %v", id, ok)
	}
	for _, bad := range []string{"", "u:", "u:0", "u:-3", "u:abc", "account:42", "42", "me"} {
		if _, ok := parseAccountKey(bad); ok {
			t.Errorf("parseAccountKey(%q) accepted", bad)
		}
	}
}

// inviteJoin brings a fresh account into gid through the invite link and returns its id.
func inviteJoin(t *testing.T, h *httpHarness, gid int64, email string) int64 {
	t.Helper()
	ctx := context.Background()
	token, _, err := h.store.RotateInvite(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	uid := seedUser(t, h.pool, email)
	if _, err := h.store.AcceptInvite(ctx, token, uid); err != nil {
		t.Fatal(err)
	}
	return uid
}

func TestAnOfficerRemovesAMemberWhoJoinedByInvite(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	officer := seedUser(t, h.pool, "acct-officer@example.com")
	seedVerifiedMember(t, h, gid, officer, "acctofficer", "officer")
	stranger := inviteJoin(t, h, gid, "leaked-link@example.com")
	path := fmt.Sprintf("/v1/guilds/%d/members/%s", gid, accountKeyFor(stranger))

	plain := seedUser(t, h.pool, "acct-plain@example.com")
	seedVerifiedMember(t, h, gid, plain, "acctplain", "member")
	h.actor = auth.Actor{UserID: plain, Role: "user", Method: "session"}
	res := h.do(http.MethodDelete, path, "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a plain member removing an invited account = %d, want 403", res.StatusCode)
	}

	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res = h.do(http.MethodDelete, path, "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("an officer removing an invited account = %d, want 200", res.StatusCode)
	}
	var rows, members int
	h.pool.QueryRow(context.Background(),
		`select count(*) from guild_characters where guild_id = $1 and user_id = $2`, gid, stranger).Scan(&rows)
	h.pool.QueryRow(context.Background(),
		`select count(*) from guild_members where guild_id = $1 and user_id = $2`, gid, stranger).Scan(&members)
	if rows != 0 || members != 0 {
		t.Fatalf("after remove: %d character rows, %d member rows; want none", rows, members)
	}

	res = h.do(http.MethodDelete, path, "")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("removing twice = %d, want 404", res.StatusCode)
	}
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/members/account:%d", gid, stranger), "")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("a malformed account key = %d, want 404", res.StatusCode)
	}
}

func TestRemovingByAccountKeepsRankProtectsRank(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	officer := seedUser(t, h.pool, "acct-officer2@example.com")
	seedVerifiedMember(t, h, gid, officer, "acctofficertwo", "officer")
	other := seedUser(t, h.pool, "acct-other-officer@example.com")
	seedVerifiedMember(t, h, gid, other, "acctotherofficer", "officer")

	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/members/%s", gid, accountKeyFor(other)), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("an officer removing another officer's account = %d, want 403", res.StatusCode)
	}
}

func TestTheHomeOffersRemoveOnAnInvitedEntryToAnOfficerOnly(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	officer := seedUser(t, h.pool, "acct-officer3@example.com")
	seedVerifiedMember(t, h, gid, officer, "acctofficerthree", "officer")
	plain := seedUser(t, h.pool, "acct-plain3@example.com")
	seedVerifiedMember(t, h, gid, plain, "acctplainthree", "member")
	invited := inviteJoin(t, h, gid, "acct-invited3@example.com")

	entry := func(viewer int64) RosterEntry {
		t.Helper()
		h.actor = auth.Actor{UserID: viewer, Role: "user", Method: "session"}
		var home struct {
			Members []RosterEntry `json:"members"`
		}
		h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), ""), &home)
		for _, e := range home.Members {
			if e.AccountKey == accountKeyFor(invited) {
				return e
			}
		}
		t.Fatalf("no invited entry in %+v", home.Members)
		return RosterEntry{}
	}
	if e := entry(officer); !e.MayRemove || e.CharacterKey != "" {
		t.Fatalf("officer's view of the invited entry = %+v, want may_remove and no character key", e)
	}
	if e := entry(plain); e.MayRemove {
		t.Fatalf("a plain member's view offers Remove: %+v", e)
	}
}
