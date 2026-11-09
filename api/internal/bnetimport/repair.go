package bnetimport

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// RepairSummary is what RepairRulesets changed, or in a dry run would
// change: the transaction is rolled back, so the counts are exact.
type RepairSummary struct {
	Checked int
	// Moves is every character re-keyed, old key to new.
	Moves []KeyMove
	// Refetched is how many profiles were fetched because none was stored.
	Refetched int
	// Unresolved is every character no profile names a ruleset for yet,
	// left for the refresh or the next sign-in.
	Unresolved []Unresolved
	// Guilds is every guild re-rulesetted or merged into its twin.
	Guilds []GuildFix
	// Rows is how many rows each table.column changed.
	Rows KeyMoveCounts
}

// KeyMove is one character moved to the key its profile's ruleset names.
type KeyMove struct {
	UserID         int64
	OldKey, NewKey string
}

// Unresolved is a character whose ruleset nothing names.
type Unresolved struct {
	UserID int64
	Key    string
}

// rulesetCandidate is one Battle.net-imported character as the repair
// reads it.
type rulesetCandidate struct {
	key, region, name, realmSlug, ruleset, profileRuleset string
	userID                                                int64
}

// RepairRulesets re-keys every Battle.net-imported character under the
// ruleset its own profile names (fetching the profile where none is
// stored), then fixes every guild left under the wrong ruleset
// (repairGuilds). It runs in one transaction, committed only when apply is
// true: a dry run does the same work and rolls it back. It is idempotent:
// a second run finds nothing to move. It migrates nothing.
func (s *Service) RepairRulesets(ctx context.Context, apply bool) (RepairSummary, error) {
	sum := RepairSummary{Rows: KeyMoveCounts{}}
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return sum, fmt.Errorf("bnetimport: repair begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	cands, err := repairCandidates(ctx, tx)
	if err != nil {
		return sum, err
	}
	sum.Checked = len(cands)
	for _, c := range cands {
		if err := s.repairCharacter(ctx, tx, c, &sum); err != nil {
			return sum, err
		}
	}
	if err := s.repairGuilds(ctx, tx, &sum); err != nil {
		return sum, err
	}
	if !apply {
		return sum, nil
	}
	if err := tx.Commit(ctx); err != nil {
		return sum, fmt.Errorf("bnetimport: repair commit: %w", err)
	}
	return sum, nil
}

func repairCandidates(ctx context.Context, tx pgx.Tx) ([]rulesetCandidate, error) {
	rows, err := tx.Query(ctx,
		`select key, region, name, coalesce(realm_slug, ''), ruleset, coalesce(bnet_profile->'ruleset'->>'name', ''),
		        coalesce(user_id, 0)
		   from characters where source = 'bnet' order by key`)
	if err != nil {
		return nil, fmt.Errorf("bnetimport: repair candidates: %w", err)
	}
	defer rows.Close()
	var out []rulesetCandidate
	for rows.Next() {
		var c rulesetCandidate
		if err := rows.Scan(&c.key, &c.region, &c.name, &c.realmSlug, &c.ruleset, &c.profileRuleset, &c.userID); err != nil {
			return nil, fmt.Errorf("bnetimport: repair candidate: %w", err)
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// repairCharacter moves one character to the key its profile's ruleset
// names, fetching and capturing the profile when none is stored.
func (s *Service) repairCharacter(ctx context.Context, tx pgx.Tx, c rulesetCandidate, sum *RepairSummary) error {
	ruleset, ok := character.RulesetNamed(c.profileRuleset)
	if !ok {
		var err error
		if ruleset, ok, err = s.refetchRuleset(ctx, tx, c); err != nil {
			return err
		}
		sum.Refetched++
	}
	if !ok {
		sum.Unresolved = append(sum.Unresolved, Unresolved{UserID: c.userID, Key: c.key})
		return nil
	}
	if ruleset == c.ruleset {
		return nil
	}
	newKey := character.Key(c.region, ruleset, c.name)
	counts, moved, err := s.moveKey(ctx, tx, c.key, newKey, ruleset)
	if err != nil {
		return err
	}
	if !moved {
		return fmt.Errorf("bnetimport: repair: %s cannot move to %s, which another row holds", c.key, newKey)
	}
	for k, n := range counts {
		sum.Rows[k] += n
	}
	sum.Moves = append(sum.Moves, KeyMove{UserID: c.userID, OldKey: c.key, NewKey: newKey})
	return nil
}

// refetchRuleset reads the character's public profile (the app token; no
// account token is needed for it) and stores it, as the enrich pass does.
func (s *Service) refetchRuleset(ctx context.Context, tx pgx.Tx, c rulesetCandidate) (string, bool, error) {
	profile, raw, err := s.Client.Character(ctx, c.region, c.realmSlug, c.name)
	if errors.Is(err, bnetapi.ErrNotFound) || errors.Is(err, bnetapi.ErrForbidden) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("bnetimport: repair profile of %s: %w", c.key, err)
	}
	if err := s.captureProfile(ctx, tx, c.key, profile, raw); err != nil {
		return "", false, err
	}
	return profile.Ruleset, profile.Ruleset != "", nil
}
