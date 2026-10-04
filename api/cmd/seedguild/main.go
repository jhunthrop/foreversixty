// api/cmd/seedguild/main.go
//
// seedguild populates ONE real guild with a realistic, clearly-tagged mock roster, raid
// nights and fight metrics, so the guild page can be designed and built against a full
// page rather than an empty one. Every row it writes is recorded in the seed_rows table
// (migration 0029) under a tag derived from the guild id, so --remove can find and delete
// exactly what --apply wrote, and restore exactly what it overwrote on the owner's own
// rows and the guild's claim state.
//
//	go run ./cmd/seedguild --guild 2 --owner-battletag 'hunthrop#1894'              # apply
//	go run ./cmd/seedguild --guild 2 --owner-battletag 'hunthrop#1894' --dry-run    # preview
//	go run ./cmd/seedguild --guild 2 --owner-battletag 'hunthrop#1894' --remove     # undo
//
// DATABASE_URL is read from the environment and never printed or logged.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "seedguild:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("seedguild", flag.ContinueOnError)
	guildID := fs.Int64("guild", 0, "the guild id to seed (required)")
	ownerBattletag := fs.String("owner-battletag", "", "the owning account's battletag (required)")
	remove := fs.Bool("remove", false, "remove a previously applied seed instead of applying one")
	dryRun := fs.Bool("dry-run", false, "print what would be written (or removed) without writing anything")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *guildID <= 0 {
		return fmt.Errorf("--guild is required and must be positive")
	}
	if *ownerBattletag == "" {
		return fmt.Errorf("--owner-battletag is required")
	}

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}
	if err := db.Migrate(url); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	pool, err := db.Connect(ctx, url)
	if err != nil {
		return err
	}
	defer pool.Close()

	if *remove {
		if *dryRun {
			return PrintRemoveDryRun(ctx, pool, *guildID)
		}
		return RemoveSeed(ctx, pool, *guildID, *ownerBattletag)
	}

	if *dryRun {
		guild, err := loadGuild(ctx, pool, *guildID)
		if err != nil {
			return err
		}
		PrintDryRun(buildSeedPlan(guild, time.Now().UTC()))
		return nil
	}
	return ApplySeed(ctx, pool, *guildID, *ownerBattletag)
}
