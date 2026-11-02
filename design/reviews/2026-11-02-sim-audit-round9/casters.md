# Sim audit round 9: casters, 2026-11-02

Specs: mage-arcane, mage-fire, mage-frost, priest-shadow, druid-balance, shaman-elemental.

## Method

- **Engine.** Lane branch `lane/r9-casters` from the pin `40021374a`. The before column runs a second worktree detached at the pin.
- **Site.** Branch `wt/r9-casters` from origin/main `55a1212f7`, with `make simdb`. The before column runs a site worktree detached at origin/main.
- **Client.** 1.60.1.70707 tables with the hotfix push layered in (fetched 2026-11-01 21:42 UTC): SpellEffect, SpellPower, SpellAuraOptions, SpellName, SkillLineAbility, SpellLevels, ItemEffect, ItemXItemEffect, and the BaseMP gametable. Wowhead Forever tooltips (nether.wowhead.com/forever, read 2026-11-02) are the second source where noted.
- **Runs.** `spec-breakdown` at 10,000 iterations, seeds 1 and 50001, both factions, bare 20 to 50 and raid 60. Grids are `before.tsv` and `after.tsv` in the lane scratch dir.
- **Downrank probes.** A scratch build of spec-breakdown that pins one rank so the rank rewrite leaves it alone (never committed). Files: `downrank.py`, `confirm.py`, `downrank-candidates.md` and `elemental-mana.md` in the lane scratch dir.
- **Tests.** Every confirmed bug got a failing test first. I saw it fail, then saw it pass.

## Assigned items

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | **Downrank candidates for the ranks lane (owner ruling 3).** Every lower learned rank of each filler, at every band, gated on the mana pace or all fight. Elemental gains the most: Lightning Bolt rank 4 is +8.2% at bare 30, +15.0% at 40 and +38.9% at 50. At raid 60, rank 9 behind the pace is +2.1%. Fire: Scorch rank 1 is +5.0% at 40 and +5.7% at 50. Balance: Wrath rank 4 behind the pace is +7.2% at 30. Shadow: Mind Blast rank 1 is +2.3% at 40, rank 2 +1.9% at 50. Arcane (+1.3%) and Frost (+0.6%) gain at 30 only, once Evocation is in (M1). Band 20 never gains, because its lower ranks pay the sub-20 penalty. | Full table: "For other lanes" below and `downrank-candidates.md`. The integrator asked for the Elemental mana model to be audited first: `elemental-mana.md`. It found one mana bug, X1, which made the shortage larger. | Measured; routed to the ranks lane | - |
| I2 | Fire's Improved Scorch upkeep line, re-measured on this lane. Dropping it moves raid 60 by +0.6 / +0.3% (alliance) and +0.3 / +0.1% (horde). | `scorch/drop.json`, 10,000 iterations, seeds 1 / 50001 | Held (under 1%) | - |
| I3 | Arcane 20 Horde and Arcane 30 Alliance. These are ranker rows: BiS sets within error of origin/main's. Nothing here re-ranks gear. M1 changes how every bare mage spends mana, so every mage band needs a re-rank before those rows can be judged again. | design/reviews/2026-11-02-sim-audit-round8/ranker.md:147 | Routed (ranker lane, after M1) | - |
| I4 | Caster debuffs by band, after round 8 moved the caster-only debuffs out. The raid preset carries Curse of the Elements, ranked by level (4/6/8/10% at 20/30/40/50; `core/raid_debuff_ranks.go`). It carries no Shadow Weaving, Improved Scorch, Improved Shadow Bolt or Stormstrike. Each caster still builds its own: raid Shadow shows Shadow Weaving (15332 talent aura) at 99.5% uptime, and raid Fire keeps its own 5 Improved Scorch stacks. Bare bands carry no raid debuffs, by design. **One flag left: Curse of Shadow (X2).** | breakdown aura tables; `TestRaidDebuffsFollowTheClientRanks` passes | 1 CONFIRMED-BUG (X2) | E 35fa13a5a |
| I5 | Crit-table findings for the crit lane: none new. Elemental Devastation is melee-only and already takes its crit back off spells (`shaman/talents.go` applyElementalDevastation). Nature's Majesty (druid) is "spells and melee" in its text, so the unified stat is right there. | client text, engine | Aligned | - |

