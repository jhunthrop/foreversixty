// Command repairrulesets re-keys every Battle.net-imported character under
// the ruleset its own profile names, fetching the public profile where none
// is stored, and moves each guild the import created under a guessed
// ruleset to its members' ruleset (bnetimport.RepairRulesets). From
// 2026-11-05 to 2026-11-09 the import keyed every new sign-up as normal.
//
// Dry run by default: the work runs in one transaction that is rolled
// back, so the summary is exact and nothing is written. -apply commits it.
// Idempotent: a second run finds nothing to move. It never migrates: it
// connects with DATABASE_URL and nothing else touches the schema. It
// prints keys and ids only.
//
//	DATABASE_URL=... BNET_CLIENT_ID=... BNET_CLIENT_SECRET=... go run ./cmd/repairrulesets [-apply]
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/bnetimport"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func main() {
	apply := flag.Bool("apply", false, "commit; without it the run is rolled back")
	flag.Parse()
	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelWarn}))
	if err := run(context.Background(), *apply, log); err != nil {
		log.Error("repairrulesets", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, apply bool, log *slog.Logger) error {
	pool, err := db.Connect(ctx, mustEnv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer pool.Close()
	svc := &bnetimport.Service{Pool: pool, Log: log, Client: bnetapi.New(bnetapi.Config{
		ClientID: mustEnv("BNET_CLIENT_ID"), ClientSecret: mustEnv("BNET_CLIENT_SECRET"),
		Game: "forever", Regions: []string{"us", "eu", "tw"}, Log: log,
	})}
	sum, err := svc.RepairRulesets(ctx, apply)
	if err != nil {
		return err
	}
	print(sum, apply)
	return nil
}

func print(sum bnetimport.RepairSummary, apply bool) {
	mode := "dry run (rolled back)"
	if apply {
		mode = "applied"
	}
	fmt.Printf("%s: %d Battle.net characters checked, %d profiles fetched\n", mode, sum.Checked, sum.Refetched)
	for _, m := range sum.Moves {
		fmt.Printf("  move user %d: %s -> %s\n", m.UserID, m.OldKey, m.NewKey)
	}
	for _, g := range sum.Guilds {
		fmt.Printf("  guild %d: %s -> %s\n", g.GuildID, g.From, g.To)
	}
	tables := make([]string, 0, len(sum.Rows))
	for t := range sum.Rows {
		tables = append(tables, t)
	}
	sort.Strings(tables)
	for _, t := range tables {
		fmt.Printf("  rows %s: %d\n", t, sum.Rows[t])
	}
	users := map[int64]bool{}
	for _, u := range sum.Unresolved {
		users[u.UserID] = true
	}
	keys := make([]string, 0, len(sum.Unresolved))
	for _, u := range sum.Unresolved {
		keys = append(keys, u.Key)
	}
	fmt.Printf("%d moved, %d guilds moved, %d unresolved on %d accounts (profile not public): %s\n",
		len(sum.Moves), len(sum.Guilds), len(sum.Unresolved), len(users), strings.Join(keys, " "))
}

func mustEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		fmt.Fprintf(os.Stderr, "repairrulesets: %s is not set\n", name)
		os.Exit(2)
	}
	return v
}
