# Spec brief for the ux-designer (2026-11-14)

Inputs: 00-facts.md (verified; quote it, never the owner's wording where they differ), 01-plan.md, 02-officer-plan.md (answer every objection in the spec, by number, in a short "Officer objections" section).

Owner rules that bind every spec:
- The addon makes no network calls. Data leaves the game only through SavedVariables, written at logout or /reload, picked up by a companion. Addon-to-addon messages (FS1 on GUILD/RAID/PARTY) work only while both players are online in game.
- Each spec has a "Data paths and latency" table: every fact shown, its source, how fresh it can be (seconds, minutes, hours, days), and what is shown when it is unknown. "Unknown" is never shown as "missing".
- Blizzard data freshness is NOT hourly: signed-in members 3 h (active) or 20 h, with enchants and suffixes; directory characters up to 72 h, no enchants, no suffixes. Say so.
- Every check and every gear row carries its time stamp.
- Existing building blocks to reuse, not reinvent: loot council candidates and award records (api/internal/guilds/loot_council_*.go, loot_candidates.go), loot values (api/internal/sims/lootvalues.go, addon loot!/LootCouncil.lua), readiness and ready check (guilds/raid_prep.go, readiness.go, ready_check.go), raid report pulls with death causes (guilds/raids_report_pulls.go), raid night (guilds/raid_night.go), Live Activity (guilds/raid_live.go), snapshot! v2 readiness fields (addon Comms.lua), RaidPrep.lua bag reading, Discord bot OpenDM, push senders. Web components from web/src/components; app tabs Home, Guild, Gear, Activity (design/specs/2026-11-10-guild-tab.md).

Relay design (owner approved; goes in prep reminder and live pulls):
- Raider addons send their in-game checks (bags, consumables, attunement, readiness: snapshot! v2) over GUILD/RAID.
- Addons whose player runs the companion announce themselves as relays (a new FS1 field or kind; 255-byte limit).
- The relay's addon collects everyone's results; its companion carries them to the site, Discord and the app, but only via SavedVariables.
- v1: a relay-only button "Send tonight's prep" that reloads the UI (a short loading screen) so the companion uploads. Out of combat only; the button says what it does. ReloadUI's rule in the Forever client (protected, or needs a hardware event) is a spike check; a button click is a hardware event either way.
- Several relays online: dedupe (one result per character, newest stamp wins; one upload per raid). No relay online: officers are told plainly. Every check stamped with its time.
- Verification task (a build-lane prerequisite, run on the Windows test box over ssh jh@192.168.1.182, never the Mac): does WoWChatLog.txt (/chatlog) record CHAT_MSG_ADDON traffic, or any channel an addon can write to live? If yes, the companion could tail it for a live path with no reload. Note ToS and spam risks. Choose nothing that puts messages in players' visible chat.

Live pulls specifics: in game, the raid leader's addon explains each pull from live combat events with no companion (new COMBAT_LOG_EVENT_UNFILTERED work; size it honestly). Site, second screen and phone need one companion in the raid; it uploads each fight as it closes (seconds). Several companions: one chosen.

Prep reminder specifics: Discord DM or phone push one hour before carries only Blizzard data (gear, enchants and resistance gear where known) plus bags from that raider's last companion upload, labelled with its age. Never imply a live bag check off game. In game: the raider's own addon checks its bags at T-1h and shows the raider directly (raid time from the companion inbox or an officer/relay addon's GUILD announcement). The officer's addon collects online raiders' checks into one in-game list, with "could not check" for the rest. iOS push is off in production today (no APNs keys); say so.

Style: no em dashes, no all-caps headings, never name the owner, game words only (players are veterans), no explainer filler. Each spec: every surface (web guild page, Discord, apps, addon), every state (default, loading, empty, error, partial data, stale, no companion, no relay), copy verbatim, data sources with accuracy, build lanes with rough size (S/M/L plus days), open owner decisions. Use the fixture world from design/specs/2026-11-10-guild-tab.md (Olympus XXVII, Thandrel). Target 150 to 260 lines per main spec; guild-later 50 to 90 lines.
