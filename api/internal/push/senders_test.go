package push

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"golang.org/x/oauth2"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

func testP8(t *testing.T) (*ecdsa.PrivateKey, []byte) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	der, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return key, pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})
}

type apnsCall struct {
	path    string
	header  http.Header
	body    map[string]any
	isHTTP2 bool
}

func apnsServer(t *testing.T, status int, reason string) (*httptest.Server, *[]apnsCall, *sync.Mutex) {
	t.Helper()
	var calls []apnsCall
	var mu sync.Mutex
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		var body map[string]any
		_ = json.Unmarshal(raw, &body)
		mu.Lock()
		calls = append(calls, apnsCall{path: r.URL.Path, header: r.Header.Clone(), body: body, isHTTP2: r.ProtoMajor == 2})
		mu.Unlock()
		w.WriteHeader(status)
		if reason != "" {
			_, _ = w.Write([]byte(`{"reason":"` + reason + `"}`))
		}
	}))
	srv.EnableHTTP2 = true
	srv.StartTLS()
	t.Cleanup(srv.Close)
	return srv, &calls, &mu
}

func raidMessage(token string) Message {
	return Message{
		Token: token, Platform: IOS, Kind: events.RaidPosted, Title: "Molten Core", Body: "Thursday 20:00",
		Link: "https://foreversixty.gg/guild/x/raids/9", Thread: "raid-9", Collapse: "raid.posted-9", Category: RaidSignupCategory,
		Actions: []events.Action{{ID: "signup_in", Title: "In", Method: "PUT", Path: "/v1/guilds/1/raids/9/signups/us/normal/kraggor", Body: map[string]string{"state": "in"}}},
		Data:    map[string]string{"raid_id": "9"},
	}
}

func verifyES256(t *testing.T, pub *ecdsa.PublicKey, jwt string) map[string]any {
	t.Helper()
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		t.Fatalf("jwt %q", jwt)
	}
	sig, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil || len(sig) != 64 {
		t.Fatalf("signature %v", err)
	}
	digest := sha256.Sum256([]byte(parts[0] + "." + parts[1]))
	if !ecdsa.Verify(pub, digest[:], new(big.Int).SetBytes(sig[:32]), new(big.Int).SetBytes(sig[32:])) {
		t.Fatal("jwt signature does not verify")
	}
	raw, _ := base64.RawURLEncoding.DecodeString(parts[1])
	var claims map[string]any
	_ = json.Unmarshal(raw, &claims)
	return claims
}

func TestAPNsSendsOverHTTP2WithASignedProviderToken(t *testing.T) {
	key, p8 := testP8(t)
	srv, calls, _ := apnsServer(t, http.StatusOK, "")
	s, err := NewAPNs(APNsConfig{KeyP8: p8, KeyID: "KEY123", TeamID: "TEAM45", Topic: "gg.foreversixty.app", BaseURL: srv.URL}, srv.Client())
	if err != nil {
		t.Fatal(err)
	}
	for range 2 {
		if err := s.Send(context.Background(), raidMessage("abcdef0123")); err != nil {
			t.Fatal(err)
		}
	}
	c := (*calls)[0]
	if !c.isHTTP2 || c.path != "/3/device/abcdef0123" {
		t.Fatalf("call %+v", c)
	}
	for h, want := range map[string]string{"apns-topic": "gg.foreversixty.app", "apns-push-type": "alert", "apns-collapse-id": "raid.posted-9", "apns-priority": "10"} {
		if got := c.header.Get(h); got != want {
			t.Errorf("%s = %q", h, got)
		}
	}
	jwt := strings.TrimPrefix(c.header.Get("authorization"), "bearer ")
	claims := verifyES256(t, &key.PublicKey, jwt)
	if claims["iss"] != "TEAM45" {
		t.Fatalf("claims %v", claims)
	}
	if (*calls)[1].header.Get("authorization") != c.header.Get("authorization") {
		t.Fatal("the provider token was not reused")
	}
	aps := c.body["aps"].(map[string]any)
	if aps["thread-id"] != "raid-9" || aps["category"] != RaidSignupCategory || c.body["link"] == nil || c.body["actions"] == nil {
		t.Fatalf("body %v", c.body)
	}
}

func TestAPNsRenewsTheProviderTokenAfterItsLife(t *testing.T) {
	_, p8 := testP8(t)
	srv, calls, _ := apnsServer(t, http.StatusOK, "")
	s, _ := NewAPNs(APNsConfig{KeyP8: p8, KeyID: "K", TeamID: "T", Topic: "t", BaseURL: srv.URL}, srv.Client())
	at := time.Date(2026, 11, 4, 0, 0, 0, 0, time.UTC)
	s.now = func() time.Time { return at }
	_ = s.Send(context.Background(), raidMessage("aa"))
	at = at.Add(apnsTokenLife + time.Second)
	_ = s.Send(context.Background(), raidMessage("aa"))
	if (*calls)[0].header.Get("authorization") == (*calls)[1].header.Get("authorization") {
		t.Fatal("token not renewed")
	}
}

func TestAPNsRefusals(t *testing.T) {
	_, p8 := testP8(t)
	cases := []struct {
		status  int
		reason  string
		invalid bool
	}{
		{http.StatusGone, "Unregistered", true},
		{http.StatusBadRequest, "BadDeviceToken", true},
		{http.StatusBadRequest, "DeviceTokenNotForTopic", true},
		{http.StatusForbidden, "ExpiredProviderToken", false},
		{http.StatusTooManyRequests, "TooManyRequests", false},
	}
	for _, c := range cases {
		srv, _, _ := apnsServer(t, c.status, c.reason)
		s, _ := NewAPNs(APNsConfig{KeyP8: p8, KeyID: "K", TeamID: "T", Topic: "t", BaseURL: srv.URL}, srv.Client())
		err := s.Send(context.Background(), raidMessage("aa"))
		var svc *ServiceError
		if err == nil || errors.Is(err, ErrInvalidToken) != c.invalid || !errors.As(err, &svc) || svc.Reason != c.reason {
			t.Errorf("%d %s: %v", c.status, c.reason, err)
		}
		if strings.Contains(err.Error(), "PRIVATE KEY") {
			t.Fatal("the key leaked into an error")
		}
	}
}

