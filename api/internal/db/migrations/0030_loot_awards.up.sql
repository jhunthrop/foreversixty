-- api/internal/db/migrations/0030_loot_awards.up.sql
-- The guild control centre's Loot tab (design/specs/2026-10-04-guild-page.md §4.F, §9):
-- no table persisted a loot decision before this. report_id is nullable so an officer can
-- mark a drop awarded before the kill that dropped it is logged (the contract's own
-- allowance, docs/contracts/2026-10-04-guild-centre-api.md).
create table if not exists loot_awards (
  id            bigserial primary key,
  guild_id      bigint not null references guilds (id) on delete cascade,
  encounter_id  bigint not null,
  item_id       bigint not null,
  character_key text not null,
  awarded_by    bigint references users (id) on delete set null,
  awarded_at    timestamptz not null default now(),
  report_id     text references reports (id) on delete set null
);
-- One open award per (guild, encounter, item) at a time - a second Award for the same drop
-- replaces the first via DELETE then POST, never a silent second row the Loot tab would not
-- know how to pick between.
create unique index if not exists loot_awards_drop_idx on loot_awards (guild_id, encounter_id, item_id);
create index if not exists loot_awards_character_idx on loot_awards (guild_id, character_key);
