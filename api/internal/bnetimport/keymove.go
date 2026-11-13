package bnetimport

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// characterKeyColumns is every column outside characters that holds a
// character key and has no ON UPDATE CASCADE foreign key to carry it
// (addon_sim_requests, character_syncs and guild_roster_snapshots do).
// TestCharacterKeyColumnsCoverTheSchema keeps it complete.
var characterKeyColumns = []struct{ table, column string }{
	{"activity_events", "character_key"},
	{"addon_exports", "character_key"},
	{"addon_inbox", "character_key"},
	{"character_sims", "character_key"},
	{"coach_packs", "player_key"},
	{"coach_profiles", "player_key"},
	{"coach_review_leases", "player_key"},
	{"coach_review_samples", "player_key"},
	{"coach_reviews", "player_key"},
	{"fight_metrics", "player_key"},
	{"guild_applications", "character_key"},
	{"guild_characters", "character_key"},
	{"guild_raid_attendance", "character_key"},
	{"guild_raid_attendance_overrides", "character_key"},
	{"guild_raid_signups", "character_key"},
	{"loot_awards", "character_key"},
	{"loot_interest", "character_key"},
	{"loot_votes", "character_key"},
	{"raid_assignments", "character_key"},
	{"rating_scores", "player_key"},
	{"reports", "logging_character"},
	{"spec_board_weeks", "character_key"},
	{"users", "main_character_key"},
	{"watched_characters", "character_key"},
}

// KeyMoveCounts is how many rows a key move changed, by "table.column".
type KeyMoveCounts map[string]int64

// moveKey moves a characters row from oldKey to newKey under newRuleset,
// and with it every row keyed by the character (characterKeyColumns: its
// guild membership, the account's main, its exports, scores, sims, raid
// and loot rows, and any report it is credited with), inside tx. A newKey
// already taken by another row (the same name on two realms sharing a
// ruleset) is logged and nothing moves: moved is false.
func (s *Service) moveKey(ctx context.Context, tx pgx.Tx, oldKey, newKey, newRuleset string) (KeyMoveCounts, bool, error) {
	var taken bool
	if err := tx.QueryRow(ctx, `select exists (select 1 from characters where key = $1)`, newKey).Scan(&taken); err != nil {
		return nil, false, fmt.Errorf("bnetimport: rekey check %s: %w", newKey, err)
	}
	if taken {
		s.logger().Warn("bnetimport", "op", "rekey_taken", "old_key", oldKey, "new_key", newKey)
		return nil, false, nil
	}
	counts := KeyMoveCounts{}
	tag, err := tx.Exec(ctx, `update characters set key = $1, ruleset = $2, ruleset_answered = false where key = $3`, newKey, newRuleset, oldKey)
	if err != nil {
		return nil, false, fmt.Errorf("bnetimport: rekey %s: %w", oldKey, err)
	}
	counts["characters.key"] = tag.RowsAffected()
	for _, c := range characterKeyColumns {
		// The names come from the constant table above, never from input.
		tag, err := tx.Exec(ctx, fmt.Sprintf(`update %s set %s = $1 where %s = $2`, c.table, c.column, c.column), newKey, oldKey)
		if err != nil {
			return nil, false, fmt.Errorf("bnetimport: rekey %s.%s of %s: %w", c.table, c.column, oldKey, err)
		}
		if n := tag.RowsAffected(); n > 0 {
			counts[c.table+"."+c.column] = n
		}
	}
	// The export row names its ruleset as well as its key.
	if _, err := tx.Exec(ctx, `update addon_exports set ruleset = $2 where character_key = $1`, newKey, newRuleset); err != nil {
		return nil, false, fmt.Errorf("bnetimport: rekey export ruleset of %s: %w", newKey, err)
	}
	s.logger().Info("bnetimport", "op", "rekey", "old_key", oldKey, "new_key", newKey)
	return counts, true, nil
}
