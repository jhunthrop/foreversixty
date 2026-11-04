# Guide accuracy round 10: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr10/audit (origin/main 84460cec8: guide rounds 1 to 9 plus sim round 15). Client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go; source read from gr10/engine). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, the rendered card notes in `data/curated/apl/paladin-*.json`, `data/curated/rank-choices.json` and `data/builds/1.60.1.70755/addon-builds.json`. Before starting I read the paladin reports for rounds 1 to 9 (with their fix logs), round4/paladin-levels.md, the race-sims files, ledger.md, and the sim round 15 paladin and priest reports (design/reviews/2026-11-04-sim-audit-round15/paladin.md, priest.md W10).

## What changed since round 9's audit (git diff 645caa983..HEAD)

- **Guides:** holy.md:64 (the R9-1 Benediction fix) and protection.md:18 and :111 (Expertise restored, bb4d616e0). Nothing else.
- **Retribution APL (ba590bcc7, sim W6-6):** the Crusader opener was adopted. If the target has no Judgement of the Crusader (20303), the list casts Seal of the Crusader rank 6 at -1.5 s, judges it on the first swing, and then `castPaladinPrimarySeal` swaps to the damage Seal. The old Seal of Command and Seal of Righteousness pre-pulls now run only when the debuff is already up. Measured: bare 60 +9.0 to +9.2%, 50 +4.1 to +4.3%, 40 +11.1%, 30 +11.4 to +11.8%, 20 +9.3 to +9.5%, and a 30 s pull at 60 +5.5%. Raid 60 is unchanged. Neither new line has a card note.
- **Holy APL (4faba89a9, 416b2280f; priest W10, paladin W3-2):** the tank Flash of Light line went from rank 1 (19750) to **rank 5 (19942)**, measured at raid 60 as 639.3 / 640.2 against rank 4's 635.8 / 636.2. Holy Shock went from rank 1 to rank 4 (20930). The party Flash of Light stays rank 4.
- **Protection APL (9bd2b3112):** Dwarf casts Stoneform below 70% health. The line has no note.
- **Consecration under the level-gap rule (paladin W3-1, W3-4, W6-4):** the rule now reaches Consecration's spell power share. Retribution's rank 1 rows still win: rank 1 at 40 is +1.3% over the top rank learned, and at 50 +3.0%.
- **Holy point order (paladin H1):** Illumination 20, Divine Favor 21, Illumination 22 to 24, Oath of Light 25, Illumination 26, Infusion of Light 27 and 28, Purifying Power 29 and 30, Divine Precision 31, Holy Shock 32. Levels 10 to 19 and 33 to 60 did not move.

## Builds decoded (unchanged, all legal on 70755)

- **Retribution `build`** `050250002/502/0523330100130135` (14/7/30) and **`raidBuild`** `050050002/503/05233301001301351` (12/8/31). The talents are as round 9 lists them. retribution.md:58 matches.
- **Protection `build`** `05005/15302102513501031/5003` (10/33/8). protection.md:66 to :82 match.
- **Holy `build`** `50320215012101251/50322/0503` (31/12/8). holy.md:51 to :67 match.

## Point orders (addon-builds.json, index 0 = level 10)

- **Retribution:** Benediction 10 to 14, Improved Judgement 15 and 16, Conviction 17 to 19, Champion of the Light 35 to 39. retribution.md:74 and :125 match.
- **Protection:** Divine Intellect 10, Redoubt 11 to 15, Precision 16 to 18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, Holy Shield 30, Shieldsworn 40. protection.md:133 and :139 match.
- **Holy:** Spiritual Focus 10 to 14, Healing Light 15 to 17, Reverence 18 and 19, Oath of Light 25, ahead of Infusion of Light (27) and Divine Precision (31). holy.md:117 and :123 match the new order. The round 15 lane's suggested holy.md:112 text points at a line that no longer exists, and the current text has no wrong claim.

## Round 9 fix and the Expertise line re-checked

