package guilds

import (
	"errors"
	"testing"
)

func TestLowestAttendanceIsOfficersOnlyLowestFirst(t *testing.T) {
	t.Run("order and limit", func(t *testing.T) {
		row := func(key, name string, present, nights int) RosterRow {
			return RosterRow{CharacterKey: key, Name: name, Attendance: Attendance{Present: present, Nights: nights}}
		}
		roster := map[string]RosterRow{
			"us/h/a": row("us/h/a", "Seven", 7, 8), "us/h/b": row("us/h/b", "Three", 3, 8),
			"us/h/c": row("us/h/c", "Five", 5, 8), "us/h/d": row("us/h/d", "Nonight", 0, 0),
			"us/h/e": row("us/h/e", "Anna", 5, 8),
		}
		got := lowestOf(roster, AttendanceSourceReports, 10)
		want := []string{"Three", "Anna", "Five", "Seven"}
		if len(got) != len(want) {
			t.Fatalf("lowest = %+v", got)
		}
		for i, n := range want {
			if got[i].Name != n {
				t.Fatalf("order = %+v, want %v", got, want)
			}
		}
		if got[0].Chip != "3 of 8 nights" {
			t.Fatalf("chip = %q", got[0].Chip)
		}
		if two := lowestOf(roster, AttendanceSourceReports, 2); len(two) != 2 || two[1].Name != "Anna" {
			t.Fatalf("limit 2 = %+v", two)
		}
	})
	t.Run("rank and nights from the database", func(t *testing.T) {
		w, _ := matchedNight(t)
		seedExport(t, w.h.pool, w.member, memberKeyA, "us", "hardcore", "Kraggor")
		seedExport(t, w.h.pool, w.member2, member2Key, "us", "hardcore", "Mirelle")
		got, err := w.h.svc.LowestAttendance(t.Context(), w.gid, w.officer, 25)
		if err != nil {
			t.Fatal(err)
		}
		for _, r := range got {
			if r.Chip == "" || len(r.Last) != 0 {
				t.Fatalf("record %+v has no chip or carries raid lines", r)
			}
		}
		if len(got) == 0 {
			t.Fatal("no member with a night was listed")
		}
		for _, who := range []int64{w.member, w.member2, seedUser(t, w.h.pool, "outsider@example.com")} {
			if _, err := w.h.svc.LowestAttendance(t.Context(), w.gid, who, 25); !errors.Is(err, ErrOfficerOnly) {
				t.Fatalf("viewer %d = %v", who, err)
			}
		}
	})
}
