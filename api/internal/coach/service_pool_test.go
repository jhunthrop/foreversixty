package coach

import (
	"context"
	"fmt"
	"strings"
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

// The pack version carries the pool state, so the review key moves exactly once when the
// pool fills: the "not enough logs yet" review is regenerated once, then the key holds.
func TestPackVersion_MovesOnceWhenThePoolFills(t *testing.T) {
	f := newT30(t)
	ctx := context.Background()
	persona := fixturePersona(t)
	thin, err := f.svc.PackFor(ctx, f.rep(t), t30Fight, f.furyKey)
	if err != nil || !strings.HasSuffix(thin.Version, ".pool:thin") {
		t.Fatalf("thin version %q %v", thin.Version, err)
	}
	if again, _ := f.svc.PackFor(ctx, f.rep(t), t30Fight, f.furyKey); again.Version != thin.Version {
		t.Fatalf("a thin pool keeps its version: %q -> %q", thin.Version, again.Version)
	}
	thinKey := NewReviewKey(t30Report, t30Fight, f.furyKey, thin, "m", persona)
	if err := f.svc.Store.SaveReview(ctx, StoredReview{Key: thinKey, Doc: exampleReview(), Source: SourceModel, GeneratedAt: capsNow}); err != nil {
		t.Fatal(err)
	}
	f.seedFullPool(t)
	svc := f.service(t, coachDirWith(t), false)
	full, err := svc.PackFor(ctx, f.rep(t), t30Fight, f.furyKey)
	if err != nil || full.Version != strings.TrimSuffix(thin.Version, "thin")+"ok" {
		t.Fatalf("full version %q (thin %q) %v", full.Version, thin.Version, err)
	}
	fullKey := NewReviewKey(t30Report, t30Fight, f.furyKey, full, "m", persona)
	if _, ok, _ := svc.Store.Review(ctx, fullKey); ok {
		t.Fatal("the full-pool key has no review yet: it is generated once")
	}
	cached, err := svc.PackFor(ctx, f.rep(t), t30Fight, f.furyKey)
	if err != nil || cached.Version != full.Version {
		t.Fatalf("the full-pool pack is served from the cache under the same version: %q %v", cached.Version, err)
	}
}
