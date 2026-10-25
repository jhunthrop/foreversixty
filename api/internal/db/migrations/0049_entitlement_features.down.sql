-- End every scoped grant first: once the column is gone a scoped row reads as the whole
-- guild plan, which would turn every grandfathered guild into a full guild plan.
update entitlements set status = 'canceled', updated_at = now() where features is not null;
alter table entitlements drop constraint entitlements_features_scope;
alter table entitlements drop column features;
