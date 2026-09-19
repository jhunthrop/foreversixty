package sims

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

// fakeScores records what was written instead of touching rankings.
type fakeScores struct {
	calls []scoreCall
	err   error
}

type scoreCall struct {
	reportID   string
	fightIndex int
	playerKey  string
	score      float64
}

func (f *fakeScores) SetExecutionScore(_ context.Context, reportID string, fightIndex int,
	playerKey string, score float64) error {
	if f.err != nil {
		return f.err
	}
	f.calls = append(f.calls, scoreCall{reportID, fightIndex, playerKey, score})
	return nil
}

// alwaysBuilds is CombatantBuilder with the one field a fight does not
// record filled in, which is what a race source will make possible.
type alwaysBuilds struct{ asked int }

func (a *alwaysBuilds) FightCharacter(_, class string, c summary.CombatantRow) (simapi.CharacterSpec, error) {
	a.asked++
	return simapi.CharacterSpec{
		Name: c.Name, Race: "orc", Class: class, Level: simapi.SimLevel,
		Talents: "-0550000505021051-05",
	}, nil
}

func (h *harness) scoreDeps(engine runner.Runner, build Builder, scores Scores) ScoreDeps {
	return ScoreDeps{
		Store: h.store, Scores: scores, Engine: engine, Build: build,
		EngineVersion: testEngine, Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

// validate marks one spec validated, so the scorer will act on it.
func (h *harness) validate(t *testing.T, spec string) {
	t.Helper()
	gap := 0.02
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: spec, State: SpecValidated, MedianGap: &gap, Parses: 60,
		WorstActions: []WorstAction{}, EngineVersion: testEngine}); err != nil {
		t.Fatal(err)
	}
}

func aTask() ScoreTask {
	return ScoreTask{
		ReportID: "aaaaaaaaaaaa", FightIndex: 1, PlayerKey: "us/normal/baelgrim",
		Spec: "warrior-fury", Class: "warrior", Role: roleDPS,
		ActualDPS: 960, DurationSec: 180,
		Combatant: summary.CombatantRow{GUID: "Player-1", Name: "Baelgrim",
			Talents: []int64{105958, 105957}},
	}
}

func TestAValidatedSpecIsScoredAsActualOverSimmed(t *testing.T) {
	h := newHarness(t)
	h.validate(t, "warrior-fury")
	scores := &fakeScores{}
	// 960 actual against 1200 simmed is 0.8.
	if err := Score(t.Context(),
		h.scoreDeps(&runner.Fixture{Mean: 1200}, &alwaysBuilds{}, scores), aTask()); err != nil {
		t.Fatal(err)
	}
	if len(scores.calls) != 1 {
		t.Fatalf("%d scores written", len(scores.calls))
	}
	got := scores.calls[0]
	if got.reportID != "aaaaaaaaaaaa" || got.fightIndex != 1 || got.playerKey != "us/normal/baelgrim" {
		t.Fatalf("wrote for %+v", got)
	}
	if got.score != 0.8 {
		t.Fatalf("score %v, want actual/simmed = 0.8", got.score)
	}
}

func TestTheScoredSimRunsTheFightThatHappened(t *testing.T) {
	h := newHarness(t)
	h.validate(t, "warrior-fury")
	engine := &runner.Fixture{Mean: 1200}
	task := aTask()
	task.DurationSec = 244
	if err := Score(t.Context(), h.scoreDeps(engine, &alwaysBuilds{}, &fakeScores{}), task); err != nil {
		t.Fatal(err)
	}
	asked := engine.Asked()
	if len(asked) != 1 {
		t.Fatalf("%d runs", len(asked))
	}
	req := asked[0]
	if req.Encounter.DurationSec != 244 {
		t.Errorf("the sim ran a %ds fight, want the recorded 244", req.Encounter.DurationSec)
	}
	if req.Encounter.Targets != simapi.DefaultEncounter().Targets {
		t.Errorf("the rest of the encounter is not defaulted: %+v", req.Encounter)
	}
	if req.Source.Kind != simapi.SourceFight || req.Source.Ref != "aaaaaaaaaaaa:1" {
		t.Errorf("source %+v, want the fight it came from", req.Source)
	}
	if req.EngineVersion != testEngine || req.Iterations != defaultIterations {
		t.Errorf("request %+v", req)
	}
	if err := req.Validate(); err != nil {
		t.Errorf("the scorer built a request the envelope rejects: %v", err)
	}
}

func TestARoleTheSimulatorDoesNotModelIsDropped(t *testing.T) {
	h := newHarness(t)
	h.validate(t, "warrior-fury")
	scores, build := &fakeScores{}, &alwaysBuilds{}
	task := aTask()
	task.Role = "healer"
	if err := Score(t.Context(), h.scoreDeps(&runner.Fixture{Mean: 1200}, build, scores), task); err != nil {
		t.Fatal(err)
	}
	if len(scores.calls) != 0 || build.asked != 0 {
		t.Fatalf("a healer was scored: %d scores, %d builds", len(scores.calls), build.asked)
	}
	// The role rule lives here and nowhere else, so the ingest can
	// hand over everything it has without knowing it.
	if roleDPS != "dps" {
		t.Fatalf("roleDPS = %q", roleDPS)
	}
}

