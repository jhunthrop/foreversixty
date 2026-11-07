-- Self-service account deletion (DELETE /v1/me, api/internal/account). Deleting a users row
-- lets every user-keyed table follow its own foreign key; this is the one key that blocked
-- it. A raid is the guild's record, not its creator's: when the officer who created it
-- deletes their account the raid stays, with no creator. Sign-ups are answered per account
-- by the deletion itself (the account's own go, an officer's acted_by moves to the account
-- it answered for), and guild_claim_resolutions.moderator_id keeps restrict on purpose:
-- staff accounts are closed by support, never self-served.
alter table guild_raids alter column created_by drop not null;
alter table guild_raids drop constraint if exists guild_raids_created_by_fkey;
alter table guild_raids add constraint guild_raids_created_by_fkey
  foreign key (created_by) references users (id) on delete set null;
