package guilds

import (
	"testing"
	"time"
)

// An alt a signed main came as earns no night of its own in member records: one appearance
// counts once, on the main.
func TestAnAltTheMainCameAsCountsNoNightOfItsOwn(t *testing.T) {
	h, gid, uid, key := attendanceWorld(t)
	const alt = "us/hardcore/raidalt"
	raidID := seedDoneRaid(t, h, gid, uid, time.Now().Add(-48*time.Hour),
		seedAtt{key: key, signed: SignupIn, outcome: OutcomePresent},
		seedAtt{key: alt, signed: SignupBench, outcome: OutcomePresent})
	if _, err := h.pool.Exec(t.Context(),
		`update guild_raid_attendance set came_as = 'Raidalt', came_as_key = $2 where raid_id = $1 and character_key = $3`,
		raidID, alt, key); err != nil {
		t.Fatal(err)
	}
	got, err := h.store.raidAttendance(t.Context(), gid, []string{key, alt}, attendanceWindow)
	if err != nil {
		t.Fatal(err)
	}
	if got[key] != (Attendance{Present: 1, Nights: 1}) {
		t.Fatalf("main = %+v, want 1 of 1", got[key])
	}
	if a, ok := got[alt]; ok {
		t.Fatalf("alt = %+v, want no night", a)
	}
}
