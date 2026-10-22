package app_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
)

func TestAStaleTokenOnAFirstRunReadsAsAFirstRun(t *testing.T) {
	var logs lockedBuf
	h := newHarnessWith(t, nil, func(o *app.Options) {
		o.Log = slog.New(slog.NewTextHandler(&logs, &slog.HandlerOptions{Level: slog.LevelDebug}))
	})
	// Windows still holds a token from an earlier install; this home never connected.
	if err := h.secret.SetToken("old-token-from-another-install"); err != nil {
		t.Fatal(err)
	}
	h.srv.RevokeToken()
	if !h.app.Snapshot().Paired {
		t.Fatal("the leftover token should read as held before the first call")
	}
	for i := 0; i < 3; i++ {
		h.app.Step(t.Context(), t0.Add(time.Duration(i)*time.Second))
	}
	s := h.app.Snapshot()
	if s.SignedOut != nil || s.Paired || s.Headline.Key == "signed_out" {
		t.Fatalf("a first run reads signed out: %+v, paired %v, headline %q", s.SignedOut, s.Paired, s.Headline.Key)
	}
	if s.Headline.Title == app.CopySignedOutTitle {
		t.Errorf("headline = %+v", s.Headline)
	}
	if n := logs.count("no longer accepts this device"); n != 0 {
		t.Errorf("the lapsed sign-in was logged %d times on a first run", n)
	}
	if n := logs.count("starting as a first run"); n != 1 {
		t.Errorf("the first-run line was logged %d times; want once", n)
	}
	if stored, _ := config.Load(h.dirs.ConfigFile()); stored.DeviceToken != "" {
		t.Errorf("the refused token is still stored")
	}
}

func TestAStaleTokenInAHomeThatWasConnectedStillReadsSignedOut(t *testing.T) {
	h := newHarnessWith(t, nil, func(o *app.Options) { o.Config.AccountLabel = "Thrall#1234" })
	if err := h.secret.SetToken("old-token"); err != nil {
		t.Fatal(err)
	}
	h.srv.RevokeToken()
	h.app.Step(t.Context(), t0)
	h.app.Step(t.Context(), t0.Add(time.Second))
	if s := h.app.Snapshot(); s.SignedOut == nil || s.Headline.Key != "signed_out" {
		t.Fatalf("a connected home with a refused token: %+v, headline %q", s.SignedOut, s.Headline.Key)
	}
}

func TestAFailingStepIsLoggedOnceNotEverySecond(t *testing.T) {
	var logs lockedBuf
	h := newHarnessWith(t, nil, func(o *app.Options) {
		o.Log = slog.New(slog.NewTextHandler(&logs, &slog.HandlerOptions{Level: slog.LevelDebug}))
	})
	if code, _ := h.call("POST", h.url("/api/pair"), map[string]string{"code": fakeapi.PairCode}); code != 200 {
		t.Fatal("pairing failed")
	}
	h.srv.Offline(true)
	for i := 0; i < 10; i++ {
		h.app.Step(t.Context(), t0.Add(time.Duration(i)*time.Second))
	}
	if n := logs.count("the addon sync did not complete"); n != 1 {
		t.Errorf("the same sync failure was logged %d times in 10 ticks", n)
	}
}
