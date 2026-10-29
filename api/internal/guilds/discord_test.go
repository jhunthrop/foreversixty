// api/internal/guilds/discord_test.go
package guilds

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// fakeDiscord is a scripted guild channel; defaults, when set, records every
// SetGuildDiscordDefaults as {post, remind}, and the hook remembers them.
type fakeDiscord struct {
	enabled  bool
	hook     *notify.GuildWebhook
	defaults *[][2]bool
}

func (f fakeDiscord) SetGuildDiscordDefaults(_ context.Context, _ int64, post, remind bool) error {
	if f.defaults != nil {
		*f.defaults = append(*f.defaults, [2]bool{post, remind})
	}
	if f.hook != nil {
		f.hook.PostDefault, f.hook.RemindDefault = post, remind
	}
	return nil
}

func (f fakeDiscord) GuildConnectEnabled() bool { return f.enabled }

func (f fakeDiscord) GuildChannel(context.Context, int64) (notify.GuildWebhook, bool, error) {
	if f.hook == nil {
		return notify.GuildWebhook{}, false, nil
	}
	return *f.hook, true, nil
}

func switchedOffHook() *notify.GuildWebhook {
	discordGuild, reason, lastErr := "4242", "failed five times", "Unknown Webhook"
	return &notify.GuildWebhook{
		ChannelID: "1198004568", DiscordGuildID: &discordGuild, Failures: 5,
		DisabledReason: &reason, LastError: &lastErr, PostDefault: true,
	}
}

// settingsDiscordAs reads GET settings as the current actor and returns its raw discord key.
func (w raidWorld) settingsDiscordAs(t *testing.T) (json.RawMessage, bool) {
	t.Helper()
	res := w.h.do(http.MethodGet, w.path("/settings"), "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("settings = %d: %s", res.StatusCode, readBody(res))
	}
	var view map[string]json.RawMessage
	w.h.data(res, &view)
	raw, ok := view["discord"]
	return raw, ok
}

func TestSettingsCarryTheDiscordChannelForAnOfficer(t *testing.T) {
	w := newRaidWorld(t, true)
	w.h.svc.Discord = fakeDiscord{enabled: true, hook: switchedOffHook()}
	w.as(w.officer)
	raw, ok := w.settingsDiscordAs(t)
	if !ok {
		t.Fatal("no discord key for an officer")
	}
	var got SettingsDiscord
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if !got.Available || !got.Connected || got.ChannelID == nil || *got.ChannelID != "1198004568" ||
		got.DiscordGuildID == nil || *got.DiscordGuildID != "4242" || got.DisabledReason == nil ||
		*got.DisabledReason != "failed five times" || got.LastError == nil || got.Failures != 5 || !got.PostDefault {
		t.Fatalf("discord = %s", raw)
	}

	w.h.svc.Discord = fakeDiscord{enabled: true}
	raw, _ = w.settingsDiscordAs(t)
	if string(raw) != `{"available":true,"connected":false,"channel_id":null,"discord_guild_id":null,`+
		`"disabled_reason":null,"last_error":null,"failures":0,"post_default":false,`+
		`"bot_available":false,"bot":null,"bot_refusal":null}` {
		t.Fatalf("not connected = %s", raw)
	}
}

func TestSettingsLeaveDiscordOutWhenOffOrNotAnOfficer(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)
	if _, ok := w.settingsDiscordAs(t); ok {
		t.Fatal("discord key with no channel wired")
	}
	w.h.svc.Discord = fakeDiscord{enabled: false, hook: switchedOffHook()}
	if _, ok := w.settingsDiscordAs(t); ok {
		t.Fatal("discord key with connect off")
	}
	w.h.svc.Discord = fakeDiscord{enabled: true, hook: switchedOffHook()}
	moderator := seedUser(t, w.h.pool, "mod@example.com")
	w.h.actor = auth.Actor{UserID: moderator, Role: "moderator", Method: "session"}
	if _, ok := w.settingsDiscordAs(t); ok {
		t.Fatal("a moderator saw the guild's channel")
	}
}

