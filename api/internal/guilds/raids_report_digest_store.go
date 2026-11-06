// api/internal/guilds/raids_report_digest_store.go
//
// guild_report_pulls: the store of each fight's digest (raids_report_digest.go) and the Service
// step that fills it from the bucket. A fight is read once: at the match (ReportCompleted) and,
// for a report matched before the digest existed, on the first single-raid view. A failed read
// leaves the fight without a digest, and its pull shows its deaths with no cause until a later
// view reads it.
package guilds

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	// maxDigestSummaryBytes bounds one fight summary read back for a digest (the ingest's own
	// bundle bound, sims.MaxSummaryBytes).
	maxDigestSummaryBytes = 64 << 20
	// digestFillBudget bounds one fill, so a view of a long night never waits on the bucket for long.
	digestFillBudget = 8 * time.Second
)

// SummaryGetter reads one stored object; Objects (the Service's R2) satisfies it.
type SummaryGetter interface {
	Get(ctx context.Context, key string) (io.ReadCloser, error)
}

// undigestedFights is reportID's verified fights with no digest yet, in order.
func (s *Store) undigestedFights(ctx context.Context, reportID string) ([]int, error) {
	rows, err := s.Pool.Query(ctx, `
		select f.fight_index from fights f
		where f.report_id = $1 and f.verified
		  and not exists (select 1 from guild_report_pulls p where p.report_id = f.report_id and p.fight_index = f.fight_index)
		order by f.fight_index`, reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: undigested fights: %w", err)
	}
	defer rows.Close()
	var out []int
	for rows.Next() {
		var n int
		if err := rows.Scan(&n); err != nil {
			return nil, fmt.Errorf("guilds: undigested fights: %w", err)
		}
		out = append(out, n)
	}
	return out, rows.Err()
}

// writePullDigest stores one fight's digest; a fight already digested keeps its row.
func (s *Store) writePullDigest(ctx context.Context, reportID string, index int, d pullDigest) error {
	deaths, err := json.Marshal(d.Deaths)
	if err != nil {
		return fmt.Errorf("guilds: pull digest: %w", err)
	}
	var cons []byte
	if d.Consumables != nil {
		if cons, err = json.Marshal(d.Consumables); err != nil {
			return fmt.Errorf("guilds: pull digest: %w", err)
		}
	}
	if _, err := s.Pool.Exec(ctx, `
		insert into guild_report_pulls (report_id, fight_index, deaths, consumables) values ($1, $2, $3, $4)
		on conflict (report_id, fight_index) do nothing`, reportID, index, deaths, cons); err != nil {
		return fmt.Errorf("guilds: pull digest: %w", err)
	}
	return nil
}

// storedDigest is one fight's digest as read back; Read is false for a fight with no row.
type storedDigest struct {
	Read        bool
	Deaths      []digestDeath
	Consumables []digestConsumable // nil: no character sheet in the pull
}

// pullDigests reads reportID's stored digests by fight index.
func (s *Store) pullDigests(ctx context.Context, reportID string) (map[int]storedDigest, error) {
	rows, err := s.Pool.Query(ctx,
		`select fight_index, deaths, consumables from guild_report_pulls where report_id = $1`, reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: pull digests: %w", err)
	}
	defer rows.Close()
	out := map[int]storedDigest{}
	for rows.Next() {
		var n int
		var deaths, cons []byte
		if err := rows.Scan(&n, &deaths, &cons); err != nil {
			return nil, fmt.Errorf("guilds: pull digests: %w", err)
		}
		d := storedDigest{Read: true}
		if err := json.Unmarshal(deaths, &d.Deaths); err != nil {
			return nil, fmt.Errorf("guilds: pull digest deaths: %w", err)
		}
		if cons != nil {
			if err := json.Unmarshal(cons, &d.Consumables); err != nil {
				return nil, fmt.Errorf("guilds: pull digest consumables: %w", err)
			}
		}
		out[n] = d
	}
	return out, rows.Err()
}

// fillPullDigests digests reportID's fights that have none, from the bucket, within the fill
// budget, reading the summaries from R2 (the bucket the ingest wrote them to). Best effort: a failure is logged and the fight is tried again on a later fill.
func (s *Service) fillPullDigests(ctx context.Context, reportID string) {
	if s.R2 == nil {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, digestFillBudget)
	defer cancel()
	fights, err := s.Store.undigestedFights(ctx, reportID)
	if err != nil {
		s.logger().Warn("guilds", "op", "pull_digest", "report", reportID, "err", err)
		return
	}
	for _, n := range fights {
		sum, err := readFightSummary(ctx, s.R2, reportID, n)
		if err == nil {
			err = s.Store.writePullDigest(ctx, reportID, n, digestOf(sum))
		}
		if err != nil {
			s.logger().Warn("guilds", "op", "pull_digest", "report", reportID, "fight", n, "err", err)
			if ctx.Err() != nil {
				return
			}
		}
	}
}

// readFightSummary reads one fight's stored summary out of the bucket.
func readFightSummary(ctx context.Context, get SummaryGetter, reportID string, index int) (summary.Summary, error) {
	body, err := get.Get(ctx, store.Keys{ReportID: reportID}.FightSummary(index))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("guilds: read summary %s/%d: %w", reportID, index, err)
	}
	defer body.Close()
	b, err := io.ReadAll(io.LimitReader(body, maxDigestSummaryBytes))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("guilds: read summary %s/%d: %w", reportID, index, err)
	}
	var sum summary.Summary
	if err := json.Unmarshal(b, &sum); err != nil {
		return summary.Summary{}, fmt.Errorf("guilds: decode summary %s/%d: %w", reportID, index, err)
	}
	return sum, nil
}
