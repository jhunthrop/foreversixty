# mage-arcane rotation ladder

Rules this ladder runs under (Phase 1a,
docs/superpowers/specs/2026-09-28-rotation-accuracy-program-design.md):

- Levels: 10, 20, 30, 38, 40, 50, 60. 300 iterations, seed 1, the
  default encounter (a stationary target three levels above the
  character).
- Talents: the guide's level-60 FS1 build code
  (web/src/content/guides/<class>/<spec>.md), truncated to level-9
  points (0 below level 10). Each talent's guide-assigned rank is read
  by the talent's own stable id against the client build the guide
  names, then re-resolved to that talent's row in the ACTIVE build
  (data/builds/<active>/talents/<class>.json) - a talent the active
  build no longer carries is dropped rather than misaligning every
  digit after it. Points are spent walking the active build's trees
  top row down, the spec's own tree first, then the other two in the
  build code's own order (0, 1, 2, skipping the spec's own). A row the
  budget runs out before reaching is simply left at 0. If the guide
  build itself spends fewer than level-9 points, the remainder is left
  unspent.
- Gear: main_hand always, off_hand for the classes that dual-wield in
  this build (rogue, warrior-fury, shaman-enhancement, hunter),
  ranged for hunter only. Each slot picks the highest item_level
  weapon (speed > 0, so a shield never fills an off hand; damage_max >
  0, so an unfinished/placeholder weapon row this build's item table
  still carries - equipping one hangs the engine mid-sim rather than
  simulating a zero-damage weapon, see the report - is never picked)
  with required_level <= the character's level, from
  data/builds/<active>/items/<class>.json filtered to
  data/builds/<active>/simitems.json's known ids, honoring the spec's
  handedness (warrior-arms and paladin-retribution two-hand only;
  warrior-fury, rogue and shaman-enhancement one-hand only for both
  hands). druid-feral picks no weapon at all. Every other slot is
  bare. Buffs and consumables: none.
- DPS regression: each level's DPS is compared against the ladder's own
  PREVIOUS rung (not literally level-10, since the ladder's own gaps
  are uneven - 30 to 38 is 8 levels, 38 to 40 is 2). A level scoring
  lower than the rung before it is a violation.
- Unresolved: an id the engine's ComputeStats warns it cannot resolve.
  Expected when data/curated/apl/<spec>.json's own inert array names
  it; otherwise a violation.
- Zero casts: one of the curated rotation's own castSpell lines, resolved
  to the id sim/internal/spellranks.HighestLearnedSpellID says the
  engine's OWN rank rewrite actually casts at this level (not a second,
  approximate copy of that resolution - this ladder calls the same
  function sim/request's rewriteRotationRanks calls), that never fires
  in the run - unless the engine could not resolve that id at all
  (already counted as unresolved) or data/curated/apl/<spec>.json's
  expected_idle array names the LINE'S AUTHORED id with a reason.
  HighestLearnedSpellID returning "not learned" means the engine's own
  rewrite already dropped the line before the request was built, which
  is not a violation to report twice.
- Learned but unused (informational, not a violation): every damage
  ability (spellconst effect 2, 6 with aura 3, or 121/31/58 - see
  ladder.go's isDamageEffect) the class has learned by this level that
  this level's cast set never touched, regardless of whether the
  rotation names it at all. An ability whose spellranks.json rows are
  ALL rank 0 (Bloodrage, Judgement - a single always-known ability, not
  a rank progression) is not tracked by this rule at all, the same way
  it is invisible to the engine's own rank rewrite. Heroic
  Strike-shaped bonus-weapon-damage effects (type 17) are NOT covered
  by the damage-effect rule and so never appear here even when
  genuinely unused - see the report for specs where that matters.
- Strict failures (FOREVER_LADDER_STRICT=1) are the four rules above;
  see the run's own "Violations" section below for what this file's own
  run found.

## Ladder

| Level | Talents | Gear | DPS | Distinct casts | Top casts | Unresolved |
|---|---|---|---|---|---|---|
| 10 | 100000000000000000-00000000000000000-0000000000000000000 | main_hand:274271 | 7.9 | 1 | spell:5143/1=67.5 | {SpellID: 12042}, {SpellID: 12043}, {SpellID: 400589} |
| 20 | 254000000000000000-00000000000000000-0000000000000000000 | main_hand:274271 | 12.8 | 1 | spell:5144/1=72.3 | {SpellID: 12042}, {SpellID: 12043}, {SpellID: 400574}, {SpellID: 400589} |
| 30 | 255225000000000000-00000000000000000-0000000000000000000 | main_hand:274271 | 22.9 | 1 | spell:5145/1=86.1, item:5514=1.0, other:mana_gain=1.0 | {SpellID: 12042}, {SpellID: 12043}, {SpellID: 1239696}, {SpellID: 400589} |
| 38 | 255225200000011400-00000000000000000-0000000000000000000 | main_hand:274271 | 34.8 | 2 | spell:8416/1=85.6, other:mana_gain=2.0, spell:12043=1.3, item:5513=1.0, item:5514=1.0 | {SpellID: 12042}, {SpellID: 1239696} |
| 40 | 255225200000011501-00000000000000000-0000000000000000000 | main_hand:274271 | 38.6 | 3 | spell:8417/1=66.0, other:mana_gain=2.0, spell:12042=1.5, spell:12043=1.3, item:5513=1.0 | {SpellID: 1239696} |
| 50 | 255225200000011501-23050000000000000-0000000000000000000 | main_hand:12061 | 52.2 | 3 | spell:10211/1=67.1, other:mana_gain=2.0, spell:12042=1.5, spell:12043=1.3, item:5513=1.0 | {SpellID: 1239699} |
| 60 | 255225200000011501-23050000000000000-0000000000000000000 | main_hand:23577 | 65.6 | 3 | spell:10212/1=66.1, other:mana_gain=2.0, spell:12042=1.5, spell:12043=1.3, item:5513=1.0 | {SpellID: 400574} |

## Learned but unused (informational)


### Level 10

- Arcane Missile (spell 7268)
- Copy of Frostbolt (spell 29163)
- Fire Blast (spell 2136)
- Fireball (spell 143)
- Frost Nova (spell 122)
- Frostbolt (spell 205)

### Level 20

- Arcane Blast (spell 400574)
- Arcane Explosion (spell 1449)
- Arcane Missile (spell 7268)
- Copy of Frostbolt (spell 29163)
- Fire Blast (spell 2137)
- Fireball (spell 3140)
- Flamestrike (spell 2120)
- Frost Nova (spell 122)
- Frostbolt (spell 7322)
- Ice Lance (spell 1312002)
- Pyroblast (spell 11366)

### Level 30

- Arcane Blast (spell 1239696)
- Arcane Explosion (spell 8438)
- Arcane Missile (spell 7268)
- Blast Wave (spell 11113)
- Cone of Cold (spell 120)
- Copy of Frostbolt (spell 29163)
- Fire Blast (spell 8412)
- Fireball (spell 8401)
- Flamestrike (spell 2121)
- Frost Nova (spell 865)
- Frostbolt (spell 8406)
- Ice Lance (spell 400640)
- Pyroblast (spell 12522)
- Scorch (spell 8444)

### Level 38

- Arcane Blast (spell 1239696)
- Arcane Explosion (spell 8439)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13018)
- Cone of Cold (spell 8492)
- Copy of Frostbolt (spell 29163)
- Fire Blast (spell 8413)
- Fireball (spell 8402)
- Flamestrike (spell 8422)
- Frost Nova (spell 865)
- Frostbolt (spell 8408)
- Ice Lance (spell 1240044)
- Pyroblast (spell 12523)
- Scorch (spell 8445)

