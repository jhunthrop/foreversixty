# Sim-accuracy round 7: integration (2026-11-02)

The owner's round-7 rulings, the integration, the pin and the review. The lane reports beside this file have the per-spec rows: warlock.md, casters.md, physical.md, hybrids.md, healers.md, pipeline.md, ci.md, page.md and ranker.md.

## Owner rulings, and where each landed

| # | ruling | outcome | where |
|---|---|---|---|
| R1 | Hawk Eye's points go to Deadly Aspects in the Survival guide builds | **Done.** The leveling build moves Hawk Eye 3 to Deadly Aspects (+4.2 to +4.6%). The raid build already had Deadly Aspects 5/5, so its 3 points went to Lightning Reflexes 5/5 and Rapid Killing 2/2, the best of ten measured placements. No other range-only talent remains in a Survival build. A local re-rank lands Survival on Deadly Aspects 5/5 at bands 50 and 60, with no Hawk Eye or Improved Concussive Shot. | physical.md S2; S 123c7b26f |
| R2 | The Wowhead refresh carries saved records by itself; the floor applies to the carried catalogue | **Done.** fetch-wowhead always carries the saved record of an item missing from the live feed. The 90% floor (`MIN_CARRIED_COVERAGE`) is on the live-plus-carried catalogue. A record is lost only when its item leaves the client's Item table; it is listed in a `dropped` field and counts against the floor, so a truncated Item table still fails, naming the ids. Each run logs its coverage, carried and dropped counts. `--allow-shrink` is gone from fetch-wowhead, and data.yml no longer passes it there. The `allow_shrink` dispatch input remains only for normalize's deliberate item re-baselines. There are four new tests. | pipeline.md; S febd69046 |
| R3a | A CI job builds the wasm and runs the four real-engine browser tests | **Done.** `real-engine` in web.yml builds sim.wasm the way deploy does (pinned engine, sim/go.mod's toolchain, `make publish-wasm`), builds the real-data site and runs `sim-real-engine`, `sim-gear-real-engine`, `sim-weights-real-engine` and `sim-talents-real-engine`. It fails on a missing or wrong-pin wasm, and on any skipped test. Proven locally: 4 passed, and a wrong wasm id fails the job. The first GitHub run is checked after the push (below). | ci.md; S cd90cbc22 |
| R3b | Room in the /sim bundle by trimming code | **Done.** The rotation notes load only when the closed "what it does" drawer or a /sim/specs card mounts. sim-island went from 87,518 to 74,834 bytes gzipped; the 92,160 budget is unchanged. The ledger's "12 bytes" was the character page's chunk, which has 2 bytes left (18,430 of 18,432). That chunk is the shared paperdoll and is still open. | page.md P4; S 2b6201a56 |
| R3c | The 36 "An unnamed spell" rows | **Fixed.** sync-data publishes the client's whole spell and item name tables in 64 shards, fetched only for ids on screen. `engine-spell-names.json` names the 20 engine-only ids, Ignite 12654 among them. `TestEveryPublishedResultActionHasAName` sims all 28 specs at every published band on both factions (576 runs). From the review: the two set-bonus procs that showed the client's internal "S03 - Mana Proc on Cast" name now read "Sudden Insight", and a test fails on any internal name. | page.md P1, P9; S 69cf0b6c8, 36136a78e |
| R3d | Ignite and wand shots in their true schools | **Already right at the pin**, now held by a test. Ignite reads Fire; Shoot reads the wand's school. Two display bugs sat beside it, both fixed: crit rates were rounded per worker before the merge (Ignite read 0.0% crit), and zero-damage rows counted as Physical. | page.md P2, P5, P6 |
| R4 | Check bis run 41363895819's picks | **Superseded.** The owner cancelled the run and held every BiS and data refresh until a zero round. The picks were measured locally and not committed. Torch of Light is gone from every band-60 warlock set: Brilliant Wand replaces it, and bare Demonology takes Greater Eternal Wand. Survival lands on Deadly Aspects. | warlock.md B1; physical.md |
| - | Do not act on the 14 dropped loot sources; record them | **Recorded, not acted on.** No upstream feed changed. Our two loot generators disagree about battlegrounds. data.yml's full `loot` reads the client's map table, which marks Alterac Valley (30) and Arathi Basin (529) as battlegrounds, so their creature and object drops get no source. The nightly `loot-merge` files the same drops as world sources. loot.json flips on every pair of runs. The net -14 is 15 Alterac Valley world sources (Grunnda Wolfheart among them) plus renamed world-drop bands. 23 items lose every source: AV turn-ins and consumables, Horn of the Frostwolf Howler, the tabard, the ice ammo and two quivers, and Arathi Basin's Balia'mah Trophy. No published BiS set, leveling band, guide or preset references any of them. PvP ranks 5 and 10 did not drop; only their `faction_source` and `reitemised_from` change. Which generator is right is the owner's call; the full table is in pipeline.md. | pipeline.md item 2 |

## Integrator findings

| # | finding | evidence | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | **data's rank test was red on main since c0bd0bd3d.** The layered 70707 push relabels Lacerate (1235827) and Shred (22829) as rank 5. The APLs still said 3 and 4. The engine resolves a cast by id, so no number moves. | test_apl.py; spellconst druid.json at c0bd0bd3d | CONFIRMED (test red on main) | S 69bcf6c87, E 61c37c227 |
| I2 | **spellranks' Greater Blessing of Light override was dead and its test red on main.** The push removes 25890 and round 6 removed it from the engine. Deleted from spellranks and from the ladder's mirror. | TestSingleTierOverridesGateTheirIDs, TestLadderLearnLevelOverridesMirrorTheRewrites | CONFIRMED (test red on main) | S 69bcf6c87, 3b7b7aef9 |
| I3 | Two lanes fixed Mark of the Chosen. The hybrids version, with its test, stays. | int/r7 merge | merge | E (merge of lane/r7-physical) |
| I4 | The warlock and page merge left a raid-consumes level row without its buffs; regenerated from its generator. | TestRaidConsumesJSONIsCommitted | merge | S 69bcf6c87 |
| I5 | Four fork goldens move under hybrids and healers together (TestP1Mage, TestProtection, TestElemental, TestEnhancement). | fork suite | test-only | E (goldens commit on int/r7) |

Follow-ups the integrator sent back to lanes, all fixed in this round:
- The ranker's empty off hand (hybrids X3) went to a new ranker lane.
- The feral attack power in `items/<class>.json` (physical X1a).
- Mana Spring at the wrong rank below 60 (healers).

## Integration

- **Fork `int/r7`** merges lane/r7-casters, -warlock, -hybrids, -physical and -healers, plus I1's APL copies. The page lane made no engine change. The suite passes with `-tags=with_db`, except `sim/web`, which also fails setup at the old pin.
- **The pin is cd4dfdcc0.** data/proto/ENGINE_SHA follows it; the proto is unchanged. The rotation ladder goldens are regenerated at it with `make simdb` and a scratch go.work.
- **The site** merges wt/r7-ci, -pipeline, -casters, -warlock, -hybrids, -physical, -healers, -page and -ranker.
- **Gates, all exit 0:**
  - `go vet` and `go test ./...` in sim
  - the coach and sims packages in api
  - `make apl-check` against the pin
  - `make styles-check`
  - the full data pytest
  - web sync, the full vitest, astro check, lint and format:check
  - the sim, tiers, planner-share-sim and header-finished-sim e2e specs
  - The CI lane proved the four real-engine specs locally.

## Before/after

Pin b57765fb9 with origin/main (9cfb91c47), against pin cd4dfdcc0 with this branch. Both run on today's published data: the BiS sets and builds, and the simdb built from the committed inputs. Raid preset, band 60, 10,000 iterations. Healers are in effective HPS on the heal profile, tanks in threat per second.

| Spec | Alliance, seed 1 | Alliance % (1 / 50001) | Horde, seed 1 | Horde % (1 / 50001) |
|---|---|---|---|---|
| druid-balance | 565.7 to 565.9 | +0.0 / +0.1 | 567.2 to 567.3 | +0.0 / +0.0 |
| druid-feral-bear | 877.0 to 864.6 | -1.4 / -1.4 | 885.3 to 868.6 | -1.9 / -1.9 |
| druid-feral | 624.1 to 620.9 | -0.5 / -0.5 | 619.6 to 616.3 | -0.5 / -0.5 |
| druid-restoration | 743.1 to 772.7 | +4.0 / +4.0 | 740.6 to 770.3 | +4.0 / +4.1 |
| hunter-beast-mastery | 669.3 to 670.9 | +0.2 / +0.2 | 673.4 to 674.8 | +0.2 / +0.3 |
| hunter-marksmanship | 661.6 to 662.7 | +0.2 / +0.1 | 666.1 to 666.7 | +0.1 / +0.1 |
| hunter-survival | 775.0 to 769.3 | -0.7 / -0.7 | 778.4 to 772.6 | -0.7 / -0.7 |
| mage-arcane | 586.9 to 589.5 | +0.4 / +0.6 | 589.6 to 592.5 | +0.5 / +0.6 |
| mage-fire | 648.3 to 653.2 | +0.8 / +0.7 | 656.7 to 660.0 | +0.5 / +0.5 |
| mage-frost | 584.3 to 583.1 | -0.2 / -0.1 | 590.1 to 589.8 | -0.1 / -0.1 |
| paladin-holy | 580.1 to 656.9 | +13.2 / +13.4 | 563.7 to 642.0 | +13.9 / +14.1 |
| paladin-protection | 992.7 to 995.3 | +0.3 / +0.4 | 1008.2 to 1011.5 | +0.3 / +0.4 |
| paladin-retribution | 666.2 to 662.4 | -0.6 / -0.7 | 664.5 to 659.6 | -0.7 / -0.8 |
| priest-discipline | 601.1 to 610.0 | +1.5 / +1.4 | 609.3 to 614.2 | +0.8 / +0.8 |
| priest-holy | 680.4 to 682.8 | +0.4 / +0.4 | 663.2 to 670.8 | +1.1 / +1.1 |
| priest-shadow | 707.8 to 707.8 | +0.0 / -0.0 | 713.9 to 714.1 | +0.0 / -0.0 |
| rogue-assassination | 621.9 to 623.5 | +0.3 / +0.3 | 619.0 to 620.6 | +0.3 / +0.3 |
| rogue-combat | 622.7 to 618.3 | -0.7 / -0.7 | 622.8 to 618.3 | -0.7 / -0.7 |
| rogue-subtlety | 558.4 to 542.2 | -2.9 / -2.9 | 558.3 to 542.0 | -2.9 / -2.9 |
| shaman-elemental | 459.4 to 465.1 | +1.2 / +1.3 | 467.6 to 473.0 | +1.2 / +1.1 |
| shaman-enhancement | 608.0 to 600.2 | -1.3 / -1.3 | 607.3 to 599.7 | -1.3 / -1.3 |
| shaman-restoration | 653.6 to 823.3 | +26.0 / +26.0 | 650.7 to 820.7 | +26.1 / +26.2 |
| warlock-affliction | 563.7 to 573.2 | +1.7 / +1.9 | 564.9 to 574.5 | +1.7 / +1.7 |
| warlock-demonology | 627.0 to 639.2 | +1.9 / +1.9 | 629.6 to 642.5 | +2.0 / +2.1 |
| warlock-destruction | 626.1 to 636.4 | +1.6 / +1.6 | 628.9 to 639.1 | +1.6 / +1.6 |
| warrior-arms | 728.3 to 718.0 | -1.4 / -1.5 | 740.4 to 729.9 | -1.4 / -1.5 |
| warrior-fury | 771.3 to 764.1 | -0.9 / -0.9 | 775.0 to 767.8 | -0.9 / -0.9 |
| warrior-protection | 1067.7 to 1055.5 | -1.1 / -1.2 | 1055.9 to 1043.7 | -1.2 / -1.1 |

Where the big moves come from:
- **Strength of Earth Totem** gave Era's 61 Strength. Forever's rank 4 is 42 (SpellEffect 10441). The Ahn'Qiraj book's 53 stays out while IncludeAQ is off. This is about -1.2% for every melee spec on the raid preset: warriors, rogues, Survival, Enhancement, Retribution, Feral, and bear and warrior threat.
- **Mana Spring** was cancelled by Blessing of Wisdom on every raid-preset run since 86d33a100. This lifts every mana user by 0.5 to 1.3%.
- **Restoration shaman:** Healing Stream's healing was never counted, and Water Shield never procced on raid pulses.
- **Holy paladin:** Oath of Light, which Forever's push adds, did nothing.
- **Healers in general:** raid pulses now push back casts. Every healer column also carries the heal profile's retune (raid damage 1.65 to 1.85), because after the fixes the strongest raid set covered 97% of the profile, past its 60 to 85% rule. healers.md has each healer at both 1.65 and 1.85 to separate the two.
- **Warlocks:** the mana potion line waited for the execute phase. Destruction also takes Elixir of Shadow Power beside Fire Power.
- **Subtlety:** Hemorrhage uses normalized weapon damage.
- **Fury and Assassination:** Dual Wield Specialization now reaches Whirlwind's and Mutilate's off-hand hits.
- **Bear:** Mark of the Chosen is 21 stats on a 2-minute proc.

What this table does not hold. Each needs data.yml's regeneration and the bis.yml re-rank, both held by the owner until a zero round.

**Item data:**
- Three Wowhead-supplement wands still Physical. Pyric Caduceus is the band-50 caster wand; band 50 bare moves Fire +4.6%, Shadow +3.5%, Frost +3.0%, Arcane +1.0% and the warlocks.
- Held-in-off-hand items no longer replacing the main hand.
- Feral attack power on 6 items.

**Re-ranks:**
- One-hand casters get their off hand (Balance, Elemental, Shadow priest and two warlocks, +1.2 to +3.3% in the local ranker).
- The new Holy and Discipline priest builds and the Survival builds publish.
- The heal profile's coverage check repeats on the re-picked sets.
- /tiers/healer reorders. On local runs: shaman, druid, Holy priest, Holy paladin, Discipline.

## wow-player review

Verdicts on the lane reports, the before/after table and the page shots:
- **SHIP:** warlock, casters, pipeline and ci.
- **SHIP WITH FIXES:** physical, hybrids, healers and page. Each fix was settled before the push:
  - **The 3-damage, 0-hit Shoot row on Fire 60.** It is a real average: the wand fires about once in 100 fights. It now reads "<1" hits and "<0.1%" (S b6b9aa7b7). At the new pin Fire 60 no longer fires the wand.
  - **Internal spell names.** Fixed: the two published ids read "Sudden Insight", with a test (S 36136a78e).
  - **The ranker bug's owner.** It was fixed in this round by the ranker lane (S 5d34ae14e).
  - **Prose quoting Strength of Earth as 61 or 77.** None exists.
  - **The raid preset and the Ahn'Qiraj books.** `core.IncludeAQ` is false, so the books stay off.
  - **The healer move's attribution.** Split in healers.md at 1.65 and 1.85.
  - **Survival's separate raid build at +0.7% over leveling.** The reviewer would not publish two builds for 0.7%. The integrator keeps it: the owner's ruling changed it, it is the best measured build at 60, and folding it into the leveling build is a separate choice. It is listed for the owner.
- **Copy suggestions, not sim bugs, left for the page owners:**
  - a note on /tiers/healer that the healer order rests on a model with no level-60 log yet
  - a line saying the Survival raid number fell because of Strength of Earth, not the build
  - ability icons on the Damage tab (a ux-designer change)
- **Plausibility:** the pattern is coherent: melee -0.5 to -1.5%, casters flat to +2%, healers up. Restoration shaman +26% and Subtlety 13% behind Combat are the two numbers players will question first. Neither is a bug.

## Still open

- **Owner decisions:**
  - the 14 loot sources and which loot generator is right
  - Survival's separate raid build
  - the about 413 client spells with internal names that no published result shows
  - the copy suggestions above
- **Held under the adoption rule:**
  - Marksmanship +0.9%
  - Affliction's Improved Bane of Agony swap +0.7%
  - Destruction's Immolate below Conflagrate +0.1 to +0.2%
  - Fire dropping its Improved Scorch upkeep +0.4%
  - Assassination leveling +3.6%, which drops Murder and Venom
  - Discipline taking Inspiration's points (its armor is unmodelled)
- **Blocked on level-60 logs:** every round-6 row, plus the following:
  - Oath of Light: scaled or flat
  - Gift of Nature on Tranquility
  - same-op percent stacking (Rejuvenation, Healing Wave)
  - Fire Power on the Fireball and Pyroblast burns
  - a Fire wand crit and Ignite
  - the wand hit table
  - rage from damage taken by level
  - Heroic Strike as a swing
  - Battle Shout's rounding
  - Eviscerate's attack power
  - Holy Shield crits, Toughness's armor, Holy Strike's flat damage
  - Barkskin in bear, Searing Totem's 2.2 s attack, Fire Nova's range
  - the Ahn'Qiraj book ranks at trainers
- **Single source or open:**
  - Concentration Aura is unmodelled (at most +1.3%)
  - Healing Wave rank 10
  - Chain Heal's party-only jumps
  - Power in Light
  - forever-bugs #419 (shadow wands and Shadow multipliers)
- **Page and pipeline:**
  - the character-page chunk has 2 bytes of headroom
  - sim.wasm.gz is at 94.7% of its 4.5 MiB budget
  - `real-engine` does not gate deploy yet
  - `allow_shrink` remains for normalize's re-baselines
  - client.yml's hand hint still mentions it
