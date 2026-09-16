// logs/cmd/forever-logs/main.go
// Command forever-logs parses combat logs with the engine.
//
//	forever-logs parse   <file>   parse a log and write the report files
//	forever-logs fights  <file>   list the fights one per line
//	forever-logs tail    <file>   follow a growing log and print fights live
//	forever-logs conformance <dir> report unknown events, parse errors, and
//	                              inferred layouts over every file in a tree
//	forever-logs mechanics-draft <file> propose a mechanics table per
//	                              encounter from the log's evidence
//
// Deep queries are not here on purpose: the design runs them in the
// browser over the fight's Parquet file.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

const chunkSize = 1 << 20

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "forever-logs:", err)
		os.Exit(1)
	}
}

func usage() error {
	return errors.New("usage: forever-logs parse|fights|tail|conformance|mechanics-draft ...")
}

func run(args []string, out, errOut io.Writer) error {
	if len(args) == 0 {
		return usage()
	}
	switch args[0] {
	case "parse":
		return cmdParse(args[1:], out)
	case "fights":
		return cmdFights(args[1:], out)
	case "tail":
		return cmdTail(args[1:], out)
	case "conformance":
		return cmdConformance(args[1:], out, errOut)
	case "mechanics-draft":
		return runMechanicsDraft(args[1:], out, errOut)
	default:
		return usage()
	}
}

// sessionOptions builds the options every command shares.
func sessionOptions(reportID string, base time.Time, layoutName string, keepEvents bool) (session.Options, error) {
	o := session.Options{
		ReportID:   reportID,
		Base:       base,
		Infer:      true,
		KeepEvents: keepEvents,
		Units:      units.Options{ClassBySpec: units.RetailSpecClass},
		Fight:      fight.DefaultOptions(),
		Summary:    summary.DefaultOptions(),
	}
	o.Summary.SpecNames = units.RetailSpecName
	if layoutName != "" {
		found := false
		for _, r := range layout.Rows() {
			if r.Name == layoutName {
				o.Layout, found = r, true
			}
		}
		if !found {
			return o, fmt.Errorf("no layout row named %q; rows are %s", layoutName, rowNames())
		}
	}
	return o, nil
}

func rowNames() string {
	var names []string
	for _, r := range layout.Rows() {
		names = append(names, r.Name)
	}
	return strings.Join(names, ", ")
}

// baseTime is the year source for dialects with no year in the timestamp.
func baseTime(path string) time.Time {
	if fi, err := os.Stat(path); err == nil {
		return fi.ModTime().UTC()
	}
	return time.Now().UTC()
}

