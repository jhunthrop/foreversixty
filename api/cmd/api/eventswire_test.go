package main

import (
	"context"
	"log/slog"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
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
	if got := pushSenders(context.Background(), config.PushConfig{Enabled: true}, slog.Default()); len(got) != 0 {
		t.Fatalf("senders %v", got)
	}
	bad := config.PushConfig{Enabled: true, APNsKeyP8: []byte("not-a-key"), APNsKeyID: "K", APNsTeamID: "T", APNsTopic: "t",
		FCMServiceAccountJSON: []byte("{}")}
	if got := pushSenders(context.Background(), bad, slog.Default()); len(got) != 0 {
		t.Fatalf("bad keys built senders %v", got)
	}
}
