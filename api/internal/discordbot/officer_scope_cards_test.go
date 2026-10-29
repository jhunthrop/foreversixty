package discordbot

import (
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

const (
	rdyPickID = "grdy:1:raid:"
	repPickID = "grep:1:raid:"
)

func plannedRaids() []guilds.Raid {
	soon := time.Now().Add(48 * time.Hour)
	return []guilds.Raid{
		{ID: 3, Title: "Onyxia's Lair", StartsAt: soon, DurationMinutes: 180, Status: guilds.RaidPlanned},
		{ID: 4, Title: "Molten Core", StartsAt: soon.Add(72 * time.Hour), DurationMinutes: 180, Status: guilds.RaidPlanned},
	}
}

func readinessFor(raid guilds.Raid, signed int) guilds.ReadinessAnswer {
	return guilds.ReadinessAnswer{Raid: &raid, Signed: signed}
}

func TestReadinessSelect(t *testing.T) {
	w := newOfficerWorld(t)
	raids := plannedRaids()
	w.og.planned = raids
	w.og.readiness = readinessFor(raids[0], 20)
	w.og.raidReady = map[int64]guilds.ReadinessAnswer{4: readinessFailing(raids[1])}

	m := w.run(t, pressDiscord, "readiness").Message
	assertCardLimits(t, m)
	if len(m.Components) != 2 || m.Components[0][0].Style != StyleLink {
		t.Fatalf("rows = %+v", m.Components)
	}
	sel := m.Components[1][0]
	if sel.CustomID != rdyPickID || sel.Placeholder != "Another raid" || len(sel.Options) != 2 ||
		!sel.Options[0].Default || sel.Options[1].Default || !strings.HasPrefix(sel.Options[1].Label, "Molten Core · ") {
		t.Fatalf("select = %+v", sel)
	}

	rep := w.attendanceCard(t, rdyPickID, "4")
	assertCardLimits(t, rep.Message)
	if !rep.Update || !strings.Contains(rep.Message.Embeds[0].Title, "Molten Core") ||
		!strings.Contains(rep.Message.Embeds[0].Description, "11 signed") {
		t.Fatalf("pick = %+v", rep)
	}
	if o := rep.Message.Components[1][0].Options; o[0].Default || !o[1].Default {
		t.Fatalf("default after pick = %+v", o)
	}

	w.og.readyErr = guilds.ErrNotOfficer // demoted since the card was drawn
	rep = w.attendanceCard(t, rdyPickID, "4")
	if !rep.Update || rep.Message.Content != fill(officerCopy.NotOfficer, "{Guild}", "Olympus XXVII") {
		t.Fatalf("demoted pick = %+v", rep)
	}
	w.og.readyErr = nil
	if rep = w.attendanceCard(t, rdyPickID, "forged"); !rep.Update || rep.Message.Content != commandCopy.RaidNone {
		t.Fatalf("forged pick = %+v", rep)
	}

	w.og.planned = raids[:1]
	if m := w.run(t, pressDiscord, "readiness").Message; len(m.Components) != 1 {
		t.Fatalf("one raid has a select: %+v", m.Components)
	}
}

func TestReadinessRaidOption(t *testing.T) {
	for _, c := range Commands() {
		if c.Name != "readiness" {
			continue
		}
		if len(c.Options) != 1 || c.Options[0].Name != "raid" || c.Options[0].Required || !c.Options[0].Autocomplete ||
			c.Options[0].Description != "Which raid. The next one if left out." {
			t.Fatalf("options = %+v", c.Options)
		}
		w := newOfficerWorld(t)
		raids := plannedRaids()
		w.og.planned = raids
		w.og.raidReady = map[int64]guilds.ReadinessAnswer{4: readinessFor(raids[1], 9)}
		m := w.run(t, pressDiscord, "readiness", map[string]any{"name": "raid", "value": "4"}).Message
		if !strings.Contains(m.Embeds[0].Title, "Molten Core") || !m.Components[1][0].Options[1].Default {
			t.Fatalf("named raid = %+v", m)
		}
		return
	}
	t.Fatal("no /readiness")
}

func TestReportNightSelect(t *testing.T) {
	w := newOfficerWorld(t)
	w.addChar(t, w.gid, w.uid, "us/normal/zulmara", "Zulmara", true)
	last := onyxiaReport(false, true)
	older := onyxiaReport(false, false)
	older.Raid.ID, older.Raid.Title, older.Raid.Match = 6, "Molten Core", &guilds.MatchView{ReportID: "def"}
	older.Raid.StartsAt = last.Raid.StartsAt.Add(-7 * 24 * time.Hour)
	w.og.report = last
	w.og.reports = map[int64]*guilds.ReportAnswer{7: last, 6: older}
	w.og.done = []guilds.Raid{
		{ID: 7, Title: last.Raid.Title, StartsAt: last.Raid.StartsAt},
		{ID: 6, Title: older.Raid.Title, StartsAt: older.Raid.StartsAt},
	}

	m := w.run(t, pressDiscord, "report").Message
	assertCardLimits(t, m)
	if len(m.Components) != 2 || !hasShare(Reply{Message: m}) {
		t.Fatalf("rows = %+v", m.Components)
	}
	sel := m.Components[1][0]
	if sel.CustomID != repPickID || sel.Placeholder != "Another night" || len(sel.Options) != 2 ||
		!sel.Options[0].Default || sel.Options[1].Default {
		t.Fatalf("select = %+v", sel)
	}
	if d := sel.Options[0].Description; d != nightLine(*last.Raid.Night) || len([]rune(d)) > 100 {
		t.Fatalf("description = %q", d)
	}

	rep := w.attendanceCard(t, repPickID, "6")
	assertCardLimits(t, rep.Message)
	if !rep.Update || !strings.Contains(rep.Message.Embeds[0].Title, "Molten Core") || hasShare(rep) {
		t.Fatalf("private night pick = %+v", rep)
	}
	if o := rep.Message.Components[len(rep.Message.Components)-1][0].Options; o[0].Default || !o[1].Default {
		t.Fatalf("default after pick = %+v", o)
	}
	rep = w.attendanceCard(t, repPickID, "7")
	if !hasShare(rep) {
		t.Fatalf("public night pick lacks Share: %+v", rep.Message.Components)
	}

	w.og.done = w.og.done[:1]
	if m := w.run(t, pressDiscord, "report").Message; len(m.Components) != 1 {
		t.Fatalf("one night has a select: %+v", m.Components)
	}
}

func readinessFailing(raid guilds.Raid) guilds.ReadinessAnswer {
	return guilds.ReadinessAnswer{Raid: &raid, Signed: 11, Failing: 2, Rows: []string{"Vexlash · Short", "Thalwyn · Short"}}
}