## Mage: all three specs

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| M1 | **Arcane and Frost never cast Evocation.** Only Fire's rotation had the line, so a bare-band Arcane or Frost mage spent a third to a half of the fight wanding with Evocation ready. Frost bare 40 was out of mana for 60 s of 180. All three rotations now cast it below 5% mana while more than 30 s of the fight remain. Fire's raid mage also stops evocating in the last seconds, when the returned mana cannot be spent. | Evocation 12051: learned at 20 by every mage (SkillLineAbility, SpellLevels). Its effects: aura 110, +1500% regen; aura 134, 100% while casting; 8 s channel, 8 min cooldown (SpellEffect 689809/689810; tooltip "mana regeneration is active and increased by 1500%"). Swept below 5/10/20/30% mana and more than 10/20/30/45 s left: 5% with 30 s is best or level at every band, and the only gate that does not cost Arcane at raid 60. Test first: `TestBareMagesEvocateWhenTheyRunDry` (8 Arcane and Frost bands at 0 casts, then pass). The fork's copies are synced; the TestP1Mage golden moves with it. | CONFIRMED-BUG (rotation; client rows) | S fc98a54a8, E 388fbc7c8 |
| M2 | Mana gems are used smallest first. The gems share one 2 min cooldown, and each pops as soon as the pool is short by its size, so a level-60 mage spends the slot on an Agate. Tried largest-carried-first: Arcane -0.2 to +0.5%, Frost -0.1 to +0.3%, Fire raid -0.4%, bare 40 -1.6%. Waiting for a big enough deficit costs as much as it saves. | `mage/mana_gem.go`; prototype measured, then reverted | Aligned (measured; no better policy) | - |
| M3 | Checked against the client: Mage Armor ranks 6117/22782/22783 (5/10/15 magic resistance; 50% regen while casting, aura 134 base 50). Arcane Concentration's charge (also X1). Ranks cast per band: Frostbolt, Fireball, Scorch, Arcane Missiles, Arcane Blast (rank 3 at 40 is 1239697, round 8's fix) and Fire Blast are each the highest trainer rank. The Ahn'Qiraj book ranks (Frostbolt 25304, Fireball 25306, Arcane Missiles 25345) stay off with `core.IncludeAQ` (round 7). | client rows; breakdowns at every band | Aligned | - |

## Priest: Shadow

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Checked: Shadow Weaving self-stack at 99.5% uptime in raid. An Era raid gave a Shadow priest Curse of Shadow; Forever deleted it (X2) and the preset never carried it. Curse of the Elements is all-magic in Forever and carries Shadow, so nothing was lost. Mind Flay, Mind Blast, Shadow Word: Pain, Devouring Plague and Shadow Word: Death ranks per band. Touch of the Grave procs. Inner Focus into Mind Blast. Mind Flay downranks never win; Mind Blast's do (I1). | breakdowns; client 1311680 (aura 87 +10% over mask 126) | Aligned | - |

## Druid: Balance

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | Round 8 side effects. Innervate: first cast at 40, 1.00 a fight at bare 40, 50, 60 and raid; none at 20 or 30. The mana-pace Wrath keeps bare 40 to 60 under 8 s out of mana. Talons of Wrath (249441) is modelled (1248996: 50% chance on Wrath to restore 35 mana via 1302521; `druid/items.go:77`). Improved Wrath is -0.1 s a rank in wrath.go and -10% mana a rank as a mod (16814 effects: aura 107 -100 ms, aura 108 -50 mana), so nothing is counted twice. | breakdowns; client rows | Aligned | - |
| B2 | The druid conformance golden shows 24 "mismatch" rows for Faerie Fire at every level: cost 0, 6 s cooldown, 1 s GCD. That is round 8's Cat and Bear form spell, which shares the caster spell's id (Tag 1), being matched against the caster rows. The caster spell (`druid/faerie_fire.go:12`) is registered at rank 4 (9907) at every level. No published caster rotation casts it, since the raid preset's level-ranked aura covers it. | `sim/core/testdata/conformance/druid.golden.md`; `druid/faerie_fire.go:12,145` | Routed (hybrids lane: report artifact; caster Faerie Fire's rank) | - |

