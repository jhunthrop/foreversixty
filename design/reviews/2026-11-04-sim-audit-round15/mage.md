# Sim-accuracy round 15: mage

Specs: mage-arcane, mage-fire, mage-frost. Site wt/r15-mage on int/r15 (fd320b60e), fork lane/r15-mage on int/r15 (d1d9f2775). Client 1.60.1.70755.

Method:
- Seed items first: every row the guide-accuracy round 1 mage report hands to sim (SIM-1 to SIM-6; `origin/land/guides-r1:design/reviews/guide-accuracy/round1/mage.md`).
- Then the re-audit. 70755 changed nothing in mage talents, trainables or spellconst; only the Missile Barrage and Improved Frostbolt tooltips dropped Frost Blast (build.md X7; guide report). So the re-audit went to what earlier rounds had not settled.
- Spec-breakdown at raid 60, 10,000 iterations, seeds 1 and 50001, both factions. Exploration runs at 3,000. Every run went through `$R/sim.sh`.
- Owner rule, round 15: when the sim ties, search the web first. I did, for Arcane Power timing and for the Fire leveling toll point.

## Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Arcane Power and Presence of Mind went out on the pull.** The rotation pressed both at 0 s. Its own Arcane Power note says "On a fresh pull, give the tank a Sunder or two first". The guide's opener (land/guides-r1 arcane.md:73) blasts twice, then presses the cooldowns. The engine has no threat, so nothing stopped it. Both lines now wait 3 s, about two Arcane Blasts. Delays of 3 / 4.5 / 6 s measured 619.6 / 619.8 / 619.8 against 619.4 ± 0.7 (horde, seed 1, 3,000 iterations): a tie. Web: Warcraft Tavern's Forever Arcane guide calls Arcane Power the "BIG DAMAGE" button but gives no pull timing (https://www.warcrafttavern.com/forever/guides/arcane-mage/, no date shown, read 2026-11-03). So the guide and the rotation note decide. | `data/curated/apl/mage-arcane.json` items 3 and 4; test `sim/request/mage_pull_timing_test.go` failed, then passed; apl-check green | CONFIRMED-BUG | old, guides-handoff (SIM-2) | S 57b90e032, E d77b35469 |
| A2 | The racial, unlike the cooldowns in A1, still fires on the pull through `autocastOtherCooldowns`. I measured explicit Blood Fury and Eureka! lines gated at 3 s: horde 619.6 → 620.1, alliance 613.1 → 613.9, a tie. A new rotation line would show on the guide without a note, and notes belong to the guide lane, so I left it. | 3,000 iterations, seed 1 | Aligned (tie) | - | - |
| A3 | **Arcane Concentration from each Arcane Missile** (log-blocked since round 7). The engine never procs Clearcasting from a missile. Beta logs: 0 Clearcastings in 320 missile hits, against 32 in 313 Arcane Blast hits (10.2%) and 13 in 325 Arcane Explosion hits. Read from 32 mage timelines on foreverlogs.gg, round 13's cache. A gain was credited to the mage's own damage event 0 to 80 ms before it (the server applies the buff about 18 ms after the hit). | `r15/mage/ccmiss.py` over `r13/casters/tl/`; `mage/talents.go:244` | Aligned (log); closes the carried row | - | - |
| A4 | Arcane Geometry 1 (range) in both builds is not needed for any tier. It is kept on purpose (raid-keeps.json, the guide's "extra range for pulling") and is not provably dead. | talents/mage.json tiers | Aligned | - | - |
| A5 | Re-read against the 70755 rows, all aligned (rounds 7, 10, 11 and 14 also checked these): <ul><li>every rank-scaled talent constant, rank by rank (`talents.go:583-690`): Arcane Subtlety 15%, Arcane Impact 2%, Arcane Mind 2% / 20%, Arcane Instability 1% / 1%, Arcane Meditation 17/33/50, Arcane Concentration 2% a rank with the 1 s proc recovery;</li><li>Missile Barrage 20/40% and 10/20%;</li><li>the pre-pull Frostbolt at -2.5 s, which needs Improved Frostbolt 5/5: both builds take it;</li><li>Missile Barrage on a miss and Clearcasting kept through a Barrage (round 13 logs).</li></ul> | talents/mage.json; engine | Aligned | - | - |

## Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **The leveling toll point: Impact or Burning Soul (SIM-6).** The engine models neither, so the sim ties. Forever sources: Warcraft Tavern's Forever Fire guide calls Burning Soul "vital to most Fire Builds, for both dungeon and solo play" (https://www.warcrafttavern.com/forever/guides/fire-mage/, no date shown, read 2026-11-03). Mobalytics' Forever level-30 build takes two Burning Soul and no Impact (https://mobalytics.gg/wow-forever/classes/fire-mage-guide, search summary 2026-11-03; the page returned 403 to a direct fetch). Classic's Icy Veins path takes Impact first (https://icy-veins.com/wow-classic/fire-mage-leveling-talent-build-from-1-to-60), but Forever sources come first. The leveling build moves its Impact point to Burning Soul: `51521001230133031` → `51520101230133031`. Accelerant stays. It speeds the Fireball filler until Scorch at 22 (band 20 builds take it at level 15), so it is not a dead point. | fire.md `build:`; vitest `Fire levels with Burning Soul, not Impact` failed, then passed | CONFIRMED (consensus pick, no DPS change) | old, guides-handoff (SIM-6) | S 071e7b59e |
| F2 | **Fire's negative haste weight (SIM-1).** It is the stale published weight, not a model fault (round 14 M2: 673.6 to 674.5 across 0 to 10% cast speed). Every Fire cast sits on the unhasted 1.5 s global: Scorch, the three-stack Pyroblast and the instant Fire Blast. No cast-timing code changed since (`git log 5195a8889..d1d9f2775 -- sim/core/cast.go sim/mage`: only the batch fix and Combustion's stack cap). | `core/cast.go:254` | Aligned (MR25) | - | - |
| F3 | Fire's raid BiS file still carries Accelerant 1 and Impact 1 (`...-51521001230133031-...`) where the guide's raid build has Burning Soul 2. No DPS difference. The refresh re-ranks it (round 14 X2, SIM-4). | bis/mage-fire.json | Measure at refresh | - | - |
| F4 | Re-read against the 70755 rows, all aligned: <ul><li>Critical Mass 2%, Incineration 2%, Improved Flamestrike 5%, Accelerant 0.1 s, Improved Fire Blast 1 s, Burning Soul 10% threat a rank;</li><li>Fire Power 1/2/3% per effect, cap 4/8/12% (`fire_power.go`); Ignite 8% a rank; Master of Elements 10/20/30%;</li><li>Improved Scorch 33/67/100%; Heating Up -25%, 3 stacks, 20 s;</li><li>Combustion starts at one stack, ends on 3 non-periodic crits, caps at 10;</li><li>the pre-pull Pyroblast at -6 s.</li></ul> | talents/mage.json; engine | Aligned | - | - |

## Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Three Fire points did nothing for Frost (SIM-5).** Accelerant 1 speeds Fireball, Frostfire Bolt and Arcane Wrath. Flame Throwing 2 adds range to Fire spells. Frost casts none of these. Incineration 5 already opens Fire's second row, so neither was a prerequisite. The round-7 raid-build pass called Flame Throwing "solo-leveling (Fire range on a mage who casts Arcane and Frost)" for Arcane and dropped it there. The points move to **Arcane Impact 3** (+6% Arcane crit, Arcane Missiles included). Frost casts about 11 free Missile Barrage Arcane Missiles a fight, 15% of its damage, and their crit goes from 15.3% to 20.7%. The build is legal: Arcane has 12 points before tier 2. New string `230005203/5/253510130101301221` (15/5/31), the same for build and raidBuild. It costs nothing, keeps every cast, and gains +0.41 to +0.43% at raid 60. Flame Throwing also leaves Frost's line in raid-keeps.json. | frost.md `build:`/`raidBuild:`; vitest `Frost spends no point on a talent for spells it never casts` failed, then passed; breakdown (Arcane Missiles 96.8 → 99.5 DPS) | CONFIRMED-BUG | old (round 12 adopted the build; talent-search 2026-10 kept Flame Throwing by a keep-list entry), guides-handoff (SIM-5) | S 79fb5536d |
| R2 | Frost Warding 2, Frostbite 1, Permafrost 1 and Ice Block 1 also do nothing in the raid fight. They are row fillers: without them Ice Barrier's tier 6 is short of 30 Frost points, and no Frost talent left does damage at 30 yd (Improved Cone of Cold and Arctic Reach serve a cone the caster never reaches). | talents/mage.json tiers | Aligned (fillers) | - | - |
| R3 | Re-read against the 70755 rows, all aligned: <ul><li>Fingers of Frost: 15% at both ranks, 1 or 2 charges (`fingers_of_frost.go:43-54`); Shatter 17/33/50;</li><li>Winter's Chill 20% a rank; Ice Shards 20% a rank; Frost Channeling 5% / 10%; Improved Frostbolt 0.1 s; Precision 2% on every school;</li><li>Ice Barrier 15% + 5%;</li><li>the pre-pull Ice Barrier at -4 s and Frostbolt at -2.5 s.</li></ul> | talents/mage.json; engine | Aligned | - | - |
| R4 | Frost's `expected_idle` reason for Shoot (5019) says "Arcane Blast costs no mana and sits ahead of the wand line". Frost's rotation casts no Arcane Blast. This is engineering text, not player text, but it is wrong. | data/curated/apl/mage-frost.json `expected_idle[0]` | Aligned (stale text; guide lane) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Mana gems and Demonic Rune share a 2 minute cooldown (SIM-3).** The client says so. ItemEffect 1.60.1.70755: spells 5405, 10052, 10057 and 10058 (the four gems), 16666 (Demonic Rune) and 27869 (Dark Rune) are all SpellCategoryID 1153 with CategoryCoolDownMSec 120000. SpellCategories puts the same spells and the Robe's 18385 in category 30, where SpellCooldowns gives each a 60,000 ms CategoryRecoveryTime. The Robe's own item cooldown is 300,000. The engine matches (`core/character.go:675`, `mage/mana_gem.go:62`, round 14 M4). | wago ItemEffect, SpellCategories and SpellCooldowns csv at 70755 (cached in `r15/mage/dl/`) | Aligned | - | - |
| X2 | With the rune on the gems' timer, a geared mage never uses its rune. A Mana Agate (375 to 425) takes the second slot, because each item fires once the pool is short by its own maximum. Frost even uses the Agate before the Ruby (Agate 1.00, Ruby 0.81 uses a fight). I tried two policies. "Lesser gems give way to a carried rune": Arcane +0.06%, Frost -0.09%, Fire -0.45% (the rune waits for a 1,500 deficit and fires only 0.34 times a fight). "Give way only when the rune would fire now": no change. Round 9 M2 found the same before the rune shared the timer: waiting for a bigger deficit costs what it saves. The policy is a tie or worse, so the engine stays. | prototype measured, then reverted | Aligned (policy measured) | - | - |

