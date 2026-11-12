# Sim accuracy 2026-11-12: Arms, Fury and the three hunters against the client and real fights

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main 39898b097 (rebased onto 731209677), fork from master 5f1b39450 (the pin). Client on the site: 1.60.1.70790; live client 1.60.1.70888 with every hotfix push layered (wago.tools `useHotfixes=1`, fetched 12 November; wago still lists 70888 as the newest Forever build). The first full engine-against-client pass for these five specs since round 29 (2026-11-07); the 11 November physical-logs lane compared them with logs only and the hotfixes lane carried Arcane Shot and Summon Hawk by hand.

**Totals: 2 confirmed engine bugs, both old, both fixed: a rage refund paid 80% of the base cost whatever the talents had cut it to (shared core; Arms raid 60 -0.4%, Fury -0.1%), and Resourcefulness cut Lacerate's cost although the client's talent mask leaves Lacerate out (Survival bare 40 -2.2%, bare 50 -1.5%, raid unchanged). Builds and rows under the adoption rule: the Arms level 60 build takes Spearing Strike for a point of Deflection (+1.3 to +1.7% at bare 60 on both seeds, gains at every level from 25); the Beast Mastery level 60 build takes Efficiency 5 and Improved Tracking 4 for Improved Stings, Rapid Killing and Mortal Shots (+1.2 to +1.6%, gains from 48); Serpent Sting rows for both ranged hunters (Beast Mastery rank 1 from 26 to bare 60, +2.3 to +3.9%) and Survival's Raptor Strike rows re-judged on the Lacerate fix. 0 new data: nothing in the 70888 client or the 6 to 12 November hotfixes moves a number these cards use (Sweeping Strikes' charges and Volley's procs change only pack fights). Held: Survival's level 60 build with a pet over Lone Wolf (+0.9% Alliance, under the bar), Summon Hawk (-2.8 to -4.5%). Engine pin 5f1b39450 to 15e2c6bae.**

## Method

