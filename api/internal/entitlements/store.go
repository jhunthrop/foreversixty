// api/internal/entitlements/store.go
package entitlements

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrNotFound is returned by Revoke when the named subject/plan has no
// entitlements row to revoke.
var ErrNotFound = errors.New("entitlements: not found")

// Store is every entitlements table read and write.
type Store struct{ Pool *pgxpool.Pool }

// activeStatusClause lists the entitlements.status values that keep every
// feature on (spec §1.3 rule 4): a status still inside Stripe's own
// trial/retry window counts; canceled, unpaid, incomplete, and
// incomplete_expired do not.
const activeStatusClause = "status in ('active', 'trialing', 'past_due')"

// Can answers whether userID may use feature right now, and why. See
// entitlements.go's Feature/Reason docs and spec §1.3 for the full rule.
func (s *Store) Can(ctx context.Context, userID int64, feature Feature) (bool, Reason, error) {
	if userID == 0 {
		return false, ReasonSignInRequired, nil
	}
	if feature == FeatureOfficerViews {
		return s.canOfficerViews(ctx, userID)
	}
	ok, err := s.hasStandardEntitlement(ctx, userID)
	if err != nil {
		return false, "", err
	}
	if !ok {
		return false, ReasonNoPlan, nil
	}
	return true, ReasonEntitled, nil
}

// hasStandardEntitlement is every feature's rule except officer_views
// (spec §1.3 rule 3): an active personal premium entitlement, or verified
// membership (any rank) in a guild that currently has the guild plan.
// IsSupporter reuses this exact query (spec: the supporter mark and every
// non-officer feature share one threshold).
func (s *Store) hasStandardEntitlement(ctx context.Context, userID int64) (bool, error) {
	var ok bool
	err := s.Pool.QueryRow(ctx, `
		select exists (
		  select 1 from entitlements
		  where user_id = $1 and plan = 'premium' and `+activeStatusClause+`
		) or exists (
		  select 1 from guild_members m
		  join entitlements e on e.guild_id = m.guild_id and e.plan = 'guild' and e.`+activeStatusClause+`
		  where m.user_id = $1 and m.verified_at is not null
		)`, userID).Scan(&ok)
	if err != nil {
		return false, fmt.Errorf("entitlements: standard check %d: %w", userID, err)
	}
	return ok, nil
}

// canOfficerViews is spec §1.3 rule 2: true only for a verified
// officer/leader of a guild that has the plan; ReasonNotOfficer when the
// caller is a verified member of such a guild but not at officer rank;
// ReasonNoPlan when no guild they verifiably belong to has the plan at all.
func (s *Store) canOfficerViews(ctx context.Context, userID int64) (bool, Reason, error) {
	var officer bool
	if err := s.Pool.QueryRow(ctx, `
		select exists (
		  select 1 from guild_members m
		  join entitlements e on e.guild_id = m.guild_id and e.plan = 'guild' and e.`+activeStatusClause+`
		  where m.user_id = $1 and m.verified_at is not null and m.rank in ('officer', 'leader')
		)`, userID).Scan(&officer); err != nil {
		return false, "", fmt.Errorf("entitlements: officer check %d: %w", userID, err)
	}
	if officer {
		return true, ReasonEntitled, nil
	}
	member, err := s.hasStandardEntitlement(ctx, userID)
	if err != nil {
		return false, "", err
	}
	// hasStandardEntitlement also counts a personal premium entitlement,
	// which never implies guild membership; re-check the guild half alone
	// so a premium-only account with no guild plan correctly gets
	// ReasonNoPlan, not ReasonNotOfficer.
	if member {
		var guildMember bool
		if err := s.Pool.QueryRow(ctx, `
			select exists (
			  select 1 from guild_members m
			  join entitlements e on e.guild_id = m.guild_id and e.plan = 'guild' and e.`+activeStatusClause+`
			  where m.user_id = $1 and m.verified_at is not null
			)`, userID).Scan(&guildMember); err != nil {
			return false, "", fmt.Errorf("entitlements: officer-guild check %d: %w", userID, err)
		}
		if guildMember {
			return false, ReasonNotOfficer, nil
		}
	}
	return false, ReasonNoPlan, nil
}

