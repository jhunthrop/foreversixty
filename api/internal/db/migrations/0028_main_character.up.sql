-- The account's main character: the one the site opens with, everywhere. Every other
-- character the account owns is an alt. A key, not a foreign key: characters can be
-- re-keyed by the Battle.net refresh (bnetimport.rekeyInPlace), which updates this too, and
-- a stale key simply reads as "no main set".
alter table users add column if not exists main_character_key text;
