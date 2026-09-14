package builds

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func testStore(t *testing.T) *Store {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	if _, err := pool.Exec(context.Background(), `truncate builds`); err != nil {
		t.Fatal(err)
	}
	return &Store{Pool: pool}
}

func sampleBuild(t *testing.T, title string) Build {
	t.Helper()
	b, err := New(Input{
		ClassID: 1, RaceID: 1, TreeVersion: "test-1",
		PointOrder: []int{101, 101, 101, 101, 101, 201},
		Gear:       map[string]int{"head": 12640},
		Title:      title,
	})
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestStoreSaveInsertsThenReturnsTheExistingRow(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()

	want := sampleBuild(t, "Arms leveling")
	first, created, err := s.Save(ctx, want)
	if err != nil || !created {
		t.Fatalf("first save: created=%v err=%v", created, err)
	}
	if first.ID != want.ID || first.CreatedAt.IsZero() || first.Views != 0 {
		t.Fatalf("first save returned %+v, want the record's own id %s", first, want.ID)
	}

	// Same build, different title: the id is a content hash of everything
	// but the title, so this must not insert and the first title must win.
	second, created, err := s.Save(ctx, sampleBuild(t, "A different title"))
	if err != nil || created {
		t.Fatalf("second save: created=%v err=%v", created, err)
	}
	if second.ID != first.ID || second.Title != "Arms leveling" {
		t.Fatalf("second save returned %+v, want the first row", second)
	}
	if !second.CreatedAt.Equal(first.CreatedAt) {
		t.Fatalf("created_at changed: %v then %v", first.CreatedAt, second.CreatedAt)
	}
}

// TestStoreSaveRefusesAnIDCollision forces the case the id length makes
// possible but a test cannot find honestly: a second, different build
// landing on an id that is already taken. Save must refuse rather than hand
// the caller a link to the stored build, which is not theirs.
func TestStoreSaveRefusesAnIDCollision(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()

	first := sampleBuild(t, "Arms leveling")
	if _, _, err := s.Save(ctx, first); err != nil {
		t.Fatal(err)
	}

	collider := sampleBuild(t, "Fury leveling")
	collider.ID = first.ID
	collider.ClassID = 2
	collider.PointOrder = []int{201, 201}
	collider.Gear = map[string]int{"neck": 18404}

	got, created, err := s.Save(ctx, collider)
	if !errors.Is(err, ErrIDCollision) {
		t.Fatalf("err = %v, want ErrIDCollision; got %+v created=%v", err, got, created)
	}
	if created || got.ID != "" {
		t.Fatalf("a refused save must return no record: %+v created=%v", got, created)
	}

	// The stored row is untouched.
	stored, err := s.Get(ctx, first.ID)
	if err != nil {
		t.Fatal(err)
	}
	if stored.ClassID != 1 || stored.Title != "Arms leveling" {
		t.Fatalf("the stored build changed: %+v", stored)
	}
}

// TestStoreSaveAllowsATitleOnlyDifference guards the other side of the
// collision check: a resave that differs only in its title is the same
// build and must still dedupe rather than be read as a collision. The
// broader dedupe behaviour is covered above; this pins the comparison
// itself to the hashed fields.
func TestStoreSaveAllowsATitleOnlyDifference(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()

	first, _, err := s.Save(ctx, sampleBuild(t, ""))
	if err != nil {
		t.Fatal(err)
	}
	again, created, err := s.Save(ctx, sampleBuild(t, "Now with a title"))
	if err != nil || created {
		t.Fatalf("created=%v err=%v, want a dedupe", created, err)
	}
	if again.ID != first.ID || again.Title != "" {
		t.Fatalf("got %+v, want the first row back", again)
	}
}

func TestStoreGetRoundTripsEveryColumn(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	saved, _, err := s.Save(ctx, sampleBuild(t, "Arms leveling"))
	if err != nil {
		t.Fatal(err)
	}

	got, err := s.Get(ctx, saved.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.ClassID != 1 || got.RaceID != 1 || got.TreeVersion != "test-1" {
		t.Fatalf("got %+v", got)
	}
	if len(got.PointOrder) != 6 || got.PointOrder[0] != 101 || got.PointOrder[5] != 201 {
		t.Fatalf("point_order = %v", got.PointOrder)
	}
	if got.Gear["head"] != 12640 {
		t.Fatalf("gear = %v", got.Gear)
	}
	if got.Title != "Arms leveling" || got.Views != 0 || got.CreatedAt.IsZero() {
		t.Fatalf("got %+v", got)
	}
}

func TestStoreGetReportsAnUnknownID(t *testing.T) {
	s := testStore(t)
	if _, err := s.Get(context.Background(), "nosuchid"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

func TestStoreSaveKeepsAnEmptyTitleNull(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	saved, _, err := s.Save(ctx, sampleBuild(t, ""))
	if err != nil {
		t.Fatal(err)
	}
	var isNull bool
	if err := s.Pool.QueryRow(ctx, `select title is null from builds where id = $1`, saved.ID).Scan(&isNull); err != nil {
		t.Fatal(err)
	}
	if !isNull {
		t.Fatal("an empty title must be stored as NULL, not an empty string")
	}
}

func TestStoreAddViewsIncrementsEachID(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	one, _, err := s.Save(ctx, sampleBuild(t, "one"))
	if err != nil {
		t.Fatal(err)
	}
	other, err := New(Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", PointOrder: []int{201}})
	if err != nil {
		t.Fatal(err)
	}
	two, _, err := s.Save(ctx, other)
	if err != nil {
		t.Fatal(err)
	}

	if err := s.AddViews(ctx, map[string]int64{one.ID: 3, two.ID: 1}); err != nil {
		t.Fatal(err)
	}
	if err := s.AddViews(ctx, nil); err != nil {
		t.Fatalf("an empty batch must be a no-op: %v", err)
	}

	got, err := s.Get(ctx, one.ID)
	if err != nil || got.Views != 3 {
		t.Fatalf("views = %d err=%v, want 3", got.Views, err)
	}
	got, err = s.Get(ctx, two.ID)
	if err != nil || got.Views != 1 {
		t.Fatalf("views = %d err=%v, want 1", got.Views, err)
	}
}
