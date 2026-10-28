package coach

import (
	"context"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/jobs"
)

type fakeValidated map[string]bool

func (f fakeValidated) Validated(_ context.Context, spec string) (bool, error) { return f[spec], nil }

// resolverFor is a resolver over an empty pool and a sim cache whose job runner records
// what it was asked to queue.
func resolverFor(t *testing.T, validated fakeValidated) (*Resolver, *jobs.Fake) {
	t.Helper()
	db := simRefTestDB(t)
	fake := &jobs.Fake{}
	return &Resolver{
		Pool: &Pool{DB: db}, Sims: &SimRefs{DB: db, Jobs: fake}, Validated: validated,
		EngineVersion: "e1", Build: "1.60",
	}, fake
}

func TestResolveRefs_UnvalidatedSpecSkipsSim(t *testing.T) {
	r, fake := resolverFor(t, fakeValidated{"warrior-fury": false})
	refs, err := r.Refs(context.Background(), onyxia40Fury, "us/normal/x", RoleDPS, 60, 232_000)
	if err != nil || refs.Sim != nil || refs.SimPreparing || len(fake.Ran()) != 0 {
		t.Fatalf("%+v %v ran %v", refs, err, fake.Ran())
	}
}

func TestResolveRefs_HealersAndTanksNeverSimmed(t *testing.T) {
	for _, role := range []string{RoleHealer, RoleTank} {
		r, fake := resolverFor(t, fakeValidated{"warrior-fury": true})
		refs, err := r.Refs(context.Background(), onyxia40Fury, "us/normal/x", role, 60, 232_000)
		if err != nil || refs.Sim != nil || refs.SimPreparing || len(fake.Ran()) != 0 {
			t.Fatalf("%s: %+v %v ran %v", role, refs, err, fake.Ran())
		}
	}
}

func TestResolveRefs_NoLevelNoSim(t *testing.T) {
	r, fake := resolverFor(t, fakeValidated{"warrior-fury": true})
	refs, err := r.Refs(context.Background(), onyxia40Fury, "us/normal/x", RoleDPS, 0, 232_000)
	if err != nil || refs.Sim != nil || refs.SimPreparing || len(fake.Ran()) != 0 {
		t.Fatalf("%+v %v ran %v", refs, err, fake.Ran())
	}
}

func TestResolveRefs_ValidatedDPSWithoutPoolPreparesTheSim(t *testing.T) {
	r, fake := resolverFor(t, fakeValidated{"warrior-fury": true})
	refs, err := r.Refs(context.Background(), onyxia40Fury, "us/normal/x", RoleDPS, 63, 232_000)
	if err != nil || refs.Sim != nil || !refs.SimPreparing || len(fake.Ran()) != 1 {
		t.Fatalf("%+v %v ran %v", refs, err, fake.Ran())
	}
	want := SimRefKey{EngineVersion: "e1", Build: "1.60", Spec: "warrior-fury", Band: 60, DurationBucket: 240}
	if got, _ := ParseSimRefKey(fake.Ran()[0][1:]); got != want {
		t.Fatalf("queued %+v, want %+v", got, want)
	}
}

func TestResolveRefs_PoolAnswersBeforeTheSim(t *testing.T) {
	r, fake := resolverFor(t, fakeValidated{"warrior-fury": true})
	seedPlayers(t, r.Pool.DB, onyxia40Fury, 25)
	refs, err := r.Refs(context.Background(), onyxia40Fury, poolReviewed, RoleDPS, 60, 232_000)
	if err != nil || refs.Pool == nil || refs.Sim != nil || refs.SimPreparing || len(fake.Ran()) != 0 {
		t.Fatalf("%+v %v ran %v", refs, err, fake.Ran())
	}
}

func TestLevelBand(t *testing.T) {
	for level, want := range map[int]int{0: 0, 19: 0, 20: 20, 39: 30, 60: 60, 63: 60} {
		if got := LevelBand(level); got != want {
			t.Errorf("LevelBand(%d) = %d, want %d", level, got, want)
		}
	}
}