func TestRaidViewCarriesTheGuildChannel(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	if got := w.getRaid(t, raid.ID).Discord; got != (RaidDiscord{}) {
		t.Fatalf("no channel wired: discord = %+v", got)
	}
	w.h.svc.Discord = fakeDiscord{enabled: true}
	if got := w.getRaid(t, raid.ID).Discord; got != (RaidDiscord{Available: true}) {
		t.Fatalf("not connected: discord = %+v", got)
	}
	w.h.svc.Discord = fakeDiscord{enabled: true, hook: switchedOffHook()}
	w.as(w.member)
	if got := w.getRaid(t, raid.ID).Discord; got != (RaidDiscord{Available: true, Connected: true, PostDefault: true}) {
		t.Fatalf("connected: discord = %+v", got)
	}
	w.as(w.officer)
	created := w.createRaid(t)
	if created.Discord != (RaidDiscord{Available: true, Connected: true, PostDefault: true}) {
		t.Fatalf("create answer: discord = %+v", created.Discord)
	}
}

// TestRaidViewCountsAnActiveBotAsConnected pins a guild with only the bot (spec state 5's
// primary path): the raid view offers the box, a ticked create reaches the poster (the
// bot's router), and the box is remembered on the install.
func TestRaidViewCountsAnActiveBotAsConnected(t *testing.T) {
	w, events, hookDefaults := postingWorld(t)
	w.h.svc.Discord = fakeDiscord{enabled: true, defaults: hookDefaults}
	botDefaults := &[][2]bool{}
	w.h.svc.Bot = fakeBot{on: true, defaults: botDefaults}
	raid := w.createRaid(t)
	if got := w.getRaid(t, raid.ID).Discord; got != (RaidDiscord{Available: true}) {
		t.Fatalf("bot installed but not active: discord = %+v", got)
	}

	w.h.svc.Bot = fakeBot{on: true, active: true, postDefault: true, defaults: botDefaults}
	w.as(w.member)
	if got := w.getRaid(t, raid.ID).Discord; got != (RaidDiscord{Available: true, Connected: true, PostDefault: true}) {
		t.Fatalf("bot only: discord = %+v", got)
	}
	created := w.createRaids(t, postingBody(time.Now().Add(3*24*time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339), 0, "true"))
	if created[0].Discord != (RaidDiscord{Available: true, Connected: true, PostDefault: true}) || len(events.posted) != 1 {
		t.Fatalf("create answer: discord = %+v, posted %d", created[0].Discord, len(events.posted))
	}
	// The first create carried the box unticked, the second ticked.
	if len(*botDefaults) != 2 || (*botDefaults)[0][0] || !(*botDefaults)[1][0] {
		t.Fatalf("bot post_default writes = %v", *botDefaults)
	}

	// The bot without connect still counts: the router posts through it.
	w.h.svc.Discord = nil
	if got := w.getRaid(t, raid.ID).Discord; got != (RaidDiscord{Available: true, Connected: true, PostDefault: true}) {
		t.Fatalf("bot, connect off: discord = %+v", got)
	}
}

func TestGuildStandingFollowsRankPlanAndFreeze(t *testing.T) {
	ctx := context.Background()
	w := newRaidWorld(t, false)
	standing, found, err := w.h.svc.GuildStanding(ctx, w.gid, w.officer)
	if err != nil || !found || !standing.Officer || standing.Plan || standing.Frozen ||
		standing.Name != "Olympus" || standing.Path != "/guild/us/hardcore/olympus" {
		t.Fatalf("officer without plan = %+v found=%v err=%v", standing, found, err)
	}
	if _, err := w.h.pool.Exec(ctx, scopedGrantSQL, w.gid); err != nil {
		t.Fatal(err)
	}
	if _, err := w.h.pool.Exec(ctx, `update guilds set claim_contested_at = now() where id = $1`, w.gid); err != nil {
		t.Fatal(err)
	}
	if standing, _, _ := w.h.svc.GuildStanding(ctx, w.gid, w.officer); !standing.Plan || !standing.Frozen {
		t.Fatalf("officer with plan, frozen = %+v", standing)
	}
	if standing, _, _ := w.h.svc.GuildStanding(ctx, w.gid, w.member); standing.Officer || standing.Plan {
		t.Fatalf("member = %+v", standing)
	}
	if _, found, err := w.h.svc.GuildStanding(ctx, w.gid+999, w.officer); found || err != nil {
		t.Fatalf("missing guild: found=%v err=%v", found, err)
	}
}

func TestGuildPagePathEscapesTheSlug(t *testing.T) {
	if got := GuildPagePath(Guild{Region: "US", Ruleset: "Normal", Name: "Les Pêcheurs/Nuit"}); got != "/guild/us/normal/les-p%C3%AAcheurs%2Fnuit" {
		t.Fatalf("path = %q", got)
	}
}
