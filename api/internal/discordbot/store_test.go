package discordbot

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
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
	if _, err := pool.Exec(context.Background(), `truncate users, guilds, reports, fight_metrics cascade`); err != nil {
		t.Fatal(err)
	}
	return &Store{Pool: pool}
}

var seq atomic.Int64

func seedUser(t *testing.T, pool *pgxpool.Pool) int64 {
	t.Helper()
	var id int64
	err := pool.QueryRow(context.Background(), `insert into users (email) values ($1) returning id`,
		fmt.Sprintf("d-%d@example.com", seq.Add(1))).Scan(&id)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func seedGuild(t *testing.T, pool *pgxpool.Pool, name string) int64 {
	t.Helper()
	var id int64
	err := pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us', 'normal', $1) returning id`, name).Scan(&id)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func addMember(t *testing.T, pool *pgxpool.Pool, guildID, userID int64, verified bool) {
	t.Helper()
	q := `insert into guild_members (guild_id, user_id) values ($1, $2)`
	if verified {
		q = `insert into guild_members (guild_id, user_id, verified_at) values ($1, $2, now())`
	}
	if _, err := pool.Exec(context.Background(), q, guildID, userID); err != nil {
		t.Fatal(err)
	}
}

func TestBindServerOneServerOneGuild(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	u := seedUser(t, st.Pool)
	a, b := seedGuild(t, st.Pool, "A"), seedGuild(t, st.Pool, "B")

	if r, err := st.BindServer(ctx, a, "srv1", "Srv One", u); err != nil || r.Outcome != BindInstalled || r.HeldBy != 0 {
		t.Fatalf("first bind = %+v err=%v", r, err)
	}
	if r, err := st.BindServer(ctx, b, "srv1", "Srv One", u); err != nil || r.Outcome != BindTaken || r.HeldBy != a {
		t.Fatalf("other guild = %+v err=%v", r, err)
	}
	if r, err := st.BindServer(ctx, a, "srv2", "Srv Two", u); err != nil || r.Outcome != BindTaken || r.HeldBy != a {
		t.Fatalf("other server = %+v err=%v", r, err)
	}
	if err := st.SetChannel(ctx, a, "chan1", "raids"); err != nil {
		t.Fatal(err)
	}
	if r, err := st.BindServer(ctx, a, "srv1", "Srv One Renamed", u); err != nil || r.Outcome != BindInstalled {
		t.Fatalf("rebind = %+v err=%v", r, err)
	}
	in, ok, err := st.InstallFor(ctx, a)
	if err != nil || !ok || in.ChannelID == nil || *in.ChannelID != "chan1" || in.ServerName != "Srv One Renamed" {
		t.Fatalf("install = %+v ok=%v err=%v", in, ok, err)
	}
	if by, ok, _ := st.InstallByServer(ctx, "srv1"); !ok || by.GuildID != a {
		t.Fatalf("by server = %+v ok=%v", by, ok)
	}
}

// A new install takes the webhook's remembered "Post to Discord" box; a rebind of the same
// server keeps the install's own.
func TestBindServerCopiesTheWebhookPostDefault(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	u := seedUser(t, st.Pool)
	ticked, bare := seedGuild(t, st.Pool, "Ticked"), seedGuild(t, st.Pool, "Bare")
	webhookOf(t, st.Pool, ticked, u, &notify.FakePoster{})
	if _, err := st.Pool.Exec(ctx, `update guild_notification_webhooks set post_default = true where guild_id = $1`, ticked); err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		guild  int64
		server string
		want   bool
	}{{ticked, "srv-ticked", true}, {bare, "srv-bare", false}} {
		if _, err := st.BindServer(ctx, c.guild, c.server, "S", u); err != nil {
			t.Fatal(err)
		}
		if in, _, _ := st.InstallFor(ctx, c.guild); in.PostDefault != c.want {
			t.Fatalf("guild %d: post_default = %v, want %v", c.guild, in.PostDefault, c.want)
		}
	}
	if err := st.SetPostDefault(ctx, ticked, false); err != nil {
		t.Fatal(err)
	}
	if _, err := st.BindServer(ctx, ticked, "srv-ticked", "S", u); err != nil {
		t.Fatal(err)
	}
	if in, _, _ := st.InstallFor(ctx, ticked); in.PostDefault {
		t.Fatal("a rebind copied the webhook's box over the install's own")
	}
}

