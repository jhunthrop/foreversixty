package events

import (
	"context"
	"fmt"
	"os"
	"slices"
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
	if _, err := pool.Exec(context.Background(), `truncate users, guilds cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

func exec(t *testing.T, pool *pgxpool.Pool, q string, args ...any) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), q, args...); err != nil {
		t.Fatalf("%s: %v", q, err)
	}
}

func seedUser(t *testing.T, pool *pgxpool.Pool, name string) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(), `insert into users (email) values ($1) returning id`, name+"@example.com").Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// guildWorld is a guild with a leader, an officer, two members and one who has not been
// verified, and one raid the leader and a member signed up for.
type guildWorld struct {
	pool                                    *pgxpool.Pool
	guild, raid                             int64
	leader, officer, inMember, out, pending int64
}

func newGuildWorld(t *testing.T) guildWorld {
	t.Helper()
	pool := testPool(t)
	w := guildWorld{pool: pool}
	w.leader, w.officer = seedUser(t, pool, "leader"), seedUser(t, pool, "officer")
	w.inMember, w.out, w.pending = seedUser(t, pool, "in"), seedUser(t, pool, "out"), seedUser(t, pool, "pending")
	if err := pool.QueryRow(context.Background(), `insert into guilds (region, ruleset, name) values ('us','normal','G') returning id`).Scan(&w.guild); err != nil {
		t.Fatal(err)
	}
	for uid, rank := range map[int64]string{w.leader: "leader", w.officer: "officer", w.inMember: "member", w.out: "member"} {
		exec(t, pool, `insert into guild_members (guild_id, user_id, rank, verified_at) values ($1, $2, $3, now())`, w.guild, uid, rank)
	}
	exec(t, pool, `insert into guild_members (guild_id, user_id) values ($1, $2)`, w.guild, w.pending)
	if err := pool.QueryRow(context.Background(), `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, 'mc', 'MC', now() + interval '1 day', 180, $2) returning id`, w.guild, w.leader).Scan(&w.raid); err != nil {
		t.Fatal(err)
	}
	signups := []struct {
		uid   int64
		key   string
		state string
	}{{w.leader, "us/normal/lead", "late"}, {w.inMember, "us/normal/inny", "in"}, {w.inMember, "us/normal/alt", "bench"}, {w.out, "us/normal/outy", "out"}}
	for _, s := range signups {
		exec(t, pool, `insert into guild_raid_signups (raid_id, character_key, name, class, spec, role, state, signed_by, acted_by)
			values ($1, $2, $2, 'warrior', 'Arms', 'dps', $3, $4, $4)`, w.raid, s.key, s.state, s.uid)
	}
	return w
}

func TestRecipientsOfEveryAudience(t *testing.T) {
	w := newGuildWorld(t)
	st := &Store{Pool: w.pool}
	ids := func(a ...int64) []int64 { slices.Sort(a); return a }
	cases := map[AudienceKind][]int64{
		ToGuildMembers:  ids(w.leader, w.officer, w.inMember, w.out),
		ToGuildOfficers: ids(w.leader, w.officer),
		ToRaidSigned:    ids(w.leader, w.inMember, w.out),
		ToRaidInLate:    ids(w.leader, w.inMember),
		ToUnanswered:    ids(w.officer),
	}
	for kind, want := range cases {
		got, err := st.Recipients(context.Background(), Audience{Kind: kind, GuildID: w.guild, RaidID: w.raid})
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, want) {
			t.Errorf("%s = %v, want %v", kind, got, want)
		}
	}
	// Another guild's id never reads this guild's raid.
	got, err := st.Recipients(context.Background(), Audience{Kind: ToRaidSigned, GuildID: w.guild + 1, RaidID: w.raid})
	if err != nil || len(got) != 0 {
		t.Fatalf("cross guild = %v %v", got, err)
	}
	if got, _ := st.Recipients(context.Background(), Users(3, 1, 3)); !slices.Equal(got, []int64{1, 3}) {
		t.Fatalf("users = %v", got)
	}
}

func TestSettingsRoundTrip(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid := seedUser(t, pool, "p")
	ctx := context.Background()
	if p, err := st.Prefs(ctx, uid); err != nil || p.TimeZone != DefaultTimeZone || p.Quiet.Start != DefaultQuietStart || len(p.Overrides) != 0 {
		t.Fatalf("fresh = %+v %v", p, err)
	}
	zone := "Europe/Berlin"
	if err := st.SaveSettings(ctx, uid, SettingsChange{TimeZone: &zone, Prefs: []PrefChange{{RaidPosted, ChannelPush, false}}}); err != nil {
		t.Fatal(err)
	}
	quiet := QuietHours{Enabled: true, Start: 60, End: 120}
	if err := st.SaveSettings(ctx, uid, SettingsChange{Quiet: &quiet, Prefs: []PrefChange{{BisChanged, ChannelPush, true}}}); err != nil {
		t.Fatal(err)
	}
	p, err := st.Prefs(ctx, uid)
	if err != nil {
		t.Fatal(err)
	}
	if p.TimeZone != zone || p.Quiet != quiet || p.Allows(RaidPosted, ChannelPush) || !p.Allows(BisChanged, ChannelPush) {
		t.Fatalf("saved = %+v", p)
	}
	if st.SaveSettings(ctx, uid, SettingsChange{}) == nil {
		t.Fatal("an empty change was saved")
	}
}

func TestClaimIsOncePerUserKeyAndChannel(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid := seedUser(t, pool, "c")
	ctx := context.Background()
	first, err := st.Claim(ctx, uid, "k", ChannelPush)
	second, _ := st.Claim(ctx, uid, "k", ChannelPush)
	other, _ := st.Claim(ctx, uid, "k", ChannelDiscord)
	if err != nil || !first || second || !other {
		t.Fatalf("claims %v %v %v %v", first, second, other, err)
	}
	if err := st.Unclaim(ctx, uid, "k", ChannelPush); err != nil {
		t.Fatal(err)
	}
	if again, _ := st.Claim(ctx, uid, "k", ChannelPush); !again {
		t.Fatal("unclaimed key not claimable")
	}
	exec(t, pool, `update notification_deliveries set sent_at = now() - interval '40 days' where channel = 'discord'`)
	n, err := st.PruneDeliveries(ctx, time.Now(), 100)
	if err != nil || n != 1 {
		t.Fatalf("pruned %d %v", n, err)
	}
}

func TestRouterOverTheStore(t *testing.T) {
	w := newGuildWorld(t)
	sink := &fakeSink{}
	r := &Router{Store: &Store{Pool: w.pool}, Sinks: map[Channel]Sink{ChannelPush: sink}}
	e := Event{Kind: RaidStarting, Audience: Audience{Kind: ToRaidInLate, GuildID: w.guild, RaidID: w.raid},
		DedupeKey: fmt.Sprintf("raid.starting:%d", w.raid)}
	for range 2 {
		if err := r.Publish(context.Background(), e); err != nil {
			t.Fatal(err)
		}
	}
	if len(sink.got) != 2 {
		t.Fatalf("sent %v", sink.got)
	}
}
