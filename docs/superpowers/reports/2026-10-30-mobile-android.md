# Mobile lane N: the Android scaffold (2026-10-30)

STATUS: done, green locally. Branch `wt/mobile-android`, not pushed.

## What landed

`apps/android`, a Gradle 9.8.2 wrapper build with a version catalog (`gradle/libs.versions.toml`) and convention plugins in `build-logic`:

- AGP 9.5.0 (built-in Kotlin), Kotlin 2.4.21, Compose BOM 2026.09.00, Material 3.
- compileSdk and targetSdk 37 (Android 17, the newest stable platform), minSdk 28, JDK 21 toolchain.
- Kotlin warnings are errors, and lint runs with warnings as errors.
- Application id `gg.foreversixty.app`; debug builds are `gg.foreversixty.app.debug`.

| Module | What it holds |
|---|---|
| `:app` | `ForeverApplication` and `AppGraph` (hand-wired dependencies), `MainActivity` (edge to edge, routes the sign-in App Link), and `ForeverTabScaffold`: a Material 3 NavigationBar with Character, BiS, Raids, Sims and Account. Character shows the placeholder home from `/v1/me`; the other tabs are placeholders. Also the network security config (no cleartext) and backup exclusions for the token files. |
| `:core:model` | `@Serializable` models: `Me`, `User`, `Character` (the character card), `CharacterBuild`, `BuildTalents`, `CharacterGuild`, `Attendance`, `Guild`, `GuildPlan`, `Entitlements`, `Billing`, `FinishedSim`, `ServerSims`, `Notifications`, `DiscordLink`, `AppExchangeRequest` and `AppExchangeResponse`. Each mirrors its Go struct in `api/internal/auth`. `ApiJson` is the one JSON configuration: unknown keys are ignored, a null takes the field's default, and absent fields are null. |
| `:core:network` | `ApiClient` (Ktor; get, post, put, delete) returns `ApiResult<T>` and unwraps the `{ok,data,error,request_id}` envelope. `ApiError` is one of `Server`, `Network` or `Decoding`. `ForeverHttpClient` sends `Authorization: Bearer <fsd_…>` from a `TokenProvider`, with a 15 s request and 10 s connect timeout. `ApiConfig.fromBuildConfig()` reads `API_BASE_URL` and `API_MODE`. Fixture mode uses `FixtureEngine`, `FixtureRoutes`, `AssetFixtureSource` and `ClasspathFixtureSource`. `AccountApi.me()` is the example endpoint. |
| `:core:auth` | `Pkce` (S256, checked against the RFC 7636 test vector). `SignInContract` holds the start URL, the redirect `https://foreversixty.gg/app/auth/callback`, and redirect parsing that checks state, scheme, host and path. `AppSignIn` runs `begin(browser)` then `complete(redirect)`, returning a `SignInOutcome`, plus `signOut()`. `CustomTabsLauncher` opens the browser. `TokenStore` (also the `TokenProvider`) has two implementations: `EncryptedTokenStore` and `AndroidTokenStore.create(context)`. The second is DataStore holding Tink AES-256-GCM ciphertext, with a keyset wrapped by an Android Keystore master key. |
| `:core:designsystem` | `Tokens.kt` placeholder at the path lane T owns, using values copied from `web/src/styles/tokens.css`. `ForeverTheme` is a dark-only Material 3 theme mapped onto the tokens. |

The contract fixtures come from the `forever.contract-fixtures` Gradle plugin. Its `copyContractFixtures` task copies `docs/contracts/fixtures/app/*.json` to `fixtures/app/` on every module's unit-test classpath and into the debug APK's assets. Release bundles carry no fixtures; I checked the bundle.

Debug builds default to fixture mode. `-Pforever.apiMode=live` and `-Pforever.apiBaseUrl=…` override it. Release builds are always live against `https://api.foreversixty.gg`.

## Commands (from `apps/android`)

```sh
export JAVA_HOME=/opt/homebrew/opt/openjdk@21/libexec/openjdk.jdk/Contents/Home
echo "sdk.dir=$HOME/Library/Android/sdk" > local.properties   # once; CI uses ANDROID_HOME
./gradlew test                          # every JVM test, debug and release variants
./gradlew :app:verifyRoborazziDebug     # screenshot test against the goldens
./gradlew :app:recordRoborazziDebug     # re-record goldens (app/src/test/screenshots)
./gradlew lint
./gradlew assembleDebug                 # app/build/outputs/apk/debug/app-debug.apk
./gradlew bundleRelease                 # unsigned unless FOREVER_UPLOAD_* env vars are set
```

## Verification

