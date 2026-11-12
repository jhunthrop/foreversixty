# druid-feral-bear rotation ladder

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
  this build (rogue, warrior-fury, hunter), a shield in the off_hand
  for shaman-elemental and shaman-enhancement (shamans never dual
  wield in Forever), ranged for hunter only. Each slot picks the
  highest item_level weapon (speed > 0, so a shield never fills a weapon slot; damage_max >
  0, so an unfinished/placeholder weapon row this build's item table
  still carries - equipping one hangs the engine mid-sim rather than
  simulating a zero-damage weapon, see the report - is never picked)
  with required_level <= the character's level, from
  data/builds/<active>/items/<class>.json filtered to
  data/builds/<active>/simitems.json's known ids, honoring the spec's
  handedness (warrior-arms and paladin-retribution two-hand only;
  warrior-fury and rogue one-hand only for both hands; the two shaman
  shield specs one-hand only). druid-feral picks no weapon at all.
  Every caster spec (priest-shadow, mage's three specs, warlock's three specs) fills
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
  build HAS spent points on it and the id still will not resolve. An
  aura such a talent registers (talentOwnedAuraIDs: Arcane Blast's
  stacking buff, Fingers of Frost's charge buff) is held to the same
  rule as the talent's own spell id.
  Anything else is a violation.
- Zero casts: one of the curated rotation's own castSpell lines, resolved
  to the id the engine's OWN rank rewrite actually casts at this level
  (sim/request's rankResolver: the spec's per-band rank choice, else
  sim/internal/spellranks.LearnedSpellID; not a second, approximate copy
  of that resolution - this ladder calls the same function
  sim/request's rewriteRotationRanks calls), that never fires
  in the run - unless the engine could not resolve that id at all
  (already counted as unresolved) or data/curated/apl/<spec>.json's
  expected_idle array names the LINE'S AUTHORED id with a reason.
  LearnedSpellID returning "not learned" means the engine's own
  rewrite already dropped the line before the request was built, which
  is not a violation to report twice. At level 60 the request skips the
  rewrite (the rotation was authored for level 60), so the id a line
  casts there is the authored one.
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
| 10 | 0000000000000000-10000000000000000000-0000000000000000 | main_hand:284166 | 9.5 | 4 | spell:6807/1=12.9, spell:6807=11.5, other:attack/1=8.3, spell:99=2.5, spell:1259799=1.0 | - |
| 20 | 0000000000000000-51100301000000000000-0000000000000000 | main_hand:7230 | 22.6 | 6 | spell:6808/1=19.5, spell:6808=18.5, other:rage_gain=10.2, spell:1735=2.5, spell:770/1=2.2 | - |
| 30 | 0000000000000000-51100332020102100000-0000000000000000 | main_hand:13045 | 59.7 | 7 | spell:1328617=28.4, other:rage_gain=14.5, other:attack/1=12.0, spell:6809/1=8.1, spell:6809=7.8 | - |
| 38 | 0000000000000000-51100332023102123000-0000000000000000 | main_hand:13045 | 65.6 | 8 | spell:1328616=22.6, other:attack/1=14.3, other:rage_gain=13.9, spell:8972/1=6.6, spell:8972=5.6 | - |
| 40 | 0000000000000000-51200332023102123010-0000000000000000 | main_hand:13045 | 77.7 | 9 | spell:1328616=18.3, other:attack/1=14.3, other:rage_gain=14.3, spell:407995=7.9, spell:8972/1=6.0 | - |
| 50 | 0000000000000000-51200332023102125510-0300000000000000 | main_hand:13045 | 110.2 | 10 | spell:1235826=19.3, other:rage_gain=18.0, spell:9880/1=11.0, spell:9880=10.3, other:attack/1=9.5 | - |
| 60 | 0000000000000000-51230332023102125510-0550000000000000 | main_hand:22632 | 258.5 | 10 | spell:1235827=73.7, spell:9881/1=71.5, spell:9881=69.9, other:rage_gain=36.3, spell:1238070=30.2 | - |

## Learned but unused (informational)


### Level 10

- Entangling Roots (spell 339)
- Moonfire (spell 8924)
- Wrath (spell 5177)

### Level 20

- Claw (spell 1082)
- Entangling Roots (spell 1062)
- Moonfire (spell 8925)
- Rip (spell 1079)
- Shred (spell 5221)
- Starfire (spell 2912)
- Swipe (spell 779)
- Wrath (spell 5178)

### Level 30

- Claw (spell 3029)
- Entangling Roots (spell 5195)
- Ferocious Bite (spell 1328618)
- Insect Swarm (spell 24974)
- Moonfire (spell 8927)
- Rake (spell 1822)
- Rip (spell 9492)
- Shred (spell 6800)
- Starfire (spell 8949)
- Swipe (spell 780)
- Wrath (spell 5180)

### Level 38

- Claw (spell 5201)
- Entangling Roots (spell 5196)
- Ferocious Bite (spell 22568)
- Insect Swarm (spell 24974)
- Moonfire (spell 8928)
- Rake (spell 1823)
- Ravage (spell 6785)
- Rip (spell 9493)
- Shred (spell 8992)
- Starfire (spell 8950)
- Swipe (spell 769)
- Wrath (spell 6780)

### Level 40

- Claw (spell 5201)
- Entangling Roots (spell 5196)
- Ferocious Bite (spell 22827)
- Hurricane (spell 16914)
- Insect Swarm (spell 24975)
- Moonfire (spell 8929)
- Rake (spell 1823)
- Ravage (spell 6785)
- Rip (spell 9493)
- Shred (spell 8992)
- Starfire (spell 8950)
- Swipe (spell 769)
- Wrath (spell 6780)

### Level 50

- Claw (spell 9849)
- Entangling Roots (spell 9852)
- Ferocious Bite (spell 22828)
- Hurricane (spell 17401)
- Insect Swarm (spell 24976)
- Moonfire (spell 9833)
- Rake (spell 1824)
- Ravage (spell 9866)
- Rip (spell 9752)
- Shred (spell 9829)
- Starfire (spell 9875)
- Swipe (spell 9754)
- Wrath (spell 8905)

### Level 60

- Claw (spell 9850)
- Entangling Roots (spell 9853)
- Ferocious Bite (spell 22829)
- Hurricane (spell 17402)
- Insect Swarm (spell 24977)
- Moonfire (spell 9835)
- Primal Bite (spell 1238073)
- Rake (spell 9904)
- Ravage (spell 9867)
- Rip (spell 9896)
- Shred (spell 9830)
- Starfire (spell 9876)
- Swipe (spell 9908)
- Test Maul (spell 24042)
- Wrath (spell 9912)

## Violations found in this run

- druid-feral-bear level=20 kind=zero_casts spell="Swipe" id=779 authored=9908
- druid-feral-bear level=30 kind=zero_casts spell="Swipe" id=780 authored=9908
- druid-feral-bear level=38 kind=zero_casts spell="Swipe" id=769 authored=9908
- druid-feral-bear level=40 kind=zero_casts spell="Swipe" id=769 authored=9908
- druid-feral-bear level=50 kind=zero_casts spell="Swipe" id=9754 authored=9908
- druid-feral-bear level=60 kind=zero_casts spell="Swipe" id=9908 authored=9908
