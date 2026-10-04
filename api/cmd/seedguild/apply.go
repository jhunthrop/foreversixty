// api/cmd/seedguild/apply.go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ApplySeed writes the whole mock guild: the owner's claim and promotion to leader, the
// 24-character roster (22 accounts - two accounts carry a main and an alt), and the four
// raid nights with their fights and fight_metrics. It is a no-op (after printing a
// notice) if tag already has recorded rows, so a second run never duplicates or
// partially re-applies the seed.
func ApplySeed(ctx context.Context, pool *pgxpool.Pool, guildID int64, ownerBattletag string) error {
	guild, err := loadGuild(ctx, pool, guildID)
	if err != nil {
		return err
	}
	ownerID, err := resolveOwner(ctx, pool, ownerBattletag)
	if err != nil {
		return err
	}
	tag := seedTag(guildID)
	applied, err := alreadyApplied(ctx, pool, tag)
	if err != nil {
		return err
	}
	if applied {
		fmt.Printf("seedguild: %q already has seeded rows; nothing to do (use --remove first to reseed)\n", tag)
		return nil
	}

	plan := buildSeedPlan(guild, time.Now().UTC())

	if err := claimGuildForOwner(ctx, pool, tag, guild, ownerID); err != nil {
		return err
	}
	if err := promoteOwner(ctx, pool, tag, guildID, ownerID); err != nil {
		return err
	}

	// accountUsers tracks which account key already has a user row, so the second of a
	// main-and-alt pair (roster.go's Account field) reuses that account's user id
	// instead of getting one of its own.
	accountUsers := map[string]int64{}
	for _, pc := range plan.Characters {
		if err := writeMockCharacter(ctx, pool, tag, guild, pc, accountUsers); err != nil {
			return err
		}
	}

	var ownerCharKey string
	if err := pool.QueryRow(ctx,
		`select character_key from guild_characters where guild_id = $1 and user_id = $2 limit 1`,
		guildID, ownerID).Scan(&ownerCharKey); err != nil {
		return fmt.Errorf("seedguild: read owner character key: %w", err)
	}

	for _, pr := range plan.Reports {
		if err := writeMockReport(ctx, pool, tag, guildID, ownerID, ownerCharKey, pr, plan.Characters); err != nil {
			return err
		}
	}

	fmt.Printf("seedguild: applied %q to guild %d (%s)\n", tag, guildID, guild.Name)
	fmt.Printf("  roster: %d characters across %d accounts (%d officers, %d unverified)\n",
		len(plan.Characters), len(accountUsers), plan.officerCount(), plan.unverifiedCount())
	fmt.Printf("  raid nights: %d reports, %d fights\n", len(plan.Reports), plan.fightCount())
	fmt.Println("  ratings: not written. rating_scores and fight_metrics.execution_score are only ever " +
		"produced from a stored fight summary object in R2 (api/internal/rating.Backfill reads it via " +
		"Store.Summaries; see api/internal/rating/backfill.go's readSummary) - this seed writes fight_metrics " +
		"rows directly and never ran a real ingest, so no such object exists for these fights. Running " +
		"`rating-backfill` against them would find the stale fights, fail to read their summary, log the " +
		"error and skip every one - it would not write anything, honest or otherwise. Fabricating a summary " +
		"object would mean fabricating a fake combat log, which this tool does not do. See README.md's " +
		"Readiness/Ratings section.")
	return nil
}

func writeMockCharacter(ctx context.Context, pool *pgxpool.Pool, tag string, guild guildInfo, pc plannedCharacter, accountUsers map[string]int64) error {
	dir, err := bisDir()
	if err != nil {
		return err
	}
	slots, err := loadBisSlots(dir, pc.Mock.BisFile)
	if err != nil {
		return err
	}
	export := buildFS1(classSlugFor(pc.Mock.Class), pc.Mock.RaceSlug, talentStringFor(pc.Mock),
		gearWithReadiness(pc.Mock, slots), professionsFor(pc.Mock.Class), bagsFor(pc.Mock))

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: character %s: begin: %w", pc.Key, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	userID, alreadyHasAccount := accountUsers[pc.AccountKey]
	if !alreadyHasAccount {
		userID, err = insertMockUser(ctx, tx, tag, pc.Battletag)
		if err != nil {
			return err
		}
	}
	if err := insertMockCharacter(ctx, tx, guild, userID, pc, export); err != nil {
		return err
	}
	if err := upsertMockMembership(ctx, tx, guild.ID, userID, pc); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("seedguild: character %s: commit: %w", pc.Key, err)
	}
	accountUsers[pc.AccountKey] = userID
	return nil
}

func writeMockReport(ctx context.Context, pool *pgxpool.Pool, tag string, guildID, ownerID int64, ownerCharKey string, pr plannedReport, chars []plannedCharacter) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: report %s: begin: %w", pr.ID, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if err := recordRow(ctx, tx, tag, "reports", pr.ID, nil); err != nil {
		return err
	}
	if err := insertReport(ctx, tx, guildID, ownerID, ownerCharKey, pr); err != nil {
		return err
	}
	if err := insertFightsAndMetrics(ctx, tx, pool, pr, attendees(chars, pr.Plan.Tag)); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("seedguild: report %s: commit: %w", pr.ID, err)
	}
	return nil
}

// attendees is chars filtered to those present at the raid night tagged reportTag
// (roster.go's SkipReports/attends).
func attendees(chars []plannedCharacter, reportTag string) []plannedCharacter {
	out := make([]plannedCharacter, 0, len(chars))
	for _, pc := range chars {
		if pc.Mock.attends(reportTag) {
			out = append(out, pc)
		}
	}
	return out
}

// classSlugFor is identity today (mockCharacter.Class is already the FS1 class slug);
// kept as its own function so a future class-name mismatch has one place to fix.
func classSlugFor(class string) string { return class }
