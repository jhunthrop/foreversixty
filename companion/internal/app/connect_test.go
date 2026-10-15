package app_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/companion/internal/secret"
)

// connectView is the snapshot's sign-in as a test reads it.
func (h *harness) connectView() app.ConnectView { return h.app.Snapshot().Connect }

// until polls a condition for up to three seconds. The sign-in polls
// every two milliseconds under test, so this is a handful of laps.
func until(t *testing.T, what string, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(3 * time.Second)
	for !cond() {
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for %s", what)
		}
		time.Sleep(time.Millisecond)
	}
}

func (h *harness) startConnect() map[string]any {
	h.t.Helper()
	code, body := h.call(http.MethodPost, h.url("/api/connect"), nil)
	if code != http.StatusOK || body["ok"] != true {
		h.t.Fatalf("connect = %d %v", code, body)
	}
	return body["data"].(map[string]any)["connect"].(map[string]any)
}

func TestConnectMintsOpensTheBrowserAndWaits(t *testing.T) {
	h := newHarness(t)
	view := h.startConnect()
	if view["state"] != "waiting" || view["user_code"] != fakeapi.ConnectCode ||
		view["url"] == "" || view["expires_at"] == nil || view["trouble"] != false {
		t.Fatalf("connect view = %v", view)
	}
	if got := h.openedLinks(); len(got) != 1 || got[0] != view["url"] {
		t.Fatalf("the browser was asked for %v, want the connect url once", got)
	}
	if h.srv.ConnectName(fakeapi.ConnectID) == "" {
		t.Error("the connect was started without a device name")
	}
	// The connect id is the secret: it is never in the snapshot.
	if raw := h.app.Snapshot(); raw.Headline.Key != "connecting" {
		t.Errorf("headline = %+v, want the sign-in wait", raw.Headline)
	}
}

func TestOnlyOneConnectRunsAtATime(t *testing.T) {
	h := newHarness(t)
	h.startConnect()
	h.startConnect()
	h.startConnect()
	if n := h.srv.Connects(); n != 1 {
		t.Fatalf("%d connects were minted, want 1", n)
	}
	if got := h.openedLinks(); len(got) != 1 {
		t.Fatalf("the browser was opened %d times", len(got))
	}
}

func TestApprovalStoresTheTokenAndWhoItBelongsTo(t *testing.T) {
	h := newHarness(t)
	h.startConnect()
	h.srv.ConnectDecide("approved")
	until(t, "the device to be paired", func() bool { return h.app.Snapshot().Paired })

	s := h.app.Snapshot()
	if s.AccountLabel != fakeapi.AccountLabel {
		t.Errorf("account label = %q", s.AccountLabel)
	}
	if s.Connect.State != "idle" {
		t.Errorf("connect = %+v, want idle once connected", s.Connect)
	}
	stored, err := config.Load(h.dirs.ConfigFile())
	if err != nil {
		t.Fatal(err)
	}
	if stored.DeviceToken != fakeapi.Token || stored.AccountLabel != fakeapi.AccountLabel {
		t.Fatalf("stored = token %q, label %q", stored.DeviceToken, stored.AccountLabel)
	}
	// A device that is connected cannot start another sign-in.
	code, body := h.call(http.MethodPost, h.url("/api/connect"), nil)
	if code != http.StatusConflict || body["error"].(map[string]any)["code"] != "already_connected" {
		t.Fatalf("connect while connected = %d %v", code, body)
	}
	// Forgetting clears the label too.
	if code, _ := h.call(http.MethodPost, h.url("/api/unpair"), map[string]string{}); code != http.StatusOK {
		t.Fatal("unpair failed")
	}
	if got := h.app.Snapshot().AccountLabel; got != "" {
		t.Errorf("label after forgetting = %q", got)
	}
	if stored, _ = config.Load(h.dirs.ConfigFile()); stored.AccountLabel != "" {
		t.Errorf("stored label after forgetting = %q", stored.AccountLabel)
	}
}

