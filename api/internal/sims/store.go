package sims

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
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

// MaxPage bounds the page a list may be asked for. Past it
// (page-1)*PerPage overflows int and Postgres is handed a negative
// OFFSET, which is an error where an empty page is the true answer.
// Ten million rows deep is far past any history anyone has, so the
// clamp costs nothing real.
const MaxPage = 100_000

// clampPage holds a requested page inside the range whose offset is
// representable, so an absurd page number is an empty page rather than
// a failed query.
func clampPage(page int) int {
	switch {
	case page < 1:
		return 1
	case page > MaxPage:
		return MaxPage
	default:
		return page
	}
}

// Store is every sims-table read and write.
type Store struct{ Pool *pgxpool.Pool }

// Row is one line of the caller's own sim history.
type Row struct {
	SimID string `json:"sim_id"`
	Spec  string `json:"spec"`
	// Kind is which tool produced this row: run, gear, talents, drops
	// or weights (contract 1.1).
	Kind string  `json:"kind"`
	DPS  float64 `json:"dps"`
	// Headline is the one line the list shows, composed by Headline at
	// write time and stored, because composing it on read would mean
	// detoasting the whole result blob for every row on the page.
	Headline string `json:"headline"`
	// State is the row's own state: done, error, running or queued. A
	// failed run has no headline to show, and without this the page
	// could not tell that from a row whose headline is simply missing.
	State         string    `json:"state"`
	EngineVersion string    `json:"engine_version"`
	CreatedAt     time.Time `json:"created_at"`
	Title         string    `json:"title"`
}

// Page is a page of that history.
type Page = httpx.Page[Row]

// Progress is what a page following a server run is shown. The three
// stage fields are zero for a plain run, which has no stages (contract
// 2); a bulk run's carry the last tick the job wrote.
type Progress struct {
	State          string   `json:"state"`
	IterationsDone int      `json:"iterations_done"`
	DPS            *float64 `json:"dps,omitempty"`
	Stage          int      `json:"stage"`
	CombosDone     int      `json:"combos_done"`
	CombosTotal    int      `json:"combos_total"`
}

// Tick is one progress report on its way into the row. It is this
// package's own vocabulary rather than the runner's callback type, so
// the store and its tests do not move when the engine's progress
// stream gains a field.
type Tick struct {
	IterationsDone int
	Mean           float64
	Stage          int
	CombosDone     int
	CombosTotal    int
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
		`insert into sims (id, user_id, spec, kind, headline, engine_version, lane,
		   dps_mean, dps_error, iterations, title, result, state)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		 on conflict (id) do nothing`,
		id, userID, res.Request.Spec, res.Request.Kind(), Headline(res), res.EngineVersion,
		res.Lane, res.DPS.Mean, res.DPS.Error, res.IterationsRun, t, body, StateDone)
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
		`insert into sims (id, user_id, spec, kind, engine_version, lane, dps_mean, dps_error,
		   iterations, result, state)
		 values ($1, $2, $3, $4, $5, $6, 0, 0, $7, $8, $9)`,
		id, userID, req.Spec, req.Kind(), req.EngineVersion, simapi.LaneServer,
		req.Iterations, body, StateQueued)
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

// ForBuild answers the newest done sim run against a build: the row
// whose stored request names buildID as its "build" source. This is
// how the build's shared card and unfurl page find what a build
// simmed to, once the planner has posted one. ok is false when the
// build has never been simmed, or only a queued or errored run
// exists for it.
//
// The filter walks the stored result's JSON with a path expression
// (result -> 'request' -> 'source' ->> 'kind'/'ref') rather than a
// generated column and index: EXPLAIN ANALYZE against 200k synthetic
// rows still answers in about 15ms, and this query only ever runs
// behind the build card's week-long cache, never on a hot path. See
// the task report for the full measurement.
func (s *Store) ForBuild(ctx context.Context, buildID string) (simapi.SimResult, bool, error) {
	var (
		id   string
		body []byte
	)
	err := s.Pool.QueryRow(ctx,
		`select id, result from sims
		 where state = $1
		   and result -> 'request' -> 'source' ->> 'kind' = $2
		   and result -> 'request' -> 'source' ->> 'ref' = $3
		 order by created_at desc, id desc
		 limit 1`,
		StateDone, simapi.SourceBuild, buildID).Scan(&id, &body)
	if isNoRows(err) {
		return simapi.SimResult{}, false, nil
	}
	if err != nil {
		return simapi.SimResult{}, false, fmt.Errorf("sims: for build %s: %w", buildID, err)
	}
	var out simapi.SimResult
	if err := json.Unmarshal(body, &out); err != nil {
		return simapi.SimResult{}, false, fmt.Errorf("sims: decode for build %s: %w", buildID, err)
	}
	out.SimID = id
	return out, true, nil
}

