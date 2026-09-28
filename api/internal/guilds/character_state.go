// api/internal/guilds/character_state.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// CharacterGuildState is the private guild summary a character's own
// account is entitled to see: enough for the addon's Guild tab state
// block and the "guild" inbox message (docs/superpowers/specs/2026-09-
// 28-addon-character-aware-design.md §3, §4), and nothing more — this is
// not the public roster/progression the data addon already carries.
type CharacterGuildState struct {
	Name             string
	Region           string
	Ruleset          string
	Claim            ClaimStateView
	// PendingApprovals is the guild's unverified-roster count. It is
	// left at zero for a member who is not an officer or leader:
	// approving a character is an officer action, so a member has no
	// use for another member's queue and should not be shown its size.
	PendingApprovals int
	// Rank is this character's own rank, read off its guild_characters
	// row rather than the verified-only guild_members table
	// (auth.Store.GuildRank): a freshly-synced, not-yet-verified
	// character still has a rank worth showing in game.
	Rank string
}

// CharacterGuildState reads the guild a character currently belongs to,
// if any. ok is false when the character is in no guild — including one
// whose guild_characters row was just cleared by a re-sync with no
// guild= section, which is a real, common state (PutExports/syncGuild),
// not an error.
func (s *Store) CharacterGuildState(ctx context.Context, characterKey string) (CharacterGuildState, bool, error) {
	var guildID int64
	var rank string
	err := s.Pool.QueryRow(ctx,
		`select guild_id, rank from guild_characters where character_key = $1`, characterKey).
		Scan(&guildID, &rank)
	if errors.Is(err, pgx.ErrNoRows) {
		return CharacterGuildState{}, false, nil
	}
	if err != nil {
		return CharacterGuildState{}, false, fmt.Errorf("guilds: character guild state %s: %w", characterKey, err)
	}
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return CharacterGuildState{}, false, fmt.Errorf("guilds: character guild state %s: %w", characterKey, err)
	}
	state := CharacterGuildState{
		Name: g.Name, Region: g.Region, Ruleset: g.Ruleset,
		Claim: claimState(g, time.Now()), Rank: rank,
	}
	if rank != "officer" && rank != "leader" {
		return state, true, nil
	}
	// The same unverified-count the unverified_idx migration comment
	// (0018_guild_membership.up.sql) exists for, and the same rows
	// HomeRoster shows as Verified == false — synthetic account:-
	// prefixed invite rows carry no rank an officer ever holds, and
	// this branch only runs once rank is already known to be officer
	// or leader, so excluding them here would change nothing; skipped
	// anyway for the same reason HomeRoster's own query does.
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from guild_characters
		 where guild_id = $1 and verified_at is null and character_key not like 'account:%'`,
		guildID).Scan(&state.PendingApprovals); err != nil {
		return CharacterGuildState{}, false, fmt.Errorf("guilds: pending approvals for guild %d: %w", guildID, err)
	}
	return state, true, nil
}
