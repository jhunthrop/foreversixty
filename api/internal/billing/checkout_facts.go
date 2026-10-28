// api/internal/billing/checkout_facts.go
//
// The two per-guild checkout facts, one predicate each: checkGuildCheckout's refusal
// (PendingGuildCheckout), the trial record (GuildTrialUsed) and the subscribe step on GET
// /v1/me (GuildCheckoutFacts) all read them, so they cannot drift.
package billing

import (
	"context"
	"fmt"
)

// openCheckoutExists is whether the guild guildRef names has an unexpired pending Checkout
// Session. guildRef is a SQL placeholder or column, never input.
func openCheckoutExists(guildRef string) string {
	return `exists (select 1 from pending_checkouts p where p.guild_id = ` + guildRef + ` and p.expires_at > now())`
}

// trialUsedExists is whether the guild guildRef names has already had its one trial.
func trialUsedExists(guildRef string) string {
	return `exists (select 1 from guild_trials t where t.guild_id = ` + guildRef + `)`
}

// CheckoutFacts is one guild's checkout facts.
type CheckoutFacts struct {
	Open      bool // a Checkout Session is in flight (PendingGuildCheckout)
	TrialUsed bool // the guild has had its trial (GuildTrialUsed)
}

// GuildCheckoutFacts reads CheckoutFacts for every one of guildIDs in one query.
func (s *Store) GuildCheckoutFacts(ctx context.Context, guildIDs []int64) (map[int64]CheckoutFacts, error) {
	rows, err := s.Pool.Query(ctx, `select g.id, `+openCheckoutExists("g.id")+`, `+trialUsedExists("g.id")+`
		from unnest($1::bigint[]) as g(id)`, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("billing: guild checkout facts: %w", err)
	}
	defer rows.Close()
	out := make(map[int64]CheckoutFacts, len(guildIDs))
	for rows.Next() {
		var id int64
		var f CheckoutFacts
		if err := rows.Scan(&id, &f.Open, &f.TrialUsed); err != nil {
			return nil, fmt.Errorf("billing: guild checkout facts: %w", err)
		}
		out[id] = f
	}
	return out, rows.Err()
}
