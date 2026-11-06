-- api/internal/db/migrations/0071_loot_awards_unmatched_by_award.up.sql
-- An award with no night is its own row. 0052's partial unique index on (guild_id, encounter_id,
-- item_id) where raid_id is null was the conflict target of every raid-less award, so the same
-- item handed out on two nights that never matched a raid or a log kept only the newer award.
-- From here an award with neither a raid nor a log is keyed by its id alone. An award filed under
-- a logged night (report_id, no planned raid) is that night's, like a raid's: one per drop per
-- log, re-awarding replaces it. Raid-attached awards keep loot_awards_drop_idx unchanged.
-- Existing rows already satisfy the new index: the old one was stricter (it ignored report_id).
create unique index if not exists loot_awards_drop_night_idx on loot_awards (guild_id, report_id, encounter_id, item_id)
  where raid_id is null and report_id is not null;
drop index if exists loot_awards_drop_unmatched_idx;
