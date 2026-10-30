# Forever Sixty apps: privacy answers for the stores

The answers for Apple's App Privacy labels (App Store Connect > App Privacy) and Google Play's
Data safety form (Play Console > Policy > App content). They describe both apps as built, and
they must change in the same commit as any change to what the apps collect or send.

Lane P2 wrote the sections on crash reporting and client identity. The account rows
describe what sign-in (lane A1) and the API already do. Push tokens (lane A2, F7) are marked
as pending until that lane ships. Wave 2's store-listing lane fills in the forms from this file.

## What the apps send, and to whom

| Data | Sent to | Why | Linked to the player? |
|---|---|---|---|
| Battle.net account id and BattleTag, from sign-in | Our API | Sign-in, the account screen | Yes: this is the account |
| Email address, only if the player added one on the site | Our API (shown, never collected by the app) | The account screen | Yes |
| Character names, realms, classes, gear and talents | Our API | The app's purpose | Yes |
| Device name (for example "iPhone"), at sign-in | Our API | The account's device list, so a player can revoke a device | Yes |
| Device bearer token (`fsd_…`) | Our API | Authentication | Yes |
| `X-App`: platform, app version and build | Our API, every request | Usage and error counts per version, and the forced-update gate | No |
| `X-App-Install`: a random UUID made on first launch | Our API, every request | Counting installs per version without any user id | No (see below) |
| Crash and hang reports | Sentry | Fixing crashes and freezes | No |
| Push token | Our API | Notifications (**pending lane A2 and F7**) | Yes |

**The install id is not linked to the player.** It is generated on the device, stored where an
uninstall deletes it, and excluded from backups. iOS keeps it in UserDefaults, not the
Keychain. Android keeps it in `forever_install`, excluded in `backup_rules.xml` and
`data_extraction_rules.xml`. The API logs it with the request's platform and version, never
with the account. It is not an advertising identifier, and nothing joins it to one.

## Crash and hang reporting (Sentry)

**Vendor: Sentry (sentry.io).** Both apps use it: iOS through `sentry-cocoa`, Android through
`sentry-android-core`. There is no NDK, replay or performance module. Lane P2 chose Sentry over
Firebase Crashlytics for these reasons:

- Sentry reports iOS app hangs. Crashlytics does not.
- Sentry's project setting stops IP storage. Crashlytics has no such setting.
- iOS needs no Firebase SDK at all. Its push goes to APNs directly.
- One vendor can later also cover the site and the Go API.

**When it runs.** Only in a release build that talks to the live API and was built with a DSN.
Debug builds, fixture mode and the empty placeholder DSN never start the SDK
(`CrashReportingPolicy`, tested on both platforms). Android's automatic initialisation is
turned off in the manifest.

**What a report contains:**
- The stack trace and the exception type and message, scrubbed.
- The app version and build.
- The OS version and device model.
- Free memory and disk, and whether the app was in the foreground.
- Breadcrumbs of screen changes and network calls. Each URL keeps only its scheme, host and
  path. Each message is scrubbed.
- Sentry's own random install id, which is regenerated on reinstall. It feeds the crash-free
  user count.
- Session start and end markers, which give the crash-free rate.

**Never in a report:**
- Account id, BattleTag or email. The apps never call `setUser`, and `beforeSend` drops every
  user field except Sentry's install id.
- IP address:
  - Both SDKs run with `sendDefaultPii` off, so they send `infer_ip: never`.
  - The project also has **Prevent Storing of IP Addresses** turned on (see the settings
    below).
- Device name, server name, request bodies or headers.
- Screenshots or the view hierarchy. Both are off.
- Query strings and fragments. These carry the sign-in callback's `code` and `state`, and both
  are stripped.
- Device tokens (`fsd_…`), bearer credentials, and any `code=`, `state=`, `token=`,
  `code_verifier=` or `key=` value in free text. These are replaced with `[redacted]` (the
  `Scrubber` on each platform).
- Performance traces or profiles. Sampling is zero.
- Any advertising identifier. The apps have no ad SDKs.

