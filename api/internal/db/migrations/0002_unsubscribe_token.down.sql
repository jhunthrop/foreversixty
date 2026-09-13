drop index if exists subscribers_unsubscribe_token_idx;
alter table subscribers drop column if exists unsubscribe_token;
