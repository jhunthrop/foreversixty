// api/internal/bnetimport/unknown_realm.go
package bnetimport

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/directory"
)

// realmAnswerQuorum is how many players must give a realm the same answer,
// with none giving another, before the import keys that realm's characters by
// it for everyone (realmConsensus).
const realmAnswerQuorum = 2

// ErrInvalidAnswer is AnswerRealm's refusal for a ruleset that is not one of
// the four.
var ErrInvalidAnswer = errors.New("bnetimport: not a ruleset")

// ErrNotAwaitingAnswer is AnswerRealm's refusal for a key that is not one of
// the player's own characters on an unknown realm.
var ErrNotAwaitingAnswer = errors.New("bnetimport: not a character of yours on an unknown realm")

// ErrAnswerKeyTaken is AnswerRealm's refusal when the answered key already
// belongs to another account (the same name under that ruleset).
var ErrAnswerKeyTaken = errors.New("bnetimport: that ruleset's key belongs to another account")

// rowQuerier is what both a pool and a transaction answer.
type rowQuerier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// realmAnswerCounts is how many players gave each ruleset for a realm.
func realmAnswerCounts(ctx context.Context, q rowQuerier, region string, realmID int64) (map[string]int, error) {
	rows, err := q.Query(ctx,
		`select ruleset, count(*) from realm_ruleset_answers where region = $1 and realm_id = $2 group by ruleset`,
		region, realmID)
	if err != nil {
		return nil, fmt.Errorf("bnetimport: realm answers %s/%d: %w", region, realmID, err)
	}
	counts := map[string]int{}
	for rows.Next() {
		var ruleset string
		var n int
		if err := rows.Scan(&ruleset, &n); err != nil {
			rows.Close()
			return nil, fmt.Errorf("bnetimport: scan realm answers %s/%d: %w", region, realmID, err)
		}
		counts[ruleset] = n
	}
	return counts, rows.Err()
}

// realmConsensus is the ruleset the players of a realm agree on: ok only when
// one ruleset holds realmAnswerQuorum or more answers and no other has any.
// A disagreement resolves nothing; AnswerRealm logs it.
func realmConsensus(ctx context.Context, q rowQuerier, region string, realmID int64) (string, bool, error) {
	if realmID == 0 {
		return "", false, nil
	}
	counts, err := realmAnswerCounts(ctx, q, region, realmID)
	if err != nil || len(counts) != 1 {
		return "", false, err
	}
	for ruleset, n := range counts {
		return ruleset, n >= realmAnswerQuorum, nil
	}
	return "", false, nil
}

// rulesetFromOwnRow is the ruleset this character's own Battle.net row is
// keyed by, when it is keyed by one: the player's answer (Answered) or what an
// earlier import's source named. A row still on character.RulesetUnknown names
// nothing.
func (s *Service) rulesetFromOwnRow(ctx context.Context, userID int64, region string, bnetID int64) (resolvedRuleset, bool, error) {
	var resolved resolvedRuleset
	err := s.Pool.QueryRow(ctx,
		`select ruleset, ruleset_answered from characters
		  where bnet_character_id = $1 and user_id = $2 and region = $3 and ruleset <> $4`,
		bnetID, userID, region, character.RulesetUnknown).Scan(&resolved.Ruleset, &resolved.Answered)
	if errors.Is(err, pgx.ErrNoRows) {
		return resolvedRuleset{}, false, nil
	}
	if err != nil {
		return resolvedRuleset{}, false, fmt.Errorf("bnetimport: own row ruleset of %d: %w", bnetID, err)
	}
	return resolved, true, nil
}

// carryProvisionalKey moves this character's row, with every row keyed by it
// (moveKey), to ref's key when the row was keyed provisionally: under
// character.RulesetUnknown, or by the player's answer, which a realm source
// has now overruled. Any other differing key is rekeyIfNeeded's to settle.
func (s *Service) carryProvisionalKey(ctx context.Context, tx pgx.Tx, ref characterRef) error {
	var oldKey string
	err := tx.QueryRow(ctx,
		`select key from characters
		  where bnet_character_id = $1 and user_id = $2 and region = $3 and key <> $4
		    and (ruleset = $5 or ruleset_answered)`,
		ref.BnetID, ref.UserID, ref.Region, ref.Key, character.RulesetUnknown).Scan(&oldKey)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("bnetimport: provisional key of %s: %w", ref.Key, err)
	}
	_, _, err = s.moveKey(ctx, tx, oldKey, ref.Key, ref.Ruleset)
	return err
}

