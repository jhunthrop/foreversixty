# Mobile Wave 3 F14: widgets and live surfaces (2026-10-31)

Branch `wt/mobile-widgets`. One lane: both platforms and the API, from the same contract fixtures. Spec: `design/specs/2026-10-31-mobile-wave3.md` §5.6 (§9 and §5.6 amended by the build review). Boards: `design/mocks/2026-10-31-mobile-wave3/renders/f14-surfaces-{ios,and}.png`.

## API (W3-C12, W3-C13, W3-C14)

| | |
|---|---|
| `GET /v1/me/widget` | `api/internal/livesurface/widget.go`. One snapshot: the character (asked key, else main, else first), its guild, the next raid on the raids tab's calendar (`guilds.WidgetRaidFor`: cancelled ones included, `signed_up` = In + Late, the character's own state, raid spec and role, `can_sign_up`), the last raid, the Home answer's first three upgrades with its count and total (`HomeUpgrades` over `app.Service.HomeOf`, null while `APP_API` is off), running server runs. `generated_at` to the minute so the ETag holds between quick refreshes. Private, no-cache. |
| `POST`, `DELETE /v1/me/live-activities` | `live.go`, `store.go`. Migration **0063** `live_activity_tokens`. A raid only for a verified member of its guild and before its end (`guilds.LiveRaidEnd`); a sim needs `server_sims` (402) and your own server run (`sims.OwnsServerRun`). 30 a minute, 20 kept per account, re-register refreshes, swept hourly once expired; device rows go with the device. Tokens never logged. |
| Pushes | `payloads.go`: APNs Live Activity start/update/end (Unix-second content-state), the silent `widget.refresh`, FCM data messages. Pinned exactly to `surface.*` fixtures (named `surface.` because the `push.*` family is held to the alert-notification fields). |
| Switch | `APP_WIDGETS=true` mounts the routes (404 otherwise). |

Fixtures `widget.get`, `widget.get.no-guild`, `live.register`, `live.register.error-premium-required`, `surface.*`. OpenAPI, `docs/contracts/app-v1.md` "Live surfaces", generated Swift and Kotlin clients.

## Apps

| | iOS | Android |
|---|---|---|
| Target | `ForeverSixtyWidgets` WidgetKit extension in `project.yml` (XcodeGen), App Group `group.gg.foreversixty.app`, entitlements under the existing ad-hoc simulator signing (no Team ID); device token in the app group's keychain group for the sign-up intent | Glance widgets in the app module, DataStore snapshot |
| Next raid | small, medium (In/Late/Out App Intent answering in place, optimistic, `Not saved` 10 s), inline, rectangular, circular ring (`24` / `of 40`), circular countdown | 2x2, 4x2 (actionRunCallback buttons), keyguard category |
| Next upgrade | small, medium, rectangular | 2x2, 4x2, picker previews |
| Live | Raid countdown and server-sim Live Activities, Dynamic Island (ticking compact timer, seal, `38 of 40`) | `Raid starting` and `Server sims` ongoing notifications (chronometer, gold bar, seal icon), quick-settings tile |
| Refresh | foreground, after a sign-up, BGAppRefreshTask ~30 min, `widget.refresh` handler; cleared at sign-out | foreground, after a sign-up, WorkManager ~30 min, FCM payload parser; cleared at sign-out |
| Settings | You › Live surfaces, both on by default | same |

**One source (the wow-player SHOULD).** Every surface computes its words from the one snapshot with one clock through `WidgetFigures` (iOS `apps/ios/Widgets/Shared/WidgetFigures.swift`, Android `widgets/logic/WidgetFigures.kt`), held on both platforms to the golden vectors `apps/shared/widgets/figures.json` (countdown, tags, inline, signed-up, stale, the Live Activity window and title, upgrades). A test also checks the Live Activity's figures equal the widget's.

## Reviews

Round 1: ux-designer HOLD (6 MUSTs), wow-player HOLD (4 MUSTs). Round 2: both **SHIP** — `design/reviews/2026-10-31-mobile-wave3-widgets/ux-verdict-build.md`, `player-verdict-build.md`; captures in `build/` (iOS 55, Android 42, default and largest text). Rulings accepted: widget text caps (iOS largest non-AX size, Android 1.3x) with nothing truncated; ticking compact island timer; upgrade tap opens Home (no item deep link yet); `Open the app` when the Home answer is off.

## Gates

- Go: `go test -count=1 -p 1` guilds, sims, livesurface, db, contract, server, config, app, cmd/api: all ok, exit 0 (private DB `forever_test_mobile_widgets`).
- Codegen drift: `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` 35 passed, exit 0.
- iOS: `make -C apps/ios test` TEST SUCCEEDED (345 app tests plus the package suites).
- Android: `test`, `lint`, `:app:verifyRoborazziDebug`, `assembleDebug` each exit 0.

## Switching it on

`apps/shared/features.json` keeps `widgets` default false: the repo's codegen test holds every bundled default dark, and a feature turns on from the API. To turn it on: `APP_WIDGETS=true` and `widgets` in `APP_FEATURES` on the API deploy (with `APP_API=true` for the upgrades). Not done from this lane (production config).

## Gaps

1. **No remote live updates.** Apple's APNs key does not exist, and lane A2's senders send alert pushes only (a Live Activity needs `apns-push-type: liveactivity` and its own topic; FCM data messages need a data-only send). The tokens are stored server-side and swept; no scheduler produces `live.*` or `widget.refresh` yet. Until then the apps start and update the Live Activity and ongoing notification while they run, and widgets tick on their timelines (30-minute entries, hourly past 2 days).
2. **Android has no FCM client yet** (lane F7), so its live registrations are not sent; the payload parser and handler are in place.
3. **Device checks before release:** the system tint on a real lock screen and iOS Tinted home screen (captures are approximations), the inline widget with `Temple of Ahn'Qiraj`, the compact island width; the real Android tile and widget picker; the background refresh tasks.
4. **iOS snapshot goldens** were recorded on the local iOS 26.5 simulator; a CI record run (`make -C apps/ios snapshots`) should confirm them.
5. `invite_at` is null until raids carry an invite time; the countdown then says `until the pull at {time}` and `{zone} starts in {n} minutes`.
6. No item deep link, so an upgrade tap opens Home.
7. The `widget.get` fixture's last raid (3 Dec) predates the 9 Dec raid opening; move it before any store screenshot.
8. The raids lane's own Android notifications still use the launcher icon (blank square); F14's use the new seal icon.
