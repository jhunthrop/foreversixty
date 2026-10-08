# Heart of the Lion: the hunter's +10% stats and the area aura it sends (2026-10-08)

Lane `lion`. Fork branch `lion` (proto, sim/core, sim/hunter, conformance golden), site branch `lion` (preset, kit, vocabulary, ladder goldens). Both are unmerged; the site's `go.mod` replace pointed at the fork worktree during the work and is not committed, so the site's Go tests and goldens in this branch are right only once the engine pin moves to a fork commit that carries `heart_of_the_lion` (`make engine-pin`, then `python -m pipeline simproto --engine <fork>`).

## Client rows (build 1.60.1.70009, read from `data/builds/1.60.1.70009/raw`)

Heart of the Lion, 409580 (the hunter's spell):

- `SkillLineAbility` 49612: SkillLine 51 (Survival), Spell 409580, MinSkillLineRank 1, ClassMask 4 (hunter), SupercedesSpell 0, AcquireMethod 0. `SpellLevels` 139818: SpellLevel 1, MaxLevel 0.
- `SpellEffect` 1074547, effect 0: Effect 6 (apply aura), aura 137 (mod total stat percentage), EffectBasePointsF 10, misc -1 (all stats), targets 1 / 0 (the caster).
- `SpellEffect` 1084334, effect 1: aura 226 (periodic dummy), EffectAuraPeriod 14000 ms, EffectBasePointsF 10, target 1 (the caster). This is the re-application of the area buff.
- `SpellMisc` 699825: duration index 21 (-1, until cancelled), cast time index 1 (instant), school mask 8.
- Cost and timing, which the brief read as "no cost": `SpellPower` 304862 has PowerType 0 (mana), ManaCost 0 and **PowerCostPct 8**, so the cast costs 8% of base mana; `SpellCooldowns` 84993 has StartRecoveryTime 1500 (the ordinary 1.5 s global cooldown) and no cooldown. The conformance report reads it as 16.48 mana at level 10.
- Description: "increasing total stats by $409583s1% and Attack Power by $409583s2 for all nearby allies, and increasing total stats for the Hunter by an additional $409580s1%".

Heart of the Lion, 409583 (the area buff the periodic dummy re-sends):

- `SpellLevels` 139819: SpellLevel 1, MaxLevel 60.
- `SpellEffect` 1074550, effect 0: aura 137, EffectBasePointsF 10, misc -1, targets 18 / 31, EffectRadiusIndex_1 12.
- `SpellEffect` 1172631, effect 1: aura 99 (attack power), EffectBasePointsF 40, EffectRealPointsPerLevel 4, misc -1, targets 18 / 31, radius index 12.
- `SpellEffect` 1172632, effect 2: aura 124 (ranged attack power), EffectBasePointsF 40, EffectRealPointsPerLevel 4, misc -1, targets 18 / 31, radius index 12.
- `SpellMisc`: duration index 8 (15000 ms), range index 6. The buff lasts 15 s and the hunter's periodic dummy re-sends it every 14 s, so an ally never loses it while in range.

## Scaling

The client's per-level rule (the same one the conformance report reads, `sim/conformance/golden.go`): amount = EffectBasePointsF + EffectRealPointsPerLevel x (caster level - the spell's level), the caster level capped at SpellLevels.MaxLevel. For 409583 that is 40 + 4 x (level - 1) with the level capped at 60: 40 at level 1, 156 at 30, 272 at 59, **276 at 60**, and 276 above. The stat effect has no per-level term: +10% at every level. Melee and ranged attack power are the same number (aura 99 and aura 124 are separate effects with the same row values), as Trueshot Aura already is in the engine.

Engine: `core.HeartOfTheLionRanks` (`sim/core/buff_ranks.go`), one row, `BuffRank.At` floors the amount as it does for every other buff. The site's `sim/leveling/buff_ranks_client_test.go` pins the row against the client's spell constants like the other rank tables (spell level, base, points per level, max level, and the stat effect of both spells).

## Target reading

`ImplicitTarget_0` 18 is "destination: the caster"; `ImplicitTarget_1` 31, with radius index 12, is the area around that destination. SpellRadius row 12 is 100 yards (Radius 100, per level 0, min 0, max 100).

The engine has no handling of target ids at all: nothing in the fork or in the conformance tool reads `ImplicitTarget`, so there is no existing reading of 31 to copy. What the engine does is apply every raid buff to every raid member whatever the client's target says, including the ones the client limits to the party (Battle Shout is target 20; Windfury Totem is a party aura). So the engine's convention already is raid-wide, and the client's own table agrees for 31: the party targets in this table are 20, 33 and 34 (Battle Shout, Ancestral Guidance), and 31 appears paired with 18 on 23 spells (Commanding Shout is one). Reading: **31 is every ally within 100 yards of the caster, not the party**. That is the whole raid for any raid standing within 100 yards of its hunter, which is the case the engine simulates.

