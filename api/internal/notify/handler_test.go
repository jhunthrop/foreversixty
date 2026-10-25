package notify

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

const (
	testHook  = "https://discord.com/api/webhooks/123456789012345678/SECRETtoken-abc_DEF"
	testToken = "SECRETtoken-abc_DEF"
	testID    = "123456789012345678"
	testChan  = "998877665544332211"
	invalidMS = "That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/."
)

type apiHarness struct {
	svc     *Service
	poster  *FakePoster
	server  *httptest.Server
	uid     int64
	session string
	t       *testing.T
}

func newAPI(t *testing.T) *apiHarness {
	t.Helper()
	pool := testPool(t)
	sealer, err := NewSealer(make([]byte, 32))
	if err != nil {
		t.Fatal(err)
	}
	poster := &FakePoster{LookupChannelID: testChan}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	ents := &entitlements.Store{Pool: pool}
	svc := &Service{
		Store: &Store{Pool: pool}, Sealer: sealer, Entitlements: ents,
		Poster: poster, PublicBaseURL: "https://foreversixty.gg", Log: quiet,
	}
	authStore := &auth.Store{Pool: pool}
	authn := &auth.Authenticator{Store: authStore, Log: quiet}
	accounts := &auth.Service{
		Store: authStore, Auth: authn, Entitlements: ents,
		Notifications: svc, PublicBaseURL: "https://foreversixty.gg", Log: quiet,
	}
	mux := http.NewServeMux()
	auth.Mount(mux, accounts, 0)
	Mount(mux, svc)
	srv := httptest.NewServer(authn.Middleware(mux))
	t.Cleanup(srv.Close)

	uid := seedUser(t, pool)
	sess, err := authStore.CreateSession(context.Background(), uid, "email", auth.SessionTTL)
	if err != nil {
		t.Fatal(err)
	}
	return &apiHarness{svc: svc, poster: poster, server: srv, uid: uid, session: sess.ID, t: t}
}

func (h *apiHarness) premium() {
	h.t.Helper()
	pool := h.svc.Store.Pool
	if _, err := pool.Exec(context.Background(),
		`insert into entitlements (user_id, plan, source, status) values ($1, 'premium', 'grant', 'active')
		 on conflict do nothing`, h.uid); err != nil {
		h.t.Fatal(err)
	}
}

func (h *apiHarness) lapse() {
	h.t.Helper()
	if _, err := h.svc.Store.Pool.Exec(context.Background(),
		`delete from entitlements where user_id = $1`, h.uid); err != nil {
		h.t.Fatal(err)
	}
}

func (h *apiHarness) call(method, path, body string) (int, string, http.Header) {
	h.t.Helper()
	req, err := http.NewRequest(method, h.server.URL+path, strings.NewReader(body))
	if err != nil {
		h.t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: auth.SessionCookie, Value: h.session})
	req.AddCookie(&http.Cookie{Name: auth.CSRFCookie, Value: "csrf"})
	req.Header.Set(auth.CSRFHeader, "csrf")
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		h.t.Fatal(err)
	}
	defer res.Body.Close()
	raw, _ := io.ReadAll(res.Body)
	return res.StatusCode, string(raw), res.Header
}

func putBody(url string, sim, upload bool) string {
	b := map[string]any{"events": map[string]bool{"sim_finished": sim, "upload_parsed": upload}}
	if url != "" {
		b["webhook_url"] = url
	}
	raw, _ := json.Marshal(b)
	return string(raw)
}

type blockEnv struct {
	Data auth.DiscordBlock `json:"data"`
}

type errEnv struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (h *apiHarness) me() (map[string]json.RawMessage, string) {
	h.t.Helper()
	status, raw, _ := h.call("GET", "/v1/me", "")
	if status != 200 {
		h.t.Fatalf("me status %d: %s", status, raw)
	}
	var env struct {
		Data map[string]json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(raw), &env); err != nil {
		h.t.Fatal(err)
	}
	return env.Data, raw
}

func (h *apiHarness) saveHook() {
	h.t.Helper()
	h.premium()
	if status, raw, _ := h.call("PUT", "/v1/notifications/discord", putBody(testHook, true, true)); status != 200 {
		h.t.Fatalf("save status %d: %s", status, raw)
	}
}

