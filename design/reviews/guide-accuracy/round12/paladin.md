# Guide accuracy round 12: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr12/audit (main a8b147a6f: guide rounds 1 to 11, sim round 15, leveling-order, guide-header-build). Client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go; source read from gr10/engine). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, the card notes in `data/curated/apl/paladin-*.json`, `data/curated/rank-choices.json`, `data/builds/1.60.1.70755/addon-builds.json`. Before starting I read the paladin reports for rounds 10 and 11 with their fix logs, and the round 1 to 9 decisions they carry.

## What changed since round 11's audit (git diff d54de48aa..HEAD)

- **Guides:** only the round 11 fixes (91da94f1e): retribution.md:133 (R11-1) and holy.md:117 (the Flash of Light at 20 SHOULD).
- **guide-header-build (3bd24d3a2, 3dff67836):** on a guide with its own raid build, the header card is now titled "Raid build" and loads the raid code; the rotation and cooldown cards read the raid build's talents and say "For the raid build."; key-talent chips show both builds' ranks where they differ. For paladin this touches Retribution only (Holy and Protection have no raidBuild).
- **Card notes, rank-choices.json, addon-builds.json:** unchanged for paladin.

## Builds decoded (unchanged, all legal on 70755)

- **Retribution** `build` 050250002/502/0523330100130135 (14/7/30), `raidBuild` 050050002/503/05233301001301351 (12/8/31). The header card now shows the raid build. retribution.md:58 names both builds and the exact move (Reverence 2 to Precision 3 and Unsealed Wrath), so the prose agrees with the new header card.
- **Protection** `build` 05005/15302102513501031/5003 (10/33/8). protection.md:66 to :82 match.
- **Holy** `build` 50320215012101251/50322/0503 (31/12/8). Rows 1 to 5 hold 24 points without Divine Precision and 25 with it, so "Divine Precision is the point that opens Holy Power's row" (holy.md:67) holds.

## Header card check (round 12 focus)

- retribution.md: no sentence says "the build above" or speaks of the header build as the leveling one. :58 names the leveling build's split and the raid move; the bullets give the same point counts in both builds except Unsealed Wrath (1 point, raid only, named as "the raid capstone"). Reverence and Precision, the two talents whose ranks differ, sit in the prose paragraph at :72 and are described without a rank. The rotation text at :78 to :99 covers both leveling and raid Seals, so it does not contradict "For the raid build." No contradiction.
- holy.md and protection.md: unchanged behaviour (no raidBuild).

## Round 11 edits re-checked

- **holy.md:117:** correct. Flash of Light is learned at 20 (trainables), and Spiritual Focus covers it.
- **retribution.md:133 (R11-1):** wrong. See PALADIN-R12-1.

## Checked and correct (no finding)

- Every talent bullet on all three spec pages against talents/paladin.json descriptions (Holy Conduit, Vengeance "non-periodic critical strike", Repentance 6 s / 60 s / Humanoids, Holy Shield 4 charges and 20% threat, Redoubt 10 s or 5 blocks, Shieldsworn, Boundless Light, Holy Power 175%, Divine Precision 4% per point).
- Divine Shield rank 2 (1020): 12 s, 50% less damage dealt, 5 min, Forbearance; the Protection card note at :108 matches.
- Voice of Truth (1310897): 6 s, 3 min (holy.md:89).
- Seal costs behind the Retribution twist note ("about 190 mana"): Seal of Command rank 5 210 and Seal of Righteousness rank 8 200 mana, less Benediction's 10%.
- Righteous Fury (25780): 30% base mana, 30 min (protection.md:86).
- Light's Vigil "almost as much as ten top-rank Flash of Lights" (holy.md:40): 1,340 against 10 x 140.
- rank-choices.json against the prose: Holy's tank Flash of Light is rank 1 only in the 26 to 39 and 50 to 59 bands and rank 5 at 60, so holy.md:78, :95 and :127, which all say "at 60", hold. Retribution Consecration rank 1 below 60 and rank 4 bare at 60 match :95.
- Card notes: all paladin notes re-read; nothing false. The Holy Strike "Twenty mana" note on Retribution (18 with Benediction) stays the carried SHOULD.
- Source links: all nine distinct URLs returned 200 with the right content types (fetched 2026-11-04), and every label matches its page.
- Voice: no em or en dashes, no site or sim talk, no stat ranking in the Stat priority sections, no nightly numbers.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R12-1 | CONFIRMED (round 11 edit) | leveling / rotation | retribution.md:133 | "Through your 30s, rank 1 at 135 mana does about as much for its mana as rank 2, so keep it. From 40, on a pack, cast the newest rank you have: each new rank hits every mob far harder for its mana, and rank 1 also loses its spell damage bonus as you outlevel it." | R11-1 compared base damage per mana only (0.36 / 0.34 / 0.43 / 0.51) and left out spell power, which is the same 0.042 a tick on every rank (engine consecration.go, `consecrationFirstTargetsCoefficient`, applied per target per tick) before the level-gap rule (core.LevelPenalty: 17 levels of grace, then 5% a level). Rank 1 carries that full share for 135 mana, so spell power swings the per-mana comparison its way. Per target per cast, at the Retribution gear the page's own tables use (bis/paladin-retribution.json, Human bare): band 30 (89 spell power from items, no Champion of the Light yet) rank 1 0.58 damage per mana against rank 2 0.47, so rank 1 is about 20% ahead, not "about as much". Band 40 (99 spell power + 60% of 137 Intellect = about 181): rank 1 (15% off its share) 100 damage for 135 mana = 0.74 a mana, rank 3 197 for 320 = 0.62. Rank 1 is about 20% ahead for its mana at 40, the reverse of "far harder for its mana". Even in strength gear (about 47 spell power) rank 3 is only about 4% ahead. The crossover is around 46 (rank 1 40% to 50% off its share). Band 50 (about 246 spell power, rank 1 65% off): rank 4 0.71 against rank 1 0.57, so from 50 the newest rank does win per mana. What is true from 40 is damage per global: rank 3 hits each mob about twice as hard as rank 1 (197 against 100; 152 against 62 in strength gear). | "From level 20 Consecration is your AoE: two or three mobs on you is the time for it. Rank 1 at 135 mana carries as much spell damage as any rank for a fraction of the cost, so until your mid 40s it gives you the most damage for your mana. From 40, when a pack is on you and you have the mana, cast the newest rank instead: it hits every mob about twice as hard for the same global. From your late 40s the newest rank wins on mana too, because rank 1 has lost most of its spell damage bonus by then. On a single target, rank 1 stays the rank to use until 60, for the reason in the Consecration paragraph of the rotation." |

