drop table if exists sim_specs;
drop table if exists sims;
drop index if exists fight_metrics_execution_idx;
alter table fight_metrics drop column if exists execution_score;
alter table users drop column if exists premium;
