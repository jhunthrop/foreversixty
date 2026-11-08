package guilds

import (
	"net/http"
	"testing"
	"time"
)

// The officer's not-answered rows carry the account's answer for the guild's raid of the week
// before (missed when it gave none) and the row's role.
func TestNotSignedRowsCarryLastWeekAndRole(t *testing.T) {
	w := newRaidWorld(t, true)
	first := time.Now().Add(3 * 24 * time.Hour).UTC().Truncate(time.Minute)
	earlier := w.createRaids(t, raidBody(first.Format(time.RFC3339), 0))[0]
	later := w.createRaids(t, raidBody(first.Add(7*24*time.Hour).Format(time.RFC3339), 0))[0]
	w.as(w.member)
	if res := w.putSignup(earlier.ID, memberKeyA, `{"state":"late"}`); res.StatusCode != http.StatusOK {
		t.Fatalf("sign-up = %d", res.StatusCode)
	}
	w.as(w.officer)
	v := w.getRaid(t, later.ID)
	if v.NotSignedRaiders == nil {
		t.Fatal("no not-signed rows for the officer")
	}
	got := map[string]NotSignedRaider{}
	for _, r := range *v.NotSignedRaiders {
		got[r.AccountKey] = r
	}
	m, m2 := got[accountKeyFor(w.member)], got[accountKeyFor(w.member2)]
	if m.LastWeek == nil || *m.LastWeek != "late" || m2.LastWeek == nil || *m2.LastWeek != lastWeekMissed {
		t.Fatalf("last week %+v / %+v", m, m2)
	}
	// A priest with no spec on file stands in the ranged group, a damage role.
	if m2.Role != RoleDPS {
		t.Fatalf("Mirelle's role %q", m2.Role)
	}
	// The first raid has no raid the week before it.
	for _, r := range *w.getRaid(t, earlier.ID).NotSignedRaiders {
		if r.LastWeek != nil {
			t.Fatalf("first raid last week %+v", r)
		}
	}
}
