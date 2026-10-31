# Mobile lane I: the iOS scaffold

Branch `wt/mobile-ios`. Plan: docs/superpowers/plans/2026-10-30-mobile-apps-lanes.md.

## What landed

`apps/ios` is a SwiftUI app named "Forever Sixty". Its bundle id is `gg.foreversixty.app`, it targets iOS 26 and it runs on iPhone only.

- **Project generation: XcodeGen 2.46.0** from `apps/ios/project.yml`. The `.xcodeproj` and `Generated/Info.plist` are gitignored. Why XcodeGen over Tuist:
  - It is one YAML file and one binary. CI downloads that binary at a pinned version.
  - It produces a plain Xcode project with no runtime, cache or manifest compile step.
  - The modules are a plain SwiftPM package, which XcodeGen links directly. Tuist's extra features (caching, its own module graph) buy nothing at this size, and they add a second dependency system on top of SwiftPM.
- **Modules** live in a local SwiftPM package, `apps/ios/Package.swift` (ForeverKit). Swift 6 language mode applies, with complete concurrency checking, `ExistentialAny`, and warnings treated as errors in the app.
  - `Models`: Codable `Me`, `User`, `CharacterCard`, `CharacterBuild`, `BuildTalents`, `CharacterGuild`, `Attendance`, `Guild`, `Entitlements`, `Billing` and `Device`. Shapes come from api/internal/auth (store.go, me_view.go, handler.go, me_billing.go). A Go nil slice that serialises as `null` decodes as `[]`, and unknown keys are ignored.
  - `Networking`: `APIClient`, `Endpoint<Response>`, `Envelope`, `ErrorBody`, `APIError`, `HTTPTransport` (`URLSessionTransport`, `FixtureTransport`), `FixtureCatalog`, `BearerTokenSource`, `APIConfiguration`/`APIMode`, and `WireFormat` (snake_case keys, RFC 3339 dates with or without fractional seconds).
  - `Auth`: `TokenStore` (`KeychainTokenStore`, `InMemoryTokenStore`), `PKCE`, `AppAuthContract`, `AppSignIn`, `WebAuthenticating` (`SystemWebAuthenticator` over ASWebAuthenticationSession), and `AuthError`.
  - `DesignTokens`: target path `apps/ios/Shared`. It holds only `Shared/DesignTokens.swift`, a **placeholder** that says lane T replaces it.
  - `DesignSystem`: `Theme` (Palette, Spacing, Radius), `Color(hex:)` and `.cardSurface()`. This is the only reader of DesignTokens.
- **App** (`apps/ios/App`): a TabView with Character, Raids, BiS, Sims and Account (lane D sets the final set). The home tab loads `/v1/me` and shows the main character's card; the other tabs are placeholders. `AppEnvironment` builds the client and token store from Info.plist.
- **Build settings:** `API_BASE_URL` (default `https://api.foreversixty.gg`) and `API_MODE` (Debug `fixture`, Release `live`) reach Info.plist as `FSAPIBaseURL` and `FSAPIMode`. Override on the command line, for example `xcodebuild … API_MODE=live`. The configuration fails fast when the URL is missing or is plain http outside localhost.
- **Fixtures:** project.yml copies `docs/contracts/fixtures/app` into the app bundle as `Fixtures` through a folder-reference resource, which is a copy-resources phase. Tests read the repo folder in place.
  - **STUB:** lane C had not committed yet, so `me.json` and `auth-app-exchange.json` there are stubs written by this lane. Each is a full response body, envelope included. Lane C should overwrite them.
  - If lane C's fixtures are bare `data` and not envelopes, change `FixtureTransport` in one place.
- **Signing:** ad-hoc (`CODE_SIGN_IDENTITY "-"`, no team, no account). A fully unsigned build has no entitlements, and the simulator Keychain then refuses it with -34018.
- **TODOs left in place:**
  - Associated Domains `applinks:foreversixty.gg`, in project.yml and WebAuthenticator.swift. It waits on the Team ID, and the https sign-in callback needs it to complete.
  - APNs entitlement and push-token registration (lane A2, F7), in project.yml and AppEnvironment.swift.
  - Server-side device revoke on sign out (lane A1), in AppSignIn.swift.

## The sign-in contract this lane assumed

Lane A1 should confirm it; every name sits in `AppAuthContract`.

1. Open `GET /v1/auth/app/start?platform=ios&code_challenge=<S256>&code_challenge_method=S256&state=<random>` in ASWebAuthenticationSession.
2. The API redirects to `https://foreversixty.gg/app/auth/callback?code=…&state=…`, or `?error=…&state=…` on failure.
3. Call `POST /v1/auth/app/exchange` with the body `{code, code_verifier, platform: "ios", device_name}`, unauthenticated. It returns `data: {token: "fsd_…", device: {id, name, platform, created_at, last_seen_at?}}`.

