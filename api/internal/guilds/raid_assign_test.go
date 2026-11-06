package guilds

import (
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"
)

// Every boss the client carries has a template, so a new boss fails here until its rows exist.
func TestEveryClientBossHasAnAssignmentTemplate(t *testing.T) {
	roles := map[string]bool{slotRoleTank: true, slotRoleHealer: true, slotRoleMelee: true, slotRoleRanged: true, slotRoleAny: true}
	for _, raid := range firstTierRaids {
		for _, boss := range raid.Bosses {
			slots, ok := templateOf(boss.EncounterID)
			if !ok || len(slots) == 0 {
				t.Errorf("%s (%d) has no template", boss.Name, boss.EncounterID)
			}
			seen := map[string]bool{}
			for _, s := range slots {
				if s.Key == "" || len(s.Key) > maxSlotKeyLen || seen[s.Key] {
					t.Errorf("%s: bad or repeated key %q", boss.Name, s.Key)
				}
				seen[s.Key] = true
				if s.Label == "" || s.Purpose == "" || s.Warning == "" || s.Count < 1 || !roles[s.Role] {
					t.Errorf("%s/%s: incomplete slot %+v", boss.Name, s.Key, s)
				}
				if strings.Contains(s.Purpose, "—") || strings.Contains(s.Label, "—") {
					t.Errorf("%s/%s: em dash in copy", boss.Name, s.Key)
				}
			}
		}
	}
	if _, ok := templateOf(1084); !ok {
		t.Fatal("Onyxia has no template")
	}
}

// assignWorld is a raid world with an Onyxia raid: Thornhide (officer) and Kraggor in as tanks,
// Mirelle in as a healer; Sunderfel has not signed.
func assignWorld(t *testing.T) (raidWorld, RaidView) {
	t.Helper()
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.officer)
	w.signup(t, raid.ID, officerKey, `{"state":"in","role":"tank"}`)
	w.as(w.member)
	w.signup(t, raid.ID, memberKeyA, `{"state":"in","role":"tank"}`)
	w.as(w.member2)
	w.signup(t, raid.ID, member2Key, `{"state":"in","role":"healer"}`)
	w.as(w.officer)
	return w, w.getRaid(t, raid.ID)
}

func slotPath(w raidWorld, raidID int64, slot string) string {
	return w.path("/raids/%d/assignments/1084/%s", raidID, slot)
}

func slotNamed(t *testing.T, v RaidView, key string) SlotAssignment {
	t.Helper()
	if v.Assignments == nil || len(v.Assignments.Encounters) != 1 {
		t.Fatalf("assignments = %+v", v.Assignments)
	}
	for _, s := range v.Assignments.Encounters[0].Slots {
		if s.Key == key {
			return s
		}
	}
	t.Fatalf("no slot %s", key)
	return SlotAssignment{}
}

func TestAssignmentsBlockFollowsTheZone(t *testing.T) {
	w, raid := assignWorld(t)
	a := raid.Assignments
	if a == nil || len(a.NotInClient) != 0 || len(a.Encounters) != 1 || a.PostedAt != nil {
		t.Fatalf("onyxia assignments = %+v", a)
	}
	e := a.Encounters[0]
	if e.EncounterID != 1084 || e.Name != "Onyxia" || e.RaidZone != "Onyxia's Lair" || len(e.Slots) != 12 || len(e.RaidWarning) != 0 {
		t.Fatalf("encounter = %+v", e)
	}
	if len(a.Candidates) != 3 || a.Candidates[0].Group != "tank" || a.Candidates[2].Name != "Mirelle" {
		t.Fatalf("candidates = %+v", a.Candidates)
	}
	for _, zone := range []string{"raid:barrow-deeps", "raid:hyjal-summit"} {
		body := strings.Replace(raidBody(futureStart(), 0), raidZoneOny, zone, 1)
		v := w.createRaids(t, body)[0]
		if v.Assignments == nil || len(v.Assignments.Encounters) != 0 || len(v.Assignments.NotInClient) != 1 ||
			v.Assignments.NotInClient[0] != v.ZoneName {
			t.Errorf("%s assignments = %+v", zone, v.Assignments)
		}
	}
	body := `{"zone":"dungeon:scholomance","title":"","starts_at":"` + futureStart() + `","duration_minutes":180,` +
		`"notes":"","timezone":"America/Chicago","repeat_weeks":0,"post_to_discord":false}`
	if v := w.createRaids(t, body)[0]; v.Assignments != nil {
		t.Errorf("dungeon assignments = %+v", v.Assignments)
	}
}

