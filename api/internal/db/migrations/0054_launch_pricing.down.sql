delete from entitlement_anomalies where kind in ('unexpected_transfer', 'unknown_price');
alter table entitlement_anomalies drop constraint entitlement_anomalies_kind_check;
alter table entitlement_anomalies add constraint entitlement_anomalies_kind_check
  check (kind in ('duplicate_subscription', 'orphan_subscription'));

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
