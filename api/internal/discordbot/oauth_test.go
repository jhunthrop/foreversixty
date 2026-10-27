package discordbot

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
)

func mustQuery(t *testing.T, raw string) url.Values {
	t.Helper()
	u, err := url.Parse(raw)
	if err != nil || u.Host != "discord.com" || u.Path != "/oauth2/authorize" {
		t.Fatalf("url = %q err=%v", raw, err)
	}
	return u.Query()
}

func TestInstallAndLinkURLs(t *testing.T) {
	o := NewHTTPOAuth("cid", "sec", "https://x.test/cb")
	q := mustQuery(t, o.InstallURL("st&1"))
	if q.Get("scope") != "bot applications.commands" || q.Get("permissions") != "84992" || q.Get("state") != "st&1" ||
		q.Get("client_id") != "cid" || q.Get("redirect_uri") != "https://x.test/cb" || q.Get("response_type") != "code" {
		t.Fatalf("install = %v", q)
	}
	l := mustQuery(t, o.LinkURL("s"))
	if l.Get("scope") != "identify" || l.Get("state") != "s" || l.Get("permissions") != "" {
		t.Fatalf("link = %v", l)
	}
	if InstallPermissions != 1<<10|1<<11|1<<14|1<<16 {
		t.Fatal("permissions drifted")
	}
}

func TestExchangeIdentityPrefersGlobalNameAndDropsToken(t *testing.T) {
	var meCalls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth2/token":
			_ = r.ParseForm()
			if r.Form.Get("code") != "the-code" || r.Form.Get("client_secret") != "sec" || r.Form.Get("grant_type") != "authorization_code" {
				w.WriteHeader(400)
				return
			}
			_, _ = w.Write([]byte(`{"access_token":"bearer-1","token_type":"Bearer","scope":"identify"}`))
		case "/users/@me":
			meCalls.Add(1)
			if r.Header.Get("Authorization") != "Bearer bearer-1" {
				w.WriteHeader(401)
				return
			}
			_, _ = w.Write([]byte(`{"id":"42","username":"zul","global_name":"zulmara"}`))
		}
	}))
	defer srv.Close()
	o := NewHTTPOAuth("cid", "sec", "https://x.test/cb")
	o.BaseURL = srv.URL
	who, err := o.ExchangeIdentity(context.Background(), "the-code")
	if err != nil || who != (Identity{ID: "42", DisplayName: "zulmara"}) || meCalls.Load() != 1 {
		t.Fatalf("who=%+v err=%v calls=%d", who, err, meCalls.Load())
	}
}

func TestExchangeIdentityFallsBackToUsername(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/oauth2/token" {
			_, _ = w.Write([]byte(`{"access_token":"b"}`))
			return
		}
		_, _ = w.Write([]byte(`{"id":"7","username":"zul","global_name":null}`))
	}))
	defer srv.Close()
	o := NewHTTPOAuth("c", "s", "r")
	o.BaseURL = srv.URL
	who, err := o.ExchangeIdentity(context.Background(), "x")
	if err != nil || who.DisplayName != "zul" {
		t.Fatalf("who=%+v err=%v", who, err)
	}
}

func TestExchangeInstallReadsTheTokenAnswersGuild(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"access_token":"a","guild":{"id":"srv1","name":"Olympus Discord"}}`))
	}))
	defer srv.Close()
	o := NewHTTPOAuth("c", "s", "r")
	o.BaseURL = srv.URL
	got, err := o.ExchangeInstall(context.Background(), "secret-code")
	if err != nil || got != (InstalledServer{ID: "srv1", Name: "Olympus Discord"}) {
		t.Fatalf("got=%+v err=%v", got, err)
	}
}

func TestExchangeErrorsNeverContainTheCode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(`{"error":"invalid_grant"}`))
	}))
	o := NewHTTPOAuth("c", "s", "r")
	o.BaseURL = srv.URL
	_, err1 := o.ExchangeInstall(context.Background(), "secret-code")
	_, err2 := o.ExchangeIdentity(context.Background(), "secret-code")
	srv.Close()
	_, err3 := o.ExchangeInstall(context.Background(), "secret-code")
	for _, err := range []error{err1, err2, err3} {
		if err == nil || strings.Contains(err.Error(), "secret-code") {
			t.Fatalf("err = %v", err)
		}
	}
	noGuild := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"access_token":"a"}`))
	}))
	defer noGuild.Close()
	o.BaseURL = noGuild.URL
	if _, err := o.ExchangeInstall(context.Background(), "secret-code"); err == nil {
		t.Fatal("a token answer without a guild must fail")
	}
}
