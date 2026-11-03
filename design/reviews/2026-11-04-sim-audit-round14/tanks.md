# Sim audit round 14: tanks (2026-11-04)

This report covers Protection warrior, Feral bear and Protection paladin: owner rulings 1 and 4, then a fresh re-audit of each spec.

**Setup:**
- Client 1.60.1.70753 with the 2026-11-03 push layered in. The spellconst files already carry the push, and only hunter moved.
- Raw tables: round 11's c70753 copy and wago.tools (SpellEquippedItems, read 2026-11-03).
- Engine `lane/r14-tanks` and site `wt/r14-tanks`, both from int/r14.
- Measurements: 10,000 iterations at seeds 1 and 50001 on the tank fight (`spec-breakdown -tank`), both factions. "Before" is `$R/sb-before` run from `$R/site-before/sim`.
- Nothing nightly-owned was regenerated. The fork goldens regenerated here are listed under "Files to regenerate".

**In short: 4 CONFIRMED-BUGs, all fixed. One owner exception is applied, and one APL adoption is flagged for the owner.**

## Owner rulings

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| O1 | **Protection paladin, band 20: Templar's Bulwark replaces Sacred Duty's second point** (ruling 1).<br>**What changed:** `ladder-first.json` reserves Templar's Bulwark rank 1 first in the list, ahead of Holy Shield and Shieldsworn. It must go first because the walk takes the longest legal prefix of the list, and Holy Shield cannot fit below 30. Band 20 becomes `-15301100…` (it was `-15302000…`). Bands 19 and 25 to 60 are unchanged; the test pins every one of them.<br>**Measured, bare 20 (seeds 1 and 50001 agree within 0.3):**<br>- Chance of death: 75.6 to 62.6% horde, 73.5 to 58.5% alliance.<br>- Damage taken: -2.0%.<br>- Threat: -2.8% (39.7 to 38.6 horde, 36.2 to 35.2 alliance).<br>**This is a stated owner exception to the threat metric, for this band only.** The guide's Leveling section names no band points, so no guide edit was needed. | Test first: `TestProtectionTakesTemplarsBulwarkAtTwenty` (sim/leveling) failed on band 20 and now passes. `TestProtectionTakesHolyShieldAtThirty` now compares bands 40 to 60 against the plain walk, since band 20 is the new test's. Data: `bulwark20.tsv` | owner ruling | - | S 6be4298c3 |
| O2 | **The tank boss had no attack power, and the engine's creature melee followed Wrath's formula.**<br>**What was wrong:**<br>- The engine computed a swing as `min x (roll + AP / (14 x 177))`. Under that, a level 63 boss's 270 attack power was under 9% of its swing, so a demoralize could never matter.<br>- `tank-encounter.json` gave the boss no attack power at all.<br>**What Classic does:** a swing is weapon damage x (1 + 0.15873% x AP), and a level 63 boss has about 270 AP. That makes attack power about 30% of the swing. Two sources agree:<br>- Blizzard Classic forums, "Boss Attacks and You", 2020-03-03. It also gives Broodlord as 2240 ± 448 before attack power.<br>- vmangos `Creature::UpdateDamagePhysical`: weapon damage x (0.7 + 0.3 x AP / database AP), the same slope.<br>The Vanilla wiki's Magmadar and Maexxna test (-16.3% from an untalented 146) fits the same model: 146/630/1.43 = 16.2%.<br>**What changed:**<br>- Engine: the creature swing is now `min x roll x (1 + AP / 630)`. The level 60 and Vaelastrasz presets now carry 268 attack power, replacing Wrath's 805.<br>- Site: the boss gets `attack_power_by_level`, cmangos classic-db's `creature_template_classlevelstats` BaseMeleeAttackPower for the warrior creature class: 268 at level 63, 218 at 53, 162 at 43, 112 at 33 and 80 at 23. cmangos also gives Ragnaros, Nefarian and the other level 63 warrior-class raid bosses 268. The boss's level is the character's plus 3.<br>- The curated minimum swing already includes the boss's full attack power. The engine's weapon roll is that minimum divided by (1 + AP/630), so the mean swing is unchanged until something lowers the boss's attack power.<br>**Effect:** at 60, Demoralizing Shout's or Roar's 204 takes 22.7% off every swing. On its own this change moves no published number: all three tanks are bit-identical at raid 60. | Tests first, both failed and now pass:<br>- `TestEnemyMeleeFollowsClassicAttackPower` (sim/core): ratio 1.0965, want 1.4286.<br>- `TestTankBossCarriesClassicAttackPower` (sim/request): attack power 0.<br>Fork tank goldens regenerated for the changed enemy swings: dtps, plus the weights that damage taken drives. | CONFIRMED-BUG (engine formula); owner ruling (boss attack power) | old (upstream Wrath constant) | E bfa9c25cd, S 82d2129f7 |
| O3 | **The raid debuff options `demoralizing_shout` and `demoralizing_roar` used Era values.**<br>**What was wrong:** they took Era's 146 and 138 attack power, scaled by Era's Improved talents.<br>**What the client says:** Forever's ranks (warrior 1160 to 11556, druid 99 to 9898) are base points plus 1.0 or 1.4 a level up to the rank's max level. Both give 204 at 60. Forever has no Improved Demoralizing Shout or Roar talent, so the "improved" option is the same debuff.<br>**What changed:** both options now take the rank the raid's level has learned, at that level, the way the other raid debuffs do. This was round 13's "Not done" (W9).<br>**Effect:** no preset carries either option, so no published number moves. | Test first: `TestRaidDemoralizeIsTheClients` (sim/core). It did not build before the fix and passes now. | CONFIRMED-BUG | old (upstream) | E a3800aac6 |
| O4 | **Protection now keeps Demoralizing Shout up, and the bear keeps Demoralizing Roar up, refreshed in the last 3 seconds.** The line sits after Thunder Clap for Protection and after Faerie Fire for the bear. The guides already tell players to keep these up. Placing them after Sunder Armor or Maul means they are almost never cast. Raid 60 results:<br>- Protection: damage taken -22.8%, threat -2.5%.<br>- Bear (with B1 and B2 below): damage taken -20.2%, chance of death 98% to 10%, threat -5.4%.<br>**Flagged for the owner:** this is adopted under ruling 4 for its mitigation, but it costs threat, so the plain threat metric would reject it. To hold to the metric, revert S 27edbc55e and E 75338e087.<br>Placing the roar ahead of Primal Bite (priority slot 3) cut damage taken more (CoD 7.5%) but cost -6.8% threat. That is not adopted. | Variants: `rot/prot-demo*.json`, `rot/bear-roar*.json`; results: `ba-item2.tsv`, `final.tsv` | adopted under ruling 4 (owner to confirm) | - | S 27edbc55e, E 75338e087 |

