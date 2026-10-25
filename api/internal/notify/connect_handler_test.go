package notify

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

const (
	connectedHook  = "https://discord.com/api/webhooks/111222333444555666/NEWtoken-xyz_1"
	connectedToken = "NEWtoken-xyz_1"
	connectedGuild = "777888999000111222"
	authCode       = "SECRETauthcode42"
	accountBase    = "https://foreversixty.gg/account?discord="
)

func connectFake() *FakeOAuthClient {
	return &FakeOAuthClient{Webhook: OAuthWebhook{
		ID: "111222333444555666", Token: connectedToken, URL: connectedHook,
		ChannelID: testChan, GuildID: connectedGuild, Name: "Forever Sixty #raid",
	}}
}

// get issues a GET that keeps redirects unfollowed, with the session and the given cookies.
func (h *apiHarness) get(path string, cookies ...*http.Cookie) (*http.Response, string) {
	h.t.Helper()
	req, err := http.NewRequest("GET", h.server.URL+path, nil)
	if err != nil {
		h.t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: auth.SessionCookie, Value: h.session})
	for _, c := range cookies {
		req.AddCookie(c)
	}
	client := &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	res, err := client.Do(req)
	if err != nil {
		h.t.Fatal(err)
	}
	res.Body.Close()
	return res, res.Header.Get("Location")
}

// getAnonymous issues a GET with no cookies at all.
func (h *apiHarness) getAnonymous(path string) (*http.Response, string) {
	h.t.Helper()
	client := &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	res, err := client.Get(h.server.URL + path)
	if err != nil {
		h.t.Fatal(err)
	}
	res.Body.Close()
	return res, res.Header.Get("Location")
}

// startConnect runs the connect leg and returns the state cookie and the state.
func (h *apiHarness) startConnect() (*http.Cookie, string) {
	h.t.Helper()
	res, loc := h.get("/v1/notifications/discord/connect")
	if res.StatusCode != http.StatusFound {
		h.t.Fatalf("connect status %d, Location %q", res.StatusCode, loc)
	}
	u, err := url.Parse(loc)
	if err != nil {
		h.t.Fatal(err)
	}
	for _, c := range res.Cookies() {
		if c.Name == discordStateCookie {
			return c, u.Query().Get("state")
		}
	}
	h.t.Fatal("no state cookie")
	return nil, ""
}

func (h *apiHarness) callback(query string, cookie *http.Cookie) (*http.Response, string) {
	h.t.Helper()
	if cookie == nil {
		return h.get("/v1/notifications/discord/callback?" + query)
	}
	return h.get("/v1/notifications/discord/callback?"+query, cookie)
}

func (h *apiHarness) assertOutcome(loc, want string) {
	h.t.Helper()
	if loc != accountBase+want+"#discord" {
		h.t.Fatalf("Location = %q, want outcome %q", loc, want)
	}
}

func (h *apiHarness) stored() Webhook {
	h.t.Helper()
	w, ok, err := h.svc.Store.Get(context.Background(), h.uid, KindDiscord)
	if err != nil || !ok {
		h.t.Fatalf("stored webhook: ok=%v err=%v", ok, err)
	}
	return w
}

func (h *apiHarness) assertNothingStored() {
	h.t.Helper()
	if _, ok, err := h.svc.Store.Get(context.Background(), h.uid, KindDiscord); err != nil || ok {
		h.t.Fatalf("a webhook was stored (ok=%v err=%v)", ok, err)
	}
}

func TestConnectRedirectsToDiscordWithAStateCookie(t *testing.T) {
	h := newAPIWithOAuth(t, connectFake())
	h.premium()
	cookie, state := h.startConnect()
	if state == "" || cookie.Value != state {
		t.Fatalf("state %q, cookie %q", state, cookie.Value)
	}
	if !cookie.HttpOnly || cookie.SameSite != http.SameSiteLaxMode || cookie.Path != "/v1/notifications/discord/callback" || cookie.MaxAge != 600 {
		t.Fatalf("cookie = %+v", cookie)
	}
}

