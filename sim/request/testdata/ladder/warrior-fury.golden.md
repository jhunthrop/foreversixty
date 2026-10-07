# warrior-fury rotation ladder

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
| 10 | 00000000000000000-10000000000000000-000000000000000000 | main_hand:1927 off_hand:6969 | 11.7 | 5 | other:attack/2=107.1, other:attack/1=56.2, other:rage_gain=38.1, spell:1680=18.1, spell:284/1=14.2 | {SpellID: 12328}, {SpellID: 23894} |
| 20 | 00000000000000000-35030000000000000-000000000000000000 | main_hand:1482 off_hand:2236 | 25.9 | 5 | other:attack/2=121.3, other:rage_gain=87.1, other:attack/1=49.4, spell:285/1=18.5, spell:285=18.2 | {SpellID: 12328}, {SpellID: 23894} |
| 30 | 00000000000000000-35051105010000000-000000000000000000 | main_hand:6692 off_hand:9457 | 45.2 | 6 | other:rage_gain=105.1, other:attack/2=79.3, other:attack/1=55.9, spell:1608/1=20.4, spell:1608=20.2 | {SpellID: 12328}, {SpellID: 23894} |
| 38 | 00000000000000000-35051105050010300-000000000000000000 | main_hand:868 off_hand:6829 | 67.2 | 7 | other:rage_gain=115.4, other:attack/2=85.6, other:attack/1=66.6, spell:11564/1=23.3, spell:11564=23.1 | {SpellID: 23894} |
| 40 | 00000000000000000-35051105050010500-000000000000000000 | main_hand:2164 off_hand:9359 | 75.0 | 8 | other:rage_gain=120.0, other:attack/2=83.6, other:attack/1=82.5, spell:11565/1=24.5, spell:11565=24.3 | {SpellID: 23881} |
| 50 | 35100000000000000-35051105050010501-000000000000000000 | main_hand:810 off_hand:2163 | 107.2 | 9 | other:attack/2=139.2, other:rage_gain=135.2, other:attack/1=86.1, spell:23892=19.0, spell:20661=15.8 | - |
| 60 | 35311103002000000-35051105050010501-000000000000000000 | main_hand:22736 off_hand:23054 | 251.0 | 9 | other:rage_gain=101.7, other:attack/2=73.0, other:attack/1=63.9, spell:23894=19.2, spell:20662=14.6 | - |

## Learned but unused (informational)


### Level 10

- Hamstring (spell 1715)
- Rend (spell 6546)
- Thunder Clap (spell 6343)

### Level 20

- Hamstring (spell 1715)
- Mocking Blow (spell 694)
- Overpower (spell 7384)
- Rend (spell 6547)
- Revenge (spell 6572)
- Shield Bash (spell 72)
- Thunder Clap (spell 8198)

### Level 30

- Hamstring (spell 1715)
- Mocking Blow (spell 7400)
- Overpower (spell 7887)
- Rend (spell 6548)
- Revenge (spell 6574)
- Shield Bash (spell 72)
- Thunder Clap (spell 8204)

### Level 38

- Hamstring (spell 7372)
- Mocking Blow (spell 7402)
- Overpower (spell 7887)
- Pummel (spell 6552)
- Rend (spell 6548)
- Revenge (spell 7379)
- Shield Bash (spell 1671)
- Thunder Clap (spell 8205)

### Level 40

- Bloodthirst (spell 23881)
- Hamstring (spell 7372)
- Mocking Blow (spell 7402)
- Mortal Strike (spell 12294)
- Overpower (spell 7887)
- Pummel (spell 6552)
- Rend (spell 11572)
- Revenge (spell 7379)
- Shield Bash (spell 1671)
- Shield Slam (spell 23922)
- Thunder Clap (spell 8205)

### Level 50

- Devastate (spell 20243)
- Hamstring (spell 7372)
- Mocking Blow (spell 20559)
- Mortal Strike (spell 21551)
- Overpower (spell 11584)
- Pummel (spell 6552)
- Rend (spell 11573)
- Revenge (spell 11600)
- Shield Bash (spell 1671)
- Shield Slam (spell 23923)
- Thunder Clap (spell 11580)

### Level 60

- Devastate (spell 20243)
- Hamstring (spell 7373)
- Mocking Blow (spell 20560)
- Mortal Strike (spell 21553)
- Overpower (spell 11585)
- Pummel (spell 6554)
- Recycle (spell 458882)
- Rend (spell 11574)
- Revenge (spell 11601)
- Shield Bash (spell 1672)
- Shield Slam (spell 23925)
- Test Strike W35 (spell 23850)
- Test Strike W50 (spell 23848)
- Thunder Clap (spell 11581)

## Violations found in this run

None.
