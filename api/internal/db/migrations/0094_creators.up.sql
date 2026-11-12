-- The creator program (design/specs/2026-11-08-creator-program.md): creators with a tracked link
-- and a code of their own (a free month at checkout), the accounts each brought (one creator per account, for good), the
-- revenue and refunds that earn a share, and the monthly payouts the owner approves. Nothing about a
-- player is stored here beyond the account id and the dates the rules read.

create table if not exists creators (
  id                       bigserial primary key,
  name                     text not null check (length(btrim(name)) between 1 and 60),
  channel_url              text not null check (channel_url ~ '^https://[^[:space:]]+$' and length(channel_url) <= 200),
  slug                     text not null unique references go_links (slug),
  code                     text not null unique check (code ~ '^[A-Z0-9]{4,20}$'),
  starts_on                date not null,
  cap_cents                integer not null default 50000 check (cap_cents between 0 and 1000000),
  -- The creator's own site account, bound by accepting the invite: never their own player.
  user_id                  bigint unique references users (id) on delete set null,
  stripe_account_id        text unique,
  stripe_status            text not null default 'none'
                             check (stripe_status in ('none', 'onboarding', 'needs_info', 'ready')),
  stripe_checked_at        timestamptz,
  -- Removed from the program: no new players, the creator Premium ended; the history stays.
  ended_at                 timestamptz,
  created_at               timestamptz not null default now(),
  updated_at               timestamptz not null default now()
);

-- One-time invites the owner sends; only the hash is kept.
create table if not exists creator_invites (
  token_hash bytea primary key,
  creator_id bigint not null references creators (id) on delete cascade,
  created_at timestamptz not null default now(),
  expires_at timestamptz not null,
  used_at    timestamptz
);
create index if not exists creator_invites_creator_idx on creator_invites (creator_id);

-- A month's statement once the owner approved it. pending: the transfer is being made (a retry
-- finds it by its transfer group); paid: Stripe took it; settled: nothing to transfer.
create table if not exists creator_payouts (
  id            bigserial primary key,
  creator_id    bigint not null references creators (id),
  month         date not null check (extract(day from month) = 1),
  status        text not null check (status in ('pending', 'paid', 'settled')),
  bounty_count  integer not null default 0,
  bounty_cents  bigint not null default 0,
  share_cents   bigint not null default 0,
  refund_cents  bigint not null default 0,
  capped_cents  bigint not null default 0,
  total_cents   bigint not null default 0 check (total_cents >= 0),
  transfer_id   text unique,
  approved_at   timestamptz not null default now(),
  paid_at       timestamptz,
  unique (creator_id, month)
);

-- The accounts a creator brought. source: their link (touched_at is the click) or their code
-- (touched_at is the redemption). second_day_at: the first signed-in request on a later UTC day
-- inside 14 days of sign-up. bounty_payout_id: the payout that paid the $2.
create table if not exists creator_players (
  user_id               bigint primary key references users (id) on delete cascade,
  creator_id            bigint not null references creators (id),
  source                text not null check (source in ('link', 'code')),
  touched_at            timestamptz not null,
  signed_up_at          timestamptz not null,
  second_day_at         timestamptz,
  first_paid_at         timestamptz,
  qualified_character   text unique,
  bounty_payout_id      bigint references creator_payouts (id),
  created_at            timestamptz not null default now()
);
create index if not exists creator_players_creator_idx on creator_players (creator_id);

-- One row per paid invoice of a creator's player inside the 12-month window. base_cents is the
-- payment's net after Stripe's fees and tax; share_cents is the creator's half of it.
create table if not exists creator_revenue (
  invoice_id  text primary key,
  creator_id  bigint not null references creators (id),
  user_id     bigint references users (id) on delete set null,
  charge_id   text not null unique,
  paid_at     timestamptz not null,
  amount_cents bigint not null,
  tax_cents   bigint not null default 0,
  fee_cents   bigint not null default 0,
  base_cents  bigint not null,
  share_cents bigint not null,
  payout_id   bigint references creator_payouts (id),
  created_at  timestamptz not null default now()
);
create index if not exists creator_revenue_creator_idx on creator_revenue (creator_id) where payout_id is null;

-- A refund of a recorded invoice: a negative line in the month it happened.
create table if not exists creator_refunds (
  refund_id    text primary key,
  invoice_id   text not null references creator_revenue (invoice_id),
  creator_id   bigint not null references creators (id),
  refunded_at  timestamptz not null,
  amount_cents bigint not null check (amount_cents >= 0),
  share_cents  bigint not null check (share_cents >= 0),
  payout_id    bigint references creator_payouts (id),
  created_at   timestamptz not null default now()
);
create index if not exists creator_refunds_creator_idx on creator_refunds (creator_id) where payout_id is null;

-- Applications to the program. The contact email is optional and shown to the owner only; it is
-- never logged. One open (pending) application per account.
create table if not exists creator_applications (
  id            bigserial primary key,
  user_id       bigint not null references users (id) on delete cascade,
  channel_url   text not null check (channel_url ~ '^https://[^[:space:]]+$' and length(channel_url) <= 200),
  platform      text not null check (platform in ('twitch', 'youtube')),
  audience      integer check (audience between 0 and 100000000),
  about         text not null check (length(btrim(about)) between 1 and 500),
  contact_email text check (length(contact_email) <= 254),
  status        text not null default 'pending' check (status in ('pending', 'approved', 'declined')),
  creator_id    bigint references creators (id),
  created_at    timestamptz not null default now(),
  decided_at    timestamptz
);
create unique index if not exists creator_applications_open_idx on creator_applications (user_id) where status = 'pending';

-- Each account's one Premium trial (a creator's code gives a free month), as guild_trials holds a
-- guild's: a second trial for the same account is refused at checkout.
create table if not exists premium_trials (
  user_id                bigint primary key references users (id) on delete cascade,
  stripe_subscription_id text not null,
  created_at             timestamptz not null default now()
);

-- A creator's own Premium (entitlements.GrantCreatorPremium): its own source, so it never replaces
-- a paid subscription or a hand grant, and ending it never touches one.
alter table entitlements drop constraint if exists entitlements_source_check;
alter table entitlements add constraint entitlements_source_check check (source in ('stripe', 'grant', 'trial', 'creator'));
