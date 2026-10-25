package guilds

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

func homeViewFor(t *testing.T, h *httpHarness, guildID, viewerID int64) HomeView {
	t.Helper()
	h.actor = auth.Actor{UserID: viewerID, Role: "user", Method: "session"}
	var view HomeView
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", guildID), ""), &view)
	return view
}

func TestHomeViewerOfficerTools(t *testing.T) {
	const grantSQL = `insert into entitlements (guild_id, plan, source, status, current_period_end, features)
		values ($1,'guild','grant','active', now()+$2::interval,'{officer_views}')`
	tests := []struct {
		name     string
		interval string // "" seeds no grant
		rank     string
		want     bool
	}{
		{"officer with a live grant", "30 days", "officer", true},
		{"officer after the grant expired", "-1 day", "officer", false},
		{"officer with no grant", "", "officer", false},
		{"member with a live grant", "30 days", "member", false},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := newHTTPHarness(t)
			h.svc.Entitlements = &entitlements.Store{Pool: h.pool}
			gid := seedGuild(t, h.pool, fmt.Sprintf("OfficerTools%d", i))
			viewer := seedUser(t, h.pool, fmt.Sprintf("ot-viewer-%d@example.com", i))
			seedVerifiedMember(t, h, gid, viewer, fmt.Sprintf("otviewer%d", i), tc.rank)
			if tc.interval != "" {
				if _, err := h.pool.Exec(t.Context(), grantSQL, gid, tc.interval); err != nil {
					t.Fatal(err)
				}
			}
			if got := homeViewFor(t, h, gid, viewer).Viewer.OfficerTools; got != tc.want {
				t.Fatalf("viewer.officer_tools = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestHomeViewerOfficerToolsFalseWithoutEntitlements(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "OfficerToolsNil")
	viewer := seedUser(t, h.pool, "ot-nil@example.com")
	seedVerifiedMember(t, h, gid, viewer, "otnil", "officer")
	if homeViewFor(t, h, gid, viewer).Viewer.OfficerTools {
		t.Fatal("nil Entitlements must read officer_tools false")
	}
}
