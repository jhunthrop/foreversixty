# Warrior builds, 2026-10-07

Source: the talent-search reports regenerated today on engine d1af3529f (read from the main checkout; the copies in this branch's base are the stale pre-sweep versions). Adoption rule: keep every talent the engine does not model, hold the spec's own tree as the majority, gain at least 1% beyond the combined error, drop nothing the rotation casts.

## Arms (adopted)

- Old: `FS1:1.60.1.70009:warrior:human:05325213032310001/0505/055:` (31/10/10)
- New: `FS1:1.60.1.70009:warrior:human:03325213032515001/0505/005:` (36/10/5)
- Variant: "guide, modeled non-damage points re-spent + 2 Deflection -> Weaponmaster" (finalist 3). Added +2 Bloodthrill, +5 Weaponmaster; removed -2 Deflection, -5 Shield Specialization.
- Rule check:
  1. Unmodeled talents kept: the report's "Drops unmodeled guide talents" column is "-". The engine-gap list is Improved Hamstring and Iron Will; Iron Will stays at 5. Deflection stays at 3 of 5 and it is now classified modeled, no damage, so moving two points is allowed by the "modeled non-damage points re-spent" clause. Shield Specialization is modeled, no damage.
  2. Own tree majority: 36/10/5.
  3. Gain: +25.0 DPS (+10.5%) against a combined error of 0.9, so beyond error by far more than 1%.
  4. Rotation casts: level-60 distinct casts are 6 before and after; Unresolved column is "-" at level 60; Mortal Strike (21553) is still the top spell cast; the ladder Zero-casts check reports no violation. Heroic Strike is not cast by the current Arms list and is not in the build either. No new level-60 zero cast.
- Search gain: about +10.5% (guide 238.8, winner 263.8, both in the report).
- Ladder level 60: about +12.3% (golden DPS 187.6 to 210.7).
- Ladder note, not a level-60 issue: at level 40 the 9-points-per-level walk now reaches Mortal Strike later (level 50 instead of 40), so the ladder lists Mortal Strike as an expected unresolved line at level 40 and level-40 DPS falls from 88.0 to 78.6. The ladder test passes (the rule treats a not-yet-reached row as expected). Leveling order is a follow-up for the leveling guide, not a reason to hold the level-60 build.
- Rejected: finalists 1, 2 (+25.1 to +25.2) and 5 drop Iron Will (unmodeled), finalist 4 drops two Iron Will. The strongest report-wide variants are all within error of finalist 3 (+25.0), so nothing is lost by keeping Iron Will.

## Fury (adopted)

- Old: `FS1:1.60.1.70009:warrior:human:35311103002/35051105050010501/0:` (19/32/0)
- New: `FS1:1.60.1.70009:warrior:human:34320003002/05153105022011501/2:` (17/32/2)
- Variant: "deep Fury" (finalist 1). Added +1 Improved Charge, +1 Lingering Rage, +2 Furious Precision, +2 Improved Execute, +1 Improved Intercept, +2 Improved Bloodrage; removed -1 Deflection, -1 Improved Tactical Mastery, -1 Improved Overpower, -3 Booming Voice, -3 Enrage.
- Rule check:
  1. Unmodeled talents kept: "Drops unmodeled guide talents" is "-". The engine-gap list is Lingering Rage, Improved Intercept and Gore Drinker; the build takes Lingering Rage and Improved Intercept, none removed. Booming Voice, Enrage, Deflection, Improved Tactical Mastery and Improved Overpower that are removed are all modeled in this run (the report lists them as non-gaps).
  2. Own tree majority: Fury 32 of 51.
  3. Gain: +23.2 DPS (+8.6%) against a combined error of 0.9.
  4. Rotation casts: level-60 distinct casts are 9 before and after, Unresolved "-". Bloodthirst, Death Wish and Flurry are in the build, and Whirlwind, Heroic Strike and Execute are untouched (Execute is cheaper by 5 Rage). No new zero cast. The data confirms Death Wish, Flurry and Bloodthirst all have non-zero ranks in the new string.
- Search gain: about +8.6%.
- Ladder level 60: about +9.5% (golden DPS 217.7 to 238.3).
- Rejected: nothing needed rejecting; the "keeps every unmodeled talent" variant and the overall winner are the same build. Next-best finalists (guide with modeled non-damage points re-spent, +20.1) were not needed.

## Verification

- `FOREVER_UPDATE_GOLDEN=1 go test ./request/ -run TestRotationLadder` regenerated warrior-arms and warrior-fury goldens; `go test ./request/` passes. Smoke expectations in `sim/request/rotations_smoke_test.go` need no change: the warrior-arms entry is for a zero-talent run and warrior-fury reads the engine's own reference talents.
- `go test ./...` under sim/ passes (one first-run failure in cmd/leveling-bis did not reproduce on rerun; unrelated to warriors).
- `npx vitest run src/content src/lib/guides`: 14 files, 341 tests passed.
- Protection is untouched (no search winner requested).