## Before/after

Raid 60, 10,000 iterations. Before: `$R/sb-before` / `$R/before.tsv`. After: the lane engine (E d77b35469) plus site (S 071e7b59e).

| Spec | Faction | s1 before → after | s1 % | s50001 before → after | s50001 % |
|---|---|---|---|---|---|
| mage-arcane (A1) | alliance | 614.2 → 613.4 | -0.13 | 613.8 → 614.0 | +0.03 |
| mage-arcane (A1) | horde | 619.2 → 619.4 | +0.03 | 619.4 → 619.9 | +0.08 |
| mage-frost (R1), raid | alliance | 628.0 → 630.6 | +0.41 | 626.6 → 629.3 | +0.43 |
| mage-frost (R1), raid | horde | 632.3 → 635.0 | +0.43 | 632.0 → 634.6 | +0.41 |
| mage-frost (R1), bare 60 | alliance | 365.3 → 367.0 | +0.47 | 365.4 → 367.1 | +0.47 |
| mage-frost (R1), bare 60 | horde | 367.2 → 368.9 | +0.46 | 367.1 → 368.8 | +0.46 |

- Frost was run with `-talents` at old and new strings. The published BiS raid entry is the old guide string, so the "before" matches before.tsv to the decimal.
- Fire does not move: F1 is a talent the engine does not model.
- No spec outside mage moves. No core or shared code changed.

## Files to regenerate

- `data/builds/1.60.1.70755/addon-builds.json`: the Frost and Fire guide builds changed. `TestCommittedAddonBuildsFollowTheGuides` is red until it is regenerated.
- `sim/request/testdata/ladder/mage-frost.golden.md` and `mage-fire.golden.md`: the ladder walks the guide builds. `web/src/lib/guides/leveling-builds.test.ts` (mage-frost) is red until they are regenerated. Also `mage-arcane.golden.md`, if the ladder's level-60 rung moves with A1.
- Coach starters and addon-data for mage, if they carry the build strings.

