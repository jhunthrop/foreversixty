# Guide accuracy round 6: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr6/audit (land/guides-r5, head 367cbadb7). Client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, plus the rendered card notes in `data/curated/apl/paladin-*.json`. Before starting I read the paladin reports from rounds 1 to 5 (both round 5 fix logs), round4/paladin-levels.md, the race-sims files for rounds 2 to 5, and ledger.md.

## Builds decoded (talents[] order in talents/paladin.json; unchanged since round 1, all legal on 70755)

- **Retribution `build`** `050250002/502/0523330100130135` (14/7/30). Holy: Divine Strength 5, Reverence 2, Improved Seals 5, Holy Conduit 2. Protection: Divine Intellect 5, Precision 2. Retribution: Benediction 5, Improved Judgement 2, Conviction 3, Sanctified Judgement 3, Vindication 3, Seal of Command 1, Sacred Arbiter 1, Two-Handed Weapon Specialization 3, Repentance 1, Vengeance 3, Champion of the Light 5.
- **Retribution `raidBuild`** `050050002/503/05233301001301351` (12/8/31): minus Reverence 2, plus Precision 3 and Unsealed Wrath 1. retribution.md:58 matches.
- **Protection `build`** `05005/15302102513501031/5003` (10/33/8). Holy: Divine Strength 5, Improved Seals 5. Protection: Divine Intellect 1, Redoubt 5, Precision 3, Sacred Duty 2, Templar's Bulwark 1, Shield Specialization 2, Toughness 5, Swift Judgement 1, Improved Righteous Fury 3, Anticipation 5, Holy Shield 1, Iron Creed 3, Shieldsworn 1. Retribution: Deflection 5, Conviction 3. Left out as protection.md:81 says: Guardian's Favor, Improved Hammer of Justice, Stoicism, Reckoning.
- **Holy `build`** `50320215012101251/50322/0503` (31/12/8). Holy: Spiritual Focus 5, Healing Light 3, Reverence 2, Purifying Power 2, Divine Favor 1, Illumination 5, Divine Precision 1, Infusion of Light 2, Oath of Light 1, Holy Shock 1, Boundless Light 2, Holy Power 5, Light's Vigil 1. Protection: Divine Intellect 5, Precision 3, Guardian's Favor 2, Sacred Duty 2. Retribution: Benediction 5, Conviction 3.

## Leveling prose against the shipped point order (addon-builds.json, index 0 = level 10, decoded again)

- **Retribution:** Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication 23-25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28-30, Repentance 31, Vengeance 32-34, Champion of the Light 35-39, then Holy and Protection (raid order: Unsealed Wrath at 40). retribution.md:74 and :125 match.
- **Protection:** Divine Intellect 10, Redoubt 11-15, Precision 16-18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, Shield Specialization 22-23, Toughness 24-28, Swift Judgement 29, Holy Shield 30, Improved Righteous Fury 31-33, Anticipation 34-38, Iron Creed 39, Shieldsworn 40, Iron Creed 41-42. protection.md:132 and :138 match.
- **Holy:** Spiritual Focus 10-14, Healing Light 15-17, Reverence 18-19, Purifying Power 20-21, Divine Favor 22, Illumination 23-24, Oath of Light 25 (the 16th point, the first its row allows), Illumination 26-28, Divine Precision 29, Infusion of Light 30-31, Holy Shock 32, Boundless Light 33-34, Holy Power 35-39, Light's Vigil 40. holy.md:117 and :123 match.

## Stat paragraphs against the page's table (bis band 60, raid preset, the first recommended race's faction, ordered as GuideStatTable/railStatRows order it)

- **Retribution** (alliance raid): Strength 1.00, crit 0.93, spell power 0.73, Intellect 0.71, Agility 0.70, attack power 0.41, then melee haste per 1%, then Hit "Not significant" (hit_to_cap baseline 12.9, specials 0, so "reached"). retribution.md:103 gives the same order with stable reasons. The one problem is a position word, R6-1.
- **Protection** (alliance raid; Expertise is not in statPriority, so the table does not show it): Defense 14.72, Stamina 12.65, Dodge 7.81, Agility 7.07, block value 6.32, Parry 5.88, Strength, crit, armor, attack power, then Block and Hit "Not significant". protection.md:110 gives the same order. See R6-5 for one label.
- **Holy** (alliance raid; spell haste is not in statPriority): MP5 1.65, healing power 1.01, Intellect 1.00, crit 1.00, Spirit 0.74. holy.md:95 matches.
- No stat paragraph quotes a number. No gear prose names a piece for a nightly reason (none of the four pages names a gear piece).

