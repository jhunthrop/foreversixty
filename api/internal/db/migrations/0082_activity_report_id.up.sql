-- Feed rows a report wrote (an upload, a raid's first kill) carry their report, so the feed
-- reads them only while that report is still evidence (presence.EvidenceReport) and public:
-- a report flagged or made private after it completed drops out of /v1/feed, and a deleted
-- report takes its rows along. Existing rows are linked from their /reports/<id> href.
alter table activity_events
  add column if not exists report_id text references reports (id) on delete cascade;

update activity_events e set report_id = r.id
  from reports r
  where e.report_id is null and e.href = '/reports/' || r.id;

create index if not exists activity_events_report_idx on activity_events (report_id)
  where report_id is not null;
