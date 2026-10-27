-- Launch pricing (design 2026-10-27): the entitlement row carries the price facts Stripe
-- sent, a guild tier, and the tables for seats, trials, plan transfers, comebacks and
-- renewal notices. Migration 0053 belongs to the discord-bot lane; this is 0054.

alter table entitlements
  add column tier text,
  add column price_lookup_key text,
  add column unit_amount_cents integer,
  add column trial_end timestamptz;

-- Every guild row that exists today (whole-plan grants and the grandfathering scoped
-- grants) keeps the loot council and readiness, which the raid tier carries.
update entitlements set tier = 'raid_guild' where plan = 'guild';

alter table entitlements add constraint entitlements_tier check (
  (plan = 'premium' and tier is null)
  or (plan = 'guild' and tier is not null and tier in ('guild', 'raid_guild'))
);

create table premium_seats (
  guild_id                  bigint not null references guilds(id) on delete cascade,
  slot                      smallint not null check (slot between 1 and 3),
  state                     text not null check (state in ('empty', 'held', 'offered', 'accepted')),
  holder_user_id            bigint references users(id) on delete cascade,
  given_by                  bigint references users(id) on delete set null,
  given_at                  timestamptz,
  offered_at                timestamptz,
  starts_at                 timestamptz,
  movable_at                timestamptz,
  holder_gave_up_lookup_key text,
  returned_name             text,
  returned_at               timestamptz,
  primary key (guild_id, slot),
  check ((state = 'empty') = (holder_user_id is null)),
  check (state <> 'accepted' or starts_at is not null),
  check (state <> 'held' or (given_at is not null and movable_at is not null))
);
create unique index premium_seats_holder_idx
  on premium_seats (guild_id, holder_user_id) where holder_user_id is not null;

-- Cascades with the guild: a re-created guild may trial again.
create table guild_trials (
  guild_id               bigint primary key references guilds(id) on delete cascade,
  stripe_subscription_id text not null,
  started_at             timestamptz not null default now()
);

create table plan_transfers (
  guild_id             bigint primary key references guilds(id) on delete cascade,
  from_user_id         bigint not null references users(id) on delete cascade,
  to_user_id           bigint not null references users(id) on delete cascade,
  from_subscription_id text not null,
  state                text not null check (state in ('offered', 'declined', 'accepted')),
  offered_at           timestamptz not null default now(),
  decided_at           timestamptz
);

create table seat_comebacks (
  id               bigserial primary key,
  user_id          bigint not null references users(id) on delete cascade,
  guild_name       text not null,
  ended_at         timestamptz not null,
  price_lookup_key text not null,
  expires_at       timestamptz not null,
  used_at          timestamptz
);

create table billing_notices (
  kind           text not null check (kind in ('renewal', 'size')),
  entitlement_id bigint not null references entitlements(id) on delete cascade,
  period_end     timestamptz not null,
  sent_at        timestamptz not null default now(),
  primary key (kind, entitlement_id, period_end)
);

create unique index seat_comebacks_once_idx on seat_comebacks (user_id, guild_name, ended_at);
create index seat_comebacks_user_idx on seat_comebacks (user_id);

-- Billing anomalies the launch webhook raises: a transfer with no offer behind
-- it, and a subscription whose price is not in the catalog for its plan.
alter table entitlement_anomalies drop constraint entitlement_anomalies_kind_check;
alter table entitlement_anomalies add constraint entitlement_anomalies_kind_check
  check (kind in ('duplicate_subscription', 'orphan_subscription', 'unexpected_transfer', 'unknown_price'));
