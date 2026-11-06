package guilds

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
)

// itemKeys is the items' keys in order.
func itemKeys(items []ChecklistItem) []string {
	out := make([]string, len(items))
	for i, it := range items {
		out[i] = it.Key
	}
	return out
}

func itemOf(items []ChecklistItem, key string) ChecklistItem {
	for _, it := range items {
		if it.Key == key {
			return it
		}
	}
	return ChecklistItem{}
}

func TestBuildChecklistOrderAndVerdicts(t *testing.T) {
	sixty, fifty := 60, 50
	ony, _ := ZoneByKey(raidZoneOny)
	ubrs, _ := ZoneByKey("dungeon:upper-blackrock-spire")
	band := bis.Band{Slots: []bis.Slot{{Slot: "hands", ItemID: 10}, {Slot: "finger_1", ItemID: 20}, {Slot: "chest", ItemID: 30}}}
	gear := map[string]int{"hands": 11, "finger_1": 21, "chest": 30}
	geared := RosterRow{Consent: "gear_bags", gear: gear, bags: []int{12344}}
	ready := CharacterReadiness{Level: &sixty, GearChecked: true, EnchantsChecked: true}
	upgrades := ready
	upgrades.GearGap = bis.GearGap{Upgrades: 2}
	upgrades.GearUnit = gearUnitDPS
	cases := map[string]struct {
		in      checklistInput
		verdict string
		keys    int
	}{
		"ready":             {checklistInput{zone: &ony, readiness: ready, row: geared, onRoster: true}, VerdictReady, 6},
		"two upgrades":      {checklistInput{zone: &ony, readiness: upgrades, row: geared, onRoster: true, band: band, hasBand: true}, VerdictAlmost, 6},
		"leveling":          {checklistInput{zone: &ony, readiness: CharacterReadiness{Level: &fifty, Leveling: true}}, VerdictGettingStarted, 6},
		"no export":         {checklistInput{zone: &ony}, VerdictGettingStarted, 6},
		"gear not shared":   {checklistInput{zone: &ony, readiness: CharacterReadiness{Level: &sixty}, row: RosterRow{Consent: "roster"}, onRoster: true}, VerdictAlmost, 6},
		"no raid planned":   {checklistInput{readiness: ready, row: geared, onRoster: true}, VerdictReady, 4},
		"keyed run, key on": {checklistInput{zone: &ubrs, readiness: ready, row: geared, onRoster: true}, VerdictReady, 5},
	}
	for name, tc := range cases {
		items, verdict := buildChecklist(tc.in)
		if verdict != tc.verdict || len(items) != tc.keys {
			t.Errorf("%s: verdict %s, items %v", name, verdict, itemKeys(items))
		}
		// An export with no raid-prep sections: flask and food, and fire resistance where the
		// zone asks for it, stay not tracked and never fail.
		last := CheckFlaskFood
		if tc.in.zone != nil && tc.in.zone.Key == raidZoneOny {
			last = CheckFireResistance
		}
		if items[0].Key != CheckLevel || items[len(items)-1].Key != last {
			t.Errorf("%s: order %v", name, itemKeys(items))
		}
		for _, key := range []string{CheckFlaskFood, CheckFireResistance} {
			if it := itemOf(items, key); it.Key != "" && it.State != CheckNotTracked {
				t.Errorf("%s: %s tracked: %+v", name, key, it)
			}
		}
	}
	items, _ := buildChecklist(checklistInput{zone: &ony, readiness: upgrades, row: geared, onRoster: true, band: band, hasBand: true})
	gap := itemOf(items, CheckGearGap)
	if gap.State != CheckFail || *gap.Upgrades != 2 || fmt.Sprint(gap.Slots) != "[hands finger_1]" || *gap.Unit != gearUnitDPS {
		t.Fatalf("gear gap = %+v", gap)
	}
	items, _ = buildChecklist(checklistInput{zone: &ubrs, readiness: ready, row: RosterRow{Consent: "gear"}, onRoster: true})
	if a := itemOf(items, CheckAttunement); a.State != CheckNotTracked || a.Reason != reasonKeysNotInExport || a.KeyItem.ID != 12344 {
		t.Fatalf("attunement without the key on file = %+v", a)
	}
	if lv := itemOf(items, CheckLevel); *lv.RequiredLevel != 59 {
		t.Fatalf("UBRS level = %+v", lv)
	}
}

