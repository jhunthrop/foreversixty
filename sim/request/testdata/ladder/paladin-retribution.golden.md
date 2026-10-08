# paladin-retribution rotation ladder

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
| 10 | 00000000000000000-0000000000000000-01000000000000000 | main_hand:263407 | 14.3 | 5 | other:attack/1=61.0, spell:25740=52.3, spell:20271=18.6, spell:20280=18.6, spell:679=18.6 | - |
| 20 | 00000000000000000-0000000000000000-05024000000000000 | main_hand:6953 | 30.8 | 5 | other:attack/1=57.3, spell:25739=48.9, spell:1866=18.6, spell:20271=18.6, spell:20281=18.6 | - |
| 30 | 00000000000000000-0000000000000000-05025331001100000 | main_hand:13045 | 48.4 | 5 | other:attack/1=63.1, spell:25738=54.0, other:mana_gain=18.6, spell:20271=18.6, spell:20282=18.6 | - |
| 38 | 00000000000000000-0000000000000000-05025331001330300 | main_hand:10758 | 71.6 | 5 | other:attack/1=55.5, spell:25737=47.5, other:mana_gain=18.6, spell:20271=18.6, spell:20283=18.6 | - |
| 40 | 00000000000000000-0000000000000000-05025331001330311 | main_hand:1982 | 81.0 | 5 | other:attack/1=65.3, spell:25737=55.8, other:mana_gain=18.6, spell:20271=18.6, spell:20283=18.6 | - |
| 50 | 52003000000000000-0000000000000000-05025331001330311 | main_hand:2915 | 117.2 | 6 | other:attack/1=79.3, spell:25735=67.7, other:mana_gain=18.6, spell:20271=18.6, spell:20285=18.6 | - |
| 60 | 52003003000000000-0520000000000000-05025331001330311 | main_hand:22798 | 197.6 | 6 | other:attack/1=48.2, spell:25713=41.5, other:mana_gain=18.6, spell:20271=18.6, spell:20286=18.6 | - |

## Learned but unused (informational)


### Level 10

- Judgement of Fury (spell 1311650)

### Level 20

- Consecration (spell 26573)
- Exorcism (spell 879)
- Judgement of Command (spell 20425)
- Judgement of Fury (spell 1311655)

### Level 30

- Consecration (spell 20116)
- Exorcism (spell 5614)
- Judgement of Command (spell 20962)
- Judgement of Fury (spell 20183)

### Level 38

- Consecration (spell 20116)
- Exorcism (spell 5615)
- Judgement of Command (spell 20962)
- Judgement of Fury (spell 20411)

### Level 40

- Consecration (spell 20922)
- Exorcism (spell 5615)
- Judgement of Command (spell 20961)
- Judgement of Fury (spell 20411)

### Level 50

- Consecration (spell 20923)
- Exorcism (spell 10312)
- Holy Wrath (spell 2812)
- Judgement of Command (spell 20965)
- Judgement of Fury (spell 20413)

### Level 60

- Consecration (spell 20924)
- Exorcism (spell 10314)
- Holy Wrath (spell 10318)
- Judgement of Command (spell 20966)
- Judgement of Fury (spell 20414)

## Violations found in this run

None.
