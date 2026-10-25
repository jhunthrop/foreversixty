-- A server run remembers when it finished, why it failed and whether its
-- owner has looked at the result yet.
alter table sims add column if not exists finished_at    timestamptz;
alter table sims add column if not exists failure_reason text not null default '';
alter table sims add column if not exists seen_at        timestamptz;

-- "Done while away": the member's unseen finished server runs, newest first.
create index if not exists sims_user_unseen_idx on sims (user_id, finished_at desc)
  where lane = 'server' and state = 'done' and seen_at is null;

-- The server-run estimate: a spec's recent done runs, newest first.
create index if not exists sims_server_done_spec_idx on sims (spec, finished_at desc nulls last)
  where lane = 'server' and state = 'done';
