-- api/internal/db/migrations/0052_guild_attendance.up.sql
-- Lane G2, attendance and the raid-night report (docs/superpowers/specs/2026-10-26-guild-suite-design.md,
-- design/specs/2026-10-27-guild-attendance.md §9). The whole lane's schema lands here.

-- The zone the engine read from the log itself (ZONE_CHANGE), as distinct from reports.zone, which
-- a client types at create and a whole-file upload leaves empty. Null until a parse or the live
-- ingest learns it; the matcher compares it with the curated zone's log name.
alter table reports add column log_zone text;

-- The guild the matcher reads a report as, by its owner's standing (reports.AttributeGuild).
-- Distinct from guild_id, which only the officer PATCH sets and which carries edit rights, the
-- guild leaderboards and character verification; the matcher reads coalesce(guild_id,
-- matched_guild_id) and nothing else reads this column.
alter table reports add column matched_guild_id bigint references guilds (id) on delete set null;
create index reports_matched_guild_idx on reports (matched_guild_id) where matched_guild_id is not null;

-- A raid record learns its log. pre_done_status is what status returns to when an officer clears
-- the match; matched_by says whether the auto-matcher or an officer picked the log; matched_at is
-- when; report_posted_at guards the one report_ready post per raid.
alter table guild_raids add column pre_done_status text check (pre_done_status in ('planned', 'locked'));
alter table guild_raids add column matched_by text check (matched_by in ('auto', 'hand'));
alter table guild_raids add column matched_at timestamptz;
alter table guild_raids add column report_posted_at timestamptz;
-- One log, one raid (R-A7).
create unique index guild_raids_report_idx on guild_raids (report_id) where report_id is not null;

-- A cleared match: the auto-matcher never picks this report for this raid again.
create table guild_raid_rejected_reports (
  raid_id     bigint not null references guild_raids (id) on delete cascade,
  report_id   text not null references reports (id) on delete cascade,
  rejected_by bigint references users (id) on delete set null,
  rejected_at timestamptz not null default now(),
  primary key (raid_id, report_id)
);

-- One row per character per matched raid. signed and role are the sign-up as it stood (null for a
-- character who joined without one); came_as names the other character of the same account that
-- was in the log (R-A2) and came_as_key its character key; excused is an officer's flag and survives a recompute.
create table guild_raid_attendance (
  raid_id             bigint not null references guild_raids (id) on delete cascade,
  character_key       text not null,
  name                text not null,
  class               text not null,
  spec                text not null default '',
  signed              text check (signed in ('in', 'late', 'bench', 'out')),
  role                text,
  outcome             text not null check (outcome in ('present', 'late', 'absent', 'no_show')),
  came_as             text,
  came_as_key         text,
  first_seen_offset_s int,
  pulls_present       int not null default 0,
  pulls_total         int not null default 0,
  excused             boolean not null default false,
  computed_at         timestamptz not null default now(),
  primary key (raid_id, character_key)
);

-- Loot remembers the raid it dropped in, so re-awarding the same item next week is a new row
-- rather than an overwrite of last week's.
alter table loot_awards add column raid_id bigint references guild_raids (id) on delete set null;
-- True for an award the raid's match attached (it had no raid and was made in the match
-- window); clearing the match detaches only these, never an award filed under the raid itself.
alter table loot_awards add column attached_by_match boolean not null default false;
drop index if exists loot_awards_drop_idx;
-- One award per drop per raid. Postgres keeps NULL raid_id rows distinct in this index, so the
-- awards made with no raid keep their one-per-drop rule through the partial index below, which is
-- the conflict target CreateLootAward upserts on.
create unique index loot_awards_drop_idx on loot_awards (guild_id, raid_id, encounter_id, item_id);
create unique index loot_awards_drop_unmatched_idx on loot_awards (guild_id, encounter_id, item_id)
  where raid_id is null;
