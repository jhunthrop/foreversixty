-- api/internal/db/migrations/0070_guild_invite_shown.up.sql
-- Invite links an officer can show again (design/specs/2026-11-05-guild-gating-and-invites.md
-- §4.1). invite_token_hash stays the only lookup key; invite_token_enc is the same token sealed
-- with the Discord webhook envelope (AES-256-GCM, NOTIFY_ENCRYPTION_KEY) so a verified officer can
-- read it back, and invite_rotated_by is who made it, so the join page can name the inviter.
-- Additive and nullable: a running build that does not know these columns keeps rotating and
-- redeeming by the hash alone, and the reader ignores a sealed token whose hash no longer matches.
alter table guilds
  add column if not exists invite_token_enc bytea,
  add column if not exists invite_rotated_by bigint references users (id) on delete set null;
