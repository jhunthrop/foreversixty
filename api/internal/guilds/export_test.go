package guilds

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// exportFixture is a guild with one officer who is signed in to the harness.
type exportFixture struct {
	h       *httpHarness
	gid     int64
	officer int64
}

func newExportFixture(t *testing.T, guildName string) exportFixture {
	t.Helper()
	h := newHTTPHarness(t)
	h.svc.Entitlements = &entitlements.Store{Pool: h.pool}
	gid := seedGuild(t, h.pool, guildName)
	officer := seedUser(t, h.pool, fmt.Sprintf("export-officer-%d@example.com", gid))
	seedVerifiedMember(t, h, gid, officer, fmt.Sprintf("exportofficer%d", gid), "officer")
	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	return exportFixture{h: h, gid: gid, officer: officer}
}

func (f exportFixture) get(kind string) *http.Response {
	f.h.t.Helper()
	return f.h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/export/%s.csv", f.gid, kind), "")
}

func (f exportFixture) rows(kind string) [][]string {
	f.h.t.Helper()
	res := f.get(kind)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		f.h.t.Fatalf("%s export = %d, want 200", kind, res.StatusCode)
	}
	rows, err := csv.NewReader(res.Body).ReadAll()
	if err != nil {
		f.h.t.Fatal(err)
	}
	return rows
}

func (f exportFixture) seedRaid(title string) int64 {
	f.h.t.Helper()
	var id int64
	if err := f.h.pool.QueryRow(f.h.t.Context(), `
		insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, status, created_by)
		values ($1, 'onyxia', $2, '2026-10-20 19:00+00', 180, 'done', $3) returning id`,
		f.gid, title, f.officer).Scan(&id); err != nil {
		f.h.t.Fatal(err)
	}
	return id
}

func (f exportFixture) seedAttendance(raidID int64, key, name, class, outcome string) {
	f.h.t.Helper()
	if _, err := f.h.pool.Exec(f.h.t.Context(), `
		insert into guild_raid_attendance (raid_id, character_key, name, class, signed, outcome, excused, pulls_present, pulls_total)
		values ($1, $2, $3, $4, 'in', $5, true, 4, 5)`, raidID, key, name, class, outcome); err != nil {
		f.h.t.Fatal(err)
	}
}

func TestExportAttendance_HeaderAndRows(t *testing.T) {
	f := newExportFixture(t, "Export Rows")
	raid := f.seedRaid("Onyxia night")
	f.seedAttendance(raid, "us/hardcore/zed", "Zed", "mage", "present")
	f.seedAttendance(raid, "us/hardcore/amy", "Amy", "hunter", "late")

	rows := f.rows("attendance")
	want := [][]string{
		{"date", "raid", "character", "class", "signed", "outcome", "excused", "pulls_present", "pulls_total"},
		{"2026-10-20", "Onyxia night", "Amy", "hunter", "in", "late", "true", "4", "5"},
		{"2026-10-20", "Onyxia night", "Zed", "mage", "in", "present", "true", "4", "5"},
	}
	if fmt.Sprint(rows) != fmt.Sprint(want) {
		t.Fatalf("rows = %v, want %v", rows, want)
	}
}

func TestExportLoot_HeaderOnlyWhenEmpty(t *testing.T) {
	f := newExportFixture(t, "Export Empty")
	rows := f.rows("loot")
	want := "date,raid,boss,item_id,item,character"
	if len(rows) != 1 || strings.Join(rows[0], ",") != want {
		t.Fatalf("rows = %v, want header only %q", rows, want)
	}
	if got := f.rows("attendance"); len(got) != 1 {
		t.Fatalf("attendance rows = %v, want header only", got)
	}
}

func TestExportLoot_RowsResolveNames(t *testing.T) {
	f := newExportFixture(t, "Export Loot")
	raid := f.seedRaid("Onyxia night")
	if _, err := f.h.pool.Exec(t.Context(), `
		insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by, awarded_at, raid_id)
		values ($1, 1084, 16963, 'us/hardcore/zed', $2, '2026-10-20 21:00+00', $3)`, f.gid, f.officer, raid); err != nil {
		t.Fatal(err)
	}
	seedExport(t, f.h.pool, f.officer, "us/hardcore/zed", "us", "hardcore", "Zed")
	rows := f.rows("loot")
	if len(rows) != 2 || rows[1][0] != "2026-10-20" || rows[1][1] != "Onyxia night" ||
		rows[1][2] != encounterName || rows[1][3] != "16963" || rows[1][4] != "Helm of Wrath" || rows[1][5] != "Zed" {
		t.Fatalf("rows = %v", rows)
	}
}

