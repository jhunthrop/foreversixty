# Mobile Wave 1 F6: Talents (read only), both apps (2026-10-31)

Branch `wt/mobile-f6-talents`. Spec `design/specs/2026-10-31-mobile-talents.md` (binding over §5.6 of the app spec); boards `f6-talents`, `f6-rank`, `f6-states`, `f6-xl`. Mock verdicts: ux-designer SHIP, wow-player SHIP after one fix round. Build verdicts are in `design/reviews/2026-10-31-mobile-talents/`.

## What shipped

- **One screen, pushed from Home's Talents row**, behind the `talents` switch (on in debug and fixture builds). No tab, no deep link, never signed out.
- **Data:** only the static talent and BiS files and `/v1/me`. The guide is the planner's `selectBand` (raid, else bare) for the character's band (60, else the decade from 20; none under 20) and faction. Your build is `build.talents.trees`, drawn on the active build's trees only when its own build lays them out the same (else the stale state).
- **The compare** is the planner's `diffAgainstBand`, on device, one pass: `diffCount`, `missingCount`, the dashed cells, the signed chips and the tab dots. Matches means no dashed cell. Totals are summed from the drawn ranks.
- **Home's Talents row** reads the same status from one shared `TalentsModel` per character, so Home and F6 always agree. The home endpoint's `talents` field is a fixture while lane A3 is held and is not read.
- **Rank sheet:** the talent file's rank text verbatim, Next rank, the game's unmet requirement lines, and on Your build the guide's rank. A system sheet, so it never clips; the tapped cell scrolls clear and keeps its ring.
- **Shared files touched:** Home (row tap, `HomeCharacter.talents`, the presenter's optional F6 status), `AppGraph` / `HomeDependencies`, BiS's static-file factory (now shared, Android), `en.json` (`talents.*`, and both differ lines say "guide build"), `cdn.talents.warrior` (the generator's full file) and `me.json` (a legal Fury build for the fixture character).

## Tests

- Same lists on both platforms: `TalentRules` (16), `TalentsPresenter` (14, including Home's row), `TalentsModel` (8).
- Snapshots at default and the largest size: iOS 17 goldens on the iOS 26.5 shared simulator (CI's runtime), Android 18 Roborazzi goldens.
- Gates, each exit 0: `make -C apps/ios test`; `./gradlew test :app:verifyRoborazziDebug lint assembleDebug`; codegen drift, the mobile vitest, `npm run check`, `lint`, `format:check`; Go `contract` and `server` on a private database, created and dropped.

## Open, for the owner

1. **Band strings hold `band - 9` points**, so a level 58 against the 41-point 50 to 59 build always reads Differs. This matches the site planner. The screen explains it with `talents.beyond`. A fairer rule would change `diffAgainstBand` on the site and in both apps. **Resolved 2026-10-31:** points beyond the guide's total no longer count as differences (site and both apps); a level 58 holding the whole 41-point build reads Matches, `+8 beyond the guide`. Bands stay sized to their starting level by design: the BiS generator sims character, gear and talents at the band's first level (`sim/leveling/talents.go` `level - 9`).
2. **At 360 dp the three tree tabs stack**, because the labels do not fit one line. This is the spec's fallback, and it costs about two tiers of the first screen.
3. **When lane A3 ships the home endpoint**, a test should assert its `talents` equal `TalentCompare` on the fixtures, and the endpoint should count the same way.
