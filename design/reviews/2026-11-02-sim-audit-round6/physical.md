# Sim audit, round 6: physical group (2026-11-02)

This lane covers Arms, Fury, Protection, Assassination, Combat, Subtlety, Beast Mastery, Marksmanship and Survival (melee), plus hunter pets, on client 1.60.1.70707.

- **Baseline:** engine pin 8baffe5fc with the committed simdb.bin. That file is byte-identical to round 5's final local regeneration, so data.yml has already applied round 5's item fixes.
- **Branches:** fork fixes are on `lane/r6-physical` (E) and site fixes on `wt/r6-physical` (S).
- **Measurements:** `sim/cmd/spec-breakdown`, 10000 iterations, seeds 1 and 50001.
- **"After" simdb.bin:** regenerated locally from this branch. It is not committed, because the nightly owns it.

**The owner's question.** The two points in Improved Concussive Shot did nothing; they only filled Marksmanship's first row, and the build search wrongly protected them as an untested talent. They now go to Deadly Aspects: about +4% on the leveling build and about +8% on a new raid build. Efficiency is modelled correctly but gains nothing, because Resourcefulness already cuts melee mana costs by 60% and a Survival hunter never runs out of mana in a 180 s fight.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Post-nightly checks (assigned item 1).** data.yml's d40f4d6a4 (2026-11-01 20:07 UTC) wrote a simdb.bin byte-identical to round 5's final local build (`cmp`). So round 5's four fixes are already live: hotfix item rows, shield block value, Wowhead on-equip stats, and stats 36/37. bis.yml 21d97955f re-ranked at 8baffe5fc, and the band-60 Survival set wears Heartseeker 12783 in the off hand on both factions. | `scratchpad/r5/physical/simdb.r5d.bin` vs `data/builds/1.60.1.70707/simdb.bin`; `bis/hunter-survival.json` | aligned | - |
| X2 | **Half-way stats rounded to even (assigned item 2).** A 1.60 curve stat is `budget × StatPercentEditor / 10000`. Python's `round()` rounds an exact half to the even neighbour, so 22.5 became 22. The client rounds half up. Theramore Gauntlets 271941: 30 × 7500 / 10000 = 22.5 Strength, which Wowhead Forever shows as "+23 Strength". | SimulationCraft `item_database::scaled_stat` uses `util::round` = `floor(x + 0.5)` (`engine/dbc/sc_item_data.cpp`, `engine/util/util.cpp:3191`). Wowhead Forever shows the higher value on all 32 exact halves with a record (tooltip 271941 read 2026-11-01). Test `test_a_curve_stat_exactly_half_way_rounds_up` failed first. | CONFIRMED-BUG (two sources) | S ed00b5d4e |
| X3 | **The "17 items with a stat gap" (assigned item 2), what remains after X2.** None is a physical item, so they are left for the casters lane. | Re-run of the round-5 comparison on the regenerated simdb | aligned (physical) | - |
| X3a | Negative stats: Cowl of Necromancy -5 Stamina, Doomsayer's Robe -10 Stamina. | Era's literal stats agree with the client. Wowhead's feed simply omits negative stats. | aligned | - |
| X3b | Crackling Staff 19102 and Atiesh: spell damage and healing split by school. | Casters lane | OPEN (casters) | - |
| X3c | Rune of Perfection 21565/21566 carries spell penetration as -10 / -6. | Casters lane | OPEN (casters) | - |
| X4 | **Keen Machete 18610's quality (assigned item 2).** Wowhead says common. Both 70532 hotfix pushes say uncommon at item level 10 (pushes 16795826 and 112920, ItemSearchName field 2; 70450's pushes disagreed with each other). We follow the newest push. No published set wears it. | `wago.tools/hotfixes?product=wowf&search=Keen%20Machete` (read 2026-11-01); tooltip 18610 | aligned (the push beats the web) | - |
| X5 | **Block value checked against two sources (assigned item 3).** The 1.60 client states no shield block value: ItemSparse has no column, the hotfix push uses the same schema, and the gametable is empty. Wowhead's `blockamount` and cmangos classic-db's `item_template.block` agree on 190 of the 200 shields both carry. Each of the 10 that differ is a Forever item-level change, or an on-equip block bonus that Wowhead folds in (The Immovable Object). simdb takes the larger of the two and never adds them. **New finding: the 36 shields Wowhead lacks still blocked nothing.** The classic-db extract never read `block`: Skullflame Shield 1168 (40), Crest of Retribution 13375 (38), Forcestone Buckler 7748, Ironbark Shield 20502. 40 shields change on a local regeneration. | `raw/classicdb/item_template.json` re-extracted; tests `test_extract_records_reads_a_shields_block` and `test_classic_db_shields_carry_their_item_template_block` failed first | CONFIRMED-BUG | S 5cfa13092 |
| X6 | **Block value semantics.** In the engine (`unit.go:396`), block value is `max(0, Str/20 - 1) + mult × item`; blocks subtract it (`spell_outcome.go:795, 944`) and Shield Slam adds it (`shield_slam.go:84`). cmangos `Player::GetShieldBlockValue` (`Player.cpp:4884`) is `floor((item + Str/20 - 1) × pct)`. For a warrior `pct` is 1, so the two differ only by the floor: 61.4 against 61 on the published set. For the Protection paladin they differ more: the paladin's Shield Specialization multiplier scales the item part only in the engine, but the strength part too in cmangos. That is for the hybrids lane's +50% decomposition. | cmangos-classic source (read 2026-11-01) | aligned (warrior); OPEN (paladin, hybrids) | - |
| X7 | **Mana regen ignored "N% continues while casting" auras (owner question).** `ManaTick` used per-tick rates cached at reset and on stat changes. An aura that changed only `SpiritRegenRateCasting` never refreshed the cache, so it did nothing: Resourcefulness, Rapid Recuperation, Improved Stormstrike, Spirit Tap. The rates are now read at each tick. | test `TestResourcefulnessRegenContinuesWhileCasting` failed first (0.00 mana a tick, want 16.33); TestEnhancement goldens move (Improved Stormstrike) | CONFIRMED-BUG | E 6fac3a6bc |
| X8 | **Rune of the Guard Captain 19120 simmed as 42 attack power / 84 ranged attack power; it should be 14 / 14 (fresh, rogue and hunter re-audits).** Two causes. First, 1318000 states its +14 with both aura 99 and aura 124, and `statmap.stat_array` also mirrors generic attack power into ranged, so the ranged part counted twice. Second, 1287704 is Forever's terrain rider ("In a Forest or Grassland area, Attack Power increased by an additional $w1"), which the sim counted unconditionally. 35 client spells gate on a terrain this way. The published numbers fight in no such area (the same rule as slaying gear on the type-neutral target), so they are no longer on-equip stats. The rune is trinket 1 on every horde rogue, hunter, Fury and Arms set. The fix also corrects Cleansed Vilebranch Medallion 270275, Graverobber's Shovel 273817, the Enriched and Darkspear Voodoo Seals' mp5, and Jewel of the Guard Captain's spell power. | Wowhead Forever 19120 and 270275: "+14 Attack Power. This effect is tripled in Forest and Grassland areas." Client Spell.csv AuraDescription of 1287704; ItemXItemEffect 97534 / 170007. Tests `test_one_spells_attack_power_and_ranged_attack_power_are_one_generic_line` and `test_area_conditional_equip_spells_are_named_by_their_client_text` failed first. | CONFIRMED-BUG (client text + server tooltip) | S f46050ab8 |