1. **Client.** Two read-only audits, one per class, every spell, rank, talent, stance, aspect, pet ability, proc and set bonus the five cards cast or the builds and leveling orders take, line by line: engine (fork worktree at 5f1b39450) against the 70888 hotfixed tables (SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellAuraOptions, SpellCategories, SpellClassOptions, SpellLevels, SpellDuration, SpellCastTimes, SpellProcsPerMinute, SkillLineAbility and the rest) and the site's 70790 `spellconst`, `spelltips` and `talents`. A table diff of 70790 against 70888 and of the 11 November cache against today's tables: the only warrior or hunter rows that moved are Arcane Shot's and Summon Hawk's 11% (already in), Frost Trap's trigger and Volley's attributes (neither cast). The full check tables are `melee-hunters/client-warrior.md` (64 rows aligned) and `melee-hunters/client-hunter.md` (about 115 rows aligned).
2. **Blizzard.** "WoW: Forever Hotfixes" (t/2403040) read to its newest post (post 8, 12 November 00:48 UTC), "Class Tuning Incoming" (t/2404092; Kaivax's only post is the 6 November one) and the patch notes thread (t/2360696, unchanged since 4 November). Warrior and hunter lines since 6 November: the Reckless Swings tooltip names Gore Drinker (a heal); Thunder Clap's cooldown display; "Sweeping Strikes no longer loses charges if no second target is hit"; pets at 115% health; Arcane Shot and Summon Hawk at 11%; the hawk's swing +80%; "Volley was activating procs" and its channel while moving. None moves a single-target number: the engine spends a Sweeping Strikes charge on a single target, where the copy does nothing, and Volley is on no card (both open below for pack fights).
3. **Real fights.** Our API's public feed still holds the two known reports (20 Oct, 8 Nov); the 8 Nov warrior tanked and it has no hunter. foreverlogs.gg: all 25 public reports, 2815 to 2853 (2850 and 2853 new since the physical lane), every warrior, hunter and pet event. Still no Arms, Fury or melee Survival player with a usable fight (Fury fragments at 40 to 42 and Arms-leaning ones at 30 to 36 in 2834 with 17 to 44 swings each; the most melee-heavy hunters are a level 8 to 9 Raptor Strike user and a level 20 to 23 with Aspect of the Monkey). The log analysis is `melee-hunters/logs-foreverlogs.md`.
4. **Sims** on the Windows box only (`tools/remote-test.sh --pin-engine` or `--engine-sha`, one job at a time, `nice -n 15`, GOMAXPROCS=5): `spec-breakdown` at 10,000 iterations (its request limit; seed 1, since `-seed` reproduced seed 1 until 073621f5a landed mid-lane; the two adopted builds were re-run at seed 50001 after it), the talent searches at 60 (raid with `raid-keeps.json`, bare; both factions), the leveling walks at the 18 checkpoint levels for every candidate build, and rank scans: Serpent Sting every level 18 to 60 at every rank to 3 and the newest for Troll and Dwarf (every rank at 42 to 60 for Marksmanship), Raptor Strike 24 to 59 for every hunter race at the row levels and 40 to 59. Job lists, the runner, the walk script and every result table are in `melee-hunters/` and `melee-hunters/results/`.

## Confirmed and fixed

| # | finding | source | origin | fix |
|---|---|---|---|---|
| B1 | **A failed rage attack refunded 80% of the base cost, not of what it paid.** `newRageCost` fixed the refund at registration (`Refund * Cost`), so Heroic Strike under Improved Heroic Strike 3 (12282, -3) refunded 12 of the 12 paid (9.6 is 80%), and Execute under Improved Execute 2 (20502, -5) refunded 12 of 10 paid: a missed Execute gained rage. The energy bar already refunds a share of `CurCast.Cost`. | client 12282 and 20502 (aura 107, misc 14); `core/energy.go` | old | E ca895b36d, `core/rage.go` (`RefundAmount` is the share of the cost paid); goldens E 15e2c6bae |
| H1 | **Resourcefulness cut Lacerate's mana cost.** 440529 effect 0 (aura 108 misc 14, -60) names its spells by class mask: 194 in word 0 (Raptor Strike, Mongoose Bite, Wing Clip, the traps), 524288 in word 1 (Counterattack), 2097152 in word 2 (Strider Kick). Lacerate (24118, 24119, 24120, 1299332) carries 4096 in word 1, which Savage Strikes and Predator's Edge name (528384) and Resourcefulness does not. The engine cut every melee special, so Lacerate cost 16 to 38 instead of 40 to 95. Rounds 12, 14 and 15 ruled from the tooltip's "melee abilities", not the mask. | client 70790 and 70888 SpellEffect, SpellClassOptions | old | E 7bca76b85, `hunter/talents.go` |

Red first on the box: `TestRageRefundIsEightyPercentOfTheCostPaid` (refunds 12, 12 and 12 where 9.6, 11.2 and 8 are due) and `TestResourcefulnessLeavesLacerateAtItsFullCost` (multiplier 40, want 100; it replaces the round 12 test that asserted the cut). Green on the fix; the fork's core, warrior, druid, hunter, conformance and top-level packages pass with `-tags with_db`; the suite goldens that moved are `TestP1DPSWarrior`, `TestP1TankWarrior` and `TestFeralTank` (the bear's Maul and Lacerate refunds are rage too).

**Before and after** (spec-breakdown, 10,000 iterations, seed 1, the published entries; `results/job_e_pin.txt`, `job_e_fix.txt`):

| spec | raid 60 Horde / Alliance | bare 60 Horde / Alliance | bare bands (Horde) |
|---|---|---|---|
| Arms | -0.37% / -0.38% | 0.00% / 0.00% | 20 -0.31%, 30 -0.16%, 40 and 50 0.00% |
| Fury | -0.11% / -0.15% | -0.25% / -0.19% | 20 and 30 0.00%, 40 -0.47%, 50 -0.30% |
| Survival | 0.00% / 0.00% | -0.05% / -0.05% | 20 and 30 0.00%, 40 -2.17%, 50 -1.48% |
| Beast Mastery, Marksmanship | unchanged | unchanged | unchanged |

Protection warrior and the bear move a little in threat with B1 (suite goldens; their published numbers wait for the refresh, MR65).

## Checked, nothing to change

The full tables are in `client-warrior.md` and `client-hunter.md`. In brief:

