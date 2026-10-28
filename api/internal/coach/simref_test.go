package coach

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/jobs"
)

func simRefTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()
	pool := poolTestDB(t)
	if _, err := pool.Exec(context.Background(), `truncate coach_sim_refs`); err != nil {
		t.Fatal(err)
	}
	return pool
}

var furyRefKey = SimRefKey{EngineVersion: "e1", Build: "1.60", Spec: "warrior-fury", Band: 60, DurationBucket: 240}

func TestDurationBucket(t *testing.T) {
	for ms, want := range map[int64]int{232_000: 240, 14_000: 30, 0: 30, 44_000: 30, 46_000: 60, 90_000: 90} {
		if got := DurationBucket(ms); got != want {
			t.Errorf("DurationBucket(%d) = %d, want %d", ms, got, want)
		}
	}
}

func TestSimRefs_MissQueuesOnceAndPrepares(t *testing.T) {
	fake := &jobs.Fake{}
	s := &SimRefs{DB: simRefTestDB(t), Jobs: fake}
	for i := 0; i < 2; i++ {
		ref, preparing, err := s.Get(context.Background(), furyRefKey)
		if err != nil || ref != nil || !preparing {
			t.Fatalf("get %d: %+v %v %v", i, ref, preparing, err)
		}
	}
	ran := fake.Ran()
	if len(ran) != 1 || ran[0][0] != SimRefJobCommand {
		t.Fatalf("ran %v", ran)
	}
	parsed, err := ParseSimRefKey(ran[0][1:])
	if err != nil || parsed != furyRefKey {
		t.Fatalf("job args %v parse to %+v %v", ran[0], parsed, err)
	}
}

func TestSimRefs_QueueFailureIsPreparingNotAnError(t *testing.T) {
	for name, s := range map[string]*SimRefs{
		"launch error": {DB: simRefTestDB(t), Jobs: &jobs.Fake{Err: errors.New("no cloud")}},
		"no runner":    {DB: simRefTestDB(t)},
	} {
		if ref, preparing, err := s.Get(context.Background(), furyRefKey); err != nil || !preparing || ref != nil {
			t.Fatalf("%s: a job that cannot be queued answers preparing: %v %v %v", name, ref, preparing, err)
		}
	}
}

func TestSimRefs_HitReturnsTheStoredSummary(t *testing.T) {
	pool := simRefTestDB(t)
	_, err := pool.Exec(context.Background(), `insert into coach_sim_refs
		(engine_version, build, spec, band, duration_bucket, summary) values ('e1','1.60','warrior-fury',60,240,
		'{"cpm":{"bloodthirst":9.5},"uptime":{"12292":87.5}}')`)
	if err != nil {
		t.Fatal(err)
	}
	fake := &jobs.Fake{}
	ref, preparing, err := (&SimRefs{DB: pool, Jobs: fake}).Get(context.Background(), furyRefKey)
	if err != nil || preparing || ref == nil {
		t.Fatalf("%+v %v %v", ref, preparing, err)
	}
	if ref.Band != 60 || ref.CastsPerMin["bloodthirst"] != 9.5 || ref.Uptime[12292] != 87.5 || len(fake.Ran()) != 0 {
		t.Fatalf("ref %+v ran %v", ref, fake.Ran())
	}
}

func TestParseSimRefKey_RejectsBadArguments(t *testing.T) {
	for _, args := range [][]string{{"a"}, {"e", "b", "s", "x", "30"}, {"e", "b", "s", "60", "y"}} {
		if _, err := ParseSimRefKey(args); err == nil {
			t.Errorf("%v parsed", args)
		}
	}
}
