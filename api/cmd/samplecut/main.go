// api/cmd/samplecut/main.go
//
// samplecut cuts the site's sample report from a real combat log, offline. It runs the same
// parse the API's whole-file job runs (engine.SessionOptions over a session, the store
// Publisher writing report.json and each fight's summary and events), after the log has been
// pseudonymised: every player GUID and name is replaced before the engine sees a byte, so no
// original name reaches a summary, a unit list or a Parquet file.
//
//	go run ./cmd/samplecut -log WoWCombatLog.txt -out ../web/public/sample-report/sample
//
// Fights with no player in them are dropped and the rest renumbered from 1. The log is only
// read. -names writes the originals to a file (kept outside the repo) for the leak check.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

const (
	reportID = "sample"
	dataBase = "/sample-report/sample"
	title    = "Beta night, 17 Sept 2026"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "samplecut:", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("samplecut", flag.ContinueOnError)
	logPath := fs.String("log", "", "the combat log to cut (read only)")
	outDir := fs.String("out", "", "the directory to write meta.json, report.json and fights/")
	seed := fs.String("seed", "forever-sample-2026-09-17", "the pseudonym seed")
	names := fs.String("names", "", "write the original unit names here, for the leak check")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *logPath == "" || *outDir == "" {
		return fmt.Errorf("-log and -out are required")
	}
	raw, err := os.ReadFile(*logPath)
	if err != nil {
		return err
	}
	ps := NewPseudonymiser(*seed)
	log := ps.Log(raw)

	// The log's own header carries the year; Base only matters for year-less dialects.
	s := session.New(engine.SessionOptions(reportID, time.Date(2026, 9, 17, 0, 0, 0, 0, time.UTC), true))
	var closed []session.Closed
	const chunk = 1 << 20
	for off := 0; off < len(log); off += chunk {
		res, err := s.Feed(log[off:min(off+chunk, len(log))], int64(off))
		if err != nil {
			return err
		}
		closed = append(closed, res.Closed...)
	}
	res, err := s.Close()
	if err != nil {
		return err
	}
	closed = append(closed, res.Closed...)

	if err := os.RemoveAll(*outDir); err != nil {
		return err
	}
	pub := store.Publisher{Keys: store.Keys{ReportID: reportID}, Put: store.NewDir(*outDir)}
	var entries []store.FightEntry
	ctx := context.Background()
	for _, c := range closed {
		if len(c.Fight.Players) == 0 {
			continue
		}
		n := len(entries) + 1
		c.Fight.Index, c.Summary.FightIndex = n, n
		trim(&c.Summary)
		if err := pub.WriteFight(ctx, n, c.Summary, c.Events); err != nil {
			return err
		}
		entries = append(entries, store.EntryOf(c.Fight))
	}
	if len(entries) == 0 {
		return fmt.Errorf("the log holds no fight with a player in it")
	}
	rep := store.Report{
		ReportID: reportID, EngineVersion: engine.Version, Health: s.Health(),
		Fights: entries, Units: s.Units().All(),
	}
	if err := pub.WriteReport(ctx, rep); err != nil {
		return err
	}
	// The store keys objects under reports/<id>/; the static layout is flat.
	if err := lift(*outDir); err != nil {
		return err
	}
	if err := writeMeta(*outDir, rep); err != nil {
		return err
	}
	if *names != "" {
		if err := os.WriteFile(*names, []byte(strings.Join(ps.Originals, "\n")+"\n"), 0o600); err != nil {
			return err
		}
	}
	leaks, err := leakCheck(*outDir, ps)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "fights kept %d of %d, players %d, units %d\n", len(entries), len(closed), len(playersOf(entries)), len(rep.Units))
	fmt.Fprintf(out, "pseudonymised %d unit names; leaks: %d\n", len(ps.Originals), len(leaks))
	for _, l := range leaks {
		fmt.Fprintln(out, "  leak:", l)
	}
	if len(leaks) > 0 {
		return fmt.Errorf("an original unit name survived in the output")
	}
	return nil
}

// lift moves reports/<id>/* up to the output root.
func lift(dir string) error {
	src := filepath.Join(dir, "reports", reportID)
	ents, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, e := range ents {
		if err := os.Rename(filepath.Join(src, e.Name()), filepath.Join(dir, e.Name())); err != nil {
			return err
		}
	}
	return os.RemoveAll(filepath.Join(dir, "reports"))
}

func playersOf(entries []store.FightEntry) []string {
	seen := map[string]bool{}
	var out []string
	for _, e := range entries {
		for _, g := range e.Players {
			if !seen[g] {
				seen[g] = true
				out = append(out, g)
			}
		}
	}
	sort.Strings(out)
	return out
}

// groupAffiliation is the unit flags' MINE, PARTY and RAID bits (logs/engine/units).
const groupAffiliation = units.FlagAffiliationMine | units.FlagAffiliationParty | units.FlagAffiliationRaid

// groupOf is the players in the logger's group, by each unit's last-seen flags, sorted. The
// sample's meta carries it so the island's first table scopes to the group before report.json
// (every unit) arrives; web/src/lib/report/group.ts reads the same bits off the units after.
func groupOf(us []units.Unit) []string {
	out := []string{}
	for _, u := range us {
		if u.Kind == units.KindPlayer && u.Flags&groupAffiliation != 0 {
			out = append(out, u.GUID)
		}
	}
	sort.Strings(out)
	return out
}

// writeMeta writes meta.json in the shape of GET /v1/reports/{id}: public, no owner, the
// night's own date and zone, and the static data base.
func writeMeta(dir string, rep store.Report) error {
	zones := map[string]int{}
	zone := ""
	for _, f := range rep.Fights {
		if f.Zone != "" {
			zones[f.Zone]++
			if zones[f.Zone] > zones[zone] {
				zone = f.Zone
			}
		}
	}
	meta := map[string]any{
		"id": reportID, "title": title, "visibility": "public", "owner": nil,
		"zone": zone, "status": "complete", "engine_version": rep.EngineVersion,
		"fights": rep.Fights, "players": playersOf(rep.Fights), "group": groupOf(rep.Units),
		"created_at": rep.Fights[0].Start.UTC(), "data_base_url": dataBase, "phase": "launch",
	}
	b, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "meta.json"), append(b, '\n'), 0o644)
}

// leakCheck scans every written file, Parquet decoded back to JSON, for any original unit
// name or bare name word. A hit on a bare word can be a name that is also an ordinary word in
// the log (a spell, an NPC); the caller reads the list.
func leakCheck(dir string, ps *Pseudonymiser) ([]string, error) {
	var leaks []string
	needles := ps.Originals
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".parquet") {
			evs, err := parquet.Unmarshal(b)
			if err != nil {
				return err
			}
			if b, err = json.Marshal(evs); err != nil {
				return err
			}
		}
		for _, n := range needles {
			if bytes.Contains(b, []byte(n)) {
				leaks = append(leaks, fmt.Sprintf("%s in %s", n, path))
			}
		}
		return nil
	})
	return leaks, err
}
