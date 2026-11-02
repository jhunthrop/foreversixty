# Sim audit round 7: ranker (2026-11-02)

This lane owns `sim/cmd/leveling-bis` for one bug, which the hybrids lane routed here (hybrids X3). It did no per-spec audit and made no engine change.

**Setup:**
- Site: `wt/r7-ranker`, started from the integration branch `wt/sim-round7` (20c3612a7). Engine: the integration fork worktree `scratchpad/r7/int/engine` (afbbba6b2), read-only.
- Local simdb: the committed `data/builds/1.60.1.70707/simdb.bin` with S 1bd3920b5's held-item typing applied to it. The raw CSVs are not in the worktree, so `python -m pipeline simdb` cannot run there. The hybrids lane's `patchheld.py` types the 178 held-in-off-hand items as the engine's off hand. Its output is byte-identical to the hybrids lane's `simdb-held.bin`. Without it, a held item still replaces the main hand in a sim, and an orb cannot be measured fairly (three different maces beside an orb all measured 229.6).
- Runs: `leveling-bis -spec <spec> -bands 60` for all 28 specs, before (the parent of S 5d34ae14e) and after, on the same simdb. Output is in `scratchpad/r7/ranker/{before,after}/` and is not committed (owner ruling: rank locally only).

**In short: 1 CONFIRMED-BUG, fixed.** It moves the band-60 sets of Balance, Shadow, Elemental, Affliction, Demonology and Feral bear.

## Ranker

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| K1 | **A one-hander that replaced a two-hander left the off hand empty, and it was measured one hand short.**<br>- pick() leaves the off hand empty under a two-handed staff. The effect tournament (`rank.go rankSlotWithEffects`) and the verify swap (`verify.go applySwaps`) can then put a one-hander in the main hand. Neither one refilled the off hand. The swap only ever cleared it, for a two-hander.<br>- Both trials also measured the one-hander with nothing in the off hand, so a one-hander had to beat a staff on its own to win.<br>- Result: Balance, Shadow and Elemental published band 60 as Hand of Edward the Odd with no orb or tome. The report said `empty_reason: no_sourced_item`, although 61 to 65 held items were in the pool.<br>- Fix: pick()'s off-hand list is now `offHandCandidates`, shared with a new `fillEmptyOffHand`. That list is shields and held items, or one-hand weapons for dual-wielders; shamans never dual wield. A main-hand trial wears the off hand the candidate would get (`trialGear`). The promoted pick gets that off hand too, and so do the post-effect picks (`main.go`) and set completion (`sets.go`). The final verify DPS is measured with it. | `verify.go:399-408` (only a two-hander touched off_hand), `rank.go` trial via `swapSlot`, `pick.go` off_hand case. Tests first, in `offhand_fill_test.go`: `TestApplySwapsFillsTheOffHandWhenAOneHanderReplacesATwoHander`, `TestVerifyBandMeasuresAOneHandRunnerUpWithItsOffHand`, `TestRankSlotWithEffectsMeasuresAOneHanderWithItsOffHand`. They failed to build, then with the verify and rank changes reverted they failed on behaviour, then they passed. Shaman, dual-wield and no-change cases are covered as well. `go test ./cmd/leveling-bis` passes. | CONFIRMED-BUG (local before/after runs below) | S 5d34ae14e |

## Before/after (band 60, local ranker, same simdb)

These are the ranker's own set DPS at 300 to guard iterations, seed 7. 22 of the 28 specs are unchanged and bit-identical. Before this fix, no spec had a one-hander beside an empty off hand except Balance, Shadow and Elemental. After it, none do.

| spec | preset | faction | before | after | change | picks |
|---|---|---|---|---|---|---|
| druid-balance | bare | alliance | 270.7 | 279.7 | +3.31% | Edward + Spirit of Aquementas (was empty) |
| druid-balance | bare | horde | 266.1 | 275.0 | +3.31% | Edward + Spirit of Aquementas |
| druid-balance | raid | alliance | 572.0 | 585.5 | +2.36% | Edward + Spirit of Aquementas |
| druid-balance | raid | horde | 567.4 | 576.9 | +1.68% | Edward + Spirit of Aquementas; shoulder, chest and trinkets re-picked |
| priest-shadow | bare | alliance | 360.4 | 367.4 | +1.94% | Edward + Mordresh's Lifeless Skull |
| priest-shadow | bare | horde | 356.9 | 363.6 | +1.89% | Edward + Spirit of Aquementas |
| priest-shadow | raid | alliance | 712.8 | 723.1 | +1.44% | Edward + Spirit of Aquementas |
| priest-shadow | raid | horde | 712.6 | 722.8 | +1.43% | Edward + Spirit of Aquementas |
| shaman-elemental | bare | alliance | 147.3 | 152.0 | +3.14% | Edward + Garrett Family Crest |
| shaman-elemental | bare | horde | 145.7 | 149.2 | +2.41% | Edward + Garrett Family Crest |
| shaman-elemental | raid | alliance | 473.4 | 483.8 | +2.21% | Edward + Spirit of Aquementas |
| shaman-elemental | raid | horde | 473.5 | 484.1 | +2.24% | Edward + Spirit of Aquementas |
| warlock-affliction | raid | alliance | 577.4 | 588.9 | +1.99% | Kindling Stave -> Inventor's Focal Sword + Spirit of Aquementas; shoulder |
| warlock-affliction | raid | horde | 575.0 | 585.1 | +1.76% | Kindling Stave -> Inventor's Focal Sword + Spirit of Aquementas |
| warlock-demonology | bare | alliance | 326.4 | 333.5 | +2.17% | Spire of Hakkar -> Inventor's Focal Sword + Spirit of Aquementas; head |
| warlock-demonology | bare | horde | 325.6 | 333.2 | +2.34% | Kindling Stave -> Inventor's Focal Sword + Spirit of Aquementas; wrist, waist |
| warlock-demonology | raid | alliance | 646.7 | 654.2 | +1.16% | Kindling Stave -> Inventor's Focal Sword + Spirit of Aquementas |
| warlock-demonology | raid | horde | 644.0 | 651.7 | +1.20% | Kindling Stave -> Inventor's Focal Sword + Spirit of Aquementas |
| druid-feral-bear | raid | alliance | 316.2 | 308.3 | -2.49% DPS | Twig of the World Tree -> Greenhammer + Furbolg Medicine Pouch |
| druid-feral-bear | raid | horde | 314.8 | 307.3 | -2.38% DPS | same |

- The warlock moves come from fairer measurement: a one-hander is now tried with its off hand, so it can beat a staff.
- Feral bear ranks on its tank score, not DPS. That score went from 30236.8 to 30440.7 (+0.67%) on Alliance, and from 31809.6 to 31761.0 on Horde, which is within noise.
- In the six changed specs, the incumbent guard kept no previous published set: every after set beat the committed one.

## Still open

- These picks depend on S 1bd3920b5's held-item typing reaching the committed `simdb.bin` at the next data run. Until then, a nightly re-rank would wear an orb in place of the main hand. The re-rank of every caster and Feral bear should follow that data run.
- Ranker output is not committed (owner ruling). The published `data/builds/1.60.1.70707/bis/` still shows the empty off hands until the bis run.
