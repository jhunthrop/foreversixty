# druid-feral rotation ladder

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
| 10 | 0000000000000000-10000000000000000000-0000000000000000 | main_hand:1933 | 8.4 | 4 | spell:5177=44.1, other:attack/1=31.3, spell:8924=5.5, spell:1259799=1.5, spell:29166=1.0 | {SpellID: 1322605}, {SpellID: 9850} |
| 20 | 0000000000000000-55100000000000000000-0000000000000000 | bare | 49.9 | 4 | other:attack/1=182.1, spell:1082=46.4, spell:1079=8.7, spell:1259799=1.5, spell:5215=1.0 | {SpellID: 1322605}, {SpellID: 9830} |
| 30 | 0000000000000000-55232031000000000000-0000000000000000 | bare | 72.2 | 5 | other:attack/1=182.1, spell:6800=44.4, spell:9492=8.2, spell:5217=6.6, spell:1259799=1.5 | {SpellID: 1322605} |
| 38 | 0000000000000000-55232032121030000000-0000000000000000 | bare | 98.6 | 7 | other:attack/1=182.1, spell:8992=50.6, spell:9493=9.8, spell:1322605=8.3, spell:5217=6.6 | - |
| 40 | 0000000000000000-55232032121032000000-0000000000000000 | bare | 100.9 | 7 | other:attack/1=182.1, spell:8992=50.1, spell:9493=10.5, spell:1322605=8.3, spell:5217=6.6 | - |
| 50 | 0100000000000000-55232032121032012001-5000000000000000 | bare | 131.2 | 8 | other:attack/1=182.1, spell:9829=48.8, spell:9752=11.5, spell:1322605=8.0, spell:5217=6.6 | - |
| 60 | 0100000000000000-55232032121032012001-5053200000000000 | bare | 173.9 | 8 | other:attack/1=182.1, spell:9830=50.6, spell:9896=11.7, spell:1322605=9.8, spell:5217=6.6 | - |

## Learned but unused (informational)


### Level 10

- Entangling Roots (spell 339)
- Maul (spell 6807)

### Level 20

- Entangling Roots (spell 1062)
- Insect Swarm (spell 5570)
- Maul (spell 6808)
- Moonfire (spell 8925)
- Starfire (spell 2912)
- Swipe (spell 779)
- Wrath (spell 5178)

### Level 30

- Claw (spell 3029)
- Entangling Roots (spell 5195)
- Insect Swarm (spell 24974)
- Maul (spell 6809)
- Moonfire (spell 8927)
- Primal Bite (spell 407995)
- Rake (spell 1822)
- Starfire (spell 8949)
- Swipe (spell 780)
- Wrath (spell 5180)

### Level 38

- Claw (spell 5201)
- Entangling Roots (spell 5196)
- Ferocious Bite (spell 22568)
- Insect Swarm (spell 24974)
- Maul (spell 8972)
- Moonfire (spell 8928)
- Primal Bite (spell 1238069)
- Rake (spell 1823)
- Starfire (spell 8950)
- Swipe (spell 769)
- Wrath (spell 6780)

### Level 40

- Claw (spell 5201)
- Entangling Roots (spell 5196)
- Ferocious Bite (spell 22827)
- Hurricane (spell 16914)
- Insect Swarm (spell 24975)
- Maul (spell 8972)
- Moonfire (spell 8929)
- Primal Bite (spell 1238069)
- Rake (spell 1823)
- Starfire (spell 8950)
- Swipe (spell 769)
- Wrath (spell 6780)

### Level 50

- Claw (spell 9849)
- Entangling Roots (spell 9852)
- Ferocious Bite (spell 22828)
- Hurricane (spell 17401)
- Insect Swarm (spell 24976)
- Lacerate (spell 1235826)
- Maul (spell 9880)
- Moonfire (spell 9833)
- Primal Bite (spell 1238070)
- Rake (spell 1824)
- Starfire (spell 9875)
- Swipe (spell 9754)
- Wrath (spell 8905)

### Level 60

- Claw (spell 9850)
- Entangling Roots (spell 9853)
- Ferocious Bite (spell 22829)
- Hurricane (spell 17402)
- Insect Swarm (spell 24977)
- Lacerate (spell 1235827)
- Maul (spell 9881)
- Moonfire (spell 9835)
- Primal Bite (spell 1238073)
- Rake (spell 9904)
- Starfire (spell 9876)
- Swipe (spell 9908)
- Test Maul (spell 24042)
- Wrath (spell 9912)

## Violations found in this run

- druid-feral level=38 kind=zero_casts spell="Ferocious Bite" id=22568 authored=22829
- druid-feral level=40 kind=zero_casts spell="Ferocious Bite" id=22827 authored=22829
- druid-feral level=50 kind=zero_casts spell="Ferocious Bite" id=22828 authored=22829
- druid-feral level=60 kind=zero_casts spell="Ferocious Bite" id=22829 authored=22829
