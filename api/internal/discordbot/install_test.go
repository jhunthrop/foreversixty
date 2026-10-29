package discordbot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const installPublic = "https://site.test"

// fakeGuilds answers standings by guild id and settings from the bot's own read.
type fakeGuilds struct {
	standings map[int64]notify.GuildStanding
	svc       *Service
}

func (f *fakeGuilds) GuildStanding(_ context.Context, guildID, _ int64) (notify.GuildStanding, bool, error) {
	st, ok := f.standings[guildID]
	return st, ok, nil
}

func (f *fakeGuilds) SettingsDiscordFor(ctx context.Context, guildID, viewerID int64) (*guilds.SettingsDiscord, error) {
	bot, refusal, err := f.svc.SettingsBot(ctx, guildID, viewerID)
	return &guilds.SettingsDiscord{BotAvailable: true, Bot: bot, BotRefusal: refusal}, err
}

// stubConnect only switches the notify service's guild connect on.
type stubConnect struct{}

func (stubConnect) AuthURL(string) string { return "" }

func (stubConnect) Exchange(context.Context, string) (notify.OAuthWebhook, error) {
	return notify.OAuthWebhook{}, nil
}

type installHarness struct {
	t       *testing.T
	svc     *Service
	rest    *FakeREST
	oauth   *FakeOAuth
	guilds  *fakeGuilds
	mux     *http.ServeMux
	log     *bytes.Buffer
	uid     int64
	gid     int64
	otherID int64
	// expire makes callback take its state already expired.
	expire bool
}

func officerStanding(name, path string) notify.GuildStanding {
	return notify.GuildStanding{Officer: true, Plan: true, Name: name, Path: path}
}

func newInstallHarness(t *testing.T) *installHarness {
	t.Helper()
	st := testStore(t)
	h := &installHarness{
		t: t, rest: &FakeREST{}, oauth: &FakeOAuth{Installed: InstalledServer{ID: "srv1", Name: "Olympus"}},
		mux: http.NewServeMux(), log: &bytes.Buffer{}, uid: seedUser(t, st.Pool),
	}
	h.gid = seedGuild(t, st.Pool, "Olympus XXVII")
	h.otherID = seedGuild(t, st.Pool, "Mograine Raiders")
	h.guilds = &fakeGuilds{standings: map[int64]notify.GuildStanding{
		h.gid:     officerStanding("Olympus XXVII", "/guild/us/normal/olympus-xxvii"),
		h.otherID: {Name: "Mograine Raiders", Path: "/guild/us/normal/mograine-raiders"},
	}}
	h.svc = New(Deps{
		Store: st, REST: h.rest, OAuth: h.oauth, PublicBaseURL: installPublic, APIBaseURL: "https://api.test",
		Guilds: h.guilds, Log: slog_(h.log),
		Notify: &notify.Service{OAuth: stubConnect{}, Guilds: h.guilds, GuildStore: &notify.GuildStore{Pool: st.Pool}},
	})
	h.guilds.svc = h.svc
	Mount(h.mux, h.svc, 0)
	base := PermViewChannel | PermSendMessages | PermEmbedLinks | PermReadMessageHistory
	h.rest.Layout = ServerLayout{
		BotUserID: "bot", Roles: []Role{{ID: "srv1", Permissions: base}},
		Channels: []RawChannel{
			{ID: "c2", Name: "general", Type: 0, Position: 2},
			{ID: "c1", Name: "raid-signups", Type: 0, Position: 1},
			{ID: "v1", Name: "voice", Type: 2},
			{ID: "locked", Name: "mods", Type: 0, Position: 3,
				Overwrites: []Overwrite{{ID: "srv1", Type: OverwriteRole, Deny: PermSendMessages}}},
		},
	}
	return h
}

func (h *installHarness) path(guildID int64, suffix string) string {
	return fmt.Sprintf("/v1/guilds/%d/discord/bot%s", guildID, suffix)
}

func (h *installHarness) do(method, target, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	req = req.WithContext(auth.WithActor(req.Context(), auth.Actor{UserID: h.uid}))
	rec := httptest.NewRecorder()
	h.mux.ServeHTTP(rec, req)
	return rec
}

