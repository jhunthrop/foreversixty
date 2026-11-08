// api/internal/guilds/raids_ping_claim.go
//
// The ping's per-member throttle (D16): a ping mentions only the raiders nobody mentioned for
// this raid in the last six hours (guild_raid_ping_mentions, migration 0090). One transaction
// per ping, serialised per raid, reads who is still blocked, stamps the rest and the raid's last
// ping (guild_raid_pings, the raid page's state); a post that does not land gives the stamps back.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// mentionClaim is one ping's claim: the targets it stamps, the ones still blocked, the earliest
// time a blocked one was mentioned (when the next one comes free, less pingEvery), and the
// raid's last ping before this one.
type mentionClaim struct {
	kept, skipped []pingTarget
	earliest      *time.Time
	prevPing      *time.Time
}

// claimMentions stamps at for every target not mentioned for raidID since at minus pingEvery,
// and the raid's last ping when it stamped any.
func (s *Store) claimMentions(ctx context.Context, raidID, by int64, targets []pingTarget, at time.Time) (mentionClaim, error) {
	var claim mentionClaim
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, `select pg_advisory_xact_lock(hashtextextended('raid-ping:' || $1::bigint::text, 0))`, raidID); err != nil {
			return err
		}
		blocked, err := recentMentions(ctx, tx, raidID, at.Add(-pingEvery))
		if err != nil {
			return err
		}
		for _, t := range targets {
			if when, ok := blocked[t.userID]; ok {
				claim.skipped = append(claim.skipped, t)
				if claim.earliest == nil || when.Before(*claim.earliest) {
					w := when
					claim.earliest = &w
				}
				continue
			}
			claim.kept = append(claim.kept, t)
		}
		if len(claim.kept) == 0 {
			return nil
		}
		return stampMentions(ctx, tx, raidID, by, at, &claim)
	})
	if err != nil {
		return mentionClaim{}, fmt.Errorf("guilds: claim ping: %w", err)
	}
	return claim, nil
}

// recentMentions is each account mentioned for the raid since since, with when.
func recentMentions(ctx context.Context, tx pgx.Tx, raidID int64, since time.Time) (map[int64]time.Time, error) {
	rows, err := tx.Query(ctx, `select user_id, mentioned_at from guild_raid_ping_mentions
		where raid_id = $1 and mentioned_at > $2`, raidID, since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := map[int64]time.Time{}
	for rows.Next() {
		var uid int64
		var at time.Time
		if err := rows.Scan(&uid, &at); err != nil {
			return nil, err
		}
		out[uid] = at
	}
	return out, rows.Err()
}

// stampMentions writes at for the claim's kept targets and as the raid's last ping, keeping the
// raid's previous ping for a release.
func stampMentions(ctx context.Context, tx pgx.Tx, raidID, by int64, at time.Time, claim *mentionClaim) error {
	ids := make([]int64, len(claim.kept))
	for i, t := range claim.kept {
		ids[i] = t.userID
	}
	if _, err := tx.Exec(ctx, `insert into guild_raid_ping_mentions (raid_id, user_id, mentioned_at)
		select $1, uid, $2 from unnest($3::bigint[]) as uid
		on conflict (raid_id, user_id) do update set mentioned_at = excluded.mentioned_at`, raidID, at, ids); err != nil {
		return err
	}
	err := tx.QueryRow(ctx, `select last_pinged_at from guild_raid_pings where raid_id = $1`, raidID).Scan(&claim.prevPing)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	_, err = tx.Exec(ctx, `insert into guild_raid_pings (raid_id, last_pinged_at, pinged_by, mentioned) values ($1, $2, $3, $4)
		on conflict (raid_id) do update set last_pinged_at = excluded.last_pinged_at,
		  pinged_by = excluded.pinged_by, mentioned = excluded.mentioned`, raidID, at, by, len(claim.kept))
	return err
}

// releaseMentions gives back a claim stamped at at: the kept targets' stamps go and the raid's
// last ping returns to the one before. A failure is logged: those raiders then wait six hours,
// never the other way.
func (s *Service) releaseMentions(ctx context.Context, raidID int64, at time.Time, claim mentionClaim) {
	ids := make([]int64, len(claim.kept))
	for i, t := range claim.kept {
		ids[i] = t.userID
	}
	err := pgx.BeginFunc(ctx, s.Store.Pool, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, `delete from guild_raid_ping_mentions
			where raid_id = $1 and mentioned_at = $2 and user_id = any($3)`, raidID, at, ids); err != nil {
			return err
		}
		if claim.prevPing == nil {
			_, err := tx.Exec(ctx, `delete from guild_raid_pings where raid_id = $1 and last_pinged_at = $2`, raidID, at)
			return err
		}
		_, err := tx.Exec(ctx, `update guild_raid_pings set last_pinged_at = $3 where raid_id = $1 and last_pinged_at = $2`,
			raidID, at, *claim.prevPing)
		return err
	})
	if err != nil {
		s.logger().Error("guilds", "op", "raid_ping_release", "raid_id", raidID, "err", err)
	}
}
