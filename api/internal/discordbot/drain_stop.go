package discordbot

import (
	"context"
	"time"
)

// defaultStopWithin is the drainer's share of a stop when Drainer.StopWithin is unset: the
// API's 9 s shutdown budget (cmd/api sets StopWithin from it).
const defaultStopWithin = 9 * time.Second

// landingKey carries a drain's landing context down to its detached writes.
type landingKey struct{}

// landing is ctx without its cancel, ended within after ctx ends: the row a stop lets land,
// its store writes and the release of the rows not reached all end inside the stop's budget,
// however their own budgets add up. end releases it.
func landing(ctx context.Context, within time.Duration) (land context.Context, end func()) {
	land, cancel := context.WithCancel(context.WithoutCancel(ctx))
	stopWatch := context.AfterFunc(ctx, func() {
		grace := time.NewTimer(within)
		defer grace.Stop()
		select {
		case <-grace.C:
			cancel()
		case <-land.Done():
		}
	})
	end = func() {
		stopWatch()
		cancel()
	}
	return context.WithValue(land, landingKey{}, land), end
}

// detached is a context for work that must outlive ctx's own cancel and deadline (a store
// write after Discord answered, a release): bounded by budget and, inside a drain, by the
// drain's landing context, so a stop still caps it.
func detached(ctx context.Context, budget time.Duration) (context.Context, context.CancelFunc) {
	base, ok := ctx.Value(landingKey{}).(context.Context)
	if !ok {
		base = context.WithoutCancel(ctx)
	}
	return context.WithTimeout(base, budget)
}