func (h *installHarness) setStanding(change func(*notify.GuildStanding)) {
	st := h.guilds.standings[h.gid]
	change(&st)
	h.guilds.standings[h.gid] = st
}

// callback runs the shared callback's install leg as actor with a fresh state.
func (h *installHarness) callback(guildID int64, query string, actor auth.Actor) string {
	h.t.Helper()
	ctx := context.Background()
	state, err := h.svc.d.Notify.StartState(ctx, notify.StateKindBotInstall, h.uid, &guildID)
	if err != nil {
		h.t.Fatal(err)
	}
	if h.expire {
		_, err := h.svc.d.Store.Pool.Exec(ctx, `update notify_connect_states set expires_at = now() - interval '1 minute' where state = $1`, state)
		if err != nil {
			h.t.Fatal(err)
		}
	}
	cs, ok, err := h.svc.d.Notify.GuildStore.TakeConnectState(ctx, state)
	if err != nil || !ok {
		h.t.Fatalf("take state: ok=%v err=%v", ok, err)
	}
	req := httptest.NewRequest(http.MethodGet, "/v1/notifications/discord/callback?"+query, nil)
	req = req.WithContext(auth.WithActor(req.Context(), actor))
	rec := httptest.NewRecorder()
	h.svc.installCallback(rec, req, cs)
	return rec.Header().Get("Location")
}

func (h *installHarness) officerActor() auth.Actor { return auth.Actor{UserID: h.uid} }

func (h *installHarness) wantOutcome(loc, path, outcome string) {
	h.t.Helper()
	if want := installPublic + path + "/settings?discord_bot=" + outcome + "#discord"; loc != want {
		h.t.Fatalf("Location = %q, want %q", loc, want)
	}
}

const olympusPath = "/guild/us/normal/olympus-xxvii"

func (h *installHarness) settings(guildID int64) guilds.SettingsDiscord {
	h.t.Helper()
	view, err := h.guilds.SettingsDiscordFor(context.Background(), guildID, h.uid)
	if err != nil {
		h.t.Fatal(err)
	}
	return *view
}

func (h *installHarness) installFor(guildID int64) (Install, bool) {
	h.t.Helper()
	in, ok, err := h.svc.d.Store.InstallFor(context.Background(), guildID)
	if err != nil {
		h.t.Fatal(err)
	}
	return in, ok
}

func TestInstallStartGate(t *testing.T) {
	h := newInstallHarness(t)
	rec := h.do(http.MethodGet, h.path(h.gid, "/install"), "")
	loc, err := url.Parse(rec.Header().Get("Location"))
	if rec.Code != http.StatusFound || err != nil {
		t.Fatalf("status %d, Location %q", rec.Code, rec.Header().Get("Location"))
	}
	state := loc.Query().Get("state")
	if rec.Header().Get("Location") != h.oauth.InstallURL(state) {
		t.Fatalf("Location = %q", rec.Header().Get("Location"))
	}
	var kind string
	var guildID int64
	err = h.svc.d.Store.Pool.QueryRow(context.Background(),
		`select kind, guild_id from notify_connect_states where state = $1`, state).Scan(&kind, &guildID)
	if err != nil || kind != "bot_install" || guildID != h.gid {
		t.Fatalf("state row = %q guild %d (%v)", kind, guildID, err)
	}

	cases := []struct {
		name   string
		change func(*notify.GuildStanding)
		status int
		code   string
	}{
		{"member", func(s *notify.GuildStanding) { s.Officer = false }, http.StatusNotFound, "not_found"},
		{"no plan", func(s *notify.GuildStanding) { s.Plan = false }, http.StatusForbidden, "guild_plan_required"},
		{"frozen", func(s *notify.GuildStanding) { s.Frozen = true }, http.StatusConflict, "claim_contested"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := newInstallHarness(t)
			g.setStanding(tc.change)
			rec := g.do(http.MethodGet, g.path(g.gid, "/install"), "")
			if rec.Code != tc.status || !strings.Contains(rec.Body.String(), `"`+tc.code+`"`) {
				t.Fatalf("status %d: %s", rec.Code, rec.Body.String())
			}
		})
	}
}

