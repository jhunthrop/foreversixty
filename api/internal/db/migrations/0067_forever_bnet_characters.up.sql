-- api/internal/db/migrations/0067_forever_bnet_characters.up.sql
-- Battle.net import becomes the primary character path on Forever's own
-- namespace, profile-forever-{region} (2026-11-03).
--
-- Every characters row with source = 'bnet' was written from classic1x (Era
-- and Season of Discovery) before Blizzard served Forever; none is a Forever
-- character. The owner ruled that data replaceable: those rows, and what was
-- derived from them, go, and the next sign-in imports the real ones. Users,
-- sessions, go links, grants and guild records stay.
update users set main_character_key = null
 where main_character_key in (select key from characters where source = 'bnet');
delete from guild_characters
 where character_key in (select key from characters where source = 'bnet');
delete from addon_exports
 where source = 'blizzard' and character_key in (select key from characters where source = 'bnet');
delete from characters where source = 'bnet';

-- A character's stable identity is its region and Blizzard character id; the
-- key (<region>/<ruleset>/<name-surname>) is what the addon and the combat
-- log can compute, and it follows the id through a rename or ruleset change.
create unique index if not exists characters_region_bnet_id_key
  on characters (region, bnet_character_id) where bnet_character_id is not null;