// IsSupporter is a lighter question than Can (spec §1.3): does this
// account show the supporter mark. Signed-out never does.
func (s *Store) IsSupporter(ctx context.Context, userID int64) (bool, error) {
	if userID == 0 {
		return false, nil
	}
	return s.hasStandardEntitlement(ctx, userID)
}

// Billing is a personal premium entitlement's state, as GET /v1/me's
// EntitlementsView.Billing needs it (spec §1.4).
type Billing struct {
	Plan              string
	Status            string
	CurrentPeriodEnd  *time.Time
	CancelAtPeriodEnd bool
}

// PersonalBilling reads userID's own premium entitlement row, regardless
// of status (a canceled row still shows on the account page until it
// ages out) — nil only when no such row exists at all.
func (s *Store) PersonalBilling(ctx context.Context, userID int64) (*Billing, error) {
	var b Billing
	err := s.Pool.QueryRow(ctx, `
		select plan, status, current_period_end, cancel_at_period_end
		from entitlements where user_id = $1 and plan = 'premium'`, userID).
		Scan(&b.Plan, &b.Status, &b.CurrentPeriodEnd, &b.CancelAtPeriodEnd)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("entitlements: personal billing %d: %w", userID, err)
	}
	return &b, nil
}

// GuildBilling is a guild-plan entitlement's state, currently active or
// not (spec §1.4's GuildBillingView, and billing's own §2.3 conflict
// check reuse this one method — see Task 2's Interfaces block).
type GuildBilling struct {
	Status            string
	CurrentPeriodEnd  *time.Time
	CancelAtPeriodEnd bool
	BillingUserID     *int64
}

// GuildBilling reads guildID's guild-plan row, nil when it has none, with
// no status filter: a canceled guild plan still needs to be visible (for
// example, to a settings page) exactly like a personal one does.
func (s *Store) GuildBilling(ctx context.Context, guildID int64) (*GuildBilling, error) {
	var g GuildBilling
	err := s.Pool.QueryRow(ctx, `
		select status, current_period_end, cancel_at_period_end, billing_user_id
		from entitlements where guild_id = $1 and plan = 'guild'`, guildID).
		Scan(&g.Status, &g.CurrentPeriodEnd, &g.CancelAtPeriodEnd, &g.BillingUserID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("entitlements: guild billing %d: %w", guildID, err)
	}
	return &g, nil
}

// StripeSubscriptionIDs lists every stripe_subscription_id whose row is
// still believed live (source = 'stripe', status not canceled) — spec
// §2.8's reconciliation job input, and part of the Entitlements
// interface billing.Service depends on (Task 4).
func (s *Store) StripeSubscriptionIDs(ctx context.Context) ([]string, error) {
	rows, err := s.Pool.Query(ctx,
		`select stripe_subscription_id from entitlements
		 where source = 'stripe' and status != 'canceled' and stripe_subscription_id is not null`)
	if err != nil {
		return nil, fmt.Errorf("entitlements: list stripe subscription ids: %w", err)
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("entitlements: list stripe subscription ids: %w", err)
		}
		ids = append(ids, id)
	}
	return ids, rows.Err()
}

// withTx runs fn in a transaction, committing on success and rolling back
// otherwise. Every write method below uses this so a row write and its
// audit row commit atomically or not at all.
func (s *Store) withTx(ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("entitlements: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := fn(tx); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("entitlements: commit: %w", err)
	}
	return nil
}

// recordAudit writes one entitlement_audit row inside tx (spec §1.1,
// §3's "audit logging of entitlement changes").
func recordAudit(ctx context.Context, tx pgx.Tx, entitlementID int64, userID, guildID *int64,
	plan string, oldStatus *string, newStatus, actor string) error {
	if _, err := tx.Exec(ctx, `
		insert into entitlement_audit (entitlement_id, user_id, guild_id, plan, old_status, new_status, actor)
		values ($1, $2, $3, $4, $5, $6, $7)`,
		entitlementID, userID, guildID, plan, oldStatus, newStatus, actor); err != nil {
		return fmt.Errorf("entitlements: audit: %w", err)
	}
	return nil
}