func TestInstallCallbackOutcomes(t *testing.T) {
	t.Run("installed, and again keeps the channel", func(t *testing.T) {
		h := newInstallHarness(t)
		h.wantOutcome(h.callback(h.gid, "code=c0de", h.officerActor()), olympusPath, "installed")
		if codes := h.oauth.InstallCodes(); len(codes) != 1 || codes[0] != "c0de" {
			t.Fatalf("codes = %v", codes)
		}
		if err := h.svc.d.Store.SetChannel(context.Background(), h.gid, "c1", "raid-signups"); err != nil {
			t.Fatal(err)
		}
		h.wantOutcome(h.callback(h.gid, "code=c0de2", h.officerActor()), olympusPath, "installed")
		in, ok := h.installFor(h.gid)
		if !ok || in.ChannelID == nil || *in.ChannelID != "c1" || in.ServerName != "Olympus" {
			t.Fatalf("install = %+v", in)
		}
	})
	t.Run("another guild's server is taken and says who, once", func(t *testing.T) {
		h := newInstallHarness(t)
		if _, err := h.svc.d.Store.BindServer(context.Background(), h.otherID, "srv1", "Olympus", h.uid); err != nil {
			t.Fatal(err)
		}
		h.wantOutcome(h.callback(h.gid, "code=c", h.officerActor()), olympusPath, "taken")
		first := h.settings(h.gid)
		if first.BotRefusal == nil || first.BotRefusal.GuildName != "Mograine Raiders" || first.Bot != nil {
			t.Fatalf("settings = %+v", first)
		}
		if again := h.settings(h.gid); again.BotRefusal != nil {
			t.Fatalf("refusal read twice: %+v", again.BotRefusal)
		}
	})
	t.Run("a guild with a server installing another names itself", func(t *testing.T) {
		h := newInstallHarness(t)
		h.wantOutcome(h.callback(h.gid, "code=c", h.officerActor()), olympusPath, "installed")
		h.oauth.Installed = InstalledServer{ID: "srv2", Name: "Elsewhere"}
		h.wantOutcome(h.callback(h.gid, "code=c", h.officerActor()), olympusPath, "taken")
		if got := h.settings(h.gid); got.BotRefusal == nil || got.BotRefusal.GuildName != "Olympus XXVII" ||
			got.Bot == nil || got.Bot.ServerID != "srv1" {
			t.Fatalf("settings = %+v", got)
		}
	})
	t.Run("denied", func(t *testing.T) {
		h := newInstallHarness(t)
		h.wantOutcome(h.callback(h.gid, "error=access_denied", h.officerActor()), olympusPath, "denied")
		if _, ok := h.installFor(h.gid); ok || len(h.oauth.InstallCodes()) != 0 {
			t.Fatal("a denied install bound or exchanged")
		}
	})
	t.Run("a Discord error other than a refusal fails", func(t *testing.T) {
		h := newInstallHarness(t)
		h.wantOutcome(h.callback(h.gid, "error=server_error&code=c", h.officerActor()), olympusPath, "failed")
		if len(h.oauth.InstallCodes()) != 0 {
			t.Fatal("exchanged on a Discord error")
		}
	})
	t.Run("failed", func(t *testing.T) {
		cases := map[string]func(*installHarness) auth.Actor{
			"demoted": func(h *installHarness) auth.Actor {
				h.setStanding(func(s *notify.GuildStanding) { s.Officer = false })
				return h.officerActor()
			},
			"no plan": func(h *installHarness) auth.Actor {
				h.setStanding(func(s *notify.GuildStanding) { s.Plan = false })
				return h.officerActor()
			},
			"frozen": func(h *installHarness) auth.Actor {
				h.setStanding(func(s *notify.GuildStanding) { s.Frozen = true })
				return h.officerActor()
			},
			"expired state": func(h *installHarness) auth.Actor {
				h.expire = true
				return h.officerActor()
			},
			"other session": func(h *installHarness) auth.Actor { return auth.Actor{UserID: h.uid + 1000} },
			"exchange error": func(h *installHarness) auth.Actor {
				h.oauth.InstallErr = ErrExchange
				return h.officerActor()
			},
		}
		for name, setup := range cases {
			t.Run(name, func(t *testing.T) {
				h := newInstallHarness(t)
				actor := setup(h)
				h.wantOutcome(h.callback(h.gid, "code=c", actor), olympusPath, "failed")
				if _, ok := h.installFor(h.gid); ok {
					t.Fatal("a failed install bound the server")
				}
			})
		}
	})
}

