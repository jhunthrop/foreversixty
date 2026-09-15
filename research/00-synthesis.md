# World of Warcraft: Forever — Research Synthesis
*Generated: 2026-09-12 (announcement day) | Sources: ~120 across five tracks | Confidence: High on facts, Medium on strategy (same-day data)*
*Patched 2026-09-14 after the second research pass. Facts corrected in place are marked; the full second pass is in [06-since-announcement.md](06-since-announcement.md).*

Detail and citations live in the numbered files in this folder:
01 official facts · 02 community reaction · 03 competitive landscape · 04 data/API/addons · 05 business/legal/GTM · 06 everything since the announcement.

## Executive summary

WoW: Forever is the official "Classic+": a permanent, level-60-capped, vanilla-era game with new zones, dungeons, raids, a paid race, and a content roadmap into 2027. Beta opens **Sept 17** (cap 30), launch is **Nov 4**, first raids **Dec 9**, the Hardcore ruleset "winter." It is bundled with the normal WoW subscription.

**Correction (Sept 14): Forever has no realms.** Players pick one of four rulesets — Normal, PvP, Roleplaying, Hardcore — and each is one large ecosystem. Read every "realm" below as "ruleset."

The incumbents moved within hours. Wowhead, Icy Veins, Warcraft Tavern, and wowtbc.gg all have live Forever verticals; Wowhead even has an empty talent-calculator stub. A new entrant cannot win news, the item database, or head-term SEO. Every independent site that has ever broken through in the Classic ecosystem did so with a narrow, free tool that solved one real problem (Deathlog, RestedXP, Sixty Upgrades, ironforge.pro), often distributed via an in-game addon, and monetized on the side. No new site has broken through since 2022; Forever is the first genuine re-opening of the field since 2021, because it is a new game system rather than a new Classic phase.

The community's loudest asks right now are a talent calculator, a "which realm do I roll on" answer, an addon-compatibility answer, leveling routes for zones that have never existed, and a Hardcore deathlog. **Two of those moved on Sept 13–14:** Wowhead, Icy Veins and talentsforever.com all shipped talent calculators, and the realm question was answered out of existence by the ruleset system. The loudest grievance is the $29.99 Skyborne paywall. The loudest fear is "it'll die like SoD."

The biggest unknown is whether Forever allows addons at all. Nothing official says either way. It is testable the day beta opens and it gates every addon-fed tool (population, deathlog, prices, Legacy tracking).

## Key dates

| Date | Event | Why it matters |
|---|---|---|
| Sept 13, 10:00 PDT | Deep Dive panel; Hardcore panel 4:15 PDT | Happened: rulesets, Legacy, camping, transmog, items, Paladin; Hardcore confirmed with no date |
| Sept 17 | Beta opens (cap 30); live dev Q&A 10:00 PDT on Twitch/YouTube | Datamine client; test addon support |
| Oct 21 | Beta ends | |
| Oct 27 | Name reservation (up to 3 characters, pack owners) | Two-part names, unique per region; "which name" traffic, not "which realm" |
| Nov 4, 15:00 PST | Launch | Peak traffic; queues |
| Dec 9 | Barrow Deeps, Hyjal Summit, Onyxia | Raid guides, logs |
| Winter | Hardcore ruleset | Deathlog window |
| Spring / Summer 2027 | More raids, dungeons, zones | Retention content |

## What we know is true (high confidence)

- Permanent level 60, no flying, no level scaling, no expansions.
- New: Mount Hyjal, Zephras Isle, Riverglades, Shen'dralar; 9 dungeons; Barrow Deeps (10), Hyjal Summit (20); Darkspear Islands 15v15; 1000+ quests.
- Skyborne neutral race, $29.99 minimum; Warrior/Hunter/Rogue/Druid + faction-locked Shaman/Mage.
- Six new race/class pairs: Gnome Priest, Human Hunter, Dwarf Shaman, Orc Mage, Troll Warlock, Undead Paladin. **Correction (Sept 14):** new classes are explicitly *not* ruled out — Clay Stone said nothing is off the table.
- **Correction (Sept 14):** transmog is opt-*out*, on by default, disabled automatically if you pick Classic Mode at first login.
- Talent revamp (milestones at 11/16/21/31), reworked racials (two active, two passive each), Legacy account progression (65 points earnable, 16 spendable per character), Camping, HD/SD toggle, gamepad support.

