drop index if exists characters_bnet_refreshed_idx;
alter table users
  drop column if exists bnet_token_expires_at,
  drop column if exists bnet_token_box,
  drop column if exists bnet_synced_at;
