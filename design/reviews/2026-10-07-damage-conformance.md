# Damage conformance: the report now sees damage

2026-10-07. Lane `damage-conformance`; site branch `damage-conformance`, fork branch `damage-conformance`.

## The convention

The client states a direct-damage or periodic-damage effect as a centre, a width and a growth rate:

- `EffectBasePointsF` is the centre (`amount` in spellconst).
- `Variance` is the whole width of the roll: the roll is centre x (1 - Variance/2) through centre x (1 + Variance/2). Fireball 12 is 483 with Variance 0.2419, so 424.6-541.4, not 596-760.
- `EffectRealPointsPerLevel` is added to the centre per caster level above the spell's own level, counting levels only up to `SpellLevels.MaxLevel` (0 means no cap).
- `EffectBonusCoefficient` is the spell-power share and is compared only when the table states it (a zero is replaced by the vanilla convention in spellconst, and the report prints it as `(convention)` without comparing).

What shipped:

| Layer | Change |
|---|---|
| Pipeline (`data/pipeline/simconst.py`, `models.py`) | each effect gains `variance` and `points_per_level`, each spell gains `max_level`, all appended last; `data/builds/1.60.1.70009/spellconst/*.json` regenerated |
| `sim/core/spellconst` | the three fields are optional on decode; `Spell.DamageRange(effectIndex, casterLevel)` returns (min, max), tested against Fireball 12 and Lightning Bolt 10 |
| Generator | `<Spell>BaseDamage` is now the client's real {min, max} at the spell's own level; new `<Spell>PointsPerLevel` and `<Spell>MaxLevel` arrays; every `constants_auto_gen.go` regenerated |
| `sim/conformance` | each row gains Damage min-max, Coefficient and a Damage status; `SUMMARY.md` gains a generated per-class block; nine goldens regenerated |

The engine has no uniform accessor for "the base damage this ability rolls", so `core.SpellConfig` and `core.Spell` gain an optional `ClientBaseDamage [2]float64` (per tick for a periodic effect, before spell power, talents and level scaling). It changes nothing at runtime. The report states one of four things per row:

- `n/a`: the client spell has no school-damage or periodic-damage effect.
- `not declared`: it has one and the ability file sets no `ClientBaseDamage`.
- `declared, matches`: both ends of the range within 1, and the coefficient within 0.005 when the table states it.
- `declared, differs`: anything else, with the Diff naming range and coefficient.

A fix lane closes a row by setting `ClientBaseDamage` from the regenerated constants (`<Spell>BaseDamage[rank]` is already the own-level roll) and making the ability roll that range with level growth. The comparison is at the preset's level, so an own-level table at a caster below 60 will correctly read `differs` for any spell with `points_per_level`.

## Counts at level 60

No ability file declares yet (this lane edits none), so every row with a client damage effect is `not declared`. That is the baseline the fix lanes move.

| Class | Declared | Matching | Differing | Not declared | n/a |
|---|---|---|---|---|---|
| Hunter | 0 | 0 | 0 | 11 | 27 |
| Mage | 0 | 0 | 0 | 77 | 24 |
| Warlock | 0 | 0 | 0 | 104 | 55 |
| Paladin | 0 | 0 | 0 | 59 | 83 |
| Warrior | 0 | 0 | 0 | 12 | 38 |
| Druid | 0 | 0 | 0 | 61 | 25 |
| Priest | 0 | 0 | 0 | 43 | 9 |
| Shaman | 0 | 0 | 0 | 123 | 160 |
| Rogue | 0 | 0 | 0 | 6 | 25 |
| **Total** | 0 | 0 | 0 | 496 | 446 |

Rows are spec-and-rank pairs, so a spell two specs register counts twice. The generated copy lives in `sim/core/testdata/conformance/SUMMARY.md` in the fork and is kept current by `TestConformanceSummaryDamageBlock`.

## Ranked engine-vs-client disagreements

