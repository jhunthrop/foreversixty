package coach

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

type fixedTables struct{ tables Tables }

func (f fixedTables) TablesFor(context.Context, summary.RosterRow, int64, time.Time) (Tables, error) {
	return f.tables, nil
}

type fixedCharacters struct{}

func (fixedCharacters) Character(context.Context, string, string, int) (simapi.CharacterSpec, error) {
	return simapi.CharacterSpec{Class: "warrior", Level: 60}, nil
}

type recordingEngine struct {
	res simapi.SimResult
	got simapi.SimRequest
}

func (e *recordingEngine) Run(_ context.Context, req simapi.SimRequest, _ runner.Progress) (simapi.SimResult, error) {
	e.got = req
	return e.res, nil
}

func TestRunSimRef_StoresTheSummaryMeasuredLikeALog(t *testing.T) {
	in := furyFight(t, 60_000, fixtureBloodthirstID, []int64{0, 6000, 12000, 18000, 24000}, repeat(500, 60))
	in.Fight.Roster = []summary.RosterRow{in.Row}
	pool := simRefTestDB(t)
	engine := &recordingEngine{res: simapi.SimResult{Summary: in.Fight}}
	deps := SimRefDeps{DB: pool, Engine: engine, Characters: fixedCharacters{}, Tables: fixedTables{in.Tables}}
	if err := RunSimRef(context.Background(), deps, furyRefKey); err != nil {
		t.Fatal(err)
	}
	if engine.got.Spec != "warrior-fury" || engine.got.Encounter.DurationSec != 240 {
		t.Fatalf("request %+v", engine.got)
	}
	ref, preparing, err := (&SimRefs{DB: pool, Jobs: &jobs.Fake{}}).Get(context.Background(), furyRefKey)
	if err != nil || preparing || ref == nil {
		t.Fatalf("%+v %v %v", ref, preparing, err)
	}
	if got, want := ref.CastsPerMin["Bloodthirst"], BuildProfile(in).CastsPerMin["Bloodthirst"]; got != want || got == 0 {
		t.Fatalf("cpm %v, want %v", got, want)
	}
}

func TestRunSimRef_NoPlayerIsAnError(t *testing.T) {
	deps := SimRefDeps{DB: simRefTestDB(t), Engine: &recordingEngine{}, Characters: fixedCharacters{}, Tables: fixedTables{}}
	if err := RunSimRef(context.Background(), deps, furyRefKey); err == nil {
		t.Fatal("a summary with no player must not store a reference")
	}
}

// The reference is the published number's setup (the band's BiS gear, preset and race), so
// it fights the published numbers' target: no creature type (sim/leveling RankerEncounter).
// An empty type is the engine's Humanoid, where a rogue's Murder adds 4% the tier list
// never shows.
func TestSimRefRequest_FightsThePublishedNumbersTypeNeutralTarget(t *testing.T) {
	k := SimRefKey{EngineVersion: "abc", Build: "1.60.1.70707", Spec: "rogue-combat", Band: 60, DurationBucket: 180}
	req := simRefRequest(k, simapi.CharacterSpec{Class: "rogue", Level: 60}, time.Unix(0, 0))
	if req.Encounter.TargetType != simapi.TargetTypeUnknown {
		t.Errorf("target type = %q, want %q", req.Encounter.TargetType, simapi.TargetTypeUnknown)
	}
	if req.Encounter.DurationSec != 180 {
		t.Errorf("duration = %d, want the bucket's 180", req.Encounter.DurationSec)
	}
}
