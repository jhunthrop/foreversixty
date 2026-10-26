package guilds

import (
	"encoding/json"
	"net/http"
	"testing"
)

// homeNextRaid reads /home as uid and returns the raw next_raid, ok false when absent.
func (w raidWorld) homeNextRaid(t *testing.T, uid int64) (json.RawMessage, bool) {
	t.Helper()
	w.as(uid)
	res := w.h.do(http.MethodGet, w.path("/home"), "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("home = %d: %s", res.StatusCode, readBody(res))
	}
	raw, ok := rawData(t, res)["next_raid"]
	return raw, ok
}

func TestHomeNextRaid(t *testing.T) {
	w := newRaidWorld(t, true)
	if _, ok := w.homeNextRaid(t, w.member); ok {
		t.Fatal("next_raid present with no raid planned")
	}
	raid := w.createRaid(t)

	raw, ok := w.homeNextRaid(t, w.member)
	if !ok {
		t.Fatal("next_raid absent for a member with a planned raid")
	}
	var next NextRaidView
	if err := json.Unmarshal(raw, &next); err != nil {
		t.Fatal(err)
	}
	if next.ID != raid.ID || next.Own != nil || next.ZoneName != "Onyxia's Lair" || next.Size == nil || *next.Size != 40 {
		t.Fatalf("next_raid = %s", raw)
	}
	var keys map[string]json.RawMessage
	if err := json.Unmarshal(raw, &keys); err != nil {
		t.Fatal(err)
	}
	if string(keys["own"]) != "null" {
		t.Fatalf("own = %s, want null when unsigned", keys["own"])
	}

	w.as(w.member)
	w.signup(t, raid.ID, memberKeyB, `{"state":"in"}`)
	w.as(w.member2)
	w.signup(t, raid.ID, member2Key, `{"state":"late","role":"healer"}`)

	raw, _ = w.homeNextRaid(t, w.member)
	if err := json.Unmarshal(raw, &next); err != nil {
		t.Fatal(err)
	}
	if next.Own == nil || next.Own.CharacterKey != memberKeyB || next.Own.State != SignupIn || next.Own.Role != RoleDPS {
		t.Fatalf("own = %+v", next.Own)
	}
	if next.Status != RaidPlanned {
		t.Fatalf("status = %q, want planned", next.Status)
	}
	if next.Counts != (NextRaidCounts{In: 1, Late: 1, Tanks: 0, Healers: 0, DPS: 1}) {
		t.Fatalf("counts = %+v", next.Counts)
	}
}
