# Rogue builds, 2026-10-07

Source: the talent-search reports regenerated today (engine d1af3529f, client 1.60.1.70009, seed 7, 800-iteration finalists), read from the main checkout's working tree (`design/reviews/talent-search/rogue-*.md`). Adoption rule: keep every unmodeled guide talent, hold the spec's own tree as the majority, gain at least 1% beyond the combined error, drop nothing the rotation casts. Ladder movement is the level-60 row of `sim/request/testdata/ladder/rogue-*.golden.md` (leveling gear, 300 iterations), relative.

## Assassination: adopted (a variant constructed from the report's rule-abiding row)

- Old: `FS1:1.60.1.70009:rogue:night-elf:32500000551501051/3252/51:` (33/12/6)
- New: `FS1:1.60.1.70009:rogue:night-elf:32502110551501001/302303/512:` (32/11/8)
- Taken: +2 Murder, +1 Improved Slice and Dice, +1 Relentless Strikes, +1 Puncturing Wounds, +3 Precision, +2 Opportunity. Dropped: -5 Seal Fate, -2 Improved Sinister Strike, -3 Lightning Reflexes.
- Search gain: guide 215.7, new 257.4, +41.7 (+19.3%), error 0.4 each, combined about 0.6.
- Ladder level 60: +19.9% (170.6 to 204.6).

Why not the report's own rule-abiding winner ("guide, modeled non-damage points re-spent", +46.2, +21.4%): it also drops **Venom** (-1 Venom to pay for Relentless Strikes). The engine scores Venom at about -2 DPS per point, but the rotation casts Venom (spell 1310703), so the variant drops something the rotation needs and is rejected. The report's other finalists with a larger gain all drop unmodeled guide talents (Improved Gouge, Remorseless Attacks, Camouflage, Master of Deception) and fail the rule.

Because the report's finalist list has no Venom-keeping variant of that size, I built one with the report's own harness (throwaway test in `sim/cmd/talent-search`, not committed): the same re-spend, minus the Venom removal, with one point of Improved Slice and Dice or Opportunity given up to pay for Venom. Variants simmed at 800 iterations, seed 7: ISAD 2 / Opp 1 +18.7%; ISAD 1 / Opp 2 +19.3% (chosen); +1 Ruthlessness in place of ISAD +19.2%; Ruthlessness 2 / Opp 1 +18.2%; Ruthlessness 3 +17.8%. If the owner prefers a build taken only from the report's tables, the next admissible row there is the screened "guide: 1 Seal Fate -> Relentless Strikes" (+13.0 at 300 iterations, +6%).

Rule check:
1. Unmodeled talents kept: Improved Gouge 3, Remorseless Attacks 2, Camouflage 5, Master of Deception 1 all unchanged. Improved Kidney Shot is not in the guide. Pass.
2. Own tree majority: 32 of 51 points in Assassination. Pass.
3. Gain: +41.7 against a combined error of about 0.6, 19.3% against the 1% bar. Pass.
4. Rotation casts: ladder level-60 row unused column is `-` before and after; the violations section is `None.` before and after; Mutilate (1241584) and Venom (1310703) both appear in the new row's top casts. Cold Blood is kept (one point), Eviscerate and Slice and Dice are untouched. Pass.

## Combat: adopted

- Old: `FS1:1.60.1.70009:rogue:night-elf:32531/32531300000515201/51:` (14/31/6)
- New: `FS1:1.60.1.70009:rogue:night-elf:32531/32530300001515201/51:`, one point moved, Deflection to Flawless Execution (14/31/6).
- Search gain: guide 224.9, new 230.5, +5.6 (+2.5%), combined error 0.6.
- Ladder level 60: +2.0% (206.2 to 210.4).

Rule check:
1. Unmodeled talents kept: the report's "Drops unmodeled guide talents" column is `-` (Deflection is modeled, zero damage). Pass.
2. Own tree majority: 31 of 51 in Combat. Pass.
3. +5.6 against error 0.6, 2.5% against 1%. Pass.
4. Rotation casts: Deflection is passive; the level-60 unused column is `-` before and after, the only violations are the pre-existing Backstab zero_casts at every level (unchanged). Pass.

Rejected: the report's finalists 1 to 6 (deep Combat, +18.1%) drop Improved Gouge, Remorseless Attacks, Camouflage and Master of Deception, which the engine cannot price; finalists 7 to 10 (+5.2%) each drop one such talent.

## Subtlety: adopted

- Old: `FS1:1.60.1.70009:rogue:night-elf:005/325131/5322210310013011051:` (5/15/31)
- New: `FS1:1.60.1.70009:rogue:night-elf:005323101014/0/5323220310013011031:` (20/0/31)
- Taken: +3 Ruthlessness, +2 Murder, +3 Improved Slice and Dice, +1 Relentless Strikes, +1 Lethality, +1 Cold Blood, +4 Improved Poisons, +1 Setup, +1 Dirty Tricks. Dropped: all 15 Combat points (Improved Eviscerate, Improved Sinister Strike, Lightning Reflexes, Puncturing Wounds, Deflection, Precision) and -2 Cutthroat.
- Search gain: guide 187.8, new 222.3, +34.5 (+18.4%), combined error 0.5.
- Ladder level 60: +18.2% (143.5 to 169.6).

Rule check:
1. Unmodeled talents kept: the report lists no dropped unmodeled guide talent. Camouflage 5, Master of Deception 3, Dirty Tricks (rises to 2), Opportunity 2 and the rest are all kept. Pass.
2. Own tree majority: 31 of 51 in Subtlety. Pass.
3. +34.5 against 0.5, 18.4% against 1%. Pass.
4. Rotation casts: Ghostly Strike, Premeditation, Hemorrhage, Preparation and Thousand Cuts all kept. Ladder level-60 row: Hemorrhage (16511) and Ghostly Strike (14278) still the top talent casts, unused column `-`; violations unchanged (the pre-existing Eviscerate zero_casts at levels 20 and 30). The new level-30 row lists Premeditation (14183) as learned but unused, which is informational and limited to the level-30 point budget (the old level-30 row carried the same kind of note for Hemorrhage); both builds keep it in the level-60 build. Pass.

Rejected: deep Combat (+18.2%) drops Camouflage, Dirty Tricks and Master of Deception, and deep Assassination (+10.7%) the same; both fail the unmodeled rule.

## Checks run

`go test ./...` under `sim/` passes; `cd web && npx vitest run src/content src/lib/guides` passes (14 files). The smoke exclusions in `sim/request/rotations_smoke_test.go` describe the bare no-talent build, so they needed no edit. Guide prose quotes relative gains only.
