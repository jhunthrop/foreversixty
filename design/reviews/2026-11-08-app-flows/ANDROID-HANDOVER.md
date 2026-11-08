# Android build lane: handover

State of `apps/android` against `design/specs/2026-11-08-app-flows.md` (with `HANDOFF.md` beside this file), for the next Android lane.

## Done (on main)
| Step | What a player gets | Where |
|---|---|---|
| 1 Push | FCM on sign-in (`onRegistered`, firebase-messaging 26; `POST /v1/me/push-tokens`, gated by `notifications`); seven channels, one per group (Raids, Gear and sims, Logs, Guild ladder, Tiers, Officer, plus Account); RAID_SIGNUP shows In / Late / Bench / Out answered in the background; after Late, How late? 15 / 30 / 60 min sends `late_minutes`; the server's `confirmation` replaces the notification; taps land on their answer | `push/` (PushMessage, PushGroup, PushNotifier, PushRegistrar, ForeverMessagingService), `raids/push/` |
| 2 Navigation | Four tabs Home, Raid, Gear, Activity with badges (Activity: unread); the profile button (Battle.net portrait, faction ring, faction badge; person glyph signed out) on every tab root; the profile menu as a `ModalNavigationDrawer` (You, Switch character with the low-level fold, Tools, Account, footer); menu rows push their screen under the current tab | `navigation/ForeverAppScaffold.kt`, `navigation/ForeverTab.kt`, `menu/` |
| 5 Activity | The inbox (`GET /v1/me/activity`), filters, Today / Yesterday / This week / Earlier, unread dots, Mark all read, each row's game icon (`ActivityRow.icon`), each row landing on its answer | `activity/` |
| Links | App Links for every F6 site page (one autoVerify filter); each `DeepLink` site case routed in `MainActivity.handleAppLink`; `DeepLink.SITE_ONLY` pages and unknown links open the browser | `AndroidManifest.xml`, `MainActivity.kt` |

Parity tests: `menu/MenuParityTest.kt` (menu, menu-one, menu-low, menu-out) and `activity/ActivityParityTest.kt` (activity, activity-empty, activity-out), each with 1.5x / 2x golden-only captures, through `testing/ParityCheck.kt`.

## Next, in the spec's order
Note: the title `letterSpacing` token change (0em to 0.01em) arrives with land/ios-shell, which re-records the Android goldens it moves; merge main after it lands before starting.

3. **Home** in its six moments from `/v1/home` `sections` (head with the class crest, never a trigger; Next upgrades with the Best run line computed from the cached BiS and loot files; next raid card or line with the four answers and Late's chips; raid live; Last night; New for you = the three newest Activity rows with `exclude_kinds`; Find a guild), plus first run (Pick your main, then the notification ask in context: E6, F8). The old Home (`home/HomeScreen.kt`, `HomeAnswer.kt`) is still what Home shows; replace it.
4. **Raid**: next raid card, Needs you (officer), Not answered with the mentions, raid night with Ready check, the applicant. Raid badge: `GET /v1/guilds/{id}/raid-badge` (`TabBadges.raid`, now always 0).
6. **Gear**: summary line with Change, Where to go tonight, Every slot, the item sheet with `GET /v1/guilds/{id}/items/{item}/needs`. Gear still shows the lane F3 BiS list (`bis/`).
7. Glance widgets, LiveNotifier (the raid-night ongoing notification), the QS tile, restyled.
8. Restyle the secondary screens reached from the menu. Notifications has no in-app screen yet: the menu row opens Android's channel settings; build the mock's settings (the groups with `group` on the notification event rows, `raid_reminder` "HH:MM"), and give pushed account screens a back arrow.

Smaller open items: the menu's Notifications line (`MenuLines.notifications`) is empty until that screen exists; the test fixtures have no avatar painter, so portraits render as their ring in goldens.

## Gotchas
- **Toolchain**: `JAVA_HOME=/opt/homebrew/opt/openjdk@21 ANDROID_HOME=/Users/jh/Library/Android/sdk ./gradlew ...`; run only the touched tests (`:app:testDebugUnitTest --tests ...`, `:app:lintDebug`). Lint is strict (NewApi for `LocalDate.ofInstant`, `remember` returning Unit, unused resources, split `<data>` tags).
- **Goldens are recorded on CI, never locally**: push a non-land branch, `gh workflow run android.yml --ref <branch> -f record=true`, then `gh run download <run> -n android-snapshots -D apps/android/app/src/test/screenshots` and commit. Running `recordRoborazziDebug` locally is fine for seeing region diffs, but `git restore` the PNGs before committing. Text goldens (`.txt`) are renderer-independent.
- **android.yml sparse checkout** left out `design/mocks/`; the parity tests read `design/mocks/app/regions/`, so the checkout now adds that folder back and a change to `android-*.json` runs the Android checks.
- **No drift entries** (coordinator rule, same as iOS): fix the measurement or the fixture instead. A scrolling row counts its whole content (`ParityScrollRegion`). Text that sits on a column's edge wraps differently in Chromium and Skia: move the mock text (`design/mocks/ws_app_flows.py`) off the edge and re-render the regions on the runner (push a `mocks/**` branch or dispatch `app-mocks.yml`, then `gh run download <run> --name app-mocks --dir design`). Never render regions on a Mac.
- **Region inventory**: the portrait and its faction badge are two images (the mock's two `img`s); a portrait without an avatar is one (the badge).
- **Copy**: only `apps/shared/copy/en.json` is edited by hand (keys under the spec's screen ids: `home.*`, `raid.*`, `gear.*`, `activity.*`, `menu.*`, `push.*`, `first.*`), wording verbatim from the mocks, no em dashes; run `node web/scripts/gen-mobile-copy.mjs` (it also writes the iOS files). On a merge conflict in a generated file, take the union of both sides' keys in `en.json` and re-run the generator. Reuse a key the iOS lane added instead of a near-duplicate.
- **Tokens, models, links and features are generated** (`Tokens.kt`, `ApiModels.kt`, `ApiEndpoints.kt`, `DeepLink.kt`, `Feature.kt`): ask the coordinator for API or link changes.
- **google-services.json** is never committed (git-ignored; CI writes it from `GOOGLE_SERVICES_JSON`). For a local build with push, fetch it with the Firebase Management API (`GET https://firebase.googleapis.com/v1beta1/projects/foreversixty/androidApps/{appId}/config`, `x-goog-user-project: foreversixty`); one config carries both packages.
- **Shared scratchpad**: the session scratchpad is shared with the iOS lane; prefix every file there with `android-lane-` (or keep commit messages in a file in your own worktree).
- **Launch mode** skips `android.yml` on `land/**` pushes, so a landing does not run the Android tests: run the CI record on a non-land branch first and land only a tree whose record run is green.
- **Landing**: only one `land/*` branch at a time; merge origin/main, push `land/android-<step>`, ask the coordinator for the go, then `tools/land.sh land/android-<step>`. The repo's commit hook rejects some flag spellings on `git commit`; plain `git commit -m` works.
