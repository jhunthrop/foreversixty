// api/cmd/api/jobs_extra_test.go
package main

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"reflect"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
)

func TestRunJobTable_RunsTheRowThenExitsZero(t *testing.T) {
	var got []string
	table := []jobCommand{{"a", func(_ context.Context, _ *slog.Logger, args []string) error { got = args; return nil }}}
	var exits []int
	runJobTable(slog.Default(), table, []string{"a", "x"}, func(c int) { exits = append(exits, c) })
	if !reflect.DeepEqual(got, []string{"x"}) || !reflect.DeepEqual(exits, []int{0}) {
		t.Fatalf("run args %v, exits %v; want [x], [0]", got, exits)
	}
}

func TestRunJobTable_ErrorExitsOneAndLogsTheName(t *testing.T) {
	var buf bytes.Buffer
	table := []jobCommand{{"boom", func(context.Context, *slog.Logger, []string) error { return errors.New("bad") }}}
	var exits []int
	runJobTable(slog.New(slog.NewTextHandler(&buf, nil)), table, []string{"boom"}, func(c int) { exits = append(exits, c) })
	if !reflect.DeepEqual(exits, []int{1}) || !strings.Contains(buf.String(), "boom") || !strings.Contains(buf.String(), "bad") {
		t.Fatalf("exits %v, log %q; want [1] and a line naming boom", exits, buf.String())
	}
}

func TestRunJobTable_UnknownOrEmptyReturns(t *testing.T) {
	ran := false
	table := []jobCommand{{"a", func(context.Context, *slog.Logger, []string) error { ran = true; return nil }}}
	exit := func(int) { t.Fatal("exit must not be called") }
	runJobTable(slog.Default(), table, []string{"serve-nothing"}, exit)
	runJobTable(slog.Default(), table, nil, exit)
	if ran {
		t.Fatal("no row should have run")
	}
}

func TestExtraJobsRegistersTheRebuild(t *testing.T) {
	for _, j := range extraJobs {
		if j.name == rating.RebuildPreparationJobCommand {
			return
		}
	}
	t.Fatal("extraJobs lacks the rating-rebuild-preparation row")
}

func TestParseRebuildFlags(t *testing.T) {
	o, err := parseRebuildFlags([]string{"-dry-run", "-batch", "50", "-after", "abc:3"}, &bytes.Buffer{})
	if err != nil || !o.DryRun || o.BatchSize != 50 || o.After == nil || o.After.ReportID != "abc" || o.After.FightIndex != 3 {
		t.Fatalf("opts %+v, err %v", o, err)
	}
	for _, bad := range [][]string{{"-batch", "0"}, {"-after", "nocolon"}, {"-after", "r:x"}, {"-nope"}} {
		if _, err := parseRebuildFlags(bad, &bytes.Buffer{}); err == nil {
			t.Errorf("%v: want an error", bad)
		}
	}
}