## Protection warrior

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | Shield Slam: rank 4 is 640 to 670 damage (655 with variance 0.0457) plus block value, 20 rage, 6 s cooldown, with no AP term. The engine's comment says the roll "carries no die width", but `damage.Roll` does use the variance. Only the comment is out of date. | 23925 spellconst; Wowhead Forever tooltip 23925 (read 2026-11-03) | Aligned | - | - |
| W2 | Revenge: rank 5 is 109 to 133 plus 0.25 x AP, 5 rage, 5 s cooldown, Defensive Stance only, lit by a block, dodge or parry. Rank 6 (25288) is held back as a trainer-book rank, which ruling 5 keeps pending. | 11601 spellconst + tooltip; `revenge.go` | Aligned (25288 pending, ruling 5) | - | - |
| W3 | Defensive Stance: -10% damage taken, -10% damage dealt, +30% threat. Defiance 3 adds 15%, multiplied in. Shield Block: +75% block for 7 s or 2 blocks, 10 rage, 5 s cooldown. Thunder Clap: 103 damage, 20% attack-speed slow. Sunder Armor rank 5: 450 armor, 206 threat. | 71, 2565, 11581, 11597 rows and tooltips; `stances.go`, `shield_block.go`, `sunder_armor.go` | Aligned | - | - |
| W4 | Guide-build talents checked against their 70753 rank text: Shield Specialization (5%, 100% for 5 rage), Master of Defense (100% for 5 rage), Defiance 15%, Improved Bloodrage +50%, Improved Revenge +60%, Anticipation 20, Deflection 5%, Bastion 10%, Focused Rage -3. All match `talents_protection.go` and `talents.go`. The `talents.go:67` comment still reads "Bastion 1" while the build has 5; that is a comment only. | talents/warrior.json rank descriptions | Aligned | - | - |
| W5 | The engine's Demoralizing Shout is 204.4 at 60 (196 + 6 x 1.4). Wowhead Forever's tooltip shows 210, which is the rank's level-64 cap. Wowhead does the same for the roar (207, the level-62 value). SpellLevels gives BaseLevel 54 (52 for the roar), so the client's own value at 60 is 204. | SpellLevels 101269 / 100331; tooltips 11556, 9898 | Aligned (the tooltip computes at the cap; an in-game tooltip would settle it) | - | - |

