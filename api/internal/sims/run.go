package sims

import "context"

// Premiumer reads the premium flag. auth.Store satisfies it; the
// tests use a stub so the handler can be exercised without an
// accounts table.
type Premiumer interface {
	Premium(ctx context.Context, userID int64) (bool, error)
}
