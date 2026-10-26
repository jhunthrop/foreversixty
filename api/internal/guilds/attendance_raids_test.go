package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// seedAtt is one attendance row a test stores directly.
type seedAtt struct {
	key, signed, outcome string
	excused              bool
}

// seedDoneRaid stores a done raid starting at startsAt with the given attendance rows.
func seedDoneRaid(t *testing.T, h *httpHarness, guildID, createdBy int64, startsAt time.Time, rows ...seedAtt) int64 {
	t.Helper()
	ctx := context.Background()
	var raidID int64
	if err := h.pool.QueryRow(ctx, `
		insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, status, created_by)
		values ($1, 'raid:onyxias-lair', 'Night', $2, 180, 'done', $3) returning id`,
		guildID, startsAt, createdBy).Scan(&raidID); err != nil {
		t.Fatal(err)
	}
	for _, r := range rows {
		var signed *string
		if r.signed != "" {
			signed = &r.signed
		}
		if _, err := h.pool.Exec(ctx, `
			insert into guild_raid_attendance (raid_id, character_key, name, class, signed, outcome, excused)
			values ($1, $2, $2, 'warrior', $3, $4, $5)`, raidID, r.key, signed, r.outcome, r.excused); err != nil {
			t.Fatal(err)
		}
	}
	return raidID
}

func attendanceWorld(t *testing.T) (*httpHarness, int64, int64, string) {
	t.Helper()
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Raiders")
	uid := seedUser(t, h.pool, "raid-attendance@example.com")
	key := "us/hardcore/raidattender"
	seedCharacter(t, h.pool, gid, uid, key, "member", true)
	seedExport(t, h.pool, uid, key, "us", "hardcore", "raidattender")
	recomputeMembership(t, h.pool, gid, uid)
	return h, gid, uid, key
}

func TestHomeReportsModeBeforeAnyDoneRaid(t *testing.T) {
	h, gid, uid, key := attendanceWorld(t)
	row := rosterRowOf(t, h, gid, uid, key)
	if row.Attendance30 != nil {
		t.Errorf("attendance_30 = %+v, want absent in report mode", row.Attendance30)
	}
	var view HomeView
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	h.data(res, &view)
	if view.AttendanceSource != "reports" {
		t.Errorf("attendance_source = %q, want reports", view.AttendanceSource)
	}
}

func TestHomeSwitchesToRaidsOnceAGuildHasADoneRaid(t *testing.T) {
	h, gid, uid, key := attendanceWorld(t)
	now := time.Now()
	// Newest first: came, late (counts), excused (neither), signed in but no-show, bench (not a night).
	seedDoneRaid(t, h, gid, uid, now.Add(-1*24*time.Hour), seedAtt{key: key, signed: "in", outcome: "present"})
	seedDoneRaid(t, h, gid, uid, now.Add(-2*24*time.Hour), seedAtt{key: key, signed: "late", outcome: "late"})
	seedDoneRaid(t, h, gid, uid, now.Add(-3*24*time.Hour), seedAtt{key: key, signed: "in", outcome: "absent", excused: true})
	seedDoneRaid(t, h, gid, uid, now.Add(-4*24*time.Hour), seedAtt{key: key, signed: "in", outcome: "no_show"})
	seedDoneRaid(t, h, gid, uid, now.Add(-5*24*time.Hour), seedAtt{key: key, signed: "bench", outcome: "absent"})
	seedDoneRaid(t, h, gid, uid, now.Add(-6*24*time.Hour), seedAtt{key: key, outcome: "present"})

	row := rosterRowOf(t, h, gid, uid, key)
	var view HomeView
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	h.data(res, &view)
	if view.AttendanceSource != "raids" {
		t.Fatalf("attendance_source = %q, want raids", view.AttendanceSource)
	}
	if row.Attendance != (Attendance{Present: 3, Nights: 4}) {
		t.Errorf("attendance = %+v, want 3 of 4 (excused and bench-absent are not nights)", row.Attendance)
	}
	if row.Attendance30 == nil || *row.Attendance30 != (Attendance{Present: 3, Nights: 4}) {
		t.Errorf("attendance_30 = %+v, want 3 of 4", row.Attendance30)
	}
}

func TestRaidAttendanceWindowsAndACharacterWithNoNights(t *testing.T) {
	h, gid, uid, key := attendanceWorld(t)
	other := "us/hardcore/neverraided"
	otherUID := seedUser(t, h.pool, "never-raided@example.com")
	seedCharacter(t, h.pool, gid, otherUID, other, "member", true)
	seedExport(t, h.pool, otherUID, other, "us", "hardcore", "neverraided")
	recomputeMembership(t, h.pool, gid, otherUID)

	now := time.Now()
	// 10 logged raids: the oldest two are outside the 8 window; the character missed the oldest.
	for i := 0; i < 10; i++ {
		outcome := "present"
		if i == 9 || i == 3 {
			outcome = "no_show"
		}
		seedDoneRaid(t, h, gid, uid, now.Add(-time.Duration(i+1)*24*time.Hour), seedAtt{key: key, signed: "in", outcome: outcome})
	}
	row := rosterRowOf(t, h, gid, uid, key)
	if row.Attendance != (Attendance{Present: 7, Nights: 8}) {
		t.Errorf("attendance = %+v, want 7 of 8", row.Attendance)
	}
	if row.Attendance30 == nil || *row.Attendance30 != (Attendance{Present: 8, Nights: 10}) {
		t.Errorf("attendance_30 = %+v, want 8 of 10", row.Attendance30)
	}
	none := rosterRowOf(t, h, gid, uid, other)
	if none.Attendance != (Attendance{}) {
		t.Errorf("a character in no raid reads %+v, want 0 of 0", none.Attendance)
	}
}

func TestCharacterAttendanceIsAbsentInReportMode(t *testing.T) {
	h, gid, uid, key := attendanceWorld(t)
	got, err := h.store.CharacterAttendance(context.Background(), gid, []string{key})
	if err != nil || len(got) != 0 {
		t.Fatalf("report mode = %v, %v; want empty", got, err)
	}
	seedDoneRaid(t, h, gid, uid, time.Now().Add(-time.Hour), seedAtt{key: key, signed: "in", outcome: "present"})
	got, err = h.store.CharacterAttendance(context.Background(), gid, []string{key})
	if err != nil || got[key].Present != 1 || got[key].Nights != 1 {
		t.Fatalf("raid mode = %v, %v; want 1 of 1", got, err)
	}
}