// Mine answers one page of a user's own sims, newest first. kind, when
// set, is one of simapi.Kinds and narrows the list to that tool; "" is
// every kind. The caller validates it — an unknown kind reaching here
// would simply return nothing, which reads as "you have none" rather
// than as the typo it is.
func (s *Store) Mine(ctx context.Context, userID int64, page int, kind string) (Page, error) {
	page = clampPage(page)
	out := Page{Rows: []Row{}, Page: page, PerPage: PerPage}
	// One predicate, two queries: the count and the page must agree, and
	// a literal `$2 = '' or kind = $2` would make the planner ignore
	// sims_user_kind_idx for the filtered case, which is the case the
	// index exists for.
	where, args := "user_id = $1", []any{userID}
	if kind != "" {
		where += " and kind = $2"
		args = append(args, kind)
	}
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from sims where `+where, args...).Scan(&out.Total); err != nil {
		return Page{}, fmt.Errorf("sims: count: %w", err)
	}
	rows, err := s.Pool.Query(ctx,
		`select id, spec, kind, dps_mean, headline, state, engine_version, created_at,
		        coalesce(title, '')
		 from sims where `+where+
			fmt.Sprintf(" order by created_at desc, id limit $%d offset $%d",
				len(args)+1, len(args)+2),
		append(args, PerPage, (page-1)*PerPage)...)
	if err != nil {
		return Page{}, fmt.Errorf("sims: list: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var r Row
		if err := rows.Scan(&r.SimID, &r.Spec, &r.Kind, &r.DPS, &r.Headline, &r.State,
			&r.EngineVersion, &r.CreatedAt, &r.Title); err != nil {
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
		`select state, dps_mean, iterations, stage, combos_done, combos_total
		 from sims where id = $1`, id).Scan(&p.State, &mean, &iter,
		&p.Stage, &p.CombosDone, &p.CombosTotal)
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
// figure refines and its stage line moves while the job is still
// going. A row already in a terminal state is left alone: a progress
// tick that arrives after the result, or after a failure, would
// otherwise walk the run backwards or revive an errored one.
func (s *Store) Advance(ctx context.Context, id string, t Tick) error {
	_, err := s.Pool.Exec(ctx,
		`update sims set state = $2, iterations = $3, dps_mean = $4,
		   stage = $5, combos_done = $6, combos_total = $7
		 where id = $1 and state in ($8, $9)`,
		id, StateRunning, t.IterationsDone, t.Mean,
		t.Stage, t.CombosDone, t.CombosTotal, StateQueued, StateRunning)
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
	// Only a run that succeeded gets a headline. Headline reads a failed
	// result's zero fields as an answer — "0 DPS", or "no combinations" —
	// and those are sentences claiming the run finished and found
	// nothing. A failed row carries its state instead, the same way
	// Fail's does, so the two failure modes read alike.
	headline := ""
	if state == StateDone {
		headline = Headline(res)
	}
	_, err = s.Pool.Exec(ctx,
		`update sims set engine_version = $2, dps_mean = $3, dps_error = $4,
		   iterations = $5, result = $6, state = $7, headline = $8 where id = $1`,
		id, res.EngineVersion, res.DPS.Mean, res.DPS.Error, res.IterationsRun, body, state,
		headline)
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
