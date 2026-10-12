// api/internal/guilds/jobs.go
package guilds

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// AgeingThreshold is how stale a guild_characters row's refreshed_at may
// get before AgeOut removes it — the same 45 days as the first draft
// (RULING 3, kept per the coordinator's instruction).
const AgeingThreshold = 45 * 24 * time.Hour

// SweepEvery is the MembershipJob's default tick.
const SweepEvery = 5 * time.Minute

// AgeOut removes every guild_characters row whose refreshed_at is older
// than AgeingThreshold as of at, then recomputes guild_members for every
// account it touched, in one transaction — "the account row follows": it
// downgrades or disappears exactly as if that character's export had
// gone unguilded.
func (s *Store) AgeOut(ctx context.Context, at time.Time) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: age out: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	rows, err := tx.Query(ctx,
		`delete from guild_characters where refreshed_at < $1 returning guild_id, user_id`,
		at.Add(-AgeingThreshold))
	if err != nil {
		return fmt.Errorf("guilds: age out: delete: %w", err)
	}
	affected, err := scanGuildUserPairs(rows)
	if err != nil {
		return fmt.Errorf("guilds: age out: %w", err)
	}
	// An aged-out row can take an account's last verified character with
	// it, exactly as an unguilded export would — so, unlike VerifyByLogs
	// (which only ever adds verification and can never cost a claim),
	// ageing also releases a claim the affected account no longer
	// qualifies for.
	for p := range affected {
		uid := p.userID
		if err := RecomputeMembership(ctx, tx, p.guildID, &uid); err != nil {
			return err
		}
		if err := ReleaseClaimIfLost(ctx, tx, p.guildID, uid); err != nil {
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("guilds: age out: commit: %w", err)
	}
	return nil
}

// VerifyByLogs corroborates every currently-unverified guild_characters
// row against that guild's own uploaded logs: the character must be
// present (logNights: a rated or ranked row read by its key, so a boss
// fight in a report that is not private) in reports carrying this
// guild_id on at least two distinct report dates within a trailing
// 30-day window. fights.players is never read: it holds the log's GUIDs.
// A single forged export cannot satisfy this on its own — it needs two genuinely separate
// nights of the guild's own real combat-log uploads naming the same
// character.
//
// One fixed, always-on independence rule (fourth security review
// response, replacing two narrower rules - a 14-day/contested switch,
// then a two-account corroboration test reading recorded report-owner
// evidence - each found gameable by a squatter in turn): a report never
// counts toward a character's own two-distinct-dates threshold if that
// report is owned by the character's own account. This still lets a
// squatter verify sockpuppets into their own guild shell from their
// OWN uploaded reports, since nothing here stops account A's reports
// from verifying account B's character - accepted, because that grants
// the sockpuppets access to nothing but the squatter's own reports (see
// mayView/mayEdit), and a moderator's release un-verifies every
// claim-derived row regardless.
//
// Separately: while a guild's claim is contested, this verifies nobody
// in that guild at all - membership cannot shift under a moderator's
// feet while they are looking at an open dispute, the same guarantee
// approve's freeze already gives the roster (item 3, fourth security
// review response).
func (s *Store) VerifyByLogs(ctx context.Context) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: verify by logs: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	rows, err := tx.Query(ctx, `
		select gc.guild_id, gc.character_key, gc.user_id
		from guild_characters gc join guilds g on g.id = gc.guild_id
		where gc.verified_at is null and g.claim_contested_at is null
		order by gc.guild_id, gc.character_key`)
	if err != nil {
		return fmt.Errorf("guilds: verify by logs: candidates: %w", err)
	}
	candidates := map[int64]map[string]int64{}
	var guildOrder []int64
	for rows.Next() {
		var guildID, userID int64
		var key string
		if err := rows.Scan(&guildID, &key, &userID); err != nil {
			rows.Close()
			return fmt.Errorf("guilds: verify by logs: candidates: %w", err)
		}
		if candidates[guildID] == nil {
			candidates[guildID] = map[string]int64{}
			guildOrder = append(guildOrder, guildID)
		}
		candidates[guildID][key] = userID
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return fmt.Errorf("guilds: verify by logs: candidates: %w", err)
	}

	affected := map[guildUserPair]bool{}
	for _, guildID := range guildOrder {
		nights, err := logNights(ctx, tx, guildID, candidates[guildID])
		if err != nil {
			return fmt.Errorf("guilds: verify by logs: %w", err)
		}
		var verify []string
		for key, n := range nights {
			if n >= logNightsToVerify {
				verify = append(verify, key)
			}
		}
		if len(verify) == 0 {
			continue
		}
		rows, err := tx.Query(ctx, `
			update guild_characters set verified_at = now(), verified_by = 'logs'
			where guild_id = $1 and character_key = any($2) and verified_at is null
			  and exists (select 1 from guilds g where g.id = $1 and g.claim_contested_at is null)
			returning guild_id, user_id`, guildID, verify)
		if err != nil {
			return fmt.Errorf("guilds: verify by logs: update: %w", err)
		}
		pairs, err := scanGuildUserPairs(rows)
		if err != nil {
			return fmt.Errorf("guilds: verify by logs: %w", err)
		}
		for p := range pairs {
			affected[p] = true
		}
	}
	if err := recomputeEach(ctx, tx, affected); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("guilds: verify by logs: commit: %w", err)
	}
	return nil
}

