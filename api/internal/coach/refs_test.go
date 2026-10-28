package coach

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/jobs"
)

type fakeValidated map[string]bool

func (f fakeValidated) Validated(_ context.Context, spec string) (bool, error) { return f[spec], nil }

var fightAt = time.Date(2026, 12, 9, 1, 0, 0, 0, time.UTC)

// buildAt is a manifest list with 1.60 fetched at the fight.
func buildAt(fought time.Time) string {
	if fought.Before(fightAt) {
		return "1.59"
	}
	return "1.60"
}

// resolverFor is a resolver over an empty pool and a sim cache whose job runner records
// what it was asked to queue.
func resolverFor(t *testing.T, validated fakeValidated) (*Resolver, *jobs.Fake) {
	t.Helper()
	db := simRefTestDB(t)
	fake := &jobs.Fake{}
	return &Resolver{
		Pool: &Pool{DB: db}, Sims: &SimRefs{DB: db, Jobs: fake}, Validated: validated,
		EngineVersion: "e1", BuildAt: buildAt,
	}, fake
}

func TestResolveRefs_UnvalidatedSpecSkipsSim(t *testing.T) {
	r, fake := resolverFor(t, fakeValidated{"warrior-fury": false})
	refs, err := r.Refs(context.Background(), onyxia40Fury, "us/normal/x", RoleDPS, 60, 232_000, fightAt)
	if err != nil || refs.Sim != nil || refs.SimPreparing || len(fake.Ran()) != 0 {
		t.Fatalf("%+v %v ran %v", refs, err, fake.Ran())
	}
}

func TestResolveRefs_HealersAndTanksNeverSimmed(t *testing.T) {
	for _, role := range []string{RoleHealer, RoleTank} {
		r, fake := resolverFor(t, fakeValidated{"warrior-fury": true})
		refs, err := r.Refs(context.Background(), onyxia40Fury, "us/normal/x", role, 60, 232_000, fightAt)
		if err != nil || refs.Sim != nil || refs.SimPreparing || len(fake.Ran()) != 0 {
			t.Fatalf("%s: %+v %v ran %v", role, refs, err, fake.Ran())
		}
	}
}

func TestResolveRefs_NoLevelNoSim(t *testing.T) {
	r, fake := resolverFor(t, fakeValidated{"warrior-fury": true})
	refs, err := r.Refs(context.Background(), onyxia40Fury, "us/normal/x", RoleDPS, 0, 232_000, fightAt)
	if err != nil || refs.Sim != nil || refs.SimPreparing || len(fake.Ran()) != 0 {
		t.Fatalf("%+v %v ran %v", refs, err, fake.Ran())
	}
	if !refs.SimNoLevel {
		t.Fatal("a validated DPS spec with no level says the sim went unchecked (CannotLevelSim)")
	}
	r, _ = resolverFor(t, fakeValidated{"warrior-fury": false})
	if refs, err := r.Refs(context.Background(), onyxia40Fury, "us/normal/x", RoleDPS, 0, 232_000, fightAt); err != nil || refs.SimNoLevel {
		t.Fatalf("an unvalidated spec is never simmed, so nothing went unchecked: %+v %v", refs, err)
	}
}

func TestResolveRefs_ValidatedDPSWithoutPoolPreparesTheSim(t *testing.T) {
	r, fake := resolverFor(t, fakeValidated{"warrior-fury": true})
	refs, err := r.Refs(context.Background(), onyxia40Fury, "us/normal/x", RoleDPS, 63, 232_000, fightAt)
	if err != nil || refs.Sim != nil || !refs.SimPreparing || len(fake.Ran()) != 1 {
		t.Fatalf("%+v %v ran %v", refs, err, fake.Ran())
	}
	want := SimRefKey{EngineVersion: "e1", Build: "1.60", Spec: "warrior-fury", Band: 60, DurationBucket: 240}
	if got, _ := ParseSimRefKey(fake.Ran()[0][1:]); got != want {
		t.Fatalf("queued %+v, want %+v", got, want)
	}
}

const furyRefRow = `insert into coach_sim_refs (engine_version, build, spec, band, duration_bucket, summary)
	values ('e1', '1.60', 'warrior-fury', 60, 240, '{"cpm":{"Bloodthirst":9.5,"Whirlwind":4},"uptime":{}}')`

func TestResolveRefs_SimFillsWhatThePoolLacks(t *testing.T) {
	r, fake := resolverFor(t, fakeValidated{"warrior-fury": true})
	seedPlayers(t, r.Pool.DB, onyxia40Fury, 25)
	if _, err := r.Sims.DB.Exec(context.Background(), furyRefRow); err != nil {
		t.Fatal(err)
	}
	refs, err := r.Refs(context.Background(), onyxia40Fury, poolReviewed, RoleDPS, 60, 232_000, fightAt)
	if err != nil || refs.Pool == nil || refs.Sim == nil || refs.SimPreparing || len(fake.Ran()) != 0 {
		t.Fatalf("%+v %v ran %v", refs, err, fake.Ran())
	}
	in := furyInput(t)
	in.Refs = refs
	pool := func(p Profile) (float64, bool) { return p.CastsPerMinOf("Whirlwind") }
	sim := func(s SimRef) (float64, bool) { v, ok := s.CastsPerMin["Whirlwind"]; return v, ok }
	if ref, b, thin := comparison(in, pool, sim, true); ref != 4 || b == nil || b.Kind != BasisSim || thin != "" {
		t.Fatalf("whirlwind: %v %+v %q", ref, b, thin)
	}
	pool = func(p Profile) (float64, bool) { return p.CastsPerMinOf("Bloodthirst") }
	if _, b, _ := comparison(in, pool, sim, true); b == nil || b.Kind != BasisTopLogs {
		t.Fatalf("bloodthirst should keep the top logs: %+v", b)
	}
}

func TestResolveRefs_TheBuildIsTheOneInForceAtTheFight(t *testing.T) {
	r, fake := resolverFor(t, fakeValidated{"warrior-fury": true})
	before, after := fightAt.Add(-time.Hour), fightAt.Add(time.Hour)
	for _, at := range []time.Time{before, after} {
		if _, err := r.Refs(context.Background(), onyxia40Fury, "us/normal/x", RoleDPS, 60, 232_000, at); err != nil {
			t.Fatal(err)
		}
	}
	builds := []string{}
	for _, call := range fake.Ran() {
		key, err := ParseSimRefKey(call[1:])
		if err != nil {
			t.Fatal(err)
		}
		builds = append(builds, key.Build)
	}
	if len(builds) != 2 || builds[0] != "1.59" || builds[1] != "1.60" {
		t.Fatalf("builds %v", builds)
	}
}

func TestLevelBand(t *testing.T) {
	for level, want := range map[int]int{0: 0, 19: 0, 20: 20, 39: 30, 60: 60, 63: 60} {
		if got := LevelBand(level); got != want {
			t.Errorf("LevelBand(%d) = %d, want %d", level, got, want)
		}
	}
}