// conflictTarget returns the unique-constraint columns UpsertStripe/Grant
// upsert against for subj: exactly one of UserID/GuildID must be set.
func conflictTarget(subj Subject) (string, error) {
	if (subj.UserID == nil) == (subj.GuildID == nil) {
		return "", fmt.Errorf("entitlements: exactly one of UserID/GuildID must be set")
	}
	if subj.UserID != nil {
		return "user_id, plan", nil
	}
	return "guild_id, plan", nil
}

// existingStatus reads the current status of subj's row for plan, inside
// tx, for the audit's old_status — "" and isNew=true when no row exists
// yet.
func existingStatus(ctx context.Context, tx pgx.Tx, subj Subject, plan string) (status string, isNew bool, err error) {
	var row pgx.Row
	if subj.UserID != nil {
		row = tx.QueryRow(ctx, `select status from entitlements where user_id = $1 and plan = $2`, *subj.UserID, plan)
	} else {
		row = tx.QueryRow(ctx, `select status from entitlements where guild_id = $1 and plan = $2`, *subj.GuildID, plan)
	}
	err = row.Scan(&status)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", true, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("entitlements: read existing: %w", err)
	}
	return status, false, nil
}

// StripeUpsert is upsertFromSubscription's (billing package) sole way to
// write entitlements from Stripe state (spec §2.6 RULING 8): one row per
// (subject, plan), written by every subscription/invoice event with a
// freshly re-fetched Subscription.
type StripeUpsert struct {
	UserID               *int64
	GuildID              *int64
	Plan                 string
	Status               string
	CurrentPeriodEnd     *time.Time
	CancelAtPeriodEnd    bool
	StripeSubscriptionID string
	BillingUserID        int64
	Actor                string
}

// UpsertStripe is the only writer of entitlements from Stripe state.
func (s *Store) UpsertStripe(ctx context.Context, p StripeUpsert) error {
	subj := Subject{UserID: p.UserID, GuildID: p.GuildID}
	target, err := conflictTarget(subj)
	if err != nil {
		return err
	}
	return s.withTx(ctx, func(tx pgx.Tx) error {
		oldStatus, isNew, err := existingStatus(ctx, tx, subj, p.Plan)
		if err != nil {
			return err
		}
		var id int64
		if err := tx.QueryRow(ctx, `
			insert into entitlements (user_id, guild_id, plan, source, status, current_period_end,
				cancel_at_period_end, stripe_subscription_id, billing_user_id)
			values ($1, $2, $3, 'stripe', $4, $5, $6, $7, $8)
			on conflict (`+target+`) do update set
				source = 'stripe', status = excluded.status, current_period_end = excluded.current_period_end,
				cancel_at_period_end = excluded.cancel_at_period_end,
				stripe_subscription_id = excluded.stripe_subscription_id,
				billing_user_id = excluded.billing_user_id, updated_at = now()
			returning id`,
			p.UserID, p.GuildID, p.Plan, p.Status, p.CurrentPeriodEnd, p.CancelAtPeriodEnd,
			p.StripeSubscriptionID, p.BillingUserID).Scan(&id); err != nil {
			return fmt.Errorf("entitlements: upsert stripe: write: %w", err)
		}
		var oldPtr *string
		if !isNew {
			oldPtr = &oldStatus
		}
		return recordAudit(ctx, tx, id, p.UserID, p.GuildID, p.Plan, oldPtr, p.Status, p.Actor)
	})
}

