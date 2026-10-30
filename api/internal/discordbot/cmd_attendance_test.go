package discordbot

import (
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

func TestAttendanceRules(t *testing.T) {
	w := newOfficerWorld(t)
	w.exec(t, `update guilds set raid_timezone = 'America/Chicago' where id = $1`, w.gid)
	segment := func(i int, fill map[string]string) string {
		return specString(t, pressHeading, "`/attendance`", i, fill)
	}
	at := func(d int) time.Time { return time.Date(2026, 12, d, 2, 0, 0, 0, time.UTC) } // the evening before, in Chicago
	w.og.records = []guilds.AttendanceRecord{
		{Name: "Zulmara", Class: "hunter", Chip: "Came 7 of 8 · 26 of 30 raids", Last: []guilds.AttendanceLine{
			{StartsAt: at(11), Outcome: guilds.OutcomePresent},
			{StartsAt: at(4), Outcome: guilds.OutcomeLate, Excused: true},
			{StartsAt: time.Date(2026, 11, 29, 2, 0, 0, 0, time.UTC), Outcome: guilds.OutcomeNoShow},
		}},
		{Name: "Duskstrider", Class: "rogue", Chip: "Came 1 of 1 raids", Last: []guilds.AttendanceLine{
			{StartsAt: at(11), Outcome: guilds.OutcomePresent, CameAs: "Zulmara"}}},
	}
	rep := w.run(t, pressDiscord, "attendance")
	assertMessageLimits(t, rep.Message)
	if len(rep.Message.Embeds) != 2 || hasShare(rep) || !rep.Message.Ephemeral {
		t.Fatalf("answer = %+v", rep.Message)
	}
	main := rep.Message.Embeds[0]
	if main.Title != segment(0, map[string]string{"Name": "Zulmara"}) || main.Color != classColorOf("hunter") ||
		main.Description != "Came 7 of 8 · 26 of 30 raids" || main.Fields[0].Name != segment(1, nil) {
		t.Fatalf("main = %+v", main)
	}
	if got := main.Fields[0].Value; got != "Thu 10 Dec Present\nThu 3 Dec Late, excused\nSat 28 Nov Not in the log" ||
		strings.Contains(strings.ToLower(got), "no-show") {
		t.Fatalf("last raids = %q", got)
	}
	if got := rep.Message.Embeds[1].Fields[0].Value; got != "Thu 10 Dec Came as Zulmara" {
		t.Fatalf("alt = %q", got)
	}
	if b := rep.Message.Components[0][0]; b.Label != segment(2, nil) || !strings.HasSuffix(b.URL, "/roster") {
		t.Fatalf("row = %+v", rep.Message.Components)
	}
	olympus := map[string]string{"Guild": "Olympus XXVII"}
	for _, c := range []struct {
		err  error
		want string
	}{
		{guilds.ErrOfficerOnly, segment(3, nil)},
		{guilds.ErrNoSuchMember, segment(4, olympus)},
		{guilds.ErrNotMember, segment(7, olympus)},
	} {
		w.og.recordsErr = c.err
		if got := w.run(t, pressDiscord, "attendance", opt(optAttendanceName, "Vexlash")).Message.Content; got != c.want {
			t.Errorf("%v: %q, want %q", c.err, got, c.want)
		}
	}
	w.og.recordsErr = nil
	w.og.records = []guilds.AttendanceRecord{{Name: "Zulmara", Class: "hunter"}}
	if got := w.run(t, pressDiscord, "attendance").Message.Content; got != segment(5, nil) {
		t.Fatalf("no raids = %q", got)
	}
	rep = w.run(t, strangerDiscord, "attendance")
	if rep.Message.Content != segment(6, nil) || rep.Message.Components[0][0].Label != linkCopy.LinkButton {
		t.Fatalf("unlinked = %+v", rep.Message)
	}
}
