package sims

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// Premiumer reads the premium flag. auth.Store satisfies it; the
// tests use a stub so the handler can be exercised without an
// accounts table.
type Premiumer interface {
	Premium(ctx context.Context, userID int64) (bool, error)
}

// failCompensationTimeout bounds the compensating write below: it
// deliberately runs on a context the request's own cancellation
// cannot touch, so it needs its own bound instead.
const failCompensationTimeout = 5 * time.Second

// run dispatches one premium server-lane sim. The row is written
// before the job is started, so the job has the request to read and
// the page has something to poll.
func (s *Service) run(w http.ResponseWriter, r *http.Request) {
	actor := auth.ActorFrom(r.Context())
	premium, err := s.Accounts.Premium(r.Context(), actor.UserID)
	if err != nil {
		s.fail(w, r, "premium", err, "could not check your account just now")
		return
	}
	if !premium {
		// 402, not 403: the account is fine, the feature is paid for.
		httpx.WriteError(w, r, http.StatusPaymentRequired, "premium_required",
			"running on our servers is a premium feature", nil)
		return
	}
	var req simapi.SimRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxRequestBytes)).Decode(&req); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"the body must be a sim request of at most 256 KB", nil)
		return
	}
	// The deployment's own pin wins: a page holding a stale bundle
	// must not pick which engine our servers run.
	req.EngineVersion = s.EngineVersion
	req.Encounter = withEncounterDefaults(req.Encounter)
	if req.Bulk != nil {
		// The lane's cap is the server's to set. A request echoes the
		// cap that bounded it (contract 1.3) so a saved request says
		// what it ran under; it is not a control the client holds.
		req.Bulk.Cap = simapi.Caps[simapi.LaneServer]
	}
	// ValidateLane, not Validate: the envelope's plain Validate checks
	// the largest lane's cap, and this is the server lane (contract A1).
	if err := req.ValidateLane(simapi.LaneServer); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", err.Error(), nil)
		return
	}
	// Bulk only, not req.Kind() != KindRun: runner.Planner's own
	// contract refuses a request with no Bulk block (ErrBadInput), and
	// a weights request carries req.Weights instead - sending one
	// through the planner would turn every stat-weights submit into a
	// 500 rather than skip a check it has no combinations to answer.
	if req.Bulk != nil && s.checkSize(w, r, req) {
		return
	}

	id := auth.Base32ID(auth.ReportIDChars)
	if err := s.Store.Queue(r.Context(), id, actor.UserID, req); err != nil {
		s.fail(w, r, "queue", err, "could not start that run just now")
		return
	}
	if err := s.Jobs.Run(r.Context(), SimRunJobCommand, id); err != nil {
		// The row exists and carries the request, so the run can be
		// retried; the row says it failed rather than sitting on
		// "queued" forever. The client disconnecting is one of the
		// reasons Jobs.Run can fail in the first place, so this write
		// must not ride the request's own (already-cancelled) context
		// - that would make the exact outcome this branch exists to
		// prevent. context.WithoutCancel keeps the request's values
		// (request id, actor) without its cancellation; its own short
		// timeout stands in for the one the request context would
		// otherwise have provided.
		s.logger().Error("sims", "id", httpx.RequestIDFrom(r.Context()), "op", "job",
			"sim", id, "err", err)
		failCtx, cancel := context.WithTimeout(context.WithoutCancel(r.Context()), failCompensationTimeout)
		defer cancel()
		if err := s.Store.Fail(failCtx, id, "the job could not be started"); err != nil {
			s.fail(w, r, "job", err, "could not start that run just now")
			return
		}
		httpx.WriteError(w, r, http.StatusBadGateway, "upstream",
			"the run was recorded but could not be started", nil)
		return
	}
	httpx.WriteOK(w, r, http.StatusAccepted, map[string]string{"sim_id": id})
}

// planTimeout bounds the plan-only subprocess. Expansion loads the item
// database and walks the candidates; it runs no iterations, so thirty
// seconds bounds something pathological rather than budgeting the work.
const planTimeout = 30 * time.Second

// checkSize counts what the request would expand to, without running any
// of it, and refuses the two ways it can be too big. It reports whether
// it has already written a response.
func (s *Service) checkSize(w http.ResponseWriter, r *http.Request, req simapi.SimRequest) bool {
	ctx, cancel := context.WithTimeout(r.Context(), planTimeout)
	defer cancel()
	plan, err := s.Planner.Plan(ctx, req)
	if err != nil {
		s.fail(w, r, "plan", err, "could not size that run just now")
		return true
	}
	if plan.Cap > 0 && plan.Combinations > plan.Cap {
		// Both numbers, because the page says how far over it is and by
		// how much to trim. error.Fields is map[string]string, so they
		// go over as decimal strings (contract 10.6).
		httpx.WriteError(w, r, http.StatusBadRequest, "cap_exceeded",
			fmt.Sprintf("that is %s combinations; a run on our servers is at most %s",
				withThousands(int64(plan.Combinations)), withThousands(int64(plan.Cap))),
			map[string]string{
				"cap":          strconv.Itoa(plan.Cap),
				"combinations": strconv.Itoa(plan.Combinations),
			})
		return true
	}
	return false
}
