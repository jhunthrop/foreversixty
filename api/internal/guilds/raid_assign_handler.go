// api/internal/guilds/raid_assign_handler.go
//
// The raid assignment officer routes (spec 2026-11-05-web-layout §7.2): PUT and DELETE
// …/raids/{raid}/assignments/{encounter}/{slot} set and clear one slot's names; POST
// …/raids/{raid}/assignments/post is raid_assign_post.go. Each runs the officer raid gate
// (verified officer, the plan, no freeze) and answers the raid view. Manual only: nothing here
// picks a name for the raid leader.
package guilds

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const assignmentsFieldKeys = "character_keys"

// mountRaidAssignmentRoutes registers the assignment routes.
func mountRaidAssignmentRoutes(mux *http.ServeMux, s *Service, officerLimit func(http.Handler) http.Handler) {
	const slot = "/v1/guilds/{id}/raids/{raid}/assignments/{encounter}/{slot}"
	mux.Handle("PUT "+slot, officerLimit(raidRouteSession(s.putAssignment)))
	mux.Handle("DELETE "+slot, officerLimit(raidRouteSession(s.deleteAssignment)))
	mux.Handle("POST /v1/guilds/{id}/raids/{raid}/assignments/post", officerLimit(raidRouteSession(s.postAssignments)))
}

// assignmentRequest is PUT …/assignments/{encounter}/{slot}'s body: the slot's names in order.
type assignmentRequest struct {
	CharacterKeys *[]string `json:"character_keys"`
}

// openAssignmentRaid is officerRaid for the assignment routes: the raid must be open (planned
// or locked).
func (s *Service) openAssignmentRaid(w http.ResponseWriter, r *http.Request, op string) (int64, Raid, bool) {
	guildID, raid, ok := s.officerRaid(w, r, op)
	if !ok {
		return 0, Raid{}, false
	}
	if raid.Status != RaidPlanned && raid.Status != RaidLocked {
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "That raid is over, so its assignments are set.", nil)
		return 0, Raid{}, false
	}
	return guildID, raid, true
}

// routeSlot is the {encounter}/{slot} of the raid's zone; false once it has answered 404.
func routeSlot(w http.ResponseWriter, r *http.Request, raid Raid) (slotRef, slotTemplate, bool) {
	id, err := strconv.ParseInt(r.PathValue("encounter"), 10, 64)
	key := r.PathValue("slot")
	if err == nil && len(key) <= maxSlotKeyLen {
		if _, ok := encounterOf(raid.Zone, id); ok {
			if t, ok := slotOf(id, key); ok {
				return slotRef{EncounterID: id, SlotKey: key}, t, true
			}
		}
	}
	httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such assignment slot", nil)
	return slotRef{}, slotTemplate{}, false
}

func (s *Service) putAssignment(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.openAssignmentRaid(w, r, "raid_assignment")
	if !ok {
		return
	}
	ref, t, ok := routeSlot(w, r, raid)
	if !ok {
		return
	}
	var req assignmentRequest
	if err := decodeJSON(w, r, &req); err != nil {
		return
	}
	if req.CharacterKeys == nil {
		writeInvalid(w, r, map[string]string{assignmentsFieldKeys: "List the characters for this slot."})
		return
	}
	signups, err := s.Store.RaidSignups(r.Context(), raid.ID)
	if err != nil {
		s.fail(w, r, "raid_assignment", err, raidFailMessage)
		return
	}
	if msg := validateSlotKeys(t, *req.CharacterKeys, signups); msg != "" {
		writeInvalid(w, r, map[string]string{assignmentsFieldKeys: msg})
		return
	}
	if err := s.Store.SetSlot(r.Context(), raid.ID, ref, *req.CharacterKeys, auth.ActorFrom(r.Context()).UserID); err != nil {
		s.fail(w, r, "raid_assignment", err, raidFailMessage)
		return
	}
	s.writeRaidView(w, r, "raid_assignment", guildID, raid, nil)
}

func (s *Service) deleteAssignment(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.openAssignmentRaid(w, r, "raid_assignment_clear")
	if !ok {
		return
	}
	ref, _, ok := routeSlot(w, r, raid)
	if !ok {
		return
	}
	if err := s.Store.ClearSlot(r.Context(), raid.ID, ref); err != nil {
		s.fail(w, r, "raid_assignment_clear", err, raidFailMessage)
		return
	}
	s.writeRaidView(w, r, "raid_assignment_clear", guildID, raid, nil)
}

// validateSlotKeys is why keys cannot fill the slot, empty when they can: at most the slot's
// count, each once, each signed In, Late or Bench for the raid. The role filter never refuses.
func validateSlotKeys(t slotTemplate, keys []string, signups []Signup) string {
	if len(keys) > t.Count {
		return fmt.Sprintf("%s takes at most %d.", t.Label, t.Count)
	}
	state := make(map[string]string, len(signups))
	for _, sg := range signups {
		state[sg.CharacterKey] = sg.State
	}
	seen := make(map[string]bool, len(keys))
	for _, key := range keys {
		if seen[key] {
			return "Pick each character once."
		}
		seen[key] = true
		if !assignable(state[key]) {
			return "Only characters signed In, Late or Bench can take a slot."
		}
	}
	return ""
}