## Measure at refresh

- Frost bare and raid bands on `230005203/5/253510130101301221` (R1); Fire bands on the new leveling build (F1, no DPS change); Fire raid on the guide raid build (F3, SIM-4); Arcane bare 60 on the guide build (SIM-4).
- Fire's haste weight should come out near 0 (F2, MR25).

## For other lanes

- None. The gem/rune policy (X2) is mage code and stays as is.

## For the guide lane

- **frost.md:66** (land/guides-r1). Replace "Twelve go into Arcane for **Arcane Concentration** (a free spell 10% of the time) and Missile Barrage. The last eight go into Fire: five ranks of **Incineration**, which adds 10% crit to Ice Lance, two of Flame Throwing for 6 more yards on Fire Blast, and one in Accelerant." with "Fifteen go into Arcane for **Arcane Concentration** (a free spell 10% of the time), Missile Barrage and three ranks of **Arcane Impact**, 6% more crit on your free Arcane Missiles. The last five go into Fire for five ranks of **Incineration**, which adds 10% crit to Ice Lance."
- **frost.md, leveling paragraph:** the point order after 40 now ends in Arcane Impact instead of Flame Throwing and Accelerant. Read the regenerated addon-builds.json for the exact levels.
- **fire.md:69.** Replace "The 60 build pays that toll with a point in **Impact**, a 3% chance to stun a mob for 2 seconds while you farm, and a point in Accelerant." with "The 60 build pays that toll with a point in **Burning Soul**, a 23% chance to keep a cast from being pushed back while mobs hit you, and a point in Accelerant, which speeds up Fireball until you have Scorch."
- **fire.md:109.** Replace "and Impact now and then stops a caster or a runner for you" with "and Burning Soul keeps some of your casts from being pushed back when a mob reaches you".
- **arcane.md, frost.md, fire.md mana lines (MAGE-11):** the claim fix pass 1 removed as unverified is verified (X1). "Mana gems and Demonic Runes share one 2 minute cooldown" is right, from ItemEffect category 1153 at 120,000 ms. The Robe of the Archmage adds a 1 minute lockout with them (category 30) and has its own 5 minutes.
- **mage-frost.json `expected_idle[0]`** (R4): the Shoot reason names Arcane Blast. Suggested: "Frostbolt is cheap enough under Frost Channeling that this geared character never drops below the wand line's 20% mana gate; the level 20 and 30 rungs show it firing."

## Wave 2: leveling orders (guide-accuracy round 1, O2)

Method:
- Checkpoint builds come from `LadderTalentString` (scratch test, not committed), at every level the ladder or ladder-order checks: 10, 20, 25, 30, 35, 38, 40, 45, 50, 55 and 60.
- Spec-breakdown, bare preset, 10,000 iterations, seeds 1 and 50001, both factions.
- Bands 10 and 25 have no BiS entry, so they ran from a scratch `-bis-dir`. Band 10 wears the ladder's gear (main hand 284287, ranged 286750). Band 25 wears the band-20 set on both sides.
- Learn levels (trainables 70755): Arcane Blast rank 1 (400574) at 20, Ice Lance rank 1 (1312002) at 20. Both are 10-point-row talents, so the 11th point at level 20 can hold either.
- All three candidates move a checkpoint, so they belong in ladder-first.json, which needs a measured gain. None is a ladder-order.json case.

