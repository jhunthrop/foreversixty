# Sim-accuracy round 10: ranks lane (owner ruling 2)

Ruling: "Retribution's Consecration rank 1 at 30/40 was adopted only because the rotation rations Consecration behind a mana check, and Ret never runs out of mana there. That is a rotation artefact. Remove the needless mana gate, re-measure the rank honestly, and keep the rank only if it wins beyond error."

Where it was measured: site branch `wt/r10-ranks`, fork branch `lane/r10-ranks` (the rotation copy only; no engine code changed). Horde, the published gear and preset per band, seeds 1 and 50001, 10,000 iterations unless marked. "Before" is `$R/sb-before` run from `$R/site-before/sim` (origin/main bb619a188, engine 4d3129cae). Tools and raw output: scratchpad `r10/ranks/` (`m.sh`, `reconfirm.py`, `crosscheck.py`, `ret1.txt` to `ret7.txt`, `ba.txt`, `reconfirm.jsonl`; the mana tracer `zz_manatrace_scratch_test.go` reads the engine's debug log of one iteration, and is not committed).

## Retribution: the gate and the rank

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **The gate.** Consecration, the last line of `data/curated/apl/paladin-retribution.json` (fork copy `ui/retribution_paladin/apls/forever_retribution.apl.json`, identical), was cast only while mana was above 65%. | APL line 368 on origin/main. | - | - |
| R2 | **With the gate, Retribution never runs dry, and it wastes mana.** Three seeded single-iteration traces per band: the lowest the pool falls is 48% (30), 50% (40), 6 to 17% (50), 23 to 30% (bare 60), 19 to 22% (raid 60), and it ends every fight with that much unspent. | Trace: `Spent/Gained ... (a --> b)` lines of the engine's debug log, scratchpad `r10/ranks/`. | CONFIRMED-BUG (rotation) | S d6cf5dca3, E 2994f28d5 |
| R3 | **Without the gate, Retribution does run dry.** The owner's premise holds only because of the gate. With no guard, Consecration drains the pool to zero and the seal lapses: out of mana 112 s (bare 20), 115 s (30), 100 s (40), 76 s (50) of a 180 s fight at the default rank; 28 s at bare 60. Plain removal changes the published number by -5.6% (20), -0.1% (30, with rank 1), +1.7% (40, rank 1), +0.1% (50, rank 1), -1.0% (bare 60), +2.45% (raid 60). | `ret1.txt`. Traces fall to 0 to 16 mana at every band. | side effect | - |
| R4 | **Replacement line (one rotation serves every band).** Consecration is cast while the mana left after it still pays for the seal, Judgement and Holy Strike, and while mana percent is at least 0.8 times the fight's remaining percent. No flat mana percentage. The pool now runs down to 0 to 7% by the end of each fight, without the seal lapsing early. | Tried at 3,000 iterations: flat 30% and 45% lines, pure pace (x1.0), pace x0.8, reserve only, reserve plus pace x1.0, x0.8, x0.5 (`ret2.txt` to `ret5.txt`). Reserve plus pace x0.8 is the only one that gains at every band but 20. Test: `TestRetributionConsecrationSpendsThePoolOnThePace` (failed first on the 65% line). | CONFIRMED-BUG fixed | S d6cf5dca3, ffaa83593; E 2994f28d5, 557e5a0ec |
| R5 | **Rank 1 at 30 to 59 is a real mana-bound win, kept.** On the new line, rank 1 against the band's top rank: +1.13% / +1.13% (30), +2.39% / +2.49% (40), +7.57% / +7.56% (50). With no guard at all it also wins: +2.8% (30), +1.8% (40), +12.8% (50), where the top rank runs dry for 75 to 115 s. Rank 2 at 40 and ranks 2 and 3 at 50 lose to rank 1. | `reconfirm-ret.jsonl`, `ret1.txt`, `ret7.txt`. Errors are 0.03 to 0.06%. | Aligned (row kept) | S 776f5679e |
| R6 | **New bare-60 row: Consecration rank 4.** On the new line the bare-60 fight runs dry for 17 s at rank 5; rank 4 (20923, 435 mana against 565) gains +1.60% / +1.64%; rank 3 +1.1%. Raid 60 keeps rank 5: ranks 2 to 4 lose 1.0 to 3.4%, and rank 1 loses 4.4%. | `ret7.txt`, `reconfirm-ret.jsonl`. Test: `TestRetributionCastsTheReMeasuredConsecrationRanks` (failed first without the row). | adopted | S 776f5679e |
| R7 | **Ranks' own amounts.** Consecration ranks 1 to 5 in `sim/paladin/consecration.go` (cost 135/235/320/435/565, tick 2/3/6/8/12, first-targets 4/7/11/20/27, coefficient 0.042) match the 70707 client: spellconst/paladin.json for 26573/20116/20922/20923/20924, and SpellEffect rows for 1280345 to 1280349 (effect 0 and 1 base points, effect 1 coefficient 0.04199999943). The hotfix push carries no rows for them. | scratchpad `tcp/db2-70707/SpellEffect.csv`; `data/builds/1.60.1.70707/hotfix-push.json`. | Aligned | - |
| R8 | **Guide and notes.** The Retribution guide said Consecration was absent from the bare rotation and cast only above 65% mana. The one rotation casts it last on every band. The guide now describes the paced line and the lower ranks, with no nightly numbers. The APL top note's old "Consecration is left out" paragraph is gone. The line's note leads with the rule, because the addon and coach show only its first sentence. | `web/src/content/guides/paladin/retribution.md` lines 51, 73, 91. | CONFIRMED-BUG (stale prose) | S eb8f88c65, ffaa83593 |

Why the 65% line was wrong: it rationed a filler by a flat share of the pool. That share was too high for the raid fight, which has mana to spare, and it never let the bare fights spend their last third. The paced line spends the pool by the end of the fight whatever the pool's size, and the reserve keeps the core buttons castable.

## Every other non-healer row: an artefact of a gate?

I listed every APL line that compares mana (`currentManaPercent`, `currentMana` or `remainingTimePercent`), then checked each rank-choice row against it. A row is an artefact only if the spell sits behind a gate and the spec never runs dry.

| # | spec, rows | gate on the row's line? | runs dry without rows? | verdict |
|---|---|---|---|---|
| G1 | druid-balance 30: Moonfire 3, Wrath 4, Insect Swarm 1 | Wrath is the "behind the pace" fallback (`not (mana% > remaining% or Innervate ready)`). The gate decides when Wrath is cast; the rank decides its cost. | yes, OOM 49 s | real, kept |
| G2 | hunter-beast-mastery 40 and bare 60; hunter-marksmanship bare 60: Aimed Shot 1 | No. Aimed Shot is only the pre-pull opener (the 25% gate sits on Arcane Shot). | yes, OOM 86 to 104 s (BM), 46 s (MM) | real, kept |
| G3 | mage-arcane 30: Frostbolt 4 | No (the pace gate is on Arcane Blast's stack line, not Frostbolt). | yes, OOM 22 s | real, kept |
| G4 | mage-fire 30 to 59: Scorch 1; 40: Fire Blast 4 | No (Evocation only). | yes, OOM 56 to 76 s | real, kept |
| G5 | paladin-protection 40 to 59 and bare 60: Consecration 1 | Yes: above 50% mana and Consecration plus Holy Shield affordable. | yes, OOM 5 to 19 s **with** the gate | real, kept. Re-measured with the 50% part removed: rank 1 still wins (+1.9% at 40, +1.9% at 50, +2.6% at bare 60, 3,000 iterations), and removing the 50% part loses 1.9 to 3.5% threat at every band and rank, so that gate stays. |
| G6 | priest-shadow 30 to bare 60 | No mana gate on any line. | yes, OOM 19 to 92 s | real, kept |
| G7 | shaman-elemental 30 to bare 60 | No mana gate on any line. | yes, OOM 80 to 102 s | real, kept |
| G8 | warlock-destruction 30: Shadow Bolt 4 | No (Life Tap below 10% only). | yes, OOM 2.8 s | real, kept |
| G9 | shaman-enhancement | 50% gate on Earth Shock; no rows (round 9 did not adopt Strength of Earth rank 1). | no | nothing to do |

Every row, re-confirmed on origin/main bb619a188 against the same tree with no rows (`sb-norows`): every one still passes the beyond-error rule at both seeds, with the round-9 gains unchanged to the hundredth. That covers every non-healer row and the bare-60 healer rows, which are mine (druid-restoration, paladin-holy, priest-discipline, priest-holy and shaman-restoration at bare 60: +5.2%, +4.8%, +6.9%, +11.0% and +12.8%). Sub-60 healer rows belong to the healprofile lane and were not touched. Source: `reconfirm.jsonl`.

## Do all consumers resolve the same rank? (after round 9)

| # | finding | evidence | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | /sim, the ranker, the nightly, spec-breakdown and the ladder share one function. /sim builds through `request.rotation` with `rankPreset(req)`. The ranker sends `Preset: spec.rankPreset()` (raid pass raid, every other pass bare; `sim/cmd/leveling-bis/preset.go:68`). spec-breakdown sends `-preset` and rewrites `-rotation` files through `RewriteRotationRanks`. The ladder calls `rankResolver(..., PublishedPreset(level))`. /sim sends the picked preset (7291b4b9d). Custom names none and falls back to the level's published preset. | code read | Aligned | - |
| C2 | For every spec with rows, at levels 20/30/38/40/50/60 under both presets, `request.rotation` and `RewriteRotationRanks` on the curated file produce the same cast ids. Every rank-choice row's chosen id appears in both. That is 200 checks: 0 differ. | `crosscheck.py` over a dump from a scratch test | Aligned | - |
| C3 | The addon (`data/pipeline/addonrotation.py build_rotation`) casts every sub-60 bare row's id. The one miss is hunter-beast-mastery's Aimed Shot at 40, which the addon never lists at any band because it is a pre-pull opener; the addon shows priority lines only. At 60 the addon shows the raid list, so bare-60 rows do not reach it, by design. | `crosscheck.py` | Aligned | - |
| C4 | The committed `addon/ForeverSixty/Data.lua` still shows Retribution's Consecration at rank 2/3/4 below 60 and the 65% note. It predates round 9's rows, and the data workflow is held. The pipeline output is right (C3). It regenerates at the next nightly. | `Data.lua:1327-1374` | nightly-owned, not committed | - |

## Before/after (paladin-retribution, published rotation)

| band | before, seed 1 | after, seed 1 | change seed 1 / 50001 | out of mana, s (before to after) |
|---|---|---|---|---|
| bare 20 | 38.97 | 38.79 | -0.47% / -0.26% | 0.1 to 7.2 |
| bare 30 | 70.4 | 71.9 | +2.1% / +2.1% | 0.0 to 0.0 |
| bare 40 | 101.8 | 107.0 | +5.1% / +5.2% | 0.0 to 0.0 |
| bare 50 | 165.9 | 166.3 | +0.24% / +0.24% | 0.1 to 6.2 |
| bare 60 | 241.4 | 248.2 | +2.8% / +2.9% | 0.0 to 9.4 |
| raid 60 | 621.3 | 636.2 | +2.4% / +2.6% | 0.0 to 0.2 |

Bare 20 has only rank 1 and a tiny pool. The 65% line rationed it slightly better than the paced line does (-0.47% / -0.26%, beyond error). No rank choice can help there. I kept one line for every band rather than add a level-specific gate.

At 60, without the rank row: removing the 65% line outright is a gain at raid (+2.45%) and a loss at bare (-1.0%). The paced line gains at both (raid +2.4%, bare +1.2% at rank 5, +2.8% with the new rank-4 row).

No other spec moves: the change is to paladin-retribution's rotation, its rank rows and its guide.

## For other lanes

- **Integrator:**
  - The `TestRotationLadder/paladin-retribution` golden moves. Do not regenerate here.
  - The ladder's level-60 row falls 187.3 to 180.2 and its 50 row 103.4 to 101.8. The ladder's character has weak gear and casts raid ranks at 60, and its smaller pool paces lower.
  - The ladder also gains one more `zero_casts` for Seal of Command at 60; it already had them at 30 to 50. The twist lines wait below 30% mana, and the paced line keeps the pool lower for longer. On the published raid gear the twist still fires (22.9 Seal of Command casts against 27.4 before), and the net is +2.4%.
  - Preview diff: scratchpad `r10/ranks/ladder-ret.diff`.
- **Integrator / healprofile:** both lanes edit the one-line `notes` of data/curated/rank-choices.json. This lane rewrote the "adopted where" sentence and added the gate sentence. Take both sides when merging, then run `make rank-choices-sync`.
- **Platform (coach):** `tools/coach/starter_overrides.py:79-83` still says "Consecration last, and only while mana is above sixty-five percent. Against one target it is a loss without the raid's mana support." Both halves are now wrong. Suggested: "Consecration last, while the mana left after it still pays for your seal, Judgement and Holy Strike. Do not let your mana fall far behind the fight."
- **Hybrids / crit:** at raid 60, Consecration crits 25.7% while the melee table shows its own crit. I did not audit it; it is outside this lane.

## Still open

- Bare 20 Retribution loses 0.3 to 0.5% to the old 65% line. A level-aware line could recover it, but I chose not to add a level gate for 0.5%.
- The pace factor 0.8 was picked from 0.5 / 0.8 / 1.0 on three presets. Bare 60 would prefer a stricter line (a flat 45% line measured +3.8% over the old line there, but sat 1.1% below the paced line at raid). An exact per-preset optimum would need a preset value in the APL, which the engine does not have.