## Warrior: Fury and Arms

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Fresh re-audit: every ability, talent, cooldown, stance, racial and gear proc in the rotations and builds matches the 70707 rows. The full list is under "What the warrior re-audit checked and found aligned" below. The hotfix push changes only SpellMisc flags for these spells. | SpellEffect, SpellAuraOptions, `spellconst/warrior.json`; `scratchpad/r6/physical/audit-warrior/` | aligned | - |
| F2 | The horde sets lose about 0.9% at the next data run: Rune of the Guard Captain (X8). | X8 | CONFIRMED-BUG | S f46050ab8 |
| A1 | Arcanite Champion's Strength proc (16916, +120 for 30 s) is an assumed 1 PPM. The client and classic-db state no rate: ItemEffect 99501 has no SpellAuraOptions or PPM row, and classic-db's `spellppmRate_1` is 0. It is the band-60 main hand on both factions. | `sim/common/item_effects.go:218` | OPEN (needs a log or a stated rate) | - |

## Warrior: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Defensive Stance's +30% and Defiance's +15% threat were added (1.45). They multiply: 1.495 (fresh).** Both are SPELL_AURA_MOD_THREAT. cmangos `Aura::HandleModThreat` multiplies each with ApplyPercentModFloatVar, as TrinityCore does, and upstream wowsims/classic `stances.go` uses 1.3 × 1.15. This is the same evidence as round 5's Righteous Fury ruling. Threat rises 3.0%. | client 7376 effect 2 (aura 10, +30) and 12792 (aura 10, +15); test `TestDefensiveStanceAndDefianceThreatMultiply` failed first; TestP1TankWarrior goldens move | CONFIRMED-BUG (two server implementations + upstream) | E 518c266aa |
| P2 | Block value (X5, X6): no change for the warrior. Evergreen Shield already carries 44. | X5 | aligned | - |
| P3 | Revenge threat is 2.25 × damage + 243. classic-db's spell_threat gives 11601 a flat 315 at multiplier 1. If that is right, Protection threat falls about 15%. | `revenge.go:9-19` (marked unconfirmed); classic-db spell_threat | OPEN (launch threat log) | - |
| P4 | Shield Slam's flat threat is 508 (2 × 254). classic-db gives 250 × 1. Heroic Strike uses its classic-db 145 unscaled, so the engine's own conventions disagree. Thunder Clap is 2.5 × damage against classic-db's flat 180 (small). Sunder's 5% attack power threat is carried open from round 5. | `shield_slam.go:16-22`; `heroic_strike_cleave.go:77`; `thunder_clap.go:92` | OPEN (launch threat log) | - |

