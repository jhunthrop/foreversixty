// api/internal/dataaddon/job.go
package dataaddon

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"strings"
	"time"
)

// Logger is the subset of *slog.Logger this package needs -- *slog.Logger
// satisfies it structurally, with no adapter, so api/cmd/api/main.go can
// pass its own logger straight through.
type Logger interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

// Deps is everything one run of the job needs.
type Deps struct {
	Store  *Store
	Upload Uploader // nil (or an empty Bucket) means this deployment has no bucket; the run still succeeds and logs why the upload was skipped
	Bucket string
	Build  string // the data build the ratings were computed against
	Now    time.Time
	Log    Logger
}

func (d Deps) logger() Logger {
	if d.Log != nil {
		return d.Log
	}
	return noopLogger{}
}

type noopLogger struct{}

func (noopLogger) Info(string, ...any)  {}
func (noopLogger) Warn(string, ...any)  {}
func (noopLogger) Error(string, ...any) {}

// Result is what one run produced.
type Result struct {
	Characters  int               // rows written
	Guilds      int               // rows written
	SkippedRows int               // rows a bad player_key could not be formatted for
	Files       map[string]int    // filename -> byte size (the dispatch's "size line")
	Rendered    map[string][]byte // filename -> content, for a caller (the golden test) that needs the bytes themselves
}

// objectKeyPrefix is where every published file (and the manifest) lives
// inside the bucket.
const objectKeyPrefix = "data-addon/"

// Run reads the database, aggregates, renders Data.lua (or its region-split
// form), and uploads it, in that order. It never fails the whole run over
// one bad row (dispatch: "skipped with a logged reason and counted"); it
// only returns an error for something that stops the run outright -- a
// database read failing, a render failing, or an upload failing.
func Run(ctx context.Context, d Deps) (Result, error) {
	since := d.Now.Add(-ratingWindow)

	characters, skipped, err := aggregateCharacters(ctx, d, since)
	if err != nil {
		return Result{}, err
	}

	guilds, err := aggregateGuilds(ctx, d, since)
	if err != nil {
		return Result{}, err
	}

	files, err := Render(Data{Generated: d.Now, Build: d.Build, Characters: characters, Guilds: guilds})
	if err != nil {
		return Result{}, fmt.Errorf("dataaddon: render: %w", err)
	}

	sizes, err := publish(ctx, d, files)
	if err != nil {
		return Result{}, err
	}

	return Result{
		Characters: len(characters), Guilds: len(guilds), SkippedRows: skipped,
		Files: sizes, Rendered: files,
	}, nil
}

// aggregateCharacters reads every in-window character fight, groups them by
// player key, and aggregates each player's fights into a characterRow. A
// player_key that cannot be split into region/ruleset/name is logged and
// counted as skipped rather than failing the run.
func aggregateCharacters(ctx context.Context, d Deps, since time.Time) (map[string]characterRow, int, error) {
	fights, err := d.Store.characterFights(ctx, since)
	if err != nil {
		return nil, 0, err
	}
	byPlayer := map[string][]fightScore{}
	for _, f := range fights {
		byPlayer[f.PlayerKey] = append(byPlayer[f.PlayerKey], fightScore{
			Overall: f.Overall, FoughtAt: f.FoughtAt, Components: decodeComponents(f.Components),
		})
	}

	var skipped int
	characters := map[string]characterRow{}
	for playerKey, fs := range byPlayer {
		region, ruleset, name, ok := splitCharacterKey(playerKey)
		if !ok {
			d.logger().Warn("dataaddon", "op", "character", "err", "unparseable player_key "+playerKey)
			skipped++
			continue
		}
		row, ok := aggregateCharacter(fs)
		if !ok {
			continue
		}
		characters[characterKey(region, ruleset, name)] = row
	}
	return characters, skipped, nil
}

// aggregateGuilds reads every guild with at least one verified member,
// along with its verified members, raid nights, and progression, and
// aggregates each into a guildRow.
func aggregateGuilds(ctx context.Context, d Deps, since time.Time) (map[string]guildRow, error) {
	guildIdentities, err := d.Store.guildsWithVerifiedMembers(ctx)
	if err != nil {
		return nil, err
	}
	guilds := map[string]guildRow{}
	if len(guildIdentities) == 0 {
		return guilds, nil
	}

	ids := make([]int64, len(guildIdentities))
	for i, g := range guildIdentities {
		ids[i] = g.ID
	}
	members, err := d.Store.verifiedMembers(ctx, ids)
	if err != nil {
		return nil, err
	}
	nights, err := d.Store.nightsByGuild(ctx, ids, since)
	if err != nil {
		return nil, err
	}
	killed, total, err := d.Store.progressionByGuild(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, g := range guildIdentities {
		guilds[guildKey(g.Region, g.Ruleset, g.Name)] = buildGuildRow(
			g, members[g.ID], nights[g.ID], killed[g.ID], total[g.ID])
	}
	return guilds, nil
}

// publish logs each rendered file's size and, unless this deployment has no
// bucket configured, uploads every file plus the manifest under
// objectKeyPrefix. It returns the filename -> byte size map for Result.Files
// regardless of whether the upload happened.
func publish(ctx context.Context, d Deps, files map[string][]byte) (map[string]int, error) {
	sizes := map[string]int{}
	for name, body := range files {
		sizes[name] = len(body)
		d.logger().Info("dataaddon", "op", "render", "file", name, "bytes", len(body))
	}

	if d.Upload == nil || d.Bucket == "" {
		d.logger().Warn("dataaddon", "op", "upload",
			"err", "no-op: DATA_ADDON_BUCKET is not set, this deployment publishes no addon data")
		return sizes, nil
	}

	for name, body := range files {
		if err := d.Upload.Upload(ctx, d.Bucket, objectKeyPrefix+name, body); err != nil {
			return nil, fmt.Errorf("dataaddon: upload %s: %w", name, err)
		}
	}
	if err := d.Upload.Upload(ctx, d.Bucket, objectKeyPrefix+"manifest.txt", manifestOf(files)); err != nil {
		return nil, fmt.Errorf("dataaddon: upload manifest: %w", err)
	}
	return sizes, nil
}

// manifestOf lists every published filename, one per line, sorted -- what
// addon-data-release.yml reads to know which files to download and which
// TOC lines to write, without guessing from a fixed list of regions that
// may not all be present.
func manifestOf(files map[string][]byte) []byte {
	names := make([]string, 0, len(files))
	for name := range files {
		names = append(names, name)
	}
	sort.Strings(names)
	var buf bytes.Buffer
	for _, name := range names {
		buf.WriteString(name)
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}

// splitCharacterKey reads region/ruleset/name back out of a player_key
// ("us/normal/thoradin" -> "us", "normal", "thoradin", true) -- mirrors
// api/internal/rating/backfill.go's own splitPlayerKeyRegionRuleset,
// extended to also return the name segment this package needs.
func splitCharacterKey(key string) (region, ruleset, name string, ok bool) {
	parts := strings.SplitN(key, "/", 3)
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", false
	}
	return parts[0], parts[1], parts[2], true
}
