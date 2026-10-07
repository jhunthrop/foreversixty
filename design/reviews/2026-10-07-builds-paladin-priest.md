# Builds: Retribution Paladin and Shadow Priest, 2026-10-07

Source: the talent-search reports re-run on engine d1af3529f (`design/reviews/talent-search/paladin-retribution.md`, `priest-shadow.md`, read from the main checkout; they are not part of this branch). Adoption rule: keep every unmodeled talent, hold the spec's own tree as the majority, gain at least 1% beyond the combined error, drop nothing the rotation casts. Ladder goldens regenerated with `FOREVER_UPDATE_GOLDEN=1`; `go test ./...` under sim/ and `vitest run src/content src/lib/guides` pass.

## Retribution Paladin: adopted (deep Retribution)

- Old: `FS1:1.60.1.70009:paladin:human:54003/323/0502533100133032:` (12/8/31)
- New: `FS1:1.60.1.70009:paladin:human:52003003/052/05025331001330311:` (13/7/31)
- Changes: +3 Reverence, +3 Redoubt, +1 Twist of Light; -2 Divine Intellect, -3 Toughness, -1 Precision, -1 Instrument of Law.
- Search gain: +14.3 DPS, +6.0%, combined error 0.7.

Rule check:
1. Unmodeled talents kept: the report's "Drops unmodeled guide talents" column is empty and the build has no unmodeled talents. The engine-gap list (Pursuit of Justice, Eye for an Eye, Repentance) holds talents the previous guide build no longer takes. Pass.
2. Spec tree is the majority: Retribution 31 of 51 points. Pass.
3. Gain: 14.3 less 0.7 error is 13.6 DPS, about 5.7% of the guide, above 1%. Pass.
4. Rotation casts everything: the regenerated ladder's `zero_casts` column is `-` at every level, before and after. Level-60 top-five shares move only because Holy Strike (10333, 18.1% before) falls just under Judgement (20271) and no longer shows in the top five; it is not idle. Pass.

Dropped talents are modeled at zero damage (Toughness) or low value (Precision, Divine Intellect, Instrument of Law's second point). The worktree's own copy of the Retribution report was the stale pre-sweep file (it named a different winner); the fresh report was read from the main checkout.

Ladder level 60: 175.3 to 188.0, about +7.2% relative. Level 50 also rises (about +9.5%); levels 10 to 40 are unchanged.

## Shadow Priest: adopted (modeled non-damage points re-spent + 1 Improved Mind Blast -> Mental Agility)

- Old: `FS1:1.60.1.70009:priest:gnome:5241110013/0/543110401201300251:` (18/0/33)
- New: `FS1:1.60.1.70009:priest:gnome:5240110313/0/543120301201300051:` (20/0/31)
- Changes: +3 Mental Agility, +1 Improved Shadow Word: Pain; -1 Silent Resolve, -1 Improved Mind Blast, -2 Early Demise.
- Search gain: +19.4 DPS, +6.9%, combined error 0.8.

Rule check:
1. Unmodeled talents kept: Power in Light 5, Holy Precision 1, Improved Power Word: Shield 1, Blackout 4 all unchanged (the finalist's "Drops unmodeled guide talents" is empty). Pass.
2. Spec tree is the majority: Shadow 31 of 51, still above the 30 Shadowform needs. Pass.
3. Gain: 19.4 less 0.8 error is 18.6 DPS, about 6.6%. Pass.
4. Rotation casts everything: the `zero_casts` column is `-` at level 60 before and after; at level 40 it improves (Shadowform is now reached, leaving only Inner Focus idle, as expected for the truncated build). Mind Blast cadence: Improved Mind Blast is probed at +0.00 DPS per point (modeled, no damage), so the 4 to 3 change does not move the cadence beyond noise; the ladder's level-60 row keeps the same action set and shares (wand 72.2 to 69.3, Mind Flay 25.9 to 27.5, Shadow Word: Pain 2.1 unchanged). Early Demise, the dropped talent, is probed at +0.03 DPS/point and Shadow Word: Death is still cast. Pass.

Ladder level 60: 195.0 to 201.5, about +3.3% relative. Level 40 rises about 6%, level 38 about 4%.

## Rejected

- Retribution "deep Retribution" is the winner here; no variant was rejected on rule grounds. Holy and Protection variants lose 13% to 20%+ and were never candidates.
- Shadow "deep Shadow + 1 Shadow Affinity/Blackout/Spirit Tap -> Mental Strength" (+9.2%) drops Improved Power Word: Shield, Holy Precision and Power in Light points (unmodeled): rejected by rule 1. The "all non-damage points re-spent" family (+8.4% to +9.2%) rejected for the same reason; "modeled non-damage + 3 Power in Light -> Mental Strength" (+7.4%) drops 3 Power in Light, also rejected. The next admissible variant, adopted above, is the best that keeps all four unmodeled talents.
- `sim/request/rotations_smoke_test.go` needed no change: both specs' entries are for bare builds and still pass.
