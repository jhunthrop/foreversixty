-- One row per (subject, plan): a user's premium entitlement, or a guild's guild-plan
-- entitlement. Polymorphic by two nullable FKs rather than a single untyped subject_id, so
-- referential integrity is real and a query never has to trust an unenforced subject_type
-- string. See spec 2026-09-21-entitlements-and-payments-design.md §1.1, RULING 1.
create table entitlements (
  id                      bigserial primary key,
  user_id                 bigint references users(id) on delete cascade,
  guild_id                bigint references guilds(id) on delete cascade,
  plan                    text not null check (plan in ('premium', 'guild')),
  source                  text not null check (source in ('stripe', 'grant', 'trial')),
  status                  text not null check (status in
    ('active', 'trialing', 'past_due', 'canceled', 'incomplete', 'incomplete_expired', 'unpaid')),
  current_period_end      timestamptz,
  cancel_at_period_end    boolean not null default false,
  -- Retention grace only (proposal's "2 years, then 90 days after a 30-day grace"), never a
  -- feature-access grace: feature access follows status/current_period_end alone (§1.3 rule 4).
  grace_until             timestamptz,
  stripe_subscription_id  text unique,
  -- Set only for source = 'stripe': the user whose Stripe Customer backs this subscription.
  -- For plan = 'premium' this is always user_id itself; for plan = 'guild' it is the officer
  -- who bought it, which can differ from every guild_members row it benefits (§2.9).
  billing_user_id         bigint references users(id) on delete set null,
  granted_by              bigint references users(id) on delete set null,
  grant_note              text,
  created_at              timestamptz not null default now(),
  updated_at              timestamptz not null default now(),
  constraint entitlements_one_subject check (
    (user_id is not null and guild_id is null and plan = 'premium') or
    (guild_id is not null and user_id is null and plan = 'guild')
  ),
  -- NULLs are distinct in a unique constraint: a guild-plan row (user_id null) never collides
  -- with a premium row here, and vice versa.
  constraint entitlements_user_plan_unique unique (user_id, plan),
  constraint entitlements_guild_plan_unique unique (guild_id, plan)
);
create index entitlements_guild_idx on entitlements (guild_id) where guild_id is not null;
create index entitlements_billing_user_idx on entitlements (billing_user_id)
  where billing_user_id is not null;

-- One Stripe Customer per site account, created lazily on first checkout (§2.4).
create table stripe_customers (
  user_id             bigint primary key references users(id) on delete cascade,
  stripe_customer_id  text not null unique,
  created_at          timestamptz not null default now()
);

-- The webhook's idempotency log and its audit trail in one: every event Stripe ever sent us,
-- keyed by Stripe's own event id, so a redelivery is detected by a unique-key conflict rather
-- than re-applied (§2.6).
create table stripe_events (
  id            text primary key,
  type          text not null,
  payload       jsonb not null,
  received_at   timestamptz not null default now(),
  processed_at  timestamptz
);

-- Our own interpretation of every entitlement change, independent of source (§3's "audit
-- logging of entitlement changes").
create table entitlement_audit (
  id              bigserial primary key,
  entitlement_id  bigint references entitlements(id) on delete set null,
  user_id         bigint references users(id) on delete set null,
  guild_id        bigint references guilds(id) on delete set null,
  plan            text not null,
  old_status      text,
  new_status      text not null,
  -- 'stripe_webhook:<event_id>' | 'stripe_reconcile' | 'cli_grant:<operator note>' |
  -- 'cli_revoke:<operator note>'.
  actor           text not null,
  created_at      timestamptz not null default now()
);
create index entitlement_audit_entitlement_idx on entitlement_audit (entitlement_id, created_at desc);

-- Backfill (spec §1.2 step 1): one entitlements row per account that has premium = true today.
-- users.premium is left in place by this migration — the Go code deployed alongside it stops
-- reading the column, but the column itself is dropped only by a later 0021, after this lane's
-- code has been confirmed healthy in production (not built in this lane).
insert into entitlements (user_id, plan, source, status, granted_by, grant_note)
select id, 'premium', 'grant', 'active', null, 'migrated from users.premium'
from users where premium = true
on conflict (user_id, plan) do nothing;
