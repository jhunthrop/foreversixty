# Mobile Wave 1 integration A: F2 Home and F1 sign-in on main (2026-10-31)

Branch `wt/mobile-integrate-w1a`, from origin/main 4761fe711. Both lanes merged with `--no-ff`, no rebase. Lane reports: `2026-10-31-mobile-home.md` (F2) and `2026-10-31-mobile-signin.md` (F1).

## Commits

| Commit | What |
|---|---|
| 86602c73f | Merge `wt/mobile-f2-home`: Home, the tooltip sheet, the four-tab shell |
| 41438c992 | Merge `wt/mobile-f1-signin`, on F2's shell. You is plugged into the slots, Home runs on the session with the selector, there is one feature gate, sign-out clears Home's cache, and the age copy is `common.age.*` |
| 7d92515f5 | iOS fixture mode reads the bundled fixtures from `app/`, with tests |
| 3cc827e80 | One design system for Home and the account screens |
| cb9343391 | The tooltip stops printing the item name twice |
| c105bca16 | Re-recorded screenshots and refreshed tooltip review captures |
| 3298ea630 | Merge origin/main (the sample chart lane: web/ and `api/cmd/samplecut` only) |

## The merge

- **Conflicts.** There were five, all in the shell: `RootView.swift`, `ForeverTab.kt`, `ForeverAppScaffold.kt`, `MainActivity.kt` and `AppGraph.kt`. F2's four tabs won: `home, raids, bis, you`, with catalog labels and F2's icons. F1's five-tab enum, its `shell.tab.*` labels and its `Character` / `Sims` cases are gone.
- **You.**
  - iOS: `RootView.screen(for:)` has `case .you where signInAvailable: YouScreen(...)`.
  - Android: the scaffold's You branch is `YouRoute(account, openBis)`.
  - With the sign-in gate off, You falls back to the placeholder.
- **Home with the session.**
  - Home sits inside F1's `HomeAccountGate`, so the signed-out, importing and no-characters states come from F1.
  - Signed in, `HomeTab` / `HomeRoute` follows the session's current character (`CharacterSelection.current`, spec §3.4), not `/v1/me`'s main. While the session's `/v1/me` has not answered yet, Home falls back to its own read.
  - Selector on iOS: `.characterSelector()` in Home's `NavigationStack`, the call the F1 report gives. It replaces the name title.
  - Selector on Android: Home has no TopAppBar. A new `CharacterSelectorBar(session)` (the session-bound twin of the iOS modifier) sits above the pull-to-refresh answer. Raids and BiS can make the same one-line call.
- **One feature gate.**
  - `isAvailable(feature, mode, isDebugBuild)` is now a member of `AppConfigStore` / `AppConfigRepository`. F1's `AppConfigStore+DevBuilds.swift` and `FeatureAvailability.kt` are deleted.
  - The app calls it in one place each: `AppEnvironment.isAvailable(_:)` (the only `#if DEBUG`) and `AppGraph.isAvailable(_)` (with `BuildConfig.DEBUG`).
  - Home's `character_home` uses it too. Before, it was on in fixture builds only; now it is also on in debug live builds, as the brief says.
  - Tests cover it on both platforms (the Android test is new).
- **Sign-out clears Home's answer cache.** F1's concern 5. iOS adds `FileHomeCache.standard()?.removeAll()` to `clearCaches`. Android passes `clearCaches = { home.cache.removeAll() }`, so `SessionManager.clearCaches` is now `suspend`.

## Reconciled: one of each

| Shared thing | Before | Now |
|---|---|---|
| Type roles | iOS `HomeType` + `ScaledType`; Android `HomeType` + `AccountType` | iOS `ScaledType`, Android `FsType` (design system), plus `stat`, `lead`, `tooltip`, `mono` |
| Buttons | F2 `PrimaryButton`/`SecondaryButton` + F1 `BrandButtonStyle`/`BrandButton` | One `BrandButtonStyle` / `BrandButton` (all four kinds, hug or fill, optional glyph). iOS keeps `PrimaryButton`/`SecondaryButton` as thin wrappers on it |
| Tag | F2 `TagView`/`Tag`, F1's inline capsule on Founder | `TagView` / `Tag` in the design system. Founder uses it |
| Section label | F2 `MicroLabel`, F1 inline micro text | `MicroLabel` in the design system |
| Offline strip | F2 only, in the Home folder | Design system, taking its line as text |
| Cards | F2 `quietCard` + F1 `accountCard` / `AccountCard` border (identical) | `quietCard`; `primaryCard` and `heroBand` move too |
| Game colours | F1 `ClassColor` / `classColor` (item quality is F2-only) | `GameColors` in the design system; `ItemQuality` stays in Items |
| Brand images | F1 only (crests, seal) | Unchanged: one `forever.brand-art` plugin, nothing checked in twice |
| Feature gate | F1 helper files + F2 inline check | The AppConfig store (above) |
| Coil | One catalog entry (F2) | Unchanged, once |
| Gradle plugins | `forever.brand-art` (F1) | Unchanged, one plugin |
| Copy keys | `home.age.*` and `account.age.*` (same strings) | `common.age.*`, one `RelativeAge.phrase` per platform (`AgeCopy`) |
| openapi.yaml + fixtures | Both lanes' additions auto-merged | Kept both; every mobile generator re-run, no drift |
| iOS fixtures | Read from `Fixtures/`, which does not exist | Read from `app/` (below) |

## iOS fixture loading

