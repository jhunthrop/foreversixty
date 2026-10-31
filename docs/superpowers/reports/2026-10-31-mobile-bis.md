# Mobile Wave 1, lane F3: Best in slot (2026-10-31)

Branch `wt/mobile-f3-bis`, from main at 4ba7c1b70, merged with main at 5d86a8dc7 (the shared
device lock). Not pushed. F3 builds the spec's §5.3 on both platforms: the list (F3.1), the picker
sheet (F3.2), offline (F3.3), loading and error, and the item tooltip on a tap.

## Commits

| Commit | What |
|---|---|
| 81563700e | `fix(contracts)`: the BiS and item schemas match the generator's files; a contract test validates every BiS and items file of the active build; fuller warrior-fury fixtures |
| a21f436f1 | `feat(ios)`: the BiS tab, the picker, the offline copy and the minimal tooltip, the `bis.*` copy, and the `-FSTab` debug launch argument |
| f66ec35aa | `feat(android)`: the same on Compose |
| 265b8ad06 (F2's 7bb25a4d5, cherry-picked) | F2's contract and copy commit, with the spec conflict resolved to the union |
| 2b652c4bc | Reuse F2's `items.slot.*` and `items.quality.*` words; Coil at F2's 3.0.4 |
| last on the branch | iOS fixture mode finds its bundled fixtures (a bug on main); the system minimised search; `Bare` for the bare preset; sheet insets; captures; this report |

## What was built

**Data.** Static CDN files only, never re-ranked:
- `/data/<build>/bis/<spec>.json`, `items/<class>.json` and `classes.json`.
- The build is `GET /v1/app/config`'s `active_build`. Before any config, the app falls back to the bundled `web/src/data/active-build.json`. (`/data/active-build.json` answers 404 on the CDN today; see app-v1.md's open item.)
- Every good fetch is saved:
  - iOS: `Application Support/static-data`, excluded from backup.
  - Android: `noBackupFilesDir/static-data`.
- A fetch that fails for lack of a network answers the saved copy, stamped with its save time. Any other failure is the error state.
- The saved copy draws at once on launch, then refreshes.
- This is a generic `CachedStaticDataClient` in Networking (iOS) and core/network (Android), for F2 and F6 to reuse.

**The list.**
- 17 rows in the file's order. Each row has the item icon with its quality ring, and the name in the quality colour.
- The source line is in the site's words: `Quest:`, `Crafted:`, `Vendor:`. Every other kind is printed as published.
- `Next: {item}, −1.3 DPS`, or `same DPS` under 0.05, or `Nothing else ranked`.
- `Gain not modelled yet` when the slot has `effect_unmodelled`.
- `You wear it` when the item is any of the character's worn ids.
- The off-hand-under-a-two-hander and no-pick lines.
- The set line with `Updated d Mon yyyy`.
- Healer bands read HPS and tank bands read points; the band's `role` was added to the spec.
- `Yours` shows when the selection is the character's spec, range (by the site's `bandForLevel`; 60 opens the raid list) and faction.
- Search filters the rows by pick and runner-up names.
- Pull to refresh.

**The picker.**
- Nine crests in class colours, from `classes.json`, falling back to the catalogue.
- Specs with icon and role, from the bundled `data/curated/specs.json`, which the CDN does not publish.
- The six level ranges and both factions.
- Each choice applies at once.
- `Back to {name}'s list` shows only when the picks differ.
- Signed out, the first visit opens the picker with Warrior preselected.

**States.** Loading (17 skeleton rows), offline (the strip and its stamp), not saved yet, no list for this range, and error with Try again.

**Gate.** Feature `bis`: on in debug and fixture builds, else the config's switch. Off shows the old placeholder.

## Files created for the shared tooltip (merge with F2)

F2's tooltip was not on main, so F3 built a minimal one with an obvious name:
- iOS: `apps/ios/App/Items/ItemTooltipSheet.swift` (`ItemTooltipContent`, `ItemTooltipAlternative`, `ItemTooltipSheet(content:)`).
- Android: `apps/android/app/src/main/kotlin/gg/foreversixty/app/items/ItemTooltipSheet.kt` (`ItemTooltipContent`, `ItemTooltipAlternative`, `ItemTooltipBody(content)`, shown in a `ModalBottomSheet`).

It draws:
- the header (icon 56, the name, then quality · slot);
- a panel with name, slot, armor, effect text, Requires Level and Item Level;
- Where it comes from;
- The next best (up to two);
- the pick footer.

