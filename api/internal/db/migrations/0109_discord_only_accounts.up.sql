-- Discord sign-up by character pick (docs/superpowers/specs/2026-11-14-character-directory-discovery.md
-- §4): a Discord user with no site account signs up by picking their character from the guild's
-- character directory roster. The bot creates a Discord-only account on first use (a users row
-- with no Battle.net identity, no email, no BattleTag, discord_only true) plus its discord_links
-- row. The pick writes characters and guild_characters with source and verified_by 'discord':
-- signable, marked for officers, never above member rank. A Battle.net sign-in that owns the
-- character takes the claim over; linking the same Discord account merges the rest.
alter table users add column if not exists discord_only boolean not null default false;

alter table characters drop constraint if exists characters_source_check;
alter table characters add constraint characters_source_check check (source in ('export', 'bnet', 'discord'));

alter table guild_characters drop constraint if exists guild_characters_source_check;
alter table guild_characters add constraint guild_characters_source_check
  check (source in ('export', 'invite', 'claim', 'bnet', 'discord'));
alter table guild_characters drop constraint if exists guild_characters_verified_by_check;
alter table guild_characters add constraint guild_characters_verified_by_check
  check (verified_by in ('claim', 'officer', 'invite', 'logs', 'bnet', 'discord'));
