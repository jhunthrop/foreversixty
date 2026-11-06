-- Raid assignments (design/specs/2026-11-05-web-layout.md §7.2 Raid planner, §7.4, D13).
--
-- raid_assignments is the raid leader's pick for one slot of one encounter's template
-- (Go: api/internal/guilds/raid_assign_templates.go): main tank, whelps west, and so on.
-- position orders the names inside a slot. A character appears once per slot; it may hold
-- several slots (a priest heals the main tank and keeps Fear Ward up). Assignments are
-- manual only: nothing writes these rows but an officer's PUT.
create table if not exists raid_assignments (
  raid_id       bigint not null references guild_raids (id) on delete cascade,
  encounter_id  bigint not null,
  slot_key      text not null check (char_length(slot_key) between 1 and 40),
  position      smallint not null check (position >= 0),
  character_key text not null,
  set_by        bigint references users (id) on delete set null,
  set_at        timestamptz not null default now(),
  primary key (raid_id, encounter_id, slot_key, position),
  unique (raid_id, encounter_id, slot_key, character_key)
);

-- The addon inbox reads a character's assignments for its next raid by key.
create index if not exists raid_assignments_character on raid_assignments (character_key);

-- raid_assignment_posts is the last "Post to Discord" of a raid's assignments: one row per
-- raid, so the short re-post throttle is one atomic upsert (as guild_raid_pings).
create table if not exists raid_assignment_posts (
  raid_id        bigint primary key references guild_raids (id) on delete cascade,
  last_posted_at timestamptz not null,
  posted_by      bigint references users (id) on delete set null
);
