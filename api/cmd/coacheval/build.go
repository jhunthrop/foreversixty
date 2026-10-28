package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// PoolReader is the read-only pool lookup (*coach.Pool satisfies it); nil offline.
type PoolReader interface {
	Reference(ctx context.Context, k coach.PoolKey, exclude string) (int, *coach.PoolRef, error)
}

// TableSource is *coach.FileTables.
type TableSource interface {
	TablesFor(ctx context.Context, row summary.RosterRow, encounterID int64, foughtAt time.Time) (coach.Tables, error)
}

// Built is one target's pack with what generation and validation need beside it.
type Built struct {
	Target  Target
	Pack    coach.Pack
	Persona coach.Persona
	Roster  []string
}

// errSkip marks a target that cannot be evaluated (no data tables, no persona); the run
// records it and goes on.
var errSkip = errors.New("skipped")

// Builder builds packs exactly as the service does: coach.Review over FileTables data.
type Builder struct {
	Tables   TableSource
	Pool     PoolReader
	Personas *coach.Personas
}

// Build assembles the Input for t and runs Review.
func (b Builder) Build(ctx context.Context, t Target) (Built, error) {
	f, row := t.Fight, t.Row
	tables, err := b.Tables.TablesFor(ctx, row, f.Summary.EncounterID, f.FoughtAt)
	if errors.Is(err, coach.ErrNoTables) {
		return Built{}, fmt.Errorf("%w: %v", errSkip, err)
	}
	if err != nil {
		return Built{}, fmt.Errorf("tables: %w", err)
	}
	persona, err := b.Personas.For(rating.DataSpecSlug(row.Class, row.Spec))
	if err != nil {
		return Built{}, fmt.Errorf("%w: %v", errSkip, err)
	}
	refs, err := b.refs(ctx, t)
	if err != nil {
		return Built{}, err
	}
	in := coach.Input{Fight: f.Summary, Row: row, Size: f.Size, BossName: bossName(tables), Refs: refs, Tables: tables}
	p := coach.Review(in)
	p.Version = coach.PackVersion(nil, nil)
	return Built{Target: t, Pack: p, Persona: persona, Roster: rosterNames(t)}, nil
}

// refs reads the top-logs pool when there is one and the player's key is known; the sim
// reference is never requested, because asking for it may queue a job (a write).
func (b Builder) refs(ctx context.Context, t Target) (coach.Refs, error) {
	key, ok := t.Fight.PlayerKeys[t.Row.Name]
	if b.Pool == nil || !ok {
		return coach.Refs{}, nil
	}
	k := coach.PoolKey{EncounterID: t.Fight.Summary.EncounterID, Difficulty: t.Fight.Summary.Difficulty,
		Size: t.Fight.Size, Spec: t.Row.Spec, Phase: phase.At(t.Fight.FoughtAt)}
	count, ref, err := b.Pool.Reference(ctx, k, key)
	if err != nil {
		return coach.Refs{}, fmt.Errorf("pool: %w", err)
	}
	return coach.Refs{Pool: ref, PoolCount: count}, nil
}

func bossName(t coach.Tables) string {
	if t.Mechanics == nil {
		return ""
	}
	return t.Mechanics.Name
}

func rosterNames(t Target) []string {
	names := make([]string, 0, len(t.Fight.Summary.Roster))
	for _, r := range t.Fight.Summary.Roster {
		names = append(names, r.Name)
	}
	return names
}

// mechanicsFor loads the mechanics tables of every encounter in fights, for the name dictionary.
func mechanicsFor(fights []Fight) []mechanics.Table {
	seen := map[int64]bool{}
	var out []mechanics.Table
	for _, f := range fights {
		id := f.Summary.EncounterID
		if seen[id] {
			continue
		}
		seen[id] = true
		if t, ok := mechanics.Load(id); ok {
			out = append(out, t)
		}
	}
	return out
}
