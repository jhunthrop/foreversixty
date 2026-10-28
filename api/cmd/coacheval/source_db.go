package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// pickFightsSQL lists stored kills on reports that are still evidence, in a stable
// pseudo-random order so the sample mixes encounters, sizes and specs.
var pickFightsSQL = `
select s.report_id, s.fight_index, s.size, s.fought_at
from (select distinct on (report_id, fight_index) report_id, fight_index, size, fought_at
      from rating_scores where kill and encounter_id > 0 and spec is not null
      order by report_id, fight_index, fought_at) s
join reports r on r.id = s.report_id
where ` + presence.EvidenceReport("r") + `
order by md5(s.report_id || ':' || s.fight_index::text)
limit $1`

const playerKeysSQL = `select player_name, player_key from rating_scores where report_id = $1 and fight_index = $2`

// readOnly makes the server refuse any write in the transaction.
var readOnly = pgx.TxOptions{AccessMode: pgx.ReadOnly}

// dbSource reads candidate fights from the database in read-only transactions and their
// summaries from the bucket. It issues no write of any kind.
type dbSource struct {
	pool *pgxpool.Pool
	get  sims.Getter
}

func (d dbSource) Fights(ctx context.Context, n int) ([]Fight, error) {
	fights, err := d.candidates(ctx, n)
	if err != nil {
		return nil, err
	}
	for i := range fights {
		if err := d.load(ctx, &fights[i]); err != nil {
			return nil, err
		}
	}
	return fights, nil
}

func (d dbSource) candidates(ctx context.Context, n int) ([]Fight, error) {
	tx, err := d.pool.BeginTx(ctx, readOnly)
	if err != nil {
		return nil, fmt.Errorf("begin read-only transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	rows, err := tx.Query(ctx, pickFightsSQL, n)
	if err != nil {
		return nil, fmt.Errorf("pick fights: %w", err)
	}
	defer rows.Close()
	var out []Fight
	for rows.Next() {
		var f Fight
		if err := rows.Scan(&f.ReportID, &f.FightIndex, &f.Size, &f.FoughtAt); err != nil {
			return nil, fmt.Errorf("scan fight: %w", err)
		}
		out = append(out, f)
	}
	return out, rows.Err()
}

func (d dbSource) load(ctx context.Context, f *Fight) error {
	s, err := d.summary(ctx, f.ReportID, f.FightIndex)
	if err != nil {
		return err
	}
	f.Summary = s
	f.PlayerKeys, err = d.playerKeys(ctx, f.ReportID, f.FightIndex)
	return err
}

func (d dbSource) playerKeys(ctx context.Context, reportID string, fightIndex int) (map[string]string, error) {
	tx, err := d.pool.BeginTx(ctx, readOnly)
	if err != nil {
		return nil, fmt.Errorf("begin read-only transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	rows, err := tx.Query(ctx, playerKeysSQL, reportID, fightIndex)
	if err != nil {
		return nil, fmt.Errorf("player keys: %w", err)
	}
	defer rows.Close()
	keys := map[string]string{}
	for rows.Next() {
		var name, key string
		if err := rows.Scan(&name, &key); err != nil {
			return nil, fmt.Errorf("scan player key: %w", err)
		}
		keys[name] = key
	}
	return keys, rows.Err()
}

func (d dbSource) summary(ctx context.Context, reportID string, fightIndex int) (summary.Summary, error) {
	body, err := d.get.Get(ctx, store.Keys{ReportID: reportID}.FightSummary(fightIndex))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("read summary %d: %w", fightIndex, err)
	}
	defer body.Close()
	raw, err := io.ReadAll(io.LimitReader(body, sims.MaxSummaryBytes))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("read summary %d: %w", fightIndex, err)
	}
	var s summary.Summary
	if err := json.Unmarshal(raw, &s); err != nil {
		return summary.Summary{}, fmt.Errorf("decode summary %d: %w", fightIndex, err)
	}
	return s, nil
}
