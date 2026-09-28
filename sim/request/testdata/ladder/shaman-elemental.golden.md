# shaman-elemental rotation ladder

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
| 10 | 1000000000000000-000000000000000000-0000000000000000 | main_hand:22348 | 23.8 | 5 | spell:3606=61.7, spell:529=28.4, other:attack/1=27.8, spell:3599=5.6, spell:20572=2.0 | - |
| 20 | 4520000000000000-000000000000000000-0000000000000000 | main_hand:22348 | 30.1 | 6 | spell:6350=65.7, other:attack/1=27.2, spell:915=15.4, spell:6363=5.2, spell:8052=3.5 | - |
| 30 | 4532310300000000-000000000000000000-0000000000000000 | main_hand:22348 | 40.3 | 6 | spell:6351=65.1, other:attack/1=30.0, spell:943=16.2, spell:6364=4.5, spell:8053=2.0 | - |
| 38 | 4532310300103031-000000000000000000-0000000000000000 | main_hand:22348 | 53.0 | 6 | spell:6351=67.2, other:attack/1=42.0, spell:8053=10.3, spell:10391=6.0, spell:6364=4.6 | - |
| 40 | 4532310300103031-000000000000000000-2000000000000000 | main_hand:22348 | 51.6 | 6 | spell:6352=60.1, other:attack/1=32.0, spell:10391=13.7, spell:6365=3.7, spell:10447=2.0 | - |
| 50 | 4532310300103031-000000000000000000-5520000000000000 | main_hand:22348 | 66.2 | 6 | spell:10435=68.4, other:attack/1=36.4, spell:15207=9.9, spell:10447=4.5, spell:10437=3.8 | - |
| 60 | 4532310300103031-000000000000000000-5533220000000000 | main_hand:22589 | 111.9 | 6 | spell:10436=70.0, other:attack/1=42.6, spell:15208=14.3, spell:10438=3.7, spell:29228=2.0 | - |

## Learned but unused (informational)


### Level 10

- Earth Shock (spell 8044)
- Flametongue Attack (spell 10444)
- Stormstrike (spell 410156)

### Level 20

- Flametongue Attack (spell 10444)
- Frost Shock (spell 8056)
- Frostbrand Attack (spell 8034)
- Stormstrike (spell 410156)

### Level 30

- Flametongue Attack (spell 10444)
- Frost Shock (spell 8056)
- Frostbrand Attack (spell 8037)
- Stormstrike (spell 410156)

### Level 38

- Chain Lightning (spell 421)
- Flametongue Attack (spell 10444)
- Frost Shock (spell 8058)
- Frostbrand Attack (spell 10458)
- Stormstrike (spell 410156)

### Level 40

- Chain Lightning (spell 930)
- Flametongue Attack (spell 10444)
- Frost Shock (spell 8058)
- Frostbrand Attack (spell 10458)
- Lava Burst (spell 408490)
- Stormstrike (spell 410156)

### Level 50

- Chain Lightning (spell 2860)
- Flametongue Attack (spell 10444)
- Frost Shock (spell 10472)
- Frostbrand Attack (spell 16352)
- Lava Burst (spell 1238299)
- Stormstrike (spell 410156)

### Level 60

- Chain Lightning (spell 10605)
- Flametongue Attack (spell 10444)
- Frost Shock (spell 10473)
- Frostbrand Attack (spell 16353)
- Lava Burst (spell 1238300)
- Stormstrike (spell 410156)

## Violations found in this run

- shaman-elemental level=10 kind=zero_casts spell="Earth Shock" id=8044 authored=10414
- shaman-elemental level=40 kind=dps_regression dps=51.6 prev_dps=53.0
