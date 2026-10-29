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

type fakeSchedule struct {
	run   guilds.ScheduleRun
	err   error
	order *[]string
}

func (f fakeSchedule) RunRaidSchedule(context.Context, time.Time) (guilds.ScheduleRun, error) {
	*f.order = append(*f.order, "schedule")
	return f.run, f.err
}

type fakeDrain struct {
	order *[]string
	err   error
}

func (f *fakeDrain) DrainOnce(context.Context) error {
	*f.order = append(*f.order, "drain")
	return f.err
}

func TestRaidPostsRunsTheScheduleThenDrains(t *testing.T) {
	var order []string
	var buf bytes.Buffer
	log := slog.New(slog.NewTextHandler(&buf, nil))
	svc := raidPostsServices{
		Guilds:  fakeSchedule{run: guilds.ScheduleRun{Posted: 1, Reminded: 2}, order: &order},
		Drainer: &fakeDrain{order: &order},
	}
	if err := raidPosts(context.Background(), log, svc, time.Now()); err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(order, ","); got != "schedule,drain,drain,drain,drain,drain" {
		t.Fatalf("order %s", got)
	}
	out := buf.String()
	for _, want := range []string{"posted=1", "reminded=2", "failed=0"} {
		if !strings.Contains(out, want) {
			t.Fatalf("log %q lacks %s", out, want)
		}
	}
}

func TestRaidPostsWithoutTheBotOnlyRunsTheSchedule(t *testing.T) {
	var order []string
	svc := raidPostsServices{Guilds: fakeSchedule{order: &order}}
	if err := raidPosts(context.Background(), slog.Default(), svc, time.Now()); err != nil {
		t.Fatal(err)
	}
	if len(order) != 1 {
		t.Fatalf("order %v", order)
	}
}

func TestRaidPostsFailsOnAScheduleErrorButStillDrains(t *testing.T) {
	var order []string
	boom := errors.New("boom")
	svc := raidPostsServices{Guilds: fakeSchedule{err: boom, order: &order}, Drainer: &fakeDrain{order: &order}}
	err := raidPosts(context.Background(), slog.Default(), svc, time.Now())
	if !errors.Is(err, boom) {
		t.Fatalf("err %v", err)
	}
	if len(order) != 1+raidPostsDrainPasses {
		t.Fatalf("order %v", order)
	}
}

func TestRaidPostsStopsDrainingOnAnError(t *testing.T) {
	var order []string
	boom := errors.New("drain boom")
	svc := raidPostsServices{Guilds: fakeSchedule{order: &order}, Drainer: &fakeDrain{order: &order, err: boom}}
	if err := raidPosts(context.Background(), slog.Default(), svc, time.Now()); !errors.Is(err, boom) {
		t.Fatalf("err %v", err)
	}
	if len(order) != 2 {
		t.Fatalf("order %v", order)
	}
}

func TestNewRaidPostsServicesHasNoDrainerWithoutTheBot(t *testing.T) {
	svc, err := newRaidPostsServices(config.Config{}, nil, slog.Default())
	if err != nil {
		t.Fatal(err)
	}
	if svc.Guilds == nil || svc.Drainer != nil {
		t.Fatalf("%+v", svc)
	}
}
