# Sim-accuracy round 29: walks-dps lane

Site `wt/r29-walksdps` from int/r29 75d6264a5. Fork: **no engine change** (no `lane/r29-walksdps` branch; every number runs on the pinned engine 276ecc361 through `$R/base-engine`). Client 1.60.1.70790. Ship mode: touched tests only, no regens. Every sim went through `$R/sim.sh`, one at a time.

**Result: 13 confirmed bugs, all fixed, all `old`; one site code change that makes the fixes expressible.** Every one of the 13 leveling walks spent points on talents that do nothing in the fight (range, armor, pushback, a kill, a stun, a pack spell, a spell the rotation never casts) or were tolls taken earlier than their row needed, while a live talent waited. Five specs also had a live-against-live reorder that measured a gain at every level it touches, and those ship with the same order (arcane, fire and Enhancement hit, Marksmanship's Careful Aim, Elemental's Alacrity, Enhancement's Mental Dexterity). Bands that move are listed under "Re-rank needed".

Method (round 27 orders.md, round 28 casters.md and hybrids.md): every walk printed 10 to 60 with each point's client tooltip, triaged against the engine's reader and the curated APL; candidates built as `ladder-first.json` reservations and compared with `spec-breakdown -preset bare -band <20..50> -level N -talents <walk> -bis-dir $R/bis-frozen` at 3,000 iterations (seed 1, Horde), every touched level; then the chosen order at 10,000 iterations, seeds 1 and 50001, Horde at every selected level (checkpoints, first and last touched level, and every level whose 3,000-iteration gain was under 2%) and Alliance as well at bands 20, 30, 40 and 50. Tanks (Protection warrior, Feral bear) are judged on threat on the tank fight (`-tank`), with chance of death reported. Outputs: `$R/walksdps/out/` (`*-3k.txt`, `f2-<spec>.tsv`), walks in `$R/walksdps/w/`.

Logs, last looked 03:10 UTC 7 Nov: `GET https://api.foreversixty.gg/v1/reports/recent` still lists only `z2opt2x2gs5u` (20 Oct, no kills). A leveling order is not settled by raid logs.

## Verdict per spec

| spec | verdict | levels moved | bands that move |
|---|---|---|---|
| warrior-fury | 1 confirmed, fixed | 35 to 40, 52 to 55 | 40 |
| warrior-protection (threat) | 1 confirmed, fixed | 17 to 45 | 20, 30, 40 |
| mage-arcane | 1 confirmed, fixed | 21 to 51 | 30, 50 |
| mage-fire | 1 confirmed, fixed | 41 to 59 | 50 |
| warlock-affliction | 1 confirmed, fixed | 22 to 59 | 30, 40, 50 |
| hunter-beast-mastery | 1 confirmed, fixed | 20 to 33 | 20, 30 |
| hunter-marksmanship | 1 confirmed, fixed (plus one live reorder) | 15 to 23, 46 to 49 | 20 |
| druid-balance | 1 confirmed, fixed | 42 to 58 | 50 |
| druid-feral-bear (threat) | 1 confirmed, fixed | 16 to 59 | 20, 30, 40, 50 |
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
| WP1 | **Anticipation at 17 to 21 before Improved Revenge and Defiance.** Anticipation is Defense; Defiance (20% more threat in Defensive Stance with a shield, client 12792) opens at 20 (10 points in the rows above: Improved Bloodrage 2, Shield Specialization 5, Improved Revenge 3) and waited until 31 to 33. Now Improved Revenge 17 to 19, Defiance 20 to 22, Master of Defense 23 and 24, Improved Sunder Armor 25 to 27, Improved Thunder Clap 28 to 30, Focused Rage 31 to 33, Anticipation 34 to 38, Concussion Blow 39 (Shield Slam's prerequisite), Shield Slam 40, Bastion 41 to 45, Last Stand 46. Threat +5 to +30% from 17 to 39, +1 to +5% at 40 to 45. **Anticipation's place is the lane's call:** taking it after Bastion (order A) adds 1 to 5% threat at 34 to 44 over this order, but the band 30 boss then kills a level 35 to 39 tank 25 to 87% of the time against 2 to 47% before (this order: 14 to 65%); Anticipation is the build's defensive talent, not a dead point, so it comes once the threat talents that waited on it are in. **Moves bands 20, 30 and 40.** | `out/prot-A-3k.txt`, the chance-of-death runs (orders A, B, C, 3,000 iterations), `f2-warrior-protection.tsv` | CONFIRMED-BUG | old | S 9588348b3 |
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
| DF1 | **Feral Instinct and Thick Hide's points spent ahead of Sharpened Claws, Rend and Tear and Naturalist.** Feral Instinct only adds to Swipe, which the APL casts on two or more targets; the 10-point row needs 4 points beyond Ferocity 5 and Furor 1, and the walk spent Feral Swiftness 2, Feral Instinct 3 and Thick Hide 3. The row is now paid with Thick Hide 3 and Feral Swiftness 1 (16 to 19: threat ties the old points to the last digit at 20, 90.57 against 90.58, and chance of death at 20 falls from 42% to 22%), then Savage Fury 20 and 21, Sharpened Claws 22 and 23, Shredding Attacks, Lacerate 25, Blood Frenzy, Leader of the Pack 30, Predatory Strikes, Natural Instinct, Rend and Tear 36 to 39 and 41, Primal Bite 40, Heart of the Wild 42 to 46, Naturalist 47 to 51 (10% damage, so threat), Natural Reaction 52 to 56, the last Feral Swiftness and Feral Instinct points 57 to 60. Threat +2 to +10% at 21 to 59. Tie at 16 to 20 decided on survival and Forever practice (Warcraft Tavern's Forever feral guide calls Thick Hide the tanking talent; Icy Veins treats Feral Swiftness and Feral Instinct as Cat Form points). Chance of death after 40: order D 30% at 50 against 17% before; order B (Thick Hide last) was 59%, order C (Natural Reaction first) kept 15% but gave only +0.3% threat at 50. **Moves bands 20, 30, 40 and 50.** Card note: Lacerate "from 36 the bleed turns on Rend and Tear". | `out/bear-3k.txt`, orders A to D with chance of death, `f2-druid-feral-bear.tsv` | CONFIRMED-BUG | old | S 9588348b3 |
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

### warrior-protection

threat/s, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 17 | 86.85 | 91.68 | +5.56% / +5.50% | - |
| 20 | 98.33 | 121.22 | +23.28% / +23.21% | +22.70% / +22.78% |
| 25 | 166.05 | 202.29 | +21.82% / +21.77% | - |
| 30 | 191.57 | 238.61 | +24.55% / +24.45% | +24.64% / +24.55% |
| 34 | 305.40 | 351.05 | +14.95% / +14.73% | - |
| 35 | 311.82 | 352.43 | +13.02% / +12.89% | - |
| 38 | 337.69 | 362.18 | +7.25% / +7.00% | - |
| 40 | 515.80 | 533.96 | +3.52% / +3.45% | +3.79% / +3.57% |
| 42 | 548.59 | 553.87 | +0.96% / +0.95% | - |
| 44 | 595.22 | 600.85 | +0.95% / +0.97% | - |
| 45 | 600.77 | 606.14 | +0.89% / +0.89% | - |

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

### druid-feral-bear

threat/s, 10,000 iterations. Absolute values Horde seed 1; change Horde 1 / 50001 (Alliance 1 / 50001 at bands 20, 30, 40, 50).

| level | before (H s1) | after (H s1) | Horde s1 / s50001 | Alliance s1 / s50001 |
|---|---|---|---|---|
| 16 | 63.91 | 64.12 | +0.33% / +0.39% | - |
| 20 | 90.53 | 90.58 | +0.06% / +0.01% | +0.02% / +0.08% |
| 25 | 132.18 | 140.35 | +6.18% / +6.07% | - |
| 30 | 194.38 | 206.87 | +6.43% / +6.47% | +6.59% / +6.70% |
| 35 | 245.15 | 251.24 | +2.48% / +2.56% | - |
| 38 | 247.27 | 266.37 | +7.72% / +7.87% | - |
| 40 | 314.32 | 344.44 | +9.58% / +9.40% | +9.13% / +9.30% |
| 42 | 376.75 | 401.63 | +6.60% / +6.70% | - |
| 44 | 393.10 | 402.76 | +2.46% / +2.56% | - |
| 45 | 400.46 | 402.68 | +0.55% / +0.62% | - |
| 46 | 404.96 | 407.00 | +0.50% / +0.58% | - |
| 47 | 410.58 | 420.32 | +2.37% / +2.41% | - |
| 50 | 470.51 | 503.17 | +6.94% / +7.04% | +6.73% / +6.54% |
| 55 | 493.99 | 541.49 | +9.62% / +9.87% | - |
| 59 | 589.16 | 599.82 | +1.81% / +1.81% | - |

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

## Files to regenerate

- `sim/request/testdata/ladder/` goldens (`TestRotationLadder`, rungs 10, 20, 30, 38, 40, 50, 60) where a walk moves: warrior-fury 38, 40; warrior-protection 20, 30, 38, 40; mage-arcane 30, 38, 50; mage-fire 50; warlock-affliction 30, 38, 40, 50; hunter-beast-mastery 20, 30; hunter-marksmanship 20; druid-balance 50; druid-feral-bear 20, 30, 38, 40, 50; paladin-retribution 30; priest-shadow 50; shaman-elemental 20, 30; shaman-enhancement 20, 38, 40, 50. Until then `web/src/lib/guides/leveling-builds.test.ts` fails on them.
- `addon-builds.json` (the 13 leveling orders), `addon-data.json` (four card notes), then `Data.lua`. `data/tests/test_card_note_levels.py::test_every_allowlist_line_still_matches_a_note` fails on four mage-fire allowlist lines until addon-data is regenerated (the curated Fire notes changed, so no addon-data card line matches them yet; the allowlisted sentences themselves are unchanged).
- `dps-per-point` / stat weights for every band that moves (Re-rank needed).
- Coach starters (`persona_starter.py --all`) after addon-builds.
- **Engine fork APL copies (hand sync, no engine worktree here):** `ui/mage/apls/forever_arcane.apl.json` (Missile Barrage "the second at 22"), `ui/mage/apls/forever_fire.apl.json` (two notes, "at 54"), `ui/feral_druid/apls/forever_feral_bear.apl.json` or its bear copy (Lacerate "from 36"), `ui/elemental_shaman/apls/forever_elemental.apl.json` (Magma Totem "from 34"). `sim/request/apl/<spec>.apl.json` updated by hand to match.
- Hand copies done, not regens: `sim/leveling/ladder-first.json` = `data/curated/ladder-first.json`, `sim/leveling/ladder-order.json` = `data/curated/ladder-order.json`.

## Re-rank needed

One batched dispatch for the band builds that change:

- warrior-fury band 40
- warrior-protection bands 20, 30, 40
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
