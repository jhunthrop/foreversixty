# Rogue parity: why the raid rogues sit at 480 to 540 (2026-10-07)

Branches: site `rogue-parity` (this repo), fork `rogue-parity` (`/Users/jh/code/wowsims-forever/.worktrees/rogue-parity`, from `forever` at 642871a9c). Evidence tool: `sim/cmd/spec-breakdown` (new, committed), which sims a published BiS entry under the request the ranker builds (gear, race and talents from `data/builds/1.60.1.70009/bis/<spec>.json`, then `leveling.KitConsumes`/`KitBuffs`, then the raid preset's `Layer`), 10,000 iterations, one seed, level-63 target of no creature type, 180 s. Its total equals the ranker's `set_dps` for the same entry. Client data: `data/builds/1.60.1.70009/raw` (Spell.csv, SpellEffect.csv, SpellAuraOptions.csv, SpellCooldowns.csv, SpellPower.csv), `talents/rogue.json`, and the conformance goldens in the fork.

## 1. Answer

The three causes the owner asked about are all partly true, in this order of size:

1. **The preset left real damage on the table: about 11% for Combat, 7% for Assassination, 11% for Subtlety.** The raid preset spelled Windfury Totem as a main-hand imbue, so the kit's poisons always won the slot and no rogue ever got Windfury. In the 1.60 client the totem is an aura on every party member, not a weapon enchant (section 4). Moved to a raid buff, it is worth +7 to +9% to a rogue (with its client cooldown). Separately, the kit carried Instant Poison on both weapons for Combat and Subtlety; Deadly Poison in the main hand is 4 to 10% stronger (section 5).
2. **The engine contradicted the client in three rogue talents and one proc cooldown** (section 3): Aggression missed Backstab, Murder was half the client value and applied twice on a crit, Serrated Blades ignored 4.3% armor per rank instead of 3%, and the Windfury proc cooldown was vanilla's 1.5 s where the client states 100 ms. On the published builds the talent fixes move only Subtlety, which takes three Serrated Blades ranks: -0.8% (519.9 to 515.9 under the old kit). Combat takes no Aggression, and Assassination's Murder only acts on Humanoid and Giant targets, so neither moves; the fixes matter for the builds the talent search proposes.
3. **What remains is real in the engine, not a parity miss.** After all of the above the three rogues are at 551 / 579 / 531 against Fury 787 (section 2). The damage tables show where it goes (section 6). Everything the client's tables can check agrees with the engine; the checks the tables cannot make are listed in section 7 with how far each could move the gap. The Blizzard statement that rogues will be the top single-target class is not reproduced by any number the client data gives us.

Movement, raid preset, band 60 alliance entry, same gear unless the ranker changed it:

| Spec | Before | Totem aura + 100 ms | + talent fixes | + kit poisons (final) | Ranker, regenerated set |
|---|---|---|---|---|---|
| Combat raid | 495.4 | 530.0 | 530.0 | 551.0 | 549.2 (was 495.7) |
| Assassination raid | 539.5 | 578.9 | 578.9 | 578.9 | 574.9 (was 539.5) |
| Subtlety raid | 478.4 | 519.9 | 515.9 | 530.5 | 528.5 (was 478.8) |
| Combat bare | 238.2 | n/a | n/a | 257.2 | 253.3 (was 237.7) |
| Assassination bare | 267.9 | n/a | n/a | 267.9 | 268.4 (was 268.4) |
| Subtlety bare | 233.1 | n/a | n/a | 248.8 | 248.6 (was 233.3) |
| Fury raid | 777.8 | 786.7 | n/a | n/a | not re-ranked |

Horde entries: Combat raid 496.7 to 547.9, Assassination 533.7 to 569.1, Subtlety 478.8 to 527.2. The ranker's regenerated sets differ from the published ones in a few slots (head, shoulder, gloves, a trinket), because the weights moved with the buff; the first five columns hold the published gear fixed. The regenerated files are in the scratch directory, not committed (`bis/*.json` is the nightly's).

