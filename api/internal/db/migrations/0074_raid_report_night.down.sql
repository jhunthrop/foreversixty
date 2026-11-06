-- api/internal/db/migrations/0074_raid_report_night.down.sql
-- Puts each overridden attendance row back to what the log said before dropping the fixes, so
-- the attendance counts read as if no fix was ever made.
update guild_raid_attendance a set outcome = o.auto_outcome, excused = o.auto_excused
from guild_raid_attendance_overrides o
where o.raid_id = a.raid_id and o.character_key = a.character_key;
drop table if exists guild_raid_attendance_overrides;
drop table if exists guild_report_pulls;