- **R9-1** (holy.md:64): correct. Holy Shock, Cleanse, the Blessings, the Seals and Judgement all sit inside Benediction 20101's class mask.
- **protection.md:111 Expertise:** the line holds. Holy Strike rolls `OutcomeMeleeWeaponSpecialHitAndCrit` (holy_strike.go:167), so expertise reaches it. The late ilvl 98 Lightbreaker, Inquisition, Duskwraith and Waywatcher sets carry it (items.json; bb4d616e0). One loose word is a SHOULD below.

## Checked and correct (no finding)

I checked every line of the four pages and every card note against spelltips/paladin.json, trainables/paladin.json, talents/paladin.json, spellranks.json, rank-choices.json and races.json, as round 9 did. Re-derived this round:

- **Holy Shock ranks:** 30 / 40 / 48 / 56, at 160 / 225 / 275 / 325 mana. "Use Holy Shock rank 3 or 4. Rank 1 costs half as much" (holy.md:81) holds, and it agrees with the rank 4 card.
- **Flash of Light ranks:** 20 / 26 / 34 / 42 / 50 / 58. Rank 4 costs 90 mana and heals 151, rank 5 costs 115 and heals 223, rank 6 costs 140 and heals 303. The party line at rank 4 (holy.md:79) matches the card.
- **Seal of the Crusader rank 6** (20308, level 52, 160 mana): 306 attack power, 40% faster, lighter swings (retribution.md:93).
- **Kings and Heart of the Lion** do not stack (protection.md:86). The sim lane confirmed this in W3 with two SoD sources.
- **Consecration ranks at 60:** rank 4 bare and rank 5 raid (retribution.md:95, last sentence) match rank-choices.json.
- **Stoneform** (retribution.md:111): 10% physical, 8 s (W6-1).
- **Source links:** all nine distinct URLs return 200, with the right content types and labels (fetched 2026-11-04). The source-link rule holds on all four pages.
- **Voice:** no em or en dashes, no site or sim talk, and no stat ranking in the Stat priority sections.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R10-1 | CONFIRMED | rotation | retribution.md:78 (read with :93) | "About a second and a half before the pull, put up your Seal: Seal of Righteousness while leveling, Seal of Command with Unsealed Wrath and raid buffs. Walk in and Judge on your first swing." | R3-6 was left pending a sim, with the rule that the guide adopts the opener if it gains beyond error. The sim gained, and the opener is now the card's (ba590bcc7; round 15 W6-6). With no Judgement of the Crusader on the target, the list seals Crusader at -1.5 s, judges on the first swing, then swaps to the damage Seal: bare 60 +9.0 to +9.2%, 20 to 50 +4.1 to +11.8%, a 30 s pull +5.5%, raid unchanged. The mechanism is in the client: Judgement keeps the Seal (20271 "Does not consume the Seal"), the debuff lasts 40 s, and your melee and Sacred Arbiter refresh it. Mobalytics' Forever Retribution guide opens the same way (quote in round3/paladin.md Sources). The page's opener now tells a solo player to skip it, against the card on the same page. | "**The opener.** Put a Blessing of Might on yourself if nobody has given you one. If no other paladin has Judgement of the Crusader on the target, put up Seal of the Crusader about a second and a half before the pull, Judge it on your first swing, then put up your damage Seal: Seal of Righteousness while leveling, Seal of Command with Unsealed Wrath and raid buffs. Your swings and Holy Strike keep the Crusader debuff up, so that one Judgement adds Holy damage to everything you do for the rest of the fight. If the raid already has Crusader on the boss, open straight with your damage Seal and Judge on your first swing." Also end :93 with: "On your own, always open with Crusader as above: one Seal and one Judgement cover the whole fight." |
| PALADIN-R10-2 | CONFIRMED | rotation (open sim item, now measured) | retribution.md:95 | "Before 40, rank 1 does more over a fight than the higher ranks because it costs so little; from 40 move up to the newest rank you can afford." | This is the "leveling Consecration rank under the level-gap curve" item, carried since round 1. Round 3 kept "from 40 move up" because the rank-1 rows predated the curve. The curve now reaches Consecration (round 15 W3-1, W6-4; core.LevelPenalty on the ground effect), and the rows were re-measured with it: rank 1 at 40 is 112.8 against the top rank learned at 111.3 (+1.3%), and at 50 it is 166.5 against 161.6 (+3.0%). Ranks 2 and 3 lose too (W3-4, 10,000 iterations). rank-choices.json keeps rank 1 for 30 to 59 on that measurement. Rank 1's base is tiny, so the curve takes little from it. The mana it saves goes into Seals, Judgements and Holy Strike. | "Below 60, rank 1 does the most over a fight: it costs 135 mana, and the mana it saves goes into Seals, Judgements and Holy Strike. At 60, use rank 4 without raid buffs and rank 5, the top rank, with them. On two or more targets, cast it freely." |
| PALADIN-R10-3 | LIKELY | leveling | retribution.md:131 and :133 (and :89) | :131 "Pull with Judgement, Holy Strike on cooldown, and let Seal of Righteousness do the rest." :133 "rank 1 at 135 mana is the rank to keep on your bar through your 30s ... so from 40 move up to the newest rank you can afford, and keep moving up as you train them." | The same two measurements as R10-1 and R10-2, applied to the leveling text. The Crusader opener gains at every bare band from 20 to 60 and on a 30 s pull. Consecration rank 1 wins single-target at 40 and 50 with the curve on. Both were measured on boss-length or 30 s fights, not on 2 or 3 questing mobs, so this is LIKELY. A Crusader Judgement on a mob that dies in 15 s still pays for one Seal; a Consecration rank on a pack is not measured (sim 1). | :131: "Pull with Seal of the Crusader and a Judgement on your first swing, then Seal of Righteousness and Holy Strike on cooldown. Sanctified Judgement refunds most of each Seal, so on one or two mobs at a time you rarely drink." :133: "...two or three mobs on you is the time for it, and rank 1 at 135 mana is the rank to keep on your bar until 60: a rank loses its spell damage bonus once you are far above it, but rank 1 has so little to lose that its low cost wins." :89 first sentence: "After the opening Crusader Judgement, stay in Seal of Righteousness." |
| PALADIN-R10-4 | CONFIRMED | rotation (card rank) | holy.md:78, :127 (and :33, :95) | :78 "**Flash of Light rank 4** on the tank below 90%. This is your staple. ... rank 4 is the cheapest rank that still carries nearly all of your healing power at 60." :127 "rank 4 under Oath of Light does the work". :33 "a middle rank Flash of Light on them". :95 "your cheap heals are Flash of Light rank 4 and Holy Shock rank 3" | Round 15 moved the card's tank line to Flash of Light **rank 5** (19942; 4faba89a9). It measured 639.3 / 640.2 against rank 4's 635.8 / 636.2 at raid 60, 10,000 iterations (priest W10). The card now says rank 5 on the tank, and the prose says rank 4 on the same page. Rank 5 is learned at 50 and keeps its whole healing bonus at 60, for 115 mana against rank 6's 140. Holy Shock's card is rank 4 (W3-2: rank 4 +0.5% over rank 1), and :95 names rank 3. The party line stays rank 4 on both. | :78: "**Flash of Light rank 5** on the tank below 90%. This is your staple. Oath of Light adds the same flat amount to every rank, and rank 5 carries all of your healing power at 60 for less mana than rank 6. Rank 1 is a trap there: Forever takes the healing power off ranks learned far below your level, so it heals for little more than its base and the Oath." :127: "...rank 5 under Oath of Light does the work for less mana than the top rank." :33: "...so a lower rank Flash of Light on them heals for most of what the top rank does..." :95: "...so at 60 your cheap heals are Flash of Light rank 4 or 5 and Holy Shock rank 3 or 4, not rank 1." |