Windfury Totem moves every melee spec the same way (before to after, raid, same gear and seed): Combat +7.0%, Assassination +7.3%, Subtlety +8.7%, Fury +1.1% (it already had the imbue; the rest is the 100 ms cooldown and the main-hand Dense stone), Arms +0.5%, Survival +3.3%, Enhancement 0 (Windfury Weapon on the main hand disables the totem, as the client text says), Feral 0 (excluded, see section 8).

## 2. Breakdown, side by side (raid preset, final engine and kit, band 60 alliance gear)

Damage per second by source, share of total in brackets. "Glance DPS" is the damage the glancing swings dealt after the penalty.

| Source | Combat | Assassination | Subtlety | Fury |
|---|---|---|---|---|
| Total | 551.2 | 579.5 | 531.0 | 786.8 |
| Main-hand whites | 144.7 (26.3%) | 139.8 (24.1%) | 141.7 (26.7%) | 119.4 (15.2%) |
| Off-hand whites | 87.7 (15.9%) | 80.1 (13.8%) | 69.9 (13.2%) | 134.5 (17.1%) |
| Windfury extra main-hand attacks | 36.8 (6.7%) | 49.2 (8.5%) | 50.4 (9.5%) | 50.0 (6.4%) |
| Builder | Backstab 150.9 (27.4%) | Mutilate 182.5 (31.5%) | Hemorrhage 107.5 + Ghostly Strike 25.3 (25.0%) | Bloodthirst 123.0 + Heroic Strike 155.1 + Whirlwind 46.5 (41.4%) |
| Finisher | Eviscerate 71.3 (12.9%) | Eviscerate 24.2 (4.2%) | Rupture 39.7 + Eviscerate 23.0 (11.8%) | Execute 138.0 (17.5%) |
| Poisons | Deadly 37.3 + Instant 9.8 (8.6%) | Deadly 51.3 + Instant 31.6 (14.3%) | Deadly 44.0 + Instant 15.2 (11.2%) | n/a |
| Item procs | 12.7 | 20.8 | 12.8 | 20.3 |

Hit tables of the white swings (percent of swings): Combat crit 39.1, miss 16.2 to 16.3, dodge 4.5, glance 40.0; Assassination crit 29.2, miss 11.3, dodge 6.4, glance 40.0; Subtlety crit 31.0, miss 14.6 to 14.8, dodge 6.5, glance 40.0; Fury main hand crit 38.1, miss 12.0, off hand crit 39.1, miss 1.5 (the off-hand weapon carries more weapon skill), glance 40.0.

Glancing blows: 40% of white swings in every column, mean damage 65% of a full hit (see section 7). Lost to the penalty: Combat 36 DPS (6.6% of its total), Assassination 39 (6.7%), Subtlety 38 (7.2%), Fury 38 (4.9%).

Special-attack tables: Backstab 29.1 landed of 30.4 cast, crit 67.7% (Puncturing Wounds adds 30), average 935; Eviscerate five points crit 38.1%, average 1,639; Mutilate crit 42%, 27.6 landed hits per hand per fight, average 585 to 607; Bloodthirst 23 landed of 24.4, average 971. No yellow attack misses (the specials sit at the 8% cap); dodge 4.4 to 6.5%.

Cooldown, aura and resource rows:

