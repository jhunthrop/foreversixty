// api/internal/rating/viewer_test.go
package rating

import "context"

// stubAccounts is a test double for Accounts (reports.GuildRanker). The visibility rule
// itself is tested once, in reports.TestMayView_TheRule.
type stubAccounts struct {
	rank string
	ok   bool
	err  error
}

func (s stubAccounts) GuildRank(_ context.Context, _, _ int64) (string, bool, error) {
	return s.rank, s.ok, s.err
}
