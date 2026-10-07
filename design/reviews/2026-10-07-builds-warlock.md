# Builds lane: warlock, 2026-10-07

Source: the talent-search reports regenerated on engine d1af3529f (`design/reviews/talent-search/warlock-*.md`, read from the main checkout). Ladder movement is the level-60 row of `sim/request/testdata/ladder/warlock-*.golden.md`, before and after, as a relative change (300 iterations, seed 1, so only indicative).

All three specs adopt the rule-abiding winner. The overall winners (deep Affliction for Affliction and Demonology, deep Destruction for Destruction) each drop unmodeled guide talents, and Demonology's also leaves its own tree, so none is taken.

## Affliction

- Old: `25552000130201051/235522/0`. New: `25552300120201351/200522/003` (37/11/3).
- Change: +3 Improved Drains, +3 Soul Siphon, +3 Bane; -1 Pandemic, -3 Improved Imp, -5 Demonic Embrace.
- Search gain: +60.2 DPS on 419.2, about +14.4%, combined error 1.0 (about 60 sigma).
- Rule check: (1) keeps every unmodeled guide talent: the finalist row lists none dropped; Soul Harvest, Improved Health Funnel, Demonic Aegis and Improved Voidwalker all stay at full rank. (2) Own tree majority: 37 of 51. (3) Gain beyond 1% and beyond error: yes on both. (4) Rotation: Pandemic, Improved Imp and Demonic Embrace are modeled non-damage or measured zero (probe diff 0.0); Wrack, Siphon Life, Amplify Curse (prepulled) and Nightfall are untouched.
- Ladder level 60: 310.3 to 343.2, about +10.6%. Top-casts list the same five abilities; no `zero_casts` line before or after.
- Rejected: deep Affliction (+24.7%) drops Soul Harvest, Improved Health Funnel, Demonic Aegis, Improved Voidwalker. The variants with "all non-damage points re-spent" (+86 to +89 DPS) drop Aegis, Voidwalker and Health Funnel, all unmodeled. Rows 5 to 7 each drop one unmodeled talent.

## Demonology

- Old: `255323/2352113101200001351/0`. New: `25532/233211310122000135/004` (17/30/4).
- Change: +2 Decimation, +4 Bane; -3 Improved Drains, -2 Demonic Embrace, -1 Demonic Pact.
- Search gain: +35.4 DPS on 359.7, about +9.8%, combined error 0.9.
- Rule check: (1) no unmodeled guide talent dropped (finalist 11: none); Demonic Aegis, Improved Health Funnel, Improved Voidwalker, Soul Harvest kept. (2) Own tree majority: 30 of 51. (3) Beyond 1% and error: yes. (4) Rotation: Demonic Embrace (-0.4 +- 1.6), Improved Drains and Demonic Pact probe at zero; Demonic Sacrifice and Soul Link are kept; Decimation and Bane add damage.
- Ladder level 60: 276.6 to 300.0, about +8.5%. Top-casts list unchanged in membership (Shadow Bolt rises); no `zero_casts` line.
- Rejected: deep Affliction (+17.0%) and deep Demonology (+14.4%) both drop Soul Harvest; deep Affliction also leaves the spec tree. Rows 8 to 10 ("all non-damage points re-spent") drop Soul Harvest.

## Destruction

- Old: `255323/0/2353225100101051`. New: `25532/0/2053225103101351` (17/0/34).
- Change: +3 Agonizing Flames, +3 Fire and Brimstone; -3 Improved Drains, -3 Improved Shadow Bolt.
- Search gain: +33.5 DPS on 392.8, about +8.5%, combined error 1.0.
- Rule check: (1) no unmodeled guide talent dropped (Destructive Reach, Molten Skin, Soul Harvest kept). (2) Own tree majority: 34 of 51. (3) Beyond 1% and error: yes. (4) Rotation: Improved Shadow Bolt probes at +0.00 in this guide's context and Improved Drains at 0.00, both modeled non-damage. Conflagrate, Shadowburn, Incinerate, Bane untouched.
- Improved Shadow Bolt check: the rotation fills with Incinerate (golden `zero_casts` for Shadow Bolt at levels 50 and 60 predates this change and is unchanged), so the talent is worth nothing to a solo caster here. Shadow Bolt remains castable, and its level-40 zero_casts line disappears after the change; levels 50 and 60 keep the same idle Shadow Bolt as before. No newly idle ability.
- Ladder level 60: 300.1 to 322.3, about +7.4%. Top-casts list identical.
- Rejected: deep Destruction (+11.3%) drops Destructive Reach, Molten Skin, Soul Harvest. Rows 2 to 4, 7 drop Molten Skin or Soul Harvest. Rows 5 and 8 (Molten Skin or Destructive Reach into Fire and Brimstone) each drop one unmodeled talent.

## Verification

`go test ./...` under sim/ passes; `FOREVER_UPDATE_GOLDEN=1 go test ./request/ -run TestRotationLadder` regenerated the three goldens; smoke expectations needed no change (the smoke build is separate and still passes). `npx vitest run src/content src/lib/guides` passes (341 tests). Guide prose quotes relative gains only.
