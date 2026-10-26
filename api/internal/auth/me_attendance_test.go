package auth

import (
	"context"
	"net/http"
	"testing"
)

// fakeGuildAttendance answers GuildAttendanceReader from a table keyed by character key; an
// empty table is a guild still counting reports.
type fakeGuildAttendance struct{ byKey map[string]Attendance }

func (f *fakeGuildAttendance) CharacterAttendance(_ context.Context, _ int64, keys []string) (map[string]Attendance, error) {
	out := map[string]Attendance{}
	for _, k := range keys {
		if a, ok := f.byKey[k]; ok {
			out[k] = a
		}
	}
	return out, nil
}

func TestMeCarriesAttendanceOnlyInRaidMode(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "me-attendance@example.com")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`insert into characters (key, region, ruleset, name, class, user_id, realm_name, level, faction, source, imported_at, refreshed_at)
		 values ('us/pvp/attender', 'us', 'pvp', 'Attender', 'warrior', $1, 'Whitemane', 60, 'alliance', 'bnet', now(), now())`,
		uid); err != nil {
		t.Fatal(err)
	}
	gid := h.seedGuild(t, "Attendance Guild")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`insert into guild_characters (guild_id, character_key, user_id, rank, rank_index, source, verified_by, verified_at)
		 values ($1, 'us/pvp/attender', $2, 'member', 3, 'bnet', 'bnet', now())`, gid, uid); err != nil {
		t.Fatal(err)
	}
	read := func() *Attendance {
		res := h.sessionRequest(t, http.MethodGet, "/v1/me", uid)
		var body struct {
			Characters []struct {
				Guild *struct {
					Attendance *Attendance `json:"attendance"`
				} `json:"guild"`
			} `json:"characters"`
		}
		h.decode(t, res, &body)
		if len(body.Characters) != 1 || body.Characters[0].Guild == nil {
			t.Fatalf("characters = %+v, want one guilded character", body.Characters)
		}
		return body.Characters[0].Guild.Attendance
	}

	if got := read(); got != nil {
		t.Fatalf("with no reader wired attendance = %+v, want absent", got)
	}
	fake := &fakeGuildAttendance{byKey: map[string]Attendance{}}
	h.svc.GuildAttendance = fake
	if got := read(); got != nil {
		t.Fatalf("in report mode attendance = %+v, want absent", got)
	}
	fake.byKey["us/pvp/attender"] = Attendance{Present: 7, Nights: 8}
	if got := read(); got == nil || *got != (Attendance{Present: 7, Nights: 8}) {
		t.Fatalf("in raid mode attendance = %+v, want 7 of 8", got)
	}
}
