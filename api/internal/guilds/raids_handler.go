// api/internal/guilds/raids_handler.go
//
// The raid calendar's read and officer routes (design/specs/2026-10-26-guild-raids.md §9):
// GET …/raids/{raid}, POST …/raids, PATCH …/raids/{raid} and POST …/raids/{raid}/lock|unlock|
// cancel. Sign-ups are in raids_signup_handler.go. R1: a raid is told to verified members
// only; anyone else gets 404, never 403.
package guilds

import (
	"errors"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// raidPlanRequiredMessage is the guild_plan_required body for the officer raid writes.
const raidPlanRequiredMessage = "Planning and running raids are part of the guild plan."

const raidFailMessage = "could not save that raid just now"

// raidStatusMoves is which statuses each officer action moves a raid from, and to.
var raidStatusMoves = map[string]struct {
	from []string
	to   string
}{
	"lock":   {from: []string{RaidPlanned}, to: RaidLocked},
	"unlock": {from: []string{RaidLocked}, to: RaidPlanned},
	"cancel": {from: []string{RaidPlanned, RaidLocked}, to: RaidCancelled},
}

func writeRaidNotFound(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such raid", nil)
}

func writeInvalid(w http.ResponseWriter, r *http.Request, fields map[string]string) {
	httpx.WriteError(w, r, http.StatusUnprocessableEntity, "invalid", "check the highlighted fields", fields)
}

// raidIDFrom parses the {raid} path value.
func raidIDFrom(r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("raid"), 10, 64)
	return id, err == nil && id > 0
}

// memberRaid resolves the guild and raid of a member route: ok false once it has answered
// (404 for a non-member, a missing raid or a bad id; 500 on a failed read).
func (s *Service) memberRaid(w http.ResponseWriter, r *http.Request, op string) (int64, Raid, bool) {
	guildID, gok := guildIDFrom(r)
	raidID, rok := raidIDFrom(r)
	if !gok || !rok {
		writeRaidNotFound(w, r)
		return 0, Raid{}, false
	}
	_, verified, err := s.Accounts.GuildRank(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, op, err, "could not load that raid just now")
		return 0, Raid{}, false
	}
	if !verified {
		writeRaidNotFound(w, r)
		return 0, Raid{}, false
	}
	raid, found, err := s.Store.GetRaid(r.Context(), guildID, raidID)
	if err != nil {
		s.fail(w, r, op, err, "could not load that raid just now")
		return 0, Raid{}, false
	}
	if !found {
		writeRaidNotFound(w, r)
		return 0, Raid{}, false
	}
	return guildID, raid, true
}

// officerRaidGate runs the officer write checks in order: a verified officer or leader
// (forbidden), the guild plan (guild_plan_required), and the claim freeze (claim_contested).
func (s *Service) officerRaidGate(w http.ResponseWriter, r *http.Request, op string) (int64, bool) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return 0, false
	}
	actor := auth.ActorFrom(r.Context())
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, op, err, raidFailMessage)
		return 0, false
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to plan raids", nil)
		return 0, false
	}
	if !s.requireGuildPlan(w, r, guildID, actor.UserID, raidPlanRequiredMessage, op, raidFailMessage) {
		return 0, false
	}
	frozen, err := s.freezeCheck(r.Context(), guildID, actor.IsModerator())
	if err != nil {
		s.fail(w, r, op, err, raidFailMessage)
		return 0, false
	}
	if frozen {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
		return 0, false
	}
	return guildID, true
}

// officerRaid is officerRaidGate plus the {raid} the route names, 404 when the guild has none.
func (s *Service) officerRaid(w http.ResponseWriter, r *http.Request, op string) (int64, Raid, bool) {
	guildID, ok := s.officerRaidGate(w, r, op)
	if !ok {
		return 0, Raid{}, false
	}
	raidID, ok := raidIDFrom(r)
	if !ok {
		writeRaidNotFound(w, r)
		return 0, Raid{}, false
	}
	raid, found, err := s.Store.GetRaid(r.Context(), guildID, raidID)
	if err != nil {
		s.fail(w, r, op, err, raidFailMessage)
		return 0, Raid{}, false
	}
	if !found {
		writeRaidNotFound(w, r)
		return 0, Raid{}, false
	}
	return guildID, raid, true
}

func (s *Service) getRaid(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.memberRaid(w, r, "raid")
	if !ok {
		return
	}
	s.writeRaidView(w, r, "raid", guildID, raid, http.StatusOK)
}