func TestChannelsListAndRemovedInDiscord(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	rec := h.do(http.MethodGet, h.path(h.gid, "/channels"), "")
	var got struct {
		Data struct {
			Channels []Channel `json:"channels"`
		} `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil || rec.Code != http.StatusOK {
		t.Fatalf("status %d: %s", rec.Code, rec.Body.String())
	}
	if len(got.Data.Channels) != 2 || got.Data.Channels[0].Name != "raid-signups" || got.Data.Channels[1].Name != "general" {
		t.Fatalf("channels = %+v", got.Data.Channels)
	}

	h.rest.ErrFor = map[string]error{"ServerLayout": &APIError{Status: http.StatusForbidden, Code: codeMissingAccess, Message: "Missing Access"}}
	rec = h.do(http.MethodGet, h.path(h.gid, "/channels"), "")
	if rec.Code != http.StatusBadGateway || !strings.Contains(rec.Body.String(), "discord_unavailable") {
		t.Fatalf("status %d: %s", rec.Code, rec.Body.String())
	}
	if bot := h.settings(h.gid).Bot; bot == nil || !bot.RemovedInDiscord {
		t.Fatalf("bot = %+v", bot)
	}
	if rec := h.do(http.MethodGet, h.path(h.otherID, "/channels"), ""); rec.Code != http.StatusNotFound {
		t.Fatalf("not an officer there: status %d", rec.Code)
	}
	h2 := newInstallHarness(t)
	if rec := h2.do(http.MethodGet, h2.path(h2.gid, "/channels"), ""); rec.Code != http.StatusNotFound {
		t.Fatalf("no install: status %d", rec.Code)
	}
}

func TestSaveChannelRefusesAnUnusableChannel(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	if err := h.svc.d.Store.MarkRemovedInDiscord(context.Background(), h.gid, "gone"); err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"v1", "nope", "locked"} {
		rec := h.do(http.MethodPut, h.path(h.gid, "/channel"), `{"channel_id":"`+id+`"}`)
		if rec.Code != http.StatusUnprocessableEntity || !strings.Contains(rec.Body.String(), "channel_unusable") {
			t.Fatalf("%s: status %d: %s", id, rec.Code, rec.Body.String())
		}
	}
	if rec := h.do(http.MethodPut, h.path(h.gid, "/channel"), `{"channel_id":"c1","extra":1}`); rec.Code != http.StatusBadRequest {
		t.Fatalf("unknown field: status %d", rec.Code)
	}
	rec := h.do(http.MethodPut, h.path(h.gid, "/channel"), `{"channel_id":"c1"}`)
	var got struct {
		Data guilds.SettingsDiscord `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil || rec.Code != http.StatusOK {
		t.Fatalf("status %d: %s", rec.Code, rec.Body.String())
	}
	bot := got.Data.Bot
	if bot == nil || bot.ChannelName == nil || *bot.ChannelName != "raid-signups" || bot.RemovedInDiscord {
		t.Fatalf("bot = %+v", bot)
	}
}

func TestBotTestMessage(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	if rec := h.do(http.MethodPost, h.path(h.gid, "/test"), ""); rec.Code != http.StatusNotFound {
		t.Fatalf("no channel: status %d", rec.Code)
	}
	if err := h.svc.d.Store.SetChannel(context.Background(), h.gid, "c1", "raid-signups"); err != nil {
		t.Fatal(err)
	}
	if err := h.svc.d.Store.RecordPostFailure(context.Background(), h.gid, "boom"); err != nil {
		t.Fatal(err)
	}

	h.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: 403, Code: codeMissingPerms, Message: "Missing Permissions see https://discord.com/x"}}
	rec := h.do(http.MethodPost, h.path(h.gid, "/test"), "")
	if rec.Code != http.StatusBadGateway || !strings.Contains(rec.Body.String(), "Missing Permissions") ||
		strings.Contains(rec.Body.String(), "https://") {
		t.Fatalf("refused test: %d %s", rec.Code, rec.Body.String())
	}
	if in, _ := h.installFor(h.gid); in.Failures != 1 || in.RemovedInDiscord {
		t.Fatalf("a test counted: failures = %d", in.Failures)
	}

	// The server read decides: the bot cannot read the server, so it was removed.
	gone := &APIError{Status: 403, Code: codeMissingAccess, Message: "Missing Access"}
	h.rest.ErrFor = map[string]error{"CreateMessage": gone, "ServerLayout": gone}
	if rec := h.do(http.MethodPost, h.path(h.gid, "/test"), ""); rec.Code != http.StatusBadGateway {
		t.Fatalf("gone bot: status %d", rec.Code)
	}
	if bot := h.settings(h.gid).Bot; bot == nil || !bot.RemovedInDiscord {
		t.Fatalf("a test to a server the bot left did not mark it: %+v", bot)
	}

	h.rest.ErrFor = nil
	if rec := h.do(http.MethodPost, h.path(h.gid, "/test"), ""); rec.Code != http.StatusNoContent {
		t.Fatalf("status %d: %s", rec.Code, rec.Body.String())
	}
	calls := h.rest.Calls()
	last := calls[len(calls)-1]
	if last.Method != "CreateMessage" || last.ChannelID != "c1" || len(last.Message.Embeds) != 1 ||
		last.Message.Embeds[0].Title != "Forever Sixty will post raids for Olympus XXVII here." {
		t.Fatalf("call = %+v", last)
	}
	if in, _ := h.installFor(h.gid); in.Failures != 0 || in.LastError != nil {
		t.Fatalf("success did not reset: %+v", in)
	}
}

