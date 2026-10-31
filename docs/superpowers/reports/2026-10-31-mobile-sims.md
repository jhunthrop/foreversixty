# Mobile Wave 1 F5: Sims (2026-10-31)

Branch `wt/mobile-f5-sims`. One lane, both platforms, from the same contract fixtures. Spec: `design/specs/2026-10-30-mobile-apps.md` §5.5, revised for today's API. Mock boards: `design/mocks/2026-10-30-mobile-apps/renders/f5-*` and the XL.5/XL.6 cells of `xl-*`.

## Design

- ux-designer revised §5.5 for the API as it is (the mobile API lanes stay held until after Nov 4) and re-rendered the boards.
- Mock verdicts: `design/reviews/2026-10-31-mobile-sims/ux-verdict-mock.md` (SHIP after one fix round) and `player-verdict-mock.md` (SHIP after Vis'kag became a +2 DPS Spritecaster Cape).

## What was built

| | iOS (`apps/ios/App/Sims`) | Android (`app/.../sims`) |
|---|---|---|
| Rules | `SimsLogic.swift`: the site's seven row states, estimate words, newest five, poll targets, combos bar, rerun status | `logic/SimsLogic.kt`, the same |
| Words | `SimsCopy.swift` | `logic/SimsCopy.kt` |
| Data | `SimsStore` (both lists, cache first, each section fails on its own, paging, 5 s poll, rerun, seen), `SimResultStore` | `SimsStore`, `SimResultStore` |
| Result | `SimResultModel`: title, kind, meta, figure or headline, error line, gear in frame order, Top Gear's changed slot | `SimResultModel` |
| Screens | `SimsListView`, `SimRowCards`, `SimResultView`, `SimsScreen` | `ui/SimsList`, `ui/SimRows`, `ui/SimResult`, `SimsFeature` |

- **Entry.** Home's Sims row now pushes Sims (only that row is tappable; Talents is F6's). iOS uses a `SimsRouter` in the environment; Android a `SimsNavigator` on `SimsFeature`, created inside `HomeDependencies.create`.
- **Links.** `/app/sim/{id}` selects Home and pushes Sims, then the result (iOS `RootView.openLink`, Android `MainActivity` plus the scaffold's `wantsHome`).
- **Gate.** The `sims` feature key (on in debug and fixture builds).
- **Entitlements (decision G).** Server runs shows for `entitlements.server_sims`, or for an account that still has server rows. `Run again` exists only on failed server runs, as on the web; a 402 shows `Server sims are part of Premium.`
- **Cache.** Answers on disk beside the raids cache (`sims/`), cleared at sign-out on both platforms.
- **Copy.** New `sims.*` keys in `apps/shared/copy/en.json`; `sims.runAgain` removed; `sims.error` takes the iteration count as a formatted string (`10,000`).

## Build review

Captures (the snapshot goldens, default and largest text, both platforms) are in `design/reviews/2026-10-31-mobile-sims/`. ux-designer SHIP (`ux-verdict-build.md`) and wow-player SHIP (`player-verdict-build.md`), each after one fix round: Android `Starting…` label, Android age stacking, iOS condensed-caps titles and `14 Dec` dates, pushed captures with the back control, cached icons on the offline result.

## Gates

- `make -C apps/ios test` (shared iOS 26.5 simulator): TEST SUCCEEDED, 225 app-hosted tests plus the package suites.
- `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21): BUILD SUCCESSFUL, 342 JVM tests, 0 failures.
- Copy codegen re-run; `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` 28 passed; `npm run check`, `lint`, `format:check` exit 0. No openapi or fixture change, so no Go contract run.

## Shared-file edits (small, additive)

`HomeScreen.swift`/`HomeScreen.kt`, `HomeAnswerView.swift`/`HomeAnswer.kt` (the Sims row and `OpenNext`), `RootView.swift`, `AppEnvironment.swift` and `AppGraph.kt` (sign-out cache), `MainActivity.kt`, `ForeverAppScaffold.kt`, `en.json` plus its generated outputs.

## Gaps (today's API, noted in spec §5.5)

1. No rerun of a finished run; F5.3 has no `Run again on the server`.
2. No best-in-slot reference figure or preset on a result.
3. No iteration total for a plain run, so only bulk runs get a bar.
4. `SimResult` lacks `finished_at`, `kind`, `headline`: a deep link reads the list's first page for them.
5. **See plans** in the lapsed line is wired fail-closed (`mayLinkToPlans` returns false) until lane F8's storefront read (`PlansLink`) is on main; then it is a one-line change in `SimsDependencies.live` and `SimsFeature.mayLinkToPlans`.
