package notify

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func testPool(t *testing.T) *pgxpool.Pool {
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
	if _, err := pool.Exec(context.Background(), `truncate users cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

func seedUser(t *testing.T, pool *pgxpool.Pool) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into users (email) values ($1) returning id`,
		"n-"+time.Now().Format("150405.000000000")+"@example.com").Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

var bothEvents = map[EventName]bool{EventSimFinished: true, EventUploadParsed: true}

func saved(t *testing.T) (*Store, int64) {
	t.Helper()
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid := seedUser(t, pool)
	if err := st.Save(context.Background(), uid, KindDiscord, "555", []byte("box-1"), bothEvents); err != nil {
		t.Fatal(err)
	}
	return st, uid
}

func TestSaveUpsertsAndResetsFailures(t *testing.T) {
	ctx := context.Background()
	st, uid := saved(t)
	for i := 0; i < MaxFailures; i++ {
		if _, err := st.RecordFailure(ctx, uid, KindDiscord, "boom", true); err != nil {
			t.Fatal(err)
		}
	}
	events := map[EventName]bool{EventSimFinished: true}
	if err := st.Save(ctx, uid, KindDiscord, "777", []byte("box-2"), events); err != nil {
		t.Fatal(err)
	}
	w, ok, err := st.Get(ctx, uid, KindDiscord)
	if err != nil || !ok {
		t.Fatalf("Get = %v, %v", ok, err)
	}
	if w.ChannelID != "777" || w.Failures != 0 || w.LastError != "" || w.DisabledAt != nil {
		t.Fatalf("not reset: %+v", w)
	}
	if !w.Events[EventSimFinished] || w.Events[EventUploadParsed] {
		t.Fatalf("events = %v", w.Events)
	}
	box, err := st.TokenBox(ctx, uid, KindDiscord)
	if err != nil || string(box) != "box-2" {
		t.Fatalf("TokenBox = %q, %v", box, err)
	}
}

func TestGetReportsAbsence(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid := seedUser(t, pool)
	if _, ok, err := st.Get(context.Background(), uid, KindDiscord); ok || err != nil {
		t.Fatalf("Get = %v, %v; want absent", ok, err)
	}
	if _, err := st.TokenBox(context.Background(), uid, KindDiscord); err != ErrNotFound {
		t.Fatalf("TokenBox err = %v; want ErrNotFound", err)
	}
}

func TestSaveEventsKeepsTheTokenAndCounters(t *testing.T) {
	ctx := context.Background()
	st, uid := saved(t)
	if _, err := st.RecordFailure(ctx, uid, KindDiscord, "boom", true); err != nil {
		t.Fatal(err)
	}
	if err := st.SaveEvents(ctx, uid, KindDiscord, map[EventName]bool{EventUploadParsed: true}); err != nil {
		t.Fatal(err)
	}
	w, _, _ := st.Get(ctx, uid, KindDiscord)
	box, _ := st.TokenBox(ctx, uid, KindDiscord)
	if w.Failures != 1 || w.Events[EventSimFinished] || string(box) != "box-1" {
		t.Fatalf("w = %+v box = %q", w, box)
	}
	if err := st.SaveEvents(ctx, uid+999, KindDiscord, bothEvents); err != ErrNotFound {
		t.Fatalf("SaveEvents on a missing row err = %v", err)
	}
}

func TestRecordFailureDisablesAtTheFifthAndSuccessResets(t *testing.T) {
	ctx := context.Background()
	st, uid := saved(t)
	for i := 1; i <= MaxFailures; i++ {
		disabled, err := st.RecordFailure(ctx, uid, KindDiscord, "boom", true)
		if err != nil {
			t.Fatal(err)
		}
		if disabled != (i == MaxFailures) {
			t.Fatalf("failure %d disabled = %v", i, disabled)
		}
	}
	w, _, _ := st.Get(ctx, uid, KindDiscord)
	if w.Failures != MaxFailures || w.DisabledAt == nil || w.LastError != "boom" {
		t.Fatalf("after five: %+v", w)
	}
	if err := st.RecordSuccess(ctx, uid, KindDiscord); err != nil {
		t.Fatal(err)
	}
	w, _, _ = st.Get(ctx, uid, KindDiscord)
	if w.Failures != 0 || w.LastError != "" || w.DisabledAt != nil {
		t.Fatalf("after success: %+v", w)
	}
}

func TestUncountedFailureOnlyUpdatesLastError(t *testing.T) {
	ctx := context.Background()
	st, uid := saved(t)
	for i := 0; i < MaxFailures+2; i++ {
		disabled, err := st.RecordFailure(ctx, uid, KindDiscord, "rate limited", false)
		if err != nil || disabled {
			t.Fatalf("uncounted failure disabled=%v err=%v", disabled, err)
		}
	}
	w, _, _ := st.Get(ctx, uid, KindDiscord)
	if w.Failures != 0 || w.DisabledAt != nil || w.LastError != "rate limited" {
		t.Fatalf("w = %+v", w)
	}
}

func TestDeleteCascadesWithTheUser(t *testing.T) {
	ctx := context.Background()
	st, uid := saved(t)
	if err := st.Delete(ctx, uid, KindDiscord); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := st.Get(ctx, uid, KindDiscord); ok {
		t.Fatal("row survived Delete")
	}
	if err := st.Save(ctx, uid, KindDiscord, "1", []byte("b"), bothEvents); err != nil {
		t.Fatal(err)
	}
	if _, err := st.Pool.Exec(ctx, `delete from users where id = $1`, uid); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := st.Get(ctx, uid, KindDiscord); ok {
		t.Fatal("row survived the user")
	}
}
