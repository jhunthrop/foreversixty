-- Live surfaces (mobile Wave 3 lane F14, contract W3-C13). One row per Live Activity (iOS) or
-- ongoing notification (Android) the app asks the server to keep current: a raid's countdown or
-- a server sim's progress. The push senders read rows by (kind, object_id); a row goes when
-- its object ends (expires_at), when its token fails, or when the app deletes it.
--
-- push_token is the ActivityKit update token (iOS) or the FCM registration token (Android);
-- push_to_start_token is iOS's push-to-start token. Both are kept as sent because the senders
-- need them; they are never logged and never answered back.
create table live_activity_tokens (
  id                  bigserial primary key,
  user_id             bigint not null references users (id) on delete cascade,
  device_id           text references devices (id) on delete cascade,
  kind                text not null check (kind in ('raid', 'sim')),
  platform            text not null check (platform in ('ios', 'android')),
  object_id           text not null check (length(object_id) between 1 and 64),
  push_token          text not null,
  push_to_start_token text,
  created_at          timestamptz not null default now(),
  updated_at          timestamptz not null default now(),
  expires_at          timestamptz not null,
  unique (kind, object_id, push_token)
);
create index live_activity_tokens_object_idx on live_activity_tokens (kind, object_id);
create index live_activity_tokens_user_idx on live_activity_tokens (user_id, created_at);
create index live_activity_tokens_expiry_idx on live_activity_tokens (expires_at);
