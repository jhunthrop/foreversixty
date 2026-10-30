# Mobile Wave 0: integration (2026-10-30)

Branch `wt/mobile-integrate`, from origin/main (43ae25c4c). Not pushed.

## Merges

| Commit | Branch | Resolution |
|---|---|---|
| fde8e27c6 | wt/mobile-tokens (lane T) | clean |
| 5c86addc3 | wt/mobile-contracts (lane C) | clean |
| fc21d911d | wt/mobile-android (lane N) | `Tokens.kt` taken from lane T, `me.json` from lane C, the stub `auth-app-exchange.json` dropped |
| 7c57982a1 | wt/mobile-ios (lane I) | `DesignTokens.swift` taken from lane T, `me.json` from lane C, the stub `auth-app-exchange.json` dropped |

`wt/mobile-auth` and `wt/mobile-push` are not merged. Their contracts are read from their reports and folded into the docs and the apps.

## Fix commits

- **44349ebfb: tokens, fonts and icon.**
  - **iOS.**
    - `gen-mobile-tokens.mjs` now emits `public` Swift tokens. `DesignTokens` is its own SwiftPM module, so the internal enums could not be read.
    - `Theme.swift` maps its roles onto `FSColor`, `FSFont`, `FSTextSize`, `FSSpacing` and `FSRadius`, plus a new `Theme.Typography`.
    - `project.yml` bundles `apps/shared/fonts` and lists the six faces under `UIAppFonts`.
  - **Android.**
    - Two convention plugins in `SharedAssetsConventionPlugins.kt` copy assets at build time, so nothing is checked in twice:
      - `forever.shared-fonts` copies the fonts to `res/font` with lowercase names.
      - `forever.app-icon` copies the adaptive icon.
    - `ForeverTheme` builds the Barlow, Barlow Condensed and Cinzel families from `FsColor` and `FsFontWeight`.
  - **App icon.**
    - `actool` ignores a symlinked appiconset. `build-app-icon.mjs` now writes a real catalog at `apps/shared/brand/ios/Brand.xcassets`, and iOS compiles it in place. Rerunning the generator reproduces the PNG byte for byte.
    - XcodeGen's `createIntermediateGroups` is off: XcodeGen 2.46 mis-pathed sources outside `apps/ios`.
  - **gradlew.** Lane N committed `gradlew` as 100644. CI runs `./gradlew`, so it would have failed. The file is executable again.
  - Both web scripts are formatted with Prettier. They were unformatted on lane T.
- **77a150641: fixtures and sign-in.**
  - **Fixture routing.**
    - Each app has one route table from method and path to lane C's names: `FixtureCatalog` on iOS, `FixtureRoutes` on Android.
    - Routes take `{name}` segments, including `{spec}.json`. They cover the `cdn.*` files and the 204 routes.
    - Variants select the `.empty`-style success files, the `.error-<code>` files with their status, and the shared `error-unauthorized`.
    - Fixtures are served exactly as lane C wrote them: envelopes for the API and bare JSON for the CDN. Android no longer wraps a bare fixture.
    - A new `StaticDataClient` on both platforms reads `/data/*` without the token.
  - **Model tests.** Both apps decode lane C's `me`, `me.premium`, `me.empty`, `me.patch`, `character.card`, `character.card.empty` and `app.signin.exchange` fixtures. A new `CharacterProfile` model reads `GET /v1/characters/{key}`.
  - **Sign-in (A1).**
    - The callback is exactly `https://foreversixty.gg/app/auth`, and no `redirect_uri` is sent.
    - The exchange body is `{code, code_verifier, device_name}`. The 201 answer `{token, device{id,name,platform}, me}` decodes.
    - Sign-out posts `/v1/auth/app/logout`, then clears the token whatever the answer.
    - A 401 (`unauthorized` or `session_expired`) to a request that carried the token discards the token in the client. `browser_session_required` is flagged on both platforms.
    - Android's client now accepts a 204 with no body.
- **6a35be174: contract.** `app-v1.md` now carries:
  - A1's final sign-in: the confirm page and continue, the `/app/auth` landing, the exchange and logout, and the codes `session_expired` and `browser_session_required`.
  - A2's routes: `GET` and `PATCH /v1/me/notifications` (now with `time_sensitive`), and `POST` and `DELETE /v1/me/push-tokens`.
  - The push payload shapes: thread `raid-<id>` and collapse `<kind>-<raid id>`.
  - The `RAID_SIGNUP` category, with In, Late and Out and the `{character}` substitution.
  - The officer lock, unlock, cancel and bench moves.

  New fixtures are `push.raid-posted.apns`, `push.raid-posted.fcm` and `push.raid-starting.apns`. The envelope test skips `push.*`, and a new contract test holds those fixtures to the payload fields.

## Screenshots, re-recorded

- **Android, both goldens.** They now read lane C's `me.json` (Obnoxious Yell, not the stub's Thrallson) and render in Barlow and Barlow Condensed.
- **iOS home.** It now renders in Barlow Condensed and Barlow.
  - The old reference still passed, because the text changes less than 1% of the pixels.
  - It was re-recorded anyway so the reference shows what the app draws.

## Verification

| Check | Result |
|---|---|
| `make -C apps/ios test` (iPhone 17, iOS 27.0) | 58 tests passed: Models 10, Networking 27, Auth 19, app-hosted 2 |
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21) | BUILD SUCCESSFUL: 62 JVM tests, lint clean, debug APK built |
| `vitest run src/styles/mobile-tokens.test.ts` (node_modules symlinked, then removed) | 5 passed |
| `go test -race -p 1` contract, auth, builds, guilds, rankings, reports and sims (private DB `forever_gc_integrate`, dropped) | all ok |
| `go vet ./...` in api | clean |

The simulator's service hub failed twice: the test runner hung at `_dyld_start`. Stopping my own run and restarting it fixed it. No code was involved.

## Concerns

1. **The contract is not machine-checked for A1 and A2.**
   - The `app.*` and `push.*` fixtures are held only to the envelope rule and the payload fields.
   - When `wt/mobile-auth` and `wt/mobile-push` merge, add `app_contract_test.go` beside their handlers. The push builders (`apnsBody`, `fcmBody`) should then be compared with the `push.*` fixtures.
2. **Push fixtures are hand-written** from A2's code (`push/apns.go`, `push/fcm.go`, `guilds/raids_events.go`). The copy and the dates in them are examples.
3. **`GET /v1/app/config` still has no handler.** The publication of `active-build.json` and `dates.json` is still open (lane C items 1 and 2). `StaticDataClient.activeBuild()` reads a file the CDN does not serve yet.
4. **Two route tables, one per app.** iOS and Android keep their own route tables. Each has a test that every route resolves to an existing fixture, but nothing checks the two against each other.
5. **Android icon layers.** The 1024 px layers go into `mipmap-nodpi`, and the launcher scales them down. The debug APK is still 42 MB.
6. **`createIntermediateGroups: false`** flattens the iOS project navigator. It has no effect on the build.
7. **A1 owner TODOs still stand.** The Team ID and the SHA-256 for universal links are needed before `APP_SIGN_IN` is turned on, and the hand-off needs testing on real devices.
