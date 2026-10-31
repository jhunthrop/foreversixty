# Mobile Wave 1 integration C: F4 raids on main (2026-10-31)

Branch `wt/mobile-integrate-w1c`, from origin/main ed1a4fdd6, where integration B (F3 BiS) had landed. F4 was merged with `--no-ff` and no rebase. The lane's own report is `2026-10-31-mobile-raids.md`. Later merges of origin/main brought in the iOS CI record mode (e9edcdd45) and remove-spine. The A1 to A3 API lanes stay held.

## Reconciled: one of each

| Shared thing | Before | Now |
|---|---|---|
| Raids tab | F4's own five-tab shell, its own gates (`raidsEnabled`, `RaidsFeature.enabled` over appConfig and mode), `/v1/me`'s main character | Plugged into iOS `RootView.screen(for: .raids)` and Android `ForeverAppScaffold` behind `isAvailable(.raids)`. Raid links and notifications still select the tab. F4's gate code is deleted |
| Character selector | Home and BiS only | Also on the Raids root: iOS `.characterSelector(isOn:)`, Android `CharacterSelectorBar(session)`. The tab follows the session's current character (`RaidsTabModel.currentKey`) and reloads when it changes. On iOS the Raids root no longer shows an inline title beside the selector, so the bar matches Home and BiS |
| Android selector row | A row that appeared only once `/v1/me` answered, so the screen shifted | A full-width row with a 48 dp minimum height at every font scale. It is reserved while a signed-in session's characters arrive, on Home (including Home's own `/v1/me` path), Raids and BiS. BiS lost its extra 4 dp top padding, so the pill sits at the same y on all three. This is recorded as a spec deviation in §3.4 |
| Sign-out | Raids cache untouched | `RaidCache.removeAll()` on both platforms, called from the session's `clearCaches` |
| Design tokens and roles | iOS `Theme+Game.swift`, Android `RaidType` and `RaidColors` | Deleted. Type roles live in `ScaledType` / `FsType`, which gain `cardHead`, `segment`, `monoMeta` and `monoDay`. Colours come from `FSColor` / `FsColor`, `ClassColor` / `classColor` and `ItemQuality`. `GameColors` gains `rowWash`, `parse()` and, on Android, `offlineGround` |
| Crests | F4 `CrestView` / `Crest`, a `crests` bundle folder and a `forever.crests` Gradle plugin | B's `ClassCrest`, from the one `forever.bundled-data` copy. F4's plugin and its folder entry are deleted |
| Game icons | F4 `RaidIcons` plus AsyncImage, and a second Coil path | The shared loader: iOS `GameIconImage` (with the 200 MB URLCache) and Android `GameIconImage` (with `ItemIconLoader`). Both are now internal, with `fallback` and `dimmed`. `LocalRaidIconBuild` / `\.raidIconBuild` carries the build. The icon URL test reads the shared loader |
| Tags, offline strip, cards | F4 copies (`StatusTag`, `OfflineStrip`, `primaryCard`, `quietCard`) | The DS `TagView` / `Tag` (plus a `TagView(chip:)` / `ChipTag` mapping), the DS `OfflineStrip` and the DS cards. Android DS cards now clip their content to the panel shape |
| Secondary buttons | F4's own uppercase styles | DS `BrandButtonStyle` / `BrandButton`, in sentence case. A new `accent` kind (gold label and border) serves the reminder ask's Turn on. Android `BrandButton` gains `enabled`, and a disabled button dims. "Not now" is now in sentence case |
| Offline stamp | F4 `RaidClock.stripTime` with its own words | Delegates to `OfflineStamp` |
| Coil version | Declared twice in libs.versions.toml | Declared once |

## The follow-ups folded in

1. **Bless contrast (reviewer must-fix).** Before, a missing buff faded the whole tile, label included, to 35%. Now `dimmed` greys and fades only the art. The label is drawn at full opacity: `muted` on a dimmed tile (about 6:1 on `#131824`) and `strong` otherwise. The label is a fixed 11 pt/dp, so it never outgrows the tile at large text.
2. **Site data.** F4's change to `sync-data.mjs` (`ROOT_FILES`) publishes `/data/active-build.json` and `dates.json`.
   - `prune` touches only directories, so the root files survive.
   - `_headers` gives `/data/*` one hour plus stale-while-revalidate, so a new build is never cached for a year.
   - The vitest case passes, and so does the full suite.
3. **Android CharacterSelectorBar.** The row is reserved as described above.
   - Captured closed on Home at 1.0, 1.3 and 2.0, and on Raids and BiS. The band below it starts at y=310 px at all three scales, and the pill top is at y=207 px on all three tabs.
   - Captured after a relaunch, and on the switch to a second character, where Raids follows to that character's guild state.
4. **iOS `you-premium`.** The CI record run (`make -C apps/ios snapshots`, which rewrites only failing goldens) did not rewrite it: it already matches CI's iOS 26.5 runtime since e9edcdd45. It is unchanged.
5. **Band backgrounds.** One `surfaceBand(bottomLine:)` per platform, modelled on heroBand: `#14100b`, an ember radial at 22% with a 320 pt/dp radius from the top centre, and an optional soft bottom line.
   - Home's hero bands use it with the line. The signed-out screens use it without; they were bg with a 28% glow, which drifted from the spec's `--surface-band`.
   - Android now sizes the radial from its own bounds, not fixed pixels. `heroBand`, iOS `BandBackground` and Android `bandBackground` are deleted.
