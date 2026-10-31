package billing

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// The apps have no billing in v1: every billing write refuses an app token with a plain
// pointer to the site, while the reads an app may show stay open to it.
func TestBillingWritesRefuseAnAppToken(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "phone@example.com")
	gid := h.seedGuild(t, "Phone Guild", true)
	app := auth.Actor{UserID: uid, Role: "member", Method: "device", DeviceID: "d", DeviceKind: auth.DeviceKindApp}
	for _, path := range []string{
		"/v1/billing/checkout", "/v1/billing/portal",
		guildPath(gid, "/upgrade"), guildPath(gid, "/transfer"),
		guildPath(gid, "/transfer/withdraw"), guildPath(gid, "/transfer/decline"),
	} {
		r := h.sessionRequest(t, http.MethodPost, path, uid, map[string]string{})
		w := h.do(t, r.WithContext(auth.WithActor(r.Context(), app)))
		e := decodeEnvelope(t, w)
		if w.Code != http.StatusForbidden || e.Error == nil || e.Error.Code != "browser_session_required" {
			t.Errorf("POST %s as an app = %d %+v, want 403 browser_session_required", path, w.Code, e.Error)
		}
	}
	r := h.sessionRequest(t, http.MethodGet, guildPath(gid, ""), uid, nil)
	if w := h.do(t, r.WithContext(auth.WithActor(r.Context(), app))); w.Code == http.StatusForbidden || w.Code == http.StatusUnauthorized {
		t.Errorf("GET guild billing as an app = %d, want the read to reach its handler", w.Code)
	}
}

func guildPath(id int64, suffix string) string {
	return "/v1/guilds/" + strconv.FormatInt(id, 10) + "/billing" + suffix
}