## SHOULDs (wording, not findings)

- protection.md:111: "Expertise, which only a few late sets carry". The paladin Tier 1 Protection 4P bonus (client 1301083) also gives expertise, and Molten Core is not late. "which only a few raid sets carry" covers both.
- protection.md:119: Stoneform as "a free defensive". It costs no resource, but it is on the global cooldown (W6-1). "a defensive that costs no mana" is exact.
- Carried: the Holy Strike card's base cost against Benediction, and holy.md:117 naming Flash of Light at 10 to 14.

## For the component and data lanes (not guide findings)

1. **Retribution opener list now renders three Seals, ungated.** `prepullCasts` (web/src/lib/guides/rotation-lists.ts `gateOf`) only reads a bare `auraIsKnown` or `not auraIsKnown`. The Seal of Command and Seal of Righteousness pre-pulls are now `and[auraIsKnown(Unsealed Wrath), auraIsActive(20303)]` and the `not` form. So all three Seals (Crusader, Command, Righteousness) print as "1.5 s before the pull" with no "with/without Unsealed Wrath" gate. The opener card states something false: three Seals at once. Either gateOf reads the `and`, or the opener prints only the gate-free Crusader line.
2. **Missing card notes** for the Crusader pre-pull, the primary-Seal line (paladin-retribution.json priorityList 2) and the Protection Stoneform line. Proposed in round 15 paladin.md Wave 6: "On your own, Seal of the Crusader before the pull: its Judgement on your first swing makes the target take more Holy damage for the whole fight, because your swings keep it up. Skip it when another paladin's Crusader is already on the target." / "Right after that Judgement, swap to your damage Seal." / "Dwarf: Stoneform when you drop under 70% health. Ten percent less physical damage for 8 seconds, on the global cooldown, every 3 minutes."
3. **The Holy APL's top-level note** still names Holy Shock rank 1 and the tank Flash of Light at rank 1, and it says the engine has "no downrank penalty" (flagged in round 15 priest.md). It is not shown in the guide prose.

