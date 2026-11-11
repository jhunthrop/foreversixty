-- A live Stripe subscription billing beside a live grant that holds the row is recorded for a
-- human to settle, never written over the grant (issue #3 M5).
alter table entitlement_anomalies drop constraint entitlement_anomalies_kind_check;
alter table entitlement_anomalies add constraint entitlement_anomalies_kind_check
  check (kind in ('duplicate_subscription', 'orphan_subscription', 'unexpected_transfer', 'unknown_price',
                  'grant_conflict'));
