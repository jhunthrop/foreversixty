# Sim-accuracy round 16: priest

This is lane priest, working on site wt/r16-priest and fork lane/r16-priest, both cut from int/r16 (site c62d77ab8, fork 9f7f1f6e4). The client is 1.60.1.70755.

Specs: Discipline, Holy and Shadow, plus the shared low-rank penalty (core.LevelPenalty).

Every sim ran through the lock wrapper on frozen BiS (`$R/bis-frozen`). Exploration runs used 3,000 iterations. Reported numbers use 10,000 iterations at seeds 1 and 50001.

## 1. The low-rank penalty: what changed, and the test the owner can run

### What the second look found

**No launch-day logs exist yet.**
- foreverlogs.gg: its public report ids stop at 2813, the beta "shield" reports of 2026-10-29. Ids from 2814 to 4000 return "Report not found". The list endpoints now need an API key.
- api.foreversixty.gg/v1/reports/recent holds one beta report (2026-10-20).

**One new independent measurement, from HealRankAdvisor.** This is a Forever addon (github.com/mendeleevforever/HealRankAdvisor, commits 2026-10-14 to 10-24). It measures each rank's spell-power share from the in-game tooltip at two Spell Healing values.
- **Setup:** a level-30 priest, at 16 and 151 Spell Healing.
- **What it says:** the learn-level + 17 formula "matched the in-game tooltips for 9 of 10 ranks tested. Lesser Heal rank 1 kept about 24% where the formula predicts 40%."

**ForeverChanges' own table has the same anomaly** (foreverchanges.pro/downranking, 2026-10-16):
- "Healing Touch rank 1 and Holy Light rank 1, learned at level 1, run one level ahead of it, as if learned at level 0."
- Its cells: Healing Touch 1 is 85% at 20, 79% at 21, 75% at 22 and 34% at 30. Holy Light 1 is 94% at 18 and 85% at 20.

**One client field explains all three outliers: SpellLevels.MaxLevel** (wago 70755).
- Almost every heal rank has MaxLevel = learn level + 5. For those ranks, "17 levels past the learn level" is the same as "12 levels past MaxLevel".
- The three ranks that ran early are exactly the ones whose MaxLevel is lower:
  - Lesser Heal 1 (2050): MaxLevel 3.
  - Healing Touch 1 (5185) and Holy Light 1 (635): MaxLevel 5.
- "MaxLevel + 12" fits each of them:
  - Lesser Heal 1 at 30: 25%. Measured 24%; the old rule gave 40%.
  - Healing Touch 1: 85% at 20, 80% at 21, 75% at 22, 35% at 30.
  - Holy Light 1: 95% at 18, 85% at 20.
- It still fits every other measured rank: Rejuvenation 1 and 2, Healing Touch 2, Regrowth 1, and the other nine HealRankAdvisor ranks.
- It also explains Ozgar's "16 levels". Ozgar cites no measurement, and a flat 16 contradicts Rejuvenation 1, which is still 100% at 21.

**Where the evidence stops:**
- **Healing Wave 1** (MaxLevel 5) fits neither rule: 100% at 20, then down 6 to 8 points a level. ForeverChanges calls it unexplained. It is left open.
- **The top-rank exemption** (Ozgar) gets no support. If HealRankAdvisor's ten ranks were a level-30 priest's ten heal ranks (Lesser Heal 1-3, Heal 1-3, Renew 1-4), then Lesser Heal 3, the top Lesser Heal rank, was penalised as the formula says. That is likely, not certain.
- **Damage spells and levels above 30** are still assumed.

### The fix (CONFIRMED on two measurements; E 18815822c)

The share is now `clamp(1 - (L - MaxLevel - 12) x 0.05, 0, 1)`.
- MaxLevel is the rank's client SpellLevels.MaxLevel.
- Where the client has no other MaxLevel, the learn level + 5 stands in.

