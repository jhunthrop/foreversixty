package reports

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestRecentCursorRoundTrips(t *testing.T) {
	want := time.Date(2026, 9, 21, 12, 30, 0, 123456789, time.UTC)
	got, ok := decodeRecentCursor(encodeRecentCursor(want, "abcdef123456"))
	if !ok || !got.CreatedAt.Equal(want) || got.ID != "abcdef123456" {
		t.Fatalf("round trip = %+v, %v; want %v, abcdef123456", got, ok, want)
	}
}

func TestRecentCursorRejectsGarbage(t *testing.T) {
	for _, bad := range []string{"", "not-base64!!", "aGVsbG8", encodeRecentCursor(time.Now(), "")} {
		if _, ok := decodeRecentCursor(bad); ok {
			t.Fatalf("decodeRecentCursor(%q) = ok, want rejected", bad)
		}
	}
}

// TestStoreRecentListsOnlyCompletePublicNewestFirst is the store-level
// test the spec calls for, in the style of the neighbouring report
// store tests: it seeds rows directly and reads the store back, with no
// HTTP layer between the assertion and the SQL.
func TestStoreRecentListsOnlyCompletePublicNewestFirst(t *testing.T) {
	h := newHarness(t)
	base := time.Now().UTC().Truncate(time.Second)
	h.seedReport("public00pass", "", "Blackrock Depths", Public, StatusComplete, base)
	h.seedReport("public00zone", "Progress night", "Molten Core", Public, StatusComplete, base.Add(time.Minute))
	h.seedReport("private000hd", "Hidden", "Onyxia's Lair", Private, StatusComplete, base.Add(2*time.Minute))
	h.seedReport("unlisted0001", "Unlisted", "Zul'Gurub", Unlisted, StatusComplete, base.Add(3*time.Minute))
	h.seedReport("stilllivepub", "Still running", "Blackwing Lair", Public, StatusLive, base.Add(4*time.Minute))

	rows, err := h.store.Recent(h.t.Context(), nil, RecentPerPage)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 {
		t.Fatalf("rows = %d, want the 2 complete public reports: %+v", len(rows), rows)
	}
	if rows[0].ID != "public00zone" || rows[1].ID != "public00pass" {
		t.Fatalf("order = %+v, want newest first", rows)
	}
	if rows[1].Title != "Blackrock Depths" {
		t.Fatalf("title should fall back to zone when it is empty: %q", rows[1].Title)
	}
}

// TestStoreRecentBreaksTiesByID pins the keyset's second sort key: two
// reports created in the same instant must still page deterministically,
// and the cursor built from the first row must resume after it exactly.
func TestStoreRecentBreaksTiesByID(t *testing.T) {
	h := newHarness(t)
	same := time.Now().UTC().Truncate(time.Second)
	h.seedReport("tie-a-report", "A", "Zone", Public, StatusComplete, same)
	h.seedReport("tie-b-report", "B", "Zone", Public, StatusComplete, same)

	rows, err := h.store.Recent(h.t.Context(), nil, RecentPerPage)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 || rows[0].ID != "tie-b-report" || rows[1].ID != "tie-a-report" {
		t.Fatalf("rows = %+v, want id desc breaking the created_at tie", rows)
	}

	cursor := recentCursor{CreatedAt: rows[0].CreatedAt, ID: rows[0].ID}
	next, err := h.store.Recent(h.t.Context(), &cursor, RecentPerPage)
	if err != nil {
		t.Fatal(err)
	}
	if len(next) != 1 || next[0].ID != "tie-a-report" {
		t.Fatalf("next = %+v, want just tie-a-report", next)
	}
}