// SetCanceled is customer.subscription.deleted's writer (spec §2.6's
// table): sets status = canceled and, when graceUntil is non-nil, the
// retention grace (spec §2.7). A subscription id with no matching row is
// not an error — it was never created through this integration, or was
// already handled by an earlier retry — the call still succeeds.
func (s *Store) SetCanceled(ctx context.Context, stripeSubscriptionID string, graceUntil *time.Time, actor string) error {
	return s.withTx(ctx, func(tx pgx.Tx) error {
		var id int64
		var userID, guildID *int64
		var plan, oldStatus string
		err := tx.QueryRow(ctx,
			`select id, user_id, guild_id, plan, status from entitlements where stripe_subscription_id = $1`,
			stripeSubscriptionID).Scan(&id, &userID, &guildID, &plan, &oldStatus)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		if err != nil {
			return fmt.Errorf("entitlements: set canceled: read: %w", err)
		}
		if _, err := tx.Exec(ctx,
			`update entitlements set status = 'canceled', grace_until = $2, updated_at = now() where id = $1`,
			id, graceUntil); err != nil {
			return fmt.Errorf("entitlements: set canceled: write: %w", err)
		}
		return recordAudit(ctx, tx, id, userID, guildID, plan, &oldStatus, "canceled", actor)
	})
}

// Subject names exactly one of a user's personal entitlement or a
// guild's, for Grant/Revoke — mirrors entitlements' own polymorphic
// (user_id, guild_id) shape (spec §1.1).
type Subject struct {
	UserID  *int64
	GuildID *int64
}

// Grant upserts a source = 'grant', status = 'active' row (spec §1.5's
// CLI). until is nil for an open-ended grant; grantedBy is nil when the
// operator running the CLI is not themselves a users row (Ruling G).
func (s *Store) Grant(ctx context.Context, subj Subject, plan string, until *time.Time,
	grantedBy *int64, note, actor string) error {
	target, err := conflictTarget(subj)
	if err != nil {
		return err
	}
	return s.withTx(ctx, func(tx pgx.Tx) error {
		oldStatus, isNew, err := existingStatus(ctx, tx, subj, plan)
		if err != nil {
			return err
		}
		var id int64
		if err := tx.QueryRow(ctx, `
			insert into entitlements (user_id, guild_id, plan, source, status, current_period_end, granted_by, grant_note)
			values ($1, $2, $3, 'grant', 'active', $4, $5, $6)
			on conflict (`+target+`) do update set
				source = 'grant', status = 'active', current_period_end = excluded.current_period_end,
				granted_by = excluded.granted_by, grant_note = excluded.grant_note, updated_at = now()
			returning id`,
			subj.UserID, subj.GuildID, plan, until, grantedBy, note).Scan(&id); err != nil {
			return fmt.Errorf("entitlements: grant: write: %w", err)
		}
		var oldPtr *string
		if !isNew {
			oldPtr = &oldStatus
		}
		return recordAudit(ctx, tx, id, subj.UserID, subj.GuildID, plan, oldPtr, "active", actor)
	})
}

// Revoke sets an existing entitlement's status to canceled. ErrNotFound
// when subj has no row for plan at all.
func (s *Store) Revoke(ctx context.Context, subj Subject, plan, actor string) error {
	return s.withTx(ctx, func(tx pgx.Tx) error {
		oldStatus, isNew, err := existingStatus(ctx, tx, subj, plan)
		if err != nil {
			return err
		}
		if isNew {
			return ErrNotFound
		}
		var id int64
		var target string
		if subj.UserID != nil {
			target = `user_id = $1 and plan = $2`
		} else {
			target = `guild_id = $1 and plan = $2`
		}
		var subjVal int64
		if subj.UserID != nil {
			subjVal = *subj.UserID
		} else {
			subjVal = *subj.GuildID
		}
		if err := tx.QueryRow(ctx,
			`update entitlements set status = 'canceled', updated_at = now() where `+target+` returning id`,
			subjVal, plan).Scan(&id); err != nil {
			return fmt.Errorf("entitlements: revoke: write: %w", err)
		}
		return recordAudit(ctx, tx, id, subj.UserID, subj.GuildID, plan, &oldStatus, "canceled", actor)
	})
}
