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

### Wave 0.5: one source of truth, operable apps (owner 2026-10-30: "do what's necessary to have scalable and operatable software")

Two lanes start once the Wave 0 integration lands.

| Lane | Delivers |
|---|---|
| **P1. Generated, not hand-written** | (1) `api/openapi.yaml` covers every app endpoint. A test checks every contract fixture against its OpenAPI schema, so the spec, fixtures and handlers cannot drift. The Swift and Kotlin models and endpoint stubs are generated from it and checked in; CI fails on drift, the same way it does for tokens. (2) One copy catalog (`apps/shared/copy/*.json`) for every on-screen string, generated into a Swift strings catalog and Android string resources, with plural and placeholder rules. (3) One deep-link table (`apps/shared/links.json`) generates `apple-app-site-association`, `assetlinks.json` and both apps' routers. (4) Feature switches in `GET /v1/app/config`, read once at launch and cached, so a feature ships dark in both apps and turns on together. |
| **P2. Operable** | (1) Each app sends `X-App: ios|android/<version> (<build>)` on every request. The API logs it and counts requests and errors per platform and version. The minimum-version gate uses the same header and returns `426` with an update link. (2) Crash and hang reporting on both platforms, privacy-first: no user ids, no ad SDKs, disclosed in the store privacy forms. Crashlytics or Sentry; pick one for both apps. (3) One version file (`apps/VERSION`) and one changelog. One command bumps both apps, builds, signs and uploads to the Play internal track, and to TestFlight once Apple enrollment clears (fastlane or the platform CLIs). (4) A mobile runbook: release, staged rollout, halt, forced update, rotating push keys, and what to do when push or sign-in fails. |

Gate to Wave 1: lane D at SHIP from both personas, lane C's contracts merged, and I and N building green in CI.

### Wave 1: features (one iOS lane and one Android lane per feature, at the same time)

Each feature is independent. Up to four features run at once per platform, so eight lanes are live together. Each pair works from the same spec section and the same contract fixtures, and each pair's captures are reviewed side by side.

| Feature | Needs | Scope (subject to lane D's spec) |
|---|---|---|
| **F1. Sign-in, account, character switcher** | A1 | Battle.net sign-in; the header's character selector as the app's switcher; the plan row; sign out; the device list. **In-app account deletion** (App Store guideline 5.1.1(v): an app that creates accounts must let people delete them in the app). Needs an API endpoint for deletion with a confirmation step and the same data removal as today's emailed request. Today deletion is only by emailed request. |
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

## Wave 3: everything the web does that v1 lacks, and what only a phone does well

Owner, 2026-10-31: "make the mobile app as feature rich as possible", then "build the server features now - stop deferring things". The experience spec is `design/specs/2026-10-31-mobile-wave3.md` (boards in `design/mocks/2026-10-31-mobile-wave3/renders/`, verdicts in `design/reviews/2026-10-31-mobile-wave3/`). It extends the v1 spec and changes the app to five tabs (Home, Raids, BiS, Guides, You) with a search button on every root bar.

Each lane is one feature on both platforms plus its server and static-data changes (the W3-C contract changes of spec section 10: OpenAPI, fixtures, contract test, generated clients, a `features` switch in `GET /v1/app/config`). The freeze no longer splits a lane: the endpoints, migrations and apps ship together, the migration in the same deploy as its handler.

| Lane | Feature | Depends on | Contract changes |
|---|---|---|---|
| **L-F14** | Home-screen and lock-screen widgets, controls, raid Live Activity, server-sim Live Activity and ongoing notifications | v1 F2, F4; A2 senders | W3-C12 `GET /v1/me/widget`, W3-C13 live-activity registration (migration), W3-C14 events |
| **L-F16** | Roster with attendance, raid history, calendar sync and a calendar link | v1 F4 | W3-C16 `role` on the guild home, W3-C17 calendar feed (migration) |
| **L-F12** | Universal search and the item page; the static data store | v1 F2.6, F3 | W3-C8 static `loot`, `search`, `item-uses` |
| **L-F10** | Your reports: list, report, your parses, meters, coach when entitled | v1 F4.22; A2 | W3-C2 `reports/{id}/me`, W3-C3 meters, W3-C4 coach routes documented |
| L-F9 | Guides, tier lists, what changed (the fifth tab) | L-F12's store | W3-C1 static guides, tiers, changelog |
| L-F11 | Rankings, character pages, watching, the Paperdoll | L-F12 | W3-C5 rankings documented, W3-C6 encounters, W3-C7 watching (migration) |
| L-F13 | Talent planner: build, share, open a shared link | v1 F6; L-F9 | W3-C9 `DELETE /v1/builds/{id}`, W3-C10 races and combos documented, W3-C11 link rows |
| L-F15 | App Intents and Siri, Spotlight, Android shortcuts, share, Offline data, stale-data strip | L-F12, L-F14, L-F16, L-F11 | W3-C15 feature switches |

**Order, at most four features at once:** L-F14, L-F16, L-F12 and L-F10 start together (most valuable first: the home-screen answer, the guild around the raid, search with the item page, the after-raid loop). L-F9 and L-F11 start when L-F12's store merges, then L-F13, then L-F15 last.

**Premium:** no mobile-only Premium extra. The coach card (`log_coach`), the sim Live Activity (`server_sims`) and a Parse history row (`history`) follow the web's flags; everything else is free (spec section 1.1).

This overrides the guard rail "API lanes merge after launch" for Wave 3 lanes (owner, 2026-10-31).

Each lane keeps the house rules: tests first for logic, captures on iPhone 17 and Pixel 9 at the default and largest text sizes, ux-designer and wow-player SHIP before merge, direct push to main behind its switch.

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
| L | Where See plans links | **Decided 2026-10-31 (coordinator, owner delegated): iOS links in the US storefront only; Android links nowhere until the Play Console is enrolled in Google's US external content links programme.** `GET /v1/app/config` answers `web_checkout_regions` by the `X-App` platform: `ios` gets `["US"]`, `android` gets `[]`. Enrolling is an owner click in Play Console; after it, the Android list becomes `["US"]` in server config with no app release. The Play Billing library stays: it reads the storefront country the rule needs, and the programme works through it. The Data safety form says no in-app products. |
| W | Premium outside link regions (spec §17 Q3) | **Decided 2026-10-31: accept the review risk.** Paying players keep every Premium feature in every country; outside link regions the app shows status only and no purchase word. If App Review objects, answer that the app sells nothing and is free; change the behaviour only on a rejection. |
| U | Supporter mark for seat holders | **Decided 2026-10-31: the apps draw `supporter_mark` exactly as the API resolves it.** The API's rule already counts a held guild seat (`entitlements.standardEntitlementSQL`), the same as the web, so seat holders show the mark in both. |
| Q | The other spec §17 questions | **Decided 2026-10-31.** Q1 approving characters: v1.1, on the officer tools already in v1. Q4 signed-in checkout handoff: after launch, as its own lane with a security review, because See plans is US-iOS-only at first. Q5 Discord DMs: not in v1; the Discord option is the existing channel events. Q6 tablets: the phone layout at launch. |
