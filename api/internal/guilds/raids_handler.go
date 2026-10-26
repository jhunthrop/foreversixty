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
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

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

// officerPlanForbidden is the forbidden body for a non-officer on the officer raid routes.
const officerPlanForbidden = "you must be a verified officer of this guild to plan raids"

// officerRaidGate runs the officer write checks in order: a verified officer or leader
// (forbidden, with forbiddenMessage), the guild plan (guild_plan_required), and the claim
// freeze (claim_contested).
func (s *Service) officerRaidGate(w http.ResponseWriter, r *http.Request, op, forbiddenMessage string) (int64, bool) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return 0, false
	}
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, op, err, raidFailMessage)
		return 0, false
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", forbiddenMessage, nil)
		return 0, false
	}
	if !s.requireGuildPlan(w, r, guildID, auth.ActorFrom(r.Context()).UserID, notify.RaidPlanRequiredMessage, op, raidFailMessage) {
		return 0, false
	}
	if !s.officerNotFrozen(w, r, guildID, op) {
		return 0, false
	}
	return guildID, true
}

// officerNotFrozen answers 409 claim_contested and returns false while the guild's claim is
// contested (freezeCheck); true when officer powers may run.
func (s *Service) officerNotFrozen(w http.ResponseWriter, r *http.Request, guildID int64, op string) bool {
	frozen, err := s.freezeCheck(r.Context(), guildID, auth.ActorFrom(r.Context()).IsModerator())
	if err != nil {
		s.fail(w, r, op, err, raidFailMessage)
		return false
	}
	if frozen {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			notify.ClaimContestedMessage, nil)
		return false
	}
	return true
}

// writeRaidWriteErr answers a failed raid write; true when err is nil.
func (s *Service) writeRaidWriteErr(w http.ResponseWriter, r *http.Request, op string, err error) bool {
	switch {
	case err == nil:
		return true
	case errors.Is(err, ErrNotFound):
		writeRaidNotFound(w, r)
	case errors.Is(err, ErrRaidStatus):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "that raid's status does not allow this", nil)
	default:
		s.fail(w, r, op, err, raidFailMessage)
	}
	return false
}

// officerRaid is officerRaidGate plus the {raid} the route names, 404 when the guild has none.
func (s *Service) officerRaid(w http.ResponseWriter, r *http.Request, op string) (int64, Raid, bool) {
	guildID, ok := s.officerRaidGate(w, r, op, officerPlanForbidden)
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
	s.writeRaidView(w, r, "raid", guildID, raid, nil)
}

// writeRaidView answers 200 with raid's view; post, when set, runs on the view first (the
// Discord posts read its counts at posting).
func (s *Service) writeRaidView(w http.ResponseWriter, r *http.Request, op string, guildID int64, raid Raid, post func(RaidView)) {
	view, err := s.raidViewOf(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID, raid)
	if err != nil {
		s.fail(w, r, op, err, "could not load that raid just now")
		return
	}
	if post != nil {
		post(view)
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}

func (s *Service) createRaid(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.officerRaidGate(w, r, "raid_create", officerPlanForbidden)
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
	raids, err := s.Store.CreateRaids(r.Context(), guildID, plan.timezone, inputs)
	if errors.Is(err, ErrRaidCap) {
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "raid_cap", "A guild plans at most 12 raids ahead.", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "raid_create", err, raidFailMessage)
		return
	}
	views, err := s.raidViews(r.Context(), guildID, actor.UserID, raids)
	if err != nil {
		s.fail(w, r, "raid_create", err, raidFailMessage)
		return
	}
	s.logger().Info("guilds", "op", "raid_create", "guild_id", guildID, "raids", len(raids), "user_id", actor.UserID)
	s.rememberPostDefault(r.Context(), guildID, req.PostToDiscord)
	if ticked(req.PostToDiscord) {
		s.postCreated(r.Context(), guildID, views)
	}
	httpx.WriteOK(w, r, http.StatusCreated, map[string][]RaidView{"raids": views})
}

func (s *Service) patchRaid(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.officerRaid(w, r, "raid_update")
	if !ok {
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
	// The store refuses a cancelled or done raid in the same write (409 conflict).
	updated, err := s.Store.UpdateRaid(r.Context(), guildID, raid.ID, plan.timezone, plan.input)
	if !s.writeRaidWriteErr(w, r, "raid_update", err) {
		return
	}
	s.rememberPostDefault(r.Context(), guildID, req.PostToDiscord)
	var post func(RaidView)
	// Only a time change posts; a title or notes edit never does.
	if was := raid.StartsAt; ticked(req.PostToDiscord) && !updated.StartsAt.Equal(was) {
		post = func(v RaidView) { s.postChanged(r.Context(), guildID, v, &was) }
	}
	s.writeRaidView(w, r, "raid_update", guildID, updated, post)
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
		var req cancelRequest
		if action == "cancel" {
			if err := decodeOptionalJSON(w, r, &req); err != nil {
				return
			}
		}
		// Guarded by the current status in the write itself, so racing moves cannot overwrite
		// each other (409 conflict when the status no longer allows this one).
		updated, err := s.Store.SetRaidStatus(r.Context(), guildID, raid.ID, move.from, move.to)
		if !s.writeRaidWriteErr(w, r, op, err) {
			return
		}
		s.logger().Info("guilds", "op", op, "guild_id", guildID, "raid_id", raid.ID,
			"user_id", auth.ActorFrom(r.Context()).UserID)
		// The cancel box only pre-ticks from the raid form's default; it never writes it.
		var post func(RaidView)
		if ticked(req.PostToDiscord) {
			post = func(v RaidView) { s.postChanged(r.Context(), guildID, v, nil) }
		}
		s.writeRaidView(w, r, op, guildID, updated, post)
	}
}

// plannedSection is the members-only half of GET …/raids.
func (s *Service) plannedSection(r *http.Request, guildID, viewerID int64) (*PlannedSection, error) {
	raids, err := s.Store.PlannedRaids(r.Context(), guildID, time.Now())
	if err != nil {
		return nil, err
	}
	views, err := s.raidViews(r.Context(), guildID, viewerID, raids)
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
