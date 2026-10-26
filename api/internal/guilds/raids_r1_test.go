package guilds

import (
	"net/http"
	"testing"
)

// R1 on Lane G2's routes: anyone who is not a verified member of the guild (signed out, a
// visitor, an unverified member) is told 404, never 401 or 403; a verified member who is not an
// officer keeps 403.
func TestTheG2RoutesAnswer404ToAnyoneButAVerifiedMember(t *testing.T) {
	w, raid := matchedNight(t)
	future := w.createRaid(t)
	visitor := seedUser(t, w.h.pool, "visitor@example.com")
	unverified := w.seedRaider(t, "unverified@example.com", "member",
		raidChar{"us/hardcore/notyet", "Notyet", "rogue", false})
	routes := []struct{ method, path, body string }{
		{http.MethodGet, w.path("/raids/%d/match-candidates", raid.ID), ""},
		{http.MethodPut, w.path("/raids/%d/match", raid.ID), `{"report_id":"night"}`},
		{http.MethodDelete, w.path("/raids/%d/match", raid.ID), ""},
		{http.MethodPut, w.path("/raids/%d/attendance/%s", raid.ID, officerKey), `{"excused":true}`},
		{http.MethodPost, w.path("/raids/%d/report/post", raid.ID), ""},
		{http.MethodGet, w.path("/readiness?raid=%d", future.ID), ""},
	}
	for name, uid := range map[string]int64{"anonymous": 0, "visitor": visitor, "unverified": unverified} {
		w.as(uid)
		for _, rt := range routes {
			res := w.h.do(rt.method, rt.path, rt.body)
			if res.StatusCode != http.StatusNotFound {
				t.Errorf("%s %s %s = %d, want 404", name, rt.method, rt.path, res.StatusCode)
			}
			res.Body.Close()
		}
	}
	w.as(w.member)
	for _, rt := range routes[:5] {
		expectError(t, w.h.do(rt.method, rt.path, rt.body), http.StatusForbidden, "forbidden")
	}
}
