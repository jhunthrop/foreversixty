// api/cmd/samplecut/recount.go
//
// recount refreshes a cut sample's kill counts without its original log, which is not kept:
// every fight's events.parquet holds all the events the segmenter assigned to that fight, so
// fight.Counter over them gives the same mob_deaths, group_kills, group_deaths and group the
// engine would have written. Only those fields change, in report.json and meta.json; the
// summaries and engine_version stay as they were cut.
//
//	go run ./cmd/samplecut -recount ../web/public/sample-report/sample
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

func recount(dir string) error {
	var rep store.Report
	if err := readJSON(filepath.Join(dir, "report.json"), &rep); err != nil {
		return err
	}
	fights := make([]store.FightEntry, 0, len(rep.Fights))
	for _, e := range rep.Fights {
		counted, err := countFight(dir, e)
		if err != nil {
			return err
		}
		fights = append(fights, counted)
	}
	rep.Fights = fights
	b, err := json.Marshal(rep)
	if err != nil {
		return fmt.Errorf("samplecut: marshal report.json: %w", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "report.json"), b, 0o644); err != nil {
		return err
	}
	return recountMeta(dir, fights)
}

// countFight returns e with the counts of its own events.
func countFight(dir string, e store.FightEntry) (store.FightEntry, error) {
	path := filepath.Join(dir, "fights", strconv.Itoa(e.Index), "events.parquet")
	raw, err := os.ReadFile(path)
	if err != nil {
		return e, err
	}
	events, err := parquet.Unmarshal(raw)
	if err != nil {
		return e, fmt.Errorf("samplecut: %s: %w", path, err)
	}
	c := fight.NewCounter()
	for _, ev := range events {
		c.Observe(ev)
	}
	var f fight.Fight
	c.Apply(&f)
	e.MobDeaths, e.GroupKills, e.GroupDeaths = &f.MobDeaths, &f.GroupKills, &f.GroupDeaths
	e.Group = f.Group
	return e, nil
}

// recountMeta swaps meta.json's fight lines and keeps every other key as it was.
func recountMeta(dir string, fights []store.FightEntry) error {
	path := filepath.Join(dir, "meta.json")
	meta := map[string]json.RawMessage{}
	if err := readJSON(path, &meta); err != nil {
		return err
	}
	lines, err := json.Marshal(fights)
	if err != nil {
		return fmt.Errorf("samplecut: marshal meta fights: %w", err)
	}
	meta["fights"] = lines
	b, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return fmt.Errorf("samplecut: marshal meta.json: %w", err)
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}

func readJSON(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("samplecut: %s: %w", path, err)
	}
	return nil
}
