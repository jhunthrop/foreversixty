package coach

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Review sources (coach_reviews.source).
const (
	SourceModel = "model"
	SourceRules = "rules"
)

// ReviewKey is one generation's identity: a review is written once per key (§14).
type ReviewKey struct {
	ReportID       string
	FightIndex     int
	PlayerKey      string
	PackVersion    string
	ModelVersion   string
	PersonaVersion string
}

// NewReviewKey is the key for a review of p written by modelSlug in persona's voice. The
// model version (slug, rules, schema, prompt template) and the persona version (the persona
// file's content hash) are separate parts: ModelVersion does not cover the persona body.
func NewReviewKey(reportID string, fightIndex int, playerKey string, p Pack, modelSlug string, persona Persona) ReviewKey {
	return ReviewKey{ReportID: reportID, FightIndex: fightIndex, PlayerKey: playerKey, PackVersion: p.Version,
		ModelVersion: ModelVersion(modelSlug), PersonaVersion: persona.Version}
}

// StoredReview is one coach_reviews row. Tokens and cost are summed over every model
// attempt; ModelSlug and RequestID are the last attempt's.
type StoredReview struct {
	Key           ReviewKey
	Doc           ReviewDoc
	Source        string // SourceModel or SourceRules
	Reason        string
	Retryable     bool
	TokensIn      int
	TokensOut     int
	CostUSD       float64
	CostEstimated bool
	ModelSlug     string
	RequestID     string
	AccountID     int64 // 0 is no account (stored as null)
	GeneratedAt   time.Time
	// Validation is the last model attempt's problems; empty for a pass or a rules row
	// that never called the model.
	Validation []Problem
}

// Store reads and writes the coach's stored reviews and their leases.
type Store struct {
	DB *pgxpool.Pool
}

const saveReviewSQL = `
insert into coach_reviews (report_id, fight_index, player_key, pack_version, model_version, persona_version,
  review, source, reason, retryable, validation, tokens_in, tokens_out, cost_usd, cost_estimated,
  model_slug, gateway_request_id, account_id, generated_at)
values ($1, $2, $3, $4, $5, $6, $7, $8, nullif($9, ''), $10, $11, $12, $13, $14, $15,
  nullif($16, ''), nullif($17, ''), nullif($18, 0), $19)
on conflict (report_id, fight_index, player_key, pack_version, model_version, persona_version) do update set
  review = excluded.review, source = excluded.source, reason = excluded.reason, retryable = excluded.retryable,
  validation = excluded.validation, tokens_in = excluded.tokens_in, tokens_out = excluded.tokens_out,
  cost_usd = excluded.cost_usd, cost_estimated = excluded.cost_estimated, model_slug = excluded.model_slug,
  gateway_request_id = excluded.gateway_request_id, account_id = excluded.account_id,
  generated_at = excluded.generated_at
  where excluded.source = 'model' or coach_reviews.source <> 'model'`

const reviewSQL = `
select review, source, coalesce(reason, ''), retryable, validation, tokens_in, tokens_out, cost_usd::float8,
  cost_estimated, coalesce(model_slug, ''), coalesce(gateway_request_id, ''), coalesce(account_id, 0), generated_at
from coach_reviews
where report_id = $1 and fight_index = $2 and player_key = $3 and pack_version = $4 and model_version = $5
  and persona_version = $6`

// SaveReview upserts the row on its key. A rules row never replaces a stored model review
// for the same key: that save is a no-op (a refusal or model error must not erase one).
func (s *Store) SaveReview(ctx context.Context, r StoredReview) error {
	if r.Source != SourceModel && r.Source != SourceRules {
		return fmt.Errorf("coach: review source %q is neither model nor rules", r.Source)
	}
	doc, err := json.Marshal(r.Doc)
	if err != nil {
		return fmt.Errorf("coach: encode review: %w", err)
	}
	validation, err := json.Marshal(storedProblemsOf(r.Validation))
	if err != nil {
		return fmt.Errorf("coach: encode validation: %w", err)
	}
	k := r.Key
	_, err = s.DB.Exec(ctx, saveReviewSQL, k.ReportID, k.FightIndex, k.PlayerKey, k.PackVersion, k.ModelVersion,
		k.PersonaVersion, doc, r.Source, r.Reason, r.Retryable, validation, r.TokensIn, r.TokensOut, r.CostUSD,
		r.CostEstimated, r.ModelSlug, r.RequestID, r.AccountID, r.GeneratedAt)
	if err != nil {
		return fmt.Errorf("coach: save review %s/%d: %w", k.ReportID, k.FightIndex, err)
	}
	return nil
}

