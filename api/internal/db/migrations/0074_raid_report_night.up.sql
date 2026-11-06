-- api/internal/db/migrations/0074_raid_report_night.up.sql
-- The raid report's night (spec 2026-11-05-web-layout §7.2, §7.4). Two new stores, nothing
-- existing changes shape, so the running API keeps working against this schema.
--
-- 1. An officer's manual attendance fix. guild_raid_attendance keeps holding the effective
--    outcome and excused flag, so every reader (the raid view, the roster's "3 of 4", the CSV
--    export, the Discord reads) reads the fix with no change of its own. This row holds what the
--    officer set and what the log said underneath (auto_outcome, refreshed on every recompute;
--    auto_excused, the officer's excused flag as it stood), so clearing it restores the auto value.
create table guild_raid_attendance_overrides (
  raid_id       bigint not null references guild_raids (id) on delete cascade,
  character_key text not null,
  status        text not null check (status in ('present', 'absent', 'bench', 'excused')),
  auto_outcome  text not null check (auto_outcome in ('present', 'late', 'absent', 'no_show')),
  auto_excused  boolean not null default false,
  set_by        bigint references users (id) on delete set null,
  set_at        timestamptz not null default now(),
  primary key (raid_id, character_key)
);

-- 2. The per-pull digest of a report: each fight's deaths with their cause (killing blow and the
--    last hits) and what each raider brought into the pull, read once from the fight's stored
--    summary so a report view never re-reads the bucket. consumables is null when the pull's log
--    carried no character sheet (COMBATANT_INFO), which is the only honest pull-start reading.
create table guild_report_pulls (
  report_id   text not null references reports (id) on delete cascade,
  fight_index int not null,
  deaths      jsonb not null default '[]',
  consumables jsonb,
  computed_at timestamptz not null default now(),
  primary key (report_id, fight_index)
);
