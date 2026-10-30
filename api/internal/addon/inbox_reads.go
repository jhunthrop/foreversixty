package addon

// The rows the inbox v2 messages are built from (docs/contracts/addon-v2.md contract 3):
// one query per message type, reading stored rows only. Nothing here reads a fight summary
// or a sim's whole result: the companion asks every ten minutes.

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// LastRaidWindow is how far back a raid can be and still be the character's last raid.
const LastRaidWindow = 14 * 24 * time.Hour

// LastRaidFight is one rated boss fight of a character's last raid.
type LastRaidFight struct {
	CharacterKey, ReportID, Zone, Boss string
	// FightIndex is the fight's index in its report, counted from one.
	FightIndex int
	// Class and Spec are the log's own names ("Warrior", "Arms").
	Class, Spec       string
	FoughtAt          time.Time
	Overall           float64
	BiggestLoss       string
	LossSeconds       int
	Deaths            int
	ConsumablesMissed []string
	// Adherence is the stored rotation adherence; nil on a fight rated before migration 0037.
	Adherence *float64
}

// LastRaidFights reads, for each of keys, every rated boss fight of its last raid: the
// newest report that is evidence (presence.EvidenceReport: not private, not flagged) holding
// a rated fight of the character fought since since, oldest fight first. A card the engine
// marked insufficient is not a rated fight here: its overall is a zero it cannot support. The
// boss's name is the fight's own name, else the newest name any fight recorded for the
// encounter (the names rankings.Store.EncounterNames reads).
//
// A fight rated before migration 0036 has no stored facts (deaths is NULL) and is left out
// rather than shown as zeros; the report is still the one chosen, so a last raid whose
// every fight predates 0036 reads as no bosses and gets no message. A backfill re-rate
// fills the facts; whether to run one is the ops team's call.
func (s *Store) LastRaidFights(ctx context.Context, keys []string, since time.Time) ([]LastRaidFight, error) {
	rows, err := s.Pool.Query(ctx,
		`with mine as (
		   select rs.player_key, rs.report_id, rs.fight_index, rs.encounter_id, rs.fought_at, rs.overall,
		          coalesce(rs.class, '') as class, coalesce(rs.spec, '') as spec,
		          coalesce(rs.biggest_loss, '') as biggest_loss, coalesce(rs.loss_seconds, 0) as loss_seconds,
		          coalesce(rs.deaths, 0) as deaths, coalesce(rs.consumables_missed, '{}') as consumables_missed,
		          rs.adherence::float8 as adherence,
		          rs.deaths is not null as known
		     from rating_scores rs join reports r on r.id = rs.report_id
		    where rs.player_key = any($1) and `+presence.EvidenceReport("r")+`
		      and not coalesce(rs.insufficient, false) and rs.fought_at >= $2
		 ), newest as (
		   select distinct on (player_key) player_key, report_id
		     from mine order by player_key, fought_at desc, report_id desc
		 )
		 select m.player_key, m.report_id, m.fight_index, coalesce(r.zone, ''),
		        coalesce(nullif(f.name, ''),
		          (select f2.name from fights f2 where f2.encounter_id = m.encounter_id and f2.name <> ''
		            order by f2.start_ms desc limit 1), ''),
		        m.class, m.spec, m.fought_at, m.overall, m.biggest_loss, m.loss_seconds, m.deaths,
		        m.consumables_missed, m.adherence
		   from mine m
		   join newest n on n.player_key = m.player_key and n.report_id = m.report_id
		   join reports r on r.id = m.report_id
		   left join fights f on f.report_id = m.report_id and f.fight_index = m.fight_index
		  where m.known
		  order by m.player_key, m.fought_at, m.fight_index`, keys, since)
	if err != nil {
		return nil, fmt.Errorf("addon: read last raids: %w", err)
	}
	defer rows.Close()
	var out []LastRaidFight
	for rows.Next() {
		var f LastRaidFight
		if err := rows.Scan(&f.CharacterKey, &f.ReportID, &f.FightIndex, &f.Zone, &f.Boss, &f.Class, &f.Spec, &f.FoughtAt,
			&f.Overall, &f.BiggestLoss, &f.LossSeconds, &f.Deaths, &f.ConsumablesMissed, &f.Adherence); err != nil {
			return nil, fmt.Errorf("addon: read last raids: %w", err)
		}
		out = append(out, f)
	}
	return out, rows.Err()
}

// simResultRow is a character's newest finished in-game sim request's run.
type simResultRow struct {
	CharacterKey, Spec string
	DPS, Error         float64
	CreatedAt          time.Time
	// Combos is the stored result's ranked combinations: present only for a bulk run (Top
	// Gear and its kin); a plain run has none.
	Combos json.RawMessage
}

