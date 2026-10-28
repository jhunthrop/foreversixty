package coach

import (
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const t28Report = "t28-report"

// reviewPool is capsPool plus the report every t28 key points at.
func reviewPool(t *testing.T) (*pgxpool.Pool, *Store) {
	t.Helper()
	pool := capsPool(t)
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `delete from reports where id = $1`, t28Report); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into reports (id, visibility, status) values ($1, 'public', 'complete')`, t28Report); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(context.Background(), `delete from reports where id = $1`, t28Report) })
	return pool, &Store{DB: pool}
}

func t28Key() ReviewKey {
	return ReviewKey{ReportID: t28Report, FightIndex: 3, PlayerKey: "us/normal/p", PackVersion: "pack-v1.example",
		ModelVersion: "m.abc", PersonaVersion: "persona-x"}
}

func TestReviewStore_RoundTripAndUpsert(t *testing.T) {
	pool, store := reviewPool(t)
	ctx := context.Background()
	user := capsUser(t, pool, "t28-store@example.com")
	row := StoredReview{Key: t28Key(), Doc: exampleReview(), Source: SourceRules, Reason: "validation",
		TokensIn: 14000, TokensOut: 1600, CostUSD: 0.004, CostEstimated: true, ModelSlug: "m", RequestID: "g2",
		AccountID: user, GeneratedAt: capsNow, Validation: []Problem{{Kind: "number", Value: "12", Near: "12 casts"}}}
	if err := store.SaveReview(ctx, row); err != nil {
		t.Fatal(err)
	}
	got, ok, err := store.Review(ctx, row.Key)
	if err != nil || !ok {
		t.Fatalf("ok %v err %v", ok, err)
	}
	got.GeneratedAt = got.GeneratedAt.UTC()
	if !reflect.DeepEqual(got, row) {
		t.Fatalf("got %+v\nwant %+v", got, row)
	}
	row.Source, row.Reason, row.Validation = SourceModel, "", nil
	if err := store.SaveReview(ctx, row); err != nil {
		t.Fatal(err)
	}
	got, _, _ = store.Review(ctx, row.Key)
	if got.Source != SourceModel || got.Reason != "" || len(got.Validation) != 0 {
		t.Fatalf("the upsert replaces the row: %+v", got)
	}
}

func TestReviewStore_MissingAndBadSource(t *testing.T) {
	_, store := reviewPool(t)
	ctx := context.Background()
	if _, ok, err := store.Review(ctx, t28Key()); ok || err != nil {
		t.Fatalf("ok %v err %v", ok, err)
	}
	if err := store.SaveReview(ctx, StoredReview{Key: t28Key(), Source: "guess"}); err == nil {
		t.Fatal("an unknown source is refused")
	}
}

func TestReviewStore_KeyCarriesModelAndPersonaVersions(t *testing.T) {
	pool, store := reviewPool(t)
	ctx := context.Background()
	persona := fixturePersona(t)
	edited, err := ParsePersona([]byte(strings.Replace(fixture(t), "Short sentences.", "Short sentence.", 1)))
	if err != nil || edited.Version == persona.Version {
		t.Fatalf("a persona body change moves its version: %v", err)
	}
	keyOf := func(slug string, p Persona) ReviewKey {
		return NewReviewKey(t28Report, 3, "us/normal/p", examplePack(), slug, p)
	}
	base := keyOf("m", persona)
	if ModelVersion("m") != base.ModelVersion || keyOf("m", edited).ModelVersion != base.ModelVersion {
		t.Fatal("the model version ignores the persona body")
	}
	saved := 0
	for _, k := range []ReviewKey{base, keyOf("m", edited), keyOf("m2", persona), keyOf("m", persona)} {
		if _, ok, _ := store.Review(ctx, k); ok {
			continue // an unchanged triple reuses the stored review
		}
		if err := store.SaveReview(ctx, StoredReview{Key: k, Source: SourceRules, GeneratedAt: capsNow}); err != nil {
			t.Fatal(err)
		}
		saved++
	}
	var n int
	if err := pool.QueryRow(ctx, `select count(*) from coach_reviews where report_id = $1`, t28Report).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 3 || saved != 3 {
		t.Fatalf("rows %d saved %d, want 3: base, persona change, model change; the repeat reuses", n, saved)
	}
}

func TestTakeLease_OneAtATime(t *testing.T) {
	pool, store := reviewPool(t)
	ctx := context.Background()
	k := t28Key()
	until := time.Now().Add(5 * time.Minute)
	if ok, err := store.TakeLease(ctx, k, "h1", until); !ok || err != nil {
		t.Fatalf("first: %v %v", ok, err)
	}
	if ok, err := store.TakeLease(ctx, k, "h2", until); ok || err != nil {
		t.Fatalf("second while the lease holds: %v %v", ok, err)
	}
	if _, err := pool.Exec(ctx, `update coach_review_leases set leased_until = now() - interval '1 second' where report_id = $1`, k.ReportID); err != nil {
		t.Fatal(err)
	}
	if ok, err := store.TakeLease(ctx, k, "h2", until); !ok || err != nil {
		t.Fatalf("after the lease passes: %v %v", ok, err)
	}
	if err := store.ReleaseLease(ctx, k, "h1"); err != nil {
		t.Fatal(err)
	}
	if ok, _ := store.TakeLease(ctx, k, "h3", until); ok {
		t.Fatal("the old holder cannot release the new holder's lease")
	}
	if err := store.ReleaseLease(ctx, k, "h2"); err != nil {
		t.Fatal(err)
	}
	if ok, _ := store.TakeLease(ctx, k, "h3", until); !ok {
		t.Fatal("the holder's release frees the key")
	}
	if _, err := store.TakeLease(ctx, k, "", until); err == nil {
		t.Fatal("a lease needs a holder")
	}
}