func TestSaveThenMeShowsTheBlockAndNoCredential(t *testing.T) {
	h := newAPI(t)
	h.saveHook()
	status, raw, hdr := h.call("PUT", "/v1/notifications/discord", putBody(testHook, true, false))
	if status != 200 {
		t.Fatalf("status %d: %s", status, raw)
	}
	if got := hdr.Get("Cache-Control"); got != "private, no-store" {
		t.Fatalf("Cache-Control = %q", got)
	}
	var env blockEnv
	if err := json.Unmarshal([]byte(raw), &env); err != nil {
		t.Fatal(err)
	}
	b := env.Data
	if b.ChannelID != testChan || !b.Events.SimFinished || b.Events.UploadParsed || b.Failures != 0 || b.DisabledReason != nil || b.LastError != nil {
		t.Fatalf("block = %+v", b)
	}
	data, meRaw := h.me()
	var notes struct {
		Discord auth.DiscordBlock `json:"discord"`
	}
	if err := json.Unmarshal(data["notifications"], &notes); err != nil || notes.Discord.ChannelID != testChan {
		t.Fatalf("me notifications = %s (%v)", data["notifications"], err)
	}
	for _, body := range []string{raw, meRaw} {
		if strings.Contains(body, testToken) || strings.Contains(body, testID) || strings.Contains(body, "webhooks") {
			t.Fatalf("a body leaks the credential: %s", body)
		}
	}
}

func TestMeReadsNullBeforeAnythingIsSavedAndAfterDelete(t *testing.T) {
	h := newAPI(t)
	data, _ := h.me()
	if string(data["notifications"]) != `{"discord":null}` {
		t.Fatalf("notifications = %s", data["notifications"])
	}
	h.saveHook()
	if status, _, _ := h.call("DELETE", "/v1/notifications/discord", ""); status != 204 {
		t.Fatalf("delete status %d", status)
	}
	if status, _, _ := h.call("DELETE", "/v1/notifications/discord", ""); status != 204 {
		t.Fatalf("second delete status %d, want idempotent 204", status)
	}
	data, _ = h.me()
	if string(data["notifications"]) != `{"discord":null}` {
		t.Fatalf("notifications after delete = %s", data["notifications"])
	}
}

func TestBadURLsAreRefusedWithoutEchoingTheInput(t *testing.T) {
	h := newAPI(t)
	h.premium()
	for _, bad := range []string{
		"https://evil.example/api/webhooks/1/abcSECRET",
		"http://discord.com/api/webhooks/1/abc",
		testHook + "?wait=true",
		"not a url SECRETINPUT",
	} {
		status, raw, _ := h.call("PUT", "/v1/notifications/discord", putBody(bad, true, true))
		var e errEnv
		_ = json.Unmarshal([]byte(raw), &e)
		if status != 400 || e.Error.Code != "invalid" || e.Error.Message != invalidMS {
			t.Fatalf("%q -> %d %s", bad, status, raw)
		}
		if strings.Contains(raw, "SECRET") || strings.Contains(raw, "evil") {
			t.Fatalf("response echoes input: %s", raw)
		}
	}
}

func TestLookupFailureIs422AndStoresNothing(t *testing.T) {
	h := newAPI(t)
	h.premium()
	h.poster.LookupErr = &DiscordError{Status: 404, Body: "Unknown Webhook"}
	status, raw, _ := h.call("PUT", "/v1/notifications/discord", putBody(testHook, true, true))
	var e errEnv
	_ = json.Unmarshal([]byte(raw), &e)
	if status != 422 || e.Error.Code != "invalid" || e.Error.Message != invalidMS {
		t.Fatalf("%d %s", status, raw)
	}
	assertNoCredential(t, raw)
	if _, ok, err := h.svc.Store.Get(context.Background(), h.uid, KindDiscord); err != nil || ok {
		t.Fatalf("a row was stored (ok=%v err=%v)", ok, err)
	}
}

func TestFreeMembersGet402OnPutAndTestButMayDelete(t *testing.T) {
	h := newAPI(t)
	for _, c := range []struct{ method, path, body string }{
		{"PUT", "/v1/notifications/discord", putBody(testHook, true, true)},
		{"POST", "/v1/notifications/discord/test", ""},
	} {
		status, raw, _ := h.call(c.method, c.path, c.body)
		var e errEnv
		_ = json.Unmarshal([]byte(raw), &e)
		if status != 402 || e.Error.Code != "premium_required" {
			t.Fatalf("%s: %d %s", c.method, status, raw)
		}
	}
	if status, _, _ := h.call("DELETE", "/v1/notifications/discord", ""); status != 204 {
		t.Fatalf("delete status %d", status)
	}
}

func TestEventsOnlyPut(t *testing.T) {
	h := newAPI(t)
	h.premium()
	if status, _, _ := h.call("PUT", "/v1/notifications/discord", putBody("", true, true)); status != 404 {
		t.Fatalf("no row: status %d, want 404", status)
	}
	h.saveHook()
	status, raw, _ := h.call("PUT", "/v1/notifications/discord", putBody("", false, true))
	var env blockEnv
	_ = json.Unmarshal([]byte(raw), &env)
	if status != 200 || env.Data.Events.SimFinished || !env.Data.Events.UploadParsed {
		t.Fatalf("%d %s", status, raw)
	}
}

