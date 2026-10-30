# Native mobile apps: the lane plan

Owner, 2026-10-30: "we'll run both builds simultaneously". iOS (SwiftUI) and Android (Kotlin/Compose) are built in parallel, both native, on the existing API (decision 2026-10-24, no web wrapper). The Play Console account exists; Apple Developer enrollment is pending.

## What already exists

- **Device bearer tokens.** `auth/session.go` already authenticates `Authorization: Bearer fsd_…` device tokens on every `/v1` route, through the `devices` table the companion pairs with. App sessions reuse this; they need no new session type.
- **Static data.** BiS, talents, items, icons and guides live under `/data/<build>/…` on the CDN. Both apps read these directly, the same files the site reads.
- **The JSON API.** It serves the character, sims, guilds, raids, reports, builds and `/v1/me`. Only the guild centre, crest and addon contracts are written down in `docs/contracts`.
- **Notifications.** `api/internal/notify` fans out to Discord only.

## Waves and lanes

Lanes inside a wave run at the same time. A wave's later lanes wait only on the gates named.

### Wave 0: foundations (starts now, all parallel)

| Lane | Owner | Delivers | Blocks |
|---|---|---|---|
| **O. Accounts and keys** | Owner | Apple Developer enrollment as Commish LLC (D-U-N-S); Play Console app record with Play App Signing; a Firebase project for FCM; an APNs auth key (.p8) and an FCM service account, both into GSM; the app's sign-in redirect registered with Battle.net. | Push in F7, store builds in Wave 2. Never blocks scaffolding. |
| **D. Experience spec** | ux-designer, then wow-player review | One app spec for both platforms: one information architecture with native idioms (tab bar / navigation bar, Material 3 navigation). Covers the v1 screens, states, copy and notification types. Mocks for both platforms at phone sizes. Lives in `design/specs/2026-10-30-mobile-apps.md` and `design/mocks/2026-10-30-mobile-apps/`. | Feature UI in Wave 1. Scaffolds, tokens and API do not wait. |
| **A1. App sign-in** | API | `GET /v1/auth/app/start`: Battle.net OAuth with PKCE, opened in the system browser (ASWebAuthenticationSession, Custom Tabs). The redirect lands on the app's link with a one-time code. `POST /v1/auth/app/exchange` swaps the code for a device token with `platform=ios|android`. `/account` lists app devices with Revoke. Tests cover replay, expiry and a stolen code without the verifier. | F1 |
| **A2. Push and channel-agnostic events** | API | Builds the lifecycle catalogue below as typed events, with the producers wired at each source (raid save and roster, signups, sims, reports, ranking, addon sync, nightly BiS) and the scheduler for the time-based events. Adds a `push_tokens` table with register and unregister endpoints (device token required). Notifications become events (`raid.reminder`, `raid.posted`, `sim.finished`, `report.ready`) routed to channels: Discord as today, plus push. APNs and FCM senders sit behind one interface, with a fake in tests. Each user sets preferences per event and channel. Ships dark: no sender is enabled until lane O's keys exist. | F7 |
| **C. Contracts and fixtures** | API | `docs/contracts/app-v1.md` for every endpoint the apps call, with JSON fixtures under `docs/contracts/fixtures/app/`. A Go contract test pins each live handler's response shape to its fixture, so either app can build against fixtures before or without the API. | Both apps' networking layers. |
| **T. Tokens and assets** | Shared | A generator from `web/src/styles/tokens.css` to `DesignTokens.swift` and `Tokens.kt` (colour, type scale, radius, spacing), checked in CI so the three never drift. Fonts are Barlow, Barlow Condensed and Cinzel for the wordmark only, all OFL. The HD class crests and app icon come from `design/logo`. Item icons load from the CDN. | Feature UI polish. Scaffolds start with placeholders. |
| **I. iOS scaffold** | iOS | `apps/ios` with SwiftUI and iOS 26+. A typed API client from lane C's contracts. Fixture mode (no network) for previews and tests. Keychain storage for the device token. Swift Testing plus snapshot tests. A GitHub Actions macOS job builds and tests on the simulator. Needs no Apple account until signing. | Wave 1 iOS lanes |
| **N. Android scaffold** | Android | Install JDK 21 and Gradle on this Mac. `apps/android` with Compose, Material 3 and minSdk 28. kotlinx.serialization and a Ktor client from lane C's contracts, with a fixture mode. EncryptedSharedPreferences or DataStore with Tink for the token. Unit and Compose screenshot tests. A CI job builds, tests and uploads to the Play internal track through a Play Developer API service account. | Wave 1 Android lanes |

