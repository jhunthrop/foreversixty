# hunter-marksmanship rotation ladder

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
| 10 | 0000000000000000-10000000000000000-000000000000000000 | main_hand:1927 off_hand:1287 ranged:3036 | 52.0 | 3 | other:shoot=83.8, spell:3044=17.5, spell:13549=17.2, other:move=1.0, spell:13165=1.0 | {SpellID: 20904} |
| 20 | 0000000000000000-35300000000000000-000000000000000000 | main_hand:1482 off_hand:2236 ranged:3021 | 64.9 | 5 | other:shoot=77.6, spell:2643=13.5, spell:19434=7.4, spell:13550=3.1, spell:14282=1.9 | - |
| 30 | 0000000000000000-35305500000000000-000000000000000000 | main_hand:6692 off_hand:23168 ranged:274748 | 74.6 | 6 | other:shoot=76.5, spell:2643=11.2, spell:20900=6.5, spell:14283=2.8, spell:13551=2.2 | - |
| 38 | 0000000000000000-35305500115001000-000000000000000000 | main_hand:869 off_hand:6829 ranged:2825 | 96.7 | 6 | other:shoot=80.0, spell:2643=12.1, spell:20901=4.8, spell:13552=2.2, spell:14284=1.9 | - |
| 40 | 0000000000000000-35305500115003000-000000000000000000 | main_hand:2164 off_hand:9465 ranged:2825 | 100.9 | 6 | other:shoot=80.0, spell:20901=8.6, spell:2643=7.9, spell:13552=2.2, spell:14284=2.0 | - |
| 50 | 5500000000000000-35305500115003000-000000000000000000 | main_hand:2163 off_hand:6660 ranged:2824 | 115.7 | 5 | other:shoot=158.7, spell:2643=15.7, spell:13554=1.1, other:move=1.0, spell:14321=1.0 | - |
| 60 | 5522000000000000-35305500115003000-510000000000000000 | main_hand:22736 off_hand:23054 ranged:22811 | 208.5 | 6 | other:shoot=83.5, spell:2643=11.6, spell:20904=3.0, spell:25295=2.1, other:move=1.0 | - |

## Learned but unused (informational)


### Level 10

- Bite (spell 17255)
- Claw (spell 16828)
- Lightning Breath (spell 24844)
- Mine! (spell 1265054)
- Raptor Strike (spell 14260)
- Scorpid Poison (spell 24640)
- Widow Bite (spell 26226)
- Wyvern Strike (spell 458482)

### Level 20

- Bite (spell 17256)
- Claw (spell 16829)
- Dismember (spell 1264758)
- Immolation Trap Effect (spell 13797)
- Lightning Breath (spell 25008)
- Mine! (spell 1265055)
- Mongoose Bite (spell 1495)
- Pinch (spell 1264735)
- Raptor Strike (spell 14261)
- Savage Rend (spell 1265065)
- Scorpid Poison (spell 24640)
- Sonic Blast (spell 1264478)
- Swipe (spell 1264494)
- Tendon Rip (spell 1265038)
- Web (spell 1265843)
- Widow Bite (spell 26226)
- Wing Clip (spell 2974)
- Wyvern Strike (spell 458482)

### Level 30

- Bite (spell 17257)
- Claw (spell 16830)
- Counterattack (spell 1242634)
- Dismember (spell 1264927)
- Immolation Trap Effect (spell 14298)
- Lacerate (spell 24118)
- Lightning Breath (spell 25009)
- Mine! (spell 1265056)
- Mongoose Bite (spell 14269)
- Pinch (spell 1264736)
- Raptor Strike (spell 14262)
- Savage Rend (spell 1265066)
- Scorpid Poison (spell 24583)
- Sonic Blast (spell 1264479)
- Strider Kick (spell 1317257)
- Summon Hawk (spell 1293241)
- Swipe (spell 1264497)
- Tendon Rip (spell 1265039)
- Thunderstomp (spell 26090)
- Web (spell 1265878)
- Widow Bite (spell 26226)
- Wing Clip (spell 2974)
- Wyvern Strike (spell 458482)

### Level 38

