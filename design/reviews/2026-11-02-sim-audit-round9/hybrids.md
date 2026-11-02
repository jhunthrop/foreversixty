# Sim audit round 9: hybrids (2026-11-02)

This report covers Feral cat, Feral bear (threat), Retribution, Protection paladin (threat) and Enhancement.

**Setup:**
- Client 1.60.1.70707, layered tables (hotfix push included) from the 2026-11-01 21:42Z fetch.
- Published sets and builds: the committed `data/builds/1.60.1.70707/bis/` (8baffe5fc nightly). Nothing nightly-owned was regenerated or committed.
- Engine: pin 40021374a plus `lane/r9-hybrids`. Site: `wt/r9-hybrids` on origin/main 55a1212f7.
- Second source: Wowhead Forever tooltips (nether.wowhead.com/forever/tooltip/spell/<id>, read 2026-11-02).
- Method:
  - every band (20 to 60 bare, 60 raid) of every spec was run and its cast and aura tables read;
  - a per-level scan listed every trainable spell (and its triggered spells) with EffectRealPointsPerLevel, checked against the engine;
  - DefenseType (crit table) was read for every hit;
  - every talent in every published build was checked for engine code.

Fix-sha column: "E" is a fork commit on `lane/r9-hybrids`, "S" a site commit on `wt/r9-hybrids`.

**In short: 5 CONFIRMED-BUGs, all fixed.** Two of them move published numbers: Enhancement's Air totem (bare 40 to 60 and raid 60) and its imbue attack power at 20 and 30. The other three move nothing measurable in my specs:
- the Clearcasting and Maelstrom consume guard (Restoration druid and Balance move a little);
- the display ids.

