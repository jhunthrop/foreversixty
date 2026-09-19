drop index if exists builds_user_idx;
alter table builds drop column if exists user_id;
