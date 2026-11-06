package guilds

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// A rogue synced with no talent spent and no fight logged still reads as a rogue on the roster:
// the class is the export's own class field, whatever its talents (guild round 3, item 6).
func TestRosterClassComesFromTheExportWithNoTalentsSpent(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := t.Context()
	gid := seedGuild(t, h.pool, "Lifecycle Test")
	uid := seedUser(t, h.pool, "driftwood@example.com")
	key := "us/pvp/driftwood"
	seedCharacter(t, h.pool, gid, uid, key, "member", true)
	if _, err := h.pool.Exec(ctx, `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, 'us', 'pvp', 'Driftwood', $3, now(), now())`, key, uid,
		"FS1:1.60.1.70790:rogue:human:0/0/0:|level=60|professions=skinning,leatherworking|guild=Lifecycle%20Test:3"); err != nil {
		t.Fatal(err)
	}
	recomputeMembership(t, h.pool, gid, uid)
	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
	var view HomeView
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), ""), &view)
	if len(view.Roster) != 1 || view.Roster[0].Class == nil || *view.Roster[0].Class != "rogue" {
		t.Fatalf("roster = %+v, want Driftwood a rogue", view.Roster)
	}
	if view.Roster[0].Spec != nil {
		t.Fatalf("spec = %q, want none: no talent spent and no fight to name one", *view.Roster[0].Spec)
	}
	found := false
	for _, e := range view.Members {
		if e.Name == "Driftwood" {
			found = true
			if e.Class == nil || *e.Class != "rogue" {
				t.Fatalf("members entry = %+v, want class rogue", e)
			}
		}
	}
	if !found {
		t.Fatalf("members = %+v, want a Driftwood entry", view.Members)
	}
}