// Review reads the stored row for k; false when there is none.
func (s *Store) Review(ctx context.Context, k ReviewKey) (StoredReview, bool, error) {
	r := StoredReview{Key: k}
	var doc, validation []byte
	err := s.DB.QueryRow(ctx, reviewSQL, k.ReportID, k.FightIndex, k.PlayerKey, k.PackVersion, k.ModelVersion,
		k.PersonaVersion).Scan(&doc, &r.Source, &r.Reason, &r.Retryable, &validation, &r.TokensIn, &r.TokensOut,
		&r.CostUSD, &r.CostEstimated, &r.ModelSlug, &r.RequestID, &r.AccountID, &r.GeneratedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return StoredReview{}, false, nil
	}
	if err != nil {
		return StoredReview{}, false, fmt.Errorf("coach: read review %s/%d: %w", k.ReportID, k.FightIndex, err)
	}
	if err := json.Unmarshal(doc, &r.Doc); err != nil {
		return StoredReview{}, false, fmt.Errorf("coach: decode review %s/%d: %w", k.ReportID, k.FightIndex, err)
	}
	var stored []storedProblem
	if err := json.Unmarshal(validation, &stored); err != nil {
		return StoredReview{}, false, fmt.Errorf("coach: decode validation %s/%d: %w", k.ReportID, k.FightIndex, err)
	}
	r.Validation = problemsOf(stored)
	return r, true, nil
}

// storedProblem is a Problem as coach_reviews.validation holds it.
type storedProblem struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
	Near  string `json:"near,omitempty"`
}

// storedProblemsOf never returns nil, so an empty list is stored as [].
func storedProblemsOf(ps []Problem) []storedProblem {
	out := make([]storedProblem, 0, len(ps))
	for _, p := range ps {
		out = append(out, storedProblem(p))
	}
	return out
}

func problemsOf(ss []storedProblem) []Problem {
	if len(ss) == 0 {
		return nil
	}
	out := make([]Problem, 0, len(ss))
	for _, s := range ss {
		out = append(out, Problem(s))
	}
	return out
}

// takeLeaseSQL inserts the lease, or takes over one whose leased_until has passed (by the
// database clock); a live lease returns no row.
const takeLeaseSQL = `
insert into coach_review_leases (report_id, fight_index, player_key, pack_version, model_version, persona_version,
  holder, leased_until)
values ($1, $2, $3, $4, $5, $6, $7, $8)
on conflict (report_id, fight_index, player_key, pack_version, model_version, persona_version) do update
  set holder = excluded.holder, leased_until = excluded.leased_until
  where coach_review_leases.leased_until < now()
returning holder`

const releaseLeaseSQL = `
delete from coach_review_leases
where report_id = $1 and fight_index = $2 and player_key = $3 and pack_version = $4 and model_version = $5
  and persona_version = $6 and holder = $7`

// TakeLease claims k for holder until the given instant. It is false while another
// holder's lease is live; an expired lease is taken over. holder is a token unique to the
// claim (a uuid), so a worker whose lease was taken over cannot release the new one.
func (s *Store) TakeLease(ctx context.Context, k ReviewKey, holder string, until time.Time) (bool, error) {
	if holder == "" {
		return false, errors.New("coach: a lease needs a holder")
	}
	var got string
	err := s.DB.QueryRow(ctx, takeLeaseSQL, k.ReportID, k.FightIndex, k.PlayerKey, k.PackVersion, k.ModelVersion,
		k.PersonaVersion, holder, until).Scan(&got)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("coach: take lease %s/%d: %w", k.ReportID, k.FightIndex, err)
	}
	return true, nil
}

// ReleaseLease deletes k's lease only when holder still holds it.
func (s *Store) ReleaseLease(ctx context.Context, k ReviewKey, holder string) error {
	_, err := s.DB.Exec(ctx, releaseLeaseSQL, k.ReportID, k.FightIndex, k.PlayerKey, k.PackVersion, k.ModelVersion,
		k.PersonaVersion, holder)
	if err != nil {
		return fmt.Errorf("coach: release lease %s/%d: %w", k.ReportID, k.FightIndex, err)
	}
	return nil
}