Method: for every spell id the curated rotations cast (`data/curated/apl/*.json`, action spell ids), the client roll at level 60 (centre plus `points_per_level` x levels above the spell's level, capped at `max_level`) against the hand table in the ability file. Percent is engine centre over client centre; dots compare total over the duration. A script parsed the hand tables; spells whose damage is weapon-based compare the flat term only. Not a substitute for the report once the lanes declare.

### A. Rolling vanilla tooltip numbers (the lanes that matter)

| # | Spell (rank) | Engine hand table | Client at L60 | Engine over client |
|---|---|---|---|---|
| 1 | Mage Scorch (7) | 237-280 | 166.4-196.4 | +42% |
| 2 | Mage Fireball (12) | 596-760 direct, dot 76 total | 424.6-541.4 direct, 15 x 4 ticks = 60 | +40% direct, +27% dot |
| 3 | Druid Rip (6) | tick base 17 | tick 15 | +13% |
| 4 | Mage Frostbolt (11) | 515-555 | 457.2-492.8 | +13% |
| 5 | Priest Shadow Word: Pain (8) | 852 total (6 ticks) | 127 x 6 = 762 | +12% |
| 6 | Mage Arcane Missiles (8) | 230 per missile | 209 per missile (trigger spell 25346) | +10% |
| 7 | Priest Mind Flay (6) | 426 total | 130 x 3 = 390 | +9% |
| 8 | Paladin Hammer of Wrath (3) | 504-566 | 473.6-522.4 | +7% |
| 9 | Priest Devouring Plague (6) | 904 total | 106 x 8 ticks = 848 | +7% |
| 10 | Priest Mind Blast (9) | 508-537 | 476.9-503.5 | +7% |
| 11 | Mage Fire Blast (7) | 446-524 | 415.4-490.6 | +7% |
| 12 | Paladin Exorcism (6) | 505-563 | 474.7-529.3 | +6% |

Mage, priest and paladin are the real work: the Fireball, Frostbolt, Fire Blast, Scorch, Arcane Missiles, Mind Blast, Mind Flay, SW:P, Devouring Plague, Exorcism and Hammer of Wrath tables are all Classic Era values. Druid Rip and Rake are hand-typed ladders (Rake 4: 58 initial and 32 per tick against 61 and 34, -5%/-6%) that predate the build. Frostbolt's ladder is the Era one even though its cost and cast time already conform.

### B. Client centre but no level growth and no variance (small, systematic)

These were moved to client amounts by hand and sit at the client's own-level centre. They miss `points_per_level` above the spell's level and roll a flat number where the client rolls the Variance width. The mean error is small, but the roll width is wrong and the report will flag them `differs` once declared.

| Spell (rank) | Engine | Client at L60 | Gap |
|---|---|---|---|
| Druid Wrath (8) | 91 flat | 91.6-102.4 (centre 97) | -6%, no width |
| Druid Moonfire (10) | 135 flat | 128.7-150.5 (centre 139.6) | -3%, no width |
| Warlock Shadowburn (6) | 266 flat | 258.3-288.1 | -3%, no width |
| Shaman Chain Lightning (4) | 123 flat | 119.2-133.2 | -3%, no width |
| Shaman Lightning Bolt (10) | 196 flat | 189.9-211.7 | -2%, no width |
| Priest Shadow Word: Death (4) | 448 flat | 444.1-471.9 | -2%, no width |

### C. Right centre, flat roll (width only)

Arcane Blast (5) rolls a flat 394 against 364.2-423.8; Shadow Bolt (10) 268 against 253.3-282.7; Conflagrate (6) 282 against 251.1-312.9; Incinerate (3) 217 against 200.7-233.3; Starfire (7) 381 against 350.0-412.0; Earth Shock (7) 301 against 293.1-308.9; Holy Strike (8) 93 against 81.4-104.6. Mean damage is right, variance is not.

### D. Conformant on the hand tables

Warrior (Rend, Mortal Strike, Execute, Bloodthirst, Overpower, Whirlwind read the generated tables or the client amount), Warlock Immolate, Corruption and Bane of Agony, Hunter Arcane Shot, Aimed Shot, Serpent Sting, Raptor Strike, Mongoose Bite, Counterattack and Immolation Trap, Rogue Sinister Strike, Eviscerate, Ambush, Druid Ferocious Bite, Claw, Shred, Ravage, Insect Swarm, Magma Totem. They still need `ClientBaseDamage` to read `declared, matches`.

### Not comparable from the tables

Searing Totem's client row is the summon (effect 28), its attack is a separate spell. Shaman Stormstrike, Judgement and the Hunter Multi-Shot and Strider Kick rows have weapon-percent effects only. Warlock Wrack (spell 1316697, in the Affliction rotation) has no engine spell at all; that is an implementation gap, not a number.

## Brief for the fix lanes

1. Mage: Scorch, Fireball (direct and dot), Frostbolt, Fire Blast, Arcane Missiles; then flat-to-roll for Arcane Blast. Regenerated `<Spell>BaseDamage` holds the own-level roll to use.
2. Priest: SW:P, Mind Flay, Devouring Plague, Mind Blast, then SW:D width.
3. Paladin: Hammer of Wrath, Exorcism (both tables carry their own level scaling, `scale`, which the client's `points_per_level` supersedes), Holy Strike width.
4. Druid: Rip, Rake, then Wrath/Moonfire/Starfire level growth and width.
5. Everyone: set `ClientBaseDamage` on every registered spell so the goldens read `declared, matches`; a lane is done when its class's `Differing` and `Not declared` counts are zero apart from the not-comparable rows above.
6. Shaman and warlock: width and growth only.

Rows to keep in mind: do not copy the regenerated own-level pair at a caster below the spell's level without the `points_per_level` term, since the report compares at the preset's level.