It leaves out stats, the weapon line, Binds, Unique, armor type and sets, which are F2's. Both files carry a MERGE NOTE: F2's body wins, and F3 reads only the content type and the entry point.

Other shared pieces F3 added, each named for reuse. If F2 or F1 drew the same thing, fold one into the other:
- **iOS.**
  - `App/Items/ItemIcon.swift`: `ItemIcon`, `SpecIcon`, `RemoteIcon`, `IconLoader` (200 MB URLCache), `IconSource`.
  - `App/Items/ItemWords.swift`.
  - `App/Common/OfflineStrip.swift` (with `OfflineStamp`).
  - `App/Common/Controls.swift`: `Tag`, `SecondaryButton`, `RowPressStyle`, `pillSurface`.
  - `App/Common/BrandImages.swift`: `ClassCrest`, `FactionMark`.
  - `App/Common/WrapLayout.swift`.
  - `Sources/DesignSystem/GamePalette.swift`, and roles added to `Theme`.
- **Android.**
  - `items/ItemIcon.kt` (Coil).
  - `items/ItemWords.kt`.
  - `common/OfflineStrip.kt`.
  - `common/Controls.kt`.
  - `common/BrandImages.kt`.
  - `designsystem/GamePalette.kt`.
  - The `forever.bundled-data` Gradle plugin: specs, the fallback build, crests and faction marks.

## Tests

One list, both platforms:
- the band rule;
- band selection;
- file order;
- the set line;
- the next line and its threshold;
- units;
- source words;
- the unmodelled and worn flags;
- the empty rows;
- search;
- every real published BiS file of the active build decodes;
- the character's own list;
- Yours and Back;
- the bundled resources;
- the model states (signed in, signed out first visit, choice and Back, offline saved, offline never fetched, error and retry, no list);
- the offline stamp;
- the cache (fresh, offline, offline without a copy, a 404 is not offline, the saved copy, the disk cache).

Screenshots cover the list, offline, loading, error, not saved, the picker and the tooltip, at default and largest text: iOS AX5 and Android font scale 2.0.

| Check | Exit | Result |
|---|---|---|
| `make -C apps/ios test` (shared simulator) | 0 | TEST SUCCEEDED: 133 tests (app-hosted 39, Models 21, Networking 45, Auth 19, Routing 4, Telemetry 5) |
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21) | 0 | 129 JVM tests, 0 failures; goldens verified; lint clean (warnings as errors); debug APK built |
| `go test ./internal/contract/`, `go vet` | 0 | ok, including the new `TestCDNGeneratorFilesMatchTheOpenAPISpec` (38 files). It failed on the old spec and passes on the fix. |
| `vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` | 0 | 28 passed |

## Captures

All captures are in `design/reviews/2026-10-31-mobile-bis/`. The scripts that took them, `capture-ios.sh` and `capture-and.sh`, run under `tools/mobile-device.sh`.

**Live, on the shared devices.** These are the debug fixture build, signed in as the fixture's level-58 Fury warrior, so the list is the 50-to-59 Horde list marked Yours. Icons come from the real CDN.
- iOS: `list-default-ios.png` and `list-large-ios.png` (AX5). The `-FSTab bis` launch argument opens the tab, because simctl cannot tap.
- Android: `list-default-and.png`, `list-large-and.png` (font scale 2.0), `picker-default-and.png` and `tooltip-default-and.png`. The script drives these with adb taps.

**Rendered by the screenshot tests.** These cover the states a fixture build cannot reach on a device: offline, loading, error and not saved. They also cover the picker, the tooltip, and both text sizes on both platforms. Files are named `render-<screen>-<state>-<size>-{ios,and}.png`. They come from the same views on the same simulator (iOS) and from Robolectric (Android), with no icon source, so each tile shows only its quality ring.

## Concerns

1. **Fixture mode was broken on iOS on main.** `FixtureTransport.bundled()` looked for a `Fixtures` folder. XcodeGen's `name:` labels only the group, so the folder lands in the bundle as `app/`. Every fixture-mode call failed, including Home's `/v1/me`. This lane fixes it with `FixtureTransport.bundledFolder = "app"`, plus an app-hosted test. Other lanes on main still have the bug until this merges.
2. **Contract overlap with F2.**
   - Both lanes fixed the same four schema kinds.
   - F2's commit is cherry-picked here, and the spec conflict is resolved to the union: F2's descriptions plus F3's `role`, `alternatives[].verified` and `low_value`.
   - Whichever lane merges second will see the same hunks already applied.
   - F3 also adds `TestCDNGeneratorFilesMatchTheOpenAPISpec`, which validates every real BiS and items file of the active build against the schemas.
