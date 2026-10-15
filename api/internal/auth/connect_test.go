package auth

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type connectMade struct {
	ConnectID string `json:"connect_id"`
	UserCode  string `json:"user_code"`
	URL       string `json:"url"`
	ExpiresIn int    `json:"expires_in"`
	Interval  int    `json:"interval"`
}

// envData reads a response's envelope and returns its data and error
// code, whatever the HTTP status was.
func envData(t *testing.T, res *http.Response) (map[string]any, string) {
	t.Helper()
	defer res.Body.Close()
	var env struct {
		Data  map[string]any `json:"data"`
		Error *struct {
			Code string `json:"code"`
		} `json:"error"`
	}
	if res.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
			t.Fatal(err)
		}
	}
	code := ""
	if env.Error != nil {
		code = env.Error.Code
	}
	return env.Data, code
}

func (h *harness) startConnect(t *testing.T, body string) connectMade {
	t.Helper()
	res, err := http.Post(h.server.URL+"/v1/devices/connect", "application/json", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("POST /v1/devices/connect = %d, want 201", res.StatusCode)
	}
	var out connectMade
	h.decode(t, res, &out)
	return out
}

func (h *harness) pollConnect(t *testing.T, id string) (int, map[string]any) {
	t.Helper()
	res, err := http.Get(h.server.URL + "/v1/devices/connect/" + id)
	if err != nil {
		t.Fatal(err)
	}
	status := res.StatusCode
	data, _ := envData(t, res)
	return status, data
}

// ageConnects pushes every connect past its expiry: the tests' clock.
func (h *harness) ageConnects(t *testing.T) {
	t.Helper()
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`update device_connects set expires_at = now() - interval '1 minute'`); err != nil {
		t.Fatal(err)
	}
}

func TestConnectHappyPathCollectsTheTokenOnce(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")
	c := h.startConnect(t, `{"name":"  Raid PC ","platform":"windows/amd64"}`)

	if len(c.ConnectID) != 43 || c.ExpiresIn != 600 || c.Interval != 2 {
		t.Fatalf("connect = %+v", c)
	}
	if len(c.UserCode) != 7 || c.UserCode[3] != '-' {
		t.Fatalf("user_code = %q, want KQ4-7TM shape", c.UserCode)
	}
	if want := "https://foreversixty.gg/connect?code=" + c.UserCode; c.URL != want {
		t.Fatalf("url = %q, want %q", c.URL, want)
	}

	if st, d := h.pollConnect(t, c.ConnectID); st != 200 || d["status"] != "pending" {
		t.Fatalf("poll = %d %v, want pending", st, d)
	}

	res := h.do(t, http.MethodGet, "/v1/devices/connect/lookup?code="+url.QueryEscape(c.UserCode), "")
	var look struct {
		Name      string `json:"name"`
		Platform  string `json:"platform"`
		ExpiresIn int    `json:"expires_in"`
	}
	h.decode(t, res, &look)
	if look.Name != "Raid PC" || look.Platform != "windows/amd64" || look.ExpiresIn < 590 || look.ExpiresIn > 600 {
		t.Fatalf("lookup = %+v", look)
	}

	res = h.do(t, http.MethodPost, "/v1/devices/connect/approve", `{"code":"`+c.UserCode+`"}`)
	var ap struct{ Name, Platform string }
	h.decode(t, res, &ap)
	if ap.Name != "Raid PC" || ap.Platform != "windows/amd64" {
		t.Fatalf("approve = %+v", ap)
	}

	// Approving made no device yet: the poll that collects it does.
	res = h.do(t, http.MethodGet, "/v1/devices", "")
	var none []Device
	h.decode(t, res, &none)
	if len(none) != 0 {
		t.Fatalf("a device exists before collection: %+v", none)
	}

	st, d := h.pollConnect(t, c.ConnectID)
	if st != 200 || d["status"] != "approved" {
		t.Fatalf("poll = %d %v, want approved", st, d)
	}
	token, _ := d["token"].(string)
	if !strings.HasPrefix(token, DeviceTokenPrefix) || d["device_id"] == "" || d["account_label"] != "raider@example.com" {
		t.Fatalf("collected = %v", d)
	}

	if st, d := h.pollConnect(t, c.ConnectID); st != http.StatusGone || d["status"] != "used" {
		t.Fatalf("second poll = %d %v, want 410 used", st, d)
	}

	res = h.do(t, http.MethodGet, "/v1/devices", "")
	var list []Device
	h.decode(t, res, &list)
	if len(list) != 1 || list[0].ID != d["device_id"] {
		t.Fatalf("devices = %+v", list)
	}
	if list[0].Name != "Raid PC" || list[0].Platform != "windows/amd64" {
		t.Fatalf("device = %+v, want the row's name and platform", list[0])
	}

	// The token authenticates as that device.
	dev, err := h.svc.Store.DeviceByToken(t.Context(), TokenHash(token))
	if err != nil || dev.ID != list[0].ID {
		t.Fatalf("device by token = %+v, %v", dev, err)
	}

	// The connect id is at rest only as its hash, and no token is at rest.
	var n int
	if err := h.svc.Store.Pool.QueryRow(t.Context(),
		`select count(*) from device_connects where id_hash = $1`, TokenHash(c.ConnectID)).Scan(&n); err != nil || n != 1 {
		t.Fatalf("id hash row = %d, %v", n, err)
	}
}

