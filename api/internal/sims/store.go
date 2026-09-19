package sims

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// ErrNotFound is returned for a sim id that was never saved.
var ErrNotFound = errors.New("sims: not found")

// isNoRows is pgx's empty-result error, which several reads here treat
// as an answer rather than a failure.
func isNoRows(err error) bool { return errors.Is(err, pgx.ErrNoRows) }

// The states a sims row can be in. A browser result is written done;
// a server run walks queued to running to done or error.
const (
	StateQueued  = "queued"
	StateRunning = "running"
	StateDone    = "done"
	StateError   = "error"
)

// PerPage is the page size of the caller's own sim list, the same
// hundred the rankings and report lists use.
const PerPage = 100

// Store is every sims-table read and write.
type Store struct{ Pool *pgxpool.Pool }

// Row is one line of the caller's own sim history.
type Row struct {
	SimID         string    `json:"sim_id"`
	Spec          string    `json:"spec"`
	DPS           float64   `json:"dps"`
	EngineVersion string    `json:"engine_version"`
	CreatedAt     time.Time `json:"created_at"`
	Title         string    `json:"title"`
}

// Page is a page of that history.
type Page struct {
	Rows    []Row `json:"rows"`
	Total   int   `json:"total"`
	Page    int   `json:"page"`
	PerPage int   `json:"per_page"`
}

// Progress is what a page following a server run is shown.
type Progress struct {
	State          string   `json:"state"`
	IterationsDone int      `json:"iterations_done"`
	DPS            *float64 `json:"dps,omitempty"`
}

// Save writes a finished result. userID may be nil: an anonymous
// browser run is saved and shareable, it simply has no owner and so
// never appears in anyone's history. A repeated id keeps the first
// save, the way a repeated build save does.
func (s *Store) Save(ctx context.Context, id string, userID *int64, title string, res simapi.SimResult) error {
	res.SimID = id
	body, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("sims: encode %s: %w", id, err)
	}
	var t *string
	if title != "" {
		t = &title
	}
	_, err = s.Pool.Exec(ctx,
		`insert into sims (id, user_id, spec, engine_version, lane, dps_mean, dps_error,
		   iterations, title, result, state)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		 on conflict (id) do nothing`,
		id, userID, res.Request.Spec, res.EngineVersion, res.Lane, res.DPS.Mean, res.DPS.Error,
		res.IterationsRun, t, body, StateDone)
	if err != nil {
		return fmt.Errorf("sims: save %s: %w", id, err)
	}
	return nil
}

// Queue writes the placeholder row a server run starts from, so the
// page has something to poll before the job has produced anything and
// the job has the request to read back.
func (s *Store) Queue(ctx context.Context, id string, userID int64, req simapi.SimRequest) error {
	stub := simapi.SimResult{
		SimID: id, EngineVersion: req.EngineVersion, Lane: simapi.LaneServer, Request: req,
	}
	body, err := json.Marshal(stub)
	if err != nil {
		return fmt.Errorf("sims: encode %s: %w", id, err)
	}
	_, err = s.Pool.Exec(ctx,
		`insert into sims (id, user_id, spec, engine_version, lane, dps_mean, dps_error,
		   iterations, result, state)
		 values ($1, $2, $3, $4, $5, 0, 0, $6, $7, $8)`,
		id, userID, req.Spec, req.EngineVersion, simapi.LaneServer, req.Iterations, body, StateQueued)
	if err != nil {
		return fmt.Errorf("sims: queue %s: %w", id, err)
	}
	return nil
}

// Get reads one stored result.
func (s *Store) Get(ctx context.Context, id string) (simapi.SimResult, error) {
	var body []byte
	err := s.Pool.QueryRow(ctx, `select result from sims where id = $1`, id).Scan(&body)
	if isNoRows(err) {
		return simapi.SimResult{}, ErrNotFound
	}
	if err != nil {
		return simapi.SimResult{}, fmt.Errorf("sims: read %s: %w", id, err)
	}
	var out simapi.SimResult
	if err := json.Unmarshal(body, &out); err != nil {
		return simapi.SimResult{}, fmt.Errorf("sims: decode %s: %w", id, err)
	}
	out.SimID = id
	return out, nil
}

