-- Per-account model spend (log-coach Task 45): each account's UTC day row carries the USD its
-- reviews cost, beside the generations it counts; an account's month is the sum of its days.
alter table coach_account_day add column if not exists usd numeric(12,6) not null default 0;
