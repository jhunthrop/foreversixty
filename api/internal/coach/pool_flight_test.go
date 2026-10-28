package coach

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func TestPool_ReadErrorIsNotCached(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 25)
	dead, err := db.Connect(context.Background(), os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	dead.Close()
	p := &Pool{DB: dead, Now: func() time.Time { return poolNow }}
	if _, _, err := p.Reference(context.Background(), onyxia40Fury, poolReviewed); err == nil {
		t.Fatal("a read against a closed pool must fail")
	}
	p.DB = pool
	n, ref, err := p.Reference(context.Background(), onyxia40Fury, poolReviewed)
	if err != nil || n != 25 || ref == nil {
		t.Fatalf("the failure was cached: n=%d ref=%v err=%v", n, ref, err)
	}
}

func TestPool_CancelledCallerFailsOnlyItself(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 25)
	p := poolAt(pool)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, _, err := p.Reference(ctx, onyxia40Fury, poolReviewed); err == nil {
		t.Fatal("a cancelled caller must get an error")
	}
	n, ref, err := p.Reference(context.Background(), onyxia40Fury, poolReviewed)
	if err != nil || n != 25 || ref == nil {
		t.Fatalf("the next caller must read normally: n=%d ref=%v err=%v", n, ref, err)
	}
}

func TestPool_KeysDoNotShareAnEntry(t *testing.T) {
	pool := poolTestDB(t)
	arms := onyxia40Fury
	arms.Spec = "Arms"
	seedPlayers(t, pool, onyxia40Fury, 25)
	seedPlayers(t, pool, arms, 21)
	p := poolAt(pool)
	nf, _, err := p.Reference(context.Background(), onyxia40Fury, poolReviewed)
	na, ref, err2 := p.Reference(context.Background(), arms, poolReviewed)
	if err != nil || err2 != nil || nf != 25 || na != 21 || ref == nil || ref.N != 21 {
		t.Fatalf("fury %d arms %d ref %+v errs %v %v", nf, na, ref, err, err2)
	}
}

func TestPool_ConcurrentReadsOfTwoKeys(t *testing.T) {
	pool := poolTestDB(t)
	arms := onyxia40Fury
	arms.Spec = "Arms"
	seedPlayers(t, pool, onyxia40Fury, 30)
	seedPlayers(t, pool, arms, 22)
	p := poolAt(pool)
	var wg sync.WaitGroup
	errs := make(chan string, 16)
	for i := 0; i < 8; i++ {
		for k, want := range map[PoolKey]int{onyxia40Fury: 30, arms: 22} {
			wg.Add(1)
			go func() {
				defer wg.Done()
				n, ref, err := p.Reference(context.Background(), k, poolReviewed)
				if err != nil || n != want || ref == nil {
					errs <- fmt.Sprintf("%s: n=%d ref=%v err=%v", k.Spec, n, ref, err)
				}
			}()
		}
	}
	wg.Wait()
	close(errs)
	for e := range errs {
		t.Error(e)
	}
}

// Two equal-metric kills by one player resolve to the lowest report id, every refresh.
func TestPool_TieBreakIsStableAcrossRefreshes(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 19)
	seedPoolRow(t, pool, poolRow{key: onyxia40Fury, player: "us/normal/tied", metric: 50, cpm: 100, report: "rep-b"})
	seedPoolRow(t, pool, poolRow{key: onyxia40Fury, player: "us/normal/tied", metric: 50, cpm: 1, report: "rep-a"})
	for i := 0; i < 3; i++ {
		_, ref, err := poolAt(pool).Reference(context.Background(), onyxia40Fury, poolReviewed)
		if err != nil || ref == nil {
			t.Fatalf("ref=%v err=%v", ref, err)
		}
		// cpm 1..19 plus the tied player's: rep-a (1) gives 9.5, rep-b (100) would give 10.5.
		if v, _ := ref.Median.CastsPerMinOf("Bloodthirst"); v != 9.5 {
			t.Fatalf("refresh %d: median %v, want 9.5 (report rep-a)", i, v)
		}
	}
}
