# Sim-accuracy round 17: shaman lane

Specs: Elemental, Enhancement, Restoration. Client 1.60.1.70755 with the late hotfix re-layer. Site branch wt/r17-shaman (off int/r17 d7fc51fe7), fork branch lane/r17-shaman (off int/r17 f637ec32c). Gear is the frozen BiS ($R/bis-frozen); talents are the guide builds.

**2 confirmed bugs, both fixed.**
- Restoration 1: the leveling Healing Wave rows at 38 to 46. Origin R16 54a017916.
- Enhancement 1: the hit weight is 0 past the 9% weapon cap. Origin old. It is a shared core fix, so it moves the hit weights of other specs (Retribution, Marksmanship), not their DPS.
- Elemental: 0.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Leveling Restoration: a falling tank gets Lesser Healing Wave, not a cheap Healing Wave.**<br>The sim has no tank deaths, so it can only value healing per mana. It prefers a rank-4 Healing Wave on a tank under 25% (round 16 R3: bare 40, 170.5 vs 159.3). In a real fight, the 1 s that Lesser Healing Wave saves is what keeps a tank alive.<br>Forever sources agree:<br>- Lesser Healing Wave for a target that "needs a heal ASAP and can't wait for a longer cast".<br>- Lesser Healing Wave "when the longer cast would land too late".<br>- Healing Wave is the main heal, with Lesser Healing Wave as the fast heal for party members in danger.<br>Client rows: Lesser Healing Wave is a 1.5 s cast at 0.429. Healing Wave ranks 4 and up are 3.0 s at 0.857, and 2.5 s with Improved Healing Wave.<br>**Decision: keep the Lesser Healing Wave line (tank under 30%).** Not adopted, because the sim cannot weigh the risk of a dead tank. | [Warcraft Tavern, Forever Restoration Shaman guide](https://www.warcrafttavern.com/forever/guides/restoration-shaman/) (fetched 2026-11-04)<br>[Mobalytics, Forever Restoration Shaman guide](https://mobalytics.gg/wow-forever/classes/restoration-shaman-guide) (updated 2026-11-04)<br>[Icy Veins, Forever Restoration Shaman guide](https://www.icy-veins.com/wow-forever/restoration-shaman-healer-pve-guide) (updated 2026-11-04)<br>[wowtbc.gg, Forever Restoration Shaman](https://wowtbc.gg/warcraftforever/class-guides/restoration-shaman/)<br>spellconst/shaman.json 8004 to 10468 and 913 to 25357 | Aligned (decided by research) | - | - |
| S2a | **Regression: Healing Wave per-level rows.** See R1. | | CONFIRMED | R16 54a017916 | S d153f2bdf |
| S2b | **Regression: Lightning Bolt per-level rows.** Every level from 30 to 59 was re-probed with the row's rank and both neighbours. Each row rank wins or ties (for example 39: rank 5 64.1 vs rank 4 63.7; 45: rank 6 107.9 vs 107.2; 57: rank 8 156.5 vs rank 7 153.6). At bare 60, rank 8 236.1 beats rank 9 226.1. | $R/shaman/ele-probe-3k.txt (horde, 3000 iterations) | Aligned | - | - |
| S2c | **Regression: the low-rank penalty key moved to MaxLevel + 12** (priest E 18815822c, merged after the shaman lane). It does not move a shaman row. Every Healing Wave and Lightning Bolt rank a row uses has MaxLevel = learn level + 5, so its grace still ends 17 levels after it is learned. The only shaman ranks in clientSpellMaxLevel are:<br>- Healing Wave 1 (unused)<br>- Riptide 408521/1239242/1239243<br>- Flame Shock 6<br>None of them reaches the penalty inside its band. | sim/core/spell_max_level_auto_gen.go; spell.go:824 | Aligned | - | - |
| S2d | **Regression: the Nature's Swiftness save gate.** Line 5 still needs `auraIsKnown 16188`. The save casts rank 7 at 44 to 47, 8 at 48 to 55 and 9 from 56. | TestRestorationShamanSwiftnessSaveNeedsTheTalent and TestRestorationShamanSwiftnessHealingWaveIsTheTopRank pass | Aligned | - | - |
| S2e | **Regression: the Rockbiter swap under Windfury Totem.** It still fires (TestLayerSwapsWindfuryWeaponForRockbiterUnderTheTotem). Raid 60 Enhancement is 855.3 / 855.6, the same as before. | $R/shaman/after-raid60.tsv vs $R/before.tsv | Aligned | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **The leveling tank Healing Wave cast the wrong rank at 38 to 46.**<br>What was wrong:<br>- Rank 5 at 38 to 41 lost to rank 4 by 0.5 to 3.3%.<br>- Rank 6 at 44 to 46 lost to rank 5 by 1.0 to 3.9%.<br>Round 16 picked these rows at 04:34, before its own 04:52 commit that gated the Nature's Swiftness save and scoped the rows to the filler line. Its numbers do not reproduce, either on the round-16 lane binary in its own worktree or now (38: rank 5 104.0, not 107.3).<br>Why the cheaper rank wins: every leveling fight runs dry. Rank 4 is learned at 18 and keeps its full bonus until 35, then 5% less a level. It still heals more before the mana runs out than the 200-mana rank 5.<br>New rows (line 17):<br>- rank 2 to 25<br>- rank 3 to 36<br>- **rank 4 to 41**<br>- **rank 5 from 42 to 46**<br>- rank 6 from 47 to 56<br>- rank 7 at 57 and 58<br>- rank 8 at 59<br>The guide's leveling paragraph and card note 17 follow. | Effective HPS, 10,000 iterations, seeds 1 / 50001 ($R/shaman/resto-10k.txt):<br>- 38 horde: rank 4 107.4 vs rank 5 104.0 (+3.3%)<br>- 38 alliance: 107.7 vs 104.6 (+3.0%)<br>- 41: 161.5 vs 160.7 / 160.8 (+0.5%)<br>- 42: rank 5 160.0 vs rank 4 159.8 / 159.7 (+0.1%, rank 5 kept)<br>- 44 horde: rank 5 166.4 vs rank 6 160.2 (+3.9%)<br>- 44 alliance: 167.1 vs 160.9 (+3.9%)<br>- 46: 163.9 vs 162.3 / 162.4 (+1.0%)<br>- 47: rank 6 163.6 / 163.5 vs 162.7 / 162.6<br>The 3000-iteration sweep of levels 20 to 59, each row rank against its neighbours, is in $R/shaman/resto-probe-3k.txt. Every other row holds. | CONFIRMED | R16 54a017916 | S d153f2bdf |
| R2 | Mana Tide Totem.<br>- Ranks 16190, 17354, 17359 and 1325256 are learned at 1 (talent), 38, 48 and 58.<br>- The ticks restore 88 / 133 / 197 / 290 every 3 s for 13 s, which is 4 ticks.<br>- The cooldown is 5 min.<br>The engine matches. | spellconst 16191/17355/17360/1325257; mana_tide_totem.go | Aligned | - | - |
| R3 | Restoration talents against the 70755 text:<br>- Improved Healing Wave -0.5 s<br>- Tidal Focus 5% cost and 5% hit<br>- Tidal Mastery 5%<br>- Purification 10%<br>- Healing Way 8/17/25%<br>- Mindfulness 17/33/50%<br>- Restorative Totems 25% / 50%<br>All match.<br>Healing Focus 3 (pushback) does nothing in the heal profile, but it is real value for a healer who gets hit. It is not a dead point. | talents/shaman.json; sim/shaman/talents_restoration.go | Aligned | - | - |
| R4 | Water Shield: 3 globes, 2% of maximum mana, 10 min, on a hit taken or a crit. The engine matches. | spelltips; water_shield.go | Aligned | - | - |
| R5 | Limitation, not a bug: `-level` keeps the band's floor talents. At 44 to 49 the band-40 string has no Nature's Swiftness, so the save does not fire in those probes. The 50+ probes include it. | sim/cmd/spec-breakdown | note | - | - |

## Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | **Hit was published as worth 0 past the 9% weapon cap, though the shaman's spells still miss until 16%.**<br>What was wrong: Enhancement's raid band has 11% hit. Its Earth Shocks, Maelstrom Lightning Bolts, Searing Totem and Fire Nova (about 22% of its damage) miss 7.2 to 7.7% of the time. Hit is one stat for both tables (SpellHitChance reads it). But the stat-weight sweep (core hitStepWindow) stopped measuring hit at the weapon cap, so the weight was a flat 0 ("insignificant", error 0) and the ranker gave hit gear nothing past 9%.<br>Fix: for a physical character who casts damage spells, the window now runs up to the spell cap.<br>Same zero weight elsewhere in the frozen BiS:<br>- Retribution raid 60 (12.9% hit)<br>- Marksmanship alliance raid 60 (9.3%)<br>- Arms raid 60 (10.8%). A warrior casts no spell, so its sweep reads identical results and stays 0.<br>Published DPS does not move; weights do. Guide accuracy round 13 raised this ("Sim findings"); no sim round took it up. | $R/shaman/sb breakdown, raid 60 horde (spell miss 7.2 to 7.7%, hit 11%)<br>bis-frozen hit_to_cap {baseline 11, specials 0}<br>TestHitStepWindowKeepsMeasuringPastTheWeaponCapWhileSpellsStillMiss | CONFIRMED | old | E 7c8f710d0 |
| N2 | Earth Shock rank while leveling (the line casts the highest rank learned):<br>- Bare 50: rank 6 217.0 vs rank 5 214.7.<br>- Bare 60: rank 7 312.9 vs 307.6 vs 301.2.<br>- 40: rank 5 154.8 vs rank 4 153.5.<br>- 47: rank 4 164.4 vs 164.1 (a tie).<br>The top rank holds. | 3000 iterations, horde | Aligned | - | - |
| N3 | Raid build utility points (Improved Ghost Wolf 2, Mindfulness 3, Ancestral Knowledge at 3). These are deliberate picks the guide explains (enhancement.md:52). None is a point that does nothing for the spec in play. | | Aligned | - | - |
| N4 | The relic slot is empty in both factions' published raid 60 entries. Totem of Thunder (+1% Lightning Bolt crit) is modelled, but on Enhancement's roughly 3% Lightning Bolt share its gain is under the ranker's noise gate (relics.go), so the slot publishes empty. A player would still wear it. This is ranker policy, not an engine bug. | sim/cmd/leveling-bis/relics.go | note (Measure at refresh) | - | - |
| N5 | Strength of Earth rank 4 and Grace of Air rank 2: rank 5 (25361) and rank 3 (25359) have trainer rows at 60. These are the level-60 book ranks, which are pending with the owner. Not reopened. | trainables/shaman.json | pending owner | - | - |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | **Guide-accuracy "Sims requested 1"** (carried through rounds 6 to 13): Earth Shock rank 7 as a filler after Flame Shock. It loses:<br>- Gated on Flame Shock above 6 s and mana ahead of the fight: 666.1 to 656.1 (-1.5%; out of mana 1.7 to 6.4 s).<br>- Ungated: -5.4%.<br>The request is closed: no Earth Shock line. | Raid 60 horde, 10,000 iterations, seeds 1 / 50001 (666.6 to 655.9) | Aligned | - | - |
| E2 | Tier 1 Storm 5-piece ("Lava Burst cooldown") has no spell to reach on 70755. The engine marks it NoSim. | item_sets_tier1.go | Aligned | - | - |
| E3 | Raid build points: Elemental Warding 3 and Elemental Devastation 1 are the toll to Elemental Fury. Every other point in tiers 1 to 3 (Reverberation, Trial by Fire, Eye of the Storm, Earthbound) does nothing in a single-target fight either, so no legal swap takes a point that does something. Pending owner (O2) as before. | talents/shaman.json | Aligned | - | - |
| E4 | Effects on the published BiS are registered:<br>- Totem of Thunder (228176)<br>- Frozen Heart of the Mountain (249469)<br>- Thunderbrew's Boot Flask (744)<br>- Hand of Edward the Odd (2243)<br>- Chthonic Power<br>Rune of the Guard Captain is plain stats. | sim/common/item_effects.go, client_procs.go, sim/shaman/items.go | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **N1 is in shared core (sim/core/hitprofile.go, 7 lines).** It changes only the stat-weight sweep's hit window for physical characters with damage spells. It moves no DPS. It moves the hit weight for Enhancement and Retribution, and Marksmanship if Arcane Shot rolls the spell table. **Flagged for the paladin and hunter lanes.** | E 7c8f710d0 | tool | - | E 7c8f710d0 |
| X2 | The rotation ladder's "Learned but unused" list shows Lava Burst at 30 to 60 (spellranks.json still carries it). 70755 removed every Lava Burst trainer row. This is informational only and no rotation casts it. For the platform/data lane. | sim/request/testdata/ladder/shaman-restoration.golden.md; data/builds/1.60.1.70755/spellranks.json | note | - | - |

## Before/after (10,000 iterations, frozen BiS, guide builds)

| Spec, band | Faction | Seed 1 before to after | % seed 1 / 50001 |
|---|---|---|---|
| Restoration bare 40 | horde | 159.4 to 162.9 | +2.20 / +2.20 |
| Restoration bare 40 | alliance | 160.3 to 163.7 | +2.12 / +2.12 |
| Elemental raid 60 | horde / alliance | 666.1 / 669.4, unchanged | 0.00 |
| Enhancement raid 60 | both | 855.3, unchanged | 0.00 |
| Restoration raid 60 | horde / alliance | 757.1 / 759.4, unchanged | 0.00 |

Inside the band, new row against old row:

| Level | Horde | Alliance |
|---|---|---|
| 38 | +3.3% | +3.0% |
| 44 | +3.9% | +3.9% |

Bare 20, 30, 50 and 60 are untouched (their rows did not change).

## Files to regenerate

- sim/request/testdata/ladder/shaman-restoration.golden.md (the 38 and 40 rungs now cast Healing Wave rank 4).
- addon-data.json, addon-builds.json, Data.lua: Restoration rank rows and card note 17.
- Coach starters, if they carry shaman ranks.
- dps-per-point and BiS weights, for the hit weights of Enhancement, Retribution and Marksmanship (refresh).
- Hand-synced, not regenerated:
  - sim/rankchoices/rank-choices.json (`make rank-choices-sync`)
  - sim/request/apl/shaman-restoration.apl.json (note 17 only)
  - The fork's restoration APL copy carries no notes and no action changed, so it needs no sync.

## Measure at refresh

- Enhancement, Retribution and Marksmanship raid 60 hit weights and hit-gear picks with the spell-cap window (N1).
- Restoration bare 40 picks on the new rows.
- Enhancement's empty relic slot (N4): Totem of Thunder or Totem of the Storm.
- Carried: Enhancement band 30 with Stormstrike; Enhancement raid weights with Rockbiter.

## For other lanes

- **Paladin:** Retribution's raid-60 hit weight was a flat 0 at 12.9% hit. With E 7c8f710d0 the sweep measures Exorcism, Judgement and the other spell-table damage up to the 16% spell cap. Weights only.
- **Hunter:** Marksmanship alliance raid 60 (9.3% hit) also published 0. Whether it moves depends on Arcane Shot's table.
- **Platform/data:** spellranks.json keeps Lava Burst rows that 70755 removed (X2).
- **Ranker:** a modelled relic whose gain is under the noise gate publishes an empty slot (N4). Consider "equip the best modelled relic when any gain is positive".

## For the guide lane

Shipped with this lane. wow-player gave SHIP on the rendered page (talents tree hydrated, leveling paragraph):
- restoration.md:120: "rank 4 until 41, rank 5 until 46".
- Card note 17 (data/curated/apl/shaman-restoration.json).

Optional SHOULDs from the reviewer, not done:
- "rank 5 from 42 to 46" is less ambiguous than "until 46".
- Give a one-clause reason why a just-trained rank is skipped.
- The note's "ranks 1 to 4 keep none of it at 60" and the prose's "learned at 23 or earlier" mean the same thing; phrase them alike.

## Still open

- Pending owner, not reopened:
  - the Healing Wave 2 tooltip on a geared 60
  - Lightning Overload rank spells
  - level-60 book ranks (Strength of Earth 5, Grace of Air 3)
  - Elemental Devastation 1 as the Elemental toll
- Warcraft Logs is blocked (gcloud login expired). Not worked around.