func TestAPNsRefusesABadKeyWithoutQuotingIt(t *testing.T) {
	for _, p8 := range [][]byte{[]byte("secret-not-pem"), pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: []byte("secret")})} {
		_, err := NewAPNs(APNsConfig{KeyP8: p8, KeyID: "K", TeamID: "T", Topic: "t", BaseURL: "https://x"}, nil)
		if err == nil || strings.Contains(err.Error(), "secret") {
			t.Fatalf("err %v", err)
		}
	}
	if _, err := NewAPNs(APNsConfig{}, nil); err == nil {
		t.Fatal("empty config accepted")
	}
}

func fcmServer(t *testing.T, status int, body string) (*httptest.Server, *[]map[string]any, *[]http.Header) {
	t.Helper()
	var bodies []map[string]any
	var headers []http.Header
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/projects/forever-app/messages:send" {
			t.Errorf("path %s", r.URL.Path)
		}
		raw, _ := io.ReadAll(r.Body)
		var b map[string]any
		_ = json.Unmarshal(raw, &b)
		bodies, headers = append(bodies, b), append(headers, r.Header.Clone())
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return srv, &bodies, &headers
}

func TestFCMSendsTheV1Message(t *testing.T) {
	srv, bodies, headers := fcmServer(t, http.StatusOK, `{"name":"x"}`)
	s := newFCM("forever-app", srv.URL, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "at-1"}), srv.Client())
	m := raidMessage("fcm-token")
	m.Platform = Android
	if err := s.Send(context.Background(), m); err != nil {
		t.Fatal(err)
	}
	if (*headers)[0].Get("Authorization") != "Bearer at-1" {
		t.Fatal("no bearer")
	}
	msg := (*bodies)[0]["message"].(map[string]any)
	data := msg["data"].(map[string]any)
	android := msg["android"].(map[string]any)
	if msg["token"] != "fcm-token" || data["link"] == nil || data["raid_id"] != "9" || data["kind"] != "raid.posted" {
		t.Fatalf("message %v", msg)
	}
	var actions []events.Action
	if err := json.Unmarshal([]byte(data["actions"].(string)), &actions); err != nil || actions[0].ID != "signup_in" {
		t.Fatalf("actions %v %v", data["actions"], err)
	}
	if android["collapse_key"] != "raid.posted-9" || android["notification"].(map[string]any)["tag"] != "raid.posted-9" {
		t.Fatalf("android %v", android)
	}
}

func TestFCMRefusals(t *testing.T) {
	cases := []struct {
		status  int
		body    string
		invalid bool
	}{
		{http.StatusNotFound, `{"error":{"status":"NOT_FOUND","details":[{"errorCode":"UNREGISTERED"}]}}`, true},
		{http.StatusBadRequest, `{"error":{"status":"INVALID_ARGUMENT","details":[{"errorCode":"UNREGISTERED"}]}}`, true},
		{http.StatusForbidden, `{"error":{"status":"PERMISSION_DENIED","details":[{"errorCode":"SENDER_ID_MISMATCH"}]}}`, true},
		{http.StatusServiceUnavailable, `{"error":{"status":"UNAVAILABLE"}}`, false},
	}
	for _, c := range cases {
		srv, _, _ := fcmServer(t, c.status, c.body)
		s := newFCM("forever-app", srv.URL, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "x"}), srv.Client())
		err := s.Send(context.Background(), raidMessage("t"))
		if err == nil || errors.Is(err, ErrInvalidToken) != c.invalid {
			t.Errorf("%d: %v", c.status, err)
		}
	}
}

func TestNewFCMRefusesABadServiceAccount(t *testing.T) {
	for _, raw := range []string{`nope`, `{"type":"service_account"}`, `{"project_id":"p","private_key":"secret"}`} {
		_, err := NewFCM(context.Background(), []byte(raw), FCMBaseURL, nil)
		if err == nil || strings.Contains(err.Error(), "secret") {
			t.Errorf("%s: %v", raw, err)
		}
	}
}

func TestTimeSensitiveMessagesBreakThrough(t *testing.T) {
	m := raidMessage("aa")
	m.TimeSensitive = true
	raw, _ := apnsBody(m)
	if !strings.Contains(string(raw), `"interruption-level":"time-sensitive"`) {
		t.Fatalf("apns %s", raw)
	}
	plain, _ := apnsBody(raidMessage("aa"))
	if strings.Contains(string(plain), "interruption-level") {
		t.Fatalf("plain apns %s", plain)
	}
	hi, _ := fcmBody(m)
	lo, _ := fcmBody(raidMessage("aa"))
	if !strings.Contains(string(hi), `"priority":"HIGH"`) || !strings.Contains(string(lo), `"priority":"NORMAL"`) {
		t.Fatalf("fcm %s / %s", hi, lo)
	}
	if !MessageOf(events.Event{Kind: events.RaidRosterMoved}, Token{}).TimeSensitive || MessageOf(events.Event{Kind: events.RaidPosted}, Token{}).TimeSensitive {
		t.Fatal("MessageOf time-sensitive")
	}
}