## Shaman: Elemental

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| E1 | Mana model audit at bare 30, 40, 50 and 60 and raid 60, asked for by the integrator before any downrank. Checked equal to the client: Intellect to mana (20 + 15 a point above 20); 17 + Spirit/5 a tick; the 5 s rule; Mindfulness 17/33/50%; Mana Spring ranks and pulses (4/6/8/10 every 2 s; cost 80/100 at ranks 3/4); Convection; and every Lightning Bolt, Chain Lightning, Lava Burst, Flame Shock and Searing Totem cost (SpellPower). The pool probe at bare 50, 3735 = 1255 + 20 + 15 x 164, matches. Water Shield is a Restoration talent; Lightning Shield and Totemic Mastery are not cast or not in Forever; Eye of the Storm is pushback only. Found X1 and X3. | `elemental-mana.md` in the lane scratch dir | 2 CONFIRMED-BUG (X1, X3) | see X1, X3 |
| E2 | Checked: Hand of Edward the Odd (2243) is a raid main hand for Elemental, Balance and Shadow, picked for Forever's +90 spell damage. Its "chance on hit" Focus (18803) is a melee proc, and these casters never swing, so it never fires. Right. Kindling Stave's Ignition is modelled (round 8 warlock lane) and shows at bare 50. | Wowhead Forever item 2243; breakdowns | Aligned | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Clearcasting gave a second free spell after an instant.** Elemental Focus (16246) and Arcane Concentration (12536) are one charge each (ProcCharges 1, "your next damage spell"). The consume guard, there to spare the cast that procced it, skipped every cast that completed while the aura still had its full duration. An instant at the proc's own moment (Elemental's "Fire Nova on Clearcasting" line, Flame Shock; a mage's Fire Blast or Ice Lance after a hit) was therefore free and left the charge for the next spell too. Bare 50 Elemental counted 6.7 free casts a fight from 4.3 procs; after the fix, 4.0 from 4.0. `core.ProcSource` records the granting cast and its moment, and only that cast at that moment keeps the charge. | SpellAuraOptions 16246 (ProcCharges 1); mage 12536 (round 7 test). Measured: with Elemental Focus off, every cast pays exactly 0.9 of its SpellPower cost, so the extra free casts come from the charge. Tests first: `TestAnInstantAtTheProcsMomentSpendsClearcasting` in shaman/elemental and in mage (both failed, then passed). | CONFIRMED-BUG (engine) | E ff4bae666 |
| X2 | **The raid's Curse of Shadow flag gave Era's curse.** Forever deleted Curse of Shadow: 17862 and 17937 have no SpellName or SpellEffect row (their SkillLineAbility rows are orphans), and Wowhead Forever returns "Entity not found". Curse of the Elements (1311680) is the all-magic curse. The flag still gave +10% Shadow and Arcane damage taken and lowered every resistance by 75 twice. With Curse of the Elements on, the two share a school-damage category, so no published number or golden moves. Alone, from the custom /sim panel, it added 10%. The flag now applies nothing. | client tables; tooltip 17937 (2026-11-02). Test first: `TestTheRaidsCurseOfShadowIsGone`. | CONFIRMED-BUG (custom panel only) | E 35fa13a5a |
| X3 | **Base mana by level came from wowhead's planner, not the client.** The planner's mana disagreed with the client's BaseMP table at 8 class levels: shaman 1, 22, 26, 30 and 52 (718 against 665 at 30), priest 2 and 12, and warlock 1. Mage, druid, paladin and hunter agreed at every level. The generator now takes every level's Mana from basemp.txt. Health stays the planner's: the client has no base-health table. | `gametables/basemp.txt` (1.60.1.70707); `tools/base_stats_parser.py`. Test first: `TestBaseManaAtEveryLevelIsTheClients` (8 rows failed, then pass). The shaman conformance golden moves: Stormstrike's 8% of base mana at 30 is now 53.20. | CONFIRMED-BUG (client table) | E 918047272 |

## Before/after

Before: pin 40021374a on origin/main 55a1212f7. After: lane engine 918047272 on site fc98a54a8, committed simdb. 10,000 iterations, seeds 1 / 50001. Shadow and Balance are 0.0% at every band on both factions and are left out. Enhancement is shown because X1 and X3 touch shaman code; it does not move. Restoration shaman at band 30: X3, under For other lanes.

