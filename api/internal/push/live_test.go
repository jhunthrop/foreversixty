package push

import (
	"context"
	"net/http"
	"testing"
)

// A Live Activity push goes on the liveactivity push type, to the bundle's liveactivity topic,
// with the payload as given.
func TestAPNsSendsALiveActivityPush(t *testing.T) {
	_, p8 := testP8(t)
	srv, calls, _ := apnsServer(t, http.StatusOK, "")
	s, err := NewAPNs(APNsConfig{KeyP8: p8, KeyID: "KEY123", TeamID: "TEAM45", Topic: "gg.foreversixty.app", BaseURL: srv.URL}, srv.Client())
	if err != nil {
		t.Fatal(err)
	}
	payload := map[string]any{"aps": map[string]any{"event": "start", "timestamp": 1}}
	if err := s.SendLive(context.Background(), "abcdef0123", payload); err != nil {
		t.Fatal(err)
	}
	c := (*calls)[0]
	for h, want := range map[string]string{"apns-topic": "gg.foreversixty.app.push-type.liveactivity",
		"apns-push-type": "liveactivity", "apns-priority": "10"} {
		if got := c.header.Get(h); got != want {
			t.Errorf("%s = %q, want %q", h, got, want)
		}
	}
	if c.path != "/3/device/abcdef0123" || c.body["aps"].(map[string]any)["event"] != "start" {
		t.Fatalf("call %+v", c)
	}
}
