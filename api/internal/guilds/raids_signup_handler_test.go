package guilds

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

type signupAnswer struct {
	Raid  RaidView `json:"raid"`
	Moved *string  `json:"moved"`
}

func (w raidWorld) signup(t *testing.T, raidID int64, key, body string) signupAnswer {
	t.Helper()
	res := w.putSignup(raidID, key, body)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("PUT %s = %d: %s", key, res.StatusCode, readBody(res))
	}
	var a signupAnswer
	w.h.data(res, &a)
	return a
}

// TestSignupMovesTheOtherOwnCharacter is R3 over HTTP: the answer names the benched character.
func TestSignupMovesTheOtherOwnCharacter(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	if a := w.signup(t, raid.ID, memberKeyA, `{"state":"in","role":"tank"}`); a.Moved != nil {
		t.Fatalf("first in moved %q", *a.Moved)
	}
	a := w.signup(t, raid.ID, memberKeyB, `{"state":"in"}`)
	if a.Moved == nil || *a.Moved != "Kraggor" {
		t.Fatalf("moved = %v, want Kraggor", a.Moved)
	}
	states := map[string]string{}
	for _, s := range a.Raid.Signups {
		states[s.CharacterKey] = s.State
	}
	if states[memberKeyA] != SignupBench || states[memberKeyB] != SignupIn {
		t.Fatalf("states = %v", states)
	}
}

func TestSignupOwnership(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	cases := []struct {
		name     string
		actor    int64
		key      string
		status   int
		wantCode string
	}{
		{"member signs another member's character", w.member, member2Key, http.StatusForbidden, "forbidden"},
		{"member signs their unverified character", w.member, memberKeyC, http.StatusForbidden, "forbidden"},
		{"member signs a character not in the guild", w.member, "us/hardcore/nobody", http.StatusNotFound, "not_found"},
		{"officer signs anyone", w.officer, member2Key, http.StatusOK, ""},
		{"member signs their own", w.member, memberKeyA, http.StatusOK, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w.as(tc.actor)
			res := w.putSignup(raid.ID, tc.key, `{"state":"in"}`)
			if tc.wantCode != "" {
				expectError(t, res, tc.status, tc.wantCode)
				return
			}
			if res.StatusCode != tc.status {
				t.Fatalf("status = %d: %s", res.StatusCode, readBody(res))
			}
			res.Body.Close()
		})
	}
	t.Run("visitor and stranger get 404", func(t *testing.T) {
		stranger := seedUser(t, w.h.pool, "signup-stranger@example.com")
		w.as(stranger)
		expectError(t, w.putSignup(raid.ID, memberKeyA, `{"state":"in"}`), http.StatusNotFound, "not_found")
		w.as(0)
		res := w.putSignup(raid.ID, memberKeyA, `{"state":"in"}`)
		res.Body.Close()
		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("visitor = %d, want 401", res.StatusCode)
		}
	})
	t.Run("an officer's sign-up for a member counts against the member's account", func(t *testing.T) {
		w.as(w.officer)
		w.signup(t, raid.ID, officerKey, `{"state":"in"}`)
		a := w.signup(t, raid.ID, memberKeyB, `{"state":"in"}`)
		if a.Moved == nil || *a.Moved != "Kraggor" {
			t.Fatalf("moved = %v, want the member's Kraggor, never the officer's own", a.Moved)
		}
	})
}

// TestSignupRaces is Review Focus 3: a lock or cancel between page load and tap answers 409.
func TestSignupRaces(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.officer)
	res := w.h.do(http.MethodPost, w.path("/raids/%d/lock", raid.ID), "")
	res.Body.Close()
	w.as(w.member)
	expectError(t, w.putSignup(raid.ID, memberKeyA, `{"state":"in"}`), http.StatusConflict, "raid_locked")
	w.as(w.officer)
	if a := w.signup(t, raid.ID, memberKeyA, `{"state":"in"}`); a.Raid.Status != RaidLocked {
		t.Fatalf("officer on locked: status %q", a.Raid.Status)
	}
	res = w.h.do(http.MethodPost, w.path("/raids/%d/cancel", raid.ID), "")
	res.Body.Close()
	w.as(w.member)
	expectError(t, w.putSignup(raid.ID, memberKeyA, `{"state":"out"}`), http.StatusConflict, "raid_cancelled")
}

func TestSignupValidation(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	cases := []struct{ name, body, field string }{
		{"no state", `{}`, "state"},
		{"unknown state", `{"state":"maybe"}`, "state"},
		{"unknown role", `{"state":"in","role":"bard"}`, "role"},
		{"long note", fmt.Sprintf(`{"state":"in","note":%q}`, strings.Repeat("n", 141)), "note"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := expectError(t, w.putSignup(raid.ID, memberKeyA, tc.body), http.StatusUnprocessableEntity, "invalid")
			if _, ok := e.Fields[tc.field]; !ok {
				t.Fatalf("fields = %v, want %q", e.Fields, tc.field)
			}
		})
	}
	expectError(t, w.putSignup(raid.ID+999, memberKeyA, `{"state":"in"}`), http.StatusNotFound, "not_found")
}

