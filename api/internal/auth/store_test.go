package auth

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
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
	if _, err := pool.Exec(context.Background(),
		`truncate users, login_tokens, pairing_codes cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

func testStore(t *testing.T) *Store { return &Store{Pool: testPool(t)} }

func TestUpsertBnetUserIsIdempotentAndRefreshesTheBattletag(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()

	first, err := s.UpsertBnetUser(ctx, "12345", "Baelgrim#1234")
	if err != nil {
		t.Fatal(err)
	}
	if first.ID == 0 || first.Role != "user" {
		t.Fatalf("user = %+v, want an id and the default role", first)
	}
	again, err := s.UpsertBnetUser(ctx, "12345", "Renamed#4321")
	if err != nil {
		t.Fatal(err)
	}
	if again.ID != first.ID {
		t.Fatalf("second sign-in made user %d, want %d", again.ID, first.ID)
	}
	if again.Name() != "Renamed#4321" {
		t.Fatalf("battletag = %v, want the new one", again.Battletag)
	}
}

func TestSessionsResolveUntilTheyExpireOrAreDeleted(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}

	sess, err := s.CreateSession(ctx, u.ID, "email", SessionTTL)
	if err != nil {
		t.Fatal(err)
	}
	got, err := s.SessionUser(ctx, sess.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != u.ID {
		t.Fatalf("session resolved to user %d, want %d", got.ID, u.ID)
	}

	expired, err := s.CreateSession(ctx, u.ID, "email", -time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.SessionUser(ctx, expired.ID); err != ErrNotFound {
		t.Fatalf("err = %v, want ErrNotFound for an expired session", err)
	}

	if err := s.DeleteSession(ctx, sess.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.SessionUser(ctx, sess.ID); err != ErrNotFound {
		t.Fatalf("err = %v, want ErrNotFound after sign-out", err)
	}
}

func TestLoginTokensAreSingleUseAndCounted(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	token := NewSessionID()
	if err := s.CreateLoginToken(ctx, "raider@example.com", TokenHash(token), LoginTokenTTL); err != nil {
		t.Fatal(err)
	}
	n, err := s.RecentLoginTokens(ctx, "raider@example.com", time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("recent = %d, want 1", n)
	}
	email, err := s.ConsumeLoginToken(ctx, TokenHash(token))
	if err != nil {
		t.Fatal(err)
	}
	if email != "raider@example.com" {
		t.Fatalf("email = %q", email)
	}
	if _, err := s.ConsumeLoginToken(ctx, TokenHash(token)); err != ErrNotFound {
		t.Fatalf("err = %v, want ErrNotFound on the second click", err)
	}
}

func TestExpiredLoginTokenCannotBeSpent(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	token := NewSessionID()
	if err := s.CreateLoginToken(ctx, "raider@example.com", TokenHash(token), -time.Minute); err != nil {
		t.Fatal(err)
	}
	if _, err := s.ConsumeLoginToken(ctx, TokenHash(token)); err != ErrNotFound {
		t.Fatalf("err = %v, want ErrNotFound for an expired link", err)
	}
}

func TestPairingCodeIsSpentOnce(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	code, err := s.CreatePairingCode(ctx, u.ID, PairingCodeTTL)
	if err != nil {
		t.Fatal(err)
	}
	if len(code) != PairingCodeChars {
		t.Fatalf("code = %q, want %d characters", code, PairingCodeChars)
	}
	got, err := s.ClaimPairingCode(ctx, code)
	if err != nil || got != u.ID {
		t.Fatalf("claim = %d, %v", got, err)
	}
	if _, err := s.ClaimPairingCode(ctx, code); err != ErrNotFound {
		t.Fatalf("err = %v, want ErrNotFound on a second claim", err)
	}
}

func TestDevicesResolveByTokenUntilRevoked(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	token, hash := NewDeviceToken()
	d, err := s.CreateDevice(ctx, u.ID, "Raid PC", "windows/amd64", hash)
	if err != nil {
		t.Fatal(err)
	}
	if len(d.ID) != DeviceIDChars {
		t.Fatalf("device id = %q, want %d characters", d.ID, DeviceIDChars)
	}
	got, err := s.DeviceByToken(ctx, TokenHash(token))
	if err != nil || got.ID != d.ID {
		t.Fatalf("device = %+v, err = %v", got, err)
	}
	if got.LastSeenAt == nil {
		t.Fatal("resolving a device should stamp last_seen_at")
	}
	list, err := s.Devices(ctx, u.ID)
	if err != nil || len(list) != 1 {
		t.Fatalf("devices = %v, err = %v", list, err)
	}
	ok, err := s.RevokeDevice(ctx, u.ID, d.ID)
	if err != nil || !ok {
		t.Fatalf("revoke = %v, %v", ok, err)
	}
	if _, err := s.DeviceByToken(ctx, TokenHash(token)); err != ErrNotFound {
		t.Fatalf("err = %v, want ErrNotFound for a revoked device", err)
	}
	if ok, _ := s.RevokeDevice(ctx, u.ID, d.ID); ok {
		t.Fatal("revoking twice should report nothing changed")
	}
}

func TestCharactersAndGuildsComeBackForMe(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if err := s.LinkCharacter(ctx, u.ID, Character{
		Key: "us/hardcore/baelgrim", Region: "us", Ruleset: "hardcore", Name: "Baelgrim", Class: "Warrior",
	}); err != nil {
		t.Fatal(err)
	}
	var guildID int64
	if err := s.Pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Forever') returning id`).
		Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Pool.Exec(ctx,
		`insert into guild_members (guild_id, user_id, rank) values ($1, $2, 'officer')`, guildID, u.ID); err != nil {
		t.Fatal(err)
	}

	chars, err := s.Characters(ctx, u.ID)
	if err != nil || len(chars) != 1 || chars[0].Class != "Warrior" {
		t.Fatalf("characters = %v, err = %v", chars, err)
	}
	guilds, err := s.Guilds(ctx, u.ID)
	if err != nil || len(guilds) != 1 || guilds[0].Rank != "officer" {
		t.Fatalf("guilds = %v, err = %v", guilds, err)
	}
	rank, ok, err := s.GuildRank(ctx, guildID, u.ID)
	if err != nil || !ok || rank != "officer" {
		t.Fatalf("rank = %q, %v, %v", rank, ok, err)
	}
	if _, ok, _ := s.GuildRank(ctx, guildID, u.ID+999); ok {
		t.Fatal("a stranger should not have a rank")
	}
}

