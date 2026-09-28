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
  below).
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
  character carries no consumes, so this can never resolve, at any
  level, any spec; (3) it is a talent-granted spell
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
| 10 | 000000000000000000-10000000000000000-0000000000000000000 | main_hand:24071 ranged:19108 | 6.1 | 2 | spell:2136=23.0, spell:143=15.0 | {SpellID: 11129}, {SpellID: 12873} |
| 20 | 000000000000000000-23510000000000000-0000000000000000000 | main_hand:24071 ranged:19108 | 10.9 | 2 | spell:2137=19.0, spell:3140=10.0 | {SpellID: 11129}, {SpellID: 12873} |
| 30 | 000000000000000000-23552110020000000-0000000000000000000 | main_hand:24071 ranged:19108 | 21.6 | 3 | spell:8444=20.7, spell:8401=5.6, spell:8412=4.5, item:5514=1.0, other:mana_gain=1.0 | {SpellID: 11129} |
| 38 | 000000000000000000-23552110030003040-0000000000000000000 | main_hand:24071 ranged:19108 | 36.0 | 3 | spell:8445=17.4, spell:8413=7.5, spell:8402=6.0, other:mana_gain=2.0, item:5513=1.0 | {SpellID: 11129} |
| 40 | 000000000000000000-23552110030003051-0000000000000000000 | main_hand:24071 ranged:19108 | 43.5 | 4 | spell:8446=17.0, spell:8413=6.8, spell:8402=6.1, other:mana_gain=2.0, spell:11129=1.0 | - |
| 50 | 253000000000000000-23552110030003051-0000000000000000000 | main_hand:24071 ranged:19108 | 57.9 | 4 | spell:10205=17.5, spell:10197=7.1, spell:10149=5.0, other:mana_gain=2.0, spell:11129=1.1 | - |
| 60 | 255115100000000000-23552110030003051-0000000000000000000 | main_hand:20279 ranged:20335 | 150.7 | 4 | spell:10207=19.5, spell:25306=17.0, spell:10199=12.1, other:mana_gain=2.0, spell:11129=1.1 | - |

## Learned but unused (informational)


### Level 10

- Arcane Missile (spell 7268)
- Frost Nova (spell 122)
- Frostbolt (spell 205)

### Level 20

- Arcane Blast (spell 400574)
- Arcane Explosion (spell 1449)
- Arcane Missile (spell 7268)
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
- Cone of Cold (spell 8492)
- Flamestrike (spell 8422)
- Frost Nova (spell 865)
- Frostbolt (spell 8408)
- Ice Lance (spell 1240044)
- Pyroblast (spell 12523)

### Level 40

- Arcane Blast (spell 1239696)
- Arcane Explosion (spell 8439)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13018)
- Cone of Cold (spell 8492)
- Flamestrike (spell 8423)
- Frost Nova (spell 6131)
- Frostbolt (spell 8408)
- Frostfire Bolt (spell 401502)
- Ice Lance (spell 1240044)
- Pyroblast (spell 12523)

### Level 50

- Arcane Blast (spell 1239699)
- Arcane Explosion (spell 10201)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13019)
- Cone of Cold (spell 10160)
- Flamestrike (spell 10215)
- Frost Nova (spell 6131)
- Frostbolt (spell 10180)
- Frostfire Bolt (spell 1237312)
- Ice Lance (spell 1240046)
- Pyroblast (spell 12525)

### Level 60

- Arcane Blast (spell 1239700)
- Arcane Explosion (spell 10202)
- Arcane Missile (spell 7268)
- Blast Wave (spell 13021)
- Cone of Cold (spell 10161)
- Debug Frost Spell (spell 29607)
- Flamestrike (spell 10216)
- Frost Nova (spell 10230)
- Frostbolt (spell 25304)
- Frostfire Bolt (spell 1237313)
- Ice Lance (spell 1240047)
- Pyroblast (spell 18809)

## Violations found in this run

- mage-fire level=10 kind=unresolved_id action={SpellID: 12873}
- mage-fire level=20 kind=unresolved_id action={SpellID: 12873}
