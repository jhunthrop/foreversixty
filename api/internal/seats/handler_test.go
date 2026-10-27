package seats

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func seatPath(gid int64, tail string) string { return fmt.Sprintf("/v1/guilds/%d/seats/%s", gid, tail) }

type seatCall struct {
	status int
	code   string
	data   map[string]any
}

func (h *harness) call(t *testing.T, method string, gid int64, tail string, uid int64, body any) seatCall {
	t.Helper()
	rec := h.request(method, seatPath(gid, tail), uid, body)
	c := seatCall{status: rec.Code}
	if rec.Code == http.StatusNoContent {
		return c
	}
	env := envelope(t, rec)
	if env.Error != nil {
		c.code = env.Error.Code
	}
	c.data, _ = env.Data.(map[string]any)
	return c
}

func TestHandler_StatusCodes(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	off := h.member(t, gid, "Off", "officer")
	bob := h.member(t, gid, "Bob", "member")
	cy := h.member(t, gid, "Cy", "member")
	amy := h.member(t, gid, "Amy", "member")
	outsider := h.user(t, "out")
	h.payingPremium(t, amy, testNow.Add(24*time.Hour))
	to := func(id int64) map[string]int64 { return map[string]int64{"user_id": id} }
	none := map[string]int{}

	expect := func(name string, c seatCall, status int, code string) {
		t.Helper()
		if c.status != status || c.code != code {
			t.Fatalf("%s: %d %q, want %d %q", name, c.status, c.code, status, code)
		}
	}

	expect("non-payer", h.call(t, "POST", gid, "1/give", off, to(bob)), 403, "forbidden")
	expect("not member", h.call(t, "POST", gid, "1/give", leader, to(outsider)), 409, "not_member")
	expect("no slot", h.call(t, "POST", gid, "9/give", leader, to(bob)), 409, "seat_count")
	given := h.call(t, "POST", gid, "1/give", leader, to(bob))
	expect("give", given, 200, "")
	if given.data["state"] != "held" {
		t.Fatalf("body = %v", given.data)
	}
	expect("taken", h.call(t, "POST", gid, "1/give", leader, to(cy)), 409, "conflict")
	expect("locked", h.call(t, "POST", gid, "1/move", leader, to(cy)), 409, "seat_locked")
	rec := h.request("POST", seatPath(gid, "1/move"), leader, to(cy))
	if env := envelope(t, rec); env.Error.Fields["movable_at"] == "" {
		t.Fatalf("no movable_at: %+v", env.Error)
	}
	expect("withdraw held", h.call(t, "POST", gid, "1/withdraw", leader, none), 404, "not_found")
	expect("no offer", h.call(t, "POST", gid, "offer/accept", cy, none), 404, "not_found")

	cands := h.call(t, "GET", gid, "candidates", leader, nil)
	expect("candidates", cands, 200, "")
	if l, _ := cands.data["candidates"].([]any); len(l) != 4 {
		t.Fatalf("candidates = %v", cands.data)
	}
	expect("candidates non-payer", h.call(t, "GET", gid, "candidates", bob, nil), 403, "forbidden")

	h.exec(t, `delete from premium_seats`)
	expect("offer", h.call(t, "POST", gid, "1/give", leader, to(amy)), 200, "")
	expect("decline", h.call(t, "POST", gid, "offer/decline", amy, none), 204, "")
	expect("offer again", h.call(t, "POST", gid, "1/give", leader, to(amy)), 200, "")
	h.svc.Canceller = nil
	expect("accept no stripe", h.call(t, "POST", gid, "offer/accept", amy, none), 503, "billing_unavailable")
	h.svc.Canceller = h.cancel
	expect("accept", h.call(t, "POST", gid, "offer/accept", amy, none), 204, "")

	h.svc.Guilds = fakeFrozen{frozen: true}
	expect("frozen", h.call(t, "POST", gid, "1/give", leader, to(bob)), 409, "claim_contested")

	h.svc.Guilds = nil
	h.exec(t, `update entitlements set status = 'canceled' where guild_id = $1`, gid)
	expect("no plan", h.call(t, "POST", gid, "1/give", leader, to(bob)), 409, "no_plan")
}
