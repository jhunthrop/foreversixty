# Hunter builds, 2026-10-07

Source: the talent-search reports re-run on engine d1af3529f (client 1.60.1.70009), read from the main checkout. Rule: a winner replaces a guide build only if it keeps every unmodeled guide talent, holds the spec's own tree as the majority, gains at least 1% beyond the combined error, and drops nothing the rotation casts. Ladder movement is the level-60 row of the regenerated golden (relative, same seed).

## Beast Mastery: no change

The guide (`5420001505001251/0053502001/4`) is the best build found; every alternative is 4% or more behind. Golden untouched, level-60 row identical.

## Marksmanship: adopted

- Old: `FS1:1.60.1.70009:hunter:dwarf:55200004/005055001150305/5:`
- New: `FS1:1.60.1.70009:hunter:dwarf:5320000501/005155000150305/5:`
- Change: 2 Endurance Training and 1 Lone Wolf became +1 Unleashed Fury (5/5), +1 Ferocity, +1 Improved Stings.
- Rule check:
  - Unmodeled talents: the engine gaps are Hawk Eye, Improved Concussive Shot and Scatter Shot. The guide takes none of them, so it keeps all of them. The report lists no dropped unmodeled talent. Endurance Training and Lone Wolf are modeled and measure no damage (threat/utility class), so the rule lets them move.
  - Own tree: 16/30/5, Marksmanship holds the majority.
  - Gain: +3.6 DPS (+1.5%) with combined error 0.5, so 3.1 beyond error, about 1.3% of the guide's DPS. Passes the 1% bar, narrowly.
  - Rotation: the level-60 top casts and the violations list are unchanged (shoot, Multi-Shot, Aimed Shot, Serpent Sting, move); no new zero_casts row. The one violation in the file (Arcane Shot at level 50) is present before and after.
- Ladder level 60: +1.5% (221.2 to 224.6).
- Not taken: the deep Beast Mastery variants (+3.1 DPS, finalists 3 and 4) are 30/16/5, so Marksmanship is no longer the spec's own majority tree, and they are no better than the +3.6 winner.

## Survival: adopted, but a smaller variant than the report's best

- Old: `FS1:1.60.1.70009:hunter:dwarf:0/32005500005/500230131051120151:`
- New: `FS1:1.60.1.70009:hunter:dwarf:0/32500500005/500230131051120151:`
- Change: 5 Efficiency became 5 Lethal Attacks. Efficiency cuts shot mana cost; the melee rotation casts no shots.
- Rejected, finalist 8 (1 Counterattack to Resourcefulness, +25.1 DPS, about +9.9%). It passes the report's four conditions as the report evaluates them, but `web/src/content/guides/builds.test.ts` ("takes its rotation-required signature talents") fails: the curated rotation casts Counterattack, so the build must take it. The ladder cannot show this, because the harness target never attacks and Counterattack is in `expected_idle`, so the golden alone would have hidden the drop. I did not loosen that test. Resourcefulness is worth about +36 DPS per point in the probe, so the engine lane or owner may want to revisit the signature-talent requirement; if the test is relaxed, the swap is a one-digit change (`...131050220151`, 0/32005500005).
- Rejected, finalists 1 to 7 (deep Survival, deep Marksmanship): each drops Hawk Eye or Deterrence, or leaves the own-tree majority.
- Taken: finalist 9, modeled non-damage points re-spent (5 Efficiency to 5 Lethal Attacks), +12.6 DPS (+5.0%) at combined error 0.7.
- Rule check:
  - Unmodeled talents: Deterrence, Hawk Eye, Improved Concussive Shot are all kept (the report's "drops unmodeled" column is empty for this variant).
  - Own tree: 0/20/31, Survival holds the majority.
  - Gain: 12.6 beyond error 0.7, about 4.7% beyond error.
  - Rotation: level-60 top casts identical (Raptor Strike/auto attacks, Mongoose Bite 14271, Strider Kick 1317257, Immolation Trap 14266); Violations: None; Counterattack stays in expected_idle as before. Counterattack is still taken.
- Ladder level 60: +5.7% (263.3 to 278.2).
- Guide prose: the Talents section gains one sentence on the Efficiency to Lethal Attacks swap, relative gain only.

## Verification

- `go test ./...` under sim/: pass. No smoke expectation changed (Survival's Counterattack warning is for the bare smoke build and is unchanged).
- `npx vitest run src/content src/lib/guides`: 14 files pass, including the build-legality and signature-talent tests.
- Goldens regenerated for marksmanship and survival only; Beast Mastery regenerates identically.