func (s *Service) writeRaidView(w http.ResponseWriter, r *http.Request, op string, guildID int64, raid Raid, status int) {
	view, err := s.Store.RaidViewOf(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID, raid)
	if err != nil {
		s.fail(w, r, op, err, "could not load that raid just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, status, view)
}

func (s *Service) createRaid(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.officerRaidGate(w, r, "raid_create")
	if !ok {
		return
	}
	var req raidRequest
	if err := decodeJSON(w, r, &req); err != nil {
		return
	}
	plan, fields := validateRaid(req, time.Now(), nil, true)
	if fields != nil {
		writeInvalid(w, r, fields)
		return
	}
	actor := auth.ActorFrom(r.Context())
	inputs := weeklyInputs(plan)
	for i := range inputs {
		inputs[i].GuildID, inputs[i].CreatedBy = guildID, actor.UserID
	}
	raids, err := s.Store.CreateRaids(r.Context(), guildID, inputs)
	if errors.Is(err, ErrRaidCap) {
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "raid_cap", "A guild plans at most 12 raids ahead.", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "raid_create", err, raidFailMessage)
		return
	}
	if !s.saveRaidTimezone(w, r, guildID, plan.timezone, "raid_create") {
		return
	}
	views, err := s.Store.RaidViews(r.Context(), guildID, actor.UserID, raids)
	if err != nil {
		s.fail(w, r, "raid_create", err, raidFailMessage)
		return
	}
	s.logger().Info("guilds", "op", "raid_create", "guild_id", guildID, "raids", len(raids), "user_id", actor.UserID)
	httpx.WriteOK(w, r, http.StatusCreated, map[string][]RaidView{"raids": views})
}

// saveRaidTimezone stores tz as the guild's raid timezone when it differs.
func (s *Service) saveRaidTimezone(w http.ResponseWriter, r *http.Request, guildID int64, tz, op string) bool {
	current, err := s.Store.GuildRaidTimezone(r.Context(), guildID)
	if err == nil && current != tz {
		err = s.Store.SetGuildRaidTimezone(r.Context(), guildID, tz)
	}
	if err != nil {
		s.fail(w, r, op, err, raidFailMessage)
		return false
	}
	return true
}

func (s *Service) patchRaid(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.officerRaid(w, r, "raid_update")
	if !ok {
		return
	}
	if raid.Status != RaidPlanned && raid.Status != RaidLocked {
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "a cancelled or logged raid cannot be edited", nil)
		return
	}
	var req raidRequest
	if err := decodeJSON(w, r, &req); err != nil {
		return
	}
	plan, fields := validateRaid(req, time.Now(), &raid.StartsAt, false)
	if fields != nil {
		writeInvalid(w, r, fields)
		return
	}
	updated, err := s.Store.UpdateRaid(r.Context(), guildID, raid.ID, plan.input)
	if errors.Is(err, ErrNotFound) {
		writeRaidNotFound(w, r)
		return
	}
	if err != nil {
		s.fail(w, r, "raid_update", err, raidFailMessage)
		return
	}
	if !s.saveRaidTimezone(w, r, guildID, plan.timezone, "raid_update") {
		return
	}
	s.writeRaidView(w, r, "raid_update", guildID, updated, http.StatusOK)
}

// raidStatusAction serves POST …/raids/{raid}/{action} for one of raidStatusMoves.
func (s *Service) raidStatusAction(action string) http.HandlerFunc {
	move := raidStatusMoves[action]
	op := "raid_" + action
	return func(w http.ResponseWriter, r *http.Request) {
		guildID, raid, ok := s.officerRaid(w, r, op)
		if !ok {
			return
		}
		if !slices.Contains(move.from, raid.Status) {
			httpx.WriteError(w, r, http.StatusConflict, "conflict", "that raid is "+raid.Status, nil)
			return
		}
		updated, err := s.Store.SetRaidStatus(r.Context(), guildID, raid.ID, move.to)
		if errors.Is(err, ErrNotFound) {
			writeRaidNotFound(w, r)
			return
		}
		if err != nil {
			s.fail(w, r, op, err, raidFailMessage)
			return
		}
		s.logger().Info("guilds", "op", op, "guild_id", guildID, "raid_id", raid.ID,
			"user_id", auth.ActorFrom(r.Context()).UserID)
		s.writeRaidView(w, r, op, guildID, updated, http.StatusOK)
	}
}

// plannedSection is the members-only half of GET …/raids.
func (s *Service) plannedSection(r *http.Request, guildID, viewerID int64) (*PlannedSection, error) {
	raids, err := s.Store.PlannedRaids(r.Context(), guildID, time.Now())
	if err != nil {
		return nil, err
	}
	views, err := s.Store.RaidViews(r.Context(), guildID, viewerID, raids)
	if err != nil {
		return nil, err
	}
	section := &PlannedSection{Planned: views, Zones: zoneViews()}
	tz, err := s.Store.GuildRaidTimezone(r.Context(), guildID)
	if err != nil {
		return nil, err
	}
	if tz != "" {
		section.GuildTimezone = &tz
	}
	return section, nil
}
