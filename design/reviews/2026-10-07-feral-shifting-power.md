# Feral: Shifting Power in the cat rotation (2026-10-07)

Branches: site `feral-shift`, fork `feral-shift` (both off `main` / `forever`, not merged).

## Model

Shifting Power (Feral node 104951, spell 1322605, hotfix_only; Improved Shifting Power node
113563, spell 1322670, hotfix_only) is registered in `sim/druid/shifting_power.go`: instant, in
cat form, 55% of base mana (reduced 10% a rank by Natural Shapeshifter) for 40 energy, 16 s
cooldown, minus 4/8 s from Improved Shifting Power. That registration was already APL-castable;
this lane added the two rotation paths that use it.

- The curated APL (`data/curated/apl/druid-feral.json`, synced to the fork's
  `ui/feral_druid/apls/forever_feral.apl.json` and `sim/request/apl/druid-feral.apl.json`) casts
  1322605 above Tiger's Fury whenever `currentEnergy <= 60`, so the whole 40 fits under the 100 cap.
  Without the talent the spell is not registered and the line does nothing (guide-build baseline
  is unchanged at 221.2 +- 0.3).
- The engine's hard-coded cat logic (the `catOptimalRotation` action, `sim/druid/feral/rotation.go`)
  powershifted out of cat form whenever it was energy-starved. It now casts Shifting Power
  first when the spell is ready, affordable and the full 40 fits (`shiftingPowerFits`), and falls
  back to the powershift otherwise. The curated rotation does not use that action; it is the
  fork UI preset path.

## Client values

| Value | Source | State |
|---|---|---|
| 40 energy | Wowhead text on the overlaid tree | live text |
| 55% of base mana, 16 s cooldown | lane brief from Blizzard's 1 Oct patch notes (the text prints "0 Mana", an unresolved hotfix cost) | not in the client |
| Improved Shifting Power -4/-8 s | Wowhead text | live text |
| Spell rows 1322605 and 1322670 | `data/builds/1.60.1.70009/spellconst/druid.json` | **absent**: hotfix-only, no client row |
| Global cooldown | no `gcd_ms` to read | **unconfirmed**, treated as on the GCD |

Rip and Ferocious Bite, against the client rows: both spells' `gcd_ms` is 1000 and costs are
30 and 35, matching the engine. Ferocious Bite's effect 1 is a dummy whose amount is 100/150/200/
250/270 for ranks 1-5, which is exactly the engine's damage-per-energy table (1.0/1.5/2.0/2.5/2.7);
this is now pinned by `TestFerociousBiteDamagePerEnergyMatchesClient`. **Neither spell's rows
carry a per-combo-point step** (Rip's second effect is a zero-amount dummy; Bite has no such
effect), so `ripTickPerComboPoint` and `ferociousBiteDamagePerComboPoint` stay the Era figures;
there is nothing to correct them to. Comments in `rip.go` and `ferocious_bite.go` say so.

## Measurements

`rotation-search -spec druid-feral`, level 60, committed alliance band gear, seed 7, 800
iterations. The guide's FS1 build is `0/54232212120032010001/55532`, which spends 1 point in
Shredding Attacks and none in Shifting Power. The search holds the build fixed, and it does not
list Shifting Power as an insertion candidate (its spell list comes from the client table), so
the line was tested by editing the curated file.

To measure the spell, the guide build was temporarily changed (not committed) to
`0/51232232121032010001/55532`: Heart of the Wild 4 to 1, Shredding Attacks 1 to 3, Shifting
Power 1. That build is the measurement bed; the first row is the same rotation without the line.

| Build | Rotation | DPS | +- |
|---|---|---|---|
| guide (FS1) | curated before | 221.2 | 0.3 |
| guide (FS1) | curated with the line (inert) | 221.2 | 0.3 |
| bed | without the line | 232.9 | 0.3 |
| bed | line at energy <= 60 (adopted) | 259.6 | 0.3 |
| bed | line at energy <= 40 | 259.6 | 0.3 |
| bed | line at energy <= 20 | 258.3 | 0.3 |

The line's own contribution on the bed is +26.7 +- 0.4 (+10 to 11%, removal probe). Shred,
Rip and Rake keep casting; Ferocious Bite casts as rarely as before (never in this solo run,
which the curated file already records as expected-idle). The full search on the bed found one
further move, dropping Rake, +10.3 +- 0.6 (270.5): not adopted (see below).

Fork tests (`--tags=with_db`): `sim/druid/...` pass. New: `TestShiftingPowerIsCastableFromTheAPL`,
`TestOptimalCatRotationUsesShiftingPower` (red before the `rotation.go` change),
`TestFerociousBiteDamagePerEnergyMatchesClient`. Site: `go test ./request/` passes after adopting
the druid-feral ladder golden, whose only change is the unresolved-id column now listing 1322605
(zero-talent rows cannot learn it; DPS is unchanged) and a `smokeBuildWarnings` entry for it.
`make apl-sync` / `apl-check` clean.

## What is open

1. **The guide build does not take Shifting Power**, so the adopted line is inert for the
   build the site shows. Taking Shredding Attacks 3 and Shifting Power is worth far more than the
   points they replace on the bed (232.9 vs 221.2 from the Shredding Attacks points alone, then
   the line on top), but the build is the talent-search lane's call: the committed
   `design/reviews/talent-search/druid-feral.md` predates the live tree and needs a re-run with
   this rotation. The guide prose says the build has not taken it yet.
2. Rake: removing it gains 10 DPS on the bed but Rake helps +3.4 on the guide build. It depends
   on the talent shape, so the rotation was not changed; decide it with the build.
3. GCD: if Shifting Power is off the GCD the real gain is slightly higher than modelled.
   The 55% and 16 s are patch-note figures, not client values; a client row would settle both.
4. Improved Shifting Power (8 s cooldown) and mana sustain with it were not measured; at 8 s the
   mana pool may bind.
5. The search tool cannot propose hotfix-only spells; teach it the overlaid tree's spell ids.
6. Operational: `make apl-sync` writes to `ENGINE_DIR` (default the main fork checkout). It
   modified the main checkout's `ui/feral_druid/apls/forever_feral.apl.json` once before being
   rerun with `ENGINE_DIR` set to the worktree; that stray edit was reverted with
   `git checkout --` and the main checkout is clean.