// logNightsToVerify is how many distinct report dates of log evidence verify a character.
const logNightsToVerify = 2

// logNights counts, per candidate character key (key -> the account that holds it), the
// distinct report dates (reports.created_at's date) within the trailing 30 days on which a
// report of guildID that is evidence (presence.EvidenceReport: not private, not flagged as
// tampered or unverified) and not owned by the candidate's own account, has the character
// present. Presence is presence.In - a rated or ranked row read by the character's key -
// never fights.players, which holds the log's GUIDs. Those rows exist only for boss fights,
// so log evidence is an unflagged, non-private boss fight: the same evidence the Forever
// Score trusts, since every score input selects its reports through EvidenceReport too. Candidates with no such date are absent from the result. q is the pool or the
// caller's transaction.
func logNights(ctx context.Context, q character.Querier, guildID int64, candidates map[string]int64) (map[string]int, error) {
	out := map[string]int{}
	if len(candidates) == 0 {
		return out, nil
	}
	rows, err := q.Query(ctx, `
		select r.id, r.owner_id, to_char(r.created_at::date, 'YYYY-MM-DD') from reports r
		where r.guild_id = $1 and `+presence.EvidenceReport("r")+`
		  and r.created_at >= now() - interval '30 days'`, guildID)
	if err != nil {
		return nil, fmt.Errorf("log nights: %w", err)
	}
	type report struct {
		owner *int64
		date  string
	}
	reports := map[string]report{}
	var ids []string
	for rows.Next() {
		var id string
		var r report
		if err := rows.Scan(&id, &r.owner, &r.date); err != nil {
			rows.Close()
			return nil, fmt.Errorf("log nights: %w", err)
		}
		reports[id] = r
		ids = append(ids, id)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("log nights: %w", err)
	}

	keys := make([]string, 0, len(candidates))
	for key := range candidates {
		keys = append(keys, key)
	}
	present, err := presence.In(ctx, q, ids, keys)
	if err != nil {
		return nil, fmt.Errorf("log nights: %w", err)
	}
	dates := map[string]map[string]bool{}
	for _, p := range present {
		r := reports[p.ReportID]
		// A report never counts toward its own account's character (VerifyByLogs' rule).
		if r.owner != nil && *r.owner == candidates[p.Key] {
			continue
		}
		if dates[p.Key] == nil {
			dates[p.Key] = map[string]bool{}
		}
		dates[p.Key][r.date] = true
	}
	for key, d := range dates {
		out[key] = len(d)
	}
	return out, nil
}

type guildUserPair struct{ guildID, userID int64 }

