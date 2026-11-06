package guilds

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

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
	if _, err := f.h.pool.Exec(t.Context(), `
		insert into guild_raid_attendance_overrides (raid_id, character_key, status, auto_outcome)
		values ($1, 'us/hardcore/amy', 'bench', 'late')`, raid); err != nil {
		t.Fatal(err)
	}

	rows := f.rows("attendance")
	want := [][]string{
		{"date", "raid", "character", "class", "signed", "outcome", "excused", "pulls_present", "pulls_total", "fixed"},
		{"2026-10-20", "Onyxia night", "Amy", "hunter", "in", "late", "true", "4", "5", "bench"},
		{"2026-10-20", "Onyxia night", "Zed", "mage", "in", "present", "true", "4", "5", ""},
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

func TestExportLoot_AnInvitedMembersAwardReadsAsTheirCharacter(t *testing.T) {
	f := newExportFixture(t, "Export Invited")
	uid := seedUser(t, f.h.pool, "export-invited@example.com")
	key := syntheticKey(uid)
	seedCharacter(t, f.h.pool, f.gid, uid, key, "member", true)
	if _, err := f.h.pool.Exec(t.Context(), `
		insert into characters (key, user_id, region, ruleset, name, class)
		values ('us/hardcore/brightwing', $1, 'us', 'hardcore', 'Brightwing', 'priest')`, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := f.h.pool.Exec(t.Context(), `
		insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by, awarded_at)
		values ($1, 1084, 16963, $2, $3, '2026-10-20 21:00+00')`, f.gid, key, f.officer); err != nil {
		t.Fatal(err)
	}
	rows := f.rows("loot")
	if len(rows) != 2 || rows[1][5] != "Brightwing" {
		t.Fatalf("rows = %v, want the award named Brightwing", rows)
	}
}

// A logged night no planned raid holds exports the characters its log shows, the rows behind
// the Raids tab's "N of N present", next to a planned raid's rows, in night order.
func TestExportAttendance_LoggedNightWithNoPlannedRaid(t *testing.T) {
	f := newExportFixture(t, "Export Logged Night")
	raid := f.seedRaid("Onyxia night")
	f.seedAttendance(raid, "us/hardcore/zed", "Zed", "mage", "present")
	night := time.Date(2026, 10, 27, 19, 0, 0, 0, time.UTC)
	seedNight(t, f.h.pool, f.gid, "export-night-1", "Onyxia's Lair", night, []seedFight{
		{encounterID: 1084, kill: true, duration: 4 * time.Minute, players: []seedPlayer{
			{key: "us/hardcore/amy", name: "Amy", class: "hunter", spec: "marksmanship", role: "dps"},
			{key: "us/hardcore/bo", name: "Bo", class: "priest", spec: "holy", role: "healer"},
		}},
	})
	if _, err := f.h.pool.Exec(t.Context(), `update reports set zone = 'Onyxia''s Lair' where id = 'export-night-1'`); err != nil {
		t.Fatal(err)
	}
	rows := f.rows("attendance")
	want := [][]string{
		{"date", "raid", "character", "class", "signed", "outcome", "excused", "pulls_present", "pulls_total", "fixed"},
		{"2026-10-20", "Onyxia night", "Zed", "mage", "in", "present", "true", "4", "5", ""},
		{"2026-10-27", "Onyxia's Lair", "Amy", "hunter", "", "present", "false", "1", "1", ""},
		{"2026-10-27", "Onyxia's Lair", "Bo", "priest", "", "present", "false", "1", "1", ""},
	}
	if fmt.Sprint(rows) != fmt.Sprint(want) {
		t.Fatalf("rows = %v, want %v", rows, want)
	}
}

// An award made from a report is dated by that report's night and named by it; one with
// neither a raid nor a report says so in the raid cell.
func TestExportLoot_NightFromTheReportElseSaysSo(t *testing.T) {
	f := newExportFixture(t, "Export Loot Night")
	night := time.Date(2026, 10, 27, 19, 0, 0, 0, time.UTC)
	seedNight(t, f.h.pool, f.gid, "export-loot-night", "Onyxia's Lair", night, nil)
	if _, err := f.h.pool.Exec(t.Context(), `update reports set zone = 'Onyxia''s Lair' where id = 'export-loot-night'`); err != nil {
		t.Fatal(err)
	}
	if _, err := f.h.pool.Exec(t.Context(), `
		insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by, awarded_at, report_id)
		values ($1, 1084, 16963, 'us/hardcore/zed', $2, '2026-10-29 12:00+00', 'export-loot-night'),
		       ($1, 1084, 17067, 'us/hardcore/zed', $2, '2026-10-30 12:00+00', null)`, f.gid, f.officer); err != nil {
		t.Fatal(err)
	}
	rows := f.rows("loot")
	if len(rows) != 3 {
		t.Fatalf("rows = %v, want two awards", rows)
	}
	if rows[1][0] != "2026-10-27" || rows[1][1] != "Onyxia's Lair" {
		t.Fatalf("report award = %v, want the report's night", rows[1])
	}
	if rows[2][0] != "2026-10-30" || rows[2][1] != lootNightUnknown {
		t.Fatalf("raidless award = %v, want the award date and %q", rows[2], lootNightUnknown)
	}
}

// A Wednesday 7 pm raid in Chicago is 01:00 UTC Thursday: both CSVs date it Wednesday, in the
// guild's raid zone, and so does a logged night no planned raid holds.
func TestExport_DatesInTheGuildsRaidZone(t *testing.T) {
	f := newExportFixture(t, "Export Zone")
	if err := f.h.svc.Store.SetGuildRaidTimezone(t.Context(), f.gid, "America/Chicago"); err != nil {
		t.Fatal(err)
	}
	var raid int64
	if err := f.h.pool.QueryRow(t.Context(), `
		insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, status, created_by)
		values ($1, 'onyxia', 'Onyxia night', '2026-12-10 01:00+00', 180, 'done', $2) returning id`,
		f.gid, f.officer).Scan(&raid); err != nil {
		t.Fatal(err)
	}
	f.seedAttendance(raid, "us/hardcore/zed", "Zed", "mage", "present")
	if _, err := f.h.pool.Exec(t.Context(), `
		insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by, awarded_at, raid_id)
		values ($1, 1084, 16963, 'us/hardcore/zed', $2, '2026-12-10 02:00+00', $3)`, f.gid, f.officer, raid); err != nil {
		t.Fatal(err)
	}
	logged := time.Date(2026, 12, 17, 1, 30, 0, 0, time.UTC) // Wed 16 Dec, 7:30 pm in Chicago
	seedNight(t, f.h.pool, f.gid, "export-zone-night", "Onyxia's Lair", logged, []seedFight{
		{encounterID: 1084, kill: true, duration: 4 * time.Minute, players: []seedPlayer{
			{key: "us/hardcore/amy", name: "Amy", class: "hunter", spec: "marksmanship", role: "dps"},
		}},
	})
	att := f.rows("attendance")
	if len(att) != 3 || att[1][0] != "2026-12-09" || att[2][0] != "2026-12-16" {
		t.Fatalf("attendance = %v, want the planned night 2026-12-09 and the logged night 2026-12-16", att)
	}
	if loot := f.rows("loot"); len(loot) != 2 || loot[1][0] != "2026-12-09" {
		t.Fatalf("loot = %v, want the award dated 2026-12-09", loot)
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

// serveDirect runs serveExport on stream as the fixture's officer.
func (f exportFixture) serveDirect(stream rowStream) *httptest.ResponseRecorder {
	f.h.t.Helper()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/export/loot.csv", f.gid), nil)
	r.SetPathValue("id", strconv.FormatInt(f.gid, 10))
	r = r.WithContext(auth.WithActor(r.Context(), f.h.actor))
	w := httptest.NewRecorder()
	f.h.svc.serveExport(w, r, "loot", lootExportHeader, stream)
	return w
}

func emitRows(n int, err error) rowStream {
	return func(_ context.Context, _ int64, emit func([]string) error) error {
		for i := range n {
			if e := emit([]string{strconv.Itoa(i)}); e != nil {
				return e
			}
		}
		return err
	}
}

// A query that fails before the first flush answers an error, not a short file.
func TestExport_EarlyFailureAnswersError(t *testing.T) {
	f := newExportFixture(t, "Export Early Failure")
	boom := errors.New("boom")
	for name, stream := range map[string]rowStream{
		"first row": emitRows(0, boom),
		"some rows": emitRows(exportFlushEvery-1, boom),
	} {
		w := f.serveDirect(stream)
		if w.Code != http.StatusInternalServerError || w.Header().Get("Content-Disposition") != "" ||
			strings.HasPrefix(w.Body.String(), "date,") {
			t.Fatalf("%s: status = %d, disposition = %q, body = %.60q", name, w.Code,
				w.Header().Get("Content-Disposition"), w.Body.String())
		}
	}
}

// Past the first flush the file is already going out: a failure ends it short.
func TestExport_LateFailureKeepsTheFile(t *testing.T) {
	f := newExportFixture(t, "Export Late Failure")
	w := f.serveDirect(emitRows(exportFlushEvery+1, errors.New("boom")))
	if w.Code != http.StatusOK || w.Header().Get("Content-Type") != csvContentType {
		t.Fatalf("status = %d, type = %q", w.Code, w.Header().Get("Content-Type"))
	}
	if rows := strings.Count(w.Body.String(), "\n"); rows != exportFlushEvery+1 {
		t.Fatalf("lines = %d, want header plus %d", rows, exportFlushEvery)
	}
}

func TestExport_SuccessUnderOneFlushSendsEverything(t *testing.T) {
	f := newExportFixture(t, "Export Small")
	w := f.serveDirect(emitRows(3, nil))
	if w.Code != http.StatusOK || strings.Count(w.Body.String(), "\n") != 4 {
		t.Fatalf("status = %d, body = %q", w.Code, w.Body.String())
	}
}
