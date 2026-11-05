// api/internal/guilds/readiness_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// TestReadinessRequiresMembership checks the member/officer-only gate (design spec §4.0:
// "Not shown" to the public).
func TestReadinessRequiresMembership(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	stranger := seedUser(t, h.pool, "readiness-stranger@example.com")
	h.actor = auth.Actor{UserID: stranger, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/readiness", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a non-member reading readiness = %d, want 403", res.StatusCode)
	}
}

// TestReadinessOnlyShowsVerifiedRowsAndGatesNudgeByRole checks: an unverified character
// never appears, a member sees a verified row with no nudge_text, and an officer sees the
// same row with nudge_text set.
func TestReadinessOnlyShowsVerifiedRowsAndGatesNudgeByRole(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")

	member := seedUser(t, h.pool, "readiness-member@example.com")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/readinessmember", "member", true)
	seedExport(t, h.pool, member, "us/hardcore/readinessmember", "us", "hardcore", "readinessmember")
	recomputeMembership(t, h.pool, gid, member)

	officer := seedUser(t, h.pool, "readiness-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/readinessofficer", "officer", true)
	seedExport(t, h.pool, officer, "us/hardcore/readinessofficer", "us", "hardcore", "readinessofficer")
	recomputeMembership(t, h.pool, gid, officer)

	unverified := seedUser(t, h.pool, "readiness-unverified@example.com")
	seedCharacter(t, h.pool, gid, unverified, "us/hardcore/readinessunverified", "member", false)
	seedExport(t, h.pool, unverified, "us/hardcore/readinessunverified", "us", "hardcore", "readinessunverified")
	recomputeMembership(t, h.pool, gid, unverified)

	h.actor = auth.Actor{UserID: unverified, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/readiness", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a member waiting for approval reading readiness = %d, want 403", res.StatusCode)
	}

	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/readiness", gid), "")
	var view ReadinessView
	h.data(res, &view)
	if len(view.Rows) != 2 {
		t.Fatalf("rows = %+v, want exactly the 2 verified rows", view.Rows)
	}
	for _, row := range view.Rows {
		if row.NudgeText != nil {
			t.Fatalf("row %s carries nudge_text for a plain member, want omitted", row.CharacterKey)
		}
	}

	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/readiness", gid), "")
	h.data(res, &view)
	for _, row := range view.Rows {
		if row.NudgeText == nil {
			t.Fatalf("row %s missing nudge_text for an officer", row.CharacterKey)
		}
	}
}

