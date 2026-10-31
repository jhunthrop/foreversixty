import Foundation
import Models
@testable import Networking
import TestSupport
import Testing

@Suite("App config")
struct AppConfigStoreTests {
    private let base = URL(string: "https://api.example.test")!

    private func store(_ transport: any HTTPTransport, cache: InMemoryAppConfigCache = InMemoryAppConfigCache()) -> AppConfigStore {
        AppConfigStore(client: APIClient(baseURL: base, transport: transport), cache: cache)
    }

    private var offline: RecordingTransport { RecordingTransport(status: 503, body: "") }

    @Test("before any answer every feature takes its bundled default")
    func bundledDefaults() {
        let config = store(offline)
        for feature in Feature.allCases {
            #expect(config.isEnabled(feature) == feature.bundledDefault)
        }
        #expect(config.requiresUpdate(version: "0.0.1") == false)
        #expect(config.mayLinkToWebCheckout(storefront: "US") == false)
    }

    @Test("the API's answer switches features, and a key it leaves out keeps the default")
    func apiOverrides() async {
        let config = store(FixtureTransport(directory: FixtureLocation.directory))
        #expect(await config.refresh() == nil)
        #expect(config.isEnabled(.raids))
        #expect(config.isEnabled(.signIn) == false)
        #expect(config.isEnabled(.sims) == Feature.sims.bundledDefault)
        #expect(config.mayLinkToWebCheckout(storefront: "us"))
        #expect(config.mayLinkToWebCheckout(storefront: "DE") == false)
    }

    @Test("the answer is cached, so the next launch reads it offline")
    func cachedOffline() async {
        let cache = InMemoryAppConfigCache()
        _ = await store(FixtureTransport(directory: FixtureLocation.directory), cache: cache).refresh()
        let nextLaunch = store(offline, cache: cache)
        #expect(await nextLaunch.refresh() != nil)
        #expect(nextLaunch.isEnabled(.raids))
        #expect(nextLaunch.current?.activeBuild.isEmpty == false)
    }

    @Test("an unreadable cache is ignored")
    func badCache() {
        let config = store(offline, cache: InMemoryAppConfigCache(Data("nope".utf8)))
        #expect(config.current == nil)
        #expect(config.isEnabled(.raids) == Feature.raids.bundledDefault)
    }

    @Test("the config is fetched without the device token")
    func noToken() async {
        let transport = RecordingTransport(body: #"{"ok":true,"data":{"web_checkout_regions":[],"min_app_version":{"ios":"2.0.0","android":"1.0.0"},"active_build":"b","features":{}},"error":null,"request_id":"r"}"#)
        let client = APIClient(baseURL: base, transport: transport, tokens: FixedToken(token: "fsd_x"))
        let config = AppConfigStore(client: client, cache: InMemoryAppConfigCache())
        _ = await config.refresh()
        #expect(transport.requests.first?.value(forHTTPHeaderField: "Authorization") == nil)
        #expect(config.requiresUpdate(version: "1.9.9"))
        #expect(config.requiresUpdate(version: "2.0.0") == false)
    }

    @Test("versions compare numerically, part by part", arguments: [
        ("1.9.2", "1.10.0", true), ("1.10.0", "1.9.2", false), ("1.0", "1.0.0", false),
        ("1.0.0", "1.0.1", true), ("2", "1.99", false),
    ])
    func versions(version: String, minimum: String, below: Bool) {
        #expect(AppVersion.isBelow(version, minimum) == below)
    }
}

@Suite("Features in development builds")
struct DevBuildFeatureTests {
    private let store = AppConfigStore(
        client: APIClient(baseURL: URL(string: "https://api.example.test")!, transport: RecordingTransport(body: "{}")),
        cache: InMemoryAppConfigCache()
    )

    @Test("debug and fixture builds turn every feature on; a live release follows the config")
    func availability() {
        #expect(store.isAvailable(.signIn, mode: .fixture, isDebugBuild: false))
        #expect(store.isAvailable(.signIn, mode: .live, isDebugBuild: true))
        #expect(!store.isAvailable(.signIn, mode: .live, isDebugBuild: false))
    }
}