// TestRecentHandlerFiltersLinksAndCounts is the handler-level test the
// spec calls for: guild name, fight and kill counts, and the title
// fallback, all read back through GET /v1/reports/recent.
func TestRecentHandlerFiltersLinksAndCounts(t *testing.T) {
	h := newHarness(t)
	base := time.Now().UTC().Truncate(time.Second)
	h.seedReport("public00pass", "", "Blackrock Depths", Public, StatusComplete, base)
	h.seedReport("public00zone", "Progress night", "Molten Core", Public, StatusComplete, base.Add(time.Minute))
	h.seedReport("private000hd", "Hidden", "Onyxia's Lair", Private, StatusComplete, base.Add(2*time.Minute))

	var guildID int64
	if err := h.store.Pool.QueryRow(h.t.Context(),
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Forever') returning id`).
		Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`update reports set guild_id = $2 where id = $1`, "public00zone", guildID); err != nil {
		t.Fatal(err)
	}
	h.reportID = "public00zone"
	h.seedFight(1, 9001, 0, 500)  // an encounter kill
	h.seedFight(2, 0, 500, 900)   // trash, not counted as a kill

	res := h.do(http.MethodGet, "/v1/reports/recent", "", nil)
	var page RecentPage
	h.data(res, &page)
	if len(page.Rows) != 2 {
		t.Fatalf("rows = %d, want 2: %+v", len(page.Rows), page.Rows)
	}
	newest := page.Rows[0]
	if newest.ID != "public00zone" || newest.Title != "Progress night" {
		t.Fatalf("newest row = %+v", newest)
	}
	if newest.GuildName != "Forever" {
		t.Fatalf("guild_name = %q, want Forever", newest.GuildName)
	}
	if newest.FightCount != 2 || newest.KillCount != 1 {
		t.Fatalf("counts = %d fights, %d kills; want 2, 1", newest.FightCount, newest.KillCount)
	}
	if page.Rows[1].GuildName != "" {
		t.Fatalf("a report with no guild should carry no guild_name: %q", page.Rows[1].GuildName)
	}
	if page.NextCursor != "" {
		t.Fatalf("next_cursor = %q, want empty for a page short of the limit", page.NextCursor)
	}
}

// TestRecentHandlerPaginatesWithoutDuplicatesOrGaps proves the cursor
// the handler hands back actually resumes the feed: a page size of 10
// against 15 rows must be walked in exactly two pages covering every row
// once.
func TestRecentHandlerPaginatesWithoutDuplicatesOrGaps(t *testing.T) {
	h := newHarness(t)
	base := time.Now().UTC().Truncate(time.Second)
	const total = RecentPerPage + 5
	for i := range total {
		id := fmt.Sprintf("recentseed%02d", i)
		h.seedReport(id, fmt.Sprintf("Report %d", i), "Blackrock Depths", Public, StatusComplete,
			base.Add(time.Duration(i)*time.Minute))
	}

	res := h.do(http.MethodGet, "/v1/reports/recent", "", nil)
	var first RecentPage
	h.data(res, &first)
	if len(first.Rows) != RecentPerPage {
		t.Fatalf("first page rows = %d, want %d", len(first.Rows), RecentPerPage)
	}
	if first.NextCursor == "" {
		t.Fatal("a full page should carry a next_cursor")
	}

	res = h.do(http.MethodGet, "/v1/reports/recent?cursor="+first.NextCursor, "", nil)
	var second RecentPage
	h.data(res, &second)
	if len(second.Rows) != total-RecentPerPage {
		t.Fatalf("second page rows = %d, want %d", len(second.Rows), total-RecentPerPage)
	}
	if second.NextCursor != "" {
		t.Fatal("the last page should carry no next_cursor")
	}

	seen := map[string]bool{}
	for _, row := range append(first.Rows, second.Rows...) {
		if seen[row.ID] {
			t.Fatalf("report %s appeared on both pages", row.ID)
		}
		seen[row.ID] = true
	}
	if len(seen) != total {
		t.Fatalf("saw %d distinct reports across both pages, want %d", len(seen), total)
	}
}

func TestRecentRejectsAGarbageCursor(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodGet, "/v1/reports/recent?cursor=not-a-real-cursor", "", nil)
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", res.StatusCode)
	}
	if h.errorFields(res)["cursor"] == "" {
		t.Fatal("the 400 should name the cursor field")
	}
}
