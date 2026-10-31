# Mobile Wave 1, lane F1: sign-in, account, character switcher (2026-10-31)

Branch `wt/mobile-f1-signin`, from origin/main, with origin/main (5d86a8dc7, the shared device lock) merged in. Not pushed. No migration, no API handler.

## Fix round 1 (2026-10-31, both reviewers HOLD)

The verdicts are `design/reviews/2026-10-31-mobile-signin/ux-verdict.md` and `player-verdict.md`, and they are committed with this round. The coordinator's six items:

1. **Region and ruleset on every character.** Forever has no realms: a character is its region, ruleset and name.
   - Switcher rows read `Obnoxious Yell · US · PvP`, with the race, spec, class and level below it and the guild as secondary text.
   - You gains a Character row in the same form.
   - The ruleset words are the site's own (`web/src/lib/characters.ts`: Normal, PvP, Roleplay, Hardcore), stored as `account.ruleset.*`. An unknown ruleset is shown as the API sends it.
2. **Delete account says what a paying player loses.** `DeletionConsequence.of(me)` (iOS Auth, Android core:auth) reads `/v1/me` and shows only the lines that apply:
   - an active Premium plan is cancelled;
   - the founder price is gone for good;
   - a guild plan the account pays for (`guilds[].plan`, scope `plan`, `you_are_billing_contact`) is cancelled and its seats end for every member;
   - a seat held from a guild ends.

   The lines appear under "Billing and seats" and again in the confirmation. The confirm button now reads `Delete account`.

   To read the seat, `seat` is now typed in the spec as `MeSeat` (it was open JSON). New fixtures `me.guild-payer` and `me.seat` are claimed by `getMe`. Each case is tested on both platforms.
3. **Guild reports. PENDING THE OWNER'S DECISION.**
   - One key, `account.delete.reports`: "Guild reports stay for the raiders in them. Your name on them becomes Former raider." This is the coordinator's recommended default.
   - The logs item now reads only "The combat logs you uploaded", so whichever way the owner decides, the change is a one-line edit to that key.
   - The handler note below must follow the decision.
4. **You (F8.1).**
   - Privacy and Terms open foreversixty.gg/privacy and /terms in the system browser, with the external glyph.
   - About Forever Sixty pushes the seal, the version and the site footer's disclaimer. `about.notAffiliated` is now byte-equal to `web/src/lib/footer-copy.ts`.
   - The version line reads `Version {v} ({build}) · data {build}`. The version and build come from apps/VERSION through P2's `ClientIdentity`; the data build comes from the app config's `active_build`. Before the config answers, the line falls back to `account.version.noData`.
   - Notifications, Characters and Discord wait for their lanes.
5. **The deletion notice is neutral.** It uses the line border and a gold check, not wipe red, and reads "Your account is deleted. Everything is removed by 29 Nov 2026." Red stays for F1.4 and the 401 line.
6. **Tab bar: minimal additive approach.** F2's shell fix and slot API have not landed: `wt/mobile-f2-home` touches only its own Home case.
   - My edits to the tab label and icon are reverted on both platforms; the tab reads "Account" again until F2 lands.
   - The case is renamed to the name F2 will use: iOS `AppTab.you` (was `account`), Android `ForeverTab.You` (was `Account`). The spec's names are `home`, `raids`, `bis`, `you`, in Kotlin's PascalCase on Android.
   - My only other tab-setup lines are the `.home` / `.you` cases that plug in `HomeAccountGate` and `YouScreen` / `YouRoute`. Merging F2's four-tab shell should only touch those lines.

Reviewer shoulds also taken:
- Delete account moved below Sign out as a quiet red text row.
- The plan line is meta under Plan, with the Founder tag on the right.
- Revoke has a 44 pt hit target.
- Companions are drawn with a monitor and phones with a smartphone glyph on Android (vector drawables, since `material-icons-core` has neither).
- The confirmation body is "Your characters, logs and sims are removed within thirty days. Every device is signed out now."
- The timing line now says "Everything listed is deleted within thirty days".
- The iOS switcher sheet title stops growing at AX2.

Left for the owner or a later lane:
- Sign-in copy suggestions (read-only Battle.net, "Pick your class to see…", "Forever lists a character…"). These change spec-owned keys and need the spec author.
- The Premium tag on the account head (F8).
- Whether signing in again during the 30 days restores the account. This is an API decision; it is added to the handler note.
- An inset-aware Android capture and iOS alert captures. These need the shared emulator (no system image installed on this Mac) or a UI-test target.