- **Warrior (64 rows).** Every rank, cost, cooldown, category and GCD of Mortal Strike, Slam (Improved Slam's replacement spells 1310196 to 1310200 included), Heroic Strike (rank 9 held as a book), Rend, Overpower and its 5 s window, Execute and its conversion, Bloodthirst (45% attack power, category 971 with Mortal Strike), Whirlwind and its off-hand strike, Spearing Strike, Reckless Swings, Hamstring, Charge, Bloodrage, Berserker Rage, Recklessness, Death Wish, Battle Shout by level, Sweeping Strikes; stance passives, Tactical Mastery and Improved Tactical Mastery; every talent the three builds and two leveling orders take (Impale's mask reaches Rend and every card ability; Weaponmaster's three effects; Deep Wounds' 20% a rank over 12 s); no Arms or Fury BiS set wears two pieces of any set.
- **Hunter (about 115 rows).** Aimed Shot ranks and its 2.0 s cast hasted by ranged speed; Multi-Shot's one rank and shared category; Arcane Shot at 11%; Serpent Sting ranks plus 5% ranged attack power a tick; Sniper Shot; Hunter's Mark; Rapid Fire on ranged and melee; Aspects of the Hawk and the Beast; Trueshot Aura; Raptor Strike, Mongoose Bite, Lacerate and Lacerating Strikes, Strider Kick, Wing Clip, Immolation Trap with Clever Traps and Survivalist's Discipline; every talent the builds take; Claw and Bite rows, the cat at 1.2 s and its inheritance.
- **One audit item withdrawn.** The hunter audit listed Efficiency's mask as including Lacerate. 19416's word 1 is 0x80084000 (bits 14, 19 and 31), not bit 12, so the engine's Efficiency list, which leaves Lacerate out, is right.

## Logs against the engine

foreverlogs.gg, 25 reports, levels from the ranks cast (`logs-foreverlogs.md`). Aligned with the engine:

| check | logs | engine |
|---|---|---|
| Serpent Sting | a tick every 3.000 s (5,701 ticks, 40 hunters), 5 ticks in 411 of 449 full applications | 5 x 3 s |
| Arcane Shot, Multi-Shot, Aimed Shot cooldowns | shortest gaps 5.94 s, 6.32 s, 7.30 s (6 s plus the cast) | 6 s, 6 s shared |
| Aimed Shot cast | 1.82 s with a quiver, 2.00 without, 1.38 under Quick Shots, 1.29 under Rapid Fire, 1.00 under both | 2.0 s over ranged speed |
| Rapid Fire and Quick Shots | Auto Shot interval x0.708 to 0.717 and x0.762 to 0.777; together x0.549 to 0.553 (they multiply) | 1/1.4 and 1/1.3, multiplied |
| Traps | Immolation Trap 5 ticks of 3 s; shortest recast 31.2 s | 30 s shared |
| Raptor Strike | replaces a swing (3 of 562 with a white hit within 100 ms); shortest gap 5.94 s | next melee, 6 s |
| Mongoose Bite, pet Bite | shortest gaps 4.99 s and 10.0 s | 5 s, 10 s |
| Cat swing | 1.20 s for three cats | 1.2 (`PetAttackSpeed`) |
| Rend, Deep Wounds | 3 s ticks; 12/15/18 s for ranks 2 to 4; Deep Wounds 4 ticks over 12 s | the same |
| Overpower | 64 of 65 within 5 s of a dodge | 5 s window |
| Unbridled Wrath | 1 rage, white swings only (106 of 108), 0.50 to 0.71 a white swing for the warriors who look 5/5 | 12% a rank |
| Slam | 1.50 s cast | 1.5 s |

Not comparable or open: Multi-Shot shows a 0.46 s windup between its cast start and success, hasted like Auto Shot's (client: instant since 70425). The engine casts it instant; inside a 1.5 s global the windup only delays the hit, and autos keep their timer (physical lane L17), so nothing moves. No hawk after the 11 November +80% (one Summon Hawk cast in 2849 at 36 or more with no hawk unit). No two-hand or off-hand rage, no Arms, Fury or melee Survival fight.

## Builds, rotations and rank rows

Talent searches at 60 on the fixed engine (`results/talent-search.txt`; both factions):

| spec | raid 60 | bare 60 (the level 60 build) |
|---|---|---|
| Arms | guide best | clean winner: Spearing Strike for a point of Deflection or Improved Charge, +1.4 to +1.9% |
| Fury | guide best (+0.2%, within error) | guide best |
| Beast Mastery | Improved Stings 3 for Efficiency, +0.6 to +0.7% (under 1%) | Efficiency 5 and Improved Tracking 4 for Improved Stings, Rapid Killing and Mortal Shots, +1.4 to +1.6% |
| Marksmanship | guide best | a deep Beast Mastery build (+1.2%), which drops the spec's tree |
| Survival | guide best | a pet over Lone Wolf (+1.0 to +1.8%) |

