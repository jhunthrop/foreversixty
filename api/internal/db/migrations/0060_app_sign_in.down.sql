-- Without devices.kind an app token would read as a companion, so every live app token is
-- revoked first: the apps sign in again once the migration is reapplied.
update devices set revoked_at = now() where kind = 'app' and revoked_at is null;
alter table devices drop column kind;
drop table app_auth_flows;