func TestADeniedOrExpiredConnectFailsWithItsReasonAndCanBeTriedAgain(t *testing.T) {
	for _, c := range []struct{ decide, reason string }{
		{"denied", "denied"}, {"expired", "expired"},
	} {
		h := newHarness(t)
		h.startConnect()
		h.srv.ConnectDecide(c.decide)
		until(t, c.decide, func() bool { return h.connectView().State == "failed" })
		if v := h.connectView(); v.Reason != c.reason || v.UserCode != "" || v.URL != "" {
			t.Fatalf("%s: view = %+v", c.decide, v)
		}
		if h.app.Snapshot().Paired {
			t.Fatalf("%s: paired", c.decide)
		}
		if h.app.Snapshot().Headline.Key != "not_connected" {
			t.Errorf("%s: headline = %+v", c.decide, h.app.Snapshot().Headline)
		}
		// Try again: a failed connect does not block a new one.
		h.startConnect()
		if h.srv.Connects() != 2 {
			t.Fatalf("%s: retry minted %d connects", c.decide, h.srv.Connects())
		}
	}
}

func TestCancelStopsPollingAndReturnsToIdle(t *testing.T) {
	h := newHarness(t)
	h.startConnect()
	until(t, "a poll", func() bool { return h.srv.ConnectPolls(fakeapi.ConnectID) > 0 })
	code, body := h.call(http.MethodPost, h.url("/api/connect/cancel"), nil)
	if code != http.StatusOK || body["data"].(map[string]any)["connect"].(map[string]any)["state"] != "idle" {
		t.Fatalf("cancel = %d %v", code, body)
	}
	time.Sleep(20 * time.Millisecond) // a poll in flight may still land
	n := h.srv.ConnectPolls(fakeapi.ConnectID)
	time.Sleep(40 * time.Millisecond)
	if after := h.srv.ConnectPolls(fakeapi.ConnectID); after != n {
		t.Fatalf("polling went on after cancel: %d -> %d", n, after)
	}
	// An approval after the cancel changes nothing.
	h.srv.ConnectDecide("approved")
	time.Sleep(20 * time.Millisecond)
	if h.app.Snapshot().Paired {
		t.Fatal("a cancelled connect paired")
	}
}

func TestReopenOpensTheSameLinkAndNeverMintsAnother(t *testing.T) {
	h := newHarness(t)
	if code, body := h.call(http.MethodPost, h.url("/api/connect/reopen"), nil); code != http.StatusConflict ||
		body["error"].(map[string]any)["code"] != "no_connect" {
		t.Fatalf("reopen with nothing to reopen = %d %v", code, body)
	}
	view := h.startConnect()
	if code, _ := h.call(http.MethodPost, h.url("/api/connect/reopen"), nil); code != http.StatusOK {
		t.Fatal("reopen failed")
	}
	got := h.openedLinks()
	if len(got) != 2 || got[0] != got[1] || got[1] != view["url"] {
		t.Fatalf("links opened = %v", got)
	}
	if h.srv.Connects() != 1 {
		t.Fatalf("reopen minted a connect: %d", h.srv.Connects())
	}
}

func TestABrowserThatWillNotOpenLeavesTheLinkOnScreenAndKeepsPolling(t *testing.T) {
	h := newHarness(t)
	h.openErr = errors.New("no browser")
	view := h.startConnect()
	if view["state"] != "blocked" || view["url"] == "" || view["user_code"] != fakeapi.ConnectCode {
		t.Fatalf("view = %v", view)
	}
	if h.app.Snapshot().Headline.Key != "connecting" {
		t.Error("the headline lost the wait")
	}
	h.srv.ConnectDecide("approved")
	until(t, "pairing without a browser", func() bool { return h.app.Snapshot().Paired })
}

