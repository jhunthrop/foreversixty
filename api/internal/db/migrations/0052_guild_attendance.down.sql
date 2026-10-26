drop index if exists loot_awards_drop_unmatched_idx;
drop index if exists loot_awards_drop_idx;
-- Collapse to one row per drop before the narrower unique index returns: the newest award stays.
delete from loot_awards a using loot_awards b
 where a.guild_id = b.guild_id and a.encounter_id = b.encounter_id and a.item_id = b.item_id
   and (a.awarded_at, a.id) < (b.awarded_at, b.id);
create unique index loot_awards_drop_idx on loot_awards (guild_id, encounter_id, item_id);
alter table loot_awards drop column if exists attached_by_match;
alter table loot_awards drop column if exists raid_id;
drop table if exists guild_raid_attendance;
drop table if exists guild_raid_rejected_reports;
drop index if exists guild_raids_report_idx;
alter table guild_raids drop column if exists report_posted_at;
alter table guild_raids drop column if exists matched_at;
alter table guild_raids drop column if exists matched_by;
alter table guild_raids drop column if exists pre_done_status;
drop index if exists reports_matched_guild_idx;
alter table reports drop column if exists matched_guild_id;
alter table reports drop column if exists log_zone;
