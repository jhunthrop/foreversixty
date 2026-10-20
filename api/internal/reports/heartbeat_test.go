package reports

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// A report with no fights answers its creation as last_upload_at and a null heartbeat_at; once
// fights are written it answers the newest fight's write; a heartbeat fills heartbeat_at in.
func TestReportMetaCarriesLastUploadAndHeartbeat(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	created := time.Now().Add(-2 * time.Hour).UTC().Truncate(time.Second)
	h.seedReport("hb-meta", "Tuesday", "Molten Core", Private, StatusLive, created)

	meta := func() View {
		var v View
		h.data(h.do(http.MethodGet, "/v1/reports/hb-meta", "", nil), &v)
		return v
	}

	v := meta()
	if v.LastUploadAt == nil || !v.LastUploadAt.Equal(created) || v.HeartbeatAt != nil {
		t.Fatalf("without fights: last_upload_at %v, heartbeat_at %v; want %v and null", v.LastUploadAt, v.HeartbeatAt, created)
	}

	h.reportID = "hb-meta"
	h.seedFight(0, 0, 0, 10)
	h.seedFight(1, 0, 10, 20)
	older := time.Now().Add(-90 * time.Minute).UTC().Truncate(time.Second)
	newer := time.Now().Add(-30 * time.Minute).UTC().Truncate(time.Second)
	for i, at := range []time.Time{older, newer} {
		if _, err := h.store.Pool.Exec(ctx, `update fights set uploaded_at = $1 where report_id = 'hb-meta' and fight_index = $2`, at, i); err != nil {
			t.Fatal(err)
		}
	}
	if v = meta(); v.LastUploadAt == nil || !v.LastUploadAt.Equal(newer) {
		t.Fatalf("with fights: last_upload_at = %v, want the newest fight's %v", v.LastUploadAt, newer)
	}

	// A fight sent again is the companion talking: its write moves last_upload_at.
	h.seedFight(0, 0, 0, 10)
	if v = meta(); v.LastUploadAt == nil || !v.LastUploadAt.After(newer) {
		t.Fatalf("after a re-sent fight: last_upload_at = %v, want later than %v", v.LastUploadAt, newer)
	}

	before := time.Now().Add(-time.Second)
	if res := h.do(http.MethodPost, "/v1/reports/hb-meta/heartbeat", "", nil); res.StatusCode != http.StatusNoContent {
		t.Fatalf("heartbeat status = %d, want 204", res.StatusCode)
	}
	if v = meta(); v.HeartbeatAt == nil || v.HeartbeatAt.Before(before) {
		t.Fatalf("heartbeat_at = %v, want a time just now", v.HeartbeatAt)
	}
}

// The heartbeat route is the owner's device only: a session or no actor is 401, another
// account's device sees 404 (the route does not say which ids exist), and a second beat inside
// the minute is 429.
func TestHeartbeatRoute(t *testing.T) {
	frozen := time.Now()
	httpx.Now = func() time.Time { return frozen }
	t.Cleanup(func() { httpx.Now = time.Now })

	h := newHarness(t)
	h.seedReport("hb-route", "Tuesday", "Molten Core", Public, StatusLive, time.Now())
	post := func() int {
		res := h.do(http.MethodPost, "/v1/reports/hb-route/heartbeat", "", nil)
		res.Body.Close()
		return res.StatusCode
	}

	owner := h.actor
	h.actor = auth.Actor{}
	if got := post(); got != http.StatusUnauthorized {
		t.Fatalf("no actor: status = %d, want 401", got)
	}
	h.asSession()
	if got := post(); got != http.StatusUnauthorized {
		t.Fatalf("browser session: status = %d, want 401 (device bearer only)", got)
	}

	h.actor = auth.Actor{UserID: h.owner + 1000, Role: "user", Method: "device", DeviceID: "device-other"}
	if got := post(); got != http.StatusNotFound {
		t.Fatalf("another account's device: status = %d, want 404", got)
	}

	// The limiter counts every request that reaches it, so each probe below that is not the owner's
	// real beat uses a device of its own.
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "device", DeviceID: "device-probe"}
	missing := h.do(http.MethodPost, "/v1/reports/no-such-report/heartbeat", "", nil)
	missing.Body.Close()
	if missing.StatusCode != http.StatusNotFound {
		t.Fatalf("unknown report: status = %d, want 404", missing.StatusCode)
	}
	h.actor = owner
	if got := post(); got != http.StatusNoContent {
		t.Fatalf("owner's device: status = %d, want 204", got)
	}
	if got := post(); got != http.StatusTooManyRequests {
		t.Fatalf("second beat inside a minute: status = %d, want 429", got)
	}
	// The budget is per device: the same account's other companion is not spent by this one.
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "device", DeviceID: "device-2"}
	if got := post(); got != http.StatusNoContent {
		t.Fatalf("a second device: status = %d, want 204", got)
	}
}
