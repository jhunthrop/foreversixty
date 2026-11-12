# Sim accuracy 2026-11-12: warlocks against the post-hotfix logs (Affliction, Demonology, Destruction)

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main a38608c67, fork from master 15e2c6bae (the pin). Client on the site: 1.60.1.70790 (the 2026-11-11 lane's hand rows carry the 10 Nov push). This lane re-checks that lane's hotfix model against every warlock fight logged on or after 10 Nov, and takes round 29's open Master Summoner item.

**Totals: 3 confirmed bugs, all old, all fixed (the pre-pull summon's mana, the Imp's cast cadence, the Imp's Firebolt crit). 0 new data: Blizzard's 11 and 12 Nov posts carry no warlock line, and our API holds no report newer than 8 Nov. The 10 Nov Firebolt rows (W5) are now confirmed by live Imps at ranks 2, 3 and 4. No build or rotation moves (Master Summoner holds under the adoption rule). No published raid number moves; Demonology's ladder rungs at 38 to 60 move by -0.4 to +0.3%. Engine pin 15e2c6bae to a63d30e26.**

## Method

1. **Real fights first.**
   - Our API (`/v1/reports/recent`): still gxrgvz4yzsi2 (8 Nov) and z2opt2x2gs5u (20 Oct). Nothing after the hotfix.
   - foreverlogs.gg: 25 public reports at fetch time (2815 to 2853; 2853 is the newest, 12 Nov 01:37 UTC). The reports starting 10 Nov or later are 2840, 2841, 2842, 2845, 2847, 2849, 2850 and 2853. Every warlock in them and every pet, with the damage, cast, resource, buff and healing streams (`meta_events`, 40 ids a page, paced; `combatants-roster` and `meta_actors` for the players and pets), fetched once into the scratchpad (`wl2/fl/data/<id>.json`). The 2026-11-11 lane's cache of the 23 earlier reports (damage events only) served the by-date crit table.
   - Warlocks with a usable fight after the hotfix: Sträyed (2845 on 10 Nov and 2853 on 11 to 12 Nov; Demonology at 36 to 37 by the ranks he casts, Succubus Zahnia, Searing Pain and Demonic Brand, 607 casts and 97 encounters in 2853), Lacra (2850, 11 Nov; Affliction at 36 or more, Imp Karkin, 258 casts), Valpuri and Freohr (2847, 10 Nov 19:45 UTC on; Imps, re-read with the cast and resource streams this time), Det (2853, about 24, Imp Jubkin, 50 casts). The rest are at 10 to 20 with a handful of casts.
   - Levels come from the ranks cast; no roster snapshot carries a level or a sheet (`primary_stats` and `snapshots` are empty on every row).
2. **Blizzard since 10 Nov.** The hotfixes thread (t/2403040): the 11 Nov post is the 10 Nov warlock hotfixes the last lane modelled; the two 12 Nov posts (the Auction House note and the 11 Nov hotfixes) have no warlock, pet or demon line.
3. **Engine.** Every warlock row the logs touch was read in the fork: `wrack.go`, `talents.go` (Soul Siphon, Master Summoner, Improved Bane of Agony), `curses.go`, `corruption.go`, `siphon_life.go`, `immolate.go`, `drain_life.go`, `lifetap.go`, `imp.go`, `succubus.go`, `pet.go`, `demonic_brand.go`, `summon_demon.go`, and `core/cast.go` and `core/mana.go` for the cooldown and cost paths.
4. **Sims** on the Windows box only (`tools/remote-test.sh --engine-sha`, one job at a time, `nice -n 15`, GOMAXPROCS=5): `spec-breakdown` at 10,000 iterations, seeds 1 and 50001, through `warlocks-logs/run.sh` and the job files beside it; the fork's warlock and core packages; the three warlock ladder goldens.

## The two Affliction spec-breakdown tests

`TestAfflictionCastsWrackWhenItsBuildTakesIt` and `TestAfflictionRaidRotationRunsAgonyAllFightWithoutDoom` (with the other two warlock tests in `sim/cmd/spec-breakdown`) are green on origin/main a38608c67 at pin 15e2c6bae on the box, and green again at the new pin. 5b4dc79a1 (the 2026-11-11 lane) had already moved them to the landed card.

## Logs against the engine

| # | finding | evidence | engine | result |
|---|---|---|---|---|
| L1 | **Firebolt ranks 2, 3 and 4 carry the 10 Nov push.** Det's Imp (rank 2, 7799, owner about 24): 14 a hit non-crit, 7 to 15. The hand row reads 11 at 8, 0.4 a level to 13: 13, spread 11.8 to 14.2, plus the Imp's share of his spell power; the 70790 row would be 8. Valpuri's Imp (rank 3, 7800, owner 36 or more): 19, 5 to 21, against 17 and 15 to 19 plus spell power; 70790 would be 13.5. Lacra's Imp (rank 4, 7801, owner 36 or more): 28, 24 to 31, against 20 plus 8 from his roughly 143 spell power (L5 reads the same figure off his Agony); 70790 would be 19. Freohr's Imp (rank 4 at 32): 52 with Demonic Knowledge and Improved Imp. | 2847, 2850, 2853 | `imp.go` FireboltDamage (W5) | **Aligned.** The last lane's "no log check yet" item for W5 is closed. |
| L2 | **The Imp casts a Firebolt every 2.43 s, not 2.2 s.** Hit-to-hit gaps at a standing target: Noklop 337 gaps, median 2.431 s (quartiles 2.41 to 2.46); Tarjub 180, 2.423 (2.41 to 2.44); Karkin 157, 2.431 (2.42 to 2.44). Firebolt's flight is constant at a standing target, so hit gaps are cast gaps. The Succubus swings every 2.00 s (933 swings) and Lashes every 12.17 s (12 s cooldown, min 12.08). | 2847, 2850, 2853 | `imp.go`: 2 s cast, then the upstream "artificial" 200 ms cooldown counted from the cast's start, so one every 2.2 s | **Bug (old), fixed: B2.** |
| L3 | **The live Imp crits 5.9% of its Firebolts.** Thirteen Imps since launch (2816 to 2853, owners at 12 to 40, 3,501 Firebolts): 207 crits, each Imp between 3.6 and 7.7%. The other demons' spells read as the engine has them: Lash of Pain 60 in 1,061 by eight Succubi (5.7%; engine 6.4%), Demonic Brand 28 in 394 (7.1%). | all cached reports | `pet.go` impSpellCritPercent 2.2 (beta logs: 53 in 2,383) | **Bug (old), fixed: B3.** Blizzard's own line on forever-bugs #331, "base spell crit is around normal levels", said as much. |
| L4 | **DoT clocks.** Corruption ticks every 3.0 s (Sträyed 175 of 177 gaps, Lacra 152 of 156), Bane of Agony every 2.0 s (1,553 of 1,560) for 12 ticks (62 of 62 full runs), Immolate 3.0 s (110 of 112), Drain Life 1.0 s (68 of 70). | 2850, 2853 | `corruption.go`, `curses.go`, `immolate.go`, `drain_life.go` | Aligned |
| L5 | **Agony's ramp is on the base tick alone.** Sträyed's 62 full Agony runs (rank 3, 6217): stage medians 17, 24 and 31 a tick, every run stepping by the same 7 from ticks 1 to 4, 5 to 8 and 9 to 12. Half, whole and one and a half of a 14 base plus a flat 10 fits exactly; the client's rank 3 tick is 14. The flat 10 is his spell power share (0.133 a tick, about 75 spell power, which his Shadow Bolt at 151 also gives). Lacra: 27, 35, 43, a base of 16 and a flat 19 (Improved Bane of Agony 2 on a 14 base and about 143 spell power). A 0.7, 1.0, 1.3 ramp of the whole tick does not fit either (0.68 to 0.77 and 1.22 to 1.29 across four players). | 2845, 2847, 2850, 2853 | `curses.go` agonyRampStage on the base, spell power re-read each tick | Aligned |
| L6 | **Life Tap.** Rank 3 (1456) by Sträyed at 36 to 37: 226 on 52 taps and 218 on 8 (130 plus 10 or 11 for the levels above 26, plus 85 or 86 Spirit; a Spirit buff fell off for the 8). Valpuri rank 2 (1455) 174 on 91. Freohr rank 3 261 on 10 (the last lane's L1). Two of the three at 36 or more still tap rank 3; rank 4 at 36 is a free 70 a tap. | 2847, 2853 | `lifetap.go` | Aligned |
| L7 | **Soul Siphon on Drain Life.** Lacra's non-crit Drain Life ticks: 47 with one counted DoT on the target (5 ticks), 50 with two (77). Rank 2 (8% an effect) predicts 50.5, rank 3 (12%) 52, rank 1 (4%) 49. Valpuri's ticks do not move with his DoTs (26, 27, 26 at 0, 1, 2): no points in it. | 2847, 2850 | `talents.go` soulSiphonMultiplier (W4) | Consistent with the 10 Nov model; the rank is not in the log |
| L8 | **Demonic Brand.** Zahnia (Succubus) consumes brands with her swings and Lash alike: 394 consumes in 2853 (3.3 a Searing Pain), 32 a hit non-crit (2 to 36), crits 1.5x at 7%. The talent's roll at 36 is 29 to 32 plus 0.078 of his Shadow power. 20% of her damage; 13% of his whole. | 2845, 2853 | `demonic_brand.go`: any landed pet hit consumes, Shadow for the Succubus | Aligned (round 22 S8 stays closed) |
| L9 | **Pet share.** Sträyed 35.7% (Succubus; Searing Pain and Demonic Brand build, 44,317 of her 65,124 from swings), Freohr 29.5% (Imp; Searing Pain and Brand), Lacra 11.6% and Valpuri 10.6% (Imp; DoT players). The bare cards at 30 and 40 put the Succubus at 18.3% and 16.3% with a Shadow Bolt filler and no Demonic Brand, raid 60 at 13.0%. Her swing (34 to 46 a hit non-crit on the bare bands; Zahnia 45 to 46 at 36 to 37) and Lash (35 at rank 2; Zahnia 36 to 40) agree. | 2845, 2847, 2850, 2853 | breakdowns in `warlocks-logs/results/` | Not like for like: the logged Demonology leveling style (Searing Pain into Demonic Brand) is not the card's. Open, O2. |
| L10 | **Life Tap rate and mana flow.** Logged: 3% of casts (Freohr), 10% (Sträyed), 12% (Lacra), 19% (Valpuri), on dungeon pulls with drinking between them; Shoot is 11 to 41% of their casts. The bare cards tap 27% of casts over a 180 s pull with no drinking and shoot only out of mana. The per-tap amount is exact (L6); the rate is the fight model's, not a bug. | 2847, 2850, 2853 | the bare breakdowns | Explained |
| L11 | **Succubus melee crit.** Zahnia crits 4.0% of 1,206 swings (both reports, one owner); the Voidwalkers in the same reports crit 0 of 733. The engine gives every demon 0.7% plus its owner's gear crit. | 2840 to 2853 | `pet.go` demonMeleeCritPercent | **Open, O1**: her owner's sheet decides whether it is his gear or a Succubus rate (about 0.3% of Demonology's DPS either way). |
| L12 | **Wrack, the doubled DoT clock under it, and the raid-60 Affliction card.** No logged warlock is 40 or more with Wrack: nobody casts 1316697 in any public report, and the dungeons logged run to Zul'Farrak. | all reports | W1 to W3 | **Still unchecked by a log.** Open, O3. |

## Confirmed and fixed

| # | finding | source | origin | fix |
|---|---|---|---|---|
| B1 | **The pre-pull summons took their mana out of the fight (round 29 W2).** Summon Imp and Summon Succubus cost the warlock 659 and 824 mana at -60 and -44 s, and only the prepull's own regen ticks gave it back (1,356 of 1,835 in the Demonology raid breakdown; `OtherActionManaRegen#2`). In play the demons are summoned a minute out, the warlock Life Taps back to full and a healer tops the health up (demonology.md's pre-pull list, step 4), so the pull starts at full mana. A summon during the fight (a dead Succubus) pays its cost, cut by Master Summoner. | the breakdowns' resource rows; demonology.md | old | E ef85ff869 (`summon_demon.go`: `prepullSummonCost` wraps the mana cost and pays nothing before the pull); tests E dbf4d1f4a, red on the box (688 took 1,098 and 712 1,373 before the pull), per-cast Master Summoner check corrected in E 1279751ce |
| B2 | **The Imp cast a Firebolt every 2.2 s; live Imps cast one every 2.43 s** (L2). | 2847, 2850 | old (upstream's 200 ms) | E b8539dcd6 (`imp.go`: `impCastDelay` 430 ms); test E 1279751ce and 029a92d73, red at 2.2 s (37 casts in 80 s where 34 are due) |
| B3 | **The Imp's Firebolt crit was the beta's 2.2%; live Imps crit 5.9%** (L3). | 13 Imps, 3,501 Firebolts | old | E 5c2705564 (`pet.go`); test E 029a92d73 |

B2 and B3 move no published number: the published bands summon the Succubus from 20, the Voidwalker from 10, and the ladder starts at 10, so the Imp is out in no published fight (Affliction's raid prepull summons it only to sacrifice it). They matter to a player who sims with the Imp on the site and to every leveling walk below 10.

## Rotation and builds (adoption rule)

**Master Summoner (round 29 W2, the Demonology raid build's 2 points).** On the fixed engine its sim value is gone (the summon's cost no longer touches the fight), so the two points were measured against Improved Sayaad 3 with the last point in Demonic Brand or Demonic Energies (`jobs_ms.txt`, 10,000 iterations):

| build | horde s1 / s50001 | alliance s1 / s50001 |
|---|---|---|
| raid build (Master Summoner 2) | 795.50 / 796.58 | 807.26 / 807.72 |
| Improved Sayaad 3, Demonic Brand 1 | 796.06 / 797.15 | - |
| Improved Sayaad 3, Demonic Energies 1 | 796.06 / 797.15 | 807.81 / 808.26 |

+0.07% at both seeds and factions, inside the error of 0.46. The build holds: the gain is under the rule's 1%, and Master Summoner's value is the 6 s resummon when the Succubus dies, which the sim does not play. The guide's line on it ("4 seconds off every summon and 40% off its mana") stays true. Round 29's W2 is closed as a model fix with no build change.

**The 2026-11-11 Affliction card** (Agony all fight, the under-3 s recasts from 40) cannot be judged against a log yet (L12). Nothing in the logs contradicts the DoT, Soul Siphon or Life Tap rows it rests on.

## Before and after

Raid 60, guide raid builds, published BiS, 10,000 iterations, seeds 1 / 50001, pin 15e2c6bae against ef85ff869 (B1; B2 and B3 reach no published fight). `results/before-*.txt`, `after-*.txt`.

| spec | horde | alliance |
|---|---|---|
| Affliction | 728.11 / 729.27, unchanged | 741.03 / 742.24, unchanged |
| Demonology | 795.50 / 796.58, unchanged | 807.26 / 807.72, unchanged |
| Destruction | 869.42 / 869.16, unchanged | - |

Bare, horde: Demonology 30 and 40 unchanged, 50 -0.03% / +0.02%, 60 +0.05% / -0.01%; Affliction 60 unchanged. The prepull's own regen had already refilled most of the cost, so the pull's missing mana was about 130 at raid 60 and no Life Tap moved.

Ladder goldens (the fight the rungs use): Demonology 38 +0.1%, 40 -0.4%, 50 +0.3%, 60 -0.07%; 20 unchanged; Affliction and Destruction byte-identical.

## Sim changes in this landing

- Fork `lane/sim-warlock2` on master 15e2c6bae: dbf4d1f4a (tests, red), ef85ff869 (B1), 1279751ce (tests), b8539dcd6 (B2), 029a92d73 (tests), 5c2705564 (B3), a63d30e26 (the demon crit test reads the live rate; the cadence test inside the Imp's mana). Site pin a63d30e26 (the lane's head on 15e2c6bae). The fork master had moved meanwhile (the mage lane's 71476b429 and 28524da82, which that lane pins itself), so the lane is merged into master rather than fast-forwarded; the site's pin stays the lane head, so no mage change rides in here.
- Shared code: none. `core` is untouched; the fork's core, warlock and warlock/dps packages pass.
- Goldens: `sim/request/testdata/ladder/warlock-demonology.golden.md` regenerated on the box (the other two are unchanged).
- Guides and cards: unchanged. No build or rotation moved, and the Demonology guide's pre-pull list already says what B1 models. The curated cards' prepull rows carry no mana note.

## Open

- **O1.** The Succubus's melee crit (L11): 4.0% on one owner against the engine's 0.7%; needs his sheet, or a second Succubus owner with 500 or more swings.
- **O2.** A Searing Pain and Demonic Brand leveling style for Demonology (L9): the logged players at 32 to 37 play it with the Succubus at 30 to 36% of the damage; the card's filler is Shadow Bolt and the leveling build has no Demonic Brand. A rotation and talent search at 30 and 40, for a later lane.
- **O3.** Wrack and the doubled DoT clock (W1 to W3) still wait for a level 40 or more Affliction log.
- The last lane's leftovers stand: the level-by-level Affliction talent list, the Execute idea, Curse of the Elements' wording, Curse of Recklessness' 505, the Improved Shadow Bolt pick.

## Measure at refresh (ledger MR66)

- Demonology bare bands 38 to 60 on the pre-pull summon's mana (the ladder moved -0.4 to +0.3%); the published raid number does not move.

## Gates

- Fork, on the box: `go test` for sim/warlock/..., sim/warlock/dps and sim/core at a63d30e26 (ok); each new test red first at its test commit (`results/box-fork-tests*.txt`).
- Site, on the box at pin a63d30e26: `go test` for sim/request (the warlock ladder goldens match), sim/cmd/spec-breakdown (the four warlock tests, the two Affliction ones among them) and sim/leveling (ok, `results/gates3.txt`); vitest on `src/lib/sim/version.test.ts`, `api.test.ts` and `details.test.ts` (46 tests, the pin agreement among them). No APL, guide or card changed, so no apl-check, guide suite or wow-player pass was due. Rebased on origin/main and pushed once.

## Sources

- foreverlogs.gg reports 2840 to 2853 (`/api/reports/<id>`, `/combatants-roster`, `/meta_actors`, `/meta_events`), and the 2026-11-11 lane's cache of 2815 to 2849.
- foreversixty API `/v1/reports/recent`.
- Blizzard, WoW: Forever Hotfixes, https://us.forums.blizzard.com/en/wow/t/2403040 (posts to 12 Nov).
- Blizzard on forever-bugs #331 (demon crit), as `pet.go` cites it.
- Client 1.60.1.70790 talent text (`data/builds/1.60.1.70790/talents/warlock.json`: Master Summoner, Improved Bane of Agony, Soul Siphon).