## Rogue: Assassination, Combat, Subtlety

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Fresh re-audit of every rank, cooldown, talent value, proc rate and the hit and crit tables against the 70707 rows and the push: aligned. The full list is under "What the rogue re-audit checked and found aligned" below. The push touches no rogue damage, cost or cooldown row. The raid searches on the lane engine find each guide build best (Subtlety's best variant is +0.2%, within error). | `scratchpad/r6/physical/audit-rogue/` | aligned | - |
| R2 | Horde sets: Rune of the Guard Captain (X8), -1.0 to -1.1%. | X8 | CONFIRMED-BUG | S f46050ab8 |
| R3 | Subtlety's Ghostly Strike starves behind Hemorrhage: 1.4 casts a fight. Pooling for it loses 0.15%, so the list is right. | variant at 10000 iterations, both seeds | aligned | - |
| R4 | Venom is still not taken (-7% if it were). Rupture's attack-power share uses TBC's 4/10/18/21/24%, while the client's BonusCoefficientFromAP is 0 and the tooltip only says "increased by your Attack Power". Combat's Slice and Dice is round 2's C2 hold. | `venom.go`, `rupture.go:96`; SpellEffect 687822 | OPEN (carried; each needs a level-60 log) | - |

## Hunter: Survival

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | **Improved Concussive Shot (owner question).** 19407 is a proc trigger (aura 42) that stuns through 19410, attached to Concussive Shot 5116, a snare the rotation never casts. The engine correctly gives it nothing. The talent search's "survival-only" rule accepted only spell-modifier rows, so the adoption rule protected the two points as an "unmodeled guide talent" (the round 5 H2 hold). A proc-only talent is now survival-only when its tooltip names only non-damage spells and no gain (damage, attack, crit, haste, power, a resource, a stat or threat). Across every class this rule change affects only Improved Concussive Shot. | spellconst 19407, 5116; tests `TestSurvivalOnlyAcceptsACrowdControlProcOnANonDamageSpell` and `TestHunterSurvivalOnlyTalentsFromTheActiveBuild` failed first | CONFIRMED-BUG (search tool) | S 34afb4634 |
| S2 | **Builds adopted under the rule (owner question).** The searches ran on the lane engine (horde, 3000 iterations) and were re-measured at 10000 × 2 seeds, both factions, after S5. All three adopted variants keep every unmodeled guide talent, keep the spec's own tree largest and drop nothing the rotation casts. Search reports are in `design/reviews/talent-search/hunter-*.md`. | | adopted | S 34afb4634 |
| S2a | Leveling build `2/305005001/500230231051222151`: 2 Improved Concussive Shot become 2 Deadly Aspects. | bare +4.0%, raid +4.1 / +4.3% | adopted | S 34afb4634 |
| S2b | New raid build `5/305005101/500230231051022131`: Deadly Aspects 5/5 and 1 Rapid Killing replace Resourcefulness 2 and 2 Lightning Reflexes. It needs Judgement of Wisdom's mana; bare it runs dry for 67 of 180 s and loses 3.7%, so it is the raid build only. | raid +7.9 / +8.0% | adopted | S 34afb4634 |
| S2c | Hawk Eye 3 → Deadly Aspects 3 would add another +5.6% bare. Hawk Eye is range utility and a Survival raid keep, the same class as Affliction's Grim Reach. | bare search row 6 | OPEN (owner) | - |
| S3 | **Efficiency (owner question).** It is modelled: Shots, Stings and the melee strikes at -15% (`talents.go:380`). It gains nothing because a Survival hunter never runs dry. Resourcefulness cuts melee and trap costs by 60%, Judgement of Wisdom returns about 5000 mana a fight, and the free globals go to Wing Clip under Windfury. Out of mana is 0 s at 180 s, 300 s and 600 s raid, and at 180 s bare. The fight length (180 s) and the mana model are not the cause, apart from the regen bug X7, which this lane fixed. The guide wrongly said Efficiency was shot-only; corrected. | breakdowns at `-duration 180/300/600`, both presets | aligned (and the guide text fixed) | S 34afb4634 |
| S4 | **Resourcefulness's regen procced on any crit (fresh).** The text is "melee and ranged ability Critical Strikes", and SpellAuraOptions 242512 gives 440529 ProcTypeMask 272 (16 \| 256). White swings and bleed ticks no longer proc it. | test `TestResourcefulnessProcsOnlyFromAbilityCrits` failed first | CONFIRMED-BUG | E b06eb1870 |
| S5 | **Aspect of the Beast was never registered (fresh, hunter re-audit).** It is Forever's trainable melee aspect: 50/70/90/110 melee attack power for 50/70/90/110 mana at levels 30/40/50/60 (13161, 1299445-1299447). Its proc trigger to Quick Strikes 1299448 fires through Deadly Aspects. Falcon is level 60 only, so Survival fought with no aspect before 60. Beast is now registered, Deadly Aspects procs under it, and the list casts it before the pull. | client rows; `trainables/hunter.json`; Wowhead Forever 1299447 ("Melee Attack Power by 110"); test `TestAspectOfTheBeastGivesMeleeAttackPowerAndDeadlyAspectsMelee` failed first; conformance golden gains the four ranks, all "match" | CONFIRMED-BUG (client + trainer + tooltip) | E ddedcd84c, defc01ed9; S 2d4d7ae55, d56f399ee |
| S5a | Gain from S5 by band. | horde, 10000 iterations, seeds 1 / 50001: band 30 +3.6 / +3.6%; 40 +4.1 / +4.0%; 50 +4.8 / +4.8%; 60 bare +2.1 / +2.1%; 60 raid +1.8 / +1.8% | measured | - |
| S6 | Aspect of the Falcon's own row has ProcChance 5 with mask 20 (melee auto and ability). The engine procs Quick Strikes only through Deadly Aspects; wowsims/sod also gates it on the talent. | SpellAuraOptions 239200 | OPEN (a level-60 melee log) | - |

## Hunter: Beast Mastery and Marksmanship

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | **Bestial Discipline's mana half was missing (fresh, hunter re-audit).** 19590 rank 2 carries aura 134 at 50: "allows 50% of your Mana Regeneration to continue while casting". Only the pet focus half was modelled. Bare at band 60, BM runs dry for 108 of 180 s. Aspect of the Viper is not learnable in Forever (415423 inactive), so the rest of that time is real. | test `TestBestialDisciplineLetsHalfTheManaRegenContinueWhileCasting` failed first | CONFIRMED-BUG | E c7cc2bdec |
| B2 | **Beast Mastery raid build adopted.** 2 Mortal Shots replace 1 Efficiency and 1 Rapid Killing: +1.3 / +1.5% raid (10000 iterations, both seeds). Bare it is +0.7%, so the leveling build is unchanged. Marksmanship's winner is +1.0% alliance and +0.9% horde, under the 1% bar, so it is held. | `design/reviews/talent-search/hunter-*-raid.md` | adopted (BM) / HOLD (MM) | S 34afb4634 |
| B3 | Arcane Shot (0.429) and Serpent Sting (0.2 a tick) carry spell coefficients that the rows 691069 and 690655 do not state. Effect under 0.1%. | conformance golden: "0.429 (convention)" | OPEN (single source) | - |
| B4 | In the raid run, Serpent Sting ticks average 143 at 28% crit, about 16% above what the known multipliers explain. The bare run fits. Hawk's proc mask is 320 (ranged auto and ability), while the text and the engine use Auto Shot only. | breakdown `audit-hunter/beast-mastery-raid.txt`; SpellAuraOptions 126992 | OPEN | - |
| B5 | Horde sets: Rune of the Guard Captain (X8). MM horde -2.0%. | X8 | CONFIRMED-BUG | S f46050ab8 |

## Hunter pets

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| PET1 | **Claw ranks 1-3 and 5 used rank 4's 16-22 (fresh, hunter re-audit).** The client gives 4-6, 8-12, 12-16 and 21-29. Pet damage was too high below level 24 and too low at levels 32-39. | SpellEffect 16827-16831 (base points and variance); Wowhead "4 to 6", "12 to 16", "21 to 29"; test `TestPetClawRanksRollTheClientRange` failed first | CONFIRMED-BUG | E 18a6ef22f |
| PET2 | These are carried open from round 5. | | OPEN (carried) | - |
| PET2a | Pet family. | | OPEN (carried) | - |
| PET2b | Pet haste. | | OPEN (carried) | - |
| PET2c | Pet expertise. | | OPEN (carried) | - |

## What the warrior re-audit checked and found aligned

**Fury and Arms abilities:**
- Bloodthirst 23894: 48 + 45% attack power, 6 s.
- Mortal Strike 21553: weapon damage + 160.
- Slam 11605: +87, 1.5 s cast, 18 s cooldown. Improved Slam -0.5 s and -3 s.
- Whirlwind hits with both hands.
- Overpower: +35, 78% crit.
- Execute: 600 + 15 a point of rage.
- Rend: 147 over 21 s, +35% with Improved Rend. It crits, per the owner's ruling.
- Hamstring 45. Heroic Strike rank 8 (rank 9 is an Ahn'Qiraj book).
- Battle Shout: 116.

**Cooldowns:** Death Wish 12328 and Recklessness 1719.

**Talents:**
- Bastion (aura 79, physical only).
- Two-Handed Weapon Specialization, Impale, Cruelty, Furious Precision, Dual Wield Specialization, Enrage, Flurry, Weaponmaster.
- Unbridled Wrath: 60% for 1 rage.
- Improved Bloodrage.

**Racials:** Blood Fury, Berserking 15 s, Ruthless 1%, Tauren Endurance.

**Gear procs:** Molten and Frozen Heart of the Mountain, Teebu's Blazing Longsword.

## What the rogue re-audit checked and found aligned

**Ranks:**
- Backstab 11281 (150% + 140, normalized).
- Sinister Strike 11294 (+68).
- Mutilate 1241584: 2 combo points, +20% on poisoned targets.
- Hemorrhage 16511: 100% (145% with a dagger), +15% Rupture damage taken, 35 energy.
- Ghostly Strike 14278, Ambush 11269.
- Eviscerate 11300 and Backstab 25300 sit behind `IncludeAQ`.

**Cooldowns:** Cold Blood, Adrenaline Rush, Blade Flurry, Premeditation, Preparation.

**Talents:**
- Lethality, Aggression, Opportunity, Improved Eviscerate, Serrated Blades, Vigor, Initiative.
- Flawless Execution, Thousand Cuts, Quietus.
- Murder: 2%, Humanoid and Giant only.

**Procs:**
- Ruthlessness: 20% a rank.
- Relentless Strikes: 20% per combo point, 25 energy.
- Seal Fate: 100% with a 500 ms ICD (AO 126847).
- Puncturing Wounds (AO 247392).
- Hack and Slash: 5% with a 200 ms ICD (AO 125982).

**Hit and crit:**
- Weapon Expertise counts as 2% expertise.
- Poisons miss with melee hit. Deadly Poison IV is 18 a tick, stacking to 5.
- Deadly Poison and Rupture ticks crit.

**Gear:**
- Stormshroud's Lightning at 10% and Revitalize at 4%.
- Devilsaur 2-piece: +2% hit. Shadowcraft 3-piece: +30 attack power.
- Heartseeker and Scarlet Kris against Wowhead Forever.

**Racials:** troll Berserking, 1.5 casts a fight.

## Before/after (band 60, published sets and builds, 10000 iterations; seed 1 / seed 50001)

**Before:** pin 8baffe5fc with the committed simdb.bin.

**After:** `lane/r6-physical` at defc01ed9, plus `wt/r6-physical`, plus a simdb.bin regenerated locally from this branch (X2, X5 and X8).

The guides' new builds (S2, B2) are not in this table. The published sets still carry the old talents until bis.yml re-ranks. Their gains are given in S2 and B2.

| Spec | alliance % | horde % | before → after (seed 1, alliance / horde) |
|---|---|---|---|
| Arms | 0.00 / 0.00 | -0.86 / -0.84 | 728.3 → 728.3 / 746.8 → 740.4 |
| Fury | 0.00 / 0.00 | -0.92 / -0.93 | 772.2 → 772.2 / 782.2 → 775.0 |
| Protection (threat) | +3.06 / +3.06 | +3.05 / +3.05 | 1036.0 → 1067.7 / 1024.6 → 1055.9 |
| Assassination | 0.00 / 0.00 | -0.96 / -0.96 | 621.9 → 621.9 / 625.0 → 619.0 |
| Combat | 0.00 / 0.00 | -0.97 / -0.97 | 622.7 → 622.7 / 628.9 → 622.8 |
| Subtlety | 0.00 / 0.00 | -1.13 / -1.13 | 558.4 → 558.4 / 564.7 → 558.3 |
| Beast Mastery | +1.56 / +1.67 | -0.12 / -0.25 | 659.0 → 669.3 / 674.2 → 673.4 |
| Marksmanship | 0.00 / 0.00 | -2.00 / -2.01 | 661.6 → 661.6 / 679.7 → 666.1 |
| Survival | +1.78 / +1.78 | +0.87 / +0.86 | 771.5 → 785.2 / 781.8 → 788.6 |

What drives each move:
- **Horde losses:** Rune of the Guard Captain (X8).
- **Beast Mastery's gain:** Bestial Discipline (B1).
- **Survival's gain:** Aspect of the Beast (S5).
- **Protection:** Defiance (P1).

Outside the group, the specs that moved beyond 0.05% are all horde sets wearing the rune: Feral -0.83 / -0.81, Retribution -0.73 / -0.71, Enhancement -0.80 / -0.82. Every other spec measures 0.00.

The mana regen fix (X7) moves no published band-60 number, but it does move Enhancement's fork goldens. The rounding (X2), the classic-db shield block (X5) and the terrain riders on caster items (X8) can move picks at the next bis.yml.

## Still open

- **Owner calls:**
  - Hawk Eye for Survival (S2c), which is in the same range-utility class as Grim Reach.
  - Whether a raid build may rely on Judgement of Wisdom's mana (S2b).
- **Need a launch log:**
  - Revenge, Shield Slam and Thunder Clap threat values (P3, P4) and Sunder's attack power threat.
  - Arcanite Champion's proc rate (A1).
  - Rupture's attack power coefficient, and Venom (R4).
  - Falcon's untalented 5% (S6).
  - Hawk's proc mask and Serpent Sting's raid tick (B4).
  - The round 5 carries: rage-crit factor, Rogue Armor Energize, Raptor Strike and the dual-wield penalty, Fire traps scaling with attack power.
- **Hybrids lane:** the paladin block-value formula (X6).
- **Casters lane:**
  - School-split and negative-penetration item rows (X3).
  - Terrain riders on caster items, now dropped from simdb (X8). The planner's `items.json` still shows them.
- **Integrator:** TestRotationLadder's Survival and hunter-pet rows move (S5, PET1). They are left for the integrator to regenerate at the pin.
