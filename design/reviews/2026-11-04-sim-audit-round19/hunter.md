# Round 19, lane hunter: Beast Mastery, Marksmanship, Survival (melee)

Client 1.60.1.70755 with the late re-layer. Site wt/r19-hunter from int/r19; fork lane/r19-hunter from int/r19 6ea675594.
Method: spec-breakdown through `$R/sim.sh`, guide builds (`$R/rt`), `-bis-dir $R/bis-frozen`. 3000 iterations to explore; 10000 at seeds 1 and 50001, both factions, for every reported number. Probe binaries that are not committed are named where used: `sb-tl` reads a target level from an env var, `sb-pet` sets the pet family and speed, `probe/sb-*` use a temporary rank-choices copy, and `probe/sb-down` gives a melee special the next rank down's numbers.

**Result: 3 CONFIRMED, all fixed, all `old`.** (1) Only the top Serpent Sting rank was registered, which makes round 17's "Serpent Sting downranks lose" an invalid measurement (E d9794e20c). Re-measured, the downranks still lose. (2) Only the top Raptor Strike rank was registered (E 62d36b277). (3) Survival at 46-47 now queues Raptor Strike rank 5 (S 4f5b4edee), with guide text and a wow-player SHIP. Seed: Lone Wolf against the pet is decided and written (S 3cf7480e0, wow-player SHIP). In a raid, dismiss the pet. Out in the world at 60, keep it.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Lone Wolf against the pet at 60 (survival.md:140, guide round 15 sim request 2).** The pet build moves the Lone Wolf point to Rapid Killing, which does nothing in these fights (`5-0050051-500230231051222151`), so the cat comes out. 10000 iterations, frozen BiS, seeds 1 and 50001, both factions. **Raid** (level-63 boss, 180 s): Lone Wolf leads, Horde 860.5/861.7 against 838.5/838.1 and Alliance 860.2/858.8 against 839.2/838.3. That is -2.4 to -2.7% with the pet. **Out in the world** (no raid buffs, a level-60 mob): the pet leads at every fight length and every run. At 20 s it is +0.5% Horde and +0.9 to +1.0% Alliance. At 40 s it is +1.0 to +1.6%, at 60 s +1.5 to +1.9%, and at 180 s +2.0 to +2.3%. The level-63 bare fight gives +1.8 to +1.9%. The pet also takes the hits. The reason: the raid's buffs reach the hunter, so Lone Wolf's 20% multiplies a bigger number, while the pet takes no player buffs (round 16 H1). Lone Wolf is the build's last point, [2,3,4] at 60 in addon-builds, so the question exists only at 60. **Web:** no Forever source addresses the world case. Warcraft Tavern (https://www.warcrafttavern.com/forever/guides/hunter/, read 2026-11-04) ties Lone Wolf to Marksmanship. Icy Veins and Mobalytics returned 403; their search snippets say a pet cuts downtime while questing. The sim is clear, so I wrote the guide: survival.md:47, 63, 71, 86 and 140. In a raid, dismiss the pet. Out in the world at 60, keep it, because the talent does nothing while it is out. | `$R/hunter/lw.tsv`, `$R/hunter/lw.sh` (probe `sb-tl`, env target level); data/builds/1.60.1.70755/addon-builds.json survival order | Decided (guide text; no build, APL or rank change) | guides-handoff | S 3cf7480e0 |
| H2 | **Cat against raptor (guides owner question 4).** Forever raptors have Savage Rend (1265065-1265069, 26 a tick at rank 5, 6 ticks, 60 s cooldown), and the engine models it (`pet_family_abilities.go:55`; bleed crits per owner law). BM raid 60, 10000 iterations, both pets at the same 1.2 s speed: raptor +0.15 to +0.31% (Horde 779.6/779.6 → 781.5/780.9; Alliance 782.8/781.9 → 784.0/784.4). At 2.0 s: cat 776.9/777.4, raptor 778.9/778.6. That is a tie under any build bar, and the request carries no pet field, so the cat stays. Either pet is fine. A fast pet beats a 2.0 s one by about 0.3%. | `$R/hunter/catraptor.tsv` (probe `sb-pet`) | Aligned (tie; owner question closed: either) | - | - |
| H3 | **Ranged Weapon Specialization on Serpent Sting: still OPEN.** The tick talent applies only at 34+ (tier 5). The beta logs (2705 Huntscar and Salami, 2712 and 2745 Miss) have Serpent Sting ticks and Arcane Shot hits for the same hunter. The ratios are sting rank 4 38-40 against base 26, and Arcane Shot rank 4 118-125 against 85. They don't decompose. The talents are unknown, partial resists spread the ticks (35 to 40), and Arcane Shot carries an unexplained extra factor of about 1.1 beyond Lone Wolf and RWS. That last factor touches the owner's pending Arcane Shot tooltip question, which I did not reopen. Worth at most 0.3% (round 14 H4). | `$R/hunter/tools/stingratio.py`, `$R/hunter/dmg/` (10 damage timelines fetched) | OPEN (a level-60 log of a hunter with known talents, sting ticks with and without RWS) | - | - |
| H4 | **Lacerating Strikes on refresh; the queued Raptor Strike and the off-hand miss penalty: still OPEN.** Lacerating Strikes is tier 6 (level 39+). Beta logs stop at 35, and foreverlogs does not separate main-hand from off-hand white swings. The raid breakdown shows the shape the model assumes: main hand 19.0% miss, off hand 17.3%, because a queued Raptor Strike lifts the off-hand penalty while it waits. | talents/hunter.json tiers; `$R/hunter/bd-survival.txt` | OPEN (level-40+ Survival log) | - | - |
| H5 | **40-59 rank rows, level by level (shared method lead plus the integrator's any-gain rule).** BM and MM run dry for 70 to 100 s at bare 40-59, and Survival for 6 to 15 s. Swept at floor, +3, +6 and +9, 3000 iterations to explore and 10000 for every candidate. **Serpent Sting** N-1 and N-2 lose at every level from 26 to 59 on both specs (-0.5 to -2.3%), measured after H6's fix. **Arcane Shot** N-1 and N-2 lose or tie (BM rank 7 at 58-59 is +0.04%, inside error). **Sniper Shot** is not in the leveling build. **Survival:** Mongoose Bite and Lacerate down a rank lose at every level (-0.06 to -1.9%), and Immolation Trap down a rank loses 3 to 4%. **Raptor Strike rank 5 at 46-47 wins**: see S1. MM Aimed Shot at 50-51 stays without a row (0.0 to 0.1 DPS against 0.1 error, a true tie, not a sub-1% keep). No hunter row had been kept for being under 1%. | `$R/hunter/probe/explore2.tsv`, `explore3.tsv`, `svdown.tsv`, `svrs.tsv` | Aligned except S1 | - | - |
| H6 | **Serpent Sting registered only its top rank.** `registerSerpentStingSpell` stopped at the highest learned rank, so a lower rank named by a rotation or a rank-choice row found no spell. The site rewriter rewrites both the cast and its `dotRemainingTime`, and the engine nils a missing spell, so the line silently dropped and Serpent Sting was never cast. Round 17 B1 measured "Serpent Sting rank 3 loses at every band" through exactly this path, so that conclusion was void; H5 re-measures it properly. Arcane Shot, Aimed Shot and the traps already register every rank. Test first: `TestSerpentStingRegistersEveryLearnedRank` (red: ranks 1, 5 and 6 missing at 53; green after). TestP1Hunter is unchanged. | `sim/hunter/serpent_sting.go:83`; client trainables rank levels 4/10/18/26/34/42/50/58 | CONFIRMED-BUG (model gap; no published number moves, no row names a lower rank) | old (round 17 B1 measured through it and missed it) | E d9794e20c |
| H7 | **Low-rank proc watch (priest P1).** No hunter rank row casts a share-0 rank of a proc spell. Improved Wing Clip is in no hunter build (round 16 S2), and Wing Clip has no rank row. Aimed Shot rank 1, Serpent Sting and Raptor Strike carry no on-hit proc of their own. | rank-choices.json hunter rows | Flag: none | - | - |

## Beast Mastery

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Raid 60 breakdown (guide raid build, frozen BiS, Horde, 3000 it, 780.0). Auto Shot 44%, cat 23% (194 swings at 1.2 s, 40% glancing against a 63, dodge 6.4% from behind, crit 33.8% = the hunter's crit plus Ferocity). Multi-Shot 20.6, Arcane Shot 27.0, Serpent Sting 12.7, Aimed Shot 7.6, Bestial Wrath 2, Rapid Fire 1. 0 s out of mana. Same shape as rounds 16 to 18. | `$R/hunter/bd-beast-mastery.txt` | Aligned | - | - |
| B2 | **Coefficient check (the mage lane's Ice Lance lesson).** No hunter spell uses a generator default. The engine reads no `*SpellCoeff` from constants_auto_gen.go. Arcane Shot, Serpent Sting and Volley are set to 0, matching the client's 0 on every trained rank (spellconst 3044/14281-14287, 1978/13549-13555, 1279715-1279721). Shots and melee specials use weapon damage with physical BonusCoefficient 1. Every BonusCoefficientFromAP column is 0. The one scaling not in a column is Summon Hawk's dive, 5% of RAP: the tooltip reads `${108+($rap*(5/100))}` (spelltips 1293527), and dummy effect 3 carries 5. No guide build casts it. | spellconst/hunter.json (all `ap_coefficient` 0); spelltips/hunter.json; `grep SpellCoeff sim/hunter` | Aligned | - | - |
| B3 | Pet families: raptor's Savage Rend tick (26) lands at 65 on average at raid 60. That is the pet's 1.25 base, family 1.10, Unleashed Fury 1.15, Bestial Wrath uptime and crit at the bleed rule, so it is consistent. | `pet.go:91-94`; breakdown | Aligned | - | - |

## Marksmanship

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | Raid 60 breakdown (801.1, Horde, 3000 it): Auto Shot 53%, Multi-Shot 21.8, Sniper Shot 9.0 at 1846, Serpent Sting 11.6, Aimed Shot 5.8, Arcane Shot 17.6. 0% miss (Surefooted 3). 4.3 s out of mana. Unchanged since round 16. | `$R/hunter/bd-marksmanship.txt` | Aligned | - | - |
| M2 | Round 18's Aimed Shot rank 1 row (28-49, 52-59, bare 60) holds after H6. The Serpent Sting and Arcane Shot downranks around it lose (H5). | H5 | Aligned | - | - |

## Survival (melee)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Raptor Strike rank 5 at 46-47.** Immolation Trap rank 4 arrives at 46 and costs more, and out of mana goes from 7 s at 45 to 14 s at 46 on the bare fight. Raptor Strike rank 5 (+35, 55 mana; 22 after Resourcefulness) instead of rank 6 (+40, 70 mana; 28 after) buys back traps. Bare band 40, 10000 iterations, before (`$R/sb-before`) → after (lane head): **46** Horde 186.0/186.1 → 186.7/186.7, Alliance 188.0/188.0 → 188.2/188.2; **47** Horde 187.8/187.8 → 188.4/188.4, Alliance 189.6/189.6 → 189.7/189.8. Error is 0.1. The gain is beyond error at both seeds for Horde, and for Alliance at 46; Alliance at 47 is +0.1/+0.2, never a loss, so the row has no race key. 45 (176.6 / 177.6) and 48 (191.8 / 193.9) don't change. At 40-45 rank 6 wins or ties, and from 48 rank 7 wins. The engine needed every rank registered first (S2). New row: hunter-survival Raptor Strike 14266, 46-47, rank 5. Test first: `TestRound19SurvivalRaptorStrikeRankFiveAt46` (red at 46/47, green after). Guide: survival.md Levels 20 to 60. The rotation paragraph now has Immolation Trap. A new paragraph says to keep rank 5 for 46-47, drag it back from the spellbook, and put rank 7 on the bar at 48. wow-player: SHIP WITH FIXES, all four SHOULDs applied, then **SHIP**. | `$R/hunter/probe/svrs.tsv`, `$R/hunter/rs-after.tsv`; spec-breakdown breakdowns at 45/46 (Immolation Trap mana 504 → 673) | CONFIRMED-BUG (rank row beyond error) | old (no round swept Survival's melee ranks; the engine could not cast them) | S 4f5b4edee (needs E 62d36b277) |
| S2 | **Raptor Strike registered only its top rank.** Each learned rank now registers its own on-next-swing queue (tag 3) and hit (tag 1), and all ranks share one 6 s cooldown timer. The hit is bound per rank: the old code read `hunter.RaptorStrikeHit`, which the last registered rank overwrote. `hunter.RaptorStrike` stays the top rank. Test first: `TestRaptorStrikeRegistersEveryLearnedRankOnOneCooldown` (red: ranks 1-5 missing at 46). The whole `sim/hunter` package passes with `-tags=with_db`, including TestP1Hunter (no golden moves). SV raid 60 is unchanged at 860.5, and bare 46 with no row is unchanged at 186.0. | `sim/hunter/raptor_strike.go` | CONFIRMED-BUG (model gap) | old | E 62d36b277 |
| S3 | **Other single-rank registrations (integrator ask).** These register only the top rank: Mongoose Bite, Lacerate, Sniper Shot, Summon Hawk, Wing Clip, Counterattack, Hunter's Mark, the aspects, and Volley (whose loop breaks at the top rank). Multi-Shot is one rank in Forever. Checked by giving the cast the next rank down's numbers (`probe/sb-down`, bare 40-59, every 3 levels): Mongoose Bite loses (-0.06 to -0.41%) and Lacerate loses (-0.96 to -1.86%). Sniper Shot is not in the leveling build, and at 60 the rotation is authored. No build casts Summon Hawk or Volley. Wing Clip is cast only under Windfury (raid 60). Counterattack is rare from behind, and the aspects and Hunter's Mark have no rank trade-off. None needs a row, so I left them single-rank. A future row for any of them needs the same engine change first, or the line silently drops (H6). | `$R/hunter/probe/svdown.tsv`; register functions in `sim/hunter/*.go` | Aligned (noted as a trap) | - | - |
| S4 | Raid 60 breakdown (861.0, Horde, 3000 it): Raptor Strike 24 hits, Mongoose Bite 20.3, Strider Kick 19.1, Immolation Trap 9.7, Lacerate 9.5 plus Lacerating Strikes 53 ticks, Wing Clip 61.8 under Windfury. 0 s out of mana. Flat amounts against 70755 tooltips: Raptor Strike +70, Mongoose Bite +57, Counterattack 50% + 110, Wing Clip 50, Lacerate 406 over 21 s (58 × 7), Strider Kick 100% normalized and 5.81% base mana. All match the engine. | spelltips/hunter.json 14266, 14271, 20910, 14268, 1299332, 1317257 | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | `go vet ./sim/hunter/` still reports the test-only lock copy at `talents_test.go:496` (round 18 X1). `go vet ./...` on the api module is clean. | go vet | Aligned (note) | - | - |
| X2 | **Live logs.** Last looked at 2026-11-04 19:13 UTC, before the 23:00 launch. foreverlogs `reports/2815` returns "Report not found", and `/v1/reports/recent` lists only the 2026-10-20 report. No launch-day or level-60 log exists, so H3, H4, the hawk swing and Heart of the Lion on the pet stay open. Warcraft Logs is still blocked (gcloud). | foreverlogs API; api.foreversixty.gg | OPEN (no data) | - | - |

## Before/after

Raid 60, guide builds, frozen BiS, 10000 iterations, against `$R/before.tsv`: **no change** for any hunter spec. SV raid 60 Horde seed 1 is 860.5 → 860.5. Neither engine commit moves a number unless a row asks for a lower rank, and the only new row is bare 46-47.

| spec | preset, level | faction | seed 1 | seed 50001 |
|---|---|---|---|---|
| Survival | bare 46 | Horde | 186.0 → 186.7 (+0.38%) | 186.1 → 186.7 (+0.32%) |
| Survival | bare 46 | Alliance | 188.0 → 188.2 (+0.11%) | 188.0 → 188.2 (+0.11%) |
| Survival | bare 47 | Horde | 187.8 → 188.4 (+0.32%) | 187.8 → 188.4 (+0.32%) |
| Survival | bare 47 | Alliance | 189.6 → 189.7 (+0.05%) | 189.6 → 189.8 (+0.11%) |

No commit moves DPS for a spec outside the hunter class: both engine commits are in `sim/hunter` only, and the site row is keyed to hunter-survival.

## Files to regenerate

- addon-data / Data.lua and the coach starter for hunter-survival (the new 46-47 Raptor Strike row).
- Ladder goldens (`sim/request TestRotationLadder`) only if a hunter-survival checkpoint covers 46 or 47. The checkpoints are 10, 20, 25, 30, 35, 38, 40, 45, 50, 55 and 60, so probably none.
- Fork: none (TestP1Hunter passes unchanged).

## Measure at refresh

- Round 15's guide list, carried: MM race order, Survival Crusader against Agility on both weapons, the BM and SV race re-runs and BM/MM with the best non-BoP two-hander, all on the re-ranked gear.
- Re-run the 46-47 Raptor Strike row on the re-ranked band-40 gear.

## For other lanes

- **The single-rank trap (H6, S3).** A rank-choice row naming a rank the engine never registered drops the whole line without an error. Any lane adding a downrank row should check that its class's register function loops over every rank. The integrator has already relayed this.

## For the guide lane

- survival.md was rewritten in this lane (H1, S1). Lines 47, 63, 71, 86, 140 and Levels 20 to 60 now carry: the pet in the world and Lone Wolf in a raid; Immolation Trap in the leveling rotation; Raptor Strike rank 5 at 46-47. Take this lane's version if a guide branch touches the same lines.
- Cat or raptor (H2): either serves, so the guide may say so. A fast pet is the small edge.

## Guide text

Done in this lane: survival.md (H1, S1). `_prose.test.ts` passes. wow-player reviewed the rendered pages (astro dev on this branch, 1440, `$R/hunter/shots/sv*-*.png`). Lone Wolf took three passes: the overview contradiction was fixed, then the "it" wording, then **SHIP**. Raptor Strike rank took two passes: four SHOULDs applied, then **SHIP**. wow-player asked for one check on a built page with data: the blank bars above "Levels 20 to 60" in the captures are the leveling-ladder island's loading skeleton on a dev server with no API. They were there before these edits, and these edits don't touch that component.

## Guide text not done

None.

## Still open

- H3 (RWS on Serpent Sting) and H4 (Lacerating Strikes refresh; queued Raptor Strike and off-hand miss): need level-40+ and level-60 logs.
- From round 18, not reopened: the hawk's level-60 swing (owner check), Heart of the Lion on the pet. Pending owner: level-60 book ranks, Arcane Shot's tooltip with spell power, Heart of the Lion 280/276.
