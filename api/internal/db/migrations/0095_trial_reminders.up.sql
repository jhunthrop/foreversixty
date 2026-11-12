-- Free trials told plainly (design/specs/2026-11-08-premium-trial.md): the reminder three days
-- before a trial's first charge and the receipt after it are billing notices, once each, and the
-- reminder's one-click cancel link is a token stored only as its SHA-256.

alter table billing_notices drop constraint if exists billing_notices_kind_check;
alter table billing_notices add constraint billing_notices_kind_check
  check (kind in ('renewal', 'size', 'trial_ending', 'trial_receipt'));

create table billing_cancel_tokens (
  token_hash             bytea primary key,
  entitlement_id         bigint not null references entitlements(id) on delete cascade,
  stripe_subscription_id text not null,
  purpose                text not null check (purpose in ('cancel_trial')),
  expires_at             timestamptz not null,
  used_at                timestamptz,
  created_at             timestamptz not null default now()
);
create index billing_cancel_tokens_entitlement_idx on billing_cancel_tokens (entitlement_id);
