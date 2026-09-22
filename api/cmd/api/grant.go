// api/cmd/api/grant.go
//
// grant/revoke are the CLI, not an HTTP endpoint (spec RULING 4): a
// money-adjacent write path run by hand against a DATABASE_URL, never
// exposed on the public internet.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/user"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

const grantRevokeUsage = `usage:
  api grant --user <id|battletag> --plan premium [--until 2026-12-31] --note "beta tester"
  api grant --guild <id> --plan guild [--until 2026-12-31] --note "beta guild"
  api revoke --user <id|battletag> --plan premium --note "requested by support"
  api revoke --guild <id> --plan guild --note "requested by support"`

// grantFlags is grant/revoke's small, manual flag set — no external flag
// package is used elsewhere in this binary's subcommands.
type grantFlags struct {
	user, guild, plan, until, note string
}

func parseGrantFlags(args []string) (grantFlags, error) {
	var f grantFlags
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
		case "--user":
			f.user, err = next()
		case "--guild":
			f.guild, err = next()
		case "--plan":
			f.plan, err = next()
		case "--until":
			f.until, err = next()
		case "--note":
			f.note, err = next()
		default:
			err = fmt.Errorf("unknown flag %q", arg)
		}
		if err != nil {
			return grantFlags{}, err
		}
	}
	if (f.user == "") == (f.guild == "") {
		return grantFlags{}, fmt.Errorf("exactly one of --user or --guild is required")
	}
	if f.plan != entitlements.PlanPremium && f.plan != entitlements.PlanGuild {
		return grantFlags{}, fmt.Errorf("--plan must be premium or guild")
	}
	if f.note == "" {
		return grantFlags{}, fmt.Errorf("--note is required")
	}
	return f, nil
}

// parseUntil parses --until as an ISO date (UTC midnight), nil for an
// open-ended grant.
func parseUntil(s string) (*time.Time, error) {
	if s == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return nil, fmt.Errorf("--until must be YYYY-MM-DD: %w", err)
	}
	t = t.UTC()
	return &t, nil
}

// resolveUserRef accepts a numeric id or a battletag, refusing ambiguity
// (spec §1.5: "refusing ambiguity (ties) or a miss with a clear message
// rather than guessing" — battletag has no database-level uniqueness
// constraint, so this is a real, reachable case).
func resolveUserRef(ctx context.Context, pool *pgxpool.Pool, ref string) (int64, error) {
	if id, err := strconv.ParseInt(ref, 10, 64); err == nil {
		var exists bool
		if err := pool.QueryRow(ctx, `select exists(select 1 from users where id = $1)`, id).Scan(&exists); err != nil {
			return 0, fmt.Errorf("grant: look up user id %d: %w", id, err)
		}
		if !exists {
			return 0, fmt.Errorf("grant: no account with id %d", id)
		}
		return id, nil
	}
	rows, err := pool.Query(ctx, `select id from users where battletag = $1`, ref)
	if err != nil {
		return 0, fmt.Errorf("grant: look up battletag %q: %w", ref, err)
	}
	defer rows.Close()
	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("grant: look up battletag %q: %w", ref, err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return 0, fmt.Errorf("grant: look up battletag %q: %w", ref, err)
	}
	switch len(ids) {
	case 0:
		return 0, fmt.Errorf("grant: no account with battletag %q", ref)
	case 1:
		return ids[0], nil
	default:
		return 0, fmt.Errorf("grant: battletag %q is ambiguous (%d accounts); use the numeric id instead", ref, len(ids))
	}
}

func resolveGrantSubject(ctx context.Context, pool *pgxpool.Pool, f grantFlags) (entitlements.Subject, error) {
	if f.guild != "" {
		gid, err := strconv.ParseInt(f.guild, 10, 64)
		if err != nil {
			return entitlements.Subject{}, fmt.Errorf("grant: --guild must be a numeric id: %w", err)
		}
		if f.plan != entitlements.PlanGuild {
			return entitlements.Subject{}, fmt.Errorf("grant: --guild requires --plan guild")
		}
		return entitlements.Subject{GuildID: &gid}, nil
	}
	if f.plan != entitlements.PlanPremium {
		return entitlements.Subject{}, fmt.Errorf("grant: --user requires --plan premium")
	}
	uid, err := resolveUserRef(ctx, pool, f.user)
	if err != nil {
		return entitlements.Subject{}, err
	}
	return entitlements.Subject{UserID: &uid}, nil
}

// operatorTag names whoever is running this command, for the audit
// trail (spec §1.5). See Ruling G for why this never becomes
// entitlements.granted_by.
func operatorTag() string {
	if u, err := user.Current(); err == nil && u.Username != "" {
		return u.Username
	}
	if v := os.Getenv("USER"); v != "" {
		return v
	}
	return "unknown"
}

func runGrant(ctx context.Context, log *slog.Logger, args []string) error {
	f, err := parseGrantFlags(args)
	if err != nil {
		return fmt.Errorf("%w\n%s", err, grantRevokeUsage)
	}
	until, err := parseUntil(f.until)
	if err != nil {
		return err
	}
	_, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	subj, err := resolveGrantSubject(ctx, pool, f)
	if err != nil {
		return err
	}
	operator := operatorTag()
	note := fmt.Sprintf("%s: %s", operator, f.note)
	actor := fmt.Sprintf("cli_grant:%s %s", operator, f.note)
	if err := (&entitlements.Store{Pool: pool}).Grant(ctx, subj, f.plan, until, nil, note, actor); err != nil {
		return fmt.Errorf("grant: %w", err)
	}
	log.Info("grant", "plan", f.plan, "operator", operator)
	return nil
}

func runRevoke(ctx context.Context, log *slog.Logger, args []string) error {
	f, err := parseGrantFlags(args)
	if err != nil {
		return fmt.Errorf("%w\n%s", err, grantRevokeUsage)
	}
	_, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	subj, err := resolveGrantSubject(ctx, pool, f)
	if err != nil {
		return err
	}
	operator := operatorTag()
	actor := fmt.Sprintf("cli_revoke:%s %s", operator, f.note)
	if err := (&entitlements.Store{Pool: pool}).Revoke(ctx, subj, f.plan, actor); err != nil {
		return fmt.Errorf("revoke: %w", err)
	}
	log.Info("revoke", "plan", f.plan, "operator", operator)
	return nil
}
