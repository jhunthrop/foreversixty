package notify

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func seedGuild(t *testing.T, pool *pgxpool.Pool) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us', 'normal', $1) returning id`,
		"Notify "+time.Now().Format("150405.000000000")).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func savedGuild(t *testing.T) (*GuildStore, int64, int64) {
	t.Helper()
	pool := testPool(t)
	st := &GuildStore{Pool: pool}
	uid, gid := seedUser(t, pool), seedGuild(t, pool)
	discordGuild := "4242"
	if err := st.Save(context.Background(), GuildSaveInput{
		GuildID: gid, ChannelID: "555", DiscordGuildID: &discordGuild, TokenBox: []byte("box-1"), ConnectedBy: uid,
	}); err != nil {
		t.Fatal(err)
	}
	return st, gid, uid
}

func TestGuildStoreSaveGetBoxDelete(t *testing.T) {
	ctx := context.Background()
	st, gid, uid := savedGuild(t)
	w, ok, err := st.Get(ctx, gid)
	if err != nil || !ok {
		t.Fatalf("Get = %v, %v", ok, err)
	}
	if w.ChannelID != "555" || w.DiscordGuildID == nil || *w.DiscordGuildID != "4242" || w.ConnectedBy == nil ||
		*w.ConnectedBy != uid || w.Failures != 0 || w.DisabledReason != nil || w.LastError != nil || !w.PostDefault || !w.RemindDefault {
		t.Fatalf("stored = %+v", w)
	}
	box, err := st.TokenBox(ctx, gid)
	if err != nil || string(box) != "box-1" {
		t.Fatalf("TokenBox = %q, %v", box, err)
	}
	if err := st.Delete(ctx, gid); err != nil {
		t.Fatal(err)
	}
	if _, ok, err := st.Get(ctx, gid); ok || err != nil {
		t.Fatalf("after delete: ok=%v err=%v", ok, err)
	}
	if _, err := st.TokenBox(ctx, gid); err != ErrNotFound {
		t.Fatalf("TokenBox after delete = %v, want ErrNotFound", err)
	}
	if err := st.Delete(ctx, gid); err != nil {
		t.Fatalf("deleting an absent channel: %v", err)
	}
}

func TestGuildStoreFiveCountedFailuresSwitchOff(t *testing.T) {
	ctx := context.Background()
	st, gid, _ := savedGuild(t)
	for i := 1; i <= MaxFailures; i++ {
		off, err := st.RecordFailure(ctx, gid, "boom", true)
		if err != nil {
			t.Fatal(err)
		}
		if off != (i == MaxFailures) {
			t.Fatalf("failure %d: off = %v", i, off)
		}
	}
	w, _, _ := st.Get(ctx, gid)
	if w.Failures != MaxFailures || w.DisabledReason == nil || *w.DisabledReason != reasonFailedFiveTimes ||
		w.LastError == nil || *w.LastError != "boom" {
		t.Fatalf("after five = %+v", w)
	}
}

func TestGuildStoreUncountedFailureKeepsTheCounter(t *testing.T) {
	ctx := context.Background()
	st, gid, _ := savedGuild(t)
	for i := 0; i < MaxFailures+2; i++ {
		if off, err := st.RecordFailure(ctx, gid, "slow down", false); err != nil || off {
			t.Fatalf("uncounted: off=%v err=%v", off, err)
		}
	}
	if w, _, _ := st.Get(ctx, gid); w.Failures != 0 || w.DisabledReason != nil {
		t.Fatalf("uncounted failures counted: %+v", w)
	}
}

