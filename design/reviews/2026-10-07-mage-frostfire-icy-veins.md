# Mage: Frostfire Bolt registered, Icy Veins ruled out (2026-10-07)

Lane `mage-ffb`. Fork branch `mage-ffb` (wowsims-forever), site branch `mage-ffb`. Client build 1.60.1.70009.

## What the client says

**Frostfire Bolt** has three trainer ranks, not two (SkillLineAbility AcquireMethod 0, Fire skill line):

| Rank | Spell id | Level | Mana | Direct (client roll at 60) | Dot |
|---|---|---|---|---|---|
| 1 | 401502 | 40 | 205 | 101.9-118.9 | 9 x 3 ticks |
| 2 | 1237312 | 50 | 285 | 180.6-210.6 | 13 x 3 ticks |
| 3 | 1237313 | 60 | 370 | 269.6-314.4 | 19 x 3 ticks |

All three: 3 s cast, 1.5 s GCD, coefficient 0.814 on the direct hit and 0 on the periodic effect (effect 2, aura 3, period 3000 ms, duration 9000 ms), school mask Fire|Frost, 40% snare (effect 0). The generated rank 0 (401735) is the level-1 "Gain the ability" teaching spell and is not castable.

The live tooltip figure for rank 3 (286 to 331) is higher than the table's roll at level 60 (269.6 to 314.4) by 8 levels of scaling (2.1 a level): the tooltip is evaluated at the spell's maximum scaling level (68). Frostbolt rank 11 shows the same offset (tooltip 470-506, table 457-493, 4 levels at 3.2). The engine follows the client table at the caster's level, as it does for every other spell, and the conformance rows read `declared, matches`.

**Icy Veins** (429125) is **not** a trainable Forever spell. The SkillLineAbility table (wago, 1.60.1.70009) gives it AcquireMethod 3 with no learn level: a Season of Discovery rune row. It was briefly registered during this lane and then removed on the coordinator's correction; no engine, test, golden or rotation references it. The client text (3 % base mana, 3 min, 20 % haste, 20 s) stays in the spellconst for if Forever ever teaches it.

## What was modelled (fork `sim/mage/frostfire_bolt.go`)

