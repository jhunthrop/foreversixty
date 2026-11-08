-- A Late answer may say how late (the app's 15, 30 and 60 minute chips, spec
-- 2026-11-08-app-flows F1); null is a Late with no minutes, and every other state.
alter table guild_raid_signups
  add column if not exists late_minutes smallint
    check (late_minutes is null or (late_minutes in (15, 30, 60) and state = 'late'));
