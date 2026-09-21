// api/internal/guilds/settings_test.go
package guilds

import (
	"context"
	"errors"
	"testing"
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