| Spec, band | before (horde, seed 1) | after | change, horde (seed 1 / 50001) | change, alliance | causes |
|---|---|---|---|---|---|
| mage-arcane bare20 | 36.5 | 44.4 | +21.6 / +21.6% | +21.0 / +21.0% | M1 |
| mage-arcane bare30 | 75.6 | 82.8 | +9.5 / +9.7% | +9.8 / +9.8% | M1 |
| mage-arcane bare40 | 155.5 | 161.6 | +3.9 / +3.9% | +4.1 / +4.1% | M1 |
| mage-arcane bare50 | 196.5 | 208.4 | +6.1 / +5.8% | +5.4 / +5.3% | M1 |
| mage-arcane raid60 | 601.4 | 600.7 | -0.1 / -0.1% | -0.1 / -0.2% | M1 (within error) |
| mage-fire bare20 | 35.2 | 35.2 | +0.0 / +0.0% | +0.0 / +0.0% | - |
| mage-fire bare30 | 70.9 | 70.9 | +0.0 / +0.0% | +0.0 / +0.0% | - |
| mage-fire bare40 | 147.0 | 147.0 | +0.0 / +0.0% | +0.0 / +0.0% | - |
| mage-fire bare50 | 182.2 | 182.2 | +0.0 / +0.0% | +0.0 / +0.0% | - |
| mage-fire raid60 | 670.8 | 670.2 | -0.1 / -0.1% | +0.1 / +0.1% | M1 time gate (within error) |
| mage-frost bare20 | 36.7 | 44.7 | +21.8 / +21.8% | +21.7 / +21.7% | M1 |
| mage-frost bare30 | 62.7 | 77.0 | +22.8 / +22.8% | +19.7 / +19.7% | M1 |
| mage-frost bare40 | 144.8 | 185.9 | +28.4 / +28.4% | +26.1 / +25.9% | M1 |
| mage-frost bare50 | 193.8 | 248.3 | +28.1 / +28.2% | +24.6 / +24.6% | M1 |
| mage-frost raid60 | 601.5 | 601.5 | +0.0 / +0.0% | +0.0 / +0.0% | - |
| shaman-elemental bare20 | 22.5 | 22.5 | +0.0 / +0.0% | +0.0 / +0.0% | - |
| shaman-elemental bare30 | 46.2 | 44.0 | -4.8 / -5.0% | -5.0 / -5.0% | X1, X3 (X3 alone: -0.9%) |
| shaman-elemental bare40 | 83.2 | 79.0 | -5.0 / -4.9% | -5.0 / -4.9% | X1 |
| shaman-elemental bare50 | 108.9 | 103.2 | -5.2 / -5.1% | -4.8 / -4.9% | X1 |
| shaman-elemental raid60 | 476.8 | 465.6 | -2.3 / -2.3% | -2.5 / -2.4% | X1 |
| shaman-enhancement bare20 | 33.8 | 33.8 | +0.0 / +0.0% | +0.0 / +0.0% | - |
| shaman-enhancement bare30 | 69.0 | 69.0 | +0.0 / +0.0% | +0.1 / +0.0% | - |
| shaman-enhancement bare40 | 123.9 | 123.9 | +0.0 / +0.0% | +0.0 / +0.0% | - |
| shaman-enhancement bare50 | 160.3 | 160.3 | +0.0 / +0.0% | +0.0 / +0.0% | - |
| shaman-enhancement raid60 | 594.0 | 594.0 | +0.0 / +0.0% | +0.0 / +0.0% | - |

## For other lanes

- **Ranks lane: the downrank candidates.** 10,000 iterations, seeds 1 / 50001, horde, absolute at seed 1, on this lane's engine and site. "pace" means a lower-rank copy of the filler, ahead of it, while `currentManaPercent < remainingTimePercent`; "all fight" means the filler at that rank. The Elemental rows include X1.

