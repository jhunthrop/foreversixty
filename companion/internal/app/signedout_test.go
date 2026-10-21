package app_test

import (
	"bytes"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
)

// lockedBuf is a log sink that tests and the app's goroutines can share.
type lockedBuf struct {
	mu sync.Mutex
	b  bytes.Buffer
}

func (l *lockedBuf) Write(p []byte) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.b.Write(p)
}

func (l *lockedBuf) count(sub string) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return strings.Count(l.b.String(), sub)
}

// queued is how many files the upload queue holds on disk.
func queued(t *testing.T, dir string) int {
	t.Helper()
	n := 0
	err := filepath.WalkDir(dir, func(_ string, d os.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			n++
		}
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	return n
}

func TestARevokedDeviceSignsOutOnceAndStopsAskingEverySecond(t *testing.T) {
	var logs lockedBuf
	h := newHarnessWith(t, nil, func(o *app.Options) {
		o.Log = slog.New(slog.NewTextHandler(&logs, &slog.HandlerOptions{Level: slog.LevelDebug}))
	})
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("pairing failed")
	}
	if s := h.app.Snapshot(); s.SignedOut != nil || s.Headline.Key == "signed_out" {
		t.Fatalf("a healthy device reads signed out: %+v", s.SignedOut)
	}
	h.app.Step(t.Context(), t0) // prime the tail over the empty folder
	if err := os.WriteFile(h.log,
		[]byte(fixture.Header+fixture.Zone+fixture.Encounter(0)+fixture.Heartbeat(0)), 0o600); err != nil {
		t.Fatal(err)
	}
	// An outage first, so there is a queue on disk to keep.
	h.srv.Offline(true)
	h.app.Step(t.Context(), t0.Add(time.Second))
	held := queued(t, h.dirs.Queue)
	if held == 0 {
		t.Fatal("nothing was queued during the outage")
	}
	h.srv.Offline(false)
	h.srv.RevokeToken()

	h.app.Step(t.Context(), t0.Add(2*time.Second))
	s := h.app.Snapshot()
	if s.SignedOut == nil || s.SignedOut.Reason != "revoked" {
		t.Fatalf("signed_out = %+v", s.SignedOut)
	}
	if s.SignedOut.AccountLabel != fakeapi.AccountLabel || s.SignedOut.At.IsZero() {
		t.Errorf("signed_out = %+v; want the account label and a time", s.SignedOut)
	}
	if s.Paired || s.Headline.Key != "signed_out" || s.Headline.Tone != "bad" ||
		s.Headline.Title != app.CopySignedOutTitle || s.Headline.Sub != app.CopySignedOutSub ||
		s.Headline.Action == nil || s.Headline.Action.Page != "device" ||
		s.Headline.Action.Label != app.CopyNotConnectedBtn {
		t.Errorf("paired %v, headline %+v", s.Paired, s.Headline)
	}
	if got := h.app.Headline().Title; got != app.CopySignedOutTitle {
		t.Errorf("the tray reads %q", got)
	}
	if stored, err := config.Load(h.dirs.ConfigFile()); err != nil || stored.DeviceToken != "" {
		t.Errorf("the token is still stored: %q, %v", stored.DeviceToken, err)
	}
	if stored, _ := config.Load(h.dirs.ConfigFile()); stored.AccountLabel != fakeapi.AccountLabel {
		t.Errorf("the account label was lost: %q", stored.AccountLabel)
	}
	if n := logs.count("no longer accepts this device"); n != 1 {
		t.Fatalf("the sign-out was logged %d times", n)
	}
	syncFails := logs.count("the addon sync did not complete")

	for i := 3; i < 8; i++ {
		h.app.Step(t.Context(), t0.Add(time.Duration(i)*time.Second))
	}
	if n := logs.count("no longer accepts this device"); n != 1 {
		t.Errorf("the sign-out was logged %d times after more ticks", n)
	}
	if n := logs.count("the addon sync did not complete"); n != syncFails {
		t.Errorf("the sync kept running while signed out: %d failures, was %d", n, syncFails)
	}
	if n := queued(t, h.dirs.Queue); n != held {
		t.Errorf("the queue went from %d files to %d while signed out", held, n)
	}

	// Signing in again clears it, and the queued night goes up.
	h.srv.RestoreToken()
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("signing in again failed")
	}
	s = h.app.Snapshot()
	if s.SignedOut != nil || !s.Paired || s.Headline.Key == "signed_out" {
		t.Fatalf("after the sign-in: signed_out %+v, paired %v, headline %q", s.SignedOut, s.Paired, s.Headline.Key)
	}
	h.app.Step(t.Context(), t0.Add(10*time.Second))
	var fights int
	for _, e := range h.srv.Order() {
		if strings.HasPrefix(e, "fight ") {
			fights++
		}
	}
	if fights == 0 {
		t.Fatalf("the queued night did not upload after the sign-in: %v", h.srv.Order())
	}
}

