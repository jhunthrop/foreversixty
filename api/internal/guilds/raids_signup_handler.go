// api/internal/guilds/raids_signup_handler.go
//
// PUT /v1/guilds/{id}/raids/{raid}/signups/{region}/{ruleset}/{name}: a member answers for
// one of their own verified characters; an officer of a guild with the plan answers for any
// verified roster character (and may still write to a locked raid). The character key is the
// path's three segments, as on every roster route.
package guilds

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// signupsPerMinute is the per-account budget for sign-up writes.
const signupsPerMinute = 30

const signupFailMessage = "could not save that sign-up just now"

// signupNotYoursMessage is the forbidden body for a member answering for another's character.
const signupNotYoursMessage = "You can only sign up your own characters."

// signupAnswerView is the PUT's answer: the raid as it now stands and the character R3 moved
// to bench, null when none.
type signupAnswerView struct {
	Raid  RaidView `json:"raid"`
	Moved *string  `json:"moved"`
}

// signupRateLimit budgets sign-up writes per account; a request with no account passes
// through to RequireSession's 401.
func signupRateLimit() func(http.Handler) http.Handler {
	return httpx.RateLimitByKey(signupsPerMinute, time.Minute, func(r *http.Request) (string, bool) {
		uid := auth.ActorFrom(r.Context()).UserID
		return strconv.FormatInt(uid, 10), uid != 0
	})
}

func (s *Service) putSignup(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.memberRaid(w, r, "raid_signup")
	if !ok {
		return
	}
	key, ok := characterKeyFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character in this guild", nil)
		return
	}
	var req signupRequest
	if err := decodeJSON(w, r, &req); err != nil {
		return
	}
	req, fields := validateSignup(req)
	if fields != nil {
		writeInvalid(w, r, fields)
		return
	}
	chars, err := s.Store.raidRoster(r.Context(), guildID, key)
	if err != nil {
		s.fail(w, r, "raid_signup", err, signupFailMessage)
		return
	}
	if len(chars) == 0 {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character in this guild", nil)
		return
	}
	char := chars[0]
	asOfficer, ok := s.signupStanding(w, r, guildID, raid, char)
	if !ok {
		return
	}
	role := req.Role
	if role == "" {
		role = char.defaultRole()
	}
	// SignedBy is the character's account, not the caller: R3 groups by it, so an officer's
	// write for a member benches the member's other character, never the officer's. ActedBy
	// records the caller.
	actorID := auth.ActorFrom(r.Context()).UserID
	moved, err := s.Store.UpsertSignup(r.Context(), Signup{
		RaidID: raid.ID, CharacterKey: char.CharacterKey, Name: char.Name, Class: char.Class, Spec: char.Spec,
		Role: role, State: req.State, Note: req.Note, SignedBy: char.userID, ActedBy: actorID, AsOfficer: asOfficer,
	})
	if !s.writeSignupErr(w, r, err) {
		return
	}
	s.logger().Info("guilds", "op", "raid_signup", "guild_id", guildID, "raid_id", raid.ID,
		"user_id", char.userID, "actor", actorID, "character_key", char.CharacterKey, "as_officer", asOfficer)
	s.writeSignupAnswer(w, r, guildID, raid.ID, moved)
}

// signupStanding decides whether the caller may answer for char, and whether as an officer
// (which a locked raid needs). ok false once it has answered.
func (s *Service) signupStanding(w http.ResponseWriter, r *http.Request, guildID int64, raid Raid, char raidRosterChar) (bool, bool) {
	userID := auth.ActorFrom(r.Context()).UserID
	if char.userID != userID {
		// Someone else's character: the full officer gate (rank, plan, freeze).
		if _, ok := s.officerRaidGate(w, r, "raid_signup", signupNotYoursMessage); !ok {
			return false, false
		}
		if !char.verified {
			httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "that character is not verified in this guild", nil)
			return false, false
		}
		return true, true
	}
	if !char.verified {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "verify that character before signing it up", nil)
		return false, false
	}
	if raid.Status != RaidLocked && raid.Status != RaidDone {
		return false, true
	}
	officer, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, "raid_signup", err, signupFailMessage)
		return false, false
	}
	if !officer || !s.officerTools(r.Context(), guildID, userID) {
		return false, true
	}
	// Writing to a locked raid is an officer power, frozen like the others.
	if !s.officerNotFrozen(w, r, guildID, "raid_signup") {
		return false, false
	}
	return true, true
}

// writeSignupErr answers a failed UpsertSignup; true when err is nil.
func (s *Service) writeSignupErr(w http.ResponseWriter, r *http.Request, err error) bool {
	switch {
	case err == nil:
		return true
	case errors.Is(err, ErrRaidLocked):
		httpx.WriteError(w, r, http.StatusConflict, "raid_locked", "sign-ups for this raid are locked", nil)
	case errors.Is(err, ErrRaidCancelled):
		httpx.WriteError(w, r, http.StatusConflict, "raid_cancelled", "this raid was cancelled", nil)
	case errors.Is(err, ErrNotFound):
		writeRaidNotFound(w, r)
	default:
		s.fail(w, r, "raid_signup", err, signupFailMessage)
	}
	return false
}

// writeSignupAnswer re-reads the raid (its status may have moved meanwhile) and answers.
func (s *Service) writeSignupAnswer(w http.ResponseWriter, r *http.Request, guildID, raidID int64, moved *string) {
	raid, found, err := s.Store.GetRaid(r.Context(), guildID, raidID)
	if err == nil && !found {
		err = ErrNotFound
	}
	var view RaidView
	if err == nil {
		view, err = s.raidViewOf(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID, raid)
	}
	if err != nil {
		s.fail(w, r, "raid_signup", err, signupFailMessage)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, signupAnswerView{Raid: view, Moved: moved})
}
