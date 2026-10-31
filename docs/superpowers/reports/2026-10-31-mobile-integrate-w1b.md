# Mobile Wave 1 integration B: F3 BiS on main (2026-10-31)

Branch `wt/mobile-integrate-w1b`, from origin/main 532a4fad7 (F2 Home and F1 sign-in already on main). F3 was merged with `--no-ff`, with no rebase. The lane's report is `2026-10-31-mobile-bis.md`.

## Commits

| Commit | What |
|---|---|
| 72290ba2f | Merge `wt/mobile-f3-bis`, with every reconciliation below |
| 930de1a1d | Re-recorded the screenshots the integration changed |
| b76c1221b | Integrated review captures and the scripts that took them |
| 38c07c470 | Merge origin/main (the logs engine counts lane: api, logs, web, plus one reports schema hunk) |
| this commit | This report |

## Reconciled: one of each

| Shared thing | Before | Now |
|---|---|---|
| Item tooltip | F2's full sheet, model and catalog, and F3's minimal copies in the same files | F2's files win on both platforms. F3's call sites compile unchanged against F2's entry points |
| Next best unit | F2: `{delta} DPS vs the pick`, and `Same DPS` only at exactly 0 | `tooltip.delta(delta, unit)` and `items.sameDps(unit)` take the band's unit word, from the band's `role`: DPS, HPS or score. The figure uses the BiS list's own `signedDelta`, so "same" applies under 0.05, the same as the list. `ItemNumbers.delta` was removed |
| Pick footer | F2: `… {band} · {preset label}`, named by the item's slot | `Best in slot for {slot} · {spec} {band}`, named by the BiS row's slot (`Ring 1`). Home now passes its band head (`60 · raid`) as the band label, so Home's footer still names the raid list |
| BiS tab | F3: a five-tab shell, its own gate, and `/v1/me`'s main character | Plugged into iOS `RootView.screen(for: .bis)` and Android `ForeverAppScaffold` behind `isAvailable(.bis)`. The list follows the session's current character (`SessionStore.currentCharacter` / `SessionManager.currentCharacter`), restarts when it changes, and waits while a signed-in session's characters are still arriving, so the signed-out first-visit picker never fires early. The main-character stand-in is gone |
| Character selector | Home only | Also on the BiS root. iOS uses `.characterSelector(isOn:)`, the one helper that Home now uses too. Android puts `CharacterSelectorBar(session)` in the bar's leading slot, which the open search replaces |
| Feature gate | F3's `BisScreen.isEnabled` / `isBisEnabled` | AppConfigStore's helper through `AppEnvironment.isAvailable` / `AppGraph.isAvailable`. F3's copies were deleted |
| Tag | DS `TagView` / `Tag`, plus F3's `Tag` with a tick | The DS tag gains an optional glyph (iOS) / `checked` (Android). F3's copy was deleted |
| Secondary button | DS button (uppercase), plus F3's sentence-case button | One DS button. A secondary is now in sentence case on both platforms, per spec §4.5 (only the primary is uppercase) and the BiS UX verdict. iOS `SecondaryButton` gains `fullWidth` |
| Offline strip | DS strip (a warning glyph on Android), plus F3's strip | The DS strip, with the spec's wifi-off glyph on Android (the drawable moved to core/designsystem) and announced once on iOS. One `OfflineStamp`; `HomeTime.stamp` delegates to it |
| Game palette | F3 `GamePalette` plus its Theme roles, and main's class colours | Item quality stays with `ItemQuality`. The rest is one `GameColors` in the design system: faction, sheet, tooltip panel, icon line, control fill and pressed wash. The tooltip's hard-coded panel colours use it too. F3's Theme additions were reverted |
| Brand images | F1 crest and seal (account), plus F3 crest and faction mark | One `common/BrandImages` per platform: F1's crest (the reviewed one), the faction mark and the seal |
| Icons | F2 `ItemIcon`, plus F3 `IconSource` / `RemoteIcon` / `SpecIcon` | F2's loader. `SpecIcon` joins it on the same source and takes the build |
| Words | F2 `ItemVocabulary`, plus F3 `ItemWords` | `ItemVocabulary.slotLabel`. `ItemWords` was deleted |
| Wrap layout, pill surface, row press | F3 app-local | iOS design system (public). Android `pillSurface` is in Surfaces |
| Gradle plugins | `forever.brand-art` (crests and seal) and `forever.bundled-data` (crests, factions and data) overlapped | One `forever.bundled-data`, which now also carries the seal. `forever.brand-art` was deleted |
| Coil | Two identical blocks | One |
| FixtureBundleTests | main's file, plus F3's copy in BisListTests | F3's copy was deleted |
| Contracts | Both lanes fixed the BiS schema | Both sets of fields kept: F3's `role`, `alternatives[].verified` and `low_value`, and F2's `source_line`, `drop_chance`, `subclass` and `weapon_slot`. The hunter fixtures are main's newer copies (F3 carried an older cherry-pick). Every generator was re-run, with no drift |

