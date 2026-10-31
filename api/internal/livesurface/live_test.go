package livesurface

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

var (
	iosLiveToken = strings.Repeat("ab", 32)
	fcmToken     = "fcm_registration:token-1234567890"
)

func raidBody(raidID int64, token string) string {
	return fmt.Sprintf(`{"kind":"raid","platform":"ios","push_token":%q,"object_id":"%d"}`, token, raidID)
}

func (h *harness) registered(raw []byte) int64 {
	h.t.Helper()
	var env struct{ Data Registered }
	if err := json.Unmarshal(raw, &env); err != nil || env.Data.ID == 0 {
		h.t.Fatalf("answer %s: %v", raw, err)
	}
	return env.Data.ID
}

func TestRegisterARaidCountdown(t *testing.T) {
	h := newHarness(t).withStore()
	end := testNow.Add(50 * time.Hour)
	h.raids.raidEnd[7] = end
	res, raw := h.do(http.MethodPost, "/v1/me/live-activities", raidBody(7, iosLiveToken))
	contract.AssertFixture(t, "live.register", h.want(res, raw, http.StatusCreated), contract.Options{})
	id := h.registered(raw)

	// The same token for the same raid refreshes its row.
	res, raw = h.do(http.MethodPost, "/v1/me/live-activities", raidBody(7, iosLiveToken))
	if again := h.registered(h.want(res, raw, http.StatusCreated)); again != id {
		t.Fatalf("re-register made row %d, want %d", again, id)
	}
	subs, err := h.svc.Tokens.Subscribers(t.Context(), KindRaid, "7", testNow)
	if err != nil || len(subs) != 1 || subs[0].PushToken != iosLiveToken || subs[0].UserID != h.actor.UserID {
		t.Fatalf("subscribers = %+v, %v", subs, err)
	}
	if subs, _ := h.svc.Tokens.Subscribers(t.Context(), KindRaid, "7", end.Add(raidTail)); len(subs) != 0 {
		t.Fatalf("still subscribed after the raid's end and tail: %+v", subs)
	}
	if n, err := h.svc.Tokens.DeleteExpired(t.Context(), end.Add(raidTail)); err != nil || n != 1 {
		t.Fatalf("expired = %d, %v", n, err)
	}
}

func TestRegisterRefusesRaidsTheCallerMayNotSee(t *testing.T) {
	h := newHarness(t).withStore()
	res, raw := h.do(http.MethodPost, "/v1/me/live-activities", raidBody(9, iosLiveToken))
	h.want(res, raw, http.StatusNotFound)
	h.raids.raidEnd[9] = testNow.Add(-time.Minute)
	res, raw = h.do(http.MethodPost, "/v1/me/live-activities", raidBody(9, iosLiveToken))
	h.want(res, raw, http.StatusNotFound)
}

func TestRegisterASimNeedsServerSimsAndYourOwnRun(t *testing.T) {
	h := newHarness(t).withStore()
	body := fmt.Sprintf(`{"kind":"sim","platform":"android","push_token":%q,"object_id":"k3v9q2m8x1ab"}`, fcmToken)
	res, raw := h.do(http.MethodPost, "/v1/me/live-activities", body)
	contract.AssertFixture(t, "live.register.error-premium-required",
		h.want(res, raw, http.StatusPaymentRequired), contract.Options{})

	h.ent.premium = true
	res, raw = h.do(http.MethodPost, "/v1/me/live-activities", body)
	h.want(res, raw, http.StatusNotFound)

	h.sims.owned["k3v9q2m8x1ab"] = true
	res, raw = h.do(http.MethodPost, "/v1/me/live-activities", body)
	h.registered(h.want(res, raw, http.StatusCreated))
}

func TestRegisterValidatesEveryField(t *testing.T) {
	h := newHarness(t)
	cases := []struct {
		name, body, field string
	}{
		{"kind", `{"kind":"party","platform":"ios","push_token":"` + iosLiveToken + `","object_id":"7"}`, "kind"},
		{"platform", `{"kind":"raid","platform":"web","push_token":"` + iosLiveToken + `","object_id":"7"}`, "platform"},
		{"ios token", `{"kind":"raid","platform":"ios","push_token":"not-hex","object_id":"7"}`, "push_token"},
		{"raid id", `{"kind":"raid","platform":"ios","push_token":"` + iosLiveToken + `","object_id":"seven"}`, "object_id"},
		{"start token on android", `{"kind":"raid","platform":"android","push_token":"` + fcmToken + `","push_to_start_token":"` + iosLiveToken + `","object_id":"7"}`, "push_to_start_token"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res, raw := h.do(http.MethodPost, "/v1/me/live-activities", c.body)
			h.want(res, raw, http.StatusUnprocessableEntity)
			var env struct {
				Error struct{ Fields map[string]string }
			}
			if err := json.Unmarshal(raw, &env); err != nil || env.Error.Fields[c.field] == "" {
				t.Fatalf("fields = %v, want %s: %s", env.Error.Fields, c.field, raw)
			}
		})
	}
	res, raw := h.do(http.MethodPost, "/v1/me/live-activities", "{not json")
	h.want(res, raw, http.StatusBadRequest)
}

func TestUnregisterTouchesOnlyYourOwn(t *testing.T) {
	h := newHarness(t).withStore()
	h.raids.raidEnd[7] = testNow.Add(time.Hour)
	res, raw := h.do(http.MethodPost, "/v1/me/live-activities", raidBody(7, iosLiveToken))
	id := h.registered(h.want(res, raw, http.StatusCreated))

	if err := h.svc.Tokens.Unregister(t.Context(), h.actor.UserID+1, id); err != nil {
		t.Fatal(err)
	}
	if subs, _ := h.svc.Tokens.Subscribers(t.Context(), KindRaid, "7", testNow); len(subs) != 1 {
		t.Fatal("another account removed the registration")
	}
	res, raw = h.do(http.MethodDelete, fmt.Sprintf("/v1/me/live-activities/%d", id), "")
	h.want(res, raw, http.StatusNoContent)
	res, raw = h.do(http.MethodDelete, fmt.Sprintf("/v1/me/live-activities/%d", id), "")
	h.want(res, raw, http.StatusNoContent)
	if subs, _ := h.svc.Tokens.Subscribers(t.Context(), KindRaid, "7", testNow); len(subs) != 0 {
		t.Fatalf("still registered: %+v", subs)
	}
}

func TestRegistrationsAreCappedPerAccount(t *testing.T) {
	h := newHarness(t).withStore()
	for i := int64(1); i <= MaxActivePerUser+2; i++ {
		h.raids.raidEnd[i] = testNow.Add(time.Hour)
		if _, err := h.svc.Tokens.Register(t.Context(), Registration{
			UserID: h.actor.UserID, Kind: KindRaid, Platform: PlatformIOS, ObjectID: fmt.Sprint(i),
			PushToken: iosLiveToken, ExpiresAt: testNow.Add(time.Hour),
		}); err != nil {
			t.Fatal(err)
		}
	}
	var n int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from live_activity_tokens`).Scan(&n); err != nil || n != MaxActivePerUser {
		t.Fatalf("rows = %d, %v; want %d", n, err, MaxActivePerUser)
	}
	if subs, _ := h.svc.Tokens.Subscribers(t.Context(), KindRaid, "1", testNow); len(subs) != 0 {
		t.Fatal("the oldest registration was kept")
	}
}
