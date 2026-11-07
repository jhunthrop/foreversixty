drop index if exists activity_events_report_idx;
alter table activity_events drop column if exists report_id;