| Row | Combat | Assassination | Subtlety | Fury |
|---|---|---|---|---|
| Adrenaline Rush uptime | 8.3% (1 cast, 5 min cooldown) | not taken | not taken | n/a |
| Blade Flurry uptime | 8.3% (1 cast, 2 min cooldown) | not taken | not taken | n/a |
| Slice and Dice uptime | 8.9% (opener only, 1.0 casts) | 70.1% (6.5 casts) | 43.0% (3.8 casts) | n/a |
| Windfury Totem aura uptime (procs per fight) | 2.9% (24.3) | 2.3% (18.3) | 4.5% (31.1) | 7.3% (29.8) |
| Venom uptime | not taken | 45.4% (4.8 casts) | not taken | n/a |
| Energy from regeneration per fight | 1,951 (includes Adrenaline Rush's doubling) | 1,801 | 1,831 | n/a |
| Energy from Relentless Strikes per fight | not taken | 285 | 310 | n/a |
| Flurry / Death Wish / Recklessness uptime | n/a | n/a | n/a | 90.3% / 21.0% / 8.3% |

Talent-granted combo points: Combat's Puncturing Wounds adds 13.1 per fight and its Ruthlessness 5.3; Relentless Strikes returns 285 energy per fight to Assassination and 310 to Subtlety.

**Poisons by hand.** The engine registers one spell per poison, not per hand, so the table above cannot split them. Measured by turning each hand's poison on alone (raid preset, 10,000 iterations): Combat with no poison 503.9; Deadly main hand only 541.3; Instant main hand only 520.8; Instant off hand only 513.8; Deadly off hand only 528.0; Deadly main plus Instant off hand 551.0. Assassination with none 467.5; Deadly main only 546.9; Deadly off only 547.6; Instant main only 497.3; Instant off only 498.9; Deadly main plus Instant off 578.9. A first Deadly is worth about 80 DPS to Assassination (the Mutilate bonus) and 37 to Combat; the second Deadly adds nothing (the 5-stack cap, section 5).

## 3. Client parity, ability by ability

"Verdict" is the engine against the client file for build 1.60.1.70009.

| Check | Engine | Client | Verdict |
|---|---|---|---|
| Energy regeneration | 20.2 per 2.02 s, 10 per second; Adrenaline Rush multiplies the tick | The class power table that holds the rate (PowerType) is not in the extracted tables; Adrenaline Rush (13750) states +100% | Rate cannot be checked from the data we hold; Adrenaline Rush matches. Fully inside section 7. |
| Maximum energy | 100, +5 per Vigor rank (talents/rogue.json: Vigor 5/10) | 100 is the vanilla value; Vigor text matches | Match for Vigor; the base 100 is unchecked like the rate |
| Combo points | Builders award 1 (Mutilate 2), Puncturing Wounds +15% per rank on Backstab, Seal Fate 20% per rank on a builder's crit with a 500 ms cooldown, Ruthlessness 20% per rank | Spell texts match; SpellAuraOptions 14186 (Seal Fate) states ProcCategoryRecovery 500 ms, 14156 (Ruthlessness) 500 ms | Match. Ruthlessness has no cooldown in the engine; finishers are never within 500 ms, so it cannot matter |
| Relentless Strikes | 20% per combo point spent to restore 25 energy | talents/rogue.json: "20% chance per Combo Point to restore 25 Energy" | Match |
| Eviscerate | roll of the rank's base (97 to the top of its range at rank 8) + 151 per point + 0.03 per point times attack power; Improved Eviscerate 7/13/20% and Aggression added to the multiplier | SpellEffect 11300: base 96 with variance 1, EffectPointsPerResource 151; text "increased by Attack Power" with `$b1` and `$<mult>`, no attack power coefficient in any column (EffectBonusCoefficient is 1 on every ability, BonusCoefficientFromAP 0) | The flat terms match. The client gives no attack-power number anywhere, so 0.03 per point is vanilla's and is neither confirmed nor contradicted. Rank 8 at 5 points averages 1,639 with 38% crits. |
| Sinister Strike, Backstab | flat roll + normalised main-hand weapon damage (2.4 or the dagger's 1.7), Backstab times 1.5 | 11294 and 11281 have SpellEffect type 121 (normalised weapon damage) with a flat 69 and 141; Backstab also has type 31 at 150% | Match: Forever normalises these. The text "your normal weapon damage" on Sinister Strike is the normalised effect. |
| Instant Poison | 20% per hit, roll 88 +/- about 12 at rank 6, Nature, can crit | SpellAuraOptions 11340 and 11339: ProcChance 20; the damage spell's roll is as client_damage.go | Match |
| Deadly Poison | 30% per hit, five stacks, tick 27 at rank 4 for 12 s | 11355, 11356: ProcChance 30; the stacking aura (11353 and the rest) has CumulativeAura 5 | Match. The Phase 1 rank is IV: rank V (25351) is the AQ rank. |
| Slice and Dice | +30% melee speed at rank 2, 6 + 3 per point seconds, times 1.15/1.30/1.45 | 6774: aura type 319 with 30; text 9 to 21 seconds per point; Improved Slice and Dice 15% per rank | Match; the haste reaches both hands (pinned by `TestSliceAndDiceHastesMeleeSwingsForTheClientDuration`) |
| Blade Flurry | +20% melee speed, 15 s, 25 energy, 2 min | 13877: aura 319 with 20; SpellPower 25; SpellCooldowns 120000; conformance row matches | Match |
| Adrenaline Rush | +100% energy regeneration, 15 s, 5 min | 13750: aura 110 with 100; SpellCooldowns 300000 | Match |
| Windfury Totem | main-hand imbue Windfury (an aura applied through the imbue slot), 1.5 s cooldown | 8512, 10613, 10614 and their proc aura 8515/10612: 20% per main-hand auto or melee spell, ProcCategoryRecovery 100 ms; buff 8516/10608/10610 aura 99 (attack power 95/179/246), 1 s, 2 charges | Was a mismatch on the cooldown and on being an imbue; fixed (section 4) |
| Dual-wield miss and off-hand penalty | miss + 19% flat on whites with two weapons, none on specials; off-hand damage from Dual Wield Specialization 5% per rank | The client has no table for the flat 19%; the Forever tree's Dual Wield Specialization is damage only (talents/rogue.json) | The 19% is the engine's own marked "unconfirmed"; the 5% per rank matches |
| Glancing blows | chance 10% + 2% per defence point above skill (40% at 300 against 315), damage 35% off at 15 points (mean 65%) | Blizzard: weapon skill "still works as it always has, but items with weapon skill offer less"; racials that gave skill now give crit; no note changes glancing | Vanilla table, unconfirmed either way (section 7). |
| Aggression | Backstab had no term | talents/rogue.json: "Sinister Strike, Backstab, and Eviscerate ... by 2%/4%/6%" | Contradiction; fixed |
| Murder | 1%/2% on Humanoid, Giant, Beast, Dragonkin, also multiplied into the crit multiplier | "all damage dealt by 2%/4% against Humanoid and Giant" | Contradiction; fixed |
| Serrated Blades | `float64(5/3*rank*level)` rating, about 4.3% armor per rank | "ignore 3%/6%/9% of your target's Armor"; Rupture 10% per rank | Contradiction in the armour term (Rupture matches); fixed |
| Hack and Slash | 1% per rank extra attack on an axe or sword, 1% crit per rank with dagger or fist, 3% armour ignore per rank with a mace, 200 ms cooldown | tree text; SpellAuraOptions 1290312: chance 5 at rank 5, ProcCategoryRecovery 200, mask 20 | Match |
| Thousand Cuts, Cutthroat, Quietus, Puncturing Wounds, Precision, Malice, Opportunity, Initiative, Lethality, Dual Wield Specialization, Vigor, Hemorrhage, Ghostly Strike, Improved Eviscerate, Improved Sinister Strike, Flawless Execution, Weapon Expertise, Mutilate, Venom, Cold Blood | per-rank constants in `client_values.go` and `talents.go` | each read against talents/rogue.json | Match (most were fixed by the earlier rogue-curate lane; Thousand Cuts' 1.9 s trigger cooldown on SpellAuraOptions 1310721 is not modelled and cannot matter at one Rupture tick per 2 s) |
| Talents with no damage lever | Improved Gouge, Remorseless Attacks, Camouflage, Master of Deception, Dirty Tricks, Improved Distract, Heightened Senses, Improved Kidney Shot, Endurance, Improved Sprint, Improved Kick | | Unmodelled by design (utility or kill-triggered); the talent search lists them as engine gaps. Remorseless Attacks needs a kill the Patchwerk fight never produces. |

Restless Blades (Power Up Gaming, research/06) is not in the 1.60.1.70009 tree, so nothing models it.

## 4. Windfury Totem

Client evidence (Spell.csv, SpellEffect.csv, SpellAuraOptions.csv): the learnable totem ranks 8512 (level 32), 10613 (42), 10614 (52) read "The totem enhances the melee attacks of all party members within N yards. Each main hand hit has a 20% chance of granting the attacker 1 extra attack with ... extra melee attack power." The totem's aura (8515, 10612) carries ProcChance 20, ProcCategoryRecovery 100 ms, proc mask 20 (auto attacks and melee spells); the buff it grants (8516, 10608, 10610) is effect 6 aura 99 (attack power 95/179/246) plus effect 19 (one extra attack), two charges, one second. None of these is an enchant effect. The shaman's Windfury Weapon (8232, 8235, 10486, 16362) is the enchant (effect 360) and says "When applied to main hand, disables any benefit you personally receive from Windfury Totem".

Engine change (fork commit `feat(core): Windfury Totem is a raid-buff aura`): `RaidBuffs.windfury_totem = 40` (proto regenerated; only `common.pb.go` changed), applied through the same extra-attack aura the old imbue path built, skipped for a shaman with Windfury Weapon on the main hand and for feral combat; `extraAttackProcICD` is now the client's 100 ms. Tests: `TestWindfuryTotemRaidBuffGivesTheRogueTheExtraAttackAura`, `TestWindfuryTotemReaches`, `TestWindfuryTotemProcCooldownMatchesClient`. The existing imbue enum value is unchanged (other specs' tests use it).

Preset change (`data/curated/presets.json`): `windfury_totem` joins the raid buffs; `main_hand_imbue:windfury` leaves both consumable groups; dual wielders carry `main_hand_imbue:dense_sharpening_stone` and `off_hand_imbue:elemental_sharpening_stone` (Elemental in both hands would count its character-wide 2% crit twice in this engine; Dense is damage on one weapon). Single-weapon melee specs get no main-hand stone: the brief named dual wielders only, and giving the slot to Arms, Retribution and Survival moves other lanes' numbers; the same Dense stone is the natural choice when that lane adopts it. Tests in `sim/request/preset_test.go` pin the buff for every spec, the absence of the imbue, the rogue's two poisons next to the totem, and the dual wielder's two stones. The `windfury_totem` id is in `IDS.md`.

## 5. Poison arrangement per spec

The instruction was to sim a Windfury main hand with a poison off hand; the client finding makes the totem independent of the weapon, so what remained to choose was the poison pair. Four arrangements per spec, raid preset (bare in brackets), 10,000 iterations:

| Spec | Instant / Instant (old kit) | Instant main, Deadly off | Deadly main, Instant off | Deadly / Deadly |
|---|---|---|---|---|
| Combat | 530.0 (238.2) | 544.5 (253.8) | **551.0 (257.2)** | 547.3 (255.8) |
| Assassination | 529.1 (242.0) | 577.3 (272.3) | **578.9 (267.9)** | 552.1 (241.7) |
| Subtlety | 515.9 (232.2) | **535.4 (250.5)** | 530.5 (248.8) | 516.8 (237.6) |

Bands 30, 40, 50 (bare, 5,000 iterations): Deadly main hand with Instant off hand is the best for Combat (60.8, 113.5, 166.7 against 54.2, 104.2, 149.9 for two Instants) and Subtlety (51.0, 98.5, 142.2 against 44.2, 88.3, 127.0) at every band. Two Deadlys lose to one because Deadly caps at five stacks on a 12-second refresh, so a second source adds almost nothing while giving up Instant's independent damage. Assassination's Deadly main hand is also what Mutilate's bonus needs (Mutilate's bonus counts only a lingering poison, so the Deadly can sit in either hand; main hand is 0.3% ahead).

Kit decision (`sim/leveling/kit.go`): every rogue spec carries Deadly Poison in the main hand and Instant Poison in the off hand from level 30; below 30, two Instants. Subtlety at level 60 alone prefers the mirror by 0.9% (535.4 against 530.5); one rule is kept and the 0.9% is recorded here as the cost. Test: `TestKitConsumesRogueCarriesDeadlyInTheMainHandFromThirty`. Ladder goldens for the three rogue specs moved for this and for Aggression (Combat level 60: 210.4 to 229.9) and are regenerated.

## 6. Why Fury is still ahead

Same gear class, same buffs, same fight. What the tables show:

- **Fury has a damage window and an execute that rogues have no counterpart for**: Execute is 17.5% of its damage, Flurry is up 90% of the time, Death Wish 21%, Recklessness 8%. The rogue cooldowns are Adrenaline Rush and Blade Flurry, one use each in 180 s (5 and 2 minute cooldowns, confirmed against SpellCooldowns), 8.3% uptime each, and neither is a damage multiplier: Adrenaline Rush doubles energy and Blade Flurry is +20% speed.
- **Rogue damage is energy-limited**: 1,800 to 1,950 energy per fight, spent at 15.6 damage per energy on Backstab (60 energy, 935 average) and 46 per energy on a five-point Eviscerate (35 energy, 1,639), with the builder paying for the combo points. Warrior rage is generated by the hits themselves (about 1,600 per fight from whites alone), so a Fury's yellow-attack budget grows with its gear and its Flurry.
- **Rogue whites carry a larger penalty**: 16% of Combat's swings miss (the BiS set reaches about 10.7% hit against the 27% a dual wielder needs), against 12% for Fury's main hand and 1.5% for its off hand, and the glancing loss is 6.6% to 7.2% of rogue damage against 4.9% of Fury's.
- **Slice and Dice is worth less than the combo points it costs** in this engine, and the earlier lane measured the same: a five-point Slice and Dice buys +30% speed for 21 to 30 seconds, which scales whites and poison procs (about 45% to 50% of rogue damage), an estimated +15%, against the 1.4 Eviscerates the same points would have cast (an estimated 18%; the earlier lane's pure-finisher sims gave 223.1 for Slice and Dice only against 229.9 for Eviscerate only). Combat's rotation therefore opens Slice and Dice once and puts every later five points into Eviscerate (uptime 8.9%).

Sizing the part this cannot see (analytic, not simmed): see section 7.

## 7. What remains, and why

Real gaps and unverifiable inputs, with the size of each if it were wrong. Sizes are first-order: the share of damage the input scales, not a re-sim.

| Item | State | If wrong, rogues move by roughly |
|---|---|---|
| Energy regeneration rate and the 100 maximum | The class power table is not in the extracted data. The engine's 10 per second and 100 are the vanilla values. | A 10% faster rate scales the special-attack and finisher share (about 45% of Combat damage): +4 to 5%. |
| Dual-wield flat miss penalty | 19%, marked "unconfirmed" in `sim/core/target.go`; no client table. | Without it Combat's whites would miss about 0% instead of 16%: about +9% for Combat, +3% for Fury (its main hand misses 12%, its off hand 1.5%). This one favours rogues if Forever changed it. |
| Glancing blows | Vanilla table (40% chance, mean 65%). Blizzard's weapon-skill note does not say it changed. | Without the penalty, +6.6% to +7.2% for rogues against +4.9% for Fury. |
| Eviscerate's attack power term | 0.03 per point, vanilla's. The client text says "increased by Attack Power" and states no number. | The term is roughly a fifth of a five-point Eviscerate (attack power is not tabulated here) and Eviscerate is 13% of Combat's damage: at 0.04 per point Combat gains about 1%. |
| Windfury Totem for feral | Excluded in the engine (`FeralCombatEnabled`). The client text says "main hand hit", which a cat claw is not; unverified. | Feral +6 to 8% if the totem applies to it (by analogy with the other melee specs). |
| Single-weapon specs' main-hand stone | Not given by this change (brief named dual wielders). | Arms, Retribution, Survival a point or two. |

The Blizzard developer statement is not contradicted by anything checkable here, and is not supported by it either. If rogues are to be the top single-target class, one of the first four rows has to differ from vanilla by an amount the client tables we hold do not show. The first thing to get is the PowerType table for energy (it is a DB2 table the extraction does not take) and a Forever weapon-skill/miss measurement from the beta.

## 8. Searches and adoption

Rotation search (`go run -C sim ./cmd/rotation-search -spec rogue-<spec> -repo-root <worktree>`, seed 7, 800 iterations, bare, on the final engine and kit; reports in the scratch directory, not committed):

| Spec | Baseline | Verdict |
|---|---|---|
| Combat | 262.6 +/- 0.4 | already the best found; Rupture inserted loses 15.0 |
| Assassination | 279.2 +/- 0.4 | removing Venom gains +4.5 (+1.6%) beyond error: **held**. Sim of that rotation on the published entry: bare 272.1 against 267.9, raid 587.7 against 578.9. Venom is the spec's finisher talent and an earlier standing instruction keeps every finisher casting. |
| Subtlety | 259.2 +/- 0.3 | removing Eviscerate and Ghostly Strike gains +6.1 (+2.4%), but the winner never casts Eviscerate: **not adoptable** (the tool says so) |

No rotation was changed, so no `apl-sync` was needed; `make apl-check ENGINE_DIR=<fork worktree>` passes. The Combat guide's rotation paragraph said Slice and Dice is kept up at all times and Eviscerate is cast at four or more; the curated rotation opens it and spends five points on Eviscerate. The guide now says what the rotation does and states the poison arrangement.

Talent search (`talent-search`, same seed, scratch directory): none of the three qualifies under the owner's build rule (keeps every unmodelled talent, holds the spec tree, at least 1% beyond error, drops nothing the rotation casts).

| Spec | Guide (bare, 800 it.) | Best that keeps every unmodelled talent | Why not adopted |
|---|---|---|---|
| Combat | 262.6 | +2.1 (+0.8%): one Improved Sinister Strike point to Aggression | under 1% |
| Assassination | 279.2 | +6.9 (+2.5%): one Venom point to Ruthlessness | drops Venom, which the rotation casts |
| Subtlety | 259.2 | +1.4 (+0.5%): deep Subtlety | under 1% |

Information for the owner, not adopted: respending every non-damage point (Improved Gouge, Remorseless Attacks, Camouflage, Master of Deception, the Combat tree's utility) is worth much more than any variant that keeps them. On the published entries, raid preset: Combat 626.4 against 551.0 (+13.7%; bare 294.7 against 257.2); Assassination 626.7 against 578.9 (+8.3%; bare 292.2 against 267.9); Subtlety 532.8 against 530.5 (+0.4%). The guide builds keep those points for reasons the owner set. Aggression, Murder and Opportunity are now correct in those builds, which is why the Combat winner takes three Aggression ranks.

## 9. Changes

Fork `rogue-parity`:

- `feat(core): Windfury Totem is a raid-buff aura, with the client's 100 ms proc cooldown`: `proto/common.proto` (`RaidBuffs.windfury_totem = 40`), `sim/core/proto/common.pb.go`, `sim/core/buffs.go`, tests in `sim/core/windfury_client_test.go` and `sim/rogue/dps_rogue/windfury_totem_test.go` (the dps_rogue test builder takes raid buffs).
- `fix(rogue): Aggression reaches Backstab; Murder and Serrated Blades follow the client`: `sim/rogue/{backstab,sinister_strike,eviscerate,talents,client_values}.go`, tests in `client_values_test.go`.
- `go test --tags=with_db ./sim/rogue/... ./sim/core/` passes. The wider `./sim/...` run has the same 17 failure lines as the base commit (conformance hunter, Mage, Elemental, Enhancement and DPS Warrior stat-weight and item goldens); none is touched by this lane and no golden was adopted.

Site `rogue-parity`:

- `feat(sim): Windfury Totem is a raid-preset buff; rogues keep both poisons, with Deadly in the main hand`: `data/curated/presets.json`, `sim/leveling/kit.go` and test, `sim/request/preset_test.go`, `sim/request/IDS.md`, the three rogue ladder goldens, `sim/cmd/spec-breakdown`.
- This document, the three rogue guides, and `spec-breakdown`'s `-rotation` flag.
- `go test ./...` under `sim/` passes.
- Not committed, by instruction: `sim/go.mod`'s replace (points at the fork worktree locally), `data/builds/*/bis/*.json`, `simdb.bin`, `spellranks.json`.

Before the owner merges: the fork commit has to be pinned (`make engine-pin`) and `sim.wasm` rebuilt, or the site's request layer rejects `windfury_totem`; the nightly regenerates the BiS files, which will move (rogues up about 7 to 11% in the raid entries).
