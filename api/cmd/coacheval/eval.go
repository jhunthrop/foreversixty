package main

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

// ModelSpec names a model under test.
type ModelSpec struct {
	Slug        string
	Temperature float64
	Reasoning   string // the level the client was built with; reported only
}

// ClientFactory returns the client for a model on one built pack: the shared gateway client
// in a real run, a scripted llm.Fake in a smoke run.
type ClientFactory func(m ModelSpec, b Built) (llm.Client, error)

// Evaluation is everything a run needs, injected.
type Evaluation struct {
	Source    Source
	Tables    *coach.FileTables
	Pool      PoolReader
	CoachDir  string
	Models    []ModelSpec
	ClientFor ClientFactory
	N         int
	Now       func() time.Time
	Log       *slog.Logger
	// Names is the validator's name dictionary; built from the corpus's build when nil.
	Names *coach.NameDict
}

// Evaluate reads the fights, builds each pack once and runs every model over the same
// packs. A fight that cannot be built or reviewed is recorded as skipped, not fatal.
func (e Evaluation) Evaluate(ctx context.Context) ([]Target, []Result, error) {
	fights, err := e.Source.Fights(ctx, e.N)
	if err != nil {
		return nil, nil, fmt.Errorf("read fights: %w", err)
	}
	targets := planTargets(fights, fightLabel)
	names, err := e.nameDict(fights)
	if err != nil {
		return nil, nil, err
	}
	builder := Builder{Tables: e.Tables, Pool: e.Pool, Personas: &coach.Personas{Dir: e.CoachDir}}
	runner := Runner{Names: names, Now: e.Now}
	var results []Result
	for _, t := range targets {
		results = append(results, e.runTarget(ctx, builder, runner, t)...)
	}
	return targets, results, nil
}

func fightLabel(i int) string { return fmt.Sprintf("f%02d", i+1) }

func (e Evaluation) nameDict(fights []Fight) (*coach.NameDict, error) {
	if e.Names != nil {
		return e.Names, nil
	}
	return e.loadNames(fights)
}

func (e Evaluation) loadNames(fights []Fight) (*coach.NameDict, error) {
	build, err := e.Tables.BuildAt(e.Now())
	if err != nil {
		return nil, fmt.Errorf("build in force: %w", err)
	}
	names, err := coach.LoadNames(filepath.Join(e.Tables.BuildsDir, build), mechanicsFor(fights))
	if err != nil {
		return nil, fmt.Errorf("name dictionary: %w", err)
	}
	return names, nil
}

func (e Evaluation) runTarget(ctx context.Context, b Builder, r Runner, t Target) []Result {
	built, err := b.Build(ctx, t)
	if err != nil {
		e.Log.Warn("coacheval: fight not built", slog.String("fight", t.Label), slog.String("error", err.Error()))
		return e.skipped(t, err)
	}
	out := make([]Result, 0, len(e.Models))
	for _, m := range e.Models {
		client, err := e.ClientFor(m, built)
		if err != nil {
			out = append(out, modelError(baseResult(built, Model{Slug: m.Slug, Reasoning: m.Reasoning}), err))
			continue
		}
		out = append(out, r.Run(ctx, built, Model{Slug: m.Slug, Client: client, Temperature: m.Temperature, Reasoning: m.Reasoning}))
	}
	return out
}

// skipped is one skipped result per model for a fight that could not be built.
func (e Evaluation) skipped(t Target, cause error) []Result {
	out := make([]Result, 0, len(e.Models))
	for _, m := range e.Models {
		out = append(out, Result{Label: t.Label, Model: m.Slug, Spec: t.Row.Class + " " + t.Row.Spec, Role: t.Row.Role,
			Size: t.Fight.Size, Outcome: OutcomeSkipped, Note: cause.Error()})
	}
	return out
}
