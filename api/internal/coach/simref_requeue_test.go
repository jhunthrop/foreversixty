package coach

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

func TestSimRefs_RequeuesAfterAnHour(t *testing.T) {
	fake := &jobs.Fake{}
	now := time.Date(2026, 12, 9, 1, 0, 0, 0, time.UTC)
	s := &SimRefs{DB: simRefTestDB(t), Jobs: fake, Now: func() time.Time { return now }}
	get := func() {
		t.Helper()
		if _, preparing, err := s.Get(context.Background(), furyRefKey); err != nil || !preparing {
			t.Fatalf("%v %v", preparing, err)
		}
	}
	get()
	now = now.Add(SimRefRequeueAfter - time.Minute)
	get()
	if len(fake.Ran()) != 1 {
		t.Fatalf("queued again inside the hour: %v", fake.Ran())
	}
	now = now.Add(2 * time.Minute)
	get()
	if len(fake.Ran()) != 2 {
		t.Fatalf("not queued after the hour: %v", fake.Ran())
	}
}

func TestSimRefs_FailedQueueIsRetriedNotRemembered(t *testing.T) {
	fake := &jobs.Fake{Err: errors.New("no cloud")}
	s := &SimRefs{DB: simRefTestDB(t), Jobs: fake}
	if _, _, err := s.Get(context.Background(), furyRefKey); err == nil {
		t.Fatal("expected the failure")
	}
	fake.Err = nil
	if _, preparing, err := s.Get(context.Background(), furyRefKey); err != nil || !preparing || len(fake.Ran()) != 1 {
		t.Fatalf("%v %v ran %v", preparing, err, fake.Ran())
	}
}

func TestRunSimRef_UptimeIsStoredAsAPercent(t *testing.T) {
	in := furyRendLow(t, Refs{})
	in.Fight.Roster = []summary.RosterRow{in.Row}
	pool := simRefTestDB(t)
	deps := SimRefDeps{DB: pool, Engine: &recordingEngine{res: simapi.SimResult{Summary: in.Fight}},
		Characters: fixedCharacters{}, Tables: fixedTables{in.Tables}}
	if err := RunSimRef(context.Background(), deps, furyRefKey); err != nil {
		t.Fatal(err)
	}
	ref, _, err := (&SimRefs{DB: pool, Jobs: &jobs.Fake{}}).Get(context.Background(), furyRefKey)
	if err != nil || ref == nil {
		t.Fatalf("%+v %v", ref, err)
	}
	want, ok := BuildProfile(in).UptimeOf(f2RendID)
	if !ok || want < 15 || want > 25 {
		t.Fatalf("the fixture's Rend uptime is %v (%v), a fifth of the fight in percent", want, ok)
	}
	if got := ref.Uptime[f2RendID]; got != want {
		t.Fatalf("stored uptime %v, want %v percent", got, want)
	}
}