I checked the built `ForeverSixty.app`: the fixtures sit at `app/`, and no `Fixtures/` folder exists. An XcodeGen folder reference keeps its on-disk name, and `name: Fixtures` only labels the group. So every fixture-mode read missed, including Home's static data and F1's fixture sign-in device id (`AccountStorage` hard-coded `"Fixtures"` as well).

**The brief asked for the fixtures to land as `Fixtures`. I did the opposite, on purpose.** I applied F3's fix (aedc28eb6) byte for byte: `FixtureTransport.bundledFolder = "app"`, the reader looks there, and the project.yml comment says why. `AccountStorage` now reads through the same constant. Matching F3's text keeps F3's merge clean.

Tests are in the new `AppTests/FixtureBundleTests.swift`: the bundle carries `me.json` where fixture mode reads it, and fixture sign-in finds its device id.

**Merge note for F3:** F3 adds the same `FixtureBundleTests` suite to `BisListTests.swift`. When F3 merges, the duplicate struct name will fail to compile, so delete F3's copy and keep this file.

## Player notes from the F2 re-check

1. **Off hand "none:" prefix: left as it is.** Discord says the same thing. `api/internal/discordbot/copy_lookup.go` has `BisOffHandTwoHand: "**Off hand** · none: {main hand name} is two-handed."`, and the gear-panel payloads in `design/mocks/2026-10-29-discord-gear-panel` read `**Off hand** · none: Legionite Glaive is two-handed.` The app matches Discord word for word. If "none:" goes, Discord and the app should change in the same edit.
2. **The weapon tooltip printed the item name twice**, once in the sheet header and again on the panel's first line. The panel keeps the name, because the game's tooltip opens with it. When the tooltip is loaded, the header now shows only the icon and `Rare · Main hand`. While loading or failed, there is no panel, so the header keeps the name. The fix is the same on both platforms.

## Gates

Each command was run on its own and its exit code checked, with no pipes.

| Gate | Result |
|---|---|
| `make -C apps/ios test` (shared simulator, behind the lock) | exit 0, TEST SUCCEEDED: 172 tests; app-hosted 63 (10 suites), Auth 40, Networking 39, Models 21, Telemetry 5, Routing 4 |
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21) | exit 0, BUILD SUCCESSFUL, 177 JVM tests, 0 failures, lint clean, debug APK built |
| `node web/scripts/gen-mobile-{api,copy,features,links,tokens}.mjs` | exit 0. After the merge: no drift; after the age keys: copy regenerated |
| `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` | exit 0, 28 passed |
| `npm run check` | exit 0, 0 errors. With the charts worktree's node_modules it failed: those node_modules have no `@floating-ui/dom`, which main's `src/lib/ui/popover.ts` needs, and nothing here touches web/. The passing run used `/Users/jh/code/forever/.claude/worktrees/dropdown/web/node_modules`. The symlink was removed before every commit |
| `npm run lint`, `npm run format:check` | exit 0, exit 0 |
| `go vet ./...` (api) | exit 0 |
| `go test -count=1 -p 1 ./internal/contract/ ./internal/server/` on the private database `forever_w1a` (created, then dropped) | exit 0, ok, ok |

**After merging origin/main (3298ea630).** Main touched only web/ and `api/cmd/samplecut`. These were re-run, each exit 0: `npm run check` (0 errors), `lint`, `format:check`, the mobile drift vitest (28 passed), `go vet ./...` and `go test ./cmd/samplecut/`. Nothing under apps/ or the contract changed, so the platform gates stand.

## Screenshots re-recorded, and why

- **Android account, all 22** (`account_*`): the tab bar is now F2's four tabs (it had F1's five), plus the shared type roles, buttons and Founder tag.
- **Android Home:** `home-under20`, `home-nogear`, `home-nospec`, `home-error`. These states draw buttons, and the unified button's label padding differs slightly.
- **Android tooltip:** `tooltip-medium`, `tooltip-large`, `tooltip-large-xl`, `tooltip-vendor`, `tooltip-weapon`. The header no longer repeats the name.
- **iOS:** `tooltip-medium`, `tooltip-large`, `tooltip-large-xl`, `tooltip-vendor`, `tooltip-weapon` (the name once). Only `-large-xl` failed the comparison; the other four were inside tolerance but still showed the name twice, so they were re-recorded too.
- iOS account snapshots stayed within their 0.99 / 0.98 precision. The iOS Founder tag change (the capsule outline becomes the shared tinted tag) is drawn but not re-recorded.
- The review captures under `design/reviews/2026-10-31-mobile-home/tooltip-*` were refreshed from the new goldens.

## Concerns

1. **Two band backgrounds remain.** F1's `BandBackground` / `bandBackground()` (ember 0.28 over `bg`, full screen) and F2's `heroBand` (ember 0.22 over `#14100b`, a strip with a bottom line) both claim `--surface-band`, and they look different. Unifying them needs a design call, so I left both.
2. **Android Home has no top app bar.** The selector sits in a padded row above the answer. Nobody has reviewed it on screenshots; a reviewer should check it against the spec's §3.4 board on both platforms. On iOS the name title gives way to the selector.
3. **Founder tag shape.** F1's reviewers passed a capsule outline. It now uses the shared tag (6 pt radius, tinted ground). Worth a glance in `account_you-premium.png`.
4. **The brief asked for `Fixtures`; I did `app/`.** The reason is under iOS fixture loading above. F3's merge must drop its duplicate `FixtureBundleTests`.
5. **Pressed state.** F1's buttons now scale to 0.98 (spec §8) instead of dimming to 0.8. Snapshots cannot show this.
