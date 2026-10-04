// api/cmd/seedguild/store.go
//
// Every database read and write this tool performs. Mutations of a pre-existing row (the
// guild itself, the owner's own guild_characters row) snapshot their prior column values
// into seed_rows.prior before changing anything, so --remove can restore exactly what was
// there rather than guess at a default.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
)

// loadGuild reads the guild this tool will seed.
func loadGuild(ctx context.Context, pool *pgxpool.Pool, id int64) (guildInfo, error) {
	var g guildInfo
	g.ID = id
	err := pool.QueryRow(ctx,
		`select region, ruleset, name, claimed_by, claimed_at, officer_max_rank_index from guilds where id = $1`,
		id).Scan(&g.Region, &g.Ruleset, &g.Name, &g.ClaimedBy, &g.ClaimedAt, &g.OfficerMaxRankIndex)
	if err != nil {
		return guildInfo{}, fmt.Errorf("seedguild: load guild %d: %w", id, err)
	}
	return g, nil
}

// resolveOwner finds the account id for a battletag.
func resolveOwner(ctx context.Context, pool *pgxpool.Pool, battletag string) (int64, error) {
	var id int64
	err := pool.QueryRow(ctx, `select id from users where battletag = $1`, battletag).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("seedguild: no account with battletag %q: %w", battletag, err)
	}
	return id, nil
}

// alreadyApplied reports whether tag has any recorded rows - the whole-or-nothing guard
// that makes a second `apply` a safe no-op instead of a partial, constraint-violating
// re-run.
func alreadyApplied(ctx context.Context, pool *pgxpool.Pool, tag string) (bool, error) {
	var exists bool
	err := pool.QueryRow(ctx, `select exists(select 1 from seed_rows where tag = $1)`, tag).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("seedguild: check applied: %w", err)
	}
	return exists, nil
}

// recordRow inserts one seed_rows bookkeeping entry within tx. prior is nil for a row the
// seed created outright.
func recordRow(ctx context.Context, tx pgx.Tx, tag, table, rowKey string, prior any) error {
	var priorJSON []byte
	if prior != nil {
		b, err := json.Marshal(prior)
		if err != nil {
			return fmt.Errorf("seedguild: marshal prior for %s/%s: %w", table, rowKey, err)
		}
		priorJSON = b
	}
	if _, err := tx.Exec(ctx,
		`insert into seed_rows (tag, table_name, row_key, prior) values ($1, $2, $3, $4)
		 on conflict (tag, table_name, row_key) do nothing`,
		tag, table, rowKey, priorJSON); err != nil {
		return fmt.Errorf("seedguild: record %s/%s: %w", table, rowKey, err)
	}
	return nil
}

// guildPrior is the guilds columns this tool mutates to claim the guild for the owner.
type guildPrior struct {
	ClaimedBy           *int64     `json:"claimed_by"`
	ClaimedAt           *time.Time `json:"claimed_at"`
	OfficerMaxRankIndex int        `json:"officer_max_rank_index"`
}

// claimGuildForOwner snapshots the guild's current claim state and sets it to claimed by
// ownerID with the seed's officer threshold.
func claimGuildForOwner(ctx context.Context, pool *pgxpool.Pool, tag string, g guildInfo, ownerID int64) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: claim guild: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	prior := guildPrior{ClaimedBy: g.ClaimedBy, ClaimedAt: g.ClaimedAt, OfficerMaxRankIndex: g.OfficerMaxRankIndex}
	if err := recordRow(ctx, tx, tag, "guilds", fmt.Sprint(g.ID), prior); err != nil {
		return err
	}
	now := time.Now().UTC()
	if _, err := tx.Exec(ctx,
		`update guilds set claimed_by = $2, claimed_at = $3, officer_max_rank_index = $4 where id = $1`,
		g.ID, ownerID, now, seedOfficerMaxRankIndex); err != nil {
		return fmt.Errorf("seedguild: claim guild %d: %w", g.ID, err)
	}
	return tx.Commit(ctx)
}

// ownerCharacterPrior is one of the owner's existing guild_characters rows before this
// tool promotes it.
type ownerCharacterPrior struct {
	RankIndex  *int16     `json:"rank_index"`
	Rank       string     `json:"rank"`
	Source     string     `json:"source"`
	VerifiedBy *string    `json:"verified_by"`
	VerifiedAt *time.Time `json:"verified_at"`
}

