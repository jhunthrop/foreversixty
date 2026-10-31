package events

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

type memPrefs struct{ p map[int64]Prefs }

func (m *memPrefs) Prefs(_ context.Context, uid int64) (Prefs, error) {
	if p, ok := m.p[uid]; ok {
		return p, nil
	}
	return DefaultPrefs(), nil
}

func (m *memPrefs) SaveSettings(_ context.Context, uid int64, c SettingsChange) error {
	if c.TimeZone == nil && c.Quiet == nil && len(c.Prefs) == 0 {
		return errEmptyChange
	}
	p, _ := m.Prefs(context.Background(), uid)
	if c.TimeZone != nil {
		p.TimeZone = *c.TimeZone
	}
	if c.Quiet != nil {
		p.Quiet = *c.Quiet
	}
	for _, ch := range c.Prefs {
		p = p.With(ch.Kind, ch.Channel, ch.Enabled)
	}
	m.p[uid] = p
	return nil
}

type envelope struct {
	Data  PrefsView `json:"data"`
	Error *struct {
		Code   string            `json:"code"`
		Fields map[string]string `json:"fields"`
	} `json:"error"`
}

func call(t *testing.T, h *Handler, method, body string, uid int64) (int, envelope) {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, h)
	req := httptest.NewRequest(method, "/v1/me/notifications", strings.NewReader(body))
	if uid != 0 {
		req = req.WithContext(auth.WithActor(req.Context(), auth.Actor{UserID: uid, DeviceID: "d1"}))
	}
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	var env envelope
	if rec.Code != http.StatusUnauthorized {
		if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
			t.Fatalf("%s: %v", rec.Body.String(), err)
		}
	}
	return rec.Code, env
}

func eventView(v PrefsView, kind Kind) EventPrefView {
	for _, e := range v.Events {
		if e.Kind == kind {
			return e
		}
	}
	return EventPrefView{}
}

func TestGetShowsTheCatalogueDefaults(t *testing.T) {
	h := &Handler{Store: &memPrefs{p: map[int64]Prefs{}}}
	code, env := call(t, h, http.MethodGet, "", 5)
	if code != http.StatusOK || env.Data.TimeZone != "UTC" || env.Data.QuietHours != (QuietHoursView{Start: "22:00", End: "07:00"}) {
		t.Fatalf("%d %+v", code, env.Data)
	}
	if len(env.Data.Events) != len(Catalogue) {
		t.Fatalf("events %d", len(env.Data.Events))
	}
	if e := eventView(env.Data, RaidStarting); !e.QuietExempt || !e.Channels[ChannelPush] || e.Who != WhoInLate {
		t.Fatalf("raid.starting %+v", e)
	}
	if e := eventView(env.Data, BisChanged); e.Channels[ChannelPush] || e.Channels[ChannelDiscord] {
		t.Fatalf("bis.changed %+v", e)
	}
	if e := eventView(env.Data, RaidShort); !e.Officer {
		t.Fatalf("raid.short %+v", e)
	}
}

func TestPatchSavesAndAnswersTheNewSettings(t *testing.T) {
	h := &Handler{Store: &memPrefs{p: map[int64]Prefs{}}}
	body := `{"time_zone":"America/Chicago","quiet_hours":{"enabled":true,"start":"23:30","end":"06:15"},
		"events":[{"kind":"raid.posted","channel":"push","enabled":false},{"kind":"bis.changed","channel":"push","enabled":true}]}`
	code, env := call(t, h, http.MethodPatch, body, 5)
	if code != http.StatusOK {
		t.Fatalf("%d %+v", code, env.Error)
	}
	v := env.Data
	if v.TimeZone != "America/Chicago" || v.QuietHours != (QuietHoursView{Enabled: true, Start: "23:30", End: "06:15"}) {
		t.Fatalf("%+v", v)
	}
	if eventView(v, RaidPosted).Channels[ChannelPush] || !eventView(v, RaidPosted).Channels[ChannelDiscord] || !eventView(v, BisChanged).Channels[ChannelPush] {
		t.Fatalf("events %+v", v.Events)
	}
	// Another account is untouched.
	if _, other := call(t, h, http.MethodGet, "", 6); !eventView(other.Data, RaidPosted).Channels[ChannelPush] {
		t.Fatal("another account changed")
	}
}

func TestPatchRefusesWhatIsNotValid(t *testing.T) {
	h := &Handler{Store: &memPrefs{p: map[int64]Prefs{}}}
	cases := map[string]string{
		`{"time_zone":"Mars/Olympus"}`:                                       "time_zone",
		`{"quiet_hours":{"enabled":true,"start":"25:00","end":"07:00"}}`:     "quiet_hours",
		`{"quiet_hours":{"enabled":true,"start":"07:00","end":"07:00"}}`:     "quiet_hours",
		`{"events":[{"kind":"raid.nope","channel":"push","enabled":true}]}`:  "events[0]",
		`{"events":[{"kind":"raid.posted","channel":"sms","enabled":true}]}`: "events[0]",
		`{"events":[{"kind":"raid.posted","channel":"push"}]}`:               "events[0]",
	}
	for body, field := range cases {
		code, env := call(t, h, http.MethodPatch, body, 5)
		if code != http.StatusUnprocessableEntity || env.Error == nil || env.Error.Fields[field] == "" {
			t.Errorf("%s: %d %+v", body, code, env.Error)
		}
	}
	if code, _ := call(t, h, http.MethodPatch, `{}`, 5); code != http.StatusUnprocessableEntity {
		t.Fatalf("empty patch %d", code)
	}
	if code, _ := call(t, h, http.MethodPatch, `nope`, 5); code != http.StatusBadRequest {
		t.Fatalf("bad json %d", code)
	}
	if code, _ := call(t, h, http.MethodGet, "", 0); code != http.StatusUnauthorized {
		t.Fatalf("anonymous %d", code)
	}
}
