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
| 10 | 00000000000000000-100000000000000000-000000000000000000 | main_hand:24071 off_hand:13371 | 0.1 | 3 | other:rage_gain=38.1, spell:2687=3.5, spell:1680=2.0, spell:6673=1.5 | {SpellID: 12328}, {SpellID: 23894} |
| 20 | 00000000000000000-353000000000000000-000000000000000000 | main_hand:24071 off_hand:13371 | 0.2 | 3 | other:rage_gain=38.1, spell:2687=3.5, spell:1680=2.0, spell:5242=1.5 | {SpellID: 12328}, {SpellID: 23894} |
| 30 | 00000000000000000-353211005010000000-000000000000000000 | main_hand:24071 off_hand:13371 | 0.5 | 4 | other:rage_gain=38.1, spell:2687=3.5, spell:1680=1.6, spell:6192=1.6, spell:5308=0.4 | {SpellID: 12328}, {SpellID: 23894} |
| 38 | 00000000000000000-353211005050010030-000000000000000000 | main_hand:24071 off_hand:13371 | 1.1 | 5 | other:rage_gain=38.1, spell:2687=3.5, spell:12328=1.5, spell:11549=1.4, spell:20658=1.2 | {SpellID: 23894} |
| 40 | 00000000000000000-353211005050010050-000000000000000000 | main_hand:24071 off_hand:13371 | 1.7 | 5 | other:rage_gain=38.1, spell:2687=3.5, spell:12328=1.5, spell:11549=1.4, spell:20660=1.2 | {SpellID: 23881} |
| 50 | 35100000000000000-353211005050010051-000000000000000000 | main_hand:24071 off_hand:13371 | 1.4 | 6 | other:rage_gain=38.1, spell:2687=3.5, spell:11550=1.7, spell:12328=1.5, spell:1719=1.0 | - |
| 60 | 35311103002000000-353211005050010051-000000000000000000 | main_hand:20279 off_hand:234558 | 135.5 | 9 | other:attack/1=97.1, other:attack/2=67.9, other:rage_gain=66.0, spell:23894=15.2, spell:20662=10.3 | - |

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
- Revenge (spell 25288)
- Shield Bash (spell 1672)
- Shield Slam (spell 23925)
- Test Strike W35 (spell 23850)
- Test Strike W50 (spell 23848)
- Thunder Clap (spell 11581)

## Violations found in this run

- warrior-fury level=10 kind=zero_casts spell="Heroic Strike" id=284 authored=25286
- warrior-fury level=20 kind=zero_casts spell="Heroic Strike" id=285 authored=25286
- warrior-fury level=30 kind=zero_casts spell="Heroic Strike" id=1608 authored=25286
- warrior-fury level=38 kind=zero_casts spell="Heroic Strike" id=11564 authored=25286
- warrior-fury level=40 kind=zero_casts spell="Heroic Strike" id=11565 authored=25286
- warrior-fury level=50 kind=dps_regression dps=1.4 prev_dps=1.7
- warrior-fury level=50 kind=zero_casts id=1680 authored=1680 (untracked ability; not in spellranks.json's rank chains)
- warrior-fury level=50 kind=zero_casts spell="Heroic Strike" id=11566 authored=25286
