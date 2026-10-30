package events

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// maxPrefsBody bounds a PATCH body: every event on every channel fits well inside it.
const maxPrefsBody = 16 << 10

// PrefsStore is what the preference routes read and write (Store).
type PrefsStore interface {
	Prefs(ctx context.Context, userID int64) (Prefs, error)
	SaveSettings(ctx context.Context, userID int64, change SettingsChange) error
}

// Handler serves GET and PATCH /v1/me/notifications.
type Handler struct {
	Store PrefsStore
	Log   *slog.Logger
}

// Mount adds the preference routes. A session or a device token reaches them.
func Mount(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("GET /v1/me/notifications", auth.Require(h.get))
	mux.HandleFunc("PATCH /v1/me/notifications", auth.Require(h.patch))
}

func (h *Handler) logger() *slog.Logger {
	if h.Log != nil {
		return h.Log
	}
	return slog.Default()
}

// QuietHoursView is the quiet hours as "HH:MM" in the player's zone.
type QuietHoursView struct {
	Enabled bool   `json:"enabled"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

// EventPrefView is one event's switches.
type EventPrefView struct {
	Kind        Kind             `json:"kind"`
	Who         string           `json:"who"`
	Officer     bool             `json:"officer"`
	QuietExempt bool             `json:"quiet_exempt"`
	Channels    map[Channel]bool `json:"channels"`
}

// PrefsView is the GET body.
type PrefsView struct {
	TimeZone   string          `json:"time_zone"`
	QuietHours QuietHoursView  `json:"quiet_hours"`
	Events     []EventPrefView `json:"events"`
}

// ViewOf is p as the routes show it, every catalogue event in its order.
func ViewOf(p Prefs) PrefsView {
	zone := p.TimeZone
	if zone == "" {
		zone = DefaultTimeZone
	}
	v := PrefsView{
		TimeZone:   zone,
		QuietHours: QuietHoursView{Enabled: p.Quiet.Enabled, Start: clock(p.Quiet.Start), End: clock(p.Quiet.End)},
		Events:     make([]EventPrefView, len(Catalogue)),
	}
	for i, s := range Catalogue {
		channels := make(map[Channel]bool, len(Channels))
		for _, ch := range Channels {
			channels[ch] = p.Allows(s.Kind, ch)
		}
		v.Events[i] = EventPrefView{Kind: s.Kind, Who: s.Who, Officer: s.Officer, QuietExempt: s.QuietExempt, Channels: channels}
	}
	return v
}

func clock(minutes int) string { return fmt.Sprintf("%02d:%02d", minutes/60, minutes%60) }

// parseClock reads "HH:MM" as minutes after midnight.
func parseClock(s string) (int, bool) {
	var hh, mm int
	if len(s) != 5 || s[2] != ':' {
		return 0, false
	}
	if _, err := fmt.Sscanf(s, "%02d:%02d", &hh, &mm); err != nil || hh > 23 || mm > 59 || hh < 0 || mm < 0 {
		return 0, false
	}
	return hh*60 + mm, true
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	p, err := h.Store.Prefs(r.Context(), auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		h.logger().Error("events", "op", "prefs_get", "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not read your notification settings", nil)
		return
	}
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, http.StatusOK, ViewOf(p))
}

// patchRequest is the PATCH body; an absent field keeps the stored value.
type patchRequest struct {
	TimeZone   *string         `json:"time_zone"`
	QuietHours *QuietHoursView `json:"quiet_hours"`
	Events     []struct {
		Kind    Kind    `json:"kind"`
		Channel Channel `json:"channel"`
		Enabled *bool   `json:"enabled"`
	} `json:"events"`
}

func (h *Handler) patch(w http.ResponseWriter, r *http.Request) {
	var req patchRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxPrefsBody)).Decode(&req); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	change, fields := req.change()
	if fields != nil {
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "invalid", "some settings are not valid", fields)
		return
	}
	uid := auth.ActorFrom(r.Context()).UserID
	if err := h.Store.SaveSettings(r.Context(), uid, change); err != nil {
		if errors.Is(err, errEmptyChange) {
			httpx.WriteError(w, r, http.StatusUnprocessableEntity, "invalid", "nothing to change", nil)
			return
		}
		h.logger().Error("events", "op", "prefs_patch", "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not save your notification settings", nil)
		return
	}
	h.get(w, r)
}

// change validates the request into a SettingsChange; fields names what is wrong.
func (req patchRequest) change() (SettingsChange, map[string]string) {
	var c SettingsChange
	fields := map[string]string{}
	if req.TimeZone != nil {
		if !ValidTimeZone(*req.TimeZone) {
			fields["time_zone"] = "not a time zone"
		}
		c.TimeZone = req.TimeZone
	}
	if q := req.QuietHours; q != nil {
		start, okStart := parseClock(q.Start)
		end, okEnd := parseClock(q.End)
		quiet := QuietHours{Enabled: q.Enabled, Start: start, End: end}
		if !okStart || !okEnd || quiet.Validate() != nil {
			fields["quiet_hours"] = "start and end are HH:MM and differ"
		}
		c.Quiet = &quiet
	}
	if len(req.Events) > len(Catalogue)*len(Channels) {
		fields["events"] = "too many changes"
	}
	for i, e := range req.Events {
		if _, ok := Lookup(e.Kind); !ok || !KnownChannel(e.Channel) || e.Enabled == nil {
			fields[fmt.Sprintf("events[%d]", i)] = "a known kind, channel and enabled"
			continue
		}
		c.Prefs = append(c.Prefs, PrefChange{Kind: e.Kind, Channel: e.Channel, Enabled: *e.Enabled})
	}
	if len(fields) > 0 {
		return SettingsChange{}, fields
	}
	return c, nil
}
