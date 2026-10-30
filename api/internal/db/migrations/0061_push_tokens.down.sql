drop trigger if exists devices_revoked_drop_push_tokens on devices;
drop function if exists push_tokens_drop_revoked();
drop table if exists push_tokens;
