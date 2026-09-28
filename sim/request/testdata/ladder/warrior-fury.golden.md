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
| 10 | 00000000000000000-100000000000000000-000000000000000000 | main_hand:274265 | 24.4 | 2 | other:attack/1=65.3, other:rage_gain=38.0, spell:1680=18.5, spell:2687=3.5 | {SpellID: 12328}, {SpellID: 23894}, {SpellID: 284, Tag: 1}, {SpellID: 6673} |
| 20 | 00000000000000000-551000000000000000-000000000000000000 | main_hand:274265 off_hand:9608 | 28.1 | 2 | other:attack/2=95.8, other:attack/1=65.2, other:rage_gain=37.7, spell:1680=18.3, spell:2687=3.5 | {SpellID: 12328}, {SpellID: 23894}, {SpellID: 285, Tag: 1}, {SpellID: 5242} |
| 30 | 00000000000000000-555000005010000000-000000000000000000 | main_hand:274265 off_hand:9608 | 38.8 | 2 | other:attack/2=95.8, other:attack/1=65.2, other:rage_gain=38.0, spell:1680=13.9, spell:2687=3.5 | {SpellID: 12328}, {SpellID: 1608, Tag: 1}, {SpellID: 23894}, {SpellID: 6192} |
| 38 | 00000000000000000-555000005050010030-000000000000000000 | main_hand:274265 off_hand:9608 | 44.1 | 3 | other:attack/2=98.5, other:attack/1=67.0, other:rage_gain=38.0, spell:1680=13.5, spell:2687=3.5 | {SpellID: 11549}, {SpellID: 11564, Tag: 1}, {SpellID: 23894} |
| 40 | 00000000000000000-555000005050010050-000000000000000000 | main_hand:274265 off_hand:9608 | 44.3 | 3 | other:attack/2=100.1, other:attack/1=68.1, other:rage_gain=38.1, spell:1680=11.4, spell:2687=3.5 | {SpellID: 11549}, {SpellID: 11565, Tag: 1}, {SpellID: 23881} |
| 50 | 35100000000000000-555000005050010051-000000000000000000 | main_hand:12061 off_hand:17738 | 70.0 | 5 | other:attack/2=107.7, other:attack/1=81.0, other:rage_gain=38.1, spell:23892=11.6, spell:2687=3.5 | {SpellID: 11550}, {SpellID: 11566, Tag: 1} |
| 60 | 35310003002000000-555000005050010051-000000000000000000 | main_hand:23577 off_hand:234558 | 140.1 | 8 | other:attack/1=129.8, other:attack/2=67.8, other:rage_gain=38.0, spell:23894=15.5, spell:1680=4.3 | - |

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
- Copy of Mortal Strike (spell 26652)
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

- Copy of Mortal Strike (spell 26652)
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

- Copy of Mortal Strike (spell 26652)
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

- warrior-fury level=10 kind=unresolved_id action={SpellID: 12328}
- warrior-fury level=10 kind=unresolved_id action={SpellID: 23894}
- warrior-fury level=10 kind=unresolved_id action={SpellID: 284, Tag: 1}
- warrior-fury level=10 kind=unresolved_id action={SpellID: 6673}
- warrior-fury level=10 kind=zero_casts spell="Heroic Strike" id=284 authored=25286
- warrior-fury level=20 kind=unresolved_id action={SpellID: 12328}
- warrior-fury level=20 kind=unresolved_id action={SpellID: 23894}
- warrior-fury level=20 kind=unresolved_id action={SpellID: 285, Tag: 1}
- warrior-fury level=20 kind=unresolved_id action={SpellID: 5242}
- warrior-fury level=20 kind=zero_casts spell="Heroic Strike" id=285 authored=25286
- warrior-fury level=30 kind=unresolved_id action={SpellID: 12328}
- warrior-fury level=30 kind=unresolved_id action={SpellID: 1608, Tag: 1}
- warrior-fury level=30 kind=unresolved_id action={SpellID: 23894}
- warrior-fury level=30 kind=unresolved_id action={SpellID: 6192}
- warrior-fury level=30 kind=zero_casts spell="Execute" id=5308 authored=20662
- warrior-fury level=30 kind=zero_casts spell="Heroic Strike" id=1608 authored=25286
- warrior-fury level=38 kind=unresolved_id action={SpellID: 11549}
- warrior-fury level=38 kind=unresolved_id action={SpellID: 11564, Tag: 1}
- warrior-fury level=38 kind=unresolved_id action={SpellID: 23894}
- warrior-fury level=38 kind=zero_casts spell="Execute" id=20658 authored=20662
- warrior-fury level=38 kind=zero_casts spell="Heroic Strike" id=11564 authored=25286
- warrior-fury level=40 kind=unresolved_id action={SpellID: 11549}
- warrior-fury level=40 kind=unresolved_id action={SpellID: 11565, Tag: 1}
- warrior-fury level=40 kind=unresolved_id action={SpellID: 23881}
- warrior-fury level=40 kind=zero_casts spell="Execute" id=20660 authored=20662
- warrior-fury level=40 kind=zero_casts spell="Heroic Strike" id=11565 authored=25286
- warrior-fury level=50 kind=unresolved_id action={SpellID: 11550}
- warrior-fury level=50 kind=unresolved_id action={SpellID: 11566, Tag: 1}
- warrior-fury level=50 kind=zero_casts spell="Execute" id=20661 authored=20662
- warrior-fury level=50 kind=zero_casts spell="Heroic Strike" id=11566 authored=25286
- warrior-fury level=60 kind=zero_casts spell="Execute" id=20662 authored=20662
