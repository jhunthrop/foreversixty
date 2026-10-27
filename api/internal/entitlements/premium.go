// api/internal/entitlements/premium.go
package entitlements

import (
	"context"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// Premium resolution: who holds Premium, from a personal plan or a held
// guild seat, and the supporter mark that reads the same rule.

// heldSeatClause (its CASE mirrors catalog.SeatsForTier, which a test walks
// tier by tier) matches a premium_seats row (alias s, joined to the
// holder's verified guild_members row m and the guild's whole-plan
// entitlements row e) that currently gives its holder Premium: held or
// accepted and started, inside the seat count of the guild's tier.
var heldSeatClause = fmt.Sprintf(`s.state in ('held', 'accepted')
	and (s.starts_at is null or s.starts_at <= now())
	and s.slot <= case e.tier when '%s' then %d when '%s' then %d else 0 end`,
	TierRaidGuild, catalog.SeatsRaidGuild, TierGuild, catalog.SeatsGuild)

// standardEntitlementSQL is the boolean expression "userExpr holds Premium":
// an active personal premium row, or a held seat on a guild plan the holder
// still verifiably belongs to. userExpr is a SQL expression for the user id.
func standardEntitlementSQL(userExpr string) string {
	return `(exists (
		  select 1 from entitlements e
		  where e.user_id = ` + userExpr + ` and e.plan = 'premium' and ` + activeClause + `
		) or exists (
		  select 1 from premium_seats s
		  join guild_members m on m.guild_id = s.guild_id and m.user_id = s.holder_user_id and m.verified_at is not null
		  join entitlements e on e.guild_id = s.guild_id and e.plan = 'guild' and e.features is null and ` + activeClause + `
		  where s.holder_user_id = ` + userExpr + ` and ` + heldSeatClause + `
		))`
}

// hasStandardEntitlement is every feature's rule except the per-guild ones:
// Premium from a personal plan or a held seat. Guild membership alone grants
// nothing. IsSupporter reuses this exact rule (the supporter mark and every
// non-officer feature share one threshold).
func (s *Store) hasStandardEntitlement(ctx context.Context, userID int64) (bool, error) {
	var ok bool
	if err := s.Pool.QueryRow(ctx, `select `+standardEntitlementSQL("$1"), userID).Scan(&ok); err != nil {
		return false, fmt.Errorf("entitlements: standard check %d: %w", userID, err)
	}
	return ok, nil
}

// SupportersAmong answers IsSupporter for many accounts in one query; the map
// holds an entry for every id asked about.
func (s *Store) SupportersAmong(ctx context.Context, userIDs []int64) (map[int64]bool, error) {
	out := make(map[int64]bool, len(userIDs))
	for _, id := range userIDs {
		out[id] = false
	}
	rows, err := s.Pool.Query(ctx, `
		select u.id from unnest($1::bigint[]) as u(id) where `+standardEntitlementSQL("u.id"), userIDs)
	if err != nil {
		return nil, fmt.Errorf("entitlements: supporters among: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("entitlements: supporters among: %w", err)
		}
		out[id] = true
	}
	return out, rows.Err()
}

// IsSupporter is a lighter question than Can (spec §1.3): does this
// account show the supporter mark. Signed-out never does.
func (s *Store) IsSupporter(ctx context.Context, userID int64) (bool, error) {
	if userID == 0 {
		return false, nil
	}
	return s.hasStandardEntitlement(ctx, userID)
}