New captures: `delete-free`, `delete-premium`, `delete-guild-payer`, `delete-seat`, `delete-failed`, `delete-confirm` (Android, Premium founder), `about-default`, plus refreshed `you-*`, `switcher-*`, `devices-*` and `signedout-deleted`. The old `delete-default` and `delete-plan-failed` are removed.

## What was built

Both platforms, the same states and the same test list.

| Spec | iOS | Android |
|---|---|---|
| F1.1 signed out (Home, and You when signed out) | `App/Account/SignedOutView.swift` | `account/SignedOutScreen.kt` |
| F1.2 system browser sign-in, spinner, cancel | `ASWebAuthenticationSession` (`SystemWebAuthenticator`), `SessionStore.signIn` | Custom Tab, App Link redirect into `SessionManager.handleRedirect`; a tab closed with no redirect returns to F1.1 on `onResume` |
| F1.3 import, and the no-characters state | `ImportView`, `NoCharactersView` | `ImportScreen`, `NoCharactersScreen` |
| F1.4 sign-in failed toast | `NoticeToast` | `NoticeToast` |
| F1.5 selector and switcher sheet | `CharacterSelector`, `SwitcherSheet`, `.characterSelector()` toolbar modifier | `CharacterSelector`, `SwitcherSheetContent`, `CharacterSelectorWithSheet` |
| You: account head, Plan, Devices, Delete account, Sign out | `YouView` / `YouScreen` | `YouScreen` / `YouRoute` |
| F1.6 Devices with Revoke | `DevicesView` / `DevicesScreen` | `DevicesScreen` / `DevicesRoute` |
| In-app account deletion | `DeleteAccountView` / `DeleteAccountScreen` | `DeleteAccountScreen` / `DeleteRoute` |

- **Session logic.** It lives in the auth module on each side: iOS `Sources/Auth/Session/`, Android `core/auth/.../auth/session/`.
  - The phases are signed out (with an optional notice), signing in, importing, and signed in.
  - The current character follows spec §3.4: the one picked on this device, else the main, else the most recently synced, else the first.
  - The device id from the exchange is kept, so Devices can mark This iPhone / This phone.
  - Sign out calls `POST /v1/auth/app/logout`, then clears the token and caches.
- **401 handling.**
  - `SessionTokenSource` (iOS) and `SessionTokens` (Android) wrap the token store and report every 401 discard.
  - Any request answered `401 unauthorized` or `401 session_expired` clears the token and shows signed out with `You were signed out. Sign in again.`
  - Both codes are tested with their fixtures.
- **403 `browser_session_required`.** Revoke and delete return `needsBrowser`, and the view opens `https://foreversixty.gg/account` in the browser.
- **Sign-in failure lines.** These are the copy catalog's three, not the API's raw text:
  - `invalid_grant` gives `signin.error.expired`;
  - no connection gives `signin.error.offline`;
  - a state mismatch, a callback error or any other exchange failure gives `signin.error.bnet`;
  - cancelling shows no line.
