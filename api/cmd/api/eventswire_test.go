package main

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/push"
	"github.com/jhunthrop/foreversixty/api/internal/seats"
	"github.com/jhunthrop/foreversixty/api/internal/server"
)

func producerDeps() server.Deps {
	return server.Deps{Guilds: &guilds.Service{}, Addon: &addon.Service{}, Billing: &billing.Service{}, Seats: &seats.Service{}}
}

func TestWireEventsIsDarkWithoutTheFlag(t *testing.T) {
	deps := producerDeps()
	wireEvents(context.Background(), config.Config{}, nil, &deps, slog.Default())
	if deps.Notifications != nil || deps.PushTokens != nil || deps.Guilds.AppSignups || deps.Guilds.Events != nil {
		t.Fatalf("lit without the flag: %+v", deps)
	}
}

func TestWireEventsMountsTheRoutesWithTheFlag(t *testing.T) {
	deps := producerDeps()
	cfg := config.Config{Push: config.PushConfig{Enabled: true}}
	wireEvents(context.Background(), cfg, nil, &deps, slog.Default())
	if deps.Notifications == nil || deps.PushTokens == nil || !deps.Guilds.AppSignups {
		t.Fatal("the flag did not mount the routes")
	}
	// No sender keys and no Discord: no channel, so no producer publishes.
	if deps.Guilds.Events != nil {
		t.Fatal("producers wired with no channel")
	}
}

func TestConnectProducersReachesEveryProducer(t *testing.T) {
	deps := producerDeps()
	router := &events.Router{}
	connectProducers(&deps, router, "https://foreversixty.gg")
	if deps.Guilds.Events != router || deps.Addon.Events != router || deps.Billing.Events != router || deps.Seats.Events != router ||
		deps.Addon.PublicBaseURL != "https://foreversixty.gg" {
		t.Fatalf("%+v", deps)
	}
}

func TestPushSendersStayOffWithoutKeys(t *testing.T) {
	noADC := func(context.Context, ...string) (*google.Credentials, error) { return nil, errors.New("no ADC") }
	withFindCredentials(t, noADC)
	if got := pushSenders(context.Background(), config.PushConfig{Enabled: true}, slog.Default()); len(got) != 0 {
		t.Fatalf("senders %v", got)
	}
	bad := config.PushConfig{Enabled: true, APNsKeyP8: []byte("not-a-key"), APNsKeyID: "K", APNsTeamID: "T", APNsTopic: "t"}
	if got := pushSenders(context.Background(), bad, slog.Default()); len(got) != 0 {
		t.Fatalf("bad keys built senders %v", got)
	}
}

// With runtime credentials and no APNs key, Android sends and iOS is skipped: nothing fails.
func TestFCMSendsAsTheRuntimeIdentityWhileAPNsIsDark(t *testing.T) {
	adc := func(_ context.Context, scopes ...string) (*google.Credentials, error) {
		if len(scopes) != 1 || scopes[0] != push.FCMScope {
			t.Errorf("scopes %v", scopes)
		}
		return &google.Credentials{ProjectID: "foreversixty",
			TokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "x"})}, nil
	}
	withFindCredentials(t, adc)
	got := pushSenders(context.Background(), config.PushConfig{Enabled: true}, slog.Default())
	if got[push.Android] == nil || got[push.IOS] != nil || len(got) != 1 {
		t.Fatalf("senders %v", got)
	}
}

func withFindCredentials(t *testing.T, f func(context.Context, ...string) (*google.Credentials, error)) {
	t.Helper()
	saved := findCredentials
	findCredentials = f
	t.Cleanup(func() { findCredentials = saved })
}