func TestConnectIs404WhenUnconfigured(t *testing.T) {
	h := newAPI(t)
	h.premium()
	res, _ := h.get("/v1/notifications/discord/connect")
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("status %d, want 404", res.StatusCode)
	}
	if h.svc.ConnectEnabled() {
		t.Fatal("ConnectEnabled with no OAuth client")
	}
}

func TestConnectRedirectsToFailedWithoutEntitlement(t *testing.T) {
	h := newAPIWithOAuth(t, connectFake())
	res, loc := h.get("/v1/notifications/discord/connect")
	if res.StatusCode != http.StatusFound {
		t.Fatalf("status %d", res.StatusCode)
	}
	h.assertOutcome(loc, "failed")
	for _, c := range res.Cookies() {
		if c.Name == discordStateCookie {
			t.Fatal("a state cookie was set for an unentitled member")
		}
	}
}

func TestCallbackStoresAConnectedWebhookAndRedirectsConnected(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	var logs bytes.Buffer
	h.svc.Log = slog.New(slog.NewTextHandler(&logs, nil))
	cookie, state := h.startConnect()
	res, loc := h.callback("code="+authCode+"&state="+url.QueryEscape(state), cookie)
	if res.StatusCode != http.StatusFound {
		t.Fatalf("status %d", res.StatusCode)
	}
	h.assertOutcome(loc, "connected")
	for _, secret := range []string{authCode, connectedToken, "webhooks"} {
		if strings.Contains(loc, secret) || strings.Contains(logs.String(), secret) {
			t.Fatalf("%q leaked into the Location or the logs", secret)
		}
	}
	w := h.stored()
	if w.Source != SourceConnected || w.Name != nil || w.ChannelID != testChan || w.GuildID == nil || *w.GuildID != connectedGuild {
		t.Fatalf("stored = %+v", w)
	}
	if !w.Events[EventSimFinished] || !w.Events[EventUploadParsed] || w.Events[EventUploadPrivate] {
		t.Fatalf("default events = %v", w.Events)
	}
	box, err := h.svc.Store.TokenBox(context.Background(), h.uid, KindDiscord)
	if err != nil {
		t.Fatal(err)
	}
	if opened, err := h.svc.Sealer.Open(box); err != nil || opened != connectedHook {
		t.Fatalf("sealed url opens to %q (%v)", opened, err)
	}
	if codes := fake.Codes(); len(codes) != 1 || codes[0] != authCode {
		t.Fatalf("exchanged codes = %v", codes)
	}
	cleared := false
	for _, c := range res.Cookies() {
		cleared = cleared || (c.Name == discordStateCookie && c.MaxAge < 0)
	}
	if !cleared {
		t.Fatal("the state cookie was not cleared")
	}
}

func TestCallbackWithoutCodeIsDenied(t *testing.T) {
	h := newAPIWithOAuth(t, connectFake())
	h.premium()
	cookie, state := h.startConnect()
	_, loc := h.callback("state="+url.QueryEscape(state), cookie)
	h.assertOutcome(loc, "denied")
	h.assertNothingStored()
}

func TestCallbackAccessDeniedIsDenied(t *testing.T) {
	h := newAPIWithOAuth(t, connectFake())
	h.premium()
	cookie, state := h.startConnect()
	_, loc := h.callback("error=access_denied&state="+url.QueryEscape(state), cookie)
	h.assertOutcome(loc, "denied")
	h.assertNothingStored()
}

func TestCallbackWithWrongStateIsFailedAndStoresNothing(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	cookie, _ := h.startConnect()
	_, loc := h.callback("code="+authCode+"&state=forged", cookie)
	h.assertOutcome(loc, "failed")
	_, loc = h.callback("code="+authCode+"&state="+url.QueryEscape(cookie.Value), nil)
	h.assertOutcome(loc, "failed")
	h.assertNothingStored()
	if len(fake.Codes()) != 0 {
		t.Fatal("a code was exchanged without a valid state")
	}
}

