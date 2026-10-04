// api/cmd/seedguild/remove.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// seedRow is one bookkeeping row this tool reads back during --remove.
type seedRow struct {
	RowKey string
	Prior  []byte
}

func readSeedRows(ctx context.Context, pool *pgxpool.Pool, tag, table string) ([]seedRow, error) {
	rows, err := pool.Query(ctx, `select row_key, prior from seed_rows where tag = $1 and table_name = $2`, tag, table)
	if err != nil {
		return nil, fmt.Errorf("seedguild: read seed_rows %s/%s: %w", tag, table, err)
	}
	defer rows.Close()
	var out []seedRow
	for rows.Next() {
		var r seedRow
		if err := rows.Scan(&r.RowKey, &r.Prior); err != nil {
			return nil, fmt.Errorf("seedguild: scan seed_rows %s/%s: %w", tag, table, err)
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// RemoveSeed deletes everything --apply wrote for this guild and restores the owner's own
// guild_characters row and the guild's claim state to what they were beforehand. A no-op
// (after printing a notice) if the tag has nothing recorded.
func RemoveSeed(ctx context.Context, pool *pgxpool.Pool, guildID int64, ownerBattletag string) error {
	tag := seedTag(guildID)
	applied, err := alreadyApplied(ctx, pool, tag)
	if err != nil {
		return err
	}
	if !applied {
		fmt.Printf("seedguild: %q has no recorded rows; nothing to remove\n", tag)
		return nil
	}
	ownerID, err := resolveOwner(ctx, pool, ownerBattletag)
	if err != nil {
		return err
	}

	reportRows, err := readSeedRows(ctx, pool, tag, "reports")
	if err != nil {
		return err
	}
	for _, r := range reportRows {
		if err := removeReport(ctx, pool, tag, r.RowKey); err != nil {
			return err
		}
	}

	userRows, err := readSeedRows(ctx, pool, tag, "users")
	if err != nil {
		return err
	}
	for _, r := range userRows {
		if err := removeMockUser(ctx, pool, tag, r.RowKey); err != nil {
			return err
		}
	}

	charRows, err := readSeedRows(ctx, pool, tag, "guild_characters")
	if err != nil {
		return err
	}
	for _, r := range charRows {
		if err := restoreOwnerCharacter(ctx, pool, tag, guildID, ownerID, r); err != nil {
			return err
		}
	}

	guildRows, err := readSeedRows(ctx, pool, tag, "guilds")
	if err != nil {
		return err
	}
	for _, r := range guildRows {
		if err := restoreGuildClaim(ctx, pool, tag, r); err != nil {
			return err
		}
	}

	fmt.Printf("seedguild: removed %q from guild %d - %d reports, %d accounts, "+
		"owner membership and claim restored\n", tag, guildID, len(reportRows), len(userRows))
	return nil
}

// PrintRemoveDryRun prints what RemoveSeed would delete and restore, without writing
// anything: the same seed_rows reads RemoveSeed itself opens with, just never followed by
// a write.
func PrintRemoveDryRun(ctx context.Context, pool *pgxpool.Pool, guildID int64) error {
	tag := seedTag(guildID)
	applied, err := alreadyApplied(ctx, pool, tag)
	if err != nil {
		return err
	}
	if !applied {
		fmt.Printf("seedguild: %q has no recorded rows; nothing to remove\n", tag)
		return nil
	}
	reportRows, err := readSeedRows(ctx, pool, tag, "reports")
	if err != nil {
		return err
	}
	userRows, err := readSeedRows(ctx, pool, tag, "users")
	if err != nil {
		return err
	}
	charRows, err := readSeedRows(ctx, pool, tag, "guild_characters")
	if err != nil {
		return err
	}
	guildRows, err := readSeedRows(ctx, pool, tag, "guilds")
	if err != nil {
		return err
	}
	fmt.Printf("seedguild: dry run of --remove for guild %d, tag %q\n", guildID, tag)
	fmt.Printf("reports: %d row(s) to delete (and their fights, fight_metrics)\n", len(reportRows))
	fmt.Printf("users: %d row(s) to delete (and their characters, addon_exports, guild_characters, guild_members)\n", len(userRows))
	fmt.Printf("guild_characters: %d owner row(s) to restore to their prior rank/verification\n", len(charRows))
	fmt.Printf("guilds: %d row to restore to its prior claim state\n", len(guildRows))
	return nil
}

func removeReport(ctx context.Context, pool *pgxpool.Pool, tag, reportID string) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: remove report %s: begin: %w", reportID, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx, `delete from fight_metrics where report_id = $1`, reportID); err != nil {
		return fmt.Errorf("seedguild: remove fight_metrics for %s: %w", reportID, err)
	}
	// fights cascades from reports (migration 0005: "references reports (id) on delete cascade").
	if _, err := tx.Exec(ctx, `delete from reports where id = $1`, reportID); err != nil {
		return fmt.Errorf("seedguild: remove report %s: %w", reportID, err)
	}
	if _, err := tx.Exec(ctx,
		`delete from seed_rows where tag = $1 and table_name = 'reports' and row_key = $2`, tag, reportID); err != nil {
		return fmt.Errorf("seedguild: unrecord report %s: %w", reportID, err)
	}
	return tx.Commit(ctx)
}

func removeMockUser(ctx context.Context, pool *pgxpool.Pool, tag, userIDStr string) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: remove user %s: begin: %w", userIDStr, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	// characters.user_id is ON DELETE SET NULL (migration 0005), not cascade, so it must
	// be deleted explicitly before the user row goes; addon_exports, guild_characters and
	// guild_members all cascade from users (migrations 0005/0018).
	if _, err := tx.Exec(ctx, `delete from characters where user_id = $1::bigint`, userIDStr); err != nil {
		return fmt.Errorf("seedguild: remove characters for user %s: %w", userIDStr, err)
	}
	if _, err := tx.Exec(ctx, `delete from users where id = $1::bigint`, userIDStr); err != nil {
		return fmt.Errorf("seedguild: remove user %s: %w", userIDStr, err)
	}
	if _, err := tx.Exec(ctx,
		`delete from seed_rows where tag = $1 and table_name = 'users' and row_key = $2`, tag, userIDStr); err != nil {
		return fmt.Errorf("seedguild: unrecord user %s: %w", userIDStr, err)
	}
	return tx.Commit(ctx)
}