- `./gradlew --rerun-tasks test :app:verifyRoborazziDebug lint assembleDebug` passes. That is 46 JVM tests per variant:
  - model: 6
  - network: 16 (envelope, errors, bearer header, base URL, POST body, fixture engine)
  - auth: 22 (PKCE vector, redirect parsing including state forgery, lookalike paths and other hosts, single use, the sign-in flow against fixtures, encrypted store round trip, no token in clear on disk, a foreign key reads as signed out)
  - screenshots: 2
- Lint reports no issues in any module. It caught `URLEncoder.encode(String, Charset)` needing API 33, which is now fixed.
- `bundleRelease` builds with R8 and resource shrinking.
- I installed the debug APK on a temporary Pixel 9 / API 36 emulator (since deleted). It cold-started in 2.9 s with no crash, ran the Keystore and Tink initialisation, and rendered the home from the fixture.

## Choices

- **Screenshots: Roborazzi over Paparazzi.**
  - Roborazzi runs inside Robolectric as part of `./gradlew test`, on the same JVM test classpath and fixtures as everything else.
  - Paparazzi's layoutlib is pinned to particular AGP and compileSdk versions and historically trails new ones. That is a real risk on AGP 9.5 with SDK 37.
  - Robolectric runs at SDK 36 (`app/src/test/resources/robolectric.properties`). It is its newest android-all, and it needs `--add-exports java.base/jdk.internal.access` on JDK 21, which is set in `app/build.gradle.kts`.
  - The goldens are recorded on macOS and verified on Linux CI with a 0.5% change threshold.
- **Play upload: the Play Developer API through `r0adkll/upload-google-play`, not gradle-play-publisher.** GPP couples to AGP internals, and 4.1.1's support for AGP 9 is unproven. The action needs only the service-account JSON.
- **No DI framework yet.** `AppGraph` is about 30 lines. Hilt or Koin can come in once the graph outgrows it.

## CI (`.github/workflows/android.yml`)

- `build` runs on ubuntu-latest for pushes and PRs touching `apps/android/**`, `apps/shared/**`, `docs/contracts/**` or the workflow itself. It runs test, Roborazzi verify, lint and assembleDebug, and uploads the reports when it fails.
- `play-internal` runs only when manually dispatched with `release: true` **and** the repository variable `ANDROID_PLAY_RELEASE_ENABLED == 'true'`. Until then it is skipped, and it fails fast if a secret is missing. It builds a signed `bundleRelease` (versionCode = run number) and uploads a **draft** release to the internal track.
- The secret names are proposals for the owner (TODO in the workflow): `PLAY_SERVICE_ACCOUNT_JSON`, `ANDROID_UPLOAD_KEYSTORE_BASE64`, `ANDROID_UPLOAD_KEYSTORE_PASSWORD`, `ANDROID_UPLOAD_KEY_ALIAS` and `ANDROID_UPLOAD_KEY_PASSWORD`.
- CI has not run yet, because nothing is pushed.

## TODOs left in the code

- FCM (lane A2, F7, the owner's Firebase project): comment in `AndroidManifest.xml`.
- App Links `https://foreversixty.gg/app/*` with `autoVerify`: comment in `AndroidManifest.xml`. It needs `/.well-known/assetlinks.json` with the Play App Signing SHA-256. `MainActivity` already routes the redirect.
- `SignInContract`: confirm the query parameter names and the redirect path against lane A1 and C's contract. Only that file changes.
- `AppSignIn.signOut`: call `DELETE /v1/devices/self` (F1). `MainActivity`: surface the `SignInOutcome` (F1).
- `ForeverTheme`: fonts from lane T.

## Concerns

1. **Fixture names are guesses.** Lane C's fixtures had not landed, so I added `docs/contracts/fixtures/app/me.json` and `auth-app-exchange.json`, both in envelope form and built from the Go structs. When lane C merges, keep theirs. If their naming differs, change `FixtureRoutes.candidates` (one function). The rule today:
   - `GET /v1/x/y` reads `x-y.json`.
   - Any other method reads `x-y.<method>.json` first, then `x-y.json`.
   - A data-only fixture is wrapped in an envelope automatically.
2. **`MeFixtureTest` decodes every `me*.json` as `Me`.** If lane C adds `me-patch.json` or similar with another shape, narrow the filter.
3. **The PKCE verifier and state are held in memory only.** If the process dies while the Custom Tab is open, sign-in is rejected and has to be retried.
4. **`Tokens.kt` names are mine.** If lane T's generator emits different names, `ForeverTheme.kt` must change in the same commit.
5. **Debug APK size.** The debug APK is 42 MB (unminified Compose plus Tink). The release bundle is minified.
