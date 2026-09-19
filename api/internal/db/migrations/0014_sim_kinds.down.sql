drop index if exists sims_user_kind_idx;
alter table sims drop column if exists combos_total;
alter table sims drop column if exists combos_done;
alter table sims drop column if exists stage;
alter table sims drop column if exists headline;
alter table sims drop column if exists kind;
