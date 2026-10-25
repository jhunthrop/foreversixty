-- A guild grant may cover officer tools alone (premium spec §4.7; design 2026-10-25 §4.1).
-- null = the whole plan (every row before this migration); '{officer_views}' = that feature only.
alter table entitlements add column features text[];
alter table entitlements add constraint entitlements_features_scope check (
  features is null
  or (plan = 'guild' and source = 'grant' and cardinality(features) > 0 and features <@ array['officer_views']::text[])
);
