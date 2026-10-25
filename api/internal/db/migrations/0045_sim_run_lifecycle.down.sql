drop index if exists sims_user_unseen_idx;
alter table sims drop column if exists seen_at;
alter table sims drop column if exists failure_reason;
alter table sims drop column if exists finished_at;