func TestAssignmentWritesAreAnOfficersOnly(t *testing.T) {
	w, raid := assignWorld(t)
	path := slotPath(w, raid.ID, "main_tank")
	w.as(w.member)
	expectError(t, w.h.do(http.MethodPut, path, `{"character_keys":["`+memberKeyA+`"]}`), http.StatusForbidden, "forbidden")
	expectError(t, w.h.do(http.MethodDelete, path, ""), http.StatusForbidden, "forbidden")
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/assignments/post", raid.ID), ""), http.StatusForbidden, "forbidden")
	w.as(0)
	expectError(t, w.h.do(http.MethodPut, path, `{"character_keys":[]}`), http.StatusNotFound, "not_found")
}

func TestAssignmentValidation(t *testing.T) {
	w, raid := assignWorld(t)
	put := func(slot, body string) *http.Response {
		return w.h.do(http.MethodPut, slotPath(w, raid.ID, slot), body)
	}
	expectError(t, put("nope", `{"character_keys":[]}`), http.StatusNotFound, "not_found")
	expectError(t, w.h.do(http.MethodPut, w.path("/raids/%d/assignments/9999/main_tank", raid.ID), `{"character_keys":[]}`),
		http.StatusNotFound, "not_found")
	expectError(t, put("main_tank", `{}`), http.StatusUnprocessableEntity, "invalid")
	expectError(t, put("main_tank", `{"character_keys":["`+officerKey+`","`+memberKeyA+`"]}`), http.StatusUnprocessableEntity, "invalid")
	expectError(t, put("off_tank", `{"character_keys":["`+memberKeyA+`","`+memberKeyA+`"]}`), http.StatusUnprocessableEntity, "invalid")
	expectError(t, put("off_tank", `{"character_keys":["`+memberKeyB+`"]}`), http.StatusUnprocessableEntity, "invalid")
}

func TestAssignmentSetReadAndClear(t *testing.T) {
	w, raid := assignWorld(t)
	var v RaidView
	w.h.data(w.h.do(http.MethodPut, slotPath(w, raid.ID, "main_tank"), `{"character_keys":["`+memberKeyA+`"]}`), &v)
	// A healer in a tank slot is a warning, never a refusal.
	w.h.data(w.h.do(http.MethodPut, slotPath(w, raid.ID, "off_tank"), `{"character_keys":["`+officerKey+`","`+member2Key+`"]}`), &v)
	ot := slotNamed(t, v, "off_tank")
	if len(ot.Names) != 2 || ot.Names[0].Name != "Thornhide" || ot.Names[0].OffRole || !ot.Names[1].OffRole {
		t.Fatalf("off_tank = %+v", ot)
	}
	warn := v.Assignments.Encounters[0].RaidWarning
	if len(warn) != 2 || warn[0] != "MT: Kraggor" || warn[1] != "OT: Thornhide, Mirelle" {
		t.Fatalf("raid_warning = %q", warn)
	}

	w.as(w.member)
	mt := slotNamed(t, w.getRaid(t, raid.ID), "main_tank")
	if len(mt.Names) != 1 || !mt.Names[0].Own || mt.Names[0].CharacterKey != memberKeyA {
		t.Fatalf("member reads main_tank = %+v", mt)
	}
	// A pick who signs Out drops out of every reading.
	w.signup(t, raid.ID, memberKeyA, `{"state":"out"}`)
	if mt := slotNamed(t, w.getRaid(t, raid.ID), "main_tank"); len(mt.Names) != 0 {
		t.Fatalf("out pick still reads %+v", mt)
	}

	w.as(w.officer)
	w.h.data(w.h.do(http.MethodDelete, slotPath(w, raid.ID, "off_tank"), ""), &v)
	if ot := slotNamed(t, v, "off_tank"); len(ot.Names) != 0 {
		t.Fatalf("cleared off_tank = %+v", ot)
	}
	w.h.data(w.h.do(http.MethodPost, w.path("/raids/%d/cancel", raid.ID), `{}`), &v)
	expectError(t, w.h.do(http.MethodPut, slotPath(w, raid.ID, "main_tank"), `{"character_keys":[]}`), http.StatusConflict, "conflict")
}

