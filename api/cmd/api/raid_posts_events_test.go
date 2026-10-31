package main

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

type fakeEvents struct {
	order *[]string
	run   guilds.EventScheduleRun
	err   error
}

func (f fakeEvents) RunEventSchedule(context.Context, time.Time) (guilds.EventScheduleRun, error) {
	*f.order = append(*f.order, "events")
	return f.run, f.err
}

type fakePrune struct{ order *[]string }

func (f fakePrune) PruneDeliveries(context.Context, time.Time, int) (int64, error) {
	*f.order = append(*f.order, "prune")
	return 3, nil
}

func TestRaidPostsRunsTheEventScheduleAfterTheDrain(t *testing.T) {
	var order []string
	var buf bytes.Buffer
	svc := raidPostsServices{
		Guilds: fakeSchedule{order: &order}, Drainer: &fakeDrain{order: &order},
		Events: fakeEvents{order: &order, run: guilds.EventScheduleRun{Nudged: 2, Starting: 1}}, Prune: fakePrune{order: &order},
	}
	if err := raidPosts(context.Background(), slog.New(slog.NewTextHandler(&buf, nil)), svc, time.Now()); err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(order, ","); got != "schedule,drain,drain,drain,drain,drain,events,prune" {
		t.Fatalf("order %s", got)
	}
	for _, want := range []string{"nudged=2", "starting=1", "pruned=3"} {
		if !strings.Contains(buf.String(), want) {
			t.Fatalf("log %q lacks %s", buf.String(), want)
		}
	}
}

func TestAnEventScheduleFailureFailsTheJob(t *testing.T) {
	var order []string
	boom := errors.New("events boom")
	svc := raidPostsServices{Guilds: fakeSchedule{order: &order}, Events: fakeEvents{order: &order, err: boom}}
	if err := raidPosts(context.Background(), slog.Default(), svc, time.Now()); !errors.Is(err, boom) {
		t.Fatalf("err %v", err)
	}
}

func TestTheEventScheduleIsOffWithoutTheFlag(t *testing.T) {
	svc, err := newRaidPostsServices(config.Config{}, nil, slog.Default())
	if err != nil {
		t.Fatal(err)
	}
	if svc.Events != nil || svc.Prune != nil {
		t.Fatalf("%+v", svc)
	}
}