func TestConnectAccountLabelIsTheBattletagWhenThereIsOne(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "tag@example.com")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`update users set battletag = 'Baelgrim#1234' where email = 'tag@example.com'`); err != nil {
		t.Fatal(err)
	}
	c := h.startConnect(t, `{"name":"PC","platform":"linux"}`)
	res := h.do(t, http.MethodPost, "/v1/devices/connect/approve", `{"code":"`+c.UserCode+`"}`)
	res.Body.Close()
	_, d := h.pollConnect(t, c.ConnectID)
	if d["account_label"] != "Baelgrim#1234" {
		t.Fatalf("account_label = %v", d["account_label"])
	}
}

func TestConnectDeniedAndExpired(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")

	c := h.startConnect(t, `{"name":"PC","platform":"linux"}`)
	res := h.do(t, http.MethodPost, "/v1/devices/connect/deny", `{"code":"`+c.UserCode+`"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("deny = %d, want 204", res.StatusCode)
	}
	if st, d := h.pollConnect(t, c.ConnectID); st != 200 || d["status"] != "denied" {
		t.Fatalf("poll = %d %v, want denied", st, d)
	}
	// A decided code can no longer be looked up or approved.
	res = h.do(t, http.MethodGet, "/v1/devices/connect/lookup?code="+c.UserCode, "")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("lookup after deny = %d, want 404", res.StatusCode)
	}
	res = h.do(t, http.MethodPost, "/v1/devices/connect/approve", `{"code":"`+c.UserCode+`"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("approve after deny = %d, want 409", res.StatusCode)
	}

	e := h.startConnect(t, `{"name":"PC","platform":"linux"}`)
	h.ageConnects(t)
	if st, d := h.pollConnect(t, e.ConnectID); st != 200 || d["status"] != "expired" {
		t.Fatalf("poll = %d %v, want expired", st, d)
	}
	for _, call := range []struct{ method, path, body string }{
		{http.MethodGet, "/v1/devices/connect/lookup?code=" + e.UserCode, ""},
		{http.MethodPost, "/v1/devices/connect/approve", `{"code":"` + e.UserCode + `"}`},
		{http.MethodPost, "/v1/devices/connect/deny", `{"code":"` + e.UserCode + `"}`},
	} {
		res := h.do(t, call.method, call.path, call.body)
		res.Body.Close()
		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("%s %s on an expired code = %d, want 404", call.method, call.path, res.StatusCode)
		}
	}
}

func TestConnectApprovedButNeverCollectedExpiresOnTheSameClock(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")
	c := h.startConnect(t, `{"name":"PC","platform":"linux"}`)
	res := h.do(t, http.MethodPost, "/v1/devices/connect/approve", `{"code":"`+c.UserCode+`"}`)
	res.Body.Close()
	h.ageConnects(t)
	if st, d := h.pollConnect(t, c.ConnectID); st != 200 || d["status"] != "expired" {
		t.Fatalf("poll = %d %v, want expired", st, d)
	}
	res = h.do(t, http.MethodGet, "/v1/devices", "")
	var list []Device
	h.decode(t, res, &list)
	if len(list) != 0 {
		t.Fatalf("an expired approval made a device: %+v", list)
	}
}

func TestConnectUnknownIdIs404InOneShape(t *testing.T) {
	h := newHarness(t)
	for _, id := range []string{"nope", strings.Repeat("A", 43), "%00"} {
		res, err := http.Get(h.server.URL + "/v1/devices/connect/" + id)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("poll %q = %d, want 404", id, res.StatusCode)
		}
		_, code := envData(t, res)
		if code != "not_found" {
			t.Fatalf("poll %q error code = %q", id, code)
		}
	}
}