func TestIDShapesMatchTheContract(t *testing.T) {
	if got := NewReportID(); len(got) != 12 {
		t.Fatalf("report id = %q, want 12 characters", got)
	}
	token, hash := NewDeviceToken()
	if len(token) != len(DeviceTokenPrefix)+32 {
		t.Fatalf("device token = %q, want the fsd_ prefix and 32 characters", token)
	}
	if len(hash) != 32 {
		t.Fatalf("hash is %d bytes, want a SHA-256", len(hash))
	}
	if !SameToken(token, token) || SameToken(token, token+"x") {
		t.Fatal("SameToken must compare the whole token")
	}
	seen := map[string]bool{}
	for range 100 {
		id := NewReportID()
		if seen[id] {
			t.Fatal("report ids must not repeat")
		}
		seen[id] = true
		for _, r := range id {
			if !(r >= 'a' && r <= 'z' || r >= '2' && r <= '7') {
				t.Fatalf("report id %q is not lowercase base32", id)
			}
		}
	}
}

func TestActorHelpers(t *testing.T) {
	if (Actor{}).Signed() {
		t.Error("the zero actor is anonymous")
	}
	a := Actor{UserID: 1, Role: "admin", DeviceID: "d"}
	if !a.Signed() || !a.IsDevice() || !a.IsModerator() {
		t.Errorf("actor = %+v", a)
	}
	ctx := WithActor(context.Background(), a)
	if ActorFrom(ctx).UserID != 1 {
		t.Error("the actor should survive the context")
	}
}

func TestAnExpiredSessionCookieIsClearedAndTreatedAsAnonymous(t *testing.T) {
	store := testStore(t)
	u, err := store.UpsertEmailUser(t.Context(), "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	sess, err := store.CreateSession(t.Context(), u.ID, "email", -time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	a := &Authenticator{Store: store}
	var reached Actor
	h := a.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reached = ActorFrom(r.Context())
		w.WriteHeader(http.StatusOK)
	}))
	r := httptest.NewRequest(http.MethodGet, "/v1/me", nil)
	r.AddCookie(&http.Cookie{Name: SessionCookie, Value: sess.ID})
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if reached.Signed() {
		t.Fatalf("actor = %+v, want anonymous", reached)
	}
	if !strings.Contains(w.Header().Get("Set-Cookie"), SessionCookie+"=;") {
		t.Fatalf("the stale cookie should be cleared: %q", w.Header().Values("Set-Cookie"))
	}
}

func TestRequireDeviceRejectsASession(t *testing.T) {
	h := RequireDevice(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	r := httptest.NewRequest(http.MethodPost, "/x", nil)
	r = r.WithContext(WithActor(r.Context(), Actor{UserID: 1, Method: "session"}))
	w := httptest.NewRecorder()
	h(w, r)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", w.Code)
	}
}

func TestRequireAcceptsEitherIdentity(t *testing.T) {
	h := Require(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	for _, a := range []Actor{{UserID: 1, Method: "session"}, {UserID: 1, Method: "device", DeviceID: "d"}} {
		r := httptest.NewRequest(http.MethodGet, "/x", nil)
		r = r.WithContext(WithActor(r.Context(), a))
		w := httptest.NewRecorder()
		h(w, r)
		if w.Code != http.StatusOK {
			t.Fatalf("actor %+v was refused with %d", a, w.Code)
		}
	}
	r := httptest.NewRequest(http.MethodGet, "/x", nil)
	w := httptest.NewRecorder()
	h(w, r)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("anonymous = %d, want 401", w.Code)
	}
}

func TestMiddlewareAnswers500WhenTheSessionCannotBeRead(t *testing.T) {
	store := testStore(t)
	store.Pool.Close()
	a := &Authenticator{Store: store, Log: slog.New(slog.NewTextHandler(io.Discard, nil))}
	h := a.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("the handler should not be reached")
	}))
	r := httptest.NewRequest(http.MethodGet, "/v1/me", nil)
	r.AddCookie(&http.Cookie{Name: SessionCookie, Value: "whatever"})
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", w.Code)
	}
}