## Decisions

- **R10-1, R10-4:** decided by the round 10 rule that prose must match the new cards, by the round 15 measurements, and by round 3's own adoption rule for the Crusader opener. No owner question.
- **R10-2:** decided by the measured rows with the level-gap rule now on Consecration. That closes the sim item carried since round 1.
- **Races:** unchanged. Retribution Human (Undead), Holy Human (Undead), Protection Dwarf (Undead). No paid race named.

## Sims requested

1. Retribution leveling Consecration on 3 targets: rank 1 against the top rank learned at bands 40 and 50, bare, on a 20 s pack fight. This settles R10-3's :133 wording for packs.
2. Retribution Crusader opener on a 15 s single mob at bands 20, 30 and 40. This confirms R10-3's :131 for questing.
3. Carried: Retribution Human against Dwarf on a Beast target, and Protection threat by ability on 1 and 4 targets.
4. PROVISIONAL (final BiS re-rank): none filed. No prose names a BiS piece.

## Sim findings

None new. Round 15's Consecration level-gap (W3-1, W6-4), own-Crusader rank (W6-5) and Stoneform GCD (W6-1) fixes agree with the 70755 client.

## Sources

- 70755 client data, data/builds/1.60.1.70755: spelltips/paladin.json (20308, 19941, 19942, 19943, 1311606, 20473, 20929, 20930, 20924, 26573, 20271, 10333, 24239), trainables/paladin.json (Flash of Light, Holy Shock, Consecration, Seal of the Crusader, Holy Light), talents/paladin.json, spellranks.json, items.json (Lightbreaker, Inquisition, Duskwraith, Waywatcher sets), addon-builds.json.
- data/curated/apl/paladin-{holy,protection,retribution}.json; data/curated/rank-choices.json.
- Commits ba590bcc7, 9bd2b3112, 416b2280f, 4faba89a9, bb4d616e0.
- design/reviews/2026-11-04-sim-audit-round15/paladin.md (W3-1 to W3-5, W6-1 to W6-6, H1) and priest.md (W10).
- Engine 9f7f1f6e4: sim/paladin/holy_strike.go:167, sim/paladin/hammer_of_the_righteous.go, sim/core/hitprofile.go, sim/core/client_sets_gen.go:555.
- web/src/lib/guides/rotation-lists.ts (prepullCasts, gateOf).
- Mobalytics, Retribution Paladin Guide, WoW Forever: https://mobalytics.gg/wow-forever/classes/retribution-paladin-guide (403 on direct fetch 2026-11-04, so unknown; quoted via the search index in round3/paladin.md: "Put on Seal of the Crusader as you run up to an enemy and use Judgement on them as you do your first Auto Attack.")
- Blizzard Deep Dive and What's Next recaps, Wowhead Undead paladins, Power Up Gaming, Warcraft Tavern (all 200 on 2026-11-04; URLs as in round 9).

