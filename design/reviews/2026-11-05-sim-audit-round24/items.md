# Sim-accuracy round 24: items lane

Site branch `wt/r24-items` from int/r24 (0eb403494). No engine change (fork `lane/r24-items` = int/r24, 3b4d6b054). Build mode. Scope: the coordinator's priority seed (ranker wand credit, from the BiS lane), then the other caster raid sets.

## Cross-cutting: the BiS ranker's caster ranged slot

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Alliance Destruction raid 60 wears Torch of Light (6 spell power, 11 Spirit, 17 healing) over Brilliant Wand (9 spell power, 7 Intellect).** It is not wand damage: a raid warlock never shoots. The ranker's own harness (plainRequest, raid preset, seed 7) measures Shoot at 0 DPS on the published set with either wand, so the measured wand credit is 0 and both wands score on stats alone. Torch scores first only through that band's Spirit weight, 0.394 +/- 0.216, which the band itself marks insignificant (the horde run of the same build measured 0.139 +/- 0.253; 11 Spirit is worth well under 1 DPS to a warlock who is always inside the five-second rule). The wand tournament (`wand_credit.go` `rankWandSlot`) then kept the higher of two 300-iteration runs: Torch by 0.07 DPS (published `dps_delta -0.071`) against errors of about 2. Nothing after it can move a wand by less than 1% (`finalWandCheck` needs `swapMargin`; verify only re-tests the runner-up), so the coin flip published. Fix: each wand is measured with its error, and the leader plus every wand inside the combined error of it are measured again at the guard harness (`guardIterations`, 2,100), whose order decides. `finalWandCheck` now compares the best wand with the worn wand inside the same tournament, not with a set figure from another harness. Tests (written first; the first failed with `ranged = 279246`): `TestTheWandTournamentSettlesALeadInsideTheErrorAtTheGuardHarness` (the Destruction raid ranged slot: Torch, Brilliant, Greater Eternal at their scores; verify-harness tie; guard harness puts Brilliant first), `TestTheWandTournamentRunsNoGuardForALeadBeyondTheError`, `TestTheFinalWandCheckComparesTheWornWandInTheSameHarness`. Ranker harness on today's engine, published set, seed 7: alliance Torch 731.3 +/- 0.8, Brilliant 733.4 +/- 0.8 at 2,100 iterations. | published `data/builds/1.60.1.70790/bis/warlock-destruction.json` (main c18afbc6e) band 60 raid alliance; item rows `data/builds/1.60.1.70790/items/warlock.json` 279246, 249385; round 6 `ranker.md` W1/W2 | CONFIRMED-BUG | S f34eeb512 |

Why round 6's fix did not cover it: round 6 removed the whole-fight wand credit (W1) and added the guard harness for a runner-up ahead inside the 1% margin (W2). The wand tournament added in round 9 adopts its raw 300-iteration best before verify, and when that best is the noise winner the real winner is the runner-up and is behind, so W2's re-measure never runs. Origin: `old` (round 9 tournament; rounds 6 to 23 did not re-check a tie inside it).

**Cross-class:** S f34eeb512 changes how every wand-wearing caster's ranged slot (mage, priest, warlock, bare and raid, every band) is chosen at the next BiS re-rank. It changes no published file and no sim number now.

## Other caster raid sets holding a wand for the same reason

Published (main c18afbc6e), band 60 raid: only **Destruction alliance** (Torch of Light). Every other mage, Shadow priest and warlock raid set wears Brilliant Wand, and Frost wears Wand of Biting Cold (16 Frost spell power, a real lead). Discipline and Holy wear Torch of Light, which is right for a healer (17 healing, significant Spirit weights 0.44 to 1.1).

The frozen BiS (70755) had two more inside-the-noise picks, both already Brilliant Wand on main, measured here on the frozen set, raid, 10,000 iterations:

| set | worn | Brilliant Wand | seeds 1 / 50001 |
|---|---|---|---|
| warlock-affliction alliance | Greater Eternal Wand (9 spell power, no Intellect) 753.7 / 754.0 | 754.8 / 755.0 | +0.15% / +0.13% |
| mage-fire alliance | Pyric Caduceus (13 Fire spell power, 3 Spirit) 724.4 / 724.9 | 725.0 / 726.1 | +0.08% / +0.17% |
| mage-fire horde | Pyric Caduceus 711.2 / 712.3 | 711.7 / 713.2 | +0.07% / +0.13% |

## Before/after

Destruction raid 60 on the frozen BiS with the guide raidBuild (`25500000033-0-0050355103101351`), 10,000 iterations, the ranged slot swapped from Torch of Light (old) to Brilliant Wand (new):

| faction | seed 1 | seed 50001 |
|---|---|---|
| alliance | 778.9 -> 780.8 (+0.24%) | 778.3 -> 780.6 (+0.30%) |
| horde | 773.9 -> 776.3 (+0.31%) | 774.2 -> 776.8 (+0.34%) |

Errors +/- 0.4 each. Shoot is cast 0 times in every run.

## Files to regenerate

None. The fix lands through the queued CI BiS re-rank.

## Measure at refresh

- After the re-rank: Destruction raid 60 alliance should wear Brilliant Wand (or a wand that beats it at the guard harness). Check every caster ranged slot whose top two wands sit within about 0.5% for a pick that flips between nightlies.
- Ranker time: the guard re-measure runs only when two of the top three wands tie at the verify harness, at most three extra 2,100-iteration sims per wand-wearing band and faction (twice, with the final check).

## For other lanes

- BiS lane: the ranker's score still credits insignificant positive weights (`weights.go` `effectiveWeights`, by design). Spirit on a caster DPS spec is the usual one. The tournament now settles wand ties by measurement; other slots that tie at the verify harness rely on verify's runner-up re-measure (W2), which runs only when the runner-up is ahead.

## For the guide lane

Nothing. No guide, APL or card text changes.

## Still open

- Logs: looked 23:48 UTC 5 Nov. api.foreversixty.gg/v1/reports/recent lists only the 20 Oct test report. This finding needs no log.