func TestBlockReasons(t *testing.T) {
	h := newAPI(t)
	h.saveHook()
	pool := h.svc.Store.Pool
	if _, err := pool.Exec(context.Background(),
		`update notification_webhooks set failures = 5, last_error = 'Unknown Webhook', disabled_at = now() where user_id = $1`, h.uid); err != nil {
		t.Fatal(err)
	}
	reason := func() string {
		_, raw, _ := h.call("PUT", "/v1/notifications/discord", putBody("", true, true))
		var env blockEnv
		_ = json.Unmarshal([]byte(raw), &env)
		if env.Data.DisabledReason == nil {
			return ""
		}
		return *env.Data.DisabledReason
	}
	h.lapse()
	// A lapsed member cannot PUT, so read the block through /v1/me.
	meReason := func() string {
		data, _ := h.me()
		var n struct {
			Discord auth.DiscordBlock `json:"discord"`
		}
		_ = json.Unmarshal(data["notifications"], &n)
		if n.Discord.DisabledReason == nil {
			return ""
		}
		return *n.Discord.DisabledReason
	}
	if got := meReason(); got != "failed five times" {
		t.Fatalf("disabled and lapsed: %q", got)
	}
	if _, err := pool.Exec(context.Background(),
		`update notification_webhooks set failures = 0, disabled_at = null where user_id = $1`, h.uid); err != nil {
		t.Fatal(err)
	}
	if got := meReason(); got != "plan lapsed" {
		t.Fatalf("lapsed: %q", got)
	}
	h.premium()
	if got := reason(); got != "" {
		t.Fatalf("healthy: %q", got)
	}
}

func TestTestRoutePostsOnceAndResetsTheCounter(t *testing.T) {
	h := newAPI(t)
	h.saveHook()
	pool := h.svc.Store.Pool
	if _, err := pool.Exec(context.Background(),
		`update notification_webhooks set failures = 5, last_error = 'x', disabled_at = now() where user_id = $1`, h.uid); err != nil {
		t.Fatal(err)
	}
	if status, raw, _ := h.call("POST", "/v1/notifications/discord/test", ""); status != 204 {
		t.Fatalf("status %d: %s", status, raw)
	}
	posts := h.poster.Posts()
	if len(posts) != 1 || posts[0].URL != testHook || posts[0].Embed.Title != "Test message from Forever Sixty" {
		t.Fatalf("posts = %+v", posts)
	}
	wh, _, _ := h.svc.Store.Get(context.Background(), h.uid, KindDiscord)
	if wh.Failures != 0 || wh.DisabledAt != nil {
		t.Fatalf("after success: %+v", wh)
	}
}

func assertNoCredential(t *testing.T, raw string) {
	t.Helper()
	// The fixed invalid-URL sentence legitimately names the path prefix.
	raw = strings.ReplaceAll(raw, invalidMS, "")
	for _, leak := range []string{testToken, testID, "webhooks"} {
		if strings.Contains(raw, leak) {
			t.Fatalf("body contains %q: %s", leak, raw)
		}
	}
}

func TestFailedTestIs502LeavesCountersAndTruncates(t *testing.T) {
	h := newAPI(t)
	h.saveHook()
	for range 2 {
		if _, err := h.svc.Store.RecordFailure(context.Background(), h.uid, KindDiscord, "Unknown Webhook", true); err != nil {
			t.Fatal(err)
		}
	}
	h.poster.PostErr = &DiscordError{Status: 401, Body: strings.Repeat("é", 200)}
	status, raw, _ := h.call("POST", "/v1/notifications/discord/test", "")
	var e errEnv
	_ = json.Unmarshal([]byte(raw), &e)
	if status != 502 || e.Error.Code != "upstream" || len([]rune(e.Error.Message)) != 140 {
		t.Fatalf("%d %s", status, raw)
	}
	assertNoCredential(t, raw)
	wh, _, _ := h.svc.Store.Get(context.Background(), h.uid, KindDiscord)
	if wh.Failures != 2 || wh.LastError != "Unknown Webhook" || wh.DisabledAt != nil {
		t.Fatalf("a failed test must change nothing: %+v", wh)
	}
	h.poster.PostErr = &DiscordError{Status: 429, Body: "rate limited"}
	if status, _, _ := h.call("POST", "/v1/notifications/discord/test", ""); status != 502 {
		t.Fatalf("discord 429 -> %d, want 502", status)
	}
	h.poster.PostErr = io.ErrUnexpectedEOF
	_, raw, _ = h.call("POST", "/v1/notifications/discord/test", "")
	_ = json.Unmarshal([]byte(raw), &e)
	if e.Error.Message != "Discord did not answer" {
		t.Fatalf("message = %q", e.Error.Message)
	}
}

