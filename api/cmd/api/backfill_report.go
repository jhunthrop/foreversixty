// api/cmd/api/backfill_report.go
//
// backfill-report is the operator's repair for a live report whose fight bundles the API
// refused while its raw chunks arrived (the 10 Nov 2026 partBytes bug): it parses the raw
// chunks again through the live path's own fight writer. A CLI beside grant/revoke, never an
// HTTP endpoint, run by hand against a DATABASE_URL (docs/runbooks/report-backfill.md).
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/parse"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/logs/engine/layout"
)

const backfillReportUsage = `usage:
  api backfill-report --report <id>                       # dry run: stream the raw chunks, count the fights, write nothing
  api backfill-report --report <id> --apply               # write the fights, complete the report; never announces
  api backfill-report --report <id> --layout retail-v22   # parse with a named layout instead of the report's (inferred) one`

// backfillFlags is backfill-report's flag set, parsed in grant.go's style.
type backfillFlags struct {
	report, layout string
	apply          bool
}

func parseBackfillFlags(args []string) (backfillFlags, error) {
	var f backfillFlags
	for i := 0; i < len(args); i++ {
		switch arg := args[i]; arg {
		case "--report":
			i++
			if i >= len(args) {
				return backfillFlags{}, fmt.Errorf("%s needs a value", arg)
			}
			f.report = args[i]
		case "--layout":
			i++
			if i >= len(args) {
				return backfillFlags{}, fmt.Errorf("%s needs a value", arg)
			}
			f.layout = args[i]
		case "--apply":
			f.apply = true
		default:
			return backfillFlags{}, fmt.Errorf("unknown flag %q", arg)
		}
	}
	if f.report == "" {
		return backfillFlags{}, fmt.Errorf("--report is required: one report id, never a sweep")
	}
	if f.layout != "" {
		if _, ok := engine.LayoutByName(f.layout); !ok {
			return backfillFlags{}, fmt.Errorf("--layout %q is not a layout the parser knows; one of %s", f.layout, strings.Join(layoutNames(), ", "))
		}
	}
	return f, nil
}

// layoutNames lists the parser's registered layouts, for the --layout error.
func layoutNames() []string {
	rows := layout.Rows()
	names := make([]string, 0, len(rows))
	for _, r := range rows {
		names = append(names, r.Name)
	}
	return names
}

// connect loads the configuration and connects without migrating. Only the deployed API
// applies migrations to production (the 2026-10-04 incident, migration 0029); a repair tool
// run from a checkout ahead of the deploy must not. The schema it reads is the deployed
// one, so the tool is still run from a worktree at the deployed image's sha.
func connect(ctx context.Context) (config.Config, *pgxpool.Pool, error) {
	cfg, err := config.Load(os.Getenv)
	if err != nil {
		return config.Config{}, nil, err
	}
	pool, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return config.Config{}, nil, err
	}
	return cfg, pool, nil
}

func runBackfillReport(ctx context.Context, log *slog.Logger, args []string) error {
	f, err := parseBackfillFlags(args)
	if err != nil {
		return fmt.Errorf("%w\n%s", err, backfillReportUsage)
	}
	cfg, pool, err := connect(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	client := objects(cfg, log)
	if client == nil {
		return fmt.Errorf("backfill-report needs R2 credentials")
	}
	treeData, err := trees.Load(cfg.TreeDataDir)
	if err != nil {
		return err
	}
	deps := newBackfillDeps(cfg, pool, client, treeData, log)
	began := time.Now()
	log.Info("backfill-report", "report", f.report, "apply", f.apply, "layout", f.layout, "operator", operatorTag())
	res, err := parse.Backfill(ctx, deps, f.report, parse.BackfillOptions{Apply: f.apply, Layout: f.layout})
	parse.LogBackfill(log, res)
	if err != nil {
		return err
	}
	log.Info("backfill-report", "report", f.report, "done", true, "ms", time.Since(began).Milliseconds())
	return nil
}
