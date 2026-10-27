// api/internal/billing/notices_store.go
package billing

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// Notice kinds, as billing_notices.kind stores them.
const (
	NoticeRenewal = "renewal"
	NoticeSize    = "size"
)

// Notice windows: how far ahead of the period end each notice is sent.
const (
	renewalNoticeLead = 7 * 24 * time.Hour
	sizeNoticeLead    = 14 * 24 * time.Hour
)

// DueNotice is one live Stripe row whose notice window is open. Guild is
// empty for a Premium row.
type DueNotice struct {
	EntitlementID   int64
	BillingUserID   int64
	Tier            string
	LookupKey       string
	UnitAmountCents int64
	PeriodEnd       time.Time
	Region          string
	Ruleset         string
	Guild           string
	Members         int
}

// dueNoticeQuery selects the live (active or trialing: a transfer trials to
// its first charge), not-canceling Stripe rows with a period
// end in (now, now+lead]; %s is a kind-specific extra filter.
const dueNoticeQuery = `
	select e.id, e.billing_user_id, coalesce(e.tier, ''), coalesce(e.price_lookup_key, ''),
	       coalesce(e.unit_amount_cents, 0), e.current_period_end,
	       coalesce(g.region, ''), coalesce(g.ruleset, ''), coalesce(g.name, ''),
	       coalesce((select count(distinct m.user_id) from guild_members m
	                 where m.guild_id = e.guild_id and m.verified_at is not null), 0)
	from entitlements e left join guilds g on g.id = e.guild_id
	where e.source = 'stripe' and e.status in ('active', 'trialing') and not e.cancel_at_period_end
	  and e.billing_user_id is not null
	  and e.current_period_end > $1 and e.current_period_end <= $1 + $2::interval
	  %s
	order by e.id`

// DueRenewals lists the live yearly rows renewing within the renewal lead.
func (s *Store) DueRenewals(ctx context.Context, now time.Time) ([]DueNotice, error) {
	rows, err := s.dueRows(ctx, now, renewalNoticeLead, "")
	if err != nil {
		return nil, err
	}
	yearly := rows[:0]
	for _, r := range rows {
		if price, ok := catalog.Lookup(r.LookupKey); ok && price.Interval == catalog.Yearly {
			yearly = append(yearly, r)
		}
	}
	return yearly, nil
}

// DueSizeChecks lists the Guild-tier rows renewing within the size lead whose
// guild has more verified members than the tier covers today.
func (s *Store) DueSizeChecks(ctx context.Context, now time.Time) ([]DueNotice, error) {
	rows, err := s.dueRows(ctx, now, sizeNoticeLead, "and e.tier = 'guild'")
	if err != nil {
		return nil, err
	}
	over := rows[:0]
	for _, r := range rows {
		if r.Members > catalog.GuildTierMemberLimit {
			over = append(over, r)
		}
	}
	return over, nil
}

func (s *Store) dueRows(ctx context.Context, now time.Time, lead time.Duration, extra string) ([]DueNotice, error) {
	rows, err := s.Pool.Query(ctx, fmt.Sprintf(dueNoticeQuery, extra), now, fmt.Sprintf("%d seconds", int64(lead.Seconds())))
	if err != nil {
		return nil, fmt.Errorf("billing: due notices: %w", err)
	}
	out, err := pgx.CollectRows(rows, func(r pgx.CollectableRow) (DueNotice, error) {
		var d DueNotice
		err := r.Scan(&d.EntitlementID, &d.BillingUserID, &d.Tier, &d.LookupKey, &d.UnitAmountCents,
			&d.PeriodEnd, &d.Region, &d.Ruleset, &d.Guild, &d.Members)
		return d, err
	})
	if err != nil {
		return nil, fmt.Errorf("billing: due notices: %w", err)
	}
	return out, nil
}

// ClaimNotice records that the notice is being sent. It reports false when
// the same notice for the same period was already claimed.
func (s *Store) ClaimNotice(ctx context.Context, kind string, entitlementID int64, periodEnd time.Time) (bool, error) {
	tag, err := s.Pool.Exec(ctx,
		`insert into billing_notices (kind, entitlement_id, period_end) values ($1, $2, $3)
		 on conflict do nothing`, kind, entitlementID, periodEnd)
	if err != nil {
		return false, fmt.Errorf("billing: claim %s notice for %d: %w", kind, entitlementID, err)
	}
	return tag.RowsAffected() == 1, nil
}

// ReleaseNotice undoes a claim whose email could not be sent, so the next
// run tries again.
func (s *Store) ReleaseNotice(ctx context.Context, kind string, entitlementID int64, periodEnd time.Time) error {
	if _, err := s.Pool.Exec(ctx,
		`delete from billing_notices where kind = $1 and entitlement_id = $2 and period_end = $3`,
		kind, entitlementID, periodEnd); err != nil {
		return fmt.Errorf("billing: release %s notice for %d: %w", kind, entitlementID, err)
	}
	return nil
}
