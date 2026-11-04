# Sim-accuracy round 19: shaman lane

Specs: Elemental, Enhancement, Restoration. Client 1.60.1.70755 with the late hotfix re-layer. Site branch wt/r19-shaman (off int/r19 f91e9d3f9), fork branch lane/r19-shaman (off int/r19 6ea675594). Gear is the frozen BiS ($R/bis-frozen); raid 60 runs the guide raid build ($R/rt), bare bands the published band talents. Every rank was judged level by level with `spec-breakdown -level`, forcing a rank by authoring its own id on the line it governs (every reference on that line rewritten, harness `$R/shaman/sub.py`, `judge.sh`).

**4 confirmed bugs, all fixed** (Restoration 2, Elemental 1, Enhancement 1), plus one cosmetic fix (ladder golden header).
- Restoration: Chain Heal cast its top rank while leveling; tank Healing Wave rank 7 one level earlier, at 56, once Chain Heal is cheaper.
- Elemental: filler Lightning Bolt at 26 and 27 (guide round 15 "Sims requested 1").
- Enhancement: Earth Shock reserve at 40 to 49 (round 18 N1's open band), now expressible through a Maelstrom Weapon rank tag.
- Origins: all four `old`.

Healing results are near-deterministic in this profile: seeds 1 and 50001 and 3,000 against 10,000 iterations agree to 0.1 HPS, so gaps of 0.3% are real.

Client note: wago lists wowf 1.60.1.70790 (18:08 UTC). For shaman it carries exactly the hotfix values our 70755 spellconst already holds (43 SpellEffect rows, Riptide and Chain Lightning costs, Riptide 12 s, Stormstrike 6 s, Totemic Call 30 s: 0 mismatches; $R/shaman/w790). The integrator has it.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Restoration's Chain Heal cast the highest rank learned while leveling.**<br>What was wrong: the Chain Heal line (line 16) authors rank 3, so from 46 it cast rank 2 (315 mana) and from 54 rank 3 (405). A leveling Restoration shaman runs out of mana in every fight (mana lasts about 110 s of 300 at 43), so the cheaper rank 1 (260 mana, keeps its whole bonus until 57 on the MaxLevel + 12 rule, 90% at 59) heals more over the fight.<br>Level by level, rank 1 over the default, horde / alliance, seed 1 (seed 50001 within 0.1): 46 +0.3% / +0.3%; 49 +0.3% / +0.4%; 50 +0.4% / +0.3%; 53 +0.7% / +0.6%; 54 +1.0% / +1.0% (rank 2 +0.4%); 56 +1.0% / +1.1%; 59 +1.2% / +1.2%.<br>Bare 60: rank 2 +1.05% / +0.9%, rank 1 +0.9% / +0.7%. Raid 60: rank 3 wins (rank 2 -3.5%, rank 1 -6.3%), as round 15 R4 found.<br>Fix: rows rank 1 at 46 to 59 and rank 2 at bare 60 on line 16. Card note 16, the leveling paragraph and the Chain Heal ranks paragraph follow. | `ch-10k.txt`, `resto-sweep-h1.txt` (lane dir); `TestRestorationShamanChainHealRankOneWhileLeveling` failed, then passed | CONFIRMED | old (round 15 R4 judged Chain Heal at raid 60 only; rounds 16 to 18 judged the Healing Wave rows) | S 472dcaaf9 |
| S2 | **Tank Healing Wave rows, re-judged at 40, 43, 46, 49, 50, 53, 56, 59 (and 54 to 58 one by one).** Every round-17 row holds on its own. With S1 in, rank 7 now beats rank 6 at 56 (+0.4% / +0.7%, both seeds); rank 6 still wins at 54 (+2.1%) and 55 (+0.9%); 57 and 58 stay rank 7; 59 rank 8 beats ranks 7 and 9.<br>Fix: rank 6 at 47 to 55, rank 7 at 56 to 58. | `hw-joint-h1.txt`, `hw-fine-*.txt`, `round2.txt`; `TestRoundNineteenShamanRowsByLevel` failed, then passed | CONFIRMED (follows S1) | old | S adc341da2 |
| S3 | Lesser Healing Wave (lines 7 to 11) at 40 to 59: the highest rank learned beats every cheaper rank (rank 3 at 46 -3.6%, at 53 -7.3%; rank 4 at 56 -5.6%, at 59 -1.7%; rank 2 at 43 ties). Riptide rank 1 at 50 to 59 loses 0.7% to 1.3%. No rows. | `resto-sweep-h1.txt` | Aligned | - | - |
| S4 | **Elemental Lightning Bolt rows, re-judged at 40, 43, 44, 46, 49, 50, 53, 56, 59.** Every row beats its neighbours (rank +1 and -1). The top-rank line 7 and the published rows hold; Flame Shock rank 4 at 53 to 59 and Searing Totem rank 4 at 50 to 56 lose 0.9% to 1.7%. | `ele-sweep-h1.txt` | Aligned | - | - |
| S5 | **Elemental Lightning Bolt at 20 to 29 (guide round 15 Sims requested 1).** At 26 and 27 the just-learned rank 5 costs more than the fight pays: rank 3 beats it by +3.7% / +3.9% (horde, 26 / 27) and +4.9% / +3.9% (alliance), both seeds; rank 4 ties or trails rank 3. At 20 to 25 rank 3 ties the top rank (within 0.3%); at 28 and 29 the top rank wins (rank 3 -1.4% / -2.0%).<br>Fix: a row, rank 3 at 26 and 27 (line 8 and the pre-pull bolt). Card note 8 and the leveling prose follow. | `ele-low.txt`, `round2.txt` | CONFIRMED | old, guides-handoff (guide rounds 6 to 15 carried it) | S adc341da2, wording S ab76630b5 |
| S6 | **Enhancement Earth Shock at 40 to 49 (round 18 N1).** The quarter-mana reserve lost 0.3% to 1.2% at every level 40 to 49, both seeds and factions, because the band's Maelstrom Weapon is 3/5 and a five-stack Lightning Bolt still costs 40% of its mana. There was no way to say "below 5/5" in a rotation.<br>Fix, engine: the Maelstrom Weapon trigger aura carries the talent rank as its ActionID tag (as mage's Missile Barrage). Site: the line keeps half its mana while Maelstrom Weapon is known below 5/5, and a quarter without the talent or at 5/5. A player at 37 to 41 holds 1 to 4 points (the guide's order), so the rule is the real one, not a band artifact.<br>Result: 40 +0.4% / +0.6%, 43 +0.8% / +1.2%, 46 +0.9% / +0.8%, 49 +0.3% / +0.8% (horde / alliance); bands 30, 50, 60 and raid 60 identical to the hundredth. | `es-ba.txt`, `esgate.txt`; `TestMaelstromWeaponTriggerCarriesItsRank` (fork) and `TestEnhancementShocksDownToAQuarterOfItsManaOrHalfBelowFullMaelstrom` (site) failed, then passed; `make apl-check` clean | CONFIRMED | old (round 18 N1 left it open) | E 73b7e850e, E bd7852d9a (fork APL copy), S cff94136a |
| S7 | **Ladder golden header said Enhancement dual-wields** (`ladderRulesHeader`, sim/request/ladder.go:1283). The code gives it a shield. Header text fixed and tied to `ladderGearProfiles` by a test. No number moves. | `TestLadderRulesHeaderGearProseMatchesProfiles` failed, then passed | fixed (text) | old | S c701c6509 |
| S8 | enhancement.md:93 hit sentence: guide round 15 already rewrote it ("Your swings stop missing a raid boss at 9% hit and your spells at 16% ..."), b2aee8a38. Checked against the client hit rules (9% melee and 16% spell against a level 63 target); wow-player SHIP on the rendered page. | cap/enh-stat-priority.png | Aligned | - | - |
| S9 | Low-rank proc watch (priest P1): no shaman row casts a share-0 rank into a chance proc. The lowest share any row casts is 90% (Chain Heal 1 at 59, Lightning Bolt rows at 95% on their last level). Elemental Focus, Maelstrom Weapon and Water Shield read the cast, not the rank, in the engine. | spell_max_level_auto_gen.go, rows | Watch (unmodelled curve) | - | - |

## Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | S6 above. | | CONFIRMED | old | E 73b7e850e, S cff94136a |
| N2 | Rank sweep 40 to 59 (Earth Shock one rank down at 40, 49, 50, 53; the Maelstrom Lightning Bolt one rank down at every level): the highest rank learned wins everywhere by 0.6% to 1.4%. No rows. | `enh-sweep-h1.txt` | Aligned | - | - |
| N3 | Spell power coefficients against the client (the mage Ice Lance pattern): every damage spell Enhancement casts matches a non-zero client row (Lightning Bolt 0.714, Earth Shock, Flame Shock 0.214 / 0.1, Searing Totem attack 0.017, Magma 0.033, Fire Nova 0.143, Flametongue Attack 10444 0.1, Frostbrand 0.1, Lightning Shield 0.267). The generator-default coefficients on zero client rows (Flametongue proc dummies, Stormstrike 17364, Riptide 409954, Mana Spring, Fire Nova 409938) are read by no code. | 70790 SpellEffect (identical to the re-layered 70755); grep of the fork | Aligned | - | - |
| N4 | Raid build utility (Improved Ghost Wolf 2): round 17 N3 ruled it a deliberate raid-movement pick the guide explains; it is the 24th and 25th point that opens Maelstrom Weapon. Not reopened. | | Aligned | - | - |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | S5 above (26 and 27). | | CONFIRMED | old, guides-handoff | S adc341da2 |
| E2 | S4 above (40 to 59 rows hold). | | Aligned | - | - |
| E3 | Talents against 70755 text: Elemental Focus is one charge on the next damage spell (engine MaxStacks 1), Lightning Overload 30% at half damage, Call of Thunder 5%, Elemental Fury +100%, Lightning Mastery 0.5 s and no Chain Lightning cooldown for 15 s, Concussion 5%, Convection 10%, Tidal Focus hit on the merged hit stat. Raid 60 unchanged (666.1 horde, 669.4 alliance). | talents/shaman.json; sim/shaman/talents.go | Aligned | - | - |
| E4 | Elemental Devastation 1 and Elemental Warding 3 in the raid build stay pending owner (round 17 E4). | | pending owner | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | S1 (Chain Heal) and S2 (Healing Wave 56). | | CONFIRMED | old | S 472dcaaf9, adc341da2 |
| R2 | Chain Heal against the client: 3 targets, 0.5 a jump, 0.714 (SpellEffect EffectChainTargets 3, EffectChainAmplitude 0.5) as chain_heal.go. Riptide: no direct heal, 12 s, 150/180/230 mana, ticks 89/115/161 at 0.1. The party-only jump rule (carried) changes nothing in the heal profile, whose five members are one party. | SpellEffect rows 1064/10622/10623/408521/1239243; riptide.go | Aligned | - | - |
| R3 | Restoration raid 60 unchanged (757.1 horde, 759.4 alliance). | | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Fork tests in a fresh worktree panic with "No item with id" (TestElemental 15684, TestEnhancement 18817, restoration relic tests 279249, TestStormstrikeUsesNormalizedWeaponDamage). The base engine at 6ea675594 fails the same way, so it is the worktree's item database, not this lane. The new fork test runs without it. | `go test ./sim/shaman/...` on base and lane | note | - | - |

## Before/after (10,000 iterations, frozen BiS, seeds 1 / 50001)

Before = $R/sb-before from $R/site-before; after = lane engine bd7852d9a + site cff94136a (the wording commit after it moves no number). Restoration is effective HPS.

| Spec, level | Faction | Seed 1 before to after | % seed 1 / 50001 |
|---|---|---|---|
| Restoration 53 (bare 50) | horde | 263.6 to 265.5 | +0.72 / +0.65 |
| Restoration 53 (bare 50) | alliance | 264.2 to 265.9 | +0.64 / +0.57 |
| Restoration 56 (bare 50) | horde | 258.8 to 262.6 | +1.47 / +1.43 |
| Restoration 56 (bare 50) | alliance | 258.0 to 262.7 | +1.82 / +1.82 |
| Restoration bare 60 | horde | 323.5 to 326.9 | +1.05 / +1.05 |
| Restoration bare 60 | alliance | 330.6 to 333.7 | +0.94 / +1.00 |
| Elemental 26 (bare 20) | horde | 32.27 to 33.51 | +3.84 / +3.71 |
| Elemental 26 (bare 20) | alliance | 32.69 to 34.32 | +4.99 / +4.86 |
| Enhancement 43 (bare 40) | horde | 159.37 to 160.59 | +0.77 / +0.78 |
| Enhancement 43 (bare 40) | alliance | 158.51 to 160.34 | +1.15 / +1.04 |
| Raid 60, all three specs | both | unchanged | 0.00 |

Per level, Enhancement 40 to 49 and Restoration 46 to 59 are in S6 and S1.

## Files to regenerate

- Every rotation ladder golden (sim/request/testdata/ladder/*.golden.md): the shared rules header changed (S7). `TestRotationLadder` fails until they are regenerated. shaman-enhancement's golden also moves at 40 (Earth Shock), shaman-restoration's at 50 (none of the rungs 46 to 59 are on the ladder but 50, where Chain Heal rank 1 now casts), shaman-elemental's does not (no rung at 26 or 27).
- data/builds/1.60.1.70755/addon-data.json, then addon/ForeverSixty/Data.lua: the new rank rows and card notes 7 (Enhancement), 8 (Elemental), 16 and 17 (Restoration).
- Coach starters, if they carry these rotations.
- Hand-synced, not regenerated: sim/rankchoices/rank-choices.json (plain copy), sim/request/apl/shaman-{elemental,enhancement,restoration}.apl.json, fork ui/enhancement_shaman/apls/forever_enhancement.apl.json (E bd7852d9a).

## Measure at refresh

- Restoration bands 40 and 50 and bare 60 with the Chain Heal rows (weights shift a little toward mana).
- Enhancement band 40 with the half-mana reserve.
- Carried: Enhancement band 30 with Stormstrike; the Enhancement relic slot (round 18 S3).

## For other lanes / integrator

- **The site commit cff94136a needs the fork commit 73b7e850e.** On the old engine the 5/5 tag never resolves, so Enhancement would keep half its mana at 50, 60 and raid (a regression). Land and pin them together.
- The Maelstrom Weapon trigger now has an ActionID, so it shows in aura metrics (100% uptime). No damage changes.
- Fork test environment (X1): fresh worktrees lack the item database the fork's integration tests read.
- Platform: the "Hit cap: reached" web line (owned by platform) is unchanged here; the engine weight side was fixed in round 17.

## For the guide lane

Shipped with this lane (wow-player review on the rendered pages, captures in $R/shaman/cap):
- restoration.md leveling (Chain Heal rank 1 until 59; Healing Wave rank 6 until 55, rank 7 from 56 to 58) and the Chain Heal ranks paragraph ("starts losing your bonus at 58"): SHIP.
- enhancement.md loop item 6 and the Mana paragraph (half mana until Maelstrom Weapon has five points): SHIP. Optional, not taken: item 1's "instant and free" holds only at five points.
- enhancement.md:93 hit sentence: SHIP, unchanged.
- elemental.md leveling bolt ranks: HOLD on the first capture (28 and 29 were not named), fixed to "rank 3 at 26 and 27 ..., the top rank at 28 and 29, then rank 3 again from 30 to 33": SHIP on the re-capture (cap/ele-leveling-2.png), with the Chain Heal ranks re-capture (cap/resto-chainheal-ranks-2.png).

The rendered Priority cards still show the old notes until addon-data regenerates.

## Still open

- Pending owner, not reopened: Healing Wave 2 tooltip on a geared 60; Lightning Overload rank spells; level-60 book ranks (Strength of Earth 5, Grace of Air 3, Healing Wave 10); Elemental Devastation 1 and Elemental Warding 3 as the Elemental tolls.
- The low-rank proc curve (priest P1) stays unmodelled; no shaman row depends on it (S9).
- Logs: foreverlogs live logs start after the 23:00 UTC launch. Last looked 18:47 UTC: `api.foreversixty.gg/v1/reports/recent` lists one beta report (2026-10-20); nothing new to check against. Warcraft Logs is blocked (gcloud login expired); not worked around.