- One spell per rank, registered by level (40/50/60), ActionID the client id, `ClientBaseDamage` the client roll (conformance compares it), coefficient from the client table, flat mana cost from the client table.
- Direct hit plus a 3-tick, 3 s dot of the flat per-tick amount (`FrostfireBoltDotTickDamage` {9, 13, 19}, equal to the tooltip's 27/39/57), snapshotting like Fireball's, applied only when the bolt lands.
- `SpellSchool` Fire|Frost. The engine's multi-school support does the rest: the highest school multiplier applies (Fire Power, Piercing Ice are class-mask mods that both bind), the lowest of the two resists is used (`resistCoeff`), and every `Matches(Fire)` / `Matches(Frost)` consumer picks it up: Ignite on crit, Combustion, Master of Elements, Winter's Chill, Heating Up (mask), Ice Shards (mask), Elemental Precision (mask), Improved Fireball (mask, now real), Clearcasting (`SpellFlagMage`). Missile Barrage's 20 % roll now includes it (`SpellCode_MageFrostfireBolt`).
- Binary, like Frostbolt (it carries the same 40 % snare on the same hit roll): a resist takes the whole spell, so the lower-of-two rule decides the hit chance, and ticks take no partial resist.
- Tests (written first): by-level registration, school/mask/flags/cast time, lower-of-two resist, direct plus 9 s dot, Improved Fireball cast reduction, Heating Up feed, Missile Barrage roll, client ladder (levels, mana, coefficient, roll at seven caster levels), dot tick table, dot tick count from the client duration.
- Conformance: six new rows in `mage.golden.md`, all `declared, matches`; mage 77 to 80 rows in SUMMARY.md.

### A bug found on the way: Improved Fireball counted twice

`fireball.go` subtracted `100 ms * ImprovedFireball` from the cast time itself while `applyDeclarativeTalents` also applies the same reduction as a `CastTime_Flat` mod on the Fireball|Frostfire Bolt mask. At 5/5 Fireball cast in 2.5 s instead of 3.0 s. The arithmetic term is removed (the new test pins both spells). This moves Fire: the ladder golden `mage-fire.golden.md` (level 60 row about 7 % lower, lower rows 1-4 %) and the rotation-search baselines. The `TestP1Mage` goldens are unchanged (Frost, no Improved Fireball). The adopted-golden change is explained by this fix alone.

## Rotation search (level 60, BiS gear and guide build fixed, seed 7, 800 iterations, paired)

`sim/cmd/rotation-search` already builds its candidate list from every learned-but-uncast ability, so Frostfire Bolt enters the insertion probe and the hill-climb with no code change (Icy Veins is not learnable and is not a candidate). Insertion at the head of the list with no condition:

| Spec | Baseline | Insert Frostfire Bolt (top) | Search result |
|---|---|---|---|
| Fire | 333.7 +- 2.4 | -9.5 +- 3.3, hurts | no mutation beat the incumbent |
| Frost | 303.4 +- 1.8 | -33.3 +- 2.5, hurts | no mutation beat the incumbent |
| Arcane | 465.3 +- 1.2 | -76.4 +- 1.7, hurts | no mutation beat the incumbent |

A head insertion starves every other line, so it understates a filler swap. The direct test is the replacement: the curated rotation with its filler (Fireball 25306 for Fire, Frostbolt 25304 for Frost and Arcane) replaced by Frostfire Bolt rank 3, same seed and iterations:

| Spec | Curated | Frostfire Bolt as the filler | Delta |
|---|---|---|---|
| Fire | 333.7 +- 2.4 | 332.9 +- 2.4 | -0.8 (-0.2 %), within error |
| Frost | 303.4 +- 1.8 | 196.9 +- 1.2 | -106.5 (-35 %) |
| Arcane | 465.3 +- 1.2 | 459.1 +- 1.3 | -6.2 (-1.3 %) |

## Adopted

Nothing. No line beats the curated rotation by more than the combined error, so `data/curated/apl/mage-*.json` are unchanged and `make apl-sync` / `make apl-check` report no change. Fire is a statistical tie: Frostfire Bolt casts 0.5 s faster than Fireball but its direct hit is smaller, and the 0.814 coefficient against Fireball's 1.0 means it falls behind as spell power grows. The three guides' Rotation prose now say so (relative figures only). The rotation-search reports under `design/reviews/rotation-search/mage-*.md` are regenerated.

The mage-fire ladder golden is regenerated (Improved Fireball fix). `rotations_smoke_test.go` needed no change. Pre-existing and unrelated: the ladder logs 4 violations (hunter Arcane Shot, mage-arcane Arcane Blast unresolved at levels 20/30), reported as a skip outside strict mode.

## Open questions

1. Binary or partial: the sim treats Frostfire Bolt as binary because of the snare, like Frostbolt. If the client lets it partial-resist the damage, the effect against a resist-bearing boss is slightly smaller; moot at equal resists.
2. Tier 1 5-piece (1301488): "10% increased chance to trigger Missile Barrage, 10% crit while Combustion is active, 10% increased chance to trigger Fingers of Frost" is not implemented (the set bonus is not modelled in `item_sets_pve.go` and Fingers of Frost is absent from the engine). Frostfire Bolt is the spell it would make worthwhile for Fire, so this is the one change that could move the verdict; re-run the search after it lands.
3. Fingers of Frost is not modelled, so any Frostfire Bolt interaction with it is absent.
4. The 40 % snare and Frostfire's chill do not matter on a stationary Patchwerk target.
5. If Icy Veins is ever taught, the client text is in the spellconst: instant, off the GCD, 3 % base mana, 3 min, +20 % cast speed for 20 s; a Cold Snap reset is plausible.
