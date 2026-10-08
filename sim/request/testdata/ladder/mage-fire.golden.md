# mage-fire rotation ladder

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
  hands). druid-feral picks no weapon at all. Every caster spec
  (priest-shadow, mage's three specs, warlock's three specs) fills
  ranged with a wand instead: the same item table's ranged rows whose
  icon names them a real wand (every wand row's own damage_max is 0 in
  this build, unlike a bow or gun, so the melee pick's damage_max > 0
  check is replaced by that icon check rather than dropped), so
  OtherActionShoot/wand lines have something to resolve against.
  Every other slot is bare. Consumables: none (see the potion rule
  below). Buffs: only the class self-buff kit (ladderKitBuffs: a mage's
  Arcane Intellect, a druid's Mark of the Wild, a paladin's Blessing of
  Might from level 4), at the highest rank the level can learn; no
  raid buffs.
- DPS regression: each level's DPS is compared against the ladder's own
  PREVIOUS rung (not literally level-10, since the ladder's own gaps
  are uneven - 30 to 38 is 8 levels, 38 to 40 is 2), tolerating up to a
  1% drop as the 300-iteration run's own noise (shaman-elemental's
  level 40, 46.0 vs a level-38 46.1, is exactly this). A level scoring
  more than 1% lower than the rung before it is a violation.
- Unresolved: an id the engine's ComputeStats warns it cannot resolve,
  with three standing exceptions before anything counts as a
  violation: (1) data/curated/apl/<spec>.json's own inert array names
  it; (2) it is the potion action ({OtherID: 13}) - the ladder
  character carries no consumes (a rogue's poisons, class kit from
  level 20, are the one exception: ladderKitConsumes), so this can
  never resolve, at any level, any spec; (3) it is a talent-granted spell
  (data/builds/<build>/talents/<class>.json's own "ranks[].spell_id")
  and the ladder's own truncated build (ladderTalentString's budget
  walk) has spent zero points on that talent at this level - expected
  right up until the level this ladder's approximation of the guide's
  build actually reaches that talent's row, a violation only once the
  build HAS spent points on it and the id still will not resolve.
  Anything else is a violation.
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
| 10 | 000000000000000000-10000000000000000-0000000000000000000 | main_hand:263937 ranged:263430 | 7.9 | 3 | spell:5019=96.8, spell:2136=19.5, spell:143=10.7 | {SpellID: 11129}, {SpellID: 12873} |
| 20 | 000000000000000000-23510000000000000-0000000000000000000 | main_hand:890 ranged:5243 | 18.5 | 3 | spell:5019=91.9, spell:2137=17.7, spell:3140=8.0 | {SpellID: 11129}, {SpellID: 11366}, {SpellID: 12873}, {SpellID: 400625} |
| 30 | 000000000000000000-23552100120000000-0000000000000000000 | main_hand:249392 ranged:5213 | 25.2 | 3 | spell:5019=93.9, spell:8444=35.7, spell:8412=4.5, item:5514=1.0, other:mana_gain=1.0 | {SpellID: 11129}, {SpellID: 400625} |
| 38 | 000000000000000000-23552100130103030-0000000000000000000 | main_hand:1664 ranged:13064 | 41.0 | 4 | spell:5019=74.8, spell:8445=36.1, spell:8413=5.1, other:mana_gain=2.0, item:5513=1.0 | {SpellID: 11129} |
| 40 | 000000000000000000-23552100130103041-0000000000000000000 | main_hand:1664 ranged:5216 | 45.0 | 5 | spell:5019=86.4, spell:8446=29.8, spell:8413=4.2, other:mana_gain=2.0, spell:11129=1.3 | - |
| 50 | 205011000000000000-23552100130103051-0000000000000000000 | main_hand:812 ranged:249232 | 75.4 | 5 | spell:5019=57.4, spell:10205=37.7, spell:10197=7.6, other:mana_gain=2.0, spell:12525=1.4 | - |
| 60 | 205015100000000000-23552100130103051-0050000000000000000 | main_hand:22589 ranged:22821 | 251.1 | 5 | spell:10207=61.5, spell:5019=41.7, spell:10199=18.8, spell:18809=3.8, other:mana_gain=2.0 | - |

## Learned but unused (informational)


### Level 10

- Arcane Missile (spell 7268)
- Frost Nova (spell 122)
- Frostbolt (spell 205)

### Level 20

- Arcane Blast (spell 400574)
- Arcane Explosion (spell 1449)
- Arcane Missile (spell 7268)
- Blizzard (spell 10)
- Chill (spell 1308651)
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
- Blizzard (spell 6141)
- Chill (spell 1308651)
- Cone of Cold (spell 120)
- Fireball (spell 8401)
- Flamestrike (spell 2121)
- Frost Nova (spell 865)
- Frostbolt (spell 8406)
- Ice Lance (spell 400640)
- Pyroblast (spell 12522)

### Level 38

- Arcane Blast (spell 1239696)
- Arcane Explosion (spell 8439)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13018)
- Blizzard (spell 8427)
- Chill (spell 1308651)
- Cone of Cold (spell 8492)
- Fireball (spell 8402)
- Flamestrike (spell 8422)
- Frost Nova (spell 865)
- Frostbolt (spell 8408)
- Ice Lance (spell 1240044)

### Level 40

- Arcane Blast (spell 1239696)
- Arcane Explosion (spell 8439)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13018)
- Blizzard (spell 8427)
- Chill (spell 1308651)
- Cone of Cold (spell 8492)
- Fireball (spell 8402)
- Flamestrike (spell 8423)
- Frost Nova (spell 6131)
- Frostbolt (spell 8408)
- Frostfire Bolt (spell 401502)
- Ice Lance (spell 1240044)

### Level 50

- Arcane Blast (spell 1239699)
- Arcane Explosion (spell 10201)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13019)
- Blizzard (spell 10185)
- Chill (spell 1308651)
- Cone of Cold (spell 10160)
- Fireball (spell 10149)
- Flamestrike (spell 10215)
- Frost Nova (spell 6131)
- Frostbolt (spell 10180)
- Frostfire Bolt (spell 1237312)
- Ice Lance (spell 1240046)

### Level 60

- Arcane Blast (spell 1239700)
- Arcane Explosion (spell 10202)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13021)
- Blizzard (spell 10187)
- Chill (spell 1308651)
- Cone of Cold (spell 10161)
- Debug Frost Spell (spell 29607)
- Fireball (spell 10151)
- Flamestrike (spell 10216)
- Frost Nova (spell 10230)
- Frostbolt (spell 10181)
- Frostfire Bolt (spell 1237313)
- Ice Lance (spell 1240047)

## Violations found in this run

None.