| # | candidate | checkpoints it moves | result (a s1 / a s50001 / h s1 / h s50001) | verdict | fix sha |
|---|---|---|---|---|---|
| W1 | Frost: Improved Frostbolt before Frost Warding | 10 only (Frost Warding 1 → Improved Frostbolt 1); 20 and up identical | band 10: 10.5 → 10.8, 10.5 → 10.8, 10.0 → 10.2, 10.0 → 10.2 (+2.0 to +2.9%) | **Adopted.** The level-10 point was armor on Frost Armor; now it speeds the Frostbolt the rotation casts. | S 1e702cd9f |
| W2 | Frost: Ice Lance at 20 | 20 (Ice Shards 1 → Ice Lance) and 25 (Frostbite 1 → Ice Lance) | band 20: -0.44 / -0.44 / -0.45 / -0.67%; band 25: +0.19 / 0.00 / 0.00 / -0.19% | **Rejected: it loses at 20.** The rotation casts Ice Lance only on a Fingers of Frost charge, a level-36 talent, so the point is idle while Ice Shards' crit damage is lost. Its real leveling use, a Frost Nova shatter on a frozen mob, is outside the sim's single-target fight. Under the wave rule a loss at any band rejects it. The guide may still mention the Nova combo. | - |
| W3 | Arcane: Arcane Blast at 20 (Blast first) | 20 (Arcane Concentration 5 → 4 plus Blast) and 25 | band 20: -2.9 to -3.6%; band 25: +17.5 to +20.7% | Rejected as written (loses at 20). | - |
| W3b | Arcane: Arcane Concentration 5, then Arcane Blast | 20 (Arcane Resilience 1 → Blast; Concentration stays at 5) and 25 (Arcane Impact 2 → 1 plus Blast); 30 and up identical | band 20: 44.3 → 44.4, 44.3 → 44.4, 44.1 → 44.2, 44.2 → 44.3 (+0.23% each); band 25: 46.6 → 56.0, 46.5 → 56.1, 47.5 → 55.8, 47.5 → 56.0 (+17.5 to +20.7%) | **Adopted.** Blast at 20 in place of an armor point. The plain walk had no Blast until 27, which is where the band-25 gain comes from. | S 1e702cd9f |
| W3c | Arcane: Concentration 5, Missile Barrage 2, then Blast | 20 (Missile Barrage 1 in place of Arcane Resilience) | band 20: -17.0 to -19.0% | Rejected. | - |

Tests:
- `sim/leveling/ladder_first_mage_test.go` failed, then passed.
- `TestCuratedLadderFirstIsSound` now knows mage-arcane and mage-frost.
- The sim/leveling package is green.

**Files to regenerate (wave 2, in addition to the list above):** the mage-arcane and mage-frost ladder goldens at rungs 10 and 20, and addon-builds.json (the Arcane and Frost point orders). `web/src/lib/guides/leveling-builds.test.ts` fails for mage-arcane, mage-fire and mage-frost until the goldens are regenerated. The TS walk already reads the new rows.

**For the guide lane (wave 2):** frost.md's leveling paragraph should read Improved Frostbolt at 10 to 14, then Frost Warding. arcane.md:107 "Missile Barrage comes at 21 and Arcane Blast at 27" becomes Arcane Blast at 20; read the regenerated addon-builds.json for the exact levels. Ice Lance stays at 30 in the order.

## Wave 3: rank rows on the level-gap penalty

Setup:
- Branches merged with int/r15: site 1e96d8df2, fork 9c511a2cd, which carry the priest lane's E 8d9ca6aae and 974614b8b.
- The rule: a rank keeps its whole spell power coefficient until the caster is 17 levels past the rank's learn level. After that it loses 5% a level.
- Learn levels: spellranks.json 70755.
- Runs: 10,000 iterations, seeds 1 and 50001, both factions. Sweeps at 3,000.

Rows and fixed-rank casts:

| # | row or cast | gap at the band it covers | result | verdict | fix sha |
|---|---|---|---|---|---|
| V1 | Fire Scorch rank 1 (2948, learned 22) at 40-49 | 18 at 40 (-5%), 27 at 49 (-50%) | Sweep at bare 40, horde s1: rank 1 155.0, rank 2 155.0, rank 3 151.4, rank 4 (default) 147.5. At 10,000 iterations, rank 2 against rank 1: alliance 157.3 → 158.4 / 157.4 → 158.6 (+0.7%), horde 154.7 → 154.7 / 154.8 → 155.1 (+0.0 / +0.2%). | **Re-picked to rank 2** (8444, learned 28). It ties or gains at 40 and keeps its whole coefficient to 45, where rank 1 has already lost 30%. Test `Scorch at 40` failed, then passed. | S c094ddd7f |
| V2 | Fire bare 50, after the priest lane dropped Scorch rank 1 at 50 | - | Scorch rank 2 / 3 / 4 at 50: 181.2 / 189.7 / 188.4, against 189.8 for the default rank 5. Fire Blast rank 5 / 4 at 50: 189.1 / 187.6, against 189.8. | No new row: no lower rank gains. | - |
| V3 | Fire Scorch rank 1 at 30-39; Fire Blast rank 3 (22) at 30-39 and rank 4 (30) at 40-49 | ≤ 17 at the bands they cover | Bare 30 identical to the round start (73.6 / 73.7 / 72.0 / 72.0). | Kept. The engine change does not reach them. | - |
| V4 | Arcane Frostbolt rank 4 (7322, learned 20) at 30-39 | 10 at 30 (19 at 39, -10%) | Bare 30 identical to the round start (88.0 / 88.1 / 86.4 / 86.6). | Kept. The band runs at 30, where it is unpenalised. | - |
| V5 | Fixed-rank APL casts, all three specs: Frostbolt 10, Arcane Missiles 7, Arcane Blast 5, Fireball 11, Fire Blast 7, Scorch 7, Pyroblast 8, Cone of Cold 5, Ice Lance 6, Ice Barrier 4 | Top trainer ranks, learned at 54 to 60. Below 60 the level rewrite casts the highest learned rank. | Raid 60 unchanged against wave 1. | Aligned: no top rank is penalised at 60. | - |