// stream feeds a whole file through a session, calling onClosed for each
// fight as it closes so memory stays bounded by the open fight.
func stream(s *session.Session, f io.Reader, onClosed func(session.Closed) error) error {
	buf := make([]byte, chunkSize)
	offset := s.Offset()
	for {
		n, err := f.Read(buf)
		if n > 0 {
			res, ferr := s.Feed(buf[:n], offset)
			if ferr != nil {
				return ferr
			}
			offset += int64(n)
			for _, c := range res.Closed {
				if cerr := onClosed(c); cerr != nil {
					return cerr
				}
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
	}
	res, err := s.Close()
	if err != nil {
		return err
	}
	for _, c := range res.Closed {
		if cerr := onClosed(c); cerr != nil {
			return cerr
		}
	}
	return nil
}

func cmdParse(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("parse", flag.ContinueOnError)
	fs.SetOutput(out)
	outDir := fs.String("out", "", "directory to write the report files into; empty prints a summary only")
	reportID := fs.String("report", "local", "report id used in the object keys and the metrics rows")
	layoutName := fs.String("layout", "", "force a layout row ("+rowNames()+")")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return errors.New("usage: forever-logs parse [-out dir] [-report id] [-layout name] <file>")
	}
	path := fs.Arg(0)
	o, err := sessionOptions(*reportID, baseTime(path), *layoutName, *outDir != "")
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	s := session.New(o)
	var pub *store.Publisher
	if *outDir != "" {
		pub = &store.Publisher{Keys: store.Keys{ReportID: *reportID}, Put: store.NewDir(*outDir)}
	}
	ctx := context.Background()
	var entries []store.FightEntry
	onClosed := func(c session.Closed) error {
		entries = append(entries, store.EntryOf(c.Fight))
		fmt.Fprintf(out, "fight %d %-10s %-32s %6.1fs players=%d deaths=%d kill=%v\n",
			c.Fight.Index, c.Fight.Kind, c.Fight.Name,
			c.Fight.Duration().Seconds(), len(c.Fight.Players), c.Fight.Deaths, c.Fight.Kill)
		if pub == nil {
			return nil
		}
		return pub.WriteFight(ctx, c.Fight.Index, c.Summary, c.Events)
	}
	if err := stream(s, file, onClosed); err != nil {
		return err
	}
	h := s.Health()
	fmt.Fprintf(out, "lines=%d fights=%d layout=%s verified=%v advanced=%v parse_errors=%d unknown=%d\n",
		h.Lines, len(entries), h.Layout, h.LayoutVerified, h.AdvancedLogging, h.ParseErrors, len(h.UnknownEvents))
	if pub == nil {
		return nil
	}
	return pub.WriteReport(ctx, store.Report{
		ReportID: *reportID, EngineVersion: session.Version,
		Health: h, Fights: entries, Units: s.Units().All(),
	})
}

func cmdFights(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("fights", flag.ContinueOnError)
	fs.SetOutput(out)
	asJSON := fs.Bool("json", false, "print one JSON object per fight")
	layoutName := fs.String("layout", "", "force a layout row ("+rowNames()+")")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return errors.New("usage: forever-logs fights [-json] [-layout name] <file>")
	}
	path := fs.Arg(0)
	o, err := sessionOptions("local", baseTime(path), *layoutName, false)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(out)
	return stream(session.New(o), file, func(c session.Closed) error {
		if *asJSON {
			return enc.Encode(store.EntryOf(c.Fight))
		}
		fmt.Fprintf(out, "%d\t%s\t%s\t%.1fs\tkill=%v\tplayers=%d\n",
			c.Fight.Index, c.Fight.Kind, c.Fight.Name,
			c.Fight.Duration().Seconds(), c.Fight.Kill, len(c.Fight.Players))
		return nil
	})
}

func cmdTail(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("tail", flag.ContinueOnError)
	fs.SetOutput(out)
	poll := fs.Duration("poll", time.Second, "how often to check the file for new bytes")
	once := fs.Bool("once", false, "stop at the end of the file instead of following")
	layoutName := fs.String("layout", "", "force a layout row ("+rowNames()+")")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return errors.New("usage: forever-logs tail [-poll d] [-once] [-layout name] <file>")
	}
	path := fs.Arg(0)
	o, err := sessionOptions("local", time.Now().UTC(), *layoutName, false)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	s := session.New(o)
	buf := make([]byte, chunkSize)
	offset := int64(0)
	for {
		n, rerr := file.Read(buf)
		if n > 0 {
			res, ferr := s.Feed(buf[:n], offset)
			if ferr != nil {
				return ferr
			}
			offset += int64(n)
			for _, c := range res.Closed {
				fmt.Fprintf(out, "closed  %d %-10s %-32s %6.1fs kill=%v\n",
					c.Fight.Index, c.Fight.Kind, c.Fight.Name,
					c.Fight.Duration().Seconds(), c.Fight.Kill)
			}
			if f, sum, ok := s.Snapshot(); ok {
				fmt.Fprintf(out, "live    %d %-10s %-32s %6.1fs rows=%d\n",
					f.Index, f.Kind, f.Name,
					float64(sum.DurationMS)/1000, len(sum.DamageDone))
			}
			continue
		}
		if rerr != nil && rerr != io.EOF {
			return fmt.Errorf("read: %w", rerr)
		}
		if *once {
			res, cerr := s.Close()
			if cerr != nil {
				return cerr
			}
			for _, c := range res.Closed {
				fmt.Fprintf(out, "closed  %d %-10s %-32s %6.1fs kill=%v\n",
					c.Fight.Index, c.Fight.Kind, c.Fight.Name,
					c.Fight.Duration().Seconds(), c.Fight.Kill)
			}
			return nil
		}
		time.Sleep(*poll)
	}
}

