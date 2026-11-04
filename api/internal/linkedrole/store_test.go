package linkedrole

import (
	"context"
	"os"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func testStore(t *testing.T) *PGStore {
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
	return &PGStore{Pool: pool}
}

func seedAccount(t *testing.T, s *PGStore, battletag string) int64 {
	t.Helper()
	var id int64
	if err := s.Pool.QueryRow(context.Background(), `insert into users (battletag) values ($1) returning id`, battletag).Scan(&id); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _, _ = s.Pool.Exec(context.Background(), `delete from users where id = $1`, id) })
	return id
}

func TestPGStoreLinkLifecycle(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	a, b := seedAccount(t, s, "lr-a#1"), seedAccount(t, s, "lr-b#2")
	discordID := "lr-discord-" + t.Name()

	if err := s.Save(ctx, Link{UserID: a, DiscordUserID: discordID, RefreshBox: []byte("box0")}); err != nil {
		t.Fatal(err)
	}
	if ok, err := s.Rotate(ctx, a, []byte("stale"), []byte("x")); err != nil || ok {
		t.Fatalf("rotate from a stale box = %v, %v; want false", ok, err)
	}
	if ok, err := s.Rotate(ctx, a, []byte("box0"), []byte("box1")); err != nil || !ok {
		t.Fatalf("rotate = %v, %v", ok, err)
	}
	if err := s.MarkPushed(ctx, a, false); err != nil {
		t.Fatal(err)
	}
	p, err := s.Profile(ctx, a)
	if err != nil || p.Premium || p.Battletag != "lr-a#1" {
		t.Fatalf("profile = %+v, %v", p, err)
	}

	// The same Discord account authorized from b moves to b.
	if err := s.Save(ctx, Link{UserID: b, DiscordUserID: discordID, RefreshBox: []byte("boxB")}); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := s.Get(ctx, a); ok {
		t.Fatal("a still holds the Discord account b took")
	}
	ids, err := s.LinkedUsers(ctx, 0)
	if err != nil || !contains(ids, b) {
		t.Fatalf("linked users = %v, %v", ids, err)
	}
	if ok, err := s.DeleteIf(ctx, b, []byte("other")); err != nil || ok {
		t.Fatalf("delete with a changed box = %v, %v; want kept", ok, err)
	}
	if ok, err := s.DeleteIf(ctx, b, []byte("boxB")); err != nil || !ok {
		t.Fatalf("delete = %v, %v", ok, err)
	}
}

func TestPGStoreTakesALinkedRoleState(t *testing.T) {
	s := testStore(t)
	a := seedAccount(t, s, "lr-c#3")
	if _, err := s.Pool.Exec(context.Background(), `insert into notify_connect_states (state, kind, user_id, guild_id, expires_at)
		values ($1, 'linked_role', $2, null, now() + interval '10 minutes')`, "lr-state-"+t.Name(), a); err != nil {
		t.Fatalf("a linked_role state with no guild was refused: %v", err)
	}
}

func contains(ids []int64, id int64) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}
