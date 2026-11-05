# Sim-accuracy round 21: mage

Specs: mage-arcane, mage-fire, mage-frost. Site wt/r21-mage (from int/r21 b67caae8f). Fork lane/r21-mage (from int/r21 ad183f267). Client 1.60.1.70790.

Method:
- Seed first (Fire filler at 22 to 30, Scorch against Fireball, then the rank rows level by level), then the guide rounds' carried sim requests, then the re-audit.
- Bare bands: spec-breakdown `-preset bare -band 20|30 -level N`, 3,000 iterations to explore, 10,000 at seeds 1 and 50001 for the reported before/after. Raid 60: frozen BiS (`-bis-dir $R/bis-frozen`), guide raid builds (`$R/rt`). Every sim through `$R/sim.sh`.
- Client: data/builds/1.60.1.70790 spellconst, talents, trainables (mage). Logs: api.foreversixty.gg, foreverlogs.gg, Warcraft Logs v1.

**Confirmed and fixed: 3** (all `old`; two guides-handoff): Evocation was castable below its trainer level (X4); Fire filled with Scorch from 22, where Fireball does more for its mana and wins; the Arcane raid order spent Presence of Mind on Frostbolt at 54 to 56 with Pyroblast already known. Plus one stale test expectation on int/r21 (not a sim bug).

## Seed: Fire filler at 22 to 30 (guide round 19 MAGE-R19-9)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F-S1 | **Fire cast Scorch as its filler from 22, before the talents that make Scorch good.**<ul><li>The card's Scorch filler line had no gate, so from Scorch at 22 it replaced Fireball. Rank 1 Scorch (2948) is 39 damage for 50 mana in 1.5 s; Fireball rank 4/5 (3140/8400) is 77+12 for 95 mana and 120+16 for 140 mana: about a fifth more damage per mana. The band-20 mage is mana-bound (out of mana about 66 s of 180).</li><li>Bare band 20, 10,000 iterations, alliance (Human), Scorch filler against Fireball filler, seed 1 / seed 50001: level 22 40.1 / 40.1 to 40.4 / 40.4 (+0.7%); 25 43.6 to 43.9 (+0.7%); 28 45.6 to 46.1 (+1.1%); 29 46.2 to 46.4 (+0.4%). Horde (Undead) 22 38.3 to 38.9 (+1.6%), 28 44.6 to 45.2 (+1.3%), both seeds. Error 0.0 to 0.1 each. 3,000-iteration sweep 22 to 29: Fireball ahead or tied at every level (tie at 26).</li><li>From 30 (Heating Up, both orders take it at 30) Scorch wins by about 10%: 30 77.3 against 69.7, 31 77.9 / 70.4, 32 78.6 / 71.4, 34 89.8 / 81.2, 36 90.7 / 80.8.</li><li>Cross-check on a 60 s fight (no mana limit): Scorch is ahead at 25 to 29 by 0.9 to 2.7% on throughput alone (Incineration's 10% and Ignite), tied at 22. A leveling mage drinks between pulls, so damage per mana decides the 20s; Fireball wins the mana-bound fight and costs a fifth less per point of damage. Classic leveling practice agrees (Fireball until the Fire talents arrive).</li><li>Fix: the Scorch filler line is gated on `auraIsKnown(400625)` (Heating Up); Fireball is the filler below it. The Improved Scorch stack line still casts Scorch at 27 to 29. Level 60 raid is identical to the decimal (horde 721.0 / 721.4 before and after).</li><li>Text: both card notes, fire.md:116 (leveling) and fire.md:66 (Accelerant "until Heating Up at 30"), the APL's expected_idle reason and engineering note 4. The Scorch note builds on guide round 19's wording.</li></ul> | `$R/mage/fire-filler-3k.txt`, `fire-filler-10k.txt`; `TestFireScorchFillerWaitsForHeatingUp` (sim/request) failed, then passed | CONFIRMED-BUG | old (the 2026-10-07 Scorch filler round measured 40, 50 and 60 only) + guides-handoff (MAGE-R19-9) | S 6c6a348b9, E 422104846 (fork APL copy) |
| F-S2 | **Rank rows at 22 to 29 on the new filler, level by level** (3,000 iterations, band 20, row rank against the trained rank): Fireball rank 4 (trained) 40.4 against rank 3 39.6 at 22, 42.4 / 41.7 at 23; rank 5 against rank 4 at 24 43.5 / 42.8, 26 44.5 / 43.5, 27 45.1 / 43.8, 29 46.4 / 45.1. Fire Blast rank 3 against rank 2 at 22 40.4 / 40.1, 26 44.5 / 43.4, 29 46.4 / 45.0. Scorch rank 2 against rank 1 at 28 and 29: tie (Scorch is only the stack line there). The trained rank wins everywhere; no new row. | `$R/mage/sb -rank` | Aligned | - | - |
| F-S3 | **Fire rows at 30 to 39 judged level by level** (they had only a per-band check): Scorch rank 1 / Fire Blast rank 3 against Scorch 2, Scorch 3 and Fire Blast 4 (5 at 39): 30 77.3 against 76.0 / 76.2; 33 79.3 against 78.7 / 78.7; 36 90.7 against 88.9 / 86.1 / 89.6; 39 98.1 against 97.1 / 94.4 / 96.1 / 94.6. Both rows hold at floor, +3, +6 and +9. | same | Aligned | - | - |

## Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Frostbolt rank 4 row at 30 to 33, level by level: 30 91.7 against rank 5 90.6; 32 98.0 against rank 5 96.8 and rank 6 95.9; 33 98.4 / 97.0 / 96.6; at 34 the trained rank 103.4 against rank 4 103.1 (row correctly ends at 33). | `$R/mage/sb -rank 10181=N`, band 30 | Aligned | - | - |
| A2 | **Arcane Missiles first against Frostbolt first, 10 to 22, level by level** (guide round 18 Sims requested 3; integrator follow-up). Measured on spec-breakdown's new below-band levels (X5): the band-20 entry at levels 10 to 19 wearing only the items that level allows, the leveling order truncated level by level as `-talents`, 10,000 iterations, alliance, shipped card against the card without the Missiles-first line, seed 1 / seed 50001: 10 5.3 / 5.2; 12 5.7 / 5.6; 13 5.8 / 5.7; **14 7.7 / 7.9** (both seeds); 15 9.5 / 8.8; 16 22.8 / 20.0; 18 25.9 / 20.9; 19 26.8 / 20.9; 20 47.9 / 33.1; 21 48.1 / 33.4; 22 49.0 / 41.4 (s50001 within 0.1 of s1 everywhere). Missiles first holds at every level but 14, where Frostbolt leads by 2.6% for one level (Frostbolt rank 3 trains at 14, Missiles rank 2 at 16, and Clearcasting's first point comes at 15). No gate expresses that one level (a gate on not knowing Arcane Concentration would also move 10 to 13, where Missiles wins by 2%), so the card stays. Runs at 10 to 19 were made after X4. | `$R/mage/arc-missfirst-10k.txt`, `arc-lev-talents.tsv` | Aligned (one-level exception at 14 recorded) | - | - |
| A3 | Round 20's two Arcane engine fixes re-read for regressions: the Presence of Mind Pyroblast keeps Heating Up only when the cast time is 0 (three stacks leave 1.5 s, Accelerant does not touch Pyroblast, no haste reaches 0); Clearcasting's Barrage check is set as the channel applies and read only for Arcane Missiles, so a later non-Barrage cast spends the charge. Raid 60 breakdown: Arcane Power 1.44 casts, Presence of Mind 1.44, Pyroblast 11.2, Missile Barrage 13.8 procs, Clearcasting (12577) 12.7. | `heating_up.go`, `talents.go`, `arcane_missiles.go` (E fd5c65756, 18355c631) | Aligned | - | - |
| A4 | Spells and talents against 70790: the mage spellconst, talents and trainables files are byte-identical to 70755 with the build label masked, and hotfix-push.json carries no rows (search totals 0 for every table), so round 20's line-by-line reading (A5) stands. Re-read this round from the talent rows: Arcane Blast (12% Intellect, +100% cost, 4 stacks, 18 s), Arcane Wrath (5% of maximum mana), Missile Barrage (40% Blast, 20% the rest, -50% channel, -100% cost), Arcane Impact 6%, Arcane Instability 3%/3%, Improved Channeling (pushback only), Accelerant 0.5 s at five ranks on Fireball, Frostfire Bolt and Arcane Wrath. | masked diff; hotfix-push.json `search_totals` | Aligned | - | - |
| A5 | **The Arcane raid order spent Presence of Mind on a Frostbolt at 54 to 56 while it already had Pyroblast** (guide round 20 Sims requested 1, integrator lead).<ul><li>The raid order takes Pyroblast at 54 and Heating Up at 57 (`addon-builds.json` mage-arcane raid, indices 44 and 47). Both Presence of Mind lines were gated on Heating Up, so for three levels the instant went on a 2.5 s Frostbolt instead of a 6 s Pyroblast.</li><li>Bare band 50 on the raid order truncated by level, 10,000 iterations, alliance, Frostbolt against Pyroblast, seed 1 / 50001: 54 253.0 / 253.3 to 254.9 / 255.2 (+0.75%); 56 274.5 / 275.0 to 276.3 / 276.6 (+0.6%); horde 54 248.3 to 250.1 (+0.7%). Error 0.2 to 0.3.</li><li>Fix: both lines are gated on `spellIsKnown(18809)` (Pyroblast). The guide card settles the gate by the talent's name, so the raid card keeps the Pyroblast row and the leveling card the Frostbolt row, as before. The leveling build never takes Pyroblast; raid 60 unchanged (horde 716.8 / 716.5 before and after). Notes build on guide round 20's wording. The round-18 gate test folds into `TestArcanePresenceOfMindWaitsForPyroblastNotHeatingUp` (red, then green).</li></ul> | `$R/mage/arc-pom-10k.txt`, `arc-pompyro.json` | CONFIRMED-BUG | old (round 18's gate, written before round 20 moved Pyroblast and Heating Up apart in the raid order) + guides-handoff (guide round 20) | S 846e07025, E e1a66349f |

## Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | F-S1 (filler), F-S2 and F-S3 (rank rows). | | CONFIRMED-BUG (F-S1) | old + guides-handoff | S 6c6a348b9 |
| F2 | Guide round 18 "Sims requested 2", Fire Blast on cooldown against held as a finisher at 20 and 28. A target-health gate does nothing on the bare fight (the condition never resolves, same casts), so the finisher was modelled as the last 20% of the fight: on cooldown 39.5 against 38.8 at 20, 46.1 against 45.3 at 28. On cooldown stays; the card is right. | `$R/mage/fire-fbfinisher.json` | Aligned | - | - |
| F3 | Talents and rows read against the client (Heating Up: "Non-periodic critical strikes with Fireball, Frostfire Bolt, Fire Blast, and Scorch ... Fire Blast grants an additional stack every time it deals damage"; Incineration 10% on Fire Blast, Ice Lance, Arcane Blast and Scorch; Critical Mass 6%; Burning Soul 70% / 30%). Raid 60 breakdown: Scorch 69.1 casts, Pyroblast 18.9, Fire Blast 27.4 (8 s less Improved Fire Blast's 2 s), Combustion 1.27, Heating Up 49 procs. The toll points (Wand Specialization, Improved Channeling, Arcane Resilience, Flame Throwing, Burning Soul) were judged in rounds 17 and 20 (F3): no point in reach does damage on a boss. | talents/mage.json; raid-60 breakdown | Aligned | - | - |

## Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Ice Lance at 60: still no live log.** Needed: a Frost mage at 56 to 60 with Frostbolt rank 10/11, Frost Nova and Ice Lance rank 6 (1240047) in one fight, fitted with `design/reviews/2026-11-04-sim-audit-round19/mage-tools/fit.py`. Last looked 03:43 UTC 5 Nov (see Logs). 0.107 from nine beta players at 24-35 stands. | logs below | OPEN (log needed) | - | - |
| R2 | Round 20's Frostbite change (only Blizzard's chill rolls at its rank's level-gap share) re-read: `chillApplied` passes 1, Blizzard alone passes `SpellLevelPenalty`; no rotation casts Blizzard. Fingers of Frost (15%, 1/2 charges, spent as a direct cast resolves, channels excluded) and Shatter's crit applied only while frozen. | `chill.go`, `blizzard.go`, `fingers_of_frost.go` (E 8fb31536d) | Aligned | - | - |
| R3 | Cold Snap: no GCD (12472 gcd_ms 0) and cast at the pull where it resets nothing. Offensive use was decided in rounds 18/19 (kept for Ice Block); not reopened. | spellconst 12472; round 19 S9 | Aligned (decided) | - | - |
| R4 | Raid 60 breakdown: Frostbolt 56.6 casts, Ice Lance 6.6 on Fingers of Frost (72.9% crit), Barrage Missiles 11.3, Ice Barrier 3.5 casts at 98% uptime, Berserking 1.48 (Troll), Major Spellblasting Potion twice. Matches the card. | breakdown | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Two mage-frost tests in web/src/lib/sim/rotations.test.ts failed on int/r21: a stale expectation, not a bug.** The test pinned the old Frostbolt filler note ("Frostbolt is the filler, and every cast can roll Fingers of Frost or Missile Barrage. Never stop casting to wait for a proc."), which guide round 18 (b806d876f, MAGE-R18-2) rewrote because it was false below 36. The note on the card is right. The test now pins the note by its opening words ("Frostbolt is the filler") and the wand note in full, so a guide round's rewording cannot break it again; 6 of 6 pass. | `npx vitest run src/lib/sim/rotations.test.ts`: 2 failed, then 6 passed | Not a sim bug (stale test) | guide round 18 b806d876f | S abd123a7a |
| X2 | int/r21's other engine changes (platform: Eureka! on Execute, spellIsKnown RequiredLevel; warlock APL syncs) touch no mage code path. The Arcane spellIsKnown(Presence of Mind) gates read a talent spell the build has from 31/32, so the RequiredLevel check changes nothing for mage. | `git log d5c708834..int/r21` in the fork | Aligned | - | - |
| X3 | BiS items with effects on the frozen raid gear (Robe of the Archmage, Frozen Heart of the Mountain's Chthonic Power, mana gems, Demonic Rune, potions) fire once or twice a fight as their cooldowns allow; no set bonus on any mage raid-60 row. | breakdowns | Aligned | - | - |

| X4 | **Evocation was castable below its trainer level 20.** It registered at every level and nothing checks RequiredLevel at cast time, so a level 10 to 19 mage channelled an Evocation the trainer has not taught (the ladder's level-10 rung casts one per fight in every mage golden; it showed in the A2 runs as one Evocation a fight at 12). The engine now registers it from 20; the three curated mage APLs declare 12051 inert with a note, so the level-10 rung does not flag the unresolved line. Bands 20 and up do not move. | `TestEvocationIsNotKnownBelowLevel20` failed, then passed; fork mage package passes with `-tags=with_db`; data/tests/test_apl.py passes | CONFIRMED-BUG | old (the ladder goldens have shown spell:12051 at 10 since the ladder began; no round read the rung) | E 7d806e74e, S 36f79cba3 |
| X5 | **spec-breakdown `-level` now reaches 10 to 19 on the band-20 entry** (integrator request; touches every class's tool). The entry's slots whose items need more than the level are left empty and the kit buffs follow the level; bands 20 and up run exactly as before. Test `TestLevelBelowTheLowestBandWearsOnlyWhatTheLevelAllows` failed (no such helper), then passed with the existing level and request tests. | `sim/cmd/spec-breakdown/main.go`, `below_band_test.go` | Tool change | - | S 559e39867 |

## Logs

Last looked 03:43 UTC 5 Nov: api.foreversixty.gg `/v1/reports/recent` lists only the 20 October report; foreverlogs.gg `/api/reports/public` is empty; Warcraft Logs v1 `/v1/zones` has no Forever zone on www (66), classic (57), fresh (21) or vanilla (29). Nothing a mage item can use.

## Before/after

Before: `$R/sb-before` from `$R/site-before/sim`. After: wt/r21-mage (6c6a348b9) with lane/r21-mage. Bare band 20, the published band build, 10,000 iterations.

| Spec | Faction | Level | Seed 1 before to after | % (s1 / s50001) |
|---|---|---|---|---|
| mage-fire bare | alliance | 22 | 40.1 to 40.4 | +0.7 / +0.7 |
| mage-fire bare | alliance | 25 | 43.6 to 43.9 | +0.7 / +0.7 |
| mage-fire bare | alliance | 28 | 45.6 to 46.1 | +1.1 / +1.1 |
| mage-fire bare | alliance | 29 | 46.2 to 46.4 | +0.4 / +0.4 |
| mage-fire bare | horde | 22 | 38.3 to 38.9 | +1.6 / +1.6 |
| mage-fire bare | horde | 28 | 44.6 to 45.2 | +1.3 / +1.3 |
| mage-fire bare 30 and up, raid 60 | both | 30, 60 | unchanged (raid 60 horde 721.0 / 721.4) | 0 / 0 |
| mage-arcane bare 50, raid order | alliance | 54 | 253.0 to 254.9 | +0.75 / +0.75 |
| mage-arcane bare 50, raid order | alliance | 56 | 274.5 to 276.3 | +0.66 / +0.58 |
| mage-arcane published bands, raid 60 | both | all | unchanged (the leveling build has no Pyroblast; raid 60 horde 716.8 / 716.5) | 0 / 0 |
| mage-frost | both | all | unchanged | 0 / 0 |

No spec outside mage moves: the changes are the mage-fire and mage-arcane APLs (and their fork copies); the test change is web-only.

## Files to regenerate

- Ladder goldens mage-arcane, mage-fire and mage-frost: the level-10 rung no longer casts Evocation (X4).

- addon-data.json and Data.lua: the mage-arcane Presence of Mind rows (three notes; the raid-order bands 54 and 56 now cast Pyroblast); the mage-fire rotation rows at bands 22 to 28 (Fireball now fires as the filler) and both changed notes at every band.
- sim/request ladder golden mage-fire, only if a rung reads the changed lines' text (its rungs are 10, 20, 30...; 20 has no Scorch and 30 has Heating Up, so the casts should not move).
- data/curated/coach/mage-fire.md:29 ("Fireball only before you train Scorch at level twenty-two"), if the coach starter is regenerated from the APL.
- Bare band-20 mage-fire BiS and weights (the ranker reads the rotation at band 20): BiS refresh.

## Measure at refresh

- mage-fire band 20 BiS and weights on the Fireball filler.
- Carried: the Arcane BiS re-rank on the round-20 raid string; the Frost re-rank and weights after Ice Lance 0.107.

## For other lanes

- Every lane: spec-breakdown `-level 10..19 -band 20` now runs levels below the lowest band on the gear that level can wear (X5).
- Every lane, method: an unranked class spell registered without a level check (Evocation, X4) is castable below its trainer level, because only spellIsKnown reads RequiredLevel; check your class's unranked spells against trainables on the ladder's level-10 rung.

- Platform (component, SHOULD from the wow-player review): a below-30 row on the raid-60 card prints the raid rank ("Fireball RANK 11" beside "Below 30").
- Every lane, method: an APL condition on `currentHealthPercent` of the target does nothing on the bare and raid fights (the value does not resolve and the gate drops out); model execute or finisher phases with `remainingTimePercent`.

## For the guide lane

- Done here (Arcane): the three Presence of Mind notes and arcane.md:56 ("from 54"), on guide round 20's wording ("and with Arcane Power once you have it", "before that and with the leveling build")
- Done here (Fire): both Fire card notes (Scorch filler, Fireball), fire.md:66 (Accelerant), fire.md:116 (leveling). The Scorch note builds on land/guides-r19's MAGE-R19-1 wording; where the two branches meet on that line, this lane's text supersedes it (the row is now gated on Heating Up, so "from 30" is the row's own condition). guides-r19's Fire Blast and Evocation note changes are untouched here.
- MAGE-R19-9 (fire.md:116) is decided: Fireball through the 20s.

## Guide review

wow-player reviewed the rendered /guides/mage/fire page at 1440 (astro dev on wt/r21-mage, real data; the card's two changed rows previewed by patching the generated addon-data.json locally, then restored). Captures: `$R/mage/shots/fire-card.png`, `fire-leveling.png`, `fire-leveling-tree.png`, `fire-talents.png`, `fire-1440.png`.
- First pass: NO-SHIP on four MUSTs: "rank for rank" overstated the measurement (it is the ranks a mage has in the 20s); "from 27" read as Scorch's training level; the Fireball row on the raid card read as the 60 mage's filler; the first captures had broken icons (fixture data). All fixed: the leveling paragraph says "at the ranks you have in the 20s" and "once you take Improved Scorch at 27"; the Fireball row reads "Below 30, Fireball is the filler"; the Scorch row says "more than two Scorches" (Fireball is 3.5 s at 60); recaptured on real data.
- Second pass: **SHIP**, no MUSTs. SHOULDs: the Fireball row prints "RANK 11" beside "Below 30" (the card labels every row with the raid rank; a component point for platform); the leveling tree hydrates (checked, `fire-leveling-tree.png`).
- Arcane (Presence of Mind on Pyroblast from 54), `arcane-card.png`: **SHIP**, no MUSTs. Taken SHOULDs: the Pyroblast note reads "from 54 on the raid build: a 6 second cast without Heating Up stacks", and arcane.md:56 says "from 54; before that and with the leveling build". The reviewer asked to confirm Arcane Blast is 1.5 s at every stack: it is (400574-1239700, 1.5 s cast; the stack raises cost and spell damage, not cast time; round 20 A5), so "never" stands.

## Still open

- Ice Lance's coefficient at 60 (a level 56+ Frost log); Arcane Power's cost on a mana log (the client column already settles the model, round 20 S2).
- Carried: the level-60 book ranks (pending owner).