// promoteOwner snapshots every guild_characters row the owner already holds in this
// guild, then promotes each to verified leader - the owner's own character becomes the
// guild's leader, exactly as a real claim would leave it.
func promoteOwner(ctx context.Context, pool *pgxpool.Pool, tag string, guildID, ownerID int64) error {
	rows, err := pool.Query(ctx,
		`select character_key, rank_index, rank, source, verified_by, verified_at
		 from guild_characters where guild_id = $1 and user_id = $2`, guildID, ownerID)
	if err != nil {
		return fmt.Errorf("seedguild: read owner membership: %w", err)
	}
	var keys []string
	var priors []ownerCharacterPrior
	for rows.Next() {
		var key string
		var p ownerCharacterPrior
		if err := rows.Scan(&key, &p.RankIndex, &p.Rank, &p.Source, &p.VerifiedBy, &p.VerifiedAt); err != nil {
			rows.Close()
			return fmt.Errorf("seedguild: scan owner membership: %w", err)
		}
		keys = append(keys, key)
		priors = append(priors, p)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}
	if len(keys) == 0 {
		return fmt.Errorf("seedguild: owner %d has no guild_characters row in guild %d to promote", ownerID, guildID)
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: promote owner: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	now := time.Now().UTC()
	for i, key := range keys {
		if err := recordRow(ctx, tx, tag, "guild_characters", fmt.Sprintf("%d:%s", guildID, key), priors[i]); err != nil {
			return err
		}
		if _, err := tx.Exec(ctx,
			`update guild_characters set rank_index = $3, rank = 'leader', verified_by = 'claim', verified_at = $4
			 where guild_id = $1 and character_key = $2`,
			guildID, key, rankIndexLeader, now); err != nil {
			return fmt.Errorf("seedguild: promote %s: %w", key, err)
		}
	}
	if err := guilds.AfterGuildChange(ctx, tx, guildID, ownerID); err != nil {
		return fmt.Errorf("seedguild: recompute owner membership: %w", err)
	}
	return tx.Commit(ctx)
}

// insertMockUser creates one mock account and records it.
func insertMockUser(ctx context.Context, tx pgx.Tx, tag, battletag string) (int64, error) {
	var id int64
	if err := tx.QueryRow(ctx,
		`insert into users (battletag, role) values ($1, 'user') returning id`, battletag).Scan(&id); err != nil {
		return 0, fmt.Errorf("seedguild: insert user %s: %w", battletag, err)
	}
	if err := recordRow(ctx, tx, tag, "users", fmt.Sprint(id), nil); err != nil {
		return 0, err
	}
	return id, nil
}

// insertMockCharacter writes a character and its addon export, mirroring exactly what a
// real export (addon.Store.putOneCharacter / PutExports) would have written.
func insertMockCharacter(ctx context.Context, tx pgx.Tx, guild guildInfo, userID int64, pc plannedCharacter, export string) error {
	if _, err := tx.Exec(ctx,
		`insert into characters (key, region, ruleset, name, class, user_id, source, level, faction, refreshed_at)
		 values ($1, $2, $3, $4, $5, $6, 'export', $7, 'alliance', now())`,
		pc.Key, guild.Region, guild.Ruleset, pc.Mock.Name, pc.Mock.Class, userID, fs1Level); err != nil {
		return fmt.Errorf("seedguild: insert character %s: %w", pc.Key, err)
	}
	if _, err := tx.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, updated_at, source, captured_at)
		 values ($1, $2, $3, $4, $5, $6, $7, 'addon', $7)`,
		pc.Key, userID, guild.Region, guild.Ruleset, pc.Mock.Name, export, pc.ExportedAt); err != nil {
		return fmt.Errorf("seedguild: insert addon export %s: %w", pc.Key, err)
	}
	return nil
}

// upsertMockMembership writes pc's guild_characters row (verified for everyone but the
// seed's three pending characters), recomputes the derived guild_members row, and sets
// that account's consent.
func upsertMockMembership(ctx context.Context, tx pgx.Tx, guildID, userID int64, pc plannedCharacter) error {
	row := guilds.MembershipRow{
		GuildID: guildID, CharacterKey: pc.Key, UserID: userID,
		RankIndex: pc.RankIndex, Rank: pc.Rank, Source: "export",
	}
	if pc.Verified {
		row.Reverify, row.VerifiedBy, row.VerifiedAt = true, "officer", time.Now().UTC()
	}
	if err := guilds.UpsertCharacterMembership(ctx, tx, row); err != nil {
		return err
	}
	if err := guilds.AfterGuildChange(ctx, tx, guildID, userID); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx,
		`update guild_members set consent = $3 where guild_id = $1 and user_id = $2`,
		guildID, userID, pc.Mock.Consent); err != nil {
		return fmt.Errorf("seedguild: set consent for %s: %w", pc.Key, err)
	}
	return nil
}

// insertReport writes one raid night's report row.
func insertReport(ctx context.Context, tx pgx.Tx, guildID, ownerID int64, loggingCharacter string, pr plannedReport) error {
	if _, err := tx.Exec(ctx,
		`insert into reports (id, owner_id, guild_id, title, visibility, zone, status, engine_version,
		   logging_character, created_at, completed_at)
		 values ($1, $2, $3, $4, 'guild', $5, 'complete', 'seedguild', $6, $7, $8)`,
		pr.ID, ownerID, guildID, pr.Title, pr.Plan.Zone, loggingCharacter,
		pr.CreatedAt, pr.CompletedAt); err != nil {
		return fmt.Errorf("seedguild: insert report %s: %w", pr.ID, err)
	}
	return nil
}

// insertFightsAndMetrics writes every fight and every roster character's fight_metrics
// row for one report, in the shape rankings.Store.WriteFight itself writes (phase, state,
// empty trinkets, zero buff_count, no faction) - see that function's own insert.
func insertFightsAndMetrics(ctx context.Context, tx pgx.Tx, pool *pgxpool.Pool, pr plannedReport, chars []plannedCharacter) error {
	keys := make([]string, len(chars))
	roster := make([]mockCharacter, len(chars))
	for i, c := range chars {
		keys[i] = c.Key
		roster[i] = c.Mock
	}
	size := int64(len(chars))

	for idx, f := range pr.Plan.Fights {
		foughtAt := pr.FightStarts[idx]
		rng := seededRand(pr.Plan.Tag, idx)
		rows := buildFightMetrics(roster, keys, f, rng)
		fightDeaths := 0
		for _, row := range rows {
			fightDeaths += row.Deaths
		}

		var encounterID, difficulty *int64
		if f.EncounterID != 0 {
			enc, diff := f.EncounterID, int64(raidDifficulty)
			encounterID, difficulty = &enc, &diff
		}
		if _, err := tx.Exec(ctx,
			`insert into fights (report_id, fight_index, encounter_id, name, difficulty, size, kill,
			   duration_ms, start_ms, verified, players, deaths, npc_kills)
			 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, true, $10, $11, 0)`,
			pr.ID, idx, encounterID, f.Name, difficulty, size, f.Kill, f.DurationMS,
			foughtAt.UnixMilli(), keys, fightDeaths); err != nil {
			return fmt.Errorf("seedguild: insert fight %s/%d: %w", pr.ID, idx, err)
		}

		if err := db.EnsureMetricsPartition(ctx, pool, foughtAt); err != nil {
			return err
		}
		at := phase.At(foughtAt)
		for i, row := range rows {
			if _, err := tx.Exec(ctx,
				`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec,
				   role, ilvl, metric_dps, metric_hps, damage_taken, active_ms, deaths, encounter_id,
				   difficulty, size, duration_ms, kill, phase, fought_at, talent_split, trinkets,
				   buff_count, state)
				 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18,
				   $19, $20, $21, $22, 0, 'ok')`,
				pr.ID, idx, row.PlayerKey, row.PlayerName, row.Class, row.Spec, row.Role, row.Ilvl,
				row.MetricDPS, row.MetricHPS, row.DamageTaken, row.ActiveMS, row.Deaths,
				encounterID, difficulty, size, f.DurationMS, f.Kill, at, foughtAt,
				talentSplitFor(roster[i].Role), []int64{}); err != nil {
				return fmt.Errorf("seedguild: insert fight_metrics %s/%d/%s: %w", pr.ID, idx, row.PlayerKey, err)
			}
		}
	}
	return nil
}
