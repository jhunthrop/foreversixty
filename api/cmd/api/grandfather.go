// api/cmd/api/grandfather.go
//
// grandfather-officer-tools is the one-off operator command that gives every
// guild claimed before the officer-tools gate a time-boxed scoped grant
// (design 2026-10-25 §4.6). Dry run unless --apply.
package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

const grandfatherUsage = `usage:
  api grandfather-officer-tools --until 2026-12-09 --note "launch grandfathering" [--apply]`

const dateLayout = "2006-01-02"

type grandfatherOptions struct {
	Until    time.Time
	Note     string
	Operator string
	Apply    bool
}

func parseGrandfatherFlags(args []string) (grandfatherOptions, error) {
	var opts grandfatherOptions
	var until string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--apply" {
			opts.Apply = true
			continue
		}
		if arg != "--until" && arg != "--note" {
			return grandfatherOptions{}, fmt.Errorf("unknown flag %q", arg)
		}
		i++
		if i >= len(args) {
			return grandfatherOptions{}, fmt.Errorf("%s needs a value", arg)
		}
		if arg == "--until" {
			until = args[i]
		} else {
			opts.Note = args[i]
		}
	}
	if until == "" {
		return grandfatherOptions{}, fmt.Errorf("--until is required")
	}
	if opts.Note == "" {
		return grandfatherOptions{}, fmt.Errorf("--note is required")
	}
	t, err := parseUntil(until)
	if err != nil {
		return grandfatherOptions{}, err
	}
	opts.Until = *t
	return opts, nil
}

func runGrandfatherOfficerTools(ctx context.Context, log *slog.Logger, args []string) error {
	opts, err := parseGrandfatherFlags(args)
	if err != nil {
		return fmt.Errorf("%w\n%s", err, grandfatherUsage)
	}
	opts.Operator = operatorTag()
	_, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	if err := grandfatherOfficerTools(ctx, pool, opts, os.Stdout); err != nil {
		return err
	}
	log.Info("grandfather-officer-tools", "apply", opts.Apply, "operator", opts.Operator)
	return nil
}

// grandfatherOfficerTools grants every claimed guild without a live plan a
// scoped officer_views grant and reports one line per guild, then totals.
func grandfatherOfficerTools(ctx context.Context, pool *pgxpool.Pool, opts grandfatherOptions, out io.Writer) error {
	store := &entitlements.Store{Pool: pool}
	todo, err := store.ClaimedGuildsWithoutActivePlan(ctx)
	if err != nil {
		return fmt.Errorf("grandfather: %w", err)
	}
	planned, err := store.ClaimedGuildsWithActivePlan(ctx)
	if err != nil {
		return fmt.Errorf("grandfather: %w", err)
	}
	granted := 0
	for _, g := range todo {
		if err := grandfatherOne(ctx, store, g, opts, out); err != nil {
			return err
		}
		granted++
	}
	skipped, err := reportSkipped(ctx, store, planned, out)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(out, "%d %s, %d skipped\n", granted, grantedVerb(opts.Apply), skipped)
	return err
}

func grantedVerb(apply bool) string {
	if apply {
		return "granted"
	}
	return "would be granted"
}

func grandfatherOne(ctx context.Context, store *entitlements.Store, g entitlements.GuildRef, opts grandfatherOptions, out io.Writer) error {
	verb := "would grant"
	if opts.Apply {
		gid := g.ID
		note := fmt.Sprintf("%s: %s", opts.Operator, opts.Note)
		actor := fmt.Sprintf("cli_grandfather:%s %s", opts.Operator, opts.Note)
		until := opts.Until
		err := store.Grant(ctx, entitlements.Subject{GuildID: &gid}, entitlements.PlanGuild, &until, nil, note, actor,
			[]string{entitlements.FeatureScopeOfficerViews})
		if err != nil {
			return fmt.Errorf("grandfather: guild %d: %w", g.ID, err)
		}
		verb = "granted"
	}
	_, err := fmt.Fprintf(out, "%s %d %s until %s\n", verb, g.ID, g.Name, opts.Until.Format(dateLayout))
	return err
}

// reportSkipped prints the claimed guilds that held a live plan before the run
// and returns how many there were.
func reportSkipped(ctx context.Context, store *entitlements.Store, planned []entitlements.GuildRef, out io.Writer) (int, error) {
	for _, g := range planned {
		billing, err := store.GuildBilling(ctx, g.ID)
		if err != nil {
			return 0, fmt.Errorf("grandfather: %w", err)
		}
		reason := "on the plan"
		if billing != nil && billing.Scoped() {
			reason = "already granted"
		}
		if _, err := fmt.Fprintf(out, "skipped %d %s: %s\n", g.ID, g.Name, reason); err != nil {
			return 0, err
		}
	}
	return len(planned), nil
}
