package livesurface

import (
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// The live-surface events (W3-C14). widget.refresh is silent and collapses per user; the
// live.* events update a Live Activity (iOS) or an ongoing notification (Android).
const (
	EventWidgetRefresh = "widget.refresh"
	EventRaidStart     = "live.raid.start"
	EventRaidUpdate    = "live.raid.update"
	EventRaidEnd       = "live.raid.end"
	EventSimUpdate     = "live.sim.update"
	EventSimEnd        = "live.sim.end"
)

// The ActivityKit attribute types the iOS widget extension declares; a push-to-start names one.
const (
	RaidActivityType = "RaidActivityAttributes"
	SimActivityType  = "SimActivityAttributes"
)

// Sim run states on a live surface.
const (
	SimRunning  = "running"
	SimFinished = "finished"
	SimFailed   = "failed"
)

// RaidAttributes are the raid countdown's fixed facts (ActivityKit attributes).
type RaidAttributes struct {
	RaidID        string `json:"raid_id"`
	GuildID       string `json:"guild_id"`
	GuildName     string `json:"guild_name"`
	ZoneName      string `json:"zone_name"`
	CharacterName string `json:"character_name"`
}

// RaidState is the raid countdown's changing state (ActivityKit content-state). Times are Unix
// seconds: ActivityKit decodes content-state with a default JSONDecoder, which reads a Date
// as seconds since 2001, so plain integers keep both ends honest. The figures are the widget
// snapshot's, so the countdown and the widgets agree.
type RaidState struct {
	StartsAt int64   `json:"starts_at"`
	InviteAt *int64  `json:"invite_at"`
	Status   string  `json:"status"`
	MyState  *string `json:"my_state"`
	Role     string  `json:"role"`
	SignedUp int     `json:"signed_up"`
	Size     *int    `json:"size"`
}

// SimAttributes are a server run's fixed facts.
type SimAttributes struct {
	SimID string `json:"sim_id"`
	Title string `json:"title"`
}

// SimState is a server run's changing state. Total is null for a plain run (it shows
// `Running`); EtaSec null when no estimate exists; DPS set once finished.
type SimState struct {
	State  string   `json:"state"`
	Done   int      `json:"done"`
	Total  *int     `json:"total"`
	EtaSec *int     `json:"eta_sec"`
	DPS    *float64 `json:"dps"`
}

// Alert is the visible line a push-to-start shows.
type Alert struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// RaidLive is one raid countdown's attributes and state, from the snapshot's next raid.
func RaidLive(raid guilds.WidgetNextRaid, characterName string) (RaidAttributes, RaidState) {
	attrs := RaidAttributes{
		RaidID: strconv.FormatInt(raid.RaidID, 10), GuildID: strconv.FormatInt(raid.GuildID, 10),
		GuildName: raid.GuildName, ZoneName: raid.ZoneName, CharacterName: characterName,
	}
	state := RaidState{
		StartsAt: raid.StartsAt.Unix(), Status: raid.Status, MyState: raid.My.State, Role: raid.My.Role,
		SignedUp: raid.SignedUp, Size: raid.Size,
	}
	if raid.InviteAt != nil {
		at := raid.InviteAt.Unix()
		state.InviteAt = &at
	}
	return attrs, state
}

// APNsLive is the APNs body of a Live Activity event: "start" (push-to-start, which names the
// attribute type and carries attrs and an alert), "update" or "end" (dismissal set).
func APNsLive(event string, now time.Time, attrsType string, attrs, state any, alert *Alert, stale, dismissal *time.Time) map[string]any {
	aps := map[string]any{"timestamp": now.Unix(), "event": event, "content-state": state}
	if event == "start" {
		aps["attributes-type"], aps["attributes"] = attrsType, attrs
	}
	if alert != nil {
		aps["alert"] = alert
	}
	if stale != nil {
		aps["stale-date"] = stale.Unix()
	}
	if dismissal != nil {
		aps["dismissal-date"] = dismissal.Unix()
	}
	return map[string]any{"aps": aps}
}

// APNsWidgetRefresh is the silent background push that asks the app to fetch the snapshot and
// reload its widgets (apns-push-type background, priority 5).
func APNsWidgetRefresh() map[string]any {
	return map[string]any{"aps": map[string]any{"content-available": 1}, "kind": EventWidgetRefresh}
}

// FCMData is an FCM data message: Android has no Live Activity, so the app draws the ongoing
// notification itself from these strings. A widget refresh collapses per user.
func FCMData(token, kind string, data map[string]string, highPriority bool) map[string]any {
	payload := map[string]string{"kind": kind}
	for k, v := range data {
		payload[k] = v
	}
	priority := "NORMAL"
	if highPriority {
		priority = "HIGH"
	}
	return map[string]any{"message": map[string]any{
		"token": token, "data": payload,
		"android": map[string]any{"priority": priority, "collapse_key": collapseKey(kind, data)},
	}}
}

// collapseKey keeps one pending message per surface: per user for a refresh, per raid or run
// for a live event.
func collapseKey(kind string, data map[string]string) string {
	switch {
	case data["raid_id"] != "":
		return "live-raid-" + data["raid_id"]
	case data["sim_id"] != "":
		return "live-sim-" + data["sim_id"]
	default:
		return kind
	}
}

// RaidData is RaidLive as FCM strings.
func RaidData(attrs RaidAttributes, state RaidState) map[string]string {
	d := map[string]string{
		"raid_id": attrs.RaidID, "guild_id": attrs.GuildID, "guild_name": attrs.GuildName,
		"zone_name": attrs.ZoneName, "character_name": attrs.CharacterName,
		"starts_at": strconv.FormatInt(state.StartsAt, 10), "status": state.Status, "role": state.Role,
		"signed_up": strconv.Itoa(state.SignedUp),
	}
	setInt64(d, "invite_at", state.InviteAt)
	if state.MyState != nil {
		d["my_state"] = *state.MyState
	}
	if state.Size != nil {
		d["size"] = strconv.Itoa(*state.Size)
	}
	return d
}

// SimData is a server run's state as FCM strings.
func SimData(attrs SimAttributes, state SimState) map[string]string {
	d := map[string]string{"sim_id": attrs.SimID, "title": attrs.Title, "state": state.State,
		"done": strconv.Itoa(state.Done)}
	if state.Total != nil {
		d["total"] = strconv.Itoa(*state.Total)
	}
	if state.EtaSec != nil {
		d["eta_sec"] = strconv.Itoa(*state.EtaSec)
	}
	if state.DPS != nil {
		d["dps"] = strconv.FormatFloat(*state.DPS, 'f', 1, 64)
	}
	return d
}

func setInt64(d map[string]string, key string, v *int64) {
	if v != nil {
		d[key] = strconv.FormatInt(*v, 10)
	}
}
