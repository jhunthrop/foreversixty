// api/internal/guilds/raids_signup_handler.go
//
// PUT /v1/guilds/{id}/raids/{raid}/signups/{region}/{ruleset}/{name}: a member answers for
// one of their own verified characters; an officer of a guild with the plan answers for any
// verified roster character (and may still write to a locked raid). The character key is the
// path's three segments, as on every roster route. The rules live in SignUp
// (raids_signup_service.go); this file maps its answers to HTTP.
package guilds

import (
	"errors"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/refusal"
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

// signupRateLimit budgets sign-up writes per account against the Service's shared limiter;
// a request with no account passes through to RequireSession's 401.
func (s *Service) signupRateLimit() func(http.Handler) http.Handler {
	return httpx.RateLimitWith(s.signupLimiter(), func(r *http.Request) (string, bool) {
		uid := auth.ActorFrom(r.Context()).UserID
		return signupLimitKey(uid), uid != 0
	})
}

// officerRaidWritesPerMinute budgets the officer raid writes (create, edit, lock, unlock,
// cancel) per account; the five routes share one budget.
const officerRaidWritesPerMinute = 20

func officerRaidRateLimit() func(http.Handler) http.Handler {
	return httpx.RateLimitByKey(officerRaidWritesPerMinute, time.Minute, auth.AccountRateLimitKey)
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
	actorID := auth.ActorFrom(r.Context()).UserID
	if req.Spec != "" {
		// The spec's role is the sign-up's: saved first, the sign-up then takes its default.
		req.Role = ""
		if _, err := s.SaveRaidSpec(r.Context(), RaidSpecCommand{
			GuildID: guildID, ActorID: actorID, RaidID: raid.ID, CharacterKey: key, Spec: req.Spec,
		}); err != nil {
			s.writeSignupErr(w, r, err)
			return
		}
	}
	res, err := s.signUp(r.Context(), SignupCommand{
		GuildID: guildID, RaidID: raid.ID, ActorID: actorID,
		CharacterKey: key, State: req.State, Role: req.Role, Note: &req.Note,
	}, false)
	if err != nil {
		s.writeSignupErr(w, r, err)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, signupAnswerView{Raid: res.Raid, Moved: res.Moved})
}

// writeSignupErr answers a failed sign-up in the shapes the PUT route has always used.
func (s *Service) writeSignupErr(w http.ResponseWriter, r *http.Request, err error) {
	switch {
	case errors.Is(err, ErrNotMember), errors.Is(err, ErrNotFound):
		writeRaidNotFound(w, r)
	case errors.Is(err, ErrCharacterNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character in this guild", nil)
	case errors.Is(err, ErrNotYourCharacter):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", signupNotYoursMessage, nil)
	case errors.Is(err, errOfficerTargetUnverified):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "that character is not verified in this guild", nil)
	case errors.Is(err, ErrCharacterNotVerified):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "verify that character before signing it up", nil)
	case errors.Is(err, ErrPlanRequired):
		httpx.WriteError(w, r, http.StatusForbidden, "guild_plan_required", refusal.RaidPlanRequired, nil)
	case errors.Is(err, ErrClaimContested):
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested", refusal.ClaimContested, nil)
	case errors.Is(err, ErrSignupRateLimited):
		httpx.WriteError(w, r, http.StatusTooManyRequests, "rate_limited", "too many requests", nil)
	case errors.Is(err, ErrRaidLocked):
		httpx.WriteError(w, r, http.StatusConflict, "raid_locked", "sign-ups for this raid are locked", nil)
	case errors.Is(err, ErrSpecNotOfClass):
		writeInvalid(w, r, map[string]string{"spec": "Choose a spec of that character's class."})
	case errors.Is(err, ErrRaidCancelled):
		httpx.WriteError(w, r, http.StatusConflict, "raid_cancelled", "this raid was cancelled", nil)
	default:
		s.fail(w, r, "raid_signup", err, signupFailMessage)
	}
}
