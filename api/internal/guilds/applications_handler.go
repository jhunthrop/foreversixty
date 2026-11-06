// api/internal/guilds/applications_handler.go
//
// The officer side of applications: the list, votes, and accept or decline.
package guilds

import (
	"errors"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// officerApplications is GET /v1/guilds/{id}/applications: pending first, then recent decisions.
func (s *Service) officerApplications(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.officerGuild(w, r, "applications", officerAppsFail)
	if !ok {
		return
	}
	actor := auth.ActorFrom(r.Context())
	rows, err := s.Store.OfficerApplications(r.Context(), guildID, actor.UserID)
	if err != nil {
		s.fail(w, r, "applications", err, officerAppsFail)
		return
	}
	pending := 0
	for _, a := range rows {
		if a.Status == ApplicationPending {
			pending++
		}
	}
	roleID, err := s.Store.RecruitRoleID(r.Context(), guildID)
	if err != nil {
		s.fail(w, r, "applications", err, officerAppsFail)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"rows": rows, "pending": pending, "discord_role_id": roleID})
}

type voteInput struct {
	Vote string `json:"vote"`
}

// putVote is PUT /v1/guilds/{id}/applications/{app}/vote: yes or no, changeable.
func (s *Service) putVote(w http.ResponseWriter, r *http.Request) {
	var in voteInput
	guildID, ok := s.officerWriteGuild(w, r, "vote", officerAppsFail)
	if !ok {
		return
	}
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	if in.Vote != VoteYes && in.Vote != VoteNo {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "vote must be yes or no", map[string]string{"vote": "yes or no"})
		return
	}
	s.writeVote(w, r, guildID, in.Vote)
}

// deleteVote is DELETE /v1/guilds/{id}/applications/{app}/vote: takes the officer's vote back.
func (s *Service) deleteVote(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.officerWriteGuild(w, r, "vote", officerAppsFail)
	if !ok {
		return
	}
	s.writeVote(w, r, guildID, "")
}

func (s *Service) writeVote(w http.ResponseWriter, r *http.Request, guildID int64, vote string) {
	id, ok := applicationIDFrom(w, r)
	if !ok {
		return
	}
	votes, err := s.Store.Vote(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID, id, vote)
	if s.writeApplicationError(w, r, "vote", err) {
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, votes)
}

// writeApplicationError writes the refusal or failure for err and reports whether it wrote one.
func (s *Service) writeApplicationError(w http.ResponseWriter, r *http.Request, op string, err error) bool {
	switch {
	case err == nil:
		return false
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such application", nil)
	case errors.Is(err, ErrApplicationClosed):
		httpx.WriteError(w, r, http.StatusConflict, "application_closed", "that application is already decided or withdrawn", nil)
	default:
		s.fail(w, r, op, err, officerAppsFail)
	}
	return true
}

// decideRoute is POST .../accept (accept true) or .../decline.
func (s *Service) decideRoute(accept bool) http.HandlerFunc {
	op := "application_decline"
	if accept {
		op = "application_accept"
	}
	return func(w http.ResponseWriter, r *http.Request) {
		guildID, ok := s.officerWriteGuild(w, r, op, officerAppsFail)
		if !ok {
			return
		}
		id, ok := applicationIDFrom(w, r)
		if !ok {
			return
		}
		actor := auth.ActorFrom(r.Context())
		d, err := s.Decide(r.Context(), guildID, actor.UserID, id, accept)
		if s.writeApplicationError(w, r, op, err) {
			return
		}
		s.logger().Info("guilds", "op", op, "guild_id", guildID, "user_id", actor.UserID, "application_id", id,
			"role_granted", d.DiscordRole.Done, "dm_sent", d.DiscordDM.Done)
		httpx.WriteOK(w, r, http.StatusOK, d)
	}
}