func TestRaidWarningLineStopsAtTheChatLimit(t *testing.T) {
	names := []AssignedName{}
	for i := 0; i < 40; i++ {
		names = append(names, AssignedName{AssignmentRaider: AssignmentRaider{Name: "Abcdefghijkl"}})
	}
	line := raidWarningLine(slotTemplate{Warning: "P2 west"}, names)
	if len(line) > maxRaidWarningBytes || !strings.HasPrefix(line, "P2 west: Abcdefghijkl, ") || strings.HasSuffix(line, ", ") {
		t.Fatalf("line (%d) = %q", len(line), line)
	}
}

func TestPostAssignmentsRepliesUnderTheRaidPostPingingNoOne(t *testing.T) {
	w, raid := assignWorld(t)
	post := w.path("/raids/%d/assignments/post", raid.ID)
	expectError(t, w.h.do(http.MethodPost, post, ""), http.StatusConflict, "discord_not_connected")
	pinger := &fakePinger{posted: true}
	w.h.svc.RaidPings = pinger
	expectError(t, w.h.do(http.MethodPost, post, ""), http.StatusConflict, "nothing_assigned")

	w.h.do(http.MethodPut, slotPath(w, raid.ID, "main_tank"), `{"character_keys":["`+memberKeyA+`"]}`)
	var res AssignmentPostResult
	w.h.data(w.h.do(http.MethodPost, post, ""), &res)
	if !res.Posted || res.PostedAt == nil || res.NextAllowedAt == nil || len(pinger.pings) != 1 {
		t.Fatalf("post = %+v, pings %d", res, len(pinger.pings))
	}
	p := pinger.pings[0]
	if p.RaidID != raid.ID || len(p.MentionIDs) != 0 || !strings.Contains(p.Content, "Main tank: Kraggor") ||
		!strings.Contains(p.Content, "Off tank (phase 3, fear): open") || !strings.Contains(p.Content, "**Onyxia**") {
		t.Fatalf("ping = %+v", p)
	}
	if v := w.getRaid(t, raid.ID); v.Assignments.PostedAt == nil {
		t.Fatal("posted_at not set")
	}
	expectError(t, w.h.do(http.MethodPost, post, ""), http.StatusTooManyRequests, "assignments_throttled")
}

func TestPostAssignmentsGivesTheThrottleBackWhenDiscordRefuses(t *testing.T) {
	w, raid := assignWorld(t)
	w.h.svc.RaidPings = &fakePinger{err: errors.New("boom")}
	w.h.do(http.MethodPut, slotPath(w, raid.ID, "main_tank"), `{"character_keys":["`+memberKeyA+`"]}`)
	post := w.path("/raids/%d/assignments/post", raid.ID)
	expectError(t, w.h.do(http.MethodPost, post, ""), http.StatusBadGateway, "discord_refused")
	w.h.svc.RaidPings = &fakePinger{posted: true}
	if res := w.h.do(http.MethodPost, post, ""); res.StatusCode != http.StatusOK {
		t.Fatalf("retry = %d: %s", res.StatusCode, readBody(res))
	}
}

func TestNextRaidAssignmentsNamesTheCharactersOwnSlots(t *testing.T) {
	w, raid := assignWorld(t)
	w.h.do(http.MethodPut, slotPath(w, raid.ID, "main_tank_heals"), `{"character_keys":["`+member2Key+`"]}`)
	w.h.do(http.MethodPut, slotPath(w, raid.ID, "off_tank"), `{"character_keys":["`+officerKey+`","`+member2Key+`"]}`)
	got, err := w.h.store.NextRaidAssignments(t.Context(), []string{member2Key, memberKeyB}, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].SlotKey != "off_tank" || got[1].SlotKey != "main_tank_heals" {
		t.Fatalf("assignments = %+v", got)
	}
	if got[0].RaidID != raid.ID || got[0].Encounter != "Onyxia" || got[0].ZoneName != "Onyxia's Lair" ||
		len(got[0].With) != 1 || got[0].With[0] != "Thornhide" || got[0].Warning != "OT" || len(got[1].With) != 0 {
		t.Fatalf("first = %+v", got[0])
	}
	// Once the raid is over there is nothing to show.
	later, err := w.h.store.NextRaidAssignments(t.Context(), []string{member2Key}, raid.StartsAt.Add(4*time.Hour))
	if err != nil || len(later) != 0 {
		t.Fatalf("after the raid = %+v, %v", later, err)
	}
}
