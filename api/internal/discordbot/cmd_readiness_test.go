package discordbot

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

func TestReadinessGate(t *testing.T) {
	w := newOfficerWorld(t)
	segment := func(i int, fill map[string]string) string {
		return specString(t, pressHeading, "`/readiness`", i, fill)
	}
	olympus := map[string]string{"Guild": "Olympus XXVII"}
	for _, c := range []struct {
		err    error
		want   string
		button string
	}{
		{guilds.ErrNotOfficer, segment(7, olympus), ""},
		{guilds.ErrPlanRequired, segment(9, nil), segment(10, nil)},
		{guilds.ErrRaidGuildRequired, segment(12, nil), segment(13, nil)},
		{guilds.ErrClaimContested, segment(11, nil), ""},
	} {
		w.og.readyErr = c.err
		rep := w.run(t, pressDiscord, "readiness")
		if rep.Message.Content != c.want || hasShare(rep) {
			t.Errorf("%v: %q, want %q", c.err, rep.Message.Content, c.want)
		}
		if c.button != "" {
			if b := rep.Message.Components[0][0]; b.Label != c.button || !strings.HasSuffix(b.URL, "/settings#billing") ||
				!strings.HasPrefix(b.URL, pressPublic+"/guild/") {
				t.Errorf("%v: button %+v", c.err, b)
			}
		}
	}
	w.og.readyErr = nil
	w.og.readiness = guilds.ReadinessAnswer{}
	if got := w.run(t, pressDiscord, "readiness").Message.Content; got != commandCopy.RaidNone {
		t.Fatalf("none planned = %q", got)
	}
	rep := w.run(t, strangerDiscord, "readiness")
	if rep.Message.Content != segment(8, olympus) || rep.Message.Components[0][0].Label != linkCopy.LinkButton {
		t.Fatalf("unlinked = %+v", rep.Message)
	}

	raid := guilds.Raid{ID: 3, Title: "Onyxia's Lair", StartsAt: time.Date(2026, 12, 10, 20, 0, 0, 0, time.UTC)}
	rows := make([]string, 10)
	for i := range rows {
		rows[i] = fmt.Sprintf("Raider%d · Short · Synced 9 days before the raid", i)
	}
	w.og.readiness = guilds.ReadinessAnswer{Raid: &raid, Signed: 23, Failing: 12, Rows: rows, More: 2}
	rep = w.run(t, pressDiscord, "readiness")
	assertMessageLimits(t, rep.Message)
	e := rep.Message.Embeds[0]
	wantTitle := segment(0, map[string]string{"Title": "Onyxia's Lair", "Weekday": "Thursday", "d": "10", "Month": "December"})
	if e.Title != wantTitle || e.Description != "The 23 signed In or Late. 12 have something to fix." || e.Fields[0].Name != segment(4, nil) {
		t.Fatalf("embed = %+v", e)
	}
	lines := strings.Split(e.Fields[0].Value, "\n")
	if len(lines) != 11 || lines[0] != rows[0] || lines[10] != segment(5, map[string]string{"n": "2"}) {
		t.Fatalf("rows = %q", lines)
	}
	if b := rep.Message.Components[0][0]; b.Label != segment(6, nil) || hasShare(rep) || !strings.HasSuffix(b.URL, "/readiness") {
		t.Fatalf("row = %+v", rep.Message.Components)
	}
	w.og.readiness = guilds.ReadinessAnswer{Raid: &raid, Signed: 23, Failing: 1, Rows: rows[:1]}
	if got := w.run(t, pressDiscord, "readiness").Message.Embeds[0].Description; got != "The 23 signed In or Late. 1 has something to fix." {
		t.Fatalf("one = %q", got)
	}
	w.og.readiness = guilds.ReadinessAnswer{Raid: &raid, Signed: 23}
	if got := w.run(t, pressDiscord, "readiness").Message.Embeds[0]; got.Description != segment(3, nil) || len(got.Fields) != 0 {
		t.Fatalf("all ready = %+v", got)
	}
}