- Bite (spell 17258)
- Claw (spell 16831)
- Counterattack (spell 1242634)
- Dismember (spell 1264929)
- Explosive Trap Effect (spell 13812)
- Immolation Trap Effect (spell 14299)
- Lacerate (spell 24118)
- Lightning Breath (spell 25010)
- Mine! (spell 1265056)
- Mongoose Bite (spell 14269)
- Pinch (spell 1264739)
- Raptor Strike (spell 14263)
- Savage Rend (spell 1265067)
- Scorpid Poison (spell 24583)
- Sonic Blast (spell 1264480)
- Strider Kick (spell 1317257)
- Summon Hawk (spell 1293525)
- Swipe (spell 1264498)
- Tendon Rip (spell 1265040)
- Thunderstomp (spell 26090)
- Web (spell 1265880)
- Widow Bite (spell 26226)
- Wing Clip (spell 14267)
- Wyvern Strike (spell 458482)

### Level 40

- Bite (spell 17259)
- Claw (spell 16832)
- Counterattack (spell 1242634)
- Dismember (spell 1264929)
- Explosive Trap Effect (spell 13812)
- Immolation Trap Effect (spell 14299)
- Lacerate (spell 24119)
- Lightning Breath (spell 25010)
- Mine! (spell 1265057)
- Mongoose Bite (spell 14269)
- Pinch (spell 1264739)
- Raptor Strike (spell 14264)
- Savage Rend (spell 1265067)
- Scorpid Poison (spell 24586)
- Sniper Shot (spell 1310687)
- Sonic Blast (spell 1264480)
- Strider Kick (spell 1317257)
- Summon Hawk (spell 1293525)
- Swipe (spell 1264498)
- Tendon Rip (spell 1265040)
- Thunderstomp (spell 26187)
- Volley (spell 1510)
- Web (spell 1265880)
- Widow Bite (spell 26226)
- Wing Clip (spell 14267)
- Wyvern Strike (spell 458482)

### Level 50

- Arcane Shot (spell 14285)
- Bite (spell 17260)
- Claw (spell 3010)
- Counterattack (spell 20909)
- Dismember (spell 1264930)
- Explosive Trap Effect (spell 14314)
- Immolation Trap Effect (spell 14300)
- Lacerate (spell 24120)
- Lava Breath (spell 444681)
- Lightning Breath (spell 25011)
- Mine! (spell 1265058)
- Mongoose Bite (spell 14270)
- Pinch (spell 1264741)
- Raptor Strike (spell 14265)
- Savage Rend (spell 1265068)
- Scorpid Poison (spell 24586)
- Sniper Shot (spell 1310785)
- Sonic Blast (spell 1264481)
- Strider Kick (spell 1317257)
- Summon Hawk (spell 1293526)
- Swipe (spell 1264501)
- Tendon Rip (spell 1265041)
- Thunderstomp (spell 26188)
- Volley (spell 14294)
- Web (spell 1265881)
- Widow Bite (spell 26226)
- Wing Clip (spell 14267)
- Wyvern Strike (spell 458482)

### Level 60

- Bite (spell 17261)
- Claw (spell 3009)
- Counterattack (spell 20910)
- Dismember (spell 1264933)
- Explosive Trap Effect (spell 14315)
- Hydra Shot (spell 1293020)
- Immolation Trap Effect (spell 14301)
- Lacerate (spell 1299332)
- Lava Breath (spell 444681)
- Lightning Breath (spell 25012)
- Mine! (spell 1265058)
- Mongoose Bite (spell 14271)
- Pinch (spell 1264742)
- Raptor Strike (spell 14266)
- Savage Rend (spell 1265069)
- Scorpid Poison (spell 24587)
- Sniper Shot (spell 1310786)
- Sonic Blast (spell 1264482)
- Strider Kick (spell 1317257)
- Summon Hawk (spell 1293527)
- Swipe (spell 1264502)
- Tendon Rip (spell 1265042)
- Thunderstomp (spell 1264455)
- Volley (spell 14295)
- Web (spell 1265883)
- Widow Bite (spell 26226)
- Wing Clip (spell 14268)
- Wyvern Strike (spell 458482)

## Violations found in this run

- hunter-marksmanship level=50 kind=zero_casts spell="Arcane Shot" id=14285 authored=14287
