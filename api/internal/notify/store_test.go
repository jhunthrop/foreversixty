package notify

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
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
	if err := st.Save(context.Background(), SaveInput{UserID: uid, Kind: KindDiscord, ChannelID: "555", TokenBox: []byte("box-1"), Events: bothEvents, Source: SourcePasted}); err != nil {
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
	if err := st.Save(ctx, SaveInput{UserID: uid, Kind: KindDiscord, ChannelID: "777", TokenBox: []byte("box-2"), Events: events, Source: SourcePasted}); err != nil {
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
	if err := st.Save(ctx, SaveInput{UserID: uid, Kind: KindDiscord, ChannelID: "1", TokenBox: []byte("b"), Events: bothEvents, Source: SourcePasted}); err != nil {
		t.Fatal(err)
	}
	if _, err := st.Pool.Exec(ctx, `delete from users where id = $1`, uid); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := st.Get(ctx, uid, KindDiscord); ok {
		t.Fatal("row survived the user")
	}
}

func TestGetNeverReturnsTheToken(t *testing.T) {
	st, uid := saved(t)
	wh, ok, err := st.Get(context.Background(), uid, KindDiscord)
	if err != nil || !ok {
		t.Fatalf("Get = %v, %v", ok, err)
	}
	if strings.Contains(fmt.Sprintf("%+v", wh), "box-1") {
		t.Fatal("Get returned the sealed box")
	}
	typ := reflect.TypeOf(wh)
	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		if strings.Contains(name, "Token") || strings.Contains(name, "Box") {
			t.Fatalf("Webhook has a credential field %q", name)
		}
	}
}

func TestExistingWebhooksKeepPostingPrivateLogs(t *testing.T) {
	ctx := context.Background()
	pool := testPool(t)
	st := &Store{Pool: pool}
	old, fresh := seedUser(t, pool), seedUser(t, pool)
	if _, err := pool.Exec(ctx, `insert into notification_webhooks (user_id, kind, channel_id, token_box, events)
		values ($1, $2, '9', 'b', '{"sim_finished": true, "upload_parsed": true}')`, old, KindDiscord); err != nil {
		t.Fatal(err)
	}
	migration, err := os.ReadFile("../db/migrations/0047_notification_webhook_name_private.up.sql")
	if err != nil {
		t.Fatal(err)
	}
	backfill := func() {
		for _, stmt := range strings.Split(string(migration), ";") {
			if idx := strings.Index(stmt, "update notification_webhooks"); idx >= 0 {
				if _, err := pool.Exec(ctx, stmt[idx:]); err != nil {
					t.Fatal(err)
				}
			}
		}
	}
	backfill()
	var raw string
	if err := pool.QueryRow(ctx, `select events->>'upload_private' from notification_webhooks where user_id = $1`, old).Scan(&raw); err != nil {
		t.Fatal(err)
	}
	if raw != "true" {
		t.Fatalf("stored upload_private after the backfill = %q, want \"true\"", raw)
	}
	events := map[EventName]bool{EventSimFinished: true, EventUploadParsed: true}
	if err := st.Save(ctx, SaveInput{UserID: fresh, Kind: KindDiscord, ChannelID: "8", TokenBox: []byte("b"), Events: events, Source: SourcePasted}); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `update notification_webhooks set events = events || '{"upload_private": false}' where user_id = $1`, old); err != nil {
		t.Fatal(err)
	}
	backfill() // a second run never flips an opt-out back
	for uid, want := range map[int64]bool{old: false, fresh: false} {
		w, _, err := st.Get(ctx, uid, KindDiscord)
		if err != nil || w.Events[EventUploadPrivate] != want {
			t.Fatalf("user %d: UploadPrivate = %v (%v), want %v", uid, w.Events[EventUploadPrivate], err, want)
		}
	}
}

func TestSaveStoresACleanName(t *testing.T) {
	for name, c := range map[string]struct {
		raw  string
		want *string
	}{
		"trimmed":        {"  Raid pings  ", ptr("Raid pings")},
		"controls gone":  {"Ra\x00id\x1f \x7fpings\n", ptr("Raid pings")},
		"81 runes to 80": {strings.Repeat("é", 81), ptr(strings.Repeat("é", 80))},
		"empty is nil":   {"   ", nil},
		"only controls":  {"\x00\x01", nil},
		"c1 controls":    {"a\u0085b\u009fc", ptr("abc")},
		"separators":     {"a\u2028b\u2029c", ptr("abc")},
		"bidi":           {"a\u202ab\u202ec\u2066d\u2069e", ptr("abcde")},
		"unicode kept":   {"Gilde ⚔ Nacht", ptr("Gilde ⚔ Nacht")},
	} {
		if got := cleanWebhookName(c.raw); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: %v, want %v", name, got, c.want)
		}
	}
	st, uid := saved(t)
	ctx := context.Background()
	if err := st.Save(ctx, SaveInput{UserID: uid, Kind: KindDiscord, ChannelID: "555", Name: ptr("Raid pings"), TokenBox: []byte("box"), Events: bothEvents, Source: SourcePasted}); err != nil {
		t.Fatal(err)
	}
	if w, _, err := st.Get(ctx, uid, KindDiscord); err != nil || w.Name == nil || *w.Name != "Raid pings" {
		t.Fatalf("stored name = %v (%v)", w.Name, err)
	}
	if err := st.Save(ctx, SaveInput{UserID: uid, Kind: KindDiscord, ChannelID: "555", TokenBox: []byte("box"), Events: bothEvents, Source: SourcePasted}); err != nil {
		t.Fatal(err)
	}
	if w, _, _ := st.Get(ctx, uid, KindDiscord); w.Name != nil {
		t.Fatalf("a nil name was kept: %v", *w.Name)
	}
}

func ptr(s string) *string { return &s }

func TestAMissingUploadPrivateKeyReadsTrueAndSaveWritesAllThree(t *testing.T) {
	ctx := context.Background()
	pool := testPool(t)
	st := &Store{Pool: pool}
	legacy, optedOut, fresh := seedUser(t, pool), seedUser(t, pool), seedUser(t, pool)
	for uid, events := range map[int64]string{
		legacy: `{"sim_finished": true, "upload_parsed": true}`, optedOut: `{"upload_private": false}`,
	} {
		if _, err := pool.Exec(ctx, `insert into notification_webhooks (user_id, kind, channel_id, token_box, events)
			values ($1, $2, '9', 'b', $3)`, uid, KindDiscord, events); err != nil {
			t.Fatal(err)
		}
	}
	for uid, want := range map[int64]bool{legacy: true, optedOut: false} {
		if w, _, err := st.Get(ctx, uid, KindDiscord); err != nil || w.Events[EventUploadPrivate] != want {
			t.Fatalf("user %d: %v (%v), want %v", uid, w.Events[EventUploadPrivate], err, want)
		}
	}
	if err := st.Save(ctx, SaveInput{UserID: fresh, Kind: KindDiscord, ChannelID: "8", TokenBox: []byte("b"), Events: map[EventName]bool{EventSimFinished: true}, Source: SourcePasted}); err != nil {
		t.Fatal(err)
	}
	var keys int
	if err := pool.QueryRow(ctx, `select count(*) from jsonb_object_keys((select events from notification_webhooks where user_id = $1))`, fresh).Scan(&keys); err != nil || keys != 3 {
		t.Fatalf("stored keys = %d (%v), want 3", keys, err)
	}
	if w, _, _ := st.Get(ctx, fresh, KindDiscord); w.Events[EventUploadPrivate] {
		t.Fatal("a fresh save read back private logs on")
	}
}
