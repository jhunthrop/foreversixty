// api/internal/guilds/settings_test.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

func TestSettingsReadsTheCurrentGuildState(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")

	view, err := s.Settings(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	if view.DefaultVisibility != "public" || view.OfficerMaxRankIndex != 1 {
		t.Fatalf("view = %+v, want the schema defaults", view)
	}
	if view.ClaimedBy != nil {
		t.Fatal("an unclaimed guild should read claimed_by = nil")
	}
}

func TestUpdateSettingsValidatesVisibilityAndThreshold(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")

	private := "private"
	if _, err := s.UpdateSettings(ctx, gid, SettingsPatch{DefaultVisibility: &private}); !errors.Is(err, ErrInvalidSettings) {
		t.Fatalf("private default_visibility = %v, want ErrInvalidSettings", err)
	}
	tooHigh := 10
	if _, err := s.UpdateSettings(ctx, gid, SettingsPatch{OfficerMaxRankIndex: &tooHigh}); !errors.Is(err, ErrInvalidSettings) {
		t.Fatalf("officer_max_rank_index = 10 = %v, want ErrInvalidSettings", err)
	}
	guildVis := "guild"
	view, err := s.UpdateSettings(ctx, gid, SettingsPatch{DefaultVisibility: &guildVis})
	if err != nil || view.DefaultVisibility != "guild" {
		t.Fatalf("view = %+v, %v, want default_visibility = guild", view, err)
	}
}

func TestUpdateSettingsRederivesEveryCharactersRankAndRecomputesTheAccounts(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	uid := seedUser(t, pool, "rerank@example.com")
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank_index, rank, verified_at)
		 values ($1, 'us/hardcore/rerank', $2, 3, 'member', now())`, gid, uid); err != nil {
		t.Fatal(err)
	}

	widened := 5
	if _, err := s.UpdateSettings(ctx, gid, SettingsPatch{OfficerMaxRankIndex: &widened}); err != nil {
		t.Fatal(err)
	}
	var rank, memberRank string
	pool.QueryRow(ctx, `select rank from guild_characters where character_key = 'us/hardcore/rerank'`).Scan(&rank)
	if rank != "officer" {
		t.Fatalf("guild_characters.rank = %q, want officer once the threshold widened past rank_index 3", rank)
	}
	pool.QueryRow(ctx, `select rank from guild_members where guild_id = $1 and user_id = $2`, gid, uid).Scan(&memberRank)
	if memberRank != "officer" {
		t.Fatalf("guild_members.rank = %q, want officer (recomputed in the same call)", memberRank)
	}
}

func TestSettingsExposesClaimStateAndPopulatesClaimant(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	claimant := seedUser(t, h.pool, "settings-claimant@example.com")
	seedCharacter(t, h.pool, gid, claimant, "us/hardcore/settingsclaimant", "leader", true)
	syncMembership(t, h.pool, gid, claimant)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, claimant, gid); err != nil {
		t.Fatal(err)
	}
	h.actor = auth.Actor{UserID: claimant, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/settings", gid), "")
	var view SettingsView
	h.data(res, &view)
	if view.Claim.State != "claimed" {
		t.Fatalf("claim.state = %q, want claimed", view.Claim.State)
	}
	if view.ClaimedBy == nil || view.ClaimedBy.Battletag == "" {
		t.Fatalf("claimed_by = %v, want populated (this was the bug the review found: always null)", view.ClaimedBy)
	}
}

func TestPatchSettingsIsFrozenDuringAContestedClaim(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	officer := seedUser(t, h.pool, "frozen-settings@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/frozensettings", "officer", true)
	syncMembership(t, h.pool, gid, officer)
	if _, err := h.pool.Exec(ctx, `update guilds set claim_contested_at = now() where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodPatch, fmt.Sprintf("/v1/guilds/%d/settings", gid), `{"default_visibility":"public"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("patch settings while contested = %d, want 409", res.StatusCode)
	}
}