| spec | band | spell | rank tried | gain (seed 1 / 50001) | DPS, seed 1 | out of mana, s (before to after) |
|---|---|---|---|---|---|---|
| shaman-elemental | bare 30 | Lightning Bolt | rank 4 (915), always | +8.3% / +8.1% | 44.4 to 48.1 | 94.7 to 73.9 |
| shaman-elemental | bare 40 | Lightning Bolt | rank 4 (915), always | +14.9% / +15.0% | 79.0 to 90.8 | 101.7 to 58.4 |
| shaman-elemental | bare 50 | Lightning Bolt | rank 4 (915), pace | +39.0% / +38.8% | 103.2 to 143.4 | 98.8 to 26.1 |
| shaman-elemental | bare 50 | Lightning Bolt | rank 4 (915), always | +39.1% / +38.9% | 103.2 to 143.5 | 98.8 to 24.7 |
| shaman-elemental | raid 60 | Lightning Bolt | rank 9 (15207), pace | +2.2% / +2.0% | 465.6 to 475.7 | 8.8 to 2.9 |
| mage-fire | bare 40 | Scorch | rank 1 (2948), always | +5.0% / +5.0% | 147.0 to 154.3 | 55.6 to 25.1 |
| mage-fire | bare 50 | Scorch | rank 1 (2948), always | +5.8% / +5.6% | 182.2 to 192.7 | 58.2 to 20.9 |
| mage-fire | bare 40 | Scorch | rank 1 (2948), pace | +3.7% / +3.7% | 147.0 to 152.5 | 55.6 to 28.9 |
| mage-fire | bare 50 | Scorch | rank 1 (2948), pace | +4.4% / +4.2% | 182.2 to 190.3 | 58.2 to 24.8 |
| druid-balance | bare 30 | Wrath | rank 4 (5179), pace | +7.2% / +7.2% | 65.5 to 70.2 | 48.0 to 36.3 |
| priest-shadow | bare 40 | Mind Blast | rank 1 (8092), always | +2.3% / +2.4% | 125.6 to 128.5 | 70.8 to 63.7 |
| priest-shadow | bare 50 | Mind Blast | rank 2 (8102), pace | +1.9% / +1.8% | 189.2 to 192.8 | 67.4 to 57.2 |
| mage-arcane | bare 30 | Frostbolt | rank 4 (7322), pace | +1.3% / +1.3% | 83.0 to 84.1 | 21.8 to 12.9 |
| mage-frost | bare 30 | Frostbolt | rank 4 (7322), pace | +0.6% / +0.5% | 77.0 to 77.5 | 21.9 to 8.5 |

  - Best lower rank at each band (screen, 3,000 iterations): `downrank-candidates.md`. Where no row is listed, no lower rank beat the top one. That covers every band 20 for every spec, Balance at 40 and above, Shadow at 20, 30, 60 and raid, Fire at 20, 30, 60 and raid, and Frost and Arcane outside band 30.
  - Elemental's mana model: `elemental-mana.md`. The shortage is real: about 85 mana a second spent against about 16 regenerated.
- **Hybrids lane.** The consume guard X1 fixed is also on Omen of Clarity (`druid/talents.go:373`) and Maelstrom Weapon (`shaman/talents.go:838`). An instant at a melee proc's moment (Shred, a 5-stack instant Lightning Bolt) does not spend the charge. `core.ProcSource` (E ff4bae666) is the fix to reuse. Also B2: the conformance report matches the form Faerie Fire against the caster rows, and the caster spell is rank 4 at every level.
- **Warlock lane.** `registerCurseOfShadowSpell` (`warlock/curses.go:315`) still registers a spell Forever deleted (X2), at levels 50 and 60 only. No warlock APL casts it.
- **Healers lane.** X3 lowers Restoration shaman at band 30 by 0.8 to 0.9% (85.9 to 85.2 HPS horde, 86.5 to 85.8 alliance).
- **Ranker lane.** M1 and X1 change mage and Elemental mana at every bare band, so their stat weights and picks need a re-rank. The Arcane 20 Horde and 30 Alliance rows (I3) are part of that.
- **Platform lane.** The custom /sim panel offers `curse_of_shadow`. It now applies nothing, like the four caster-only flags round 8 emptied.
- **Integrator.** sim/request TestRotationLadder goldens move for mage-arcane and mage-frost (M1), and for every shaman spec at band 30 (X3) and Elemental (X1). I did not regenerate them.

## Still open

- X3's planner-vs-client base health: there is no client table to check it against.
- I did not re-run the raid talent searches this round. The published builds are unchanged, and M1 and X1 change mana-bound bare bands, where the adoption rule keeps the guide builds.
- Carried, blocked on level-60 logs: Kindling Stave's rate, Fire Power on the burns, Ignite's 0.73, Ice Lance's coefficient, Missile Barrage on a miss, Shadowfiend, Searing Totem's 2.2 s attack, Starfire rank 7 at trainers, Wrath's cost stacking.

## Sources

- 1.60.1.70707 client tables with the hotfix push (fetched 2026-11-01 21:42 UTC): SpellEffect, SpellPower, SpellAuraOptions, SpellName, SkillLineAbility, SpellLevels, ItemEffect, ItemXItemEffect; gametables basemp.txt.
- Wowhead Forever tooltips, read 2026-11-02: spells 12051, 17937, 18803, 20572, 22783, 1248996, 1292997, 1302521, 1322312; items 2243, 744, 5198, 249441, 273049.