## Build and test locally

The project needs Xcode 27.1, the iOS 27 simulator runtime, an iPhone 17 simulator and `brew install xcodegen`.

```
cd apps/ios
xcodegen generate
xcodebuild test -project ForeverSixty.xcodeproj -scheme ForeverSixty \
  -destination 'platform=iOS Simulator,name=iPhone 17' -derivedDataPath build/DerivedData
```

`make -C apps/ios test` (or `make -C apps/ios build`) runs the same steps.

To re-record the snapshot, run with `SNAPSHOT_TESTING_RECORD=all` in the scheme's test environment, then review the PNG at `apps/ios/AppTests/__Snapshots__/HomeSnapshotTests/`.

## Tests

44 tests, all green on the iPhone 17 simulator (`make -C apps/ios test`, 2026-10-30), with no warnings in project code.

All tests use Swift Testing.
- **ModelsTests:** `/v1/me` fixture decoding, main-character choice, null slices, and sparse and full character cards.
- **NetworkingTests:** the client's URL, method, headers and bearer token; ok, error, 401, non-envelope 5xx, decoding failure and 204; the JSON body in snake_case; fixture mode and a missing fixture; every catalogued fixture exists; configuration parsing; envelope decoding and Go timestamps.
- **AuthTests:** PKCE against the RFC 7636 Appendix B vector, verifier shape and injected randomness; the start URL; a full sign-in that stores the token; the exchange body, whose verifier matches the challenge; state mismatch; an error callback; cancel; a refused exchange; the wrong callback host; sign out.
- **ForeverSixtyTests** (hosted in the app):
  - A snapshot of the placeholder home with swift-snapshot-testing 1.19.7, pinned exactly. The reference was recorded on the iPhone 17 simulator, iOS 27.0, rendered at the iPhone 13 portrait size, with a 0.98 perceptual tolerance.
  - The Keychain round trip. It must run hosted, because the Keychain refuses an unhosted package bundle.

## CI

`.github/workflows/ios.yml` runs on `macos-latest` for changes under apps/ios, apps/shared and docs/contracts, and for changes to the workflow itself.
1. Select **Xcode 27.1** with maxim-lobanov/setup-xcode. The job fails loudly if the image lacks it; if it does, move `runs-on` to the image that carries Xcode 27.1.
2. Download XcodeGen 2.46.0.
3. Generate the project.
4. Run `xcodebuild test` on the iPhone 17 simulator.
5. On failure, upload the xcresult and the snapshots.

## Interfaces for the feature lanes

```swift
// Networking
struct APIClient { let baseURL: URL
  init(baseURL:transport:tokens:); static func make(_: APIConfiguration, tokens:, bundle:) -> APIClient
  func send<R>(_ endpoint: Endpoint<R>) async throws(APIError) -> R }
struct Endpoint<Response: Decodable & Sendable> { init(method:path:query:body:authenticated:)
  static func json(_ method:, _ path:, body: some Encodable, authenticated:) throws(APIError) -> Endpoint }
extension Endpoint where Response == Me { static var me: Endpoint }   // add routes the same way
enum APIError { transport, unauthorized(requestID:), server(status:body:requestID:), decoding, unexpectedStatus, missingFixture }
FixtureCatalog.app: [FixtureRoute]   // add one FixtureRoute per new route + lane C fixture file
protocol HTTPTransport; protocol BearerTokenSource
// Auth
protocol TokenStore: BearerTokenSource { load/save/clear throws(TokenStoreError) }
struct AppSignIn { init(client:web:store:random:); func signIn(deviceName:) async throws(AuthError) -> AppSession; func signOut() throws(AuthError) }
@MainActor final class SystemWebAuthenticator: WebAuthenticating { init(anchor: @MainActor () -> ASPresentationAnchor) }
// DesignSystem
Theme.Palette.{background, raised, line, text, strong, muted, accent}; Theme.Spacing.{xs,sm,md,lg}; Theme.Radius.card; View.cardSurface()
```

## Concerns

- Package.resolved is not committed. The snapshot library is pinned exactly, but its transitive dependencies (swift-syntax, custom-dump, issue-reporting) resolve fresh on each clean run.
- The snapshot depends on the simulator runtime. If CI's iOS runtime differs from the local 27.0, the 0.98 tolerance may not be enough. Re-record on CI's runtime if it fails, and keep one runtime as the reference.
- The first build compiles swift-syntax for the snapshot library's macros, which makes a cold local build take several minutes.
- The stub fixtures and the guessed exchange shape must be reconciled with lanes C and A1 before Wave 1.
