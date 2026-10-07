# Self-buff kit: Arcane Intellect, Mark of the Wild, Blessing of Might

2026-10-07, lane self-buffs. Fork branch `self-buffs`, site branch `self-buffs`.

## The rule

`leveling.KitBuffs(spec, level)` (sim/leveling/kit.go) is the buff twin of
`KitConsumes`, decided by spec slug:

- any mage spec, from level 1: `arcane_brilliance` (no improved form; the
  Arcane tree has no Improved Arcane Intellect in Forever).
- any druid spec, from level 1: `gift_of_the_wild`. `:improved` only when the
  spec's guide build takes Improved Mark of the Wild. It is not in the live
  Restoration tree (data/builds/1.60.1.70009/talents/druid.json); the druid
  guide says it became a baseline passive, of a rank the client tables do not
  state. The Feral and Balance builds cannot take it. Every druid spec
  therefore carries the plain buff, the floor.
- any paladin spec, from level 4: `blessing_of_might` (Holy has no Improved
  Blessing of Might in Forever).

The ladder (`ladderKitBuffs`), the ranker (`withKit`, used by `plainRequest`
and `weightsRequest`, so the published stat weights come from the same buffed
character), rotation-search and talent-search all read it. Side effect: the
ranker's weight sweep now also carries `KitConsumes` (rogue poisons, enhancement
imbue), which it did not before.

## Engine: level-aware ranks

Fork `sim/core/buff_ranks.go`. The three buffs keep the raid-buff proto path;
only the value lookup is level-aware. A rank's value is the client amount
(SpellEffect.EffectBasePointsF) scaled so the top rank equals the existing
`BuffSpellValues` entry exactly, so level 60 is bit-for-bit unchanged.

| Buff | Ranks (learn level: amount) |
|---|---|
| Arcane Intellect (1459, 1460, 1461, 10156, 10157) | 1: 2, 14: 7, 28: 15, 42: 22, 56: 31 (Intellect) |
| Blessing of Might (19740 .. 25291) | 4: 14, 12: 25, 22: 40, 32: 61, 42: 83, 52: 112, 60: 133 (attack power) |
| Mark of the Wild armor | 1: 34, 10: 88, 20: 142, 30: 203, 40: 263, 50: 324, 60: 385 |
| Mark of the Wild all stats | 10: 3, 20: 5, 30: 8, 40: 11, 50: 14, 60: 16 |
| Mark of the Wild resistances | 30: 7, 40: 14, 50: 20, 60: 27 |

Arcane Intellect's table equals the engine value at 60 (31), so it is exact at
every rank. Blessing of Might and Mark of the Wild are NOT: the client's level
60 amounts (133 AP; 385 armor, 16 stats, 27 resist) differ from the engine's
legacy level-60 values (185 AP; 285, 12, 20). The task pinned level 60, so the
lower ranks follow the client's proportions anchored on the engine value.
Owner decision pending: adopt the client's level-60 numbers (Retribution
would lose roughly a quarter of the Blessing of Might attack power).

## Movement (relative, ladder DPS, golden diff)

| Spec | 10 | 20 | 30 | 40 | 50 | 60 |
|---|---|---|---|---|---|---|
| mage-arcane | +8% | -4% | -3% | -3% | -2% | 0 |
| mage-fire | -12% | -5% | -4% | -3% | -2% | 0 |
| mage-frost | -14% | -5% | -2% | -3% | -2% | 0 |
| druid-balance | +4% | +3% | +3% | +4% | +4% | +4% |
| druid-feral | +4% | +2% | +2% | +3% | +3% | +3% |
| paladin-retribution | +6% | +5% | +6% | +6% | +7% | +7% |

Every other spec's golden moved only in the rules header. Mage did not gain
at 60: `sim/mage/mage.go` (`AddRaidBuffs`) already sets `ArcaneBrilliance` for
every mage, so the mage was never without Arcane Intellect, only always at the
level-60 value. The kit line is therefore redundant with the class package for
the engine, and the only mage movement is the lower ranks removing a level-60
Intellect bonus from low bands. The mage kit line is kept so the request
states what the character carries. Level-10 movements are small-DPS noise
amplified (a few DPS in absolute terms).

## Deliberately left out

- Inner Fire: armor only, no DPS effect.
- Lightning Shield: reactive damage, modelled by the shaman's own sim.
- Power Word: Fortitude: stamina, no DPS effect.
- Divine Spirit: a Discipline talent, not a class kit buff.
- Battle Shout (warrior casts it, already rank-aware) and Demon Armor (set as
  an option); not changed.
- Improved Mark of the Wild (see above) and any other rank tables.