func TestRemoveLeavesTheServerWithoutThePlan(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	pool := h.svc.d.Store.Pool
	if _, err := pool.Exec(context.Background(),
		`insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, created_by)
		 values ($1, 'mc', 'Molten Core', now(), 180, $2)`, h.gid, h.uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(context.Background(),
		`insert into guild_raid_discord_posts (raid_id, guild_id, channel_id)
		 select id, guild_id, 'c1' from guild_raids where guild_id = $1`, h.gid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(context.Background(),
		`insert into guild_notification_webhooks (guild_id, token_box, channel_id) values ($1, '\x00', 'w1')`, h.gid); err != nil {
		t.Fatal(err)
	}
	h.setStanding(func(s *notify.GuildStanding) { s.Plan = false })
	if rec := h.do(http.MethodDelete, h.path(h.gid, ""), ""); rec.Code != http.StatusNoContent {
		t.Fatalf("status %d: %s", rec.Code, rec.Body.String())
	}
	calls := h.rest.Calls()
	if len(calls) != 1 || calls[0].Method != "LeaveServer" || calls[0].ServerID != "srv1" {
		t.Fatalf("calls = %+v", calls)
	}
	if _, ok := h.installFor(h.gid); ok {
		t.Fatal("install survived")
	}
	var posts int
	if err := h.svc.d.Store.Pool.QueryRow(context.Background(),
		`select count(*) from guild_raid_discord_posts where guild_id = $1`, h.gid).Scan(&posts); err != nil || posts != 0 {
		t.Fatalf("posts = %d (%v)", posts, err)
	}

	var hooks int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from guild_notification_webhooks where guild_id = $1`, h.gid).Scan(&hooks); err != nil || hooks != 1 {
		t.Fatalf("webhook rows = %d (%v): remove must leave the connected webhook", hooks, err)
	}

	h.setStanding(func(s *notify.GuildStanding) { s.Frozen = true })
	if rec := h.do(http.MethodDelete, h.path(h.gid, ""), ""); rec.Code != http.StatusConflict {
		t.Fatalf("frozen: status %d", rec.Code)
	}
}

func TestActiveNeedsAChannelAndAPresentBot(t *testing.T) {
	h := newInstallHarness(t)
	ctx := context.Background()
	if _, ok, _ := h.svc.ActiveInstall(ctx, h.gid); ok {
		t.Fatal("active with no install")
	}
	h.callback(h.gid, "code=c", h.officerActor())
	if _, ok, _ := h.svc.ActiveInstall(ctx, h.gid); ok {
		t.Fatal("active with no channel")
	}
	if err := h.svc.d.Store.SetChannel(ctx, h.gid, "c1", "raid-signups"); err != nil {
		t.Fatal(err)
	}
	if in, ok, err := h.svc.ActiveInstall(ctx, h.gid); err != nil || !ok || in.ServerID != "srv1" {
		t.Fatalf("active = %+v %v %v", in, ok, err)
	}
	if err := h.svc.d.Store.MarkRemovedInDiscord(ctx, h.gid, "gone"); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := h.svc.ActiveInstall(ctx, h.gid); ok {
		t.Fatal("active after removal in Discord")
	}
}

// TestActivePortCarriesThePostDefault pins guilds.DiscordBot's Active to ActiveInstall's rule
// and the box SetPostDefault remembers on the install.
func TestActivePortCarriesThePostDefault(t *testing.T) {
	h := newInstallHarness(t)
	ctx := context.Background()
	if err := h.svc.SetPostDefault(ctx, h.gid, true); err != nil {
		t.Fatalf("no install: %v", err)
	}
	h.callback(h.gid, "code=c", h.officerActor())
	if active, _, err := h.svc.Active(ctx, h.gid); err != nil || active {
		t.Fatalf("no channel: active %v err %v", active, err)
	}
	if err := h.svc.d.Store.SetChannel(ctx, h.gid, "c1", "raid-signups"); err != nil {
		t.Fatal(err)
	}
	if active, def, err := h.svc.Active(ctx, h.gid); err != nil || !active || !def {
		t.Fatalf("channel saved: active %v default %v err %v", active, def, err)
	}
	if err := h.svc.SetPostDefault(ctx, h.gid, false); err != nil {
		t.Fatal(err)
	}
	if active, def, _ := h.svc.Active(ctx, h.gid); !active || def {
		t.Fatalf("default off: active %v default %v", active, def)
	}
	if err := h.svc.SetPostDefault(ctx, h.gid, true); err != nil {
		t.Fatal(err)
	}
	if active, def, _ := h.svc.Active(ctx, h.gid); !active || !def {
		t.Fatalf("default on: active %v default %v", active, def)
	}
	unusable := ChannelUnusable{Reason: UnusableDeleted}
	if err := h.svc.d.Store.MarkChannelUnusable(ctx, h.gid, unusable, "Unknown Channel"); err != nil {
		t.Fatal(err)
	}
	if active, def, _ := h.svc.Active(ctx, h.gid); active || def {
		t.Fatalf("channel unusable: active %v default %v", active, def)
	}
}

func TestSettingsBotCountsLinkedMembers(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	addMember(t, h.svc.d.Store.Pool, h.gid, h.uid, true)
	if err := h.svc.d.Store.SaveLink(context.Background(), h.uid, "d-1", "zul"); err != nil {
		t.Fatal(err)
	}
	bot := h.settings(h.gid).Bot
	if bot == nil || bot.LinkedMembers != 1 || bot.Members != 1 || bot.ServerName != "Olympus" {
		t.Fatalf("bot = %+v", bot)
	}
}

func TestInstallReturnsThroughTheSharedCallback(t *testing.T) {
	h := newInstallHarness(t)
	mux := http.NewServeMux()
	notify.Mount(mux, h.svc.d.Notify)
	state, err := h.svc.d.Notify.StartState(context.Background(), notify.StateKindBotInstall, h.uid, &h.gid)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/v1/notifications/discord/callback?code=c0de&state="+url.QueryEscape(state), nil)
	req = req.WithContext(auth.WithActor(req.Context(), h.officerActor()))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	h.wantOutcome(rec.Header().Get("Location"), olympusPath, "installed")
	if _, ok := h.installFor(h.gid); !ok {
		t.Fatal("the shared callback did not install")
	}
}

func TestChannelRoutesAreRateLimitedPerAccount(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	var last int
	for i := 0; i <= layoutReadsPerMinute; i++ {
		last = h.do(http.MethodGet, h.path(h.gid, "/channels"), "").Code
	}
	if last != http.StatusTooManyRequests {
		t.Fatalf("call %d: status %d, want 429", layoutReadsPerMinute+1, last)
	}
}
