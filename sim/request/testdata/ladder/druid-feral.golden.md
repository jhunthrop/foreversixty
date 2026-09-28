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
  hands). druid-feral picks no weapon at all. Every other slot is
  bare. Buffs and consumables: none.
- DPS regression: each level's DPS is compared against the ladder's own
  PREVIOUS rung (not literally level-10, since the ladder's own gaps
  are uneven - 30 to 38 is 8 levels, 38 to 40 is 2). A level scoring
  lower than the rung before it is a violation.
- Unresolved: an id the engine's ComputeStats warns it cannot resolve.
  Expected when data/curated/apl/<spec>.json's own inert array names
  it; otherwise a violation.
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
| 10 | 0000000000000000-1000000000000000000-0000000000000000 | bare | 28.4 | 1 | other:attack/1=182.1, spell:58984=1.5 | {SpellID: 5217} |
| 20 | 0000000000000000-5510000000000000000-0000000000000000 | bare | 50.8 | 3 | other:attack/1=182.1, spell:1082=46.6, spell:1079=8.6, spell:58984=1.5 | {SpellID: 5217} |
| 30 | 0000000000000000-5523231000000000000-0000000000000000 | bare | 62.5 | 6 | other:attack/1=182.1, spell:3029=26.5, spell:1822=19.8, spell:9492=8.9, spell:5217=6.6 | - |
| 38 | 0000000000000000-5523232020032000000-0000000000000000 | bare | 80.2 | 6 | other:attack/1=182.1, spell:5201=26.5, spell:1823=19.8, spell:9493=8.9, spell:5217=6.6 | {SpellID: 22568} |
| 40 | 0000000000000000-5523232020032010001-0000000000000000 | bare | 82.7 | 6 | other:attack/1=182.1, spell:5201=26.5, spell:1823=19.8, spell:9493=8.9, spell:5217=6.6 | {SpellID: 22827} |
| 50 | 0000000000000000-5523232020032010001-0550000000000000 | bare | 102.4 | 6 | other:attack/1=182.1, spell:9849=26.5, spell:1824=19.8, spell:9752=8.9, spell:5217=6.6 | {SpellID: 22828} |
| 60 | 0000000000000000-5523232020032010001-0550000000000000 | bare | 136.0 | 6 | other:attack/1=182.1, spell:9850=26.5, spell:9904=19.8, spell:9896=8.9, spell:5217=6.6 | - |

## Learned but unused (informational)


### Level 10

- Entangling Roots (spell 339)
- Maul (spell 6807)
- Moonfire (spell 8924)
- Wrath (spell 5177)

### Level 20

- Entangling Roots (spell 1062)
- Insect Swarm (spell 5570)
- Maul (spell 6808)
- Moonfire (spell 8925)
- Starfire (spell 2912)
- Swipe (spell 779)
- Wrath (spell 5178)

### Level 30

- Entangling Roots (spell 5195)
- Insect Swarm (spell 24974)
- Maul (spell 6809)
- Moonfire (spell 8927)
- Primal Bite (spell 407995)
- Starfire (spell 8949)
- Swipe (spell 780)
- Wrath (spell 5180)

### Level 38

- Entangling Roots (spell 5196)
- Ferocious Bite (spell 22568)
- Insect Swarm (spell 24974)
- Maul (spell 8972)
- Moonfire (spell 8928)
- Primal Bite (spell 1238069)
- Ravage (spell 6785)
- Starfire (spell 8950)
- Swipe (spell 769)
- Wrath (spell 6780)

### Level 40

- Entangling Roots (spell 5196)
- Ferocious Bite (spell 22827)
- Hurricane (spell 16914)
- Insect Swarm (spell 24975)
- Maul (spell 8972)
- Moonfire (spell 8929)
- Primal Bite (spell 1238069)
- Ravage (spell 6785)
- Starfire (spell 8950)
- Swipe (spell 769)
- Wrath (spell 6780)

### Level 50

- Entangling Roots (spell 9852)
- Ferocious Bite (spell 22828)
- Hurricane (spell 17401)
- Insect Swarm (spell 24976)
- Lacerate (spell 1235826)
- Maul (spell 9880)
- Moonfire (spell 9833)
- Primal Bite (spell 1238070)
- Ravage (spell 9866)
- Starfire (spell 9875)
- Swipe (spell 9754)
- Wrath (spell 8905)

### Level 60

- Entangling Roots (spell 9853)
- Ferocious Bite (spell 31018)
- Hurricane (spell 17402)
- Insect Swarm (spell 24977)
- Lacerate (spell 1235827)
- Maul (spell 9881)
- Moonfire (spell 9835)
- Primal Bite (spell 1238073)
- Ravage (spell 9867)
- Starfire (spell 25298)
- Swipe (spell 9908)
- Test Maul (spell 24042)
- Wrath (spell 9912)

## Violations found in this run

- druid-feral level=10 kind=unresolved_id action={SpellID: 5217}
- druid-feral level=20 kind=unresolved_id action={SpellID: 5217}
- druid-feral level=38 kind=unresolved_id action={SpellID: 22568}
- druid-feral level=40 kind=unresolved_id action={SpellID: 22827}
- druid-feral level=50 kind=unresolved_id action={SpellID: 22828}
- druid-feral level=60 kind=zero_casts spell="Ferocious Bite" id=31018 authored=31018