- **Fixture mode signs in for real against the fixtures.**
  - iOS `FixtureWebAuthenticator` and Android `FixtureBrowser` answer at once with the callback, so the state check, the exchange and the import all run.
  - iOS fixture mode still starts signed in (the scaffold's `fsd_fixture`), and seeds the device id from `app.signin.exchange.json`.
  - Android fixture mode starts signed out; one tap signs in.
- **Feature gate.** `sign_in` is on in debug and fixture builds:
  - iOS `AppConfigStore.isAvailable(_:mode:isDebugBuild:)`;
  - Android `AppConfigRepository.isAvailable(...)`.
  - A live release follows the config. With the gate off, Home and You fall back to the scaffold.
- **Crests and seal.**
  - Bundled from `web/public/icons/hd/crests` and `web/public/seal-400.png`: iOS through `project.yml` resources, Android through a new `forever.brand-art` Gradle plugin that copies them into generated res.
  - Nothing is checked in twice.
- **Type.** iOS `DesignSystem/ScaledType.swift` holds the spec §4.2 roles with `relativeTo`. Android `AccountType` uses `sp`, and the selector name stops at 22 sp.

## Contract changes (`api/openapi.yaml`, fixtures, `docs/contracts/app-v1.md`)

- **`DELETE /v1/me` (`deleteAccount`).**
  - Body `AccountDeletionRequest {confirm: "delete my account"}` (an enum of one value).
  - `202 AccountDeletion {requested_at, complete_by}`.
  - Fixtures `me.delete` (202), `me.delete.error-invalid` (400) and `me.delete.error-unauthorized` (401).
- **`GET /v1/devices`** gains the `devices.list` fixture: this iPhone, a Pixel 9 and a Windows companion.
- **`DELETE /v1/devices/{id}`** is now the app operation `revokeDevice` (204, plus the `devices.revoke.error-not-found` 404 fixture).
- **`error-session-expired`** joins `GET /v1/me`'s 401 fixtures.
- Models, endpoints, fixture routes and both string catalogs are regenerated.
- The Go contract tests pass on a private database, `forever_gc_f1`, created for the run and dropped after it.

## For the API lane: what the `DELETE /v1/me` handler must do

Today deletion is by emailed request. The privacy page (`web/src/content/pages/privacy.md`, "Requests and contact" and "Data retention") promises two things:
- "To see, export, correct or delete what the site holds about you, email help@foreversixty.gg ... You get an answer within thirty days, and usually much sooner."
- Uploaded logs, the reports parsed from them, saved sims and the ratings derived from them "are kept until you delete them or your account".

The site's `/account` has no deletion entry. The in-app copy (`account.delete.*`) repeats that promise and nothing more. The handler must match it:

1. **Auth and confirmation.**
   - Accept a session or an app device token: App Store 5.1.1(v) needs it in the app, so this is not a `RequireBrowserSession` route.
   - Answer 400 `invalid` with `fields.confirm` unless `confirm` is exactly `delete my account`.
   - Rate-limit it like the other writes.
2. **Sign out everywhere at once, before answering.**
   - Revoke every device (app and companion) and delete every session row.
   - The app's token must answer 401 from then on. The app clears its token itself on the 202 and does not call logout.
3. **Cancel billing at once.**
   - Cancel the Stripe subscription now, with no further renewal, and delete or detach the Stripe Customer. The refund policy applies to the current period.
   - If the account holds a guild claim or guild plan, cancel or release it the way the site's release flow does.
   - Revoke seats it gave out. Those players fall back to free.
   - The app shows `Your plan is cancelled and does not renew.` only when `entitlements.billing` is set.
4. **Remove the data by `complete_by`** (`requested_at` + 30 days). This is the privacy page's list:
   - the Battle.net identity and email;
   - linked characters, their addon exports and `character_syncs` history;
   - uploaded combat logs and the reports parsed from them;
   - saved sims and builds;
   - the ratings and ranking rows derived from them;
   - notification settings and push tokens (lane A2's tables);
   - the Discord link and its webhooks (the notifications copy already promises "we ... delete it with your account");
   - the user row itself.
   - Guild reports that list other raiders: pending the owner. The app's default copy (`account.delete.reports`) says reports stay and the deleted player's name becomes "Former raider". If that stands, rewrite the player's name on report rows to "Former raider" and keep the reports; then align the privacy page ("kept until you delete them or your account") and the account page.
5. **Idempotent.** A repeat call while a deletion is pending answers 202 with the same two dates.
   - Decide whether a Battle.net sign-in during the 30 days restores the account or starts a fresh one, and tell the app lane so the screen can say it.
6. **Contract test.** Add `app_contract_test.go` beside the handler against the three fixtures, and keep it dark until after Nov 6 like the other API lanes.

## Verification (each command on its own, exit code checked)

| Check | Exit | Result |
|---|---|---|
| `make -C apps/ios test` (shared simulator, after the merge) | 0 | TEST SUCCEEDED. Suites: app-hosted 24 (account snapshots 16), Auth 40 (session 21), Networking 35, Models 21, Routing 4, Telemetry 5 |
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21) | 0 | All JVM tests, `SessionManagerTest` (16) and 19 account goldens verified, lint clean, debug APK built |
| `go test -count=1 -p 1` contract, server and auth (private DB `forever_gc_f1`, dropped) | 0 | ok |
| `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` | 0 | 28 passed. `web/node_modules` was symlinked for the run and removed before every commit |

Nothing under `web/` was edited, so `npm run check`, `lint` and `format:check` were not needed.

## Captures

`design/reviews/2026-10-31-mobile-signin/`, as `<screen>-<state>-ios.png` and `-and.png`. The iOS images are snapshot renders hosted in the app on the shared iPhone 17 simulator, at 402 x 874. The Android images are Roborazzi renders at Pixel 9 size.

| State | Files |
|---|---|
| F1.1 | `signedout-default` |
| F1.2 | `signin-browser-open` |
| F1.3 | `import-default`, `home-no-characters` |
| F1.4 | `signin-failed` |
| The 401 line | `signedout-401` |
| After a deletion | `signedout-deleted` |
| F1.5 with three characters | `switcher-default` |
| You | `you-free`, `you-premium`; `you-signout-confirm` (Android) |
| F1.6 | `devices-default`, `devices-revoke-failed`; `devices-revoke-confirm` (Android) |
| Delete account | `delete-default`, `delete-plan-failed`; `delete-confirm` (Android) |
| Largest text (AX5 / 2.0) | `signedout-ax5`, `switcher-ax5` |

## Concerns

1. **System UI is not in the captures.**
   - The iOS confirmations (Revoke, Sign out, Delete) are SwiftUI `.alert`s and `.confirmationDialog`s, which snapshots cannot render. The Android captures show the same three dialogs with the same copy.
   - The F1.2 system consent sheet and the Custom Tab need a live start URL and a universal link / App Link, so the capture shows the button's spinner state instead.
   - No Android emulator capture. `tools/mobile-device.sh android-run` could not create the shared AVD `foreversixty-test` on this Mac.
     - First it needed `JAVA_HOME` for avdmanager.
     - With that set, avdmanager answered "Package path is not valid": no Android system image is installed.
     - No AVD was created and no emulator was left running.
     - Installing the system image and creating the shared AVD is the owner's or coordinator's job, after which `tools/mobile-device.sh` works. Until then the Android captures are the Roborazzi renders only.
2. **Shared-file touches.** All are additive, but other lanes may collide:
   - `RootView.swift`: You title and icon, the session in the environment, the Home and You cases.
   - `AppEnvironment.swift`: `session` and the `account:` parameter.
   - `ForeverAppScaffold.kt` and `ForeverTab.kt`: the You label.
   - `MainActivity.kt` and `AppGraph.kt`.
   - `project.yml` and the Gradle catalog / build-logic.
   - Two helpers every lane will want, which another lane might also write: `AppConfigStore+DevBuilds.swift` / `FeatureAvailability.kt` (feature on in debug and fixture builds), and `ScaledType.swift`. `ScaledType.swift` adds `@_exported import DesignTokens` to DesignSystem so App code can read `FSColor`.
3. **The selector is not placed on Home, Raids or BiS yet.** Those roots belong to F2, F4 and F3.
   - iOS: add `.characterSelector()` inside their `NavigationStack`. It reads `SessionStore` from the environment.
   - Android: put `CharacterSelectorWithSheet` in the `TopAppBar` title slot.
   - `Pick your class` switches to the BiS tab; F3 should open its picker sheet there.
   - Raids' signed-out panel is `SignedOutView(variant: .raids)` / `SignedOutVariant.Raids`.
4. **Logout versus the spec.** Sign out follows the A1 contract, `POST /v1/auth/app/logout`, not the spec's `DELETE /v1/devices/self`. Unregistering the push token on sign-out is a TODO for A2 and F7.
5. **Caches.** Sign-out clears `URLCache` (iOS) and the account values. Feature lanes that add disk caches must clear them on sign-out too, through the `clearCaches` hook on `SessionStore` / `SessionManager`.
6. **Universal links and App Links** are still TODO, waiting on the owner's Apple Team ID and SHA-256 fingerprint (the existing TODOs in `project.yml`, `WebAuthenticator.swift` and `AndroidManifest.xml`). Live sign-in cannot complete on either platform until then.
7. **Android pending sign-in is in memory.** If the process dies while the Custom Tab is open, the redirect is rejected and shows the Battle.net line. Persist the verifier and state if that proves common.
8. **Icons.** The Android device rows use the Material `Phone` (handset) glyph, because `material-icons-core` has no smartphone icon. iOS uses `iphone`.