**Also fixed.** iOS `ScaledType.mono` now scales with the text size. It is the same at the default size, and BiS's figure no longer stays at 15 pt at AX5.

**Found on capture.** A reused iOS icon view kept the previous row's art when the list changed spec: the Holy list's crown appeared on the Protection Faceguard. `GameIconImage` is now keyed by build and icon.

## Captures

All captures are in `design/reviews/2026-10-31-mobile-bis/`.

- **`integrated-list-default-and.png` and `integrated-list-large-and.png`.** Fixture build, signed in through You.
  - The tab bar shows Home, Raids, BiS and You on one line at both 1.0 and 2.0.
  - The selector (Obnoxious Yell) sits in the bar, and the character's own list carries `Yours`.
- **`integrated-tooltip-healer-{and,ios}.png`.** Holy Priest, 60 Phase 1 raid, Shoulder: `−8.6 HPS vs the pick` on Android and `−8.9 HPS` on iOS (Horde vs Alliance list).
- **`integrated-tooltip-tank-{and,ios}.png`.** Protection Warrior: `−275.6 score vs the pick` on Android (Neck) and `−1517.4 score` on iOS (Back). On Android the footer reads `Best in slot for Neck · Protection Warrior 60 · Phase 1 raid`.
- **How they were taken.** The tooltips come from a live-mode debug build (`forever.apiMode=live` / `API_MODE=live`), which reads the real CDN files. The fixtures carry only fury and hunter.
  - iOS cannot tap, so two debug-only launch arguments open the list and the tooltip: `-FSBisSpec <spec>` and `-FSBisTooltip <slot>`.
  - The scripts are `capture-integrated-and.sh` and `capture-integrated-ios.sh`.

## Gates

Each command was run on its own and its exit code checked, with no pipes.

| Gate | Result |
|---|---|
| `make -C apps/ios test` | exit 0, TEST SUCCEEDED: 212 tests (app-hosted 97 in 16 suites, Networking 46, Auth 39, Models 21, Telemetry 5, Routing 4) |
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21) | exit 0, BUILD SUCCESSFUL, 218 JVM tests, 0 failures |
| `node web/scripts/gen-mobile-{api,copy,features,links,tokens}.mjs` | no drift |
| `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` | exit 0, 28 passed |
| `npm run check`, `lint`, `format:check` | all three exit 0 (check: 0 errors). Run with the dropdown worktree's node_modules; the link was removed before every commit |
| `go vet ./...` | exit 0 |
| `go test -count=1 -p 1 ./internal/contract/ ./internal/server/` on `forever_w1b` (created, then dropped) | exit 0, ok, ok |

All of these were re-run after merging origin/main (38c07c470), with the same results.

## Screenshots re-recorded

- **BiS, both platforms, all states.**
  - iOS: the 14 files in `AppTests/__Snapshots__/BisSnapshotTests/`.
  - Android: the 15 `bis_*.png` files.
  - They now draw the repo's icon files (as Home's tests do), the shared tag, buttons and strip, and F2's tooltip.
- **Android account.** `signedout-default`, `-401`, `-deleted`, `-ax5`, `signin-browser-open`, `signin-failed`, `home-no-characters` and `you-premium`: secondary buttons are now in sentence case, plus the Founder tag.
- **Android Home.** `home-error`, `-nogear`, `-nospec`, `-under20` (sentence-case buttons), `home-offline` (the wifi-off glyph), and `home-default`, `-full`, `-allmatch`, `-nolist` (the shared tag's tick row and padding).
- **Android tooltip.** `tooltip-large` and `tooltip-weapon`: the footer now reads `60 · raid`.
- **No other iOS goldens were touched.** The iOS account and Home snapshots stayed inside their precision.

## Concerns

1. **Secondary buttons changed case everywhere.** Sentence case (spec §4.5, the BiS UX verdict) now also applies to F1 and F2 surfaces: `Pick your class`, `Try again` and the account buttons. It is right by the spec, but F1's and F2's reviewers saw uppercase.
2. **New iOS BiS goldens were recorded on the local simulator.** If CI's iOS runtime differs (the coordinator's iOS 27 vs 26.5 note), they may need the same treatment as the wt/ios-ci lane's fix.
3. **The iOS list was not captured signed in.** simctl cannot tap Sign in, and the main-character stand-in that F3 used for its list captures is gone by design. The Android captures show the signed-in list and the selector.
4. **The iOS tooltip opens at the medium detent.** The tank capture uses Back, whose tooltip is short enough to show the score line; Neck's set bonuses push it below the fold.
5. **The tank headline bolds `Tank` in mono.** The set line styles its first word as a figure, and for tank lists that word is "Tank". This is F3's behaviour, left unchanged.
6. **Home's tooltip footer** now reads `Beast Mastery 60 · raid` instead of `60 · Raid-ready, Phase 1`.