Before (round start, `$R/sb-before`) → engine change with the old rows → final:

| band | faction | s1 | s50001 |
|---|---|---|---|
| Fire bare 40 | alliance | 159.4 → 157.3 → 158.4 (-0.6%) | 159.6 → 157.4 → 158.6 (-0.6%) |
| Fire bare 40 | horde | 156.9 → 154.7 → 154.7 (-1.4%) | 156.9 → 154.8 → 155.1 (-1.1%) |
| Fire bare 50 | alliance | 202.0 → 195.8 (-3.1%) | 202.1 → 195.9 (-3.1%) |
| Fire bare 50 | horde | 195.5 → 189.5 (-3.1%) | 195.7 → 189.6 (-3.1%) |
| Fire / Arcane bare 30 | both | unchanged | unchanged |
| Fire raid 60 | both | unchanged (666.4, 673.0) | unchanged |
| Arcane, Frost raid 60 | both | same as wave 1 (Before/after above) | same |

- Fire bare 50's -3.1% is the engine change: the rank-1 Scorch it used to cast is gone, and no lower rank recovers any of it.
- Bare 40 recovers 0.5% for alliance and up to 0.2% for horde.

**Files to regenerate:** the mage-fire ladder golden (rung 40 casts Scorch rank 2).

## Wave 4: guide-accuracy round 2 seeds

Setup:
- Branches merged with int/r15: site 0d9241e0c, fork 86fe4bf34.
- Live hotfix check: wago SpellEffect at 1.60.1.70755, fetched 2026-11-04. Every mage effect this wave relies on matches our spellconst: Arcane Wrath 1326082, Eureka! 1259817, Arcane Blast 1239700 and its stack 400573, Fire Blast 10199, Pyroblast 18809 and 11366, Fireball 10151, Arcane Missiles 10212 and 10274, Frostbolt 10181, Scorch 10207, Heating Up 400624 and 400625, Arcane Power, Presence of Mind, Incineration, Accelerant, Ignite, Ice Lance 1240047. No live difference to record.

