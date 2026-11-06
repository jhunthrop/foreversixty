-- api/internal/db/migrations/0071_loot_awards_unmatched_by_award.down.sql
-- Back to one raid-less award per drop. Raid-less awards of the same drop that 0071 let stand
-- side by side cannot all fit 0052's index, so the newest of each drop stays (what 0052's upsert
-- kept) and the older ones are deleted. This down loses those older awards.
delete from loot_awards la
using loot_awards newer
where la.raid_id is null and newer.raid_id is null
  and newer.guild_id = la.guild_id and newer.encounter_id = la.encounter_id and newer.item_id = la.item_id
  and (newer.awarded_at, newer.id) > (la.awarded_at, la.id);
create unique index if not exists loot_awards_drop_unmatched_idx on loot_awards (guild_id, encounter_id, item_id)
  where raid_id is null;
drop index if exists loot_awards_drop_night_idx;
