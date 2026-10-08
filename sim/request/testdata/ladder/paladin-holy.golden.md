# paladin-holy rotation ladder

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
| 10 | 01000000000000000-0000000000000000-00000000000000000 | main_hand:277247 | 0.0 | 0 | - | {SpellID: 20216}, {SpellID: 25292}, {SpellID: 25890} |
| 20 | 05320001000000000-0000000000000000-00000000000000000 | main_hand:6953 | 9.6 | 1 | spell:19750=54.0 | {SpellID: 20216}, {SpellID: 25292}, {SpellID: 25890} |
| 30 | 05320003224000000-0000000000000000-00000000000000000 | main_hand:267369 | 16.3 | 1 | spell:19939=59.0, other:mana_gain=3.4 | {SpellID: 1311606}, {SpellID: 20216}, {SpellID: 25292}, {SpellID: 25890} |
| 38 | 05320003225111040-0000000000000000-00000000000000000 | main_hand:267369 | 22.8 | 2 | spell:19940=52.5, other:mana_gain=8.6, spell:20216=3.0 | {SpellID: 25292}, {SpellID: 25890} |
| 40 | 05320003225111051-0000000000000000-00000000000000000 | main_hand:7723 | 24.6 | 4 | spell:19940=55.7, other:mana_gain=9.8, spell:20216=3.0, spell:20473=0.0, spell:25914=0.0 | {SpellID: 25292}, {SpellID: 25890} |
| 50 | 05320003225111051-5500000000000000-00000000000000000 | main_hand:7723 | 38.9 | 2 | spell:19942=42.8, other:mana_gain=7.8, spell:20216=2.9 | {SpellID: 25292}, {SpellID: 25890} |
| 60 | 05320003225111051-5532500000000000-00000000000000000 | main_hand:23455 | 62.6 | 3 | spell:19943=43.1, other:mana_gain=7.6, spell:20216=2.0, spell:25890=1.0 | - |

## Learned but unused (informational)


### Level 10

- Holy Strike (spell 679)
- Judgement of Fury (spell 1311650)
- Judgement of Righteousness (spell 20280)

### Level 20

- Consecration (spell 26573)
- Exorcism (spell 879)
- Holy Strike (spell 1866)
- Judgement of Command (spell 20425)
- Judgement of Fury (spell 1311655)
- Judgement of Righteousness (spell 20281)

### Level 30

- Consecration (spell 20116)
- Exorcism (spell 5614)
- Holy Strike (spell 680)
- Judgement of Command (spell 20962)
- Judgement of Fury (spell 20183)
- Judgement of Righteousness (spell 20282)

### Level 38

- Consecration (spell 20116)
- Exorcism (spell 5615)
- Holy Strike (spell 2495)
- Judgement of Command (spell 20962)
- Judgement of Fury (spell 20411)
- Judgement of Righteousness (spell 20283)

### Level 40

- Consecration (spell 20922)
- Exorcism (spell 5615)
- Holy Strike (spell 2495)
- Judgement of Command (spell 20961)
- Judgement of Fury (spell 20411)
- Judgement of Righteousness (spell 20283)

### Level 50

- Consecration (spell 20923)
- Exorcism (spell 10312)
- Hammer of Wrath (spell 24275)
- Holy Strike (spell 5569)
- Holy Wrath (spell 2812)
- Judgement of Command (spell 20965)
- Judgement of Fury (spell 20413)
- Judgement of Righteousness (spell 20285)

### Level 60

- Consecration (spell 20924)
- Exorcism (spell 10314)
- Hammer of Wrath (spell 24239)
- Holy Strike (spell 10333)
- Holy Wrath (spell 10318)
- Judgement of Command (spell 20966)
- Judgement of Fury (spell 20414)
- Judgement of Righteousness (spell 20286)

## Violations found in this run

- paladin-holy level=10 kind=unresolved_id action={SpellID: 25292}
- paladin-holy level=10 kind=unresolved_id action={SpellID: 25890}
- paladin-holy level=20 kind=unresolved_id action={SpellID: 25292}
- paladin-holy level=20 kind=unresolved_id action={SpellID: 25890}
- paladin-holy level=30 kind=unresolved_id action={SpellID: 25292}
- paladin-holy level=30 kind=unresolved_id action={SpellID: 25890}
- paladin-holy level=38 kind=unresolved_id action={SpellID: 25292}
- paladin-holy level=38 kind=unresolved_id action={SpellID: 25890}
- paladin-holy level=38 kind=zero_casts spell="Holy Shock" id=1311606 authored=20930
- paladin-holy level=40 kind=unresolved_id action={SpellID: 25292}
- paladin-holy level=40 kind=unresolved_id action={SpellID: 25890}
- paladin-holy level=50 kind=unresolved_id action={SpellID: 25292}
- paladin-holy level=50 kind=unresolved_id action={SpellID: 25890}
- paladin-holy level=50 kind=zero_casts spell="Holy Shock" id=20929 authored=20930
- paladin-holy level=60 kind=zero_casts id=25292 authored=25292 (untracked ability; not in spellranks.json's rank chains)
- paladin-holy level=60 kind=zero_casts spell="Holy Shock" id=20930 authored=20930