**A1, Arms level 60 build: adopted.** `12323213132215201/0505/205`, Spearing Strike for the third point of Deflection (Deflection 2 and Iron Will 5 stay, so no unmodelled talent goes; the rotation already casts Spearing Strike). Bare 60, Horde and Alliance: +1.36% and +1.65% at seed 1, +1.31% and +1.36% at seed 50001. The walk (`results/walk-*.txt`) takes Deflection at 14 and 17, Improved Charge at 18 and 19, Spearing Strike at 27, and gains at every checkpoint from 25: +0.6 to +1.1% at 25, +5.0 to +5.8% at 28, +3.2 to +4.1% at 30 to 38, +2.4% at 40, +0.8 to +1.1% at 43 to 58. The Improved Charge variant gains less from 43. Raid build unchanged (it already takes Spearing Strike).

**A2, Beast Mastery level 60 build: adopted.** `5420001505001251/0050550001/4`: Efficiency 5 and Improved Tracking 4 (4% damage on the tracked type) for Improved Stings 2, Rapid Killing 1 and Mortal Shots 2, all modelled. Bare 60: +1.55% and +1.28% at seed 1, +1.41% and +1.17% at seed 50001. The walk changes from 46 and gains at every checkpoint from 48 (+0.6 to +1.5%). Raid build unchanged (its winner is under 1%).

**Held.**
- **Survival with a pet at 60.** Dropping Lone Wolf for the cat gains +1.78% Horde but +0.93% Alliance at bare 60; under the bar. The guide already tells a world player to keep the pet out.
- **Summon Hawk for Beast Mastery (MR44).** With the hotfixed hawk and the card casting Summon Hawk in Arcane Shot's slot (always, or only under two hawks), raid 60 loses 2.8 to 4.6% and bare 60 4.0 to 6.2%: the 3 November Arcane Shot (15% of ranged attack power) out-hits the dive and the hawk's swings together. Settled; no build takes it.
- **Marksmanship deep Beast Mastery** leaves the spec's tree.

**R1, Serpent Sting rows (MR44's leveling mana lines).** With Arcane Shot at 11% of base mana on every rank, every sting tick's 5% of ranged attack power makes a low rank keep most of the damage. Beast Mastery: rank 1 from 26 to 59 and at bare 60, +2.3 to +3.9% over the newest at every level for Troll and Dwarf; below 26 the newest rank wins (the old rank 1 rows at 18 to 25 lose 0.1 to 1.1%, and the rank 2 row at 38 and 39 loses to rank 1). Marksmanship: rank 1 at 18 (+0.1%), 26 to 32 and bare 60 (+3.8 to +4.3%), rank 3 at 42 to 49 (+1.6 to +3.2%), rank 4 at 50 to 59 (+2.1 to +3.3%); at 33 to 41 Rapid Recuperation keeps the newest rank ahead. In a raid the top rank wins by about 4.5% for both (Judgement of Wisdom; `results/job_k.txt`), so the rows are bare only at 60.

**R2, Survival's Raptor Strike rows on the Lacerate fix.** Every hunter race at 26 to 29, 31 and 40 to 59, the newest rank against the three below it; a row is kept where a lower rank gains 0.2% or more (one seed, so twice round 21's bar) and sits within 0.15% of that race's best (`rows_rs.py`). Horde: rank 2 at 26 to 28, 3 at 29, 4 at 40 to 45, 3 at 46 and 47, 4 at 48 and 49, 5 at 50 and 51 (a Tauren to 52), 5 at 56 to 59. Alliance: rank 3 at 26 and 27 (28 for a High Order Skyborne), 5 at 40 (a Dwarf to 45), 4 at 46 and 47, 5 at 48 (a Dwarf to 51) and at 56 to 59 (a Human only at 59). Gains 0.2 to 2.6%. Rank 1 at 31 is gone (it loses 0.1 to 0.8%), as are the old rank 5 rows at 46 to 49 (rank 3 or 4 wins now).

Arms and Fury have no rank rows. Leveling orders: the round 29 walks stand; the two adopted builds move the Arms walk from 17 and the Beast Mastery walk from 46 (`addon-builds.json` regenerated on the box).

## Guide and card