func TestConnectSessionRoutesNeedASession(t *testing.T) {
	h := newHarness(t)
	c := h.startConnect(t, `{"name":"PC","platform":"linux"}`)
	for _, call := range []struct{ method, path, body string }{
		{http.MethodGet, "/v1/devices/connect/lookup?code=" + c.UserCode, ""},
		{http.MethodPost, "/v1/devices/connect/approve", `{"code":"` + c.UserCode + `"}`},
		{http.MethodPost, "/v1/devices/connect/deny", `{"code":"` + c.UserCode + `"}`},
	} {
		req, _ := http.NewRequest(call.method, h.server.URL+call.path, strings.NewReader(call.body))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("%s %s without a session = %d, want 401", call.method, call.path, res.StatusCode)
		}
	}
}

func TestConnectCodeIsMatchedLooselyAndMalformedOnesAre404(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")
	c := h.startConnect(t, `{"name":"PC","platform":"linux"}`)
	raw := strings.ToLower(strings.Replace(c.UserCode, "-", " ", 1))
	res := h.do(t, http.MethodGet, "/v1/devices/connect/lookup?code="+url.QueryEscape(raw), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("lookup %q = %d, want 200", raw, res.StatusCode)
	}
	for _, bad := range []string{"", "AAA-AAA", "KQ4-7T", "KQ4-7TMM", "BCD-FG!"} {
		res := h.do(t, http.MethodGet, "/v1/devices/connect/lookup?code="+url.QueryEscape(bad), "")
		res.Body.Close()
		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("lookup %q = %d, want 404", bad, res.StatusCode)
		}
	}
	res = h.do(t, http.MethodPost, "/v1/devices/connect/approve", `{"code":"`+raw+`"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("approve %q = %d, want 200", raw, res.StatusCode)
	}
}

func TestConnectDecidedTwiceIs409(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")
	c := h.startConnect(t, `{"name":"PC","platform":"linux"}`)
	res := h.do(t, http.MethodPost, "/v1/devices/connect/approve", `{"code":"`+c.UserCode+`"}`)
	res.Body.Close()
	res = h.do(t, http.MethodPost, "/v1/devices/connect/approve", `{"code":"`+c.UserCode+`"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("second approve = %d, want 409", res.StatusCode)
	}
	res = h.do(t, http.MethodPost, "/v1/devices/connect/deny", `{"code":"`+c.UserCode+`"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("deny after approve = %d, want 409", res.StatusCode)
	}
}

func TestConnectIsRateLimitedPerIPAndPollingIsNot(t *testing.T) {
	h := newHarness(t)
	res, err := http.Post(h.server.URL+"/v1/devices/connect", "application/json", strings.NewReader(`nope`))
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("bad body = %d, want 400", res.StatusCode)
	}
	// That request spent one of the budget; the rest goes here.
	var last connectMade
	for i := 0; i < claimsPerHour-1; i++ {
		last = h.startConnect(t, `{"name":"PC","platform":"linux"}`)
	}
	res, err = http.Post(h.server.URL+"/v1/devices/connect", "application/json", strings.NewReader(`{}`))
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	if res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("connect over the budget = %d, want 429", res.StatusCode)
	}
	for i := 0; i < claimsPerHour+5; i++ {
		if st, _ := h.pollConnect(t, last.ConnectID); st != http.StatusOK {
			t.Fatalf("poll %d = %d", i, st)
		}
	}
}

func TestClaimAnswersWithTheAccountLabel(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")
	res := h.do(t, http.MethodPost, "/v1/devices/pair", "")
	var p struct{ Code string }
	h.decode(t, res, &p)
	claimed, err := http.Post(h.server.URL+"/v1/devices/claim", "application/json",
		strings.NewReader(`{"code":"`+p.Code+`","name":"PC","platform":"linux"}`))
	if err != nil {
		t.Fatal(err)
	}
	if claimed.StatusCode != http.StatusCreated {
		t.Fatalf("claim = %d", claimed.StatusCode)
	}
	var out map[string]string
	h.decode(t, claimed, &out)
	if out["account_label"] != "raider@example.com" || out["token"] == "" {
		t.Fatalf("claim = %v", out)
	}
}

func TestSafeNextKeepsTheConnectCodeThroughBattleNet(t *testing.T) {
	const next = "/connect?code=KQ4-7TM"
	if got := safeNext(next); got != next {
		t.Fatalf("safeNext(%q) = %q", next, got)
	}
	h := newHarness(t)
	h.withBattleNet(t, "12345", "Baelgrim#1234")
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/start?next="+url.QueryEscape(next), "")
	res.Body.Close()
	target, _ := url.Parse(res.Header.Get("Location"))
	state := target.Query().Get("state")
	res = h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state="+state, "")
	defer res.Body.Close()
	if got := res.Header.Get("Location"); got != "https://foreversixty.gg"+next {
		t.Fatalf("redirected to %q, want the connect page with its code", got)
	}
}

func TestUserCodeShapeAndNormalisation(t *testing.T) {
	for i := 0; i < 200; i++ {
		c := NewUserCode()
		if got, ok := NormalizeUserCode(FormatUserCode(c)); !ok || got != c {
			t.Fatalf("code %q does not round-trip: %q %v", c, got, ok)
		}
	}
	if got, ok := NormalizeUserCode("kq4 7tm"); !ok || got != "KQ47TM" {
		t.Fatalf("normalise = %q %v", got, ok)
	}
	if _, ok := NormalizeUserCode("AEI-OU0"); ok {
		t.Fatal("vowels and 0 are outside the alphabet")
	}
}

func TestACollidingUserCodeIsRejectedByTheIndexAndRetried(t *testing.T) {
	s := testStore(t)
	first, err := s.CreateConnect(t.Context(), "A", "linux")
	if err != nil {
		t.Fatal(err)
	}
	live, _ := NormalizeUserCode(first.UserCode)
	draws := []string{live, live, "BCDFGH"}
	old := newUserCode
	newUserCode = func() string {
		c := draws[0]
		if len(draws) > 1 {
			draws = draws[1:]
		}
		return c
	}
	t.Cleanup(func() { newUserCode = old })

	second, err := s.CreateConnect(t.Context(), "B", "linux")
	if err != nil {
		t.Fatalf("the retry should have found a free code: %v", err)
	}
	if second.UserCode != "BCD-FGH" || second.UserCode == first.UserCode {
		t.Fatalf("second code = %q", second.UserCode)
	}
	// With nothing but collisions on offer it gives up rather than spin.
	newUserCode = func() string { return live }
	if _, err := s.CreateConnect(t.Context(), "C", "linux"); err == nil {
		t.Fatal("a store with no free code should fail")
	}
}

func TestOldConnectRowsAreSweptOnInsert(t *testing.T) {
	s := testStore(t)
	c, err := s.CreateConnect(t.Context(), "A", "linux")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.Pool.Exec(t.Context(),
		`update device_connects set expires_at = now() - interval '2 hours'`); err != nil {
		t.Fatal(err)
	}
	if _, err := s.CreateConnect(t.Context(), "B", "linux"); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := s.Pool.QueryRow(t.Context(),
		`select count(*) from device_connects where id_hash = $1`, TokenHash(c.ConnectID)).Scan(&n); err != nil || n != 0 {
		t.Fatalf("swept row count = %d, %v", n, err)
	}
}

func TestEmailSignInCarriesNextThroughTheLink(t *testing.T) {
	h := newHarness(t)
	const next = "/connect?code=KQ4-7TM"
	res := h.do(t, http.MethodPost, "/v1/auth/email", `{"email":"raider@example.com","next":"`+next+`"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("email = %d", res.StatusCode)
	}
	text := h.mailer.Sent[len(h.mailer.Sent)-1].Text
	i := strings.Index(text, "https://api.foreversixty.gg")
	link, _, _ := strings.Cut(strings.TrimSpace(text[i:]), "\n")
	res = h.do(t, http.MethodGet, strings.TrimPrefix(link, "https://api.foreversixty.gg"), "")
	res.Body.Close()
	if got := res.Header.Get("Location"); got != "https://foreversixty.gg"+next {
		t.Fatalf("redirected to %q, want the connect page", got)
	}

	// An off-site next never makes it into the link as anything but the default.
	res = h.do(t, http.MethodPost, "/v1/auth/email", `{"email":"raider@example.com","next":"https://evil.example"}`)
	res.Body.Close()
	text = h.mailer.Sent[len(h.mailer.Sent)-1].Text
	i = strings.Index(text, "https://api.foreversixty.gg")
	link, _, _ = strings.Cut(strings.TrimSpace(text[i:]), "\n")
	res = h.do(t, http.MethodGet, strings.TrimPrefix(link, "https://api.foreversixty.gg"), "")
	res.Body.Close()
	if got := res.Header.Get("Location"); got != "https://foreversixty.gg/logs" {
		t.Fatalf("redirected to %q, want /logs", got)
	}
}