## Assigned items

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Side effects of round 8, every band. Checked:<br>- Shred rank 1 at 20: Claw until Bite, then Shred.<br>- Faerie Fire (Feral): ranks 770/778/9749 at 20/30/50, 40 s, about 17% spell-table misses.<br>- Stormstrike normalized.<br>- Pummeler Haste: 30 s, 1.5 uses a fight, at 30 to 50.<br>- Weapon procs: Fiery War Axe's Fireball at Enhancement 40, Devious Strike at 20.<br>- Chthonic Power and Molten Fury.<br>- Raid 60: Leader of the Pack, CoR, foods and debuffs unchanged since round 8's measure.<br>Below 60 the published preset is bare, so raid buffs by level never reach a published number. | `$R/bands/*.txt` breakdowns | Aligned | - |
| A2 | Faerie Fire's threat in Bear Form.<br>- A combat log carries no threat: foreverlogs.gg's events and ability breakdowns have no threat field, and its non-public API now returns 403.<br>- The client row has no amount (9907 effect 3, Effect 63, no points).<br>- What would settle it: an in-game read of `UnitDetailedThreatSituation` raw threat on a fresh mob before and after one bear Faerie Fire, on the beta or at launch.<br>- About 0.1% of the bear's threat. | foreverlogs.gg openapi.json; client 9907 | OPEN | - |
| A3 | Tank threat, checked against the client rows and the live text:<br>- Bear Form: +50% (21178 aura 10). Maul and Swipe: 1.75x (Era server value, no client column).<br>- Lacerate and Primal Bite: "high amount of threat", 2.0x unconfirmed (carried).<br>- Righteous Fury: 25780 effect 0 is +90% Holy and effect 4 +33% all. Both are aura 10, multiplied (cmangos' ApplyPercentModFloatVar agrees; an emulator, not Blizzard).<br>- Holy Shield: +20% ("causes 20% additional threat").<br>- Iron Creed: +10/20/30% Holy Strike threat. Trait curve 110346 is 10/20/30; the damage-taken curve 110345 is 2/4/6% (Wowhead shows the uncurved 15%).<br>- Seal of Fury, Judgement of Fury and Hammer of the Righteous: no threat text, so plain damage times Righteous Fury.<br>- Hand of Reckoning 1219206 ("Righteous Fury increased to 80%"): no SkillLineAbility row, so not learnable. | client rows, Wowhead 25780, 20928, 1311034, 1219206 | Aligned (two carried OPEN) | - |
| A4 | Crit tables: every hit's DefenseType is in `$R/hybrids/crit-notes.md`, handed to the crit lane.<br>- No hit in the group uses the wrong table today. Hammer of Wrath is the group's one ranged-table hit (DefenseType 3).<br>- For the split: Elemental Devastation's buff 30165 is aura 52 (melee only).<br>- Leader of the Pack 24932 and Moonkin Aura 24907 are aura 290 (every table), as are Conviction, Thundering Strikes, Sharpened Claws and Nature's Majesty. | SpellCategories, SpellEffect | Reported | - |
| A5 | Downranking:<br>- Enhancement bare 50 and 60 is mana-bound (Earth Shock behind a 50% mana gate). Earth Shock ranks 4 to 6 at band 60 measured 225.5 / 226.2 / 227.1 against 227.3 at rank 7: no gain.<br>- Protection bare 40 and 50 is mana-bound. Holy Shield's damage per mana is flat across ranks (0.75 / 0.73 / 0.78), Seal of Fury's low ranks lose, and the seal-before-Holy-Shield order measured -5.7% (189.0 against 200.5 threat a second at 40).<br>- Retribution is out of mana for 1.4 s at bare 50 only.<br>No candidates for the ranks lane. | spec-breakdown, 3000 iterations | No candidates | - |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | **A Shred cast at the moment its swing procced Clearcasting was free and kept the charge.**<br>- What the client says: Clearcasting 16870 is one charge.<br>- What was wrong: the consume guard skipped any cast while the aura still had its full duration. A cast completing at the proc's own moment was skipped too.<br>- What changed: `core.ProcSource` (the casters lane's ff4bae666) records the cast that procced; only that cast keeps the charge.<br>- Measured: no change for cat, bear or Enhancement at identical seeds (the coincidence does not arise in these rotations). Restoration druid -0.2 to -0.8%, Balance -0.1 to -0.3% (table below). | `druid/talents.go` applyOmenOfClarity. Tests first: `TestAShredAtTheProcsMomentSpendsClearcasting` failed on the old guard and passes now; `TestTheShredThatProccedClearcastingKeepsIt` passes. | CONFIRMED-BUG (client charge count, engine test) | E d985fbb84 (with E d3b5b796b, the cherry-picked ProcSource) |
| F2 | Rip is re-applied with 5 points when it has under 6 s left. A 2 s threshold measured +0.1% (609.8 against 609.0, seed 1), and waiting for it to fall off measured -0.8%: no change. | spec-breakdown, 3000 iterations, raid 60 and bare 40 and 50 | Aligned (measured) | - |
| F3 | Checked fresh:<br>- Cat Form AP: 3025 effect 0, 12 + 2 a level.<br>- Shifting Power 1322605: 55% base mana, 40 energy, 8 s cooldown in the layered SpellCooldowns, 1 s global.<br>- Ravage from 32; Prowl.<br>- Rip ticks crit (bleed rule).<br>- Every talent in every published build has engine code.<br>- Night Elf and Tauren racials (Quickness 1% dodge, Endurance 5% health).<br>Display only: Prowl's aura reads 9913 at every band (Era id, cosmetic). | client rows, talent scan | Aligned | - |