wow-player read the four guides, the three cards and two coach notes in full: warrior SHIP after two SHOULDs (both builds take Spearing Strike; place the level 20 points); hunter NO-SHIP, then SHIP after four MUSTs (the Survival rank list split into Horde and Alliance lists that name races and levels and match the card; a raid caveat on the Beast Mastery sting line) and SHOULDs (Efficiency bullet, no sim-style sentence, reasons on the Marksmanship rank lines, the coach note points to the guide's list).

- **Arms:** build line, the level 60 and raid paragraphs (Spearing Strike in both), the 1 to 19 list (Deflection 14 and 17, Improved Charge 18 and 19), the 20 to 60 paragraph; card and coach: Spearing Strike from 27.
- **Beast Mastery:** build line, talents (Efficiency, Improved Tracking, Mortal Shots raid only), the raid paragraph, the sting rank list; card: Rapid Killing gone from the Rapid Fire note.
- **Marksmanship:** the sting rank list.
- **Survival:** Resourcefulness names its spells and leaves Lacerate at full price (strengths, mana paragraph), the Raptor Strike rank lists; card and coach Raptor lines; the coach's stale Counterattack line (round 29 dropped Counterattack from the build) is gone.

## Open

- **Immolation Trap ignores hit.** The engine strips the hunter's whole Hit before the trap rolls (an upstream Era workaround). Forever merges spell, melee and ranged hit into one stat and Surefooted (19290) carries spell hit; whether a Forever trap uses it is server side. Needs a trap miss rate from a log with known hit (about 0.3 to 0.6% of Survival).
- **What spends the Mongoose Bite window.** 5302 and 1310726 have ProcCharges 1 and ProcTypeMask 16 (any melee ability); the engine lets only Mongoose Bite spend it. A log showing the window removed by a Raptor Strike would cut Survival's bites sharply.
- **Pets and Expertise** (patch notes: "Pets now correctly benefit from Expertise"): the pet inherits none; no hunter gear carries Expertise today.
- **Warrior server rules** with no client value: a failed Execute's extra rage, Anger Management's tick (aura 85 base 15), two-hand and off-hand rage, Reckless Swings and Unbridled Wrath on a landed swing only, Bloodthrill from Rend and Overpower themselves, Weaponmaster's sword extra-attack cooldown.
- **Pack fights:** Sweeping Strikes keeps its charge without a second target (10 Nov) and Volley no longer procs (11 Nov). The engine still spends the charge and gives Volley a spell proc mask; neither touches a single-target number.
- **No Arms, Fury or melee Survival log, no level 60 log, no hawk after 11 November.**

## Measure at refresh

- **MR61** Arms on the new level 60 build: warrior-arms bare bands 20 to 60 (talents move from 17); in `knownStaleBISFiles`.
- **MR62** Beast Mastery on the new level 60 build and sting rows: hunter-beast-mastery bare bands 50 and 60 (talents from 46), bands 30 to 60 on the rank 1 sting; in `knownStaleBISFiles`.
- **MR63** Survival on the Lacerate cost and the new Raptor rows: hunter-survival bare bands 30 to 60, picks and weights.
- **MR64** Marksmanship's sting rows: hunter-marksmanship bare bands 20 to 60.
- **MR65** the rage refund: warrior-arms, warrior-fury raid and bare; warrior-protection and druid-feral-bear threat and weights.
- MR44 is settled for its sting rows and the hawk (this lane); its Beast Mastery builds part is MR62. No 70888-only number came out of this lane.

## Sim changes in this landing

- **Fork `lane/sim-melee` on master 5f1b39450:** b7f6d3df4 (tests, red), ca895b36d (rage refund), 7bca76b85 (Resourcefulness), 15e2c6bae (suite goldens). Master fast-forwarded to 15e2c6bae.
- **Site pin:** 15e2c6bae.
- **Site:** guide frontmatter builds (Arms, Beast Mastery), `rank-choices.json` and its embedded copy, `addon-builds.json` and the seven ladder goldens regenerated on the box, `knownStaleBISFiles`, the hunter rank-row tests (the superseded round 19, 21, 22, 24 and 25 cases replaced by `rank_choices_hunter_1112_test.go`), the Arms walk tests, the low-mana test's floor (a Beast Mastery hunter at 18 now keeps the newest sting), the rotation-lists test (Spearing Strike is in both Arms builds), cards' notes (notes only, so no `apl-sync`), coach notes, four guides, this report and its scripts and tables.
- **Shared code:** `core/rage.go` (every rage spender: warrior, bear).

## Sources

- Blizzard: t/2403040 (posts 1 to 8), t/2404092, t/2360696.
- wago.tools DB2 CSVs for 1.60.1.70790 and 1.60.1.70888, `useHotfixes=1`.
- foreverlogs.gg reports 2815 to 2853; foreversixty API `/v1/reports/recent`.
