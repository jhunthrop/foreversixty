// api/cmd/api/grant_rank.go
//
// grant-rank/revoke-rank are the operator's time-boxed guild rank
// override (guild_rank_overrides), a CLI beside grant/revoke and never an
// HTTP endpoint. The in-game rank still drives guild_characters on every
// sync; the override only raises guild_members.rank while unexpired.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

const rankOverrideUsage = `usage:
  api grant-rank --guild <id> --user <id|battletag> --rank officer|leader --until YYYY-MM-DD --note "testing guild tools"
  api revoke-rank --guild <id> --user <id|battletag>
--until is exclusive: the override ends at 00:00 UTC on that date.`

// rankFlags is grant-rank/revoke-rank's flag set, parsed in grant.go's style.
type rankFlags struct {
	guild, user, rank, until, note string
}

func parseRankFlags(args []string) (rankFlags, error) {
	var f rankFlags
	for i := 0; i < len(args); i++ {
		arg := args[i]
		next := func() (string, error) {
			i++
			if i >= len(args) {
				return "", fmt.Errorf("%s needs a value", arg)
			}
			return args[i], nil
		}
		var err error
		switch arg {
		case "--guild":
			f.guild, err = next()
		case "--user":
			f.user, err = next()
		case "--rank":
			f.rank, err = next()
		case "--until":
			f.until, err = next()
		case "--note":
			f.note, err = next()
		default:
			err = fmt.Errorf("unknown flag %q", arg)
		}
		if err != nil {
			return rankFlags{}, err
		}
	}
	if f.guild == "" || f.user == "" {
		return rankFlags{}, fmt.Errorf("--guild and --user are required")
	}
	return f, nil
}

// grantRankOverride validates the grant-only flags into an override,
// relative to now.
func (f rankFlags) grantRankOverride(guildID, userID int64, operator string, now time.Time) (guilds.RankOverride, error) {
	if !guilds.ValidRankOverride(f.rank) {
		return guilds.RankOverride{}, fmt.Errorf("--rank must be %s or %s", guilds.RankOfficer, guilds.RankLeader)
	}
	if f.until == "" {
		return guilds.RankOverride{}, fmt.Errorf("--until is required")
	}
	until, err := parseUntil(f.until)
	if err != nil {
		return guilds.RankOverride{}, err
	}
	if err := requireFutureUntil(*until, now); err != nil {
		return guilds.RankOverride{}, err
	}
	if f.note == "" {
		return guilds.RankOverride{}, fmt.Errorf("--note is required")
	}
	return guilds.RankOverride{
		GuildID: guildID, UserID: userID, Rank: f.rank,
		ExpiresAt: *until, Note: f.note, CreatedBy: operator,
	}, nil
}

// resolveRankSubject resolves --guild (which must exist) and --user.
func resolveRankSubject(ctx context.Context, pool *pgxpool.Pool, f rankFlags) (int64, int64, error) {
	guildID, err := strconv.ParseInt(f.guild, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("--guild must be a numeric id: %w", err)
	}
	var exists bool
	if err := pool.QueryRow(ctx, `select exists(select 1 from guilds where id = $1)`, guildID).Scan(&exists); err != nil {
		return 0, 0, fmt.Errorf("look up guild %d: %w", guildID, err)
	}
	if !exists {
		return 0, 0, fmt.Errorf("no guild with id %d", guildID)
	}
	userID, err := resolveUserRef(ctx, pool, f.user)
	if err != nil {
		return 0, 0, err
	}
	return guildID, userID, nil
}

func runGrantRank(ctx context.Context, log *slog.Logger, args []string) error {
	f, err := parseRankFlags(args)
	if err != nil {
		return fmt.Errorf("%w\n%s", err, rankOverrideUsage)
	}
	if _, err := f.grantRankOverride(0, 0, "", time.Now()); err != nil {
		return fmt.Errorf("%w\n%s", err, rankOverrideUsage)
	}
	_, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	guildID, userID, err := resolveRankSubject(ctx, pool, f)
	if err != nil {
		return fmt.Errorf("grant-rank: %w", err)
	}
	operator := operatorTag()
	o, err := f.grantRankOverride(guildID, userID, operator, time.Now())
	if err != nil {
		return fmt.Errorf("grant-rank: %w", err)
	}
	res, err := (&guilds.Store{Pool: pool}).GrantRankOverride(ctx, o)
	if err != nil {
		return fmt.Errorf("grant-rank: %w", err)
	}
	log.Info("grant-rank", "guild_id", guildID, "user_id", userID, "rank", o.Rank,
		"expires_at", o.ExpiresAt, "note", o.Note, "operator", operator,
		"member_rank", res.Rank, "verified", res.Verified)
	if res.Rank == "" {
		log.Warn("grant-rank", "guild_id", guildID, "user_id", userID,
			"msg", "account has no character in this guild; the override applies once one syncs")
	}
	return nil
}

func runRevokeRank(ctx context.Context, log *slog.Logger, args []string) error {
	f, err := parseRankFlags(args)
	if err != nil {
		return fmt.Errorf("%w\n%s", err, rankOverrideUsage)
	}
	_, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	guildID, userID, err := resolveRankSubject(ctx, pool, f)
	if err != nil {
		return fmt.Errorf("revoke-rank: %w", err)
	}
	res, removed, err := (&guilds.Store{Pool: pool}).RevokeRankOverride(ctx, guildID, userID)
	if err != nil {
		return fmt.Errorf("revoke-rank: %w", err)
	}
	if !removed {
		return fmt.Errorf("revoke-rank: guild %d user %d has no rank override (membership recomputed anyway)", guildID, userID)
	}
	log.Info("revoke-rank", "guild_id", guildID, "user_id", userID, "operator", operatorTag(),
		"member_rank", res.Rank, "verified", res.Verified)
	return nil
}