**Sentry project settings** (Settings > Security & Privacy). Set these once, on both projects,
`forever-ios` and `forever-android`:
- Prevent Storing of IP Addresses: **on**.
- Data Scrubber: **on**. Use Default Scrubbers: **on**.
- Additional sensitive fields: `code`, `state`, `code_verifier`, `token`, `authorization`.
- Data retention: the plan default, 90 days.

## Apple: App Privacy labels

**Tracking:** No. The apps do not track. They have no IDFA, no ad SDKs and no data brokers, so
there is no App Tracking Transparency prompt.

Data types collected:

| Category > type | Collected | Linked to user | Used for tracking | Purposes |
|---|---|---|---|---|
| Identifiers > User ID (Battle.net account id, BattleTag) | Yes | Yes | No | App Functionality |
| Contact Info > Email Address | No (the app shows an email only if the player added one on the site; it never asks for one) | n/a | n/a | n/a |
| User Content > Other User Content (characters, gear, talents, raid sign-ups) | Yes | Yes | No | App Functionality |
| Identifiers > Device ID (the random `X-App-Install` id and Sentry's install id) | Yes | No | No | Analytics, App Functionality |
| Usage Data > Product Interaction (requests per app version) | Yes | No | No | Analytics |
| Diagnostics > Crash Data | Yes | No | No | App Functionality |
| Diagnostics > Performance Data (app hangs) | Yes | No | No | App Functionality |
| Diagnostics > Other Diagnostic Data (OS version, device model, memory) | Yes | No | No | App Functionality |

Not collected: location, contacts, health, financial info, purchases (there is no in-app
purchase), browsing history, search history, sensitive info, photos, audio.

The push token row is **pending lane F7.** Apple does not list a push token as a data type.
If notification preferences are tied to the account, add nothing; review this when F7 ships.

## Google Play: Data safety form

**Overview answers:**
- Does the app collect or share any of the required user data types? **Yes.**
- Is all user data encrypted in transit? **Yes.** HTTPS only. The network security config
  allows no cleartext traffic.
- Do you provide a way for users to request that their data be deleted? **Yes.** Today a
  player emails help@foreversixty.gg, as https://foreversixty.gg/privacy ("Requests and
  contact") says. That page is the deletion URL the form asks for.
  - **Open item:** Apple's guideline 5.1.1(v) requires an app that creates accounts to let the
    player start deleting the account inside the app.
  - Battle.net sign-in creates one, so lane F1's account screen needs a "Delete account" entry
    before App Review.
  - A link to the request flow is acceptable if there is no in-app deletion API.
- Shared with third parties? **No.** Sentry is a service provider processing on our behalf, so
  under Play's definition it is not "sharing".

**Data types:**

| Data type | Collected | Shared | Ephemeral | Required or optional | Purposes |
|---|---|---|---|---|---|
| Personal info > User IDs (Battle.net account id, BattleTag) | Yes | No | No | Required to sign in; the app has signed-out screens | App functionality, Account management |
| Personal info > Email address | No (shown, never collected) | n/a | n/a | n/a | n/a |
| App activity > Other user-generated content (characters, gear, talents, raid sign-ups) | Yes | No | No | Required for those features | App functionality |
| App activity > App interactions (requests per app version) | Yes | No | No | Required | Analytics |
| App info and performance > Crash logs | Yes | No | No | Required | App functionality (stability) |
| App info and performance > Diagnostics (ANRs, OS version, device model, memory) | Yes | No | No | Required | App functionality (stability) |
| Device or other IDs (the random `X-App-Install` id and Sentry's install id) | Yes | No | No | Required | Analytics, App functionality |

Not collected: location, financial info, health, messages, photos, audio, files, calendar,
contacts, web browsing, installed apps.

Push tokens are **pending lane F7.** When F7 ships, add Device or other IDs > "Firebase
installation / FCM token" for App functionality.

## Keeping this true

- A new SDK, or a new field sent to any vendor, changes this file in the same commit.
- `CrashReportingPolicy` and `Scrubber` are the code behind these answers. Their tests are the
  check: `Tests/TelemetryTests` on iOS and `core/telemetry` on Android.
- The runbook (`docs/runbooks/mobile-apps.md`) says where crash reports live and who can read
  them.
