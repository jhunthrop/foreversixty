// api/internal/dataaddon/job.go
package dataaddon

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
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
	Store *Store
	// Trees is the loaded client build data a no-logs record's spec is read from (its
	// export's talents); nil writes no spec for those records. The scores' gear needs
	// sims.SetGearData wired to the same data, as the job's command does.
	Trees  *trees.Data
	Upload Uploader // nil (or an empty Bucket) means this deployment has no bucket; the run still succeeds and logs why the upload was skipped
	Bucket string
	Build  string // the active data build, logged with the run
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
	Characters  int               // records written
	Guilds      int               // records written
	SkippedRows int               // characters or guilds left out over a bad key or a key two of them share
	Files       map[string]int    // filename -> byte size (the dispatch's "size line")
	Rendered    map[string][]byte // filename -> content, for a caller (the golden test) that needs the bytes themselves
}

// objectKeyPrefix is where every published file (and the manifest) lives
// inside the bucket.
const objectKeyPrefix = "data-addon/"

// Run reads the database, builds every record, renders Data.lua (or its region-split
// form), and uploads it, in that order. It never fails the whole run over one bad row
// (skipped with a logged reason and counted); it only returns an error for something that
// stops the run outright -- a database read failing, a render failing, or an upload
// failing.
func Run(ctx context.Context, d Deps) (Result, error) {
	data, skipped, err := collect(ctx, d)
	if err != nil {
		return Result{}, err
	}

	files, err := Render(data)
	if err != nil {
		return Result{}, fmt.Errorf("dataaddon: render: %w", err)
	}

	sizes, err := publish(ctx, d, files)
	if err != nil {
		return Result{}, err
	}

	d.logger().Info("dataaddon", "op", "run", "build", d.Build,
		"characters", len(data.Characters), "guilds", len(data.Guilds), "skipped", skipped)
	return Result{
		Characters: len(data.Characters), Guilds: len(data.Guilds), SkippedRows: skipped,
		Files: sizes, Rendered: files,
	}, nil
}

// maxCharacterFailures is how many characters a run may fail to read (a ScoreV2,
// anonymize or item-level read that errors) before it stops: past it the fault is the
// database's or the data's, not one row's. A variable so a test can lower it.
var maxCharacterFailures = 50

