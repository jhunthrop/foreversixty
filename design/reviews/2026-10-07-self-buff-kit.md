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

## Engine: client rank amounts

Fork `sim/core/buff_ranks.go`. The three class buffs keep the raid-buff proto
path and Battle Shout joins them; each takes the highest rank the character's
level can learn and applies the client's amount (coordinator decision: Forever
rebalanced these buffs, the engine's old fixed values were vanilla AQ-era
numbers, so level 60 moves too). Rows are pinned against the committed client
spellconst (SpellEffect base points, points per level, SpellLevels) by the
site test `sim/leveling/buff_ranks_client_test.go`, so a new build moves that
test, not the number.

| Buff | Ranks (learn level: amount) |
|---|---|
| Arcane Intellect (1459, 1460, 1461, 10156, 10157) | 1: 2, 14: 7, 28: 15, 42: 22, 56: 31 Intellect |
| Blessing of Might (19740, 19834 .. 19838, 25291) | 4: 14, 12: 25, 22: 40, 32: 61, 42: 83, 52: 112, 60: 133 attack power |
| Mark of the Wild armor (1126, 5232, 6756, 5234, 8907, 9884, 9885) | 1: 34, 10: 88, 20: 142, 30: 203, 40: 263, 50: 324, 60: 385 |
| Mark of the Wild all stats | 1: 0, 10: 3, 20: 5, 30: 8, 40: 11, 50: 14, 60: 16 |
| Mark of the Wild resistances | 1 to 20: 0, 30: 7, 40: 14, 50: 20, 60: 27 |
| Battle Shout (6673, 5242, 6192, 11549, 11550, 11551, 25289) | 1: 9, 12: 21, 22: 33, 32: 51, 42: 78, 52: 111, 60: 139 attack power |

Battle Shout also carries the client's points per level (0.3, 0.3, 0.3, 0.6,
0.6, 0.6, 0.6) up to the rank's max level (11, 21, 31, 41, 51, 61, 61), so a
rank grows inside its level span (rank 4 is 55 at level 40). The other three
buffs have no per-level term.

Ahn'Qiraj book ranks: Blessing of Might 25291 and Battle Shout 25289 are
flagged `AhnQiraj` and resolve only while `core.IncludeAQ` is set. With the
flag false (the coordinator flips it after this lane merges) a level-60
character resolves to rank 6: Blessing of Might 112, Battle Shout 111 plus 0.6
per level (115 at level 60). Mark of the Wild rank 7 (9885) and Arcane
Intellect rank 5 are trainer ranks either way. Both behaviours are pinned in
`sim/core/buff_ranks_test.go`. The ladder goldens here are generated with
`IncludeAQ` true and will move again when it flips.

The warrior's own Battle Shout (`sim/warrior/shouts.go`, `battleShoutGrant`)
now grants the same table's amount for the rank it casts instead of a
hard-coded 232 (test `shouts_test.go`). The paladin registers no self-buff
aura of this kind besides Seal of the Crusader, whose attack power (31, 51,
94, 145, 221, 306 at levels 6/12/22/32/42/52 plus 0.7/1.1/1.7/2.0/2.2/2.4 per
level to 12/20/30/40/50/60) already equals the client rows
(spells 21082 .. 20308); Holy Shield and Blessing of Sanctuary are not stat
buffs. Nothing to change there.

## Movement (relative, ladder DPS, golden diff, IncludeAQ true)

| Spec | 10 | 20 | 30 | 40 | 50 | 60 |
|---|---|---|---|---|---|---|
| mage-arcane | +8% | -4% | -3% | -3% | -2% | 0 |
| mage-fire | -12% | -5% | -4% | -3% | -2% | 0 |
| mage-frost | -14% | -5% | -2% | -3% | -2% | 0 |
| druid-balance | +6% | +3% | +5% | +5% | +5% | +5% |
| druid-feral | +6% | +2% | +3% | +4% | +4% | +4% |
| paladin-retribution | +4% | +4% | +5% | +5% | +5% | +5% |
| warrior-arms | -3% | -3% | -2% | -3% | -3% | -4% |
| warrior-fury | -4% | -4% | -3% | -3% | -5% | -4% |

Every other spec's golden moved only in the rules header. Warriors fall
because Battle Shout is now 139 at 60, not 232. The raid-buff test profiles in
the fork (improved Gift of the Wild, Blessing of Might and so on) moved the
TestP1Mage, TestElemental, TestEnhancement and TestP1DPSWarrior `.results`
files, adopted because this change explains them.

Mage did not gain at 60: `sim/mage/mage.go` (`AddRaidBuffs`) already sets
`ArcaneBrilliance` for every mage, so the mage was never without Arcane
Intellect, only always at the level-60 value. The kit line is redundant with
the class package for the engine; the mage movement is the lower ranks
removing a level-60 Intellect bonus from low bands. The line is kept so the
request states what the character carries. Level-10 numbers are a few DPS in
absolute terms.

## Audit: client vs engine, for the hunter and warlock lanes (report only)

Aspect of the Hawk (spells 13165, 14318, 14319, 14320, 14321, 14322, 25296;
`sim/hunter/aspects.go`, `getMaxAspectOfTheHawkAttackPower`):

| Rank | Level | Client ranged AP | Engine |
|---|---|---|---|
| 1 | 10 | 20 | 20 |
| 2 | 18 | 35 | 35 |
| 3 | 28 | 50 | 50 |
| 4 | 38 | 70 | 70 |
| 5 | 48 | 90 | 90 |
| 6 | 58 | 55 | 55 |
| 7 | 60 | 120 | 120 |

All match. Note for the hunter lane: rank 6 grants less than rank 5, and the
engine casts the highest learned rank, so levels 58 and 59 sim 55 where a
player would stay on rank 5 (90). Rank 7 is a book rank candidate under the
Ahn'Qiraj flag decision.

Demon Armor (spells 706, 1086, 11733, 11734, 11735; `sim/warlock/armors.go`):

| Rank | Level | Client armor / shadow resistance | Engine |
|---|---|---|---|
| 1 (706) | 20 | 210 / 3 | 210 / 3 |
| 2 (1086) | 30 | 300 / 6 | missing: levels 30 to 39 keep rank 1 (210 / 3) |
| 3 (11733) | 40 | 390 / 9 | 390 / 9 |
| 4 (11734) | 50 | 480 / 12 | 480 / 12 |
| 5 (11735) | 60 | 570 / 15 | 570 / 15 |

The only discrepancy is the omitted rank 2 (`demonArmorLearnLevels` and
`demonArmorRanks` skip 1086). The client also states a third effect per rank
(aura 161, 7/9/11/13/15) the engine does not model.

## Deliberately left out

- Inner Fire: armor only, no DPS effect.
- Lightning Shield: reactive damage, modelled by the shaman's own sim.
- Power Word: Fortitude: stamina, no DPS effect.
- Divine Spirit: a Discipline talent, not a class kit buff.
- Improved Mark of the Wild (see above) and any other rank tables.
