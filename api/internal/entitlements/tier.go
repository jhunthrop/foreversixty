// api/internal/entitlements/tier.go
package entitlements

import (
	"context"
	"fmt"
)

// planTierSQL ranks what userID holds, by the access rules' own predicates: 2 for a
// verified officer or leader of a guild whose live, paid whole plan is the raid_guild tier,
// 1 for the guild tier, else 0; and whether they hold Premium. Only a Stripe-billed row
// raises the tier: grants (comp and the grandfathered officer_views scope) carry raid_guild
// with no revenue behind it, so their officers fall back to Premium or to no plan.
var planTierSQL = `
select coalesce((
  select max(case e.tier when '` + TierRaidGuild + `' then 2 when '` + TierGuild + `' then 1 else 0 end)
  from guild_members m
  join entitlements e on e.guild_id = m.guild_id and e.plan = 'guild' and e.source = 'stripe'
    and e.features is null and ` + activeClause + `
  where m.user_id = $1 and m.verified_at is not null and m.rank in ('officer', 'leader')
), 0), ` + standardEntitlementSQL("$1")

// PlanTier is the plan whose price bounds userID's spend: TierRaidGuild or TierGuild for a
// verified officer of a guild paying for that plan (the highest wins), else PlanPremium for a
// Premium holder (a personal plan or a held seat), else "" for no plan.
func (s *Store) PlanTier(ctx context.Context, userID int64) (string, error) {
	var rank int
	var premium bool
	if err := s.Pool.QueryRow(ctx, planTierSQL, userID).Scan(&rank, &premium); err != nil {
		return "", fmt.Errorf("entitlements: plan tier %d: %w", userID, err)
	}
	switch {
	case rank == 2:
		return TierRaidGuild, nil
	case rank == 1:
		return TierGuild, nil
	case premium:
		return PlanPremium, nil
	}
	return "", nil
}