// collect builds the night's records (contract 5, docs/contracts/addon-v2.md). The
// characters are every one with a public, sufficient rated fight in the last 90 days and
// every verified guild member:
//
//   - a character whose account asked to be anonymous gets no record and is listed in no
//     guild (the endpoint's own rule, rating.Store.Anonymized): the file is public;
//   - otherwise rating.ScoreV2 alone decides: a score is a scored record, with the best
//     parse per boss in the window; ErrNoLogs is a no-logs record (score = nil), so the
//     guild's roster still lists the character;
//   - a character with no stored name, or no realm from Battle.net, its export or its
//     newest rated row, gets no record (its key would be a guess), logged and counted;
//   - one whose reads fail is logged and counted (fail), and the file goes on without it;
//   - every guild with a verified member gets a record.
//
// Character keys are Name-Realm and guild keys the guild's name, neither of which carries
// the region: two characters (or guilds) sharing a key would each read as the other in
// game, so both are left out, logged and counted rather than one being shown as the other.
func collect(ctx context.Context, d Deps) (Data, int, error) {
	since := d.Now.Add(-ratingWindow)
	scores := &rating.Store{Pool: d.Store.Pool}
	if l, ok := d.Log.(*slog.Logger); ok {
		scores.Log = l
	}

	parses, err := d.Store.windowParses(ctx, since)
	if err != nil {
		return Data{}, 0, err
	}
	guildIdentities, err := d.Store.guildsWithVerifiedMembers(ctx)
	if err != nil {
		return Data{}, 0, err
	}
	ids := make([]int64, len(guildIdentities))
	for i, g := range guildIdentities {
		ids[i] = g.ID
	}
	members, err := d.Store.verifiedMembers(ctx, ids)
	if err != nil {
		return Data{}, 0, err
	}

	keySet := map[string]bool{}
	for key := range parses {
		keySet[key] = true
	}
	for _, keys := range members {
		for _, key := range keys {
			keySet[key] = true
		}
	}
	keys := make([]string, 0, len(keySet))
	for key := range keySet {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	who, err := d.Store.identities(ctx, keys)
	if err != nil {
		return Data{}, 0, err
	}
	exports, err := d.Store.exports(ctx, keys)
	if err != nil {
		return Data{}, 0, err
	}

	var skipped, failures int
	// fail logs and counts one character that could not be read, so one bad row never costs
	// the whole night's file; a cancelled run, or more than maxCharacterFailures of them
	// (something wrong with the database or the data, not with one row), stops the run.
	fail := func(key, op string, err error) error {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return ctxErr
		}
		failures++
		skipped++
		d.logger().Warn("dataaddon", "op", op, "character", key, "err", err)
		if failures > maxCharacterFailures {
			return fmt.Errorf("dataaddon: more than %d characters could not be read, the last %s: %w",
				maxCharacterFailures, key, err)
		}
		return nil
	}
	records := map[string]characterRecord{} // character key -> record
	recordKeys := map[string][]string{}     // Name-Realm -> character keys
	for _, key := range keys {
		// An anonymised character is left out before anything else is read about it, and
		// silently: it is not a fault, and it is not counted.
		anon, err := scores.Anonymized(ctx, key)
		if err != nil {
			if err := fail(key, "anonymized", err); err != nil {
				return Data{}, 0, err
			}
			continue
		}
		if anon {
			continue
		}
		region, ruleset, slug, ok := splitCharacterKey(key)
		nameKey := ""
		if ok {
			nameKey = characterKey(nameAndRealm(who[key], exports[key]))
		}
		if nameKey == "" {
			d.logger().Warn("dataaddon", "op", "character", "character", key,
				"err", "no Name-Realm key: no stored name, or no realm from Battle.net, the export or the log")
			skipped++
			continue
		}
		rec := characterRecord{Region: region, Profile: profilePath(region, ruleset, slug)}
		// Scored or not is ScoreV2's answer alone: ErrNoLogs is a no-logs record.
		sc, err := scores.ScoreV2(ctx, key)
		switch {
		case errors.Is(err, rating.ErrNoLogs):
			if rec.ItemLevel, err = scores.ItemLevel(ctx, key); err != nil {
				if err := fail(key, "item_level", err); err != nil {
					return Data{}, 0, err
				}
				continue
			}
			rec.Spec = exportSpec(d.Trees, exports[key])
		case err != nil:
			if err := fail(key, "score", err); err != nil {
				return Data{}, 0, err
			}
			continue
		default:
			rec.Score = &sc
			rec.ItemLevel, rec.Spec = sc.ItemLevel, sc.Spec
			rec.BestParses = make(map[string]int, len(parses[key]))
			for boss, overall := range parses[key] {
				rec.BestParses[boss] = percentile(overall)
			}
		}
		records[key] = rec
		recordKeys[nameKey] = append(recordKeys[nameKey], key)
	}

	data := Data{Generated: d.Now, Characters: map[string]characterRecord{}, Guilds: map[string]guildRecord{}}
	listed := map[string]string{} // character key -> Name-Realm, for the guilds' members
	for nameKey, owners := range recordKeys {
		if len(owners) > 1 {
			d.logger().Warn("dataaddon", "op", "character", "key", nameKey, "characters", owners,
				"err", "characters on different regions share this Name-Realm key; none is written")
			skipped += len(owners)
			continue
		}
		data.Characters[nameKey] = records[owners[0]]
		listed[owners[0]] = nameKey
	}

	guilds, guildSkipped, err := collectGuilds(ctx, d, since, guildIdentities, ids, members, data.Characters, listed)
	if err != nil {
		return Data{}, 0, err
	}
	data.Guilds = guilds
	return data, skipped + guildSkipped, nil
}

// collectGuilds builds every guild's record from its verified members and the character
// records already built (listed: character key -> Name-Realm key of the ones written).
func collectGuilds(ctx context.Context, d Deps, since time.Time, identities []guildIdentity, ids []int64,
	members map[int64][]string, characters map[string]characterRecord, listed map[string]string,
) (map[string]guildRecord, int, error) {
	out := map[string]guildRecord{}
	if len(identities) == 0 {
		return out, 0, nil
	}
	nights, err := d.Store.nightsByGuild(ctx, ids, since)
	if err != nil {
		return nil, 0, err
	}
	kills, err := d.Store.bossesKilledByGuild(ctx, ids, since)
	if err != nil {
		return nil, 0, err
	}

	records := make(map[int64]guildRecord, len(identities))
	ranked := make([]guildStanding, 0, len(identities))
	for _, g := range identities {
		rec := guildRecord{Region: g.Region, Nights: nights[g.ID], Roster: len(members[g.ID]), Members: []string{}}
		var fractions []float64
		for _, key := range members[g.ID] {
			nameKey, ok := listed[key]
			if !ok {
				continue
			}
			rec.Members = append(rec.Members, nameKey)
			if sc := characters[nameKey].Score; sc != nil {
				fractions = append(fractions, sc.Attendance)
			}
		}
		sort.Strings(rec.Members)
		rec.Attendance = meanAttendance(fractions)
		records[g.ID] = rec
		ranked = append(ranked, guildStanding{ID: g.ID, Kills: kills[g.ID], Attendance: rec.Attendance})
	}
	standing := standings(ranked)

	byName := map[string][]int64{}
	for _, g := range identities {
		byName[g.Name] = append(byName[g.Name], g.ID)
	}
	var skipped int
	for name, owners := range byName {
		if len(owners) > 1 {
			d.logger().Warn("dataaddon", "op", "guild", "guild", name, "guilds", owners,
				"err", "guilds on different regions or rulesets share this name; none is written")
			skipped += len(owners)
			continue
		}
		rec := records[owners[0]]
		rec.Standing = standing[owners[0]]
		out[name] = rec
	}
	return out, skipped, nil
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
