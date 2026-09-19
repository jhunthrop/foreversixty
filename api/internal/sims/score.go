package sims

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/runner"
	"github.com/jhunthrop/foreversixty/sim/talents"
)

// ScoreTask is one fight's player to score, as the ingest's
// fight-close path hands it over.
type ScoreTask struct {
	ReportID   string
	FightIndex int
	PlayerKey  string
	Spec       string
	Class      string
	// Role is the metrics row's role. The simulator models DPS only at
	// launch, and Score is the one place that knows it, so the ingest
	// hands over every parse it has and this drops the rest.
	Role string
	// ActualDPS is what the player actually did, from the fight's own
	// metrics row.
	ActualDPS float64
	// DurationSec is the fight's length, so the sim runs the fight
	// that happened rather than the default three minutes.
	DurationSec int
	// Combatant is the gear, talents, consumables, and raid buffs the
	// fight recorded for this player. It is what the character is
	// built from, and it is why the score is "what your gear can do"
	// rather than "what a best-in-slot character can do".
	Combatant summary.CombatantRow
}

// Builder turns a fight's recorded combatant into the sim envelope's
// character. No protobuf is involved: sim/request builds the engine's
// message inside forever-sim.
type Builder interface {
	FightCharacter(spec, class string, c summary.CombatantRow) (simapi.CharacterSpec, error)
}

// ErrNoCharacter is what a builder returns when a fight does not
// record enough to reconstruct the character. It is not a failure: no
// score is written, and the nightly job fills the column in when a
// source for the missing half exists.
var ErrNoCharacter = errors.New("sims: the fight does not record a simmable character")

// NoBuilder refuses every character. It is what a deployment with no
// talent layout runs, and what the tests use to prove that a fight
// nothing can be built from is skipped rather than failed.
type NoBuilder struct{}

// FightCharacter always reports that the fight cannot be rebuilt.
func (NoBuilder) FightCharacter(string, string, summary.CombatantRow) (simapi.CharacterSpec, error) {
	return simapi.CharacterSpec{}, ErrNoCharacter
}

// CombatantBuilder builds a character from what a fight recorded: the
// gear, the talents through the build's layout, the buffs and
// consumables through the engine's id vocabulary, and the one level the
// engine simulates.
//
// It cannot finish, and says which field stops it. Race is recorded by
// no log dialect - COMBATANT_INFO carries a faction number and nothing
// else here stores a race - and racials change the number the score
// divides by, so this refuses rather than guessing one. The day the
// companion writes race into its export, this returns a character and
// the execution score starts filling in with no other change.
type CombatantBuilder struct {
	// Talents is the build's talent layout, loaded once at startup.
	// Nil means no layout, which is one of the reasons to refuse.
	Talents *talents.Layouts
}

// FightCharacter assembles what the fight recorded.
//
// It collects every reason it cannot finish rather than returning the
// first: a caller reading the log wants to know that a fight is short
// a race *and* a talent layout, not to fix one and meet the other
// tomorrow. The race reason is always present today, so the returned
// error is never nil and Score always skips - which is the state the
// plan is honest about until a race source lands.
func (b CombatantBuilder) FightCharacter(spec, class string, c summary.CombatantRow) (simapi.CharacterSpec, error) {
	out := simapi.CharacterSpec{
		Name: c.Name, Class: class, Level: simapi.SimLevel,
		Gear:     gearFrom(c.Gear),
		Buffs:    BuffIDs(c.RaidBuffs, nil),
		Consumes: BuffIDs(c.Consumables, nil),
	}
	var missing []string
	switch {
	case b.Talents == nil:
		missing = append(missing, "no talent layout is loaded")
	case len(c.Talents) == 0:
		// Forever's own log dialect (version 22) does not give the
		// logs engine a flat talent tuple, so a Forever fight arrives
		// with none (logs/engine/event/special.go:322-331 and
		// layout/retail.go:242-252). A fight that does carry them
		// converts here.
		missing = append(missing, "the fight recorded no talents")
	default:
		str, err := b.Talents.String(class, c.Talents)
		if err != nil {
			missing = append(missing, err.Error())
		} else {
			out.Talents = str
		}
	}
	// The field nothing records, and the reason this always fails
	// today.
	missing = append(missing, "the fight records no race, and racials change the result")
	return out, fmt.Errorf("%w: %s", ErrNoCharacter, strings.Join(missing, "; "))
}

// gearFrom turns a fight's recorded items into the envelope's gear
// list, in the order the fight recorded them, which is the client's
// slot order and the planner's.
func gearFrom(items []event.Item) []simapi.GearSlot {
	out := make([]simapi.GearSlot, 0, len(items))
	for i, it := range items {
		if it.ID == 0 {
			continue // an empty slot
		}
		slot := simapi.GearSlot{Slot: strconv.Itoa(i), ItemID: int(it.ID)}
		if len(it.Enchants) > 0 {
			slot.Enchant = int(it.Enchants[0])
		}
		out = append(out, slot)
	}
	return out
}

// Scores writes one player's execution score on one fight. The
// rankings store satisfies it; holding it as an interface keeps sims
// from importing rankings, which imports reports, which would be a
// cycle the moment reports learns about sims.
type Scores interface {
	SetExecutionScore(ctx context.Context, reportID string, fightIndex int,
		playerKey string, score float64) error
}

// ScoreIterations is how many iterations a fight-close score runs.
// Three seconds of native compute per fight is the design's budget,
// which at roughly 1,200 iterations a second is this.
const ScoreIterations = defaultIterations

