# Healer ranker robustness, 2026-10-08

Lane heal-rank. Follows design/reviews/2026-10-08-heart-of-the-lion.md item 6 (restoration druid's new pick published a lower HPS than its old set) and 2026-10-08-ranker-set-completion.md.

## 1. Reproduction (restoration druid, band 60 raid, engine f3a673d10, 2000-class iterations)

The committed set is `data/builds/1.60.1.70009/bis/druid-restoration.json` (built on engine 35be2e187); the "before" pick is the unmodified ranker at main. Both sets simmed on the same engine, talents and seeds (7, 11, 13 agree to the first decimal), 2000 iterations. Errors are one standard error.

| Faction | Set | Effective HPS | mana_lasts_sec | Overheal | Guarded score |
|---|---|---|---|---|---|
| alliance | committed set | 622.2 +/- 0.3 | 275.0 | 3.7% | 523.0 +/- 0.3 |
| alliance | before pick | 598.6 +/- 0.3 | 280.9 | 3.5% | 524.8 +/- 0.3 |
| horde | committed set | 619.3 +/- 0.3 | 272.9 | 3.7% | 512.5 +/- 0.3 |
| horde | before pick | 583.6 +/- 0.3 | 284.4 | 3.4% | 524.5 +/- 0.3 |

## 2. Cause

The brief's suspicion (noise let a worse set through) is only partly right, and the finding matters more than the suspicion.

1. **The pick was not worse by the ranker's own objective.** The healer objective is HPS times the squared share of the 300 s fight the mana lasted. At 2000 iterations the before pick scores 524.8 against 523.0 (alliance, a tie) and 524.5 against 512.5 (horde, +2.3%). What fell is the published HPS (the unguarded figure), by 3.8% and 5.7%, because the guard pays about two percent of score for every one percent of mana time, so a set that trades healing for mana wins. The ranker did what its objective says. That the objective can lower the headline number is a decision for the owner (section 6).
2. **Noise did inflate the margin.** The guard multiplies healing by a function of the mean time to empty, and that time varies about 17 s per iteration against about 1% for healing. `healEngine` returned the healing error scaled by the guard and left out the mana-lasts error, so a 300-iteration guarded score reported +/-0.7 when its real error was about 3.7 to 4.5 (0.7 to 0.9% of the score), of the order of the 1% adoption margin. The ranker saw horde 532.3 for the before pick (true 524.5) and 511.9 for the old set (true 512.5): a 4% lead that is 2.3% in truth. The ordering held here; the size of the win and every smaller decision rested on noise the error bar did not show.
3. **No incumbent check.** Nothing compared the new set with the set already published, so any chain of locally better picks, or a noisy one, could end below it.

## 3. Fix

- `inproc.HealingResult.ManaLastsError` (standard error of the time to empty) and `guardedError` (score_heal.go): the delta-method error of the guarded score, the healing's error through the guard plus the mana-lasts error through the guard's slope (independence assumed, covariance left out). Flat above the fight length.
- `healIterationScale = 7`: a healer's ranking runs (screens, tournaments, verification, set completion, incumbent check) use 7 times the requested iterations, 300 becomes 2100, 100 becomes 700. The guarded error falls to about 0.3% of the score, below half the 1% margin. Published metrics run 2000 iterations (was 1000).
- incumbent.go, `keepIncumbent`: after the last pass, a healer band reads the previous published set for its band, preset and faction from the committed report (read before this run can overwrite it), simulates it and the new set under the final harness (verification seed 7), and keeps the incumbent when it beats the new set by more than the two errors combined. An incumbent the candidate pools cannot reproduce, or one identical to the new set, costs no run. A kept band publishes `kept_incumbent` ({incumbent_score, incumbent_error, new_score, new_error, differing_slots}) and its changed slots are sim-decided at the incumbent's score. Healers only; the damage and tank paths are untouched.
- Tests (incumbent_test.go, synthetic): kept when ahead beyond error, not kept inside error or when worse, no run when identical or unreproducible, a slot the incumbent left empty, report loading, the error formula, and the guarded error under the adoption margin at the healer harness while 300 unscaled iterations is not.

## 4. Results, band 60, all five healers

Cells are published HPS / mana_lasts_sec. Committed is the file in the repo (engine 35be2e187), before is the ranker at main, after is this branch; before and after are the same engine, data and seeds. Under the final harness (2100-iteration guarded score, same seed) the after set is at or above the committed set in every row beyond error except two bare paladin rows, where the gap is inside the combined error (alliance 10.94 +/- 0.05 against 10.99; horde 10.57 against 10.59).

| Healer | Preset | Faction | Committed | Before | After | Guarded, committed set to after set (final harness) | Kept |
|---|---|---|---|---|---|---|---|
| druid-restoration | raid | alliance | 600.0 / 271 | 598.5 / 281 | 613.5 / 283 | 523.2 to 548.0 | |
| druid-restoration | raid | horde | 596.5 / 269 | 583.3 / 285 | 611.7 / 282 | 512.6 to 542.5 | |
| druid-restoration | bare | alliance | 324.7 / 228 | 324.7 / 228 | 295.7 / 247 | 186.1 to 201.2 | |
| druid-restoration | bare | horde | 297.7 / 241 | 297.7 / 241 | 302.1 / 241 | 191.6 to 194.6 | |
| paladin-holy | raid | alliance | 510.0 / 239 | 518.0 / 249 | 521.4 / 250 | 350.8 to 362.5 | |
| paladin-holy | raid | horde | 508.8 / 236 | 516.5 / 249 | 516.6 / 249 | 342.2 to 354.8 | |
| paladin-holy | bare | alliance | 175.3 / 75 | 173.1 / 76 | 173.1 / 75 | 10.99 to 10.94 (inside error) | |
| paladin-holy | bare | horde | 173.6 / 74 | 173.6 / 74 | 176.7 / 73 | 10.59 to 10.57 (inside error) | |
| priest-discipline | raid | alliance | 583.9 / 188 | 628.3 / 209 | 628.4 / 209 | 282.6 to 305.9 | |
| priest-discipline | raid | horde | 580.3 / 184 | 622.6 / 209 | 622.3 / 208 | 266.2 to 300.0 | |
| priest-discipline | bare | alliance | 373.4 / 101 | 373.4 / 101 | 373.5 / 101 | 42.3 to 42.3 | |
| priest-discipline | bare | horde | 367.7 / 100 | 367.7 / 100 | 367.8 / 100 | 41.1 to 41.1 | kept |
| priest-holy | raid | alliance | 659.2 / 231 | 707.6 / 257 | 707.2 / 257 | 480.5 to 519.1 | |
| priest-holy | raid | horde | 650.7 / 227 | 682.9 / 244 | 683.3 / 244 | 446.9 to 451.9 | |
| priest-holy | bare | alliance | 331.2 / 96 | 331.2 / 96 | 330.9 / 96 | 34.0 to 34.0 | kept |
| priest-holy | bare | horde | 323.2 / 92 | 323.2 / 92 | 323.2 / 92 | 30.3 to 30.3 | |
| shaman-restoration | raid | alliance | 568.4 / 291 | 576.2 / 333 | 582.6 / 356 | 582.6 to 582.6 | kept |
| shaman-restoration | raid | horde | 567.8 / 295 | 573.5 / 313 | 582.3 / 358 | 582.4 to 582.4 | kept |
| shaman-restoration | bare | alliance | 211.7 / 79 | 211.7 / 79 | 211.4 / 78 | 14.4 to 14.4 | |
| shaman-restoration | bare | horde | 211.2 / 78 | 211.2 / 78 | 210.9 / 78 | 14.2 to 14.2 | |

The incumbent check fired on four rows. On shaman-restoration raid, both factions, the before ranker had published a set the sim measured 5 to 6 guarded points below the committed one (577.6 and 576.5 against 582.6 and 582.4); that is the regression class the lane exists for, caught.

Published HPS is below the committed figure in six rows: druid bare alliance 324.7 to 295.7 (guarded 186.1 to 201.2), paladin bare alliance 175.3 to 173.1 (-1.3%, guarded inside error), shaman bare both factions and priest-holy bare alliance (-0.1 to -0.3%, equal sets or one-slot differences inside error), and priest-discipline raid horde 580.3 is above, not below. Only the druid row is a guard trade beyond error. Restoration druid raid, the row that started the lane, now publishes 613.5 and 611.7 against 600.0 and 596.5 committed, but on the new engine the old set is 622.3 and 619.2 effective HPS: the new pick still has less raw HPS than the old set and more guarded score.

## 5. Spot check, tanks and DPS

druid-feral-bear, paladin-protection, warrior-protection, hunter-marksmanship, mage-fire and rogue-combat, band 60, binary at main against this branch, same engine and data: all four bands of every spec (bare 60, raid 60, both factions) publish the identical set_dps and the identical item in every slot. The damage and tank path is unchanged by construction; the only shared edit is the `kept_incumbent` field, omitted when absent.

## 6. For the owner

The guard decides that druid-restoration bare alliance ships 295.7 HPS over a set that heals 324.7, because it lasts 247 s of mana against 228 s. That is the guard as designed (squared share of the fight), not a bug, but the headline number falls. If the site should never publish a lower HPS than the set before, the objective needs a floor on raw HPS or a milder guard, a separate decision from this lane.