## The model

- `proto.RaidBuffs.heart_of_the_lion` (field 41): the area buff, `core.HeartOfTheLionAreaAura`: +10% multiplicative on stamina, agility, strength, intellect and spirit, plus the ranked attack power and ranged attack power, one exclusive category so two sources never add. Applied through `applyBuffEffects`, so a hunter's pet receives it through the pet path exactly as it receives Trueshot Aura.
- The hunter (`sim/hunter/heart_of_the_lion.go`): `AddRaidBuffs` always sets the raid flag (a hunter has the spell from level 1, so a raid with a hunter always has the aura), and `ApplyTalents` registers 409580: a permanent self aura (a further +10% on the five stats, exclusive category of its own) and the castable spell with the client's cost (8% of base mana) and global cooldown, which the conformance report compares (`Heart of the Lion | 409580 | match`, damage columns n/a because the spell has no damage effect, so there is no `declared, matches` to earn; every compared field matches). The spell is never in a rotation: the hunter is wearing it before the pull, and a pre-pull cast spends no combat mana.
- The preset (`data/curated/presets.json`) lists `heart_of_the_lion` after Trueshot Aura with the reason "any Phase 1 raid has a hunter". It is what every non-hunter receives; for a hunter it is redundant, because the hunter sets the same flag.
- The kit (`sim/leveling/kit.go`): every hunter spec carries `heart_of_the_lion` from level 1, ahead of Survival's Hunter's Mark. Also redundant in the engine for the same reason; it states the self buff in the request, where the ranker's applied-buff list and the settings bar read it.

## The stacking rule

1. A hunter character always casts Heart of the Lion on himself: the self spell (+10%) and the area aura it sends (+10%, attack power), the area aura reaching him as it reaches every ally.
2. The raid buff stands for *another* hunter's aura. It is one aura per unit (label and exclusive category), so a preset that names it, a kit that names it, and any number of hunters in the raid add exactly one copy. A hunter who is also given the raid buff reads the same stats as a hunter alone.
3. The two +10% effects multiply, as every pair of percentage stat auras here does (Blessing of Kings is a multiplicative dependency too), so a hunter reads 1.10 x 1.10 = **1.21** of his base stats, not 1.20, and a non-hunter 1.10. Kings and Mark of the Wild multiply with it likewise. The brief's "+20%" is the nominal sum; the client's aura 137 pair multiplies.

Tests: `sim/heart_of_the_lion_test.go` (fork) builds real characters through `core.ComputeStats`: a priest receives 1.10 on every stat and exactly 276 / 156 / 40 ranged attack power at levels 60 / 30 / 1 (the melee side gains the same beyond what strength adds); a hunter alone reads 1.21 of his base stats at 60 and at 30 and gains the ranked attack power beyond agility; adding the raid buff, or a second hunter, changes none of his stats. `TestHeartOfTheLionAttackPowerFollowsClientRule` pins the cap (61 reads 276).

## Movement

The engine keeps raid buffs permanent and takes none off when a unit moves; the client's buff is 15 s long and re-sent every 14 s over a 100-yard radius, so there is no gap while an ally stays within 100 yards. The hunter's own aura has no range at all. Nothing about movement is simulated for this spell: the cast is instant and made before the pull, and a raid member the engine moves is still a raid member in range.

## Ranker, band 60 (before and after)

Before is the fork at its base (`3c754f915`) with the site's unmodified preset and kit; after is the `lion` branch of both. `sim/cmd/leveling-bis -all -bands 60`, default weights iterations, headline figure `set_dps` of each entry, gear re-ranked in both runs. Bare is the class kit alone; raid is the Phase 1 raid preset.

### Alliance

