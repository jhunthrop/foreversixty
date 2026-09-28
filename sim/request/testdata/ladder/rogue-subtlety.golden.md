# rogue-subtlety rotation ladder

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
| 10 | 00000000000000000-00000000000000000-1000000000000000000 | main_hand:22377 off_hand:11863 | 39.3 | 3 | other:attack/2=138.5, other:attack/1=107.8, spell:1757=41.7, spell:6760/5=4.3, spell:5171/5=3.0 | {SpellID: 14183}, {SpellID: 14278}, {SpellID: 16511} |
| 20 | 00000000000000000-00000000000000000-5321000000000000000 | main_hand:22377 off_hand:11863 | 40.6 | 3 | other:attack/2=138.3, other:attack/1=107.7, spell:1758=42.0, spell:6761/5=4.4, spell:5171/5=2.9 | {SpellID: 14183}, {SpellID: 14278}, {SpellID: 16511} |
| 30 | 00000000000000000-00000000000000000-5322210310011000000 | main_hand:22377 off_hand:11863 | 44.0 | 4 | other:attack/2=140.8, other:attack/1=109.6, spell:1760=34.1, spell:14278=9.0, spell:5171/5=3.9 | {SpellID: 16511} |
| 38 | 00000000000000000-00000000000000000-5322210310013011040 | main_hand:22377 off_hand:11863 | 55.4 | 4 | other:attack/2=142.7, other:attack/1=111.1, spell:16511=42.7, spell:14278=8.8, spell:5171/5=4.5 | - |
| 40 | 00000000000000000-00000000000000000-5322210310013011051 | main_hand:22377 off_hand:11863 | 59.3 | 4 | other:attack/2=142.7, other:attack/1=111.1, spell:16511=42.7, spell:14278=8.8, spell:5171/5=4.5 | - |
| 50 | 00500000000000000-32000000000000000-5322210310013011051 | main_hand:22377 off_hand:11863 | 72.3 | 4 | other:attack/2=149.3, other:attack/1=116.2, spell:16511=42.5, spell:14278=8.8, spell:6774/5=4.6 | - |
| 60 | 00500000000000000-32513100000000000-5322210310013011051 | main_hand:21126 off_hand:18847 | 104.3 | 4 | other:attack/1=116.4, other:attack/2=72.4, spell:16511=42.5, spell:14278=8.7, spell:6774/5=4.6 | - |

## Learned but unused (informational)


### Level 10

- Backstab (spell 53)
- Gouge (spell 1776)
- Serrated Blades (spell 461327)

### Level 20

- Ambush (spell 8676)
- Backstab (spell 2590)
- Garrote (spell 703)
- Gouge (spell 1777)
- Kick (spell 1766)
- Rupture (spell 1943)
- Serrated Blades (spell 461327)

### Level 30

- Ambush (spell 8724)
- Backstab (spell 2591)
- Garrote (spell 8632)
- Gouge (spell 1777)
- Kick (spell 1767)
- Rupture (spell 8639)
- Serrated Blades (spell 461327)

### Level 38

- Ambush (spell 8725)
- Backstab (spell 8721)
- Garrote (spell 8633)
- Gouge (spell 8629)
- Kick (spell 1767)
- Rupture (spell 8640)
- Serrated Blades (spell 461327)
- Sinister Strike (spell 8621)

### Level 40

- Ambush (spell 8725)
- Backstab (spell 8721)
- Garrote (spell 8633)
- Gouge (spell 8629)
- Kick (spell 1767)
- Rupture (spell 8640)
- Serrated Blades (spell 461327)
- Sinister Strike (spell 8621)

### Level 50

- Ambush (spell 11268)
- Backstab (spell 11279)
- Garrote (spell 11289)
- Gouge (spell 11285)
- Kick (spell 1768)
- Rupture (spell 11273)
- Serrated Blades (spell 461327)
- Sinister Strike (spell 11293)

### Level 60

- Ambush (spell 11269)
- Backstab (spell 25300)
- Garrote (spell 11290)
- Gouge (spell 11286)
- Kick (spell 1769)
- Rupture (spell 11275)
- Serrated Blades (spell 461327)
- Sinister Strike (spell 11294)
- Test Stab R50 (spell 23959)
- Test Strike R50 (spell 23960)

## Violations found in this run

None.