## Feral bear

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **Frenzied Regeneration healed 0.3% of health per point of rage.**<br>**What the client says:** spell 22842's description is "Each point of Rage is converted into $s1% health", and effect 0 (aura 23) has amount 1. So each point of rage heals 1% of maximum health. Wowhead Forever's tooltip also reads 1%.<br>**What was wrong:** the engine used a later expansion's 0.3%.<br>**Effect:** with the roar line, bear raid 60 chance of death falls from 17.6 to 9.8%. The fix's own effect on threat is about -0.2%. | Test first: `TestFrenziedRegenerationIsOnePercentAPointOfRage` failed (0.003) and passes. Client Spell.csv 22842; tooltip 22842 | CONFIRMED-BUG | old (86dc10877, 2026-10-07; no earlier round checked the conversion) | E 809b1fc9b |
| B2 | **Demoralizing Roar's threat was twice the rank's level** (104 at rank 5).<br>**What the client says:** each rank carries effect 63, 8/16/26/34/42 plus 0.8 a level up to its max level. Rank 5 at 60 is 48.4 per target. This is the same bug round 13 fixed for Demoralizing Shout (W8).<br>**Effect:** about -0.3% bear threat. | Test first: `TestDemoralizingRoarThreatMatchesClient` did not build before the fix and passes. TestFeralTank golden: three tps rows. | CONFIRMED-BUG | old (86dc10877) | E 68002c42e |
| B3 | Bear and Dire Bear Form passives (1178, 9635): armor from items +180% / +360%, health, and attack power at 3 a level. Bear Form threat (21178) is +50%. Maul is +128, Swipe 83 on 3 targets, Lacerate rank 5 is 75 over 15 s plus 10% weapon damage per stack, Primal Bite rank 3 is weapon damage plus 77. Barkskin: -20% physical damage taken for 15 s. Enrage: 10 rage, then 20 over 10 s, with -27%/-16% armor. | spellconst druid.json; tooltips 9881, 9908, 1235827, 1238070, 22812, 5229 | Aligned | - | - |
| B4 | Thick Hide's 70753 text reads "180 additional base Armor per level". That is the client resolving `$s2` at level 60: effect 1 is base 3 with RealPointsPerLevel 3. The engine's 3 a level at rank 3 gives the same 180 at 60. | SpellEffect 1344589; `thick_hide.go` | Aligned (the text changed, the value did not) | - | - |
| B5 | Demoralizing Roar and Shout roll magic hit against the boss (`OutcomeMagicHit`), so they can miss. Only the engine states this. | `demoralizing_roar.go`, `demoralizing_shout.go` | OPEN (a Forever log would show misses on the shout or roar) | - | - |

## Protection paladin

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Holy Shield: ranks 1 to 4 (20925, 20927, 20928, 1239570) at 75/110/153/221 damage and 100/150/195/240 mana, 4 charges, +30% block, 20% extra threat. **The push's SpellEquippedItems row 255045 (1239570: armor class 4, shield mask 64) matches ranks 1 to 3.** The engine gates every rank on `CanBlock`, so this is a no-op for a paladin wearing a shield (asked by lane physical). | wago SpellEquippedItems 70753; `holy_shield.go:34-41,131` | Aligned | - | - |
| P2 | Righteous Fury: +33% threat and +90% on Holy. Templar's Bulwark: 110 mana, 5 min cooldown, 8 s absorb. Hammer of the Righteous: 6% base mana, 6 s cooldown, 3 x weapon damage per second. All match the conformance golden. Every talent in the guide build matches its 70753 rank text (Redoubt, Precision, Sacred Duty, Shield Specialization, Toughness, Swift Judgement, Improved Righteous Fury, Anticipation, Holy Shield, Iron Creed, Shieldsworn, Deflection, Conviction, Divine Strength, Improved Seals). | conformance/paladin.golden.md; talents/paladin.json; tooltips 25780, 1311015, 407632 | Aligned | - | - |
| P3 | Rank-choice rows for Consecration and Holy Strike (asked by lane data): nothing to re-run. The paladin has no attack-power reduction, and the boss's attack power leaves raid 60 bit-identical, so every rank-row comparison on this fight is unchanged. | `final.tsv` | Aligned | - | - |
| P4 | Holy Shield crit (round 11 O3): still no level 40+ Forever log. | - | OPEN | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Forever's Curse of Recklessness carries no attack-power effect (its second effect is aura 92), so the raid preset raises no boss attack power. | spellconst warlock 11717 | Aligned | - | - |
| X2 | The bear still dies in most bare fights (bare 60 chance of death 87% horde). It has no block or parry and takes 15% crushing blows. That is the Era trade round 13 B3 recorded; it is not a bug found here. | `final.tsv` | Aligned (recorded) | - | - |

## Before/after

Tank fight, 10,000 iterations. Each cell is threat / damage taken, as relative change. Horde and alliance agree within 0.3 points.

