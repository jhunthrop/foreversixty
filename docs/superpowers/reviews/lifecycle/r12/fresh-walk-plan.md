# Fresh-account walk — plan (21 Oct 2026, evening)

Goal: the owner's Battle.net account looks brand new to the site, then the first-week path is walked for real: companion install → device login → addon in game → first-login welcome → logout → the character on the site → claim the guild → send a build → sim it.

## What "fresh" means in the production database (Neon)

The account row is `users` (id, bnet_sub, battletag, email, role). Deleting it cascades `sessions`, `devices`, `login_tokens`, `pairing_codes`, consents and entitlement rows, and sets `null` on `builds.user_id`, `guilds.claimed_by` / `claim_pending_by` / `claim_contested_by`, `loot_awards.awarded_by`, `moderation.actor_id`, `reports.user_id`. `characters.user_id` is set null, not deleted, so for a true fresh arrival the owner's `characters` rows (and their `character_syncs`, `addon_exports`, `addon_inbox`, `addon_sim_requests`) go too. Reports, fights, metrics, rankings and the guild rows stay: the beta nights are history, and the guild simply becomes unclaimed.

```sql
begin;
-- scope A: the owner's account only
create temp table gone as select id from users where battletag = 'hunthrop#1894';
delete from addon_inbox        where character_key in (select key from characters where user_id in (select id from gone));
delete from addon_sim_requests where character_key in (select key from characters where user_id in (select id from gone));
delete from addon_exports      where character_key in (select key from characters where user_id in (select id from gone));
delete from character_syncs    where character_key in (select key from characters where user_id in (select id from gone));
delete from characters         where user_id in (select id from gone);
delete from users              where id in (select id from gone);
-- look before commit: select count(*) from users; select name, claimed_by from guilds;
commit;
```

Scope B (every account) is the same with `gone` = `select id from users`. Column names above are to be checked against the live schema (`\d characters`, `\d addon_inbox`) before running; a backup of the touched tables is taken first with `pg_dump -t users -t characters -t character_syncs -t addon_exports -t addon_inbox -t addon_sim_requests`.

## Blocked on the owner

1. The production connection string (Neon): put it in the repo-root `.env` as `PROD_DATABASE_URL=` (git-ignored). This box has no gcloud, so Secret Manager is not reachable from here.
2. The scope: the owner's account only (A) or every account (B). Other beta testers who signed in would lose their sessions, mains and consents under B.
3. The game: launch WoW and log in; the walk never launches the game or types credentials.
4. The site sign-in on production (Battle.net OAuth in the browser pane) and the companion's device-login approval, both the owner's clicks.

## The walk, once unblocked

Companion from the released exe into a scratch home (`FS_COMPANION_HOME`) → "Connect" → device login → addon zip from the release into the client folder (the game closed, as the updater requires) → the owner logs in → `/fs` welcome, `/fs help` → logout to the character screen → the companion uploads → `/character/us/pvp/bow-jackzon` as a stranger, then signed in → claim OLYMPUS XXVII → planner send → simulator. Captures into `docs/superpowers/reviews/lifecycle/r12/fresh-*.png`.
