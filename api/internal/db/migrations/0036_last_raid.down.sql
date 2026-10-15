-- Back to 0034's addon_sim_requests: only the requests that started a run.
delete from addon_sim_requests where state is distinct from 'queued';
alter table addon_sim_requests drop column if exists state;

alter table rating_scores
  drop column if exists consumables_missed,
  drop column if exists loss_seconds,
  drop column if exists biggest_loss,
  drop column if exists deaths;
