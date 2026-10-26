// api/cmd/api/main_test.go
package main

import (
	"io"
	"log/slog"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// TestNewReportsServiceWiresGuilds is the wiring guard item 3 of the
// third security review response asked for: dropping Guilds from the
// reports.Service literal main wires would silently disable D's
// contested-and-frozen-claim report-edit freeze rather than failing
// loudly, since reports.Service.Guilds is deliberately nil-safe
// everywhere else (every test harness that does not care about claim
// disputes leaves it nil on purpose). This asserts the one call site
// that matters for a real deployment always populates it. No database
// connection is exercised - newReportsService only assembles the
// struct.
func TestNewReportsServiceWiresGuilds(t *testing.T) {
	svc := newReportsService(&reports.Store{}, &auth.Store{}, &guilds.Store{}, &rankings.Store{}, &guilds.Service{}, config.Config{}, nil)
	if svc.Guilds == nil {
		t.Fatal("reports.Service.Guilds is nil: the contested-claim report-edit freeze (D) is silently disabled")
	}
	if svc.Completed == nil {
		t.Fatal("reports.Service.Completed is nil: a patch that sets a complete report's guild tells no one")
	}
}

// The parse-report job tells guilds of every upload it completes; a Deps without the hook would
// leave every uploaded raid night unmatched with nothing failing. No database connection is
// exercised - newParseDeps only assembles the struct.
func TestNewParseDepsWiresTheCompletionHook(t *testing.T) {
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	deps := newParseDeps(config.Config{}, nil, nil, &trees.Data{}, quiet)
	g, ok := deps.Completed.(*guilds.Service)
	if !ok || g == nil || g.Store == nil {
		t.Fatalf("parse.Deps.Completed = %#v, want a guilds.Service with a store", deps.Completed)
	}
}