func TestTestWithoutARowIs404(t *testing.T) {
	h := newAPI(t)
	h.premium()
	if status, _, _ := h.call("POST", "/v1/notifications/discord/test", ""); status != 404 {
		t.Fatalf("status %d", status)
	}
}

func TestTestRouteIsLimitedToFivePerMinute(t *testing.T) {
	h := newAPI(t)
	h.saveHook()
	for i := 1; i <= 5; i++ {
		if status, _, _ := h.call("POST", "/v1/notifications/discord/test", ""); status != 204 {
			t.Fatalf("call %d status %d", i, status)
		}
	}
	if status, _, _ := h.call("POST", "/v1/notifications/discord/test", ""); status != 429 {
		t.Fatalf("sixth call status %d, want 429", status)
	}
}

func TestMeOmitsNotificationsWhenOff(t *testing.T) {
	pool := testPool(t)
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	authStore := &auth.Store{Pool: pool}
	authn := &auth.Authenticator{Store: authStore, Log: quiet}
	mux := http.NewServeMux()
	auth.Mount(mux, &auth.Service{Store: authStore, Auth: authn, Entitlements: &entitlements.Store{Pool: pool}, Log: quiet}, 0)
	srv := httptest.NewServer(authn.Middleware(mux))
	defer srv.Close()
	sess, err := authStore.CreateSession(context.Background(), seedUser(t, pool), "email", auth.SessionTTL)
	if err != nil {
		t.Fatal(err)
	}
	req, _ := http.NewRequest("GET", srv.URL+"/v1/me", nil)
	req.AddCookie(&http.Cookie{Name: auth.SessionCookie, Value: sess.ID})
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	raw, _ := io.ReadAll(res.Body)
	var env struct {
		Data map[string]json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(raw, &env); err != nil {
		t.Fatal(err)
	}
	if _, present := env.Data["notifications"]; present {
		t.Fatalf("notifications present when off: %s", raw)
	}
}

func TestTestEmbedNamesOnlyTheEventsThatAreOn(t *testing.T) {
	cases := map[string]struct {
		events map[EventName]bool
		want   string
	}{
		"both": {bothEvents, "Your webhook works. You will get a message here when a server sim finishes or a log file you upload finishes parsing."},
		"sim":  {map[EventName]bool{EventSimFinished: true}, "Your webhook works. You will get a message here when a server sim finishes."},
		"log":  {map[EventName]bool{EventUploadParsed: true}, "Your webhook works. You will get a message here when a log file you upload finishes parsing."},
		"none": {map[EventName]bool{}, "Your webhook works. No events are on yet."},
	}
	for name, c := range cases {
		e := TestEmbed("https://foreversixty.gg/", c.events)
		if e.Description != c.want || e.URL != "https://foreversixty.gg/account#discord" || len(e.Fields) != 0 {
			t.Fatalf("%s: %+v", name, e)
		}
	}
}

func TestMalformedPutBodiesAre400(t *testing.T) {
	h := newAPI(t)
	h.saveHook()
	good := `{"events":{"sim_finished":true,"upload_parsed":true}}`
	for name, body := range map[string]string{
		"empty events":  `{"events":{}}`,
		"one key":       `{"events":{"sim_finished":true}}`,
		"no events":     `{}`,
		"unknown field": `{"events":{"sim_finished":true,"upload_parsed":true},"extra":1}`,
		"unknown event": `{"events":{"sim_finished":true,"upload_parsed":true,"x":true}}`,
		"trailing data": good + `{}`,
		"oversize":      `{"webhook_url":"` + strings.Repeat("a", 5000) + `","events":{"sim_finished":true,"upload_parsed":true}}`,
		"blank url":     `{"webhook_url":"   ","events":{"sim_finished":true,"upload_parsed":true}}`,
		"not json":      `nope`,
	} {
		status, raw, _ := h.call("PUT", "/v1/notifications/discord", body)
		var e errEnv
		_ = json.Unmarshal([]byte(raw), &e)
		if status != 400 || e.Error.Code != "invalid" {
			t.Fatalf("%s: %d %s", name, status, raw)
		}
	}
	_, raw, _ := h.call("PUT", "/v1/notifications/discord", `{"webhook_url":"   ","events":{"sim_finished":true,"upload_parsed":true}}`)
	if !strings.Contains(raw, invalidMS) {
		t.Fatalf("blank url message: %s", raw)
	}
}
