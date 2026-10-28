package coach

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// ReviewLeaseTTL is how long a queued generation holds its key before another view may
// queue it again.
const ReviewLeaseTTL = 5 * time.Minute

const (
	msgNotFound = "no such report"
	msgFlagged  = "that report is flagged, so its fights have no coach review"
	msgFailed   = "the coach couldn't read this pull just now"
)

// Mount registers the fight route; Task 31 adds the rest.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET /v1/reports/{id}/fights/{n}/coach", s.fightCoach)
}

// fightRequest is one parsed fight-route call.
type fightRequest struct {
	reportID string
	fight    int
	player   string
	line     bool // the row strip's read-only form: never queues a generation
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error) {
	s.logger().Error("coach", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", msgFailed, nil)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusNotFound, "not_found", msgNotFound, nil)
}

// fightCoach is GET /v1/reports/{id}/fights/{n}/coach?player=…[&line=1] (§17.1).
func (s *Service) fightCoach(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	req, ok := parseFightRequest(w, r)
	if !ok {
		return
	}
	rep, access, account, ok := s.access(w, r, req)
	if !ok {
		return
	}
	l, err := s.loadPack(r.Context(), rep, req.fight, req.player)
	if errors.Is(err, ErrNoFight) || errors.Is(err, ErrNoPlayer) {
		notFound(w, r)
		return
	}
	if err != nil {
		s.fail(w, r, "pack", err)
		return
	}
	if access == AccessTeaser {
		httpx.WriteOK(w, r, http.StatusOK, teaserOf(l.pack))
		return
	}
	sv, err := s.serve(r.Context(), req, l, account)
	if err != nil {
		s.fail(w, r, "serve", err)
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, fullOf(req, l, access, sv))
}

func parseFightRequest(w http.ResponseWriter, r *http.Request) (fightRequest, bool) {
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil || n < 1 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "the fight index must be a number, counted from one", nil)
		return fightRequest{}, false
	}
	player := r.URL.Query().Get("player")
	if !character.ValidKey(player) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "player must be a character key", nil)
		return fightRequest{}, false
	}
	return fightRequest{reportID: r.PathValue("id"), fight: n, player: player, line: r.URL.Query().Get("line") == "1"}, true
}

// access answers the report, then the row (§17.3): 404 not_found for a report the caller
// cannot see or a row they may not coach, 404 flagged for a flagged report they can see.
func (s *Service) access(w http.ResponseWriter, r *http.Request, req fightRequest) (reports.Report, Access, int64, bool) {
	ctx, actor := r.Context(), auth.ActorFrom(r.Context())
	rep, err := s.Reports.Get(ctx, req.reportID)
	if errors.Is(err, reports.ErrNotFound) || (err == nil && !reports.MayView(ctx, actor, rep, s.Accounts)) {
		notFound(w, r)
		return reports.Report{}, AccessNone, 0, false
	}
	if err != nil {
		s.fail(w, r, "report", err)
		return reports.Report{}, AccessNone, 0, false
	}
	if rep.Flagged != nil {
		httpx.WriteError(w, r, http.StatusNotFound, "flagged", msgFlagged, nil)
		return reports.Report{}, AccessNone, 0, false
	}
	access, account, err := s.AccessFor(ctx, actor, rep, req.player)
	if err != nil {
		s.fail(w, r, "access", err)
		return reports.Report{}, AccessNone, 0, false
	}
	if access == AccessNone {
		notFound(w, r)
		return reports.Report{}, AccessNone, 0, false
	}
	return rep, access, account, true
}

// serve picks the review for a full or officer viewer (§14, §15.4, §20.9): the pack's own
// status when it is not ready; the rule-only rendering while the switch is off; a stored
// model review, or a stored rules review that is not retryable; else the rule-only
// rendering, busy when a cap holds, writing when a generation is queued.
func (s *Service) serve(ctx context.Context, req fightRequest, l loaded, account int64) (served, error) {
	persona, hasPersona := s.personaFor(l.pack)
	rules := s.rulesServed(l.pack, persona, hasPersona)
	if l.pack.Status != StatusReady || !s.Enabled || !hasPersona || s.Generator == nil {
		rules.status = string(l.pack.Status)
		return rules, nil
	}
	key := NewReviewKey(req.reportID, req.fight, req.player, l.pack, s.Generator.Model, persona)
	stored, ok, err := s.Store.Review(ctx, key)
	if err != nil {
		return served{}, err
	}
	if ok && stored.final() {
		return served{doc: stored.Doc, source: stored.Source, status: string(StatusReady), persona: persona,
			hasPersona: true, at: stored.GeneratedAt}, nil
	}
	if req.line {
		return rules, nil
	}
	return s.generate(ctx, key, account, rules)
}

// rulesServed is the rule-only rendering, ready, written now.
func (s *Service) rulesServed(p Pack, persona Persona, hasPersona bool) served {
	return served{doc: RenderRules(p, persona.Drills), source: SourceRules, status: string(StatusReady),
		persona: persona, hasPersona: hasPersona, at: s.Now()}
}

// generate asks the caps without counting; a spend or account cap answers busy and queues
// nothing. Otherwise it queues one generation under a lease and answers writing.
func (s *Service) generate(ctx context.Context, key ReviewKey, account int64, rules served) (served, error) {
	refusal, err := s.Caps.Peek(ctx, account)
	if err != nil {
		s.logger().Error("coach: caps peek failed; serving the rule-only review", slog.String("error", err.Error()))
		return rules, nil
	}
	switch refusal {
	case RefusalSpendCap, RefusalAccountCap:
		rules.busy = true
		return rules, nil
	case RefusalNone:
	default:
		return rules, nil
	}
	if s.Jobs == nil {
		return rules, nil
	}
	if err := s.queue(ctx, key, account); err != nil {
		return served{}, err
	}
	rules.status = statusWriting
	return rules, nil
}

// queue takes the key's lease and runs the job with it; a live lease means a job is
// already writing, so nothing more is queued. A failed launch releases the lease.
func (s *Service) queue(ctx context.Context, key ReviewKey, account int64) error {
	holder, err := newHolder()
	if err != nil {
		return err
	}
	// The database clock judges a lease's expiry, so its end is wall-clock time too.
	took, err := s.Store.TakeLease(ctx, key, holder, time.Now().Add(ReviewLeaseTTL))
	if err != nil || !took {
		return err
	}
	args := []string{ReviewJobCommand, key.ReportID, strconv.Itoa(key.FightIndex), key.PlayerKey,
		strconv.FormatInt(account, 10), holder}
	if err := s.Jobs.Run(ctx, args...); err != nil {
		if rerr := s.Store.ReleaseLease(ctx, key, holder); rerr != nil {
			s.logger().Error("coach: release lease after a failed launch", slog.String("error", rerr.Error()))
		}
		return fmt.Errorf("coach: queue review: %w", err)
	}
	return nil
}

// newHolder is a lease holder token unique to one claim.
func newHolder() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", fmt.Errorf("coach: lease holder: %w", err)
	}
	return hex.EncodeToString(b[:]), nil
}
