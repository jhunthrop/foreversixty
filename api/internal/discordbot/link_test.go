package discordbot

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	linkPublic = "https://site.test"
	linkAPI    = "https://api.test"
	linkCode   = "c0de"
)

type linkHarness struct {
	t     *testing.T
	svc   *Service
	oauth *FakeOAuth
	mux   *http.ServeMux
	log   *bytes.Buffer
	uid   int64
}

func newLinkHarness(t *testing.T) *linkHarness {
	t.Helper()
	st := testStore(t)
	h := &linkHarness{
		t: t, oauth: &FakeOAuth{Who: Identity{ID: "d-77", DisplayName: "zulmara"}},
		mux: http.NewServeMux(), log: &bytes.Buffer{}, uid: seedUser(t, st.Pool),
	}
	h.svc = New(Deps{
		Store: st, REST: &FakeREST{}, OAuth: h.oauth, PublicBaseURL: linkPublic, APIBaseURL: linkAPI,
		Notify: &notify.Service{GuildStore: &notify.GuildStore{Pool: st.Pool}}, Log: slog_(h.log),
	})
	Mount(h.mux, h.svc, 0)
	return h
}

// request runs one request through the mux as the given user (0 is signed out).
func (h *linkHarness) request(method, target string, userID int64, device bool) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, target, nil)
	if userID != 0 {
		actor := auth.Actor{UserID: userID}
		if device {
			actor.DeviceID = "dev-1"
		}
		req = req.WithContext(auth.WithActor(req.Context(), actor))
	}
	rec := httptest.NewRecorder()
	h.mux.ServeHTTP(rec, req)
	return rec
}

func (h *linkHarness) newState(userID int64) string {
	h.t.Helper()
	state, err := h.svc.d.Notify.StartState(context.Background(), notify.StateKindLink, userID, nil)
	if err != nil {
		h.t.Fatal(err)
	}
	return state
}

// callback takes the state as the shared callback does and finishes the link as actor.
func (h *linkHarness) callback(state, query string, actor auth.Actor) string {
	h.t.Helper()
	cs, ok, err := h.svc.d.Notify.GuildStore.TakeConnectState(context.Background(), state)
	if err != nil || !ok {
		h.t.Fatalf("take state: ok=%v err=%v", ok, err)
	}
	req := httptest.NewRequest(http.MethodGet, "/v1/notifications/discord/callback?"+query, nil)
	req = req.WithContext(auth.WithActor(req.Context(), actor))
	rec := httptest.NewRecorder()
	h.svc.linkCallback(rec, req, cs)
	return rec.Header().Get("Location")
}

func (h *linkHarness) wantOutcome(loc, outcome string) {
	h.t.Helper()
	if want := linkPublic + "/account?discord_link=" + outcome + "#you"; loc != want {
		h.t.Fatalf("Location = %q, want %q", loc, want)
	}
}

func TestLinkStartSignedOutGoesThroughSignIn(t *testing.T) {
	h := newLinkHarness(t)
	rec := h.request(http.MethodGet, linkStartPath, 0, false)
	want := linkPublic + "/login?next=%2Faccount%3Fdiscord_link%3Dstart%23you"
	if rec.Code != http.StatusFound || rec.Header().Get("Location") != want {
		t.Fatalf("status %d, Location %q, want %q", rec.Code, rec.Header().Get("Location"), want)
	}
}

func TestLinkStartSignedInOpensIdentify(t *testing.T) {
	h := newLinkHarness(t)
	rec := h.request(http.MethodGet, linkStartPath, h.uid, false)
	loc, err := url.Parse(rec.Header().Get("Location"))
	if err != nil || rec.Code != http.StatusFound {
		t.Fatalf("status %d, Location %q", rec.Code, rec.Header().Get("Location"))
	}
	state := loc.Query().Get("state")
	if rec.Header().Get("Location") != h.oauth.LinkURL(state) {
		t.Fatalf("Location = %q, want the identify URL", rec.Header().Get("Location"))
	}
	var kind string
	var userID int64
	var guildID *int64
	err = h.svc.d.Store.Pool.QueryRow(context.Background(),
		`select kind, user_id, guild_id from notify_connect_states where state = $1`, state).Scan(&kind, &userID, &guildID)
	if err != nil || kind != "link" || userID != h.uid || guildID != nil {
		t.Fatalf("state row = %q user %d guild %v (%v)", kind, userID, guildID, err)
	}
}