// TestReadinessConsentGatesGearAndConsumables checks a roster-consent character's gear_gap
// is nil and consumables/enchants read as not-checked/unknown, while a gear_bags-consent
// character gets the full picture.
func TestReadinessConsentGatesGearAndConsumables(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")

	rosterOnly := seedUser(t, h.pool, "readiness-roster-only@example.com")
	seedCharacter(t, h.pool, gid, rosterOnly, "us/hardcore/readinessrosteronly", "member", true)
	seedExport(t, h.pool, rosterOnly, "us/hardcore/readinessrosteronly", "us", "hardcore", "readinessrosteronly")
	recomputeMembership(t, h.pool, gid, rosterOnly)
	if _, err := h.pool.Exec(ctx,
		`update guild_members set consent = 'roster' where guild_id = $1 and user_id = $2`, gid, rosterOnly); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{UserID: rosterOnly, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/readiness", gid), "")
	var view ReadinessView
	h.data(res, &view)
	if len(view.Rows) != 1 {
		t.Fatalf("rows = %+v, want 1", view.Rows)
	}
	row := view.Rows[0]
	if row.GearGap != nil {
		t.Errorf("gearGap = %+v, want nil without gear consent", row.GearGap)
	}
	if row.Enchants.Checked {
		t.Error("enchants.checked = true, want false without gear consent")
	}
	if row.Consumables.State != "unknown" {
		t.Errorf("consumables.state = %q, want unknown", row.Consumables.State)
	}
}

// seedLevelExport writes key's export at level with no talents spent and no gear.
func seedLevelExport(t *testing.T, h *httpHarness, userID int64, key, name string, level int) {
	t.Helper()
	export := fmt.Sprintf("FS1:1.60.1.70009:warrior:human:0/0/0:|level=%d", level)
	if _, err := h.pool.Exec(context.Background(), `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, 'us', 'hardcore', $3, $4, now(), now())`, key, userID, name, export); err != nil {
		t.Fatal(err)
	}
}

// A leveling character is listed, last, with its level and nothing graded; an officer's nudge
// is for graded rows only and names no day when no raid is planned (week-one spec §3.5, §4.2).
func TestReadinessLeavesLevelingCharactersUngraded(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")

	officer := seedUser(t, h.pool, "readiness-sixty@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/readinesssixty", "officer", true)
	seedLevelExport(t, h, officer, "us/hardcore/readinesssixty", "readinesssixty", 60)
	recomputeMembership(t, h.pool, gid, officer)

	leveler := seedUser(t, h.pool, "readiness-leveler@example.com")
	seedCharacter(t, h.pool, gid, leveler, "us/hardcore/readinessleveler", "member", true)
	seedLevelExport(t, h, leveler, "us/hardcore/readinessleveler", "readinessleveler", 52)
	recomputeMembership(t, h.pool, gid, leveler)

	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/readiness", gid), "")
	var view ReadinessView
	h.data(res, &view)
	if len(view.Rows) != 2 {
		t.Fatalf("rows = %+v, want 2", view.Rows)
	}
	sixty, leveling := view.Rows[0], view.Rows[1]
	if leveling.CharacterKey != "us/hardcore/readinessleveler" || leveling.Level == nil || *leveling.Level != 52 {
		t.Fatalf("last row = %+v, want the level-52 character", leveling)
	}
	if leveling.Failing != 0 || leveling.TalentPointsUnspent != 0 || leveling.GearGap != nil ||
		leveling.Enchants.Checked || leveling.Consumables.State != "unknown" || leveling.NudgeText != nil {
		t.Fatalf("leveling row = %+v, want nothing graded and no nudge", leveling)
	}
	if sixty.Failing == 0 || sixty.NudgeText == nil || strings.Contains(*sixty.NudgeText, "before") {
		t.Fatalf("level-60 row = %+v, want graded with a nudge naming no day", sixty)
	}
}

// TestReadinessSortsWorstFirst checks the readinessScore ordering: a character failing
// more checks always sorts above one failing fewer.
func TestReadinessSortsWorstFirst(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")

	clean := seedUser(t, h.pool, "readiness-clean@example.com")
	seedCharacter(t, h.pool, gid, clean, "us/hardcore/readinessclean", "member", true)
	// A clean export: no gear at all means the enchant check finds nothing equipped to
	// flag, and level 60 fully-spent talents means no unspent-points failure either.
	if _, err := h.pool.Exec(ctx, `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, 'us', 'hardcore', 'readinessclean', 'FS1:1.60.1.70009:warrior:human:555555555/33/0:|level=60', now(), now())`,
		"us/hardcore/readinessclean", clean); err != nil {
		t.Fatal(err)
	}
	recomputeMembership(t, h.pool, gid, clean)

	gapped := seedUser(t, h.pool, "readiness-gapped@example.com")
	seedCharacter(t, h.pool, gid, gapped, "us/hardcore/readinessgapped", "member", true)
	if _, err := h.pool.Exec(ctx, `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, 'us', 'hardcore', 'readinessgapped', 'FS1:1.60.1.70009:warrior:human:0/0/0:|level=60', now(), now())`,
		"us/hardcore/readinessgapped", gapped); err != nil {
		t.Fatal(err)
	}
	recomputeMembership(t, h.pool, gid, gapped)

	h.actor = auth.Actor{UserID: clean, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/readiness", gid), "")
	var view ReadinessView
	h.data(res, &view)
	if len(view.Rows) != 2 {
		t.Fatalf("rows = %+v, want 2", view.Rows)
	}
	// The unspent-talent-points character (0 points spent, vs. the clean one's full 51)
	// must sort first - it fails strictly more checks.
	if view.Rows[0].CharacterKey != "us/hardcore/readinessgapped" {
		t.Fatalf("rows[0] = %s, want the gapped character sorted worst-first; rows=%+v",
			view.Rows[0].CharacterKey, view.Rows)
	}
	if view.Rows[0].TalentPointsUnspent == 0 {
		t.Error("gapped character's talent_points_unspent = 0, want > 0")
	}
	if view.Rows[1].TalentPointsUnspent != 0 {
		t.Errorf("clean character's talent_points_unspent = %d, want 0", view.Rows[1].TalentPointsUnspent)
	}
}