func TestCallbackWithoutTheStateCookieIsFailed(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	cookie, state := h.startConnect()
	query := "code=" + authCode + "&state=" + url.QueryEscape(state)
	_, loc := h.callback(query, cookie)
	h.assertOutcome(loc, "connected")
	// The callback cleared the cookie, so a replay presents none. The state is
	// single-use by the cookie alone; there is no server-side marker (the
	// cookie is the only memory, and a second API instance would not share one).
	_, loc = h.callback(query, nil)
	h.assertOutcome(loc, "failed")
	if len(fake.Codes()) != 1 {
		t.Fatalf("codes exchanged = %d, want 1", len(fake.Codes()))
	}
}

func TestCallbackWithAMismatchedStateCookieIsFailed(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	_, state := h.startConnect()
	other := &http.Cookie{Name: discordStateCookie, Value: "another-browsers-state"}
	_, loc := h.callback("code="+authCode+"&state="+url.QueryEscape(state), other)
	h.assertOutcome(loc, "failed")
	h.assertNothingStored()
	if len(fake.Codes()) != 0 {
		t.Fatal("a code was exchanged against a mismatched cookie")
	}
}

func TestConnectAndCallbackWithoutASessionRedirectToFailed(t *testing.T) {
	h := newAPIWithOAuth(t, connectFake())
	for _, path := range []string{"/v1/notifications/discord/connect", "/v1/notifications/discord/callback?code=c&state=s"} {
		res, loc := h.getAnonymous(path)
		if res.StatusCode != http.StatusFound {
			t.Fatalf("%s: status %d, want 302", path, res.StatusCode)
		}
		h.assertOutcome(loc, "failed")
	}
}

func TestCallbackLapsedSinceConnectIsFailedWithoutAnExchange(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	cookie, state := h.startConnect()
	h.lapse()
	_, loc := h.callback("code="+authCode+"&state="+url.QueryEscape(state), cookie)
	h.assertOutcome(loc, "failed")
	h.assertNothingStored()
	if len(fake.Codes()) != 0 {
		t.Fatal("a lapsed member's code was exchanged")
	}
}

func TestCallbackOrphanedWebhookIsDeletedFromDiscord(t *testing.T) {
	cases := map[string]OAuthWebhook{
		"empty channel id": {ID: "111222333444555666", Token: connectedToken, URL: connectedHook, GuildID: connectedGuild},
		"save fails":       {ID: "111222333444555666", Token: connectedToken, URL: connectedHook, ChannelID: "bad\x00id"},
	}
	for name, hook := range cases {
		t.Run(name, func(t *testing.T) {
			h := newAPIWithOAuth(t, &FakeOAuthClient{Webhook: hook})
			h.premium()
			var logs bytes.Buffer
			h.svc.Log = slog.New(slog.NewTextHandler(&logs, nil))
			cookie, state := h.startConnect()
			_, loc := h.callback("code="+authCode+"&state="+url.QueryEscape(state), cookie)
			h.assertOutcome(loc, "failed")
			h.assertNothingStored()
			if got := h.poster.Deletes(); len(got) != 1 || got[0] != connectedHook {
				t.Fatalf("deletes = %v, want the orphan", got)
			}
			assertNoSecrets(t, logs.String())
		})
	}
}

func TestCallbackOrphanWithAnUnparseableURLIsDeletedFromItsIDAndToken(t *testing.T) {
	hook := OAuthWebhook{ID: "111222333444555666", Token: connectedToken, URL: "https://evil.example/x", ChannelID: testChan}
	h := newAPIWithOAuth(t, &FakeOAuthClient{Webhook: hook})
	h.premium()
	var logs bytes.Buffer
	h.svc.Log = slog.New(slog.NewTextHandler(&logs, nil))
	cookie, state := h.startConnect()
	_, loc := h.callback("code="+authCode+"&state="+url.QueryEscape(state), cookie)
	h.assertOutcome(loc, "failed")
	if got := h.poster.Deletes(); len(got) != 1 || got[0] != connectedHook {
		t.Fatalf("deletes = %v", got)
	}
	assertNoSecrets(t, logs.String())
}

