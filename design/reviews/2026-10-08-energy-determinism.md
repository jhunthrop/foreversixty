# Energy determinism across architectures (2026-10-08)

## Symptom

After the continuous-energy lane, `rogue-combat.golden.md` passed locally (arm64) under
`go test -race -count=8` and failed on CI (x64). Go may fuse `x*y + z` into one fused
multiply-add on arm64, never on amd64. A one-bit difference in an energy value flips an
`energy >= cost` comparison once per fight.

## Fusable expressions on the energy path

The Go spec forbids fusing across an explicit `float64(...)` conversion and across an
assignment. Audit of `sim/core/energy.go` and its feeders in the fork:

| Site | Expression | Fusable? |
| --- | --- | --- |
| `accrue` | `amount := rate*elapsed.Seconds()` then `currentEnergy+amount` | No. The product is assigned first. |
| `computeWakeAt` | `(threshold-energy)/rate`, `Ceil(seconds*float64(time.Second))` | No. Division, and a product with no sum. |
| `EnergyForTime`, `TimeForEnergy` | one product, one division | No. But callers add the result (`curEnergy + EnergyForTime(..)`), and an inlined return value is fusable. |
| `ApplyCostModifiers` | `cost*mod/100` | No. Product then division. |
| `druid/furor.go` `furorCatFormEnergy` | `a*r*last/100 + b*r*secs` | Yes. A product feeds the sum. It sets the energy on entering cat form. |
| `rogue/expose_armor.go` cost | `25 - 5*float64(rank)` | Yes (`c - a*b`). Values are whole numbers, so both results agree today, but it is a latent hazard. |
| `druid/ferocious_bite.go` damage | `... + damagePerEnergy*excessEnergy` | Yes, but it is damage, not a threshold. Left alone: the whole engine's damage maths has the same shape and the goldens print one decimal. |

Honest finding: the accrual in `energy.go` was not itself fusable as written, so the
diagnosis is not confirmed at that line. Rogue has no `furorCatFormEnergy`. The cross-arch
difference may come from another site in the rogue path that this audit did not reach (for
example a damage sum). The architecture-independent accrual removes energy as a suspect;
CI is the only proof that the rogue-combat difference is gone. If it still fails there, the
cause is outside the energy path and the next step is dumping the fight log on x64.

## Design

The bar stores `currentUnits`, an `int64` count of 1e-8 Energy.

- 10 Energy per second is exactly 1 unit per nanosecond, so accrual is
  `elapsed_ns * multiplier` with no rounding. Adrenaline Rush is multiplier 2.
- Spend and add convert the amount once with `Round(amount * 1e8)`; the product stands
  alone, so it cannot fuse. The fractional carry on spend is kept exactly.
- The float value is derived by one division, `units / 1e8`. For whole and one-decimal
  amounts this equals the float literal, so `energy >= cost` compares like for like.
- The wake time is integer: `ceil(missing_units / rate)`. It lands exactly on the
  threshold, so the old `+1 ns` safety is gone.
- `EnergyForTime` and `TimeForEnergy` use the same units, exact inverses at ms resolution.
- `AddEnergyRegenMultiplier` takes whole steps (`int64`) and rejects a negative multiplier.
- `APLValueEnergyThreshold` read the stored field without accruing; it now calls
  `CurrentEnergy()`.
- `furorCatFormEnergy` and the Improved Expose Armor cost wrap each product in `float64()`.

Units of 1e-8 rather than milli-energy: milli-energy would round the accrual of a 1 ns
step and the wake time. Nanosecond units keep the existing behaviour (10 per second mean,
cap, doubling, carry on spend, wake timings) bit for bit apart from the rounding below.

Other resources: the energy lane did not touch rage or focus, so they are unchanged.

## Tests

Fork `sim/core/energy_test.go`: an accrual sequence pinned with `==`
(0.35 s gives 3.5, spend 3, 0.15 s gives 2.0, then doubled regen and a spend of 42.5), a
10,000-step drift test, `EnergyForTime`/`TimeForEnergy` inverses over 20,000 ms and 20,000
hundredths, and a check that the unit constants divide exactly.

## Golden movement

Regenerated with `FOREVER_UPDATE_GOLDEN=1 go test ./request/ -run TestRotationLadder`.
Feral, druid and all other goldens did not move. Rogue moved by rounding only (DPS, per
level band, old to new):

| Spec | 30 | 38 | 40 | 50 | 60 |
| --- | --- | --- | --- | --- | --- |
| Combat | 47.8 same | 68.9 to 68.8 | 80.1 to 80.0 | 114.6 to 114.5 | 229.5 same |
| Assassination | n/a | 68.7 same | 85.2 same | 129.0 to 129.2 | 208.2 to 208.3 |
| Subtlety | 38.1 same | 62.2 same | 71.8 to 71.9 | 102.1 to 102.2 | n/a |

Largest move 0.2 DPS (0.15 percent). Ability shares moved by at most 0.3 points
(for example Combat 30 Sinister Strike 45.9 to 45.6).

## Verification

- Fork: `go test --tags=with_db ./sim/core/ ./sim/rogue/... ./sim/druid/...` passes.
- amd64 cannot run here (no Rosetta). `GOARCH=amd64 go test -c` compiles the rogue, feral
  and site request test binaries. CI on x64 is the cross-architecture proof.
