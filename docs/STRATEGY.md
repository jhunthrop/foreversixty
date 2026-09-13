# Forever Sixty — strategy

Updated 2026-09-13. This is the one document that says what the site is for, what it replaces,
in what order, and why. Phase specs under `docs/superpowers/specs/` argue from it; the research
under `research/` is its evidence.

## The goal

Forever Sixty is the one place a World of Warcraft: Forever player needs. Today a Classic
player's tools are scattered across a dozen sites and each of them was built for a different game:
Wowhead and Icy Veins for facts and guides, Sixty Upgrades for gear, Raidbots for simulation,
Warcraft Logs for raid analysis, RestedXP for leveling, ironforge.pro for population, Deathlog for
Hardcore. Forever is a new game with new talents, races, zones, dungeons, and raids, released
Nov 4, 2026, and every one of those tools has to rebuild its Forever coverage from zero. That is
the opening: build them once, together, around one shareable data model, for this game only.

## The thesis

1. **Tools first, content second.** Every independent site that broke through in the Classic
   ecosystem did it with one narrow free tool that solved a real problem, then grew. News,
   databases, and head-term guides are where the incumbents are strongest; we do not lead there.
2. **One data model, many tools.** A build (class, race, point order, gear) is the unit that
   flows between the planner, the addon, the simulator, and the logs. A character (builds,
   gear, kills, deaths) is the unit that flows between the tracker, the logs, and Hardcore.
   Because one site owns both, each tool makes the others better, which no incumbent can copy
   without owning all of them.
3. **The addon is the distribution engine.** Players who never search still install addons.
   The addon carries builds in and out of the game, follows a build in-game, and later uploads
   logs through the companion. It is the reason the site's name is on every player's screen.
4. **Every fact is dated and sourced.** The site is a reference, not a pitch. Single-source
   claims say so. This is the trust that the incumbents' ad-driven pages have lost, and it
   costs nothing but discipline.
5. **Free and fast at scale.** Static pages served from the edge cost nothing at a million
   visits, load in under two seconds on a phone, and never show a cookie banner. The first
   feature with a real bill is logs, and it arrives only when raids exist to log.

## What we replace, and with what

| Incumbent | What players use it for | Forever Sixty's answer | Phase |
|---|---|---|---|
| Wowhead, Icy Veins | Talent calculator, item and spell facts, class guides | Shareable build planner; datamined reference pages that are dated and sourced; community guides through issue templates, later accounts | 1, then 2 to 4 |
| Sixty Upgrades | Gear planning with stat weights | Gear picker on every build with curated, sourced weights | 1 to 2 |
| RestedXP | In-game leveling guidance | Addon "follow": the next talent point in-game from any shared build | 2 |
| Raidbots | Simulation and stat weights | Curated weights first; a browser-side simulator forked from WoWSims once the Forever trees settle | 2, then 4+ |
| Warcraft Logs | Upload, parse, and analyze raid logs; rankings | Companion uploader, parser, per-fight tables and share links, then rankings and guild pages | 3, then 4 |
| Deathlog, Deathmap | Hardcore deaths | Deathlog and deathmap fed by the addon, live the day Hardcore ships | 4 |
| ironforge.pro | Population and faction balance | Later, from the site's own addon and log data rather than scraping | 4+ |
| Discord, spreadsheets | Guild coordination | Guild pages on top of logs and characters | 4 |

Deliberately not pursued: news, patch drama, an item database race against Wowhead, ads that
slow the page, anything that needs Blizzard's marks in the brand.

## Sequence

Each phase ships on a date the game sets, and each one starts on the previous one's data.

| Phase | Window | Ships | Gate |
|---|---|---|---|
| 0 Foundation | done Sept 13 | Site, research pages, data pipeline, API, hosting | live |
| 1 Build planner | Sept 13 to Oct 10 | Planner with talents, point order, gear; shareable short links with preview cards; races and classes page | beta client exports talents Sept 17 |
| 2 Addon | Oct 10 to Nov 4 | Addon with import, follow, gear upgrades; stat weights; gear-release tail | beta loads addons |
| 3 Accounts, companion, logs MVP | Nov 4 to Dec 9 | Accounts, character tracker, desktop companion, log upload and parse, per-fight tables, share links | first raids Dec 9 |
| 4 Logs at scale, Hardcore | Dec 9 onward | Rankings, guild pages; deathlog and deathmap; simulator start | Hardcore realms |
| 5 Simulator, population | 2027 | WoWSims fork per class; population from own data | trees stable |

## How each piece connects

```
data pipeline (client builds) ─▶ planner ─▶ build record ─▶ addon code ─▶ in-game follow
                                     ▲            │
        addon export string ─────────┘            ├─▶ gear score (weights) ─▶ simulator input
                                                  │
accounts ─▶ characters ─▶ companion ─▶ logs ─▶ fights ─▶ rankings ─▶ guild pages
                    └──────────────────────────▶ deathlog (Hardcore)
```

## Principles for every phase

- Ship on the game's dates, not ours. A tool that lands after its moment is worth little.
- One narrow tool at a time, complete and reviewed, then the next. No half-built surfaces.
- Everything shareable: a build, a log, a death, a character has a short link that unfurls.
- Data is generated, deterministic, and keyed by client build id. A new build is one pipeline
  run, never a hand edit.
- The design system is binding: dark, high contrast, 44 px targets, class and rarity colors,
  reference voice, no ornament, no marketing.
- Costs stay near zero until logs. Logs are budgeted separately when Phase 3 is designed.

## Risks

- **Beta data and addon support** are unknown until Sept 17. Both phases 1 and 2 have
  fallbacks: hand-curated trees, and building against Classic Era until launch.
- **The game fails.** "It'll die like SoD" is the community's loudest fear. The mitigation is
  cost: a static site and a scale-to-zero API cost nothing while waiting.
- **Blizzard policy.** No Blizzard marks in the brand, API terms respected, no paywalls on
  API-backed features, no RMT. The addon and companion touch only files the game writes for
  addons and logs.
- **Logs at scale** are the first real infrastructure and cost problem. They are designed in
  Phase 3 with a budget, not assumed.
- **Simulation accuracy** depends on modeling every spell and talent, which is community work
  measured in months. It starts one class at a time and only after the trees stop moving.

## Success, phase by phase

- Phase 1: a build made on the site unfurls on Discord; the planner runs on the Forever trees
  within two days of the beta export.
- Phase 2: the addon is on CurseForge before name reservation; a shared build is followed
  in-game.
- Phase 3: a raid log uploads, parses, and shares before the first raid week ends.
- Phase 4: rankings exist for every Dec 9 encounter; the deathlog is live the day Hardcore ships.
- Overall: the tools a Forever player opens on a raid night are all on one domain.
