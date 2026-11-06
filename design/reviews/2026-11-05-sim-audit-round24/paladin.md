# Sim-accuracy round 24: paladin

Lane paladin, resumed after rate-limit kills (no earlier commits; every number below re-measured on this branch). Site wt/r24-paladin from int/r24 (a9573ab6e, platform merged), fork lane/r24-paladin from int/r24 (3d706f6a7, manaSpentPerSecond 568d4a7fe). Client 1.60.1.70790. Every sim ran through `$R/sim.sh`, frozen BiS (`$R/bis-frozen`), the guide raid build (`$R/rt`) at raid 60 and "bare 60r", the band's published build at bare levels (below 20: band 20 with the leveling order's talents), Holy with `-heal`, Protection with `-tank`, 10,000 iterations, seeds 1 and 50001, both factions (Protection Dwarf/Undead, Holy and Retribution Human/Undead). Build mode: no wow-player pass; guide text written in the guide voice. Proto re-vendored locally with `python -m pipeline simproto` for testing, not committed.

Comparator: means at both seeds against the run's reported error (2 x the combined standard error). spec-breakdown now prints that error on threat /s and effective HPS (X1).

**Result: 4 CONFIRMED-BUGs, all fixed (`old` x4).** Protection's Consecration gate (round 23 P1) is closed with the platform's spend-rate value; three rank rows came out of the rank sweep.

Lane files: `$R/paladin/` (STATE.md; run.sh, runb.sh, pgrid.sh, rsweep*.sh, mkp.py, cmp2.py; grids g3..g11.tsv, fin.tsv, fin-raid.tsv; rank sweeps rs3/rs5/rs6/rs7.tsv; Holy hl2.tsv; variants v/, vh/).

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Consecration behind "mana above 50%".** What was wrong: the last line cast Consecration whenever the bar was over half. At the band floors that spent mana Holy Strike needed; at other levels it sat on mana nothing else would use (round 23 P1). What changed: `currentMana + (manaIncomePerSecond - manaSpentPerSecond, 60 s) x remainingTime >= cost(Consecration)`, OR health below 35%; the Holy Shield reserve clause stays. Shapes tried (mkp.py): income keeps up with spend at 5/10/20/30 s, x0.8/1.2, OR mana above 80/90%; the projection at 10/20/30/60 s, reserve 1/1.5/2 casts, OR 90%; each with and without the low-health clause. Without the low-health clause every shape loses 1.3 to 3.6% at 26 to 29 (tanks die in the bare fights there; held mana is wasted); the "mana above 90%" clause costs 0.15 to 0.3% at Dwarf 27 to 29. The 60 s projection with the health clause won: 10,000 iterations, every level 20 to 60, raid 60 and bare 60r, both seeds and factions, 149 gains beyond error, 22 ties, 1 loss (Dwarf 29 seed 50001 -0.27%, seed 1 level: a 0.1 threat/s cell, not a loss at both seeds). Test first: `TestProtectionConsecrationSpendsWhatTheNetRateCovers` (red: flat gate, no net rate; green). | `$R/paladin/g3..g7.tsv`, `g10.tsv`, `fin.tsv` | CONFIRMED-BUG (line measured worse beyond error on 149 cells) | old (round 10 gate; round 22 P2 and round 23 P1 measured it and waited for the value) | S 298549a09, E b83e54898 |
| P1b | **Consecration rank rows re-judged on the new line** (`-rank 20924=r`, 10,000 iterations, both seeds, Dwarf and Undead): rank 1 beats rank 2 by 0.3 to 0.7% at 30 (Dwarf +0.4/+0.6% at 31, Undead level; 32 one seed only); ranks 2 and 3 lose 0.3 to 0.8% at 40 to 46, tie at 47 and 48; rank 3 beats rank 1 by 0.33 to 0.42% at 49 (all four runs); ranks 3 and 4 lose at 50 to 55, tie at 56 and 57; rank 4 beats rank 1 by 0.2 to 0.5% at 58 and 0.5 to 0.8% at 59 (this is what removes the 59 loss the line alone had, -0.2 to -0.5%); bare 60 keeps rank 1. Rows: 30-31 r1, 40-48 r1, 49 r3, 50-57 r1, 58-59 r4, bare 60 r1. Test: `TestProtectionCastsTheReMeasuredConsecrationRanks`. | `g4.tsv`, `g8.tsv`, `g9.tsv`, `g11.tsv` | part of P1 | old | S 298549a09 |
| P2 | Measurement note: `spellIsReady` on a Lay on Hands rank the engine does not register reads nil and drops out of an `and` (sim/paladin/lay_on_hands.go registers only the highest rank; the rank rewrite keeps an authored 633 at rank 1). A "Lay on Hands is ready" clause authored at 633 was silently "health below 35%" from 30 on; authored at 10310 it measures level with the plain health clause (171 of 172 cells tie), so the shipped line has no Lay on Hands term. Nothing shipped used it. | `g10.tsv` | Aligned (not shipped) | - | - |
| P3 | **Holy Strike rank 1 while leveling.** The sweep after rank 2 arrives at 12: rank 1 makes 2.6 to 3.6% (Dwarf) and 3.6 to 5.6% (Undead) more threat at 12 to 14, both seeds; at 15 the Undead keeps +2.4% and the Dwarf loses 2.1%; 16 to 19 rank 2 wins (0.3 to 3.1%). The 5-mana rank leaves mana for Seals and Judgements, the same reason as the 40-42 row. Rows: 12-14 r1, 15 r1 Undead. Test first: `TestProtectionHolyStrikeRankOneWhileLeveling` (red, green). | `rs6.tsv`, `rs7.tsv` | CONFIRMED-BUG (rank row beyond error, both seeds) | old (no round swept Protection below 20) | S 215ee9421 |
| P4 | Rest of the sweep: Seal of Fury rank 1 at 18/19 loses 13 to 17%; Holy Strike at 20/21 to 60, Seal of Fury 26 to 59, Holy Shield 40/41, 50/51, 60: default wins or ties (round 24 morning grid rs1, re-read). | `rs1.tsv`, `rs6.tsv` | Aligned | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | Guide round 26 R26-4 (Sims requested S1): the emergency Holy Light (row 3, tank below 40%) at the top trained rank (3 at 18 to 21, 4 at 22 to 24) against the card's rank 1 (row 4 kept at rank 1). Rank 1 wins at every level 18 to 24, both seeds and factions: 18 -7.7 to -8.5%, 20 -12.3 to -12.5%, 23 -6.6%, 24 -0.6 to -0.8% for the top rank. The card stays; holy.md:113 told the player to keep rank 3 for that heal and now says rank 1 heals more over a fight even on the tank below 40%. No split of the row needed. | `hl2.tsv` | Aligned (guide text fixed) | guides-handoff | S e35e80704 |
| H2 | Rank sweep: Holy Light rank 2 at 14/15 loses 6.6 to 14.7%; at bare 30 to 39 the fight casts no Holy Light at all (only Flash of Light: forcing ranks changes nothing, breakdown checked); 46 tie; 54 rank 7 loses 0.2 to 0.4%; 55 mixed (two seeds +0.1/+0.2, two level); Holy Shock rank 3 at 56/57 identical. No row. | `rs3.tsv` | Aligned | - | - |

## Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Hammer of Wrath rank 1 at 52 and 56.** Rank 1 beats rank 2 by 0.20 to 0.35% at 52 at all four runs and by 0.34/0.39% Human at 56 (Undead +0.05/+0.10, level); 53 two of four, 54, 55, 58 lose, 60 rank 3 beats rank 2 by 0.2 to 0.5% and rank 1 by 1%. Rows 52 r1 and 56 r1 (no race loses, so no race key). The casts differ (DPS moves), so the rank is cast. Test first: `TestRetributionHammerOfWrathRankFollowsTheSweep` (red, green). | `rs3.tsv`, `rs5.tsv` | CONFIRMED-BUG (rank row beyond error, both seeds) | old (round 16 wave 2 sweep missed 52 and 56) | S 2d776996d |
| R2 | Rest of the sweep: Seal of Righteousness rank below the new one at 10/11, 18/19, 26/27 to 58/59 loses 0 to 3.4%; Holy Strike rank below at 12/13 to 60 loses 0.8 to 1.7%; Seal of the Crusader 42/43 ties (one run +0.21%); Seal of Command ranks at 30 to 60 identical (the bare band builds never cast it; the twist is raid-only); Consecration 30/31 is already rank 1. | `rs1.tsv`, `rs3.tsv` | Aligned | - | - |
| R3 | Raid 60 on this branch against $R/before.tsv: identical at all four runs (772.7/772.5, 779.8/779.8); the platform's duration-const fix moves no paladin line. | `fin-raid.tsv` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Tool: spec-breakdown's tank and heal reports printed only the mean, one decimal; a 0.1 step is 0.25% at leveling threat. Both lines now print the mean to two decimals with its standard error (stdev / sqrt(iterations), as the DPS line). Tests `TestTankAndHealMarkdownReportTheMeanError`. Changes the text of every tank and healer spec-breakdown report (format only, no numbers). | sim/cmd/spec-breakdown/tank.go:163, heal.go:126 | tool | - | S 441c52e28, S 81d99ccf2 |
| X2 | Blizzard's Forever notes (https://us.forums.blizzard.com/en/wow/t/2360696.json, fetched 23:27 UTC 5 Nov): 10 posts, last edit 4 Nov 19:35 UTC, same as round 23; no new paladin line. | thread JSON | Aligned | - | - |
| X3 | Live logs, last looked 23:27 UTC 5 Nov: api.foreversixty.gg/v1/reports/recent lists only the 20 Oct report; foreverlogs.gg serves its web app (200) but its /api answers 403; Warcraft Logs not tried (round 23: key 401). Seal of Righteousness rank 8 procs and Oath of Light on a low rank stay log-gated. | live checks | OPEN (log-gated) | - | - |