func TestARevokedInboxAnswerAloneSignsTheDeviceOut(t *testing.T) {
	var logs lockedBuf
	h := newHarnessWith(t, nil, func(o *app.Options) {
		o.Log = slog.New(slog.NewTextHandler(&logs, nil))
	})
	h.call(http.MethodPost, h.url("/api/pair"), map[string]string{"code": fakeapi.PairCode})
	h.srv.RevokeToken()
	h.app.Step(t.Context(), t0)
	if s := h.app.Snapshot(); s.SignedOut == nil || s.Headline.Key != "signed_out" {
		t.Fatalf("snapshot = %+v", s.SignedOut)
	}
	if logs.count("no longer accepts this device") != 1 {
		t.Error("the sign-out was not logged once")
	}
}

func TestForgettingADeviceIsNotBeingSignedOutBySomeoneElse(t *testing.T) {
	h := newHarness(t)
	base := h.url("")
	h.call(http.MethodPost, base+"/api/pair", map[string]string{"code": fakeapi.PairCode})
	h.srv.RevokeToken()
	h.call(http.MethodPost, base+"/api/unpair", map[string]string{})
	if s := h.app.Snapshot(); s.SignedOut != nil || s.Headline.Key != "not_connected" {
		t.Fatalf("a deliberate forget reads signed out: %+v %q", s.SignedOut, s.Headline.Key)
	}
}

// A heartbeat the site refuses signs the device out through the same path
// as an upload does, and stops; signing in again starts it once more.
func TestARefusedHeartbeatSignsTheDeviceOutAndASignInResumesIt(t *testing.T) {
	var logs lockedBuf
	h := newHarnessWith(t, nil, func(o *app.Options) {
		o.Log = slog.New(slog.NewTextHandler(&logs, &slog.HandlerOptions{Level: slog.LevelDebug}))
		o.HeartbeatEvery = 20 * time.Millisecond
	})
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("pairing failed")
	}
	h.app.Step(t.Context(), t0) // prime the tail over the empty folder
	if err := os.WriteFile(h.log,
		[]byte(fixture.Header+fixture.Zone+fixture.Encounter(0)+fixture.Heartbeat(0)), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second)) // the report opens and is created
	wait := func(what string, cond func() bool) {
		t.Helper()
		for deadline := time.Now().Add(2 * time.Second); !cond(); time.Sleep(5 * time.Millisecond) {
			if time.Now().After(deadline) {
				t.Fatalf("timed out waiting for %s", what)
			}
		}
	}
	wait("a heartbeat", func() bool { return h.srv.Heartbeats() > 0 })
	h.srv.RevokeToken()
	wait("the refusal", func() bool {
		h.app.Step(t.Context(), t0.Add(2*time.Second))
		return h.app.Snapshot().SignedOut != nil
	})
	n := h.srv.HeartbeatRequests()
	time.Sleep(100 * time.Millisecond)
	if got := h.srv.HeartbeatRequests(); got != n {
		t.Fatalf("heartbeats went from %d to %d while signed out", n, got)
	}
	h.srv.RestoreToken()
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("signing in again failed")
	}
	wait("a heartbeat after the sign-in", func() bool { return h.srv.HeartbeatRequests() > n })
}
