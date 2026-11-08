-- The player's Activity inbox (GET /v1/me/activity): one row per notification event an
-- account was a recipient of, whatever its push and Discord choices and quiet hours, so a
-- player with pushes off still sees what changed. The group a kind belongs to lives in code
-- (api/internal/events/groups.go), so a new event needs no migration. A republished event is
-- one row (the dedupe key, as in notification_deliveries). Rows past 90 days are pruned by
-- the raid-posts job.
create table user_activity (
  id          bigserial primary key,
  user_id     bigint not null references users (id) on delete cascade,
  kind        text not null check (kind ~ '^[a-z]+\.[a-z_]+$'),
  title       text not null default '',
  body        text not null default '',
  link        text not null default '',
  dedupe_key  text not null check (char_length(dedupe_key) between 1 and 200),
  occurred_at timestamptz not null default now(),
  read_at     timestamptz,
  unique (user_id, dedupe_key)
);
-- The page read, newest first, and its cursor.
create index user_activity_user_id_idx on user_activity (user_id, id desc);
-- The tab badge's unread count.
create index user_activity_unread_idx on user_activity (user_id) where read_at is null;
-- The retention prune.
create index user_activity_occurred_idx on user_activity (occurred_at);
