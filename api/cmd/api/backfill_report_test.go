// api/cmd/api/backfill_report_test.go
package main

import (
	"io"
	"log/slog"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func TestParseBackfillFlagsRequiresAReportAndDefaultsToADryRun(t *testing.T) {
	if _, err := parseBackfillFlags(nil); err == nil || !strings.Contains(err.Error(), "--report") {
		t.Fatalf("no --report should refuse, got %v", err)
	}
	if _, err := parseBackfillFlags([]string{"--apply"}); err == nil {
		t.Fatal("--apply without --report should refuse")
	}
	if _, err := parseBackfillFlags([]string{"--report"}); err == nil {
		t.Fatal("--report without a value should refuse")
	}
	if _, err := parseBackfillFlags([]string{"--report", "3sdrcw", "--all"}); err == nil {
		t.Fatal("an unknown flag should refuse")
	}
	f, err := parseBackfillFlags([]string{"--report", "3sdrcw"})
	if err != nil || f.report != "3sdrcw" || f.apply {
		t.Fatalf("flags = %+v, %v; want a dry run of 3sdrcw", f, err)
	}
	f, err = parseBackfillFlags([]string{"--report", "3sdrcw", "--apply"})
	if err != nil || !f.apply {
		t.Fatalf("flags = %+v, %v; want apply", f, err)
	}
}

// The backfill's wiring carries no notifier: nothing announces a night that is already over.
func TestBackfillDepsAreSilent(t *testing.T) {
	deps := newBackfillDeps(config.Config{}, nil, nil, &trees.Data{}, slog.New(slog.NewTextHandler(io.Discard, nil)))
	if deps.Notify != nil {
		t.Fatal("the backfill must not announce")
	}
	if deps.Completed == nil {
		t.Fatal("the backfill must still run the completion hook")
	}
}

func TestParseBackfillFlagsValidatesTheLayout(t *testing.T) {
	if _, err := parseBackfillFlags([]string{"--report", "4q5el2", "--layout", "inferred"}); err == nil || !strings.Contains(err.Error(), "retail-v22") {
		t.Fatalf("an unknown layout should refuse and name the known ones, got %v", err)
	}
	if _, err := parseBackfillFlags([]string{"--report", "4q5el2", "--layout"}); err == nil {
		t.Fatal("--layout without a value should refuse")
	}
	f, err := parseBackfillFlags([]string{"--report", "4q5el2", "--layout", "retail-v22", "--apply"})
	if err != nil || f.layout != "retail-v22" || !f.apply {
		t.Fatalf("flags = %+v, %v", f, err)
	}
}
