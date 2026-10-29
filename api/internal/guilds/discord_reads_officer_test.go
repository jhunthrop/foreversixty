package guilds

import "testing"

func TestIsOfficerFollowsRank(t *testing.T) {
	w, _ := readinessRaidWorld(t)
	stranger := seedUser(t, w.h.pool, "nobody@example.com")
	for _, c := range []struct {
		name string
		id   int64
		want bool
	}{{"officer", w.officer, true}, {"member", w.member, false}, {"stranger", stranger, false}} {
		got, err := w.h.svc.IsOfficer(t.Context(), w.gid, c.id)
		if err != nil || got != c.want {
			t.Errorf("%s = %v %v, want %v", c.name, got, err, c.want)
		}
	}
}
