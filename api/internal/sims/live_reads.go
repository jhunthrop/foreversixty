// api/internal/sims/live_reads.go
//
// The two reads the mobile live surfaces (lane F14) need: how many of a member's server runs are
// still going, for the widget snapshot, and whether a server run is the member's own, before a
// Live Activity may subscribe to it.
package sims

import (
	"context"
	"fmt"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// RunningServerRuns counts the member's server runs that are queued or running.
func (s *Store) RunningServerRuns(ctx context.Context, userID int64) (int, error) {
	var n int
	err := s.Pool.QueryRow(ctx,
		`select count(*) from sims
		 where user_id = $1 and lane = '`+simapi.LaneServer+`'
		   and state in ('`+StateQueued+`', '`+StateRunning+`')`, userID).Scan(&n)
	if err != nil {
		return 0, fmt.Errorf("sims: running server runs: %w", err)
	}
	return n, nil
}

// OwnsServerRun reports whether id is one of the member's server runs.
func (s *Store) OwnsServerRun(ctx context.Context, userID int64, id string) (bool, error) {
	var owns bool
	err := s.Pool.QueryRow(ctx,
		`select exists (select 1 from sims where id = $1 and user_id = $2 and lane = '`+simapi.LaneServer+`')`,
		id, userID).Scan(&owns)
	if err != nil {
		return false, fmt.Errorf("sims: owns server run: %w", err)
	}
	return owns, nil
}
