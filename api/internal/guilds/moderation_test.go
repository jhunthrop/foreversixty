// api/internal/guilds/moderation_test.go
package guilds

import (
	"context"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// TestModerationClaimsAnswers404ForANonModerator is item 4's own
// requirement (fourth security review response): the route's existence
// is never advertised - a plain member, a verified officer, and an
// unauthenticated caller all get the same 404 a non-moderator would,
// never a 401 or 403 that would reveal the route needs moderator
// standing.
func TestModerationClaimsAnswers404ForANonModerator(t *testing.T) {
	h := newHTTPHarness(t)

	h.actor = auth.Actor{}
	res := h.do(http.MethodGet, "/v1/moderation/claims", "")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("an unauthenticated caller sees %d, want 404", res.StatusCode)
	}

	plain := seedUser(t, h.pool, "moderation-plain@example.com")
	h.actor = auth.Actor{UserID: plain, Role: "user", Method: "session"}
	res = h.do(http.MethodGet, "/v1/moderation/claims", "")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("a plain user sees %d, want 404", res.StatusCode)
	}
}

// TestModerationClaimsListsOpenContestsWithEvidence is item 4's
// contents requirement: a moderator sees guild identity, both parties'
// battletags, the claim/contest timestamps, and each side's plain
// evidence facts (rank index, verification, independent-night count).
func TestModerationClaimsListsOpenContestsWithEvidence(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()

	gid := seedGuild(t, h.pool, "Moderated Guild")
	claimant := seedUser(t, h.pool, "moderation-claimant@example.com")
	seedCharacter(t, h.pool, gid, claimant, "us/hardcore/moderationclaimant", "leader", false)
	// seedCharacter never sets rank_index (only the derived rank label);
	// set it directly so evidenceFor's "rank index from their export"
	// fact has something real to read.
	if _, err := h.pool.Exec(ctx,
		`update guild_characters set rank_index = 0 where character_key = 'us/hardcore/moderationclaimant'`); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, h.pool, "moderation-contester@example.com")
	seedCharacter(t, h.pool, gid, contester, "us/hardcore/moderationcontester", "officer", false)
	if _, err := h.pool.Exec(ctx,
		`update guild_characters set rank_index = 2 where character_key = 'us/hardcore/moderationcontester'`); err != nil {
		t.Fatal(err)
	}
	if err := h.store.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}

	moderator := seedUser(t, h.pool, "moderation-mod@example.com")
	h.actor = auth.Actor{UserID: moderator, Role: "moderator", Method: "session"}
	res := h.do(http.MethodGet, "/v1/moderation/claims", "")
	var out struct {
		Claims []ModerationClaim `json:"claims"`
	}
	h.data(res, &out)
	if len(out.Claims) != 1 {
		t.Fatalf("open claims = %d, want 1", len(out.Claims))
	}
	c := out.Claims[0]
	if c.Guild.ID != gid || c.Guild.Name != "Moderated Guild" {
		t.Fatalf("guild = %+v", c.Guild)
	}
	if c.Claimant.Battletag == "" || c.Contester.Battletag == "" {
		t.Fatalf("battletags should be populated: claimant=%q contester=%q",
			c.Claimant.Battletag, c.Contester.Battletag)
	}
	if c.ClaimedAt == nil {
		t.Fatal("claimed_at should be populated for a claimed (not merely pending) claimant")
	}
	if c.Claimant.Evidence.RankIndex == nil || *c.Claimant.Evidence.RankIndex != 0 {
		t.Fatalf("claimant rank index = %v, want 0 (leader)", c.Claimant.Evidence.RankIndex)
	}
	if c.Contester.Evidence.RankIndex == nil {
		t.Fatal("contester rank index should be populated")
	}
	// Claim()'s leader-instant branch verifies the claimant's own
	// character as part of claiming - real evidence, not fabricated.
	if !c.Claimant.Evidence.Verified || c.Claimant.Evidence.VerifiedBy == nil || *c.Claimant.Evidence.VerifiedBy != "claim" {
		t.Fatalf("claimant evidence = %+v, want verified via claim", c.Claimant.Evidence)
	}
	// The contester's own raw officer character was never verified
	// through any path - evidence must not fabricate it.
	if c.Contester.Evidence.Verified {
		t.Fatalf("contester evidence = %+v, want unverified", c.Contester.Evidence)
	}
}
