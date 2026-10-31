import Foundation
import Networking
import Observation
import StoreKit

/// Decision P: no in-app purchase. `See plans` links to web checkout only where the store
/// allows it, which the API names (`GET /v1/app/config`'s `web_checkout_regions`). Every unknown
/// fails closed: no config, no storefront, or the feature off means no link.
enum WebCheckoutPolicy {
    static func mayLink(regions: [String]?, storefront: String?, featureOn: Bool) -> Bool {
        guard featureOn, let regions, let country = storefront.flatMap(alpha2) else { return false }
        return regions.contains(country)
    }

    /// StoreKit names the storefront in ISO 3166-1 alpha-3 (`USA`); the config lists alpha-2
    /// (`US`). Returns nil for anything that is not a known region.
    static func alpha2(_ code: String) -> String? {
        let upper = code.uppercased()
        guard upper.count == 2 || upper.count == 3, upper.allSatisfy(\.isLetter) else { return nil }
        guard let region = Locale(identifier: "und_\(upper)").region, region.isISORegion else { return nil }
        let identifier = region.identifier
        return identifier.count == 2 && identifier.allSatisfy(\.isLetter) ? identifier : nil
    }
}

/// Where the storefront country comes from: StoreKit in the app, a fixed answer in tests and
/// captures.
protocol StorefrontReading: Sendable {
    func countryCode() async -> String?
}

struct StoreKitStorefront: StorefrontReading {
    func countryCode() async -> String? { await Storefront.current?.countryCode }
}

struct FixedStorefront: StorefrontReading {
    let code: String?
    func countryCode() async -> String? { code }
}

/// Whether this install may show `See plans` and `Manage billing`, resolved once the config and
/// the storefront have answered. Starts at no link and stays there on any failure.
@MainActor
@Observable
final class PlansLink {
    private(set) var mayLink = false
    private let appConfig: AppConfigStore
    private let storefront: any StorefrontReading
    private let featureOn: @MainActor () -> Bool

    init(appConfig: AppConfigStore, storefront: any StorefrontReading, featureOn: @escaping @MainActor () -> Bool) {
        self.appConfig = appConfig
        self.storefront = storefront
        self.featureOn = featureOn
    }

    func resolve() async {
        if appConfig.current == nil { await appConfig.refresh() }
        let country = await storefront.countryCode()
        mayLink = WebCheckoutPolicy.mayLink(regions: appConfig.current?.webCheckoutRegions, storefront: country, featureOn: featureOn())
    }
}