3. **The selector slot is empty.** The BiS bar's leading slot is reserved for F1's character selector. BiS reads the main character from `/v1/me` until F1's switcher exists.
4. **Spec deviations.**
   - `bis.setDps`, `bis.next` and `bis.nextSame` take a `{unit}`. Healer bands publish HPS and tank bands a score, and printing "DPS" there would be wrong.
   - Dungeon, rep and PvP sources print the file's own words: `Stratholme: Balzaphon` and `Warsong Outriders`. The mock draws `Stratholme · Balzaphon` and `(Exalted)`, but the BiS file carries neither the dot form nor the standing, and the app re-derives nothing.
   - Both rings read `Ring` (the site's list word, `bis.slot.ring`).
5. **The minimal tooltip.**
   - It omits stats, the weapon line, Binds, Unique, armor type and set lines. F2's full sheet replaces the body.
   - `tooltip.delta` says "DPS vs the pick" even for a healer band; F2's key has no unit.
6. **Shared pieces may duplicate F1 and F2.** These are the offline strip, tags, the secondary button, `GamePalette`, the Theme roles, `WrapLayout`, crests and faction marks, and the `forever.bundled-data` plugin. Each is listed above under an obvious name, for whoever merges second to fold into one.
7. **The feature gate is local.** `BisScreen.isEnabled` and `isBisEnabled` turn BiS on in debug and fixture builds. There is no shared "on in debug" helper, so each lane writes its own; one belongs in AppConfigStore and AppConfigRepository.
8. **Data the app bundles because the CDN lacks it.**
   - The spec catalogue comes from `data/curated/specs.json`.
   - The fallback build comes from `web/src/data/active-build.json`; `/data/active-build.json` is still 404.
   - Publishing `specs.json` per build would let the app read it like the rest.
9. **Visual notes.**
   - JetBrains Mono is not bundled, so figures use the system monospaced face.
   - At Android font scale 2.0 the shell's tab labels break inside the word ("Chara/cter"). That is the shell's NavigationBar, not BiS.
   - The search is iOS's system search, minimised to the bar's trailing button. On Android it is a field that the bar's search button opens.
10. **Offline on a device** cannot be captured from a fixture build, which never touches the network. The offline captures are the rendered ones.

## Fix round 1 (both reviewers HOLD, 2026-10-31)

The verdicts are `design/reviews/2026-10-31-mobile-bis/ux-verdict.md` and `player-verdict.md`, committed with the fixes.

| # | Coordinator item | What changed |
|---|---|---|
| 1 | iOS picker, live | The debug-only `-FSBisPicker open` launch argument opens the real sheet; `capture-ios.sh` shoots it at default and AX5 (`picker-default-ios.png`, `picker-large-ios.png`). See the hairline note below. |
| 2 | Android picker at 2.0 | The picker and tooltip contents pad by the navigation-bar window insets. `capture-and.sh` captures the picker at 2.0 scrolled to its end, Faction (`picker-large-bottom-and.png`). Back does not show on the character's own list., and the tooltip's footer above the gesture bar (`tooltip-large-bottom-and.png`). The render test draws the sheets alone, not inside the tab shell, plus `bis_picker_bottom_large` scrolled to Back. |
| 3 | Next best in the band's unit | `tooltip.delta` is now `{delta} {unit} vs the pick` and `items.sameDps` is `Same {unit} as the pick`, from the band's role: DPS, HPS or score. Tested on real holy-priest and protection-warrior files on both platforms. |
| 4 | No pipeline words | The footer is `Best in slot for {slot} · {spec} {band}`. The band label already carries the phase (`60 · Phase 1 raid`), so no preset is printed and `Bare` is gone. `tooltip.pickOf` lost its `{preset}`. |
| 5 | Tank score explained | The tank unit word is `score`, as the site's `scoreUnitWord` uses. A tank band's `set_dps` is its own damage, which the site does not headline (`headlineDpsOf`), so the line above the rows reads `Tank lists rank by score, the tank sim's blend of mitigation, risk and threat. Higher is better · Updated {date}`. The wording comes from the site's `tankCopy` doc ("a composite of mitigation, risk and threat"). The deltas read `−36.8 score`. |
| 6 | Raid chip before raids open | The chip and the pill read `60 · Phase 1 raid · opens Dec 9` while now is before the `raids` entry of `dates.json`. The date is the entry's own `key`, and its `iso` day starts at 00:00 UTC, as the site's dates line reads it. The list stays viewable. Dates come from `/data/dates.json` through the offline cache, else the bundled `web/src/data/dates.json` (the CDN path is still 404). The `cdn.dates` fixture gains the raids entry. |
| 7 | Tooltip body | F3 does not build a second full sheet. Its call sites use F2's entry point exactly: iOS `ItemTooltipSheet(request: ItemTooltipRequest, catalog: ItemCatalog)`, with the pure body `ItemTooltipContent(request:phase:retry:)`; Android `ItemTooltipSheet(request: ItemTooltipRequest, catalog: ItemCatalog, onDismiss: () -> Unit)`, with `ItemTooltipContent(request, phase, retry)`. `ItemTooltipRequest`, `HomeItemLook`/`ItemLook`, `ItemTooltipSources` and `ItemCatalog(client:)` carry F2's names and fields. F3's minimal versions live in the same file names (`Items/ItemTooltipModel`, `ItemCatalog`, `ItemTooltipSheet`), so at integration F2's three files replace them and BiS's call site compiles unchanged. |

**What F2's model must change at integration.** The renamed keys will refuse to compile until it does:
- word the next best with the band's unit: `tooltip.delta(delta:unit:)` and `items.sameDps(unit:)`;
- call `tooltip.pickOf(slot:spec:band:)` without the preset.

**Also from the verdicts:**
- The two rings read `Ring 1` and `Ring 2`, and the trinkets `Trinket 1` and `Trinket 2` (F2's `items.slot.*`). UX ruled this a fix.
- Not saved has a way out: `Try again`, muted text and a 24 pt top margin. No list and error use the same block.
- Secondary buttons are in sentence case (`Back to Obnoxious Yell's list`, `Try again`).
- The iOS `Yours` tag sits on the title's first baseline at AX5.

**Not done here:**
- The kind icon before the source line: UX accepted the gap in writing as an option.
- `60 · dungeons and PvP` for the plain 60 pill: the site has no such copy, so it needs an owner call.
- Runner-up sources and PvP rank caps: the coordinator routes these to the ranking data.
- The re-scan of launch gating after tonight's re-rank: the coordinator runs it.

**Captures from this round.** Live on the shared devices, in `design/reviews/2026-10-31-mobile-bis/`:
- **iOS.**
  - `list-default-ios`, `list-large-ios`, `picker-default-ios` and `picker-large-ios` (AX5).
  - The live picker shows **no edge hairlines** on the chips or pills, so they were an artifact of the snapshot renderer. `pillSurface` now draws with `strokeBorder` inside the edge either way.
  - At AX5 the class grid drops to two columns so names never break inside the word, a defect the live capture showed.
- **Android.**
  - `list-default-and`, `list-large-and`, `picker-default-and`, `picker-large-top-and`, `picker-large-bottom-and`, `tooltip-default-and` and `tooltip-large-bottom-and`.
  - At 2.0 the sheet sits over the whole window, the app's tab bar included, and its end (Faction) clears the gesture bar.
  - `Back to {name}'s list` shows only when the picks differ from the character's. These captures are on the character's own list, so it is absent. The render `bis_picker_bottom_large` shows it scrolled to, away from the character's list.
- **Rendered states** (offline, loading, error, not saved, tooltip) live in the test goldens: `apps/ios/AppTests/__Snapshots__/BisSnapshotTests/` and `apps/android/app/src/test/screenshots/bis_*.png`.

**Gates, fix round 1.** Each ran on its own, with its exit code checked.

| Check | Exit | Result |
|---|---|---|
| `make -C apps/ios test` (shared simulator) | 0 | TEST SUCCEEDED: 137 tests (app-hosted 43, Models 21, Networking 45, Auth 19, Routing 4, Telemetry 5) |
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21) | 0 | 134 JVM tests, 0 failures; goldens verified; lint clean; debug APK built |
| `go test -count=1 ./internal/contract/`, `go vet` | 0 | ok |
| `vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` | 0 | 28 passed |