6. **iOS press scale.** `BrandButtonStyle` scales to 0.98 only when pressed, enabled and Reduce Motion is off: `BrandButtonStyle.scale(pressed:reduceMotion:enabled:)`, with four tests. Android keeps the ripple.

**CI snapshot mechanism.** The record run uploaded only `AppTests/__Snapshots__`, so the Raids goldens (`AppTests/Raids/__Snapshots__`) were never recorded on CI. ios.yml now uploads `AppTests/**/__Snapshots__/**`, and the Makefile copies the results back under `AppTests/`.

**Xcode 26.6 strictness.** CI caught two main-actor calls from nonisolated tests that the local Xcode accepted. The move-menu test is now `@MainActor`, and `scale` is `nonisolated`.

## Gates

Each command was run on its own and its exit code checked, with no pipes.

| Gate | Result |
|---|---|
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21) | exit 0, 298 JVM tests (app 190, auth 47, network 40, model 16, telemetry 5), 0 failures |
| `make -C apps/ios test` (local iOS 27) | Every non-snapshot test passes: app-hosted 181 tests in 26 suites, plus packages 115 (Networking 46, Auth 39, Models 21, Telemetry 5, Routing 4). The only failures are snapshots, which is expected locally per the CI-recording rule |
| ios.yml dispatched on the branch (Xcode 26.6 / iOS 26.5) | green: 181 app-hosted tests and all package suites passed (the last record run, on this branch's final code) |
| `npm run check`, `lint`, `format:check` | all exit 0 (check: 0 errors), using the dropdown worktree's node_modules, unlinked before commit |
| `npx vitest run` (full web suite) | exit 0, 690 files, 7776 tests |
| `node web/scripts/gen-mobile-{api,copy,features,links,tokens}.mjs` | no drift |
| `go vet ./...` | exit 0 |
| `go test -count=1 -p 1 ./internal/contract/ ./internal/server/ ./internal/guilds/` on `forever_test_w1c` (created, then dropped) | exit 0 |

## Goldens

- **Android, re-recorded locally (Roborazzi on the JVM).**
  - All 33 raids goldens.
  - Account goldens: the signed-out band, plus sub-pixel card-clip changes on delete and devices.
  - Home and gear goldens: the band radial is now sized from the band.
- **iOS, recorded on CI (Xcode 26.6 / iOS 26.5) through `make -C apps/ios snapshots`.**
  - The 30 Raids goldens that changed.
  - The 7 signed-out and import account goldens.
  - The reminder-ask golden after "Not now" passed CI within its precision, so it was not rewritten.
  - No iOS golden was recorded on the local simulator.

## Captures

All captures are in `design/reviews/2026-10-31-mobile-raids/`.
- `selector-home-{1.0,1.3,2.0}-and`, `selector-{raids,bis}-1.0-and`, `selector-home-relaunch-loaded-and`, and `selector-{home,raids,bis}-ios`.
- `selector-open-raids-{and,ios}` and `selector-raids-switched-{and,ios}`: the switch to Yellfury, where Raids shows "No guild yet".
- `raid-composition-integrated-and` and `raid-page-roster-ios`: live icons, with readable MotW and Bless labels.

The scripts are `capture-integrated-and.sh` (Android, under `tools/mobile-device.sh android-run`) and `RaidCaptureTests.testRaidsTabUnderTheSelector` (iOS, under the `ForeverSixtyCaptures` scheme).

## Verdicts

- **ux-designer: SHIP** (`ux-verdict-integration.md`).
- **wow-player: SHIP** after one re-check (`player-verdict-integration.md`).
  - Blocker 1, proof that the selector changes the Raids guild: cleared by the switch captures.
  - Blocker 2, the Home vs BiS band: moved to the Home/BiS owners, since it is outside this merge.

## Concerns and follow-ups

1. **Home vs BiS band.** The wow-player flagged that, for the same level-58 character, Home names its list `60 · raid` while BiS opens on `50 to 59`. This predates the merge (F2 and integration B, already on main). It belongs to the Home/BiS owners.
2. **Live buff tiles carry no name.** Fort and SPro, for example, are told apart only by their icon art. A long-press name or label should follow. MotW and Bless keep their labels until the icon export publishes `spell_nature_regeneration` and `spell_holy_fistofjustice`.
3. **The iOS selector truncates its detail line.** On the Raids root at the default size it reads `Hunt...`. The selector belongs to F1 and appears on every tab.
4. **The guild head's faction mark differs by platform.** Android shows a star and iOS a shield, both from the fixture's crest. The mono digits in "11 signed up" and the Android XL note placeholder clip are carried over from F4.
5. **The reserved-row frame was not captured.** The Android emulator draws its first frame after `/v1/me` answers, so the reserved row is shown by code and by the identical relaunch position, not by a mid-load frame.