Gate to Wave 1: lane D at SHIP from both personas, lane C's contracts merged, and I and N building green in CI.

### Wave 1: features (one iOS lane and one Android lane per feature, at the same time)

Each feature is independent. Up to four features run at once per platform, so eight lanes are live together. Each pair works from the same spec section and the same contract fixtures, and each pair's captures are reviewed side by side.

| Feature | Needs | Scope (subject to lane D's spec) |
|---|---|---|
| **F1. Sign-in, account, character switcher** | A1 | Battle.net sign-in; the header's character selector as the app's switcher; the plan row; sign out; the device list. |
| **F2. Your character** | C | The home answer: upgrades for your level range, worn vs BiS with icons, quality colours and item tooltips (tap), talents status, last sync. |
| **F3. BiS lists** | C, static data | Class, spec, range and faction picker; the list with sources and runners-up. Read-only and cached offline. |
| **F4. Guild and raids** | C | Guild home, raids list, raid detail with roster and spec icons, sign up In / Late / Bench / Out, reminders. Likely the flagship mobile use. |
| **F5. Sims** | C | Your sims and results, server-run progress, rerun (Premium). |
| **F6. Talents (read)** | C, static data | Guide builds and your build on the trees, with the band compare. |
| **F7. Notifications** | A2, O keys | Permission prompt at the right moment, token registration, preferences, deep links into F4, F5 and reports. |
| **F8. Premium** | C | Show the plan and what it unlocks. No in-app purchase (decision P). Where the store allows a link to web checkout (the US storefront today, read from StoreKit / Play Billing), a See plans button opens foreversixty.gg/premium in the browser. Everywhere else there is no purchase prompt or call to action at all. The regions that may link come from the API (`GET /v1/app/config`), so a ruling change turns the link off without an app update. |

Order when capped at four per platform: F1, F2, F4 and F3 first; then F5, F6, F7 and F8. F7 starts as soon as lane O's keys land.

Every feature lane follows the house rules:
- Design before build.
- Tests first for logic.
- Captures on the iPhone 17 simulator and a Pixel 9 emulator at the default and largest text sizes.
- ux-designer and wow-player SHIP on the captures before merge.
- Direct push to main.

### Wave 2: release (parallel per platform)

- **Store listings.** Following the `design/listing` pattern: App Store and Play screenshots, descriptions, the App Privacy labels and the Play Data safety form, and the content rating.
- **Accessibility and performance.** VoiceOver and TalkBack passes; Dynamic Type and font scale; cold-start and scroll budgets measured on device.
- **Beta.** Play internal, then closed testing, available now. TestFlight once Apple enrollment clears. A beta group from the Discord.
- **Release.** Play production first if Apple is still pending; the App Store follows.

## The player's lifecycle and its notifications

Owner, 2026-10-30: push for different events, such as "Raid starting soon" and "Raid signups posted", and raid sign-up in the app instead of Discord. "Think about the full lifecycle of a player."

Raid sign-up in the app calls the same `PUT /v1/guilds/{id}/raids/{raid}/signups/{region}/{ruleset}/{name}` the site uses, so one roster serves the site, the app and the Discord post. A sign-up from the app refreshes the guild's Discord raid post the same way a Discord press does; lane A2 verifies this and closes any gap.

Every event below is one typed event in the API. Each player chooses its channel per event: push, Discord DM, both or off. **On** means on by default.

| Stage | Event | Who | Default | Action from the notification |
|---|---|---|---|---|
| Getting started | `character.linked`: the addon or companion synced a character for the first time | player | on | Open your character |
| Getting started | `guild.approved`: an officer approved your character to the roster | player | on | Open the guild |
| Levelling | `band.entered`: a character crossed into a new level range (20, 30, 40, 50, 60) | player | on | Open the new BiS list |
| Levelling | `bis.changed`: the nightly ranking changed a pick in your range for a slot you have not filled with it | player | off (weekly digest) | Open the list |
| Raids | `raid.posted`: signups opened for a new raid | guild members | on | **In / Late / Out** buttons, then Open |
| Raids | `raid.signup_nudge`: 24 h before a raid you have not answered | guild members | on | **In / Late / Out** |
| Raids | `raid.changed`: time moved or raid cancelled | signed-up players | on | Open |
| Raids | `raid.roster_set`: the officer set the roster; you are In or Bench | signed-up players | on | Open |
| Raids | `raid.starting`: starting soon, at the guild's reminder lead (60 min default) | In and Late players | on | Open; iOS Live Activity countdown (post-v1) |
| Raids | `loot.awarded`: an item was awarded to you | player | on | Open loot |
| After the raid | `report.ready`: the night's log is processed | raiders in it | on | Open report |
| After the raid | `coach.ready`: your coach notes for the night (Premium) | player | on | Open coach |
| After the raid | `ranking.best`: a new personal best on a fight | player | on | Open ranking |
| Sims | `sim.finished`: a server run finished (Premium) | player | on | Open result |
| Officers | `roster.pending`: characters waiting for approval (batched) | officers | on | Open roster |
| Officers | `raid.short`: under the raid's minimum 24 h out | officers | on | Open raid |
| Officers | `raid.late_change`: a raider changed to Out inside 24 h | officers | on | Open raid |
| Account | `premium.seat`: your guild gave you a Premium seat | player | on | Open account |
| Account | `billing.failed`: a payment failed | payer | on | Open billing on the web |
| Game | `game.milestone`: launch, phase and raid-unlock dates from `dates.json` | everyone | on, one per milestone | Open the date's guide |

Rules for every event:
- Quiet hours are set in the player's own time zone. A raid that starts inside quiet hours still sends `raid.starting`.
- A later notification replaces an earlier one of the same kind for the same raid (one thread per raid).
- At most one `bis.changed` digest a week.
- No marketing or Premium upsell pushes, ever. This is the "no nag" rule from the Premium entry spec.
- A notification never says more than the player could see on the site. Officer events go only to officers.

The time-based events (`raid.signup_nudge`, `raid.starting`, `raid.short`) come from one scheduler. It generalises today's `raid-posts` Cloud Run job, which runs every 10 minutes. The scheduler is idempotent per event, raid and player.

## Guard rails

- **API lanes merge after launch (coordinator ruling, 2026-10-30).** The A1, A2 and A3 branches change auth or add migrations, and the apps build against contract fixtures, so none of them merges to main before Nov 6. The app scaffolds, tokens, contracts and spec merge now.
- **Launch week freeze.** Lanes A1 and A2 add endpoints and a migration. No API deploy that carries their migration from Nov 3 to Nov 5 (launch Nov 4). Their code lands dark, behind unmounted routes, before or after that window. Only the deployed API migrates.
- **One behaviour, one place.** Rules stay in the API and the static data. Neither app re-derives BiS, gains or scores.
- **No shared UI layer.** Swift and Kotlin each render natively from the same spec and contracts.
- **Contracts are floors.** A fixture pins the shape, never catalogue counts.

## Decisions for the owner

| # | Question | Recommendation |
|---|---|---|
| P | In-app purchase of Premium | **Decided 2026-10-30 (owner): no in-app purchase; link to web checkout where the store allows it.** Today that means the US storefront on both stores. Elsewhere the app shows Premium status only, with no purchase prompt. The allowed regions are server config. Risk: outside the US, App Review may still ask for in-app purchase for unlocked features. |
| G | Premium gate on the apps | **Decided 2026-10-30 (owner): the apps are free and carry the web's exact entitlement rules.** Raid sign-up and every free web feature stay free in the app; guild members never pay; Premium features (server sims, coach, history charts) are Premium in the app too. Mobile-only Premium extras (a widget, a Live Activity for a server sim) are a post-v1 decision. |
| S | v1 scope | F1 to F8 as above, with F4 (raids) as the headline. |
| R | Repo | A monorepo: `apps/ios` and `apps/android` beside `api` and `web`, so contracts and tokens change in one commit. |
| B | Identifiers | Bundle id and package `gg.foreversixty.app`; app link `https://foreversixty.gg/app/…` (universal links and App Links) rather than a custom scheme. |
| M | Minimum OS | iOS 26 and Android 9 (API 28). |
