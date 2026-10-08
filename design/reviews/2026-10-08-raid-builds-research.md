# Raid builds and rotations, research pass, 2026-10-08

Lane `research` (site branch `research`, fork branch `research`, neither merged or pushed). Engine: the fork worktree at `b4fa0ac3f` plus the synced rotations, client 1.60.1.70009. This pass redoes the 2026-10-07 raid-build search on the final engine (Tidal Focus, Suppression, Life Tap by Spirit, Elemental Fury on totems, Cataclysm, additive Venom are all in it) and applies the rule to the twenty DPS specs. A stalled earlier attempt left its reports mixed (some written before its own curated edits, some after); every report below was regenerated from a clean tree, and the four rotations it had edited were reverted and only the changes the new reports justify were put back.

## The rules applied

- Raid build: a `talent-search -preset raid -keep <raid-keeps.json>` winner is adopted when it holds the spec's tree majority, keeps every talent the rotation casts or gates on (the builds test's signature table), keeps the protections in `talent-search/raid-keeps.json`, drops only utility the engine cannot model, and gains at least 1% and is beyond the search's combined error. A winner is adopted and searched again from itself until the search names no further qualifying move.
- Rotation: a `rotation-search -preset raid -build-code <raid build>` line is adopted only beyond error; a line that drops an ability only when its model was verified against the client in a review; and, as on 2026-10-07, only if it leaves the bare guide build level (the curated rotation also serves the bare entry and the ladder).

The talent search reads the rotation from the embedded `sim/request/apl` copies, so the binary was rebuilt after every `apl-sync`. Searches took 4 to 9 seconds each; no search was killed and none is unsearched.

## Per spec

Builds are the FS1 digits (tree 1 / tree 2 / tree 3). Gain is the talent search's, against the previous raid build on the rotation as finally curated, in percent of the previous build's DPS; "within" means inside the combined error.

| Spec | Previous raid build | New raid build | Moved (added / dropped) | Gain | Rotation change | Hold, with reason |
|---|---|---|---|---|---|---|
| druid-balance | `4132220115501051/05/5053` | same | none | - | none | deep Balance 35/0/16 (+2 Genesis, +3 Reflection, -5 Heart of the Wild) is +0.7%: under the 1% bar |
| druid-feral | `05002/45211031021032212001/5053` | same | none | - | none | the guide is the best build found |
| hunter-beast-mastery | `5420001505001251/0053502001/4` | same | none | - | none | the guide is the best; rotation swap of Serpent Sting and Multi-Shot is +0.4% (see Held rotation lines) |
| hunter-marksmanship | `5320000501/005155000150305/5` | same | none | - | none | best finalist (deep Beast Mastery, 30/21/0) is +0.8% and not Marksmanship-majority |
| hunter-survival | `0/32500500005/500230131051120151` | same | none | - | none | deep Survival is +0.2%: under the bar |
| mage-arcane | `153005113100011531/03/05450003` | same | none | - | none | the guide is the best |
| mage-fire | `2050151/23252100130133051/005` | same | none | - | none | the guide is the best |
| mage-frost | `203005/13102/255510032100030025` | same | none | - | none | best is +0.3%, within error, and drops Permafrost |
| paladin-retribution | `54003/053/0520533100133032` | `550030031001/0/0522533100133032` | +1 Divine Intellect, +3 Reverence, +1 Purifying Power, +1 Divine Favor, +2 Holy Conduit / -5 Redoubt, -3 Precision | +3.2% | Consecration added (below) | second pass from the new build finds +0.5% (`550030030001/0/05225331001330321`): under the bar |
| priest-shadow | `3250010313/0/523120501201300251` | same | none | - | none | best is +0.0% |
| rogue-assassination | `01532310421501/315303000015/002` | same | none | - | none | clean best is the guide; the +0.6% variant drops Remorseless Attacks, which is outside the classified utility |
| rogue-combat | `005320105/31530300001515231/002` | same | none | - | none | the guide is the best |
| rogue-subtlety | `005323101014/0/5323220310013011031` | same | none | - | none; Eviscerate removal held | deep Assassination 29/0/22 is +2.1% but Assassination holds the majority, not Subtlety (22), and drops Camouflage; with every unmodeled talent kept the guide is the best |
| shaman-elemental | `5530311300103051/052/053302` | `5530311300103051/02/053352` | +5 Tidal Focus / -3 Thundering Strikes, -2 Ancestral Knowledge | +3.7% | none | second pass: the guide is the best; deep Restoration (+2 Restorative Totems) is lower than this build |
| shaman-enhancement | `32303/255130030005102051/052` (the leveling build) | `3230031/255030031005102031/053` | +3 Elemental Devastation, +1 Elemental Focus, +1 Shamanistic Focus, +1 Mindfulness / -3 Call of Flame, -1 Guardian Totems, -2 Maelstrom Weapon | +8.6% (with Earth Shock in the loop) | Earth Shock added (below) | second pass: the guide is the best |
| warlock-affliction | `25550300100201351/00052/0550001` | same | none | - | none | the guide is the best |
| warlock-demonology | `0553/203511310122000135/005003` | `055/203511311112000135/0050051` | +1 Improved Sayaad, +2 Aftermath, +1 Ruin / -3 Malediction, -1 Master Summoner | +5.7% | none | second pass: +0.3%, within error (Improved Health Funnel to Improved Sayaad drops an unmodeled talent) |
| warlock-destruction | `255/0005/2053045103101351` | same | none | - | none | the guide is the best |
| warrior-arms | `02305213032515001/55050000001/2` | same | none | - | none | best is +0.7%: under the bar |
| warrior-fury | `34320003002/05153105022011501/2` | same | none | - | none | best is +1.0%, within error |