## Counts

| status | count |
|---|---|
| CONFIRMED | 3 (R10-1 Crusader opener, R10-2 Consecration rank from 40, R10-4 tank Flash of Light rank) |
| LIKELY | 1 (R10-3 leveling opener and Consecration text) |
| NEEDS-SIM | 2 new (3 carried) |
| NEEDS-OWNER | 0 |

Levels 1 to 19 text: 0 errors. Source links: 0 errors. Round 9 edit: 0 errors. Every confirmed error comes from sim round 15 measurements and card changes that the prose has not caught up with.

## Fix log (round 10)

- **PALADIN-R10-1** fixed: retribution.md:78 (Crusader opener, then the damage Seal; straight to the damage Seal when the raid's Crusader is up) and :93 (ends "On your own, always open with Crusader as above"). Checked against spelltips 20308 ("Your melee strikes will refresh the spell's duration", 40 s) and 20271 ("Does not consume the Seal"), and engine core/debuffs.go:323-329 (own melee proc mask refreshes, so Holy Strike counts).
- **PALADIN-R10-2** fixed: retribution.md:95, rank 1 below 60, rank 4 bare and rank 5 raid at 60, as rank-choices.json and W3-4/W6-4 measure.
- **PALADIN-R10-3** (LIKELY) fixed with the audit text: retribution.md:89, :131, :133. The single-target and 30 s pull measurements back it. A 15 s questing mob and a 3-target pack are unmeasured (sims 1 and 2 below). The wording stays true either way: it tells you how to open and which rank to keep on the bar, and it makes no claim about packs.
- **PALADIN-R10-4** fixed: holy.md:33, :78, :95, :127. Rank 5 (19942, level 50, 115 mana, heals 223) keeps its full bonus at 60 under core.LevelPenalty (grace 17 levels, 5% a level after that). Rank 4 at 42 keeps 95%.
- **SHOULD protection.md:111** applied: "which only a few raid sets carry".
- **SHOULD protection.md:119** applied: "a defensive that costs no mana". protection.md:107 now also explains the card's Stoneform line (Dwarf, under 70% health, 10% physical, 8 s, 3 min, costs a global, no Forbearance).
- **Card notes** (notes only, no actions touched):
  - paladin-protection.json:155: the Stoneform note, as the brief gave it.
  - paladin-retribution.json:188 (priorityList 2, the primary-Seal swap): the note covers the Crusader opener and the swap to your damage Seal.
  - The Crusader pre-pull cannot carry its own note. The APLPrepullAction proto (engine proto/apl.proto:32-36) has no `notes` field, no curated APL puts a note on a prepull entry, and prepullCasts (rotation-lists.ts:147) never reads one. I added a note there, then took it out so protojson parsing would not break. The opener text sits on the swap line instead, and in the prose at :78.
- **Stat priority (new rule):** all three spec pages were re-read. None of them ranks: they give the reasons and point to the table. The prose test passes for paladin (`-t paladin`). No other text on these pages ranks stats or gear by a nightly value.
- **updated:** set to 2026-11-04 on retribution, holy and protection (all three were already on that date). index.md was not changed.
- **Left:** the Holy APL's top-level engineering `notes` still names rank 1 Holy Shock and rank 1 Flash of Light, and says the engine has "no downrank penalty". It does not render in the guide, so it belongs to the data lane (audit item 3). The prepull gate rendering (audit item 1) belongs to the component lane.
- **SIM findings:** none new. Sims 1 and 2 (packs and the 15 s mob) are still open, along with the carried items.