// TestRaidViewShape checks counts, roster order, not_signed and the viewer's characters.
func TestRaidViewShape(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member2)
	w.signup(t, raid.ID, member2Key, `{"state":"in","role":"healer","note":"  Late by 10  "}`)
	w.as(w.member)
	w.signup(t, raid.ID, memberKeyA, `{"state":"bench","role":"tank"}`)
	v := w.signup(t, raid.ID, memberKeyB, `{"state":"late"}`).Raid

	want := RaidCounts{In: 1, Late: 1, Bench: 1, Out: 0, Tanks: 0, Healers: 1, DPS: 0}
	if v.Counts != want {
		t.Fatalf("counts = %+v, want %+v", v.Counts, want)
	}
	order := []string{}
	for _, s := range v.Signups {
		order = append(order, s.Name)
	}
	if strings.Join(order, ",") != "Kraggor,Mirelle,Sunderfel" {
		t.Fatalf("signup order = %v (tank, healer, dps)", order)
	}
	for _, s := range v.Signups {
		if s.Own != (s.CharacterKey == memberKeyA || s.CharacterKey == memberKeyB) {
			t.Fatalf("own wrong on %+v", s)
		}
		if s.CharacterKey == memberKeyB && (s.Role != RoleDPS || s.Class != "mage") {
			t.Fatalf("default role/class = %+v", s)
		}
		if s.CharacterKey == member2Key && s.Note != "Late by 10" {
			t.Fatalf("note = %q", s.Note)
		}
	}
	if len(v.NotSigned) != 1 || v.NotSigned[0].CharacterKey != officerKey || v.NotSigned[0].Name != "Thornhide" {
		t.Fatalf("not_signed = %+v (verified roster rows only)", v.NotSigned)
	}
	chars := map[string]ViewerCharacter{}
	for _, c := range v.Viewer.Characters {
		chars[c.CharacterKey] = c
	}
	if len(chars) != 3 || chars[memberKeyC].Verified || chars[memberKeyC].State != nil {
		t.Fatalf("viewer characters = %+v", v.Viewer.Characters)
	}
	if a := chars[memberKeyA]; !a.Verified || a.State == nil || *a.State != SignupBench || a.Role != RoleTank || a.Name != "Kraggor" {
		t.Fatalf("viewer Kraggor = %+v", a)
	}
}

func TestRoleForSpec(t *testing.T) {
	cases := map[string]string{
		"Protection": RoleTank, "protection": RoleTank, "Feral Combat": RoleTank, "Guardian": RoleTank,
		"Holy": RoleHealer, "Discipline": RoleHealer, "Restoration": RoleHealer,
		"Arms": RoleDPS, "Marksmanship": RoleDPS, "": RoleDPS,
	}
	for spec, want := range cases {
		if got := roleForSpec(spec); got != want {
			t.Errorf("roleForSpec(%q) = %q, want %q", spec, got, want)
		}
	}
}

func TestSignupNotYoursMessage(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	e := expectError(t, w.putSignup(raid.ID, member2Key, `{"state":"in"}`), http.StatusForbidden, "forbidden")
	if e.Message != "You can only sign up your own characters." {
		t.Fatalf("message = %q", e.Message)
	}
}

func TestSignupRateLimitIsPerAccount(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	for i := 0; i < signupsPerMinute; i++ {
		res := w.putSignup(raid.ID, memberKeyA, `{"state":"out"}`)
		res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Fatalf("PUT %d = %d", i, res.StatusCode)
		}
	}
	expectError(t, w.putSignup(raid.ID, memberKeyA, `{"state":"out"}`), http.StatusTooManyRequests, "rate_limited")
	w.as(w.member2)
	if a := w.signup(t, raid.ID, member2Key, `{"state":"in"}`); a.Raid.ID != raid.ID {
		t.Fatal("another account's budget was spent")
	}
}

// TestOfficerOwnSignupOnALockedRaid: an officer with the plan still answers for their own
// character on a locked raid, unless the claim is contested (officer powers frozen).
func TestOfficerOwnSignupOnALockedRaid(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.officer)
	res := w.h.do(http.MethodPost, w.path("/raids/%d/lock", raid.ID), "")
	res.Body.Close()
	if a := w.signup(t, raid.ID, officerKey, `{"state":"in"}`); a.Raid.Status != RaidLocked {
		t.Fatalf("status = %q", a.Raid.Status)
	}
	if _, err := w.h.pool.Exec(t.Context(), `update guilds set claim_contested_at = now() where id = $1`, w.gid); err != nil {
		t.Fatal(err)
	}
	expectError(t, w.putSignup(raid.ID, officerKey, `{"state":"late"}`), http.StatusConflict, "claim_contested")
	expectError(t, w.putSignup(raid.ID, member2Key, `{"state":"late"}`), http.StatusConflict, "claim_contested")
	w.as(w.member)
	expectError(t, w.putSignup(raid.ID, memberKeyA, `{"state":"in"}`), http.StatusConflict, "raid_locked")
}
