// api/internal/guilds/raids_attendance_override_handler.go
//
// PUT and DELETE …/raids/{raid}/attendance/{region}/{ruleset}/{name}/override: an officer
// marks a raider present, absent, bench or excused for the night, or clears the fix to restore
// what the log said. Both answer the raid view, as the excused route does.
package guilds

import (
	"errors"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

type overrideRequest struct {
	Status string `json:"status"`
}

func mountAttendanceOverrideRoutes(mux *http.ServeMux, s *Service, officerLimit func(http.Handler) http.Handler) {
	const route = "/v1/guilds/{id}/raids/{raid}/attendance/{region}/{ruleset}/{name}/override"
	mux.Handle("PUT "+route, officerLimit(raidRouteSession(s.putAttendanceOverride)))
	mux.Handle("DELETE "+route, officerLimit(raidRouteSession(s.deleteAttendanceOverride)))
}

func (s *Service) putAttendanceOverride(w http.ResponseWriter, r *http.Request) {
	const op = "raid_attendance_override"
	guildID, raid, key, ok := s.overrideTarget(w, r, op)
	if !ok {
		return
	}
	var req overrideRequest
	if err := decodeJSON(w, r, &req); err != nil {
		return
	}
	if !validOverride(req.Status) {
		writeInvalid(w, r, map[string]string{"status": "say present, absent, bench or excused"})
		return
	}
	err := s.Store.SetAttendanceOverride(r.Context(), guildID, raid.ID, key, req.Status, auth.ActorFrom(r.Context()).UserID)
	if !s.answeredOverrideWrite(w, r, op, err) {
		return
	}
	s.markRaid(r.Context(), raid.ID)
	s.logger().Info("guilds", "op", op, "guild_id", guildID, "raid_id", raid.ID, "status", req.Status)
	s.writeRaidView(w, r, op, guildID, raid, nil)
}

func (s *Service) deleteAttendanceOverride(w http.ResponseWriter, r *http.Request) {
	const op = "raid_attendance_override_clear"
	guildID, raid, key, ok := s.overrideTarget(w, r, op)
	if !ok {
		return
	}
	err := s.Store.ClearAttendanceOverride(r.Context(), guildID, raid.ID, key)
	if !s.answeredOverrideWrite(w, r, op, err) {
		return
	}
	s.markRaid(r.Context(), raid.ID)
	s.writeRaidView(w, r, op, guildID, raid, nil)
}

// overrideTarget is the officer gate and the route's character; ok false once answered.
func (s *Service) overrideTarget(w http.ResponseWriter, r *http.Request, op string) (int64, Raid, string, bool) {
	guildID, raid, ok := s.matchOfficerRaid(w, r, op)
	if !ok {
		return 0, Raid{}, "", false
	}
	key, ok := characterKeyFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character in this guild", nil)
		return 0, Raid{}, "", false
	}
	return guildID, raid, key, true
}

// answeredOverrideWrite answers a failed fix; true when err is nil.
func (s *Service) answeredOverrideWrite(w http.ResponseWriter, r *http.Request, op string, err error) bool {
	switch {
	case err == nil:
		return true
	case errors.Is(err, ErrNoLog):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "match a log to this raid before fixing attendance", nil)
	case errors.Is(err, ErrNotFound):
		writeInvalid(w, r, map[string]string{"character": "that character is not a verified raider of this guild"})
	default:
		s.fail(w, r, op, err, raidFailMessage)
	}
	return false
}
