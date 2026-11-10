// api/internal/guilds/setup_handler.go
//
// The setup checklist's routes: GET /v1/guilds/{id}/setup reads it; PUT …/setup/{step}
// {"state":"done"|"skipped"} marks a step and DELETE …/setup/{step} takes the mark back, each
// answering the checklist. Verified officers and the leader only (403 otherwise); the writes
// wait out a contested claim. No plan: the checklist is part of claiming a guild.
package guilds

import (
	"errors"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const (
	setupFailMessage = "could not load the setup just now"
	setupStepParam   = "step"
)

func mountSetupRoutes(mux *http.ServeMux, s *Service, officerLimit func(http.Handler) http.Handler) {
	mux.HandleFunc("GET /v1/guilds/{id}/setup", auth.RequireSession(s.getSetup))
	mux.Handle("PUT /v1/guilds/{id}/setup/{step}", officerLimit(auth.RequireSession(s.putSetupMark)))
	mux.Handle("DELETE /v1/guilds/{id}/setup/{step}", officerLimit(auth.RequireSession(s.deleteSetupMark)))
}

func (s *Service) getSetup(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.officerGuild(w, r, "setup", setupFailMessage)
	if !ok {
		return
	}
	s.writeSetup(w, r, guildID)
}

type setupMarkRequest struct {
	State string `json:"state"`
}

func (s *Service) putSetupMark(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.officerWriteGuild(w, r, "setup_mark", setupFailMessage)
	if !ok {
		return
	}
	var req setupMarkRequest
	if err := decodeJSON(w, r, &req); err != nil {
		return
	}
	err := s.Store.MarkSetup(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID, r.PathValue(setupStepParam), req.State)
	if !s.answeredSetupWrite(w, r, err) {
		return
	}
	s.writeSetup(w, r, guildID)
}

func (s *Service) deleteSetupMark(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.officerWriteGuild(w, r, "setup_mark", setupFailMessage)
	if !ok {
		return
	}
	if !s.answeredSetupWrite(w, r, s.Store.ClearSetupMark(r.Context(), guildID, r.PathValue(setupStepParam))) {
		return
	}
	s.writeSetup(w, r, guildID)
}

func (s *Service) answeredSetupWrite(w http.ResponseWriter, r *http.Request, err error) bool {
	switch {
	case err == nil:
		return true
	case errors.Is(err, ErrSetupStep):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that step does not take that mark", nil)
	default:
		s.fail(w, r, "setup_mark", err, setupFailMessage)
	}
	return false
}

func (s *Service) writeSetup(w http.ResponseWriter, r *http.Request, guildID int64) {
	view, err := s.Setup(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "setup", err, setupFailMessage)
		return
	}
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}