func TestBindServerConcurrentGuildsOneWins(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	u := seedUser(t, st.Pool)
	gs := []int64{seedGuild(t, st.Pool, "A"), seedGuild(t, st.Pool, "B")}
	results := make([]BindResult, len(gs))
	errs := make([]error, len(gs))
	var wg sync.WaitGroup
	for i := range gs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i], errs[i] = st.BindServer(ctx, gs[i], "srv1", "S", u)
		}()
	}
	wg.Wait()
	installed := 0
	for i, r := range results {
		if errs[i] != nil {
			t.Fatalf("bind %d: %v", i, errs[i])
		}
		if r.Outcome == BindInstalled {
			installed++
		} else if r.HeldBy == 0 || r.HeldBy == gs[i] {
			t.Fatalf("taken result %d = %+v, want the other guild", i, r)
		}
	}
	if installed != 1 {
		t.Fatalf("installed = %d, want exactly 1", installed)
	}
}

func TestSaveLinkConcurrentUsersOneWins(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	users := []int64{seedUser(t, st.Pool), seedUser(t, st.Pool)}
	errs := make([]error, len(users))
	var wg sync.WaitGroup
	for i := range users {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs[i] = st.SaveLink(ctx, users[i], "same", "N")
		}()
	}
	wg.Wait()
	conflicts := 0
	for _, err := range errs {
		switch {
		case err == nil:
		case errors.Is(err, ErrLinkConflict):
			conflicts++
		default:
			t.Fatalf("unexpected error: %v", err)
		}
	}
	if conflicts != 1 {
		t.Fatalf("conflicts = %d, want exactly 1", conflicts)
	}
}

func TestInstallFailureBookkeepingAndDelete(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	u := seedUser(t, st.Pool)
	g := seedGuild(t, st.Pool, "A")
	if _, err := st.BindServer(ctx, g, "srv1", "S", u); err != nil {
		t.Fatal(err)
	}
	if err := st.RecordPostFailure(ctx, g, "boom"); err != nil {
		t.Fatal(err)
	}
	if in, _, _ := st.InstallFor(ctx, g); in.Failures != 1 || in.LastError == nil || *in.LastError != "boom" {
		t.Fatalf("after failure = %+v", in)
	}
	if err := st.RecordPostSuccess(ctx, g); err != nil {
		t.Fatal(err)
	}
	if in, _, _ := st.InstallFor(ctx, g); in.Failures != 0 || in.LastError != nil {
		t.Fatalf("after success = %+v", in)
	}
	if err := st.MarkRemovedInDiscord(ctx, g, "403"); err != nil {
		t.Fatal(err)
	}
	if in, _, _ := st.InstallFor(ctx, g); !in.RemovedInDiscord {
		t.Fatalf("not marked removed: %+v", in)
	}
	if _, err := st.BindServer(ctx, g, "srv1", "S", u); err != nil {
		t.Fatal(err)
	}
	if in, _, _ := st.InstallFor(ctx, g); in.RemovedInDiscord {
		t.Fatalf("rebind should clear removed: %+v", in)
	}
	if err := st.RemoveBot(ctx, g); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := st.InstallFor(ctx, g); ok {
		t.Fatal("install survived RemoveBot")
	}
}

