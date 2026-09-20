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