func TestExport_LapsedGuildStillExports(t *testing.T) {
	f := newExportFixture(t, "Export Lapsed") // no entitlement row at all
	for _, kind := range []string{"attendance", "loot"} {
		res := f.get(kind)
		res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Fatalf("%s export on a lapsed guild = %d, want 200", kind, res.StatusCode)
		}
	}
}

func TestExport_MemberGets404(t *testing.T) {
	f := newExportFixture(t, "Export Member")
	member := seedUser(t, f.h.pool, "export-member@example.com")
	seedVerifiedMember(t, f.h, f.gid, member, "exportmember", "member")
	f.h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	for _, kind := range []string{"attendance", "loot"} {
		res := f.get(kind)
		res.Body.Close()
		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("%s export as member = %d, want 404", kind, res.StatusCode)
		}
	}
}

func TestExport_FormulaCellEscaped(t *testing.T) {
	f := newExportFixture(t, "Export Formula")
	raid := f.seedRaid("=HYPERLINK(\"x\")")
	f.seedAttendance(raid, "us/hardcore/evil", "@evil", "+mage", "present")
	row := f.rows("attendance")[1]
	if row[1] != "'=HYPERLINK(\"x\")" || row[2] != "'@evil" || row[3] != "'+mage" {
		t.Fatalf("row = %v, want formula cells apostrophe-prefixed", row)
	}
}

func TestExport_FilenameSlug(t *testing.T) {
	f := newExportFixture(t, "Olympus XXVII")
	res := f.get("loot")
	res.Body.Close()
	if got, want := res.Header.Get("Content-Disposition"), `attachment; filename="olympus-xxvii-loot.csv"`; got != want {
		t.Fatalf("Content-Disposition = %q, want %q", got, want)
	}
	if got := res.Header.Get("Content-Type"); got != "text/csv; charset=utf-8" {
		t.Fatalf("Content-Type = %q", got)
	}
}

func TestExport_ControlCharacterCellsEscaped(t *testing.T) {
	for _, lead := range []string{"\t", "\r"} {
		f := newExportFixture(t, "Export Control")
		raid := f.seedRaid("Onyxia night")
		f.seedAttendance(raid, "us/hardcore/ctl", lead+"x", "mage", "present")
		if got := f.rows("attendance")[1][2]; got != "'"+lead+"x" {
			t.Fatalf("cell = %q, want apostrophe prefix", got)
		}
	}
}

func TestExport_SpecialCharactersRoundTrip(t *testing.T) {
	f := newExportFixture(t, "Export Special")
	raid := f.seedRaid("Night, \"one\"\nsecond line")
	f.seedAttendance(raid, "us/hardcore/odd", "A,B \"C\"\nD", "mage", "present")
	row := f.rows("attendance")[1]
	if row[1] != "Night, \"one\"\nsecond line" || row[2] != "A,B \"C\"\nD" {
		t.Fatalf("row = %q", row)
	}
}

func TestExport_NonASCIINameFallsBackToGuildSlug(t *testing.T) {
	f := newExportFixture(t, "Олимп")
	res := f.get("loot")
	res.Body.Close()
	if got, want := res.Header.Get("Content-Disposition"), `attachment; filename="guild-loot.csv"`; got != want {
		t.Fatalf("Content-Disposition = %q, want %q", got, want)
	}
}

func TestExport_NonOfficersRefused(t *testing.T) {
	f := newExportFixture(t, "Export Refused")
	unverified := seedUser(t, f.h.pool, "export-unverified@example.com")
	seedCharacter(t, f.h.pool, f.gid, unverified, "us/hardcore/unverifiedofficer", "officer", false)
	stranger := seedUser(t, f.h.pool, "export-stranger@example.com")
	for name, uid := range map[string]int64{"unverified officer": unverified, "visitor": stranger} {
		f.h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
		res := f.get("attendance")
		res.Body.Close()
		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("%s = %d, want 404", name, res.StatusCode)
		}
	}
}

func TestExport_ModeratorAllowed(t *testing.T) {
	f := newExportFixture(t, "Export Moderator")
	mod := seedUser(t, f.h.pool, "export-mod@example.com")
	f.h.actor = auth.Actor{UserID: mod, Role: "moderator", Method: "session"}
	res := f.get("loot")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("moderator export = %d, want 200", res.StatusCode)
	}
}
