# Sim-accuracy round 23: paladin

Lane paladin, resumed after the 10:50 rate-limit kill (the killed run had no commits; its grid scripts were reused, every number below re-measured). Site wt/r23-paladin from int/r23 (5278e2dd5: druid and platform merged), fork lane/r23-paladin from int/r23 (c595f424b: manaIncomePerSecond 7bb3c301a, druid f52b25882). Client 1.60.1.70790. Every sim ran through `$R/sim.sh`, frozen BiS (`$R/bis-frozen`), the guide raid build (`$R/rt`) at raid 60 and "bare 60r", the band's published build at bare levels, Holy with `-heal`, Protection with `-tank`. Proto re-vendored locally with `python -m pipeline simproto` for testing, not committed. Build mode (integrator, 12:00 UTC): no wow-player pass; guide text written in the guide voice.

**Result: 1 CONFIRMED-BUG, fixed (`old`, round 22 R2).** Retribution's Consecration now spends on a mana budget (the platform's manaIncomePerSecond), and its rank rows are re-judged on the new line. Protection's half-mana gate stays OPEN: no budget shape separates the levels where Consecration should go.

Lane files: `$R/paladin/` (ex.sh, run.sh, runb.sh, mkbud.py, cmp.py, cmp2.py; grids ex-ret-g*.tsv, exp-g*.tsv; full-ret2.tsv, fin.tsv, rank-sweep.txt; variants v*/).

## Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Consecration's 0.8 pace overspent without raid mana; a mana budget wins or ties everywhere.** What was wrong: the last line cast Consecration while mana percent was at least 0.8 of the fight's remaining percent. A percentage cannot tell a raid paladin (Blessing of Wisdom, Judgement of Wisdom, Mana Spring) from a solo one, so the line spent too early at bare 60 and on most leveling levels. What changed: the pace became `currentMana + manaIncomePerSecond(10 s) x remainingTime >= 0.57 x cost(Judgement) x remainingTime + cost(Consecration)`; the existing reserve (a Seal, Judgement and Holy Strike after the cast) stays. Grid at 3,000 iterations: unit Judgement / Consecration / Seal of Righteousness, c 0.05 to 1.3, windows 5/10/15/20/30, with and without Consecration's own cost (ex-ret-g1..g7). Only Judgement-unit c 0.57 with the own cost had no loss anywhere; 0.55 lost at 54, 0.6 at 41 and 43. 10,000 iterations, seeds 1 and 50001, both factions, every level 20 to 60 plus raid 60 and bare 60 on the raid build: 166 gains beyond error, 6 ties, 0 losses (fin.tsv). Below 20 Consecration is not trained (trainables/paladin.json, rank 1 at 20). Test first: `TestRetributionConsecrationSpendsThePoolOnTheBudget` (red: no manaIncomePerSecond; green). | `$R/paladin/fin.tsv`, `full-ret2.tsv`, grids above | CONFIRMED-BUG (rotation line measured worse beyond error on 160+ cells) | old (round 10 pace; round 22 R2 measured it and waited for the value); guides-handoff no | S b28f02334, E cd9efc091 |
| R1b | **The Consecration rank rows moved with the new line.** Ranks 1 to 4 forced (`-rank 20924=r`) at every level 40 to 59, Human, Dwarf and Undead, seeds 1 and 50001, 10,000 iterations (rank-sweep.txt). Rows adopted on any gain beyond error at both seeds: rank 1 at 40 to 46 (rank 1 beats rank 2 by 0.2 to 1.7% at 44 to 46, where rank 2 used to win); Undead rank 1 at 47; Human and Dwarf rank 2 at 47 to 49 (Undead takes the default rank 3 at 48 and 49, +0.4 to +1.2% over rank 2); rank 1 at 50 to 53 for every race (round 21's Dwarf and Undead rank 4 at 53 now loses 0.3 and 1.0%); Human rank 1 at 54 and 55 (+0.6 to +1.0% over rank 4) and rank 3 at 57 (+0.4/+0.5%); Dwarf rank 3 at 54 to 56 (+0.4 to +0.8%) and rank 1 at 58 and 59 (+0.4 to +0.6%); Undead rank 3 at 54 to 58 (+0.4 to +1.0%). Rank 1 still wins at 30 to 39 (3,000-iteration check at 30, 36, 39); bare 60 rank 4 and raid 60 rank 5 unchanged. Tests: `TestRetributionCastsTheReMeasuredConsecrationRanks`, `TestRetributionConsecrationRankFollowsTheRace` (red then green), `TestChosenRanksReadsTheCuratedRowsAtALevel` (Dwarf's row at 59). | `$R/paladin/rank-sweep.txt`, `rank-recheck.txt` | part of R1 | old | S b28f02334 |
| R2 | **The twist (lines 6 to 8) keeps the 0.8 pace.** Budgets on the twist (Judgement unit c 0.3 to 1.2, windows 3/5/10/20, on top of R1) either lose at bare 60 or lose 0.1 to 0.4% at raid 60 at all four seed/faction runs (tw.tsv, tw2.tsv): at a full bar the income reads 0, so a budget holds the twist back in the opening seconds of a raid fight. t 0.6 gains bare 60 +0.3 to +1.1% but costs raid 60 0.10 to 0.14% at every run; not adopted. | `$R/paladin/tw.tsv`, `tw2.tsv` | Aligned (pace stays; a budget does not win everywhere) | - | - |
| R3 | Guide-accuracy round 24, Sims requested 1: Consecration rank 1 casts at the band floors on the new line, against about 22 available in a 180 s fight (8 s cooldown): 30: 9.5, 40: 18.2, 50: 19.1. "Almost every time" was not true in the 30s; both sentences (retribution.md:97 and :135) now say rank 1 goes down "far more often" than the higher ranks. | bare band floors, 10,000 iterations | Aligned (wording fixed with R1) | guides-handoff | S (report commit) |
| R4 | Guide-accuracy round 24, Sims requested 2: Human against Dwarf on the raid build. Raid 60: Human 772.7/772.5, Dwarf 775.1/774.9 (Dwarf +0.3%); bare 60: Human 299.7/299.8, Dwarf 297.5/297.5 (Human +0.7%). retribution.md:113 "Dwarf matches Human on the raid build in a raid, and Human is ahead solo" holds (within 1%, race picks stay). | 10,000 iterations, seeds 1/50001 | Aligned | guides-handoff | - |
| R5 | Fresh look: Sacred Arbiter (+20% Holy Strike, refreshes every Judgement on the target on a landed hit, client talent 1311087) as `holy_strike.go:121-180` has it; round 22's unknown-aura change (fork 30bf8bdb0) re-read against every Retribution aura condition: the Unsealed Wrath lines are behind auraIsKnown, the Crusader opener reads an aura the raid preset or your own Judgement applies, Seal of Righteousness is known from level 1. Nothing flips. | engine read | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Half-mana Consecration gate (row 12): still OPEN.** Budgets (unit Judgement, Holy Strike or Consecration; c 0.1 to 1.3; windows 10/30/60; with own cost; OR'd with the half-mana gate; OR'd with mana above 80/90/95%) all gain at the band floors 20, 30, 40, 50 and bare 60 (+2 to +7%) and raid 60 (+0.2%), and all lose 1.3 to 3.6% at 26 to 29 and 0.8 to 1.9% at 36 to 39 (exp-g1..g5). The breakdowns say why: at 20 and 30 the half-mana Consecrations take mana from Holy Strike (16.6 against 18.1 casts at 20), at 26 and 39 the same two or three Consecrations are free (core casts unchanged, +1.4 to 2 threat per second). The difference follows the band build and the Holy Strike rank, not the mana income, so an income budget cannot read it. Needs a spend-rate value (mana spent per second over a window) or per-level rows; see For other lanes. | `$R/paladin/exp-*.tsv`; variants `vp*/` | OPEN | - | - |
| P2 | Fresh look: the unknown-aura change (30bf8bdb0) touches no Protection line (no aura conditions). Tank deaths at bare 23 to 39 are 98 to 100% without a healer, as before; threat is measured over the fight either way. | breakdowns | Aligned | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | Guide-accuracy round 22, Sims requested 1: the emergency Holy Light (line 3, tank below 40%) at 20 to 24 at the top rank trained instead of rank 1 (a temporary build scoping the 20 to 24 row to line 4 only). Rank 1 wins: 20: 48.8 against 42.8 effective HPS, 22: 48.4 against 44.3, 24: 50.4 against 50.1, both seeds and factions. The row stays; the card note fix (R23-1) stands. | 10,000 iterations, seeds 1/50001, both factions | Aligned | guides-handoff | - |
| H2 | Guide-accuracy round 22, Sims requested 2: the party Flash of Light (19941) ranks 3, 4 and 5 at 50, 52, 54: every rank within 0.1 HPS of the default at both seeds and factions (the bare fight barely casts the party lines). No row. | 10,000 iterations | Aligned | guides-handoff | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | Blizzard's Forever notes (https://us.forums.blizzard.com/en/wow/t/2360696.json, fetched 12:50 UTC 5 Nov): 10 posts, the last edit 4 Nov 19:35 UTC, identical to round 22's read; no new paladin line. A web search for later Forever hotfix notes found none. | thread JSON | Aligned | - | - |
| N2 | Live logs, last looked 12:51 UTC 5 Nov: api.foreversixty.gg/v1/reports/recent lists only the 20 Oct report; foreverlogs.gg answers 403. Seal of Righteousness rank 8 procs and Oath of Light on a low rank stay log-gated. | live checks | OPEN (log-gated) | - | - |