// Mine answers one page of a user's own sims, newest first.
func (s *Store) Mine(ctx context.Context, userID int64, page int) (Page, error) {
	if page < 1 {
		page = 1
	}
	out := Page{Rows: []Row{}, Page: page, PerPage: PerPage}
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from sims where user_id = $1`, userID).Scan(&out.Total); err != nil {
		return Page{}, fmt.Errorf("sims: count: %w", err)
	}
	rows, err := s.Pool.Query(ctx,
		`select id, spec, dps_mean, engine_version, created_at, coalesce(title, '')
		 from sims where user_id = $1
		 order by created_at desc, id limit $2 offset $3`,
		userID, PerPage, (page-1)*PerPage)
	if err != nil {
		return Page{}, fmt.Errorf("sims: list: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var r Row
		if err := rows.Scan(&r.SimID, &r.Spec, &r.DPS, &r.EngineVersion, &r.CreatedAt, &r.Title); err != nil {
			return Page{}, fmt.Errorf("sims: scan: %w", err)
		}
		out.Rows = append(out.Rows, r)
	}
	return out, rows.Err()
}

// Progress reads how far a run has got. A queued run reports no
// iterations and no figure: there is nothing yet to show.
func (s *Store) Progress(ctx context.Context, id string) (Progress, error) {
	var (
		p    Progress
		mean float64
		iter int
	)
	err := s.Pool.QueryRow(ctx,
		`select state, dps_mean, iterations from sims where id = $1`, id).Scan(&p.State, &mean, &iter)
	if isNoRows(err) {
		return Progress{}, ErrNotFound
	}
	if err != nil {
		return Progress{}, fmt.Errorf("sims: progress %s: %w", id, err)
	}
	if p.State == StateRunning || p.State == StateDone {
		p.IterationsDone, p.DPS = iter, &mean
	}
	return p, nil
}

// Advance records a running job's partial estimate, so the page's DPS
// figure refines while the job is still going. A row already done is
// left alone: a progress tick that arrives after the result would
// otherwise walk the run backwards.
func (s *Store) Advance(ctx context.Context, id string, done int, mean float64) error {
	_, err := s.Pool.Exec(ctx,
		`update sims set state = $2, iterations = $3, dps_mean = $4 where id = $1 and state <> $5`,
		id, StateRunning, done, mean, StateDone)
	if err != nil {
		return fmt.Errorf("sims: advance %s: %w", id, err)
	}
	return nil
}

// Finish writes a server run's result over its placeholder. A result
// carrying its own Error is stored as an error: the engine refusing a
// request is an answer, and the page must be told rather than left
// polling.
func (s *Store) Finish(ctx context.Context, id string, res simapi.SimResult) error {
	res.SimID = id
	body, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("sims: encode %s: %w", id, err)
	}
	state := StateDone
	if res.Error != "" {
		state = StateError
	}
	_, err = s.Pool.Exec(ctx,
		`update sims set engine_version = $2, dps_mean = $3, dps_error = $4,
		   iterations = $5, result = $6, state = $7 where id = $1`,
		id, res.EngineVersion, res.DPS.Mean, res.DPS.Error, res.IterationsRun, body, state)
	if err != nil {
		return fmt.Errorf("sims: finish %s: %w", id, err)
	}
	return nil
}

// Fail marks a run that could not be started or could not complete.
// why is recorded in the error it returns and logged by the caller;
// the row carries only the state, because the page is shown it.
func (s *Store) Fail(ctx context.Context, id, why string) error {
	_, err := s.Pool.Exec(ctx, `update sims set state = $2 where id = $1`, id, StateError)
	if err != nil {
		return fmt.Errorf("sims: fail %s (%s): %w", id, why, err)
	}
	return nil
}
