# Sim-accuracy round 16: shaman lane

Specs: Elemental, Enhancement, Restoration. Client 1.60.1.70755 with the late hotfix re-layer. Site branch wt/r16-shaman (off int/r16 c62d77ab8), fork branch lane/r16-shaman (off int/r16 9f7f1f6e4).

**4 confirmed bugs, all fixed.** By lane: Elemental 1, Restoration 3, Enhancement 0.

By origin:
- R15 a0497d1bc: 2. Elemental and Restoration rank rows were measured only at the band floor.
- R15 a0497d1bc plus cf5fe63bb: 1. The Nature's Swiftness save shared the filler's rows; guides-handoff (guide round 10 "Sims requested 2").
- old: 1. The Nature's Swiftness save fired without the talent.

Two shared tool changes made the fixes possible:
- spec-breakdown `-level`.
- Line-scoped rank rows.

Neither moves any spec outside shaman.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| 1 | **Rank rows were measured only at a band's lowest level.** Under the level-gap rule (full bonus until 17 levels past the learn level, then 5% less a level), a rank picked at 30 is 8 levels deeper in the penalty at 39. The best rank climbs within the band. **Decision: split the bands, measured level by level, not one midpoint pick.** The best rank changes 1 to 3 times inside each ten-level band, so no single pick holds a band. New tool: `spec-breakdown -level` runs a band's published gear and talents at any level inside the band (S 22d545cb5). Details in E1 and R1. | Elemental filler, horde, 10,000 iterations, seeds 1 / 50001 (probes in $R/shaman/ele-lb-10k.txt):<br>- level 39: rank 3 57.3 vs rank 5 64.0 (+11.8%)<br>- level 49: rank 4 85.6 vs rank 6 111.1 (+29.7%)<br>- level 59: rank 6 132.3 vs rank 8 158.0 (+19.6%)<br>Restoration filler (resto-hw-10k*.txt):<br>- level 39: rank 3 103.7 vs rank 5 107.3 (+3.5%)<br>- level 49: rank 5 164.2 vs rank 6 172.9 (+5.3%) | CONFIRMED | R15 a0497d1bc | S 54a017916 |
| 1b | **Shared part, flagged for other lanes.**<br>- Every class's rows carry the same floor-only caveat.<br>- The addon resolves ranks only at its rungs 10/20/30/38/40/50/60 (data/pipeline/addonrotation.py LEVEL_BANDS, tied to the ladder rungs by test). So a row that starts mid-band reaches the addon only at the next rung. The log coach already reads rows per level (rankchoices.ChosenRanks). | see "For other lanes" | - | - | - |
| 2a | Rockbiter under Windfury Totem (sim/request/preset.go swapImbueUnderWindfuryTotem).<br>- The swap fires only when the buffs carry `windfury_totem`, from level 32 (rank 1 of the totem).<br>- It replaces only `main_hand_imbue:windfury_weapon`.<br>- The engine's Rockbiter is client 16313 (554 + 16.5 a level to 62: 653 at 60), times Elemental Weapons 3 (+20%, the 70755 text "Rockbiter Weapon by 20%"). | rockbiter_weapon.go, talents/shaman.json 104750. Raid 60 Enhancement unchanged by this lane (855.3 / 855.6 both sides). | Aligned | - | - |
| 2b | Lightning Bolt rank 10 at raid 60. The filler is now authored at rank 10, so a raid casts it as written and the raid-60 row is gone. Same casts as before. | Raid 60 Elemental: 666.1 / 666.6 (horde) and 669.4 / 669.5 (alliance), identical before and after. | Aligned | - | - |
| 2c | Stormstrike at 25. RequiredLevel 40 (stormstrike.go) feeds only the physical-exempt spell-level penalty. | -level 30 and 39 on the band-30 entry with Stormstrike: 29.8 and 29.3 casts per fight. The frozen published band-30 string (055030310040) still lacks the point; it refreshes with the BiS run. | Aligned | - | - |
| 2d | Nature's Swiftness Healing Wave rank 9. It held at 50 to 60. **But below 50 it took the filler's rank row** (rank 5 at 44 to 49 on the shared authored id 10396), which guide round 10 also flagged. Fixed by scoping the rows to the filler's line: the save now casts rank 7 at 44 to 47, rank 8 at 48 to 55 and rank 9 from 56. | TestRestorationShamanSwiftnessHealingWaveIsTheTopRank now also checks 44 and 49. | CONFIRMED | R15 a0497d1bc (on cf5fe63bb's shared id), guides-handoff | S e332dc8f5, 6eb3d7083, 54a017916 |
| 2e | Riptide order and rows. It sits right after the save, as published. The client's 408521/1239242/1239243 match riptide.go: 15 s, 5 ticks of 89/115/161 at 0.1, cost 150/180/230, a 12 s shared cooldown, and effect 2 (+25% to Chain Heal). The ranks are never more than 17 levels behind inside their bands, so they need no row. | spellconst/shaman.json 70755 | Aligned | - | - |
| 2f | Re-picked Elemental and Restoration rows: items 1 and E1/R1. Flame Shock's rank-3 row at 40 to 49 now loses to rank 4 and is dropped (E2). | | see rows | | |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | **The filler Lightning Bolt's rows were wrong for most levels of each band.**<br>New rows: rank 3 to 33, 4 to 38, 5 to 44, 6 to 50, 7 to 56, 8 from 57 and at bare 60.<br>The filler (line 8) and the opener are now authored at rank 10, with their rows scoped to those two entries (`lines: [8]`, `prepull: [1]`). The top-rank line 7 keeps the highest rank learned, and the addon at 60 shows rank 10, the raid cast. Before, it showed rank 9, which no fight casts. | Each crossover measured on both neighbours, horde, seeds 1 / 50001:<br>- 33: rank 3 ties rank 4.<br>- 34: rank 4 60.8 vs 59.3.<br>- 39: rank 5 64.0 vs 63.6.<br>- 44: rank 5.<br>- 45: rank 6 106.7 vs 106.2.<br>- 50: rank 6 148.3 vs 147.8.<br>- 51: rank 7 149.0 vs 146.3.<br>- 56: tie.<br>- 57: rank 8 156.1 vs 153.2.<br>Band floors, both factions: 40 +0.9 to +1.2%; 30, 50, bare 60 and raid 60 unchanged. | CONFIRMED | R15 a0497d1bc | S 54a017916, E a79fc82e2 |
| E2 | Flame Shock rank 3 at 40 to 49. With the filler on rank 5 from 39, rank 4 wins at every level measured. The row is dropped. | 40: rank 4 108.3 / 108.4 vs rank 3 107.6 / 107.9 (+0.5 to +0.7%). 45: +1.0%. | CONFIRMED (folded into E1's measurement) | R15 a0497d1bc | S 54a017916 |
| E3 | Fresh re-audit of the engine values against 70755: every Elemental cast's spellconst row (Lightning Bolt 1 to 10, Chain Lightning 1 to 4, Flame/Earth/Frost Shock, Searing and Magma Totem, Fire Nova, Call of Thunder 5%) matches the layered tables of the latest re-layer to the point. The late re-layer (43dc0f074) changed only whitespace in shaman item and set text, no shaman spell or talent value. Talent curves are identical between the pre- and post-re-layer raw tables. The 2026-11-04 wago CSVs are the shipped rows; the pushed values are in spellconst. Elemental Devastation 1 remains the toll (O2, round 15). | Script diff in this lane: spellconst vs bri/site raw 70755 (layered), 0 shaman differences. Lightning Overload copy ranks are 14 vs 14.5 (rounding, pending owner). | Aligned | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **The tank filler Healing Wave's rows were wrong for most levels of each band.**<br>New rows (scoped to line 17): rank 2 to 25, 3 to 36, 4 at 37, 5 to 43, 6 to 56, 7 at 57 and 58, 8 at 59, then the top rank at 60.<br>Rank 2 is new: round 15 compared only rank 3 against the top. | Horde, seeds 1 / 50001:<br>- 20: rank 2 68.6 vs rank 3 66.6 (+3.0%; alliance +4.1%).<br>- 25: +0.9%.<br>- 26: rank 3.<br>- 37: rank 4 +0.4%.<br>- 38: rank 5 +0.4%.<br>- 44: rank 6 ties rank 5.<br>- 50: rank 6 248.8 vs rank 8 243.9 (+2.0%; alliance +2.2%).<br>- 57: rank 7 260.7 vs 258.1.<br>- 59: rank 8 267.7 vs 266.5.<br>Every fight still runs dry (at 29, rank 3 lasts 300.4 of 300 s: it ends empty). | CONFIRMED | R15 a0497d1bc | S 54a017916 |
| R2 | **The Nature's Swiftness save fired without Nature's Swiftness.** Its condition asked only `auraIsActive 16188`. The engine drops a condition on an aura the character does not have, so below 44 the line was a plain "Healing Wave on the tank under 25%" that no leveling shaman casts. At 40 it fired about 13 times a fight with rank 5 and lifted the published number. The line now also needs `auraIsKnown 16188`, the way Enhancement gates Maelstrom Weapon. | Bare 40, both factions: 168.1 to 159.4 (**-5.2%**), the honest number for the rotation as written. 20, 30, 50 and 60: no change from this line. | CONFIRMED | old (rounds 1 to 15 measured with the phantom line) | S 54a017916, E a79fc82e2 |
| R3 | Lead from R2: a cheap Healing Wave on a tank under 25% beats the Lesser Healing Wave line at bare 40. Rank 4 there gives 170.5 vs 159.3. The heal profile has no tank deaths, so it cannot weigh Lesser Healing Wave's 1.5 s answer to a falling tank. Classic practice keeps Lesser Healing Wave as the emergency heal. Not changed. | 3000 iterations, seed 1 | OPEN (the sim cannot model tank death) | - | - |
| R4 | Chain Heal party rule. The 70755 tooltip on every rank says: "If cast on a party member, the heal will only jump to other party members." The engine jumps raid-wide (chain_heal.go mostInjuredMember). The list casts Chain Heal only on the tank, which the fake raid seats in a second party (core.AddHealingFakeRaid), and from there raid-wide jumps are right. A Chain Heal cast on a party member would be wrong; nothing casts one. | spelltips/shaman.json 1064/10622/10623; client ImplicitTarget 45 (chain-heal ally) | Aligned for the published list (the engine lacks the party branch) | - | - |

## Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | Re-audit against the 70755 talent text and the engine:<br>- Stormstrike: 150% normalized weapon damage, 6 s cooldown, 8% base mana; its 20% one-charge bonus on Lightning Bolt / Chain Lightning / Earth Shock is applied as a multiplicative spell mod (1.20).<br>- Flurry; Mental Dexterity 50/100/150% Intellect to attack power.<br>- Mental Quickness 15/30% Intellect to spell damage and healing (aura 174/175).<br>- Maelstrom Weapon 4% a stack per rank (20% at 5/5); Shamanistic Focus -45% on shocks and Lightning Shield.<br>- Elemental Devastation 3/6/9% melee crit for 10 s on a direct spell crit.<br>- Rage of the Farseer: 30% attack speed, 25 s, 3 min, no global cooldown (spellconst 425336 aura 342).<br>- Elemental Weapons 7/13/20% Rockbiter, 13/27/40% Windfury.<br>All match. No rank rows; every fixed cast is a top rank with a small spell-power share. | talents/shaman.json 70755; stormstrike.go, talents.go:196-720 | Aligned | - | - |

## Cross-cutting (shared code, flagged)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Tool: spec-breakdown `-level`** (band to band+9). Rotation ranks and the heal profile resolve at that level. Lets any lane judge a row at the top of its band. | sim/cmd/spec-breakdown TestLevelRunsTheBandEntryAtThatLevel | tool | - | S 22d545cb5 |
| X2 | **Rank rows may be scoped to `lines` (priority-list entries) and `prepull` (pre-pull actions).** Two lines can cast the same authored id for different jobs.<br>Touches:<br>- sim/request (rewrite and the ladder harness's fired check)<br>- sim/rankchoices (types)<br>- data/pipeline/addonrotation.py (addon export)<br>A test holds each scoped line to a cast of the authored id. Only shaman rows use it, so no other spec moves. | TestRewriteScopesARankChoiceToItsLines, TestParseRankChoicesScopesRowsToLines, TestCuratedRankChoiceLinesCastTheirAuthoredSpell; pytest test_a_rank_choice_scoped_to_* | tool | - | S e332dc8f5, 6eb3d7083 |
| X3 | api/internal/coach TestWebFixtures_NoDrift is still red ("Battle Shout up 100%"), as in round 15's X1. Not from this lane. | go test ./internal/coach | pre-existing | - | - |

## Before/after (bare bands and raid 60, 10,000 iterations, frozen BiS, band entries)

| Spec, band | Faction | Seed 1 before to after | % seed 1 / 50001 |
|---|---|---|---|
| Elemental bare 40 | horde | 107.4 to 108.6 | +1.12 / +0.93 |
| Elemental bare 40 | alliance | 108.2 to 109.3 | +1.02 / +1.20 |
| Elemental bare 30, 50, 60; raid 60 | both | unchanged (e.g. raid 666.1 / 669.4) | 0.00 |
| Restoration bare 20 | horde | 66.6 to 68.6 | +3.00 / +3.00 |
| Restoration bare 20 | alliance | 68.4 to 71.2 | +4.09 / +4.09 |
| Restoration bare 30 | both | 115.4 / 116.1, unchanged | 0.00 / -0.09 |
| Restoration bare 40 | horde | 168.1 to 159.4 | -5.18 / -5.23 (R2, phantom line removed) |
| Restoration bare 40 | alliance | 168.8 to 160.3 | -5.04 / -5.04 |
| Restoration bare 50 | horde | 243.9 to 248.8 | +2.01 / +1.97 |
| Restoration bare 50 | alliance | 244.1 to 249.4 | +2.17 / +2.09 |
| Restoration bare 60, raid 60 | both | unchanged (raid 757.1 / 759.4) | 0.00 |
| Enhancement raid 60 | both | 855.3, unchanged | 0.00 |

Inside the bands, the published floors do not show the gains. Each is the new row against the old row at that level:

| Spec | Level | Gain |
|---|---|---|
| Elemental | 39 | +11.8% |
| Elemental | 49 | +29.7% |
| Elemental | 59 | +19.6% |
| Restoration | 39 | +3.5% |
| Restoration | 49 | +5.3% |
| Restoration | 55 | +2.4% |

## Files to regenerate

- sim/request/testdata/ladder/shaman-elemental.golden.md and shaman-restoration.golden.md (rows, the rank-10 authoring and the save's gate; not run here).
- addon-data.json, addon-builds.json and Data.lua. These carry the rank rows, the line scope, the rank-10 line 8 and the new card notes. Until then the guide card reads "Lightning Bolt rank 9 / Rank 9 for everything else".
- Coach starters, if they carry shaman ranks.
- Hand-synced, not regenerated:
  - sim/rankchoices/rank-choices.json (make rank-choices-sync)
  - sim/request/apl/shaman-elemental.apl.json and shaman-restoration.apl.json
  - fork ui/elemental_shaman/apls/forever_elemental.apl.json and ui/restoration_shaman/apls/forever_restoration.apl.json

## Measure at refresh

- Elemental bare 40 and Restoration bare 20, 40 and 50 picks and weights on the new rows and the gated save.
- Enhancement band 30 with Stormstrike (the frozen string still has Flurry 4 and no Stormstrike).
- Carried from round 15: Enhancement raid-60 weights with Rockbiter; presets.raid `elemental_sharpening_stone` vs raid-consumes `windfury_weapon`.

## For other lanes

- **Every lane with rank rows (priest, paladin, druid, warlock, mage, hunter):**
  - Each row was measured at its band floor. Use `spec-breakdown -level <band+9>` to judge it at the band's top.
  - For shaman the best rank moved 1 to 3 times per band, worth up to +30%.
  - Rows that share an authored id with a line that must not downrank can now name `lines` / `prepull`.
- **Platform:** the addon resolves ranks only at the rungs 10/20/30/38/40/50/60 (addonrotation LEVEL_BANDS, pinned to the ladder rungs). Mid-band rows (shaman 33, 34, 39, 45, 51, 57) reach the addon only at the next rung. Proposal: per-level rows or extra rungs at row starts.
- **Phantom lines (same bug as R2), in lines gated on a talent's aura with no `auraIsKnown`.** Each fires unconditionally at bands where the talent is not taken:
  - druid-restoration lines 2 and 3 (Nature's Swiftness 17116)
  - mage-arcane lines 4 and 5 (Presence of Mind 12043)
  - paladin-holy line 1 (Oath of Light 1328548)
  - priest-shadow line 3 (Vampiric Embrace 15286)

  Each lane should check its bands.

## For the guide lane

Shipped with this lane (wow-player SHIP on both rendered pages, no MUST):
- elemental.md Lightning Bolt ranks paragraph and Levels 20 to 60.
- restoration.md leveling paragraph.
- Card notes: Elemental 7 and 8, Restoration 17.

**Collision:** the uncommitted guide round 10 tree (scratchpad gr4/int) edits the same lines: elemental.md:80 and :116, Elemental notes 7 and 8, Restoration note 17. This lane took round 10's wording for :80 and notes 7 and 8 and extended it with the per-level ranks. Keep this lane's text on merge.

The reviewer's optional notes:
- Name Mana Tide Totem in "The talent teaches its first rank" (restoration.md:120).
- Give rank 8's mana cost beside rank 10 and rank 9 (elemental.md:80).

## Still open

- R3: Lesser Healing Wave vs a cheap Healing Wave as the leveling emergency heal (the sim has no tank deaths).
- Pending owner, not reopened: the Healing Wave 2 tooltip on a geared 60, Lightning Overload rank spells, level-60 book ranks, Elemental Devastation 1 as the Elemental toll.

Sources:
- client 70755 spelltips/shaman.json (Chain Heal)
- wago SpellEffect/SpellCooldowns 70755 (fetched 2026-11-04)
- [warcraft.wiki.gg Chain Heal](https://warcraft.wiki.gg/wiki/Chain_Heal)
- [wowclassicdb Chain Heal](https://wowclassicdb.com/tbc/spell/1064)