func TestChecklistIsTheViewersOwn(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	seedLevelExport(t, w.h, w.member, memberKeyB, "Sunderfel", 60)
	w.as(w.member)
	if res := w.putSignup(raid.ID, memberKeyB, `{"state":"in"}`); res.StatusCode != http.StatusOK {
		t.Fatalf("signup = %d: %s", res.StatusCode, readBody(res))
	}

	var v ChecklistView
	w.h.data(w.h.do(http.MethodGet, w.path("/me/checklist"), ""), &v)
	if v.Character.Name != "Sunderfel" || v.Raid == nil || v.Raid.ID != raid.ID || v.Raid.ZoneKey != raidZoneOny {
		t.Fatalf("member checklist = %+v", v)
	}
	if lv := itemOf(v.Items, CheckLevel); lv.State != CheckOK || *lv.Level != 60 {
		t.Fatalf("level = %+v", lv)
	}
	if v.Verdict == VerdictReady {
		t.Fatalf("verdict %s with no gear on file", v.Verdict)
	}

	// Another member reads their own character, never the first member's; there is no way to
	// name someone else's.
	w.as(w.member2)
	w.h.data(w.h.do(http.MethodGet, w.path("/me/checklist?raid=%d", raid.ID), ""), &v)
	if v.Character.CharacterKey != member2Key {
		t.Fatalf("member2 reads %+v", v.Character)
	}
	if v.Verdict != VerdictGettingStarted {
		t.Fatalf("no level on file reads %s", v.Verdict)
	}

	outsider := seedUser(t, w.h.pool, "outsider@example.com")
	w.as(outsider)
	expectError(t, w.h.do(http.MethodGet, w.path("/me/checklist"), ""), http.StatusForbidden, "forbidden")

	w.as(w.member)
	expectError(t, w.h.do(http.MethodGet, w.path("/me/checklist?raid=999999"), ""), http.StatusNotFound, "not_found")
}

func TestChecklistWithNoRaidPlanned(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.member2)
	var v ChecklistView
	w.h.data(w.h.do(http.MethodGet, w.path("/me/checklist"), ""), &v)
	if v.Raid != nil || v.Character.Name != "Mirelle" || len(v.Items) != 4 {
		t.Fatalf("checklist = %+v", v)
	}
}

func TestOfficerReadinessRowsCarryAttuned(t *testing.T) {
	w := newRaidWorld(t, true)
	body := `{"zone":"dungeon:upper-blackrock-spire","title":"","starts_at":"` + futureStart() + `","duration_minutes":180,` +
		`"notes":"","timezone":"America/Chicago","repeat_weeks":0,"post_to_discord":false}`
	run := w.createRaids(t, body)[0]
	if _, err := w.h.pool.Exec(t.Context(), `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, 'us', 'hardcore', 'Mirelle', 'FS1:1.60.1.70009:priest:human:0/0/0:|level=60|bags=12344', now(), now())`,
		member2Key, w.member2); err != nil {
		t.Fatal(err)
	}
	if _, err := w.h.pool.Exec(t.Context(), `update guild_members set consent = 'gear_bags' where user_id = $1`, w.member2); err != nil {
		t.Fatal(err)
	}
	if res := w.signupAs(t, w.member2, run.ID, member2Key); res != http.StatusOK {
		t.Fatalf("signup = %d", res)
	}
	read := func(uid int64) map[string]*bool {
		w.as(uid)
		var v ReadinessView
		w.h.data(w.h.do(http.MethodGet, w.path("/readiness?raid=%d", run.ID), ""), &v)
		out := map[string]*bool{}
		for _, r := range v.Rows {
			out[r.CharacterKey] = r.Attuned
		}
		return out
	}
	if got := read(w.officer); got[member2Key] == nil || !*got[member2Key] {
		t.Fatalf("officer attuned = %+v", got)
	}
	for key, a := range read(w.member) {
		if a != nil {
			t.Fatalf("member reads attuned on %s", key)
		}
	}
}

// signupAs signs key up In for raidID as uid and returns the status.
func (w raidWorld) signupAs(t *testing.T, uid, raidID int64, key string) int {
	t.Helper()
	w.as(uid)
	res := w.putSignup(raidID, key, `{"state":"in"}`)
	defer res.Body.Close()
	return res.StatusCode
}
