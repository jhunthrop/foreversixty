// api/internal/guilds/raids_overview_handler.go
//
// The Overview's officer raid routes (spec 2026-11-05-web-layout §7.1): PUT and DELETE
// …/raids/{raid}/comp-target set and clear the raid's own comp target; POST …/raids/{raid}/ping
// is raids_ping.go. Each runs the officer raid gate (verified officer, the plan, no freeze).
package guilds

import (
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// mountRaidOverviewRoutes registers the Overview's officer raid routes.
func mountRaidOverviewRoutes(mux *http.ServeMux, s *Service, officerLimit func(http.Handler) http.Handler) {
	mux.Handle("PUT /v1/guilds/{id}/raids/{raid}/comp-target", officerLimit(raidRouteSession(s.putCompTarget)))
	mux.Handle("DELETE /v1/guilds/{id}/raids/{raid}/comp-target", officerLimit(raidRouteSession(s.deleteCompTarget)))
	mux.Handle("POST /v1/guilds/{id}/raids/{raid}/ping", officerLimit(raidRouteSession(s.pingRaid)))
}

// compTargetRequest is PUT …/comp-target's body: all four counts, each required.
type compTargetRequest struct {
	Tanks   *int `json:"tanks"`
	Healers *int `json:"healers"`
	Melee   *int `json:"melee"`
	Ranged  *int `json:"ranged"`
}

// counts is the request's counts, or the fields it left out.
func (q compTargetRequest) counts() (RoleCounts, map[string]string) {
	missing := map[string]string{}
	for name, p := range map[string]*int{"tanks": q.Tanks, "healers": q.Healers, "melee": q.Melee, "ranged": q.Ranged} {
		if p == nil {
			missing[name] = "Give a number, 0 or more."
		}
	}
	if len(missing) > 0 {
		return RoleCounts{}, missing
	}
	return RoleCounts{Tanks: *q.Tanks, Healers: *q.Healers, Melee: *q.Melee, Ranged: *q.Ranged}, nil
}

// compTargetRaid is officerRaid for the comp-target routes: the raid must be open (planned or
// locked) and its zone must have a default to edit.
func (s *Service) compTargetRaid(w http.ResponseWriter, r *http.Request, op string) (int64, Raid, Zone, bool) {
	guildID, raid, ok := s.officerRaid(w, r, op)
	if !ok {
		return 0, Raid{}, Zone{}, false
	}
	if raid.Status != RaidPlanned && raid.Status != RaidLocked {
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "That raid is over, so its comp is set.", nil)
		return 0, Raid{}, Zone{}, false
	}
	zone := zoneOf(raid.Zone)
	if _, has := zoneCompDefault(zone); !has {
		httpx.WriteError(w, r, http.StatusConflict, "no_comp_target", "This run has no comp target to set.", nil)
		return 0, Raid{}, Zone{}, false
	}
	return guildID, raid, zone, true
}

func (s *Service) putCompTarget(w http.ResponseWriter, r *http.Request) {
	guildID, raid, zone, ok := s.compTargetRaid(w, r, "raid_comp_target")
	if !ok {
		return
	}
	var req compTargetRequest
	if err := decodeJSON(w, r, &req); err != nil {
		return
	}
	counts, fields := req.counts()
	if fields == nil {
		fields = validateCompTarget(zone, counts)
	}
	if fields != nil {
		writeInvalid(w, r, fields)
		return
	}
	if err := s.Store.SetCompTarget(r.Context(), raid.ID, auth.ActorFrom(r.Context()).UserID, counts); err != nil {
		s.fail(w, r, "raid_comp_target", err, raidFailMessage)
		return
	}
	s.writeRaidView(w, r, "raid_comp_target", guildID, raid, nil)
}

func (s *Service) deleteCompTarget(w http.ResponseWriter, r *http.Request) {
	guildID, raid, _, ok := s.compTargetRaid(w, r, "raid_comp_target_clear")
	if !ok {
		return
	}
	if err := s.Store.ClearCompTarget(r.Context(), raid.ID); err != nil {
		s.fail(w, r, "raid_comp_target_clear", err, raidFailMessage)
		return
	}
	s.writeRaidView(w, r, "raid_comp_target_clear", guildID, raid, nil)
}
