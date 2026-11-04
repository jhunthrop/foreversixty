// api/internal/bnetimport/launch_capacity_test.go
package bnetimport

import (
	"context"
	"errors"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
)

// TestAcquireEnrichSlotCapsConcurrentEnrichments: with a cap of one, a second
// enrichment waits until the first gives its slot back, and gives up when its
// context ends first.
func TestAcquireEnrichSlotCapsConcurrentEnrichments(t *testing.T) {
	svc := &Service{MaxConcurrentEnrichments: 1}
	release, err := svc.acquireEnrichSlot(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	if _, err := svc.acquireEnrichSlot(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("second acquire = %v, want it to wait out its deadline", err)
	}

	release()
	second, err := svc.acquireEnrichSlot(context.Background())
	if err != nil {
		t.Fatalf("acquire after release = %v, want the freed slot", err)
	}
	second()
}

// TestAcquireEnrichSlotWithoutACapNeverWaits: the nightly refresh sets no cap.
func TestAcquireEnrichSlotWithoutACapNeverWaits(t *testing.T) {
	svc := &Service{}
	for range 3 {
		release, err := svc.acquireEnrichSlot(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		defer release()
	}
}

// TestImportAccountStopsEnrichingOnA429: once Battle.net answers 429 the run
// stops, reports every unenriched row as pending and flags the rate limit, and
// never calls Blizzard for the characters after it.
func TestImportAccountStopsEnrichingOnA429(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	f := newBlizzardFixture(t)
	f.handlers[http.MethodGet+" "+accountCharactersPath] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"wow_accounts":[{"id":1,"characters":[
			{"name":"Mainone","id":11,"realm":{"slug":"whitemane","name":"Whitemane"},
			 "playable_class":{"name":"Mage"},"faction":{"type":"HORDE"},"level":60},
			{"name":"Alttwo","id":12,"realm":{"slug":"whitemane","name":"Whitemane"},
			 "playable_class":{"name":"Rogue"},"faction":{"type":"HORDE"},"level":20}
		]}]}`))
	}
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"realms": []any{}})
	var profileCalls atomic.Int32
	for _, name := range []string{"mainone", "alttwo"} {
		f.handlers[http.MethodGet+" /profile/wow/character/whitemane/"+name+"?namespace=profile-classic1x-us"] =
			func(w http.ResponseWriter, r *http.Request) {
				profileCalls.Add(1)
				w.WriteHeader(http.StatusTooManyRequests)
			}
	}

	svc := newTestService(t, pool, f)
	svc.MaxConcurrentEnrichments = 1
	summary, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token")
	if !errors.Is(err, bnetapi.ErrRateLimited) {
		t.Fatalf("err = %v, want bnetapi.ErrRateLimited", err)
	}
	if !summary.RateLimited || summary.Pending != 2 || summary.Characters != 2 {
		t.Fatalf("summary = %+v, want both rows listed, both pending, rate limited", summary)
	}
	if got := profileCalls.Load(); got != 1 {
		t.Fatalf("profile calls = %d, want 1: the run must stop at the first 429", got)
	}
}