| spec | preset | horde s1 | horde s50001 | alliance s1 | alliance s50001 | horde s1 absolute (threat, dt, CoD) |
|---|---|---|---|---|---|---|
| warrior-protection | raid 60 | -2.44% / -22.9% | -2.63% / -22.7% | -2.60% / -22.7% | -2.49% / -22.8% | 1068.8→1042.7, 340.7→262.7, 0→0% |
| warrior-protection | bare 20/30/40/50/60 | -2.3/-3.2/-3.0/-4.0/-4.0% threat; -8.3/-10.5/-11.9/-16.3/-21.1% dt | same within 0.1 | same within 0.2 | same within 0.2 | 60: 653.6→627.2, 376.6→297.3, 4.3→0% |
| druid-feral-bear | raid 60 | -5.49% / -20.0% | -5.39% / -20.2% | -5.49% / -20.2% | -5.39% / -20.2% | 1011.2→955.7, 592.1→473.5, 97.8→10.1% |
| druid-feral-bear | bare 20/30/40/50/60 | -2.4/-2.9/-3.2/-6.0/-6.2% threat; -10.8/-12.2/-11.9/-15.7/-19.7% dt | same within 0.15 | same within 0.15 | same within 0.2 | 60: 444.0→416.7, 639.3→513.3, 100→87% |
| paladin-protection | raid 60 | 0.00% | 0.00% | 0.00% | 0.00% | 1125.5→1125.5 |
| paladin-protection | bare 20 (O1) | -2.8% / -2.0% | -2.5% / -2.2% | -2.8% / -2.0% | -2.8% / -1.9% | 39.7→38.6, 92.4→90.6, 75.6→62.6% |

No spec outside the tanks moves. The engine changes reach only enemy swings, the two raid debuff options and druid bear spells. The core golden tests pass, and no DPS spec's target swings.

## Files to regenerate

- Ladder goldens (`TestRotationLadder`): warrior-protection and druid-feral-bear (the new APL line, the boss's attack power, and Frenzied Regeneration and roar threat on the bear), and paladin-protection band 20 (O1).
- BiS, weights and leveling-strip data for all three tanks: `bis/warrior-protection.json`, `bis/druid-feral-bear.json`, `bis/paladin-protection.json` (band 20 talents), `addon-builds.json` (Protection band 20), dps-per-point, and the coach starters that carry the tank rotations.
- The apl-sync copies: done by hand for warrior-protection and druid-feral-bear (`sim/request/apl/*.apl.json` and the fork's `ui/*/apls/forever_*.apl.json`), so `make apl-check` should pass.
- Fork goldens regenerated here: `sim/warrior/tank_warrior/TestP1TankWarrior.results`, `sim/druid/tank/TestFeralTank.results` and `sim/paladin/protection/TestProtection.results` (E bfa9c25cd, E 75338e087, E 68002c42e).

## Measure at refresh

- All three tanks' weights and gear picks: with the shout and roar up, damage taken falls about 20%, so avoidance and armor weigh differently.
- Protection paladin band 20: the new talents' gear and weights.

## For other lanes

- **Owner / coordinator:** O4. The demoralize lines cost 2.5% (warrior) and 5.4% (bear) threat for about 20% less damage taken. Adopt them under ruling 4, or revert S 27edbc55e and E 75338e087 to hold the threat metric.
- **Guides:** the top-level APL `notes` of warrior-protection and druid-feral-bear said the shout and roar were left out because the boss had no attack power. That claim was false, so I replaced it with a one-sentence engineering note. Please review the wording. The new line notes ("Keep it on the boss: 204 attack power off it at 60 takes about a fifth of its melee damage away") pass `_prose.test.ts`'s banned list (checked by hand; vitest could not start in this worktree). Optional: one sentence in the Protection paladin guide's Leveling section, "take Templar's Bulwark as soon as it opens, ahead of a second point of Sacred Duty".
- **Core:** E bfa9c25cd changes `core/attack.go`'s `EnemyWeaponDamage` and `constants.go`'s `EnemyAutoAttackAPCoefficient` (to 1/630). Only a swinging target uses it. Merge carefully if your tie-break work touches enemy auto-attacks.

## Still open

1. O4 needs the owner's confirmation (above).
2. Revenge rank 6 and Heroic Strike rank 9 trainer-book ranks: pending the owner (ruling 5).
3. B5: whether Demoralizing Shout and Roar roll the spell-hit table against a raid boss.
4. W5: confirm the shout's and roar's in-game tooltip value at 60 (204 from the client, 210/207 on Wowhead).
5. P4: Holy Shield crit.
6. Lacerate and Primal Bite threat x2.0, and Maul and Swipe x1.75: still single-source (round 13).
