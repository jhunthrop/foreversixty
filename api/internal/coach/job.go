package coach

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
)

// reviewJobArgs is coach-review's argument count after the command.
const reviewJobArgs = 5

// reviewJob is one parsed coach-review call.
type reviewJob struct {
	reportID string
	fight    int
	player   string
	account  int64
	holder   string
}

func parseReviewJob(args []string) (reviewJob, error) {
	if len(args) != reviewJobArgs {
		return reviewJob{}, fmt.Errorf("coach-review: want <report> <fight> <player_key> <account_id> <holder>, got %d args", len(args))
	}
	fight, err := strconv.Atoi(args[1])
	if err != nil || fight < 1 {
		return reviewJob{}, fmt.Errorf("coach-review: fight %q is not a fight index", args[1])
	}
	account, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil || account < 1 {
		return reviewJob{}, fmt.Errorf("coach-review: account %q is not an account id", args[3])
	}
	if args[0] == "" || args[2] == "" || args[4] == "" {
		return reviewJob{}, errors.New("coach-review: report, player and holder are required")
	}
	return reviewJob{reportID: args[0], fight: fight, player: args[2], account: account, holder: args[4]}, nil
}

// RunReviewJob is the coach-review job: coach-review <report> <fight> <player_key>
// <account_id> <holder>. It loads the pack (cached, else built), the persona and the
// fight's roster, writes the review through a Generator of its own, and releases the lease
// the route took under holder when it ends, on every path: by holder, not by key, so an
// early error, a pack that is not ready, or a version that moved since the route leased
// never leaves "writing" held.
func RunReviewJob(ctx context.Context, s *Service, args []string) error {
	job, err := parseReviewJob(args)
	if err != nil {
		return err
	}
	defer s.release(job.holder)
	return s.runReview(ctx, job)
}

func (s *Service) runReview(ctx context.Context, job reviewJob) error {
	if s.Generator == nil {
		return errors.New("coach-review: no generator")
	}
	rep, err := s.Reports.Get(ctx, job.reportID)
	if err != nil {
		return fmt.Errorf("coach-review: report %s: %w", job.reportID, err)
	}
	entry, err := s.fight(ctx, rep.ID, job.fight)
	if err != nil {
		return err
	}
	sum, err := s.summary(ctx, rep.ID, job.fight)
	if err != nil {
		return err
	}
	l, ok, err := s.cachedPack(ctx, rep.ID, entry, job.player)
	if err == nil && !ok {
		l, err = s.buildPack(ctx, rep, entry, sum, job.player)
	}
	if err != nil {
		return err
	}
	persona, hasPersona := s.personaFor(l.pack)
	if l.pack.Status != StatusReady || !hasPersona {
		return nil // nothing to write: the route serves the pack's own status
	}
	key := NewReviewKey(rep.ID, job.fight, job.player, l.pack, s.Generator.Model, persona)
	return s.write(ctx, key, l.pack, persona, fightPlayerNames(sum), job.account)
}

// write generates unless a model review, or a rules review that is not retryable, is
// already stored for the key.
func (s *Service) write(ctx context.Context, key ReviewKey, p Pack, persona Persona, roster []string, account int64) error {
	stored, ok, err := s.Store.Review(ctx, key)
	if err != nil {
		return err
	}
	if ok && stored.final() {
		return nil
	}
	g := *s.Generator
	g.Validator = Validator{Names: s.Names, Roster: roster}
	if _, err := g.Generate(ctx, key, p, persona, account); err != nil {
		return fmt.Errorf("coach-review: %w", err)
	}
	return nil
}

// final reports a stored review that is served as is and never regenerated: a model
// review, or a rules review that is not retryable (§14).
func (r StoredReview) final() bool { return r.Source == SourceModel || !r.Retryable }

// release deletes the holder's lease; it runs on its own context so a cancelled job still
// frees it.
func (s *Service) release(holder string) {
	if err := s.Store.ReleaseHolder(context.Background(), holder); err != nil {
		s.logger().Error("coach-review: release lease", slog.String("error", err.Error()))
	}
}

const releaseHolderSQL = `delete from coach_review_leases where holder = $1`

// ReleaseHolder deletes every lease holder holds. A holder token is unique to one claim,
// so this frees exactly the lease the route took, whatever key the job ends up computing.
func (s *Store) ReleaseHolder(ctx context.Context, holder string) error {
	if holder == "" {
		return errors.New("coach: release needs a holder")
	}
	if _, err := s.DB.Exec(ctx, releaseHolderSQL, holder); err != nil {
		return fmt.Errorf("coach: release lease of %s: %w", holder, err)
	}
	return nil
}
