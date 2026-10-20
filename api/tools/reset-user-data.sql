-- Reset every account and everything it produced, so the site can be walked as a
-- first-time user. Reference data stays: builds (shared planner links the guides point
-- at), sim_specs (the engine's validation state), subscribers (the mailing list),
-- schema_migrations. Run inside one transaction; take a dump first:
--   pg_dump "$DIRECT_URL" --no-owner --no-privileges -Fc -f "backup-$(date +%F).dump"
--   psql "$DIRECT_URL" -v ON_ERROR_STOP=1 -f api/tools/reset-user-data.sql
begin;

truncate table
  -- accounts and their sign-in state
  sessions, login_tokens, devices, device_connects, pairing_codes,
  -- characters and the addon's data
  characters, character_syncs, addon_exports, addon_inbox, addon_sim_requests, sims,
  -- logs and everything derived from them
  reports, fights, fight_metrics, raw_chunks, uploads, kill_duration_digests,
  percentile_digests, rating_percentile_digests, rating_scores, moderation,
  -- guilds
  guilds, guild_members, guild_characters, guild_claim_attempts, guild_claim_resolutions,
  guild_roster_snapshots, loot_awards, seed_rows,
  -- billing (the Stripe side keeps its customers; see the note in the ledger)
  entitlements, entitlement_audit, entitlement_anomalies, pending_checkouts,
  stripe_customers, stripe_events,
  -- the accounts themselves, last
  users
restart identity cascade;

commit;

-- Afterwards: select count(*) from users; -- 0