// ConformanceRow is one file's result in the conformance report.
type ConformanceRow struct {
	Path          string         `json:"path"`
	Bytes         int64          `json:"bytes"`
	Lines         int64          `json:"lines"`
	Fights        int            `json:"fights"`
	Layout        string         `json:"layout"`
	Verified      bool           `json:"layout_verified"`
	Inferred      bool           `json:"layout_inferred"`
	Advanced      bool           `json:"advanced_logging"`
	ParseErrors   int64          `json:"parse_errors"`
	UnknownEvents map[string]int `json:"unknown_events"`
	Error         string         `json:"error,omitempty"`
}

func cmdConformance(args []string, out, errOut io.Writer) error {
	fset := flag.NewFlagSet("conformance", flag.ContinueOnError)
	fset.SetOutput(out)
	asJSON := fset.Bool("json", false, "print one JSON object per file")
	if err := fset.Parse(args); err != nil {
		return err
	}
	if fset.NArg() != 1 {
		return errors.New("usage: forever-logs conformance [-json] <dir>")
	}
	root := fset.Arg(0)

	var rows []ConformanceRow
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		switch strings.ToLower(filepath.Ext(path)) {
		case ".txt", ".log":
		default:
			return nil
		}
		rows = append(rows, conform(path))
		return nil
	})
	if err != nil {
		return err
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].Path < rows[j].Path })

	totalUnknown := map[string]int{}
	var totalErrors int64
	enc := json.NewEncoder(out)
	for _, r := range rows {
		for k, v := range r.UnknownEvents {
			totalUnknown[k] += v
		}
		totalErrors += r.ParseErrors
		if *asJSON {
			if err := enc.Encode(r); err != nil {
				return err
			}
			continue
		}
		fmt.Fprintf(out, "%s\n  layout=%s verified=%v inferred=%v advanced=%v lines=%d fights=%d parse_errors=%d\n",
			r.Path, r.Layout, r.Verified, r.Inferred, r.Advanced, r.Lines, r.Fights, r.ParseErrors)
		if r.Error != "" {
			fmt.Fprintf(out, "  error: %s\n", r.Error)
		}
		for _, name := range sortedKeys(r.UnknownEvents) {
			fmt.Fprintf(out, "  unknown %-32s %d\n", name, r.UnknownEvents[name])
		}
	}
	if !*asJSON {
		fmt.Fprintf(out, "\n%d files, %d parse errors, %d distinct unknown events\n",
			len(rows), totalErrors, len(totalUnknown))
		for _, name := range sortedKeys(totalUnknown) {
			fmt.Fprintf(out, "  %-32s %d\n", name, totalUnknown[name])
		}
	}
	if len(rows) == 0 {
		fmt.Fprintf(errOut, "no .txt or .log files under %s\n", root)
	}
	return nil
}

func conform(path string) ConformanceRow {
	row := ConformanceRow{Path: path, UnknownEvents: map[string]int{}}
	fi, err := os.Stat(path)
	if err != nil {
		row.Error = err.Error()
		return row
	}
	row.Bytes = fi.Size()
	o, err := sessionOptions("conformance", fi.ModTime().UTC(), "", false)
	if err != nil {
		row.Error = err.Error()
		return row
	}
	f, err := os.Open(path)
	if err != nil {
		row.Error = err.Error()
		return row
	}
	defer f.Close()
	s := session.New(o)
	if err := stream(s, f, func(session.Closed) error { row.Fights++; return nil }); err != nil {
		row.Error = err.Error()
	}
	h := s.Health()
	row.Lines, row.Layout = h.Lines, h.Layout
	row.Verified, row.Inferred, row.Advanced = h.LayoutVerified, h.LayoutInferred, h.AdvancedLogging
	row.ParseErrors, row.UnknownEvents = h.ParseErrors, h.UnknownEvents
	return row
}

func sortedKeys(m map[string]int) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