// assertNoSecrets fails when the logs carry the code, the token, the URL, or any
// given sealed box (raw or base64, standard or URL-safe).
func assertNoSecrets(t *testing.T, logs string, boxes ...[]byte) {
	t.Helper()
	secrets := []string{authCode, connectedToken, connectedHook, "webhooks/"}
	for _, box := range boxes {
		if len(box) == 0 {
			continue
		}
		secrets = append(secrets, string(box), base64.StdEncoding.EncodeToString(box), base64.URLEncoding.EncodeToString(box), base64.RawURLEncoding.EncodeToString(box))
	}
	for _, secret := range secrets {
		if strings.Contains(logs, secret) {
			t.Fatalf("%q leaked into the logs: %s", secret, logs)
		}
	}
}

func TestCallbackExchangeFailureIsFailed(t *testing.T) {
	cases := map[string]OAuthWebhook{
		"error":         {},
		"empty url":     {ChannelID: testChan},
		"foreign url":   {URL: "https://evil.example/api/webhooks/1/abc", ChannelID: testChan},
		"no channel id": {URL: connectedHook},
	}
	for name, hook := range cases {
		t.Run(name, func(t *testing.T) {
			fake := &FakeOAuthClient{Webhook: hook}
			if name == "error" {
				fake.Err = errors.New("exchange failed")
			}
			h := newAPIWithOAuth(t, fake)
			h.premium()
			var logs bytes.Buffer
			h.svc.Log = slog.New(slog.NewTextHandler(&logs, nil))
			cookie, state := h.startConnect()
			_, loc := h.callback("code="+authCode+"&state="+url.QueryEscape(state), cookie)
			h.assertOutcome(loc, "failed")
			h.assertNothingStored()
			assertNoSecrets(t, logs.String())
		})
	}
}

func (h *apiHarness) connectWith(fake *FakeOAuthClient) {
	h.t.Helper()
	cookie, state := h.startConnect()
	_, loc := h.callback("code="+authCode+"&state="+url.QueryEscape(state), cookie)
	h.assertOutcome(loc, "connected")
}

func TestSuccessfulConnectLogsNeitherSecretsNorTheSealedBox(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	var logs bytes.Buffer
	h.svc.Log = slog.New(slog.NewTextHandler(&logs, nil))
	h.connectWith(fake)
	box, err := h.svc.Store.TokenBox(context.Background(), h.uid, KindDiscord)
	if err != nil || len(box) == 0 {
		t.Fatalf("sealed box = %d bytes, err %v", len(box), err)
	}
	assertNoSecrets(t, logs.String(), box)
}

func TestReconnectDeletesTheOldWebhookFromDiscordAfterSaving(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	h.connectWith(fake)
	// Turn one event off: a reconnect keeps the member's events.
	events := map[EventName]bool{EventSimFinished: false, EventUploadParsed: true, EventUploadPrivate: true}
	if err := h.svc.Store.SaveEvents(context.Background(), h.uid, KindDiscord, events); err != nil {
		t.Fatal(err)
	}
	if len(h.poster.Deletes()) != 0 {
		t.Fatal("the first connect deleted something")
	}
	const second = "https://discord.com/api/webhooks/999888777666555444/SECONDtoken_2"
	fake.Webhook.URL = second
	h.connectWith(fake)
	if got := h.poster.Deletes(); len(got) != 1 || got[0] != connectedHook {
		t.Fatalf("deletes = %v, want the first webhook", got)
	}
	box, _ := h.svc.Store.TokenBox(context.Background(), h.uid, KindDiscord)
	if opened, _ := h.svc.Sealer.Open(box); opened != second {
		t.Fatalf("stored url = %q", opened)
	}
	if w := h.stored(); w.Events[EventSimFinished] || !w.Events[EventUploadPrivate] {
		t.Fatalf("events were not kept: %v", w.Events)
	}
}

func TestReconnectOverAPastedWebhookDeletesNothingInDiscord(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.saveHook()
	h.connectWith(fake)
	if got := h.poster.Deletes(); len(got) != 0 {
		t.Fatalf("a pasted webhook was deleted from Discord: %v", got)
	}
}

func TestPastingOverAConnectedWebhookDeletesItFromDiscord(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	h.connectWith(fake)
	if status, raw, _ := h.call("PUT", "/v1/notifications/discord", putBody(testHook, true, true)); status != 200 {
		t.Fatalf("status %d: %s", status, raw)
	}
	if got := h.poster.Deletes(); len(got) != 1 || got[0] != connectedHook {
		t.Fatalf("deletes = %v", got)
	}
	if w := h.stored(); w.Source != SourcePasted {
		t.Fatalf("source = %q", w.Source)
	}
}

