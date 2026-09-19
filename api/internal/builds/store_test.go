package builds

import (
	"context"
	"errors"
	"math"
	"os"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
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
	if _, err := pool.Exec(context.Background(), `truncate builds, users cascade`); err != nil {
		t.Fatal(err)
	}
	return &Store{Pool: pool}
}

// storeWithUser is testStore with a signed-in account already on hand,
// for the tests that save builds on somebody's behalf.
func storeWithUser(t *testing.T) (*Store, int64) {
	t.Helper()
	s := testStore(t)
	u, err := (&auth.Store{Pool: s.Pool}).UpsertEmailUser(context.Background(), "owner@example.com")
	if err != nil {
		t.Fatal(err)
	}
	return s, u.ID
}

// anotherUser adds a second account to s's database, for the tests that
// check one saver's claim does not reach another player's list.
func anotherUser(t *testing.T, s *Store) int64 {
	t.Helper()
	u, err := (&auth.Store{Pool: s.Pool}).UpsertEmailUser(context.Background(), "second@example.com")
	if err != nil {
		t.Fatal(err)
	}
	return u.ID
}

// aBuild is a fixed, valid build for the ownership tests, which read
// better naming the build than the title alone.
func aBuild(title string) Build {
	b, _ := New(Input{
		ClassID: 1, RaceID: 1, TreeVersion: "test-1",
		PointOrder: []int{101, 101, 101, 101, 101, 201},
		Gear:       map[string]int{"head": 12640},
		Title:      title,
	})
	return b
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
	first, created, err := s.Save(ctx, want, nil)
	if err != nil || !created {
		t.Fatalf("first save: created=%v err=%v", created, err)
	}
	if first.ID != want.ID || first.CreatedAt.IsZero() || first.Views != 0 {
		t.Fatalf("first save returned %+v, want the record's own id %s", first, want.ID)
	}

	// Same build, different title: the id is a content hash of everything
	// but the title, so this must not insert and the first title must win.
	second, created, err := s.Save(ctx, sampleBuild(t, "A different title"), nil)
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
	if _, _, err := s.Save(ctx, first, nil); err != nil {
		t.Fatal(err)
	}

	collider := sampleBuild(t, "Fury leveling")
	collider.ID = first.ID
	collider.ClassID = 2
	collider.PointOrder = []int{201, 201}
	collider.Gear = map[string]int{"neck": 18404}

	got, created, err := s.Save(ctx, collider, nil)
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

	first, _, err := s.Save(ctx, sampleBuild(t, ""), nil)
	if err != nil {
		t.Fatal(err)
	}
	again, created, err := s.Save(ctx, sampleBuild(t, "Now with a title"), nil)
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
	saved, _, err := s.Save(ctx, sampleBuild(t, "Arms leveling"), nil)
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
	saved, _, err := s.Save(ctx, sampleBuild(t, ""), nil)
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

func TestStoreGetManyReadsSeveralRowsInOneQuery(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	one, _, err := s.Save(ctx, sampleBuild(t, "one"), nil)
	if err != nil {
		t.Fatal(err)
	}
	other, err := New(Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", PointOrder: []int{201}})
	if err != nil {
		t.Fatal(err)
	}
	two, _, err := s.Save(ctx, other, nil)
	if err != nil {
		t.Fatal(err)
	}

	// A missing id is simply absent from the result, not an error: the
	// caller (the addon inbox) treats that as "no code for this entry".
	got, err := s.GetMany(ctx, []string{one.ID, two.ID, "nosuchid"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d rows, want 2: %+v", len(got), got)
	}
	if got[one.ID].Title != "one" || got[two.ID].ClassID != 1 {
		t.Fatalf("got = %+v", got)
	}
}

func TestStoreGetManyOfNoIDsIsEmptyAndMakesNoQuery(t *testing.T) {
	s := testStore(t)
	got, err := s.GetMany(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Fatalf("got %+v, want empty", got)
	}
}

// TestStoreSaveRoundTripsASixDigitTalentID pins the fix for the beta
// client's build failures: its trees use trait node ids running into six
// digits (see data/builds/1.60.1.69893/talents/hunter.json), past a
// smallint's range, which made every save on those trees 500.
func TestStoreSaveRoundTripsASixDigitTalentID(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()

	b, err := New(Input{
		ClassID: 1, RaceID: 1, TreeVersion: "test-1",
		PointOrder: []int{104960},
	})
	if err != nil {
		t.Fatal(err)
	}
	saved, created, err := s.Save(ctx, b, nil)
	if err != nil || !created {
		t.Fatalf("save: created=%v err=%v", created, err)
	}

	got, err := s.Get(ctx, saved.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(got.PointOrder, []int{104960}) {
		t.Fatalf("point_order = %v, want [104960]", got.PointOrder)
	}
}

// TestStoreSaveRejectsATalentIDAboveTheIntegerColumn guards the other side
// of the widened column: a talent id past math.MaxInt32 still cannot be
// stored, and Save must say so clearly rather than truncate it silently.
func TestStoreSaveRejectsATalentIDAboveTheIntegerColumn(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()

	b, err := New(Input{
		ClassID: 1, RaceID: 1, TreeVersion: "test-1",
		PointOrder: []int{math.MaxInt32 + 1},
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := s.Save(ctx, b, nil); err == nil {
		t.Fatal("want an error for a talent id that does not fit an integer column")
	}
}

func TestStoreAddViewsIncrementsEachID(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	one, _, err := s.Save(ctx, sampleBuild(t, "one"), nil)
	if err != nil {
		t.Fatal(err)
	}
	other, err := New(Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", PointOrder: []int{201}})
	if err != nil {
		t.Fatal(err)
	}
	two, _, err := s.Save(ctx, other, nil)
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

// The rest of this file tests the smallint/integer conversion helpers
// directly, so the bounds check runs in every environment even without
// TEST_DATABASE_URL set.

func TestIntegerAcceptsTheFullInt32RangeIncludingASixDigitTalentID(t *testing.T) {
	for _, v := range []int{math.MinInt32, 0, 104960, math.MaxInt32} {
		got, err := integer(v, "talent id")
		if err != nil {
			t.Fatalf("integer(%d): %v", v, err)
		}
		if int(got) != v {
			t.Fatalf("integer(%d) = %d, want %d", v, got, v)
		}
	}
}

func TestIntegerRejectsValuesOutsideInt32(t *testing.T) {
	for _, v := range []int{math.MaxInt32 + 1, math.MinInt32 - 1} {
		if _, err := integer(v, "talent id"); err == nil {
			t.Fatalf("integer(%d): want an error, got nil", v)
		}
	}
}

func TestIntegersConvertsEveryElement(t *testing.T) {
	got, err := integers([]int{101, 104960, 201})
	if err != nil {
		t.Fatal(err)
	}
	want := []int32{101, 104960, 201}
	if !slices.Equal(got, want) {
		t.Fatalf("integers = %v, want %v", got, want)
	}
}

func TestIntegersRejectsAnOutOfRangeElement(t *testing.T) {
	if _, err := integers([]int{101, math.MaxInt32 + 1}); err == nil {
		t.Fatal("integers: want an error for an out-of-range talent id")
	}
}

func TestSmallintStillRejectsAndAcceptsInt16Bounds(t *testing.T) {
	if _, err := smallint(math.MaxInt16+1, "class_id"); err == nil {
		t.Fatal("smallint: want an error above math.MaxInt16")
	}
	got, err := smallint(1, "class_id")
	if err != nil || got != 1 {
		t.Fatalf("smallint(1) = %d, err=%v", got, err)
	}
}

func TestASavedBuildRemembersWhoSavedIt(t *testing.T) {
	s, owner := storeWithUser(t)
	b := aBuild("Fury")
	if _, _, err := s.Save(t.Context(), b, &owner); err != nil {
		t.Fatal(err)
	}
	page, err := s.Mine(t.Context(), owner, 1)
	if err != nil {
		t.Fatal(err)
	}
	if page.Total != 1 || len(page.Rows) != 1 || page.Rows[0].ID != b.ID {
		t.Fatalf("mine: %+v", page)
	}
	if page.Rows[0].Title != "Fury" || page.Rows[0].TreeVersion != b.TreeVersion {
		t.Errorf("the row is not the whole build: %+v", page.Rows[0])
	}
}

func TestAnAnonymousBuildBelongsToNobodyUntilSomeoneSavesIt(t *testing.T) {
	s, owner := storeWithUser(t)
	b := aBuild("Fury")
	// Saved by nobody first.
	if _, created, err := s.Save(t.Context(), b, nil); err != nil || !created {
		t.Fatalf("created %v err %v", created, err)
	}
	page, err := s.Mine(t.Context(), owner, 1)
	if err != nil {
		t.Fatal(err)
	}
	if page.Total != 0 {
		t.Fatalf("an unowned build showed in a history: %+v", page)
	}
	// The same build saved again by a signed-in player claims the row:
	// ids are content hashes, so this is the same row, and leaving it
	// ownerless would mean a player could never list a build they saved.
	if _, created, err := s.Save(t.Context(), b, &owner); err != nil || created {
		t.Fatalf("created %v err %v; a content-hash id is saved once", created, err)
	}
	page, err = s.Mine(t.Context(), owner, 1)
	if err != nil {
		t.Fatal(err)
	}
	if page.Total != 1 {
		t.Fatalf("the second saver did not claim the unowned row: %+v", page)
	}
}

func TestAnOwnedBuildIsNotReassignedByASecondSaver(t *testing.T) {
	s, first := storeWithUser(t)
	second := anotherUser(t, s)
	b := aBuild("Fury")
	if _, _, err := s.Save(t.Context(), b, &first); err != nil {
		t.Fatal(err)
	}
	if _, _, err := s.Save(t.Context(), b, &second); err != nil {
		t.Fatal(err)
	}
	mine, err := s.Mine(t.Context(), first, 1)
	if err != nil {
		t.Fatal(err)
	}
	theirs, err := s.Mine(t.Context(), second, 1)
	if err != nil {
		t.Fatal(err)
	}
	if mine.Total != 1 || theirs.Total != 0 {
		t.Fatalf("the first saver keeps it: mine %d theirs %d", mine.Total, theirs.Total)
	}
}
