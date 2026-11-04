# Round 18, lane hunter: Beast Mastery, Marksmanship, Survival (melee)

Client 1.60.1.70755 with the late re-layer. Site wt/r18-hunter from int/r18 4600b7d82; fork lane/r18-hunter from int/r18 b2c4a2d05.
Method: spec-breakdown through `$R/sim.sh`, guide builds (`$R/rt`), `-bis-dir $R/bis-frozen`. 3000 iterations to explore; 10000 at seeds 1 and 50001, both factions, for every reported number.
Log tool: `hunter-tools/hawktiming.py` next to this file (round 16's `hunter-tools/` fetch and analyses are reused).

**Result: 1 CONFIRMED model gap, closed in the engine (E 8c24d6202). It moves no published number.** Summon Hawk's 18-second hawks are now in the sim. With them, the Summon Hawk build gains +0.6 to +0.7% at raid 60 on the floor projection. That is short of the 1% the adoption rule needs, so no build, rotation or guide text changes. Heart of the Lion on the pet: still no log. Fresh re-audit: aligned.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **No level-60 hawk data exists yet.** foreverlogs.gg has nothing after report 2814 ("Test", 2026-11-04 12:24 UTC, 0 encounters); 2815 and up return "Report not found". api.foreversixty.gg `/v1/reports/recent` still lists only the 2026-10-20 report. The web has no measured level-60 hawk swing (search 2026-11-04). Guides give only the dive's text, or say the beta stopped at 30. Forever launches at 23:00 UTC, so a live log cannot exist before this round ends. | foreverlogs `reports/2813..2900` (2026-11-04 16:10 UTC); `/v1/reports/recent`; https://mobalytics.gg/wow-forever/classes/beast-mastery-hunter-guide, https://www.icy-veins.com/wow-forever/beast-mastery-hunter-ranged-dps-pve-guide, https://wow.gg/guides/hunter-beast-mastery-forever-overview (read 2026-11-04) | OPEN (no data) | - | - |
| H2 | **Summon Hawk's 18 s assault was missing from the engine. It is now modelled from the beta logs.** The talent text says the hawk "dive-bomb[s] your targeted enemy ... and continu[es] its assault for 18 sec. Only 2 hawks can be active at once." The engine had only the dive. The client has no data for the rest: all four ranks' effect 1 trigger the same 1312639, a dummy. The 48 beta encounters (8 BM hunters, levels 25 to 35, 718 hawk swings) give the whole kit: (a) a "Hawk" pet unit arrives a median 1.07 s after the cast (quartiles 1.02 to 1.17 s); (b) it swings in melee every 2.26 s, or 1.62 s inside Rapid Fire; (c) on a target that lives, it lands 8 swings a summon (the mode, 27 of 134 summons; 9 to 11 inside Rapid Fire); (d) 0 glancing blows against level 25-35 mobs; it misses 3.6%, is dodged 4.4% and parried 1.2%; (e) it crits 0 to 1% without Ferocity and 10.6 to 11.6% with it, so it takes Ferocity but not the hunter's crit; (f) it takes no buffs (no `buff_applied` ever lands on a Hawk). **New in the engine:** `summon_hawk_assault.go`. Each cast lands a guardian 1.07 s later. The guardian swings every 2.26 s for 18 s and is hasted 40% by Rapid Fire. It has Ferocity's crit and Unleashed Fury's damage, inherits no stats or attack power, and takes no buffs. When a third hawk arrives, the oldest leaves. A hunter aura, Hawk Assault (1312639), counts the hawks out, so a rotation can read `auraNumStacks`. **Level projection:** a swing is 0.32 × level (19.2 at 60, before Unleashed Fury and armor), ±15% (the logs' non-crit spread). This is the floor of the readings. The level-35 talented mean (12.8 on level-35 mobs) is taken as unarmored and divided by Unleashed Fury 5: 12.8 / 1.15 / 35 = 0.32. It grows in proportion to level, as a hunter pet's weapon does (`hunterPetDamageLevelScale`). The pooled through-origin fit of all six levels gives 0.35 a level. Correcting for the beta mobs' armor, or letting the hawk scale with attack power like the dive (the hawk-to-dive ratio holds at 0.21 to 0.29 from 25 to 35), gives about 0.45 to 0.50. Tests first: `TestSummonHawkSendsAHawkThatSwingsForEighteenSeconds` (arrival, 18 s, 8 swings), `TestOnlyTwoHawksAtOnce`, `TestRapidFireHastesTheHawks` (2.26 to 1.614 s), `TestHawkTakesFerocityAndUnleashedFuryNotTheHuntersCrit`, `TestNoHawksWithoutTheTalent`. They were red: no hawks existed. Now green, and the whole `sim/hunter` package passes with `-tags=with_db`. | foreverlogs 2695 to 2764 timelines (`$SP/r16/hunter/full`), `hunter-tools/hawktiming.py`, round 16 `analyze.py`/`ratio.py`; spellconst 1293241/1293525/1293526/1293527 effect 1 → 1312639 (no body); talents/hunter.json Ferocity, Unleashed Fury ("your pets and hawks") | CONFIRMED (model gap). The engine left out half the talent. No guide build takes Summon Hawk, so no published number moves. | old (round 15 B2/W6-2 and round 16 H2 found it and could not price it) | E 8c24d6202 |
| H3 | **The build decision with the hawks modelled: unchanged.** Candidate: one Endurance Training point (4 → 3, the point round 15 and the guides' S7 named) into Summon Hawk. Endurance Training keeps 3 points, so the unmodelled talent stays. Summon Hawk takes Arcane Shot's slot, with the same mana and Multi-Shot gates. It is cast while fewer than two hawks are out; casting it on every cooldown and casting it only below two hawks give the same number. Raid 60, frozen BiS, 10000 iterations. On the floor (0.32): Alliance 782.8 → 787.3 (+0.57%) and 781.9 → 787.3 (+0.69%); Horde 779.6 → 785.0 (+0.69%) and 779.6 → 784.9 (+0.68%). That is short of 1% at both seeds, so **not adopted**. Sensitivity, seed 1 Horde, 3000 iterations: 0.35 a level gives +0.78%, and 0.48 gives +1.37% (10000 iterations: +1.30 to +1.42% on all four runs). The break-even is about **0.40 a level, 24 a swing at 60 before Unleashed Fury and armor**. At bare 60 the hawk build loses (-0.4%; it is out of mana for 78 s of 180), so the leveling build stays without it. The hawk melee itself is 11 DPS at raid 60 on the floor; the dive is 27.5 DPS. | spec-breakdown (`$R/hunter/hawk-raid.txt`, `$R/hunter/rot/bm-hawk-*.json`, sensitivity binaries `sb-k0.35`, `sb-k0.48`, not committed) | Aligned under the adoption rule. The uncertainty is the projection: the floor and central readings fall short of 1%, and the armor-corrected or attack-power reading clears it. | - | - |
| H4 | **Owner's 2-minute check, now with a number.** At 60, with a Beast Mastery hunter with Unleashed Fury 5, cast Summon Hawk on any target. In the combat log, read ten of the hawk's non-crit, non-glancing swings and the dive's non-crit hit on the same target; both carry Unleashed Fury and the target's armor. **Adopt Summon Hawk (Endurance Training 4 → 3) if hawk swing / dive ≥ 24 / (108 + 0.05 × your ranged attack power)**: 0.14 at 1300 RAP, 0.12 at 1800. The floor model assumes 19.2 / (108 + 0.05 × RAP). Either way, set `hawkSwingDamagePerLevel` (summon_hawk_assault.go) to the measured ratio × (108 + 0.05 × RAP) ÷ 60. The beta ratio at 25 to 35 was 0.21 to 0.29 against the rank-1 dive, so the dive comparison is the cleanest test. | H2, H3 | OPEN (owner check) | - | - |
| H5 | **Heart of the Lion on the pet (round 17 H2): still no direct evidence.** No new logs (H1). A web search finds no Forever source either way; Season of Discovery sources describe the aura only as a party/raid "kings". The round-17 basis stands: client 409583 is unflagged and targets every nearby ally, and the unflagged Trueshot Aura lands on pets in the beta logs. | foreverlogs (H1); web search 2026-11-04 (https://warcraft.wiki.gg/wiki/Heart_of_the_Lion, https://virtgold.com/blog/hunter-dps-guide-for-season-of-discovery-phase-2/) | OPEN, owner check (hover the pet's buff bar with Heart of the Lion on) | - | - |

## Beast Mastery

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Raid 60 guide build is unchanged by the hawk commit: 782.8/781.9 Alliance, 779.6/779.6 Horde, identical to `$R/before.tsv`. The hawks register only for a hunter with the talent. | spec-breakdown 10000 it | Aligned | - | - |
| B2 | Pet talents against 70755 text, re-read: Ferocity 2% a rank on pet and hawk (`talents.go:25`), Unleashed Fury 3% a rank (pet Tamed Pet Passive; hawk DamageDealtMultiplier), Rapid Fire 40% ranged and melee (3045 auras 140/319), now also on hawks. Summon Hawk ranks 25/36/48/60, 32/47/85/108 + 5% RAP, 80/105/135/190 mana, 6 s shared with Arcane Shot, all match spellconst. | talents/hunter.json; spellconst/hunter.json; constants_auto_gen.go:4653-4663 | Aligned | - | - |
| B3 | Ladders and band talents: no hunter ladder (ladder-first/ladder-order) or guide build takes Summon Hawk, so no band number moves. The site adapter handles several pets (`sim/adapter/adapter.go:168, 286, 653`), so a build with hawks would report them. | grep data/curated, sim/adapter | Aligned | - | - |

## Marksmanship

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | Raid build talents re-read against the 70755 text: Lethal Attacks 1%/rank, Careful Aim 20% of Intellect/rank, Lone Wolf 20%, Surefooted 1% hit/rank, Improved Tracking 1%/rank, Survivalist 5 (health) as the gate that opens Surefooted. Barrage, Sniper Shot, Trueshot Aura, Mortal Shots and Ranged Weapon Specialization were checked in rounds 15 to 17 (M2, B3) and are untouched since. No pet in the raid (Lone Wolf), so the hawk change cannot reach it. | talents/hunter.json; talents.go:132, 485-540, 662-690 | Aligned | - | - |

## Survival (melee)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Every Survival talent in the build, re-read against the 70755 text: Savage Strikes 2%/rank (4% at 2), Predator's Edge crit damage 6%/rank and off hand 10%/rank (30%/50% at 5), Clever Traps 15%/rank on Immolation and Explosive Trap (30% at 2), Survivalist's Discipline 40% trap cooldown (9.7 traps in 180 s on an 18 s cooldown), Resourcefulness 30%/rank cost and 50%/rank regen, Lightning Reflexes 3%/rank, Surefooted 1%/rank, Improved Tracking 1%/rank. **Survival Tactics** (no earlier round named it) is trap and Feign Death hit only, and the build takes no points in it. | talents/hunter.json; talents.go:39-46, 66, 132, 164-191, 226, 325-380, 662-740 | Aligned | - | - |
| S2 | Raid 60 rotation behaviour, both factions: Mongoose Bite 20.3 casts (the Expose Prey window opens 5302, so the APL's `auraIsActive 5302` covers both windows), Strider Kick 19.1 casts on its 8 s cooldown, Lacerate kept up (9.5 casts, 58 ticks), Immolation Trap 9.7, Wing Clip only under Windfury (61.8). Immolation Trap rolls the boss's spell hit without the hunter's hit (about 21% of ticks lost, matching 17% plus the level gap), as traps do. 0 s out of mana. | spec-breakdown 3000 it | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | `go vet ./sim/hunter/` reports a lock copy at `talents_test.go:496` (`target := *core.DefaultTargetProtoLvl60`). It predates this lane (line 489 at b2c4a2d05) and is test-only. | go vet | Aligned (note) | - | - |
| X2 | Guide-accuracy rounds 12 to 14 carried four hunter LIKELYs to the sim. Falcon's proc (closed in round 15 W2), Sniper Shot haste (closed in round 17 B3) and the stale top-level APL notes (round 17 X2, not rendered) are settled. Summon Hawk is now modelled (H2). | guide-accuracy/round14/hunter.md:85-90 | Aligned | guides-handoff | - |

## Before/after

Raid 60, guide builds, frozen BiS, 10000 iterations (before = `$R/before.tsv`):

| spec | faction | seed 1 | seed 50001 |
|---|---|---|---|
| Beast Mastery | alliance | 782.8 to 782.8, 0.00% | 781.9 to 781.9, 0.00% |
| Beast Mastery | horde | 779.6 to 779.6, 0.00% | 779.6 to 779.6, 0.00% |
| Marksmanship | both | unchanged (no Summon Hawk, no pet) | unchanged |
| Survival | both | unchanged (no Summon Hawk, no pet) | unchanged |

Report-only (H3): the Summon Hawk build is +0.57 to +0.69% on the floor projection and +1.30 to +1.42% on the 0.48 reading.

No commit moves a spec outside the hunter class: the change is in `sim/hunter` only, and core is untouched.

## Files to regenerate

- None. No build, APL, rank row or published number changed. The fork golden `sim/hunter/TestP1Hunter.results` passes as is (`-tags=with_db`).

## Measure at refresh

- If talent-search or a ranker now looks at Summon Hawk, it sees the floor model. Any pick that rests on the hawk should wait for the owner check (H4).

## For other lanes

- None.

## For the guide lane

- None from this lane. If the owner's check clears the bar (H4), the BM raid build gains Summon Hawk for the fourth Endurance Training point. beast-mastery.md:81 (talent paragraph) and the rotation section would then need a line on keeping two hawks out in Arcane Shot's place.

## Guide text not done

None (no guide-affecting change).

## Still open

- H4: the hawk's level-60 swing (owner check, with the bar above). H5: Heart of the Lion on the pet's buff bar.
- Carried, not reopened: cat vs raptor, Ranged Weapon Specialization on Serpent Sting, Lacerating Strikes on refresh, the queued Raptor Strike and the off-hand miss penalty. Pending owner: level-60 book ranks, Arcane Shot's tooltip with spell power, Heart of the Lion 280/276.
- Warcraft Logs is still blocked (gcloud login).
