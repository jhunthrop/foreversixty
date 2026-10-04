// api/internal/reports/guild_reports_test.go
package reports

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// seedGuild inserts a guild row and returns its id.
func (h *harness) seedGuild(name, ruleset string) int64 {
	h.t.Helper()
	var id int64
	if err := h.store.Pool.QueryRow(h.t.Context(),
		`insert into guilds (region, ruleset, name) values ('us', $1, $2) returning id`,
		ruleset, name).Scan(&id); err != nil {
		h.t.Fatal(err)
	}
	return id
}

// addGuildMember verifies userID as a member of guildID at the given
// rank, inserting the user row first if it is not already there.
func (h *harness) addGuildMember(guildID, userID int64, rank string) {
	h.t.Helper()
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`insert into users (id, email) values ($1, $2) on conflict (id) do nothing`,
		userID, fmt.Sprintf("u%d@example.com", userID)); err != nil {
		h.t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`insert into guild_members (guild_id, user_id, rank, verified_at) values ($1, $2, $3, now())`,
		guildID, userID, rank); err != nil {
		h.t.Fatal(err)
	}
}

// attachReportToGuild gives an existing report a guild_id directly,
// bypassing PATCH so the test does not also need officer standing.
func (h *harness) attachReportToGuild(reportID string, guildID int64) {
	h.t.Helper()
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`update reports set guild_id = $2 where id = $1`, reportID, guildID); err != nil {
		h.t.Fatal(err)
	}
}

// TestGuildReportsMemberSeesGuildAndPublicRowsNewestFirst covers the
// member path of GET /v1/guilds/{id}/reports (design/specs/
// 2026-10-04-logs-landing.md §4.C.1): a verified member sees the
// guild's own guild-visibility reports plus its public ones, newest
// first, and the response is shaped exactly like MinePage.
func TestGuildReportsMemberSeesGuildAndPublicRowsNewestFirst(t *testing.T) {
	h := newHarness(t)
	guildID := h.seedGuild("Forever", "hardcore")

	now := time.Now().UTC()
	// Oldest to newest, so sorting is actually exercised.
	h.seedReport("night1", "Night one", "Molten Core", GuildTo, StatusComplete, now.Add(-2*time.Hour))
	h.attachReportToGuild("night1", guildID)
	h.seedReport("night2", "Night two", "Molten Core", Public, StatusComplete, now.Add(-1*time.Hour))
	h.attachReportToGuild("night2", guildID)
	// Not this guild's: must never appear.
	h.seedReport("other1", "Someone else's", "Molten Core", Public, StatusComplete, now)
	// This guild's, but private: must never appear even to a member.
	h.seedReport("night3priv", "Private night", "Molten Core", Private, StatusComplete, now)
	h.attachReportToGuild("night3priv", guildID)
	// This guild's, unlisted: must never appear - only guild and public do.
	h.seedReport("night4unl", "Unlisted night", "Molten Core", Unlisted, StatusComplete, now)
	h.attachReportToGuild("night4unl", guildID)

	member := h.owner + 1
	h.addGuildMember(guildID, member, "member")
	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}

	res := h.do(http.MethodGet, "/v1/guilds/"+strconv.FormatInt(guildID, 10)+"/reports", "", nil)
	var page MinePage
	h.data(res, &page)

	if len(page.Rows) != 2 {
		t.Fatalf("rows = %d, want 2 (guild + public only): %+v", len(page.Rows), page.Rows)
	}
	if page.Rows[0].ID != "night2" || page.Rows[1].ID != "night1" {
		t.Fatalf("rows = %v, want night2 then night1 (newest first)", []string{page.Rows[0].ID, page.Rows[1].ID})
	}
	if page.Total != 2 || page.Page != 1 || page.PerPage != MinePerPage {
		t.Fatalf("page = %+v, want total 2, page 1, per_page %d", page, MinePerPage)
	}
}

// TestGuildReportsNonMemberGets404 is the guild-visibility rule's own
// shape, same as a stranger reading a private report: a signed-in
// caller with no standing in the guild never learns whether it exists.
func TestGuildReportsNonMemberGets404(t *testing.T) {
	h := newHarness(t)
	guildID := h.seedGuild("Walled Off", "normal")
	h.seedReport("walled1", "Walled night", "Molten Core", GuildTo, StatusComplete, time.Now().UTC())
	h.attachReportToGuild("walled1", guildID)

	stranger := h.owner + 1
	h.actor = auth.Actor{UserID: stranger, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, "/v1/guilds/"+strconv.FormatInt(guildID, 10)+"/reports", "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("a non-member sees %d, want 404", res.StatusCode)
	}
}

// TestGuildReportsAnonymousGets401 covers the signed-out case: no
// session at all is refused before the guild is even looked up.
func TestGuildReportsAnonymousGets401(t *testing.T) {
	h := newHarness(t)
	guildID := h.seedGuild("No Session", "normal")

	h.actor = auth.Actor{}
	res := h.do(http.MethodGet, "/v1/guilds/"+strconv.FormatInt(guildID, 10)+"/reports", "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("an anonymous caller sees %d, want 401", res.StatusCode)
	}
}

// TestGuildReportsEmptyGuildReturnsEmptyRows covers a verified member
// of a guild that has logged nothing yet: the rows array is empty, not
// null or an error.
func TestGuildReportsEmptyGuildReturnsEmptyRows(t *testing.T) {
	h := newHarness(t)
	guildID := h.seedGuild("Fresh Guild", "rp")
	member := h.owner + 1
	h.addGuildMember(guildID, member, "officer")
	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}

	res := h.do(http.MethodGet, "/v1/guilds/"+strconv.FormatInt(guildID, 10)+"/reports", "", nil)
	var page MinePage
	h.data(res, &page)
	if page.Rows == nil || len(page.Rows) != 0 {
		t.Fatalf("rows = %#v, want an empty, non-nil array", page.Rows)
	}
	if page.Total != 0 {
		t.Fatalf("total = %d, want 0", page.Total)
	}
}

// TestGuildReportsRejectsAnInvalidGuildID guards the path parameter the
// same way a malformed report id already answers 404 rather than 500:
// a non-numeric id can never match a guild, so it is refused the same
// way an unknown one is.
func TestGuildReportsRejectsAnInvalidGuildID(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	res := h.do(http.MethodGet, "/v1/guilds/not-a-number/reports", "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", res.StatusCode)
	}
}
