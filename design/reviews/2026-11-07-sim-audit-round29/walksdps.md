# Sim-accuracy round 29: walks-dps lane

Site `wt/r29-walksdps` from int/r29 75d6264a5. Fork: **no engine change** (no `lane/r29-walksdps` branch; every number runs on the pinned engine 276ecc361 through `$R/base-engine`). Client 1.60.1.70790. Ship mode: touched tests only, no regens. Every sim went through `$R/sim.sh`, one at a time.

**Result: 13 confirmed bugs, all fixed, all `old`; one site code change that makes the fixes expressible.** Every one of the 13 leveling walks spent points on talents that do nothing in the fight (range, armor, pushback, a kill, a stun, a pack spell, a spell the rotation never casts) or were tolls taken earlier than their row needed, while a live talent waited. Five specs also had a live-against-live reorder that measured a gain at every level it touches, and those ship with the same order (arcane, fire and Enhancement hit, Marksmanship's Careful Aim, Elemental's Alacrity, Enhancement's Mental Dexterity). Bands that move are listed under "Re-rank needed".

Method (round 27 orders.md, round 28 casters.md and hybrids.md): every walk printed 10 to 60 with each point's client tooltip, triaged against the engine's reader and the curated APL; candidates built as `ladder-first.json` reservations and compared with `spec-breakdown -preset bare -band <20..50> -level N -talents <walk> -bis-dir $R/bis-frozen` at 3,000 iterations (seed 1, Horde), every touched level; then the chosen order at 10,000 iterations, seeds 1 and 50001, Horde at every selected level (checkpoints, first and last touched level, and every level whose 3,000-iteration gain was under 2%) and Alliance as well at bands 20, 30, 40 and 50. Tanks (Protection warrior, Feral bear) are judged on threat on the tank fight (`-tank`), with chance of death reported. Outputs: `$R/walksdps/out/` (`*-3k.txt`, `f2-<spec>.tsv`), walks in `$R/walksdps/w/`.

Logs, last looked 03:10 UTC 7 Nov: `GET https://api.foreversixty.gg/v1/reports/recent` still lists only `z2opt2x2gs5u` (20 Oct, no kills). A leveling order is not settled by raid logs.

## Tank rule (coordinator ruling, after the first report)

Leveling tanks quest solo and run dungeons with weak healers, so: **a threat talent goes ahead of a survival talent only where the tank fight's chance of death rises by less than 5 points absolute at every level the swap touches; otherwise the survival talent stays until a level where the swap passes.** The first versions of both tank orders failed it (Protection warrior 2% to 14% at 35, bear 17% to 30% at 50) and were replaced. Both orders below pass at every level, 10,000 iterations, seeds 1 and 50001, both factions (tables in Before/after).

## Verdict per spec

| spec | verdict | levels moved | bands that move |
|---|---|---|---|
| warrior-fury | 1 confirmed, fixed | 35 to 40, 52 to 55 | 40 |
| warrior-protection (threat) | 1 confirmed, fixed (revised under the tank rule) | 17 to 38 | 20, 30 |
| mage-arcane | 1 confirmed, fixed | 21 to 51 | 30, 50 |
| mage-fire | 1 confirmed, fixed | 41 to 59 | 50 |
| warlock-affliction | 1 confirmed, fixed | 22 to 59 | 30, 40, 50 |
| hunter-beast-mastery | 1 confirmed, fixed | 20 to 33 | 20, 30 |
| hunter-marksmanship | 1 confirmed, fixed (plus one live reorder) | 15 to 23, 46 to 49 | 20 |
| druid-balance | 1 confirmed, fixed | 42 to 58 | 50 |
| druid-feral-bear (threat) | 1 confirmed, fixed (revised under the tank rule) | 16 to 59 | 20, 30, 40, 50 |
| paladin-retribution | 1 confirmed, fixed | 24 to 33 | 30 |
| priest-shadow | 1 confirmed, fixed | 31 to 35, 48 to 59 | 50 |
| shaman-elemental | 1 confirmed, fixed (plus two live reorders) | 15 to 21, 30 to 33, 41 to 49 | 20, 30 |
| shaman-enhancement | 1 confirmed, fixed (plus two live reorders) | 15 to 21, 33 to 57 | 20, 40, 50 |

## Warrior: Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| WF1 | **Improved Berserker Rage's second point at 35 and Improved Whirlwind at 36 and 37 before Flurry.** The APL casts Berserker Rage only from 36 (`data/curated/apl/warrior-fury.json`, "From 36, on cooldown under 80 rage"), so the first point at 34 is the toll Flurry's row needs and the second at 35 does nothing. Flurry (30% attack speed after a crit, client 12319) waited until 38 to 41. Now Flurry 35 to 37, Improved Whirlwind 38 and 39 (Whirlwind trains at 36), Bloodthirst 40, Improved Berserker Rage 2 at 41. A sweep of the 36 and 37 points (Improved Whirlwind first, Berserker Rage first, Flurry first) put Flurry first at every level. **Moves band 40** (Flurry 3 in place of Improved Berserker Rage 2). | `out/fury-A-3k.txt`, walks D to G (3,000 iterations), `f2-warrior-fury.tsv` | CONFIRMED-BUG | old | S 9588348b3 |
| WF2 | **Arms: Improved Charge's second point at 49 before Improved Tactical Mastery, Anger Management and Deep Wounds.** Fury never casts Charge in a fight; one point of Improved Charge plus Improved Heroic Strike 3, Improved Rend 3 and Improved Tactical Mastery 3 pay the 10-point row, and the second point is only the toll Impale's row needs. Now Improved Tactical Mastery 49 to 51, Anger Management 52, Deep Wounds 53 to 55, Improved Charge 2 at 56. 49 to 51 tie to the last digit (Improved Tactical Mastery is as dead as Improved Charge on one target: no stance swap after 36). Improved Bloodrage stays last: placed before Impale it lost 1 to 2% at 57 to 59. | walk H, `f2-warrior-fury.tsv` | CONFIRMED-BUG (part of WF1's row) | old | S 9588348b3 |
| WF3 | Enrage at 23 and 24 (needs the target to hit you) and Improved Rend, Improved Tactical Mastery and Improved Charge's first point are tolls with no live alternative in the build's rows (15 points of Fury rows 1 to 3 hold exactly Cruelty, Unbridled Wrath, Furious Precision and Enrage 2; Deep Wounds needs Improved Rend 3, Anger Management needs Improved Tactical Mastery 3). | client `talents/warrior.json` prerequisites | Aligned (tolls) | - | - |

## Warrior: Protection (threat)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| WP1 | **Anticipation at 17 to 21 before Improved Revenge and Defiance.** Anticipation is Defense; Defiance (20% more threat in Defensive Stance with a shield, client 12792) opens at 20 (Improved Bloodrage 2, Shield Specialization 5 and Improved Revenge 3 pay its row) and waited until 31 to 33. Now Improved Revenge 17 to 19, Defiance 20 to 22, Master of Defense 23 and 24, Anticipation 25 to 27, Last Stand 28, Anticipation 29 and 30, Improved Sunder Armor 31 to 33, Improved Thunder Clap 34 to 36, Focused Rage 37 and 38; Concussion Blow 39, Shield Slam 40, Focused Rage 41 and Bastion 42 to 46 are unchanged. Threat +17 to +30% at 20 to 30, +1 to +11% at 31 to 38, both seeds and factions; chance of death at most 1.8 points higher (29). Under the tank rule (above) Anticipation and Last Stand must be in by 28 and 29, where the band 20 boss starts to kill a tank: Anticipation after Bastion took a level 35 tank from 2% to 25%, after Focused Rage 2% to 14%, and without Last Stand at 29 it was 2.6% to 8.5%; all three fail. **Moves bands 20 and 30** (band 40 no longer moves). | `out/prot-A-3k.txt`, `out/prot2-3k.txt` (orders B, E to K), `out/tk-warrior-protection.tsv` | CONFIRMED-BUG | old | S 84542481c |
| WP2 | The band 30 tank fight's chance of death climbs with the level inside the band even on the old order (0% at 30, 47% at 39): the same encounter-tuning question as round 28's P4, for walks-heal's item 3. | chance-of-death runs | OPEN (walks-heal item 3) | - | - |

## Mage: Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| MA1 | **Arcane Resilience at 21 and Arcane Geometry at 24 before Missile Barrage, Arcane Impact and Arcane Meditation.** The engine reads neither (fork `sim/mage/talents.go:103-108`: armor from Intellect and range). Arcane Power needs exactly 30 points in the rows above at 40, so both are tolls, but not at 21 and 24: Missile Barrage now 21 and 22, Arcane Impact 23 to 25, Arcane Meditation 26 to 28, Arcane Resilience 29, Presence of Mind 30 (it trains at 30 and the rotation casts it), Arcane Mind 31 to 35, Arcane Instability 36 to 38, Arcane Geometry 39, Arcane Power 40. +0.5 to +5.9% at 21 to 38. | `out/arcane-3k.txt`, `f2-mage-arcane.tsv` | CONFIRMED-BUG | old | S 9588348b3 |
| MA2 | **Precision after Arcane Power.** From 41 the walk took Incineration and Improved Frostbolt before Precision (6% spell hit, client 29438) at 50 to 52. Precision first at 41 to 43: +2 to +6% at 41 to 51. A live-against-live reorder with a gain at every level it touches. **With MA1, moves bands 30 and 50.** | `out/arcane-3k.txt`, `f2-mage-arcane.tsv` | CONFIRMED-BUG (order) | old | S 9588348b3 |
| MA3 | Precision even earlier (10 to 13, or 21 to 23 after Arcane Blast) measured +2 to +7% at most levels, but every point of it before 40 costs Presence of Mind at 30 and Arcane Power at 40, which the rotation casts; the talent adoption rule forbids dropping them. The hit is a build question (fewer Arcane points before 40), not an order. | `out/arcane-prec-3k.txt` | OPEN (build lane) | - | - |

## Mage: Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| MF1 | **Burning Soul, Improved Channeling and Accelerant's last point at 43 to 48 before Precision at 58 to 60.** Burning Soul and Improved Channeling are pushback and threat (fork `sim/mage/talents.go` reads only the threat mod), and Accelerant shortens Fireball, which the Scorch filler does not cast from 30 until Arcane Blast. Arcane Subtlety only reduces resistance the target does not have (and the engine models only its threat half). Precision (6% spell hit) waited for the last three levels, Arcane Blast for 57. Now Precision 41 to 43, Wand Specialization 44 and 45, Improved Channeling 46 to 48 (the Arcane row toll), Arcane Concentration 49 to 53, Arcane Blast 54, Missile Barrage 55 and 56, Accelerant 5 at 57, Burning Soul 58 and 59, Arcane Subtlety 60. +2.3 to +19.5% at 41 to 59. Wand Specialization moved from ladder-order.json into this list (the round 24 wand row is kept, after Precision). **Moves band 50.** Two card notes name the level (Scorch filler "until Arcane Blast at 54", Fireball filler "the leveling build takes at 54"). | `out/fire-3k.txt`, `f2-mage-fire.tsv` | CONFIRMED-BUG | old | S 9588348b3 |

## Warlock: Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| WA1 | **Grim Reach, Soul Harvest and Improved Drains at 22 to 28 before Amplify Curse, Pandemic and Malevolence.** Grim Reach is range; Soul Harvest needs a Drain Soul kill; Improved Drains boosts Drain Life and Drain Soul, which the rotation never channels (it fills with Shadow Bolt, then Wrack). Pandemic's row opens at 13 points (Improved Life Tap, Suppression and Improved Corruption already pay it). Now Amplify Curse 22, Pandemic 23 to 25, Malevolence 26 and 27, Grim Reach 28 and 29 (Siphon Life's toll), Siphon Life 30, Soul Siphon 31 to 33 (dead until Wrack, but the toll Shadow Mastery needs, and live at 40), one Soul Harvest point 34, Shadow Mastery 35 to 39, Wrack 40, then Soul Harvest and Improved Drains. A first try with Shadow Mastery before Soul Siphon lost 0.5 to 4.5% at 40 to 42 (Soul Siphon feeds Wrack), so Soul Siphon pays the row. +0.5 to +6% at 22 to 43. | `out/aff-3k.txt`, `f2-warlock-affliction.tsv` | CONFIRMED-BUG | old | S 9588348b3 |
| WA2 | **Demonology and Destruction after 44: Improved Health Funnel, Demonic Aegis and Improved Voidwalker before Bane.** None of the three does anything in a fight (no Health Funnel cast, armor, a Voidwalker the Imp-out fight never summons); Bane shortens Immolate. Now Unholy Power 45 to 49 (the Imp is out), Bane 50 to 54, then the three toll points. +0.2 to +2.8% at 45 to 59. **With WA1, moves bands 30, 40 and 50.** | `f2-warlock-affliction.tsv` | CONFIRMED-BUG | old | S 9588348b3 |

## Hunter: Beast Mastery

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| HB1 | **Endurance Training's fourth point at 20 and Bestial Swiftness at 21 before Unleashed Fury and Ferocity.** Pet health and pet run speed do nothing in the fight; Endurance Training 3 already pays the 10-point row. Now Unleashed Fury 20 to 24, Ferocity 25 to 29, Bestial Discipline 30 and 31, Bestial Swiftness 32 and Intimidation 33 (the prerequisite chain to Bestial Wrath; Intimidation is never cast), Endurance Training 4 at 34, Frenzy 35 to 39. +0.7 to +5.8% at 20 to 33. **Moves bands 20 and 30.** | `out/bm-3k.txt`, `f2-hunter-beast-mastery.tsv` | CONFIRMED-BUG | old | S 9588348b3 |

## Hunter: Marksmanship

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| HM1 | **Endurance Training at 46 to 48 before Focused Fire.** Endurance Training only pays Unleashed Fury's row; Focused Fire (2% all damage, Beast Mastery row 2) needs only Deadly Aspects 5. Focused Fire now 46 and 47, Endurance Training 48 to 50. +1 to +2% at 46 to 49. (Improved Tracking before Deadly Aspects lost 0.4 to 1.2% at 41 to 45; not taken.) | `out/mm-3k.txt`, `f2-hunter-marksmanship.tsv` | CONFIRMED-BUG | old | S 9588348b3 |
| HM2 | **Careful Aim before Efficiency** (both row 2): +1.4 to +5.5% at 15 to 23. Live against live, a gain at every level it touches. **Moves band 20.** | 3,000-iteration run (orders A/C), `f2-hunter-marksmanship.tsv` | CONFIRMED-BUG (order) | old | S 9588348b3 |

## Druid: Balance

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| DB1 | **Nature's Focus at 42 to 46 before Heart of the Wild and Naturalist.** Nature's Focus is pushback protection on a target that never hits you (engine `applyNaturesFocus`, pushback only); Naturalist is 2% damage a point (client 17069) and waited until 52 to 56. Now Heart of the Wild 42 to 46, Naturalist 47 to 51, Reflection 52 to 54, Nature's Focus 55 to 59, Subtlety 60. +0.3 to +10% at 42 to 58. This needed a deeper off-tree reservation (Naturalist sits in Restoration's second row), see X1. **Moves band 50.** | `out/bal-3k.txt`, `f2-druid-balance.tsv` | CONFIRMED-BUG | old | S 9588348b3 |

## Druid: Feral bear (threat)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| DF1 | **Feral Instinct and Feral Swiftness spent ahead of Sharpened Claws and Rend and Tear.** Feral Instinct only adds to Swipe, which the APL casts on two or more targets; the 10-point row needs 4 points beyond Ferocity 5 and Furor 1, and the walk spent Feral Swiftness 2 and Feral Instinct 2 there, Feral Instinct 3 at 22 and Thick Hide at 23, 24 and 26. The row is now paid with Thick Hide 3 and Feral Swiftness 1 (16 to 19: threat ties, chance of death falls 18 to 22 points at 19 to 21), then Savage Fury 20 and 21, Sharpened Claws 22 and 23, Shredding Attacks, Lacerate 25, Blood Frenzy 28 and 29, Leader of the Pack 30, Predatory Strikes 31 to 33, Natural Instinct 34 and 35, Rend and Tear 36 to 38, Feral Swiftness 2 at 39, Primal Bite 40, Rend and Tear 41 and 42, Natural Reaction 43 to 47, Heart of the Wild 48 to 52, Naturalist 53 to 57, Feral Instinct 58 to 60. Threat +2.4 to +7.9% at 22 to 39, gains at 40 to 49 and 53 to 59; 16 to 21 and 50 to 52 tie on threat with a lower chance of death; no level's chance of death rises more than 1 point. Under the tank rule, Naturalist before Natural Reaction (first version) took death at 50 from 17% to 30% and Naturalist before Heart of the Wild's stamina 17% to 23%, and Feral Swiftness 2 after 40 took 40 from 81% to 89%; none is taken. Tie at 16 to 21 decided on survival and Forever practice (Warcraft Tavern's Forever feral guide calls Thick Hide the tanking talent; Icy Veins treats Feral Swiftness and Feral Instinct as Cat Form points). **Moves bands 20, 30, 40 and 50.** Card note: Lacerate "from 36 the bleed turns on Rend and Tear". | `out/bear-3k.txt`, `out/bear2-3k.txt` (orders D to G), `out/tk-druid-feral-bear.tsv` | CONFIRMED-BUG | old | S 84542481c |
| DF2 | The Maul note's 30-rage gate at 40 to 53 was tuned on the old order (Natural Reaction at 46 to 50). Not re-tuned here. | APL note | OPEN (next pass) | - | - |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| PR1 | **Vindication 2 and 3 at 24 and 25 before Seal of Command, Sacred Arbiter and Two-Handed Weapon Specialization.** Vindication lowers the target's attack power; nothing in the fight hits you. One point pays the 15-point row with Seal of Command. Now Vindication 1 at 23, Seal of Command 24, Sacred Arbiter 25, Two-Handed Weapon Specialization 26 to 28, Vindication 2 at 29, Vengeance 30 to 32, Vindication 3 at 33, Repentance 34 (Champion of the Light's toll). +1.6 to +6.2% at 25 to 33; 24 ties exactly (Seal of Command alone adds nothing that level). Precision with Divine Intellect after 40 lost 1 to 12% (not taken). **Moves band 30.** | `out/ret-3k.txt`, `f2-paladin-retribution.tsv` | CONFIRMED-BUG | old | S 9588348b3 |

## Priest: Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| PS1 | **Blackout at 31 and 32 and Vampiric Embrace at 33 before Shadow Weaving.** Blackout is a stun proc; Vampiric Embrace costs a global every 30 seconds for healing only, so the level it is taken loses about 6% (71.1 to 67.1 at 33). Both pay Darkness's row and Shadowform's prerequisite. Now Shadow Weaving 31 to 33, Blackout 34 and 35, Vampiric Embrace 36, Darkness 37 to 41 (unchanged). +3 to +14% at 31 to 35. | `out/shadow-3k.txt`, `f2-priest-shadow.tsv` | CONFIRMED-BUG | old | S 9588348b3 |
| PS2 | **Discipline: Silent Resolve at 48 before Mental Strength.** Silent Resolve is threat; Wand Specialization 2 and Twin Disciplines 5 already pay Mental Strength's row. Mental Strength 48 to 52, Mental Agility, Inner Focus and Meditation one level each earlier, Silent Resolve at 60. +0.5 to +6% at 48 to 59. **Moves band 50.** | `f2-priest-shadow.tsv` | CONFIRMED-BUG | old | S 9588348b3 |
| PS3 | Spirit Tap at 12 to 16 does nothing on one target (no kills), but round 15 decided it on Forever guides and real leveling (every kill), like round 28's Feral Charge. Kept. | ladder-first reason (round 15) | Aligned (real play) | - | - |

## Shaman: Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| SE1 | **Elemental Warding and both points of Trial by Fire at 27 to 30 before Call of Thunder and Lightning Overload.** Elemental Warding is damage taken, Trial by Fire is Fire Nova and Magma Totem, both two-target lines (APL expected_idle). Three of the four are tolls for the 20-point row; the fourth (Trial by Fire 2) now waits until 34. Call of Thunder 30, Lightning Overload 31 to 33. +1.7 to +3.2% at 30 to 33. Card note: Magma Totem "40 from 34 with both points of Trial by Fire". | `out/ele-3k.txt`, `f2-shaman-elemental.tsv` | CONFIRMED-BUG | old | S 9588348b3 |
| SE2 | **Elemental Alacrity before Concussion** (+1.6 to +7.7% at 15 to 20, 21 ties) and **Tidal Focus before Thundering Strikes** after 40 (+0.3 to +1.5% at 41 to 49). Live against live, gains at every touched level. **Moves bands 20 and 30.** | `f2-shaman-elemental.tsv` | CONFIRMED-BUG (order) | old | S 9588348b3 |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| SH1 | **Improved Ghost Wolf at 33 and 34 before Mental Quickness and Maelstrom Weapon.** Improved Ghost Wolf is travel. The APL casts Lightning Bolt only at five Maelstrom stacks with all five points, so Maelstrom Weapon 1 to 4 do nothing either (36 to 38 tie to the last digit); the order only matters for when the fifth point lands. Now Mental Quickness 33 and 34, Maelstrom Weapon 35 to 39, Rage of the Farseer 40, Improved Ghost Wolf 41 and 42: the instant bolt arrives at 39 instead of 42, +24.6% at 39, +10% at 40 and 41. **Moves band 40.** | `out/enh-3k.txt`, `f2-shaman-enhancement.tsv` | CONFIRMED-BUG | old | S 9588348b3 |
| SH2 | **Mental Dexterity before Ancestral Knowledge** (+2.5 to +7.4% at 15 to 21) and **Tidal Focus (5% hit for your swings) after one point of Convection** (44 ties, +1.5 to +6.5% at 45 to 57; Tidal Focus before that Convection point lost 0.5% at 43). Live against live. **Moves bands 20 and 50.** | `f2-shaman-enhancement.tsv` | CONFIRMED-BUG (order) | old | S 9588348b3 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **A reservation could not name another tree's deeper talent** (`reserveOffTree` took only first-row talents), so Balance's Naturalist, Fire's Arcane Concentration, Fury's Deep Wounds, Marksmanship's Focused Fire and Shadow's Mental Strength could not be ordered. A deeper off-tree rank is now reserved once the ranks the reservation holds in that tree pay its row and prerequisite (`offTreeRankIsPaid`, `sim/leveling/talents.go`). No walk moves from this alone (all 28 specs printed 10 to 60 before and after: identical). Red first: `TestReservationSpendsADeeperOffTreeRankOnceItsRowIsPaid`. | test; walk print | CONFIRMED (tool gap) | old | S f54cac4a8 |
| X2 | Shamans' Maelstrom Weapon APL line waits for all five points (36 to 38 tie): with fewer points a five-stack bolt is still 20 to 80% faster and cheaper. Whether a partial-rank bolt line pays is an APL question for the shaman lane. | APL line; ties | OPEN (APL) | - | - |
| X3 | Fresh re-audit: limited to the 13 walks 10 to 60 (every point against its client row, the engine's reader and the APL) and the build choices the walks touch. The full class re-audit was not repeated. Rank rows on the moved levels were not re-judged (see Measure at refresh). | - | not done (scope) | - | - |

## Before/after

Bare band gear at the level (`$R/bis-frozen`), the guide build's walk at the level, engine 276ecc361. Before = the int/r29 walk, after = this lane's walk. Levels not listed either tie by construction (identical builds) or were confirmed at 3,000 iterations only (gain above 2%, `out/*-3k.txt`). Published races: Horde/Alliance per spec's faction default.

### warrior-fury

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 35 | 79.64 | 81.96 | +2.91% / +3.01% | +3.26% / +3.28% |
| 36 | 86.45 | 89.49 | +3.51% / +3.54% | - |
| 37 | 89.57 | 92.36 | +3.11% / +3.09% | - |
| 38 | 92.71 | 94.79 | +2.24% / +2.06% | +2.46% / +2.48% |
| 39 | 95.99 | 98.07 | +2.17% / +2.09% | - |
| 40 | 143.06 | 147.21 | +2.90% / +2.89% | +3.08% / +2.89% |
| 49 | 170.73 | 170.73 | +0.00% / +0.00% | - |
| 50 | 222.37 | 222.37 | +0.00% / +0.00% | +0.00% / +0.00% |
| 51 | 222.94 | 222.94 | +0.00% / +0.00% | - |
| 52 | 227.12 | 229.81 | +1.18% / +1.29% | - |
| 55 | 244.16 | 250.42 | +2.56% / +2.57% | - |

### mage-arcane

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 21 | 45.57 | 46.25 | +1.49% / +1.56% | - |
| 22 | 46.64 | 46.95 | +0.66% / +0.83% | - |
| 23 | 48.42 | 48.70 | +0.57% / +0.58% | - |
| 25 | 56.66 | 57.41 | +1.31% / +1.30% | - |
| 30 | 90.01 | 90.48 | +0.52% / +0.22% | +0.41% / +0.44% |
| 31 | 90.36 | 91.49 | +1.25% / +1.36% | - |
| 35 | 107.64 | 108.77 | +1.05% / +1.18% | - |
| 38 | 117.36 | 119.37 | +1.71% / +1.72% | - |
| 41 | 172.41 | 176.16 | +2.18% / +1.93% | - |
| 45 | 182.50 | 192.87 | +5.68% / +5.66% | - |
| 50 | 247.22 | 254.38 | +2.89% / +3.05% | +2.76% / +2.62% |
| 51 | 254.30 | 258.16 | +1.52% / +1.61% | - |

### mage-fire

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 41 | 161.26 | 165.27 | +2.49% / +2.34% | - |
| 45 | 165.46 | 178.22 | +7.71% / +7.85% | - |
| 50 | 200.97 | 219.63 | +9.29% / +9.56% | +9.60% / +9.56% |
| 55 | 223.58 | 251.02 | +12.27% / +12.27% | - |
| 59 | 274.96 | 282.20 | +2.63% / +2.58% | - |

### warlock-affliction

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 22 | 62.51 | 62.90 | +0.63% / +0.63% | - |
| 23 | 63.31 | 63.89 | +0.90% / +0.91% | - |
| 25 | 68.05 | 69.05 | +1.47% / +1.48% | - |
| 30 | 100.71 | 102.92 | +2.20% / +2.20% | +2.03% / +2.01% |
| 33 | 104.53 | 105.65 | +1.07% / +1.08% | - |
| 34 | 106.52 | 107.10 | +0.55% / +0.55% | - |
| 35 | 107.93 | 109.43 | +1.39% / +1.39% | - |
| 40 | 188.00 | 199.43 | +6.08% / +6.08% | +6.12% / +6.12% |
| 41 | 194.00 | 202.69 | +4.48% / +4.48% | - |
| 42 | 199.99 | 205.86 | +2.94% / +2.94% | - |
| 43 | 206.39 | 209.37 | +1.44% / +1.44% | - |
| 45 | 218.92 | 219.43 | +0.23% / +0.23% | - |
| 46 | 226.57 | 227.64 | +0.47% / +0.47% | - |
| 48 | 242.62 | 243.76 | +0.47% / +0.47% | - |
| 50 | 280.38 | 283.15 | +0.99% / +0.96% | +0.96% / +0.97% |
| 55 | 303.56 | 311.84 | +2.73% / +2.80% | - |
| 57 | 320.25 | 324.90 | +1.45% / +1.54% | - |
| 58 | 334.92 | 337.77 | +0.85% / +0.97% | - |
| 59 | 339.68 | 340.82 | +0.33% / +0.37% | - |

### hunter-beast-mastery

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 20 | 87.75 | 88.35 | +0.68% / +0.68% | +0.69% / +0.69% |
| 21 | 89.32 | 90.55 | +1.38% / +1.38% | - |
| 25 | 104.68 | 106.08 | +1.33% / +1.34% | - |
| 30 | 137.20 | 143.72 | +4.75% / +4.75% | +4.79% / +4.72% |
| 33 | 151.27 | 152.99 | +1.14% / +1.18% | - |

### hunter-marksmanship

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 15 | 57.98 | 58.80 | +1.42% / +1.41% | - |
| 20 | 90.18 | 92.80 | +2.90% / +2.92% | +2.77% / +2.82% |
| 23 | 100.82 | 101.31 | +0.48% / +0.47% | - |
| 46 | 260.52 | 263.13 | +1.00% / +1.00% | - |
| 49 | 274.85 | 277.57 | +0.99% / +0.99% | - |

### druid-balance

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 42 | 176.23 | 176.56 | +0.19% / +0.31% | - |
| 43 | 177.04 | 177.94 | +0.51% / +0.44% | - |
| 44 | 178.16 | 179.44 | +0.72% / +0.80% | - |
| 45 | 178.58 | 180.30 | +0.96% / +0.92% | - |
| 50 | 217.23 | 235.11 | +8.23% / +8.33% | +8.25% / +8.41% |
| 55 | 246.12 | 261.33 | +6.18% / +6.18% | - |
| 58 | 275.13 | 279.30 | +1.51% / +1.44% | - |

### paladin-retribution

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 24 | 52.59 | 52.59 | +0.00% / +0.00% | - |
| 25 | 52.86 | 53.81 | +1.80% / +1.79% | - |
| 29 | 61.59 | 62.58 | +1.60% / +1.60% | - |
| 30 | 83.02 | 85.59 | +3.09% / +3.08% | +3.30% / +3.30% |
| 33 | 97.41 | 100.05 | +2.71% / +2.70% | - |

### priest-shadow

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 31 | 69.23 | 71.27 | +2.94% / +2.96% | - |
| 35 | 71.26 | 77.09 | +8.19% / +8.13% | - |
| 48 | 195.69 | 197.10 | +0.72% / +0.72% | - |
| 49 | 198.32 | 199.42 | +0.55% / +0.58% | - |
| 50 | 243.30 | 245.22 | +0.79% / +0.72% | +0.97% / +0.94% |
| 53 | 258.03 | 259.33 | +0.50% / +0.52% | - |
| 55 | 264.04 | 266.13 | +0.79% / +0.78% | - |
| 59 | 318.80 | 337.33 | +5.81% / +5.79% | - |

### shaman-elemental

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 15 | 15.16 | 15.42 | +1.70% / +1.57% | - |
| 20 | 30.38 | 30.82 | +1.45% / +1.47% | +1.38% / +1.28% |
| 21 | 31.41 | 31.45 | +0.13% / +0.17% | - |
| 30 | 61.02 | 62.04 | +1.67% / +1.68% | +1.63% / +1.63% |
| 33 | 65.19 | 67.16 | +3.02% / +2.94% | - |
| 41 | 110.14 | 110.51 | +0.34% / +0.32% | - |
| 42 | 111.69 | 112.38 | +0.62% / +0.47% | - |
| 43 | 113.38 | 114.43 | +0.93% / +0.84% | - |
| 45 | 112.53 | 114.14 | +1.42% / +1.23% | - |
| 48 | 121.81 | 122.46 | +0.53% / +0.44% | - |
| 49 | 124.52 | 124.77 | +0.20% / +0.16% | - |

### shaman-enhancement

DPS, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 15 | 26.07 | 26.72 | +2.50% / +2.48% | - |
| 16 | 26.67 | 28.00 | +4.99% / +5.12% | - |
| 17 | 30.02 | 32.03 | +6.71% / +6.67% | - |
| 18 | 31.11 | 33.37 | +7.25% / +7.47% | - |
| 19 | 32.33 | 34.55 | +6.87% / +6.81% | - |
| 20 | 36.23 | 38.15 | +5.30% / +5.38% | +5.93% / +6.00% |
| 21 | 37.69 | 38.60 | +2.40% / +2.36% | - |
| 33 | 103.13 | 103.91 | +0.75% / +0.76% | - |
| 34 | 103.66 | 105.23 | +1.51% / +1.51% | - |
| 35 | 104.79 | 105.58 | +0.75% / +0.75% | +0.87% / +0.87% |
| 36 | 107.78 | 107.78 | +0.00% / +0.00% | - |
| 37 | 108.36 | 108.36 | +0.00% / +0.00% | - |
| 38 | 110.38 | 110.38 | +0.00% / +0.00% | +0.00% / +0.00% |
| 39 | 111.02 | 138.28 | +24.56% / +25.07% | - |
| 40 | 157.19 | 173.01 | +10.06% / +10.37% | +10.33% / +10.27% |
| 41 | 157.57 | 173.82 | +10.31% / +10.41% | - |
| 44 | 187.01 | 187.16 | +0.08% / +0.15% | - |
| 45 | 188.08 | 190.90 | +1.50% / +1.60% | - |
| 50 | 237.80 | 251.73 | +5.86% / +5.64% | +5.49% / +5.68% |
| 55 | 255.73 | 265.77 | +3.93% / +3.98% | - |
| 57 | 274.25 | 277.90 | +1.33% / +1.41% | - |

### warrior-protection (threat and chance of death, tank fight)

10,000 iterations. Threat change and death before -> after, Horde seed 1 / 50001 and Alliance seed 1 / 50001. Absolute threat is Horde seed 1.

| level | threat before -> after (H s1) | threat change H s1 / s50001, A s1 / s50001 | death H s1 | death H s50001 | death A s1 | death A s50001 | max death rise |
|---|---|---|---|---|---|---|---|
| 17 | 86.85 -> 91.68 | +5.56 / +5.50, +5.48 / +5.39% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 18 | 91.82 -> 101.38 | +10.41 / +10.33, +10.34 / +10.36% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 19 | 94.79 -> 109.87 | +15.91 / +15.88, +15.40 / +15.42% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 20 | 98.33 -> 121.22 | +23.28 / +23.21, +22.70 / +22.78% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 21 | 99.23 -> 128.70 | +29.70 / +29.84, +29.12 / +29.32% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 22 | 131.90 -> 171.90 | +30.33 / +30.20, +29.87 / +29.82% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 23 | 139.45 -> 174.60 | +25.21 / +25.31, +25.08 / +25.00% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 24 | 164.23 -> 198.28 | +20.73 / +20.72, +20.78 / +20.67% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 25 | 166.05 -> 199.28 | +20.01 / +19.94, +20.04 / +20.09% | 0.0 -> 0.1% | 0.0 -> 0.0% | 0.0 -> 0.1% | 0.0 -> 0.1% | +0.1 |
| 26 | 168.80 -> 201.42 | +19.32 / +19.17, +19.25 / +19.28% | 0.0 -> 0.2% | 0.0 -> 0.3% | 0.0 -> 0.3% | 0.1 -> 0.3% | +0.3 |
| 27 | 171.29 -> 203.05 | +18.54 / +18.38, +18.42 / +18.36% | 0.2 -> 0.7% | 0.3 -> 0.9% | 0.3 -> 1.0% | 0.2 -> 1.0% | +0.8 |
| 28 | 176.58 -> 209.87 | +18.85 / +18.85, +18.90 / +18.95% | 0.5 -> 1.2% | 0.5 -> 1.6% | 0.4 -> 1.2% | 0.4 -> 1.5% | +1.1 |
| 29 | 179.68 -> 211.73 | +17.84 / +17.89, +17.83 / +17.88% | 2.7 -> 4.2% | 2.9 -> 4.1% | 2.7 -> 4.5% | 3.0 -> 4.3% | +1.8 |
| 30 | 191.57 -> 224.09 | +16.98 / +16.77, +16.98 / +16.83% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 31 | 205.54 -> 228.64 | +11.24 / +11.28, +11.26 / +11.31% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 32 | 225.32 -> 242.45 | +7.60 / +7.58, +7.71 / +7.46% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | 0.0 -> 0.0% | +0.0 |
| 33 | 240.07 -> 248.22 | +3.39 / +3.35, +3.52 / +3.58% | 0.1 -> 0.2% | 0.2 -> 0.1% | 0.1 -> 0.2% | 0.2 -> 0.3% | +0.1 |
| 34 | 305.40 -> 312.10 | +2.19 / +2.00, +2.18 / +2.02% | 0.4 -> 0.5% | 0.4 -> 0.4% | 0.7 -> 0.7% | 0.7 -> 0.6% | +0.1 |
| 35 | 311.82 -> 315.29 | +1.11 / +0.95, +1.06 / +0.99% | 2.2 -> 2.5% | 2.0 -> 1.9% | 2.9 -> 2.9% | 3.0 -> 3.0% | +0.3 |
| 37 | 320.64 -> 331.08 | +3.26 / +3.24, +3.35 / +3.44% | 17.0 -> 17.0% | 16.5 -> 17.0% | 20.1 -> 20.6% | 19.9 -> 20.5% | +0.6 |
| 38 | 337.69 -> 349.67 | +3.55 / +3.38, +3.54 / +3.52% | 30.2 -> 29.6% | 30.0 -> 30.9% | 34.3 -> 35.2% | 35.1 -> 34.8% | +0.9 |

### druid-feral-bear (threat and chance of death, tank fight)

10,000 iterations. Threat change and death before -> after, Horde seed 1 / 50001 and Alliance seed 1 / 50001. Absolute threat is Horde seed 1.

| level | threat before -> after (H s1) | threat change H s1 / s50001, A s1 / s50001 | death H s1 | death H s50001 | death A s1 | death A s50001 | max death rise |
|---|---|---|---|---|---|---|---|
| 16 | 63.91 -> 64.12 | +0.33 / +0.39, +0.36 / +0.24% | 12.6 -> 12.1% | 12.4 -> 12.5% | 11.4 -> 11.5% | 12.0 -> 11.8% | +0.1 |
| 17 | 66.85 -> 67.36 | +0.76 / +0.84, +0.51 / +0.68% | 3.9 -> 3.7% | 4.4 -> 3.3% | 3.1 -> 2.4% | 3.5 -> 2.5% | -0.2 |
| 18 | 85.26 -> 85.58 | +0.38 / +0.42, +0.39 / +0.28% | 29.6 -> 21.3% | 30.6 -> 20.0% | 24.9 -> 17.7% | 26.5 -> 18.1% | -7.2 |
| 19 | 85.94 -> 85.92 | -0.02 / -0.09, -0.05 / -0.06% | 63.5 -> 42.2% | 63.6 -> 43.7% | 59.2 -> 37.3% | 59.2 -> 38.1% | -19.9 |
| 20 | 90.53 -> 90.58 | +0.06 / +0.01, +0.02 / +0.08% | 41.4 -> 21.7% | 42.4 -> 22.2% | 34.5 -> 16.9% | 35.4 -> 17.9% | -17.5 |
| 21 | 94.57 -> 94.63 | +0.06 / +0.07, +0.04 / +0.10% | 74.5 -> 56.7% | 75.2 -> 57.6% | 67.6 -> 49.2% | 68.7 -> 50.2% | -17.6 |
| 22 | 95.12 -> 97.52 | +2.52 / +2.49, +2.40 / +2.44% | 92.5 -> 83.6% | 93.2 -> 84.4% | 88.7 -> 77.6% | 89.3 -> 78.8% | -8.8 |
| 23 | 95.64 -> 100.33 | +4.90 / +4.86, +4.68 / +4.76% | 96.5 -> 95.4% | 97.0 -> 96.0% | 94.6 -> 93.0% | 95.0 -> 93.6% | -1.0 |
| 24 | 98.38 -> 103.13 | +4.83 / +4.79, +4.64 / +4.70% | 98.5 -> 99.0% | 98.9 -> 99.1% | 97.6 -> 98.4% | 97.9 -> 98.5% | +0.8 |
| 25 | 132.18 -> 140.35 | +6.18 / +6.07, +6.06 / +6.02% | 99.9 -> 100.0% | 99.9 -> 100.0% | 99.9 -> 99.9% | 99.9 -> 99.9% | +0.1 |
| 26 | 141.39 -> 151.09 | +6.86 / +6.92, +6.94 / +6.91% | 99.9 -> 100.0% | 99.9 -> 100.0% | 99.9 -> 100.0% | 99.9 -> 100.0% | +0.1 |
| 27 | 143.57 -> 152.65 | +6.32 / +6.39, +6.37 / +6.39% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | +0.0 |
| 28 | 145.66 -> 154.30 | +5.93 / +5.93, +6.13 / +6.11% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | +0.0 |
| 29 | 147.27 -> 155.87 | +5.84 / +5.74, +6.16 / +6.05% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | +0.0 |
| 30 | 194.38 -> 206.87 | +6.43 / +6.47, +6.59 / +6.70% | 99.1 -> 99.7% | 99.3 -> 99.8% | 98.6 -> 99.5% | 98.5 -> 99.5% | +1.0 |
| 31 | 200.61 -> 210.88 | +5.12 / +5.16, +5.71 / +5.62% | 99.8 -> 99.9% | 99.8 -> 100.0% | 99.6 -> 99.9% | 99.7 -> 99.8% | +0.3 |
| 32 | 206.72 -> 215.14 | +4.07 / +4.00, +4.82 / +4.88% | 99.9 -> 100.0% | 99.9 -> 100.0% | 99.9 -> 100.0% | 99.8 -> 99.9% | +0.1 |
| 33 | 211.01 -> 219.23 | +3.90 / +3.92, +4.71 / +4.71% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | +0.0 |
| 34 | 240.93 -> 248.29 | +3.05 / +3.06, +3.53 / +3.60% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | +0.0 |
| 35 | 245.15 -> 251.24 | +2.48 / +2.56, +2.81 / +2.90% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | +0.0 |
| 36 | 242.07 -> 252.89 | +4.47 / +4.60, +4.63 / +4.68% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | +0.0 |
| 37 | 243.82 -> 259.55 | +6.45 / +6.50, +6.36 / +6.44% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | +0.0 |
| 38 | 247.27 -> 266.37 | +7.72 / +7.87, +7.43 / +7.53% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | +0.0 |
| 39 | 250.14 -> 266.72 | +6.63 / +6.63, +6.34 / +6.34% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | 100.0 -> 100.0% | +0.0 |
| 40 | 314.32 -> 334.86 | +6.53 / +6.54, +6.52 / +6.52% | 81.2 -> 81.2% | 82.2 -> 82.2% | 79.3 -> 79.3% | 79.6 -> 79.6% | +0.0 |
| 41 | 323.99 -> 344.83 | +6.43 / +6.44, +6.41 / +6.41% | 90.1 -> 90.1% | 90.5 -> 90.5% | 88.7 -> 88.7% | 89.1 -> 89.1% | +0.0 |
| 42 | 376.75 -> 399.01 | +5.91 / +5.92, +5.89 / +5.90% | 80.3 -> 80.3% | 81.3 -> 81.3% | 79.4 -> 79.4% | 79.3 -> 79.3% | +0.0 |
| 43 | 388.40 -> 403.12 | +3.79 / +3.90, +3.93 / +4.13% | 88.8 -> 86.2% | 89.9 -> 86.4% | 88.1 -> 84.8% | 88.0 -> 85.0% | -2.6 |
| 44 | 393.10 -> 401.74 | +2.20 / +2.40, +2.36 / +2.35% | 77.1 -> 68.0% | 77.9 -> 69.2% | 75.4 -> 66.7% | 76.6 -> 68.8% | -7.8 |
| 45 | 400.46 -> 402.69 | +0.56 / +0.63, +0.74 / +0.81% | 85.5 -> 74.7% | 85.9 -> 74.9% | 84.8 -> 73.9% | 84.8 -> 74.0% | -10.8 |
| 46 | 404.96 -> 408.66 | +0.91 / +0.86, +0.91 / +1.06% | 89.5 -> 80.6% | 89.7 -> 80.3% | 88.3 -> 79.0% | 88.6 -> 79.5% | -8.9 |
| 47 | 410.58 -> 414.37 | +0.92 / +1.09, +1.24 / +1.09% | 92.2 -> 84.4% | 92.5 -> 84.6% | 91.1 -> 83.1% | 91.1 -> 83.7% | -7.4 |
| 48 | 415.36 -> 418.21 | +0.69 / +0.76, +0.87 / +0.82% | 94.2 -> 88.4% | 94.0 -> 88.6% | 93.5 -> 88.3% | 93.7 -> 88.7% | -5.0 |
| 49 | 420.42 -> 421.96 | +0.37 / +0.33, +0.37 / +0.37% | 95.2 -> 91.9% | 95.3 -> 92.0% | 94.7 -> 91.6% | 94.7 -> 92.0% | -2.7 |
| 50 | 470.51 -> 471.06 | +0.12 / +0.14, +0.05 / +0.05% | 17.2 -> 11.3% | 16.9 -> 11.7% | 16.8 -> 11.6% | 16.4 -> 11.3% | -5.1 |
| 51 | 474.76 -> 474.84 | +0.02 / +0.06, +0.07 / +0.02% | 24.4 -> 17.8% | 24.8 -> 17.7% | 22.2 -> 16.0% | 22.4 -> 16.2% | -6.2 |
| 52 | 474.04 -> 474.17 | +0.03 / +0.06, +0.10 / +0.15% | 13.1 -> 8.7% | 13.5 -> 9.1% | 12.2 -> 8.3% | 12.7 -> 8.6% | -3.9 |
| 53 | 478.36 -> 487.97 | +2.01 / +1.94, +1.97 / +2.03% | 18.7 -> 15.2% | 18.8 -> 15.6% | 17.7 -> 14.4% | 17.1 -> 14.1% | -3.0 |
| 54 | 489.39 -> 508.48 | +3.90 / +3.90, +3.93 / +3.89% | 24.5 -> 22.5% | 24.9 -> 23.3% | 23.7 -> 21.5% | 23.0 -> 20.9% | -1.6 |
| 55 | 493.99 -> 522.83 | +5.84 / +5.84, +5.85 / +5.85% | 32.0 -> 32.0% | 31.8 -> 31.8% | 29.8 -> 29.8% | 28.9 -> 28.9% | +0.0 |
| 56 | 507.74 -> 536.83 | +5.73 / +5.73, +5.73 / +5.74% | 43.5 -> 43.5% | 42.8 -> 42.8% | 39.8 -> 39.8% | 39.8 -> 39.8% | +0.0 |
| 57 | 522.11 -> 551.47 | +5.62 / +5.62, +5.63 / +5.63% | 52.4 -> 52.4% | 52.2 -> 52.2% | 50.1 -> 50.1% | 48.7 -> 48.7% | +0.0 |
| 58 | 577.44 -> 598.73 | +3.69 / +3.69, +3.69 / +3.69% | 62.7 -> 62.7% | 63.5 -> 63.5% | 58.8 -> 58.8% | 58.9 -> 58.9% | +0.0 |
| 59 | 589.16 -> 599.82 | +1.81 / +1.81, +1.81 / +1.81% | 72.1 -> 72.1% | 71.8 -> 71.8% | 67.5 -> 67.5% | 69.0 -> 69.0% | +0.0 |

## Files to regenerate

- `sim/request/testdata/ladder/` goldens (`TestRotationLadder`, rungs 10, 20, 30, 38, 40, 50, 60) where a walk moves: warrior-fury 38, 40; warrior-protection 20, 30, 38; mage-arcane 30, 38, 50; mage-fire 50; warlock-affliction 30, 38, 40, 50; hunter-beast-mastery 20, 30; hunter-marksmanship 20; druid-balance 50; druid-feral-bear 20, 30, 38, 40, 50; paladin-retribution 30; priest-shadow 50; shaman-elemental 20, 30; shaman-enhancement 20, 38, 40, 50. Until then `web/src/lib/guides/leveling-builds.test.ts` fails on them.
- `addon-builds.json` (the 13 leveling orders), `addon-data.json` (four card notes), then `Data.lua`. `data/tests/test_card_note_levels.py::test_every_allowlist_line_still_matches_a_note` fails on four mage-fire allowlist lines until addon-data is regenerated (the curated Fire notes changed, so no addon-data card line matches them yet; the allowlisted sentences themselves are unchanged).
- `dps-per-point` / stat weights for every band that moves (Re-rank needed).
- Coach starters (`persona_starter.py --all`) after addon-builds.
- **Engine fork APL copies (hand sync, no engine worktree here):** `ui/mage/apls/forever_arcane.apl.json` (Missile Barrage "the second at 22"), `ui/mage/apls/forever_fire.apl.json` (two notes, "at 54"), `ui/feral_druid/apls/forever_feral_bear.apl.json` or its bear copy (Lacerate "from 36"), `ui/elemental_shaman/apls/forever_elemental.apl.json` (Magma Totem "from 34"). `sim/request/apl/<spec>.apl.json` updated by hand to match.
- Hand copies done, not regens: `sim/leveling/ladder-first.json` = `data/curated/ladder-first.json`, `sim/leveling/ladder-order.json` = `data/curated/ladder-order.json`.

## Re-rank needed

One batched dispatch for the band builds that change:

- warrior-fury band 40
- warrior-protection bands 20, 30
- mage-arcane bands 30, 50
- mage-fire band 50
- warlock-affliction bands 30, 40, 50
- hunter-beast-mastery bands 20, 30
- hunter-marksmanship band 20
- druid-balance band 50
- druid-feral-bear bands 20, 30, 40, 50
- paladin-retribution band 30
- priest-shadow band 50
- shaman-elemental bands 20, 30
- shaman-enhancement bands 20, 40, 50

Raid 60 and bare 60 do not move for any spec (every walk ends on the guide build; the raid builds are untouched).

## Measure at refresh

- Every re-ranked band above on its new gear; the before/after here is on frozen gear.
- Rank rows on the moved levels were not re-judged on the new talents (round 28 found three specs' rows moved with their orders). Candidates: Arcane Frostbolt and Arcane Missiles 21 to 30, Affliction Shadow Bolt and Corruption 22 to 39, Retribution Consecration 24 to 33, Shadow Mind Flay and Shadow Word: Pain 31 to 35, Elemental Lightning Bolt 15 to 21 and 30 to 33, Balance Wrath and Starfire 42 to 58, Fire Scorch and Pyroblast 41 to 53.

## For other lanes

- **walks-heal (item 3):** WP2; the band 30 tank fight's chance of death climbs from 0% at 30 to 47% at 39 for a Protection warrior on the old order, and the band 20 to 40 bear fights kill a bare bear 90 to 100% of the time at 22 to 48.
- **Build lane:** MA3 (Precision before 40 on Arcane measures +2 to +7% but needs a build with fewer Arcane points).
- **Shaman lane:** X2 (Maelstrom Weapon line waits for all five points).
- **Bear APL:** DF2 (Maul's 30-rage gate at 40 to 53 was tuned on the old order).
- **Coach (generated; do not hand-edit):** starters that name a talent's level for these 13 specs follow addon-builds at the next `persona_starter.py --all`.
- **Integrator:** sync the four fork APL copies listed above.

## For the guide lane

Done here, `_prose.test.ts` and `_sections.test.ts` pass (145 tests): warrior/protection.md (Levels 1 to 19 and 20 to 60 talent order), warrior/fury.md (Levels 20 to 60), mage/arcane.md (Levels 20 to 60), mage/fire.md (the post-40 paragraph), warlock/affliction.md (Amplify Curse from 22, Soul Harvest from 34, the talent order), hunter/beast-mastery.md and hunter/marksmanship.md (Levels 1 to 19, and Marksmanship's post-40 points), druid/balance.md (post-40 points), druid/feral-bear.md (Levels 1 to 19 and 20 to 60), paladin/retribution.md, priest/shadow.md, shaman/elemental.md (Levels 1 to 19 and the build paragraph), shaman/enhancement.md (Levels 1 to 19 and 20 to 60). Four card notes (Arcane Missiles, Fire's two filler lines, Lacerate, Magma Totem). No wow-player SHIP was run (build mode).

## Guide text not done

- None.

## Still open

- WP2, MA3, DF2, X2, X3. Logs: still only the 20 Oct report (03:10 UTC 7 Nov). Owner pending items untouched.

## Sources

- Client 1.60.1.70790: `talents/{warrior,mage,warlock,hunter,druid,paladin,priest,shaman}.json` (rank texts and prerequisites quoted per row). Fork 276ecc361: `sim/mage/talents.go` (Arcane Geometry, Improved Channeling, Arcane Resilience unread; Arcane Subtlety and Burning Soul threat only), `sim/core/spell_resistances.go` (spell penetration reduces only real resistance), `sim/druid/talents.go` (Nature's Focus pushback only), `sim/shaman/talents.go` (Maelstrom Weapon per-rank stacks). Curated APLs `data/curated/apl/<spec>.json` for what each rotation casts.
- Warcraft Tavern, Forever Feral Druid guide (Thick Hide as the tanking talent): https://www.warcrafttavern.com/forever/guides/feral-druid/ (search summary, 7 Nov).
- Icy Veins, Forever Feral druid builds (Feral Swiftness and Feral Instinct as Cat Form points): https://www.icy-veins.com/wow-forever/feral-druid-melee-dps-and-tank-pve-spec-builds-talents (search summary, 7 Nov).
- Mobalytics, Forever feral druid tank guide: https://mobalytics.gg/wow-forever/classes/feral-druid-tank-guide (search summary, 7 Nov).