| Spec | Bare before -> after | Raid before -> after |
|---|---|---|
| druid-balance | 225.7 -> 225.7 (+0.0%) | 516.5 -> 521.0 (+0.9%) |
| druid-feral | 302.9 -> 302.9 (+0.0%) | 625.7 -> 706.6 (+12.9%) |
| druid-feral-bear | 138.1 -> 138.1 (+0.0%) | 285.8 -> 329.3 (+15.2%) |
| druid-restoration | 324.7 -> 324.7 (+0.0%) | 600.0 -> 598.5 (-0.2%) |
| hunter-beast-mastery | 243.7 -> 315.8 (+29.6%) | 673.6 -> 835.1 (+24.0%) |
| hunter-marksmanship | 238.0 -> 306.8 (+28.9%) | 656.9 -> 804.5 (+22.5%) |
| hunter-survival | 243.9 -> 299.8 (+22.9%) | 694.9 -> 833.0 (+19.9%) |
| mage-arcane | 484.6 -> 484.6 (+0.0%) | 709.1 -> 712.7 (+0.5%) |
| mage-fire | 411.6 -> 411.6 (+0.0%) | 750.0 -> 757.4 (+1.0%) |
| mage-frost | 329.4 -> 329.4 (+0.0%) | 563.7 -> 566.7 (+0.5%) |
| paladin-holy | 173.1 -> 173.1 (+0.0%) | 510.2 -> 518.0 (+1.5%) |
| paladin-protection | 186.4 -> 186.4 (+0.0%) | 390.0 -> 442.2 (+13.4%) |
| paladin-retribution | 285.0 -> 285.0 (+0.0%) | 624.4 -> 684.6 (+9.6%) |
| priest-discipline | 373.4 -> 373.4 (+0.0%) | 583.9 -> 628.3 (+7.6%) |
| priest-holy | 331.2 -> 331.2 (+0.0%) | 668.9 -> 707.6 (+5.8%) |
| priest-shadow | 326.2 -> 326.2 (+0.0%) | 592.0 -> 592.4 (+0.1%) |
| rogue-assassination | 289.9 -> 289.9 (+0.0%) | 679.4 -> 774.2 (+14.0%) |
| rogue-combat | 263.9 -> 263.9 (+0.0%) | 660.0 -> 741.6 (+12.4%) |
| rogue-subtlety | 263.2 -> 263.2 (+0.0%) | 571.7 -> 659.5 (+15.4%) |
| shaman-elemental | 162.2 -> 162.2 (+0.0%) | 484.1 -> 496.8 (+2.6%) |
| shaman-enhancement | 234.8 -> 234.8 (+0.0%) | 594.6 -> 681.1 (+14.6%) |
| shaman-restoration | 211.7 -> 211.7 (+0.0%) | 564.9 -> 576.2 (+2.0%) |
| warlock-affliction | 464.0 -> 464.0 (+0.0%) | 887.5 -> 912.6 (+2.8%) |
| warlock-demonology | 418.2 -> 418.2 (+0.0%) | 837.8 -> 864.1 (+3.1%) |
| warlock-destruction | 426.6 -> 426.6 (+0.0%) | 806.4 -> 836.6 (+3.7%) |
| warrior-arms | 267.0 -> 267.0 (+0.0%) | 703.6 -> 797.9 (+13.4%) |
| warrior-fury | 301.6 -> 301.6 (+0.0%) | 788.9 -> 916.7 (+16.2%) |
| warrior-protection | 162.8 -> 162.8 (+0.0%) | 278.6 -> 330.1 (+18.5%) |

### Horde