// resolveUnknownInPlace is the nightly refresh's step for a row still on
// character.RulesetUnknown: the learned realm table, the character directory,
// the committed realm table or the realm's agreed answers may name its ruleset
// since the import, and the row moves to it. Returns the key and ruleset to
// carry on with.
func (s *Service) resolveUnknownInPlace(ctx context.Context, tx pgx.Tx, key, region, ruleset, name string,
	realmID, bnetID int64) (string, string, error) {
	if ruleset != character.RulesetUnknown {
		return key, ruleset, nil
	}
	resolved, ok, err := s.unknownRealmRuleset(ctx, tx, region, realmID, bnetID)
	if err != nil || !ok {
		return key, ruleset, err
	}
	newKey := character.Key(region, resolved, name)
	_, moved, err := s.moveKey(ctx, tx, key, newKey, resolved)
	if err != nil || !moved {
		return key, ruleset, err
	}
	return newKey, resolved, nil
}

// unknownRealmRuleset is the first source that names a ruleset for a row on an unknown realm,
// in rulesetFor's order without a Blizzard call.
func (s *Service) unknownRealmRuleset(ctx context.Context, tx pgx.Tx, region string, realmID, bnetID int64) (string, bool, error) {
	if ruleset, ok, err := directoryRuleset(ctx, tx, region, realmID, bnetID); err != nil || ok {
		return ruleset, ok, err
	}
	if realmID == 0 {
		return "", false, nil
	}
	if ruleset, ok := s.rulesetForRealmID(region, nil, realmID); ok {
		return ruleset, true, nil
	}
	return realmConsensus(ctx, tx, region, realmID)
}

// directoryRuleset is what the character directory knows without a Blizzard call: the realm
// id's learned ruleset, then the directory's own row for the character id.
func directoryRuleset(ctx context.Context, q directory.Querier, region string, realmID, bnetID int64) (string, bool, error) {
	if ruleset, ok, err := directory.RealmRuleset(ctx, q, region, realmID); err != nil || ok {
		return ruleset, ok, err
	}
	return directory.CharacterRuleset(ctx, q, region, bnetID)
}

// awaitingRow is a character on an unknown realm, as AnswerRealm reads it.
type awaitingRow struct {
	Key     string
	UserID  int64
	Region  string
	Name    string
	BnetID  int64
	RealmID int64
}

// AnswerRealm is the player's answer to "Which realm type is <character> on?":
// the character (one of userID's own, still on character.RulesetUnknown) moves
// to ruleset's key with every row keyed by it, and so do the player's other
// characters on the same realm. The answer is recorded as evidence for the
// realm; once realmAnswerQuorum players agree with none disagreeing, every
// character on the realm still unknown moves too, and the import keys the
// realm by it from then on. A disagreement is logged and applies nothing
// beyond the answering player's own characters. Returns the new key.
func (s *Service) AnswerRealm(ctx context.Context, userID int64, key, ruleset string) (string, error) {
	if !character.ValidRuleset(ruleset) {
		return "", ErrInvalidAnswer
	}
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("bnetimport: begin realm answer: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	var row awaitingRow
	err = tx.QueryRow(ctx,
		`select key, user_id, region, name, coalesce(bnet_character_id, 0), coalesce(bnet_realm_id, 0)
		   from characters where key = $1 and user_id = $2 and ruleset = $3 for update`,
		key, userID, character.RulesetUnknown).
		Scan(&row.Key, &row.UserID, &row.Region, &row.Name, &row.BnetID, &row.RealmID)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrNotAwaitingAnswer
	}
	if err != nil {
		return "", fmt.Errorf("bnetimport: read awaiting %s: %w", key, err)
	}
	newKey, err := s.placeAnswer(ctx, tx, row, ruleset)
	if err != nil {
		return "", err
	}
	if row.RealmID != 0 {
		if err := s.recordRealmAnswer(ctx, tx, row, ruleset); err != nil {
			return "", err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("bnetimport: commit realm answer %s: %w", key, err)
	}
	return newKey, nil
}

// placeAnswer moves the answered character to its ruleset's key and marks the
// ruleset as the player's answer. A key the same account already holds (its
// addon row of the same character) takes the Battle.net identity instead; a
// key another account holds is refused.
func (s *Service) placeAnswer(ctx context.Context, tx pgx.Tx, row awaitingRow, ruleset string) (string, error) {
	newKey := character.Key(row.Region, ruleset, row.Name)
	_, moved, err := s.moveKey(ctx, tx, row.Key, newKey, ruleset)
	if err != nil {
		return "", err
	}
	if !moved {
		if err := s.mergeIntoOwnRow(ctx, tx, row, newKey); err != nil {
			return "", err
		}
	}
	if _, err := tx.Exec(ctx, `update characters set ruleset_answered = true where key = $1`, newKey); err != nil {
		return "", fmt.Errorf("bnetimport: mark answered %s: %w", newKey, err)
	}
	return newKey, nil
}

// mergeIntoOwnRow folds an unknown-realm row into the account's own row at
// newKey (rekeyIfNeeded drops the unknown row), carrying the Battle.net
// identity across; ErrAnswerKeyTaken when newKey belongs to anyone else.
func (s *Service) mergeIntoOwnRow(ctx context.Context, tx pgx.Tx, row awaitingRow, newKey string) error {
	var owner *int64
	if err := tx.QueryRow(ctx, `select user_id from characters where key = $1`, newKey).Scan(&owner); err != nil {
		return fmt.Errorf("bnetimport: answer key owner %s: %w", newKey, err)
	}
	if owner == nil || *owner != row.UserID {
		return ErrAnswerKeyTaken
	}
	if err := s.rekeyIfNeeded(ctx, tx, row.UserID, row.Region, newKey, row.BnetID); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx,
		`update characters set bnet_character_id = nullif($2::bigint, 0), bnet_realm_id = nullif($3::bigint, 0) where key = $1`,
		newKey, row.BnetID, row.RealmID); err != nil {
		return fmt.Errorf("bnetimport: answer merge %s: %w", newKey, err)
	}
	return nil
}

