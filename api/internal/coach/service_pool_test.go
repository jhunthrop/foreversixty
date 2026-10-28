package coach

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
)

// seedFullPool fills the t30 fight's comparison bracket with PoolMinPlayers other players.
func (f *t30) seedFullPool(t *testing.T) {
	t.Helper()
	if err := db.EnsureCoachProfilesPartition(context.Background(), f.pool, poolFoughtAt); err != nil {
		t.Fatal(err)
	}
	in := furyThreeFindings(t)
	k := PoolKey{EncounterID: in.Fight.EncounterID, Difficulty: in.Fight.Difficulty, Size: 40,
		Spec: furyRowOf(t, in.Fight.Roster).Spec, Phase: phase.At(time.UnixMilli(capsNow.UnixMilli()))}
	for i := 1; i <= PoolMinPlayers; i++ {
		seedPoolRow(t, f.pool, poolRow{key: k, player: fmt.Sprintf("us/normal/t30-pool-%d", i), metric: float64(i), cpm: float64(i)})
	}
}

// A pack whose pool is below the floor is served but never stored, so the next view
// rebuilds it and picks up the top-logs basis once the pool fills.
func TestPackFor_AThinPoolPackIsServedButNotCached(t *testing.T) {
	f := newT30(t)
	ctx := context.Background()
	thin, err := f.svc.PackFor(ctx, f.rep(t), t30Fight, f.furyKey)
	if err != nil || thin.Status != StatusReady || len(thin.Findings) == 0 {
		t.Fatalf("served: %+v %v", thin, err)
	}
	if _, cached, _ := f.svc.Store.Pack(ctx, t30Report, t30Fight, f.furyKey, thin.Version); cached {
		t.Fatal("a thin-pool pack is not stored")
	}
	reads := f.bucket.reads()
	f.seedFullPool(t)
	svc := f.service(t, coachDirWith(t), false) // a fresh pool cache, as the next hour's read
	full, err := svc.PackFor(ctx, f.rep(t), t30Fight, f.furyKey)
	if err != nil || f.bucket.reads() != reads+1 {
		t.Fatalf("the next view rebuilds: reads %d -> %d, %v", reads, f.bucket.reads(), err)
	}
	if _, cached, _ := svc.Store.Pack(ctx, t30Report, t30Fight, f.furyKey, full.Version); !cached {
		t.Fatal("a pack with a full pool is stored")
	}
}
