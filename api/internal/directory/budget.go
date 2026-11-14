// api/internal/directory/budget.go
package directory

import (
	"context"
	"sync/atomic"

	"golang.org/x/time/rate"
)

// Budget paces Blizzard calls to a rate per hour and counts them. One Budget is shared by
// every pass of a process, so passes that overlap never add up past the rate.
type Budget struct {
	limiter *rate.Limiter
	calls   atomic.Int64
}

// NewBudget allows perHour calls an hour, evenly spaced (a burst of one).
func NewBudget(perHour int) *Budget {
	return &Budget{limiter: rate.NewLimiter(rate.Limit(float64(perHour)/3600), 1)}
}

// Spend waits for the next call's slot, or for ctx to end.
func (b *Budget) Spend(ctx context.Context) error {
	if err := b.limiter.Wait(ctx); err != nil {
		return err
	}
	b.calls.Add(1)
	return nil
}

// Calls is how many calls the budget has paid for so far.
func (b *Budget) Calls() int64 { return b.calls.Load() }