## Before/after

10,000 iterations, seeds 1 / 50001, frozen BiS, before = $R/before.tsv (raid 60) and this branch before 298549a09 (levels):
- Protection raid 60 (threat /s): Dwarf 1125.6/1125.8 to 1129.2/1127.8 (+0.32/+0.17%), Undead 1123.8/1123.3 to 1127.1/1126.8 (+0.30/+0.32%).
- Protection bare 60r: Dwarf 484.3 to 492.2 (+1.6 to +1.7%), Undead 439.5 to 449.8 (+2.35 to +2.5%).
- Protection levels: 20-29 -0.27 to +2.7% (the one loss above), 30-39 +0.05 to +5.4%, 40-49 +0.04 to +2.8%, 50-59 +0.25 to +2.7%. 12-15: +2.4 to +5.6% (P3).
- Holy and Retribution raid 60 unchanged (701.6/702.0, 684.6/683.9; 772.7/772.5, 779.8/779.8). Retribution bare 52 +0.2 to +0.35%, 56 +0.35% Human.
- No commit moves DPS for specs outside paladin (X1 changes report text only).

## Files to regenerate

- data/proto (apl.proto, apl_pb2.py, ENGINE_SHA) via `pipeline simproto` at a pin carrying 568d4a7fe: the Protection APL now uses manaSpentPerSecond, and `data/tests/test_apl.py` rejects it until then (passes locally on the re-vendored proto).
- addon-data, then Data.lua (addonrotation): Protection's Consecration line and note, the Holy Strike note, the Consecration and Holy Strike rank rows; Retribution's Hammer of Wrath note and rows.
- sim/request ladder goldens testdata/ladder/paladin-protection.golden.md and paladin-retribution.golden.md (TestRotationLadder).
- Coach starters: tools/coach/starter_overrides.py:260 still says "Consecration only above half mana, to spend mana you will not need." (rendered at data/curated/coach/paladin-protection.md:32). Not touched (tools/coach is off limits to lanes); suggested: "Consecration last, while your mana plus what you regain less what you spend still pays for it by the end of the fight, or below 35% health." then `persona_starter.py --all`.

## Measure at refresh

- Protection BiS and weights on the new Consecration line (mana stats are worth more now that Consecration reads the income).
- Carried: Retribution gear under the paced twist; Holy raid 60 party Flash of Light row.

## For other lanes

- **Platform / integrator:** the card has no words for the adopted shape (projection: pool plus income less spend times remaining time against a cost, OR'd with a health clause); the note's first sentence carries the rule. A reader in `web/src/lib/bis/apl-condition.ts` for `currentMana + (income - spend) x remainingTime >= cost` would let the card print it.
- **Platform / engine:** `spellIsReady` (and any spell value) on a rank the class registers only at its top rank reads nil and drops out of an `and` silently (P2, Lay on Hands). A validation warning when an APL names an unregistered rank of a known spell would have caught it.
- Anyone parsing spec-breakdown's tank or heal text: the threat /s and effective HPS lines now read `mean +/- error` with two decimals (`$R/measure.sh` takes the first number and is unaffected).

## For the guide lane

- Shipped: holy.md:113 (the emergency heal keeps rank 1; R26-4), protection.md:99 step 8 (the net-rate rule, the low-health clause and the Consecration ranks), protection.md:131 (Holy Strike rank 1 to 14, Undead 15), retribution.md:99 (Hammer of Wrath rank 1 at 52 and 56); card notes for Protection Consecration and Holy Strike and Retribution Hammer of Wrath. `_prose.test.ts` passes.
- Not touched, guide round 26's own rows on the same paragraph: R26-1 ("carries all of your healing power" at 18 to 29 is false) and R26-2 ("From 30 your newest rank takes over"), holy.md:113.

## Still open

- Seal of Righteousness rank 8 procs and Oath of Light on a low rank (logs).
- Round 20 P3 (Protection BiS use effects, platform).
- Guide round 24 optional 3: Retribution raid order twist at 40 to 58 on the bare band (not run).

## Sources

- Client 70790: data/builds/1.60.1.70790/spellranks.json (Consecration 20/30/40/50/60, Holy Strike 6/12/20/28/36/44/52/60, Hammer of Wrath 44/52/60, Lay on Hands 10/30/50), trainables/paladin.json.
- Engine: sim/paladin/lay_on_hands.go (top rank only, drains the bar), sim/core/apl_values_spell.go:60 (unregistered spell reads nil).
- Platform note `$R/platform/for-paladin.md` (manaSpentPerSecond semantics and shapes).
- Guide round 26 paladin report (gr26 scratch), R26-4 and Sims requested S1.
