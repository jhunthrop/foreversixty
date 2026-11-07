-- Automatic Battle.net refresh (api/internal/bnetimport/visit.go, loop.go).
--
-- bnet_synced_at is when an account-wide Battle.net sync last started: a sign-in import, or
-- a visit refresh. A visit claims it with a compare-and-set, so the column is both the
-- per-account cooldown and the dedupe across Cloud Run instances, and a recent value marks
-- the account active for the background loop's shorter window.
--
-- bnet_token_box is the sign-in's Battle.net access token, sealed with the notify key
-- (notify.Sealer), and bnet_token_expires_at its expiry. Battle.net issues no refresh
-- token, so a visit uses the sealed token until it expires and then refreshes character by
-- character from public profiles. api/internal/bnetimport writes these columns; sign-out
-- (auth.Store.ForgetBnetToken) clears the token, and deleting the users row takes it along.
alter table users
  add column if not exists bnet_synced_at timestamptz,
  add column if not exists bnet_token_box bytea,
  add column if not exists bnet_token_expires_at timestamptz;

create index if not exists characters_bnet_refreshed_idx on characters (refreshed_at) where source = 'bnet';