func restoreOwnerCharacter(ctx context.Context, pool *pgxpool.Pool, tag string, guildID, ownerID int64, r seedRow) error {
	var prior ownerCharacterPrior
	if err := json.Unmarshal(r.Prior, &prior); err != nil {
		return fmt.Errorf("seedguild: unmarshal owner prior for %s: %w", r.RowKey, err)
	}
	// row_key is "guildID:characterKey"; the character key is everything after the first colon.
	_, characterKey, ok := strings.Cut(r.RowKey, ":")
	if !ok {
		return fmt.Errorf("seedguild: malformed guild_characters row_key %q", r.RowKey)
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: restore owner character %s: begin: %w", characterKey, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx,
		`update guild_characters set rank_index = $3, rank = $4, source = $5, verified_by = $6, verified_at = $7
		 where guild_id = $1 and character_key = $2`,
		guildID, characterKey, prior.RankIndex, prior.Rank, prior.Source, prior.VerifiedBy, prior.VerifiedAt); err != nil {
		return fmt.Errorf("seedguild: restore owner character %s: %w", characterKey, err)
	}
	if err := guilds.AfterGuildChange(ctx, tx, guildID, ownerID); err != nil {
		return fmt.Errorf("seedguild: recompute owner membership: %w", err)
	}
	if _, err := tx.Exec(ctx,
		`delete from seed_rows where tag = $1 and table_name = 'guild_characters' and row_key = $2`,
		tag, r.RowKey); err != nil {
		return fmt.Errorf("seedguild: unrecord owner character %s: %w", characterKey, err)
	}
	return tx.Commit(ctx)
}

func restoreGuildClaim(ctx context.Context, pool *pgxpool.Pool, tag string, r seedRow) error {
	var prior guildPrior
	if err := json.Unmarshal(r.Prior, &prior); err != nil {
		return fmt.Errorf("seedguild: unmarshal guild prior for %s: %w", r.RowKey, err)
	}
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: restore guild %s: begin: %w", r.RowKey, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx,
		`update guilds set claimed_by = $2, claimed_at = $3, officer_max_rank_index = $4 where id = $1::bigint`,
		r.RowKey, prior.ClaimedBy, prior.ClaimedAt, prior.OfficerMaxRankIndex); err != nil {
		return fmt.Errorf("seedguild: restore guild %s: %w", r.RowKey, err)
	}
	if _, err := tx.Exec(ctx,
		`delete from seed_rows where tag = $1 and table_name = 'guilds' and row_key = $2`, tag, r.RowKey); err != nil {
		return fmt.Errorf("seedguild: unrecord guild %s: %w", r.RowKey, err)
	}
	return tx.Commit(ctx)
}
