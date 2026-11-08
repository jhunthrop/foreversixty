-- The member's opt-in to Discord DMs from their guild's officers (spec 2026-11-05-web-layout D16:
-- mentions in the raid channel are the default, DMs are opt-in). Off until the member turns it
-- on; only then does the officers' not-answered list carry a link to message them.
alter table notification_settings add column if not exists officer_dms boolean not null default false;