// recordRealmAnswer stores the player's answer as evidence for the realm (one
// row per player per realm, a later answer replacing an earlier one), logs
// a disagreement, and moves the realm's other unknown characters: every
// player's once the realm reaches its quorum, the answering player's own
// otherwise.
func (s *Service) recordRealmAnswer(ctx context.Context, tx pgx.Tx, row awaitingRow, ruleset string) error {
	if _, err := tx.Exec(ctx,
		`insert into realm_ruleset_answers (region, realm_id, user_id, ruleset) values ($1, $2, $3, $4)
		 on conflict (region, realm_id, user_id) do update set ruleset = excluded.ruleset, answered_at = now()`,
		row.Region, row.RealmID, row.UserID, ruleset); err != nil {
		return fmt.Errorf("bnetimport: record realm answer %s/%d: %w", row.Region, row.RealmID, err)
	}
	counts, err := realmAnswerCounts(ctx, tx, row.Region, row.RealmID)
	if err != nil {
		return err
	}
	if len(counts) > 1 {
		s.logger().Warn("bnetimport", "op", "realm_answers_disagree", "region", row.Region,
			"realm_id", row.RealmID, "answers", formatCounts(counts))
	}
	agreed, ok, err := realmConsensus(ctx, tx, row.Region, row.RealmID)
	if err != nil {
		return err
	}
	everyone := ok && agreed == ruleset
	s.logger().Info("bnetimport", "op", "realm_answer", "region", row.Region, "realm_id", row.RealmID,
		"ruleset", ruleset, "answers", formatCounts(counts), "resolved", everyone)
	return s.moveRealmUnknowns(ctx, tx, row, ruleset, everyone)
}

// moveRealmUnknowns moves the realm's characters still on
// character.RulesetUnknown to ruleset: every account's when everyone is true,
// else the answering player's own. The answering player's are marked as their
// answer; the rest are keyed by the realm's agreed answers.
func (s *Service) moveRealmUnknowns(ctx context.Context, tx pgx.Tx, answered awaitingRow, ruleset string, everyone bool) error {
	rows, err := tx.Query(ctx,
		`select key, user_id, name from characters
		  where region = $1 and bnet_realm_id = $2 and ruleset = $3 and ($4 or user_id = $5)
		  order by key for update`,
		answered.Region, answered.RealmID, character.RulesetUnknown, everyone, answered.UserID)
	if err != nil {
		return fmt.Errorf("bnetimport: realm unknowns %s/%d: %w", answered.Region, answered.RealmID, err)
	}
	var pending []awaitingRow
	for rows.Next() {
		var r awaitingRow
		if err := rows.Scan(&r.Key, &r.UserID, &r.Name); err != nil {
			rows.Close()
			return fmt.Errorf("bnetimport: scan realm unknown: %w", err)
		}
		pending = append(pending, r)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("bnetimport: realm unknowns %s/%d: %w", answered.Region, answered.RealmID, err)
	}
	for _, r := range pending {
		newKey := character.Key(answered.Region, ruleset, r.Name)
		_, moved, err := s.moveKey(ctx, tx, r.Key, newKey, ruleset)
		if err != nil {
			return err
		}
		if moved && r.UserID == answered.UserID {
			if _, err := tx.Exec(ctx, `update characters set ruleset_answered = true where key = $1`, newKey); err != nil {
				return fmt.Errorf("bnetimport: mark answered %s: %w", newKey, err)
			}
		}
	}
	return nil
}

// formatCounts is a realm's answers for a log line: "normal:2,pvp:1".
func formatCounts(counts map[string]int) string {
	parts := make([]string, 0, len(counts))
	for ruleset, n := range counts {
		parts = append(parts, ruleset+":"+strconv.Itoa(n))
	}
	sort.Strings(parts)
	return strings.Join(parts, ",")
}