## What nobody knows yet

Addon policy, Era transfers, the Hardcore date and rules, Honor numbers, full talent rank scaling, Battle.net API namespace for a realmless game, whether a living character can change ruleset, and the Merchant Favor crate economy named in a Legacy tooltip. (Realm list and types is closed: there are no realms.)

## Where a new site can actually win (ranked)

1. **Hardcore deathlog + public data layer.** The single best-attested playbook (2023 Deathlog/Deathmap went viral off streamer deaths). Hardcore ships "winter," so there is unusual lead time. Requires addon support.
2. **Legacy / alt progression tracker.** Matches Forever's alt-friendly design pillar; no incumbent tracks cross-character account progress. Buildable from beta data.
3. **Build planner for new race/class combos and revamped talents**, with shareable links and paywall transparency. Wowhead will own "talent calculator" but not "undead paladin build" or "skyborne druid."
4. **Leveling routes for the four new zones.** RestedXP's head start does not carry over. Only buildable during beta (7-week hard floor).
5. ~~**Realm picker**~~ / launch-day queue tracker. **Superseded Sept 13:** rulesets replace realms, so there is no realm to pick. The queue tracker survives; the picker does not. A two-part-name availability checker is the closer analogue.
6. **Exploration / quest-completion tracker** for the expanded world.
7. **Realm/community finder** (vibe, not raid logistics). True green field, but no historical proof players want a tool for it.

Not recommended as a lead: generic population tracking (5+ trackers already), news, item database, class guides (most contested).

## Constraints that shape the build

- **Domain:** no Blizzard marks. wowforever.* is taken anyway. Pick a brand-neutral name; use "Forever" descriptively in paths and titles.
- **Data source:** datamine the beta client via wago.tools / WoWDBDefs from Sept 17. Until then, talentsforever.com publishes its demo transcription at https://talentsforever.com/data.json under CC BY 4.0 (credit link required).
- **Rankings and character pages key on region + ruleset, not realm.** Bootstrap vanilla baseline from cmangos/classic-db (GPL). Battle.net API has no talent endpoint for any Classic variant and no Forever namespace yet; keep it configurable and non-load-bearing.
- **Addons cannot make network calls.** Any addon-fed feature needs a desktop companion (Tauri/Electron reading SavedVariables) or a paste-import UX.
- **API ToU:** never paywall API-backed features; 30-day retention; 36K calls/hour.
- **Monetization:** ads (Raptive/Mediavine once traffic qualifies), Patreon, premium cosmetics/features not tied to API data. No RMT.
- **Reddit:** participate for weeks before promoting. Addon distribution via CurseForge/Wago reaches players who never search.

## Audience sizing

Classic 2019 peaked at 1.17M concurrent Twitch viewers and tripled subs. SoD went from ~500K raiders to ~2K in two years. Project Epoch hit 25K concurrent with zero marketing before Blizzard shut it down. A subscription-bundled official Classic+ should comfortably exceed Epoch and plausibly approach 2019 in launch week. A small team's realistic ceiling is ~1M visits/month (Warcraft Tavern); Icy Veins sold for ~$8M on ~$2.4M revenue.

## Recommended shape

One brand-neutral destination that ships three narrow, free, shareable tools in sequence, each timed to a fixed date, with a companion addon as the distribution engine:

- **Before beta (Sept 17):** landing page, countdown, "everything we know," realm-picker shell, newsletter/Discord capture.
- **During beta (Sept 17–Oct 21):** build planner from datamined talents; new-zone leveling routes; addon prototype (if addons work); Legacy tracker.
- **Launch window (Oct 27–Nov 11):** realm list and queue tracker; day-one leveling guides; addon on CurseForge.
- **Dec 9:** raid guides and comp tools timed to the unlock.
- **Winter:** deathlog and deathmap flipped on the day Hardcore ships.

## Methodology

Five parallel research agents, ~140 tool calls, ~120 unique sources. Sub-questions: official facts; community reaction and unmet needs; competitive landscape and precedent; data/API/addon ecosystem; audience, monetization, legal, GTM. Same-day limitations: Reddit megathreads unindexed, r/wow blocked, no streamer reaction videos yet, ironforge.pro and NexusHub status unverified.