// ScoreTimeout bounds one score, so a slow engine cannot hold the
// queue.
const ScoreTimeout = 60 * time.Second

// ScoreDeps is what the scorer needs.
type ScoreDeps struct {
	Store  *Store
	Scores Scores
	Engine runner.Runner
	Build  Builder
	// EngineVersion is the pin the score was computed against.
	EngineVersion string
	Log           *slog.Logger
}

func (d ScoreDeps) logger() *slog.Logger {
	if d.Log != nil {
		return d.Log
	}
	return slog.Default()
}

// Score runs one fight's sim and writes the ratio. A spec that is not
// validated is skipped silently: the score would be a number nobody
// should read.
func Score(ctx context.Context, d ScoreDeps, t ScoreTask) error {
	if t.Role != roleDPS {
		// Tanks and healers are research problems and stay out of the
		// first cut (design, "Scope at launch"). The ingest does not
		// know that; this does.
		return nil
	}
	ok, err := d.Store.Validated(ctx, t.Spec)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	character, err := d.Build.FightCharacter(t.Spec, t.Class, t.Combatant)
	if errors.Is(err, ErrNoCharacter) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("sims: build character %s/%d: %w", t.ReportID, t.FightIndex, err)
	}
	req := simapi.SimRequest{
		EngineVersion: d.EngineVersion, Spec: t.Spec, Iterations: ScoreIterations,
		Source: simapi.CharacterSource{
			Kind: simapi.SourceFight,
			Ref:  fmt.Sprintf("%s:%d", t.ReportID, t.FightIndex),
		},
		Character: character,
		Encounter: withEncounterDefaults(simapi.EncounterSpec{DurationSec: t.DurationSec}),
	}
	runCtx, cancel := context.WithTimeout(ctx, ScoreTimeout)
	defer cancel()
	res, err := d.Engine.Run(runCtx, req, nil)
	if err != nil {
		return fmt.Errorf("sims: score %s/%d: %w", t.ReportID, t.FightIndex, err)
	}
	if res.DPS.Mean <= 0 {
		// A sim that produced nothing is not a score of zero; it is
		// no score, and the column stays null.
		d.logger().Warn("sims", "op", "score", "report", t.ReportID, "fight", t.FightIndex,
			"err", "the sim produced no dps")
		return nil
	}
	return d.Scores.SetExecutionScore(ctx, t.ReportID, t.FightIndex, t.PlayerKey,
		t.ActualDPS/res.DPS.Mean)
}

// ScoreQueueDepth is how many fights may wait to be scored. Beyond it
// a fight-close drops its score rather than blocking the ingest: the
// nightly job picks it up.
const ScoreQueueDepth = 512

// Scorer runs scores out of band, so the companion's fight-close call
// returns at once. It is parse.Worker's shape, for parse.Worker's
// reason: the consumer has to outlive the handlers that feed it, so
// Close - not a cancelled context - is the only thing that stops it,
// and the shutdown order is Shutdown then Close.
type Scorer struct {
	Deps  ScoreDeps
	queue chan ScoreTask
	done  chan struct{}
	once  sync.Once

	// mu guards closed against a Schedule that races Close, exactly
	// as parse.Worker's does: without it a Schedule racing Close is a
	// send on a closed channel, which is a panic rather than a drop.
	mu     sync.RWMutex
	closed bool
}

// NewScorer returns a scorer with an empty queue.
func NewScorer(d ScoreDeps) *Scorer {
	return &Scorer{Deps: d, queue: make(chan ScoreTask, ScoreQueueDepth), done: make(chan struct{})}
}

// Schedule queues one fight. It never blocks, and a Schedule that
// loses the race with Close drops its task the same way one that
// finds the queue full does, rather than panicking.
func (s *Scorer) Schedule(t ScoreTask) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.closed {
		s.Deps.logger().Warn("sims", "op", "score", "report", t.ReportID, "err", "scorer is closed")
		return
	}
	select {
	case s.queue <- t:
	default:
		s.Deps.logger().Warn("sims", "op", "score", "report", t.ReportID, "err", "queue is full")
	}
}

// Run consumes the queue until Close is called. A task already queued
// when Close runs is still drained and scored: closing a channel does
// not discard what is already buffered in it. Scores run on a context
// the shutdown signal does not cancel, each bounded by ScoreTimeout.
func (s *Scorer) Run(ctx context.Context) {
	defer close(s.done)
	work := context.WithoutCancel(ctx)
	for t := range s.queue {
		if err := Score(work, s.Deps, t); err != nil {
			s.Deps.logger().Error("sims", "op", "score", "report", t.ReportID,
				"fight", t.FightIndex, "err", err)
		}
	}
}

// Close stops the scorer and waits for the score in flight, if any,
// to finish. No Schedule started after Close returns reaches the
// queue, and calling it twice is fine.
func (s *Scorer) Close() {
	s.once.Do(func() {
		s.mu.Lock()
		s.closed = true
		close(s.queue)
		s.mu.Unlock()
	})
	<-s.done
}

// FightAt is the fight-close hand-off in the shape the ingest sends
// it. It is a separate type from ScoreTask only so that neither
// package has to import the other; the fields are the same, and the
// conversion in ScheduleFight compiles only while they stay that way.
type FightAt struct {
	ReportID    string
	FightIndex  int
	PlayerKey   string
	Spec        string
	Class       string
	Role        string
	ActualDPS   float64
	DurationSec int
	Combatant   summary.CombatantRow
}

// ScheduleFight queues a fight the ingest closed.
func (s *Scorer) ScheduleFight(f FightAt) {
	s.Schedule(ScoreTask(f))
}
