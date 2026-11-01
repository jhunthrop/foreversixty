# Sim-accuracy round 6: ranker lane

Lane `ranker`, site branch `wt/r6-ranker` (from origin/main). Owns `sim/cmd/leveling-bis` this round. No engine change. Engine for every run: fork `lane/r6-warlock` at ee0393589 (adds `SimItem.damage_school` (24) and casts Shoot in the wand's school); "before" binaries were built at 8baffe5fc.

Re-rank method: the ranker at origin/main ("old") against this branch ("new"), on the same engine and the same local item database. That database is the committed `simdb.bin` with two local patches. The 52 client wands get their `damage_school` through this branch's `_apply_wand_school`. The three items in S2 get the ranged-AP fix. Neither patch is committed. `-bands 40,60 -weights-iterations 200`, both factions. Then each published set is measured with spec-breakdown at 10,000 iterations, seeds 1 and 50001, built on the same database. No BiS json, simdb or items json is committed.

## Item 1: wand credit (warlock T3)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| W1 | **score() gave a wand its whole tooltip DPS whenever the APL had a Shoot line.** Code: `score.go:184`, `c.DPS / referenceDPSPerPoint` when `castsShoot`. On the band-60 raid sets a warlock Shoots 0 times a fight, and so do Fire and Frost mages and Shadow priests. Even so, Torch of Light (76.8 DPS) beat Brilliant Wand (59.6) by 17 DPS of credit, and it held every warlock, mage and Shadow priest ranged slot at band 60. Fix: credit is now measured per band. The first pass picks at the whole-fight credit, so the set wears a wand. A sim of that set (verify harness, 300 iterations, seed 7) reads Shoot's damage per second from the player's action metrics (`wand_credit.go` `shootDPS`). The band is then picked again at `shootDPS / wandDPS / referenceDPSPerPoint`. That ratio is the Shoot time share times hit, crit, partial resists and Wand Specialization. Measured credit, as a share of the whole-fight credit: 0 on every band-60 raid set except Arcane's (0.6 to 0.8%); 0 on every warlock band 40; Arcane bare 60 22 to 24%; Frost bare 60 39 to 46%; Shadow bare 40 33 to 35%. Tests (written first; failed to compile: `measureWandCredit`, `shootDPS`, `wandScorePerDPS` undefined): `TestWandScorePerDPSIsTheMeasuredShootDamageOverTheWornWand`, `TestScoreCreditsAWandOnlyAtTheBandsWandCredit`, `TestShootDPSIsShootsShareOfThePlayersOwnDPS`, `TestMeasureWandCreditSimsTheBandsSetAndReadsItsShootDamage`, `TestMeasureWandCreditIsZeroWithoutAWandAndRunsNothing`. | ranker logs `wand credit` lines (scratch `r6/ranker/out-new/*.log`); `sim/cmd/leveling-bis/score.go`, `wand_credit.go`, `band.go` `pickBand` | CONFIRMED-BUG | S 413808063 |
| W2 | **The verify pass needed a flat 1% gain.** `verifyBand` promoted a runner-up only when `beatsByMargin` (1% at 300 iterations) and the gain beat the error. Alliance Affliction's Brilliant Wand measured ahead of Torch beyond the error (+1.9 ± 0.4 DPS at 10,000 iterations, +0.33%), but under 1%. So Torch stood, and `swapMeasuredDelta` clamped the positive gap to "dps_delta 0, verified". Horde did not clear the 300-iteration error, so it published the score's -15.0 instead. A single slot's gap is almost always under 1% of a set. Fix: a runner-up that runs ahead at 300 iterations without clearing the margin is re-measured at the guard harness (`guardIterations`, 2,100). The baseline is measured the same way once per band. The runner-up is promoted when it is ahead beyond the combined error there. A runner-up that loses costs no run. Verify time over the 7 specs, bands 40 and 60: 266 s to 318 s (+19%). Tests: `TestVerifyBandPromotesARunnerUpAheadInsideTheMarginOnceTheGuardHarnessConfirmsIt` (failed: Beat false at 571.9 against 570), `TestVerifyBandKeepsThePickWhenTheGuardHarnessCallsItATie`, `TestVerifyBandDoesNotReMeasureARunnerUpThatLoses`. | `verify.go` `verifyBand`/`confirmSwap`; `report.go:1190` `swapMeasuredDelta`; published `bis/warlock-affliction.json` band 60 raid | CONFIRMED-BUG | S 2c379af92 |
| W3 | **Found on the re-rank: a trinket slot published empty when a trinket gains.** `rankByMeasuredGain` orders the shortlist at 100 iterations (several DPS of error on a raid set) and refines only the winner's gain. A winner that measured at or below the empty slot was never refined (`trinketAdaptiveGain` needs gain > 0), and the low-gain gate then emptied the slot. With W1 and W2 alone, warlock-demonology horde band 60 raid published trinket2 empty, at "0.00 ± 5.71". There Rune of the Guard Captain sims +0.7% at 10,000 iterations, and mage-fire alliance band 60 raid lost about 1% the same way. The published `druid-balance` band 60 raid alliance trinket2 is empty for the same reason. Fix: when the winner's gain is not significant, the empty slot and the three best-measured candidates are re-measured at the 2,000-iteration ceiling, and the largest significant gain takes the slot. Tests: `TestRankByMeasuredGainFallsBackToARealGainWhenTheNoisyWinnerHasNone` (failed: Noisy kept at -0.5 ± 4.2), `TestRankByMeasuredGainKeepsTheWinnerWhenNothingGainsAtTheCeiling`. | `trinkets.go` `gainAtCeiling`; `out-new/warlock-demonology.json` trinket2 `empty_reason: no_dps_value`; spec-breakdown `-gear trinket2:19120` 653.1 against 648.7 | CONFIRMED-BUG | S 698d66ce3 |

### Ranged slot, band 60 (old ranker / new ranker)

| spec | raid alliance | raid horde | bare alliance | bare horde |
|---|---|---|---|---|
| Affliction | Torch / Brilliant Wand | Torch / Brilliant Wand | Torch / Brilliant Wand | Torch / Brilliant Wand |
| Demonology | Torch / Brilliant Wand | Torch / Brilliant Wand | Torch / Brilliant Wand | Torch / Greater Eternal Wand |
| Destruction | Torch / Brilliant Wand | Torch / Brilliant Wand | Torch / Brilliant Wand | Torch / Brilliant Wand |
| Arcane | Torch / Brilliant Wand | Torch / Brilliant Wand | Torch / Torch | Torch / Torch |
| Fire | Torch / Pyric Caduceus | Torch / Pyric Caduceus | Torch / Torch | Torch / Torch |
| Frost | Torch / Brilliant Wand | Torch / Brilliant Wand | Torch / Torch | Torch / Torch |
| Shadow | Torch / Brilliant Wand | Torch / Brilliant Wand | Torch / Brilliant Wand | Torch / Brilliant Wand |

Band 40 bare: warlocks move from Jaina's Firestarter to Twisted Nether Wand, and Arcane to Nether Force Wand. Fire, Frost and Shadow keep their wands. Where a mage keeps Torch bare, it is earned: it shoots 22 to 46% of the fight, and the verify run backs Torch (Brilliant Wand -3.1 to -5.2 DPS). The new sets also change other slots, through the guard-harness promotions (W2) and the trinket fallback (W3). Per-band slot diffs are in scratch `r6/ranker/diffs.txt`.

## Item 2: wand school in the data pipeline (warlock T4)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | `pipeline/simdb/items.py` writes `SimItem.damage_school` for wands only (Item class 2, subclass 19). ItemSparse `DamageType` 1/2/3/4/5/6 maps to SpellSchool Holy/Fire/Nature/Frost/Shadow/Arcane (proto 4/2/5/3/6/1). 0 is physical and is never written. An unknown value raises. Classic-db rows read `item_template.dmg_type1` into a new `ClassicDbItem.damage_type` (default 0, so an older extract still loads). data.yml re-extracts the dump on every run (`fetch-classic-sources`). 52 client wands get a school on 70707. Classic-db-only wands get theirs at the next extract. Tests: `test_a_wand_carries_the_school_its_item_sparse_damage_type_states` (7 cases), `test_only_a_wand_carries_a_damage_school`, `test_a_classicdb_wand_carries_its_dmg_type1_school`, `test_an_extract_written_before_dmg_type1_was_read_is_physical`, `test_extract_records_reads_a_wands_dmg_type1`. All failed first. | ItemSparse 70707 DamageType; Wowhead Forever 279246 "Arcane Damage" | done (feature for T4) | S e395edd50 |
| D2 | **The vendored proto moves only at the pin.** `python -m pipeline simproto --engine <checkout>` rewrites `data/proto`, `pipeline/simproto/common_pb2.py` and `data/proto/ENGINE_SHA`. `test_genproto` requires ENGINE_SHA to equal `sim/enginever`, so the regeneration was not committed. On today's bindings, 9 of the new tests fail (no field 24), and so would any data.yml simdb run that reaches a wand with a school: `item.damage_school` raises AttributeError. All 108 pass on bindings regenerated from ee0393589. **Merge e395edd50 only together with the pin that carries ee0393589, and re-run simproto in the same change.** | local run with and without the regenerated `common_pb2.py` | integrator step | - |

## Item 3: excluded caster items (warlock C2)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | Briarwood Reed (12930) is excluded as `client_unconfirmed`, and that is right. The 70707 Item table has the id, but ItemSparse has no row for it. Neither does the 70707 hotfix push, nor `raw/wowhead/items.json`, and the Wowhead Forever tooltip returns 404. Its +29 comes from classic-db alone. It is also a rare-spawn drop (Jed Runewatcher, UBRS, rank 2, classic-db chance 0.0). Burst of Knowledge (11832) and Bonecreeper Stylus (13938) are in the same position: in the Item table, missing from ItemSparse, 404 on Wowhead Forever. No source confirms Forever's stats for any of them, so the L1.6 rule holds. | wago ItemSparse/Item 70707; https://nether.wowhead.com/forever/tooltip/item/12930 and /11832 (404, 2026-11-01); https://nether.wowhead.com/classic/tooltip/item/12930 (+29); `raw/classicdb/sources.json` | Aligned (exclusion right) | - |
| X2 | Talisman of Ascendance (22678) is confirmed: ItemSparse row, Forever tooltip. Its only source is the Argent Dawn "Epic Armaments of Battle" turn-ins (quests 9222/9224/9225/9228; Insignia 22523/22524). `content-phases.json` gates those as `argent-dawn-armaments`, phase "later", from original patch 1.11, the Naxxramas lead-in. The 2026-10-30 phase audit lists it as not available. Forever launches with Phase 1 content, so the gate is right. | `data/curated/content-phases.json:230-248`; `raw/classicdb/sources.json` 22678; `design/reviews/2026-10-30-bis-phase-audit.md:143` | Aligned (exclusion right) | - |
| X3 | Integrator item (a): Benediction (18608) has no source in `loot.json`. It is the reward of the Eye of Divinity chain. The Eye (18646) drops only in Molten Core: Majordomo Executus, Cache of the Firelord, and Flamewaker Healer and Elite (classic-db map 409; `loot.json` `raid:molten-core`, opens "later"). Molten Core is not in Forever's announced content, the same finding as round 5's Rhok'delar (RK1). So it stays out of the priest candidates. The simdb row carries the server's stats (195 healing, 65 spell damage, matching the Forever tooltip). `items/priest.json` shows Intellect, Stamina, Spirit and Shadow resistance but not the 195/65 equip line, a planner display gap (see Still open). | `loot.json`; classic-db sources 18646; https://nether.wowhead.com/forever/tooltip/item/18608 | Aligned (exclusion right) | - |

## Integrator item (b): doubled ranged attack power

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **A generic "+N Attack Power" spell counted its ranged half twice.** The client states it as one spell with aura 99 N and aura 124 N: Era 14027 "Attack Power 24"; Forever 1318000/1287704 "Authority" on Rune of the Guard Captain (14/14 and 28/28), 1318001/1317999 on Cleansed Vilebranch Medallion, and 1319047 on Graverobber's Shovel (12/12). `equip.spell_bonus` kept both rows, and `statmap.stat_array` mirrors `attack_power` into ranged attack power as well. Result: simdb 19120 and 270275 carried 42 AP and 84 RAP, 273817 12 and 24, and the Zandalar Signet of Might enchant (2606) 30 and 60. The `items/<class>.json` rows also carried an extra `ranged_attack_power` line, so the ranker weighed the RAP twice for hunters. Two sources agree on single counting. The Wowhead Forever tooltip for 19120 shows one line, "+14 Attack Power. This effect is tripled in Forest and Grassland areas" (no ranged line). The fork's db.json states the rune at Era 20/20, ranged equal to melee. Classic-db's `equip_stats` already drops the 124 row. Fix: within one spell, the 124 amount that the 99 amount covers is the mirror, and only an excess stays ranged-only. The parity test and the archived 70374 build test still pin that build's committed bytes; their notes now say why. Tests: `test_a_generic_attack_power_spell_grants_ranged_attack_power_once`, `test_ranged_attack_power_beyond_a_spells_melee_amount_is_ranged_only`, and the 14027 fixture expectation. All failed first. | wago SpellEffect/ItemEffect/ItemXItemEffect 70707; https://nether.wowhead.com/forever/tooltip/item/19120 (2026-11-01); `pipeline/simdb/equip.py`; `tests/test_simdb_fork_parity.py` | CONFIRMED-BUG | S b9fc1cbf5 |
| R2 | Rune of the Guard Captain sums both of its spells, 14 + 28 = 42 AP. The 28 is the "tripled in Forest and Grassland" part, so the item is credited with its zone bonus everywhere. Whether a fight's zone counts is not settled. Barrow Deeps, Hyjal Summit and Onyxia's Lair, and the type-neutral target, have no zone type in the sim. | tooltip text above; spells 1318000 / 1287704 | OPEN (needs the spell's area condition and a ruling) | - |

Horde hunters wear the rune in every band-50 and band-60 set, so R1 moves them, outside this group (raid / bare, seed 1 / 50001):

| spec (horde, band 60) | raid | bare |
|---|---|---|
| Marksmanship | 679.7 to 671.5 (-1.21%) / 680.3 to 672.1 (-1.21%) | 312.5 to 308.1 (-1.41%) / 312.2 to 307.8 (-1.41%) |
| Beast Mastery | 674.2 to 666.8 (-1.10%) / 674.9 to 667.6 (-1.08%) | 307.6 to 303.8 (-1.24%) / 307.8 to 304.0 (-1.23%) |

These are the published sets on the patched simdb. The nightly re-rank decides whether the rune still holds trinket1. Survival fights in melee and was not re-measured. Melee classes are unaffected: their weights give ranged attack power 0. Zandalar Signet of Might is phase 4 and is in no published preset.

## Before/after (published set by the old ranker against the new ranker, same engine and database; 10,000 iterations)

Demonology and Fire include W3 (S 698d66ce3). The others were ranked before W3 landed, so their rows show W1 and W2 only.

| spec | faction | band 60 raid, seed 1 (abs) / 50001 | band 60 bare | band 40 bare |
|---|---|---|---|---|
| Affliction | alliance | 572.4 to 577.7 (+0.93%) / +1.01% | +1.05% / +0.99% | +1.43% / +1.43% |
| Affliction | horde | 573.5 to 575.4 (+0.32%) / +0.31% | +0.22% / +0.21% | +1.05% / +1.06% |
| Demonology | alliance | 644.3 to 649.6 (+0.83%) / +0.82% | +0.93% / +0.84% | +1.18% / +1.14% |
| Demonology | horde | 646.8 to 648.6 (+0.27%) / +0.15% | +0.09% / +0.12% | +0.98% / +0.91% |
| Destruction | alliance | 670.2 to 678.6 (+1.25%) / +1.26% | +0.34% / +0.32% | +0.87% / +0.93% |
| Destruction | horde | 673.4 to 679.1 (+0.84%) / +0.70% | +0.36% / +0.35% | +0.94% / +0.75% |
| Arcane | alliance | 585.4 to 595.1 (+1.66%) / +1.61% | 0.00% / 0.00% | +1.30% / +1.29% |
| Arcane | horde | 588.4 to 591.4 (+0.50%) / +0.55% | 0.00% / 0.00% | +0.76% / +0.71% |
| Fire | alliance | 693.8 to 708.3 (+2.09%) / +1.99% | 0.00% / 0.00% | +1.92% / +2.11% |
| Fire | horde | 699.7 to 706.4 (+0.96%) / +0.93% | 0.00% / 0.00% | 0.00% / 0.00% |
| Frost | alliance | 542.0 to 546.8 (+0.88%) / +0.90% | 0.00% / 0.00% | +0.55% / +0.44% |
| Frost | horde | 546.9 to 548.9 (+0.37%) / +0.38% | 0.00% / 0.00% | +0.40% / +0.55% |
| Shadow | alliance | 707.5 to 712.7 (+0.74%) / +0.71% | +1.55% / +1.54% | +1.63% / +1.57% |
| Shadow | horde | 713.0 to 715.1 (+0.29%) / +0.31% | +1.65% / +1.65% | 0.00% / 0.00% |

No cell loses. A 0.00% cell is a band that re-picked the same set, or where the guard kept the published set (Fire bare 60 both factions: "kept the previous published set").

## Still open

- **Pin step for D1/D2:** pin the engine with ee0393589 and run `python -m pipeline simproto --engine <checkout>` in the same change as e395edd50. Before that, a data.yml simdb run raises on the first wand.
- R2: Rune of the Guard Captain's zone-tripled 28 AP counts in every fight.
- Generic-AP mirror on 99-only spells: 24 item spells on 70707 state aura 99 with no 124 row (e.g. 26153 on Mace of Unending Life, worded as feral forms only), and `stat_array` still mirrors them into ranged attack power. They are not generic AP lines, so a hunter wearing one gets RAP it may not have. Not checked one by one.
- `items/priest.json` Benediction (18608) lacks its "+195 healing, +65 damage" equip line (simdb has it). Display only today, because the item is unobtainable.
- The re-rank was local. The nightly bis.yml decides the published sets. Expect every caster ranged slot to change (Brilliant Wand, Pyric Caduceus for Fire raid, other wands below 60), some sub-1% slot promotions from W2, and filled trinket slots from W3, including druid-balance band 60 raid alliance.
- sim/request TestRotationLadder goldens were not touched.

Sources (read 2026-11-01): wago.tools db2 ItemSparse, Item, ItemEffect, ItemXItemEffect, SpellEffect, SpellName at build 1.60.1.70707 (warlock lane's copies, `r6/warlock/db2/`); https://nether.wowhead.com/forever/tooltip/item/19120, /12930, /11832, /22678, /18608; https://nether.wowhead.com/classic/tooltip/item/12930.