func TestRemoveDeletesAConnectedWebhookFromDiscordButNotAPastedOne(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	h.connectWith(fake)
	h.poster.DeleteErr = errors.New("discord is down") // best effort: the row still goes
	if status, _, _ := h.call("DELETE", "/v1/notifications/discord", ""); status != 204 {
		t.Fatalf("status %d", status)
	}
	if got := h.poster.Deletes(); len(got) != 1 || got[0] != connectedHook {
		t.Fatalf("deletes = %v", got)
	}
	h.assertNothingStored()

	h.saveHook()
	if status, _, _ := h.call("DELETE", "/v1/notifications/discord", ""); status != 204 {
		t.Fatalf("status %d", status)
	}
	if got := h.poster.Deletes(); len(got) != 1 {
		t.Fatalf("a pasted webhook was deleted from Discord: %v", got)
	}
}

func TestRemoveDeletesTheRowBeforeCallingDiscord(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	h.connectWith(fake)
	rowStoredAtDelete := true
	h.poster.OnDelete = func(string) {
		_, rowStoredAtDelete, _ = h.svc.Store.Get(context.Background(), h.uid, KindDiscord)
	}
	if status, _, _ := h.call("DELETE", "/v1/notifications/discord", ""); status != 204 {
		t.Fatalf("status %d", status)
	}
	if len(h.poster.Deletes()) != 1 {
		t.Fatalf("deletes = %v", h.poster.Deletes())
	}
	if rowStoredAtDelete {
		t.Fatal("Discord was called while the row was still stored")
	}
}

func TestLapsedRemoveDeletesNothingInDiscord(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	h.connectWith(fake)
	h.lapse()
	if status, _, _ := h.call("DELETE", "/v1/notifications/discord", ""); status != 204 {
		t.Fatalf("status %d", status)
	}
	if got := h.poster.Deletes(); len(got) != 0 {
		t.Fatalf("a lapsed remove deleted from Discord: %v", got)
	}
	h.assertNothingStored()
}

func TestMeCarriesConnectSourceAndGuildID(t *testing.T) {
	fake := connectFake()
	h := newAPIWithOAuth(t, fake)
	h.premium()
	data, _ := h.me()
	if string(data["notifications"]) != `{"connect":true,"discord":null}` {
		t.Fatalf("notifications = %s", data["notifications"])
	}
	h.connectWith(fake)
	data, raw := h.me()
	for _, want := range []string{`"connect":true`, `"source":"connected"`, `"guild_id":"` + connectedGuild + `"`, `"name":null`} {
		if !strings.Contains(string(data["notifications"]), want) {
			t.Fatalf("notifications = %s, missing %s", data["notifications"], want)
		}
	}
	if strings.Contains(raw, connectedToken) {
		t.Fatal("the token is in /v1/me")
	}
}

func TestPastedWebhookStoresSourceAndLookupGuildID(t *testing.T) {
	h := newAPI(t)
	h.poster.LookupGuildID = connectedGuild
	h.saveHook()
	w := h.stored()
	if w.Source != SourcePasted || w.GuildID == nil || *w.GuildID != connectedGuild {
		t.Fatalf("stored = %+v", w)
	}
}

func TestMigration0048ReadsAnOldRowAsPasted(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	if _, err := pool.Exec(context.Background(), `insert into notification_webhooks (user_id, kind, channel_id, token_box, events)
		values ($1, $2, '9', 'b', '{"sim_finished": true}')`, uid, KindDiscord); err != nil {
		t.Fatal(err)
	}
	w, ok, err := (&Store{Pool: pool}).Get(context.Background(), uid, KindDiscord)
	if err != nil || !ok || w.Source != SourcePasted || w.GuildID != nil {
		t.Fatalf("old row = %+v ok=%v err=%v", w, ok, err)
	}
	if _, err := pool.Exec(context.Background(), `update notification_webhooks set source = 'other' where user_id = $1`, uid); err == nil {
		t.Fatal("the source check accepted an unknown value")
	}
}