Elemental takes Tidal Focus 5/5 (the Restoration talent is +1% spell hit a rank, applied by the engine as `stats.Hit`; see `2026-10-08-tidal-focus.md`), but the earlier expectation of a Totemic Focus swap did not hold: the search's two candidates are deep Elemental with Tidal Focus and deep Restoration, and the first is better. Enhancement's gain exists only because Earth Shock is back in the loop: Elemental Devastation does nothing without a shock, and the same build scores the same as the old one on the old rotation. Demonology's build moved for Searing Pain as the filler: Improved Shadow Bolt is no longer taken by the raid build, Malediction (a modeled talent) leaves for Aftermath, Ruin and Improved Sayaad.

Reports: `design/reviews/talent-search/<spec>-raid.md` (all twenty regenerated).

## Rotation changes

Two lines adopted, both in `data/curated/apl`, synced to the site and fork copies and checked by `make apl-check`.

| Spec | Line | Raid-ready gain | Bare guide build | Model check | Decision |
|---|---|---|---|---|---|
| Retribution | Consecration (rank 5) last, while mana is above 65% | +8.9% (ungated: +10.5%) | level (254.3 before, 253.6 after, error 0.5) | the client puts the damage on a companion row per rank; `TestConsecrationDamageMatchesClient` holds the engine's tick damage, extra damage on the first targets, coefficient and target count to those rows | **adopted** (this closes the 2026-10-07 hold, which was for want of a client-stated damage) |
| Enhancement | Earth Shock (rank 7) last, while mana is above 50% | +18.7% on the old build, from 498.2 to 591.3 together with the new build | level (226.4 before, 227.0 after, error 0.8) | `TestEarthShockDamageMatchesClient` (and Flame and Frost Shock) | **adopted** |

The mana gate is the guard on the bare entry: ungated, Consecration costs the bare character 8.5% and Earth Shock 5%. Gates tried (raid / bare, percent of the ungated raid result): Consecration 20% 100 / -5.3%, 35% 100 / -4.0%, 50% 99.2 / -1.5%, 65% 98.6 / -0.3%; Earth Shock 20% 100 / -5.3%, 40% 99.7 / -1.2%, 50% 96.3 / level, 60% 94.9 / level. The Flame Shock line the stalled attempt had also added is not in the Enhancement search winner and is not adopted.

### Held rotation lines

