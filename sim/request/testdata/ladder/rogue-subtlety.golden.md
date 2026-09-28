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
| 10 | 00000000000000000-00000000000000000-1000000000000000000 | main_hand:274271 | 17.9 | 0 | other:attack/1=101.4 | {SpellID: 14183}, {SpellID: 14278}, {SpellID: 16511} |
| 20 | 00000000000000000-00000000000000000-5321000000000000000 | main_hand:274271 | 18.9 | 0 | other:attack/1=101.4 | {SpellID: 14183}, {SpellID: 14278}, {SpellID: 16511} |
| 30 | 00000000000000000-00000000000000000-5323221300000000000 | main_hand:274271 | 20.2 | 0 | other:attack/1=101.4 | {SpellID: 14183}, {SpellID: 14278}, {SpellID: 16511} |
| 38 | 00000000000000000-00000000000000000-5323221310003001030 | main_hand:274271 | 25.7 | 3 | other:attack/1=101.4, spell:14278=9.6, spell:8623/5=1.3, spell:5171/5=0.0 | {SpellID: 14183}, {SpellID: 16511} |
| 40 | 00000000000000000-00000000000000000-5323221310003001050 | main_hand:274271 | 27.0 | 3 | other:attack/1=101.4, spell:14278=9.6, spell:8624/5=1.3, spell:5171/5=0.0 | {SpellID: 14183}, {SpellID: 16511} |
| 50 | 32100000000000000-31000000000000000-5323221310003001050 | main_hand:12061 off_hand:17738 | 40.6 | 3 | other:attack/2=101.3, other:attack/1=76.1, spell:14278=9.6, spell:11299/5=1.4, spell:6774/5=0.0 | {SpellID: 14183}, {SpellID: 16511} |
| 60 | 32100000000000000-32003000000000000-5323221310003001050 | main_hand:23577 off_hand:234558 | 62.4 | 3 | other:attack/1=121.4, other:attack/2=63.0, spell:14278=9.6, spell:31016/5=1.3, spell:6774/5=0.0 | {SpellID: 14183}, {SpellID: 16511} |

## Learned but unused (informational)


### Level 10

- Backstab (spell 53)
- Eviscerate (spell 6760)
- Gouge (spell 1776)
- Serrated Blades (spell 461327)
- Sinister Strike (spell 1757)

### Level 20

- Ambush (spell 8676)
- Backstab (spell 2590)
- Eviscerate (spell 6761)
- Garrote (spell 703)
- Gouge (spell 1777)
- Kick (spell 1766)
- Rupture (spell 1943)
- Serrated Blades (spell 461327)
- Sinister Strike (spell 1758)

### Level 30

- Ambush (spell 8724)
- Backstab (spell 2591)
- Eviscerate (spell 6762)
- Garrote (spell 8632)
- Gouge (spell 1777)
- Kick (spell 1767)
- Rupture (spell 8639)
- Serrated Blades (spell 461327)
- Sinister Strike (spell 1760)

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
- Copy of Deadly Poison IV (spell 25348)
- Garrote (spell 11290)
- Gouge (spell 11286)
- Kick (spell 1769)
- Rupture (spell 11275)
- Serrated Blades (spell 461327)
- Sinister Strike (spell 11294)
- Test Stab R50 (spell 23959)
- Test Strike R50 (spell 23960)

## Violations found in this run

- rogue-subtlety level=10 kind=unresolved_id action={SpellID: 14183}
- rogue-subtlety level=10 kind=unresolved_id action={SpellID: 14278}
- rogue-subtlety level=10 kind=unresolved_id action={SpellID: 16511}
- rogue-subtlety level=10 kind=zero_casts spell="Eviscerate" id=6760 authored=31016
- rogue-subtlety level=10 kind=zero_casts spell="Slice and Dice" id=5171 authored=6774
- rogue-subtlety level=20 kind=no_damage_cast dps=18.9
- rogue-subtlety level=20 kind=unresolved_id action={SpellID: 14183}
- rogue-subtlety level=20 kind=unresolved_id action={SpellID: 14278}
- rogue-subtlety level=20 kind=unresolved_id action={SpellID: 16511}
- rogue-subtlety level=20 kind=zero_casts spell="Ambush" id=8676 authored=11269
- rogue-subtlety level=20 kind=zero_casts spell="Eviscerate" id=6761 authored=31016
- rogue-subtlety level=20 kind=zero_casts spell="Slice and Dice" id=5171 authored=6774
- rogue-subtlety level=30 kind=no_damage_cast dps=20.2
- rogue-subtlety level=30 kind=unresolved_id action={SpellID: 14183}
- rogue-subtlety level=30 kind=unresolved_id action={SpellID: 14278}
- rogue-subtlety level=30 kind=unresolved_id action={SpellID: 16511}
- rogue-subtlety level=30 kind=zero_casts spell="Ambush" id=8724 authored=11269
- rogue-subtlety level=30 kind=zero_casts spell="Eviscerate" id=6762 authored=31016
- rogue-subtlety level=30 kind=zero_casts spell="Slice and Dice" id=5171 authored=6774
- rogue-subtlety level=38 kind=unresolved_id action={SpellID: 14183}
- rogue-subtlety level=38 kind=unresolved_id action={SpellID: 16511}
- rogue-subtlety level=38 kind=zero_casts spell="Ambush" id=8725 authored=11269
- rogue-subtlety level=40 kind=unresolved_id action={SpellID: 14183}
- rogue-subtlety level=40 kind=unresolved_id action={SpellID: 16511}
- rogue-subtlety level=40 kind=zero_casts spell="Ambush" id=8725 authored=11269
- rogue-subtlety level=50 kind=unresolved_id action={SpellID: 14183}
- rogue-subtlety level=50 kind=unresolved_id action={SpellID: 16511}
- rogue-subtlety level=50 kind=zero_casts spell="Ambush" id=11268 authored=11269
- rogue-subtlety level=60 kind=unresolved_id action={SpellID: 14183}
- rogue-subtlety level=60 kind=unresolved_id action={SpellID: 16511}
- rogue-subtlety level=60 kind=zero_casts spell="Ambush" id=11269 authored=11269
