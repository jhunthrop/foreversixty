-- addon_exports.updated_at only moves when the export string changes, so the inbox ack, which
-- told the addon when the site last saw the character, kept saying the old time after an
-- unchanged re-upload. seen_at is the time of the last accepted upload, changed or not;
-- updated_at keeps meaning "the export last changed". Existing rows start at updated_at.
alter table addon_exports add column if not exists seen_at timestamptz;
update addon_exports set seen_at = updated_at where seen_at is null;
alter table addon_exports alter column seen_at set not null;
alter table addon_exports alter column seen_at set default now();

-- A live report's waiting screen needs to say when the companion last uploaded and when it was
-- last heard. fights.uploaded_at is when the fight's row was last written (existing rows start
-- at their report's creation, the best time known); reports.heartbeat_at is the companion's
-- own "still tailing" ping, null until the first one.
alter table fights add column if not exists uploaded_at timestamptz;
update fights f set uploaded_at = r.created_at from reports r where r.id = f.report_id and f.uploaded_at is null;
alter table fights alter column uploaded_at set not null;
alter table fights alter column uploaded_at set default now();
alter table reports add column if not exists heartbeat_at timestamptz;