How it is built:
- A generated table holds the 1,165 exceptions: `sim/core/spell_max_level_auto_gen.go`, written by `tools/spellmaxlevel/gen.py` (`make spellmaxlevel`, now a step in tools/follow.sh).
- Wrack (MaxLevel 60) is never below the caster on its own row, so its `ScalesToCasterLevel` flag is gone.
- Conformance now mirrors the client row's own MaxLevel.
- Conformance also stopped reading a share of 0 as "no penalty". Before, 163 rows across six classes showed a false "differs"; now all 771 declared rows match.
- Tests:
  - `TestSpellLevelPenaltyKeysOnTheRanksMaxLevel` and `TestACastRankTakesItsMaxLevelShare` failed (build), then passed.
  - `TestAFullyPenalisedRankExpectsNoCoefficient` failed, then passed.

What it moves:
- **Raid 60:** nothing, for any of the 28 specs, both factions, both seeds (identical to `$R/before.tsv` within 0.005%).
- **Bands** (rank rows only):
  - Holy priest Prayer of Healing rank 1 (MaxLevel 39) keeps 55% at bare 60 instead of 35%, and 100% at 50 instead of 85%.
  - Paladin Holy Light rank 1 keeps 85% at 20 instead of 90%.
  - The other changed rows move only at their band's last level, which no sim runs: mage Frostbolt 4 and Scorch 1 at 39, Scorch 2 at 49, Holy Light 1 at 29, and Prayer of Healing 1 at 49 and 59.

### The in-game test (for the owner)

**Character:** a level-60 priest in any spec. The ranks stay in the spellbook.

**Gear swap:** one item, or a set, that changes Spell Healing by at least 100. Read the character sheet's Spell Healing with the item on (H1) and off (H2).

**What to read:** for each spell below, read the tooltip's heal range both ways, take the mid-point and compute ΔHeal. Do the same for **Greater Heal rank 4**, the reference. It is learned at 58, so every rule keeps it whole.

Share kept = (ΔHeal_rank / coefficient_rank) / (ΔHeal_GH4 / 0.857). Dividing by the reference cancels Spiritual Healing and the other talents.

| Spell (id) | coefficient | learned / MaxLevel | learn + 17 (old) | flat 16 (Ozgar) | MaxLevel + 12 (now) | top rank exempt |
|---|---|---|---|---|---|---|
| Prayer of Healing 1 (596) | 0.286 | 30 / 39 | 35% | 30% | **55%** | 35% |
| Prayer of Healing 2 (996) | 0.286 | 40 / 49 | 85% | 80% | **100%** | 85% |
| Heal 4 (6064) | 0.857 | 34 / 39 | 55% | 50% | 55% | **100%** (top Heal) |
| Lesser Heal 3 (2053) | 0.714 | 10 / 15 | 0% | 0% | 0% | **100%** (top Lesser Heal) |
| Flash Heal 4 (9474) | 0.429 | 38 / 43 | 75% | 70% | 75% | 75% |
| Greater Heal 1 (2060) | 0.857 | 40 / 45 | 85% | 80% | 85% | 85% |

How to read the result:
- **Prayer of Healing 1 and 2** separate MaxLevel from the learn level.
- **Flash Heal 4 and Greater Heal 1** separate 16 from 17.
- **Heal 4 and Lesser Heal 3** test the top-rank exemption. Lesser Heal 3 is the cleanest of the two: 0% against 100%. If it keeps everything, the exemption is real.

**Damage, the same way:** use a Spell Damage or Shadow Damage swap, with Mind Blast rank 9 (10947, learned at 58) as the reference.

| Spell (id) | learned / MaxLevel | learn + 17 | flat 16 | MaxLevel + 12 |
|---|---|---|---|---|
| Mind Flay 3 (17312) | 36 / 44 | 65% | 60% | **80%** |
| Holy Fire 4 (15264) | 36 / 42 | 65% | 60% | **70%** |
| Mind Blast 5 (8105) | 34 / 39 | 55% | 50% | 55% |

**A leveling character, for a quick check:** a priest at level 30 reads Lesser Heal 1 at about 25%. The old rule gives 40% and a flat 16 gives 35%.

**Caveat:** this assumes the tooltip carries the penalty. HealRankAdvisor and ForeverChanges both say it does, and Blizzard's character-sheet text names it.

## Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| SH1 | **Regression targets from round 15 are intact:**<br>- Wand Specialization first (`data/curated/ladder-first.json`). The ladder point order is Wand 10-11, Spirit Tap 12-16, Shadow Focus 17-21, Improved Mind Blast 22-25, Improved Shadow Word: Pain 26-27, Mind Flay 28, Improved Mind Flay 29-30, Blackout 31-32, Vampiric Embrace 33, Shadow Weaving 34-36, Darkness 37-41, Shadowform 42.<br>- The prepull Mind Blast at -1.5 s.<br>- The Starshards 30 s cooldown (fork tests pass).<br>- The raid build (no Blackout or Spirit Tap). | fork priest tests pass; LadderPointOrder on the guide build | Aligned | - | - |
| SH2 | **Band 40 represents level 40 exactly.** Wand-first gives Shadowform at 42, so the band-40 number (-36%) describes two of that band's ten levels; levels 42 to 49 have Shadowform. This is how bands are evaluated, not a priest bug. | ladder | Note (platform) | - | - |
| SH3 | **The guide prose still described round 14's builds.** Round 15's changes had shipped without text:<br>- Spirit Tap 3 / Blackout 4, "from level 21";<br>- Wand Specialization "right after Shadowform";<br>- the old opener;<br>- the raid trades;<br>- "Shadowform is your point at 40".<br>Rewritten to the shipped builds. | shadow.md :34, :61-62 (now :61-63, plus Devouring Contagion, Shadow Reach and Silence lines), :69, :73-75, :127, :133, :135 | CONFIRMED (guide) | R15 925a3b1e0 / 75cae4b93 / 5b46c799d (text not shipped) | S 58bff07ca |
| SH4 | **Fresh re-audit, raid 60 on the guide build** (horde, 3,000 iterations):<br>- Mind Flay 39%, Shadow Word: Pain 20%, Mind Blast 19%, Devouring Plague 12%, Shadow Word: Death 9%.<br>- No time out of mana.<br>- Devouring Plague 7.6 casts in 180 s (24 s cooldown).<br>- Dark Inspiration uptime 99%.<br>- Every priest conformance row (amount, cost, cast, cooldown, coefficient) matches 70755, now with no false "differs". | spec-breakdown; conformance golden | Aligned | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Notes and prose named the round-14 filler.**<br>- The Holy list's two filler notes said "Heal is the filler: rank 1 in a raid". They now sit on Greater Heal rank 4 lines (round 15 W6 moved the filler).<br>- holy.md :26 and :70 said "downranked Greater Heal, rank 1 or 2".<br>Both are rewritten: Greater Heal at the top trainer rank, with Heal 4 as the before-40 and low-mana cast. At raid 60, Heal 4 still fires 17 times a fight when Greater Heal is unaffordable. | `data/curated/apl/priest-holy.json`, holy.md | CONFIRMED (guide) | R15 4faba89a9 (text not shipped) | S 58bff07ca |
| H2 | Prayer of Healing rank 1 rows (50-59, bare 60) gain from the new key: bare 60 +0.15 to +0.27%, band 50 +0.04 to +0.20%. The rows still stand. | sims (below) | Aligned | - | - |
| H3 | **Fresh re-audit:**<br>- Prayer of Mending 24 casts and 79 heals a fight, the best healing per mana (10.6).<br>- Renew rank 7 ticks crit (PERIODIC_CAN_CRIT).<br>- Holy Binding fires off every Flash Heal.<br>- Litany of Light mana.<br>- Conformance matches. | spec-breakdown, raid 60 | Aligned | - | - |

