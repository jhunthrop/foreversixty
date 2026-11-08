-- The weekly tier move (spec 2026-11-08-app-flows F4): the grade a spec's row last told its
-- mains about, so the hourly snapshot publishes each move once.
alter table tier_grade_weeks add column if not exists notified_grade text;