func TestGuildStoreSuccessAndResaveReset(t *testing.T) {
	ctx := context.Background()
	st, gid, uid := savedGuild(t)
	for i := 0; i < MaxFailures; i++ {
		if _, err := st.RecordFailure(ctx, gid, "boom", true); err != nil {
			t.Fatal(err)
		}
	}
	if err := st.RecordSuccess(ctx, gid); err != nil {
		t.Fatal(err)
	}
	if w, _, _ := st.Get(ctx, gid); w.Failures != 0 || w.DisabledReason != nil || w.LastError != nil {
		t.Fatalf("success did not reset: %+v", w)
	}
	if _, err := st.Pool.Exec(ctx, `update guild_notification_webhooks set post_default = true, failures = 3 where guild_id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	if err := st.Save(ctx, GuildSaveInput{GuildID: gid, ChannelID: "777", TokenBox: []byte("box-2"), ConnectedBy: uid}); err != nil {
		t.Fatal(err)
	}
	w, _, _ := st.Get(ctx, gid)
	if w.ChannelID != "777" || w.DiscordGuildID != nil || w.Failures != 0 || !w.PostDefault {
		t.Fatalf("re-save = %+v (post_default must survive, the counter must not)", w)
	}
	if _, err := st.RecordFailure(ctx, gid+999, "x", true); err != ErrNotFound {
		t.Fatalf("absent row = %v, want ErrNotFound", err)
	}
}

func TestConnectStatesAreSingleUseAndExpire(t *testing.T) {
	ctx := context.Background()
	st, gid, uid := savedGuild(t)
	if err := st.CreateConnectState(ctx, "state-1", uid, gid, time.Minute); err != nil {
		t.Fatal(err)
	}
	got, ok, err := st.TakeConnectState(ctx, "state-1")
	if err != nil || !ok || got.UserID != uid || got.GuildID != gid || got.Expired {
		t.Fatalf("take = %+v ok=%v err=%v", got, ok, err)
	}
	if _, ok, err := st.TakeConnectState(ctx, "state-1"); ok || err != nil {
		t.Fatalf("second take: ok=%v err=%v", ok, err)
	}
	if err := st.CreateConnectState(ctx, "state-2", uid, gid, -time.Minute); err != nil {
		t.Fatal(err)
	}
	if got, ok, _ := st.TakeConnectState(ctx, "state-2"); !ok || !got.Expired {
		t.Fatalf("expired take = %+v ok=%v", got, ok)
	}
	if _, ok, _ := st.TakeConnectState(ctx, ""); ok {
		t.Fatal("an empty state matched")
	}
}

func TestGuildCheckRejectsANonLinkStateWithoutAGuild(t *testing.T) {
	st, _, uid := savedGuild(t)
	for _, kind := range []string{StateKindGuildWebhook, StateKindBotInstall} {
		if err := st.CreateState(context.Background(), "nog-"+kind, kind, uid, nil, time.Minute); err == nil {
			t.Fatalf("kind %s with a null guild was accepted", kind)
		}
	}
	if err := st.CreateState(context.Background(), "bad-kind", "other", uid, nil, time.Minute); err == nil {
		t.Fatal("unknown kind was accepted")
	}
}

func TestCreateStateRoundTripsKind(t *testing.T) {
	ctx := context.Background()
	st, gid, uid := savedGuild(t)
	if err := st.CreateState(ctx, "s1", StateKindLink, uid, nil, time.Minute); err != nil {
		t.Fatal(err)
	}
	cs, ok, err := st.TakeConnectState(ctx, "s1")
	if err != nil || !ok || cs.Kind != StateKindLink || cs.GuildID != 0 || cs.UserID != uid {
		t.Fatalf("link state = %+v ok=%v err=%v", cs, ok, err)
	}
	if err := st.CreateConnectState(ctx, "s2", uid, gid, time.Minute); err != nil {
		t.Fatal(err)
	}
	cs, ok, err = st.TakeConnectState(ctx, "s2")
	if err != nil || !ok || cs.Kind != StateKindGuildWebhook || cs.GuildID != gid {
		t.Fatalf("webhook state = %+v ok=%v err=%v", cs, ok, err)
	}
	if err := st.CreateState(ctx, "s3", StateKindBotInstall, uid, &gid, time.Minute); err != nil {
		t.Fatal(err)
	}
	if cs, _, _ := st.TakeConnectState(ctx, "s3"); cs.Kind != StateKindBotInstall || cs.GuildID != gid {
		t.Fatalf("install state = %+v", cs)
	}
}

func TestDefaultsStartOnAndAreRemembered(t *testing.T) {
	ctx := context.Background()
	st, gid, _ := savedGuild(t)
	if err := st.SetDefaults(ctx, gid, false, true); err != nil {
		t.Fatal(err)
	}
	w, _, err := st.Get(ctx, gid)
	if err != nil || w.PostDefault || !w.RemindDefault {
		t.Fatalf("after SetDefaults(false,true) = %+v, %v", w, err)
	}
	if err := st.SetDefaults(ctx, gid, true, false); err != nil {
		t.Fatal(err)
	}
	if w, _, _ = st.Get(ctx, gid); !w.PostDefault || w.RemindDefault {
		t.Fatalf("after SetDefaults(true,false) = %+v", w)
	}
}