| Spec | Line | Gain | Reason |
|---|---|---|---|
| Retribution, Enhancement | drop the mana gate | +1.5%, +0.5% | the bare character loses 8.5% and 5% |
| Demonology, Destruction | remove the potion line (mana below 75% in the execute phase) | +1.0%, part of +2.1% | drops an ability whose model is not verified in any review; the engine also drinks potions on cooldown |
| Destruction | replace Bane of Doom with Curse of the Elements | +9.3 DPS | the 2026-10-08 hold stands: the raid preset's own Curse of the Elements collides with the cast one, and the action probe prices Bane of Doom at +51 |
| Destruction | Shoot gate from 40% to 50% | +0.8 DPS alone, within error | the +5.4 in the search was measured after two other mutations |
| Subtlety | remove Eviscerate (and then Ghostly Strike) | +3.6% | the search itself marks the winner "not adoptable" (it never casts Eviscerate), and Eviscerate's attack power term is the unverified part of its model (`2026-10-07-rogue-parity.md`); tables otherwise match the client |
| Beast Mastery | swap Serpent Sting and Multi-Shot | +0.4% | outside this lane's set of specs (build and engine models unchanged); not applied |
| Marksmanship | remove Serpent Sting | +1.2% | same, and a drop with no model review |
| Survival | Lacerate refresh below 3 s | +0.4% | wins the raid, loses the bare build (2026-10-07 precedent) |
| Shadow | Shadow Word: Death without its execute gate | +3.4% | unpriced self-damage (2026-10-07 hold) |
| Fury | swap Execute order; remove Death Wish | +1.5% | outside the lane's set; the drop has no model review |
| Elemental, Mage Fire (Frost Nova) | swap Lava Burst and Mana Spring; insert Frost Nova | +0.2%, +0.6% | within error |

Every other rotation report reads "the curated rotation is already the best found". Reports: `design/reviews/rotation-search/<spec>-raid.md` (all twenty, regenerated with `-build-code <raid build>` on the final builds).

## BiS ranker, before and after

Band 60, one `leveling-bis -spec <spec> -bands 60` run per spec per side into a scratch directory: before is the clean `HEAD` of the branch (cut from `main`) built and run on its own tree, after is this branch, both on the fork worktree engine. Set DPS of the published picks.

| Spec | Raid A before | Raid A after | Delta | Raid H before | Raid H after | Delta | Bare A before | Bare A after | Delta |
|---|---|---|---|---|---|---|---|---|---|
| paladin-retribution | 538.5 | 579.6 | +7.6% | 551.5 | 583.1 | +5.7% | 254.5 | 257.0 | +1.0% |
| shaman-elemental | 472.8 | 484.1 | +2.4% | 472.7 | 481.4 | +1.8% | 157.5 | 157.5 | +0.0% |
| shaman-enhancement | 515.0 | 580.4 | +12.7% | 523.2 | 596.7 | +14.0% | 226.0 | 227.9 | +0.8% |
| warlock-demonology | 792.6 | 837.8 | +5.7% | 784.0 | 828.4 | +5.7% | 418.2 | 418.2 | +0.0% |

The ranker re-picks gear for the new build, so the raid columns can differ from the fixed-gear searches above (Enhancement before is 515 on the ranker's own gear, 498 on the search's); the bare column is the guard.

## Checks

`go test ./...` under `sim/` (ladder goldens regenerated for Retribution and Enhancement), `cd data && uv run pytest tests/test_apl.py`, `make apl-check ENGINE_DIR=<fork worktree>`, `npx prettier --check` on the four touched guides and `npx vitest run src/content src/lib/guides`: results in the lane's final message. `sim/go.mod`'s replace points at the fork worktree locally and is not committed; the nightly owns the generated BiS files, which are not touched.

## Follow-ups

- Engine: Elemental Devastation and Shamanistic Focus changed Enhancement by tens of DPS once a shock is cast; worth a conformance review against the client rows (the probe credits both as zero on a rotation with no shock).
- Held lines above that need a model review before a drop is adoptable: the Potion line, Eviscerate's attack power term.
