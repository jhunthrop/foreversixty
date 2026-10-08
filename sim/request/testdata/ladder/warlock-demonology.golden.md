# warlock-demonology rotation ladder

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
| 10 | 00000000000000000-1000000000000000000-0000000000000000 | main_hand:263937 ranged:263430 | 8.8 | 5 | other:mana_gain=57.5, spell:1454=57.5, spell:5019=56.9, spell:695=16.4, spell:172=14.2 | {OtherID: 13}, {SpellID: 603} |
| 20 | 00000000000000000-2332100000000000000-0000000000000000 | main_hand:890 ranged:5243 | 17.7 | 5 | spell:5019=48.0, other:mana_gain=40.0, spell:1455=40.0, spell:1088=17.1, spell:1094=11.9 | {OtherID: 13}, {SpellID: 603} |
| 30 | 00000000000000000-2332113101220000000-0000000000000000 | main_hand:249392 ranged:5213 | 28.3 | 5 | spell:5019=45.1, other:mana_gain=43.1, spell:1456=43.1, spell:1106=25.6, spell:2941=12.3 | {OtherID: 13}, {SpellID: 603} |
| 38 | 00000000000000000-2332113101220001340-0000000000000000 | main_hand:1664 ranged:13064 | 52.9 | 5 | spell:5019=30.9, spell:7641=28.1, other:mana_gain=28.0, spell:11687=28.0, spell:2941=13.1 | {OtherID: 13}, {SpellID: 603} |
| 40 | 10000000000000000-2332113101220001350-0000000000000000 | main_hand:1664 ranged:5216 | 81.4 | 5 | spell:5019=32.2, spell:7641=28.6, other:mana_gain=27.6, spell:11687=27.6, spell:11665=12.9 | {OtherID: 13}, {SpellID: 603} |
| 50 | 25400000000000000-2332113101220001350-0000000000000000 | main_hand:812 ranged:249232 | 104.6 | 5 | spell:5019=32.7, spell:11659=27.7, other:mana_gain=24.0, spell:11688=24.0, spell:11667=13.3 | {OtherID: 13}, {SpellID: 603} |
| 60 | 25532000000000000-2332113101220001350-0040000000000000 | main_hand:22630 ranged:22821 | 306.6 | 6 | spell:11661=37.9, spell:5019=37.8, other:mana_gain=22.8, spell:11689=22.8, spell:11668=13.8 | {OtherID: 13} |

## Learned but unused (informational)


### Level 10

- Bane of Agony (spell 980)
- Drain Soul (spell 1120)
- Firebolt (spell 7799)
- Shadow Cleave (spell 403839)

### Level 20

- Bane of Agony (spell 1014)
- Drain Soul (spell 1120)
- Firebolt (spell 7800)
- Health Funnel (spell 3698)
- Lash of Pain (spell 7814)
- Rain of Fire (spell 5740)
- Searing Pain (spell 5676)
- Shadow Cleave (spell 403841)
- Shadowburn (spell 17877)

### Level 30

- Bane of Agony (spell 6217)
- Conflagrate (spell 1293817)
- Drain Soul (spell 8288)
- Firebolt (spell 7801)
- Health Funnel (spell 3699)
- Hellfire (spell 1949)
- Hellfire Effect (spell 5857)
- Lash of Pain (spell 7815)
- Rain of Fire (spell 5740)
- Searing Pain (spell 17919)
- Shadow Cleave (spell 403842)
- Shadowburn (spell 18867)

### Level 38

- Bane of Agony (spell 11711)
- Conflagrate (spell 1293818)
- Drain Soul (spell 8289)
- Firebolt (spell 7802)
- Health Funnel (spell 3700)
- Hellfire (spell 1949)
- Hellfire Effect (spell 5857)
- Lash of Pain (spell 7816)
- Rain of Fire (spell 6219)
- Searing Pain (spell 17920)
- Shadow Cleave (spell 403843)
- Shadowburn (spell 18868)

### Level 40

- Bane of Agony (spell 11711)
- Conflagrate (spell 17962)
- Drain Soul (spell 8289)
- Firebolt (spell 7802)
- Haunt (spell 403501)
- Health Funnel (spell 3700)
- Hellfire (spell 1949)
- Hellfire Effect (spell 5857)
- Incinerate (spell 412758)
- Lash of Pain (spell 7816)
- Rain of Fire (spell 6219)
- Searing Pain (spell 17920)
- Shadow Cleave (spell 403843)
- Shadowburn (spell 18869)
- Unstable Affliction (spell 427717)
- Wrack (spell 1316697)

### Level 50

- Bane of Agony (spell 11712)
- Conflagrate (spell 18930)
- Drain Soul (spell 8289)
- Firebolt (spell 11762)
- Haunt (spell 1293693)
- Health Funnel (spell 11693)
- Hellfire (spell 11683)
- Hellfire Effect (spell 11681)
- Incinerate (spell 1293812)
- Lash of Pain (spell 11778)
- Rain of Fire (spell 11677)
- Searing Pain (spell 17922)
- Shadow Cleave (spell 403844)
- Shadowburn (spell 18870)
- Soul Fire (spell 6353)
- Unstable Affliction (spell 1242972)
- Wrack (spell 1316697)

### Level 60

- Bane of Agony (spell 11713)
- Conflagrate (spell 18932)
- Drain Soul (spell 11675)
- Firebolt (spell 11763)
- Haunt (spell 1293694)
- Health Funnel (spell 11695)
- Hellfire (spell 11684)
- Hellfire Effect (spell 11682)
- Incinerate (spell 1293813)
- Lash of Pain (spell 11780)
- Rain of Fire (spell 11678)
- Searing Pain (spell 17923)
- Shadow Cleave (spell 403852)
- Shadowburn (spell 18871)
- Soul Fire (spell 17924)
- Test Curse of Agony (spell 28608)
- Unstable Affliction (spell 1242972)
- Wrack (spell 1316697)

## Violations found in this run

- warlock-demonology level=10 kind=unresolved_id action={SpellID: 603}
- warlock-demonology level=20 kind=unresolved_id action={SpellID: 603}
- warlock-demonology level=30 kind=unresolved_id action={SpellID: 603}
- warlock-demonology level=38 kind=unresolved_id action={SpellID: 603}
- warlock-demonology level=40 kind=unresolved_id action={SpellID: 603}
- warlock-demonology level=50 kind=unresolved_id action={SpellID: 603}