// scanGuildUserPairs drains rows of (guild_id, user_id) into a
// deduplicated set, closing rows itself either way.
func scanGuildUserPairs(rows pgx.Rows) (map[guildUserPair]bool, error) {
	defer rows.Close()
	out := map[guildUserPair]bool{}
	for rows.Next() {
		var p guildUserPair
		if err := rows.Scan(&p.guildID, &p.userID); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		out[p] = true
	}
	return out, rows.Err()
}

// recomputeEach runs RecomputeMembership for every affected account, in tx.
func recomputeEach(ctx context.Context, tx pgx.Tx, affected map[guildUserPair]bool) error {
	for p := range affected {
		uid := p.userID
		if err := RecomputeMembership(ctx, tx, p.guildID, &uid); err != nil {
			return err
		}
	}
	return nil
}

// MembershipJob runs AgeOut and VerifyByLogs at startup and then on every
// tick, until ctx is cancelled — the same shape as db.PartitionJob.
type MembershipJob struct {
	Store *Store
	Log   *slog.Logger
	Every time.Duration
	// Now is the clock, so a test can drive the job from a fixed time.
	Now func() time.Time
}

func (j *MembershipJob) logger() *slog.Logger {
	if j.Log != nil {
		return j.Log
	}
	return slog.Default()
}

// Run performs the first sweep pass itself, returning its error so a
// startup can fail loudly on a database that cannot be swept at all;
// later passes are logged, because by then the service is serving
// traffic.
func (j *MembershipJob) Run(ctx context.Context) error {
	every, now := j.Every, j.Now
	if every <= 0 {
		every = SweepEvery
	}
	if now == nil {
		now = time.Now
	}
	log := j.logger()
	if err := j.sweepOnce(ctx, now()); err != nil {
		return err
	}
	go func() {
		t := time.NewTicker(every)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				if err := j.sweepOnce(ctx, now()); err != nil {
					log.Error("guilds", "op", "sweep", "err", err)
				}
			}
		}
	}()
	return nil
}

func (j *MembershipJob) sweepOnce(ctx context.Context, at time.Time) error {
	if err := j.Store.AgeOut(ctx, at); err != nil {
		return fmt.Errorf("guilds: age out: %w", err)
	}
	if err := j.Store.VerifyByLogs(ctx); err != nil {
		return fmt.Errorf("guilds: verify by logs: %w", err)
	}
	if err := j.Store.BackfillFactions(ctx); err != nil {
		return fmt.Errorf("guilds: backfill factions: %w", err)
	}
	return nil
}

// BackfillFactions runs RecomputeFaction for every guild that has never had it computed
// (faction_updated_at is null) — every guild that predates migration 0031, backfilled
// automatically on the sweep's own tick (SweepEvery, so within one deploy's worth of
// ticks) rather than a manual migration step. A guild's own membership changes already
// call RecomputeFaction the moment they happen (AfterGuildChange, ApproveCharacter,
// RemoveCharacter, Leave, AcceptInvite), so this only ever has guilds untouched since
// migration 0031 left to do.
func (s *Store) BackfillFactions(ctx context.Context) error {
	rows, err := s.Pool.Query(ctx, `select id from guilds where faction_updated_at is null`)
	if err != nil {
		return fmt.Errorf("guilds: backfill factions: list: %w", err)
	}
	ids, err := scanInt64s(rows)
	if err != nil {
		return fmt.Errorf("guilds: backfill factions: %w", err)
	}
	for _, id := range ids {
		if err := RecomputeFaction(ctx, s.Pool, id); err != nil {
			return fmt.Errorf("guilds: backfill factions: guild %d: %w", id, err)
		}
	}
	return nil
}

// scanInt64s drains rows of a single bigint column into a slice, closing rows itself
// either way — the same drain-before-use shape scanGuildUserPairs already uses, so a
// second query (RecomputeFaction, per id) never runs while this one is still open.
func scanInt64s(rows pgx.Rows) ([]int64, error) {
	defer rows.Close()
	var out []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		out = append(out, id)
	}
	return out, rows.Err()
}