func TestLinkStartFromADeviceFails(t *testing.T) {
	h := newLinkHarness(t)
	rec := h.request(http.MethodGet, linkStartPath, h.uid, true)
	h.wantOutcome(rec.Header().Get("Location"), linkFailed)
}

func TestLinkCallbackOutcomes(t *testing.T) {
	signedIn := func(h *linkHarness) auth.Actor { return auth.Actor{UserID: h.uid} }
	cases := []struct {
		name    string
		query   string
		outcome string
		setup   func(h *linkHarness, state string) auth.Actor
	}{
		{"linked", "code=" + linkCode, linkLinked, func(h *linkHarness, _ string) auth.Actor { return signedIn(h) }},
		{"denied at discord", "error=access_denied", linkDenied, func(h *linkHarness, _ string) auth.Actor { return signedIn(h) }},
		{"empty code", "code=", linkDenied, func(h *linkHarness, _ string) auth.Actor { return signedIn(h) }},
		{"other session user", "code=" + linkCode, linkFailed, func(h *linkHarness, _ string) auth.Actor {
			return auth.Actor{UserID: seedUser(t, h.svc.d.Store.Pool)}
		}},
		{"signed out", "code=" + linkCode, linkFailed, func(*linkHarness, string) auth.Actor { return auth.Actor{} }},
		{"device", "code=" + linkCode, linkFailed, func(h *linkHarness, _ string) auth.Actor {
			return auth.Actor{UserID: h.uid, DeviceID: "dev-1"}
		}},
		{"expired state", "code=" + linkCode, linkFailed, func(h *linkHarness, state string) auth.Actor {
			_, err := h.svc.d.Store.Pool.Exec(context.Background(),
				`update notify_connect_states set expires_at = now() - interval '1 second' where state = $1`, state)
			if err != nil {
				t.Fatal(err)
			}
			return signedIn(h)
		}},
		{"exchange error", "code=" + linkCode, linkFailed, func(h *linkHarness, _ string) auth.Actor {
			h.oauth.IdentityErr = ErrExchange
			return signedIn(h)
		}},
		{"discord account held by another", "code=" + linkCode, linkConflict, func(h *linkHarness, _ string) auth.Actor {
			other := seedUser(t, h.svc.d.Store.Pool)
			if err := h.svc.d.Store.SaveLink(context.Background(), other, "d-77", "someone"); err != nil {
				t.Fatal(err)
			}
			return signedIn(h)
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := newLinkHarness(t)
			state := h.newState(h.uid)
			actor := tc.setup(h, state)
			h.wantOutcome(h.callback(state, tc.query, actor), tc.outcome)
			link, linked, err := h.svc.d.Store.LinkByUser(context.Background(), h.uid)
			if err != nil {
				t.Fatal(err)
			}
			if wantLinked := tc.outcome == linkLinked; linked != wantLinked {
				t.Fatalf("linked = %v, want %v", linked, wantLinked)
			}
			if linked && (link.DiscordUserID != "d-77" || link.DisplayName != "zulmara") {
				t.Fatalf("link = %+v", link)
			}
		})
	}
}

func TestLinkCallbackNeverLogsTheCode(t *testing.T) {
	h := newLinkHarness(t)
	h.oauth.IdentityErr = errors.New("discordbot: discord code exchange failed: status 400")
	state := h.newState(h.uid)
	loc := h.callback(state, "code="+linkCode+"&state="+state, auth.Actor{UserID: h.uid})
	h.wantOutcome(loc, linkFailed)
	if h.log.Len() == 0 {
		t.Fatal("the exchange failure was not logged at all")
	}
	if strings.Contains(h.log.String(), linkCode) || strings.Contains(loc, linkCode) {
		t.Fatalf("the code leaked: log %q location %q", h.log.String(), loc)
	}
}

