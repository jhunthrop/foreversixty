-- not valid keeps any grant_conflict rows already recorded instead of deleting them.
alter table entitlement_anomalies drop constraint entitlement_anomalies_kind_check;
alter table entitlement_anomalies add constraint entitlement_anomalies_kind_check
  check (kind in ('duplicate_subscription', 'orphan_subscription', 'unexpected_transfer', 'unknown_price'))
  not valid;