func TestTheRealBuilderSaysEveryFieldItIsMissing(t *testing.T) {
	// CombatantBuilder assembles everything a fight records, and names
	// what it cannot: with no layout loaded it is short both the
	// talents and the race, and it says both rather than stopping at
	// the first. Race is not recorded by any log dialect, so it always
	// refuses rather than inventing an orc.
	row := summary.CombatantRow{
		Name:      "Baelgrim",
		Gear:      []event.Item{{ID: 17182, Enchants: []int64{2564}}},
		RaidBuffs: []summary.AuraRef{{SpellID: 20217, Name: "Blessing of Kings"}},
		Talents:   []int64{105958},
	}
	got, err := CombatantBuilder{}.FightCharacter("warrior-fury", "warrior", row)
	if !errors.Is(err, ErrNoCharacter) {
		t.Fatalf("err = %v, want ErrNoCharacter", err)
	}
	for _, want := range []string{"talent layout", "race"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("the error should name %q: %v", want, err)
		}
	}
	// What it could build, it did: the caller's log line is worth
	// more when it shows how close the character got.
	if got.Level != simapi.SimLevel || got.Class != "warrior" {
		t.Errorf("character: %+v", got)
	}
	if len(got.Gear) != 1 || got.Gear[0].ItemID != 17182 || got.Gear[0].Enchant != 2564 {
		t.Errorf("gear: %+v", got.Gear)
	}
	if len(got.Buffs) != 1 || got.Buffs[0] != "blessing_of_kings" {
		t.Errorf("buffs: %v", got.Buffs)
	}
	if got.Talents != "" {
		t.Errorf("talents %q, want none without a layout", got.Talents)
	}
}

func TestAnUnvalidatedSpecIsNotScoredAtAll(t *testing.T) {
	h := newHarness(t)
	scores, build := &fakeScores{}, &alwaysBuilds{}
	if err := Score(t.Context(),
		h.scoreDeps(&runner.Fixture{Mean: 1200}, build, scores), aTask()); err != nil {
		t.Fatal(err)
	}
	if len(scores.calls) != 0 {
		t.Fatal("an unvalidated spec was scored")
	}
	if build.asked != 0 {
		t.Fatal("an unvalidated spec was even built for")
	}
}

func TestWithNoBuilderNothingIsScoredAndNothingFails(t *testing.T) {
	h := newHarness(t)
	h.validate(t, "warrior-fury")
	scores := &fakeScores{}
	if err := Score(t.Context(),
		h.scoreDeps(&runner.Fixture{Mean: 1200}, NoBuilder{}, scores), aTask()); err != nil {
		t.Fatalf("a missing character source is not a failure: %v", err)
	}
	if len(scores.calls) != 0 {
		t.Fatal("something was scored with no character")
	}
}

func TestASimThatProducedNothingLeavesTheColumnNull(t *testing.T) {
	h := newHarness(t)
	h.validate(t, "warrior-fury")
	scores := &fakeScores{}
	if err := Score(t.Context(),
		h.scoreDeps(&runner.Fixture{Mean: -1}, &alwaysBuilds{}, scores), aTask()); err != nil {
		t.Fatal(err)
	}
	if len(scores.calls) != 0 {
		t.Fatal("a sim with no dps produced a score of zero instead of no score")
	}
}

func TestAnEngineFailureIsReportedRatherThanSwallowed(t *testing.T) {
	h := newHarness(t)
	h.validate(t, "warrior-fury")
	if err := Score(t.Context(),
		h.scoreDeps(&runner.Fixture{Err: errAnyway}, &alwaysBuilds{}, &fakeScores{}),
		aTask()); err == nil {
		t.Fatal("an engine failure should be reported")
	}
}

func TestTheScorerDrainsWhatWasQueuedBeforeItClosed(t *testing.T) {
	h := newHarness(t)
	h.validate(t, "warrior-fury")
	scores := &fakeScores{}
	s := NewScorer(h.scoreDeps(&runner.Fixture{Mean: 1200}, &alwaysBuilds{}, scores))
	go s.Run(t.Context())
	for i := 1; i <= 3; i++ {
		task := aTask()
		task.FightIndex = i
		s.Schedule(task)
	}
	s.Close()
	if len(scores.calls) != 3 {
		t.Fatalf("%d scores after Close, want 3: Close must drain the queue", len(scores.calls))
	}
	// A Schedule after Close drops rather than panicking.
	s.Schedule(aTask())
	s.Close() // and Close twice is fine
}