func TestNetworkErrorsRetrySilentlyAndTroubleComesAfterThree(t *testing.T) {
	h := newHarness(t)
	h.startConnect()
	until(t, "a poll", func() bool { return h.srv.ConnectPolls(fakeapi.ConnectID) > 0 })
	h.srv.Offline(true)
	until(t, "trouble", func() bool { return h.connectView().Trouble })
	if v := h.connectView(); v.State != "waiting" {
		t.Fatalf("a flaky network ended the sign-in: %+v", v)
	}
	h.srv.Offline(false)
	until(t, "trouble to clear", func() bool { return !h.connectView().Trouble })
	h.srv.ConnectDecide("approved")
	until(t, "pairing after the outage", func() bool { return h.app.Snapshot().Paired })
}

// The first two failed polls are silent: only the third says anything.
func TestTroubleIsNotShownBeforeThreeFailures(t *testing.T) {
	h := newHarnessWith(t, nil, func(o *app.Options) { o.ConnectInterval = 30 * time.Millisecond })
	h.startConnect()
	until(t, "a poll", func() bool { return h.srv.ConnectPolls(fakeapi.ConnectID) > 0 })
	h.srv.Offline(true)
	// Two intervals later at most two polls have failed.
	time.Sleep(45 * time.Millisecond)
	if h.connectView().Trouble {
		t.Fatal("trouble was reported before three failed polls")
	}
	until(t, "the third failure", func() bool { return h.connectView().Trouble })
}

func TestAConnectThatCannotStartFailsWithAReason(t *testing.T) {
	h := newHarness(t)
	h.srv.Offline(true) // answers 503: the site is up and unwell
	view := h.startConnect()
	if view["state"] != "failed" || view["reason"] != "server" {
		t.Fatalf("503: view = %v", view)
	}
	if len(h.openedLinks()) != 0 {
		t.Error("the browser was opened for a connect that did not start")
	}

	h2 := newHarness(t)
	h2.srv.Close() // nothing to connect to: the network
	view = h2.startConnect()
	if view["state"] != "failed" || view["reason"] != "offline" {
		t.Fatalf("closed: view = %v", view)
	}
}

func TestATokenThatCannotBeSavedFailsTheConnectWithoutPairing(t *testing.T) {
	h := newHarnessWith(t, nil, func(o *app.Options) {
		o.Secret = failingSecret{o.Secret}
	})
	h.startConnect()
	h.srv.ConnectDecide("approved")
	until(t, "the save failure", func() bool { return h.connectView().State == "failed" })
	if v := h.connectView(); v.Reason != "save" {
		t.Fatalf("view = %+v", v)
	}
	if h.app.Snapshot().Paired {
		t.Fatal("paired without a saved token")
	}
}

func TestThePairingFallbackSaysWhoItConnectedAsAndNamesItsFailures(t *testing.T) {
	h := newHarness(t)
	code, body := h.call(http.MethodPost, h.url("/api/pair"), map[string]string{"code": fakeapi.PairCode})
	if code != http.StatusOK || body["data"].(map[string]any)["account_label"] != fakeapi.AccountLabel {
		t.Fatalf("pair = %d %v", code, body)
	}

	for _, c := range []struct {
		name   string
		setup  func(*harness)
		code   string
		status int
		code_  string
	}{
		{"empty", func(*harness) {}, "", http.StatusBadRequest, "empty_code"},
		{"wrong", func(*harness) {}, "NOPE", http.StatusBadRequest, "bad_code"},
		{"outage", func(h *harness) { h.srv.Offline(true) }, fakeapi.PairCode, http.StatusBadGateway, "server"},
		{"closed", func(h *harness) { h.srv.Close() }, fakeapi.PairCode, http.StatusBadGateway, "offline"},
	} {
		h := newHarness(t)
		c.setup(h)
		status, body := h.call(http.MethodPost, h.url("/api/pair"), map[string]string{"code": c.code})
		e, _ := body["error"].(map[string]any)
		if status != c.status || e["code"] != c.code_ || e["message"] != app.Message(c.code_) {
			t.Errorf("%s: %d %v", c.name, status, body)
		}
	}
}

// failingSecret reads fine and cannot write: the machine that would not let the companion save.
type failingSecret struct{ secret.Store }

func (failingSecret) SetToken(string) error { return errors.New("the keychain is locked") }