## SHOULDs (wording, not findings)

- Carried: the Retribution Holy Strike card note's "Twenty mana" against Benediction's 18. The card now explicitly describes the raid build, which has Benediction 5/5, so this is a little more visible than before.
- Carried: retribution.md:111 "Human is ahead solo" (measured on the default target; Beasts unmeasured).

## For the component and data lanes (not guide findings)

- Carried: the Holy APL's top-level engineering note names rank 1 Holy Shock and rank 1 Flash of Light and says the engine has "no downrank penalty"; the Protection APL's top-level note says Righteous Fury is "+60 percent threat on Holy damage", while the 70755 tooltip (25780) says 33% on all threat and a further 90% on Holy. Neither note renders.
- Carried: trainables lists Holy Light 9, Blessing of Might 7 and Blessing of Wisdom 6 at 60 as skill_line_ability rows (AQ librams in 1.12).

## Decisions

- **R12-1:** decided from client damage rows, the engine's spell power coefficient and level-gap rule, and the gear in the page's own Retribution tables. No owner question. It corrects a round 11 edit that compared base damage only.
- **Races:** unchanged. Retribution Human (Undead), Holy Human (Undead), Protection Dwarf (Undead). No paid race named, and no race line contradicts a measured result.

## Sims requested

1. Confirm R12-1 (not needed to decide it): Retribution leveling Consecration on 3 targets, 20 s pack fight, bare gear from bis/paladin-retribution.json, rank 1 against the newest rank learned, at levels 40, 45 and 50. Report damage per mana and damage per cast.
2. Carried: the Crusader opener on a 15 s single mob at bands 20, 30 and 40 (retribution.md:131).
3. Carried: Retribution Human against Dwarf on a Beast target (retribution.md:111); Protection threat by ability on 1 and 4 targets.
4. Carried: race re-runs for Retribution on the round 15 APL (Crusader opener), raid 60 and bare 60.
5. PROVISIONAL (final BiS re-rank): none. No prose names a BiS piece or ranks a stat.

## Sim findings

None new.

## Sources

- 70755 client data, data/builds/1.60.1.70755: spelltips/paladin.json (26573, 20922, 1020, 642, 20375, 20920, 20293, 1310897, 20066, 25780), talents/paladin.json, trainables/paladin.json, levels.json (paladin Intellect at 30, 40, 50), items/paladin.json, bis/paladin-retribution.json (bands 30, 40, 50 Human bare), addon-builds.json.
- data/curated/apl/paladin-{holy,protection,retribution}.json; data/curated/rank-choices.json.
- Engine 9f7f1f6e4: sim/paladin/consecration.go (rank table, 0.042 coefficient on the first four targets), sim/core/spell.go:822 (LevelPenalty).
- Commits 91da94f1e, 3bd24d3a2, 3dff67836 (web/src/lib/guides/guide-rotation.ts, key-talents.ts, GuideActionRail.astro).
- Links fetched 2026-11-04, all 200: Blizzard Deep Dive recap https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap ; Blizzard What's Next recap https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap ; Wowhead, Undead Paladins https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823 ; Power Up Gaming https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ ; Warcraft Tavern Retribution rotation https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/ ; foreversixty.gg planner, /sim/weights, races.json, combos.json.

## Counts

| status | count |
|---|---|
| CONFIRMED | 1 (R12-1, a round 11 edit) |
| LIKELY | 0 |
| NEEDS-SIM | 0 new (1 confirming, 3 carried) |
| NEEDS-OWNER | 0 |

Header card: 0 contradictions. Levels 1 to 19: 0 errors. Source links: 0 errors. Card notes: 0 errors. Round 11 edits: 1 error (retribution.md:133).

## Fix log (round 12)

- PALADIN-R12-1: fixed, retribution.md:133. Verified against engine consecration.go (rank table, 0.042 coefficient, 8 ticks) and core LevelPenalty (17-level grace, 5% a level): per target per cast, band 30 rank 1 0.58/mana vs rank 2 0.47; band 40 rank 1 0.74 vs rank 3 0.62 (197 vs 100 damage per cast); crossover about 46; band 50 rank 4 0.71 vs rank 1 0.57. Applied the report's text (rank 1 best per mana until the mid 40s, newest rank about twice the damage per global from 40, newest rank wins on mana from the late 40s). `updated` already 2026-11-04.
- SHOULDs (Holy Strike "Twenty mana", Human solo line): left, not in this lane's scope.
