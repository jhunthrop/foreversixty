package discordbot

import (
	"context"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

const (
	attPickID = "gatt:1:pick:"
	attAskID  = "gatt:1:ask:"
	attNameID = "gatt:1:name:"
)

// attendanceCard presses a card component of /attendance (a modal submit when typed is set).
func (w *officerWorld) attendanceCard(t *testing.T, id string, values ...string) Reply {
	t.Helper()
	data := map[string]any{"custom_id": id}
	kind := typeComponent
	if len(values) > 0 {
		data["values"] = values
	}
	return w.routeCard(t, interactionOf(kind, pressDiscord, data))
}

func (w *officerWorld) attendanceModal(t *testing.T, typed string) Reply {
	t.Helper()
	data := map[string]any{"custom_id": attNameID, "components": []map[string]any{
		{"type": 1, "components": []map[string]any{{"type": 4, "custom_id": "name", "value": typed}}},
	}}
	return w.routeCard(t, interactionOf(typeModalSubmit, pressDiscord, data))
}

func (w *officerWorld) routeCard(t *testing.T, in Interaction) Reply {
	t.Helper()
	h, ok := w.svc.components[customPrefix(in.CustomID())]
	if !ok {
		t.Fatalf("no route for %q", in.CustomID())
	}
	rep, err := h(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func officerLowest(w *officerWorld) {
	w.og.lowestErr = nil
	w.og.lowest = []guilds.AttendanceRecord{
		{Name: "Thalwyn", Class: "mage", Chip: "Came 3 of 8 · 10 of 30 raids"},
		{Name: "Vexlash", Class: "rogue", Chip: "Came 5 of 8 · 20 of 30 raids"},
	}
	w.og.records = []guilds.AttendanceRecord{{Name: "Zulmara", Class: "hunter", Chip: "Came 7 of 8 · 26 of 30 raids"}}
	w.og.byName = map[string]guilds.AttendanceRecord{
		"Thalwyn": w.og.lowest[0], "Vexlash": w.og.lowest[1],
	}
}

func TestOfficerAttendanceSelect(t *testing.T) {
	w := newOfficerWorld(t)
	officerLowest(w)
	m := w.run(t, pressDiscord, "attendance").Message
	assertCardLimits(t, m)
	if len(m.Components) != 2 {
		t.Fatalf("rows = %+v", m.Components)
	}
	sel := m.Components[0][0]
	if sel.CustomID != attPickID || sel.Placeholder != guideCopy.AttendancePlaceholder || len(sel.Options) != 2 {
		t.Fatalf("select = %+v", sel)
	}
	if o := sel.Options[0]; o.Label != "Thalwyn" || o.Value != "Thalwyn" || o.Description != "Came 3 of 8 · 10 of 30 raids" {
		t.Fatalf("first option = %+v", o)
	}
	for _, o := range sel.Options {
		if o.Default {
			t.Fatalf("nothing is pre-chosen: %+v", o)
		}
	}
	last := m.Components[len(m.Components)-1]
	if len(last) != 2 || last[0].Label != "Type a name" || last[0].Style != StyleSecondary || last[0].CustomID != attAskID ||
		last[1].Label != officerCopy.Roster || last[1].Style != StyleLink {
		t.Fatalf("last row = %+v", last)
	}
	w.og.lowestErr = guilds.ErrOfficerOnly
	member := w.run(t, pressDiscord, "attendance").Message
	if len(member.Components) != 1 || len(member.Components[0]) != 1 {
		t.Fatalf("a member's answer has controls: %+v", member.Components)
	}
}

func TestOfficerAttendancePickAndDemotion(t *testing.T) {
	w := newOfficerWorld(t)
	officerLowest(w)
	rep := w.attendanceCard(t, attPickID, "Thalwyn")
	assertCardLimits(t, rep.Message)
	if !rep.Update || len(rep.Message.Embeds) != 1 || !strings.Contains(rep.Message.Embeds[0].Title, "Thalwyn") {
		t.Fatalf("pick = %+v", rep)
	}
	sel := rep.Message.Components[0][0]
	if sel.CustomID != attPickID || !sel.Options[0].Default || sel.Options[1].Default {
		t.Fatalf("select after pick = %+v", sel)
	}
	w.og.recordsErr = guilds.ErrOfficerOnly // demoted
	rep = w.attendanceCard(t, attPickID, "Thalwyn")
	if !rep.Update || rep.Message.Content != officerCopy.OfficersOnly {
		t.Fatalf("demoted pick = %+v", rep)
	}
}

func TestAttendanceModalUnknownName(t *testing.T) {
	w := newOfficerWorld(t)
	officerLowest(w)
	ask := w.attendanceCard(t, attAskID)
	if ask.Modal == nil || ask.Modal.Title != "Look up attendance" || ask.Modal.CustomID != attNameID {
		t.Fatalf("ask = %+v", ask)
	}
	rep := w.attendanceModal(t, "Nobody")
	assertCardLimits(t, rep.Message)
	want := "No member of Olympus XXVII named Nobody."
	if rep.Message.Content != want || len(rep.Message.Embeds) != 1 || rep.Message.Embeds[0].Title == "" ||
		!strings.Contains(rep.Message.Embeds[0].Title, "Zulmara") {
		t.Fatalf("unknown = %+v", rep.Message)
	}
	found := w.attendanceModal(t, "Vexlash")
	if found.Message.Content != "" || !strings.Contains(found.Message.Embeds[0].Title, "Vexlash") ||
		!found.Message.Components[0][0].Options[1].Default {
		t.Fatalf("typed name = %+v", found.Message)
	}
}