// simResults reads, for each of userID's characters that asked for a sim in game, the
// newest request whose run finished.
func (s *Store) simResults(ctx context.Context, userID int64) ([]simResultRow, error) {
	rows, err := s.Pool.Query(ctx,
		`select distinct on (a.character_key) a.character_key, s.spec, s.dps_mean, s.dps_error, s.created_at,
		        s.result -> 'combos'
		   from addon_sim_requests a join sims s on s.id = a.sim_id
		  where a.user_id = $1 and s.state = 'done'
		  order by a.character_key, a.created_at desc`, userID)
	if err != nil {
		return nil, fmt.Errorf("addon: read sim results: %w", err)
	}
	defer rows.Close()
	var out []simResultRow
	for rows.Next() {
		var r simResultRow
		var combos []byte
		if err := rows.Scan(&r.CharacterKey, &r.Spec, &r.DPS, &r.Error, &r.CreatedAt, &combos); err != nil {
			return nil, fmt.Errorf("addon: read sim results: %w", err)
		}
		r.Combos = combos
		out = append(out, r)
	}
	return out, rows.Err()
}

// ackRow is when the site last received an upload for one character (SeenAt, whether or not the
// export changed) and when the export last changed (ChangedAt), and what became of the
// character's newest in-game sim request ("" when it never asked) and when the site received
// it (nil when it never asked).
type ackRow struct {
	CharacterKey    string
	SeenAt          time.Time
	ChangedAt       time.Time
	SimRequestState string
	RequestedAt     *time.Time
}

// acks reads an ackRow for every one of userID's exports, but only when the account has a
// companion paired (a companion device that is not revoked; a signed-in app does not count): without one nothing would carry the ack
// into the game. A request's state is its run's once the run has ended - done, or failed for
// a run that errored - and the state recorded with the request (queued, or why no run started)
// until then.
func (s *Store) acks(ctx context.Context, userID int64) ([]ackRow, error) {
	rows, err := s.Pool.Query(ctx,
		`select e.character_key, e.seen_at, e.updated_at, coalesce(n.state, ''), n.created_at
		   from addon_exports e
		   left join lateral (
		     select case sm.state when 'done' then 'done' when 'error' then 'failed'
		                 else coalesce(a.state, 'queued') end as state,
		            a.created_at
		       from addon_sim_requests a left join sims sm on sm.id = a.sim_id
		      where a.character_key = e.character_key and a.user_id = e.user_id
		      order by a.created_at desc limit 1
		   ) n on true
		  where e.user_id = $1
		    and exists (select 1 from devices d where d.user_id = $1 and d.revoked_at is null and d.kind = 'companion')
		  order by e.character_key`, userID)
	if err != nil {
		return nil, fmt.Errorf("addon: read acks: %w", err)
	}
	defer rows.Close()
	var out []ackRow
	for rows.Next() {
		var r ackRow
		if err := rows.Scan(&r.CharacterKey, &r.SeenAt, &r.ChangedAt, &r.SimRequestState, &r.RequestedAt); err != nil {
			return nil, fmt.Errorf("addon: read acks: %w", err)
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// inGameWeightsRow is a character's newest finished in-game weights run.
type inGameWeightsRow struct {
	CharacterKey, Spec string
	// CreatedAt is when the run was queued (sims.created_at; the table has no finish time).
	CreatedAt time.Time
	// ExportSeenAt is the export the request was made on, nil when it was not recorded. A
	// request recorded before migration 0038 queued a plain run (no weights), so it is never
	// read here; in practice nil means the character had no stored export when it asked,
	// which the outbox path, storing the export first, does not produce.
	ExportSeenAt *time.Time
	Weights      []simapi.StatWeight
}

// inGameWeights reads, for each of userID's characters that asked for a sim in game, the
// newest request whose stat-weights run finished with weights (a done run with an empty list
// is passed over for an older one, so the character still gets its last real weights): its
// spec, weights, queue time and export.
func (s *Store) inGameWeights(ctx context.Context, userID int64) ([]inGameWeightsRow, error) {
	rows, err := s.Pool.Query(ctx,
		`select distinct on (a.character_key) a.character_key, sm.spec, sm.created_at, a.export_seen_at,
		        coalesce(sm.result -> 'weights', '[]')
		   from addon_sim_requests a join sims sm on sm.id = a.sim_id
		  where a.user_id = $1 and sm.state = 'done' and sm.kind = 'weights'
		    and case when jsonb_typeof(sm.result -> 'weights') = 'array'
		             then jsonb_array_length(sm.result -> 'weights') else 0 end > 0
		  order by a.character_key, a.created_at desc`, userID)
	if err != nil {
		return nil, fmt.Errorf("addon: read in-game weights: %w", err)
	}
	defer rows.Close()
	var out []inGameWeightsRow
	for rows.Next() {
		var r inGameWeightsRow
		var raw []byte
		if err := rows.Scan(&r.CharacterKey, &r.Spec, &r.CreatedAt, &r.ExportSeenAt, &raw); err != nil {
			return nil, fmt.Errorf("addon: read in-game weights: %w", err)
		}
		if err := json.Unmarshal(raw, &r.Weights); err != nil {
			return nil, fmt.Errorf("addon: read in-game weights of %s: %w", r.CharacterKey, err)
		}
		out = append(out, r)
	}
	return out, rows.Err()
}
