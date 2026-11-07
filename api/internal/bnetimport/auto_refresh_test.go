// api/internal/bnetimport/auto_refresh_test.go
package bnetimport

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// prefixSealer is a stand-in for notify.Sealer: reversible, and visibly not the plain token.
type prefixSealer struct{}

func (prefixSealer) Seal(plain string) ([]byte, error) { return []byte("sealed:" + plain), nil }
func (prefixSealer) Open(box []byte) (string, error) {
	s, ok := strings.CutPrefix(string(box), "sealed:")
	if !ok {
		return "", errors.New("not sealed here")
	}
	return s, nil
}

func seedBnetUser(t *testing.T, pool *pgxpool.Pool, syncedAgo string) int64 {
	t.Helper()
	uid := seedUser(t, pool)
	if _, err := pool.Exec(context.Background(),
		`update users set bnet_sub = 'sub-' || id::text,
		   bnet_synced_at = case when $2 = '' then null else now() - $2::interval end where id = $1`,
		uid, syncedAgo); err != nil {
		t.Fatal(err)
	}
	return uid
}

func seedBnetCharacter(t *testing.T, pool *pgxpool.Pool, uid int64, key, name string, bnetID int, refreshedAgo string) {
	t.Helper()
	if _, err := pool.Exec(context.Background(),
		`insert into characters (key, region, ruleset, name, user_id, realm_slug, bnet_character_id, source, level, refreshed_at)
		 values ($1, 'us', 'pvp', $2, $3, 'whitemane', $4, 'bnet', 30, now() - $5::interval)`,
		key, name, uid, bnetID, refreshedAgo); err != nil {
		t.Fatal(err)
	}
}

// publicProfile stubs the public reads one character's refresh makes.
func (f *blizzardFixture) publicProfile(slug, name string) {
	base := "/profile/wow/character/whitemane/" + slug
	f.json(http.MethodGet, base+"?namespace=profile-classic1x-us", http.StatusOK,
		map[string]any{"name": name, "level": 30, "faction": map[string]string{"type": "HORDE"},
			"character_class": map[string]string{"name": "Rogue"}, "realm": map[string]string{"slug": "whitemane"}})
	f.json(http.MethodGet, base+"/equipment?namespace=profile-classic1x-us", http.StatusOK,
		map[string]any{"equipped_items": []map[string]any{}})
	f.json(http.MethodGet, base+"/character-media?namespace=profile-classic1x-us", http.StatusOK,
		map[string]any{"assets": []map[string]string{}})
}

