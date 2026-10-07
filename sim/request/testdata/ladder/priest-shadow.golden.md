# priest-shadow rotation ladder

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
| 10 | 000000000000000000-00000000000000000-100000000000000000 | main_hand:263937 ranged:263430 | 9.9 | 3 | spell:5019=96.4, spell:8092=17.7, spell:594=10.8 | {SpellID: 15473} |
| 20 | 000000000000000000-00000000000000000-542000000000000000 | main_hand:890 ranged:5243 | 23.0 | 4 | spell:5019=94.7, spell:8102=13.2, spell:970=8.8, spell:2944=1.9 | {SpellID: 14751}, {SpellID: 15407}, {SpellID: 15473} |
| 30 | 000000000000000000-00000000000000000-543110401200000000 | main_hand:249392 ranged:5213 | 34.9 | 4 | spell:5019=77.4, spell:17311=29.1, spell:992=2.1, spell:19276=1.0 | {SpellID: 14751}, {SpellID: 15473} |
| 38 | 000000000000000000-00000000000000000-543110401201300220 | main_hand:1664 ranged:13064 | 55.9 | 4 | spell:5019=70.1, spell:17312=24.8, spell:2767=2.0, spell:19277=1.0 | {SpellID: 14751}, {SpellID: 15473} |
| 40 | 000000000000000000-00000000000000000-543110401201300240 | main_hand:1664 ranged:5216 | 57.4 | 4 | spell:5019=73.1, spell:17312=25.7, spell:2767=2.0, spell:19277=1.0 | {SpellID: 14751}, {SpellID: 15473} |
| 50 | 521000000000000000-00000000000000000-543110401201300251 | main_hand:812 ranged:249232 | 84.7 | 5 | spell:5019=60.6, spell:17313=21.3, spell:10893=2.0, spell:15473=1.0, spell:19278=1.0 | {SpellID: 14751} |
| 60 | 524111001300000000-00000000000000000-543110401201300251 | main_hand:22799 ranged:22821 | 204.2 | 6 | spell:5019=72.1, spell:18807=26.0, spell:10894=2.1, spell:14751=1.5, spell:19280=1.5 | - |

## Learned but unused (informational)


### Level 10

- Smite (spell 591)
- Starshards (spell 10797)

### Level 20

- Chastise (spell 1277331)
- Dark Sacrifice (spell 1277324)
- Holy Fire (spell 14914)
- Holy Nova (spell 15237)
- Mind Flay (spell 15407)
- Smite (spell 598)
- Starshards (spell 19296)

### Level 30

- Chastise (spell 1277332)
- Dark Sacrifice (spell 1277325)
- Holy Fire (spell 15263)
- Holy Nova (spell 15430)
- Mind Blast (spell 8104)
- Smite (spell 1004)
- Starshards (spell 19299)

### Level 38

- Chastise (spell 1277332)
- Dark Sacrifice (spell 1277325)
- Holy Fire (spell 15264)
- Holy Nova (spell 15431)
- Mind Blast (spell 8105)
- Shadow Word: Death (spell 1309595)
- Smite (spell 6060)
- Starshards (spell 19302)

### Level 40

- Chastise (spell 1277333)
- Dark Sacrifice (spell 1277326)
- Holy Fire (spell 15264)
- Holy Nova (spell 15431)
- Mind Blast (spell 8106)
- Shadow Word: Death (spell 1309633)
- Smite (spell 6060)
- Starshards (spell 19302)

### Level 50

- Chastise (spell 1277334)
- Dark Sacrifice (spell 1277327)
- Holy Fire (spell 15266)
- Holy Nova (spell 27799)
- Mind Blast (spell 10945)
- Shadow Word: Death (spell 1309635)
- Smite (spell 10933)
- Starshards (spell 19304)

### Level 60

- Chastise (spell 1277335)
- Dark Sacrifice (spell 1277328)
- Holy Fire (spell 15261)
- Holy Nova (spell 27801)
- Mind Blast (spell 10947)
- Shadow Word: Death (spell 1309636)
- Smite (spell 10934)
- Starshards (spell 19305)

## Violations found in this run

None.
