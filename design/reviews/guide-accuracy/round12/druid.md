# Guide accuracy round 12: Druid

Auditor: druid persona. Tree: gr12/audit at a8b147a6f (main after round 11 and guide-header-build, READ-ONLY), client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go). Before starting I read the round 1 to 11 druid reports (with their fix logs), round4/druid-levels.md and ledger.md.

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`), and the card notes in all four druid APLs.

## What moved since round 11 (d54de48aa to a8b147a6f)

- 91da94f1e (round 11 fixes) changed only feral.md:67 (R11-1) and feral.md:119 (Rake and Ferocious Bite "from 24").
- 3bd24d3a2 (guide-header-build): on Feral and Balance, the header build card, the key-talent chips and the rotation now follow `raidBuild`. Restoration and bear have no raidBuild, so their pages are unchanged. The commit changes only components.
- addon-builds.json, talents, trainables, spelltips, rank-choices.json and the druid bis/*.json did not change. No druid build string changed.

## Decoded builds (all legal, 51 points, tier gates and prerequisites met)

Unchanged from round 11:

- Balance leveling `5232220115101151/0/555301`, raid `5232220115501351/0/505003`.
- Feral leveling `050022/50212032123012123001/05501` (9/31/11).
- Feral raid `050022/50230032023012125001/055` (9/32/10).
- Bear `0/5123033202310212551/055` (0/41/10).
- Restoration `05302001/0/555300515311321` (11/0/40).

Feral tier check, raid build:
- Tier 0: Ferocity 5.
- Tier 1: Feral Swiftness 2 and Feral Instinct 3. Ten points, which opens tier 2.
- Tier 2: Shredding Attacks 3, Savage Fury 2 and Sharpened Claws 2. That makes 17 points, two more than tier 3 needs.

Leveling to raid moves Brutal Impact 2 (tier 1), Feral Charge 1 (tier 2) and Natural Shapeshifter 1 into Feral Instinct +2 and Rend and Tear +2.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| DRUID-R12-1 | CONFIRMED | talent (round 11 edit) | feral.md:67 | "The extra Feral Instinct points add nothing to a cat's damage; they are there because the raid build drops Brutal Impact and Feral Charge, and something has to fill that row to open the next one." | Decoded raid string (above). Feral Instinct and Brutal Impact are tier 1 (talents/druid.json tier 1, columns 1 and 2). Feral Charge is tier 2, column 2. Dropping Brutal Impact's 2 points leaves tier 1 at 8 of the 10 that open tier 2, so 2 points have to go back into that row. Dropping Feral Charge leaves tiers 0 to 2 at 17 points, two more than tier 3 needs, so it frees a point and costs no gate. That point went to Rend and Tear. The sentence blames Feral Charge for a gate it does not hold. | "The extra Feral Instinct points add nothing to a cat's damage; with Brutal Impact gone, they are the two points that keep that row at ten and open the next one." |
| DRUID-R12-2 | CONFIRMED | rotation (card note) | data/curated/apl/druid-restoration.json priorityList[7] notes (also the file's top-level notes) | "A raid fight's mana pays for the rank that heals the tank most; below 60 and without raid buffs, rank 4 keeps your mana going." | rank-choices.json, druid-restoration, tank Rejuvenation (authored 9841, rank 10). Its only band is `from_level 28, to_level 49, rank 4`. It has no 50 to 59 row and no bare 60 row, so from 50 to 59 and at bare 60 the line casts the top rank you know (rank 8 at 46, rank 9 at 52, rank 10 at 58, from trainables/druid.json). Rank 4 is learned at 22 and is not cast below 28. The page has this right (restoration.md:81, "on the tank from 28 to 49 a middle rank such as rank 4"), so the card note contradicts both its own rank rows and the prose. | "On the tank whenever it is not ticking, at rank 10 in a raid. Let it run out before you recast it, so Bloom heals him when it does. A raid fight's mana pays for the rank that heals the tank most; from 28 to 49, rank 4 keeps your mana going." In the top-level notes, change "rank 4 below 60 and without raid buffs" to "rank 4 from 28 to 49". |
| DRUID-R12-3 | LIKELY | talent (build) | feral.md:7 raidBuild, :67 | The raid build spends 2 points in Feral Instinct, which the page itself says "add nothing to a cat's damage". | Furor (tier 0) and Thick Hide (tier 1, the same row) open tier 2 just as well. Furor 2/5 (talents/druid.json rank 2): "Shapeshifting into Cat Form restores 40% of the Energy you had when last in Cat Form, plus 4 Energy per second spent outside Bear Form or Dire Bear Form, up to 40 Energy." The same paragraph tells the raid cat to shift out to Innervate, battle rez and decurse, and says "Furor saves the energy you would lose". Thick Hide gives armor in Cat Form. Feral Instinct 3 is Swipe (a bear ability) and Prowl detection. Both alternatives are at least as good as Feral Instinct, and both are worth something. The sim does not shift out of Cat Form, so it should read them as equal. | If Sims requested 1 comes back level within error, change raidBuild Feral to `52210032023012125001` (Furor 2, Feral Instinct 1). Then replace the R12-1 sentence with: "Two points of Furor fill the row that Brutal Impact leaves; every Innervate, battle rez or decurse then brings you back to Cat Form with energy." Rewrite "The raid build skips Furor" as "The raid build takes two points of Furor". |

I checked every other line against 70755 and found nothing wrong. The checks that matter most this round:

**Header card against prose (guide-header-build)**
- On Feral and Balance the header card, the key-talent chips and the Rotation section now show the raid build.
- Every Feral and Balance sentence that is about the leveling build says so ("The leveling build takes ...", "Your 31 Feral points are done at 40", the Leveling section).
- Every bolded talent with a point count has the same count in both builds. This covers Shredding Attacks 3, Savage Fury 2, Predatory Strikes 3, Predatory Discipline 1, Shifting Power 1, Leader of the Pack 1, Vengeance 5, Moonfury 5, Improved Wrath 5, Nature's Reach 2, Moonglow 3 and Moonkin Form 1.
- Talents that differ are worded per build: Rend and Tear "15% at 5", Eclipse "At full ranks".
- The Feral and Balance rotation prose uses only raid-build talents.
- No prose contradicts the header card.

**Round 11 fixes**
- feral.md:67 "Natural Shapeshifter only makes your shifts and Shifting Power cheaper" matches "Reduces the mana cost of all shapeshifting by 30%". The causal clause after it is wrong (R12-1).
- feral.md:119: Rake and Ferocious Bite both train at 24 (trainables).

**Client values re-derived on 70755 (spelltips/druid.json, talents/druid.json)**
- Cat abilities:
  - Shifting Power 1322605 costs 55% of base mana, has an 8 s cooldown and gives 40 energy.
  - Poised: 10 energy at full energy, out of combat.
  - Ravage and Pounce: "Must be prowling and behind the target".
  - Rake is 40 energy, 35 with Ferocity.
  - Claw is 45 energy, 40 with Ferocity. Shred is 54 at 20, 48 at 21 and 42 at 22, so the card's "8 to 14 more" holds.
- Bear abilities and cooldowns:
  - Enrage 5229: 10 rage up front plus 20 over 10 s, an armor penalty, a 60 s cooldown.
  - Frenzied Regeneration 22842: 10 rage a second, 1% health a rage. 100 rage heals your whole health bar.
  - Barkskin 22812: 20% physical, 15 s, 1 min.
  - Primal Bite 1238070: 20 rage (15 with Ferocity), 6 s.
  - Lacerate: 15 s, 5 stacks.
  - Swipe: 3 targets.
  - Challenging Roar 5209 is real.
  - Natural Reaction 5% plus Feral Swiftness 4% gives 9% dodge.
- Healing:
  - Tranquility 9863: every 2 s for 10 s.
  - Wild Growth 1238215: 3.0 s cast (2.5 s with Naturalist), 6 s cooldown, 7 s duration.
  - Improved Tranquility 60% takes Tranquility from 5 min to 2 min.
  - Improved Regrowth is 10% a point.
  - Swiftmend 18562: 20% of base mana, 15 s.
  - Innervate 29166 is +400%.
  - Rejuvenation ranks 1 to 11 are at 4/10/16/22/28/34/40/46/52/58/60. Regrowth rank 9 is at 60.
- Talents:
  - Genesis covers "periodic damage and healing done by your spells and abilities" (so Rip).
  - Naturalist is 10% damage and 0.5 s off Healing Touch and Wild Growth.
  - Heart of the Wild: 10% Int, 20% Stamina in bear, 10% Strength in cat.
  - Furor rank 1 is 20% for 10 rage.
  - Thick Hide text.
  - Nature's Reach 4% hit and 20% range.
  - Nature's Splendor: Rejuvenation +3 s, Regrowth +6 s.
  - Eclipse: 2 charges a Wrath, 0.5 s each, stores 4.
  - Nature's Grace: non-periodic spell crits, 3 s.
  - Moonkin: 45 yd, 3% crit, Omen +100%, no healing spells.
  - Omen of Clarity: "not consumed by Wrath".
  - Leader of the Pack: all crit, 45 yd.
- Races (races.json and earlier rounds' spelltips checks): Elune's Light 10% / 15 s / 3 min, Quickness 1% dodge, Endurance 5%, War Stomp 5 targets 2 s, Read Leyline 50% / 15 s / 2 min with 60 min Attuned (Alliance Skyborne), Wind Blessed 2% haste. Skyborne is named only as the paid option.

**Card notes**
- Balance: all seven notes match the raid build and the client values.
- Bear: all eight notes match. Demoralizing Roar is 204 at 60 (scaling, decided in rounds 2, 5 and 7).
- Feral: all notes match (Claw 40 against Shred 54 at 20 and 48 at 21).
- Restoration: every note except [7] (R12-2) matches.

**Leveling** (addon-builds.json unchanged, orders as in round 11)
- Every Levels 1 to 19 paragraph matches its order and the trainer levels.
- The Levels 20 to 60 text that follows does not contradict them.

**Sources**
- Frontmatter links checked on 2026-11-04 with a browser user agent. These returned 200:
  - both Blizzard articles
  - talentsforever.com/data.json
  - /sim/weights, /sim/drops, /bis/druid/feral-bear, /bis/druid/restoration
- The Wowhead Skyborne First Look page returned 403. Wowhead refuses automated fetches, so this is unknown, not dead. Round 11 had a 200 for it.
- All entries still follow the round 4 source-link rule.

**Voice**
- No em or en dashes. No site, sim, APL, preset or band words in prose. No quoted nightly numbers.
- Stat paragraphs explain the stats and do not rank them.

## Shoulds (wording, not findings)

- feral.md:77 "at 60 energy or less, so all 40 fits under the cap": with Wolfshead Helm (feral.md:93) a press gives 45, so the threshold is 55. Suggest adding "(55 with Wolfshead Helm)".
- feral.md:85 "Your threat is high with Leader of the Pack crits and Rend and Tear": Leader of the Pack is a crit aura, not a source of crits. Suggest "Your crits and Rend and Tear make your threat high."
- balance.md:91 "it always lands on a Starfire or a DoT": carried. "almost always" is exact.

## Decisions

1. **A round 11 fix sentence is held to the tier math** (R12-1). Only Brutal Impact's tier-1 points have to be replaced. Feral Charge sits in tier 2, which has spare points.
2. **The filler row is a build choice, so it is measured** (R12-3). If the sim reads it as a tie, take Furor 2. The page already explains that shifting out costs the cat its energy, and Furor is the only point that pays that back. Thick Hide is the fallback if the integrator wants no prose change beyond R12-1.
3. **Card rank notes follow rank-choices.json bands exactly** (R12-2). The prose is already right.
4. **Rounds 1 to 11 decisions stand:**
   - Swiftmend keeps the HoT.
   - Gift of Nature misses Tranquility.
   - Moonkin casts Rebirth, Innervate and Remove Curse.
   - Pounce and Ravage need you behind the target.
   - Cat Swipe does not exist in Forever.
   - Bear Furor is 1 point.
   - Feral Charge (Cat) 414924 is the 30 s leap.
   - Nature's Reach covers melee hit.

## Sims requested

1. New (R12-3): druid-feral raid build `050022/50230032023012125001/055` against `050022/52210032023012125001/055` (Furor 2, Feral Instinct 1). Also run Thick Hide 2 (`050022/50210232023012125001/055`). Pin 9f7f1f6e4, band 60 raid and bare, seeds 1 and 50001. Expect all three within error. If they are, adopt Furor 2 per Decision 2.
2. Carried, R10-12: druid-feral Rip against Bite at 5 points, band 60 raid and bare, seeds 1 and 50001. The prose no longer depends on the result.
3. Carried: druid-feral Bite energy pooling under Predatory Discipline (the shipped list against holding the 5-point Bite until 80 energy or more).
4. Carried: bear race re-check on `0/5123033202310212551/055`, Night Elf against Tauren (threat and chance of death, raid 60, seeds 1 and 50001). feral-bear.md:103 "slightly less damage, a little more threat" is PROVISIONAL.
5. PROVISIONAL BiS and stat order: every druid bis/*.json in the tree is still engine d1d9f2775 (generated 2026-11-04 02:22 to 02:33 UTC), from before the final re-rank. After the re-rank, check the statPriority frontmatter on all four specs and that the stat prose still reads with the table.
6. Carried S10: feral leveling, Claw to Shred gate (Shredding Attacks 3/3 at 22 against the Ferocious Bite rank 5 gate), ladder 20 to 25.

## Sim findings

- None new. Carried: S11 bear expertise weight, S12 bear Defense noise, S13 cat abilities and Starfire on 70753 values (the Shred spelltip on 70755 reads 200%; the engine notes say 155%).

## Sources

- Client 1.60.1.70755:
  - talents/druid.json: tiers, columns and rank texts for Furor, Feral Instinct, Thick Hide, Brutal Impact, Feral Charge, Nature's Reach, Natural Shapeshifter and every talent the pages name.
  - trainables/druid.json: Rejuvenation and Regrowth rank levels.
  - spelltips/druid.json: 1322605, 1328829, 9830, 9850, 9867, 9827, 9904, 9896, 1328618, 5229, 22812, 22842, 1238070, 1328617, 9908, 5209, 9863, 1238215, 18562, 29166, 16864, 17402, 21850.
  - races.json.
  - addon-builds.json.
  - bis/druid-*.json (engine_version d1d9f2775).
- data/curated/apl/druid-*.json (card notes); data/curated/rank-choices.json (druid-restoration tank Rejuvenation band 28 to 49).
- Commit 3bd24d3a2 (guide-header-build), for what the header card and Rotation now show.
- Blizzard, "World of Warcraft: Forever Class Deep Dives: Hunter and Druid", https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (200).
- Blizzard, "World of Warcraft: Forever Deep Dive Panel Recap", https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200).
- Wowhead, "Skyborne First Look", https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829 (403 to automated fetch, unknown).
- Talents Forever, https://talentsforever.com/data.json (200).
- Forever Sixty: /sim/weights, /sim/drops, /bis/druid/feral-bear, /bis/druid/restoration (all 200).

## Fix log (round 12)

- DRUID-R12-1: fixed, feral.md:67. Verified talents/druid.json: Feral Instinct and Brutal Impact are tier 1, Feral Charge tier 2. The sentence now names only Brutal Impact: "with Brutal Impact gone, they are the two points that keep that row at ten and open the next one."
- DRUID-R12-2: fixed, data/curated/apl/druid-restoration.json priorityList[7] note ("from 28 to 49, rank 4 keeps your mana going") and the top-level notes ("rank 4 from 28 to 49"). Verified rank-choices.json druid-restoration Rejuvenation 9841: one band, 28 to 49, rank 4. No action changed.
- DRUID-R12-3: left; needs the Furor/Thick Hide raid-build sim (Sims requested 1).
