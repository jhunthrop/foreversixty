-- A player may hide their own Forever Score (owner approval 2026-11-08; addon-v2 contract 4,
-- export.score_visible): false keeps the score from everyone but the character's own account,
-- on every surface. Visible by default. A new column with a default, so the running API keeps
-- working against this schema.
alter table characters add column if not exists score_visible boolean not null default true;