## Levels 1 to 19, then Levels 20 to 60 (reading order)

The page renders the 1 to 19 text, then the "Levels 20 to 60" heading, the strip, then the 20 to 60 text ([spec].astro:368-372). Every trainer level, talent rank and gear line in the three 1 to 19 blocks checks out against trainables/paladin.json, spelltips/paladin.json and talents/paladin.json (Seal of Righteousness, Holy Light and Devotion Aura 1; Judgement and Blessing of Might 4; Holy Strike and Holy Light rank 2 at 6; Hammer of Justice and Purify 8; Lay on Hands, Righteous Fury and Seal of Fury 10; Redemption 12; Blessing of Wisdom and Holy Light rank 3 at 14; Divine Protection rank 2 at 18; Battleworn Hammer 2361 a two-handed mace). Retribution reads right in order. Protection and Holy read awkwardly in order, see R6-4.

## Round 5 fixes and decisions re-checked

- R5-1 (holy.md:56 and :95, Illumination's 50% casting regeneration): correct.
- R5-2 (retribution.md:141): correct, and it agrees with :72 and :103.
- R5-3 (retribution.md:125, "every button you fight with"): correct.
- R5-4 (holy.md:123, Divine Precision hit and silence reduction): correct against 1310904 (rank 1 is 4% and 10%; the build takes one point).
- Voice of Truth paragraph (holy.md:89): correct (trainer 20, 180000 ms, "immunity to Silence and Interrupt effects for 6 sec").
- Review pass, retribution.md:103: the order is right, but "the raid gear above" is wrong (R6-1).
- Review pass, protection.md:110 and holy.md:95: correct order.
- **Round 5 decision "Dwarf even in a raid" (retribution.md:111): wrong, R6-2.** round5/race-sims.md itself shows Dwarf ahead on the raid build beyond error.

## Checked and correct (no finding)

Every other line of the four pages and every rendered card note was checked again against talents/paladin.json, trainables/paladin.json, spelltips/paladin.json, races.json, combos.json and the APLs. Highlights:

- Divine Shield rank 2: 12 s, "reduces all damage you deal by 50%" (1020). Divine Protection rank 2: 8 s, no attacking (5573). Blessing of Protection: 300 s base, rank 3 10 s, 3 min with Guardian's Favor (holy.md:47).
- Lay on Hands rank 3 (10310): heals for the paladin's maximum health, 550 mana to the target, drains all of the paladin's mana, regeneration keeps running (holy.md:91, protection.md:106).
- Light's Vigil rank 3 (1311595): 1340 mana, 1.5 s cast, next Holy Shock heals the target's party (holy.md:40, :63, :85).
- Righteous Fury 25780: 30% of base mana, 30 min, 33% threat plus 90% on Holy (protection.md:85).
- Hammer of the Righteous 407632: 6% of base mana, 6 s, target plus up to 3 more (protection.md:94, :104). Consecration rank 5 (20924): 565 mana, 8 s cooldown.
- Hammer of Wrath 24239: 1.0 s cast, 6 s, 30 yd, below 20%; Champion of the Light takes 1.0 s off (retribution.md:42, :67, :84).
- Seal of the Crusader 20308: 306 attack power, 40% faster swings, Judgement 40 s refreshed by melee (retribution.md:93).
- Concentration Aura 19746 at 22 (holy.md:36, :71). Spiritual Focus also covers Light's Vigil (talent text), see R6-3 for the wording.
- Holy Power: 50% more crit bonus, so 175% (holy.md:62). Purifying Power: Cleanse and Purify 20% cheaper, plus Exorcism and Holy Wrath cooldowns (holy.md:67 names only the Cleanse half, which is true).
- Index: Human, Dwarf and Undead only (combos.json class 2); milestone rows at 11/16/21/31 in all three trees; Holy Shield at tier 4; no Blessing of Light, Salvation or Sanctuary in trainables.
- Prepull timings: Retribution Seal at -1.5 s, Protection Seal of Fury -3.5 s and Holy Shield -2 s, Holy Oath of Light -1.5 s (retribution.md:78, protection.md:87, holy.md:71).
- Racials (races.json 70755 text): Ruthless, The Human Spirit, Will to Survive, Stoneform 10% for 8 s, Artisan's Precision 1% hit, Will of the Forsaken without immunity, Cannibalize health and mana, Touch of the Grave up to 5% of maximum health once every 15 s.
- No em or en dashes, no site or sim talk, no quoted nightly numbers.

## Source links (fetched 2026-11-04; every one returns 200)

| page | label | URL | says what the label says? |
|---|---|---|---|
| all four | Blizzard, Deep Dive panel recap | news.blizzard.com/en-us/article/24303313 | yes |
| all four | Blizzard, What's Next panel recap | news.blizzard.com/en-gb/article/24303862 | yes (Hyjal Summit 20-player, the Barrow Deeps 10-player) |
| index | Wowhead, Undead Paladins coming to Forever | wowhead.com/forever/news/...-382823 | yes |
| index | Forever Sixty, which races can be paladins | foreversixty.gg/data/1.60.1.70755/combos.json | yes (application/json) |
| index | Power Up Gaming, Forever talent changes vs Classic | powerupgaming.co.uk/2026/09/14/... | yes |
| ret, holy, prot | Forever Sixty, paladin talents | foreversixty.gg/planner?class=paladin | yes (talent-tree sources open the planner) |
| ret | Forever Sixty, what each stat is worth | foreversixty.gg/sim/weights | yes (damage spec only) |
| ret | Warcraft Tavern, Classic Retribution Paladin rotation | warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/ | yes |
| ret, holy, prot | Forever Sixty, races | foreversixty.gg/data/1.60.1.70755/races.json | yes (application/json) |

Source-link rule (ledger, round 4): compliant on all four pages.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R6-1 | CONFIRMED | stats (round 5 review-pass edit) | retribution.md:103 | "Hit matters only up to the cap, and the raid gear above already reaches it, so more hit past that adds nothing." | The best in slot table renders after the stat prose, not above it: web/src/pages/guides/[class]/[spec].astro:337-353 places StatPriorityPills, GuideStatTable and the stat prose, then the gear heading and GuideBisTable. The markdown order is the same (## Stat priority at :101, ## Best in slot at :105). | "Hit matters only up to the cap, and the raid gear below already reaches it, so more hit past that adds nothing." |
| PALADIN-R6-2 | CONFIRMED | consumable/race (round 5 decision) | retribution.md:111 | "**Dwarf** comes out even in a raid, where Artisan's Precision, 1% hit, is worth as much as Ruthless, and Stoneform ..." | round5/race-sims.md, raid build on the re-ranked gear, 10,000 iterations: Dwarf 783.39 and 782.97 against Human 780.79 and 780.26, errors 0.52. Dwarf leads by about 2.7 on both seeds, beyond error, so "even" is false for the build you raid with (only the leveling build in a raid is even). The stated reason also contradicts the same page: at raid 60 the table reads Hit "Not significant" (bis/paladin-retribution.json band 60 raid alliance, hit scale factor 0.00, insignificant, hit_to_cap reached), and :103 says hit past the cap "adds nothing", so Artisan's Precision cannot be what earns Dwarf its raid edge. | Replace the Dwarf sentence with: "**Dwarf** edges a fraction ahead in a raid and falls behind solo, in dungeons and while leveling. Stoneform takes 10% off physical damage for 8 seconds on a hard pull, and Artisan's Precision, 1% hit, helps until your gear reaches the hit cap." Keep the Human and Undead sentences. |
| PALADIN-R6-3 | LIKELY | talent | holy.md:53 | "**Spiritual Focus** (5 points): Flash of Light, Holy Light and Light's Vigil keep 70% of their casts from being pushed back." | talents/paladin.json Spiritual Focus 5/5: "a 70% chance to not lose casting time when you take damage". It is a roll on every hit, not 70% of casts: a cast that takes two hits gets through clean only about half the time. holy.md:117 words the same talent correctly ("70% of the hits you take do not push back"), so the page describes one talent two ways. | "**Spiritual Focus** (5 points): every hit you take while casting Flash of Light, Holy Light or Light's Vigil has a 70% chance not to push the cast back." |
| PALADIN-R6-4 | LIKELY | leveling (reading order) | protection.md:132 with :138; holy.md:117 with :123 | 1 to 19: "If you level as Protection, your first point at 10 is Divine Intellect ..." then, after the strip, 20 to 60: "Level as Retribution and tank dungeons in it ..." (Holy: "If you level as Holy, your first five points ..." then "Level as Retribution and heal dungeons on the way") | Since round 5 the 1 to 19 text comes first and the "Level as Retribution" advice only appears after the strip ([spec].astro:368-372). A reader at 10 meets "If you level as Protection" (or Holy) before the page has said that the usual way is Retribution, and has spent ten points by the time they read it. Nothing is false, but the two blocks no longer read right in this order. | protection.md:132, start the paragraph with: "Most tanks level as Retribution and pick these up when they respec; the Retribution guide has those points. If you level as Protection, ..." holy.md:117, start it with: "Most Holy paladins level as Retribution and respec for raids; the Retribution guide has those points. If you level as Holy, ..." Then :138 and :123 can keep their first sentences as they are. |
| PALADIN-R6-5 | LIKELY | stats (wording) | protection.md:110 | "More block chance is no gain at 60" | The table beside it reads Block "Not significant" (bis/paladin-protection.json band 60 raid, block 2.36, flagged insignificant), not "No gain". The table uses "No gain" for a stat below zero beyond its error (weight-reading.ts:33-37, rail-stats.ts), so the prose borrows the table's other label. | "More block chance adds nothing you can measure at 60, but keep the shield: ..." |
| PALADIN-R6-6 | LIKELY | spell | holy.md:81 (and the Holy Light card note, paladin-holy.json:135) | "Holy Light rank 8 is the top rank a trainer teaches; rank 9 is an Ahn'Qiraj book." | Forever's announced raids are Hyjal Summit and the Barrow Deeps (What's Next recap), plus Onyxia; there is no Ahn'Qiraj. Libram: Holy Light IX (items.json 21290) has no known source in the 70755 loot data (loot.json source "world_drop:unknown", carried over from the Classic database). A Forever healer reads "an Ahn'Qiraj book" as somewhere to go and get it. | holy.md:81: "Holy Light rank 8 is the top rank a trainer teaches, and the rank 9 book came from Ahn'Qiraj, which Forever does not have, so rank 8 is your top rank." The card note at paladin-holy.json:135 the same way: "Rank 8 is your top rank: the trainer stops there, and the rank 9 book came from Ahn'Qiraj, which Forever does not have." |

## Decisions

- **Retribution race: Human stays the Alliance pick (R6-2 changes only the Dwarf sentence).** The measurements: raid build in a raid, Dwarf about 0.35% ahead; leveling build in a raid, even; solo at 60, Human 283.1 against Dwarf 280.0 (about 1.1% ahead). Human's lead solo, in dungeons and through leveling is three times the size of Dwarf's raid lead, and Will to Survive is a stun break Dwarf has no answer to. Frontmatter stays `[human, undead]`, with `human` in both build strings.
- **R6-6:** decided on Blizzard's raid list and the 70755 loot data. No owner question.
- **Holy and Protection races:** unchanged. Holy is Human (round 4 sims, Human 676.8 against Undead 662.2 and Dwarf 661.0), and Protection is Dwarf (a tie with Human; Stoneform tips it).

## Sims requested

1. **Retribution, raid band 60, raidBuild, published gear, 10,000 iterations, seeds 1 and 50001:** Dwarf as is, against Dwarf with Artisan's Precision removed, against Human with Ruthless removed. This finds where Dwarf's raid edge comes from (base stats or the racial), so the race line can give the real reason. Until it reports, R6-2's text names no reason for the raid edge.
2. Carried, not re-filed: the Retribution Crusader opener (round 3 R3-6), and the Retribution leveling Consecration rank under the level-gap curve (retribution.md:95 and :133 say "from 40 move up", while rank-choices.json still casts rank 1 from 30 to 59).

## Sim findings

None new. Lay on Hands drain, Divine Shield's 50% damage cut, Illumination's casting regeneration, Holy Power's 75% bonus and Seal of the Crusader's refresh all agree with the 70755 client text. Open question for the sim lane, from sim request 1: with hit at the cap and Hit not significant at raid 60, the engine still puts Dwarf ahead of Human beyond error. If neither racial explains it, check the paladin base stats by race.

## Sources

- 70755 client data, data/builds/1.60.1.70755:
  - talents/paladin.json, trainables/paladin.json, spelltips/paladin.json, addon-builds.json
  - races.json, combos.json, items.json (2361, 21290), loot.json (source "world_drop:unknown")
  - bis/paladin-{retribution,protection,holy}.json (band 60 weights, hit_to_cap)
- data/curated: apl/paladin-*.json, rank-choices.json, book-ranks.json.
- Site code: web/src/pages/guides/[class]/[spec].astro (section order), web/src/components/guides/GuideStatTable.astro, web/src/lib/guides/rail-stats.ts, web/src/lib/bis/weight-reading.ts, web/src/lib/guides/class-dps.ts, web/src/lib/bis/presets.ts.
- Engine d1d9f2775: sim/core/racials.go (Ruthless, Artisan's Precision, Stoneform never auto-cast).
- design/reviews/guide-accuracy/round5/race-sims.md, round4/race-sims.md, round2/race-sims.md.
- Blizzard, World of Warcraft Forever What's Next panel recap (Hyjal Summit and the Barrow Deeps are the raids; no Ahn'Qiraj): https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
- Blizzard, World of Warcraft Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Undead Paladins Coming to World of Warcraft Forever: https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823
- Power Up Gaming, New WoW Forever Talents: Changes vs Classic Guide: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Warcraft Tavern, PvE Retribution Paladin Rotations & Cooldowns (Classic): https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/

## Counts

| status | count |
|---|---|
| CONFIRMED | 2 (R6-1, R6-2; both are round 5 edits or decisions) |
| LIKELY | 4 (R6-3 to R6-6) |
| NEEDS-SIM | 0 separate (one decomposition sim requested above, two carried) |
| NEEDS-OWNER | 0 |

The Levels 1 to 19 text has 0 confirmed errors (one LIKELY ordering, R6-4). The source links have 0 errors. The stat paragraphs give the tables' current order on all three specs.

## Fix log (round 6)

- **New stat rule, all three spec pages.** Stat priority prose rewritten with no ranking words: each stat gets its lasting reason, hit gets the cap rule (matters up to the cap, nothing past it), and the paragraph points to the table above for the order. retribution.md:103, protection.md:110, holy.md:95. Holy's "five second rule" became "once you have gone five seconds without a cast" (the test reads "second" as a ranking word). No other section on the four pages ranks stats or gear by nightly value; protection.md:81 "worth less than Conviction" compares talents, not stats, and stays. index.md has no stat prose and is unchanged.
- **PALADIN-R6-1** fixed by the new rule: the gear reference is gone, retribution.md:103.
- **PALADIN-R6-2** fixed per the integrator, retribution.md:111: Dwarf is slightly ahead on the raid build in a raid, Human is ahead solo and while leveling, which is why Human stays the pick. No reason given for the raid edge, no numbers. Stoneform kept as its own sentence.
- **PALADIN-R6-3** fixed, holy.md:53 (a 70% roll on every hit, matching the talent text and holy.md:117).
- **PALADIN-R6-4** fixed: protection.md:132 and holy.md:117 now open by saying most players level as Retribution and pointing to that guide; :138 and :123 unchanged.
- **PALADIN-R6-5** fixed, protection.md:110: "adds nothing you can measure at 60".
- **PALADIN-R6-6** fixed after checking the loot data: Libram: Holy Light IX (21290) sits only in source `world_drop:unknown` (source_origin classic-db), so no known Forever source. holy.md:81 and the Holy Light card note paladin-holy.json:135 now say the rank 9 book has no known source in Forever yet, so rank 8 is the top rank. Ahn'Qiraj is no longer named.
- **Decisions:** Retribution race stays Human (frontmatter and build strings untouched); Holy and Protection races unchanged.
- `updated:` bumped to 2026-11-04 on holy.md, protection.md, retribution.md. Builds, point orders, card actions and BiS files untouched.
- Prose check `-t paladin`: passes.
- **Sim findings:** none new. Sim request 1 (where Dwarf's raid edge comes from) stays open for the sim lane.
