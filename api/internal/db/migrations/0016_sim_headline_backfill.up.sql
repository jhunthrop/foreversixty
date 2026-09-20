-- This backfill was originally written as an addition to 0014_sim_kinds,
-- which is where headline was introduced. It had to move here instead:
-- 0014 already shipped and its startup migration already ran against the
-- production database (the part-1 API merge; GET /v1/phases and the
-- kind= filter are live on it). golang-migrate records a schema version
-- once and never re-runs a migration it has already applied, so editing
-- 0014 after the fact would be invisible to the only database that has
-- rows needing this — every environment that matters would silently skip
-- the backfill. A schema change that has already run in production can
-- only be followed up by a new migration, never edited in place.
--
-- Rows saved before 0014 have no headline, and would render as a blank
-- line in the history list. Every one of them is necessarily a plain
-- run — bulk and weights did not exist in the envelope then — so its
-- headline is exactly what sims.Headline composes for that branch: the
-- mean rounded to the nearest whole, halves away from zero, grouped in
-- threes, followed by " DPS". round() on numeric rounds half away from
-- zero, which is math.Round's own rule; dps_mean is numeric (0013), and
-- the cast keeps that true if the column's type is ever widened.
--
-- Only done rows: a queued or errored run has no result, and handing it
-- a partial figure would state a conclusion it never reached. Only rows
-- whose headline is still null or empty, so re-running this migration —
-- or running it after a row has since earned a real headline — is a
-- no-op rather than a clobber. The column is `not null default ''`, so
-- null cannot occur today; the check costs nothing and survives a future
-- widening of the column.
update sims
   set headline = to_char(round(dps_mean::numeric), 'FM999,999,999,999,999,990') || ' DPS'
 where (headline is null or headline = '') and kind = 'run' and state = 'done';