| # | finding | evidence | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | **Arcane Wrath in the engine** matches the client (MAGE-R2-1). Rank 4 is 451 + 3.0 a level, ±20%, coefficient 1.0, 3.5 s cast, 6 s cooldown, 395 mana. Effect 1 is the 5% maximum-mana share, applied as spell damage. | `arcane_wrath.go`; constants; live SpellEffect | Aligned | - | - |
| G2 | **Arcane Wrath on the old Arcane/Frost raid build** (SIM-R2-1). The point comes from Frost Channeling 1 or Incineration 1. Wand Specialization cannot give it: without it the build is illegal, because Arcane Subtlety needs 5 points above. Best line: Presence of Mind on Arcane Wrath, 616.7 to 624.3 against 613.4 to 619.9 (+0.6 to +0.7%, 10,000 iterations, both seeds). Arcane Wrath at three Blast stacks: +0.3%. | runs | Held under the adoption rule (<1%) | - | - |
| G3 | **The Arcane raid build that takes Fire second** (SIM-R2-2). Build `131205213100211531/525000012001/003`: <ul><li>Arcane 32, as before, but Arcane Meditation 2 and Arcane Wrath;</li><li>Fire 16: Incineration 5, Accelerant 2, Ignite 5, Pyroblast, Improved Fire Blast 2, Heating Up;</li><li>Frost: Precision 3.</li></ul>The rotation adds, each gated on Heating Up or the Arcane Wrath talent so the leveling build is untouched: <ul><li>Pyroblast on three Heating Up stacks, and on Presence of Mind;</li><li>Fire Blast on cooldown;</li><li>Arcane Wrath at three Blast stacks;</li><li>Fireball as the filler, ahead of Frostbolt.</li></ul>Raid 60 (alliance / horde, seeds 1 and 50001): 613.4 / 614.0 / 619.4 / 619.9 → 657.1 / 657.2 / 659.1 / 659.1 (+6.4 to +7.1%). Without Arcane Wrath (Arcane Meditation 3): 645.4 / 644.8 / 648.7 / 648.3. Bare 30 to 60 on the leveling build: unchanged. Adoption rule: <ul><li>the Arcane tree holds;</li><li>every unmodeled talent stays (Improved Channeling, Arcane Resilience, Arcane Subtlety, Arcane Geometry);</li><li>the dropped Frost points (Improved Frostbolt, Ice Shards, Frost Channeling) go with the Frostbolt they served, and Frostbolt stays as the fallback line;</li><li>the gain is beyond error at both seeds.</li></ul>Sources: Mobalytics Verified Arcane ("the expected raiding build for Arcane mages", 32/16/3) and Icy Veins Forever Arcane (Fire subspec), both cited in guide-accuracy round 2. | `TestArcaneRaidPlaysItsFireSecondTree` failed, then passed; apl-check green; builds vitest green | CONFIRMED-BUG | old, guides-handoff (SIM-R2-2) | S e01dc2dff, E b46cc2fe4 |
| G4 | **Eureka!'s mask** reaches Arcane Blast, Arcane Missiles (cast and missiles), Frostbolt, Fireball, Fire Blast, Scorch, Pyroblast, Ice Lance and Cone of Cold, for +10% damage and -10% cost. It does not reach Arcane Wrath or any periodic damage. The client says the same: <ul><li>effect 1's mask (551688947, 4096, 0, 8) misses Arcane Wrath's family bits (0x40000000, 0x8000);</li><li>effect 2, the periodic modifier, has an all-zero mask, so Pyroblast's and Fireball's burns and Ignite get nothing.</li></ul>The tooltip says "Your damaging abilities", which is broader, but the client table decides. | live SpellEffect 1259817; `core/eureka_auto_gen.go` | Aligned (client mask; tooltip broader) | - | - |

**Measure at refresh:** Arcane raid 60 gear, weights and tiers on the Fire second tree. Fire spell power and hit now count, and Kindling Stave's case changes.

**Files to regenerate:** the mage-arcane ladder golden at rung 60 (if the ladder uses the raid build), addon-builds.json (the Arcane raid order), and the coach starter for Arcane.

**For the guide lane (wave 4):**
- arcane.md's Talents section must describe the raid build's 16 Fire points:
  - Incineration 5 (Fire Blast and Arcane Blast crit), Accelerant 2 (faster Fireball);
  - Ignite 5, Pyroblast, Improved Fire Blast 2 and Heating Up (instant Pyroblasts from Fire Blast);
  - Arcane Wrath for one rank of Arcane Meditation;
  - Precision 3 in Frost.
- The raid loop: Pyroblast on three Heating Up and with Presence of Mind, Fire Blast on cooldown, Arcane Wrath at three Blast stacks, Fireball as the filler.
- The leveling build keeps Frost and Frostbolt.
- The new rotation lines have no notes. Suggested:
  - Pyroblast: "Pyroblast on three Heating Up stacks, when it is a 1.5 second cast, or with Presence of Mind."
  - Fire Blast: "Fire Blast on cooldown: every hit is a Heating Up stack."
  - Arcane Wrath: "Arcane Wrath on cooldown once you hold three Blast stacks, so their spell damage rides on it."
  - Fireball: "Fireball fills in the raid build; its crits add Heating Up."
- The Presence of Mind note ("spend it on a 2.5 second Frostbolt") becomes "spend it on Pyroblast" for the raid build.
- The threat line should say that Fire spells get no Burning Soul cut.

## Still open

- Arcane Concentration from Shoot and Kindling Stave's Ignition (round 14 M7), and Kindling Stave's proc rate (M5): carried, with no new evidence.
- Ice Lance's coefficient; cost modifiers adding; Master of Elements on a Clearcasting crit: carried (round 14 M8). The beta logs stop near level 30, too low for Master of Elements.
- Warcraft Logs is blocked (gcloud login expired). Not worked around.
