-- The Stripe subscription's own created timestamp, so a "Take over billing" transfer
-- (spec §2.9 RULING 10) can only replace a strictly older subscription (issue #3 H2).
alter table entitlements add column stripe_subscription_created timestamptz;
-- One anomaly row per (kind, subscription): a webhook retry after a failed Stripe cancel,
-- or the nightly orphan pass, no longer re-records the same subscription (issue #3 L1).
-- Existing duplicates go first, keeping the lowest id, or the unique index cannot build.
delete from entitlement_anomalies a using entitlement_anomalies b
  where a.kind = b.kind and a.stripe_subscription_id = b.stripe_subscription_id and a.id > b.id;
create unique index entitlement_anomalies_kind_sub_unique
  on entitlement_anomalies (kind, stripe_subscription_id);
