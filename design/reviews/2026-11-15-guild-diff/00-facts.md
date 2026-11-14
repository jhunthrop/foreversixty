# Guild differentiators: verified data facts (2026-11-14)

Every line below was checked against the code on origin/main. Specs must quote these, not the brief.

## Blizzard profile data

- Two stores, two cadences.
  - Signed-in members' characters (bnetimport, `characters` + `addon_exports` source `blizzard`): refreshed every 3 h when the account was active in the last 14 days, else every 20 h (api/internal/bnetimport/refresh.go `activeRefreshWindow`, `idleRefreshWindow`; loop runs hourly). Equipment is encoded with the PERMANENT enchant id and random suffix (api/internal/bnetbuild/gear.go `permanentEnchant`). Statistics are captured too (bnetimport/capture.go).
  - Everyone else, the character directory (`character_directory`, migration 0106): the loop starts a pass every hour, but a character's profile is re-read only once it is 72 h old (api/internal/directory/directory.go `profileAfter = 72h`); rosters every 24 h. `equipment` jsonb is a summary: slot, item id, name, quality, item level. NO enchants, NO random suffix (directory/parse.go `EquippedItem`). Only public profiles (`public` true once the profile answered). Rows not confirmed in 30 days are deleted. `directory_optouts` (0108) removes a character forever.
- FALSE in the brief: "Blizzard profile refreshed hourly" and "enchants from the Blizzard profile" hold only for signed-in members (3 h / 20 h, enchants yes). Directory characters are up to 72 h old and carry no enchants.
- Resistance gear: derivable from equipped item ids against our item data. Resistance enchants (cloak, ZG/AQ libram style) only for signed-in members. Whether Blizzard's classic /statistics answer carries resistances is unverified.
- Gear consent: the guild Loot tab and roster read a member's gear only from the addon export and only at gear consent (guilds/loot_candidates.go header). Using directory gear for a guild loot plan is a new use of public data: an owner decision.

## Loot value (the gain)

- api/internal/sims/lootvalues.go: a stat-line score delta (the same Gear.score arithmetic as GearScore) of the item against what the character wears in that slot, under the character's own weights run if it has one, else the build's shipped spec weights; turned into DPS with data/builds/<build>/dps-per-point.json when the spec has an entry whose reference stat weighs exactly 1. dps-per-point carries per-band (20/30/40/50/60) values; healers are marked `healing` (HPS). Tank specs have no entry: their gain is in "score", not DPS. Planned = the item is the spec's level-60 BiS pick.
- This is linear weights times the sim's DPS per point, not a full sim per item. The addon's live loot session (`loot!` value, unit dps|score) uses the same rule (LootCouncil.lua).
- data/builds/<build>/loot.json is keyed by sources (bosses) so a per-boss drop list is available. The server LootValueSet groups a raid's items together today.
- Existing guild loot features: loot council (loot_council_*.go) already ranks candidates in two tiers (gain_dps tier 0, ilvl_delta tier 1) with attendance over the last 4 raids, loot this tier, MS/OS, votes; award records (migration 0073). A character with no gear on file is no candidate.

## Companion (desktop, opt in)

- Tails WoWCombatLog live: app ticks every 1 s (companion/internal/app TickEvery), watcher reads up to 1 MiB per poll.
- Uploads each fight as it closes (client.PutFight) with metrics rows per player (DPS, HPS, damage taken, active ms, deaths, encounter id, kill) plus a live snapshot every 5 s (pipeline.LiveEvery, PutLive) and raw 4 MiB chunks. So a pull reaches the site within seconds of ENCOUNTER_END plus the API's fold; it does not depend on SavedVariables.
- One companion in the raid covers everyone in its combat log range (the log records the whole raid within range). Several companions in the same raid upload duplicate reports: the site already attributes and merges reports (reports/attribution.go); dedupe for live pulls must pick one.
- Engine already has per-pull: deaths with killing blow, last hits, heals, auras held and lost (logs/engine/summary/deaths.go); phases; mechanics per encounter (logs/engine/mechanics/encounters.json); consumables at pull and potion uses (mechanics/consumables/atpull.go). Consumables at pull are "known" only when the log has a COMBATANT_INFO row for the player; the one real Forever log in the repo has none, so no real Forever boss pull has proven the consumables reading yet (atpull_test.go comment). Without it, preparation reads unknown, never missing.
- guilds/raids_report_pulls.go already serves a pull list with death causes and consumables by pull for a matched report.
- The companion writes the addon's inbox (contract 3: raid_signup, raid_assignment, raid_roster, loot_values every 10 min) and reads the addon's outbox from SavedVariables.

## Addon (in game)

- Makes NO network calls. Data leaves the game only through SavedVariables, written at logout or /reload, read by a companion. ReloadUI is not called anywhere in the addon today; whether the Forever client requires a hardware event for it is unverified (spike check). A button click is a hardware event, so a "Send" button is safe either way.
- Addon-to-addon messages (Comms.lua, prefix FS1, 255 bytes, GUILD/RAID/PARTY): hello, share, push, loot?, loot!, snapshot?, snapshot!. snapshot! v2 already carries addonVersion, attuned (Onyxia today: quests 6502/6602 or item 16309), encounterItems, supplies (counts by kind: flask, elixir, food, fire/frost/nature/shadow/arcane protection). Works only while both players are online in game.
- Bag reading exists (RaidPrep.lua via C_Container; consumable class stacks with counts; resistances via UnitResistance; keyring; attunement quests/items). A reader with no client API answers nil: "not tracked", never "missing".
- The raider's addon knows tonight's raid only from the companion inbox (raid_signup). A raider with no companion has no raid time unless another addon announces it over GUILD.
- No live combat-event parsing for the raid exists in the addon today (Recorder.lua records the player's own data). A raid leader's in-game pull explanation needs new COMBAT_LOG_EVENT_UNFILTERED work.

## Notifications

- Discord bot can DM a linked user (discordbot/rest.go OpenDM; used today for applicant DMs). notify/events.go: per-account Discord sink has no DM sender for other event kinds yet. Raid reminder today is a day-before channel post (notify/reminder.go).
- Phone push: APNs and FCM senders exist (api/internal/push). Production api service (us-east1) has FCM_PROJECT_ID set and no APNS_* variables, so iOS push is off in production; Android depends on FCM credentials. The app's Activity tab is the push inbox (spec 2026-11-08-app-flows).
- Live Activity hooks for raid start/end exist (guilds/raid_live.go, livesurface).

## Existing guild readiness

- raid_prep.go, readiness.go, ready_check.go, me_checklist: flask/food on hand, resistance totals, attunement from the last addon export; never guesses.

## Added after the specs were drafted

- The companion reads SavedVariables on its 1 s tick (companion/internal/addon/sync.go Poll, memoised on file mtime), so an outbox written at logout or /reload reaches the site within seconds. The inbox (site to addon) is fetched every 10 min (InboxEvery) or right after an upload, and the addon reads it only at its next login or /reload.
