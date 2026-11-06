// api/internal/guilds/loot_council_handler.go
//
// The loot council routes. Nights, candidates and votes are officer-only; the officer writes
// (votes, an MS/OS flag on someone else's behalf) take the Raid guild plan and stop while the
// claim is contested, like the award. A member reads and sets only their own MS/OS flags.
package guilds

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/refusal"
)

const (
	councilReadFail    = "could not load the loot council just now"
	councilVoteFail    = "could not record that vote just now"
	councilFlagFail    = "could not record that just now"
	councilOfficerOnly = "you must be a verified officer of this guild to see the loot council"
)

func mountLootCouncilRoutes(mux *http.ServeMux, s *Service, officerLimit func(http.Handler) http.Handler) {
	mux.HandleFunc("GET /v1/guilds/{id}/loot/nights", auth.RequireSession(s.lootNights))
	mux.HandleFunc("GET /v1/guilds/{id}/loot/drops/{item_id}/candidates", auth.RequireSession(s.lootDropCandidates))
	mux.Handle("PUT /v1/guilds/{id}/loot/votes", officerLimit(auth.RequireSession(s.putLootVote)))
	mux.HandleFunc("GET /v1/guilds/{id}/loot/interest", auth.RequireSession(s.getLootInterest))
	mux.Handle("PUT /v1/guilds/{id}/loot/interest", officerLimit(auth.RequireSession(s.putLootInterest)))
}

// councilGuild reads the guild id and checks the caller is a verified officer, writing the
// refusal and returning false otherwise.
func (s *Service) councilGuild(w http.ResponseWriter, r *http.Request, op, failMessage, forbidden string) (int64, bool) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return 0, false
	}
	return guildID, s.councilOfficer(w, r, guildID, op, failMessage, forbidden)
}

// councilOfficer checks the caller is a verified officer of guildID, writing the refusal.
func (s *Service) councilOfficer(w http.ResponseWriter, r *http.Request, guildID int64, op, failMessage, forbidden string) bool {
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, op, err, failMessage)
		return false
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", forbidden, nil)
		return false
	}
	return true
}

// councilOfficerGate is councilOfficer then councilOfficerWrite.
func (s *Service) councilOfficerGate(w http.ResponseWriter, r *http.Request, guildID int64, op, failMessage, forbidden string) bool {
	return s.councilOfficer(w, r, guildID, op, failMessage, forbidden) && s.councilOfficerWrite(w, r, guildID, op, failMessage)
}

// councilOfficerWrite is the gate every council officer write shares after the officer check:
// the Raid guild plan, then the claim freeze.
func (s *Service) councilOfficerWrite(w http.ResponseWriter, r *http.Request, guildID int64, op, failMessage string) bool {
	actor := auth.ActorFrom(r.Context())
	if !s.requireGuildPlan(w, r, guildID, actor.UserID, entitlements.FeatureRaidGuildTools, refusal.AwardPlanRequired, op, failMessage) {
		return false
	}
	return s.officerNotFrozen(w, r, guildID, op, failMessage)
}

// lootNights is GET /v1/guilds/{id}/loot/nights?limit=.
func (s *Service) lootNights(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.councilGuild(w, r, "loot_nights", councilReadFail, councilOfficerOnly)
	if !ok {
		return
	}
	limit := councilNightsDefault
	if raw := r.URL.Query().Get("limit"); raw != "" {
		n, err := strconv.Atoi(raw)
		if err != nil || n < 1 || n > councilNightsMax {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "limit must be 1 to 30", nil)
			return
		}
		limit = n
	}
	view, err := s.Store.CouncilNights(r.Context(), guildID, limit, time.Now())
	if err != nil {
		s.fail(w, r, "loot_nights", err, councilReadFail)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// lootDropCandidates is GET /v1/guilds/{id}/loot/drops/{item_id}/candidates?night=.
func (s *Service) lootDropCandidates(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.councilGuild(w, r, "loot_candidates", councilReadFail, councilOfficerOnly)
	if !ok {
		return
	}
	itemID, err := strconv.Atoi(r.PathValue("item_id"))
	if err != nil || !validOnyxiaItem(itemID) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such drop", nil)
		return
	}
	night, ok := s.councilNightFrom(w, r, guildID, r.URL.Query().Get("night"), "loot_candidates", councilReadFail)
	if !ok {
		return
	}
	view, err := s.Store.CouncilDrop(r.Context(), guildID, night, itemID, auth.ActorFrom(r.Context()).UserID, time.Now())
	if err != nil {
		s.fail(w, r, "loot_candidates", err, councilReadFail)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// lootDropInput names a drop of a night and a candidate.
type lootDropInput struct {
	Night        string `json:"night"`
	ItemID       int    `json:"item_id"`
	CharacterKey string `json:"character_key"`
}

type lootVoteInput struct {
	lootDropInput
	// Vote is "yes" or "no"; null removes the officer's vote.
	Vote *string `json:"vote"`
}

// putLootVote is PUT /v1/guilds/{id}/loot/votes.
func (s *Service) putLootVote(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.councilGuild(w, r, "loot_vote", councilVoteFail, "you must be a verified officer of this guild to vote on loot")
	if !ok || !s.councilOfficerWrite(w, r, guildID, "loot_vote", councilVoteFail) {
		return
	}
	var in lootVoteInput
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	vote := ""
	if in.Vote != nil {
		if *in.Vote != voteYes && *in.Vote != voteNo {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", `vote must be "yes", "no" or null`, nil)
			return
		}
		vote = *in.Vote
	}
	night, ok := s.councilNightFrom(w, r, guildID, in.Night, "loot_vote", councilVoteFail)
	if !ok {
		return
	}
	actor := auth.ActorFrom(r.Context())
	now := time.Now()
	err := s.Store.SetLootVote(r.Context(), guildID, night, in.ItemID, in.CharacterKey, actor.UserID, vote, now)
	if !s.councilWriteOK(w, r, err, "loot_vote", councilVoteFail) {
		return
	}
	tally, err := s.Store.candidateVotes(r.Context(), guildID, night, in.ItemID, in.CharacterKey, actor.UserID, now)
	if err != nil {
		s.fail(w, r, "loot_vote", err, councilVoteFail)
		return
	}
	s.logger().Info("guilds", "op", "loot_vote", "guild_id", guildID, "item_id", in.ItemID,
		"character_key", in.CharacterKey, "user_id", actor.UserID, "vote", vote)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"night": night.key(), "item_id": in.ItemID,
		"character_key": in.CharacterKey, "votes": tally})
}

// councilWriteOK maps a council store write's error to its refusal; true when err is nil.
func (s *Service) councilWriteOK(w http.ResponseWriter, r *http.Request, err error, op, failMessage string) bool {
	switch {
	case err == nil:
		return true
	case errors.Is(err, ErrInvalidAward):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a valid item/character for this guild", nil)
	case errors.Is(err, ErrVotingClosed):
		httpx.WriteError(w, r, http.StatusConflict, "voting_closed", "that drop is already awarded", nil)
	default:
		s.fail(w, r, op, err, failMessage)
	}
	return false
}
