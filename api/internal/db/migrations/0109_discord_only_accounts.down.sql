delete from guild_raid_signups where signed_by in (select id from users where discord_only)
  or acted_by in (select id from users where discord_only);
delete from users where discord_only;
update characters set source = 'bnet' where source = 'discord';
update guild_characters set source = 'bnet' where source = 'discord';
update guild_characters set verified_by = 'officer' where verified_by = 'discord';
alter table guild_characters drop constraint if exists guild_characters_verified_by_check;
alter table guild_characters add constraint guild_characters_verified_by_check
  check (verified_by in ('claim', 'officer', 'invite', 'logs', 'bnet'));
alter table guild_characters drop constraint if exists guild_characters_source_check;
alter table guild_characters add constraint guild_characters_source_check
  check (source in ('export', 'invite', 'claim', 'bnet'));
alter table characters drop constraint if exists characters_source_check;
alter table characters add constraint characters_source_check check (source in ('export', 'bnet'));
alter table users drop column if exists discord_only;
