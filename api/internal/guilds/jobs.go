// api/internal/guilds/jobs.go
package guilds

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
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
// row against that guild's own uploaded logs: the character must appear
// in fights.players for reports carrying this guild_id on at least two
// distinct report dates within a trailing 30-day window. A single forged
// export cannot satisfy this on its own — it needs two genuinely separate
// nights of the guild's own real combat-log uploads naming the same
// character.
//
// Third security review response: while the guild's claim is young
// (established less than 14 days ago - the literal below must stay in
// sync with freezeThreshold in contest.go, the same duplication
// AutoConfirmClaimIfPending's own literal already carries against
// ClaimPendingTTL) or contested, only reports owned by neither the
// character's own account nor the current claim holder count toward
// the two-distinct-dates threshold - otherwise a squatter (the guild's
// only officer, so free to attach any report to it) could self-verify
// a throwaway account's character with their own uploaded reports
// during that risk window and manufacture, in advance, the very
// corroboration frozen() will later check for once the claim turns 14
// days old. An entirely unclaimed guild carries none of this risk (no
// claim exists yet for a contest to dispute), so it is exempt - this
// is also what lets a freshly-synced, not-yet-claimed guild's
// characters still verify normally from their own reports, exactly as
// before this response. Outside the risk window any report counts;
// either way, the report owner(s) that actually drove a fresh
// verification are recorded on log_evidence_owner_1/2, which frozen()'s
// corroboration check reads later - re-evaluated against whoever holds
// the claim at contest time, not against this sweep's snapshot of who
// held it.
func (s *Store) VerifyByLogs(ctx context.Context) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: verify by logs: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	rows, err := tx.Query(ctx, `
		with unverified as (
		  select gc.guild_id, gc.character_key, gc.user_id, g.claimed_by,
		         (g.claim_contested_at is not null
		          or (g.claimed_at is not null and now() - g.claimed_at < interval '14 days')
		         ) as restrict_independence
		  from guild_characters gc
		  join guilds g on g.id = gc.guild_id
		  where gc.verified_at is null
		),
		qualifying as (
		  select u.guild_id, u.character_key, u.user_id,
		         r.id as report_id, r.owner_id as report_owner, r.created_at::date as report_date
		  from unverified u
		  join reports r on r.guild_id = u.guild_id
		  join fights f on f.report_id = r.id and u.character_key = any(f.players)
		  where r.created_at >= now() - interval '30 days'
		    and (
		      not u.restrict_independence
		      or (r.owner_id is distinct from u.user_id
		          and (u.claimed_by is null or r.owner_id is distinct from u.claimed_by))
		    )
		),
		by_date as (
		  select distinct on (guild_id, character_key, report_date)
		    guild_id, character_key, user_id, report_date, report_owner
		  from qualifying
		  order by guild_id, character_key, report_date, report_id
		),
		ranked as (
		  select guild_id, character_key, user_id, report_owner,
		         row_number() over (partition by guild_id, character_key order by report_date) as rn,
		         count(*) over (partition by guild_id, character_key) as distinct_dates
		  from by_date
		),
		eligible as (
		  select guild_id, character_key, user_id,
		         max(report_owner) filter (where rn = 1) as owner1,
		         max(report_owner) filter (where rn = 2) as owner2
		  from ranked
		  where distinct_dates >= 2
		  group by guild_id, character_key, user_id
		)
		update guild_characters gc
		set verified_at = now(), verified_by = 'logs',
		    log_evidence_owner_1 = e.owner1, log_evidence_owner_2 = e.owner2
		from eligible e
		where gc.guild_id = e.guild_id and gc.character_key = e.character_key and gc.verified_at is null
		returning gc.guild_id, gc.user_id
	`)
	if err != nil {
		return fmt.Errorf("guilds: verify by logs: update: %w", err)
	}
	affected, err := scanGuildUserPairs(rows)
	if err != nil {
		return fmt.Errorf("guilds: verify by logs: %w", err)
	}
	if err := recomputeEach(ctx, tx, affected); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("guilds: verify by logs: commit: %w", err)
	}
	return nil
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
	return nil
}