## Before/after

Retribution, 10,000 iterations, seeds 1 / 50001, frozen BiS, before = this branch before b28f02334:
- Raid 60 (raid build): alliance 772.9/772.5 to 772.7/772.5, horde 779.8/779.8 to 779.8/779.8 (tie).
- Bare 60, raid build: alliance 297.9/298.0 to 299.7/299.8 (+0.6%), horde 296.8/297.5 to 299.3/299.2 (+0.6 to +0.8%).
- Bare band 60: 309.4 to 311.7 alliance (+0.7%), 306.3 to 308.3 horde (+0.7%).
- Levels: 20 to 29 +0.4 to +2.6%, 30 to 39 0 to +0.5%, 40 to 49 +0.15 to +2.4%, 50 to 59 +0.4 to +2.1%. No level loses.
Holy and Protection unchanged. No commit moves DPS outside paladin.

## Files to regenerate

- data/proto (apl.proto, apl_pb2.py, ENGINE_SHA) via `pipeline simproto` at the pin that carries 7bb3c301a (platform already listed it; this lane's curated APL needs it for test_apl).
- addon-data, then Data.lua (addonrotation): Retribution's Consecration line, its card note and the twist note; the Consecration rank rows.
- sim/request ladder golden testdata/ladder/paladin-retribution.golden.md (TestRotationLadder).
- Fork: nothing beyond the APL copy (cd9efc091); fork notes on that copy lag the curated notes (notes are curated-only).

## Measure at refresh

- Retribution raid 60 on the re-ranked gear with the budget line (it ties on the frozen gear).
- Carried: Retribution gear under the paced twist; Holy raid 60 party Flash of Light row.

## For other lanes

- **Platform:** Protection's Consecration needs a spend-rate value (mana spent per second over a trailing window) so a line can read "your income covers what you spend"; income alone cannot (P1). Also: manaIncomePerSecond reads 0 at a full bar, which is what keeps a budget off the Retribution twist at the pull (R2); a "mana that overflowed" or "regen rate" value would fix it.
- Coach starters (generated): tools/coach/starter_overrides.py's Retribution Consecration line is still true; no regen needed.

## For the guide lane

- Shipped: retribution.md:93 (twist no longer "the same pace as Consecration"), :97 (the budget rule in plain words, the new ranks by race, "far more often"), :135 (leveling ranks); the Consecration card note (first sentence states the rule) and the twist card note. `_prose.test.ts` passes. No wow-player pass (build mode).

## Still open

- P1 (Protection Consecration gate; needs a spend-rate value).
- Seal of Righteousness rank 8 procs and Oath of Light on a low rank (logs).
- Round 20 P3 (Protection BiS use effects, platform).

## Sources

- Blizzard, WoW Forever Beta Development Notes, https://us.forums.blizzard.com/en/wow/t/2360696 (JSON, fetched 2026-11-05 12:50 UTC).
- Client 70790: data/builds/1.60.1.70790/trainables/paladin.json (Consecration ranks 20/30/40/50/60), talents/paladin.json (Sacred Arbiter 1311087).
- Platform note `$R/platform/for-paladin.md` (manaIncomePerSecond semantics).
