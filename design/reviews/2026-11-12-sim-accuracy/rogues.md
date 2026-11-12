# Sim accuracy 2026-11-12: Assassination, Combat and Subtlety against the client and real fights

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-rogues`. Site from origin/main 39898b097, fork at the pin 5f1b39450. Client on the site: 1.60.1.70790; live client 1.60.1.70888 with its hotfix push. Last rogue passes: round 23 against the client (2026-11-05), rounds 26, 27 and 29 on orders and walks, and the 2026-11-11 physical-logs lane against logs only (`2026-11-11-sim-accuracy/physical-logs.md`, nothing found).

**Totals: 0 confirmed bugs. No new data: between 70790 and the hotfixed 70888 the only rogue rows that change are Poised's (a new enter-combat trigger to two server dummies, no number), and Blizzard's posts since 6 Nov carry two Shank tooltip lines only. One standing open item is closed by a log: Backstab's flat bonus sits inside its 150%, as the engine has it (round 23's "Backstab's flat bonus on a live log"). Every engine row checked against the client matched, the log checks aligned, 1 search candidate held under the adoption rule. No card, build, order, rank row or guide line changes. Engine pin unchanged: 5f1b39450.**

## Method

1. **Blizzard since round 29.** "WoW: Forever Hotfixes" (t/2403040) read through post 8 (12 Nov 00:48 UTC; post 1 last edited 12 Nov 00:47), "Class Tuning Incoming" (t/2404092, 1,294 posts, one Blizzard post: the 6 Nov opener, no rogue line) and the development notes (t/2360696, last post 4 Nov, read by round 23).
2. **Client and push.** Fresh wago.tools CSVs for 1.60.1.70888 with `useHotfixes=1` (SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellAuraOptions, SpellCategories, SpellClassOptions, SpellLevels, SpellName, Spell, SpellDuration, SpellProcsPerMinute, Talent, SpellEquippedItems, SpellAuraRestrictions, SpellCastingRequirements; fetched 12 Nov) against the 70790 hotfixed set. Rogue rows are every SpellClassSet 8 spell, every rogue talent rank, and two hops of their trigger spells: 506 spells. Script: `rogues/logs/client_diff.py`.
3. **Engine against the 70790 tables**, every spell, talent, poison, proc and set bonus the three cards cast or the five guide builds take (`sim/rogue/*.go`, `sim/common/item_sets` for Stormshroud and Defias) against `data/builds/1.60.1.70790/spellconst/rogue.json`, `talents/rogue.json`, `spelltips/rogue.json` and the raw SpellEffect and SpellAuraOptions rows.
4. **Real fights.** Our API: `/v1/reports/recent` still lists only gxrgvz4yzsi2 (8 Nov, used by the physical lane; its one rogue never cast Backstab) and z2opt2x2gs5u (20 Oct). foreverlogs.gg: 25 public reports; the ones after 8 Nov are 2836 to 2853. 2836 to 2849 come from the physical lane's event cache (every event of every rogue); 2850 and 2853 were fetched here (rogue players, `meta_events`, paced). Rogues with 30 or more casts after 8 Nov: Sugo (2839, Combat at 30 to 37: Sinister Strike rank 5, Shank rank 1, Blade Flurry, 1,348 casts), Jaice (2847) and Nakir (2848) at 16 to 21, Hermes (2850, Slice and Dice with Improved Slice and Dice). Nobody at 40 or more casts Mutilate, Hemorrhage, Ghostly Strike or Death Mark yet. For the Backstab question the cache's older Wailing Caverns rogue Katrina (2818, 6 Nov) is used, with her armory capture (`combatants-info`: 14 talent points, 0/2/12, Scrimshaw Dagger 11 to 21 at 1.8). Scripts: `rogues/logs/`.
5. **Searches** on the Windows box (`tools/remote-test.sh --dir sim --pin-engine`, engine 5f1b39450, one at a time, `nice`, GOMAXPROCS=5): `talent-search` for each spec at raid 60 (raid build) and bare 60 (leveling build), keep lists holding every talent the card casts or leans on. Script `rogues/search.sh`, output `rogues/search.out.txt`.

## Blizzard since round 29

| post | line | verdict |
|---|---|---|
| 6 Nov, Rogue | "Shank will now display Rank 1 and Rank 2 in the tooltips." | tooltip |
| 6 Nov, Rogue | "While Boot Knife is talented, Shank will now correctly display the number of Combo Points it grants." | tooltip; it agrees with the engine, which awards the trigger row's 2 (1329285 effect 30, 2 points) where Shank's old tooltip said 1 |
| 6 Nov, Items | "Searing Dagger is now only activated by spells that have a damaging component." | in no rogue BiS file |
| 11 Nov, Items | "Bloody Claws of Attack now correctly add +1 Fist/Unarmed Weapon skill." | in no rogue BiS file |
| 6 Nov, Classes | "The Quick and the Dead Legacy Talent" | not a rogue talent |
| 9 to 11 Nov | Judgement as a ranged attack, Volley procs, shaman, hunter, warlock, mage, druid lines | no rogue row (the totem-reach change, a8a54522e, gives a rogue nothing new: poisons never shut out Windfury Totem) |

## Client: 70790 against 70888 (hotfixed)

Of the 506 rogue spells, only Poised's rows differ:

| spell | 70790 | 70888 |
|---|---|---|
| 1328825 Poised (passive) | e0 aura 4 (dummy), bp 10; no aura options | e0 aura 42 (proc trigger) to 1330741; SpellAuraOptions ProcChance 100, ProcTypeMask 0x08000000 (enter combat) |
| 1330741 Poised | absent | new: aura 4 dummy, no cost, rogue class set |
| 1330746 Poised | absent | new: aura 226 (periodic dummy) every 500 ms |
| 1328826 Poised (the buff) | Attributes_0 262160 | Attributes_0 gains 0x80000000; the -10 Energy on the same mask is unchanged |

Both new spells are server dummies, the description text is unchanged ("While at maximum Energy, and out of combat, reduce the Energy cost of your next ability by 10"), and no note mentions Poised. The cards spend the buff on the pull's opener, the cast that enters combat, so no number moves on either reading. The logs agree with the engine's out-of-combat model: of 863 Poised gains and refreshes in the cache, 860 come with no damage from the rogue in the 3 seconds before. Recorded as MR60 for the client move.

## Engine against the client (1.60.1.70790)

Every row matched.

| area | engine | client | verdict |
|---|---|---|---|
| Backstab | 150% of normalized weapon damage plus the rank's flat, the flat inside the 150%; 60 Energy; dagger, behind | 53 to 11281: effect 121 amount 10 to 140, effect 31 at 150; the tooltip prints `$s1` | aligned (and see L1: the flat is multiplied) |
| Sinister Strike | normalized weapon damage plus 3 to 68, 45 Energy | 1752 to 11294: effect 121 3 to 68 | aligned |
| Ambush | 250% of normalized weapon damage plus 28 to 116 inside it, 60 Energy | 8676 to 11269: effect 121 28 to 116 under 250; tooltip 70 and 290 (= x 2.5) | aligned |
| Mutilate | both hands, 90% of normalized weapon damage plus 44/63/88 inside it, off-hand weapon damage halved with Dual Wield Specialization, +20% on a poisoned target, 2 points, 60 Energy, ranks at 40/50/60 | 399960/1241585/1241586 and twins: effect 121 44/63/88, effect 31 90; 399956 e3 20; tooltip `${m1*m2/100}` = 79 at rank 3 | aligned |
| Shank | off hand, 50% weapon damage, 5 Energy, 10 s, 15/25 Energy back, cannot miss, procs nothing, ranks at 24/44 | 1329283/1329549 | aligned |
| Eviscerate | roll plus 5/11/19/31/45/71/110/151 a point plus 3% of attack power a point, 35 Energy | 2098 to 11300; rank 8 tooltip 199-295 at one point and 803-899 at five (96 x (1 +/- 0.5) + 151 a point) | aligned |
| Rupture | (35 + 4.73 a point) a 2 s tick at rank 6, 3 + points ticks, 25 Energy | 11275; tooltip 159 over 8 s to 469 over 16 s | aligned |
| Garrote | 24 to 92 a 3 s tick for 18 s plus 3% of attack power a tick, 50 Energy | 703 to 11290; rank 6 tooltip 552 over 18 s | aligned |
| Slice and Dice | 20% (rank 1) and 30% (rank 2, 42) attack speed, 6 + 3 s a point, 25 Energy | 5171/6774 effect 319 20/30; tooltip 9 to 21 s | aligned (and L3) |
| Hemorrhage, Ghostly Strike | 100%/145% and 125%/180% with a dagger, normalized, 35 and 40 Energy, Ghostly 20 s cooldown and the 7 s opener aura 1329326; Hemorrhage +15% to Rupture and Garrote for 15 s | 16511, 14278, 1329326 | aligned |
| Death Mark | 12% a point stored, bleeds twice, 6 s, 10 s cooldown, 25 Energy; cap (156 + 245 a point + 4% of attack power a point) x (1 + crit) at rank 4, 63/98/138 and 102/158/217 below; ranks at 40/48/56/60 | 1316709 to 1325758; rank 4 tooltip 401, 646, 891, 1136 | aligned |
| Instant Poison | 15 to 88 by rank plus 1% of attack power, 20%, magic table | 8680 to 11337 rolls; 3 Nov notes 0.01 | aligned |
| Deadly Poison | 6/9/14/18 a 3 s tick a stack for 12 s, 5 stacks, plus 0.9% of attack power a tick, 30% | 2818/2819/11353/11354; 3 Nov notes 0.009 | aligned |
| Adrenaline Rush, Blade Flurry, Cold Blood, Premeditation, Preparation | 100% regeneration for 15 s on 5 min; 20% attack speed for 15 s on 2 min at 25 Energy; +100% crit on the next Sinister Strike, Backstab, Ambush, Eviscerate or Mutilate, 3 min; 2 points for 20 s on 2 min; 10 min | 13750, 13877, 14177 (mask 131590 + mask_1 6), 14183, 14185 | aligned |
| Malice, Precision | 1% crit a rank on every table (aura 290); 1% hit a rank, melee and spell (aura 54 and 55) | Forever merges hit into one stat and `stats.Crit` is every table, so the poisons get both | aligned |
| Ruthlessness, Relentless Strikes, Seal Fate | 20% a rank; 4% a point a rank for 25 Energy (14181); 33/66/100% with a 500 ms cooldown | 14156 (trigger 14157), 14179 (aura 109 on the finisher mask), 14186 ProcCategoryRecovery 500 | aligned |
| Lethality, Aggression, Coup de Grace, Opportunity | 10% crit damage a rank on Sinister Strike, Gouge, Backstab, Mutilate, Ghostly Strike, Hemorrhage; 5% a rank on Sinister Strike and Eviscerate; 7/13/20% on Eviscerate and Death Mark; 4% a rank on Backstab, Ambush, Mutilate and Garrote's ticks | 14128 mask 100663310 + mask_1 6; 18427 mask 131074; 14162 e0 and e1; 14057 e0 and e1 | aligned |
| Puncturing Wounds, Hack and Slash | Backstab +10% crit and 15% for a second point a rank, Mutilate +5% crit a rank; daggers and fists +1% melee crit a rank, swords and axes 1% a rank for an extra attack (200 ms cooldown), maces 3% of armor a rank | 1224716 e0 to e2; 13960 ProcCategoryRecovery 200 | aligned |
| Vile Poisons, Improved Poisons, Venom | +4% poison damage a rank; +2% poison chance a rank; Shank gives +30% (additive with Vile Poisons) and +20% chance for 6 s | 16513, 14113, 1329328 to 1310703 | aligned |
| Flash of Steel, Boot Knife | +3 Energy and -0.4 s on Shank a rank; 20% a rank for 2 points | 1329284, 1329286 to 1329285 | aligned |
| Composure, Flawless Execution, Efficient Strikes, Dirty Deeds | Poised -5/-10 more; Eviscerate -15; Sinister Strike and Hemorrhage -3/-5; Garrote -10 a rank | 1329265, 1310711, 13732, 14082 | aligned |
| Serrated Blades, Blood Splatter, Improved Slice and Dice, Initiative, Weapon Expertise, Dual Wield Specialization | 3% of armor a rank and Rupture +10% a rank; Rupture +6 s, Garrote +9 s; +15% a rank; 33/67/100%; 1% a rank; +25% off hand a rank | 14171, 1329256 (trait curves 6000/9000), 14165, 13976, 30919, 13715 | aligned |
| Stormshroud 2 and 3, Defias 4 and 5 | 10% for 15 to 25 Nature; 4% for 30 Energy; the bleed and +1 dagger skill | sets 142 and 161 (the only rogue set bonuses in any BiS band) | aligned |
| Rank rows | none | every rank of every rogue spell costs the same Energy, so the newest rank always wins and the site carries no rogue rank rows | nothing to measure |

## Engine against real fights

| # | check | log | engine | verdict |
|---|---|---|---|---|
| L1 | **Backstab's flat bonus is inside its 150%.** Katrina's non-crit Ambush over non-crit Backstab on the same targets, pooled over 10 Wailing Caverns mob types (Ambush 20 hits, Backstab 140): **1.574**. With the flat inside, Backstab rank 3 (32) and Ambush rank 1 (28 inside 250%) predict 1.564 to 1.574 for any attack power from 140 to 200; with the flat outside, 1.847 to 1.871. The ratio does not depend on Opportunity (both take it) or armor. Backstab over Sinister Strike rank 4: 2.378 against 2.384 (inside) or 2.018 (outside); Ambush over Sinister Strike 3.771 against 3.745, which also reads her 12 Subtlety points as Opportunity 5/5 | foreverlogs 2818 | flat inside the 150% (`backstab.go`), as in Ambush and Mutilate | **aligned; closes round 23's open item.** The Forever tooltip's `$s1` prints the flat before the 150% (140 at rank 8, not 210) |
| L2 | Shank: Energy and cooldown | Sugo (2839): all 175 Shanks give 15 Energy, the shortest gap between two is 10.009 s; the three earlier Shank users the same (15, one 11 at the Energy cap; shortest gaps 9.98 to 10.01 s) | rank 1 15 Energy, 10 s, no Flash of Steel at his level | aligned |
| L3 | Slice and Dice duration and Improved Slice and Dice | Sugo: 49 at 9.0 s, 32 at 12.0, 2 at 15.0; Jaice 9.0, 12.1, 15.0, 21.0; Hermes (2850, 11 Nov): 21.7 to 21.8 s five times and 26.1 twice | 6 + 3 s a point; x 1.45 at 3/3 gives 21.75 (3 points) and 26.1 (4) | aligned |
| L4 | Eviscerate's step per point | rank 3 medians step 18 to 24 a point (Jaice, Nakir), rank 5 43 to 52 (Sugo), non-crits after armor | 19 and 45 a point plus 3% of attack power a point | aligned (combo points rebuilt from the resource events; target changes make the one-point bucket noisy, so only the step is read) |
| L5 | Ruthlessness | Hermes: 18 extra points on 37 finishers | 20% a rank, at most 60% | consistent with 2 or 3 ranks |
| L6 | Cheap Shot | Hermes: every Cheap Shot point gain is 2 (12) | 2 | aligned |
| L7 | Blade Flurry's second hit | Sugo (2839): 482 hits under the spell 1230700; a non-crit over the preceding non-crit hit on another target, median 0.81 after an auto attack (303) and 0.73 after a Sinister Strike (90) | the hit's damage before armor, re-armored on the second target | not comparable on these numbers (different mobs, different armor); no card fight has a second target |
| L8 | Poised out of combat | 860 of 863 gains with no damage in the 3 s before | prepull only | aligned |
| L9 | Glancing, crit multipliers, Shank's 25% of the off-hand weapon, Instant Poison | the physical lane's L3 to L7 | | unchanged (no new rogue data beyond theirs) |

## Builds, rotations, orders and rank rows

Nothing the builds were tuned on has moved since round 29: no rogue row in the engine (the engine commits since fe29dfbd2 that touch `sim/rogue`, `sim/core` or `sim/common` are the warlock, shaman, hunter, mage and druid hotfixes, two healer relics and conformance goldens), none in the client, none in Blizzard's notes. Searches on the current engine (`search.out.txt`, 300 iterations to screen, 800 for finalists, seed 7):

| search | verdict |
|---|---|
| Assassination raid 60, raid build | the guide is the best build (869.6 ± 1.6) |
| Combat raid 60, raid build | the guide is the best build (851.3 ± 1.2) |
| Subtlety raid 60 | the guide is the best build (841.1 ± 1.2) |
| Assassination bare 60, leveling build | the guide is the best that keeps its unmodelled trash talents (the +6.8% winner drops Improved Gouge and Remorseless Attacks for Aggression; the guide takes them for trash and leveling, by design) |
| Combat bare 60, leveling build | the same (the +11.4% winner is the raid build's shape without Improved Gouge, Remorseless Attacks and Master of Deception) |
| Subtlety bare 60 | Lethality 1 into a third point of Improved Slice and Dice: +0.5% (2.2 ± 0.9). **Held:** under the adoption rule's 1% bar, and the same build is Subtlety's raid build, where the guide wins |

The cards' standing rulings (Assassination's Eviscerate at four, Combat's Slice and Dice from 60, Subtlety's Death Mark line) rest on numbers nothing here moved. Leveling orders were walked in rounds 27 and 29 on the same engine rows. Rogues carry no rank rows (see the table above).

## Held

- **Subtlety's third Improved Slice and Dice point for Lethality's second** (bare 60 +0.5%, under the 1% bar; raid 60 prefers the guide).

## Open

- Mutilate's off-hand strike (the weapon half halved, the flat whole), Death Mark's storage and cap, Hemorrhage and Ghostly Strike: no public log has a rogue at 40 or more casting them.
- The physical lane's held rows stand: Sinister Strike 6% over the client (Aggression in an unknown build) and Instant Poison's rate (poison uptime unknown).
- Poised under 70888 (MR60).

## Sim changes in this landing

- **Fork:** none. Pin 5f1b39450 unchanged.
- **Site:** this report, the run and analysis scripts under `rogues/`, the ledger row and MR60. No card, guide, coach note, build, order or golden changes, so no wow-player round.

## Sources

- Blizzard, "WoW: Forever Hotfixes - November 11, 2026": https://us.forums.blizzard.com/en/wow/t/2403040 (8 posts, fetched 12 Nov about 07:00 UTC); "Class Tuning Incoming - November 10": https://us.forums.blizzard.com/en/wow/t/2404092.
- wago.tools, build 1.60.1.70888 with `useHotfixes=1`, the tables in Method 2, fetched 12 Nov.
- foreverlogs.gg reports 2818, 2836 to 2853 (`/api/reports/public`, `/api/reports/<id>/meta_events`, `/combatants-info`).
- Our API: `GET https://api.foreversixty.gg/v1/reports/recent`.