func TestSaveLinkOneToOne(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	u1, u2 := seedUser(t, st.Pool), seedUser(t, st.Pool)
	if err := st.SaveLink(ctx, u1, "d1", "One"); err != nil {
		t.Fatal(err)
	}
	if err := st.SaveLink(ctx, u2, "d1", "Two"); !errors.Is(err, ErrLinkConflict) {
		t.Fatalf("conflict err = %v", err)
	}
	if err := st.SaveLink(ctx, u1, "d9", "One B"); err != nil {
		t.Fatal(err)
	}
	l, ok, err := st.LinkByUser(ctx, u1)
	if err != nil || !ok || l.DiscordUserID != "d9" || l.DisplayName != "One B" {
		t.Fatalf("link = %+v ok=%v err=%v", l, ok, err)
	}
	if _, ok, _ := st.UserByDiscord(ctx, "d1"); ok {
		t.Fatal("old discord id still resolves")
	}
	if id, ok, _ := st.UserByDiscord(ctx, "d9"); !ok || id != u1 {
		t.Fatalf("by discord = %d ok=%v", id, ok)
	}
	if err := st.DeleteLink(ctx, u1); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := st.LinkByUser(ctx, u1); ok {
		t.Fatal("link survived delete")
	}
}

func TestLinkedMembersCountsVerifiedOnly(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	g := seedGuild(t, st.Pool, "A")
	v1, v2, pending := seedUser(t, st.Pool), seedUser(t, st.Pool), seedUser(t, st.Pool)
	addMember(t, st.Pool, g, v1, true)
	addMember(t, st.Pool, g, v2, true)
	addMember(t, st.Pool, g, pending, false)
	if err := st.SaveLink(ctx, v1, "d1", "V1"); err != nil {
		t.Fatal(err)
	}
	if err := st.SaveLink(ctx, pending, "d3", "P"); err != nil {
		t.Fatal(err)
	}
	linked, members, err := st.LinkedMembers(ctx, g)
	if err != nil || linked != 1 || members != 2 {
		t.Fatalf("linked=%d members=%d err=%v", linked, members, err)
	}
}

func TestSuggestedGuildNeedsAChannelAndAPresentBot(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	u := seedUser(t, st.Pool)
	g := seedGuild(t, st.Pool, "Olympus XXVII")
	addMember(t, st.Pool, g, u, true)
	if _, err := st.BindServer(ctx, g, "srv1", "S", u); err != nil {
		t.Fatal(err)
	}
	if got, err := st.SuggestedGuild(ctx, u); err != nil || got != nil {
		t.Fatalf("no channel = %v err=%v", got, err)
	}
	if err := st.SetChannel(ctx, g, "c1", "raids"); err != nil {
		t.Fatal(err)
	}
	if got, err := st.SuggestedGuild(ctx, u); err != nil || got == nil || *got != "Olympus XXVII" {
		t.Fatalf("with channel = %v err=%v", got, err)
	}
	if err := st.MarkRemovedInDiscord(ctx, g, "gone"); err != nil {
		t.Fatal(err)
	}
	if got, err := st.SuggestedGuild(ctx, u); err != nil || got != nil {
		t.Fatalf("removed = %v err=%v", got, err)
	}
}

func TestTakeRefusalReadsOnceAndExpires(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	u := seedUser(t, st.Pool)
	g := seedGuild(t, st.Pool, "A")
	if err := st.SaveRefusal(ctx, u, g, "Mograine Raiders", time.Minute); err != nil {
		t.Fatal(err)
	}
	if got, err := st.TakeRefusal(ctx, u, g); err != nil || got == nil || *got != "Mograine Raiders" {
		t.Fatalf("take = %v err=%v", got, err)
	}
	if got, err := st.TakeRefusal(ctx, u, g); err != nil || got != nil {
		t.Fatalf("second take = %v err=%v", got, err)
	}
	if err := st.SaveRefusal(ctx, u, g, "Old", -time.Second); err != nil {
		t.Fatal(err)
	}
	if got, err := st.TakeRefusal(ctx, u, g); err != nil || got != nil {
		t.Fatalf("expired take = %v err=%v", got, err)
	}
}