| Spec | Bare before -> after | Raid before -> after |
|---|---|---|
| druid-balance | 227.2 -> 227.2 (+0.0%) | 508.4 -> 513.5 (+1.0%) |
| druid-feral | 304.6 -> 304.6 (+0.0%) | 627.3 -> 716.2 (+14.2%) |
| druid-feral-bear | 141.8 -> 141.8 (+0.0%) | 287.0 -> 332.1 (+15.7%) |
| druid-restoration | 297.7 -> 297.7 (+0.0%) | 596.5 -> 583.3 (-2.2%) |
| hunter-beast-mastery | 250.5 -> 321.3 (+28.3%) | 687.1 -> 847.6 (+23.4%) |
| hunter-marksmanship | 242.9 -> 314.5 (+29.5%) | 675.7 -> 827.9 (+22.5%) |
| hunter-survival | 244.7 -> 303.3 (+23.9%) | 694.4 -> 834.2 (+20.1%) |
| mage-arcane | 470.7 -> 470.6 (-0.0%) | 693.3 -> 696.8 (+0.5%) |
| mage-fire | 397.9 -> 397.9 (+0.0%) | 737.4 -> 746.4 (+1.2%) |
| mage-frost | 320.9 -> 320.9 (+0.0%) | 563.0 -> 565.5 (+0.5%) |
| paladin-holy | 173.6 -> 173.6 (+0.0%) | 504.0 -> 516.5 (+2.5%) |
| paladin-protection | 212.9 -> 212.9 (+0.0%) | 420.9 -> 484.5 (+15.1%) |
| paladin-retribution | 285.8 -> 285.8 (+0.0%) | 605.6 -> 679.6 (+12.2%) |
| priest-discipline | 367.7 -> 367.7 (+0.0%) | 580.3 -> 622.6 (+7.3%) |
| priest-holy | 323.2 -> 323.2 (+0.0%) | 651.1 -> 682.9 (+4.9%) |
| priest-shadow | 329.2 -> 329.2 (+0.0%) | 604.6 -> 605.5 (+0.1%) |
| rogue-assassination | 290.0 -> 290.0 (+0.0%) | 676.4 -> 770.7 (+13.9%) |
| rogue-combat | 265.3 -> 265.3 (+0.0%) | 651.1 -> 732.6 (+12.5%) |
| rogue-subtlety | 262.7 -> 262.7 (+0.0%) | 569.8 -> 651.0 (+14.2%) |
| shaman-elemental | 163.6 -> 163.6 (+0.0%) | 481.4 -> 493.1 (+2.4%) |
| shaman-enhancement | 243.1 -> 243.1 (+0.0%) | 615.0 -> 702.2 (+14.2%) |
| shaman-restoration | 211.2 -> 211.2 (+0.0%) | 564.4 -> 573.5 (+1.6%) |
| warlock-affliction | 451.9 -> 451.9 (+0.0%) | 874.9 -> 905.6 (+3.5%) |
| warlock-demonology | 413.0 -> 413.0 (+0.0%) | 828.4 -> 863.9 (+4.3%) |
| warlock-destruction | 424.4 -> 424.4 (+0.0%) | 800.4 -> 833.4 (+4.1%) |
| warrior-arms | 264.7 -> 264.7 (+0.0%) | 701.5 -> 798.8 (+13.9%) |
| warrior-fury | 303.4 -> 303.4 (+0.0%) | 792.9 -> 923.3 (+16.5%) |
| warrior-protection | 166.3 -> 166.3 (+0.0%) | 280.3 -> 325.8 (+16.2%) |

Reading the tables:

- Every non-hunter bare entry is unchanged, as it must be: the three hunters gain 23 to 30 percent bare (21 percent stats multiplied through agility, plus 276 ranged attack power, on a band-60 set), because a hunter always has the spell. The one 0.0 percent wobble (mage-arcane horde bare, 470.7 to 470.6) is below the ranker's own run-to-run resolution, not a HotL effect.
- Every raid entry rises except two: druid-restoration alliance (-0.2 percent) and horde (-2.2 percent). Both are reproducible (rerun twice, identical) and neither is the buff. Simming the *before* entry's exact gear under the new engine with `sim/cmd/spec-breakdown -heal` (horde, 2000 iterations) reads effective HPS 596.8 before and 619.2 after (+3.8 percent), with the same mana lasting time to within five seconds. The new ranking therefore picked a different gear set whose headline (583.3) is below what the old set now scores: the healer ranker chooses by stat weights and verifies only against the runner-up, and the stat weights moved when every stat grew by a tenth. That is a healer-ranking finding to follow up, not a defect in the spell; the rankings published after this branch merges will carry it until it is looked at.
- Casters gain least in the raid (0.1 to 4 percent; the paladin and shaman healers 1.5 to 2 percent; the two priest healers 6 to 8 percent): the attack power is useless to them and 10 percent of intellect, spirit and stamina is worth little at their stat weights. Melee and tank entries gain 10 to 18 percent, hunters 20 to 24 percent.

## Files

Fork: `proto/common.proto`, `sim/core/proto/common.pb.go`, `sim/core/buff_ranks.go`, `sim/core/buffs.go`, `sim/hunter/heart_of_the_lion.go`, `sim/hunter/hunter.go`, `sim/hunter/talents.go`, `sim/heart_of_the_lion_test.go`, `sim/core/buff_ranks_test.go`, the hunter conformance golden and SUMMARY.

Site: `data/curated/presets.json`, `sim/leveling/kit.go` and tests, `sim/leveling/buff_ranks_client_test.go`, `sim/request/IDS.md`, `sim/request/preset_test.go`, `data/tests/test_presets.py`, `data/builds/1.60.1.70009/simbuffs.json` (named from the fork's tables by the merge-night path, no curated override needed) with its manifest hash line, the three hunter ladder goldens.

Held for the merge: `make engine-pin` and `python -m pipeline simproto` (the vendored proto and `ENGINE_SHA` stay at the old pin in this branch); the fork's `forever` branch moved to `4cfab7871` while this lane ran (conformance lanes for druid, hunter, priest, rogue, warrior and paladin), so the hunter conformance golden here needs a regenerate at merge (`FOREVER_UPDATE_GOLDEN=1`). The shaman-enhancement ladder golden fails on the site's current pin against the fork's base for this lane as well, so it is not touched here.
