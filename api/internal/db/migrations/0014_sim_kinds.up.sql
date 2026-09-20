-- Every sim row says which tool produced it, carries the one line the
-- history list shows for it, and — while a bulk run is going — how far
-- through its stages it is (contract 10.6).
--
-- headline is stored rather than composed on read: the result blob is a
-- few hundred kilobytes and a history page is a hundred rows, so reading
-- it back to compose one line would detoast tens of megabytes per page.
-- A stored result never changes, so a stored headline never goes stale.

alter table sims add column if not exists kind         text not null default 'run';
alter table sims add column if not exists headline     text not null default '';
alter table sims add column if not exists stage        int  not null default 0;
alter table sims add column if not exists combos_done  int  not null default 0;
alter table sims add column if not exists combos_total int  not null default 0;

-- The history list filtered by kind. sims_user_idx still serves the
-- unfiltered list, whose ordering this index cannot satisfy.
create index if not exists sims_user_kind_idx on sims (user_id, kind, created_at desc);

-- Rows saved before this migration have no headline, and would render as
-- a blank line in the history list. Every one of them is necessarily a
-- plain run — bulk and weights did not exist in the envelope then — so
-- their headline is exactly what sims.Headline composes for that branch:
-- the mean rounded to the nearest whole, halves away from zero, grouped
-- in threes, followed by " DPS". round() on numeric rounds half away
-- from zero, which is math.Round's own rule; dps_mean is numeric (0013),
-- and the cast keeps that true if the column's type is ever widened.
--
-- Only done rows: a queued or errored run has no result, and handing it
-- a partial figure would state a conclusion it never reached. Only rows
-- whose headline is still empty, so re-running this is a no-op.
update sims
   set headline = to_char(round(dps_mean::numeric), 'FM999,999,999,999,999,990') || ' DPS'
 where headline = '' and kind = 'run' and state = 'done';
