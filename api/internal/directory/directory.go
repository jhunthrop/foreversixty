// api/internal/directory/directory.go

// Package directory is the character directory: every Forever character the site can learn
// about before its player signs in, keyed by region and Blizzard character id (spec
// docs/superpowers/specs/2026-11-14-character-directory-discovery.md §6). It grows from guild
// rosters, sign-ins, addon exports and logged names, and is enriched from Blizzard's public
// profile endpoints with an app token only. It keeps public game data, never a BattleTag or
// anything behind a user token, and it honours Blizzard's API terms: a paced call budget,
// a /status validation and a refresh or delete within 30 days of every row.
package directory

import (
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
)

const (
	// DefaultCallsPerHour is the directory's share of the app's 36,000 calls an hour (terms
	// section 2(r)); the rest stays for sign-ins and the nightly jobs.
	DefaultCallsPerHour = 20_000
	// rosterAfter is how often a guild's roster is read.
	rosterAfter = 24 * time.Hour
	// profileAfter is how old a character's profile read may grow before the next pass
	// reads it again: well inside the 30-day bound, so a missed pass never breaks it.
	profileAfter = 72 * time.Hour
	// staleAfter is Blizzard's bound: a row no source confirmed for this long is deleted.
	staleAfter = 30 * 24 * time.Hour
	// missAfter is how long a name that answered 404 is left before a seed asks again.
	missAfter = 7 * 24 * time.Hour
	// defaultWorkers is how many characters a pass reads at once; the budget paces them.
	defaultWorkers = 4
	// batchSize is how many due rows a pass reads from the table at a time.
	batchSize = 200
	// retryAfter is how long a row whose read failed waits before the next try.
	retryAfter = time.Hour
	// maxGuildsPerPass bounds the guilds one pass reads.
	maxGuildsPerPass = 5_000
	// nameSeedsPerPass bounds the name lookups (addon, log and ratings names) one pass makes.
	nameSeedsPerPass = 500
)

// Sources a directory row can name, in its sources column.
const (
	SourceRoster = "roster"
	SourceSignIn = "signin"
	SourceAddon  = "addon"
	SourceLog    = "log"
	SourceRating = "rating"
)

// Service fills and refreshes the directory.
type Service struct {
	Pool    *pgxpool.Pool
	Client  *bnetapi.Client
	Regions []string
	Log     *slog.Logger
	// Budget paces every Blizzard call the directory makes; nil gets DefaultCallsPerHour.
	Budget *Budget
	// Workers is how many characters a pass reads at once; zero gets defaultWorkers.
	Workers int

	now func() time.Time
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) budget() *Budget {
	if s.Budget == nil {
		s.Budget = NewBudget(DefaultCallsPerHour)
	}
	return s.Budget
}

func (s *Service) workers() int {
	if s.Workers > 0 {
		return s.Workers
	}
	return defaultWorkers
}

func (s *Service) clock() time.Time {
	if s.now != nil {
		return s.now()
	}
	return time.Now()
}
