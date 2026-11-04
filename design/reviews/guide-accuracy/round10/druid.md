# Guide accuracy round 10: Druid

Auditor: druid persona. Tree: gr10/audit at 84460cec8 (origin/main with all nine guide rounds and sim round 15; READ-ONLY), client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go). Engine source read at gr10/engine (9f7f1f6e4).

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`). Before starting I read the round 1 to 9 druid reports, round4/druid-levels.md, the race-sims files, ledger.md and design/reviews/2026-11-04-sim-audit-round15/druid.md.

## What moved since round 9 (645caa983 to 84460cec8)

Sim round 15 changed four of the five druid build strings and the Restoration card ranks. In each case the commit edited only the frontmatter `build:` line, so the prose around it still describes the old build:

- 6b9b19145 and 0eeab841a: cat leveling `01/50212032123012123001/535033` becomes `050022/50212032123012123001/05501`, and cat raid `05002/50230032023011125001/05503` becomes `050022/50230032023012125001/055`. feral.md changed 2 lines, both in the frontmatter.
- 6de6cbf98: bear `0/5523033202310212551/051` becomes `0/5123033202310212551/055`. feral-bear.md: build line only.
- 9b931dc3d: Restoration `05312001/0/55530051531132` becomes `05302001/0/555300515311321`. restoration.md: build line only.
- 4faba89a9: the Restoration party Rejuvenation lines cast rank 7 (8910) at raid 60, in place of rank 1. rank-choices.json keeps rank 1 for 10 to 59 and for bare 60.
- addon-builds.json was regenerated for every druid build (see below).
- Engine W8-1 (E 230a079c5): the engine now carries the 70755 values for Claw, Shred, Ravage, Rake, Rip, Ferocious Bite and Starfire. Round 9's S13 is closed at this pin. K3: Ravage now checks that you are behind the target.

Round 15's own "For the guide lane" notes list most of the feral and Restoration line fixes. No guide lane applied them.

## Decoded builds (all legal, 51 points, prerequisites met)

- **Balance** (unchanged): leveling `5232220115101151/0/555301`, raid `5232220115501351/0/505003`.
- **Feral leveling** `050022/50212032123012123001/05501`, 9/31/11:
  - Balance: Genesis 5, Nature's Majesty 2, Nature's Reach 2.
  - Feral: Ferocity 5, Feral Swiftness 2, Feral Instinct 1, Brutal Impact 2, Shredding Attacks 3, Savage Fury 2, Feral Charge 1, Sharpened Claws 2, Predatory Strikes 3, Predatory Discipline 1, Blood Frenzy 2, Leader of the Pack 1, Natural Instinct 2, Rend and Tear 3, Shifting Power 1.
  - Restoration: Heart of the Wild 5, Naturalist 5, Natural Shapeshifter 1.
- **Feral raid** `050022/50230032023012125001/055`, 9/32/10:
  - Balance: the same 9.
  - Feral: Ferocity 5, Feral Swiftness 2, Feral Instinct 3, Shredding Attacks 3, Savage Fury 2, Sharpened Claws 2, Predatory Strikes 3, Predatory Discipline 1, **Blood Frenzy 2**, Leader of the Pack 1, Natural Instinct 2, Rend and Tear 5, Shifting Power 1.
  - Restoration: Heart of the Wild 5, Naturalist 5, **Natural Shapeshifter 0**.
- **Bear** `0/5123033202310212551/055`, 0/41/10:
  - Feral: Ferocity 5, **Furor 1**, Feral Swiftness 2, Feral Instinct 3, Thick Hide 3, Shredding Attacks 3, Savage Fury 2, Sharpened Claws 2, Predatory Strikes 3, Lacerate 1, Blood Frenzy 2, Leader of the Pack 1, Natural Instinct 2, Rend and Tear 5, Natural Reaction 5, Primal Bite 1.
  - Restoration: Heart of the Wild 5, **Naturalist 5**.
- **Restoration** `05302001/0/555300515311321`, 11/0/40:
  - Balance: Genesis 5, Moonglow 3, Nature's Majesty 2, Nature's Splendor 1. **No Improved Moonfire.**
  - Restoration: Nature's Focus 5, Heart of the Wild 5, Naturalist 5, Reflection 3, Gift of Nature 5, Bloom 1, Tranquil Spirit 5, Improved Rejuvenation 3, Swiftmend 1, Nature's Swiftness 1, Living Spirit 3, Improved Tranquility 2, **Improved Regrowth 1**.

Shipped point orders (addon-builds.json at HEAD, index 0 = level 10, entries [tree, row, column] with 1-based columns):

- **Balance** (unchanged): Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, ... Moonkin Form 40.
- **Feral leveling:** Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19, Shredding Attacks 20-22, Savage Fury 23-24, Feral Charge 25, Sharpened Claws 26-27, Predatory Strikes 28-30, Predatory Discipline 31, Blood Frenzy 32-33, Leader of the Pack 34, Natural Instinct 35-36, Rend and Tear 37-39, Shifting Power 40, **Genesis 41-45, Nature's Majesty 46-47, Nature's Reach 48-49, Heart of the Wild 50-54, Naturalist 55-59, Natural Shapeshifter 60**.
- **Bear:** Ferocity 10-14, **Furor 15, Feral Swiftness 16-17, Feral Instinct 18-19**, Savage Fury 20-21, Feral Instinct 22, Thick Hide 23-24, Lacerate 25, Thick Hide 26, Shredding Attacks 27-29, Leader of the Pack 30, Sharpened Claws 31-32, Predatory Strikes 33-35, Blood Frenzy 36-37, Natural Instinct 38-39, Primal Bite 40, Rend and Tear 41-45, Natural Reaction 46-50, Heart of the Wild 51-55, Naturalist 56-60.
- **Restoration:** Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19, ... Improved Tranquility 47-48, **Improved Regrowth 49**, Genesis 50-54, Moonglow 55-57, Nature's Majesty 58-59, Nature's Splendor 60.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| DRUID-R10-1 | CONFIRMED | talent, leveling | feral.md:55 | "From 41 take Genesis, then Nature's Focus and three points of Heart of the Wild to open Restoration, then Naturalist, Natural Shapeshifter and Subtlety last. Nature's Focus keeps a mob's hits from pushing back ... Subtlety and Nature's Focus both move at the 60 respec. At 60, respec into the raid build: Brutal Impact, Feral Charge and one point of Blood Frenzy finish Rend and Tear and fill Feral Instinct, and Nature's Focus and Subtlety become the rest of Heart of the Wild, Genesis and Nature's Majesty." | The leveling build is now `050022/50212032123012123001/05501`: no Nature's Focus and no Subtlety, Heart of the Wild 5, Blood Frenzy 2. The shipped order from 41 is Genesis 41-45, Nature's Majesty 46-47, Nature's Reach 48-49, Heart of the Wild 50-54, Naturalist 55-59, Natural Shapeshifter 60. The raid diff is Brutal Impact 2, Feral Charge 1 and Natural Shapeshifter 1 out, and Feral Instinct +2 and Rend and Tear +2 in. Blood Frenzy is 2 in both builds. The page's own leveling build contradicts every clause from "From 41" on. | "From 41 take Genesis, then Nature's Majesty and Nature's Reach in Balance, then Heart of the Wild, Naturalist and a point of Natural Shapeshifter in Restoration. At 60, respec into the raid build: Brutal Impact, Feral Charge and that Natural Shapeshifter point finish Rend and Tear and fill Feral Instinct." |
| DRUID-R10-2 | CONFIRMED | talent | feral.md:60 | "a bonus combo point on builder crits (every time with both points while leveling, a 50% chance with the raid build's single point)" | The raid build has Blood Frenzy 2 (digit 13 of the Feral tree is 2). Rank 2 text: "a 100% chance to add an additional Combo Point". | "a bonus combo point on every builder crit, with both points in either build" |
| DRUID-R10-3 | CONFIRMED | talent | feral.md:67 | "The raid build's Balance points are Genesis, 10% more periodic damage (a bigger Rip), and Nature's Majesty, 4% crit. Natural Shapeshifter (3 points in both builds) makes Shifting Power cheaper, which is mana you spend every 8 seconds." | The raid build has Nature's Reach 2 ("improves your chance to hit by 4%"; melee hit aura 54 on 16819, per sim round 15 W4-1) and Natural Shapeshifter 0. The leveling build has Natural Shapeshifter 1. | "The raid build's Balance points are Genesis, 10% more periodic damage (a bigger Rip), Nature's Majesty, 4% crit, and Nature's Reach, 4% hit on your attacks as well as your spells. Shifting Power's mana holds up through a raid fight, so Natural Shapeshifter is not worth its points there." |
| DRUID-R10-4 | CONFIRMED | spell | feral.md:81 | "Shifting Power costs 55% of your base mana, 30% less with Natural Shapeshifter's three points" | No build on the page takes three points: raid 0, leveling 1 (10%, "Reduces the mana cost of all shapeshifting by 10%"). | "Shifting Power costs 55% of your base mana, and mana decides how many you get." Then keep the rest of the sentence as it is. |
| DRUID-R10-5 | CONFIRMED | talent | feral.md:65, :89 | ":65 Heart of the Wild is 6% more Strength in Cat Form at three points and 10% at the raid build's five"; ":89 Heart of the Wild adds 6% to it (10% with the raid build's five points)" | Both builds now take Heart of the Wild 5 (Restoration digit 2 is 5 in both). "Three points" was the old leveling build's figure, so the contrast with "the raid build's five" now contradicts the page. | :65 "Heart of the Wild is 10% more Strength in Cat Form at five points; Naturalist is 10% more damage of every kind." :89 "Strength is attack power, and Heart of the Wild adds 10% to it." |
| DRUID-R10-6 | CONFIRMED | talent | feral-bear.md:52 | "The build is 45 points of Feral and 6 of Restoration." | `0/5123033202310212551/055` is 41 Feral and 10 Restoration. | "The build is 41 points of Feral and 10 of Restoration." |
| DRUID-R10-7 | CONFIRMED | talent | feral-bear.md:56 (and the :64 list) | "**Furor** (5 points): 10 rage every time you shift into Bear Form. A shift out to battle rez or Innervate comes back with rage." | Furor is now 1 point. Rank 1 text: "a 20% chance to gain 10 Rage when shapeshifting into Bear Form". Naturalist is now 5 ("increases all damage you deal by 10%"), and the list does not name it. Round 15 K1 moved four Furor points into Naturalist for threat. | Replace the Furor bullet with: "**Furor** (1 point): the path point to Feral Swiftness; a tank does not shift in a fight, so the other four points go to Naturalist." Change the Heart of the Wild bullet to: "**Heart of the Wild** and **Naturalist** (5 points each): 20% more Stamina in Bear Form, and 10% more damage of every kind, which is threat." |
| DRUID-R10-8 | CONFIRMED | leveling | feral-bear.md:115 | "Furor from 15 to 19 adds 20% to the chance that a shift into Bear Form grants 10 Rage, so by 19 every shift does, and a pull opens with Maul ready." | Shipped bear order: Furor 15 (1 point), Feral Swiftness 16-17, Feral Instinct 18-19. | "Furor takes one point at 15, for a 20% chance of 10 Rage when you shift. Feral Swiftness at 16 and 17 adds 4% dodge, and Feral Instinct at 18 and 19 adds 20% to Swipe." (Feral Instinct is 10% a point, 30% at 3.) |
| DRUID-R10-9 | CONFIRMED | talent | restoration.md:51 | "The build is 39 points of Restoration and 12 of Balance." | `05302001/0/555300515311321` is 40 Restoration and 11 Balance. | "The build is 40 points of Restoration and 11 of Balance." |
| DRUID-R10-10 | CONFIRMED | talent | restoration.md:62 | "The one Balance point left after that sits in Improved Moonfire, which does nothing for a healer. A point of Improved Regrowth, 10% more Regrowth crit, heals a little more if you would rather spend it there." | The build now has Improved Moonfire 0 and Improved Regrowth 1 (sim round 15 R1). The page says the point is in Improved Moonfire, and its own build string has it in Improved Regrowth. | Cut both sentences. Add to the talent list: "- **Improved Regrowth** (1 point): 10% more crit on the Regrowth you keep on the tank." |
| DRUID-R10-11 | CONFIRMED | rotation | restoration.md:80 | "Rank 1 Rejuvenation is not a cheap heal at 60. ... Your last three ranks (8, 9 and 10) keep all of it, so on a hurt party member cast rank 8 or 9" | The card now casts party Rejuvenation at rank 7 (8910) at raid 60, in priorityList 9 to 13 (commit 4faba89a9). rank-choices.json casts rank 1 on the party from 10 to 59 and at bare 60 ("+6.2% (the fight runs dry)"). The prose names ranks the card never casts, and "rank 1 is not a cheap heal at 60" contradicts the bare-60 row. Rank levels (spellranks.json): r7 40, r8 46, r9 52, r10 58. Engine LevelPenalty (grace 17, 5% a level) leaves r7 85% of your bonus at 60 and r8 to r10 100%, so "last three keep all of it" is itself right. The tank lines (rank 4 from 28 to 49, top rank otherwise) match the prose. | "**Ranks.** A rank learned far below your level loses its share of your healing power: ranks 8, 9 and 10 keep all of it at 60, rank 7 most of it, and rank 1 none. In a raid, keep the top rank on the tank and put rank 7 on a hurt party member: most of the heal for far less mana. When there is no raid behind you and the fight will run you dry, as in a long five-man, rank 1 on the party is the most healing your mana buys, because what it costs is so small. While leveling, rank 1 is your party Rejuvenation too, and on the tank from 28 to 49 a middle rank such as rank 4 keeps the heal sizeable and your mana going. Use the top ranks of Regrowth and Healing Touch." |
| DRUID-R10-12 | NEEDS-SIM | rotation | feral.md:73, :37; druid-feral.json priorityList[6] note | "Rip at 5 points does far more over its duration than a Bite, so it always wins." (card: "does far more than a Bite") | Carried from R9-1. The engine now has the 70755 values (W8-1), but no run has compared the two. Hand count at 60 (about 1200 AP, 70755 tooltips): Rip r6 5 CP = 6×(19+5×31.9) = 1071, plus about 24% AP = about 1360, plus Genesis 10% = about 1500. Bite r5 5 CP = 855 to 919, plus about 15% AP, plus 3.1 for each spare energy (kept by Predatory Discipline) = about 1150 to 1250 before Rend and Tear's 15% (which Rip's ticks do not get). Both can crit. Rip leads by roughly 10%, not "far". | No prose change until Sims requested 1 runs. If Rip leads by under 15%, write: "Rip at 5 points comes first: it costs less, does more over its duration and keeps Rend and Tear on." |

Checked and right (no finding):

**Rest of feral.md**
- Feral Levels 1 to 19 (:111-115): Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17 and Brutal Impact 18-19 match the shipped order.
- Feral Levels 20 to 60 (:119): Shredding Attacks at 20-22 for the 42-energy Shred, Feral Charge at 25 (the Cat Form leap 414924 lands behind the target), and Ravage at 32 and Pounce at 36, both needing you behind the target. All match.
- Card notes: Claw "40 to Shred's 54 at 20 and 48 at 21" holds, because Shredding Attacks is still 20-22. "Shred hits harder than Claw at every rank" holds on the 70755 values (200% plus 22 to 80, against 140% plus 27 to 115). The Shifting Power note matches 1322605 (8 s, 40 energy).
- Weaknesses :46 and :67 "The raid build skips Furor": raid Furor is 0. Right.

**Rest of feral-bear.md**
- :55 "9% Dodge" (Natural Reaction 5% plus Feral Swiftness 4%), :57-62, the priority list and the card notes are unchanged and right.
- :121 "Lacerate arrives at 25 and Primal Bite at 40" matches the new order (Lacerate 25, Primal Bite 40).
- :91 Expertise: commit bb4d616e0 records 71 catalogue items with expertise (the ilvl-98 sets). The round 9 ledger accepts the line.

**Rest of restoration.md**
- :64 Wild Growth swap: dropping Improved Tranquility 2 for Wild Growth 1 and Improved Regrowth +1 (to 2) is still legal at 38 points before row 7, with Living Spirit 3. Wild Growth is 3.0 s, so 2.5 s with Naturalist 5. Right.
- Levels 1 to 19: Heart of the Wild 10-14, Reflection 15-17 and Nature's Focus 18-19 match the order (unchanged).
- Card notes: Swiftmend 20% base mana, Tranquility r4 when two are under 70%, and Healing Touch r10 under 55%. All right.

**balance.md and index.md**: no druid data, build or card note changed under them. I re-read every line. The Eclipse loop still beats plain Starfires on the 70755 Starfire. By hand, Wrath r8 and two 2.5 s Starfires in 6.5 s is about 196 base damage a second, against about 163 for 3.0 s plain Starfires, and it is ahead on the spell power coefficient as well. Nothing is false.

**Sources**: every frontmatter link returned 200 to a browser user agent on 2026-11-04: Blizzard deep dive, Blizzard panel recap, Wowhead Skyborne First Look, Talents Forever data.json, /sim/weights, /sim/drops, /bis/druid/feral-bear and /bis/druid/restoration. Each one follows the round 4 source-link rule.

No em or en dashes. No site, sim, APL, preset or band words in prose. No quoted nightly numbers.

## Shoulds (wording, not findings)

- balance.md:91 "it always lands on a Starfire or a DoT": "almost always" is exact (carried).
- feral.md:73 "so do not pool energy for it" (carried; see Sims requested 2).
- feral-bear.md:66 "Left out: Brutal Impact ... Feral Charge": Furor is now a single path point. A line saying the other four Furor points went to Naturalist would help the reader, and R10-7 covers that.

## Decisions

1. **A build string that moved without its prose is a CONFIRMED finding on every sentence it falsifies** (R10-1 to R10-10). The sim commits followed the adoption rule, so the builds are right. The prose is what is wrong.
2. **Restoration ranks follow the new cards** (R10-11). The level-gap rule is unchanged. The rank the card casts on the party moved from 1 to 7 at raid 60 and stays 1 everywhere else. The prose now says both. The "last three ranks keep all of it" fact is kept, because the engine's 17-level grace confirms it.
3. **Bear Furor at 1 is adopted (round 15 K1)**, and it is not an owner question. The prose follows the build.
4. **Rounds 1 to 9 decisions stand**: Swiftmend keeps the HoT, Gift of Nature misses Tranquility, Moonkin casts Rebirth, Innervate and Remove Curse, Pounce and Ravage need you behind the target, and a failed Shred refunds 80%.

## Sims requested

1. **druid-feral, Rip against Bite at 5 points (R10-12, carried R9-1)** at pin 9f7f1f6e4 (70755 values in), band 60 raid and bare, seeds 1 and 50001. Compare the shipped list (Bite only while Rip has 6 s or more) with a variant that Bites in place of a Rip refresh. Report the DPS of each and each finisher's share.
2. **druid-feral, Bite energy pooling under Predatory Discipline** (carried): the shipped list against holding the 5-point Bite until 80 energy or more.
3. **Race re-check on the new bear build** `0/5123033202310212551/055` (Night Elf against Tauren, threat and chance of death, raid 60, seeds 1 and 50001). feral-bear.md:103 says Night Elf "takes slightly less damage and holds a little more threat" from round 4's measurement on the old build. It is PROVISIONAL until this runs.
4. **PROVISIONAL BiS and stat order:** after the 13:00 UTC re-rank, check that the feral and bear statPriority frontmatter and the stat prose still read in the table's order. The engine follow (W8-1) moves the cat weights a long way.
5. **Carried (S10):** feral leveling, the switch from Claw to Shred, gated on Shredding Attacks 3/3 at 22 against the Ferocious Bite rank 5 gate, ladder 20 to 25, on the 70755 values.

## Sim findings

- Round 9's S13 is closed at 9f7f1f6e4 (W8-1, E 230a079c5), and S1 Rend and Tear 3% is closed (E e45a540db).
- The APL top-level `notes` in druid-restoration.json still say "rank 1 on the party at every level, and on the tank rank 10 in a raid". The party lines cast rank 7 at raid 60. This is an internal note and is not rendered, so it is for the sim lane to tidy.
- Carried: S7 is resolved by W4-1. S8 Feral Charge (Cat) needs no model (K4). S11 the bear expertise weight. S12 bear Defense noise.

## Sources

- Client 1.60.1.70755:
  - talents/druid.json: Furor, Naturalist, Heart of the Wild, Blood Frenzy, Natural Shapeshifter, Nature's Reach, Improved Regrowth, Improved Moonfire, Feral Instinct.
  - addon-builds.json, at HEAD and at 645caa983.
  - spellranks.json: Rejuvenation, Regrowth, Healing Touch and Tranquility rank levels.
  - spelltips/druid.json: Rip 9896, Ferocious Bite 22829 and 31018, Shifting Power 1322605.
  - data/curated/rank-choices.json (druid-restoration rows) and data/curated/apl/druid-*.json.
- Commits: 6b9b19145, 0eeab841a, 6de6cbf98, 9b931dc3d, 4faba89a9, bb4d616e0.
- design/reviews/2026-11-04-sim-audit-round15/druid.md: S1, S2, R1, W3-2, W4-1, W8-1, K1, K3, K4, and its "For the guide lane" lines.
- Engine 9f7f1f6e4: sim/core/spell.go:786-831 (LevelPenalty, grace 17, 5% a level).
- Blizzard, "World of Warcraft: Forever Class Deep Dives: Hunter and Druid", https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (200).
- Blizzard, "World of Warcraft: Forever Deep Dive Panel Recap", https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200).
- Wowhead, "Skyborne First Look", https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829 (200).
- Talents Forever, https://talentsforever.com/data.json (200).
- Forever Sixty: https://foreversixty.gg/sim/weights, https://foreversixty.gg/sim/drops, https://foreversixty.gg/bis/druid/feral-bear, https://foreversixty.gg/bis/druid/restoration (all 200).

## Fix log (round 10)

Findings:
- DRUID-R10-1: fixed, feral.md:55. Leveling order from 41 now Genesis, Nature's Majesty, Nature's Reach, Heart of the Wild, Naturalist, Natural Shapeshifter 1; the respec diff is Brutal Impact, Feral Charge and that Natural Shapeshifter point into Rend and Tear and Feral Instinct (4 out, 4 in; checked against both build strings).
- DRUID-R10-2: fixed, feral.md:60 (Blood Frenzy 2 in both builds; rank 2 text "100% chance").
- DRUID-R10-3: fixed, feral.md:67. Balance points now named for both builds, with Nature's Reach ("improves your chance to hit by 4%"). The suggested reason for dropping Natural Shapeshifter ("Shifting Power's mana holds up through a raid fight") is not verified, so the line gives the stable reason instead: the talent only trims Shifting Power's mana cost, and in a raid the point does more in Feral Instinct and Rend and Tear (the adoption-rule result, no number).
- DRUID-R10-4: fixed, feral.md:81.
- DRUID-R10-5: fixed, feral.md:65 and :89 (Heart of the Wild 5 in both builds, 10% Strength in Cat Form).
- DRUID-R10-6: fixed, feral-bear.md:52 (41 Feral, 10 Restoration).
- DRUID-R10-7: fixed, feral-bear.md:56 and :64. Departure from the exact fix: Furor is not a prerequisite of Feral Swiftness (talents/druid.json: no prereq_talent_id; Ferocity 5 opens tier 1), so "the path point to Feral Swiftness" would be false. The bullet now gives the rank 1 text (20% chance of 10 rage on shifting into Bear Form) and says a tank seldom shifts in a fight, so the rest went to Naturalist. Naturalist joins the Heart of the Wild bullet (10% more damage of every kind, which is threat).
- DRUID-R10-8: fixed, feral-bear.md:115 (Furor 15, Feral Swiftness 16-17 for 4% dodge, Feral Instinct 18-19 for 20% Swipe; tooltip 30% at 3 points).
- DRUID-R10-9: fixed, restoration.md:51.
- DRUID-R10-10: fixed, restoration.md:62-63 (Improved Moonfire sentences cut; Improved Regrowth bullet added, rank 1 text "critical effect chance of your Regrowth spell by 10%").
- DRUID-R10-11: fixed, restoration.md:81, close to the proposed wording. Checked against rank-choices.json (rank 1 on the party 10-59 and bare 60, rank 4 on the tank 28-49) and the APL (party lines cast 8910 rank 7; tank Rejuvenation rank 10). Rank levels and the 17-level grace give r7 85% and r1 0% at 60.
- DRUID-R10-12 (NEEDS-SIM): per integrator ruling, "far more" is dropped and no Rip-vs-Bite damage claim remains. feral.md:73 now says Rip goes out before any Bite because the bleed keeps Rend and Tear on for everything else; feral.md:37 Strengths no longer calls Rip "the hardest-hitting finisher" (now: bleeds the boss for its whole duration and keeps Rend and Tear on). Card note druid-feral.json priorityList[6] reworded the same way. True whatever the sim finds.

Decisions:
- 1 (prose follows moved builds): applied in every item above.
- 2 (Restoration ranks follow the new cards): applied, restoration.md:81.
- 3 (bear Furor 1 adopted): applied, feral-bear.md:56 and :115.
- 4 (rounds 1-9 decisions): unchanged; nothing touched contradicts them.

Round 10 stat rule (every spec page):
- Stat priority prose on balance.md, feral.md, feral-bear.md and restoration.md gives reasons only and points to the table. restoration.md:87 drops "count so much" (an implied ranking). No other page ranks stats or gear by nightly value. Prose test (`-t druid`) passes.

Card notes and other:
- druid-restoration.json top-level notes: the false "rank 1 on the party at every level" now says rank 7 in a raid and rank 1 below 60 and without raid buffs (notes string only).
- Party Rejuvenation lines priorityList[10-13] have no note; they are not new this round (same count at 645caa983) and repeat line 9's, so left.
- Shoulds left: balance.md:91 "always lands" and feral.md:73 "do not pool energy" (carried; pooling waits on Sims requested 2).
- `updated:` already 2026-11-04 on feral.md, feral-bear.md and restoration.md. balance.md and index.md are unchanged.

SIM findings (new or carried):
- Sims requested 1 (Rip vs Bite at 5 points) is still unrun; the prose no longer depends on its result.
- Sims requested 2 (Bite pooling) and 3 (bear race re-check: feral-bear.md:103 Night Elf "slightly less damage, a little more threat" stays PROVISIONAL) are carried.
