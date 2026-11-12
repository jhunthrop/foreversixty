drop table if exists billing_cancel_tokens;

delete from billing_notices where kind in ('trial_ending', 'trial_receipt');
alter table billing_notices drop constraint if exists billing_notices_kind_check;
alter table billing_notices add constraint billing_notices_kind_check check (kind in ('renewal', 'size'));