func TestVisitDueAfterAnHour(t *testing.T) {
	now := time.Date(2026, 11, 7, 12, 0, 0, 0, time.UTC)
	at := func(ago time.Duration) *time.Time { t := now.Add(-ago); return &t }
	cases := []struct {
		name string
		last *time.Time
		want bool
	}{
		{"never synced", nil, true},
		{"synced 30 minutes ago", at(30 * time.Minute), false},
		{"synced 59 minutes ago", at(59 * time.Minute), false},
		{"synced an hour ago", at(time.Hour), true},
		{"synced yesterday", at(26 * time.Hour), true},
	}
	for _, c := range cases {
		if got := visitDue(c.last, now); got != c.want {
			t.Errorf("%s: visitDue = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestClaimSyncIsTheCooldown(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	svc := &Service{Pool: pool}

	never := seedBnetUser(t, pool, "")
	if ok, err := svc.claimSync(ctx, never); err != nil || !ok {
		t.Fatalf("first claim = %v, %v; want claimed", ok, err)
	}
	if ok, err := svc.claimSync(ctx, never); err != nil || ok {
		t.Fatalf("second claim inside the hour = %v, %v; want refused (the cooldown)", ok, err)
	}

	recent := seedBnetUser(t, pool, "20 minutes")
	if ok, _ := svc.claimSync(ctx, recent); ok {
		t.Fatal("an account synced 20 minutes ago was claimed again")
	}
	stale := seedBnetUser(t, pool, "2 hours")
	if ok, _ := svc.claimSync(ctx, stale); !ok {
		t.Fatal("an account synced 2 hours ago was not claimed")
	}

	emailOnly := seedUser(t, pool)
	if ok, _ := svc.claimSync(ctx, emailOnly); ok {
		t.Fatal("an account with no Battle.net sign-in was claimed")
	}
}

func TestVisitedIgnoresFreshAndNonBattlenetAccounts(t *testing.T) {
	svc := &Service{}
	recent := time.Now().Add(-10 * time.Minute)
	svc.Visited(auth.User{ID: 1, BnetSub: "sub", BnetSyncedAt: &recent})
	svc.Visited(auth.User{ID: 2})
	count := 0
	svc.visiting.Range(func(any, any) bool { count++; return true })
	if count != 0 {
		t.Fatalf("visiting = %d accounts, want none started", count)
	}
}

func TestVisitedRefreshesAStaleAccountInTheBackground(t *testing.T) {
	pool := testPool(t)
	uid := seedBnetUser(t, pool, "3 hours")
	seedBnetCharacter(t, pool, uid, "us/pvp/visitor", "Visitor", 701, "5 hours")
	f := newBlizzardFixture(t)
	f.realms("us", map[string]string{"whitemane": "PVP"})
	f.publicProfile("visitor", "Visitor")
	svc := newTestService(t, pool, f)

	var synced *time.Time
	if err := pool.QueryRow(context.Background(), `select bnet_synced_at from users where id = $1`, uid).Scan(&synced); err != nil {
		t.Fatal(err)
	}
	start := time.Now()
	svc.Visited(auth.User{ID: uid, BnetSub: "sub", BnetSyncedAt: synced})
	if time.Since(start) > 50*time.Millisecond {
		t.Fatalf("Visited took %s; it must not wait on the refresh", time.Since(start))
	}
	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		var fresh bool
		if err := pool.QueryRow(context.Background(),
			`select refreshed_at > now() - interval '1 minute' from characters where key = 'us/pvp/visitor'`).Scan(&fresh); err != nil {
			t.Fatal(err)
		}
		if fresh {
			waitVisitsDone(t, svc)
			return
		}
		time.Sleep(50 * time.Millisecond)
	}
	t.Fatal("the visit refresh never refreshed the character")
}

// waitVisitsDone waits for every visit refresh svc started, so none outlives the test's pool.
func waitVisitsDone(t *testing.T, svc *Service) {
	t.Helper()
	for range 200 {
		running := false
		svc.visiting.Range(func(any, any) bool { running = true; return false })
		if !running {
			return
		}
		time.Sleep(25 * time.Millisecond)
	}
	t.Fatal("a visit refresh is still running")
}

func TestRefreshWindowForActiveAndIdleAccounts(t *testing.T) {
	now := time.Date(2026, 11, 7, 12, 0, 0, 0, time.UTC)
	at := func(ago time.Duration) *time.Time { t := now.Add(-ago); return &t }
	if got := refreshWindowFor(at(24*time.Hour), now); got != 3*time.Hour {
		t.Errorf("active yesterday: window = %s, want 3h", got)
	}
	if got := refreshWindowFor(at(13*24*time.Hour), now); got != 3*time.Hour {
		t.Errorf("active 13 days ago: window = %s, want 3h", got)
	}
	if got := refreshWindowFor(at(15*24*time.Hour), now); got != 20*time.Hour {
		t.Errorf("last active 15 days ago: window = %s, want 20h", got)
	}
	if got := refreshWindowFor(nil, now); got != 20*time.Hour {
		t.Errorf("never synced: window = %s, want 20h", got)
	}
}

func TestStaleBnetCharactersUsesTheActiveAndIdleWindows(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	active := seedBnetUser(t, pool, "1 day")
	idle := seedBnetUser(t, pool, "20 days")
	never := seedBnetUser(t, pool, "")
	seedBnetCharacter(t, pool, active, "us/pvp/active-due", "ActiveDue", 801, "4 hours")
	seedBnetCharacter(t, pool, active, "us/pvp/active-fresh", "ActiveFresh", 802, "2 hours")
	seedBnetCharacter(t, pool, idle, "us/pvp/idle-fresh", "IdleFresh", 803, "4 hours")
	seedBnetCharacter(t, pool, idle, "us/pvp/idle-due", "IdleDue", 804, "21 hours")
	seedBnetCharacter(t, pool, never, "us/pvp/never-fresh", "NeverFresh", 805, "19 hours")
	seedBnetCharacter(t, pool, never, "us/pvp/never-due", "NeverDue", 806, "21 hours")

	due, err := (&Service{Pool: pool}).staleBnetCharacters(ctx)
	if err != nil {
		t.Fatal(err)
	}
	got := map[int64]bool{}
	for _, sc := range due {
		got[sc.BnetCharacterID] = true
	}
	want := map[int64]bool{801: true, 804: true, 806: true}
	if len(got) != len(want) {
		t.Fatalf("due = %v, want %v", got, want)
	}
	for id := range want {
		if !got[id] {
			t.Fatalf("due = %v, want %v", got, want)
		}
	}
}

func readToken(t *testing.T, pool *pgxpool.Pool, uid int64) []byte {
	t.Helper()
	var box []byte
	if err := pool.QueryRow(context.Background(), `select bnet_token_box from users where id = $1`, uid).Scan(&box); err != nil {
		t.Fatal(err)
	}
	return box
}

func TestKeepTokenSealsAndNeedsASealerAndAnExpiry(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	uid := seedBnetUser(t, pool, "")
	in24h := time.Now().Add(24 * time.Hour)

	(&Service{Pool: pool}).KeepToken(ctx, uid, "tok", in24h)
	if box := readToken(t, pool, uid); box != nil {
		t.Fatalf("kept %q with no sealer", box)
	}
	svc := &Service{Pool: pool, Sealer: prefixSealer{}}
	svc.KeepToken(ctx, uid, "tok", time.Time{})
	if box := readToken(t, pool, uid); box != nil {
		t.Fatalf("kept %q with no expiry", box)
	}
	svc.KeepToken(ctx, uid, "tok", in24h)
	if box := readToken(t, pool, uid); string(box) != "sealed:tok" {
		t.Fatalf("box = %q, want the sealed token", box)
	}
	if got, err := svc.liveToken(ctx, uid); err != nil || got != "tok" {
		t.Fatalf("liveToken = %q, %v; want tok", got, err)
	}
}

func TestRefreshAccountUsesALiveTokenForTheAccountList(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	uid := seedBnetUser(t, pool, "2 hours")
	f := newBlizzardFixture(t)
	var bearer atomic.Value
	f.handlers[http.MethodGet+" "+accountCharactersPath] = func(w http.ResponseWriter, r *http.Request) {
		bearer.Store(r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"wow_accounts":[{"id":1,"characters":[
			{"name":"Lowbie","id":901,"realm":{"slug":"whitemane","name":"Whitemane"},
			 "playable_class":{"name":"Rogue"},"faction":{"type":"HORDE"},"level":7},
			{"name":"Newmain","id":902,"realm":{"slug":"whitemane","name":"Whitemane"},
			 "playable_class":{"name":"Rogue"},"faction":{"type":"HORDE"},"level":31}
		]}]}`))
	}
	f.realms("us", map[string]string{"whitemane": "PVP"})
	f.publicProfile("lowbie", "Lowbie")
	f.publicProfile("newmain", "Newmain")
	svc := newTestService(t, pool, f)
	svc.Sealer = prefixSealer{}
	svc.KeepToken(ctx, uid, "live-tok", time.Now().Add(20*time.Hour))

	if err := svc.RefreshAccount(ctx, uid); err != nil {
		t.Fatal(err)
	}
	if got, _ := bearer.Load().(string); got != "Bearer live-tok" {
		t.Fatalf("account list read with %q, want the stored token", got)
	}
	var n int
	var main *string
	if err := pool.QueryRow(ctx,
		`select (select count(*) from characters where user_id = $1), main_character_key from users where id = $1`, uid).
		Scan(&n, &main); err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("characters = %d, want both, a level-7 one included", n)
	}
	if main == nil || *main != "us/pvp/newmain" {
		t.Fatalf("main_character_key = %v, want the highest level us/pvp/newmain", main)
	}
}

func TestRefreshAccountFallsBackToPublicProfilesWhenBattlenetRefusesTheToken(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	uid := seedBnetUser(t, pool, "2 hours")
	seedBnetCharacter(t, pool, uid, "us/pvp/kept", "Kept", 911, "5 hours")
	f := newBlizzardFixture(t)
	f.json(http.MethodGet, accountCharactersPath, http.StatusUnauthorized, nil)
	f.realms("us", map[string]string{"whitemane": "PVP"})
	f.publicProfile("kept", "Kept")
	svc := newTestService(t, pool, f)
	svc.Sealer = prefixSealer{}
	svc.KeepToken(ctx, uid, "revoked-tok", time.Now().Add(20*time.Hour))

	if err := svc.RefreshAccount(ctx, uid); err != nil {
		t.Fatal(err)
	}
	assertPublicRefresh(t, pool, uid, "us/pvp/kept")
}

func TestRefreshAccountUsesPublicProfilesOnceTheTokenExpired(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	uid := seedBnetUser(t, pool, "2 hours")
	seedBnetCharacter(t, pool, uid, "us/pvp/aged", "Aged", 921, "5 hours")
	if _, err := pool.Exec(ctx,
		`update users set bnet_token_box = 'sealed:old-tok', bnet_token_expires_at = now() + interval '2 minutes' where id = $1`,
		uid); err != nil {
		t.Fatal(err)
	}
	// No account-list stub: the fixture fails the test if the expired token is used.
	f := newBlizzardFixture(t)
	f.realms("us", map[string]string{"whitemane": "PVP"})
	f.publicProfile("aged", "Aged")
	svc := newTestService(t, pool, f)
	svc.Sealer = prefixSealer{}

	if err := svc.RefreshAccount(ctx, uid); err != nil {
		t.Fatal(err)
	}
	assertPublicRefresh(t, pool, uid, "us/pvp/aged")
}

// assertPublicRefresh checks a public-profile refresh ran: the character refreshed, the
// account's Battle.net stamp moved, and the stored token is gone.
func assertPublicRefresh(t *testing.T, pool *pgxpool.Pool, uid int64, key string) {
	t.Helper()
	var fresh, stamped bool
	if err := pool.QueryRow(context.Background(),
		`select c.refreshed_at > now() - interval '1 minute', u.bnet_imported_at > now() - interval '1 minute'
		 from characters c join users u on u.id = c.user_id where c.key = $1`, key).Scan(&fresh, &stamped); err != nil {
		t.Fatal(err)
	}
	if !fresh || !stamped {
		t.Fatalf("refreshed=%v account stamped=%v, want both", fresh, stamped)
	}
	if box := readToken(t, pool, uid); box != nil {
		t.Fatalf("token box = %q, want it forgotten", box)
	}
}

func TestRateBackoffDoublesCapsAndResets(t *testing.T) {
	now := time.Date(2026, 11, 7, 12, 0, 0, 0, time.UTC)
	var b rateBackoff
	if b.blocked(now) {
		t.Fatal("a fresh backoff is blocked")
	}
	if got := b.hit(now); got != 30*time.Second {
		t.Fatalf("first pause = %s, want 30s", got)
	}
	if !b.blocked(now.Add(29*time.Second)) || b.blocked(now.Add(30*time.Second)) {
		t.Fatal("the first pause does not cover exactly 30s")
	}
	if got := b.hit(now); got != time.Minute {
		t.Fatalf("second pause = %s, want 1m", got)
	}
	for range 10 {
		b.hit(now)
	}
	if got := b.hit(now); got != 15*time.Minute {
		t.Fatalf("pause = %s, want the 15m cap", got)
	}
	b.reset()
	if got := b.hit(now); got != 30*time.Second {
		t.Fatalf("pause after reset = %s, want 30s", got)
	}
}

func TestRefreshBatchSkipsInsideABackoff(t *testing.T) {
	svc := &Service{}
	svc.backoff.hit(time.Now())
	result, err := svc.refreshBatch(context.Background(), []staleCharacter{{BnetCharacterID: 1}})
	if err != nil || !result.RateLimited || result.Refreshed != 0 {
		t.Fatalf("result = %+v, %v; want RateLimited and nothing refreshed", result, err)
	}
}