func TestUnlinkIsIdempotent(t *testing.T) {
	h := newLinkHarness(t)
	ctx := context.Background()
	if err := h.svc.d.Store.SaveLink(ctx, h.uid, "d-77", "zulmara"); err != nil {
		t.Fatal(err)
	}
	for i := range 2 {
		if rec := h.request(http.MethodDelete, linkPath, h.uid, false); rec.Code != http.StatusNoContent {
			t.Fatalf("delete %d: status %d", i, rec.Code)
		}
	}
	if _, linked, err := h.svc.d.Store.LinkByUser(ctx, h.uid); err != nil || linked {
		t.Fatalf("linked = %v (%v), want unlinked", linked, err)
	}
	if rec := h.request(http.MethodDelete, linkPath, 0, false); rec.Code != http.StatusUnauthorized {
		t.Fatalf("signed out delete: status %d, want 401", rec.Code)
	}
}

func TestLinkViewForTheAccountRead(t *testing.T) {
	h := newLinkHarness(t)
	ctx := context.Background()
	if !h.svc.LinkAvailable() {
		t.Fatal("link reads as unavailable with oauth and notify wired")
	}
	g := seedGuild(t, h.svc.d.Store.Pool, "Olympus XXVII")
	addMember(t, h.svc.d.Store.Pool, g, h.uid, true)
	if _, err := h.svc.d.Store.BindServer(ctx, g, "s-1", "Srv", h.uid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.d.Store.Pool.Exec(ctx, `update guild_discord_installs set channel_id = 'c1' where guild_id = $1`, g); err != nil {
		t.Fatal(err)
	}
	link, suggested, err := h.svc.DiscordLink(ctx, h.uid)
	if err != nil || link != nil || suggested == nil || *suggested != "Olympus XXVII" {
		t.Fatalf("unlinked = %+v %v (%v)", link, suggested, err)
	}
	if err := h.svc.d.Store.SaveLink(ctx, h.uid, "d-77", "zulmara"); err != nil {
		t.Fatal(err)
	}
	link, suggested, err = h.svc.DiscordLink(ctx, h.uid)
	if err != nil || link == nil || link.DisplayName != "zulmara" || suggested != nil {
		t.Fatalf("linked = %+v %v (%v)", link, suggested, err)
	}
}

func TestLinkCommand(t *testing.T) {
	h := newLinkHarness(t)
	ctx := context.Background()
	in := Interaction{DiscordUserID: "d-77"}
	const heading = "### 6.2 In Discord"

	reply, err := h.svc.cmdLink(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	h.wantLinkReply(reply, specString(t, heading, "`/link`", 0, nil), specString(t, heading, "`/link`", 1, nil), linkAPI+"/v1/discord/link/start")

	if err := h.svc.d.Store.SaveLink(ctx, h.uid, "d-77", "zulmara"); err != nil {
		t.Fatal(err)
	}
	reply, err = h.svc.cmdLink(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	h.wantLinkReply(reply, specString(t, heading, "`/link`", 2, nil), specString(t, heading, "`/link`", 3, nil), linkPublic+"/account#you")
}

func (h *linkHarness) wantLinkReply(reply Reply, content, label, linkURL string) {
	h.t.Helper()
	m := reply.Message
	if !m.Ephemeral || m.Content != content || len(m.Components) != 1 || len(m.Components[0]) != 1 {
		h.t.Fatalf("reply = %+v, want ephemeral %q with one button", m, content)
	}
	b := m.Components[0][0]
	if b.Style != StyleLink || b.Label != label || b.URL != linkURL {
		h.t.Fatalf("button = %+v, want link %q to %q", b, label, linkURL)
	}
}

func TestLinkCommandNeverShowsAnotherAccountsData(t *testing.T) {
	h := newLinkHarness(t)
	other := seedUser(t, h.svc.d.Store.Pool)
	if err := h.svc.d.Store.SaveLink(context.Background(), other, "d-other", "otherperson"); err != nil {
		t.Fatal(err)
	}
	reply, err := h.svc.cmdLink(context.Background(), Interaction{DiscordUserID: "d-77"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(reply.Message.Content, "otherperson") || strings.Contains(reply.Message.Content, strconv.FormatInt(other, 10)) {
		t.Fatalf("reply mentions another account: %q", reply.Message.Content)
	}
}
