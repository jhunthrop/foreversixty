package coach

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// Codes the night, recent and history routes answer with.
const (
	codeNotFound = "not_found"
	codeNoPlan   = "no_plan"
	codeFlagged  = "flagged"
	codeInvalid  = "invalid"
	codeInternal = "internal"
)

// MountLines mounts the night line, the account's recent line and a character's history
// (§17.1), and the moderators' sample route. Mount (Task 30) mounts the fight route.
func MountLines(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET /v1/reports/{id}/coach/night", auth.Require(s.nightLine))
	mux.HandleFunc("GET /v1/me/coach/recent", auth.Require(s.recentLine))
	mux.HandleFunc("GET /v1/characters/{region}/{ruleset}/{name}/coach/history", auth.Require(s.history))
	mux.HandleFunc("GET /v1/admin/coach/samples", SamplesHandler(s.Store.DB))
}

// nightDTO is the night route's body.
type nightDTO struct {
	Stem         string `json:"stem"`
	On           int    `json:"on"`
	Of           int    `json:"of"`
	FindingCount int    `json:"finding_count"`
	FightIndex   int    `json:"fight_index"`
}

// nightLine is GET /v1/reports/{id}/coach/night?player=: the owner with log_coach only (403
// no_plan without it); anyone else 404. No line answers null.
func (s *Service) nightLine(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	rep, ok := s.viewableReport(w, r)
	if !ok {
		return
	}
	playerKey := r.URL.Query().Get("player")
	access, _, err := s.AccessFor(r.Context(), auth.ActorFrom(r.Context()), rep, playerKey)
	if err != nil {
		s.linesFail(w, r, "night access", err)
		return
	}
	if !ownerGate(w, r, rep, access) {
		return
	}
	line, ok, _, err := s.nightOf(r.Context(), rep, playerKey, 0)
	if err != nil {
		s.linesFail(w, r, "night", err)
		return
	}
	if !ok {
		httpx.WriteOK(w, r, http.StatusOK, nil)
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, nightDTO{Stem: line.Stem, On: line.On, Of: line.Of,
		FindingCount: line.FindingCount, FightIndex: line.FightIndex})
}

// ownerGate lets AccessFull through; the teaser (an owner without log_coach) is 403
// no_plan; an officer, a moderator and anyone else get the 404 of an unknown report. A
// flagged report answers 404 flagged, but only to a viewer who may see it.
func ownerGate(w http.ResponseWriter, r *http.Request, rep reports.Report, access Access) bool {
	switch {
	case access != AccessNone && rep.Flagged != nil:
		httpx.WriteError(w, r, http.StatusNotFound, codeFlagged, "this report is flagged", nil)
	case access == AccessFull:
		return true
	case access == AccessTeaser:
		httpx.WriteError(w, r, http.StatusForbidden, codeNoPlan, "the coach is a Premium feature", nil)
	default:
		httpx.WriteError(w, r, http.StatusNotFound, codeNotFound, "no such report", nil)
	}
	return false
}

// viewableReport reads the path's report: 404 not_found when it does not exist.
func (s *Service) viewableReport(w http.ResponseWriter, r *http.Request) (reports.Report, bool) {
	rep, err := s.Reports.Get(r.Context(), r.PathValue("id"))
	if errors.Is(err, reports.ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, codeNotFound, "no such report", nil)
		return reports.Report{}, false
	}
	if err != nil {
		s.linesFail(w, r, "report", err)
		return reports.Report{}, false
	}
	return rep, true
}

// linesFail logs err and answers 500.
func (s *Service) linesFail(w http.ResponseWriter, r *http.Request, op string, err error) {
	s.logger().Error("coach", slog.String("op", op), slog.String("error", err.Error()))
	httpx.WriteError(w, r, http.StatusInternalServerError, codeInternal, "the coach could not answer just now", nil)
}
