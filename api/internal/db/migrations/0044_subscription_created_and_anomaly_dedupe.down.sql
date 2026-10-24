drop index if exists entitlement_anomalies_kind_sub_unique;
alter table entitlements drop column if exists stripe_subscription_created;
