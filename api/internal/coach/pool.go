package coach

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// DefaultPoolTTL is how long a bracket's pool is served from memory.
const DefaultPoolTTL = time.Hour

// PoolKey is one comparison bracket: a boss at a difficulty and size, a spec, and the
// content phase, so a player is never compared across phases.
type PoolKey struct {
	EncounterID, Difficulty int64
	Size                    int
	Spec, Phase             string
}

// poolRowsSQL lists the bracket's players, best kill first, one row each, on reports that
// are still evidence. Only the first $6 rows carry their profile; the rest are counted.
var poolRowsSQL = `
with best as (
  select distinct on (cp.player_key) cp.player_key, cp.metric, cp.profile
  from coach_profiles cp join reports r on r.id = cp.report_id
  where cp.encounter_id = $1 and cp.difficulty = $2 and cp.size = $3 and cp.spec = $4
    and cp.phase = $5 and cp.kill and ` + presence.EvidenceReport("r") + `
  order by cp.player_key, cp.metric desc
), ranked as (
  select player_key, profile, row_number() over (order by metric desc, player_key) as rn from best
)
select player_key, case when rn <= $6 then profile end from ranked order by rn`

// Pool reads the comparison pool: one best kill per player in the bracket, on reports that
// are still evidence. The bracket is read once per TTL and the reviewed player is left
// out in memory, so nobody in the pool is ever named or linked.
type Pool struct {
	DB  *pgxpool.Pool
	TTL time.Duration    // DefaultPoolTTL when zero
	Now func() time.Time // time.Now when nil

	mu    sync.Mutex
	cache map[PoolKey]poolEntry
}

// poolEntry is a bracket as read: every player's key in rank order, and the profiles of
// the first PoolTopN+1 (the extra one replaces the reviewed player when they are in the top).
type poolEntry struct {
	readAt   time.Time
	players  []string
	profiles []Profile
}

// Reference returns the distinct-player count of the bracket without exclude and, when
// that count is at least PoolMinPlayers, the median profile of the top PoolTopN by metric.
func (p *Pool) Reference(ctx context.Context, k PoolKey, exclude string) (int, *PoolRef, error) {
	e, err := p.entry(ctx, k)
	if err != nil {
		return 0, nil, err
	}
	count, top := e.without(exclude)
	if count < PoolMinPlayers {
		return count, nil, nil
	}
	return count, &PoolRef{N: len(top), Median: MedianProfile(top)}, nil
}

// without is the player count and the top profiles once exclude is removed.
func (e poolEntry) without(exclude string) (int, []Profile) {
	count := len(e.players)
	profiles := e.profiles
	if i := slices.Index(e.players, exclude); i >= 0 {
		count--
		if i < len(profiles) {
			profiles = slices.Delete(slices.Clone(profiles), i, i+1)
		}
	}
	return count, profiles[:min(len(profiles), PoolTopN)]
}

func (p *Pool) entry(ctx context.Context, k PoolKey) (poolEntry, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	now := p.now()
	if e, ok := p.cache[k]; ok && now.Sub(e.readAt) < p.ttl() {
		return e, nil
	}
	e, err := p.read(ctx, k)
	if err != nil {
		return poolEntry{}, err
	}
	e.readAt = now
	if p.cache == nil {
		p.cache = map[PoolKey]poolEntry{}
	}
	p.cache[k] = e
	return e, nil
}

func (p *Pool) read(ctx context.Context, k PoolKey) (poolEntry, error) {
	rows, err := p.DB.Query(ctx, poolRowsSQL, k.EncounterID, k.Difficulty, k.Size, k.Spec, k.Phase, PoolTopN+1)
	if err != nil {
		return poolEntry{}, fmt.Errorf("coach pool %+v: %w", k, err)
	}
	defer rows.Close()
	var e poolEntry
	for rows.Next() {
		var player string
		var raw []byte
		if err := rows.Scan(&player, &raw); err != nil {
			return poolEntry{}, fmt.Errorf("coach pool %+v: scan: %w", k, err)
		}
		e.players = append(e.players, player)
		if raw == nil {
			continue
		}
		var prof Profile
		if err := json.Unmarshal(raw, &prof); err != nil {
			return poolEntry{}, fmt.Errorf("coach pool %+v: profile of %s: %w", k, player, err)
		}
		e.profiles = append(e.profiles, prof)
	}
	if err := rows.Err(); err != nil {
		return poolEntry{}, fmt.Errorf("coach pool %+v: %w", k, err)
	}
	return e, nil
}

func (p *Pool) ttl() time.Duration {
	if p.TTL > 0 {
		return p.TTL
	}
	return DefaultPoolTTL
}

func (p *Pool) now() time.Time {
	if p.Now != nil {
		return p.Now()
	}
	return time.Now()
}
