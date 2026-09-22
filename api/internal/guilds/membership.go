// api/internal/guilds/membership.go

package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// MembershipRow is one guild_characters upsert — the single write shape
// both the export path (addon.Store's guild sync) and the Battle.net
// import path (bnetimport.ImportAccount) use, so guild_characters has one
// writer rather than two SQL statements that can drift apart (spec §4.2
// step 5).
type MembershipRow struct {
	GuildID      int64
	CharacterKey string
	UserID       int64
	RankIndex    int
	Rank         string
	Source       string
	// VerifiedBy and VerifiedAt are only written when Reverify is true.
	VerifiedBy string
	VerifiedAt time.Time
	// Reverify is false for an export re-sync: a re-sync of the same
	// character in the same guild keeps whatever verification (and
	// corroboration source) it already earned, including a 'bnet' row an
	// export must not downgrade (spec §4.3's "the freshest evidence wins
	// on guild identity; Blizzard wins on verification"). Reverify is
	// true for a Battle.net import or refresh: the roster it just read is
	// fresh evidence, so source/verified_by/verified_at are re-stamped
	// every time, even on a re-sync of an already-'bnet' row.
	Reverify bool
}

// UpsertCharacterMembership writes one guild_characters row within tx.
// Extracted from the export path's own syncGuild (spec §4.2 step 5); the
// Reverify-false branch is byte-for-byte the SQL that path always ran.
func UpsertCharacterMembership(ctx context.Context, tx pgx.Tx, row MembershipRow) error {
	if row.Reverify {
		if _, err := tx.Exec(ctx,
			`insert into guild_characters
			   (guild_id, character_key, user_id, rank_index, rank, source, verified_by, verified_at, refreshed_at)
			 values ($1, $2, $3, $4, $5, $6, $7, $8, now())
			 on conflict (guild_id, character_key) do update set
			   user_id = excluded.user_id, rank_index = excluded.rank_index, rank = excluded.rank,
			   source = excluded.source, verified_by = excluded.verified_by, verified_at = excluded.verified_at,
			   refreshed_at = now()`,
			row.GuildID, row.CharacterKey, row.UserID, row.RankIndex, row.Rank,
			row.Source, row.VerifiedBy, row.VerifiedAt); err != nil {
			return fmt.Errorf("guilds: upsert membership for %s: %w", row.CharacterKey, err)
		}
		return nil
	}
	// verified_at/verified_by/source are deliberately not in this SET
	// list: a re-sync of the same character in the same guild keeps
	// whatever verification it already earned.
	if _, err := tx.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank_index, rank, source, refreshed_at)
		 values ($1, $2, $3, $4, $5, $6, now())
		 on conflict (guild_id, character_key) do update set
		   user_id = excluded.user_id, rank_index = excluded.rank_index, rank = excluded.rank,
		   refreshed_at = now()`,
		row.GuildID, row.CharacterKey, row.UserID, row.RankIndex, row.Rank, row.Source); err != nil {
		return fmt.Errorf("guilds: upsert membership for %s: %w", row.CharacterKey, err)
	}
	return nil
}

// ResolveGuild finds or creates the guild an export or a Battle.net roster
// names, matching case-insensitively on (region, ruleset, lower(name)) —
// two sources differing only in casing must resolve to the same guilds
// row. The first writer's casing is kept as the display name; a
// concurrent insert racing on the same case-insensitive name is tolerated
// by falling back to the row the winner created.
//
// Moved here from addon.resolveGuild (spec §4.2 step 5) so the Battle.net
// import path can share it; addon.Store's own guild sync now calls this
// exported form.
func ResolveGuild(ctx context.Context, tx pgx.Tx, region, ruleset, name string) (id int64, officerMax int, err error) {
	err = tx.QueryRow(ctx,
		`select id, officer_max_rank_index from guilds where region = $1 and ruleset = $2 and lower(name) = lower($3)`,
		region, ruleset, name).Scan(&id, &officerMax)
	if err == nil {
		return id, officerMax, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return 0, 0, fmt.Errorf("guilds: read guild %s: %w", name, err)
	}
	err = tx.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ($1, $2, $3)
		 on conflict (region, ruleset, (lower(name))) do nothing
		 returning id, officer_max_rank_index`,
		region, ruleset, name).Scan(&id, &officerMax)
	if errors.Is(err, pgx.ErrNoRows) {
		// A concurrent insert of a case-variant name won the race; read
		// the row it created.
		err = tx.QueryRow(ctx,
			`select id, officer_max_rank_index from guilds where region = $1 and ruleset = $2 and lower(name) = lower($3)`,
			region, ruleset, name).Scan(&id, &officerMax)
	}
	if err != nil {
		return 0, 0, fmt.Errorf("guilds: create/read guild %s: %w", name, err)
	}
	return id, officerMax, nil
}

// DeriveRank turns a raw GetGuildInfo (or a Blizzard roster's) rank index
// into the label officer detection uses. Index 0 is always the guild
// master (server-authoritative, never configurable). Moved here from
// addon.deriveRank (spec §4.2 step 5).
func DeriveRank(rankIndex, officerMax int) string {
	switch {
	case rankIndex == 0:
		return "leader"
	case rankIndex <= officerMax:
		return "officer"
	default:
		return "member"
	}
}

// AfterGuildChange runs RecomputeMembership and the lost-claim check for
// one account in one guild — the pair of calls every caller needs after
// it changes a guild_characters row. Moved here from addon.afterGuildChange
// (spec §4.2 step 5) so the Battle.net import path can share it.
func AfterGuildChange(ctx context.Context, tx pgx.Tx, guildID, userID int64) error {
	if err := RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
		return err
	}
	return ReleaseClaimIfLost(ctx, tx, guildID, userID)
}
