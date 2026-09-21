// api/cmd/api/main_test.go
package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
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
	svc := newReportsService(&reports.Store{}, &auth.Store{}, &guilds.Store{}, &rankings.Store{}, config.Config{}, nil)
	if svc.Guilds == nil {
		t.Fatal("reports.Service.Guilds is nil: the contested-claim report-edit freeze (D) is silently disabled")
	}
}
