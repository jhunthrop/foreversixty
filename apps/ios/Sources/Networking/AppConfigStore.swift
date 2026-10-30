import Foundation
import Models
import Synchronization

/// Where the last good `GET /v1/app/config` answer is kept between launches. Not a secret, so
/// `UserDefaultsAppConfigCache` is enough; tests use `InMemoryAppConfigCache`.
public protocol AppConfigCache: Sendable {
    func load() -> Data?
    func save(_ data: Data)
}

/// The app config, read once at launch (`refresh()`), cached, and answered from the cache or
/// the bundled defaults when offline. A feature is on when the API says so, else when its
/// bundled default (apps/shared/features.json) is; a key the API sends that the build does not
/// know is ignored.
public final class AppConfigStore: Sendable {
    private let client: APIClient
    private let cache: any AppConfigCache
    private let state: Mutex<AppConfig?>

    public init(client: APIClient, cache: any AppConfigCache) {
        self.client = client
        self.cache = cache
        let cached = cache.load().flatMap { try? WireFormat.decoder().decode(AppConfig.self, from: $0) }
        state = Mutex(cached)
    }

    /// The last config the API gave, from this launch or a cached earlier one; nil before the first.
    public var current: AppConfig? { state.withLock { $0 } }

    /// Fetches the config and caches it. On failure the cached or bundled answer stands, and the
    /// error is returned for the caller to log, never shown: the app works offline.
    @discardableResult
    public func refresh() async -> APIError? {
        do throws(APIError) {
            let config = try await client.send(.getAppConfig)
            state.withLock { $0 = config }
            if let data = try? WireFormat.encoder().encode(config) { cache.save(data) }
            return nil
        } catch {
            return error
        }
    }

    public func isEnabled(_ feature: Feature) -> Bool {
        current?.features[feature.rawValue] ?? feature.bundledDefault
    }

    /// True when the API no longer supports `version` (dotted, compared numerically): block with
    /// the update screen. Unknown until the first config, so never blocks offline on first launch.
    public func requiresUpdate(version: String) -> Bool {
        guard let minimum = current?.minAppVersion.ios else { return false }
        return AppVersion.isBelow(version, minimum)
    }

    /// Whether the storefront country may show See plans (`web_checkout_regions`).
    public func mayLinkToWebCheckout(storefront country: String) -> Bool {
        current?.webCheckoutRegions.contains(country.uppercased()) ?? false
    }
}

/// Dotted version strings compared part by part as numbers: "1.10.0" is above "1.9.2".
public enum AppVersion {
    public static func isBelow(_ version: String, _ minimum: String) -> Bool {
        let a = parts(version)
        let b = parts(minimum)
        for i in 0..<max(a.count, b.count) {
            let x = i < a.count ? a[i] : 0
            let y = i < b.count ? b[i] : 0
            if x != y { return x < y }
        }
        return false
    }

    private static func parts(_ version: String) -> [Int] {
        version.split(separator: ".").map { Int($0.prefix { $0.isNumber }) ?? 0 }
    }
}

public struct UserDefaultsAppConfigCache: AppConfigCache {
    private static let key = "gg.foreversixty.app-config"
    // UserDefaults is documented thread-safe.
    private nonisolated(unsafe) let defaults: UserDefaults

    public init(defaults: UserDefaults = .standard) {
        self.defaults = defaults
    }

    public func load() -> Data? { defaults.data(forKey: Self.key) }

    public func save(_ data: Data) { defaults.set(data, forKey: Self.key) }
}

public final class InMemoryAppConfigCache: AppConfigCache {
    private let stored: Mutex<Data?>

    public init(_ data: Data? = nil) {
        stored = Mutex(data)
    }

    public func load() -> Data? { stored.withLock { $0 } }

    public func save(_ data: Data) { stored.withLock { $0 = data } }
}
