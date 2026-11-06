package guilds

import (
	"net/http"
	"testing"
)

func TestZoneCompDefaults(t *testing.T) {
	// Every first-tier raid (9 December) has a default that fills its size exactly.
	want := map[string]RoleCounts{
		"raid:barrow-deeps": {Tanks: 1, Healers: 3, Melee: 3, Ranged: 3},
		"raid:hyjal-summit": {Tanks: 2, Healers: 5, Melee: 6, Ranged: 7},
		raidZoneOny:         {Tanks: 2, Healers: 8, Melee: 10, Ranged: 20},
	}
	for key, w := range want {
		z, _ := ZoneByKey(key)
		got, ok := zoneCompDefault(z)
		if !ok || got != w || z.Size == nil || got.total() != *z.Size {
			t.Errorf("%s default = %+v, %v (size %v)", key, got, ok, z.Size)
		}
	}
	// A raid outside the first tier has no default, whatever its size.
	if _, ok := zoneCompDefault(Zone{Key: "raid:unannounced", Kind: ZoneRaid, Size: intPtr(40)}); ok {
		t.Error("an unlisted raid has a default")
	}
	for _, key := range []string{"dungeon:scholomance", "event:guild-40"} {
		z, _ := ZoneByKey(key)
		if _, ok := zoneCompDefault(z); ok {
			t.Errorf("%s has a default", key)
		}
		if compTargetOf(z, &RoleCounts{Tanks: 1}) != nil {
			t.Errorf("%s has a target", key)
		}
	}
}

func TestValidateCompTarget(t *testing.T) {
	ony, _ := ZoneByKey(raidZoneOny)
	cases := map[string]struct {
		c     RoleCounts
		field string
	}{
		"fits":     {RoleCounts{Tanks: 3, Healers: 9, Melee: 10, Ranged: 18}, ""},
		"open":     {RoleCounts{Tanks: 2, Healers: 8, Melee: 8, Ranged: 18}, ""},
		"negative": {RoleCounts{Tanks: -1, Healers: 8}, "tanks"},
		"too many": {RoleCounts{Tanks: 3, Healers: 10, Melee: 10, Ranged: 20}, "total"},
		"nobody":   {RoleCounts{}, "total"},
	}
	for name, tc := range cases {
		fields := validateCompTarget(ony, tc.c)
		if tc.field == "" && fields != nil {
			t.Errorf("%s: fields = %v", name, fields)
		}
		if tc.field != "" && fields[tc.field] == "" {
			t.Errorf("%s: fields = %v, want %s", name, fields, tc.field)
		}
	}
}

func TestCompTargetDefaultEditAndClear(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	if raid.CompTarget == nil || raid.CompTarget.Source != CompSourceDefault || raid.CompTarget.Healers != 8 {
		t.Fatalf("created comp_target = %+v", raid.CompTarget)
	}
	path := w.path("/raids/%d/comp-target", raid.ID)

	w.as(w.member)
	expectError(t, w.h.do(http.MethodPut, path, `{"tanks":3,"healers":9,"melee":10,"ranged":18}`), http.StatusForbidden, "forbidden")
	expectError(t, w.h.do(http.MethodDelete, path, ""), http.StatusForbidden, "forbidden")

	w.as(w.officer)
	expectError(t, w.h.do(http.MethodPut, path, `{"tanks":3,"healers":10,"melee":10,"ranged":20}`), http.StatusUnprocessableEntity, "invalid")
	expectError(t, w.h.do(http.MethodPut, path, `{"tanks":3,"healers":9}`), http.StatusUnprocessableEntity, "invalid")
	res := w.h.do(http.MethodPut, path, `{"tanks":3,"healers":9,"melee":10,"ranged":18}`)
	var v RaidView
	w.h.data(res, &v)
	if v.CompTarget == nil || v.CompTarget.Source != CompSourceRaid || v.CompTarget.Tanks != 3 || v.CompTarget.Ranged != 18 {
		t.Fatalf("set comp_target = %+v", v.CompTarget)
	}

	w.as(w.member)
	if got := w.getRaid(t, raid.ID).CompTarget; got == nil || got.Source != CompSourceRaid || got.Healers != 9 {
		t.Fatalf("member reads comp_target = %+v", got)
	}

	w.as(w.officer)
	w.h.data(w.h.do(http.MethodDelete, path, ""), &v)
	if v.CompTarget == nil || v.CompTarget.Source != CompSourceDefault || v.CompTarget.Tanks != 2 {
		t.Fatalf("cleared comp_target = %+v", v.CompTarget)
	}
}

func TestCompTargetRefusesARunWithoutOne(t *testing.T) {
	w := newRaidWorld(t, true)
	body := `{"zone":"dungeon:scholomance","title":"","starts_at":"` + futureStart() + `","duration_minutes":180,` +
		`"notes":"","timezone":"America/Chicago","repeat_weeks":0,"post_to_discord":false}`
	run := w.createRaids(t, body)[0]
	if run.CompTarget != nil {
		t.Fatalf("dungeon comp_target = %+v", run.CompTarget)
	}
	w.as(w.officer)
	expectError(t, w.h.do(http.MethodPut, w.path("/raids/%d/comp-target", run.ID), `{"tanks":1,"healers":1,"melee":1,"ranged":2}`),
		http.StatusConflict, "no_comp_target")
}
