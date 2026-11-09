// Command repairreports re-parses the uploaded reports the 2026-11-08 engine
// fixes change: a report with a boss fight recorded as a wipe at 0% health
// (Forever's ENCOUNTER_END was refused, so every kill read as a wipe), or an
// upload credited to nobody (the parse job now credits the log's own
// writer). The re-parse rewrites fights, rankings rows and report files,
// credits the logger and re-runs guild matching; it sends no Discord or
// push notice again.
//
// Dry run by default: it lists what it would re-parse and writes nothing.
// -apply re-parses. Idempotent: a re-parsed report no longer matches, apart
// from an upload whose log names no writer, which re-parses to the same
// result. It never migrates: it connects with DATABASE_URL and nothing
// else touches the schema.
//
//	DATABASE_URL=... R2_ACCOUNT_ID=... R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... \
//	  R2_BUCKET=foreversixty-logs TREE_DATA_DIR=../data/builds go run ./cmd/repairreports [-apply]
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/parse"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/spec"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// candidatesSQL is every uploaded report the fixes change.
const candidatesSQL = `
select r.id,
       r.logging_character is null,
       (select count(*) from fights f where f.report_id = r.id
          and f.encounter_id is not null and not f.kill and f.boss_health_pct = 0)
  from reports r
 where r.upload_id is not null and r.status = 'complete'
   and (r.logging_character is null
        or exists (select 1 from fights f where f.report_id = r.id
                     and f.encounter_id is not null and not f.kill and f.boss_health_pct = 0))
 order by r.created_at`

type candidate struct {
	id        string
	noLogger  bool
	zeroWipes int
}

func main() {
	apply := flag.Bool("apply", false, "re-parse; without it, only list")
	flag.Parse()
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))
	if err := run(context.Background(), *apply, log); err != nil {
		log.Error("repairreports", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, apply bool, log *slog.Logger) error {
	pool, err := db.Connect(ctx, mustEnv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer pool.Close()
	cands, err := candidates(ctx, pool)
	if err != nil {
		return err
	}
	for _, c := range cands {
		fmt.Printf("report %s: no logger=%v, wipes at 0%%=%d\n", c.id, c.noLogger, c.zeroWipes)
	}
	fmt.Printf("%d report(s) to re-parse\n", len(cands))
	if !apply || len(cands) == 0 {
		return nil
	}
	deps, err := parseDeps(pool, log)
	if err != nil {
		return err
	}
	for _, c := range cands {
		if err := reparse(ctx, pool, deps, c.id); err != nil {
			return err
		}
		fmt.Printf("re-parsed %s\n", c.id)
	}
	return nil
}

func candidates(ctx context.Context, pool *pgxpool.Pool) ([]candidate, error) {
	rows, err := pool.Query(ctx, candidatesSQL)
	if err != nil {
		return nil, fmt.Errorf("candidates: %w", err)
	}
	defer rows.Close()
	var out []candidate
	for rows.Next() {
		var c candidate
		if err := rows.Scan(&c.id, &c.noLogger, &c.zeroWipes); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// reparse withdraws the report's fights and rankings rows (a fight that was
// a bogus encounter is trash now, and its indexes shift, so nothing else
// would take the old rows back; no table references fights), then runs the
// parse job, which writes them all again.
func reparse(ctx context.Context, pool *pgxpool.Pool, deps parse.Deps, id string) error {
	for _, q := range []string{
		`delete from fight_metrics where report_id = $1`,
		`delete from coach_profiles where report_id = $1`,
		`delete from fights where report_id = $1`,
	} {
		if _, err := pool.Exec(ctx, q, id); err != nil {
			return fmt.Errorf("withdraw %s: %w", id, err)
		}
	}
	if err := parse.Report(ctx, deps, id); err != nil {
		return fmt.Errorf("re-parse %s: %w", id, err)
	}
	return nil
}

// parseDeps is the parse job's wiring without its notifier: the guild
// matcher is told of the completed report, Discord and push are not.
func parseDeps(pool *pgxpool.Pool, log *slog.Logger) (parse.Deps, error) {
	objects, err := r2.New(r2.Config{
		AccountID: mustEnv("R2_ACCOUNT_ID"), AccessKeyID: mustEnv("R2_ACCESS_KEY_ID"),
		SecretAccessKey: mustEnv("R2_SECRET_ACCESS_KEY"), Bucket: mustEnv("R2_BUCKET"),
	})
	if err != nil {
		return parse.Deps{}, err
	}
	data, err := trees.Load(mustEnv("TREE_DATA_DIR"))
	if err != nil {
		return parse.Deps{}, err
	}
	var specs *spec.Inferrer
	if b, ok := data.Active(); ok {
		specs = spec.New(b)
	} else {
		specs = spec.New(nil)
	}
	return parse.Deps{
		Reports: &reports.Store{Pool: pool}, Objects: objects, Log: log,
		Rank:      &rankings.Store{Pool: pool, Specs: specs},
		Completed: &guilds.Service{Store: &guilds.Store{Pool: pool}, PublicBaseURL: "https://foreversixty.gg", Log: log},
	}, nil
}

func mustEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		fmt.Fprintf(os.Stderr, "repairreports: %s is not set\n", name)
		os.Exit(2)
	}
	return v
}