## Druid: Feral (bear)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| G1 | Checked fresh:<br>- Bear and Dire Bear Form: AP 30 + 3 a level and 120 + 3 a level; health 20 + 18 and 600 + 32 a level (1178, 9635, `bear_form.go` reads the same per-level rule; Wowhead renders 9634 at its MaxLevel 70: 210 AP, 1560 health).<br>- Lacerate 1328616 / 1235826 / 1235827 by band, 10% weapon per existing stack.<br>- Primal Bite 1238069 / 1238073.<br>- Maul and Swipe ranks by band; Faerie Fire (Feral) by band.<br>- Furor's display id is fixed in X2.<br>- Threat: A2, A3. | client rows, breakdowns | Aligned | - |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Checked fresh:<br>- Seal of Righteousness proc: the dummy 1786 + 47 a level is read on `scaleLevel`, which matches each rank's MaxLevel.<br>- Judgement of Righteousness 0.429 coefficient.<br>- Seal of Command proc 20424: 70% weapon, 0.29 coefficient (about 0.2 of spell power after the 70%), 1 s proc cooldown.<br>- Judgement of Command per level.<br>- Holy Strike: effect 121 plus effect 31 (the flat amount is multiplied by the 50%, as Spell::EffectWeaponDmg does).<br>- Hammer of Wrath, Exorcism and Consecration per level.<br>- Improved Judgement: 22.9 Judgements a fight at bare 20.<br>- Vengeance 20049 (trait curve 1/2/3%; display id fixed in X2).<br>- Human Ruthless and Undead Touch of the Grave (1260198 cannot crit: Attributes_2 0x20000000). | client rows, breakdowns | Aligned | - |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | **Holy Shield's damage can crit in the engine; in 70707 it is a damage shield.**<br>- 20925 / 20927 / 20928 effect 1 is aura 43 (SPELL_AURA_DAMAGE_SHIELD, as Thorns), where Era's Holy Shield triggered a damage spell.<br>- Classic damage shields never crit.<br>- The engine rolls `OutcomeMagicCrit` (`holy_shield.go`), so with about 27% crit at 60 that is roughly +13% on Holy Shield's threat, about 2 to 3% of Protection's threat.<br>- Kept open, carried from round 7: whether Blizzard's script crits it is server behaviour.<br>- What would settle it: a Forever log with Holy Shield events at 30+ (the crit flag on its DAMAGE_SHIELD rows). | client 20928 SpellEffect; `holy_shield.go:86` | OPEN (leans engine-wrong) | - |
| T2 | Checked fresh:<br>- Consecration: 12 + 27 a second for the first four targets (companion 1280349).<br>- Seal of Fury proc flat 35 and Judgement of Fury per level.<br>- Holy Shield 30% block, 4 charges, plus Shieldsworn's shield strike.<br>- Righteous Fury and Iron Creed (A3).<br>- Redoubt's display id fixed in X2.<br>- Reckoning is not in the build.<br>- The Holy Shield-first mana order (A5).<br>- Dwarf and Undead racials. | client rows, breakdowns | Aligned | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N1 | **Windfury Weapon and Rockbiter Weapon gave one attack-power number per rank; the client's grows with level.**<br>- What the client says: each rank is the aura 99 of its buff or enchant spell, growing EffectRealPointsPerLevel to the rank's MaxLevel.<br>- Windfury 8233 (rank 1): 46 at 30, +7.2 a level to 104 at 38. 8236: 119 + 12.8 to 48. 10484: 249 + 8.3 to 58. 16361: 333 + 12.5.<br>- Wowhead Forever renders each rank at its MaxLevel: 104, 221, 315, 433, the per-level rule's numbers.<br>- Rockbiter 10400 / 15567 / 15568 / 15569 / 16311 / 16312 / 16313: 29 + 4.1, 58 + 3.5, 88 + 5, 129 + 8.1, 211 + 18, 393 + 16.1, 554 + 16.5 a level.<br>- What was wrong: a level-30 shaman got Windfury rank 1's 104 (client: 46), and a level-20 one Rockbiter rank 3's 118 (client: 108). Rockbiter rank 4 was 138 at any level (client: 129 to 194).<br>- What changed: both read the client per-level rule (`clientdamage.Effect`). Bands 40 to 60 fall exactly on each Windfury rank's own level, so they do not move. Rockbiter still loses to Windfury at 30 and 40 (61.7 against 67.5, and 115.3 against 123.7), so the kit's imbue choice stands. | `shaman/windfury_weapon.go`, `rockbiter_weapon.go`. Tests first: `TestWindfuryWeaponAttackPowerGrowsWithLevel` and `TestRockbiterWeaponAttackPowerGrowsWithLevel` (build failure, then pass). | CONFIRMED-BUG (client rows, Wowhead) | E 60205a769 |
| N2 | **Enhancement dropped its own Windfury Totem, which gave nobody anything.**<br>- What the client says: Windfury Weapon on the main hand "disables any benefit you personally receive from Windfury Totem" (16362), and the engine models it.<br>- Solo (the bare preset) there is no group. The raid preset already carries a Windfury Totem for the group.<br>- What was wrong: the totem cost its mana and a global and held the Air slot Grace of Air needs.<br>- What changed: the rotation keeps Grace of Air down from 42 (rank 2 at 56; rank 3 25359 is the held Ahn'Qiraj book), and drops no Air totem before it.<br>- Measured: +0.5% at bare 40, +3.6% at 50, +4.4% at 60 and +4.0% on the raid preset (table below).<br>- For the wow-player review: a raid Enhancement whose group has no other shaman would drop Windfury for the warriors. That is group value the published single-character number cannot see, and the preset already assumes another shaman's Windfury. | `data/curated/apl/shaman-enhancement.json` (and the fork's copy via `make apl-sync`). Test first: `TestEnhancementDropsGraceOfAirNotItsOwnWindfuryTotem` (failed on the old list). | CONFIRMED-BUG (client text, engine) | S cefa42f10, E 58c993def |
| N3 | **Maelstrom Weapon: an instant Lightning Bolt cast at the moment a swing gave the fifth stack kept all five.**<br>- What was wrong: every stack resets the aura's duration, and the guard compared remaining time to the full duration.<br>- What changed: the same `core.ProcSource` guard as F1.<br>- Measured: no change at identical seeds (the APL never meets the case). | `shaman/talents.go` applyMaelstromWeapon. Test first: `TestALightningBoltAtTheStacksMomentSpendsMaelstromWeapon` (failed, then passes). | CONFIRMED-BUG (engine test) | E d985fbb84 |
| N4 | Checked fresh:<br>- Earth Shock: the learnable ids are 8042 to 10414. 408681 to 408690 and 1220744 to 1220751 have no SkillLineAbility row (a tank variant).<br>- Stormstrike, Searing Totem and Lightning Bolt per level.<br>- Elemental Weapons: 7/13/20% Rockbiter, 13/27/40% Windfury, on the right class masks.<br>- Windfury 20%, 1.5 s proc cooldown (439431).<br>- Orc Blood Fury (10% AP and SP, 15 s, 2 min). Dwarf.<br>- Flurry's display id fixed in X2. Stoneskin rank 6 at every band (round 8 N2, cosmetic). | client rows, breakdowns | Aligned | - |

## Integrator follow-ups (ranks lane results)

Measured on int/r9 (282a87fed plus lane/r9-hybrids d985fbb84). Neither is an engine bug.

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Divine Protection and Divine Shield rank 1 raise Protection's threat a second.**<br>- The client gives the ranks different lengths: Divine Protection 498 lasts 6 s and 5573 8 s (DurationIndex 32 and 31); Divine Shield 642 lasts 10 s and 1020 12 s (DurationIndex 1 and 29). The engine's table matches.<br>- Divine Protection pacifies (aura 25: "you cannot attack", so the engine stops the auto swing). Divine Shield is aura 79 -50% damage dealt on every school ("reduces all damage you deal by 50%").<br>- So a top rank costs two more seconds of no swings or half damage, and the shorter rank 1 loses less threat. The price is two fewer seconds of immunity, which the threat metric does not count.<br>- No on-cast threat, heal threat, seal or Holy Shield path is involved.<br>- Pacify blocks only ability-flagged spells (SPELL_ATTR0_ABILITY 0x10). Holy Strike, Hammer of the Righteous, Judgement, Swift Judgement and Holy Shield all read Attributes_0 0x50000 (no ability flag), so the engine is right to let them through and only stop the auto swing.<br>- These are panic buttons below 20 to 25% health; their rank is a survival choice, not a mana one. Keep the top rank. | SpellMisc, SpellDuration, SpellEffect for 498, 5573, 642, 1020; `paladin/defensive_cooldowns.go` | Aligned (correct behaviour; not a downrank candidate) | - |
| R2 | **Consecration rank 1 beats the top rank (Retribution 50, Protection 40 to 60).** The lower rank is more mana-efficient, and both rotations ration Consecration by mana:<br>- Every rank's damage row (1280345 to 1280349) carries the same 0.042 coefficient on its first-targets part: rank 1 is 2 + 4 a second, rank 5 is 12 + 27.<br>- Every rank lasts 8 s (DurationIndex 31) with 1 s ticks and an 8 s category cooldown (8000).<br>- Mana costs 135 / 235 / 320 / 435 / 565. Rank 1 is level 20, so the only Classic coefficient penalty (spells learned below 20) does not reach it, and 1.12 has no downrank penalty.<br>- The engine matches all of it (`consecration.go`). No threat modifier exists beyond Righteous Fury's Holy +90%.<br>- Both rotations ration it by mana: Retribution casts Consecration only above 65% mana (5.5 casts a fight at bare 50 against a possible 22), and Protection only above 50% mana and after Holy Shield's reserve. A cheaper rank keeps the paladin over the line more often, so the fight is mana-bound through the gate even though out-of-mana seconds are near zero.<br>- On one target each cast deals less, but casts more often, the same per-mana trade the owner's downrank ruling allows. A ranks-lane choice. | client rows above; APL conditions in `paladin-retribution.json` and `paladin-protection.json` | Aligned (genuine efficiency; ranks lane may adopt it) | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | The crit tables (A4) and Elemental Devastation's display id (30165 at every rank; 29177 and 29178 have no SpellName row) go to the crit lane, which owns `applyElementalDevastation`. | crit-notes.md | Routed | - |
| X2 | **Talent buffs carried Era per-rank ids the 70707 client no longer has.** The breakdown showed "spell (20059)".<br>- Why: in 70707 a multi-rank talent is one spell with a trait curve, so it has one buff.<br>- Vengeance 20059 is now 20050, Redoubt 20134 is 20128, Flurry 16277 to 16280 are 16257, Furor 17058 to 17061 are 17056.<br>- Display only. | SpellName, SpellEffect triggers, TraitDefinitionEffectPoints. Tests first: `TestTalentBuffsCarryTheClientsIds`, `TestFlurryBuffCarriesTheClientsId`, `TestFurorCarriesTheClientsId` (build failures, then pass). | CONFIRMED-BUG (display) | E 261a43a91 |
| X3 | The shaman's own Windfury Totem aura (10611) and Prowl's 9913 are Era ids too. Cosmetic, not fixed: Windfury Totem is no longer cast. | breakdowns | Note | - |

## Before/after

Before: pin 40021374a on origin/main 55a1212f7. After: `lane/r9-hybrids` d985fbb84 with `wt/r9-hybrids`. 10,000 iterations. Tanks are in threat a second, Restoration in effective HPS.

Enhancement (N1, N2), horde (alliance run at raid 60; no change here depends on race):

| band | % seed 1 / 50001 | absolute, seed 1 |
|---|---|---|
| bare 20 | -1.2 / -1.2 | 33.8 to 33.4 |
| bare 30 | -1.9 / -1.9 | 69.0 to 67.7 |
| bare 40 | +0.5 / +0.5 | 123.9 to 124.5 |
| bare 50 | +3.7 / +3.6 | 160.3 to 166.2 |
| bare 60 | +4.4 / +4.5 | 217.3 to 226.9 |
| raid 60, horde | +4.0 / +4.0 | 594.0 to 617.7 |
| raid 60, alliance | +4.0 / +4.0 | 594.7 to 618.3 |

The guard fix (F1, N3), measured alone, from engine 261a43a91 to d985fbb84:
- Feral cat, bear and Enhancement: 0.00% at bare 20, 40 and 60 and raid 60, both seeds.
- Restoration druid (healers lane), effective HPS: bare 20 -0.8 / -0.5 (39.4 to 39.1); bare 40 -0.3 / -0.3; bare 60 -0.2 / -0.2; raid 60 -0.6 / -0.5 (777.3 to 772.9).
- Balance (casters lane): bare 40 -0.3 / -0.3 (156.1 to 155.6); bare 60 -0.1 / -0.1; raid 60 -0.0 / -0.0 (579.5 to 579.4).

Unchanged at raid 60, seed 1, before and after:
- Retribution 658.8.
- Protection 1018.2 threat a second.
- Feral cat 608.9.
- Bear 856.3 threat a second.

## For other lanes

- **Crit lane:** `$R/hybrids/crit-notes.md`:
  - Elemental Devastation is melee only (aura 52).
  - Leader of the Pack and Moonkin Aura are aura 290 and raise every table.
  - Hammer of Wrath is DefenseType 3 (ranged).
  - ED's display id is 30165.
- **Integrator:**
  - DPS outside my group moves only through d985fbb84: Restoration druid and Balance (numbers above).
  - The cherry-picked d3b5b796b is the casters lane's ff4bae666, so the merge should resolve to one copy.
  - `TestRotationLadder` goldens were not regenerated; the Enhancement rungs move at 20 to 60.
- **Ranker:** none. `sim/leveling/kit.go`'s Rockbiter-to-Windfury switch at 30 still holds on the corrected numbers.

## Still open

1. T1, Holy Shield crit: a level-30+ Forever log of Holy Shield; the client now states a damage shield.
2. A2, Faerie Fire's form threat: an in-game threat read; combat logs carry no threat.
3. Carried: Lacerate and Primal Bite's 2.0 threat, Righteous Fury's multiplication, the Ahn'Qiraj book ranks at 60, Holy Strike's coefficient under the 50%, the round-8 item procs (P1, T1, T2, X1 there).
