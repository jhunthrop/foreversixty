# WoW: Forever — Data, API, and Addon Ecosystem (as of Sept 12, 2026)

Almost everything Forever-specific is inferred from precedent until beta opens Sept 17.

## 1. Battle.net APIs

**Namespaces:** `{static|dynamic|profile}-{version}-{region}`. Classic splits into `classic1x-*` (Era/Hardcore/SoD) and `classic-*` (current progression expansion) since June 2021. [Forum clarification](https://us.forums.blizzard.com/en/blizzard/t/wow-classic-era-realm-apis/16812). TBC Anniversary namespace behavior is unsettled as of early 2026. [Thread](https://us.forums.blizzard.com/en/blizzard/t/wow-classic-tbc-anniversary-realms-api-issues/57076)

**Coverage on Classic:** AH, connected realm, realm index, character profile, guild, item, playable class/race, PvP leaderboards, token. **Weak or broken:** spell endpoints 404 often; **no talent-tree endpoint has ever existed for any Classic variant**; no Classic quest endpoint; AH outages. [Spells 404](https://us.forums.blizzard.com/en/blizzard/t/spells-api-returns-404-responses-for-a-lot-of-spells/5173), [AH 404](https://us.forums.blizzard.com/en/blizzard/t/404-for-all-classic-era-namespace-auction-house-endpoints/54307)

**Auth:** client-credentials for game data; user-delegated OAuth for protected profiles (no bulk armory).

**Rate limit:** 36,000 requests/hour per [Developer API Terms of Use](https://www.blizzard.com/en-us/legal/a2989b50-5f16-43b1-abec-2ae17cc09dd6/blizzard-developer-api-terms-of-use). "100/sec" is community lore, not in the ToU.

**ToU constraints:** no paid features gated on API data; no Blizzard trademarks in site name/URL; 30-day data TTL with re-validation (constrains historical armory/deathlog snapshots from API data); no selling/sharing raw or derived data; Blizzard may revoke at will.

**Forever namespace:** no signal anywhere. Precedent leans toward reusing `classic`/`classic1x` (Hardcore, SoD, Cata, MoP all did) rather than a new one; if new, expect it after launch. Make the namespace configurable.

## 2. Client datamining

**Pipeline:** CASC + TACT + DB2 + community listfile ([wowdev/wow-listfile](https://github.com/wowdev/wow-listfile)). Tools: CascView, [wow.export](https://github.com/Kruithne/wow.export) (MIT, active), wago.tools (live DB2 export, build monitor, CDN mirror). [WoWDBDefs](https://github.com/wowdev/WoWDBDefs) auto-updates within minutes of a new build; a release was tagged the same day as this announcement ([202609121904](https://github.com/wowdev/WoWDBDefs/releases/tag/202609121904)). Readers: [DBCD](https://github.com/wowdev/DBCD) (.NET), [pywowlib](https://github.com/wowdev/pywowlib) (MIT).

**How Wowhead gets data:** DB2 datamining plus the crowdsourced "Wowhead Client" Looter addon that logs loot/NPC/quest data during play. [Wowhead explainer](https://www.wowhead.com/classic/news/how-does-wowhead-get-game-data-wowhead-looter-allows-everyone-to-contribute-336588). Reusable model.

**Beta datamineable Sept 17?** Very likely. "Forever" was itself discovered as a codename via datamining of patch 12.1.5 pre-reveal. [Wowhead](https://www.wowhead.com/forever/news/everything-we-know-about-wow-forever-382827). No NDA found for Forever beta; prior Classic betas had none (inferred).

**Legal:** EULA §1.C prohibits mining, but Blizzard has tolerated it for 20+ years and enforcement targets emulators. Low but non-zero, discretionary.

## 3. Addons

**Forever addon policy: no evidence either way.** Zero mention in Blizzard's recap or any coverage. Only confirmed UI detail is gamepad support. **This is the single biggest open risk.** Testable the moment beta opens (does the AddOns folder load?).

**Distribution:** CurseForge (Overwolf), Wago.io (Method), WowUp. Classic Era TOC ~11508; Forever's unknown.

**Hard constraint: addons cannot make network calls.** Every addon-to-website pipeline uses one of:
1. SavedVariables + desktop companion app (TSM, Warcraft Logs uploader tails `WoWCombatLog.txt`). [TSM app](https://support.tradeskillmaster.com/en_US/addon)
2. Manual copy-paste export strings (WeakAuras, GatherMate2).
3. Addon-to-addon in-game sync via `SendAddonMessage` (NovaWorldBuffs, Deathlog).

A real pipeline needs an Electron/Tauri helper or a paste-import UX.

**Policy:** addons must be free, inspectable, no redistributed assets. [Official policy](https://us.forums.blizzard.com/en/wow/t/wow-user-interface-add-on-development-policy/1642). *MDY v. Blizzard* makes ToU enforceable, but passive data-export tools have run unchallenged for years.

## 4. Combat logs / Warcraft Logs

- Classic support today: vanilla.warcraftlogs.com, sod.warcraftlogs.com, classic.warcraftlogs.com. Parent: Archon.
- Added SoD support same day as launch ([source](https://www.archon.gg/classic-sod/articles/news/season-of-discovery-on-warcraft-logs)); Cata day one. Forever likely by Dec 9, unconfirmed.
- API v2 GraphQL, client-credentials, 3,600 points/hour. Commercial use reportedly requires written approval (ToS page 403'd; secondary-sourced).
- Independent parsing viable: [wowparser](https://github.com/spell02/wowparser) (MIT), [wow-combat-log-parser](https://github.com/atinylittleshell/wow-combat-log-parser), [WoWP](https://github.com/rp4rk/WoWP) (Rust); WoWAnalyzer's [CombatLogParser](https://github.com/WoWAnalyzer/CombatLogParser) runs locally.

## 5. Realm population / armory

- ironforge.pro methodology unverified (JS SPA, blank fetch); likely WCL-derived and/or census addons.
- Documented method: CensusPlus addons issue repeated `/who`, log to SavedVariables, uploader pushes to site. [CensusPlusClassic](https://github.com/christophrus/CensusPlusClassic), [WowPopUploader](https://github.com/scarecr0w12/WowPopUploader). Self-selecting sample.
- Battle.net realm endpoint exposes only a coarse population bucket.
- Scraping outside the API violates ToU; census addons run through the sanctioned Lua API and have been tolerated for two decades.

## 6. Open-source starting points

| # | Repo | Stars | License | Notes |
|---|---|---|---|---|
| 1 | [cmangos/classic-db](https://github.com/cmangos/classic-db) | 203 | GPL-3.0 | Best vanilla item/quest/NPC SQL baseline |
| 2 | [mangoszero/database](https://github.com/mangoszero/database) | 163 | Unclear, verify | Alternate vanilla dataset |
| 3 | [vmangos/core](https://github.com/vmangos/core) | 925 | GPL-2.0 | Engine + world DB |
| 4 | [Sarjuuk/aowow](https://github.com/Sarjuuk/aowow) | 233 | No LICENSE file | PHP Wowhead-clone engine used by Turtle WoW's DB |
| 5 | [nexus-devs/wow-classic-items](https://github.com/nexus-devs/wow-classic-items) | 58 | MIT | Clean JSON items/professions/zones; stale 2023 |
| 6 | [BAndysc/WoWDatabaseEditor](https://github.com/BAndysc/WoWDatabaseEditor) | 564 | MIT | Desktop DB editor |
| 7 | [Questie/QuestieDB](https://github.com/Questie/QuestieDB) | 5 | GPL-3.0 | Questie's Lua quest DB |
| 8 | [aaronma37/Deathlog](https://github.com/aaronma37/Deathlog) | 19 | GPL-3.0 | Reference deathlog architecture |
| 9 | [Hoizame/AtlasLootClassic](https://github.com/Hoizame/AtlasLootClassic) | 90 | GPL-2.0 | Loot tables |
| 10 | [shadcn-ui/taxonomy](https://github.com/shadcn-ui/taxonomy) | 19K | MIT | Next.js content-site scaffold |
| 11 | [nextjs/saas-starter](https://github.com/nextjs/saas-starter) | 16K | MIT | Next.js + Postgres starter |

No viable licensed talent-calculator, guild-board, or armory-clone repo found; build those custom.

## 7. Gamepad / console

Official gamepad support confirmed. No console SKU announced; treat any console inference as speculative.

## Recommended data architecture (8-week build)

**Day 1 (Sept 17):** datamine the beta client via wago.tools/WoWDBDefs/DBCD as the primary source for items, spells, NPCs, zones, talents. Bootstrap vanilla-60 baseline from cmangos/classic-db. Keep an MIT combat-log parser as a hedge.

**Contingent on Blizzard:** any Forever namespace (realm status, AH, profiles); talent data via API (never existed; datamine instead); WCL support (likely by Dec 9); any addon-fed pipeline (population, deathlog, prices) depends entirely on whether Forever allows addons. Verify that first.

**Legal summary:** datamining tolerated but discretionary; API usable within ToU (36K/hr, 30-day TTL, no paid API features, no marks in name); WCL API needs approval for commercial use; addon data collection tolerated but not a safe harbor; GPL emulator DBs are the same gray area every fan DB site occupies; clarify mangoszero and aowow licenses before shipping.