## Discipline

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **Notes and prose contradicted the shipped rotation:**<br>- The Power Word: Shield note said "at a low rank" (the list casts rank 10).<br>- discipline.md:68 said "a full-rank shield on a warrior tank eats the rage he needs". Blizzard's priest deep dive (2026-09-30) says absorbs no longer reduce rage.<br>- The party-filler notes said "Lesser Heal rank 1" and "a heal learned below level 20 keeps its full healing power bonus", on what are now Greater Heal 1 and Lesser Heal 3 lines.<br>All rewritten. | `priest-discipline.json`, discipline.md | CONFIRMED (guide) | R15 ee7be8d87 / 4faba89a9 (text not shipped) | S 58bff07ca |
| D2 | **The raid-60 fight runs dry at 78 s of 300.** After that, the deeper-hole lines cast Lesser Heal 3 (57 a fight, 0% bonus at 60, 2.9 healed per mana).<br>This is the known healer mana model (round 14 C4/F5). It is not reopened. | spec-breakdown | Aligned (known) | - | - |
| D3 | **Fresh re-audit:**<br>- Penance 17 casts a fight, best direct healing per mana (6.0).<br>- Renewed Hope at 4 ranks is 8% (client: 10% at 5).<br>- Soul Warding -4 s.<br>- Conformance matches. | client talents, sims | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **The low-rank penalty key** (section 1). | HealRankAdvisor; ForeverChanges; SpellLevels 70755 | CONFIRMED-BUG | old (round 15 W1 keyed it on the learn level) | E 18815822c |
| X2 | **Conformance read a share of 0 as "no penalty"** (`expectedCoefficient`). It showed 163 false "differs" rows: Starfire 1, Fireball 5, Smite 1 and others at 60. | conformance SUMMARY before/after | CONFIRMED-BUG (harness) | R15 8d9ca6aae | E 18815822c |
| X3 | **The web Leveling strip orders chips by `bandSpendOrder`** (the planner's reconstruction of the level-60 build) instead of the ladder. Shadow lists Blackout ahead of Wand Specialization and Spirit Tap. `LevelingBandStrip.astro` builds `spendOrder` from `bandSpendOrder(fullBuild)` and ignores `ladder-first.json`; the band talent strings themselves are right (`guideLevelingBuilds` uses the ladder). This is web-only, so it goes to the platform lane. | `web/src/components/guides/LevelingBandStrip.astro:40` | CONFIRMED (web) | old | handed to platform |
| X4 | **The druid conformance golden was already stale on int/r16** (Starfire rank 1 amounts). It was regenerated together with this change; the Starfire part is not explained by it. | conformance before this change | Note (druid lane) | - | E 18815822c |

## Before/after

Raid 60, all 28 specs, 10,000 iterations, seeds 1 and 50001, both factions, guide builds on frozen BiS: **no change** against `$R/before.tsv`.

Bands (bare, 10,000 iterations; before = `$R/sb-before`):

| spec, band | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|
| priest-holy, bare 60 | 386.6 → 387.2 (+0.16%) | +0.15% | 401.8 → 402.9 (+0.27%) | +0.25% |
| priest-holy, 50 | 232.1 → 232.2 (+0.04%) | +0.09% | 244.7 → 245.1 (+0.16%) | +0.20% |
| paladin-holy, 20 | 49.6 → 48.8 (-1.61%) | -1.61% | 49.7 → 48.8 (-1.81%) | -1.81% |

The site commits (notes and prose) move no number.

## Files to regenerate

- **Fork goldens, regenerated** (explained by E 18815822c): `sim/core/testdata/conformance/{druid,mage,paladin,priest,shaman,warlock}.golden.md` and `SUMMARY.md`.
- **Site, not regenerated:**
  - `data/builds/1.60.1.70755/addon-data.json`, then `addon/ForeverSixty/Data.lua`. The rendered Priority notes come from these, so the new Holy and Discipline notes appear only after the regen.
  - `addon-builds.json` (Shadow order, owed since round 15).
  - TestRotationLadder goldens for priest-holy (50, bare 60) and paladin-holy (20).
  - dps-per-point for those.
- **APL copies:** `sim/request/apl/priest-{holy,discipline}.apl.json` were updated by hand (notes only). The fork's ui copies carry no notes. `make apl-check` passes.

## Measure at refresh

- priest-holy at 50 and bare 60, and paladin-holy at 20: the new penalty key.

## For other lanes

- **Shared core change, loud:** E 18815822c changes `core.LevelPenalty`'s key for every class.
  - Raid 60 moves nothing anywhere.
  - **paladin:** the Holy Light rank 1 row at 20-29 now keeps 85% at 20. Band 20 moves -1.6 to -1.8%; re-check the row against the default rank.
  - **mage:** the Frostbolt 4 (Arcane, to 39) and Scorch 1/2 (Fire, to 39/49) rows lose 5% at their band's last level only.
  - **warlock:** `ScalesToCasterLevel` is removed (`sim/warlock/wrack.go`, comment only); Wrack is exempt through its MaxLevel 60. Demonic Brand (MaxLevel 60) and other scaling spells are now exempt the same way.
  - **paladin:** the Holy Shock heal (1311605, MaxLevel 37) keeps 45% at 60 instead of 35% (Holy casts rank 4, so no number moves).
  - **shaman:** `talents.go` AtLightningOverloadCoefficient still calls `core.LevelPenalty(RequiredLevel, …)`. Lightning Bolt ranks have MaxLevel = learn level + 5, so nothing changes; `core.SpellLevelPenalty(id, …)` is the call to use if that changes.
  - **Test files touched for the API:** `sim/paladin/holy/spells_test.go`, `sim/druid/restoration/spells_test.go`, `sim/priest/healing/spells_test.go` (`SpellLevelPenalty` with the spell id).
- **platform:** X3, the Leveling strip chip order.
- **druid:** X4, the stale Starfire golden.

## For the guide lane

Done in this lane: shadow.md, holy.md, discipline.md, and the Holy and Discipline notes. The wow-player verdict is below.

Not mine, but checked against the new key:
- shaman/restoration.md:76 ("17 levels past the level you learned it") holds for every shaman heal rank it names. Healing Wave 1 (MaxLevel 5) is the only shaman heal whose key differs.

## Still open

- Healing Wave rank 1's measured curve fits no rule.
- The top-rank exemption and damage spells above 30: the owner test in section 1 settles both.
- foreverlogs.gg has no launch logs yet, and its list API now needs a key. Warcraft Logs is still blocked.
- The pending-owner low-rank tooltip items (round 15) are covered by the section-1 test and are not reopened.

## Sources

- HealRankAdvisor (Forever addon), README and HealRankAdvisor.lua, commits 2026-10-14 to 10-24: https://github.com/mendeleevforever/HealRankAdvisor
- ForeverChanges, "WoW Forever downranking", updated 2026-10-16: https://foreverchanges.pro/downranking
- Ozgar's Downranking Guide & Tool (undated; no measurement cited): https://wowdownrank.forever.ozgar.se/
- Forever Upgrades, 2026-10-11, quoting Blizzard's 1, 7 and 8 October notes: https://foreverupgrades.com/news-article.html?a=downranking-nerfed-less-spell-power-fewer-procs
- ForeverWisp, 2026-10-10: https://www.foreverwisp.com/guides/wow-forever-downranking-spell-ranks
- SpellLevels, build 1.60.1.70755: https://wago.tools/db2/SpellLevels/csv?build=1.60.1.70755

## wow-player verdict

**SHIP**, on the second pass over the rendered shadow, holy and discipline pages (dev server, 1440 wide).

Applied from the first pass:
- The Discipline Mana line no longer calls the shield cheap.
- The shield wording is now "no longer cuts a warrior's rage gain".
- The Shadow opener's Mind Blast lands after the tank's first hit, and the Threat paragraph says the opener spends Fade.
- Spirit Tap, Devouring Contagion and Shadow Reach take the reviewer's wording.
- The Lesser Heal 3 and Greater Heal 1 notes were rewritten.
- The shield prose now carries the mana-pace condition.

Left as SHOULDs:
- The generated key-talent table lacks Dark Inspiration and shows Silent Resolve 1/3 (platform).
- The Discipline Priority list's 24 alternating Greater Heal 1 / Lesser Heal 3 rows are measured lines. Collapsing them is a later list change.
- After the addon-data regeneration, the integrator should check that the Holy rows no longer say "Heal is the filler".

## Wave 2: rank rows level by level, phantom talent lines

I merged int/r16 first: site c5915e4bb (fast-forward) and fork 65c8e36cf.

**Method.** I swept every priest rank row with `spec-breakdown -level` at the band floor, +3, +6 and +9 (3,000 iterations, horde, seed 1). Each candidate rank is cast through its own spell id, so no row overrides it. Every switch point was then confirmed at 10,000 iterations, seeds 1 and 50001. Where a row is race-keyed, both factions were confirmed. The sweep files are in the lane scratch (`w2/`).

| # | row | finding | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W2-1 | Holy Renew rank 1, 14-39 | Rank 1 (MaxLevel 13) loses its bonus from 26.<br>- **Rank 1 still wins to 31:** 68.8 against rank 3's 67.2 at 31.<br>- **From 33, rank 3 wins:** 69.5 against rank 1's 67.8 at 33, where rank 2 ties at 69.5.<br>- **At 34:** rank 3 69.2 against rank 2's 68.3. At 36: 70.8 against 68.6.<br>- Seeds 1 and 50001 are identical.<br>- **New rows:** rank 1 at 14-32, rank 3 at 33-39. | CONFIRMED | R15 4faba89a9 (floor-only measurement) | S 297b20ef7 |
| W2-2 | Discipline Renew rank 1, 14-39 | Rank 1 wins through band 20.<br>- **From 30, rank 2 wins:** 90.0 against 88.4 at 30, and 90.1 against 87.7 at 33.<br>- **Rank 3 takes over from 34:** it ties rank 2 at 34, then wins 90.4 against 88.9 at 36.<br>- **New rows:** rank 1 at 14-29, rank 2 at 30-33, rank 3 at 34-39.<br>- **Band 30 floor:** 87.8 to 89.4 alliance and 88.4 to 90.0 horde, both seeds (+1.8%). | CONFIRMED | R15 4faba89a9 | S 297b20ef7 |
| W2-3 | Shadow Devouring Plague rank 1 at 50-59 (five races) | Rank 1 (learned at 20) loses its bonus from 38.<br>- **From 52, rank 2 wins on both factions:** horde 251.5 against 247.2 at 52 and 256.2 against 252.0 at 56; alliance 223.4 against 217.5 at 52 and 228.0 against 223.8 at 56. Seeds 1 and 50001 agree.<br>- At 50, rank 1 and rank 2 tie. At 59, rank 3 ties rank 2.<br>- **New rows:** the race rows end at 51 (rank 1 for five races, rank 2 for Night Elf); rank 2 for every race at 52-59.<br>- The band floor (50) does not move. | CONFIRMED | R15 4faba89a9 (row kept at the floor) | S 297b20ef7 |
| W2-4 | The other priest rows: Shadow Word: Pain rank 3 at 30-39; Mind Blast rank 3/4 at 40-49; Devouring Plague rank 1 at 30-49; Shadow Word: Death rank 2 at 50-59; Prayer of Healing rank 1 at 40-59 and bare 60 (Holy), and at 40-49 (Discipline) | Each row's rank wins or ties at every level swept.<br>- Mind Blast ranks 4, 5 and 7 are within 0.4% from 40 to 49.<br>- Holy bare 60 Prayer of Healing: rank 2 404.7 against rank 1 403.1, inside error.<br>Kept. | Aligned | - | - |
| W2-5 | **Phantom lines** | None in priest.<br>- **priest-shadow line 3** (Vampiric Embrace), which the shaman lane flagged, already gates on `spellIsKnown` 15286, and it casts the talent spell itself.<br>- **Every other priest line** that needs a talent casts the talent's own spell: Inner Focus, Penance, Prayer of Mending, Mind Flay. An unlearned spell simply does not fire. Shadow line 5's Inner Focus sequence is followed by a plain Mind Blast line, so it adds no phantom cast.<br>- **Discipline line 7** reads Weakened Soul, a debuff, not a talent aura. | Aligned (flag withdrawn) | - | - |

- **Tests:** `TestPriestRenewRowsSwitchInsideBandThirty` and `TestShadowDevouringPlagueMovesToRankTwoAtFiftyTwo` (sim/rankchoices) failed on the old rows and now pass.
- **Guide text:** no guide prose or rendered card note names these ranks. The rows' `note` field is internal, so no page changed and no wow-player pass was needed.
- **Addon:** the addon resolves ranks only at its rungs (shaman lane, platform item). Discipline's 30-33 row reaches it at rung 30. The rows at 33/34 and 52 reach it only at the next rung.
- **Pre-existing failure, not mine:** `data/tests/test_addonrotation.py::test_the_holy_paladin_band_40_flash_of_light_pair_exports_the_ranks_it_casts` fails on int/r16 without these changes.
- **Files to regenerate:**
  - addon-data and then Data.lua (rank rows);
  - TestRotationLadder goldens: priest-discipline band 30;
  - dps-per-point for priest-discipline.
