-- What a wipe got the boss down to, from the engine's fight entry: null on a
-- kill, on trash, and on fights parsed before the engine recorded it; -1 when
-- the log never showed the boss's health.
alter table fights add column if not exists boss_health_pct real;
