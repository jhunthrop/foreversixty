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

## Still open

- Arcane Concentration from Shoot and Kindling Stave's Ignition (round 14 M7), and Kindling Stave's proc rate (M5): carried, with no new evidence.
- Ice Lance's coefficient; cost modifiers adding; Master of Elements on a Clearcasting crit: carried (round 14 M8). The beta logs stop near level 30, too low for Master of Elements.
- Warcraft Logs is blocked (gcloud login expired). Not worked around.
