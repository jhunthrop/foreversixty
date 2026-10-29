// api/cmd/api/jobs_extra.go
package main

import (
	"context"
	"log/slog"
	"slices"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
)

// jobCommand is one Cloud Run job the image runs: the first container argument names it.
type jobCommand struct {
	name string
	run  func(ctx context.Context, log *slog.Logger, args []string) error
}

// extraJobs is the table of job commands kept out of main's switch (main.go is at its line
// budget): a new command is a row here, never another case.
var extraJobs = append([]jobCommand{
	{rating.RebuildPreparationJobCommand, runRebuildPreparation},
}, coachJobs...)

// setupCommands are run by hand from a checkout, never as Cloud Run jobs, so the deploy does
// not move them; main runs them through the same table runner.
var setupCommands = []jobCommand{
	{discordEmojiCommand, runDiscordEmojiCommand},
}

// allCommands is every table-run command: the jobs, then the setup commands.
func allCommands() []jobCommand { return slices.Concat(extraJobs, setupCommands) }

// runJobTable runs the row of table named by args[0] with args[1:], then calls exit(0), or
// exit(1) after one log.Error(name, "err", err) line, as main's own cases do. It returns,
// calling nothing, when args is empty or names no row.
func runJobTable(log *slog.Logger, table []jobCommand, args []string, exit func(int)) {
	if len(args) == 0 {
		return
	}
	for _, job := range table {
		if job.name != args[0] {
			continue
		}
		if err := job.run(context.Background(), log, args[1:]); err != nil {
			log.Error(job.name, "err", err)
			exit(1)
			return
		}
		exit(0)
		return
	}
}
