-- Back to 0037: the requests keep their state, without the export they were made on.
alter table addon_sim_requests drop column if exists export_seen_at;