### Level 40

- Arcane Blast (spell 1239696)
- Arcane Explosion (spell 8439)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13018)
- Cone of Cold (spell 8492)
- Copy of Frostbolt (spell 29163)
- Fire Blast (spell 8413)
- Fireball (spell 8402)
- Flamestrike (spell 8423)
- Frost Nova (spell 6131)
- Frostbolt (spell 8408)
- Frostfire Bolt (spell 401502)
- Ice Lance (spell 1240044)
- Pyroblast (spell 12523)
- Scorch (spell 8446)

### Level 50

- Arcane Blast (spell 1239699)
- Arcane Explosion (spell 10201)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13019)
- Cone of Cold (spell 10160)
- Copy of Frostbolt (spell 29163)
- Fire Blast (spell 10197)
- Fireball (spell 10149)
- Flamestrike (spell 10215)
- Frost Nova (spell 6131)
- Frostbolt (spell 10180)
- Frostfire Bolt (spell 1237312)
- Ice Lance (spell 1240046)
- Pyroblast (spell 12525)
- Scorch (spell 10205)

### Level 60

- Arcane Blast (spell 1239700)
- Arcane Explosion (spell 10202)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13021)
- Cone of Cold (spell 10161)
- Copy of Frostbolt (spell 29163)
- Debug Frost Spell (spell 29607)
- Fire Blast (spell 10199)
- Fireball (spell 25306)
- Flamestrike (spell 10216)
- Frost Nova (spell 10230)
- Frostbolt (spell 25304)
- Frostfire Bolt (spell 1237313)
- Ice Lance (spell 1240047)
- Pyroblast (spell 18809)
- Scorch (spell 10207)

## Violations found in this run

- mage-arcane level=10 kind=unresolved_id action={SpellID: 12042}
- mage-arcane level=10 kind=unresolved_id action={SpellID: 12043}
- mage-arcane level=10 kind=unresolved_id action={SpellID: 400589}
- mage-arcane level=20 kind=unresolved_id action={SpellID: 12042}
- mage-arcane level=20 kind=unresolved_id action={SpellID: 12043}
- mage-arcane level=20 kind=unresolved_id action={SpellID: 400574}
- mage-arcane level=20 kind=unresolved_id action={SpellID: 400589}
- mage-arcane level=30 kind=unresolved_id action={SpellID: 12042}
- mage-arcane level=30 kind=unresolved_id action={SpellID: 12043}
- mage-arcane level=30 kind=unresolved_id action={SpellID: 1239696}
- mage-arcane level=30 kind=unresolved_id action={SpellID: 400589}
- mage-arcane level=38 kind=unresolved_id action={SpellID: 12042}
- mage-arcane level=38 kind=unresolved_id action={SpellID: 1239696}
- mage-arcane level=40 kind=unresolved_id action={SpellID: 1239696}
- mage-arcane level=50 kind=unresolved_id action={SpellID: 1239699}
- mage-arcane level=60 kind=unresolved_id action={SpellID: 400574}
- mage-arcane level=60 kind=zero_casts spell="Arcane Blast" id=1239700 authored=400574
- mage-arcane level=60 kind=zero_casts spell="Arcane Missiles" id=25345 authored=10212
