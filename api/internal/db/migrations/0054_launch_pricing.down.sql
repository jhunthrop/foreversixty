drop table billing_notices;
drop table seat_comebacks;
drop table plan_transfers;
drop table guild_trials;
drop table premium_seats;

alter table entitlements drop constraint entitlements_tier;
alter table entitlements
  drop column trial_end,
  drop column unit_amount_cents,
  drop column price_lookup_key,
  drop column tier;
