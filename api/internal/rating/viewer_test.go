// api/internal/rating/viewer_test.go
package rating

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// stubAccounts is a test double for Accounts. Its GuildRank signature matches the real
// auth.Store.GuildRank (context.Context, then the two int64 IDs) - the pattern mirrored
// from api/internal/reports/harness_test.go, whose own harness wires the request's actor
// onto the context via auth.WithActor rather than inventing a bespoke context type per
// test file.
type stubAccounts struct {
	rank string
	ok   bool
	err  error
}

func (s stubAccounts) GuildRank(_ context.Context, _, _ int64) (string, bool, error) {
	return s.rank, s.ok, s.err
}

// requestAs builds a GET request whose context carries actor, exactly the way
// reports/harness_test.go's own httptest.Server middleware does it:
// r.WithContext(auth.WithActor(r.Context(), actor)).
func requestAs(actor auth.Actor) *http.Request {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	return r.WithContext(auth.WithActor(r.Context(), actor))
}

func TestVisiblePublicAndUnlistedAlwaysTrue(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	for _, v := range []string{reports.Public, reports.Unlisted} {
		if !visible(r, reports.Report{Visibility: v}, nil) {
			t.Errorf("visibility %q must be visible to an anonymous request", v)
		}
	}
}

func TestVisiblePrivateFalseForAnonymous(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	if visible(r, reports.Report{Visibility: reports.Private}, nil) {
		t.Error("a private report must not be visible to an anonymous request")
	}
}

func TestVisiblePrivateFalseForASignedInStranger(t *testing.T) {
	owner := int64(1)
	r := requestAs(auth.Actor{UserID: 2, Role: "user", Method: "session"})
	rep := reports.Report{Visibility: reports.Private, OwnerID: &owner}
	if visible(r, rep, nil) {
		t.Error("a private report must not be visible to a signed-in non-owner, non-moderator")
	}
}

func TestVisiblePrivateTrueForOwner(t *testing.T) {
	owner := int64(1)
	r := requestAs(auth.Actor{UserID: owner, Role: "user", Method: "session"})
	rep := reports.Report{Visibility: reports.Private, OwnerID: &owner}
	if !visible(r, rep, nil) {
		t.Error("a private report must be visible to its owner")
	}
}

func TestVisiblePrivateTrueForModerator(t *testing.T) {
	owner := int64(1)
	r := requestAs(auth.Actor{UserID: 99, Role: "moderator", Method: "session"})
	rep := reports.Report{Visibility: reports.Private, OwnerID: &owner}
	if !visible(r, rep, nil) {
		t.Error("a private report must be visible to a moderator")
	}
}

func TestVisibleGuildRequiresVerifiedMembership(t *testing.T) {
	owner := int64(1)
	guildID := int64(42)
	r := requestAs(auth.Actor{UserID: 2, Role: "user", Method: "session"})
	rep := reports.Report{Visibility: reports.GuildTo, OwnerID: &owner, GuildID: &guildID}
	accounts := stubAccounts{rank: "", ok: false, err: nil}
	if visible(r, rep, accounts) {
		t.Error("a guild report must not be visible to a non-member, non-owner, non-moderator")
	}
}

func TestVisibleGuildTrueForVerifiedMember(t *testing.T) {
	owner := int64(1)
	guildID := int64(42)
	r := requestAs(auth.Actor{UserID: 2, Role: "user", Method: "session"})
	rep := reports.Report{Visibility: reports.GuildTo, OwnerID: &owner, GuildID: &guildID}
	accounts := stubAccounts{rank: "member", ok: true, err: nil}
	if !visible(r, rep, accounts) {
		t.Error("a guild report must be visible to a verified guild member")
	}
}
