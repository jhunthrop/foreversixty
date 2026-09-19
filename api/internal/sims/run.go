package sims

import (
	"context"
	"encoding/json"
	"net/http"

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
	if err := req.Validate(); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", err.Error(), nil)
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
		// "queued" forever.
		s.logger().Error("sims", "id", httpx.RequestIDFrom(r.Context()), "op", "job",
			"sim", id, "err", err)
		if err := s.Store.Fail(r.Context(), id, "the job could not be started"); err != nil {
			s.fail(w, r, "job", err, "could not start that run just now")
			return
		}
		httpx.WriteError(w, r, http.StatusBadGateway, "upstream",
			"the run was recorded but could not be started", nil)
		return
	}
	httpx.WriteOK(w, r, http.StatusAccepted, map[string]string{"sim_id": id})
}
